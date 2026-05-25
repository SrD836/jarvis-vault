package research

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
)

// Headline is a single news search result from Tavily.
type Headline struct {
	Title         string  `json:"title"`
	URL           string  `json:"url"`
	Content       string  `json:"content"` // snippet
	Score         float64 `json:"score"`
	PublishedDate string  `json:"published_date,omitempty"` // ISO date from Tavily (advanced search)
	Source        string  `json:"source,omitempty"`         // derived from URL host
}

// CitedDate is what the LLM asserts having based its verdict on.
type CitedDate struct {
	HeadlineTitle string `json:"headline_title"`
	Date          string `json:"date"`
}

// Verdict is the structured output of the LLM auditor synthesis.
type Verdict struct {
	Confirms    bool        `json:"confirms"`
	Contradicts bool        `json:"contradicts"`
	Silent      bool        `json:"silent"`
	Score       float64     `json:"score"`
	Summary     string      `json:"summary"`
	CitedDates  []CitedDate `json:"cited_dates,omitempty"` // pruebas temporales del veredicto
}

// CacheEntry is one row of research-cache.jsonl.
type CacheEntry struct {
	Hash      string     `json:"hash"`
	Query     string     `json:"query"`
	Timestamp string     `json:"ts"`
	Headlines []Headline `json:"top_headlines"`
	Verdict   Verdict    `json:"verdict"`
	Model     string     `json:"model_used"`
}

// Client wraps Tavily search + LLM synthesis with disk cache.
type Client struct {
	TavilyAPIKey       string
	DashboardURL       string // e.g. http://jarvis-dashboard:3000
	CachePath          string // vault/inbox/trading/research-cache.jsonl
	CacheTTL           time.Duration
	HTTPTimeout        time.Duration
	BlacklistedDomains []string // set externally per cycle; subtracted from include_domains
	cache              map[string]CacheEntry
	cacheMu            sync.Mutex
	loaded             bool
}

const (
	tavilyURL       = "https://api.tavily.com/search"
	defaultTTL      = 6 * time.Hour
	defaultTimeout  = 30 * time.Second
	maxHeadlines    = 10
	recencyDays     = 7
	llmModel        = "deepseek/deepseek-chat"
	systemPrompt    = `Eres un analista de mercados de predicción. Lees titulares de noticias RECIENTES (últimos 7 días) y decides si confirman, contradicen, o son silentes respecto a una tesis de trading.

Devuelves EXCLUSIVAMENTE un JSON válido con este shape (sin markdown, sin code-fence):

{
  "confirms": false,
  "contradicts": false,
  "silent": true,
  "score": 0.7,
  "summary": "Una frase <= 120 caracteres explicando el veredicto",
  "cited_dates": [
    {"headline_title": "titulo exacto del titular X", "date": "2026-05-18"}
  ]
}

Reglas DURAS (cualquier desviacion -> tu salida sera rechazada y se forzara silent):
1. Solo UNO de confirms/contradicts/silent puede ser true. Los otros dos son false.
2. SOLO considera titulares cuya "published_date" este dentro de los ultimos 7 dias desde hoy. Si todos son mas antiguos o sin fecha -> devuelve silent.
3. cited_dates DEBE incluir 1-3 titulares en los que basas tu veredicto, copiando su "published_date" tal cual. Si no hay headlines con fecha reciente, devuelve cited_dates: [] y silent: true.
4. "confirms" = la mayoria de titulares recientes apoya el desenlace YES de la tesis.
5. "contradicts" = la mayoria apunta al desenlace NO.
6. "silent" = sin cobertura mediatica reciente relevante.
7. score: 0=debil, 1=fuerte. Si solo hay 1 fuente, max 0.6.
8. summary: <= 120 chars, sin punto final, en espaniol, citando la fuente mas decisiva.
9. NO inventes hechos ni fechas. Si dudas -> silent.`
)

// NewClient builds a research client. Reads Tavily key from connector JSON.
func NewClient(connectorPath, dashboardURL, cachePath string) (*Client, error) {
	data, err := os.ReadFile(connectorPath)
	if err != nil {
		return nil, fmt.Errorf("read tavily connector: %w", err)
	}
	var conn struct {
		APIKey string `json:"api_key"`
	}
	if err := json.Unmarshal(data, &conn); err != nil {
		return nil, fmt.Errorf("parse tavily connector: %w", err)
	}
	if conn.APIKey == "" {
		return nil, fmt.Errorf("empty tavily api_key in %s", connectorPath)
	}
	c := &Client{
		TavilyAPIKey: conn.APIKey,
		DashboardURL: dashboardURL,
		CachePath:    cachePath,
		CacheTTL:     defaultTTL,
		HTTPTimeout:  defaultTimeout,
		cache:        map[string]CacheEntry{},
	}
	c.loadCache()
	return c, nil
}

// Evaluate runs the full pipeline: cache lookup → Tavily → LLM synthesis → cache append.
// Returns the verdict. On any error returns a Silent verdict with score=0 (defensive fallback).
func (c *Client) Evaluate(question, side string) Verdict {
	if c == nil {
		return Verdict{Silent: true, Summary: "research client nil"}
	}
	query := question
	hash := hashQuery(query)

	if cached, ok := c.lookupCache(hash); ok {
		return cached.Verdict
	}

	headlines, err := c.tavilySearch(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[research] tavily error for %q: %v\n", query, err)
		return Verdict{Silent: true, Score: 0, Summary: "tavily error"}
	}
	if len(headlines) == 0 {
		v := Verdict{Silent: true, Score: 0.5, Summary: "sin resultados Tavily"}
		c.appendCache(CacheEntry{Hash: hash, Query: query, Timestamp: time.Now().UTC().Format(time.RFC3339),
			Headlines: headlines, Verdict: v, Model: "none"})
		return v
	}

	verdict, err := c.synthesize(question, side, headlines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[research] synthesize error for %q: %v\n", query, err)
		verdict = Verdict{Silent: true, Score: 0, Summary: "synthesize error"}
	}
	c.appendCache(CacheEntry{Hash: hash, Query: query, Timestamp: time.Now().UTC().Format(time.RFC3339),
		Headlines: headlines, Verdict: verdict, Model: llmModel})
	return verdict
}

func (c *Client) tavilySearch(query string) ([]Headline, error) {
	body := map[string]interface{}{
		"api_key":        c.TavilyAPIKey,
		"query":          query,
		"max_results":    maxHeadlines,
		"search_depth":   "advanced",
		"include_answer": true,
		"days":           recencyDays,
		"topic":          "news",
		"include_domains": c.buildIncludeDomains(),
	}
	raw, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", tavilyURL, bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("build req: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: c.HTTPTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post tavily: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("tavily HTTP %d: %s", resp.StatusCode, truncate(string(body), 300))
	}
	var out struct {
		Results []Headline `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode tavily: %w", err)
	}
	for i := range out.Results {
		out.Results[i].Source = extractHost(out.Results[i].URL)
	}
	return out.Results, nil
}

func extractHost(u string) string {
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")
	u = strings.TrimPrefix(u, "www.")
	if i := strings.IndexByte(u, '/'); i >= 0 {
		u = u[:i]
	}
	return u
}

func (c *Client) synthesize(question, side string, headlines []Headline) (Verdict, error) {
	today := time.Now().UTC().Format("2006-01-02")
	userContent := fmt.Sprintf("Hoy es %s. Tesis: %s\nLado: %s\n\nTitulares (top %d, solo ultimos %d dias):\n",
		today, question, side, len(headlines), recencyDays)
	for i, h := range headlines {
		dateLbl := h.PublishedDate
		if dateLbl == "" {
			dateLbl = "sin fecha"
		}
		srcLbl := h.Source
		if srcLbl == "" {
			srcLbl = "sin fuente"
		}
		userContent += fmt.Sprintf("%d. [%s | %s] %s — %s\n",
			i+1, dateLbl, srcLbl, truncate(h.Title, 140), truncate(h.Content, 220))
	}
	userContent += "\nDevuelve el JSON con el veredicto. Recuerda: cited_dates obligatorio si no es silent."

	body := map[string]interface{}{
		"model":      llmModel,
		"system":     systemPrompt,
		"messages":   []map[string]interface{}{{"role": "user", "content": userContent}},
		"max_tokens": 400,
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

	// callLLM returns various shapes depending on provider. We look for content[0].text or messages[0].content[0].text.
	text := extractTextFromLLMResp(respBody)
	if text == "" {
		return Verdict{Silent: true, Score: 0, Summary: "empty LLM output"}, nil
	}
	// extract first JSON block
	jstart := strings.Index(text, "{")
	jend := strings.LastIndex(text, "}")
	if jstart < 0 || jend <= jstart {
		return Verdict{Silent: true, Score: 0, Summary: "no JSON in LLM output"}, nil
	}
	var v Verdict
	if err := json.Unmarshal([]byte(text[jstart:jend+1]), &v); err != nil {
		return Verdict{Silent: true, Score: 0, Summary: "JSON parse failed"}, nil
	}
	// enforce exclusive
	if (v.Confirms && v.Contradicts) || (v.Confirms && v.Silent) || (v.Contradicts && v.Silent) {
		// fallback to silent if model produced inconsistent state
		v = Verdict{Silent: true, Score: 0.3, Summary: "model gave inconsistent verdict"}
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

// extractTextFromLLMResp handles multiple provider response shapes.
func extractTextFromLLMResp(raw []byte) string {
	// shape A: { "content": [ { "text": "..." } ] }  (Anthropic-style)
	// shape B: { "messages": [ { "role": "assistant", "content": [...] } ] }
	// shape C: { "text": "..." }
	// shape D: callLLM's own wrap with { "content": "...", "model_used": "..." }
	var anyMap map[string]json.RawMessage
	if err := json.Unmarshal(raw, &anyMap); err != nil {
		return string(raw)
	}
	// D: top-level "content" as string
	if rawC, ok := anyMap["content"]; ok {
		// try string
		var s string
		if err := json.Unmarshal(rawC, &s); err == nil && s != "" {
			return s
		}
		// try array
		var arr []map[string]interface{}
		if err := json.Unmarshal(rawC, &arr); err == nil {
			for _, item := range arr {
				if t, ok := item["text"].(string); ok && t != "" {
					return t
				}
			}
		}
	}
	if rawT, ok := anyMap["text"]; ok {
		var s string
		if err := json.Unmarshal(rawT, &s); err == nil && s != "" {
			return s
		}
	}
	// C: messages
	if rawM, ok := anyMap["messages"]; ok {
		var msgs []map[string]interface{}
		if err := json.Unmarshal(rawM, &msgs); err == nil {
			for _, m := range msgs {
				if c, ok := m["content"].([]interface{}); ok {
					for _, it := range c {
						if mm, ok := it.(map[string]interface{}); ok {
							if t, ok := mm["text"].(string); ok && t != "" {
								return t
							}
						}
					}
				}
				if s, ok := m["content"].(string); ok && s != "" {
					return s
				}
			}
		}
	}
	return string(raw)
}

func (c *Client) lookupCache(hash string) (CacheEntry, bool) {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()
	e, ok := c.cache[hash]
	if !ok {
		return CacheEntry{}, false
	}
	t, err := time.Parse(time.RFC3339, e.Timestamp)
	if err != nil {
		return CacheEntry{}, false
	}
	if time.Since(t) > c.CacheTTL {
		return CacheEntry{}, false
	}
	return e, true
}

func (c *Client) loadCache() {
	if c.loaded || c.CachePath == "" {
		return
	}
	c.loaded = true
	data, err := os.ReadFile(c.CachePath)
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var e CacheEntry
		if err := json.Unmarshal([]byte(line), &e); err == nil && e.Hash != "" {
			c.cache[e.Hash] = e
		}
	}
}

func (c *Client) appendCache(e CacheEntry) {
	c.cacheMu.Lock()
	c.cache[e.Hash] = e
	c.cacheMu.Unlock()

	if c.CachePath == "" {
		return
	}
	dir := filepath.Dir(c.CachePath)
	_ = os.MkdirAll(dir, 0755)
	f, err := os.OpenFile(c.CachePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	raw, _ := json.Marshal(e)
	f.Write(raw)
	f.Write([]byte("\n"))
}

func hashQuery(q string) string {
	q = strings.ToLower(strings.TrimSpace(q))
	h := sha256.Sum256([]byte(q))
	return hex.EncodeToString(h[:8])
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}


// baseDomains is the canonical premier news set used for prediction-market research.
var baseDomains = []string{
	"reuters.com", "apnews.com", "bloomberg.com", "ft.com", "wsj.com",
	"x.com", "twitter.com", "nytimes.com", "bbc.com", "aljazeera.com",
	"theguardian.com", "cnn.com", "washingtonpost.com", "axios.com", "politico.com",
}

// buildIncludeDomains subtracts blacklisted sources from baseDomains.
// If all are blacklisted (pathological), returns baseDomains unchanged so research
// degrades gracefully instead of returning zero results.
func (c *Client) buildIncludeDomains() []string {
	if len(c.BlacklistedDomains) == 0 {
		return baseDomains
	}
	bl := map[string]bool{}
	for _, d := range c.BlacklistedDomains {
		bl[strings.ToLower(strings.TrimSpace(d))] = true
	}
	out := make([]string, 0, len(baseDomains))
	for _, d := range baseDomains {
		if !bl[d] {
			out = append(out, d)
		}
	}
	if len(out) == 0 {
		return baseDomains
	}
	return out
}


// EvaluateFull is like Evaluate but also returns the headlines used (for source attribution).
// Returns at most 3 SourceCite entries derived from verdict.CitedDates ∩ headlines (by title prefix match).
// If verdict is silent or has no cited_dates, returns an empty []SourceCite.
func (c *Client) EvaluateFull(question, side string) (Verdict, []Headline, []commontypes.SourceCite) {
	if c == nil {
		return Verdict{Silent: true, Summary: "research client nil"}, nil, nil
	}
	// No Tavily key → use Claude Max OAuth fallback (own WebSearch, no external API key).
	if c.TavilyAPIKey == "" {
		return c.evaluateFullClaudemax(question, side)
	}
	query := question
	hash := hashQuery(query)

	if cached, ok := c.lookupCache(hash); ok {
		return cached.Verdict, cached.Headlines, citesFromVerdict(cached.Verdict, cached.Headlines)
	}

	headlines, err := c.tavilySearch(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[research] tavily error for %q: %v\n", query, err)
		return Verdict{Silent: true, Score: 0, Summary: "tavily error"}, nil, nil
	}
	if len(headlines) == 0 {
		v := Verdict{Silent: true, Score: 0.5, Summary: "sin resultados Tavily"}
		c.appendCache(CacheEntry{Hash: hash, Query: query, Timestamp: time.Now().UTC().Format(time.RFC3339),
			Headlines: headlines, Verdict: v, Model: "none"})
		return v, headlines, nil
	}

	verdict, err := c.synthesize(question, side, headlines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[research] synthesize error for %q: %v\n", query, err)
		verdict = Verdict{Silent: true, Score: 0, Summary: "synthesize error"}
	}
	c.appendCache(CacheEntry{Hash: hash, Query: query, Timestamp: time.Now().UTC().Format(time.RFC3339),
		Headlines: headlines, Verdict: verdict, Model: llmModel})
	return verdict, headlines, citesFromVerdict(verdict, headlines)
}

// citesFromVerdict turns the LLM's cited_dates into SourceCite by matching headline titles.
// If the verdict is silent or has no cited_dates, returns nil (no attribution).
func citesFromVerdict(v Verdict, headlines []Headline) []commontypes.SourceCite {
	if v.Silent || len(v.CitedDates) == 0 {
		return nil
	}
	titleIdx := map[string]Headline{}
	for _, h := range headlines {
		titleIdx[strings.ToLower(strings.TrimSpace(h.Title))] = h
	}
	out := make([]commontypes.SourceCite, 0, len(v.CitedDates))
	for _, cd := range v.CitedDates {
		key := strings.ToLower(strings.TrimSpace(cd.HeadlineTitle))
		if h, ok := titleIdx[key]; ok {
			out = append(out, commontypes.SourceCite{
				Domain:        h.Source,
				URL:           h.URL,
				HeadlineTitle: h.Title,
				PublishedDate: h.PublishedDate,
			})
			if len(out) >= 3 {
				break
			}
		}
	}
	return out
}
