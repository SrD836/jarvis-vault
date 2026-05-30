#!/usr/bin/env python3
"""Honest expectancy report for the polymarket bot (v7 cohort, simulation).

Defeats the survivorship bias of "P&L desde v7" (winners close fast, longshot
losers sit open because v7 has no stop) by NETTING realized P&L of the v7 cohort
with the mark-to-market unrealized P&L of the still-open book (mark_to_market.json,
produced by the Go cmd that reuses polyclient.ExitPrice). Adds Brier on resolved
predictions and rule-firing counts, then emits a PASS/FAIL verdict against David's
explicit criterion: expectancy/trade > 0 over N>=50, Brier <= 0.25, new rules firing.

Honest by construction: it also reports a WORST-CASE where open positions we cannot
price (dead longshots with no bid) resolve to total loss — so the verdict never
hides behind "unpriceable".
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import sys
from collections import Counter, defaultdict

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import brier  # noqa: E402  reuse load_rows, brier, group_brier, parse_ts, _pnl

V7_START = os.environ.get("V7_START", "2026-05-19")
N_MIN = int(os.environ.get("EXPECTANCY_N_MIN", "50"))
BRIER_MAX = float(os.environ.get("EXPECTANCY_BRIER_MAX", "0.25"))
T = "vault/inbox/trading/"


def _entry_dt(r: dict):
    return brier.parse_ts(r.get("entry_time") or r.get("entry_timestamp") or "")


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--closed", default=T + "closed.jsonl")
    ap.add_argument("--predictions", default=T + "predictions.jsonl")
    ap.add_argument("--mtm", default=T + "mark_to_market.json")
    ap.add_argument("--blocked", default=T + "blocked.jsonl")
    ap.add_argument("--out", default="vault/outputs/expectancy-report-" + dt.date.today().isoformat() + ".md")
    ap.add_argument("--no-write", action="store_true")
    args = ap.parse_args()

    floor = dt.datetime.min.replace(tzinfo=dt.timezone.utc)
    v7_start = brier.parse_ts(V7_START + "T00:00:00Z") or floor

    # --- realized v7 cohort (by ENTRY date; incluye perdedores, no maquilla) ---
    closed = brier.load_rows(pathlib.Path(args.closed))
    cohort = [r for r in closed if (_entry_dt(r) or floor) >= v7_start]
    realized_pnl = sum(brier._pnl(r) for r in cohort)
    n_real = len(cohort)
    wins = sum(1 for r in cohort if brier._pnl(r) > 0)
    fb = [r for r in cohort if r.get("exit_price_source") == "lastTradePrice_fallback"]
    exp_real = realized_pnl / n_real if n_real else 0.0
    wr_real = 100 * wins / n_real if n_real else 0.0

    # --- mark-to-market of still-open book ---
    mtm = {}
    pm = pathlib.Path(args.mtm)
    if pm.exists():
        try:
            mtm = json.loads(pm.read_text(encoding="utf-8"))
        except json.JSONDecodeError:
            pass
    unreal = float(mtm.get("total_unrealized", 0.0))
    n_priced = int(mtm.get("n_priced", 0))
    n_unpriceable = int(mtm.get("n_unpriceable", 0))
    cost_unpriceable = float(mtm.get("cost_unpriceable", 0.0))

    # --- combined honest expectancy ---
    n_comb = n_real + n_priced
    comb_pnl = realized_pnl + unreal
    exp_comb = comb_pnl / n_comb if n_comb else 0.0
    # worst case: unpriceable open (dead longshots) resolve to 0 -> lose full stake.
    worst_pnl = comb_pnl - cost_unpriceable
    n_worst = n_comb + n_unpriceable
    exp_worst = worst_pnl / n_worst if n_worst else 0.0

    # --- Brier on resolved predictions (clean, no survivorship) ---
    preds = brier.load_rows(pathlib.Path(args.predictions))
    n_brier, b = brier.brier(preds)

    # --- rule firing ---
    blk = brier.load_rows(pathlib.Path(args.blocked))
    rule_counts = Counter(str(r.get("rule") or r.get("vetoed_by") or r.get("reason", "?")) for r in blk)
    new_rules = {k: v for k, v in rule_counts.items()
                 if k.startswith(("R1", "R3", "R5", "S1", "pnl_expectancy"))}
    rules_ok = any(k.startswith(("R1", "R3", "R5")) for k in rule_counts)

    # --- per-category P&L in cohort (qué sangra) ---
    cat: dict[str, dict] = defaultdict(lambda: {"n": 0, "pnl": 0.0})
    for r in cohort:
        a = cat[str(r.get("category") or "uncategorized")]
        a["n"] += 1
        a["pnl"] += brier._pnl(r)
    bleeders = sorted(((c, a["n"], a["pnl"]) for c, a in cat.items() if a["pnl"] < 0),
                      key=lambda x: x[2])

    # --- verdict ---
    crit_exp = exp_comb > 0 and n_comb >= N_MIN
    crit_brier = n_brier > 0 and b <= BRIER_MAX
    verdict = "PASS" if (crit_exp and crit_brier and rules_ok) else "FAIL"

    # --- render ---
    L = []
    L.append(f"# Expectancy report — v7 cohort (entry >= {V7_START}) — {dt.date.today().isoformat()}")
    L.append("")
    L.append("> Honesto: netea realizado v7 con mark-to-market del libro abierto (sin sesgo de supervivencia).")
    L.append("")
    L.append("| métrica | valor |")
    L.append("|---|---|")
    L.append(f"| realizado v7 (cerrados) | {realized_pnl:+.2f} USD · n={n_real} · wr={wr_real:.1f}% · exp/trade={exp_real:+.2f} |")
    L.append(f"| mark-to-market abiertas (priced) | {unreal:+.2f} USD · n={n_priced} |")
    L.append(f"| **expectancy COMBINADA** | **{comb_pnl:+.2f} USD · n={n_comb} · exp/trade={exp_comb:+.2f}** |")
    L.append(f"| peor caso (unpriceable→0) | {worst_pnl:+.2f} USD · n={n_worst} · exp/trade={exp_worst:+.2f} · ({n_unpriceable} sin precio, ${cost_unpriceable:.0f} en riesgo) |")
    L.append(f"| Brier (predicciones resueltas) | {b:.4f} · n={n_brier} |")
    L.append(f"| bucket fallback-bug en cohorte | n={len(fb)} (nota: precios ~0.001 de Fase 7) |")
    L.append("")
    L.append("### Reglas nuevas firing (blocked.jsonl)")
    if new_rules:
        for k, v in sorted(new_rules.items(), key=lambda kv: -kv[1]):
            L.append(f"- {k}: {v}")
    else:
        L.append("- (ninguna regla nueva ha bloqueado todavía)")
    L.append("")
    L.append("### Categorías que sangran (cohorte v7, realizado)")
    if bleeders:
        for c, n, pnl in bleeders:
            L.append(f"- {c}: {pnl:+.2f} USD (n={n})")
    else:
        L.append("- (ninguna categoría en pérdida realizada)")
    L.append("")
    L.append("### Veredicto")
    L.append(f"- expectancy combinada > 0 con n>={N_MIN}: {'SÍ' if crit_exp else 'NO'} "
             f"(exp={exp_comb:+.2f}, n={n_comb})")
    L.append(f"- Brier <= {BRIER_MAX} (estable): {'SÍ' if crit_brier else 'NO'} (Brier={b:.4f})")
    L.append(f"- reglas nuevas bloqueando: {'SÍ' if rules_ok else 'NO'}")
    L.append("")
    L.append(f"## {verdict} — {'expectancy positiva sostenida demostrable' if verdict == 'PASS' else 'NO avanzar a dinero real'}")
    if verdict == "FAIL":
        why = []
        if not crit_exp:
            why.append(f"expectancy combinada {exp_comb:+.2f} (n={n_comb})")
        if not crit_brier:
            why.append(f"Brier {b:.4f}")
        if not rules_ok:
            why.append("reglas nuevas no firing")
        L.append(f"Falla: {', '.join(why)}. Sangra: " +
                 (", ".join(f"{c} ({pnl:+.0f})" for c, _, pnl in bleeders[:3]) or "—") + ".")
    report = "\n".join(L) + "\n"
    print(report)

    if not args.no_write:
        out = pathlib.Path(args.out)
        out.parent.mkdir(parents=True, exist_ok=True)
        out.write_text(report, encoding="utf-8")
        print(f"[written] {out}", file=sys.stderr)
    return 0 if verdict == "PASS" else 1


if __name__ == "__main__":
    sys.exit(main())
