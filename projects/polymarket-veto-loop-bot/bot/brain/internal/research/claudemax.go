// claudemax.go — research client fallback that uses Claude Max OAuth
// (via dashboard /api/llm/call provider=claudemax) when the Tavily connector
// is missing. Claude Max performs its own web search instead of consuming a
// separate Tavily key, so N1/N2 vetoes can fire even with no external news API.
//
// Respects the hard rule jarvis_never_anthropic_api — the dashboard bridge
// shells out to `docker exec gateway claude -p`, which uses the subscription's
// OAuth, not an API key.
package research

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
)

// extractHostname parses a URL and returns the lowercased host without "www.".
// Returns empty string when the input is not a parseable absolute URL.
func extractHostname(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return ""
	}
	h := strings.ToLower(u.Host)
	return strings.TrimPrefix(h, "www.")
}

const (
	claudemaxModel        = "claudemax/claude-sonnet-4-6"
	claudemaxSystemPrompt = `Eres un analista de mercados de predicción polymarket. Recibes una tesis y un lado (yes/no).

Tu tarea: usar WebSearch para verificar si las noticias RECIENTES (últimos 7 días) confirman, contradicen o son silentes respecto a la tesis.

Devuelves EXCLUSIVAMENTE un JSON válido (sin markdown, sin code-fence):
{
  "confirms": false,
  "contradicts": false,
  "silent": true,
  "score": 0.5,
  "summary": "una frase breve en español",
  "cited_dates": [
    {"headline_title": "...", "date": "YYYY-MM-DD", "url": "https://www.reuters.com/..."}
  ]
}

Reglas:
- Exactamente UNO de {confirms, contradicts, silent} debe ser true.
- score ∈ [0, 1] = tu confianza.
- summary <= 120 chars.
- cited_dates obligatorio si no es silent (al menos 1 entry).
- url EN CADA cited_dates obligatorio cuando no es silent — apunta a la página real (Reuters/AP/Bloomberg/NYT/...). Sin url ficticio: si no recuerdas la URL exacta, omite ese entry, no inventes.
- NO inventes hechos ni fechas. Si dudas → silent.
- Para mercados deportivos same-day: silent salvo lesiones/cambios alineación oficiales.
- Para geopolítica: contradicts si hay anuncio oficial opuesto a la tesis.`
)

// NewClaudemaxClient builds a research client that uses Claude Max (no Tavily).
// EvaluateFull will skip the Tavily search and ask Claude to do its own search.
func NewClaudemaxClient(dashboardURL, cachePath string) *Client {
	c := &Client{
		TavilyAPIKey: "", // empty signals claudemax-only path
		DashboardURL: dashboardURL,
		CachePath:    cachePath,
		CacheTTL:     defaultTTL,
		HTTPTimeout:  90 * time.Second, // Claude Max with web search needs longer
		cache:        map[string]CacheEntry{},
	}
	c.loadCache()
	return c
}

// evaluateFullClaudemax is the no-Tavily code path. Called from EvaluateFull
// when c.TavilyAPIKey == "".
func (c *Client) evaluateFullClaudemax(question, side string) (Verdict, []Headline, []commontypes.SourceCite) {
	query := question
	hash := hashQuery(query)
	if cached, ok := c.lookupCache(hash); ok {
		return cached.Verdict, cached.Headlines, citesFromVerdict(cached.Verdict, cached.Headlines)
	}

	v, err := c.synthesizeClaudemax(question, side)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[research] claudemax error for %q: %v\n", query, err)
		v = Verdict{Silent: true, Score: 0, Summary: "claudemax error"}
	}

	// v7 P2: build Headlines with real URL + host when the LLM provided one,
	// otherwise fall back to the synthetic pseudo-domain. citesFromVerdict
	// matches by title, so we keep the existing pipeline but emit accurate
	// Domain/URL when available.
	headlines := make([]Headline, 0, len(v.CitedDates))
	for _, cd := range v.CitedDates {
		h := Headline{
			Title:         cd.HeadlineTitle,
			URL:           cd.URL,
			PublishedDate: cd.Date,
		}
		if host := extractHostname(cd.URL); host != "" {
			h.Source = host
		} else {
			h.Source = "claudemax-websearch"
		}
		headlines = append(headlines, h)
	}
	c.appendCache(CacheEntry{
		Hash: hash, Query: query, Timestamp: time.Now().UTC().Format(time.RFC3339),
		Headlines: headlines, Verdict: v, Model: claudemaxModel,
	})
	cites := citesFromVerdict(v, headlines)
	// Mark cites with no resolvable URL as Synthetic so downstream filters
	// (closer.AppendSourceStats) can drop them from the domain blacklist.
	for i := range cites {
		if cites[i].URL == "" || cites[i].Domain == "claudemax-websearch" {
			cites[i].Synthetic = true
		}
	}
	return v, headlines, cites
}

func (c *Client) synthesizeClaudemax(question, side string) (Verdict, error) {
	today := time.Now().UTC().Format("2006-01-02")
	user := fmt.Sprintf("Hoy es %s.\nTesis: %s\nLado: %s\n\nUsa WebSearch para verificar y devuelve el JSON.", today, question, side)
	body := map[string]interface{}{
		"model":      claudemaxModel,
		"system":     claudemaxSystemPrompt,
		"messages":   []map[string]interface{}{{"role": "user", "content": user}},
		"max_tokens": 600,
		"tools":      []string{"web_search"},
	}
	raw, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", c.DashboardURL+"/api/llm/call", bytes.NewReader(raw))
	if err != nil {
		return Verdict{}, fmt.Errorf("build req: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: c.HTTPTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return Verdict{}, fmt.Errorf("post llm/call: %w", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return Verdict{}, fmt.Errorf("llm HTTP %d: %s", resp.StatusCode, truncate(string(respBody), 300))
	}
	text := extractTextFromLLMResp(respBody)
	if text == "" {
		return Verdict{Silent: true, Score: 0, Summary: "empty claudemax output"}, nil
	}
	jstart := strings.Index(text, "{")
	jend := strings.LastIndex(text, "}")
	if jstart < 0 || jend <= jstart {
		return Verdict{Silent: true, Score: 0, Summary: "no JSON in claudemax output"}, nil
	}
	var v Verdict
	if err := json.Unmarshal([]byte(text[jstart:jend+1]), &v); err != nil {
		return Verdict{Silent: true, Score: 0, Summary: "JSON parse failed"}, nil
	}
	// enforce exclusivity
	if (v.Confirms && v.Contradicts) || (v.Confirms && v.Silent) || (v.Contradicts && v.Silent) {
		v = Verdict{Silent: true, Score: 0.3, Summary: "model inconsistent"}
	}
	if !v.Confirms && !v.Contradicts && !v.Silent {
		v.Silent = true
	}
	if v.Score == 0 {
		v.Score = 0.5
	}
	if len(v.Summary) > 120 {
		v.Summary = v.Summary[:117] + "..."
	}
	return v, nil
}

