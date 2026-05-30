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

// ---------------------------------------------------------------------------
// Exit-price guardrail (fix 2026-05-30): un cierre NUNCA debe ejecutarse a un
// precio que el sistema no puede validar como real. Distingue "mercado resolvió"
// (resultado real, settlement 0/1) de "no pude obtener precio" (diferir, no cerrar).
// ---------------------------------------------------------------------------

// Banda de sanidad: con mercado NO resuelto, un precio que se desploma de forma
// absurda respecto al último válido conocido (ref) se bloquea para revisión.
const (
	sanityRefFloor = 0.10 // solo aplicamos la banda si el ref era >= 10%
	sanityAbsFloor = 0.02 // un colapso a < 2% (p.ej. 0.44 -> 0.001) es sospechoso
	sanityRatio    = 0.10 // o caer a < 10% del ref
)

// ExitDecision es el precio de cierre validado para una posición.
type ExitDecision struct {
	Price    float64
	Source   string
	Resolved bool // true => mercado genuinamente resuelto (settlement real)
	OK       bool // false => sin precio fiable; el caller DEBE diferir (no cerrar)
}

// suspicious detecta un desplome absurdo de precio en un mercado NO resuelto.
func suspicious(p, ref float64) bool {
	if ref < sanityRefFloor {
		return false // ref bajo: longshots legítimamente baratos, no aplicar banda
	}
	return p < sanityAbsFloor || p < ref*sanityRatio
}

// Settlement deriva el valor de liquidación real (0 o 1) de un mercado resuelto,
// para el side dado. outcomePrices manda; si no, lastTradePrice>=0.5 como proxy.
func Settlement(q MarketQuote, side string) float64 {
	if side == "" {
		side = "yes"
	}
	if q.OutcomePricesY > 0 || q.OutcomePricesN > 0 {
		switch side {
		case "no":
			if q.OutcomePricesN >= 0.5 {
				return 1
			}
			return 0
		default: // yes
			if q.OutcomePricesY >= 0.5 {
				return 1
			}
			return 0
		}
	}
	// Sin outcomePrices: usar lastTradePrice como proxy del resultado YES.
	yesWon := q.LastTradePrice >= 0.5
	if side == "no" {
		if yesWon {
			return 0
		}
		return 1
	}
	if yesWon {
		return 1
	}
	return 0
}

// ExitPrice decide el precio de cierre SIN volcar nunca a un fallback no validado.
// ref = último precio válido conocido (entry/approved) para la banda de sanidad.
func ExitPrice(q MarketQuote, side string, ref float64) ExitDecision {
	if side == "" {
		side = "yes"
	}
	// 1. Resolución genuina: liquidar al resultado real (0 o 1).
	if q.Closed {
		return ExitDecision{Price: Settlement(q, side), Source: "resolved_settlement", Resolved: true, OK: true}
	}
	// 2. Mercado vivo: solo un precio real de orderbook es aceptable.
	if side == "yes" && q.BestBid > 0 {
		p := q.BestBid
		if suspicious(p, ref) {
			return ExitDecision{Price: p, Source: "blocked_sanity_band", OK: false}
		}
		return ExitDecision{Price: p, Source: "bestBid", OK: true}
	}
	if side == "no" && q.BestAsk > 0 {
		p := 1.0 - q.BestAsk
		if suspicious(p, ref) {
			return ExitDecision{Price: p, Source: "blocked_sanity_band", OK: false}
		}
		return ExitDecision{Price: p, Source: "bestBid_NO_derived", OK: true}
	}
	// 3. Sin precio fiable y no resuelto: NO cerrar. Nunca lastTradePrice para ejecución.
	return ExitDecision{Price: 0, Source: "no_reliable_price", OK: false}
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
