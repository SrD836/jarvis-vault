package loglosses

import (
	"path/filepath"
	"testing"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

func countLossNodes(t *testing.T, dir string) int {
	t.Helper()
	m, _ := filepath.Glob(filepath.Join(dir, "*-polymarket-loss-*.md"))
	return len(m)
}

// DASHBOARD_URL points at a closed port so the best-effort post-mortem fails
// instantly (and is swallowed) — the test only asserts node-writing behaviour.

func TestWriteDecisionLogGatedOff(t *testing.T) {
	t.Setenv("DASHBOARD_URL", "http://127.0.0.1:9")
	t.Setenv("DECISION_NODES", "")
	dir := t.TempDir()
	writeDecisionLog(dir, types.ClosedTrade{Slug: "gated-off", Question: "Q", Pnl: -10})
	if n := countLossNodes(t, dir); n != 0 {
		t.Fatalf("DECISION_NODES off: expected 0 loss nodes, got %d", n)
	}
}

func TestWriteDecisionLogGatedOn(t *testing.T) {
	t.Setenv("DASHBOARD_URL", "http://127.0.0.1:9")
	t.Setenv("DECISION_NODES", "1")
	dir := t.TempDir()
	writeDecisionLog(dir, types.ClosedTrade{Slug: "gated-on", Question: "Q", Pnl: -10})
	if n := countLossNodes(t, dir); n != 1 {
		t.Fatalf("DECISION_NODES on: expected 1 loss node, got %d", n)
	}
}
