#!/usr/bin/env python3
"""Honest expectancy report for the polymarket bot (v7 policy cohort, simulation).

Defeats the survivorship bias of "P&L desde v7" (winners close fast via
target_hit/market_closed, longshot losers sit OPEN because v7 has no stop) by
NETTING the realized v7-policy P&L with the mark-to-market unrealized P&L of the
still-open book (mark_to_market.json, produced by the Go cmd reusing
polyclient.ExitPrice). Reports a market-priced figure AND a worst case (open book
resolves to 0), plus Brier on resolved predictions and rule firing. Emits a
PASS/FAIL verdict against David's criterion: expectancy/trade > 0 over N>=50,
Brier <= 0.25, new rules firing — AND robust to the open book cratering, because
we do NOT advance to real money on a fragile snapshot.

v7 cohort = exit_reason in {market_closed, target_hit, no_remaining_edge} and not
the fallback-bug bucket — same definition as the dashboard pnl_v7 metric. The rest
(stop_loss / fallback / dead regime) is the "histórico contaminado" David excludes.
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
import brier  # noqa: E402  reuse load_rows, brier, _pnl

N_MIN = int(os.environ.get("EXPECTANCY_N_MIN", "50"))
BRIER_MAX = float(os.environ.get("EXPECTANCY_BRIER_MAX", "0.25"))
V7_REASONS = {"market_closed", "target_hit", "no_remaining_edge"}
FALLBACK = "lastTradePrice_fallback"
T = "vault/inbox/trading/"


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--closed", default=T + "closed.jsonl")
    ap.add_argument("--predictions", default=T + "predictions.jsonl")
    ap.add_argument("--mtm", default=T + "mark_to_market.json")
    ap.add_argument("--blocked", default=T + "blocked.jsonl")
    ap.add_argument("--out", default="vault/outputs/expectancy-report-" + dt.date.today().isoformat() + ".md")
    ap.add_argument("--no-write", action="store_true")
    args = ap.parse_args()

    # --- realized v7-policy cohort (winners-heavy by construction; los perdedores
    #     v7 siguen ABIERTOS) ---
    closed = brier.load_rows(pathlib.Path(args.closed))
    v7 = [r for r in closed if r.get("exit_reason") in V7_REASONS and r.get("exit_price_source") != FALLBACK]
    realized_pnl = sum(brier._pnl(r) for r in v7)
    n_real = len(v7)
    wins = sum(1 for r in v7 if brier._pnl(r) > 0)
    exp_real = realized_pnl / n_real if n_real else 0.0
    wr_real = 100 * wins / n_real if n_real else 0.0
    cont_pnl = sum(brier._pnl(r) for r in closed) - realized_pnl
    n_cont = len(closed) - n_real

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
    n_open = int(mtm.get("n", 0))
    open_stake = sum(float(p.get("size_usd", 0) or 0) for p in mtm.get("positions", []))

    # --- combined honest expectancy ---
    n_comb = n_real + n_priced
    comb_pnl = realized_pnl + unreal
    exp_comb = comb_pnl / n_comb if n_comb else 0.0
    # worst case: TODO el libro abierto resuelve a 0 -> pierde el stake completo.
    worst_pnl = realized_pnl - open_stake
    n_worst = n_real + n_open
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

    # --- per-category P&L of the open book (dónde está el riesgo latente) ---
    cat: dict[str, dict] = defaultdict(lambda: {"n": 0, "unreal": 0.0})
    for p in mtm.get("positions", []):
        a = cat[str(p.get("category") or "uncategorized")]
        a["n"] += 1
        a["unreal"] += float(p.get("unrealized_pnl") or 0.0)
    bleeders = sorted(((c, a["n"], a["unreal"]) for c, a in cat.items()), key=lambda x: x[2])

    # --- verdict: PASS solo si positivo a mercado Y robusto al peor caso ---
    crit_market = exp_comb > 0 and n_comb >= N_MIN
    crit_worst = exp_worst > 0
    crit_brier = n_brier > 0 and b <= BRIER_MAX
    verdict = "PASS" if (crit_market and crit_worst and crit_brier and rules_ok) else "FAIL"

    L = [
        f"# Expectancy report — v7 policy cohort — {dt.date.today().isoformat()}",
        "",
        "> Honesto: netea el realizado v7 (ganadores cerrados) con el mark-to-market del libro abierto",
        "> (perdedores latentes). Sin esto, el 98% wr es puro sesgo de supervivencia.",
        "",
        "| métrica | valor |",
        "|---|---|",
        f"| realizado v7 (cerrados, policy) | {realized_pnl:+.2f} USD · n={n_real} · wr={wr_real:.1f}% · exp/trade={exp_real:+.2f} |",
        f"| histórico contaminado (excluido) | {cont_pnl:+.2f} USD · n={n_cont} (stop_loss/fallback régimen muerto) |",
        f"| mark-to-market abiertas | {unreal:+.2f} USD · n={n_priced} priced / {n_open} |",
        f"| **expectancy COMBINADA (a mercado)** | **{comb_pnl:+.2f} USD · n={n_comb} · exp/trade={exp_comb:+.2f}** |",
        f"| peor caso (abiertas→0) | {worst_pnl:+.2f} USD · n={n_worst} · exp/trade={exp_worst:+.2f} · (stake abierto ${open_stake:.0f}) |",
        f"| Brier (predicciones resueltas) | {b:.4f} · n={n_brier} |",
        "",
        "### Reglas nuevas firing (blocked.jsonl)",
    ]
    if new_rules:
        L += [f"- {k}: {v}" for k, v in sorted(new_rules.items(), key=lambda kv: -kv[1])]
    else:
        L.append("- (ninguna regla nueva ha bloqueado todavía)")
    L += ["", "### Libro abierto por categoría (riesgo latente, mark-to-market)"]
    if bleeders:
        L += [f"- {c}: {u:+.2f} USD (n={n})" for c, n, u in bleeders]
    else:
        L.append("- (sin posiciones abiertas)")
    L += [
        "",
        "### Veredicto (criterio de David)",
        f"- expectancy combinada > 0 con n>={N_MIN}: {'SÍ' if crit_market else 'NO'} (exp={exp_comb:+.2f}, n={n_comb})",
        f"- robusto al peor caso (abiertas→0) > 0: {'SÍ' if crit_worst else 'NO'} (exp={exp_worst:+.2f})",
        f"- Brier <= {BRIER_MAX}: {'SÍ' if crit_brier else 'NO'} (Brier={b:.4f})",
        f"- reglas nuevas bloqueando: {'SÍ' if rules_ok else 'NO'}",
        "",
        f"## {verdict} — {'expectancy positiva sostenida demostrable' if verdict == 'PASS' else 'NO avanzar a dinero real'}",
    ]
    if verdict == "FAIL":
        why = []
        if not crit_market:
            why.append(f"combinada a mercado {exp_comb:+.2f}")
        if not crit_worst:
            why.append(f"peor caso {exp_worst:+.2f} (libro longshot abierto ${open_stake:.0f})")
        if not crit_brier:
            why.append(f"Brier {b:.4f}")
        if not rules_ok:
            why.append("reglas no firing")
        L.append(f"Falla: {'; '.join(why)}.")
        L.append(f"Sangra (abiertas): " + (", ".join(f"{c} ({u:+.0f})" for c, _, u in bleeders[:3] if u < 0) or "—") + ".")
        L.append("La estrategia v7 cierra ganadores rápido y deja correr perdedores (sin stop) → el "
                 "realizado SIEMPRE parece bueno hasta que el libro abierto resuelve. Prueba de "
                 "sostenibilidad = ventana forward con R1/R3 activos desde el inicio (ya bloquean longshots).")
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
