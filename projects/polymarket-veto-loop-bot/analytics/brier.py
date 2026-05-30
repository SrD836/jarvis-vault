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
# v8 BOT-CAL: count-based FAST auto-suspend. A category with >= SUSPEND_MIN_N
# resolved predictions AND Brier > SUSPEND_BRIER_FAST in the current window is
# suspended immediately, without waiting the 3 consecutive weekly windows. This
# is the "suspende categorias con Brier sistematicamente malo (>0.30 sobre N>=20)"
# rule — it fires as soon as enough data exists, complementing the slower 3w rule.
SUSPEND_BRIER_FAST = 0.30
SUSPEND_MIN_N = 20

# Category P&L suspension (fase categorización): una categoría que PIERDE dinero
# de forma sistemática en una VENTANA RECIENTE se suspende aunque su Brier esté
# bien (mercado eficiente / ya priceado — el caso "sports"). Ventana acotada a
# propósito para evitar el sesgo del régimen muerto pre-v7 + bucket fallback
# (lección R3). Config-gated; alimenta el MISMO suspended_categories.json que ya
# lee brain S1 (sin cambio en brain).
CATEGORY_PNL_SUSPEND = os.environ.get("CATEGORY_PNL_SUSPEND", "1").lower() not in ("0", "false", "no")
CATEGORY_PNL_MIN_N = int(os.environ.get("CATEGORY_PNL_MIN_N", "20"))
CATEGORY_PNL_MAX_LOSS = float(os.environ.get("CATEGORY_PNL_MAX_LOSS", "-1.0"))
CATEGORY_PNL_WINDOW_DAYS = int(os.environ.get("CATEGORY_PNL_WINDOW_DAYS", "30"))
# Cierres realizados que reflejan el edge de la política (excluye stop_loss del
# régimen muerto + el bucket del fallback-bug). Espejo de daily_calibration.REAL_REASONS.
REAL_REASONS = {"market_closed", "target_hit", "no_remaining_edge"}

DEFAULT_PREDICTIONS = "vault/inbox/trading/predictions.jsonl"
DEFAULT_CLOSED = "vault/inbox/trading/closed.jsonl"
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


def write_suspended(path: pathlib.Path, entries: list[tuple[str, str]], now: str) -> None:
    """entries: lista de (categoria, rule). brain S1 solo lee `category`; `rule`
    es para auditoría (distingue Brier-S1 de pnl_expectancy)."""
    payload = {
        "suspended": [
            {"category": c, "since": now, "rule": rule}
            for c, rule in entries
        ]
    }
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2), encoding="utf-8")


def _pnl(r: dict) -> float:
    for k in ("pnl_usd", "pnl"):
        v = r.get(k)
        if v is not None:
            try:
                return float(v)
            except (TypeError, ValueError):
                pass
    return 0.0


def category_pnl_expectancy(closed_rows: list[dict], since: dt.datetime | None) -> dict:
    """Expectancy realizada por categoría sobre cierres REAL_REASONS en la ventana."""
    agg: dict[str, dict] = defaultdict(lambda: {"n": 0, "pnl": 0.0, "wins": 0})
    for r in closed_rows:
        if r.get("exit_reason") not in REAL_REASONS:
            continue
        if since is not None:
            ts = parse_ts(r.get("exit_timestamp") or r.get("exit_time") or "")
            if ts is None or ts < since:
                continue
        a = agg[str(r.get("category") or "uncategorized")]
        p = _pnl(r)
        a["n"] += 1
        a["pnl"] += p
        a["wins"] += 1 if p > 0 else 0
    for a in agg.values():
        a["expectancy"] = a["pnl"] / a["n"] if a["n"] else 0.0
    return dict(agg)


def category_pnl_suspended(expect: dict, min_n: int, max_loss: float) -> list[str]:
    """Categorías que pierden de forma sistemática: n >= min_n y expectancy < max_loss."""
    return sorted(c for c, a in expect.items() if a["n"] >= min_n and a["expectancy"] < max_loss)


def render_pnl_by_category(closed_rows: list[dict]) -> str:
    """P&L realizado por categoría sobre TODO el histórico (deliverable legible)."""
    agg: dict[str, dict] = defaultdict(lambda: {"n": 0, "pnl": 0.0, "wins": 0})
    for r in closed_rows:
        a = agg[str(r.get("category") or "uncategorized")]
        p = _pnl(r)
        a["n"] += 1
        a["pnl"] += p
        a["wins"] += 1 if p > 0 else 0
    if not agg:
        return "### P&L realizado por categoría\n\n_no closed trades_\n"
    lines = ["### P&L realizado por categoría", "",
             "| categoría | n | pnl | win% | avg/trade |", "|---|---|---|---|---|"]
    for c, a in sorted(agg.items(), key=lambda kv: kv[1]["pnl"]):
        wr = 100 * a["wins"] / a["n"] if a["n"] else 0
        avg = a["pnl"] / a["n"] if a["n"] else 0
        lines.append(f"| {c} | {a['n']} | {a['pnl']:+.2f} | {wr:.0f}% | {avg:+.2f} |")
    return "\n".join(lines) + "\n"


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
    ap.add_argument("--closed", default=DEFAULT_CLOSED)
    ap.add_argument("--days", default="7", help='"7", "30", or "all"')
    ap.add_argument("--memory", default=DEFAULT_MEMORY)
    ap.add_argument("--suspended-out", default=DEFAULT_SUSPENDED)
    ap.add_argument("--shadow-ready-out", default=DEFAULT_SHADOW_READY)
    ap.add_argument("--history", default=DEFAULT_HISTORY)
    ap.add_argument("--no-write", action="store_true", help="print only, do not touch memory.md / suspended / shadow_ready")
    ap.add_argument(
        "--skips",
        choices=("include", "exclude", "only"),
        default="include",
        help=(
            "How to treat predictions whose decision was skip / skip_shadow / "
            "veto. v7 P6: passive_resolver backfills outcomes for *every* "
            "evaluated candidate, so 'include' (default) lets Brier score the "
            "LLM's negative calls — proving it acts as a filter, not just a "
            "trader. 'exclude' restricts to actual trade intents; 'only' "
            "isolates the skip-quality view (does the LLM's 'none' actually "
            "ride past edge?)."
        ),
    )
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

    # v7 P6: filter by decision intent per --skips choice. Skip-class
    # decisions are skip, skip_shadow, and explicit veto (E1/E2/M2/P*).
    # Trade-class decisions are buy_yes and buy_no.
    skip_decisions = {"skip", "skip_shadow"}
    def is_skip(r: dict) -> bool:
        d = (r.get("decision") or "").lower()
        return d in skip_decisions or d.startswith("veto") or d.startswith("skip")
    if args.skips == "exclude":
        windowed = [r for r in windowed if not is_skip(r)]
    elif args.skips == "only":
        windowed = [r for r in windowed if is_skip(r)]

    n_total, b_total = brier(windowed)

    by_cat = group_brier(windowed, "category")
    by_hor = group_brier(windowed, "horizon")
    by_edge = group_brier(windowed, "edge_type")

    print(f"# Brier — last {args.days} days (skips={args.skips})")
    print(f"resolved={n_total}  overall_brier={b_total:.4f}")
    print()
    print(render_table("Por categoría", by_cat))
    print(render_table("Por horizonte", by_hor))
    print(render_table("Por edge_type", by_edge))

    # P&L realizado por categoría (deliverable legible) + suspensión por P&L en
    # ventana reciente (config-gated, evita el sesgo del régimen muerto pre-v7).
    closed_rows = load_rows(pathlib.Path(args.closed))
    print(render_pnl_by_category(closed_rows))
    pnl_since = now - dt.timedelta(days=CATEGORY_PNL_WINDOW_DAYS)
    pnl_expect = category_pnl_expectancy(closed_rows, pnl_since)
    pnl_susp = (category_pnl_suspended(pnl_expect, CATEGORY_PNL_MIN_N, CATEGORY_PNL_MAX_LOSS)
                if CATEGORY_PNL_SUSPEND else [])
    print(f"# P&L expectancy ventana {CATEGORY_PNL_WINDOW_DAYS}d (REAL_REASONS; "
          f"suspende si n>={CATEGORY_PNL_MIN_N} y avg<{CATEGORY_PNL_MAX_LOSS}; "
          f"enabled={CATEGORY_PNL_SUSPEND})")
    for c, a in sorted(pnl_expect.items(), key=lambda kv: kv[1]["expectancy"]):
        flag = "  <= SUSPEND" if c in pnl_susp else ""
        print(f"  {c}: n={a['n']} expectancy={a['expectancy']:+.2f}{flag}")

    if args.no_write:
        return 0

    body_lines = [
        f"resolved={n_total} overall_brier={b_total:.4f}",
        "",
        render_table("Por categoría", by_cat),
        render_table("Por horizonte", by_hor),
        render_table("Por edge_type", by_edge),
        render_pnl_by_category(closed_rows),
    ]
    append_memory_section(pathlib.Path(args.memory), "\n".join(body_lines))

    week = now.strftime("%G-W%V")
    hist = update_history(pathlib.Path(args.history), week, by_cat)
    suspended_3w = compute_suspended(hist)
    # v8 BOT-CAL: count-based fast suspend on the current window.
    suspended_count = sorted({
        cat for cat, n, b in by_cat
        if n >= SUSPEND_MIN_N and b > SUSPEND_BRIER_FAST
    })
    # Une las tres fuentes en suspended_categories.json. brain S1 solo lee
    # `category`; `rule` distingue Brier vs P&L para auditoría.
    rules_map: dict[str, str] = {}
    for c in sorted(set(suspended_3w) | set(suspended_count)):
        rules_map[c] = "Brier > 0.25 × 3w"
    for c in pnl_susp:
        rules_map[c] = "pnl_expectancy" if c not in rules_map else "Brier+pnl_expectancy"
    write_suspended(pathlib.Path(args.suspended_out), sorted(rules_map.items()), now.isoformat())
    if suspended_count:
        print(f"S1 (count-based, Brier>{SUSPEND_BRIER_FAST} N>={SUSPEND_MIN_N}): {suspended_count}", file=sys.stderr)
    if suspended_3w:
        print(f"S1 (3-week, Brier>{SUSPEND_BRIER} x{SUSPEND_CONSECUTIVE}w): {suspended_3w}", file=sys.stderr)
    if pnl_susp:
        print(f"S1 (pnl_expectancy, n>={CATEGORY_PNL_MIN_N} avg<{CATEGORY_PNL_MAX_LOSS} en {CATEGORY_PNL_WINDOW_DAYS}d): {pnl_susp}", file=sys.stderr)

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
