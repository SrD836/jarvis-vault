// Package categorize infers a semantic category from a polymarket slug+question
// when the Gamma API returns an empty category field (which is the case for
// ~100% of markets). Without this inference the softrules cluster collapses
// to horizon|band and the M2 rule cannot discriminate between a politics
// loss and a sports loss — defeating the learning loop.
//
// Categories used downstream by softrules cluster keys and the per-category
// hard rules (P7-P10 in rules.go).
package categorize

import (
	"regexp"
	"strings"
)

// Categories — kept short and stable. Adding a new one means updating the
// rule layer that consumes it.
const (
	CatSportsGame       = "sports-game"
	CatSportsSeason     = "sports-season"
	CatElections        = "elections"
	CatGeopolitics      = "geopolitics"
	CatCryptoLaunch     = "crypto-launch"
	CatMarket           = "market" // BTC/ETH/WTI/SPX — handled by P6 marketcheck
	CatWeatherNatural   = "weather-natural"
	CatExecutiveAction  = "executive-action"
	CatEntertainment    = "entertainment"
	CatOther            = "other"
)

type rule struct {
	cat string
	rx  *regexp.Regexp
}

// Ordered: first match wins. Order matters when phrases overlap (e.g.
// "trump-resign" matches executive-action before geopolitics).
var rules = []rule{
	// Same-day sports first — very high volume + most precise pattern.
	{CatSportsGame, regexp.MustCompile(`(?i)will-.+-win-on-\d{4}-\d{2}-\d{2}|both-teams-to-score|end-in-a-draw|vs\.?-.+-end|win-the-(coin-toss|game)|match-result`)},

	// Season-long sports: championships, finals, MVPs.
	{CatSportsSeason, regexp.MustCompile(`(?i)\b(champion(ship)?|final[s]?|world-cup|mvp|playoff|nba-finals|nfl-(nfc|afc)|premier-league|champions-league|league-championship|trophy-this-season)\b`)},

	// Elections / politics nominations.
	{CatElections, regexp.MustCompile(`(?i)\b(election|primary|governor|presidential|nominat(ion|ee)|prime-minister-of|democratic|republican)\b`)},

	// Executive actions on named individuals — before geopolitics so
	// "trump-resign" goes here, not into geopolitics by accident.
	{CatExecutiveAction, regexp.MustCompile(`(?i)\b(out-by|resign|fired-by|impeach|indictment|sanction|recogniz(e|ed)|pardon|extradit|step-down)\b`)},

	// Crypto launch / pre-event — caught by P4 already but tagging here
	// makes softrules clusters semantically meaningful.
	{CatCryptoLaunch, regexp.MustCompile(`(?i)\b(ipo|launch|tge|mainnet|listing|airdrop|token-by|fdv-above|public-ticker)\b`)},

	// Pure market price predictions — handed off to P6 marketcheck.
	{CatMarket, regexp.MustCompile(`(?i)\b(bitcoin|btc|ethereum|eth|solana|sol-usd|wti|crude-oil|brent|nasdaq|s.?p.?500|tesla|nvidia)\b.*\b(above|below|hit|reach|over|under|dip|by)\b`)},

	// Weather / natural disasters / fringe — historically pure noise.
	{CatWeatherNatural, regexp.MustCompile(`(?i)\b(earthquake|volcanic|hurricane|tornado|hantavirus|meteor|alien|ufo|pandemic|asteroid|tsunami)\b`)},

	// Geopolitics last because many of its keywords (iran, ukraine) can
	// appear inside elections/executive-action slugs.
	{CatGeopolitics, regexp.MustCompile(`(?i)\b(iran|israel|ukraine|russia|hezbollah|gaza|hamas|airspace|ceasefire|invade|nuclear|strait-of-hormuz|nato|kurd|peacekeeping)\b`)},

	// Entertainment / awards.
	{CatEntertainment, regexp.MustCompile(`(?i)\b(oscar|grammy|olympic|opening-weekend|spotify|billion-views|emmy|billboard)\b`)},
}

// Infer returns the best-match category for a slug+question pair. Returns
// CatOther when nothing matches. Both inputs are concatenated before matching
// so that the regex can hit on either; this is intentional — Polymarket
// slugs are often abbreviated and the question text disambiguates.
func Infer(slug, question string) string {
	if slug == "" && question == "" {
		return CatOther
	}
	text := strings.ToLower(slug + " " + question)
	for _, r := range rules {
		if r.rx.MatchString(text) {
			return r.cat
		}
	}
	return CatOther
}
