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

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/decisionlog"
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
	if cfg.Mode != "simulation" {
		panic("live trading not authorized — set mode=simulation in bot/config.json")
	}
	log.Printf("executor v4: mode=%s sizeMode=%s sizeFrac=%.3f sizeMin=%.0f sizeMax=%.0f maxExposure=%.0f maxOpen=%d maxSameCat=%d maxPerDay=%d quota=%.0f/%.0f/%.0f",
		cfg.Mode, cfg.SizeMode, cfg.SizeFraction, cfg.SizeMin, cfg.SizeMax,
		cfg.MaxExposureUSD, cfg.MaxOpenPositions, cfg.MaxSameCategory, cfg.MaxNewTradesPerDay,
		cfg.HorizonQuotaShort*100, cfg.HorizonQuotaMedium*100, cfg.HorizonQuotaLong*100)

	portfolio := loadOrCreatePortfolio(cfg)
	approved := readApproved()
	if len(approved) == 0 {
		log.Println("executor: no approved candidates")
		return
	}

	todayPrefix := time.Now().UTC().Format("2006-01-02")
	todayCount := 0
	for _, p := range portfolio.Positions {
		if strings.HasPrefix(p.EntryTimestamp, todayPrefix) {
			todayCount++
		}
	}

	// Quota slots per horizon for THIS run (ceil of quota * (max - already_today))
	slotsRemaining := cfg.MaxNewTradesPerDay - todayCount
	if slotsRemaining < 0 {
		slotsRemaining = 0
	}
	slotShort := ceilQuota(cfg.HorizonQuotaShort, slotsRemaining)
	slotMedium := ceilQuota(cfg.HorizonQuotaMedium, slotsRemaining)
	slotLong := ceilQuota(cfg.HorizonQuotaLong, slotsRemaining)
	log.Printf("executor: quota slots this run: short=%d medium=%d long=%d (of %d remaining today)",
		slotShort, slotMedium, slotLong, slotsRemaining)

	// Q2 caps: portfolio-aware horizon quota. Q1 only counts trades opened today;
	// historical positions from previous days don't constrain it, which is how
	// the portfolio drifted to 1/19/17 (short/medium/long) instead of 70/15/15.
	// Q2 enforces the quota against MaxOpenPositions directly.
	capShort := ceilQuota(cfg.HorizonQuotaShort, cfg.MaxOpenPositions)
	capMedium := ceilQuota(cfg.HorizonQuotaMedium, cfg.MaxOpenPositions)
	capLong := ceilQuota(cfg.HorizonQuotaLong, cfg.MaxOpenPositions)
	openByHorizon := map[string]int{}
	for _, p := range portfolio.Positions {
		openByHorizon[p.Horizon]++
	}
	log.Printf("executor: portfolio quota caps: short=%d medium=%d long=%d (open now: short=%d medium=%d long=%d)",
		capShort, capMedium, capLong,
		openByHorizon["short"], openByHorizon["medium"], openByHorizon["long"])

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

		// Q1: horizon quota for new trades opened today.
		switch a.Horizon {
		case "short":
			if slotShort <= 0 {
				rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: "Q1_horizon_quota_full_short"})
				continue
			}
		case "medium":
			if slotMedium <= 0 {
				rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: "Q1_horizon_quota_full_medium"})
				continue
			}
		case "long":
			if slotLong <= 0 {
				rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: "Q1_horizon_quota_full_long"})
				continue
			}
		default:
			rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: "Q1_horizon_unknown"})
			continue
		}

		// Q2: portfolio-aware horizon cap (sees inherited positions, not just today).
		switch a.Horizon {
		case "short":
			if openByHorizon["short"] >= capShort {
				rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: fmt.Sprintf("Q2_portfolio_quota_full_short_%d_of_%d", openByHorizon["short"], capShort)})
				continue
			}
		case "medium":
			if openByHorizon["medium"] >= capMedium {
				rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: fmt.Sprintf("Q2_portfolio_quota_full_medium_%d_of_%d", openByHorizon["medium"], capMedium)})
				continue
			}
		case "long":
			if openByHorizon["long"] >= capLong {
				rejects = append(rejects, types.Rejection{Timestamp: nowStr, MarketID: a.CandidateID, Question: a.Question, Reason: fmt.Sprintf("Q2_portfolio_quota_full_long_%d_of_%d", openByHorizon["long"], capLong)})
				continue
			}
		}

		tradeSize := cfg.ComputeTradeSize(portfolio.Bankroll)
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

		trade := types.ActiveTrade{
			ID:               fmt.Sprintf("T-%s-%d", a.CandidateID, now.UnixMilli()),
			MarketID:         a.CandidateID,
			Slug:             a.Slug,
			Question:         a.Question,
			Side:             "yes",
			EntryPrice:       a.CurrentPriceYes,
			Size:             tradeSize,
			SizeUSD:          tradeSize,
			Category:         a.Category,
			EntryTimestamp:   nowStr,
			ApprovedPrice:    a.ApprovedPriceYes,
			DaysToResolution: a.DaysToResolution,
			Horizon:          a.Horizon,
			EndDate:          a.EndDate,
			LiquidityUSD:     a.LiquidityUSD,
			SourcesUsed:      a.SourcesUsed,
		}

		appendActive(trade)
		portfolio.Positions = append(portfolio.Positions, trade)
		portfolio.UsedExposure += tradeSize
		portfolio.Bankroll -= tradeSize
		todayCount++
		openCount++
		openByHorizon[a.Horizon]++

		switch a.Horizon {
		case "short":
			slotShort--
		case "medium":
			slotMedium--
		case "long":
			slotLong--
		}

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
