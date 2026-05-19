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
	"github.com/davidgn/polymarket-veto-loop-bot/bot/executor/internal/types"
)

const (
	dataDir       = "vault/inbox/trading"
	approvedFile  = "approved.jsonl"
	activeFile    = "active.jsonl"
	rejectFile    = "rejections.jsonl"
	portfolioFile = "portfolio.json"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("executor: load config: %v", err)
	}
	if cfg.Mode != "simulation" {
		panic("live trading not authorized — set mode=simulation in bot/config.json")
	}
	log.Printf("executor v3: config loaded mode=%s tradeSize=%.0f maxExposure=%.0f maxSameCat=%d maxPerDay=%d quota=%.0f/%.0f/%.0f",
		cfg.Mode, cfg.TradeSizeUSD, cfg.MaxExposureUSD, cfg.MaxSameCategory, cfg.MaxNewTradesPerDay,
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

	var rejects []types.Rejection
	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)

	activeIDs := map[string]bool{}
	for _, p := range portfolio.Positions {
		activeIDs[p.MarketID] = true
	}
	for _, a := range approved {
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

		// Q1: horizon quota — reject if the candidate's horizon bucket is full.
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

		if portfolio.UsedExposure+cfg.TradeSizeUSD > cfg.MaxExposureUSD {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: "exposure_cap_reached",
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
			Question:         a.Question,
			Side:             "yes",
			EntryPrice:       a.CurrentPriceYes,
			Size:             cfg.TradeSizeUSD,
			Category:         a.Category,
			EntryTimestamp:   nowStr,
			ApprovedPrice:    a.ApprovedPriceYes,
			DaysToResolution: a.DaysToResolution,
			Horizon:          a.Horizon,
			EndDate:          a.EndDate,
		}

		appendActive(trade)
		portfolio.Positions = append(portfolio.Positions, trade)
		portfolio.UsedExposure += cfg.TradeSizeUSD
		portfolio.Bankroll -= cfg.TradeSizeUSD
		todayCount++

		switch a.Horizon {
		case "short":
			slotShort--
		case "medium":
			slotMedium--
		case "long":
			slotLong--
		}
		log.Printf("executor: entered [%s] %s (%.4f, $%.0f)", a.Horizon, a.Question, trade.EntryPrice, trade.Size)
	}

	savePortfolio(portfolio)
	writeRejects(rejects)
	log.Printf("executor: %d entered, %d rejected", len(approved)-len(rejects), len(rejects))
}

func loadOrCreatePortfolio(cfg *config.BotConfig) types.Portfolio {
	p := types.Portfolio{
		Bankroll:     10000.0,
		UsedExposure: 0.0,
		MaxExposure:  cfg.MaxExposureUSD,
		MaxPerTrade:  cfg.TradeSizeUSD,
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
			Bankroll: 10000.0, UsedExposure: 0.0,
			MaxExposure: cfg.MaxExposureUSD, MaxPerTrade: cfg.TradeSizeUSD,
			MaxSameCat: cfg.MaxSameCategory, Positions: []types.ActiveTrade{},
		}
		saveJSON(path, reset)
		return reset
	}
	// Always reflect current config limits in portfolio metadata
	p.MaxExposure = cfg.MaxExposureUSD
	p.MaxPerTrade = cfg.TradeSizeUSD
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
