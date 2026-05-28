// DEPRECATED (v7, 2026-05-28): closes the K worst positions per horizon
// bucket by mark-to-market PnL. This violates the v7 thesis-exit rule
// ("sales cuando la tesis se invalida, no porque toque rebalancear" — see
// prompt maestro principle #4 and projects/polymarket-veto-loop-bot/design.md).
//
// DO NOT run on cron. Kept as a manual utility only; new callers should use
// bot/exit_monitor/cmd/force_close_all/ (closes everything, no PnL ranking)
// or rely on monitor.go's thesis-based exits.
//
// Original docstring:
//   Reads active.jsonl + portfolio.json, fetches a live quote per position,
//   computes unrealized PnL, and closes the K worst positions in each over-
//   quota bucket so the remaining counts respect ceil(MaxOpenPositions × quota).
//   Defaults to --dry-run; pass --apply to actually mutate.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
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
	flag.Parse()
	if *apply {
		*dryRun = false
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("force_close: load config: %v", err)
	}

	actives := closer.ReadActive()
	if len(actives) == 0 {
		log.Println("force_close: no active positions, nothing to do")
		return
	}

	// Caps from the same formula as executor Q2.
	caps := map[string]int{
		"short":  ceilQuota(cfg.HorizonQuotaShort, cfg.MaxOpenPositions),
		"medium": ceilQuota(cfg.HorizonQuotaMedium, cfg.MaxOpenPositions),
		"long":   ceilQuota(cfg.HorizonQuotaLong, cfg.MaxOpenPositions),
	}
	log.Printf("force_close: caps short=%d medium=%d long=%d (max_open=%d)",
		caps["short"], caps["medium"], caps["long"], cfg.MaxOpenPositions)

	byHorizon := map[string][]scored{}
	for _, a := range actives {
		quote, err := polyclient.FetchQuote(a.MarketID)
		if err != nil {
			log.Printf("force_close: quote error for %s (%s) — skipping: %v", a.MarketID, a.Slug, err)
			continue
		}
		price, src := polyclient.PriceForExecution(quote, a.Side, "sell")
		if price <= 0 {
			// Use last trade as a soft fallback for PnL ranking only (not for close).
			price = quote.LastTradePrice
			src = "last_trade_fallback"
		}
		size := a.Size
		if a.SizeUSD > 0 {
			size = a.SizeUSD
		}
		pnlUSD := (price - a.EntryPrice) / a.EntryPrice * size
		pnlPct := (price - a.EntryPrice) / a.EntryPrice * 100
		h := a.Horizon
		if h == "" {
			h = "long"
		}
		byHorizon[h] = append(byHorizon[h], scored{a, price, pnlUSD, pnlPct, quote, src})
	}

	var toClose []scored
	keepIDs := map[string]bool{}
	for h, list := range byHorizon {
		sort.Slice(list, func(i, j int) bool { return list[i].pnlUSD < list[j].pnlUSD })
		capN := caps[h]
		if len(list) <= capN {
			for _, s := range list {
				keepIDs[s.a.ID] = true
			}
			continue
		}
		excess := len(list) - capN
		log.Printf("force_close: horizon %s — %d open, cap %d → closing %d worst", h, len(list), capN, excess)
		toClose = append(toClose, list[:excess]...)
		for _, s := range list[excess:] {
			keepIDs[s.a.ID] = true
		}
	}
	// Skipped positions (quote errors) are also kept.
	for _, a := range actives {
		found := false
		for _, list := range byHorizon {
			for _, s := range list {
				if s.a.ID == a.ID {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			keepIDs[a.ID] = true
		}
	}

	if len(toClose) == 0 {
		log.Println("force_close: no excess positions; quota already satisfied")
		return
	}

	fmt.Println("\nplan:")
	fmt.Println("slug | horizon | entry → exit | pnl_usd | pnl_pct | price_src")
	for _, s := range toClose {
		fmt.Printf("%-60s | %-6s | %.4f → %.4f | %+7.2f | %+6.2f%% | %s\n",
			trunc(s.a.Slug, 60), s.a.Horizon, s.a.EntryPrice, s.price, s.pnlUSD, s.pnlPct, s.priceSrc)
	}
	fmt.Printf("\nwould close %d positions (sum unrealized PnL: $%.2f)\n",
		len(toClose), sumPnL(toClose))

	if *dryRun {
		log.Println("force_close: DRY-RUN — pass --apply to execute")
		return
	}

	// Backup before mutating.
	date := time.Now().UTC().Format("2006-01-02")
	for _, name := range []string{"active.jsonl", "portfolio.json"} {
		src := filepath.Join(closer.DataDir, name)
		dst := src + ".bak-" + date
		raw, err := os.ReadFile(src)
		if err != nil {
			log.Fatalf("force_close: read %s for backup: %v", src, err)
		}
		if err := os.WriteFile(dst, raw, 0644); err != nil {
			log.Fatalf("force_close: write backup %s: %v", dst, err)
		}
		log.Printf("force_close: backup → %s", dst)
	}

	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)
	var closedTrades []types.ClosedTrade
	for _, s := range toClose {
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
			Reason:          "horizon-quota-rebalance-" + date,
			SourcesUsed:     s.a.SourcesUsed,
			DaysOpen:        daysOpen,
			EndDate:         s.a.EndDate,
			ExitPriceSource: s.priceSrc,
			LiquidityUSD:    s.quote.LiquidityUSD,
			Horizon:         s.a.Horizon,
		}
		closedTrades = append(closedTrades, ct)
		if err := closer.AppendSourceStats(ct); err != nil {
			log.Printf("force_close: source-stats append failed: %v", err)
		}
		if err := closer.PatchDecisionMD(ct); err != nil {
			log.Printf("force_close: decision .md patch failed for %s: %v", ct.Slug, err)
		}
	}

	// Rewrite active.jsonl with kept positions.
	var remaining []types.ActiveTrade
	for _, a := range actives {
		if keepIDs[a.ID] {
			remaining = append(remaining, a)
		}
	}
	closer.AppendClosed(closedTrades)
	closer.UpdatePortfolio(closedTrades)
	closer.RewriteActive(remaining)
	memoryPath := os.Getenv("EXIT_MEMORY_PATH")
	if memoryPath == "" {
		memoryPath = "vault/agents/polymarket-bot/memory.md"
	}
	for _, ct := range closedTrades {
		loglosses.LogClose(closer.DecisionsDir, memoryPath, ct)
	}
	log.Printf("force_close: closed %d, %d remain (logged to %s)", len(closedTrades), len(remaining), memoryPath)
}

func ceilQuota(quota float64, total int) int {
	if total <= 0 || quota <= 0 {
		return 0
	}
	raw := quota * float64(total)
	rounded := int(raw)
	if raw-float64(rounded) > 0 {
		rounded++
	}
	return rounded
}

func computeDaysOpen(entryTS string, now time.Time) float64 {
	t, err := time.Parse(time.RFC3339, entryTS)
	if err != nil {
		return 0
	}
	return math.Round(now.Sub(t).Hours()/24*100) / 100
}

func sumPnL(s []scored) float64 {
	total := 0.0
	for _, x := range s {
		total += x.pnlUSD
	}
	return total
}

func trunc(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
