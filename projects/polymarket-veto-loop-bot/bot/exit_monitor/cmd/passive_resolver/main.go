// Command passive_resolver backfills outcomes in predictions.jsonl for
// markets that have closed upstream, regardless of whether the bot actually
// traded them.
//
// v7 P6: the brain in shadow mode emits one prediction per evaluated
// candidate (approve / veto / skip_shadow). Without a passive resolver,
// only positions the bot opened ever receive an outcome — leaving 99%+
// of predictions unscored and starving Brier calibration.
//
// This binary scans predictions.jsonl for unique MarketIDs that lack an
// outcome, fetches the gamma-api quote for each, and when the market is
// closed derives the YES-side outcome (0 or 1) and patches every matching
// row via predictions.BackfillOutcome.
//
// Cron: */60 (hourly). Idempotent — already-resolved rows are skipped by
// BackfillOutcome.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/predictions"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	limit := flag.Int("limit", 0, "max number of distinct markets to resolve this run (0=unlimited)")
	dryRun := flag.Bool("dry-run", false, "fetch quotes but do not patch predictions.jsonl")
	flag.Parse()

	path := os.Getenv("EXIT_PREDICTIONS_PATH")
	if path == "" {
		path = os.Getenv("BRAIN_PREDICTIONS_PATH")
	}
	if path == "" {
		path = predictions.DefaultPath
	}

	unresolved, err := collectUnresolved(path)
	if err != nil {
		log.Fatalf("passive_resolver: collect unresolved: %v", err)
	}
	log.Printf("passive_resolver: %d distinct unresolved markets in %s", len(unresolved), path)
	if len(unresolved) == 0 {
		return
	}

	now := time.Now().UTC().Format(time.RFC3339)
	closedCount, patchedCount, skippedCount, errCount := 0, 0, 0, 0
	processed := 0
	for marketID := range unresolved {
		if *limit > 0 && processed >= *limit {
			break
		}
		processed++
		q, err := polyclient.FetchQuote(marketID)
		if err != nil {
			errCount++
			log.Printf("passive_resolver: fetch %s failed: %v", marketID, err)
			continue
		}
		if !q.Closed {
			skippedCount++
			continue
		}
		outcome := outcomeYES(q)
		closedCount++
		if *dryRun {
			log.Printf("[dry-run] would patch %s outcome=%.0f (YES_price=%.3f last=%.3f)",
				marketID, outcome, q.OutcomePricesY, q.LastTradePrice)
			continue
		}
		patched, err := predictions.BackfillOutcome(path, marketID, outcome, "passive_resolver", now)
		if err != nil {
			errCount++
			log.Printf("passive_resolver: backfill %s failed: %v", marketID, err)
			continue
		}
		patchedCount += patched
		log.Printf("passive_resolver: %s closed → outcome=%.0f, patched %d row(s)",
			marketID, outcome, patched)
	}
	log.Printf("passive_resolver: done — processed=%d closed=%d patched_rows=%d still_open=%d errors=%d",
		processed, closedCount, patchedCount, skippedCount, errCount)
}

// collectUnresolved returns the set of MarketIDs that have at least one
// prediction row with Outcome == nil. MarketIDs already fully resolved are
// excluded.
func collectUnresolved(path string) (map[string]struct{}, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]struct{}{}, nil
		}
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	out := map[string]struct{}{}
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var p predictions.Prediction
		if err := json.Unmarshal(line, &p); err != nil {
			continue
		}
		if p.MarketID == "" || p.Outcome != nil {
			continue
		}
		out[p.MarketID] = struct{}{}
	}
	return out, sc.Err()
}

// outcomeYES derives the YES-side binary outcome from a closed quote.
// All predictions are stored as YES-probability regardless of decision side
// (buy_yes / buy_no / skip), so the canonical outcome is always the YES
// resolution: 1 if YES won, 0 if NO won.
//
// Resolution heuristic (mirrors monitor.outcomeFromQuote):
//  1. If outcomePrices is populated, use OutcomePricesY ≥ 0.5.
//  2. Else fall back to LastTradePrice ≥ 0.5.
//  3. Else default to 0 (conservative — unresolved-looking quote).
func outcomeYES(q polyclient.MarketQuote) float64 {
	if q.OutcomePricesY > 0 || q.OutcomePricesN > 0 {
		if q.OutcomePricesY >= 0.5 {
			return 1
		}
		return 0
	}
	if q.LastTradePrice > 0 {
		if q.LastTradePrice >= 0.5 {
			return 1
		}
		return 0
	}
	return 0
}
