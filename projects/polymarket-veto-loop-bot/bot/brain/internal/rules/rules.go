package rules

import (
	"fmt"
	"regexp"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

// VetoV1 checks volume: Veto if Vol 24h < 50,000 USD.
func VetoV1(c *types.Candidate) *types.VetoResult {
	if c.Volume24h < 50000 {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("volumen insuficiente: %.2f USD < 50000 USD", c.Volume24h),
			VetoedBy:    "V1",
		}
	}
	return nil
}

// VetoV2 checks catalyst window: Veto if resolution is within underHours of now.
func VetoV2(c *types.Candidate, underHours int) *types.VetoResult {
	if c.EndDate == "" {
		return nil
	}

	end, err := time.Parse(time.RFC3339, c.EndDate)
	if err != nil {
		end, err = time.Parse("2006-01-02T15:04:05Z", c.EndDate)
		if err != nil {
			return nil
		}
	}

	now := time.Now().UTC()
	remaining := end.Sub(now)

	if remaining > 0 && remaining < time.Duration(underHours)*time.Hour {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("catalyst dentro de ventana %dh pre-resolución: resolución en %.0f horas", underHours, remaining.Hours()),
			VetoedBy:    "V2",
		}
	}

	return nil
}

// VetoV4 checks momentum chasing. Requires 4h historical price data not available from scanner.
func VetoV4(c *types.Candidate) *types.VetoResult {
	return nil
}

// EvaluateNumeric runs V1, V2, V4 sequentially. Returns first veto or nil.
func EvaluateNumeric(c *types.Candidate, v2UnderHours int) *types.VetoResult {
	if r := VetoV1(c); r != nil {
		return r
	}
	if r := VetoV2(c, v2UnderHours); r != nil {
		return r
	}
	return VetoV4(c)
}

// ---------------------------------------------------------------------------
// v6 hardening (post-mortem Discord trade T-608399 — entry 0.066, lost -89% in 10min)
// ---------------------------------------------------------------------------

// daysUntilISO returns days until the given ISO-8601 timestamp. Returns -1 if empty/unparsable.
// Local helper to avoid an import cycle with brain package.
func daysUntilISO(endDate string) int {
	if endDate == "" {
		return -1
	}
	t, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05Z", endDate)
		if err != nil {
			return -1
		}
	}
	delta := t.Sub(time.Now().UTC())
	return int(delta.Hours() / 24)
}

// EvalPriceBand: P0_floor if YES price is below the configured floor (with relaxed
// floor for imminent binary events), P0_ceiling if above the ceiling.
//
// v7 P4 (2026-05-28): floor lowered from 0.10 → 0.05. Additional relaxation
// to 0.03 when a short-horizon candidate has liquidity above
// PriceFloorShortLiqRelaxUSD (default $20k): deep books absorb the bid/ask
// spread well enough that low-priced bets stay tradeable without the SL/TP
// asymmetric burn that motivated the v6 0.10 floor (and which v7 already
// removed by killing % stop-loss).
//
// EvalPriceBandV7 accepts the extra relaxation knobs; the old 2-arg
// EvalPriceBand stays as a thin shim so untouched callers keep compiling.
func EvalPriceBand(c *types.Candidate, floor, ceiling float64) *types.VetoResult {
	return EvalPriceBandV7(c, floor, ceiling, 0, 0)
}

func EvalPriceBandV7(c *types.Candidate, floor, ceiling, relaxFloor, relaxLiqUSD float64) *types.VetoResult {
	if floor <= 0 {
		floor = 0.05
	}
	if ceiling <= 0 {
		ceiling = 0.95
	}

	effFloor := floor
	days := daysUntilISO(c.EndDate)
	if days >= 0 && days <= 7 {
		// Imminent binary events: keep at floor (already 0.05 by default).
	}
	// v7 P4 short-horizon + deep-book relaxation.
	if relaxFloor > 0 && relaxLiqUSD > 0 &&
		days >= 0 && days <= 7 &&
		c.LiquidityUSD >= relaxLiqUSD &&
		relaxFloor < effFloor {
		effFloor = relaxFloor
	}

	if c.CurrentPriceYes < effFloor {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("price floor: %.4f < %.3f (horizon %d d, liq $%.0f)", c.CurrentPriceYes, effFloor, days, c.LiquidityUSD),
			VetoedBy:    "P0_floor",
		}
	}
	if c.CurrentPriceYes > ceiling {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("price ceiling: %.4f > %.3f", c.CurrentPriceYes, ceiling),
			VetoedBy:    "P0_ceiling",
		}
	}
	return nil
}

// EvalAbsoluteLiquidity: veto if visible orderbook depth (USD) is below an absolute
// threshold, regardless of trade size. Complements executor's Q3 Size×ratio gate.
//
// A market with $2k of total depth is a trap even with a $20 trade: the bestBid can
// collapse on a single market-sell from another participant. Discord trade had $2916
// liquidity and lost -89% in 10 minutes.
//
// Pass-through if minLiq <= 0 (rule disabled) or if c.LiquidityUSD == 0 (legacy
// candidate from pre-v5 scanner that didn't capture liquidity yet).
func EvalAbsoluteLiquidity(c *types.Candidate, minLiq float64) *types.VetoResult {
	if minLiq <= 0 {
		return nil
	}
	if c.LiquidityUSD > 0 && c.LiquidityUSD < minLiq {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("liquidity $%.0f < absolute min $%.0f", c.LiquidityUSD, minLiq),
			VetoedBy:    "P3_low_absolute_liquidity",
		}
	}
	return nil
}

// preEventSlugRegex matches Polymarket slugs for pre-launch/pre-IPO style markets.
// The \b word boundaries reduce false positives (matches "ipo-day" but not "biopolymer").
var preEventSlugRegex = regexp.MustCompile(`(?i)\b(ipo-day|ipo|tge|mainnet|listing|release-date|launch|launches|launching)\b`)

// EvalPreEventVeto: veto markets resolving around a future product launch / IPO / token
// generation event when the event is still >= minDays away.
//
// Rationale: the price of a "day-zero" event is signal-from-noise until the launch is
// imminent. There's no public news flow to research, the orderbook is thin, and price
// can move on rumour with no recovery path.
//
// As the event approaches (days_to_end < minDays), real news appears and the regular
// V1/V2/N1/N2/LLM stack can evaluate the market on its merits.
func EvalPreEventVeto(c *types.Candidate, minDays int) *types.VetoResult {
	if minDays <= 0 {
		return nil
	}
	days := daysUntilISO(c.EndDate)
	if days < minDays {
		// Close enough to the event that catalysts/news can be researched normally.
		return nil
	}
	if preEventSlugRegex.MatchString(c.Slug) {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("pre-event slug + %d d to resolution (>=%d threshold)", days, minDays),
			VetoedBy:    "P4_pre_event",
		}
	}
	return nil
}

// EvalWeatherNatural (P7): veto every "weather-natural" candidate.
// Historical wr in this cluster is < 10% — earthquake / volcanic / pandemic
// markets are pure noise with no information edge for a retail predictor.
func EvalWeatherNatural(c *types.Candidate) *types.VetoResult {
	if c.Category != "weather-natural" {
		return nil
	}
	return &types.VetoResult{
		CandidateID: c.ID, Slug: c.Slug, Blocked: true,
		Reason:   "P7: weather/natural-disaster cluster — no edge",
		VetoedBy: "P7",
	}
}

// EvalElectionFar (P8): veto elections > 90 days away with price in noisy
// band [0.30, 0.70]. Far-future elections oscillate on news; only trade
// extreme convictions.
func EvalElectionFar(c *types.Candidate, daysToEnd int) *types.VetoResult {
	if c.Category != "elections" || daysToEnd <= 90 {
		return nil
	}
	p := c.CurrentPriceYes
	if p < 0.30 || p > 0.70 {
		return nil
	}
	return &types.VetoResult{
		CandidateID: c.ID, Slug: c.Slug, Blocked: true,
		Reason: fmt.Sprintf("P8: election %dd out, price %.2f en banda ruido [0.30, 0.70]",
			daysToEnd, p),
		VetoedBy: "P8",
	}
}

// EvalGeopoliticsPump (P9): geopolitics with low price + short horizon
// is the pump-and-dump cluster (ceasefire/airspace/strait rumor spikes).
// Veto price ∈ [0.05, 0.30] AND days_to_resolution ≤ 14.
func EvalGeopoliticsPump(c *types.Candidate, daysToEnd int) *types.VetoResult {
	if c.Category != "geopolitics" || daysToEnd > 14 || daysToEnd < 0 {
		return nil
	}
	p := c.CurrentPriceYes
	if p < 0.05 || p > 0.30 {
		return nil
	}
	return &types.VetoResult{
		CandidateID: c.ID, Slug: c.Slug, Blocked: true,
		Reason:   fmt.Sprintf("P9: geopolitics pump cluster (price %.2f, %dd)", p, daysToEnd),
		VetoedBy: "P9",
	}
}

// EvalSportsChalk (P10): same-day sports with price > 0.85 has RR ≤ 0.18.
// One loss erases 5 wins. Veto.
func EvalSportsChalk(c *types.Candidate, daysToEnd int) *types.VetoResult {
	if c.Category != "sports-game" || daysToEnd > 2 || daysToEnd < 0 {
		return nil
	}
	if c.CurrentPriceYes <= 0.85 {
		return nil
	}
	return &types.VetoResult{
		CandidateID: c.ID, Slug: c.Slug, Blocked: true,
		Reason: fmt.Sprintf("P10: sports same-day chalk (price %.2f, RR≈%.2f)",
			c.CurrentPriceYes, (1-c.CurrentPriceYes)/c.CurrentPriceYes),
		VetoedBy: "P10",
	}
}
