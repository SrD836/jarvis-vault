// Package tradingview — P11 veto rule: TradingView MCP cross-check for crypto candidates.
//
// Polymarket questions about crypto assets (BTC, ETH, SOL, ...) get a second
// opinion from the tradingview-mcp container (github.com/atilaahmettaner/tradingview-mcp).
// If TradingView's coin_analysis returns a directional signal that contradicts
// the implied yes-probability of the market, we veto.
//
// Why this exists (P6 marketcheck.go uses Yahoo Finance directly): Yahoo gives
// us spot price but no signal/momentum. TradingView's coin_analysis returns a
// BUY/SELL/NEUTRAL bias with confidence — useful as second-gate veto when our
// thesis disagrees with the technical consensus.
//
// Hard rule: brain is a Go cron one-shot; we hit tradingview-mcp directly via
// HTTP (streamable-http MCP transport) without going through claude-cli. Cache
// 5min per asset to avoid hammering on rate-limited free tier.
package tradingview

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	endpoint = envOr("TRADINGVIEW_MCP_URL", "http://tradingview-mcp:8766/mcp")
	timeout  = 8 * time.Second
)

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

// Result is the gate output. Block=true means brain should veto with P11.
type Result struct {
	Asset      string
	Block      bool
	Reason     string
	Sentiment  string // "bullish" | "bearish" | "neutral" | "unknown"
	Confidence float64
}

type cacheEntry struct {
	r  *Result
	at time.Time
}

var (
	cacheMu   sync.Mutex
	cache     = map[string]cacheEntry{}
	sessionID string
	sessMu    sync.Mutex
)

// assetPatterns maps slug+question keywords to tradingview symbols.
var assetPatterns = []struct {
	re  *regexp.Regexp
	sym string
}{
	{regexp.MustCompile(`(?i)\bbitcoin\b|\bbtc\b`), "BTCUSDT"},
	{regexp.MustCompile(`(?i)\bethereum\b|\beth\b`), "ETHUSDT"},
	{regexp.MustCompile(`(?i)\bsolana\b|\bsol\b`), "SOLUSDT"},
	{regexp.MustCompile(`(?i)\bripple\b|\bxrp\b`), "XRPUSDT"},
	{regexp.MustCompile(`(?i)\bcardano\b|\bada\b`), "ADAUSDT"},
	{regexp.MustCompile(`(?i)\bdogecoin\b|\bdoge\b`), "DOGEUSDT"},
}

func detectAsset(text string) string {
	for _, p := range assetPatterns {
		if p.re.MatchString(text) {
			return p.sym
		}
	}
	return ""
}

// dependsOnUnderlying reports whether a Polymarket question's resolution
// depends on the spot price of an underlying asset (vs. a categorical event
// like "will-coinbase-list-X" or "ethereum-merge-completes"). Only price-
// dependent markets benefit from TradingView's technical signal.
//
// v7 rationale: TA over a binary-event market price is a category error.
// We restrict P11 to questions where the underlying ASSET price is the
// thing being asked about.
var underlyingPricePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\b(price|above|below|hit|reach|exceed|cross|over|under)\b[^.]{0,40}\$?\s*\d`),
	regexp.MustCompile(`(?i)\$\s*\d[\d,kKmMbB.]*\b`),
	regexp.MustCompile(`(?i)\b\d+[kKmMbB]\b`),
	regexp.MustCompile(`(?i)\b(close|closing|finish|end)\s+(above|below|at)\b`),
	regexp.MustCompile(`(?i)\b(all[- ]time[- ]high|ath|new high|new low)\b`),
}

func dependsOnUnderlying(slug, question string) bool {
	hay := strings.ToLower(slug + " " + question)
	for _, re := range underlyingPricePatterns {
		if re.MatchString(hay) {
			return true
		}
	}
	return false
}

// Evaluate returns a P11 veto verdict. nil = no opinion (not crypto or network error).
// category and currentPriceYes are used to scope and direction-check.
func Evaluate(slug, question, category string, currentPriceYes float64) *Result {
	asset := detectAsset(slug + " " + question)
	isCrypto := asset != "" || strings.HasPrefix(strings.ToLower(category), "crypto")
	if !isCrypto {
		return nil
	}
	if asset == "" {
		// Crypto category but no asset detected — skip (cannot query specific symbol)
		return nil
	}
	// v7: only run TradingView for markets whose resolution depends on the
	// underlying asset price. "Will-coinbase-list-token-x" and "ethereum-merge-
	// completes" are crypto-flavoured but TA tells us nothing about them.
	if !dependsOnUnderlying(slug, question) {
		return nil
	}

	// Cache lookup
	cacheMu.Lock()
	if e, ok := cache[asset]; ok && time.Since(e.at) < 5*time.Minute {
		cacheMu.Unlock()
		return e.r
	}
	cacheMu.Unlock()

	r := evaluateLive(asset, currentPriceYes)

	cacheMu.Lock()
	cache[asset] = cacheEntry{r: r, at: time.Now()}
	cacheMu.Unlock()
	return r
}

func evaluateLive(asset string, currentPriceYes float64) *Result {
	log.Printf("[P11] evaluateLive asset=%s priceYes=%.2f endpoint=%s", asset, currentPriceYes, endpoint)
	sid, err := ensureSession()
	if err != nil {
		log.Printf("[P11] ensureSession error: %v", err)
		return nil // best effort — no opinion on network error
	}
	raw, err := callTool(sid, "coin_analysis", map[string]any{"symbol": asset})
	if err != nil {
		log.Printf("[P11] callTool error: %v", err)
		return nil
	}
	preview := string(raw)
	if len(preview) > 500 {
		preview = preview[:500] + "...(truncated)"
	}
	log.Printf("[P11] response asset=%s bytes=%d preview=%s", asset, len(raw), strings.ReplaceAll(preview, "\n", "\\n"))
	sent, conf := parseSentiment(raw)
	log.Printf("[P11] parsed asset=%s sentiment=%s confidence=%.2f", asset, sent, conf)
	r := &Result{
		Asset:      asset,
		Sentiment:  sent,
		Confidence: conf,
	}

	// Veto logic:
	//   currentPriceYes >= 0.50 → market believes outcome is likely (bullish thesis on the question)
	//   currentPriceYes <  0.50 → market believes outcome is unlikely (bearish thesis)
	// If TradingView's directional sentiment STRONGLY contradicts the thesis → block.
	// Require confidence >= 0.60 to avoid noise.
	if conf < 0.60 {
		r.Reason = fmt.Sprintf("tradingview confidence %.2f < 0.60, no opinion", conf)
		return r
	}
	bullThesis := currentPriceYes >= 0.50
	if bullThesis && sent == "bearish" {
		r.Block = true
		r.Reason = fmt.Sprintf("tradingview %s sentiment=bearish (conf %.2f) contradicts bull thesis (price_yes=%.2f)", asset, conf, currentPriceYes)
	} else if !bullThesis && sent == "bullish" {
		r.Block = true
		r.Reason = fmt.Sprintf("tradingview %s sentiment=bullish (conf %.2f) contradicts bear thesis (price_yes=%.2f)", asset, conf, currentPriceYes)
	} else {
		r.Reason = fmt.Sprintf("tradingview %s sentiment=%s (conf %.2f) aligned with thesis", asset, sent, conf)
	}
	log.Printf("[P11] decision asset=%s block=%t reason=%s", asset, r.Block, r.Reason)
	return r
}

// ensureSession initializes the MCP session lazily. Cached for process lifetime.
func ensureSession() (string, error) {
	sessMu.Lock()
	defer sessMu.Unlock()
	if sessionID != "" {
		return sessionID, nil
	}
	body := `{"jsonrpc":"2.0","id":0,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{}}}`
	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(body))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json,text/event-stream")
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	sid := resp.Header.Get("mcp-session-id")
	if sid == "" {
		return "", fmt.Errorf("tradingview-mcp returned no session id (status %d)", resp.StatusCode)
	}
	// notifications/initialized (fire-and-forget)
	nbody := `{"jsonrpc":"2.0","method":"notifications/initialized"}`
	req2, _ := http.NewRequest("POST", endpoint, strings.NewReader(nbody))
	req2.Header.Set("content-type", "application/json")
	req2.Header.Set("accept", "application/json,text/event-stream")
	req2.Header.Set("mcp-session-id", sid)
	if r2, e := client.Do(req2); e == nil {
		r2.Body.Close()
	}
	sessionID = sid
	return sid, nil
}

func callTool(sid, name string, args map[string]any) ([]byte, error) {
	payload, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0", "id": 1, "method": "tools/call",
		"params": map[string]any{"name": name, "arguments": args},
	})
	req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json,text/event-stream")
	req.Header.Set("mcp-session-id", sid)
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// parseSentiment extracts directional bias + confidence from the tradingview-mcp
// coin_analysis response. The MCP returns SSE-framed JSON-RPC like:
//
//	event: message
//	data: {"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"<JSON>"}],"isError":false}}
//
// where <JSON> is a stringified object with rich technical analysis. We extract
// directional bias from:
//   - market_structure.trend ("Bullish"/"Bearish"/"Neutral")
//   - market_sentiment.overall_rating (numeric -3..+3)
//   - market_sentiment.buy_sell_signal ("STRONG_BUY"/"BUY"/"NEUTRAL"/"SELL"/"STRONG_SELL")
//   - market_structure.trend_strength ("Strong"/"Moderate"/"Weak")
//
// Confidence is derived from trend_strength × abs(overall_rating)/3.
func parseSentiment(raw []byte) (sentiment string, confidence float64) {
	text := string(raw)

	// SSE-framed response: extract the "data: {...}" line.
	jsonLine := text
	for _, ln := range strings.Split(text, "\n") {
		ln = strings.TrimSpace(ln)
		if strings.HasPrefix(ln, "data: ") {
			jsonLine = strings.TrimPrefix(ln, "data: ")
			break
		}
	}

	// First decode the outer JSON-RPC envelope.
	var envelope struct {
		Result struct {
			Content []struct {
				Text string `json:"text"`
			} `json:"content"`
		} `json:"result"`
	}
	if err := json.Unmarshal([]byte(jsonLine), &envelope); err != nil || len(envelope.Result.Content) == 0 {
		// Fallback to regex over the raw text (backwards compatibility / unit tests with simpler shape)
		return parseSentimentLegacy(text)
	}

	// Then decode the inner stringified JSON.
	var ta struct {
		MarketStructure struct {
			Trend         string `json:"trend"`
			TrendScore    int    `json:"trend_score"`
			TrendStrength string `json:"trend_strength"`
		} `json:"market_structure"`
		MarketSentiment struct {
			OverallRating int    `json:"overall_rating"`
			BuySellSignal string `json:"buy_sell_signal"`
			Momentum      string `json:"momentum"`
		} `json:"market_sentiment"`
	}
	if err := json.Unmarshal([]byte(envelope.Result.Content[0].Text), &ta); err != nil {
		return parseSentimentLegacy(text)
	}

	// Directional signal: prefer market_structure.trend, fallback to overall_rating sign.
	trend := strings.ToLower(ta.MarketStructure.Trend)
	switch trend {
	case "bullish":
		sentiment = "bullish"
	case "bearish":
		sentiment = "bearish"
	default:
		if ta.MarketSentiment.OverallRating > 0 {
			sentiment = "bullish"
		} else if ta.MarketSentiment.OverallRating < 0 {
			sentiment = "bearish"
		} else {
			sentiment = "neutral"
		}
	}

	// STRONG_BUY/STRONG_SELL override (rarer but more conclusive).
	switch ta.MarketSentiment.BuySellSignal {
	case "STRONG_BUY":
		sentiment = "bullish"
	case "STRONG_SELL":
		sentiment = "bearish"
	}

	// Confidence: combine trend_strength + magnitude of rating.
	strength := strings.ToLower(ta.MarketStructure.TrendStrength)
	switch strength {
	case "strong":
		confidence = 0.80
	case "moderate":
		confidence = 0.65
	case "weak":
		confidence = 0.50
	default:
		confidence = 0.40
	}
	// Boost if STRONG_* signal present and aligned with trend.
	if ta.MarketSentiment.BuySellSignal == "STRONG_BUY" || ta.MarketSentiment.BuySellSignal == "STRONG_SELL" {
		confidence = max64(confidence, 0.85)
	}
	// Boost if overall_rating has high magnitude.
	mag := ta.MarketSentiment.OverallRating
	if mag < 0 {
		mag = -mag
	}
	if mag >= 2 {
		confidence = max64(confidence, 0.65)
	}
	return sentiment, confidence
}

func max64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// parseSentimentLegacy keeps the regex-based fallback for tests and unexpected formats.
func parseSentimentLegacy(text string) (sentiment string, confidence float64) {
	textU := strings.ToUpper(text)
	sentiment = "unknown"
	switch {
	case strings.Contains(textU, "STRONG_BUY") || strings.Contains(textU, "STRONG BUY"):
		sentiment = "bullish"
		confidence = 0.85
	case strings.Contains(textU, "STRONG_SELL") || strings.Contains(textU, "STRONG SELL"):
		sentiment = "bearish"
		confidence = 0.85
	case strings.Contains(textU, `"BUY"`) || strings.Contains(textU, "RECOMMENDATION: BUY") || strings.Contains(textU, "SIGNAL: BUY"):
		sentiment = "bullish"
		confidence = 0.65
	case strings.Contains(textU, `"SELL"`) || strings.Contains(textU, "RECOMMENDATION: SELL") || strings.Contains(textU, "SIGNAL: SELL"):
		sentiment = "bearish"
		confidence = 0.65
	case strings.Contains(textU, "NEUTRAL"):
		sentiment = "neutral"
		confidence = 0.50
	}
	if m := regexp.MustCompile(`(?i)confidence[^0-9]{0,10}(\d{1,3})\s*%`).FindStringSubmatch(text); m != nil {
		var pct int
		fmt.Sscanf(m[1], "%d", &pct)
		if pct > 0 && pct <= 100 {
			confidence = float64(pct) / 100.0
		}
	}
	return sentiment, confidence
}
