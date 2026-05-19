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

// MarketQuote is the full read-out for a market needed by both entry and exit logic.
type MarketQuote struct {
	BestBid         float64
	BestAsk         float64
	LastTradePrice  float64
	OutcomePricesY  float64
	OutcomePricesN  float64
	LiquidityUSD    float64
	Closed          bool
	Active          bool
	AcceptingOrders bool
	EndDate         string
}

// FetchQuote pulls a single market and returns the structured quote + closed flag.
// Caller decides which price field to use (bestBid for sell, bestAsk for buy).
func FetchQuote(marketID string) (MarketQuote, error) {
	url := fmt.Sprintf("%s/markets/%s", gammaBase, marketID)
	resp, err := httpClient.Get(url)
	if err != nil {
		return MarketQuote{}, fmt.Errorf("gamma GET %s: %w", marketID, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return MarketQuote{}, fmt.Errorf("gamma GET %s: HTTP %d: %s", marketID, resp.StatusCode, string(body))
	}
	var mkt types.MarketPrice
	if err := json.NewDecoder(resp.Body).Decode(&mkt); err != nil {
		return MarketQuote{}, fmt.Errorf("gamma decode %s: %w", marketID, err)
	}
	q := MarketQuote{
		BestBid:         mkt.BestBid,
		BestAsk:         mkt.BestAsk,
		LastTradePrice:  mkt.LastTradePrice,
		LiquidityUSD:    mkt.LiquidityNum,
		Closed:          mkt.Closed,
		Active:          mkt.Active,
		AcceptingOrders: mkt.AcceptingOrders,
		EndDate:         mkt.EndDate,
	}
	if mkt.OutcomePrices != "" {
		var arr []string
		if err := json.Unmarshal([]byte(mkt.OutcomePrices), &arr); err == nil {
			if len(arr) > 0 {
				if v, err := strconv.ParseFloat(arr[0], 64); err == nil {
					q.OutcomePricesY = v
				}
			}
			if len(arr) > 1 {
				if v, err := strconv.ParseFloat(arr[1], 64); err == nil {
					q.OutcomePricesN = v
				}
			}
		}
	}
	return q, nil
}

// PriceForExecution returns the realistic execution price for the given intent + side.
//   intent "buy":  pay bestAsk (when side="yes", the YES bestAsk).
//   intent "sell": receive bestBid.
// Side is "yes" (default) or "no" (mirror prices using outcomePrices[1] / inverse logic).
//
// Fallback chain when orderbook is empty:
//   1. lastTradePrice  (label "lastTradePrice_fallback")
//   2. outcomePrices[idx]  (label "outcomePrices_fallback")
// Returns (price, source). price=0 means "no price available".
func PriceForExecution(q MarketQuote, side, intent string) (float64, string) {
	if side == "" {
		side = "yes"
	}
	if intent == "buy" {
		if side == "yes" && q.BestAsk > 0 {
			return q.BestAsk, "bestAsk"
		}
		if side == "no" && q.BestBid > 0 {
			// buying NO is symmetric: 1 - YES bid (Polymarket prices YES+NO=1).
			return 1.0 - q.BestBid, "bestAsk_NO_derived"
		}
	} else { // sell (default)
		if side == "yes" && q.BestBid > 0 {
			return q.BestBid, "bestBid"
		}
		if side == "no" && q.BestAsk > 0 {
			return 1.0 - q.BestAsk, "bestBid_NO_derived"
		}
	}
	// Fallbacks
	if q.LastTradePrice > 0 {
		return q.LastTradePrice, "lastTradePrice_fallback"
	}
	if side == "yes" && q.OutcomePricesY > 0 {
		return q.OutcomePricesY, "outcomePrices_fallback"
	}
	if side == "no" && q.OutcomePricesN > 0 {
		return q.OutcomePricesN, "outcomePrices_fallback"
	}
	return 0, ""
}

// GetMarketPrice is the legacy entry point. Preserves the old signature for callers
// that don't need side/liquidity (kept for backward compat — should be migrated).
// Uses mid-price (last trade or bid_ask_mid) as a neutral estimator. NEVER use for
// real exits/entries; use FetchQuote + PriceForExecution instead.
func GetMarketPrice(marketID string) (float64, bool, error) {
	q, err := FetchQuote(marketID)
	if err != nil {
		return 0, false, err
	}
	// Neutral midpoint estimate (compat with old mid behaviour).
	var price float64
	switch {
	case q.LastTradePrice > 0:
		price = q.LastTradePrice
	case q.BestBid > 0 && q.BestAsk > 0:
		price = (q.BestBid + q.BestAsk) / 2.0
	case q.OutcomePricesY > 0:
		price = q.OutcomePricesY
	}
	if price <= 0 {
		return 0, q.Closed, fmt.Errorf("gamma %s: no price available (bestBid=%.4f bestAsk=%.4f lastTrade=%.4f)",
			marketID, q.BestBid, q.BestAsk, q.LastTradePrice)
	}
	return price, q.Closed, nil
}
