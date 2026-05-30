#!/usr/bin/env python3
"""daily_rollup.py — one Obsidian node per day summarizing the polymarket bot.

Replaces the per-event decision-node firehose (gated off via DECISION_NODES in
bot/common/decisionlog/decisionlog.go + bot/exit_monitor/internal/loglosses).
Deterministic, no LLM. Reuses brier.load_rows / brier.parse_ts.

Writes <root>/03-decisions/<date>-polymarket-daily.md with:
  - decisiones del dia (predictions.jsonl): por tipo + top motivos skip/veto
  - trades abiertos hoy (approved.jsonl): por horizonte
  - cierres hoy (closed.jsonl): n, P&L, wins/losses, top-3 por |pnl|
  - calibracion actual (calibration.json): expectancy short, shrink, brier

Idempotent (overwrites the day's file). No retroactive backfill. The daily node
is the ONE node we keep; it is NOT gated by DECISION_NODES (that gates the Go
per-event writers) and is NOT touched by the archive sweep (which only moves
*-polymarket-{veto,trade,loss}-*.md).

CLI:
  python3 analytics/daily_rollup.py [--date YYYY-MM-DD] [--root vault] [--out PATH]
                                    [--dry-run]
Run from the agent-stack dir (cwd) so the default --root "vault" resolves, same
convention as brier.py / daily_calibration.py.
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import sys
from collections import Counter

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from brier import load_rows, parse_ts  # noqa: E402

DEFAULT_ROOT = "vault"
TRADING = "inbox/trading"
DECISIONS = "03-decisions"
TRADE_DECISIONS = {"buy_yes", "buy_no"}


def _day_bounds(date_str: str) -> tuple[dt.datetime, dt.datetime]:
    d = dt.datetime.strptime(date_str, "%Y-%m-%d").replace(tzinfo=dt.timezone.utc)
    return d, d + dt.timedelta(days=1)


def _in_day(ts_str: str, lo: dt.datetime, hi: dt.datetime) -> bool:
    ts = parse_ts(ts_str or "")
    return ts is not None and lo <= ts < hi


def _pnl(r: dict) -> float:
    v = r.get("pnl", r.get("pnl_usd"))
    return float(v) if isinstance(v, (int, float)) else 0.0


def gather(root: pathlib.Path, date_str: str) -> dict:
    lo, hi = _day_bounds(date_str)
    tdir = root / TRADING

    preds = [r for r in load_rows(tdir / "predictions.jsonl")
             if _in_day(r.get("timestamp", ""), lo, hi)]
    decisions = Counter((r.get("decision") or "?") for r in preds)
    skip_reasons = Counter(
        (r.get("skip_reason") or "").strip()
        for r in preds
        if (r.get("decision") or "").lower() not in TRADE_DECISIONS
        and (r.get("skip_reason") or "").strip()
    )

    approved = [r for r in load_rows(tdir / "approved.jsonl")
                if _in_day(r.get("approved_at", ""), lo, hi)]
    opened_by_h = Counter((r.get("horizon") or "?") for r in approved)

    closed = [r for r in load_rows(tdir / "closed.jsonl")
              if _in_day(r.get("exit_timestamp", "") or r.get("exit_time", ""), lo, hi)]
    pnl_total = sum(_pnl(r) for r in closed)
    wins = sum(1 for r in closed if _pnl(r) > 0)
    losses = sum(1 for r in closed if _pnl(r) < 0)
    top = sorted(closed, key=lambda r: abs(_pnl(r)), reverse=True)[:3]

    calib = {}
    cpath = tdir / "calibration.json"
    if cpath.exists():
        try:
            calib = json.loads(cpath.read_text(encoding="utf-8"))
        except json.JSONDecodeError:
            calib = {}

    return {
        "date": date_str, "preds_n": len(preds), "decisions": decisions,
        "skip_reasons": skip_reasons, "opened_n": len(approved),
        "opened_by_h": opened_by_h, "closed_n": len(closed),
        "pnl_total": pnl_total, "wins": wins, "losses": losses,
        "top": top, "calib": calib,
    }


def _table(headers: list[str], rows: list[list]) -> str:
    if not rows:
        return "_ninguno_\n"
    out = ["| " + " | ".join(headers) + " |",
           "|" + "|".join("---" for _ in headers) + "|"]
    for r in rows:
        out.append("| " + " | ".join(str(c) for c in r) + " |")
    return "\n".join(out) + "\n"


def render(g: dict) -> str:
    d = g["date"]
    fm = [
        "---",
        f'title: "Polymarket — resumen diario {d}"',
        "type: decision",
        f"date: {d}",
        "decision_type: daily-rollup",
        "tags: [polymarket, bot, daily, rollup]",
        "related:",
        '  - "[[agents/polymarket-bot]]"',
        '  - "[[agents/polymarket-bot/memory]]"',
        '  - "[[projects/polymarket-veto-loop-bot]]"',
        "---",
    ]
    decisions = sorted(g["decisions"].items(), key=lambda kv: -kv[1])
    skips = g["skip_reasons"].most_common(5)
    opened = sorted(g["opened_by_h"].items())
    top_rows = [[r.get("slug", "?"), f'{_pnl(r):+.2f}', r.get("exit_reason", "")]
                for r in g["top"]]
    c = g.get("calib") or {}
    short = c.get("short", {}) if isinstance(c, dict) else {}
    brier = c.get("brier", {}) if isinstance(c, dict) else {}

    body = [
        "",
        f"# Polymarket — resumen diario {d}",
        "",
        "> Nodo unico del dia (reemplaza el firehose de nodos por evento; el dato bruto vive en `inbox/trading/*.jsonl`, el aprendizaje en `memory.md`).",
        "",
        "## Decisiones del dia",
        "",
        f"- Evaluadas (predictions): **{g['preds_n']}**",
        f"- Abiertas hoy (approved): **{g['opened_n']}**",
        f"- Cerradas hoy (closed): **{g['closed_n']}**  ·  P&L dia: **{g['pnl_total']:+.2f} USD**  ·  wins {g['wins']} / losses {g['losses']}",
        "",
        "### Por tipo de decision",
        "",
        _table(["decision", "n"], [[k, v] for k, v in decisions]),
        "### Top motivos skip/veto",
        "",
        _table(["motivo", "n"], [[k, v] for k, v in skips]),
        "### Trades abiertos por horizonte",
        "",
        _table(["horizonte", "n"], [[k, v] for k, v in opened]),
        "### Top 3 cierres (por |P&L|)",
        "",
        _table(["slug", "pnl", "exit_reason"], top_rows),
        "## Calibracion actual (`calibration.json`)",
        "",
    ]
    if short:
        body.append(
            f"- short: n={short.get('n', 0)} wr={short.get('win_rate', 0):.3f} "
            f"expectancy={short.get('expectancy_per_trade', 0):+.4f}/trade "
            f"total_pnl={short.get('total_pnl', 0):+.2f}"
        )
        body.append(
            f"- kelly_shrink={c.get('kelly_shrink', 1.0)} "
            f"min_edge_override={c.get('min_edge_points_override', 0.0)} "
            f"brier={brier.get('score', 0):.4f} (n={brier.get('n', 0)})"
        )
    else:
        body.append("_(sin calibration.json)_")
    body += ["", "## Human notes", "", "_(no se toca por automatizacion)_", ""]
    return "\n".join(fm) + "\n" + "\n".join(body)


def _notify_day_closed(g: dict) -> None:
    """Mensaje Telegram humano del cierre del dia (solo si hubo cierres)."""
    if g.get("closed_n", 0) <= 0:
        return
    try:
        hermes = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                              "..", "..", "..", "..", "hermes")
        sys.path.insert(0, os.path.abspath(hermes))
        from humanize import notify_human
        notify_human("bot_day_closed", {
            "won": g.get("wins", 0),
            "total": g.get("closed_n", 0),
            "pnl": g.get("pnl_total", 0.0),
        })
    except Exception as e:
        sys.stderr.write(f"[daily_rollup] notify failed: {e}\n")


def main(argv=None) -> int:
    ap = argparse.ArgumentParser(description="Polymarket daily rollup node")
    ap.add_argument("--date", default=dt.datetime.now(dt.timezone.utc).strftime("%Y-%m-%d"))
    ap.add_argument("--root", default=DEFAULT_ROOT, help="vault root (default: vault)")
    ap.add_argument("--out", default=None, help="override output path")
    ap.add_argument("--dry-run", action="store_true", help="print, do not write")
    args = ap.parse_args(argv)

    root = pathlib.Path(args.root)
    g = gather(root, args.date)
    md = render(g)

    if args.dry_run:
        print(md)
        return 0

    out = pathlib.Path(args.out) if args.out else root / DECISIONS / f"{args.date}-polymarket-daily.md"
    out.parent.mkdir(parents=True, exist_ok=True)
    out.write_text(md, encoding="utf-8")
    _notify_day_closed(g)
    print(f"wrote {out}  (preds={g['preds_n']} opened={g['opened_n']} "
          f"closed={g['closed_n']} pnl={g['pnl_total']:+.2f})")
    return 0


if __name__ == "__main__":
    sys.exit(main())
