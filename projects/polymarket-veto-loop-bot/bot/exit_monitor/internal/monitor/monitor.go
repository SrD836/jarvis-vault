package monitor

import (
	"bufio"
	"encoding/json"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const (
	dataDir      = "vault/inbox/trading"
	activeFile   = "active.jsonl"
	closedFile   = "closed.jsonl"
	portfolioFile = "portfolio.json"

	stopLossUSD = 150.0
	takeProfitPct = 60.0
)

// Run executes one pass of the exit monitor.
func Run() {
	actives := readActive()
	if len(actives) == 0 {
		log.Println("exit_monitor: no active positions")
		return
	}

	var closed []types.ClosedTrade
	var remaining []types.ActiveTrade
	now := time.Now().UTC()

	for _, a := range actives {
		price, mktClosed, err := polyclient.GetMarketPrice(a.MarketID)
		if err != nil {
			log.Printf("exit_monitor: price fetch error for %s: %v", a.MarketID, err)
			remaining = append(remaining, a)
			continue
		}

		// Determine entry price (yes side)
		entryPrice := a.EntryPrice
		pnlPct := (price - entryPrice) / entryPrice * 100
		pnlUSD := (price - entryPrice) / entryPrice * a.Size // proportional P&L

		reason := ""
		switch {
		case mktClosed:
			reason = "market_closed"
		case pnlPct >= takeProfitPct:
			reason = "take_profit"
		case pnlUSD <= -stopLossUSD:
			reason = "stop_loss"
		default:
			remaining = append(remaining, a)
			continue
		}

		ct := types.ClosedTrade{
			ID:         a.ID,
			MarketID:   a.MarketID,
			EntryPrice: entryPrice,
			ExitPrice:  price,
			EntryTime:  a.EntryTimestamp,
			ExitTime:   now.Format(time.RFC3339),
			PnlUSD:     math.Round(pnlUSD*100) / 100,
			PnlPct:     math.Round(pnlPct*100) / 100,
			Reason:     reason,
		}
		closed = append(closed, ct)
		log.Printf("exit_monitor: closed %s (%s): PnL $%.2f (%.2f%%)", a.Question, reason, ct.PnlUSD, ct.PnlPct)
	}

	// Update files
	if len(closed) > 0 {
		appendClosed(closed)
		updatePortfolio(closed)
	}
	rewriteActive(remaining)
	log.Printf("exit_monitor: %d closed, %d remaining", len(closed), len(remaining))
}

func readActive() []types.ActiveTrade {
	path := filepath.Join(dataDir, activeFile)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var out []types.ActiveTrade
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var a types.ActiveTrade
		if err := json.Unmarshal(line, &a); err != nil {
			log.Printf("exit_monitor: skip bad active line: %v", err)
			continue
		}
		out = append(out, a)
	}
	return out
}

func appendClosed(closed []types.ClosedTrade) {
	path := filepath.Join(dataDir, closedFile)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("exit_monitor: cannot append closed: %v", err)
		return
	}
	defer f.Close()
	for _, c := range closed {
		_ = json.NewEncoder(f).Encode(c)
	}
}

func rewriteActive(remaining []types.ActiveTrade) {
	path := filepath.Join(dataDir, activeFile)
	f, err := os.Create(path)
	if err != nil {
		log.Printf("exit_monitor: cannot rewrite active: %v", err)
		return
	}
	defer f.Close()
	for _, a := range remaining {
		_ = json.NewEncoder(f).Encode(a)
	}
}

func updatePortfolio(closed []types.ClosedTrade) {
	path := filepath.Join(dataDir, portfolioFile)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("exit_monitor: cannot read portfolio for update: %v", err)
		return
	}

	var pf struct {
		Bankroll     float64            `json:"bankroll"`
		UsedExposure float64            `json:"used_exposure"`
		MaxExposure  float64            `json:"max_exposure"`
		MaxPerTrade  float64            `json:"max_per_trade"`
		MaxSameCat   int                `json:"max_same_category"`
		Positions    []types.ActiveTrade `json:"positions"`
	}
	if err := json.Unmarshal(data, &pf); err != nil {
		log.Printf("exit_monitor: corrupt portfolio.json: %v", err)
		return
	}

	// Build set of closed IDs
	closedIDs := make(map[string]bool)
	for _, c := range closed {
		closedIDs[c.ID] = true
	}

	// Filter out closed positions
	var newPositions []types.ActiveTrade
	for _, p := range pf.Positions {
		if !closedIDs[p.ID] {
			newPositions = append(newPositions, p)
		}
	}

	// Update bankroll and exposure
	for _, c := range closed {
		pf.Bankroll += c.PnlUSD
		// Find the original size for this trade
		for _, p := range pf.Positions {
			if p.ID == c.ID {
				pf.Bankroll += p.Size // return size capital
				pf.UsedExposure -= p.Size
				break
			}
		}
	}
	pf.Positions = newPositions

	out, _ := json.MarshalIndent(pf, "", "  ")
	_ = os.WriteFile(path, out, 0644)
}
