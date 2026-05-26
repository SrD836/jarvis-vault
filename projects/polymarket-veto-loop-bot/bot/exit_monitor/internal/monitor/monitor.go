package monitor

import (
	"log"
	"math"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/closer"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/loglosses"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const memoryPath = "vault/agents/polymarket-bot/memory.md"

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("exit_monitor: load config: %v", err)
	}
	log.Printf("exit_monitor: config loaded stopLoss=%.0f takeProfitPct=%.0f%%",
		cfg.StopLossUSD, cfg.TakeProfitPct)

	actives := closer.ReadActive()
	if len(actives) == 0 {
		log.Println("exit_monitor: no active positions")
		return
	}

	var closed []types.ClosedTrade
	var remaining []types.ActiveTrade
	now := time.Now().UTC()

	for _, a := range actives {
		quote, err := polyclient.FetchQuote(a.MarketID)
		if err != nil {
			log.Printf("exit_monitor: quote fetch error for %s: %v", a.MarketID, err)
			remaining = append(remaining, a)
			continue
		}
		// v5: sell YES at bestBid (realistic execution).
		price, priceSrc := polyclient.PriceForExecution(quote, a.Side, "sell")
		if price <= 0 {
			log.Printf("exit_monitor: no sell price for %s (bestBid=%.4f bestAsk=%.4f)", a.MarketID, quote.BestBid, quote.BestAsk)
			remaining = append(remaining, a)
			continue
		}
		mktClosed := quote.Closed

		// v5: liquidity gating — if our position is >25% of visible depth, defer exit.
		// (Selling would push price below bestBid; better to wait for deeper book.)
		sellSize := a.Size
		if a.SizeUSD > 0 {
			sellSize = a.SizeUSD
		}
		if cfg.LiquidityMinRatio > 0 && quote.LiquidityUSD > 0 && quote.LiquidityUSD < sellSize*cfg.LiquidityMinRatio && !mktClosed {
			log.Printf("exit_monitor: skip exit %s — low liquidity $%.0f < $%.0f", a.MarketID, quote.LiquidityUSD, sellSize*cfg.LiquidityMinRatio)
			remaining = append(remaining, a)
			continue
		}
		if !quote.AcceptingOrders && !mktClosed {
			log.Printf("exit_monitor: skip exit %s — orders paused", a.MarketID)
			remaining = append(remaining, a)
			continue
		}

		entryPrice := a.EntryPrice
		pnlPct := (price - entryPrice) / entryPrice * 100
		pnlUSD := (price - entryPrice) / entryPrice * sellSize

		// v5: stop_loss as % too (configurable). Either absolute USD or relative % triggers.
		stopHitPct := cfg.StopLossPct > 0 && pnlPct <= -cfg.StopLossPct
		stopHitUSD := cfg.StopLossUSD > 0 && pnlUSD <= -cfg.StopLossUSD

		reason := ""
		switch {
		case mktClosed:
			reason = "market_closed"
		case pnlPct >= cfg.TakeProfitPct:
			reason = "take_profit"
		case stopHitUSD || stopHitPct:
			reason = "stop_loss"
		default:
			remaining = append(remaining, a)
			continue
		}

		pnlR := math.Round(pnlUSD*100) / 100
		exitTs := now.Format(time.RFC3339)
		daysOpen := computeDaysOpen(a.EntryTimestamp, now)
		early := false
		if !mktClosed && a.EndDate != "" {
			if t, err := time.Parse(time.RFC3339, a.EndDate); err == nil && now.Before(t) {
				early = true
			}
		}
		ct := types.ClosedTrade{
			ID:              a.ID,
			MarketID:        a.MarketID,
			Slug:            a.Slug,
			Question:        a.Question,
			Category:        a.Category,
			Side:            a.Side,
			Size:            sellSize,
			EntryPrice:      entryPrice,
			ExitPrice:       price,
			EntryTime:       a.EntryTimestamp,
			ExitTime:        exitTs,
			Pnl:             pnlR,
			PnlUSD:          pnlR,
			PnlPct:          math.Round(pnlPct*100) / 100,
			Reason:          reason,
			SourcesUsed:     a.SourcesUsed,
			DaysOpen:        daysOpen,
			EarlyExit:       early,
			EndDate:         a.EndDate,
			ExitPriceSource: priceSrc,
			LiquidityUSD:    quote.LiquidityUSD,
			Horizon:         a.Horizon,
		}
		closed = append(closed, ct)
		earlyTag := ""
		if early {
			earlyTag = " EARLY"
		}
		log.Printf("exit_monitor: closed %s (%s%s) @ %.4f src=%s liq=$%.0f: PnL $%.2f (%.2f%%)", a.Question, reason, earlyTag, price, priceSrc, quote.LiquidityUSD, ct.Pnl, ct.PnlPct)

		loglosses.LogClose(closer.DecisionsDir, memoryPath, ct)

		// v4: attribute outcome to every source that informed the decision.
		if err := closer.AppendSourceStats(ct); err != nil {
			log.Printf("exit_monitor: source-stats append failed: %v", err)
		}
		// v4: sync Obsidian decision .md (frontmatter outcome/pnl/closed_at).
		if err := closer.PatchDecisionMD(ct); err != nil {
			log.Printf("exit_monitor: decision .md patch failed for %s: %v", ct.Slug, err)
		}
	}

	if len(closed) > 0 {
		closer.AppendClosed(closed)
		closer.UpdatePortfolio(closed)
	}
	closer.RewriteActive(remaining)
	log.Printf("exit_monitor: %d closed, %d remaining", len(closed), len(remaining))
}

func computeDaysOpen(entryTS string, now time.Time) float64 {
	t, err := time.Parse(time.RFC3339, entryTS)
	if err != nil {
		return 0
	}
	return math.Round(now.Sub(t).Hours()/24*100) / 100
}
