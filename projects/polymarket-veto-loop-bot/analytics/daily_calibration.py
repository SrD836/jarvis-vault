#!/usr/bin/env python3
"""
Daily calibration cycle for polymarket-veto-loop-bot v8.

Deterministic — no LLM, no Claude, no network. Each day it:
  1. Reads closed.jsonl, computes realized SHORT-horizon expectancy over a
     rolling window, EXCLUDING audit/rebalance/manual closes (only the v7
     thesis exits market_closed / target_hit / no_remaining_edge count, since
     those are the strategy's real outcomes).
  2. Reads predictions.jsonl, computes the short-horizon Brier score (reusing
     analytics/brier.py).
  3. Derives an adaptive Kelly shrink and an optional min_edge_points override
     and writes them to inbox/trading/calibration.json. The Go executor reads
     kelly_shrink (sizes overconfident edges down); the brain reads
     min_edge_points_override (raises the edge gate when expectancy is negative).
  4. Appends a "Calibracion diaria" section to memory.md (idempotent per day).

It NEVER edits config.json (design rule): config.json holds the static
defaults; this file is the learned adaptive layer on top. With too little data
(n < N_MIN) it deliberately emits shrink=1.0 and no override — we do not fake
calibration off two samples.
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import sys

# Reuse the Brier loop's IO + scoring instead of reimplementing it.
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import brier  # noqa: E402

DEFAULT_CLOSED = "vault/inbox/trading/closed.jsonl"
DEFAULT_PREDICTIONS = "vault/inbox/trading/predictions.jsonl"
DEFAULT_CALIBRATION = "vault/inbox/trading/calibration.json"
DEFAULT_MEMORY = "vault/agents/polymarket-bot/memory.md"
DEFAULT_CONFIG = "vault/projects/polymarket-veto-loop-bot/bot/config.json"

# Only v7 thesis exits are the strategy's real outcomes. Everything else
# (audit liquidations, horizon rebalances, manual resets) is excluded so the
# expectancy reflects the policy, not bookkeeping events.
REAL_REASONS = {"market_closed", "target_hit", "no_remaining_edge"}

N_MIN = 10            # min real short closes before we act on expectancy
SHRINK_FLOOR = 0.25   # never shrink Kelly below 1/4 of the edge
BRIER_N_MIN = 20      # min resolved short predictions before Brier blends in
MIN_EDGE_STEP = 0.01  # bump applied to the edge gate when expectancy < 0
MIN_EDGE_CAP = 0.10   # ceiling for the edge-gate override


def load_cfg_min_edge(config_path: str) -> float:
    """Read min_edge_points from config.json (default 0.05 mirrors Go applyDefaults)."""
    try:
        cfg = json.loads(pathlib.Path(config_path).read_text(encoding="utf-8"))
        v = cfg.get("min_edge_points")
        if isinstance(v, (int, float)) and v > 0:
            return float(v)
    except (OSError, json.JSONDecodeError):
        pass
    return 0.05


def expectancy_short(closed_rows: list[dict], since: dt.datetime | None) -> dict:
    """Realized expectancy of SHORT-horizon thesis closes within the window."""
    pnls: list[float] = []
    for r in closed_rows:
        reason = r.get("exit_reason") or r.get("reason")
        if reason not in REAL_REASONS:
            continue
        hz = r.get("horizon")
        d = r.get("days_open")
        is_short = (hz == "short") or (hz in (None, "") and isinstance(d, (int, float)) and d <= 2)
        if not is_short:
            continue
        ts = brier.parse_ts(r.get("exit_timestamp") or r.get("exit_time") or "")
        if since is not None and (ts is None or ts < since):
            continue
        p = r.get("pnl")
        if p is None:
            p = r.get("pnl_usd")
        if p is None:
            continue
        pnls.append(float(p))

    n = len(pnls)
    wins = [x for x in pnls if x > 0]
    losses = [x for x in pnls if x < 0]
    wr = len(wins) / n if n else 0.0
    aw = sum(wins) / len(wins) if wins else 0.0
    al = sum(losses) / len(losses) if losses else 0.0
    payoff = (aw / abs(al)) if al else 0.0
    be = (1.0 / (1.0 + payoff)) if payoff > 0 else 0.0
    exp = wr * aw + (1 - wr) * al
    return {
        "n": n, "wins": len(wins), "losses": len(losses),
        "win_rate": round(wr, 4), "avg_win": round(aw, 2), "avg_loss": round(al, 2),
        "payoff": round(payoff, 4), "breakeven_win_rate": round(be, 4),
        "expectancy_per_trade": round(exp, 4), "total_pnl": round(sum(pnls), 2),
    }


def brier_short(pred_rows: list[dict], since: dt.datetime | None) -> tuple[int, float]:
    """Short-horizon Brier over resolved predictions in the window."""
    rows = [
        r for r in pred_rows
        if r.get("horizon") == "short" and (since is None or brier.in_window(r, since))
    ]
    return brier.brier(rows)


def derive_shrink(exp: dict, brier_n: int, brier_score: float) -> tuple[float, str]:
    """Kelly shrink from realized payoff/win-rate, blended with Brier when ripe.

    shrink = clamp(win_rate / breakeven_win_rate, FLOOR, 1.0), times a Brier
    penalty when >= BRIER_N_MIN resolved. n < N_MIN -> 1.0 (no fake calibration).
    """
    n = exp["n"]
    if n < N_MIN:
        return 1.0, f"n<{N_MIN}: insufficient short closes, shrink=1.0 (no calibration)"
    payoff = exp["payoff"]
    if payoff <= 0:
        shrink = 1.0
        note = "no losers in window, shrink_exp=1.0"
    else:
        be = 1.0 / (1.0 + payoff)
        shrink = exp["win_rate"] / be if be > 0 else 1.0
        note = f"wr/breakeven={exp['win_rate']:.3f}/{be:.3f}={shrink:.3f}"
    if brier_n >= BRIER_N_MIN:
        bf = 1.0 - max(0.0, brier_score - 0.20) / 0.20
        bf = min(1.0, max(0.5, bf))
        shrink *= bf
        note += f" * brier_factor={bf:.3f} (brier={brier_score:.3f} n={brier_n})"
    shrink = min(1.0, max(SHRINK_FLOOR, shrink))
    return round(shrink, 4), note


def derive_min_edge_override(exp: dict, cfg_min_edge: float) -> float:
    """Raise the edge gate by one step (capped) when short expectancy is negative.

    Returns 0.0 (no override) when expectancy is non-negative or data is thin,
    so the brain falls back to the config default and the loop self-corrects.
    """
    if exp["n"] >= N_MIN and exp["expectancy_per_trade"] < 0:
        return round(min(cfg_min_edge + MIN_EDGE_STEP, MIN_EDGE_CAP), 4)
    return 0.0


def build_payload(now: dt.datetime, window_days, exp: dict, brier_n: int,
                  brier_score: float, shrink: float, shrink_note: str,
                  min_edge_override: float, cfg_min_edge: float) -> dict:
    rationale = (
        f"short n={exp['n']} wr={exp['win_rate']:.3f} expectancy={exp['expectancy_per_trade']:+.4f}/trade; "
        f"shrink={shrink} ({shrink_note}); "
        f"min_edge {cfg_min_edge:.3f}->{min_edge_override or cfg_min_edge:.3f}"
    )
    return {
        "generated_at": now.isoformat(),
        "window_days": window_days,
        "short": exp,
        "brier": {"n": brier_n, "score": round(brier_score, 4)},
        "kelly_shrink": shrink,
        "min_edge_points_override": min_edge_override,
        "rationale": rationale,
    }


def render_memory_section(payload: dict, stamp: str) -> str:
    e = payload["short"]
    b = payload["brier"]
    return (
        f"\n## Calibracion diaria — {stamp}\n\n"
        f"- Ventana: {payload['window_days']}d · short thesis-closes n={e['n']} "
        f"(W{e['wins']}/L{e['losses']})\n"
        f"- Win rate {e['win_rate']*100:.1f}% · avg win ${e['avg_win']} · avg loss ${e['avg_loss']} "
        f"· payoff {e['payoff']} (break-even wr {e['breakeven_win_rate']*100:.1f}%)\n"
        f"- Expectancy/trade ${e['expectancy_per_trade']} · total ${e['total_pnl']}\n"
        f"- Brier short: n={b['n']} score={b['score']}\n"
        f"- **kelly_shrink={payload['kelly_shrink']}** · "
        f"min_edge_points_override={payload['min_edge_points_override'] or 'ninguno'}\n"
        f"- {payload['rationale']}\n"
    )


def append_memory(memory_path: pathlib.Path, section: str, stamp: str) -> bool:
    """Append the section unless today's section already exists (idempotent)."""
    marker = f"## Calibracion diaria — {stamp}"
    if memory_path.exists() and marker in memory_path.read_text(encoding="utf-8"):
        return False
    memory_path.parent.mkdir(parents=True, exist_ok=True)
    with memory_path.open("a", encoding="utf-8") as f:
        f.write(section)
    return True


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--closed", default=DEFAULT_CLOSED)
    ap.add_argument("--predictions", default=DEFAULT_PREDICTIONS)
    ap.add_argument("--calibration-out", default=DEFAULT_CALIBRATION)
    ap.add_argument("--memory", default=DEFAULT_MEMORY)
    ap.add_argument("--config", default=DEFAULT_CONFIG)
    ap.add_argument("--days", default="7")
    ap.add_argument("--no-write", action="store_true", help="print only, touch nothing")
    args = ap.parse_args()

    now = dt.datetime.now(dt.timezone.utc)
    try:
        window_days = int(args.days)
        since = now - dt.timedelta(days=window_days)
    except ValueError:
        window_days, since = 0, None  # "all"

    closed_rows = brier.load_rows(pathlib.Path(args.closed))
    pred_rows = brier.load_rows(pathlib.Path(args.predictions))
    cfg_min_edge = load_cfg_min_edge(args.config)

    exp = expectancy_short(closed_rows, since)
    brier_n, brier_score = brier_short(pred_rows, since)
    shrink, shrink_note = derive_shrink(exp, brier_n, brier_score)
    min_edge_override = derive_min_edge_override(exp, cfg_min_edge)
    payload = build_payload(now, window_days, exp, brier_n, brier_score,
                            shrink, shrink_note, min_edge_override, cfg_min_edge)

    print(json.dumps(payload, indent=2))

    if args.no_write:
        return 0

    out = pathlib.Path(args.calibration_out)
    out.parent.mkdir(parents=True, exist_ok=True)
    out.write_text(json.dumps(payload, indent=2), encoding="utf-8")

    stamp = now.strftime("%Y-%m-%d")
    wrote = append_memory(pathlib.Path(args.memory), render_memory_section(payload, stamp), stamp)
    print(f"\ncalibration.json written; memory section {'appended' if wrote else 'already present'} for {stamp}",
          file=sys.stderr)
    return 0


if __name__ == "__main__":
    sys.exit(main())
