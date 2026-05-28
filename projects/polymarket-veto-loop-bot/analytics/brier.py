#!/usr/bin/env python3
"""
Brier-score loop for polymarket-veto-loop-bot v7.

Reads predictions.jsonl (one JSON object per line, schema in
bot/common/predictions/predictions.go) and computes Brier =
mean((estimated_prob - outcome)^2) for the rows whose outcome has been
backfilled by the exit monitor (i.e. the underlying market resolved).

Groups by category, horizon, and edge_type. Auto-suspend rule (S1):
when a category posts Brier > 0.25 in three consecutive weekly windows
the category is written to suspended_categories.json so the brain blocks
new candidates in that bucket.

Shadow→simulation release: when at least one category has Brier < 0.20
on the most recent 7-day window AND we have >= MIN_RESOLVED_FOR_RELEASE
resolved predictions in that category, write shadow_ready.json with the
qualifying category. The mode toggle (config.json: "shadow" → "simulation")
is intentionally manual — this script never edits config.json itself.
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import sys
from collections import defaultdict
from typing import Iterable

MIN_RESOLVED_FOR_RELEASE = 20
SUSPEND_BRIER = 0.25
SUSPEND_CONSECUTIVE = 3
RELEASE_BRIER = 0.20

DEFAULT_PREDICTIONS = "vault/inbox/trading/predictions.jsonl"
DEFAULT_MEMORY = "vault/agents/polymarket-bot/memory.md"
DEFAULT_SUSPENDED = "vault/inbox/trading/suspended_categories.json"
DEFAULT_SHADOW_READY = "vault/inbox/trading/shadow_ready.json"
DEFAULT_HISTORY = "vault/inbox/trading/brier_history.json"


def load_rows(path: pathlib.Path) -> list[dict]:
    if not path.exists():
        return []
    rows: list[dict] = []
    with path.open("r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                rows.append(json.loads(line))
            except json.JSONDecodeError:
                continue
    return rows


def parse_ts(s: str) -> dt.datetime | None:
    if not s:
        return None
    for fmt in (
        "%Y-%m-%dT%H:%M:%S.%fZ",
        "%Y-%m-%dT%H:%M:%SZ",
        "%Y-%m-%dT%H:%M:%S%z",
    ):
        try:
            t = dt.datetime.strptime(s, fmt)
            if t.tzinfo is None:
                t = t.replace(tzinfo=dt.timezone.utc)
            return t
        except ValueError:
            continue
    try:
        return dt.datetime.fromisoformat(s.replace("Z", "+00:00"))
    except ValueError:
        return None


def in_window(row: dict, since: dt.datetime | None) -> bool:
    if since is None:
        return True
    ts = parse_ts(row.get("timestamp", ""))
    return ts is not None and ts >= since


def brier(rows: Iterable[dict]) -> tuple[int, float]:
    rows = [
        r for r in rows
        if r.get("outcome") is not None
        and isinstance(r.get("estimated_prob"), (int, float))
        and 0.0 <= r["estimated_prob"] <= 1.0
        and r["outcome"] in (0, 0.0, 1, 1.0)
    ]
    n = len(rows)
    if n == 0:
        return 0, 0.0
    total = sum((r["estimated_prob"] - float(r["outcome"])) ** 2 for r in rows)
    return n, total / n


def group_brier(rows: list[dict], key: str) -> list[tuple[str, int, float]]:
    buckets: dict[str, list[dict]] = defaultdict(list)
    for r in rows:
        buckets[str(r.get(key) or "unknown")].append(r)
    out = []
    for k, v in buckets.items():
        n, b = brier(v)
        if n > 0:
            out.append((k, n, b))
    out.sort(key=lambda x: x[2])
    return out


def render_table(title: str, rows: list[tuple[str, int, float]]) -> str:
    if not rows:
        return f"### {title}\n\n_no resolved rows_\n"
    lines = [f"### {title}", "", "| group | n | brier |", "|-------|---|-------|"]
    for name, n, b in rows:
        lines.append(f"| {name} | {n} | {b:.4f} |")
    return "\n".join(lines) + "\n"


def update_history(history_path: pathlib.Path, week: str, by_cat: list[tuple[str, int, float]]) -> dict:
    hist = {"weeks": {}}
    if history_path.exists():
        try:
            hist = json.loads(history_path.read_text(encoding="utf-8"))
        except json.JSONDecodeError:
            pass
    hist.setdefault("weeks", {})
    hist["weeks"][week] = {name: {"n": n, "brier": b} for name, n, b in by_cat}
    history_path.parent.mkdir(parents=True, exist_ok=True)
    history_path.write_text(json.dumps(hist, indent=2), encoding="utf-8")
    return hist


def compute_suspended(hist: dict) -> list[str]:
    weeks = sorted(hist.get("weeks", {}).keys())
    if len(weeks) < SUSPEND_CONSECUTIVE:
        return []
    last_n = weeks[-SUSPEND_CONSECUTIVE:]
    bad_per_week: list[set[str]] = []
    for w in last_n:
        bad_per_week.append({
            cat for cat, stats in hist["weeks"][w].items()
            if stats.get("n", 0) > 0 and stats.get("brier", 0.0) > SUSPEND_BRIER
        })
    if not bad_per_week:
        return []
    consistent = bad_per_week[0].copy()
    for s in bad_per_week[1:]:
        consistent &= s
    return sorted(consistent)


def write_suspended(path: pathlib.Path, cats: list[str], now: str) -> None:
    payload = {
        "suspended": [
            {"category": c, "since": now, "rule": "Brier > 0.25 × 3w"}
            for c in cats
        ]
    }
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2), encoding="utf-8")


def append_memory_section(memory_path: pathlib.Path, body: str) -> None:
    if not memory_path.parent.exists():
        memory_path.parent.mkdir(parents=True, exist_ok=True)
    stamp = dt.datetime.now(dt.timezone.utc).strftime("%Y-%m-%d")
    section = f"\n## Brier score (semanal) — {stamp}\n\n{body}\n"
    with memory_path.open("a", encoding="utf-8") as f:
        f.write(section)


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--predictions", default=DEFAULT_PREDICTIONS)
    ap.add_argument("--days", default="7", help='"7", "30", or "all"')
    ap.add_argument("--memory", default=DEFAULT_MEMORY)
    ap.add_argument("--suspended-out", default=DEFAULT_SUSPENDED)
    ap.add_argument("--shadow-ready-out", default=DEFAULT_SHADOW_READY)
    ap.add_argument("--history", default=DEFAULT_HISTORY)
    ap.add_argument("--no-write", action="store_true", help="print only, do not touch memory.md / suspended / shadow_ready")
    args = ap.parse_args()

    pred_path = pathlib.Path(args.predictions)
    rows = load_rows(pred_path)
    now = dt.datetime.now(dt.timezone.utc)
    if args.days == "all":
        since = None
    else:
        try:
            since = now - dt.timedelta(days=int(args.days))
        except ValueError:
            print(f"invalid --days {args.days!r}", file=sys.stderr)
            return 2

    windowed = [r for r in rows if in_window(r, since)]
    n_total, b_total = brier(windowed)

    by_cat = group_brier(windowed, "category")
    by_hor = group_brier(windowed, "horizon")
    by_edge = group_brier(windowed, "edge_type")

    print(f"# Brier — last {args.days} days")
    print(f"resolved={n_total}  overall_brier={b_total:.4f}")
    print()
    print(render_table("Por categoría", by_cat))
    print(render_table("Por horizonte", by_hor))
    print(render_table("Por edge_type", by_edge))

    if args.no_write:
        return 0

    body_lines = [
        f"resolved={n_total} overall_brier={b_total:.4f}",
        "",
        render_table("Por categoría", by_cat),
        render_table("Por horizonte", by_hor),
        render_table("Por edge_type", by_edge),
    ]
    append_memory_section(pathlib.Path(args.memory), "\n".join(body_lines))

    week = now.strftime("%G-W%V")
    hist = update_history(pathlib.Path(args.history), week, by_cat)
    suspended = compute_suspended(hist)
    write_suspended(pathlib.Path(args.suspended_out), suspended, now.isoformat())
    if suspended:
        print(f"S1: suspended categories (Brier > {SUSPEND_BRIER} × {SUSPEND_CONSECUTIVE}w): {suspended}", file=sys.stderr)

    qualifying = [
        (name, n, b) for name, n, b in by_cat
        if n >= MIN_RESOLVED_FOR_RELEASE and b < RELEASE_BRIER
    ]
    shadow_payload = {
        "ready": bool(qualifying),
        "evaluated_at": now.isoformat(),
        "qualifying_categories": [
            {"category": name, "n": n, "brier": b} for name, n, b in qualifying
        ],
        "release_threshold": {"brier_lt": RELEASE_BRIER, "min_n": MIN_RESOLVED_FOR_RELEASE},
    }
    out = pathlib.Path(args.shadow_ready_out)
    out.parent.mkdir(parents=True, exist_ok=True)
    out.write_text(json.dumps(shadow_payload, indent=2), encoding="utf-8")
    if qualifying:
        print(
            "READY_TO_RESUME: "
            + ", ".join(name for name, _, _ in qualifying),
            file=sys.stderr,
        )
    return 0


if __name__ == "__main__":
    sys.exit(main())
