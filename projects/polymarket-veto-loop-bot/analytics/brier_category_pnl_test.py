"""Tests for the category P&L suspension trigger added to brier.py.

Run: pytest analytics/brier_category_pnl_test.py
"""
import datetime as dt
import os
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import brier  # noqa: E402

NOW = dt.datetime(2026, 5, 30, tzinfo=dt.timezone.utc)
RECENT = (NOW - dt.timedelta(days=5)).strftime("%Y-%m-%dT%H:%M:%SZ")
OLD = (NOW - dt.timedelta(days=120)).strftime("%Y-%m-%dT%H:%M:%SZ")
SINCE = NOW - dt.timedelta(days=30)


def _row(cat, pnl, reason="market_closed", ts=RECENT):
    return {"category": cat, "pnl": pnl, "exit_reason": reason, "exit_timestamp": ts}


def test_loser_in_window_suspended():
    rows = [_row("market", -10.0) for _ in range(20)]
    exp = brier.category_pnl_expectancy(rows, SINCE)
    assert exp["market"]["n"] == 20
    assert exp["market"]["expectancy"] == -10.0
    assert brier.category_pnl_suspended(exp, 20, -1.0) == ["market"]


def test_too_few_trades_not_suspended():
    rows = [_row("elections", -10.0) for _ in range(19)]  # n < 20
    exp = brier.category_pnl_expectancy(rows, SINCE)
    assert brier.category_pnl_suspended(exp, 20, -1.0) == []


def test_profitable_not_suspended():
    rows = [_row("geopolitics", +5.0) for _ in range(30)]
    exp = brier.category_pnl_expectancy(rows, SINCE)
    assert brier.category_pnl_suspended(exp, 20, -1.0) == []


def test_old_losses_excluded_by_window():
    # 30 fat losers but all OUTSIDE the 30d window -> not counted -> not suspended.
    rows = [_row("market", -50.0, ts=OLD) for _ in range(30)]
    exp = brier.category_pnl_expectancy(rows, SINCE)
    assert exp.get("market", {"n": 0})["n"] == 0
    assert brier.category_pnl_suspended(exp, 20, -1.0) == []


def test_dead_regime_reasons_excluded():
    # stop_loss / fallback regime is not a REAL_REASON -> ignored even if recent.
    rows = [_row("market", -50.0, reason="stop_loss") for _ in range(30)]
    exp = brier.category_pnl_expectancy(rows, SINCE)
    assert exp.get("market", {"n": 0})["n"] == 0


def test_borderline_loss_not_suspended():
    # avg exactly at threshold (-1.0) is NOT < -1.0 -> not suspended.
    rows = [_row("other", -1.0) for _ in range(25)]
    exp = brier.category_pnl_expectancy(rows, SINCE)
    assert brier.category_pnl_suspended(exp, 20, -1.0) == []
