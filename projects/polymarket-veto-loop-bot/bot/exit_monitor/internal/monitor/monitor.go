package monitor

import (
	"bufio"
	"encoding/json"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/loglosses"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const (
	dataDir       = "vault/inbox/trading"
	activeFile    = "active.jsonl"
	closedFile    = "closed.jsonl"
	portfolioFile = "portfolio.json"

	memoryPath   = "vault/agents/polymarket-bot/memory.md"
	decisionsDir = "vault/03-decisions"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("exit_monitor: load config: %v", err)
	}
	log.Printf("exit_monitor: config loaded stopLoss=%.0f takeProfitPct=%.0f%%",
		cfg.StopLossUSD, cfg.TakeProfitPct)

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

		entryPrice := a.EntryPrice
		pnlPct := (price - entryPrice) / entryPrice * 100
		pnlUSD := (price - entryPrice) / entryPrice * a.Size

		reason := ""
		switch {
		case mktClosed:
			reason = "market_closed"
		case pnlPct >= cfg.TakeProfitPct:
			reason = "take_profit"
		case pnlUSD <= -cfg.StopLossUSD:
			reason = "stop_loss"
		default:
			remaining = append(remaining, a)
			continue
		}

		pnlR := math.Round(pnlUSD*100) / 100
		exitTs := now.Format(time.RFC3339)
		ct := types.ClosedTrade{
			ID:         a.ID,
			MarketID:   a.MarketID,
			Slug:       a.Slug,
			Question:   a.Question,
			Category:   a.Category,
			Side:       a.Side,
			Size:       a.Size,
			EntryPrice: entryPrice,
			ExitPrice:  price,
			EntryTime:  a.EntryTimestamp,
			ExitTime:   exitTs,
			Pnl:        pnlR,
			PnlUSD:     pnlR,
			PnlPct:     math.Round(pnlPct*100) / 100,
			Reason:     reason,
		}
		closed = append(closed, ct)
		log.Printf("exit_monitor: closed %s (%s): PnL $%.2f (%.2f%%)", a.Question, reason, ct.Pnl, ct.PnlPct)

		if pnlR < 0 {
			loglosses.LogLoss(
				decisionsDir, memoryPath,
				ct.ID, ct.Slug, ct.Question, ct.Category,
				ct.EntryPrice, ct.ExitPrice, ct.Size, ct.Pnl,
				ct.EntryTime, ct.ExitTime, ct.Reason,
			)
		}
	}

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
		return
	}
	var p map[string]interface{}
	if err := json.Unmarshal(data, &p); err != nil {
		return
	}
	bankroll, _ := p["bankroll"].(float64)
	used, _ := p["used_exposure"].(float64)
	for _, c := range closed {
		bankroll += c.Pnl
		used -= c.Size
		if used < 0 {
			used = 0
		}
	}
	p["bankroll"] = bankroll
	p["used_exposure"] = used
	if positions, ok := p["positions"].([]interface{}); ok {
		closedIDs := map[string]bool{}
		for _, c := range closed {
			closedIDs[c.ID] = true
		}
		filtered := []interface{}{}
		for _, pos := range positions {
			if m, ok := pos.(map[string]interface{}); ok {
				if id, ok := m["id"].(string); ok && closedIDs[id] {
					continue
				}
			}
			filtered = append(filtered, pos)
		}
		p["positions"] = filtered
	}
	out, _ := json.MarshalIndent(p, "", "  ")
	_ = os.WriteFile(path, out, 0644)
}
