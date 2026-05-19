package decisionlog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// VetoDecision describes a brain veto worth persisting as an Obsidian note.
type VetoDecision struct {
	Slug          string
	Question      string
	Category      string
	PriceYes      float64
	Volume24h     float64
	EndDate       string
	Rule          string   // V1, V2, V4, M1, N1, N2, V3, V5, V6
	Reason        string   // human-readable
	ResearchSummary string // optional: news synthesis result
	MemoryHits    []string // optional: pattern matches that contributed
	Thesis        string   // optional original thesis if available
}

// LossDecision describes a closed trade that ended in red (PnL < 0).
type LossDecision struct {
	TradeID       string
	Slug          string
	Question      string
	Category      string
	EntryPrice    float64
	ExitPrice     float64
	Size          float64
	Pnl           float64
	EntryTime     string
	ExitTime      string
	ExitReason    string
}

// WriteVeto creates 03-decisions/<date>-polymarket-veto-<slug>.md.
// `baseDir` should point to vault/03-decisions/.
// Returns the path written.
func WriteVeto(baseDir string, d VetoDecision) (string, error) {
	if d.Slug == "" {
		return "", fmt.Errorf("slug required")
	}
	date := time.Now().UTC().Format("2006-01-02")
	slug := sanitize(d.Slug)
	if len(slug) > 80 {
		slug = slug[:80]
	}
	name := fmt.Sprintf("%s-polymarket-veto-%s.md", date, slug)
	path := filepath.Join(baseDir, name)

	frontmatter := buildVetoFrontmatter(d, date)
	body := buildVetoBody(d)
	full := frontmatter + "\n" + body
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, []byte(full), 0644); err != nil {
		return "", err
	}
	return path, nil
}

// WriteLoss creates 03-decisions/<date>-polymarket-loss-<slug>.md.
func WriteLoss(baseDir string, d LossDecision) (string, error) {
	if d.Slug == "" {
		return "", fmt.Errorf("slug required")
	}
	date := time.Now().UTC().Format("2006-01-02")
	slug := sanitize(d.Slug)
	if len(slug) > 80 {
		slug = slug[:80]
	}
	name := fmt.Sprintf("%s-polymarket-loss-%s.md", date, slug)
	path := filepath.Join(baseDir, name)

	frontmatter := buildLossFrontmatter(d, date)
	body := buildLossBody(d)
	full := frontmatter + "\n" + body
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, []byte(full), 0644); err != nil {
		return "", err
	}
	return path, nil
}

func buildVetoFrontmatter(d VetoDecision, date string) string {
	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"Polymarket veto — %s\"\n", yamlEscape(d.Slug))
	b.WriteString("type: decision\n")
	fmt.Fprintf(&b, "date: %s\n", date)
	fmt.Fprintf(&b, "decision: \"Veto de tesis '%s' (rule %s): %s\"\n",
		yamlEscape(truncate(d.Question, 80)), d.Rule, yamlEscape(truncate(d.Reason, 80)))
	b.WriteString("alternatives:\n")
	b.WriteString("  - \"Aprobar tesis y entrar trade simulado\"\n")
	b.WriteString("  - \"Vetar y mantener bankroll\"\n")
	b.WriteString("outcome: pending\n")
	b.WriteString("outcome_observed_after_days: 30\n")
	b.WriteString("tags: [decision, polymarket, bot, veto, ")
	b.WriteString(strings.ToLower(d.Rule))
	b.WriteString("]\n")
	b.WriteString("related:\n")
	b.WriteString("  - \"[[agents/polymarket-bot]]\"\n")
	b.WriteString("  - \"[[agents/polymarket-bot/memory]]\"\n")
	b.WriteString("  - \"[[projects/polymarket-veto-loop-bot]]\"\n")
	b.WriteString("---\n")
	return b.String()
}

func buildVetoBody(d VetoDecision) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Veto: %s\n\n", d.Question)
	b.WriteString("## Mercado\n\n")
	fmt.Fprintf(&b, "- **Slug**: `%s`\n", d.Slug)
	fmt.Fprintf(&b, "- **Categoría**: %s\n", d.Category)
	fmt.Fprintf(&b, "- **Precio YES**: %.4f\n", d.PriceYes)
	fmt.Fprintf(&b, "- **Volumen 24h**: %.2f USD\n", d.Volume24h)
	fmt.Fprintf(&b, "- **End date**: %s\n", d.EndDate)
	b.WriteString("\n")
	b.WriteString("## Razón del veto\n\n")
	fmt.Fprintf(&b, "**Regla aplicada**: `%s`\n\n", d.Rule)
	fmt.Fprintf(&b, "%s\n\n", d.Reason)

	if d.Thesis != "" {
		b.WriteString("## Tesis original\n\n")
		fmt.Fprintf(&b, "%s\n\n", d.Thesis)
	}

	if d.ResearchSummary != "" {
		b.WriteString("## Investigación de noticias\n\n")
		fmt.Fprintf(&b, "%s\n\n", d.ResearchSummary)
	}

	if len(d.MemoryHits) > 0 {
		b.WriteString("## Patterns en memoria que contribuyeron\n\n")
		for _, m := range d.MemoryHits {
			fmt.Fprintf(&b, "- %s\n", m)
		}
		b.WriteString("\n")
	}

	b.WriteString("## Human notes\n\n_(no se toca por automatización)_\n")
	return b.String()
}

func buildLossFrontmatter(d LossDecision, date string) string {
	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"Polymarket loss — %s\"\n", yamlEscape(d.Slug))
	b.WriteString("type: decision\n")
	fmt.Fprintf(&b, "date: %s\n", date)
	fmt.Fprintf(&b, "decision: \"Trade cerrado en pérdida: %s (PnL %.2f USD)\"\n",
		yamlEscape(truncate(d.Question, 80)), d.Pnl)
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
	b.WriteString("---\n")
	return b.String()
}

func buildLossBody(d LossDecision) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Loss: %s\n\n", d.Question)
	b.WriteString("## Trade\n\n")
	fmt.Fprintf(&b, "- **Trade ID**: `%s`\n", d.TradeID)
	fmt.Fprintf(&b, "- **Slug**: `%s`\n", d.Slug)
	fmt.Fprintf(&b, "- **Categoría**: %s\n", d.Category)
	fmt.Fprintf(&b, "- **Entry price**: %.4f\n", d.EntryPrice)
	fmt.Fprintf(&b, "- **Exit price**: %.4f\n", d.ExitPrice)
	fmt.Fprintf(&b, "- **Size**: $%.2f\n", d.Size)
	fmt.Fprintf(&b, "- **PnL**: $%.2f\n", d.Pnl)
	fmt.Fprintf(&b, "- **Entry timestamp**: %s\n", d.EntryTime)
	fmt.Fprintf(&b, "- **Exit timestamp**: %s\n", d.ExitTime)
	fmt.Fprintf(&b, "- **Exit reason**: %s\n", d.ExitReason)
	b.WriteString("\n")
	b.WriteString("## Análisis\n\n")
	b.WriteString("_(autogenerado tras cierre — revisar manualmente para extraer lección)_\n\n")
	b.WriteString("## Human notes\n\n_(no se toca por automatización)_\n")
	return b.String()
}

func sanitize(s string) string {
	s = strings.ToLower(s)
	out := make([]rune, 0, len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			out = append(out, r)
		case r >= '0' && r <= '9':
			out = append(out, r)
		case r == '-' || r == '_':
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

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
