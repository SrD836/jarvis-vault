package executor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/calibration"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/decisionlog"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/predictions"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/executor/internal/types"
)

const (
	dataDir       = "vault/inbox/trading"
	approvedFile  = "approved.jsonl"
	activeFile    = "active.jsonl"
	rejectFile    = "rejections.jsonl"
	portfolioFile = "portfolio.json"
)

func envOr(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("executor: load config: %v", err)
	}
	switch cfg.Mode {
	case "shadow", "simulation":
		// allowed
	case "live":
		panic("live trading not authorized — only shadow/simulation are supported in v7")
	default:
		panic(fmt.Sprintf("unknown mode %q — expected shadow|simulation|live", cfg.Mode))
	}
	log.Printf("executor v7: mode=%s sizingMode=%s kellyFrac=%.2f maxPerTrade=%.2f%% maxExposure=%.2f%% minEdge=%.3f maxOpen=%d maxSameCat=%d maxPerDay=%d quota=%.0f/%.0f/%.0f",
		cfg.Mode, cfg.SizingMode, cfg.KellyFraction, cfg.MaxPerTradePct*100, cfg.MaxTotalExposurePct*100,
		cfg.MinEdgePoints, cfg.MaxOpenPositions, cfg.MaxSameCategory, cfg.MaxNewTradesPerDay,
		cfg.HorizonQuotaShort*100, cfg.HorizonQuotaMedium*100, cfg.HorizonQuotaLong*100)

	// v8: Brier/expectancy calibration shrink (analytics/daily_calibration.py).
	// Absent file → shrink 1.0 (identical to raw Kelly). Applied to the model
	// probability before sizing so overconfident edges size down. Side and exit
	// target keep the raw estimate (shrink never crosses the implied price).
	calib := calibration.Load()
	if s := calib.Shrink(); s < 1.0 {
		log.Printf("executor v8: calibration kelly_shrink=%.3f (%s)", s, calib.Rationale)
	}

	portfolio := loadOrCreatePortfolio(cfg)
	approved := readApproved()
	if len(approved) == 0 {
		log.Println("executor: no approved candidates")
		return
	}

	// v7 shadow mode: log every approved candidate as a prediction with
	// decision="skip_shadow" and exit without mutating portfolio/active.jsonl.
	// This is the calibration warm-up before re-enabling real fills.
	if cfg.Mode == "shadow" {
		nowStr := time.Now().UTC().Format(time.RFC3339)
		predPath := envOr("EXECUTOR_PREDICTIONS_PATH", predictions.DefaultPath)
		for _, a := range approved {
			tradeSize := cfg.ComputeTradeSize(portfolio.Bankroll, calib.ShrinkProb(a.EstimatedProb, a.CurrentPriceYes), a.CurrentPriceYes)
			side := "buy_yes"
			if a.EstimatedProb < a.CurrentPriceYes {
				side = "buy_no"
			}
			p := predictions.Prediction{
				Timestamp:          nowStr,
				MarketID:           a.CandidateID,
				Slug:               a.Slug,
				Question:           a.Question,
				Category:           predictions.NormalizeCategory(a.Category),
				HorizonDays:        a.DaysToResolution,
				Horizon:            a.Horizon,
				ImpliedPrice:       a.CurrentPriceYes,
				EstimatedProb:      a.EstimatedProb,
				EdgePoints:         math.Abs(a.EstimatedProb - a.CurrentPriceYes),
				EdgeType:           a.EdgeType,
				EdgeDescription:    a.EdgeDescription,
				ChecklistPassed:    true,
				Decision:           "skip_shadow",
				SkipReason:         "mode=shadow (calibration warm-up)",
				SizeUSD:            tradeSize,
				ThesisInvalidation: a.ThesisInvalidation,
			}
			if err := predictions.Append(predPath, p); err != nil {
				log.Printf("executor shadow: predictions append failed for %s: %v", a.Slug, err)
			}
			log.Printf("executor shadow: would have opened [%s] %s @ %.4f size=$%.2f (side=%s edge=%.3f)",
				a.Horizon, a.Question, a.CurrentPriceYes, tradeSize, side, math.Abs(a.EstimatedProb-a.CurrentPriceYes))
		}
		log.Printf("executor shadow: %d candidates logged to predictions.jsonl, 0 trades opened", len(approved))
		return
	}

	todayPrefix := time.Now().UTC().Format("2006-01-02")
	todayCount := 0
	for _, p := range portfolio.Positions {
		if strings.HasPrefix(p.EntryTimestamp, todayPrefix) {
			todayCount++
		}
	}

	// v7: horizon quotas are USD caps over MaxExposureUSD, not trade counts.
	// Reason (post-mortem 2026-05-27): counting trades makes a $25 trade equal
	// to a $500 trade, which let the portfolio drift away from the intended
	// capital allocation. Caps are computed once per cycle and consumed
	// incrementally as trades are accepted below.
	usdCapShort := cfg.HorizonQuotaShort * cfg.MaxExposureUSD
	usdCapMedium := cfg.HorizonQuotaMedium * cfg.MaxExposureUSD
	usdCapLong := cfg.HorizonQuotaLong * cfg.MaxExposureUSD
	usdByHorizon := map[string]float64{}
	for _, p := range portfolio.Positions {
		sz := p.Size
		if p.SizeUSD > 0 {
			sz = p.SizeUSD
		}
		usdByHorizon[p.Horizon] += sz
	}
	log.Printf("executor v7: USD horizon caps short=$%.0f medium=$%.0f long=$%.0f (used now short=$%.0f medium=$%.0f long=$%.0f)",
		usdCapShort, usdCapMedium, usdCapLong,
		usdByHorizon["short"], usdByHorizon["medium"], usdByHorizon["long"])
	maxTotalExposureUSD := portfolio.Bankroll * cfg.MaxTotalExposurePct
	log.Printf("executor v7: total exposure cap $%.0f (bankroll=$%.0f × %.0f%%)", maxTotalExposureUSD, portfolio.Bankroll, cfg.MaxTotalExposurePct*100)

	decisionsDir := envOr("EXECUTOR_DECISIONS_DIR", "vault/03-decisions")

	var rejects []types.Rejection
	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)

	activeIDs := map[string]bool{}
	for _, p := range portfolio.Positions {
		activeIDs[p.MarketID] = true
	}
	openCount := len(portfolio.Positions)
	for _, a := range approved {
		if openCount >= cfg.MaxOpenPositions {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: fmt.Sprintf("max_open_positions_cap_reached_%d", cfg.MaxOpenPositions),
			})
			continue
		}
		if activeIDs[a.CandidateID] {
			rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: "already_active"})
			continue
		}
		if todayCount >= cfg.MaxNewTradesPerDay {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: "max_new_trades_per_day_reached",
			})
			continue
		}

		// v7 sizing: Kelly fractional needs the LLM's estimated probability and
		// the current price. If the LLM did not declare a real edge or Kelly
		// returns 0 (no edge after fees), we reject without consuming quota.
		tradeSize := cfg.ComputeTradeSize(portfolio.Bankroll, calib.ShrinkProb(a.EstimatedProb, a.CurrentPriceYes), a.CurrentPriceYes)
		if tradeSize <= 0 {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: "kelly_size_zero (edge insufficient or invalid price)",
			})
			continue
		}

		// USD horizon quota (replaces Q1/Q2 trade-count caps from v6).
		var usdCap float64
		switch a.Horizon {
		case "short":
			usdCap = usdCapShort
		case "medium":
			usdCap = usdCapMedium
		case "long":
			usdCap = usdCapLong
		default:
			rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: "Q1_horizon_unknown"})
			continue
		}
		if usdByHorizon[a.Horizon]+tradeSize > usdCap {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question,
				Reason: fmt.Sprintf("Q1_horizon_usd_cap_%s: used $%.0f + $%.0f > cap $%.0f",
					a.Horizon, usdByHorizon[a.Horizon], tradeSize, usdCap),
			})
			continue
		}

		// Total exposure cap (Kelly v7 — bankroll × max_total_exposure_pct).
		if portfolio.UsedExposure+tradeSize > maxTotalExposureUSD {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question,
				Reason: fmt.Sprintf("kelly_total_exposure_cap: used $%.0f + $%.0f > cap $%.0f",
					portfolio.UsedExposure, tradeSize, maxTotalExposureUSD),
			})
			continue
		}
		// Legacy hard floor still respected.
		if portfolio.UsedExposure+tradeSize > cfg.MaxExposureUSD {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: "exposure_cap_reached",
			})
			continue
		}

		// v5 Q3: liquidity gate. Reject if our trade would be >25% of visible orderbook.
		// Threshold: liquidity_usd >= tradeSize * cfg.LiquidityMinRatio (default 4 → ≤25% depth).
		if cfg.LiquidityMinRatio > 0 && a.LiquidityUSD > 0 && a.LiquidityUSD < tradeSize*cfg.LiquidityMinRatio {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: fmt.Sprintf("Q3_low_liquidity: $%.0f < $%.0f (size $%.0f × %.0f)", a.LiquidityUSD, tradeSize*cfg.LiquidityMinRatio, tradeSize, cfg.LiquidityMinRatio),
			})
			continue
		}

		// max_same_category: skip when category is "uncategorized" (scanner cannot classify;
		// applying the cap there would otherwise cluster all markets into one bucket and
		// kill diversity. Real categorical correlation will be enforced later by a smarter
		// taxonomy in the scanner).
		if a.Category != "" && a.Category != "uncategorized" {
			sameCat := 0
			for _, p := range portfolio.Positions {
				if p.Category == a.Category {
					sameCat++
				}
			}
			if sameCat >= cfg.MaxSameCategory {
				rejects = append(rejects, types.Rejection{
					Timestamp: nowStr, MarketID: a.CandidateID,
					Question: a.Question, Reason: "max_same_category_reached",
				})
				continue
			}
		}

		diffPct := math.Abs(a.CurrentPriceYes-a.ApprovedPriceYes) / a.ApprovedPriceYes * 100
		if diffPct > 2.0 {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: fmt.Sprintf("trigger_out_of_range: diff=%.2f%%", diffPct),
			})
			continue
		}

		// v7: side is dictated by edge direction. If LLM probability > implied
		// we buy YES; if lower we buy NO. The "target" price for the monitor
		// is the LLM probability itself (estimatedProb).
		side := "yes"
		entryPrice := a.CurrentPriceYes
		if a.EstimatedProb > 0 && a.EstimatedProb < a.CurrentPriceYes {
			side = "no"
			entryPrice = 1.0 - a.CurrentPriceYes
		}
		trade := types.ActiveTrade{
			ID:                 fmt.Sprintf("T-%s-%d", a.CandidateID, now.UnixMilli()),
			MarketID:           a.CandidateID,
			Slug:               a.Slug,
			Question:           a.Question,
			Side:               side,
			EntryPrice:         entryPrice,
			Size:               tradeSize,
			SizeUSD:            tradeSize,
			Category:           a.Category,
			EntryTimestamp:     nowStr,
			ApprovedPrice:      a.ApprovedPriceYes,
			DaysToResolution:   a.DaysToResolution,
			Horizon:            a.Horizon,
			EndDate:            a.EndDate,
			LiquidityUSD:       a.LiquidityUSD,
			SourcesUsed:        a.SourcesUsed,
			EstimatedProb:      a.EstimatedProb,
			TargetProb:         a.EstimatedProb,
			ThesisInvalidation: a.ThesisInvalidation,
			EdgeType:           a.EdgeType,
		}

		appendActive(trade)
		portfolio.Positions = append(portfolio.Positions, trade)
		portfolio.UsedExposure += tradeSize
		portfolio.Bankroll -= tradeSize
		todayCount++
		openCount++
		usdByHorizon[a.Horizon] += tradeSize

		// Persist Obsidian trade .md once per real entry. Brain used to write it
		// at approve time, which produced one .md per re-approval day (5+ duplicates
		// before the executor took the slug). Moved here so 1 .md == 1 real trade.
		var tradeSources []decisionlog.TradeSource
		for _, sc := range a.SourcesUsed {
			tradeSources = append(tradeSources, decisionlog.TradeSource{
				Domain: sc.Domain, URL: sc.URL,
				HeadlineTitle: sc.HeadlineTitle, PublishedDate: sc.PublishedDate,
			})
		}
		if _, wErr := decisionlog.WriteTrade(decisionsDir, decisionlog.TradeDecision{
			Slug: a.Slug, MarketID: a.CandidateID, Question: a.Question, Category: a.Category,
			EntryPrice: a.CurrentPriceYes, SizeUSD: tradeSize,
			Horizon: a.Horizon, DaysToResolution: a.DaysToResolution,
			SourcesUsed: tradeSources,
		}); wErr != nil {
			log.Printf("WARN: decisionlog WriteTrade failed for %s: %v", a.Slug, wErr)
		}

		log.Printf("executor: entered [%s] %s @ %.4f size=$%.2f (open=%d/%d, bankroll=$%.0f)", a.Horizon, a.Question, trade.EntryPrice, tradeSize, openCount, cfg.MaxOpenPositions, portfolio.Bankroll)
	}

	savePortfolio(portfolio)
	writeRejects(rejects)
	log.Printf("executor: %d entered, %d rejected", len(approved)-len(rejects), len(rejects))
}

func loadOrCreatePortfolio(cfg *config.BotConfig) types.Portfolio {
	p := types.Portfolio{
		Bankroll:     5000.0,
		UsedExposure: 0.0,
		MaxExposure:  cfg.MaxExposureUSD,
		MaxPerTrade:  cfg.SizeMax,
		MaxSameCat:   cfg.MaxSameCategory,
		Positions:    []types.ActiveTrade{},
	}
	path := filepath.Join(dataDir, portfolioFile)
	data, err := os.ReadFile(path)
	if err != nil {
		saveJSON(path, p)
		return p
	}
	if err := json.Unmarshal(data, &p); err != nil {
		log.Printf("executor: corrupt portfolio.json, resetting: %v", err)
		reset := types.Portfolio{
			Bankroll: 5000.0, UsedExposure: 0.0,
			MaxExposure: cfg.MaxExposureUSD, MaxPerTrade: cfg.SizeMax,
			MaxSameCat: cfg.MaxSameCategory, Positions: []types.ActiveTrade{},
		}
		saveJSON(path, reset)
		return reset
	}
	// Always reflect current config limits in portfolio metadata
	p.MaxExposure = cfg.MaxExposureUSD
	p.MaxPerTrade = cfg.SizeMax
	p.MaxSameCat = cfg.MaxSameCategory
	return p
}

func readApproved() []types.Approved {
	path := filepath.Join(dataDir, approvedFile)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var out []types.Approved
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var a types.Approved
		if err := json.Unmarshal(line, &a); err != nil {
			log.Printf("executor: skip bad approved line: %v", err)
			continue
		}
		out = append(out, a)
	}
	return out
}

func appendActive(t types.ActiveTrade) {
	path := filepath.Join(dataDir, activeFile)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("executor: cannot append active: %v", err)
		return
	}
	defer f.Close()
	_ = json.NewEncoder(f).Encode(t)
}

func writeRejects(rejects []types.Rejection) {
	if len(rejects) == 0 {
		return
	}
	path := filepath.Join(dataDir, rejectFile)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("executor: cannot write rejections: %v", err)
		return
	}
	defer f.Close()
	for _, r := range rejects {
		_ = json.NewEncoder(f).Encode(r)
	}
}

func savePortfolio(p types.Portfolio) {
	saveJSON(filepath.Join(dataDir, portfolioFile), p)
}

func saveJSON(path string, v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	_ = os.WriteFile(path, data, 0644)
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
