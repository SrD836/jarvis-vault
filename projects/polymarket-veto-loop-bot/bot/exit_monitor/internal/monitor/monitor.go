package monitor

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/predictions"
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
	log.Printf("exit_monitor v7: mode=%s thesis-exits (no SL/TP). MinEdgePoints=%.3f", cfg.Mode, cfg.MinEdgePoints)

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
		// Fix 2026-05-30: precio de cierre VALIDADO. Un cierre nunca se ejecuta a un
		// fallback no fiable (lastTradePrice de un libro muerto). ExitPrice distingue
		// "mercado resolvió" (settlement real 0/1) de "no pude obtener precio" (diferir).
		ref := a.EntryPrice
		if a.ApprovedPrice > 0 {
			ref = a.ApprovedPrice
		}
		dec := polyclient.ExitPrice(quote, a.Side, ref)
		if !dec.OK {
			if dec.Source == "blocked_sanity_band" {
				log.Printf("exit_monitor: BLOCKED close %s — precio %.4f vs ref %.4f (banda de sanidad), marcado para revision", a.MarketID, dec.Price, ref)
				enqueuePriceReview(a, dec, ref, now)
			} else {
				log.Printf("exit_monitor: sin precio fiable para %s (src=%s bestBid=%.4f bestAsk=%.4f) — se mantiene abierta", a.MarketID, dec.Source, quote.BestBid, quote.BestAsk)
			}
			remaining = append(remaining, a)
			continue
		}
		price, priceSrc := dec.Price, dec.Source
		mktClosed := dec.Resolved

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

		// v7: thesis-based exits only. SL/TP percent triggers are gone —
		// see prompt maestro principle #4 ("nunca cierres por % stop").
		//   1. market_closed   — Polymarket resolved the market.
		//   2. no_remaining_edge — under 24h left and current price already
		//      within 2pp of our estimated probability (no informational edge
		//      remains, capital better redeployed).
		//   3. target_hit      — current price reached the LLM's TargetProb.
		daysLeft := daysToResolution(a.EndDate, now)
		targetHit := false
		if a.TargetProb > 0 {
			switch a.Side {
			case "yes":
				targetHit = price >= a.TargetProb
			case "no":
				targetHit = price <= a.TargetProb
			}
		}
		noEdge := a.EstimatedProb > 0 && daysLeft >= 0 && daysLeft < 1 && math.Abs(price-a.EstimatedProb) < 0.02

		// G1 thesis_stale (Fase B, NUEVA 2026-05-30, reversible vía thesis_stale_enabled):
		// contraparte simétrica de target_hit. Cerca de resolución, si la tesis no se
		// materializó (precio no subió hacia EstimatedProb y sigue <= entrada) => cerrar.
		// Respeta "no % stop": solo cerca de expiry, nunca con tiempo por delante (sin whipsaw).
		// price ya está VALIDADO por ExitPrice (las no-fiables se difieren arriba).
		thesisStale := cfg.ThesisStaleEnabled && IsThesisStale(a.Side, price, entryPrice,
			a.EstimatedProb, daysLeft, cfg.ThesisStaleDays, cfg.ThesisStaleMargin)

		reason := ""
		switch {
		case mktClosed:
			reason = "market_closed"
		case noEdge:
			reason = "no_remaining_edge"
		case targetHit:
			reason = "target_hit"
		case thesisStale:
			reason = "thesis_stale"
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

		// v7: backfill calibration outcome only on market_closed (ground truth).
		// For target_hit / no_remaining_edge the ground truth is the eventual
		// resolution, not our early exit — those rows stay open in
		// predictions.jsonl until a later pass closes the underlying market.
		if reason == "market_closed" {
			outcome := polyclient.Settlement(quote, a.Side)
			predPath := os.Getenv("EXIT_PREDICTIONS_PATH")
			if predPath == "" {
				predPath = predictions.DefaultPath
			}
			if patched, err := predictions.BackfillOutcome(predPath, a.MarketID, outcome, reason, exitTs); err != nil {
				log.Printf("exit_monitor: backfill outcome %s: %v", a.MarketID, err)
			} else if patched > 0 {
				log.Printf("exit_monitor: backfilled %d prediction(s) for %s (outcome=%.0f)", patched, a.MarketID, outcome)
			}
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

// IsThesisStale (G1): la tesis no se materializó cerca de la resolución. Contraparte
// simétrica de target_hit. Solo dispara con 0 <= daysLeft < staleDays (ventana de expiry)
// y est > 0; nunca con tiempo por delante => respeta "no % stop" (sin whipsaw). Para YES:
// el precio sigue <= entrada (no progresó) y < est-margin (no convergió al alza). Para NO,
// simétrico (precio >= entrada y > est+margin). margin evita cortar por ruido pegado al estimado.
func IsThesisStale(side string, price, entry, est, daysLeft, staleDays, margin float64) bool {
	if est <= 0 || daysLeft < 0 || daysLeft >= staleDays {
		return false
	}
	switch side {
	case "yes":
		return price <= entry && price < est-margin
	case "no":
		return price >= entry && price > est+margin
	}
	return false
}

// daysToResolution returns days remaining until endDate. -1 if unparseable.
func daysToResolution(endDate string, now time.Time) float64 {
	if endDate == "" {
		return -1
	}
	t, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05Z", endDate)
		if err != nil {
			return -1
		}
	}
	return t.Sub(now).Hours() / 24
}

// priceReviewPath: cierres bloqueados por la banda de sanidad se encolan aquí
// para revisión humana (superficie Telegram/dashboard la consume aparte).
const priceReviewPath = "vault/inbox/trading/price_review.jsonl"

// enqueuePriceReview persiste un cierre bloqueado (precio sospechoso, mercado no
// resuelto). Append JSON; nunca aborta el monitor si falla la escritura.
func enqueuePriceReview(a types.ActiveTrade, dec polyclient.ExitDecision, ref float64, now time.Time) {
	rec := map[string]interface{}{
		"id":              a.ID,
		"market_id":       a.MarketID,
		"slug":            a.Slug,
		"entry_price":     a.EntryPrice,
		"candidate_price": dec.Price,
		"ref":             ref,
		"ts":              now.Format(time.RFC3339),
		"source":          dec.Source,
	}
	b, err := json.Marshal(rec)
	if err != nil {
		return
	}
	f, err := os.OpenFile(priceReviewPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("exit_monitor: price-review append failed: %v", err)
		return
	}
	defer f.Close()
	if _, err := f.Write(append(b, '\n')); err != nil {
		log.Printf("exit_monitor: price-review write failed: %v", err)
	}
}
