#!/usr/bin/env python3
"""Stdlib-unittest tests for daily_calibration.py (no pytest dependency)."""
import datetime as dt
import os
import pathlib
import sys
import tempfile
import unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import daily_calibration as dc  # noqa: E402

NOW = dt.datetime.now(dt.timezone.utc)
SINCE = NOW - dt.timedelta(days=7)
TS = NOW.strftime("%Y-%m-%dT%H:%M:%SZ")


def closes(wins, win_amt, losses, loss_amt, reason="market_closed", hz="short"):
    rows = [{"exit_reason": reason, "horizon": hz, "exit_timestamp": TS, "pnl": win_amt} for _ in range(wins)]
    rows += [{"exit_reason": reason, "horizon": hz, "exit_timestamp": TS, "pnl": -loss_amt} for _ in range(losses)]
    return rows


class TestExpectancy(unittest.TestCase):
    def test_negative_expectancy_shrinks_and_bumps_edge(self):
        rows = closes(6, 10, 6, 20)  # wr .5, payoff .5, expectancy -5
        exp = dc.expectancy_short(rows, SINCE)
        self.assertEqual(exp["n"], 12)
        self.assertLess(exp["expectancy_per_trade"], 0)
        shrink, _ = dc.derive_shrink(exp, 0, 0.0)
        self.assertLess(shrink, 1.0)
        self.assertGreaterEqual(shrink, dc.SHRINK_FLOOR)
        self.assertAlmostEqual(dc.derive_min_edge_override(exp, 0.05), 0.06, places=4)

    def test_positive_expectancy_no_shrink_no_override(self):
        rows = closes(8, 20, 4, 10)  # wr .667, payoff 2, expectancy +10
        exp = dc.expectancy_short(rows, SINCE)
        self.assertGreater(exp["expectancy_per_trade"], 0)
        shrink, _ = dc.derive_shrink(exp, 0, 0.0)
        self.assertEqual(shrink, 1.0)
        self.assertEqual(dc.derive_min_edge_override(exp, 0.05), 0.0)

    def test_thin_data_no_calibration(self):
        exp = dc.expectancy_short(closes(2, 10, 3, 20), SINCE)  # n=5 < N_MIN
        self.assertEqual(dc.derive_shrink(exp, 0, 0.0)[0], 1.0)
        self.assertEqual(dc.derive_min_edge_override(exp, 0.05), 0.0)

    def test_excludes_non_real_reasons(self):
        exp = dc.expectancy_short(closes(6, 10, 6, 20, reason="manual_reset_v7"), SINCE)
        self.assertEqual(exp["n"], 0)

    def test_excludes_non_short(self):
        exp = dc.expectancy_short(closes(6, 10, 6, 20, hz="long"), SINCE)
        self.assertEqual(exp["n"], 0)

    def test_brier_penalty_blends_when_ripe(self):
        rows = closes(6, 10, 6, 20)  # shrink_exp 0.75
        plain, _ = dc.derive_shrink(rows and dc.expectancy_short(rows, SINCE), 0, 0.0)
        worse, _ = dc.derive_shrink(dc.expectancy_short(rows, SINCE), 30, 0.40)  # bad Brier, n>=20
        self.assertLess(worse, plain)  # worse Brier shrinks further


class TestMemoryIdempotent(unittest.TestCase):
    def test_append_idempotent(self):
        mp = pathlib.Path(tempfile.mkdtemp()) / "memory.md"
        mp.write_text("# memory\n", encoding="utf-8")
        section = "\n## Calibracion diaria — 2026-05-29\n\nbody\n"
        self.assertTrue(dc.append_memory(mp, section, "2026-05-29"))
        self.assertFalse(dc.append_memory(mp, section, "2026-05-29"))
        self.assertEqual(mp.read_text(encoding="utf-8").count("Calibracion diaria — 2026-05-29"), 1)


if __name__ == "__main__":
    unittest.main()
