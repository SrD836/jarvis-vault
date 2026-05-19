package loglosses

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const maxRowsKeep = 100

var mu sync.Mutex

// LogLoss writes both a decision log markdown and a memory.md row
// for a trade that closed in pérdida. Best-effort: errors are logged via stderr.
func LogLoss(
	decisionsDir, memoryPath string,
	tradeID, slug, question, category string,
	entryPrice, exitPrice, size, pnl float64,
	entryTime, exitTime, exitReason string,
) {
	if pnl >= 0 {
		return // not a loss
	}
	if slug == "" {
		slug = tradeID
	}
	writeDecisionLog(decisionsDir, tradeID, slug, question, category,
		entryPrice, exitPrice, size, pnl, entryTime, exitTime, exitReason)
	appendMemoryLoss(memoryPath, slug, category, entryPrice, exitPrice, pnl, exitReason)
}

func writeDecisionLog(decisionsDir, tradeID, slug, question, category string,
	entryPrice, exitPrice, size, pnl float64, entryTime, exitTime, exitReason string,
) {
	date := time.Now().UTC().Format("2006-01-02")
	slugSan := sanitize(slug)
	if len(slugSan) > 80 {
		slugSan = slugSan[:80]
	}
	name := fmt.Sprintf("%s-polymarket-loss-%s.md", date, slugSan)
	path := filepath.Join(decisionsDir, name)

	if err := os.MkdirAll(decisionsDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "[loglosses] mkdir %s: %v\n", decisionsDir, err)
		return
	}

	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"Polymarket loss — %s\"\n", yamlEscape(slug))
	b.WriteString("type: decision\n")
	fmt.Fprintf(&b, "date: %s\n", date)
	fmt.Fprintf(&b, "decision: \"Trade cerrado en pérdida: %s (PnL %.2f USD)\"\n",
		yamlEscape(truncate(question, 80)), pnl)
	b.WriteString("alternatives:\n")
	b.WriteString("  - \"Haber vetado la tesis antes de entrar\"\n")
	b.WriteString("  - \"Cerrar antes (stop más conservador)\"\n")
	b.WriteString("outcome: confirmed_bad\n")
	b.WriteString("outcome_observed_after_days: 0\n")
	b.WriteString("tags: [decision, polymarket, bot, loss, post-mortem]\n")
	b.WriteString("related:\n")
	b.WriteString("  - \"[[agents/polymarket-bot]]\"\n")
	b.WriteString("  - \"[[agents/polymarket-bot/memory]]\"\n")
	b.WriteString("  - \"[[projects/polymarket-veto-loop-bot]]\"\n")
	b.WriteString("---\n\n")
	fmt.Fprintf(&b, "# Loss: %s\n\n", question)
	b.WriteString("## Trade\n\n")
	fmt.Fprintf(&b, "- **Trade ID**: `%s`\n", tradeID)
	fmt.Fprintf(&b, "- **Slug**: `%s`\n", slug)
	fmt.Fprintf(&b, "- **Categoría**: %s\n", category)
	fmt.Fprintf(&b, "- **Entry price**: %.4f\n", entryPrice)
	fmt.Fprintf(&b, "- **Exit price**: %.4f\n", exitPrice)
	fmt.Fprintf(&b, "- **Size**: $%.2f\n", size)
	fmt.Fprintf(&b, "- **PnL**: $%.2f\n", pnl)
	fmt.Fprintf(&b, "- **Entry timestamp**: %s\n", entryTime)
	fmt.Fprintf(&b, "- **Exit timestamp**: %s\n", exitTime)
	fmt.Fprintf(&b, "- **Exit reason**: %s\n\n", exitReason)
	b.WriteString("## Análisis\n\n_(autogenerado tras cierre — revisar manualmente para extraer lección)_\n\n")
	b.WriteString("## Human notes\n\n_(no se toca por automatización)_\n")

	if err := os.WriteFile(path, []byte(b.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "[loglosses] write %s: %v\n", path, err)
	}
}

func appendMemoryLoss(memoryPath, slug, category string, entry, exit, pnl float64, reason string) {
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
	header := "## Losses pattern"
	nextHeader := "## Reglas blandas"
	idxSec := strings.Index(text, header)
	if idxSec < 0 {
		return
	}
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

	ts := time.Now().UTC().Format(time.RFC3339)
	row := fmt.Sprintf("| %s | %s | %s | %.4f | %.4f | %.2f | %s |",
		ts, safeMd(slug), safeMd(category), entry, exit, pnl, safeMd(reason))

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
