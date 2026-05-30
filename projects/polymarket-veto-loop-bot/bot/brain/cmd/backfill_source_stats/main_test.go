package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/memory"
)

// A domain cited multiple times in one trade, and a trade processed once, must
// each count as exactly one trade for that domain. Empty domains are dropped.
func TestBuildEventsDedupsByTradeAndDomain(t *testing.T) {
	now := time.Now().UTC().Format(time.RFC3339)
	trades := []closedTrade{
		{ID: "T-1", Pnl: 10, ExitTime: now, SourcesUsed: []sourceCite{
			{Domain: "reuters.com"}, {Domain: "Reuters.com"}, {Domain: "apnews.com"},
		}},
		{ID: "T-2", Pnl: -5, ExitTime: now, SourcesUsed: []sourceCite{{Domain: " "}, {Domain: ""}}},
		{ID: "T-3", Pnl: -3, ExitTime: now, SourcesUsed: []sourceCite{{Domain: "claudemax-websearch", Synthetic: true}}},
	}
	got := buildEvents(trades)
	// reuters (x2 same trade -> 1, case-insensitive) + apnews + claudemax = 3
	if len(got) != 3 {
		t.Fatalf("want 3 deduped events, got %d: %+v", len(got), got)
	}
	count := map[string]int{}
	for _, e := range got {
		count[e.Domain]++
	}
	if count["reuters.com"] != 1 {
		t.Errorf("reuters.com events=%d want 1", count["reuters.com"])
	}
	if count["claudemax-websearch"] != 1 {
		t.Errorf("claudemax-websearch should be kept (events=%d)", count["claudemax-websearch"])
	}
}

// Full pipeline: 5 trades on one domain at 20% win-rate (>=5 trades, <30%) must
// be blacklisted; legacy is filtered at read; Human notes survive the rewrite.
func TestBackfillAggregatesAndBlacklists(t *testing.T) {
	dir := t.TempDir()
	statsPath := filepath.Join(dir, "source-stats.jsonl")
	memPath := filepath.Join(dir, "memory.md")
	seed := "# Mem\n\n## Sources stats (stale)\n\nold table\n\n## Human notes\n\nkeep me\n"
	if err := os.WriteFile(memPath, []byte(seed), 0644); err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC().Format(time.RFC3339)
	var trades []closedTrade
	for i := 0; i < 4; i++ {
		trades = append(trades, closedTrade{ID: fmt.Sprintf("L%d", i), Pnl: -10, ExitTime: now,
			SourcesUsed: []sourceCite{{Domain: "badsource.com"}}})
	}
	trades = append(trades, closedTrade{ID: "W0", Pnl: 5, ExitTime: now,
		SourcesUsed: []sourceCite{{Domain: "badsource.com"}}})
	trades = append(trades, closedTrade{ID: "X", Pnl: 1, ExitTime: now,
		SourcesUsed: []sourceCite{{Domain: "legacy"}}})

	events := buildEvents(trades)
	if err := writeEvents(statsPath, events); err != nil {
		t.Fatal(err)
	}
	stats, err := memory.LoadSourceStats(statsPath, 30*24*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := stats["legacy"]; ok {
		t.Error("legacy must be filtered by LoadSourceStats")
	}
	bs := stats["badsource.com"]
	if bs == nil || bs.NTrades != 5 {
		t.Fatalf("badsource.com NTrades want 5, got %+v", bs)
	}
	bl := memory.Blacklisted(stats, 5, 0.30)
	found := false
	for _, d := range bl {
		if d == "badsource.com" {
			found = true
		}
	}
	if !found {
		t.Errorf("badsource.com should be blacklisted (winrate %.0f%%), got %v", bs.WinRate*100, bl)
	}
	if err := memory.WriteSourceStatsSection(memPath, stats, bl); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(memPath)
	if !strings.Contains(string(out), "badsource.com") {
		t.Errorf("memory.md missing badsource.com row:\n%s", out)
	}
	if !strings.Contains(string(out), "keep me") {
		t.Error("WriteSourceStatsSection clobbered Human notes")
	}
}
