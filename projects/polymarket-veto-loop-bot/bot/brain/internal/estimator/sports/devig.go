// Package sports is the niche estimator for sports markets. It prices a
// Polymarket sports question against a sharp bookmaker (Pinnacle, fallback
// Betfair Exchange) via the-odds-api: Polymarket sports prices lag the sharp
// closing lines, so the devigged sharp probability is a real "true" reference
// — genuine information arbitrage, not a bet against an efficient market.
package sports

// Devig removes the bookmaker overround from a field of decimal odds and
// returns the implied true probabilities (same order). Multiplicative method:
// p_i = (1/odds_i) / sum(1/odds_j). Returns nil if any odd is <= 1 (invalid) or
// the field is empty. For a 2-outcome h2h this is just the same normalization.
func Devig(decimalOdds []float64) []float64 {
	if len(decimalOdds) == 0 {
		return nil
	}
	inv := make([]float64, len(decimalOdds))
	var sum float64
	for i, o := range decimalOdds {
		if o <= 1.0 {
			return nil // decimal odds must be > 1.0
		}
		inv[i] = 1.0 / o
		sum += inv[i]
	}
	if sum <= 0 {
		return nil
	}
	out := make([]float64, len(decimalOdds))
	for i := range inv {
		out[i] = inv[i] / sum
	}
	return out
}

// TrueProb devigs the field and returns the true probability at index target.
// ok=false when the field is invalid or target is out of range.
func TrueProb(decimalOdds []float64, target int) (float64, bool) {
	if target < 0 || target >= len(decimalOdds) {
		return 0, false
	}
	probs := Devig(decimalOdds)
	if probs == nil {
		return 0, false
	}
	return probs[target], true
}
