package memory

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Pattern represents a past event (veto or loss) the bot should remember.
type Pattern struct {
	Timestamp string
	Slug      string
	Category  string
	PriceYes  float64
	Rule      string  // for vetos: V1, V2, M1, N1...  | for losses: free text reason
	Source    string  // "veto" | "loss"
	Raw       string  // original table row (for debugging)
	Pnl       float64 // only for losses
	Entry     float64 // only for losses
	Exit      float64 // only for losses
}

// Memory is the parsed state of agents/polymarket-bot/memory.md.
type Memory struct {
	Path      string
	Vetos     []Pattern
	Losses    []Pattern
	SoftRules []string
	mu        sync.Mutex
}

// Match represents a memory hit when evaluating a candidate.
type Match struct {
	Pattern Pattern
	Score   float64 // 0-1
	Why     string  // human-readable reason
}

const (
	// maxRowsKeep — append rotation cap per section. Raised from 100 to 500
	// (2026-05-26) so M2 soft-rule clustering sees five times the historical
	// sample once closed.jsonl has accumulated enough trades.
	maxRowsKeep = 500
)

// Load parses memory.md from disk. Returns empty Memory on file-not-found.
func Load(path string) (*Memory, error) {
	m := &Memory{Path: path}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return m, nil
		}
		return nil, fmt.Errorf("read memory %s: %w", path, err)
	}
	parseInto(m, string(data))
	return m, nil
}

func parseInto(m *Memory, text string) {
	section := ""
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		trim := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(trim, "## Vetos pattern"):
			section = "vetos"
			continue
		case strings.HasPrefix(trim, "## Losses pattern"):
			section = "losses"
			continue
		case strings.HasPrefix(trim, "## Reglas blandas"):
			section = "rules"
			continue
		case strings.HasPrefix(trim, "## Human notes") || strings.HasPrefix(trim, "## "):
			// any other H2 ends the parsing for known sections (until next match)
			if section == "rules" || section == "vetos" || section == "losses" {
				section = ""
			}
			continue
		}

		switch section {
		case "vetos":
			if p, ok := parseTableRow(trim, "veto"); ok {
				m.Vetos = append(m.Vetos, p)
			}
		case "losses":
			if p, ok := parseTableRow(trim, "loss"); ok {
				m.Losses = append(m.Losses, p)
			}
		case "rules":
			if strings.HasPrefix(trim, "- ") {
				m.SoftRules = append(m.SoftRules, strings.TrimPrefix(trim, "- "))
			}
		}
	}
}

// parseTableRow parses "| ts | slug | category | price | rule | reason |" rows.
// Skips header, separator, and empty placeholder rows.
func parseTableRow(line, kind string) (Pattern, bool) {
	if !strings.HasPrefix(line, "|") {
		return Pattern{}, false
	}
	// skip header and separator
	if strings.Contains(line, "---") || strings.Contains(line, "timestamp") {
		return Pattern{}, false
	}
	cols := splitTableCols(line)
	if len(cols) < 4 {
		return Pattern{}, false
	}
	// skip placeholder/empty rows
	if strings.HasPrefix(cols[0], "_vacío") || cols[0] == "" {
		return Pattern{}, false
	}
	p := Pattern{
		Timestamp: cols[0],
		Slug:      cols[1],
		Category:  cols[2],
		Source:    kind,
		Raw:       line,
	}
	if kind == "veto" && len(cols) >= 6 {
		// | ts | slug | category | price_yes | rule | reason |
		fmt.Sscanf(cols[3], "%f", &p.PriceYes)
		p.Rule = cols[4]
	}
	if kind == "loss" && len(cols) >= 7 {
		// | ts | slug | category | entry | exit | pnl | reason |
		fmt.Sscanf(cols[3], "%f", &p.Entry)
		fmt.Sscanf(cols[4], "%f", &p.Exit)
		fmt.Sscanf(cols[5], "%f", &p.Pnl)
		p.Rule = cols[6]
	}
	return p, true
}

func splitTableCols(line string) []string {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "|")
	line = strings.TrimSuffix(line, "|")
	parts := strings.Split(line, "|")
	out := make([]string, len(parts))
	for i, p := range parts {
		out[i] = strings.TrimSpace(p)
	}
	return out
}

// CandidateLike is the minimum interface to score patterns against a candidate.
// The brain's types.Candidate satisfies this implicitly.
type CandidateLike struct {
	Slug            string
	Category        string
	CurrentPriceYes float64
	Question        string
}

// Match returns the strongest pattern hits for a candidate. Sorted desc by score.
// Score is in [0,1]. Threshold 0.7 is the recommended veto trigger.
func (m *Memory) Match(c CandidateLike) []Match {
	if m == nil {
		return nil
	}
	var hits []Match
	all := append([]Pattern{}, m.Vetos...)
	all = append(all, m.Losses...)
	for _, p := range all {
		score, why := patternScore(p, c)
		if score >= 0.4 {
			hits = append(hits, Match{Pattern: p, Score: score, Why: why})
		}
	}
	// sort desc by score (simple selection sort, n small)
	for i := 0; i < len(hits); i++ {
		max := i
		for j := i + 1; j < len(hits); j++ {
			if hits[j].Score > hits[max].Score {
				max = j
			}
		}
		hits[i], hits[max] = hits[max], hits[i]
	}
	if len(hits) > 5 {
		hits = hits[:5]
	}
	return hits
}

// patternScore returns similarity in [0,1] between a past pattern and a candidate.
func patternScore(p Pattern, c CandidateLike) (float64, string) {
	score := 0.0
	whys := []string{}

	// exact slug match → very strong signal
	if p.Slug != "" && strings.EqualFold(p.Slug, c.Slug) {
		return 1.0, "exact slug match"
	}

	// slug prefix (first 4 segments) → strong
	if commonSlugPrefix(p.Slug, c.Slug) >= 4 {
		score += 0.5
		whys = append(whys, "slug prefix match")
	}

	// category match
	if p.Category != "" && p.Category != "uncategorized" &&
		strings.EqualFold(p.Category, c.Category) {
		score += 0.2
		whys = append(whys, "same category")
	}

	// price bucket: both <0.05, or both >0.95, or both in middle band
	if priceBucket(p.PriceYes) == priceBucket(c.CurrentPriceYes) && p.PriceYes > 0 {
		score += 0.2
		whys = append(whys, "same price bucket "+priceBucket(p.PriceYes))
	}

	// loss pattern is heavier than veto pattern (a real loss is stronger evidence)
	if p.Source == "loss" {
		score *= 1.15
	}
	if score > 1.0 {
		score = 1.0
	}
	if score < 0.4 {
		return score, ""
	}
	return score, strings.Join(whys, "; ")
}

func commonSlugPrefix(a, b string) int {
	if a == "" || b == "" {
		return 0
	}
	pa := strings.Split(a, "-")
	pb := strings.Split(b, "-")
	n := len(pa)
	if len(pb) < n {
		n = len(pb)
	}
	common := 0
	for i := 0; i < n; i++ {
		if pa[i] == pb[i] {
			common++
		} else {
			break
		}
	}
	return common
}

func priceBucket(p float64) string {
	switch {
	case p <= 0:
		return "unknown"
	case p < 0.05:
		return "long-tail-low"
	case p < 0.20:
		return "low"
	case p < 0.80:
		return "mid"
	case p < 0.95:
		return "high"
	default:
		return "long-tail-high"
	}
}

// AppendVeto appends a row to the "Vetos pattern" table in memory.md.
// Preserves frontmatter, Human notes, and other sections verbatim.
// Truncates to maxRowsKeep most recent vetos.
func (m *Memory) AppendVeto(slug, category, rule, reason string, priceYes float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	ts := time.Now().UTC().Format(time.RFC3339)
	row := fmt.Sprintf("| %s | %s | %s | %.4f | %s | %s |",
		ts,
		safeMarkdown(slug),
		safeMarkdown(category),
		priceYes,
		safeMarkdown(rule),
		safeMarkdown(reason),
	)
	return m.appendToSection("## Vetos pattern", "## Losses pattern", row)
}

// AppendLoss appends a row to the "Losses pattern" table in memory.md.
func (m *Memory) AppendLoss(slug, category string, entry, exit, pnl float64, reason string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	ts := time.Now().UTC().Format(time.RFC3339)
	row := fmt.Sprintf("| %s | %s | %s | %.4f | %.4f | %.2f | %s |",
		ts,
		safeMarkdown(slug),
		safeMarkdown(category),
		entry, exit, pnl,
		safeMarkdown(reason),
	)
	return m.appendToSection("## Losses pattern", "## Reglas blandas", row)
}

// appendToSection inserts `row` at the end of the markdown table that follows
// `sectionHeader`, just before `nextHeader`. Truncates to maxRowsKeep rows.
func (m *Memory) appendToSection(sectionHeader, nextHeader, row string) error {
	data, err := os.ReadFile(m.Path)
	if err != nil {
		return fmt.Errorf("read memory: %w", err)
	}
	text := string(data)
	idxSec := strings.Index(text, sectionHeader)
	if idxSec < 0 {
		return fmt.Errorf("section header not found: %s", sectionHeader)
	}
	// find nextHeader (or end of file)
	idxNext := strings.Index(text[idxSec+len(sectionHeader):], nextHeader)
	var sectionEnd int
	if idxNext < 0 {
		sectionEnd = len(text)
	} else {
		sectionEnd = idxSec + len(sectionHeader) + idxNext
	}
	sectionText := text[idxSec:sectionEnd]

	// split into lines, identify table rows
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
		// skip placeholder
		if strings.Contains(t, "_vacío") {
			continue
		}
		rowIdxs = append(rowIdxs, i)
	}
	if headerIdx < 0 || sepIdx < 0 {
		return fmt.Errorf("table header/separator not found in section %s", sectionHeader)
	}

	// truncate older rows if over cap
	existing := []string{}
	for _, i := range rowIdxs {
		existing = append(existing, lines[i])
	}
	existing = append(existing, row)
	if len(existing) > maxRowsKeep {
		existing = existing[len(existing)-maxRowsKeep:]
	}

	// rebuild section: header lines (frontmatter intro), table header, separator, rows, blank line
	prefix := strings.Join(lines[:sepIdx+1], "\n") // up to & including separator
	rebuilt := prefix + "\n" + strings.Join(existing, "\n") + "\n"

	// remove leftover placeholder rows that may have been inside the original
	// (already excluded above)

	// replace section in full text
	newText := text[:idxSec] + rebuilt + text[sectionEnd:]

	// write atomically (write tmp + rename)
	tmp := m.Path + ".tmp"
	if err := os.WriteFile(tmp, []byte(newText), 0644); err != nil {
		return fmt.Errorf("write tmp: %w", err)
	}
	if err := os.Rename(tmp, m.Path); err != nil {
		return fmt.Errorf("rename: %w", err)
	}
	return nil
}

// safeMarkdown escapes pipe chars and trims long fields for table cells.
func safeMarkdown(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > 120 {
		s = s[:117] + "..."
	}
	return s
}

// CountsSummary returns N vetos / N losses for the dashboard badge.
func (m *Memory) CountsSummary() (int, int) {
	if m == nil {
		return 0, 0
	}
	return len(m.Vetos), len(m.Losses)
}
