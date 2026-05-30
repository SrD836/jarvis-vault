// Package marketcheck — P6 veto rule for market-related candidates.
//
// Detects polymarket slugs that reference tradeable assets (BTC, ETH, WTI,
// SPX, individual stocks) and queries Yahoo Finance for the current price.
// Vetoes the candidate if the live price contradicts the thesis embedded
// in the slug (e.g. "Will Bitcoin be above $78,000 on May 25?" with
// current_price_yes=0.24 → if BTC is at $76k with 1h to close, the
// implied yes-probability is way too high and the bid is dumb money).
//
// We use Yahoo Finance directly (not the tradingview MCP) because the brain
// is a Go cron one-shot — shelling out to claude -p per candidate would
// add 5-15s of latency and pollute the prompt cache. Yahoo's chart API is
// free, no auth, returns last close + intraday OHLC.
package marketcheck

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/marketdata"
)

// SlugPattern maps a regex-matchable polymarket slug fragment to a Yahoo
// Finance symbol and a price-threshold extractor.
type slugPattern struct {
	rx     *regexp.Regexp
	symbol string
}

// known asset patterns. We extract a "above $X" or "hit $X" target from the
// question text and compare it to spot.
var assetPatterns = []slugPattern{
	{regexp.MustCompile(`(?i)\bbitcoin\b|\bbtc\b`), "BTC-USD"},
	{regexp.MustCompile(`(?i)\bethereum\b|\beth\b`), "ETH-USD"},
	{regexp.MustCompile(`(?i)\bsolana\b|\bsol\b`), "SOL-USD"},
	{regexp.MustCompile(`(?i)\bwti\b|crude.?oil`), "CL=F"},  // WTI crude futures
	{regexp.MustCompile(`(?i)\bbrent\b`), "BZ=F"},
	{regexp.MustCompile(`(?i)nasdaq|\bndx\b`), "^IXIC"},
	{regexp.MustCompile(`(?i)s.p.?500|spx\b`), "^GSPC"},
	{regexp.MustCompile(`(?i)\btesla\b|\btsla\b`), "TSLA"},
	{regexp.MustCompile(`(?i)\bnvidia\b|\bnvda\b`), "NVDA"},
}

// targetPriceRE captures phrases like "above $78,000", "hit $90", "reach 6500".
var targetPriceRE = regexp.MustCompile(`(?i)\b(?:above|below|hit|reach|over|under|at)\b[^\d]*\$?(\d{1,3}(?:[,.]\d{3})*(?:\.\d+)?)`)

// Direction words. "hit"/"reach" are NEUTRAL (you can hit a high OR a low), so
// they are NOT in aboveRE — otherwise "hit (LOW) $80" would resolve to "above".
// Direction is carried by above/over/high vs below/under/low/dip/drop/fall;
// default (no word) is "above".
var aboveRE = regexp.MustCompile(`(?i)\b(above|over|high)\b`)
var belowRE = regexp.MustCompile(`(?i)\b(below|under|low|dip|drop|fall)\b`)

// CheckResult is returned by Evaluate. Block=true means brain should veto.
type CheckResult struct {
	Asset       string
	Symbol      string
	SpotPrice   float64
	TargetPrice float64
	Direction   string // "above" | "below"
	Reason      string
	Block       bool
}

// Parse extracts the tradeable asset (Yahoo symbol), the target price and the
// direction ("above"|"below") from a polymarket slug+question. ok=false when the
// candidate is not an asset-price market or the target can't be read. Shared by
// the P6 veto (Evaluate) and the crypto barrier estimator so both agree on what
// "the asset, the target and the side" are.
func Parse(slug, question string) (symbol string, target float64, dir string, ok bool) {
	text := slug + " " + question
	var pat *slugPattern
	for i := range assetPatterns {
		if assetPatterns[i].rx.MatchString(text) {
			pat = &assetPatterns[i]
			break
		}
	}
	if pat == nil {
		return "", 0, "", false
	}
	m := targetPriceRE.FindStringSubmatch(question)
	if len(m) < 2 {
		return "", 0, "", false
	}
	t, err := strconv.ParseFloat(strings.ReplaceAll(m[1], ",", ""), 64)
	if err != nil {
		return "", 0, "", false
	}
	dir = "above"
	if belowRE.MatchString(question) && !aboveRE.MatchString(question) {
		dir = "below"
	}
	return pat.symbol, t, dir, true
}

// Evaluate returns a result for market-related candidates only. Non-market
// slugs return nil (no opinion). Network errors return nil — best effort.
func Evaluate(slug, question string, currentPriceYes float64) *CheckResult {
	symbol, target, dir, ok := Parse(slug, question)
	if !ok {
		return nil
	}

	spot, err := marketdata.Spot(symbol)
	if err != nil || spot <= 0 {
		return nil
	}

	res := &CheckResult{
		Asset: symbol, Symbol: symbol, SpotPrice: spot,
		TargetPrice: target, Direction: dir,
	}

	// Distance from target as % of spot.
	gap := (target - spot) / spot
	if dir == "below" {
		gap = -gap
	}

	// Veto conditions:
	//   thesis "above target": YES needs price to rise. If price is already
	//   well above target → trivially YES → polymarket should be ~0.95+.
	//   If current_price_yes < 0.50 AND spot is already past target by >2%,
	//   the market is mispriced or close is imminent — skip (we don't trade
	//   confused books).
	//   If spot is below target by >10% AND current_price_yes > 0.30, the
	//   bet is implausible — veto.
	switch {
	case dir == "above" && spot > target*1.02 && currentPriceYes < 0.50:
		res.Block = true
		res.Reason = fmt.Sprintf("P6 market: %s spot $%.2f already > target $%.2f but yes=%.2f (confused book)", symbol, spot, target, currentPriceYes)
	case dir == "above" && gap > 0.10 && currentPriceYes > 0.30:
		res.Block = true
		res.Reason = fmt.Sprintf("P6 market: %s spot $%.2f is %.1f%% below target $%.2f, yes=%.2f implausible", symbol, spot, gap*100, target, currentPriceYes)
	case dir == "below" && spot < target*0.98 && currentPriceYes < 0.50:
		res.Block = true
		res.Reason = fmt.Sprintf("P6 market: %s spot $%.2f already < target $%.2f but yes=%.2f (confused book)", symbol, spot, target, currentPriceYes)
	case dir == "below" && gap > 0.10 && currentPriceYes > 0.30:
		res.Block = true
		res.Reason = fmt.Sprintf("P6 market: %s spot $%.2f is %.1f%% above target $%.2f, yes=%.2f implausible", symbol, spot, gap*100, target, currentPriceYes)
	}
	return res
}
