// Command backfill_source_stats rebuilds source-stats.jsonl from the canonical
// closed.jsonl, then recomputes the "## Sources stats" section of memory.md.
//
// Why: closer.AppendSourceStats writes one row per source CITATION at close time,
// so a domain cited several times in one trade (or a trade processed by both the
// exit monitor and a force_close one-shot) inflates the per-domain trade count.
// This one-shot reprocesses every closed trade and emits exactly one row per
// (trade_id, domain) pair, so SourceStats.NTrades counts trades, not citations.
//
// Unlike closer.AppendSourceStats, the backfill keeps synthetic/claudemax-websearch
// cites: they are the only attribution the historical trades carry. The display
// filter (legacy/empty dropped) is applied by memory.LoadSourceStats at read time.
//
// It never mutates closed.jsonl. Run from the agent-stack working directory:
//   ./backfill_source_stats
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/memory"
)

// closedTrade is the minimal projection of exit_monitor's types.ClosedTrade
// needed for source attribution. Defined locally to avoid a cross-module import
// (closed.jsonl is written by the exit_monitor module; this command lives in the
// brain module). Field tags mirror exit_monitor/internal/types.ClosedTrade.
type closedTrade struct {
	ID          string       `json:"id"`
	Pnl         float64      `json:"pnl"`
	DaysOpen    float64      `json:"days_open"`
	ExitTime    string       `json:"exit_timestamp"`
	SourcesUsed []sourceCite `json:"sources_used"`
}

type sourceCite struct {
	Domain    string `json:"domain"`
	Synthetic bool   `json:"synthetic"`
}

func outcomeOf(pnl float64) string {
	switch {
	case pnl > 0:
		return "win"
	case pnl < 0:
		return "loss"
	default:
		return "breakeven"
	}
}

// buildEvents turns closed trades into deduped source events: one row per
// (trade_id, domain). Empty domains are skipped; everything else is kept here so
// LoadSourceStats can apply the display filter consistently.
func buildEvents(trades []closedTrade) []memory.SourceEvent {
	seen := map[string]bool{}
	var out []memory.SourceEvent
	for _, t := range trades {
		for _, s := range t.SourcesUsed {
			dom := strings.ToLower(strings.TrimSpace(s.Domain))
			if dom == "" {
				continue
			}
			key := t.ID + "|" + dom
			if seen[key] {
				continue
			}
			seen[key] = true
			out = append(out, memory.SourceEvent{
				TS:       t.ExitTime,
				TradeID:  t.ID,
				Domain:   dom,
				Outcome:  outcomeOf(t.Pnl),
				PnlUSD:   t.Pnl,
				DaysOpen: t.DaysOpen,
			})
		}
	}
	return out
}

func readClosed(path string) ([]closedTrade, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []closedTrade
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(strings.TrimSpace(string(line))) == 0 {
			continue
		}
		var ct closedTrade
		if err := json.Unmarshal(line, &ct); err != nil {
			continue // tolerate malformed rows, same as the bot's scanners
		}
		out = append(out, ct)
	}
	return out, sc.Err()
}

// writeEvents rewrites the stats file atomically (tmp + rename).
func writeEvents(path string, events []memory.SourceEvent) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	for _, ev := range events {
		if err := enc.Encode(ev); err != nil {
			f.Close()
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func main() {
	var (
		closedPath = flag.String("closed", "vault/inbox/trading/closed.jsonl", "path to closed.jsonl (read-only)")
		statsPath  = flag.String("stats", "vault/agents/polymarket-bot/source-stats.jsonl", "path to source-stats.jsonl (rewritten)")
		memPath    = flag.String("memory", "vault/agents/polymarket-bot/memory.md", "path to memory.md (## Sources stats rewritten)")
		windowDays = flag.Int("window-days", 30, "rolling window (days) for stats aggregation")
		minTrades  = flag.Int("min-trades", 5, "blacklist gate: minimum trades")
		minWinRate = flag.Float64("min-winrate", 0.30, "blacklist gate: win-rate floor")
		dryRun     = flag.Bool("dry-run", false, "compute and print without writing")
	)
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	trades, err := readClosed(*closedPath)
	if err != nil {
		log.Fatalf("read closed: %v", err)
	}
	events := buildEvents(trades)
	log.Printf("closed trades=%d -> deduped source events=%d", len(trades), len(events))

	if !*dryRun {
		if err := writeEvents(*statsPath, events); err != nil {
			log.Fatalf("write stats: %v", err)
		}
		log.Printf("wrote %d events to %s", len(events), *statsPath)
	}

	window := time.Duration(*windowDays) * 24 * time.Hour
	stats, err := memory.LoadSourceStats(*statsPath, window)
	if err != nil {
		log.Fatalf("load source stats: %v", err)
	}
	bl := memory.Blacklisted(stats, *minTrades, *minWinRate)
	log.Printf("domains after display filter=%d  blacklisted=%d  %v", len(stats), len(bl), bl)
	for _, st := range memory.SortedByVolume(stats) {
		log.Printf("  %-26s trades=%d wins=%d losses=%d winrate=%.1f%% pnl=%+.2f",
			st.Domain, st.NTrades, st.NWins, st.NLosses, st.WinRate*100, st.TotalPnl)
	}

	if !*dryRun {
		if err := memory.WriteSourceStatsSection(*memPath, stats, bl); err != nil {
			log.Fatalf("write memory section: %v", err)
		}
		log.Printf("rewrote ## Sources stats in %s", *memPath)
	}
}
