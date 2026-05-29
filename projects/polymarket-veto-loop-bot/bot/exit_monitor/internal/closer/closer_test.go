package closer

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

// A domain cited several times within one closed trade must yield exactly one
// source-stats row; legacy/claudemax-websearch/synthetic stay filtered.
func TestAppendSourceStatsDedupsDomainWithinTrade(t *testing.T) {
	dir := t.TempDir()
	orig := SourceStatsPath
	SourceStatsPath = filepath.Join(dir, "source-stats.jsonl")
	defer func() { SourceStatsPath = orig }()

	ct := types.ClosedTrade{
		ID:       "T-1",
		Pnl:      -10,
		ExitTime: "2026-05-29T00:00:00Z",
		SourcesUsed: []commontypes.SourceCite{
			{Domain: "reuters.com"},
			{Domain: "reuters.com"}, // duplicate within the same trade
			{Domain: "apnews.com"},
			{Domain: "legacy"},              // filtered
			{Domain: "claudemax-websearch"}, // filtered
			{Domain: "x.com", Synthetic: true}, // filtered
		},
	}
	if err := AppendSourceStats(ct); err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(SourceStatsPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	rows := 0
	domains := map[string]int{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.TrimSpace(sc.Text()) == "" {
			continue
		}
		rows++
		for _, d := range []string{"reuters.com", "apnews.com", "legacy", "claudemax-websearch"} {
			if strings.Contains(sc.Text(), `"domain":"`+d+`"`) {
				domains[d]++
			}
		}
	}
	if rows != 2 {
		t.Fatalf("want 2 rows (reuters once + apnews once), got %d", rows)
	}
	if domains["reuters.com"] != 1 {
		t.Errorf("reuters.com rows=%d want 1 (deduped)", domains["reuters.com"])
	}
	if domains["legacy"] != 0 || domains["claudemax-websearch"] != 0 {
		t.Errorf("legacy/claudemax must stay filtered: %v", domains)
	}
}
