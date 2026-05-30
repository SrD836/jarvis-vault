// Package marketdata is a tiny free-tier Yahoo Finance client shared by the
// brain's marketcheck veto (P6) and the crypto barrier estimator. Both need
// live spot; the estimator also needs realized volatility. Centralizing the
// client here avoids two copies of the Yahoo plumbing and keeps the niche
// estimator dependency-free (no new external API — Yahoo is already in use).
//
// Yahoo's chart API is free, unauthenticated, and returns last price + intraday
// OHLC. All functions are best-effort: on any error they return a zero value +
// error so callers can decline (the estimator falls back to the LLM, the veto
// becomes a no-op).
package marketdata

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	chartBase  = "https://query1.finance.yahoo.com/v8/finance/chart/"
	httpUA     = "Mozilla/5.0 (polymarket-veto-bot)"
	httpTimeout = 8 * time.Second
)

// Spot returns the current regular-market price for a Yahoo symbol (e.g.
// "BTC-USD", "^GSPC"). Returns 0 + error on any failure.
func Spot(symbol string) (float64, error) {
	var out struct {
		Chart struct {
			Result []struct {
				Meta struct {
					RegularMarketPrice float64 `json:"regularMarketPrice"`
				} `json:"meta"`
			} `json:"result"`
		} `json:"chart"`
	}
	if err := fetchChart(symbol, "1d", "1d", &out); err != nil {
		return 0, err
	}
	if len(out.Chart.Result) == 0 {
		return 0, fmt.Errorf("marketdata: no chart result for %s", symbol)
	}
	p := out.Chart.Result[0].Meta.RegularMarketPrice
	if p <= 0 {
		return 0, fmt.Errorf("marketdata: non-positive spot for %s", symbol)
	}
	return p, nil
}

// DailyVol returns the ANNUALIZED volatility of daily log-returns over the last
// ~3 months (annualized with 365 calendar days so it composes with a barrier
// tau expressed in calendar-year fractions: sigma*sqrt(tau) = sigmaDaily*sqrt(days)).
// Returns 0 + error when there is not enough clean history (<20 returns).
func DailyVol(symbol string) (float64, error) {
	var out struct {
		Chart struct {
			Result []struct {
				Indicators struct {
					Quote []struct {
						Close []float64 `json:"close"`
					} `json:"quote"`
				} `json:"indicators"`
			} `json:"result"`
		} `json:"chart"`
	}
	if err := fetchChart(symbol, "1d", "3mo", &out); err != nil {
		return 0, err
	}
	if len(out.Chart.Result) == 0 || len(out.Chart.Result[0].Indicators.Quote) == 0 {
		return 0, fmt.Errorf("marketdata: no history for %s", symbol)
	}
	closes := out.Chart.Result[0].Indicators.Quote[0].Close
	// Yahoo returns nulls (decoded as 0) for non-trading days; drop them.
	clean := make([]float64, 0, len(closes))
	for _, c := range closes {
		if c > 0 {
			clean = append(clean, c)
		}
	}
	return AnnualizedVol(clean)
}

// AnnualizedVol computes the annualized stdev of daily log-returns from a close
// series (oldest→newest). Exported so it can be unit-tested without the network.
func AnnualizedVol(closes []float64) (float64, error) {
	if len(closes) < 21 { // need >=20 returns for a usable estimate
		return 0, fmt.Errorf("marketdata: insufficient history (%d closes)", len(closes))
	}
	rets := make([]float64, 0, len(closes)-1)
	for i := 1; i < len(closes); i++ {
		rets = append(rets, math.Log(closes[i]/closes[i-1]))
	}
	var sum float64
	for _, r := range rets {
		sum += r
	}
	mean := sum / float64(len(rets))
	var ss float64
	for _, r := range rets {
		d := r - mean
		ss += d * d
	}
	variance := ss / float64(len(rets)-1) // sample variance
	sigmaDaily := math.Sqrt(variance)
	return sigmaDaily * math.Sqrt(365.0), nil
}

func fetchChart(symbol, interval, rng string, dst interface{}) error {
	url := fmt.Sprintf("%s%s?interval=%s&range=%s", chartBase, symbol, interval, rng)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", httpUA)
	resp, err := (&http.Client{Timeout: httpTimeout}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("marketdata: yahoo HTTP %d for %s", resp.StatusCode, symbol)
	}
	return json.Unmarshal(body, dst)
}
