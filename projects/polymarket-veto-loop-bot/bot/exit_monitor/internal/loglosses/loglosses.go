package loglosses

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/postmortem"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const maxRowsKeep = 500

var mu sync.Mutex

// LogClose handles both losses and wins. PnL < 0 → losses section, else wins.
// Writes the .md decision log only for losses (post-mortem). Wins go only to memory.md.
func LogClose(decisionsDir, memoryPath string, ct types.ClosedTrade) {
	if ct.Pnl < 0 {
		writeDecisionLog(decisionsDir, ct)
		appendMemoryRow(memoryPath, ct, true)
	} else {
		appendMemoryRow(memoryPath, ct, false)
	}
}

// LogLoss is the back-compat shim used by existing callers.
func LogLoss(
	decisionsDir, memoryPath string,
	tradeID, slug, question, category string,
	entryPrice, exitPrice, size, pnl float64,
	entryTime, exitTime, exitReason string,
) {
	if pnl >= 0 {
		return
	}
	ct := types.ClosedTrade{
		ID: tradeID, Slug: slug, Question: question, Category: category,
		EntryPrice: entryPrice, ExitPrice: exitPrice, Size: size, Pnl: pnl,
		EntryTime: entryTime, ExitTime: exitTime, Reason: exitReason,
	}
	LogClose(decisionsDir, memoryPath, ct)
}

// nodesEnabled reports whether per-decision Obsidian .md nodes should be written.
// Default OFF (mirrors decisionlog.nodesEnabled): memory.md + the inbox/trading
// *.jsonl logs already capture every close, and the post-mortem anti-pattern
// learning lands in memory.md regardless of the node. Set DECISION_NODES=1
// (true/on/yes) to re-enable nodes.
func nodesEnabled() bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("DECISION_NODES"))) {
	case "1", "true", "on", "yes":
		return true
	}
	return false
}

func writeDecisionLog(decisionsDir string, ct types.ClosedTrade) {
	date := time.Now().UTC().Format("2006-01-02")
	slug := ct.Slug
	if slug == "" {
		slug = ct.ID
	}
	slugSan := sanitize(slug)
	if len(slugSan) > 80 {
		slugSan = slugSan[:80]
	}
	name := fmt.Sprintf("%s-polymarket-loss-%s.md", date, slugSan)
	path := filepath.Join(decisionsDir, name)
	// Obsidian loss node gated behind DECISION_NODES (default OFF). The claudemax
	// post-mortem below STILL runs: its anti-pattern learning lands in memory.md
	// (## Anti-patterns identificados, consumed by the brain via M3), independent
	// of the .md node — so gating the node does NOT break learning.
	writeNode := nodesEnabled()
	if writeNode {
		if err := os.MkdirAll(decisionsDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "[loglosses] mkdir %s: %v\n", decisionsDir, err)
			writeNode = false
		}
	}

	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"Polymarket loss — %s\"\n", yamlEscape(slug))
	b.WriteString("type: decision\n")
	fmt.Fprintf(&b, "date: %s\n", date)
	fmt.Fprintf(&b, "decision: \"Trade cerrado en pérdida: %s (PnL %.2f USD)\"\n",
		yamlEscape(truncate(ct.Question, 80)), ct.Pnl)
	b.WriteString("alternatives:\n  - \"Haber vetado la tesis antes de entrar\"\n  - \"Cerrar antes (stop más conservador)\"\n")
	b.WriteString("outcome: confirmed_bad\n")
	b.WriteString("outcome_observed_after_days: 0\n")
	b.WriteString("tags: [decision, polymarket, bot, loss, post-mortem]\n")
	b.WriteString("related:\n  - \"[[agents/polymarket-bot]]\"\n  - \"[[agents/polymarket-bot/memory]]\"\n  - \"[[projects/polymarket-veto-loop-bot]]\"\n")
	if len(ct.SourcesUsed) > 0 {
		b.WriteString("sources_used: [")
		for i, s := range ct.SourcesUsed {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(s.Domain)
		}
		b.WriteString("]\n")
	}
	if ct.Horizon != "" {
		fmt.Fprintf(&b, "horizon: %s\n", ct.Horizon)
	}
	fmt.Fprintf(&b, "days_held: %.2f\n", ct.DaysOpen)
	b.WriteString("---\n\n")
	fmt.Fprintf(&b, "# Loss: %s\n\n", ct.Question)
	b.WriteString("## Trade\n\n")
	fmt.Fprintf(&b, "- **Trade ID**: `%s`\n", ct.ID)
	fmt.Fprintf(&b, "- **Slug**: `%s`\n", ct.Slug)
	fmt.Fprintf(&b, "- **Categoría**: %s\n", ct.Category)
	fmt.Fprintf(&b, "- **Entry price**: %.4f\n", ct.EntryPrice)
	fmt.Fprintf(&b, "- **Exit price**: %.4f\n", ct.ExitPrice)
	fmt.Fprintf(&b, "- **Size**: $%.2f\n", ct.Size)
	fmt.Fprintf(&b, "- **PnL**: $%.2f (%.2f%%)\n", ct.Pnl, ct.PnlPct)
	fmt.Fprintf(&b, "- **Days held**: %.2f\n", ct.DaysOpen)
	fmt.Fprintf(&b, "- **Horizon**: %s\n", ct.Horizon)
	fmt.Fprintf(&b, "- **Entry timestamp**: %s\n", ct.EntryTime)
	fmt.Fprintf(&b, "- **Exit timestamp**: %s\n", ct.ExitTime)
	fmt.Fprintf(&b, "- **Exit reason**: %s\n", ct.Reason)
	if len(ct.SourcesUsed) > 0 {
		b.WriteString("\n## Fuentes consultadas al aprobar\n\n")
		for _, s := range ct.SourcesUsed {
			fmt.Fprintf(&b, "- **%s**: %s\n", s.Domain, s.URL)
		}
	}
	b.WriteString("\n## Análisis\n\n_(autogenerado tras cierre — revisar manualmente para extraer lección)_\n\n")
	b.WriteString("## Human notes\n\n_(no se toca por automatización)_\n")

	if writeNode {
		if err := os.WriteFile(path, []byte(b.String()), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "[loglosses] write %s: %v\n", path, err)
		}
	}

	// Synchronous post-mortem via claudemax (~30-60s). exit_monitor runs as a
	// one-shot cron every 5 min so we have ample headroom; goroutine would be
	// killed when main() returns. Best-effort: postmortem.Run swallows errors.
	dashboardURL := os.Getenv("DASHBOARD_URL")
	if dashboardURL == "" {
		dashboardURL = "http://jarvis-dashboard:3000"
	}
	memoryPath := os.Getenv("EXIT_MEMORY_PATH")
	if memoryPath == "" {
		memoryPath = "vault/agents/polymarket-bot/memory.md"
	}
	postmortem.Run(dashboardURL, path, memoryPath, postmortem.LossContext{
		Slug: ct.Slug, Question: ct.Question, Category: ct.Category,
		EntryPrice: ct.EntryPrice, ExitPrice: ct.ExitPrice,
		SizeUSD: ct.Size, PnlUSD: ct.Pnl, DaysHeld: ct.DaysOpen,
		Horizon: ct.Horizon, ExitReason: ct.Reason,
		SourcesUsed: ct.SourcesUsed,
		Date:        date,
	})
}

func appendMemoryRow(memoryPath string, ct types.ClosedTrade, isLoss bool) {
	if memoryPath == "" {
		return
	}
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(memoryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[loglosses] read memory %s: %v\n", memoryPath, err)
		return
	}
	text := string(data)

	var header string
	nextHeader := "## Reglas blandas"
	if isLoss {
		header = "## Losses pattern"
	} else {
		header = "## Wins pattern"
	}

	if !strings.Contains(text, header) {
		idx := strings.Index(text, nextHeader)
		if idx < 0 {
			fmt.Fprintf(os.Stderr, "[loglosses] cannot find insertion anchor %q\n", nextHeader)
			return
		}
		block := header + " (últimos 100, append-only, rotación a tail)\n\n" +
			"| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |\n" +
			"|---|---|---|---|---|---|---|---|---|---|---|\n\n"
		text = text[:idx] + block + text[idx:]
	}

	idxSec := strings.Index(text, header)
	idxNext := strings.Index(text[idxSec+len(header):], nextHeader)
	var sectionEnd int
	if idxNext < 0 {
		sectionEnd = len(text)
	} else {
		sectionEnd = idxSec + len(header) + idxNext
	}
	sectionText := text[idxSec:sectionEnd]
	lines := strings.Split(sectionText, "\n")
	headerIdx, sepIdx := -1, -1
	rowIdxs := []int{}
	for i, ln := range lines {
		t := strings.TrimSpace(ln)
		if !strings.HasPrefix(t, "|") {
			continue
		}
		if headerIdx < 0 {
			headerIdx = i
			continue
		}
		if sepIdx < 0 && strings.Contains(t, "---") {
			sepIdx = i
			continue
		}
		if strings.Contains(t, "_vacío") {
			continue
		}
		rowIdxs = append(rowIdxs, i)
	}
	if headerIdx < 0 || sepIdx < 0 {
		return
	}

	sources := joinSources(ct.SourcesUsed)
	ts := time.Now().UTC().Format(time.RFC3339)
	tradeID := ct.ID
	if tradeID == "" {
		tradeID = ct.Slug
	}
	row := fmt.Sprintf("| %s | %s | %s | %.4f | %.4f | %.2f | %s | %s | %.2f | %s | %.2f |",
		ts, safeMd(tradeID), safeMd(ct.Category), ct.EntryPrice, ct.ExitPrice, ct.Pnl,
		safeMd(ct.Reason), safeMd(sources), ct.Size, safeMd(ct.Horizon), ct.DaysOpen)

	existing := []string{}
	for _, i := range rowIdxs {
		existing = append(existing, lines[i])
	}
	existing = append(existing, row)
	if len(existing) > maxRowsKeep {
		existing = existing[len(existing)-maxRowsKeep:]
	}

	prefix := strings.Join(lines[:sepIdx+1], "\n")
	rebuilt := prefix + "\n" + strings.Join(existing, "\n") + "\n"
	newText := text[:idxSec] + rebuilt + text[sectionEnd:]

	tmp := memoryPath + ".tmp"
	if err := os.WriteFile(tmp, []byte(newText), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "[loglosses] write tmp: %v\n", err)
		return
	}
	if err := os.Rename(tmp, memoryPath); err != nil {
		fmt.Fprintf(os.Stderr, "[loglosses] rename: %v\n", err)
	}
}

func joinSources(ss []commontypes.SourceCite) string {
	if len(ss) == 0 {
		return ""
	}
	domains := make([]string, 0, len(ss))
	seen := map[string]bool{}
	for _, s := range ss {
		d := strings.TrimSpace(s.Domain)
		if d == "" || seen[d] {
			continue
		}
		seen[d] = true
		domains = append(domains, d)
	}
	return strings.Join(domains, ",")
}

func sanitize(s string) string {
	s = strings.ToLower(s)
	out := make([]rune, 0, len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '-', r == '_':
			out = append(out, r)
		case r == ' ':
			out = append(out, '-')
		}
	}
	return strings.Trim(string(out), "-")
}

func yamlEscape(s string) string {
	s = strings.ReplaceAll(s, "\"", "'")
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}

func safeMd(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > 120 {
		s = s[:117] + "..."
	}
	return s
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
