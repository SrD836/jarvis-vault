package polyclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const gammaBase = "https://gamma-api.polymarket.com"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// GetMarketPrice fetches a single market from Gamma API and returns current yes price.
// Returns price, closed bool, error.
func GetMarketPrice(marketID string) (float64, bool, error) {
	url := fmt.Sprintf("%s/markets/%s", gammaBase, marketID)
	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, false, fmt.Errorf("gamma GET %s: %w", marketID, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, false, fmt.Errorf("gamma GET %s: HTTP %d: %s", marketID, resp.StatusCode, string(body))
	}

	var mkt types.MarketPrice
	if err := json.NewDecoder(resp.Body).Decode(&mkt); err != nil {
		return 0, false, fmt.Errorf("gamma decode %s: %w", marketID, err)
	}

	price, source := extractPrice(&mkt)
	if price <= 0 {
		return 0, mkt.Closed, fmt.Errorf("gamma %s: no price available (currentPrice=%q lastTradePrice=%.4f bestBid=%.4f bestAsk=%.4f outcomePrices=%q)",
			marketID, mkt.CurrentPrice, mkt.LastTradePrice, mkt.BestBid, mkt.BestAsk, mkt.OutcomePrices)
	}
	_ = source // available for caller debugging
	return price, mkt.Closed, nil
}

// extractPrice returns the best available YES price using a fallback chain:
//   1. currentPrice (string) — primary, but often empty in gamma-api
//   2. lastTradePrice (float64) — last fill in CLOB, usually populated
//   3. outcomePrices (JSON-string array) — "["0.285","0.715"]", index 0 is YES
//   4. midpoint of bestBid + bestAsk — orderbook mid when no recent trades
// Returns (price, source_label). source is empty if no price found.
func extractPrice(m *types.MarketPrice) (float64, string) {
	if m.CurrentPrice != "" {
		if p, err := strconv.ParseFloat(m.CurrentPrice, 64); err == nil && p > 0 {
			return p, "currentPrice"
		}
	}
	if m.LastTradePrice > 0 {
		return m.LastTradePrice, "lastTradePrice"
	}
	if m.OutcomePrices != "" {
		var arr []string
		if err := json.Unmarshal([]byte(m.OutcomePrices), &arr); err == nil && len(arr) > 0 {
			if p, err := strconv.ParseFloat(arr[0], 64); err == nil && p > 0 {
				return p, "outcomePrices[0]"
			}
		}
	}
	if m.BestBid > 0 && m.BestAsk > 0 {
		return (m.BestBid + m.BestAsk) / 2.0, "bid_ask_mid"
	}
	return 0, ""
}
