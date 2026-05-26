// Package closer holds the side-effect helpers that materialize a ClosedTrade:
// append to closed.jsonl, rewrite active.jsonl, mutate portfolio.json, patch
// the trade decision .md, and attribute the outcome to each source domain.
//
// Extracted from monitor.go so that monitor and the force_close_horizon_excess
// one-shot can share the same closure logic. No behavior change vs. the previous
// private functions in monitor.
package closer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

// Paths mirror monitor.go's constants. Exposed so callers can override in tests.
var (
	DataDir         = "vault/inbox/trading"
	ActiveFile      = "active.jsonl"
	ClosedFile      = "closed.jsonl"
	PortfolioFile   = "portfolio.json"
	DecisionsDir    = "vault/03-decisions"
	SourceStatsPath = "vault/agents/polymarket-bot/source-stats.jsonl"
)

// ReadActive parses active.jsonl. Returns nil on missing file.
func ReadActive() []types.ActiveTrade {
	path := filepath.Join(DataDir, ActiveFile)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var out []types.ActiveTrade
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var a types.ActiveTrade
		if err := json.Unmarshal(line, &a); err != nil {
			log.Printf("closer: skip bad active line: %v", err)
			continue
		}
		out = append(out, a)
	}
	return out
}

// AppendClosed appends one or more ClosedTrade records to closed.jsonl.
func AppendClosed(closed []types.ClosedTrade) {
	path := filepath.Join(DataDir, ClosedFile)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("closer: cannot append closed: %v", err)
		return
	}
	defer f.Close()
	for _, c := range closed {
		_ = json.NewEncoder(f).Encode(c)
	}
}

// RewriteActive replaces active.jsonl with the remaining open positions.
func RewriteActive(remaining []types.ActiveTrade) {
	path := filepath.Join(DataDir, ActiveFile)
	f, err := os.Create(path)
	if err != nil {
		log.Printf("closer: cannot rewrite active: %v", err)
		return
	}
	defer f.Close()
	for _, a := range remaining {
		_ = json.NewEncoder(f).Encode(a)
	}
}

// UpdatePortfolio mutates portfolio.json to credit realized PnL + principal
// back to bankroll, decrement used_exposure, and drop the closed positions.
func UpdatePortfolio(closed []types.ClosedTrade) {
	path := filepath.Join(DataDir, PortfolioFile)
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

// AppendSourceStats writes one row per cited source so the brain's v4 source
// blacklist can decay bad outlets. No-op when SourcesUsed is empty.
func AppendSourceStats(ct types.ClosedTrade) error {
	if len(ct.SourcesUsed) == 0 {
		return nil
	}
	outcome := "breakeven"
	if ct.Pnl > 0 {
		outcome = "win"
	} else if ct.Pnl < 0 {
		outcome = "loss"
	}
	if err := os.MkdirAll(filepath.Dir(SourceStatsPath), 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(SourceStatsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

// PatchDecisionMD rewrites the frontmatter (outcome/pnl/closed_at/exit_reason/days_open)
// on every <date>-polymarket-trade-<slug>.md so historical duplicates (created by
// brain on consecutive re-approvals before WriteTrade was moved to executor) also
// reflect the real outcome. Silently skips when no matching file exists.
func PatchDecisionMD(ct types.ClosedTrade) error {
	if ct.Slug == "" {
		return nil
	}
	entries, err := os.ReadDir(DecisionsDir)
	if err != nil {
		return err
	}
	suffix := "-polymarket-trade-" + ct.Slug + ".md"
	var patched int
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), suffix) {
			continue
		}
		path := filepath.Join(DecisionsDir, e.Name())
		if err := patchOne(path, ct); err != nil {
			log.Printf("closer: patch %s: %v", path, err)
			continue
		}
		patched++
	}
	if patched > 1 {
		log.Printf("closer: patched %d duplicate .md for slug %s", patched, ct.Slug)
	}
	return nil
}

func patchOne(path string, ct types.ClosedTrade) error {
	raw, err := os.ReadFile(path)
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
	patched = patchYAMLLine(patched, "pnl_usd", fmt.Sprintf("%.2f", ct.Pnl))
	patched = patchYAMLLine(patched, "closed_at", ct.ExitTime)
	patched = patchYAMLLine(patched, "exit_reason", ct.Reason)
	patched = patchYAMLLine(patched, "days_open", fmt.Sprintf("%.2f", ct.DaysOpen))
	out := "---\n" + patched + "\n---\n" + rest
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(out), 0644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
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
