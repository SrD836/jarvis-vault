// Package estimator routes a candidate to a niche-specific probability
// estimator. The brain's generic LLM (DeepSeek opining) has no real edge
// against an efficient market; a niche estimator instead derives estimated_prob
// from real data (live spot + volatility for crypto/index price-level markets).
//
// The router is deliberately minimal: try estimators in order, first one that
// returns OK wins; if none opine, the brain falls back to the LLM path. Adding
// a niche (e.g. sports-vs-sharp-books) = append another Estimator — no router
// changes.
package estimator

import (
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

// Estimate is a niche estimator's verdict on a candidate.
type Estimate struct {
	EstimatedProb   float64 // P(YES resolves true) from the niche model
	EdgeType        string  // e.g. "model-barrier"
	EdgeDescription string  // human-readable, carries the model inputs for traceability
	Confidence      float64 // 0..1
	OK              bool    // false => this estimator has no opinion => fall back to the LLM
}

// Estimator produces a data-derived probability for candidates in its niche.
type Estimator interface {
	// Estimate returns OK=false when the candidate is outside this estimator's
	// niche or its data inputs are unavailable (so the caller can fall back).
	Estimate(c *types.Candidate) Estimate
	Name() string
}

// Route tries each estimator in order and returns the first OK estimate. If none
// opine it returns a zero Estimate (OK=false) so the brain uses the LLM path.
func Route(c *types.Candidate, ests []Estimator) Estimate {
	for _, e := range ests {
		if est := e.Estimate(c); est.OK {
			return est
		}
	}
	return Estimate{OK: false}
}
