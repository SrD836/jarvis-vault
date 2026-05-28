#!/usr/bin/env python3
"""
P4 retrospective: did v6's P0_floor=0.10 cut profitable trades, or block
genuine traps?

Reads closed.jsonl, splits trades by entry_price band, reports PnL net and
win rate per band. Output goes to stdout AND optionally appended to a
markdown report under vault/analysis/performance/.

The bands mirror the v7 floor decision:
   below 0.03   — always rejected (P0 always applies)
   0.03 — 0.05  — newly visible only when horizon≤7 AND liquidity≥$20k (v7 relax)
   0.05 — 0.10  — newly visible under v7 (was vetoed under v6)
   0.10+        — always visible under both v6 and v7

Usage:
    python3 audit_p0_floor.py [--closed PATH] [--report PATH]
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import pathlib
import sys
from collections import defaultdict

DEFAULT_CLOSED = "vault/inbox/trading/closed.jsonl"


def load_rows(path: pathlib.Path) -> list[dict]:
    if not path.exists():
        return []
    out = []
    with path.open("r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                out.append(json.loads(line))
            except json.JSONDecodeError:
                continue
    return out


def band(p: float) -> str:
    if p < 0.03:
        return "<0.03"
    if p < 0.05:
        return "0.03-0.05"
    if p < 0.10:
        return "0.05-0.10"
    if p < 0.30:
        return "0.10-0.30"
    if p < 0.70:
        return "0.30-0.70"
    if p < 0.95:
        return "0.70-0.95"
    return ">=0.95"


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--closed", default=DEFAULT_CLOSED)
    ap.add_argument(
        "--report",
        default="vault/analysis/performance/p0_audit_2026-05.md",
        help="markdown report path; pass empty string to skip",
    )
    args = ap.parse_args()

    rows = load_rows(pathlib.Path(args.closed))
    if not rows:
        print(f"no rows in {args.closed}", file=sys.stderr)
        return 1

    # Exclude trades closed by manual_reset_v7 and horizon-quota-rebalance —
    # those are systemic cleanups that don't reflect strategy edge.
    excluded_reasons = {"manual_reset_v7"}
    filtered = []
    for r in rows:
        reason = (r.get("exit_reason") or "").lower()
        if any(reason.startswith(x) for x in ("horizon-quota", "manual_reset")):
            continue
        if reason in excluded_reasons:
            continue
        if not isinstance(r.get("entry_price"), (int, float)):
            continue
        filtered.append(r)

    stats: dict[str, dict[str, float]] = defaultdict(
        lambda: {"n": 0, "wins": 0, "pnl_sum": 0.0, "pnl_min": 0.0, "pnl_max": 0.0}
    )
    for r in filtered:
        b = band(float(r["entry_price"]))
        s = stats[b]
        pnl = float(r.get("pnl", 0) or 0)
        s["n"] += 1
        s["pnl_sum"] += pnl
        if pnl > 0:
            s["wins"] += 1
        s["pnl_min"] = min(s["pnl_min"], pnl)
        s["pnl_max"] = max(s["pnl_max"], pnl)

    band_order = [
        "<0.03",
        "0.03-0.05",
        "0.05-0.10",
        "0.10-0.30",
        "0.30-0.70",
        "0.70-0.95",
        ">=0.95",
    ]
    lines = []
    lines.append(f"# P0 floor audit — {dt.date.today().isoformat()}")
    lines.append("")
    lines.append(f"Total trades (excluding manual_reset_v7 and horizon-quota rebalance): {len(filtered)}")
    lines.append("")
    lines.append("| band | n | wins | win_rate | pnl_sum | avg_pnl | min_pnl | max_pnl |")
    lines.append("|------|---|------|----------|---------|---------|---------|---------|")
    for b in band_order:
        if b not in stats:
            continue
        s = stats[b]
        wr = s["wins"] / s["n"] if s["n"] else 0.0
        avg = s["pnl_sum"] / s["n"] if s["n"] else 0.0
        lines.append(
            f"| {b} | {int(s['n'])} | {int(s['wins'])} | {wr*100:.0f}% | "
            f"{s['pnl_sum']:+.2f} | {avg:+.2f} | {s['pnl_min']:+.2f} | {s['pnl_max']:+.2f} |"
        )
    lines.append("")
    sub = sum(stats[b]["pnl_sum"] for b in ("<0.03", "0.03-0.05", "0.05-0.10") if b in stats)
    n_sub = sum(int(stats[b]["n"]) for b in ("<0.03", "0.03-0.05", "0.05-0.10") if b in stats)
    lines.append("## v6 → v7 floor decision")
    lines.append("")
    lines.append(
        f"Trades that WOULD HAVE been blocked by v6 P0_floor=0.10 but pass v7 P0_floor=0.05: "
        f"**{n_sub}** trades, net PnL **${sub:+.2f}**."
    )
    if sub > 0:
        lines.append(
            "→ v6 floor was costing edge. v7 relaxation justified."
        )
    elif sub < -100:
        lines.append(
            "→ v6 floor was protecting capital. **Consider keeping P0_floor=0.10** — v7 relaxation may reintroduce losses."
        )
    else:
        lines.append(
            "→ Near break-even. Floor relaxation is neutral; v7 edge gate (E1/E2) should do the filtering."
        )

    report = "\n".join(lines)
    print(report)

    if args.report:
        out = pathlib.Path(args.report)
        out.parent.mkdir(parents=True, exist_ok=True)
        out.write_text(report + "\n", encoding="utf-8")
        print(f"\nwrote {out}", file=sys.stderr)
    return 0


if __name__ == "__main__":
    sys.exit(main())
