// Package postmortem — async DeepSeek post-mortem analysis on a closed loss.
//
// Replaces the hardcoded "## Análisis _(revisar manualmente)_" placeholder
// in the loss .md with a one-line lesson + 3-5 anti-pattern tags extracted
// by DeepSeek from the trade context.
//
// Also appends a row to memory.md under "## Anti-patterns identificados"
// which the brain reads on the next cycle (M3 rule) to veto future
// candidates matching the same anti-pattern.
package postmortem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
)

const (
	dashboardCall = "/api/llm/call"
	// DeepSeek, NOT claudemax: this post-mortem fires from the exit_monitor cron
	// (*/5), so it must never use Claude Max — automated Max usage is the ban risk
	// the model-routing guardrail eliminates. Lesson+tag extraction needs no frontier.
	postmortemModel = "deepseek/deepseek-chat"
	systemPrompt    = `Eres un analista de post-mortem de trading. Recibes una operación PERDEDORA en polymarket y extraes una lección concisa + tags de anti-pattern que servirán para vetar futuros candidatos similares.

Devuelves EXCLUSIVAMENTE un JSON válido (sin markdown, sin code-fence):
{
  "lesson": "una frase breve en español, max 200 chars",
  "anti_pattern_tags": ["tag-1", "tag-2", "tag-3"]
}

Reglas:
- 3-5 tags, kebab-case, en inglés, max 30 chars cada uno.
- Tags deben ser PATRONES textuales reusables, no detalles específicos del trade.
  Buenos ejemplos: "ceasefire-rumor", "low-liquidity-pump", "election-far-noise", "same-day-chalk-bet", "geopolitics-short-horizon".
  Malos ejemplos: "iran-may-31", "trump-resign-june", "ucl-final-2026".
- lesson en español, accionable: "evitar X cuando Y".
- NO inventes hechos.`
)

var apMu sync.Mutex

type Output struct {
	Lesson          string   `json:"lesson"`
	AntiPatternTags []string `json:"anti_pattern_tags"`
}

// Run extracts the lesson + tags via DeepSeek, then:
//  1. Rewrites the loss .md replacing the Análisis placeholder.
//  2. Appends a row to memory.md "## Anti-patterns identificados".
//
// Best-effort: any error is logged to stderr and swallowed (we don't want
// post-mortem failures to corrupt the loss .md or stall exit_monitor).
func Run(dashboardURL, lossMdPath, memoryPath string, ct LossContext) {
	out, err := callPostmortemLLM(dashboardURL, ct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[postmortem] llm error: %v\n", err)
		return
	}
	if out == nil || (out.Lesson == "" && len(out.AntiPatternTags) == 0) {
		return
	}
	rewriteAnalysis(lossMdPath, out)
	appendAntiPatterns(memoryPath, ct, out)
}

// LossContext is the trade summary handed to the LLM. Kept narrow so the
// prompt is cheap and focused.
type LossContext struct {
	Slug        string
	Question    string
	Category    string
	EntryPrice  float64
	ExitPrice   float64
	SizeUSD     float64
	PnlUSD      float64
	DaysHeld    float64
	Horizon     string
	ExitReason  string
	SourcesUsed []commontypes.SourceCite
	Date        string
}

func callPostmortemLLM(dashboardURL string, ct LossContext) (*Output, error) {
	srcLines := ""
	for _, s := range ct.SourcesUsed {
		srcLines += fmt.Sprintf("- %s · %s · %s\n", s.Domain, s.PublishedDate, s.HeadlineTitle)
	}
	user := fmt.Sprintf(`Operación PERDEDORA en polymarket:

- Pregunta: %s
- Categoría: %s · Horizonte: %s
- Entry: %.4f · Exit: %.4f · Size: $%.2f
- PnL: $%.2f · Days held: %.2f
- Exit reason: %s

Fuentes consultadas al aprobar:
%s

Extrae lesson + 3-5 anti-pattern tags (kebab-case, reusables).`,
		ct.Question, ct.Category, ct.Horizon,
		ct.EntryPrice, ct.ExitPrice, ct.SizeUSD,
		ct.PnlUSD, ct.DaysHeld, ct.ExitReason,
		strings.TrimSpace(srcLines))

	body := map[string]interface{}{
		"model":      postmortemModel,
		"system":     systemPrompt,
		"messages":   []map[string]interface{}{{"role": "user", "content": user}},
		"max_tokens": 400,
	}
	raw, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", dashboardURL+dashboardCall, bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncate(string(respBody), 200))
	}
	text := extractText(respBody)
	jstart := strings.Index(text, "{")
	jend := strings.LastIndex(text, "}")
	if jstart < 0 || jend <= jstart {
		return nil, fmt.Errorf("no JSON in response")
	}
	var out Output
	if err := json.Unmarshal([]byte(text[jstart:jend+1]), &out); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	// normalise tags
	clean := make([]string, 0, len(out.AntiPatternTags))
	for _, t := range out.AntiPatternTags {
		t = strings.ToLower(strings.TrimSpace(t))
		t = strings.ReplaceAll(t, " ", "-")
		if len(t) > 0 && len(t) <= 32 {
			clean = append(clean, t)
		}
	}
	out.AntiPatternTags = clean
	if len(out.Lesson) > 220 {
		out.Lesson = out.Lesson[:217] + "..."
	}
	return &out, nil
}

func rewriteAnalysis(mdPath string, out *Output) {
	data, err := os.ReadFile(mdPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[postmortem] read %s: %v\n", mdPath, err)
		return
	}
	text := string(data)
	placeholder := "## Análisis\n\n_(autogenerado tras cierre — revisar manualmente para extraer lección)_\n\n"
	replacement := fmt.Sprintf("## Análisis\n\n%s\n\n**Anti-patterns**: %s\n\n",
		out.Lesson, strings.Join(out.AntiPatternTags, " · "))
	if !strings.Contains(text, placeholder) {
		// nothing to replace (file was already patched or shape changed)
		return
	}
	text = strings.Replace(text, placeholder, replacement, 1)
	if err := os.WriteFile(mdPath, []byte(text), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "[postmortem] write %s: %v\n", mdPath, err)
	}
}

func appendAntiPatterns(memoryPath string, ct LossContext, out *Output) {
	if memoryPath == "" || len(out.AntiPatternTags) == 0 {
		return
	}
	apMu.Lock()
	defer apMu.Unlock()
	data, err := os.ReadFile(memoryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[postmortem] read memory %s: %v\n", memoryPath, err)
		return
	}
	text := string(data)
	header := "## Anti-patterns identificados"
	if !strings.Contains(text, header) {
		// append at EOF
		text += fmt.Sprintf("\n\n%s\n\n_(autogenerada por exit_monitor tras cada pérdida — brain consume vía M3)_\n\n", header)
	}
	row := fmt.Sprintf("- %s — visto en %s (%s, pnl $%.2f)\n",
		strings.Join(out.AntiPatternTags, " · "),
		ct.Slug, ct.Date, ct.PnlUSD)
	// Insert immediately after the header (or before the next "## " block).
	idx := strings.Index(text, header)
	if idx < 0 {
		return
	}
	rest := text[idx+len(header):]
	// position right after first \n\n following header
	insAt := idx + len(header)
	if j := strings.Index(rest, "\n\n"); j >= 0 {
		insAt += j + 2
	} else {
		insAt += 1
	}
	text = text[:insAt] + row + text[insAt:]
	if err := os.WriteFile(memoryPath, []byte(text), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "[postmortem] write memory: %v\n", err)
	}
}

func extractText(raw []byte) string {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(raw, &m); err != nil {
		return string(raw)
	}
	if rc, ok := m["content"]; ok {
		var s string
		if err := json.Unmarshal(rc, &s); err == nil && s != "" {
			return s
		}
		var arr []map[string]interface{}
		if err := json.Unmarshal(rc, &arr); err == nil {
			for _, item := range arr {
				if t, ok := item["text"].(string); ok && t != "" {
					return t
				}
			}
		}
	}
	if rt, ok := m["text"]; ok {
		var s string
		if err := json.Unmarshal(rt, &s); err == nil && s != "" {
			return s
		}
	}
	return string(raw)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
