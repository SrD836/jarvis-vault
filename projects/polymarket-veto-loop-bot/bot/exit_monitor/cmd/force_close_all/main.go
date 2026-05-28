// One-shot: close every position in active.jsonl regardless of horizon / PnL.
//
// v7 safety halt: before flipping the bot into shadow mode we purge all open
// positions. Defaults to --dry-run; pass --apply to mutate. Always writes
// .bak-YYYY-MM-DD next to active.jsonl/portfolio.json before mutating.
//
// Differs from force_close_horizon_excess (deprecated, do NOT cron) which
// selected the "worst" N positions per horizon bucket — that command violates
// the v7 thesis-exit principle and is kept only as a manual utility.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/closer"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/loglosses"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

type scored struct {
	a        types.ActiveTrade
	price    float64
	pnlUSD   float64
	pnlPct   float64
	quote    polyclient.MarketQuote
	priceSrc string
}

func main() {
	dryRun := flag.Bool("dry-run", true, "print the plan without mutating (default)")
	apply := flag.Bool("apply", false, "execute the plan; required together with --dry-run=false")
	reason := flag.String("reason", "manual_reset_v7", "exit_reason written to closed.jsonl")
	flag.Parse()
	if *apply {
		*dryRun = false
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	actives := closer.ReadActive()
	if len(actives) == 0 {
		log.Println("force_close_all: no active positions, nothing to do")
		return
	}

	var rows []scored
	for _, a := range actives {
		quote, err := polyclient.FetchQuote(a.MarketID)
		if err != nil {
			log.Printf("force_close_all: quote error for %s (%s) — skipping: %v", a.MarketID, a.Slug, err)
			continue
		}
		price, src := polyclient.PriceForExecution(quote, a.Side, "sell")
		if price <= 0 {
			price = quote.LastTradePrice
			src = "last_trade_fallback"
		}
		size := a.Size
		if a.SizeUSD > 0 {
			size = a.SizeUSD
		}
		pnlUSD := (price - a.EntryPrice) / a.EntryPrice * size
		pnlPct := (price - a.EntryPrice) / a.EntryPrice * 100
		rows = append(rows, scored{a, price, pnlUSD, pnlPct, quote, src})
	}

	fmt.Println("\nplan: close ALL positions")
	fmt.Println("slug | horizon | entry → exit | pnl_usd | pnl_pct | price_src")
	total := 0.0
	for _, s := range rows {
		fmt.Printf("%-60s | %-6s | %.4f → %.4f | %+7.2f | %+6.2f%% | %s\n",
			trunc(s.a.Slug, 60), s.a.Horizon, s.a.EntryPrice, s.price, s.pnlUSD, s.pnlPct, s.priceSrc)
		total += s.pnlUSD
	}
	fmt.Printf("\nwould close %d positions (sum unrealized PnL: $%.2f)\n", len(rows), total)

	if *dryRun {
		log.Println("force_close_all: DRY-RUN — pass --apply to execute")
		return
	}

	date := time.Now().UTC().Format("2006-01-02")
	for _, name := range []string{"active.jsonl", "portfolio.json"} {
		src := filepath.Join(closer.DataDir, name)
		dst := src + ".bak-" + date
		raw, err := os.ReadFile(src)
		if err != nil {
			log.Fatalf("force_close_all: read %s for backup: %v", src, err)
		}
		if err := os.WriteFile(dst, raw, 0644); err != nil {
			log.Fatalf("force_close_all: write backup %s: %v", dst, err)
		}
		log.Printf("force_close_all: backup → %s", dst)
	}

	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)
	var closedTrades []types.ClosedTrade
	for _, s := range rows {
		size := s.a.Size
		if s.a.SizeUSD > 0 {
			size = s.a.SizeUSD
		}
		daysOpen := computeDaysOpen(s.a.EntryTimestamp, now)
		ct := types.ClosedTrade{
			ID:              s.a.ID,
			MarketID:        s.a.MarketID,
			Slug:            s.a.Slug,
			Question:        s.a.Question,
			Category:        s.a.Category,
			Side:            s.a.Side,
			Size:            size,
			EntryPrice:      s.a.EntryPrice,
			ExitPrice:       s.price,
			EntryTime:       s.a.EntryTimestamp,
			ExitTime:        nowStr,
			Pnl:             math.Round(s.pnlUSD*100) / 100,
			PnlUSD:          math.Round(s.pnlUSD*100) / 100,
			PnlPct:          math.Round(s.pnlPct*100) / 100,
			Reason:          *reason,
			SourcesUsed:     s.a.SourcesUsed,
			DaysOpen:        daysOpen,
			EndDate:         s.a.EndDate,
			ExitPriceSource: s.priceSrc,
			LiquidityUSD:    s.quote.LiquidityUSD,
			Horizon:         s.a.Horizon,
		}
		closedTrades = append(closedTrades, ct)
		if err := closer.AppendSourceStats(ct); err != nil {
			log.Printf("force_close_all: source-stats append failed: %v", err)
		}
		if err := closer.PatchDecisionMD(ct); err != nil {
			log.Printf("force_close_all: decision .md patch failed for %s: %v", ct.Slug, err)
		}
	}

	closer.AppendClosed(closedTrades)
	closer.UpdatePortfolio(closedTrades)
	closer.RewriteActive(nil)

	memoryPath := os.Getenv("EXIT_MEMORY_PATH")
	if memoryPath == "" {
		memoryPath = "vault/agents/polymarket-bot/memory.md"
	}
	for _, ct := range closedTrades {
		loglosses.LogClose(closer.DecisionsDir, memoryPath, ct)
	}
	log.Printf("force_close_all: closed %d positions, 0 remain (reason=%s, logged to %s)",
		len(closedTrades), *reason, memoryPath)
}

func computeDaysOpen(entryTS string, now time.Time) float64 {
	t, err := time.Parse(time.RFC3339, entryTS)
	if err != nil {
		return 0
	}
	return math.Round(now.Sub(t).Hours()/24*100) / 100
}

func trunc(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
