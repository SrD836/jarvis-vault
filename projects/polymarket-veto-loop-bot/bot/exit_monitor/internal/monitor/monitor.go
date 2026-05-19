package monitor

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
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/loglosses"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const (
	dataDir       = "vault/inbox/trading"
	activeFile    = "active.jsonl"
	closedFile    = "closed.jsonl"
	portfolioFile = "portfolio.json"

	memoryPath      = "vault/agents/polymarket-bot/memory.md"
	decisionsDir    = "vault/03-decisions"
	sourceStatsPath = "vault/agents/polymarket-bot/source-stats.jsonl"
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
		}
		closed = append(closed, ct)
		earlyTag := ""
		if early {
			earlyTag = " EARLY"
		}
		log.Printf("exit_monitor: closed %s (%s%s) @ %.4f src=%s liq=$%.0f: PnL $%.2f (%.2f%%)", a.Question, reason, earlyTag, price, priceSrc, quote.LiquidityUSD, ct.Pnl, ct.PnlPct)

		if pnlR < 0 {
			loglosses.LogLoss(
				decisionsDir, memoryPath,
				ct.ID, ct.Slug, ct.Question, ct.Category,
				ct.EntryPrice, ct.ExitPrice, ct.Size, ct.Pnl,
				ct.EntryTime, ct.ExitTime, ct.Reason,
			)
		}

		// v4: attribute outcome to every source that informed the decision.
		if err := appendSourceStats(ct); err != nil {
			log.Printf("exit_monitor: source-stats append failed: %v", err)
		}
		// v4: sync Obsidian decision .md (frontmatter outcome/pnl/closed_at).
		if err := patchDecisionMD(ct); err != nil {
			log.Printf("exit_monitor: decision .md patch failed for %s: %v", ct.Slug, err)
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
		// v4: return the principal (Size) plus realized PnL to bankroll.
		// Previously only c.Pnl was added, which caused bankroll to bleed by ~Size every cycle.
		bankroll += c.Size + c.Pnl
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


func computeDaysOpen(entryTS string, now time.Time) float64 {
	t, err := time.Parse(time.RFC3339, entryTS)
	if err != nil {
		return 0
	}
	return math.Round(now.Sub(t).Hours()/24*100) / 100
}

func appendSourceStats(ct types.ClosedTrade) error {
	if len(ct.SourcesUsed) == 0 {
		return nil
	}
	outcome := "breakeven"
	if ct.Pnl > 0 {
		outcome = "win"
	} else if ct.Pnl < 0 {
		outcome = "loss"
	}
	if err := os.MkdirAll(filepath.Dir(sourceStatsPath), 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(sourceStatsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	for _, s := range ct.SourcesUsed {
		dom := strings.ToLower(strings.TrimSpace(s.Domain))
		if dom == "" || dom == "legacy" {
			continue
		}
		_ = enc.Encode(map[string]interface{}{
			"ts":        ct.ExitTime,
			"trade_id":  ct.ID,
			"domain":    dom,
			"outcome":   outcome,
			"pnl_usd":   ct.Pnl,
			"days_open": ct.DaysOpen,
		})
	}
	return nil
}

// patchDecisionMD looks for vault/03-decisions/<date>-polymarket-trade-<slug>.md
// (created by brain.go at approve time) and rewrites its frontmatter outcome/pnl/closed_at.
// Silently skips if the file is not found (legacy trades).
func patchDecisionMD(ct types.ClosedTrade) error {
	if ct.Slug == "" {
		return nil
	}
	// Search across all decisions to handle date mismatch.
	entries, err := os.ReadDir(decisionsDir)
	if err != nil {
		return err
	}
	var found string
	suffix := "-polymarket-trade-" + ct.Slug + ".md"
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), suffix) {
			found = filepath.Join(decisionsDir, e.Name())
			break
		}
	}
	if found == "" {
		return nil // brain didn't create one (older approve), silent skip
	}
	raw, err := os.ReadFile(found)
	if err != nil {
		return err
	}
	text := string(raw)
	if !strings.HasPrefix(text, "---\n") {
		return nil
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return nil
	}
	fm := text[4 : 4+end]
	rest := text[4+end+5:]

	outcome := "breakeven"
	if ct.Pnl > 0 {
		outcome = "win"
	} else if ct.Pnl < 0 {
		outcome = "loss"
	}
	patched := patchYAMLLine(fm, "outcome", outcome)
	patched = ensureYAMLLine(patched, "pnl_usd", fmt.Sprintf("%.2f", ct.Pnl))
	patched = ensureYAMLLine(patched, "closed_at", ct.ExitTime)
	patched = ensureYAMLLine(patched, "exit_reason", ct.Reason)
	patched = ensureYAMLLine(patched, "days_open", fmt.Sprintf("%.2f", ct.DaysOpen))
	out := "---\n" + patched + "\n---\n" + rest
	tmp := found + ".tmp"
	if err := os.WriteFile(tmp, []byte(out), 0644); err != nil {
		return err
	}
	return os.Rename(tmp, found)
}

func patchYAMLLine(fm, key, value string) string {
	lines := strings.Split(fm, "\n")
	prefix := key + ":"
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimLeft(l, " \t"), prefix) {
			lines[i] = key + ": " + value
			return strings.Join(lines, "\n")
		}
	}
	return strings.Join(append(lines, key+": "+value), "\n")
}

func ensureYAMLLine(fm, key, value string) string {
	return patchYAMLLine(fm, key, value)
}
