package tradingview

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// resetState wipes module-level cache/session so tests are independent.
func resetState() {
	cacheMu.Lock()
	cache = map[string]cacheEntry{}
	cacheMu.Unlock()
	sessMu.Lock()
	sessionID = ""
	sessMu.Unlock()
}

// fakeMCP serves a minimal MCP streamable-http endpoint.
func fakeMCP(t *testing.T, responder func(toolName string, args map[string]any) string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		method, _ := body["method"].(string)
		switch method {
		case "initialize":
			w.Header().Set("mcp-session-id", "test-session-1")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"jsonrpc":"2.0","id":0,"result":{"protocolVersion":"2024-11-05","capabilities":{}}}`)
		case "tools/call":
			params, _ := body["params"].(map[string]any)
			name, _ := params["name"].(string)
			args, _ := params["arguments"].(map[string]any)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, responder(name, args))
		default:
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":{}}`)
		}
	}))
	return srv
}

func TestEvaluate_NonCrypto_ReturnsNil(t *testing.T) {
	resetState()
	r := Evaluate("will-bernie-sanders-win-2026", "Sanders 2026?", "elections", 0.42)
	if r != nil {
		t.Fatalf("expected nil for non-crypto, got %+v", r)
	}
}

func TestEvaluate_BullThesis_ContradictedByBearishStrongSignal_Blocks(t *testing.T) {
	resetState()
	srv := fakeMCP(t, func(name string, args map[string]any) string {
		return `{"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"Recommendation: STRONG_SELL\nConfidence: 80%"}]}}`
	})
	defer srv.Close()
	endpoint = srv.URL

	r := Evaluate("will-btc-be-above-100k-by-june-30", "BTC > $100k?", "crypto", 0.65)
	if r == nil {
		t.Fatal("expected non-nil for crypto")
	}
	if !r.Block {
		t.Fatalf("expected Block=true (bear signal vs bull thesis); got %+v", r)
	}
	if r.Asset != "BTCUSDT" {
		t.Errorf("expected asset=BTCUSDT, got %s", r.Asset)
	}
	if !strings.Contains(strings.ToLower(r.Reason), "contradicts") {
		t.Errorf("reason should mention contradiction: %s", r.Reason)
	}
}

func TestEvaluate_BearThesis_NeutralSignal_NoBlock(t *testing.T) {
	resetState()
	srv := fakeMCP(t, func(name string, args map[string]any) string {
		return `{"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"Recommendation: NEUTRAL\nConfidence: 50%"}]}}`
	})
	defer srv.Close()
	endpoint = srv.URL

	r := Evaluate("will-eth-be-above-10k-by-june", "ETH 10k?", "crypto", 0.10)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	if r.Block {
		t.Fatalf("neutral signal should NOT block; got %+v", r)
	}
}

func TestEvaluate_CacheHit_NoSecondHTTP(t *testing.T) {
	resetState()
	calls := 0
	srv := fakeMCP(t, func(name string, args map[string]any) string {
		calls++
		return `{"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"NEUTRAL\nConfidence: 50%"}]}}`
	})
	defer srv.Close()
	endpoint = srv.URL

	Evaluate("btc-test", "BTC?", "crypto", 0.5)
	// First call → 2 HTTP (initialize + tools/call). Second call hits cache.
	first := calls
	Evaluate("btc-test", "BTC?", "crypto", 0.5)
	if calls != first {
		t.Errorf("cache miss: expected %d calls, got %d", first, calls)
	}
}

func TestParseSentiment_StrongBuy(t *testing.T) {
	s, c := parseSentiment([]byte(`{"content":[{"text":"STRONG_BUY signal\nConfidence: 88%"}]}`))
	if s != "bullish" {
		t.Errorf("expected bullish, got %s", s)
	}
	if c != 0.88 {
		t.Errorf("expected 0.88, got %.2f", c)
	}
}
