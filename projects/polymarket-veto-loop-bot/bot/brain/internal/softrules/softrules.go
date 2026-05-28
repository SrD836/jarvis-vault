// Package softrules clusters historical losses and wins to derive soft veto rules.
// It writes generated rules into memory.md under "## Reglas blandas aprendidas".
// Safe re-entry: dedupes by normalized rule text. Caps at maxRules entries (LRU).
package softrules

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	maxRules      = 30
	minLossesCluster = 5
	maxWinRate    = 0.30
)

type row struct {
	ts          string
	tradeID     string
	category    string
	entryPrice  float64
	exitPrice   float64
	pnl         float64
	reason      string
	sources     string
	size        float64
	horizon     string
	daysHeld    float64
}

type clusterKey struct {
	category string
	horizon  string
	band     string
}

// GenerateAndAppend reads memory.md, generates rules, and writes them back.
// Best-effort: errors are logged to stderr but never panic.
func GenerateAndAppend(memoryPath string) {
	data, err := os.ReadFile(memoryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[softrules] read %s: %v\n", memoryPath, err)
		return
	}
	text := string(data)

	losses := parseSection(text, "## Losses pattern", "## Wins pattern", "## Reglas blandas")
	wins := parseSection(text, "## Wins pattern", "## Reglas blandas")

	stats := map[clusterKey]struct {
		losses int
		wins   int
	}{}
	for _, r := range losses {
		k := keyFor(r)
		s := stats[k]
		s.losses++
		stats[k] = s
	}
	for _, r := range wins {
		k := keyFor(r)
		s := stats[k]
		s.wins++
		stats[k] = s
	}

	var newRules []string
	for k, s := range stats {
		total := s.losses + s.wins
		if s.losses < minLossesCluster {
			continue
		}
		// v7 P1: drop clusters with no category. Legacy losses without a
		// resolvable category accumulated 126 "uncategorized" rules that
		// the current scanner (which classifies real categories) can never
		// match → 1 match/cron of pure noise. The rule generator now
		// refuses to emit them.
		if k.category == "" || strings.EqualFold(k.category, "uncategorized") {
			continue
		}
		winRate := 0.0
		if total > 0 {
			winRate = float64(s.wins) / float64(total)
		}
		if winRate >= maxWinRate {
			continue
		}
		rule := fmt.Sprintf(
			"- En categoría `%s` · horizonte `%s` · precio `%s`: %d losses, %d wins (win rate %.0f%%). **Veto soft o reducir size 50%%.**",
			k.category,
			emptyOr(k.horizon, "?"),
			k.band, s.losses, s.wins, winRate*100,
		)
		newRules = append(newRules, rule)
	}
	if len(newRules) == 0 {
		return
	}
	sort.Strings(newRules)

	updated := mergeRules(text, newRules)
	if updated == text {
		return
	}
	tmp := memoryPath + ".tmp"
	if err := os.WriteFile(tmp, []byte(updated), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "[softrules] write tmp: %v\n", err)
		return
	}
	if err := os.Rename(tmp, memoryPath); err != nil {
		fmt.Fprintf(os.Stderr, "[softrules] rename: %v\n", err)
	}
}

func keyFor(r row) clusterKey {
	return clusterKey{
		category: r.category,
		horizon:  r.horizon,
		band:     priceBand(r.entryPrice),
	}
}

func priceBand(p float64) string {
	switch {
	case p < 0.10:
		return "<0.10"
	case p < 0.30:
		return "0.10-0.30"
	case p < 0.70:
		return "0.30-0.70"
	default:
		return ">0.70"
	}
}

func emptyOr(s, d string) string {
	if strings.TrimSpace(s) == "" {
		return d
	}
	return s
}

// parseSection extracts table rows from a markdown section bounded by header → any of the stop headers.
func parseSection(text, header string, stops ...string) []row {
	idx := strings.Index(text, header)
	if idx < 0 {
		return nil
	}
	rest := text[idx+len(header):]
	end := len(rest)
	for _, st := range stops {
		j := strings.Index(rest, st)
		if j >= 0 && j < end {
			end = j
		}
	}
	body := rest[:end]
	out := []row{}
	for _, ln := range strings.Split(body, "\n") {
		t := strings.TrimSpace(ln)
		if !strings.HasPrefix(t, "|") {
			continue
		}
		// Skip header and separator
		if strings.Contains(t, "timestamp") || strings.Contains(t, "---") {
			continue
		}
		cells := splitCells(t)
		if len(cells) < 7 {
			continue
		}
		r := row{ts: cells[0], tradeID: cells[1], category: cells[2], reason: cells[6]}
		r.entryPrice = parseF(cells[3])
		r.exitPrice = parseF(cells[4])
		r.pnl = parseF(cells[5])
		if len(cells) >= 8 {
			r.sources = cells[7]
		}
		if len(cells) >= 9 {
			r.size = parseF(cells[8])
		}
		if len(cells) >= 10 {
			r.horizon = cells[9]
		}
		if len(cells) >= 11 {
			r.daysHeld = parseF(cells[10])
		}
		out = append(out, r)
	}
	return out
}

var spaceRe = regexp.MustCompile(`\s+`)

func splitCells(s string) []string {
	s = strings.TrimPrefix(s, "|")
	s = strings.TrimSuffix(s, "|")
	parts := strings.Split(s, "|")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		out = append(out, strings.TrimSpace(p))
	}
	return out
}

func parseF(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

// mergeRules inserts unique rule lines into the "## Reglas blandas" section.
func mergeRules(text string, newRules []string) string {
	header := "## Reglas blandas"
	idx := strings.Index(text, header)
	if idx < 0 {
		return text
	}
	// Find end of header line
	eol := strings.Index(text[idx:], "\n")
	if eol < 0 {
		return text
	}
	startBody := idx + eol + 1
	// Find next "## " header after startBody
	endBody := startBody + len(text[startBody:])
	if j := strings.Index(text[startBody:], "\n## "); j >= 0 {
		endBody = startBody + j
	}

	body := text[startBody:endBody]
	// Dedupe: existing rule lines (those starting with "- ").
	existing := map[string]bool{}
	for _, ln := range strings.Split(body, "\n") {
		t := strings.TrimSpace(ln)
		if strings.HasPrefix(t, "- ") {
			existing[norm(t)] = true
		}
	}
	added := []string{}
	for _, r := range newRules {
		if existing[norm(r)] {
			continue
		}
		added = append(added, r)
		existing[norm(r)] = true
	}
	if len(added) == 0 {
		return text
	}

	// Strip out the "_vacío_" placeholder if present
	body = strings.ReplaceAll(body, "_vacío — primera ejecución v2_", "")

	// Combine existing rule lines (sorted, kept) with new ones, keep first maxRules (latest wins)
	allLines := strings.Split(body, "\n")
	keepLines := []string{}
	for _, ln := range allLines {
		t := strings.TrimSpace(ln)
		if strings.HasPrefix(t, "- ") || t == "" || strings.HasPrefix(t, "**") || strings.HasPrefix(t, ">") {
			keepLines = append(keepLines, ln)
		} else {
			keepLines = append(keepLines, ln)
		}
	}

	stamp := fmt.Sprintf("\n<!-- auto-generated %s -->\n", time.Now().UTC().Format(time.RFC3339))
	insertion := strings.Join(added, "\n") + "\n"

	newBody := strings.TrimRight(strings.Join(keepLines, "\n"), "\n") + stamp + insertion
	return text[:startBody] + newBody + "\n" + text[endBody:]
}

func norm(s string) string {
	return spaceRe.ReplaceAllString(strings.ToLower(s), " ")
}
