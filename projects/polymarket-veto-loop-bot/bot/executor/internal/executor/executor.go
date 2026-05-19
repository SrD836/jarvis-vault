package executor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/executor/internal/types"
)

const (
	dataDir      = "vault/inbox/trading"
	approvedFile = "approved.jsonl"
	activeFile   = "active.jsonl"
	rejectFile   = "rejections.jsonl"
	portfolioFile = "portfolio.json"

	tradeSize    = 150.0
	maxExposure  = 1000.0
	maxSameCat   = 2
)

// Run executes one pass of the executor.
func Run() {
	portfolio := loadOrCreatePortfolio()
	approved := readApproved()
	if len(approved) == 0 {
		log.Println("executor: no approved candidates")
		return
	}

	var rejects []types.Rejection
	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)

	for _, a := range approved {
		// Check max exposure cap
		if portfolio.UsedExposure+tradeSize > maxExposure {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: "exposure_cap_reached",
			})
			continue
		}

		// Check correlation: max 2 same category
		sameCat := 0
		for _, p := range portfolio.Positions {
			if p.Category == a.Category {
				sameCat++
			}
		}
		if sameCat >= maxSameCat {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: "max_same_category_reached",
			})
			continue
		}

		// Check trigger ±2% of approved price
		diffPct := math.Abs(a.CurrentPriceYes-a.ApprovedPriceYes) / a.ApprovedPriceYes * 100
		if diffPct > 2.0 {
			rejects = append(rejects, types.Rejection{
				Timestamp: nowStr, MarketID: a.CandidateID,
				Question: a.Question, Reason: fmt.Sprintf("trigger_out_of_range: diff=%.2f%%", diffPct),
			})
			continue
		}

		// Enter trade
		trade := types.ActiveTrade{
			ID:             fmt.Sprintf("T-%s-%d", a.CandidateID, now.UnixMilli()),
			MarketID:       a.CandidateID,
			Question:       a.Question,
			Side:           "yes",
			EntryPrice:     a.CurrentPriceYes,
			Size:           tradeSize,
			Category:       a.Category,
			EntryTimestamp: nowStr,
			ApprovedPrice:  a.ApprovedPriceYes,
		}

		appendActive(trade)
		portfolio.Positions = append(portfolio.Positions, trade)
		portfolio.UsedExposure += tradeSize
		portfolio.Bankroll -= tradeSize // capital committed

		log.Printf("executor: entered %s (%.4f, $%.0f)", a.Question, trade.EntryPrice, trade.Size)
	}

	savePortfolio(portfolio)
	writeRejects(rejects)
	log.Printf("executor: %d entered, %d rejected", len(approved)-len(rejects), len(rejects))
}

func loadOrCreatePortfolio() types.Portfolio {
	p := types.Portfolio{
		Bankroll:     10000.0,
		UsedExposure: 0.0,
		MaxExposure:  maxExposure,
		MaxPerTrade:  tradeSize,
		MaxSameCat:   maxSameCat,
		Positions:    []types.ActiveTrade{},
	}
	path := filepath.Join(dataDir, portfolioFile)
	data, err := os.ReadFile(path)
	if err != nil {
		// Create default
		saveJSON(path, p)
		return p
	}
	if err := json.Unmarshal(data, &p); err != nil {
		log.Printf("executor: corrupt portfolio.json, resetting: %v", err)
		saveJSON(path, types.Portfolio{
			Bankroll: 10000.0, UsedExposure: 0.0, MaxExposure: maxExposure,
			MaxPerTrade: tradeSize, MaxSameCat: maxSameCat, Positions: []types.ActiveTrade{},
		})
		return p
	}
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
	path := filepath.Join(dataDir, portfolioFile)
	saveJSON(path, p)
}

func saveJSON(path string, v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	_ = os.WriteFile(path, data, 0644)
}
