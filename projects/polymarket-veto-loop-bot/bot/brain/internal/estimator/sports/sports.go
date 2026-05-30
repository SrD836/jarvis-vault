package sports

import (
	"fmt"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/estimator"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

// Estimator prices Polymarket outright-winner sports questions against the
// devigged sharp line (Pinnacle/Betfair via the-odds-api). Declines (OK=false)
// for non-sports questions, uncovered competitions (e.g. Champions League),
// unmatched teams, or when no API key is configured.
type Estimator struct {
	client *Client
}

// New builds the estimator. If no odds API key is configured the client is nil
// and the estimator declines everything (safe to wire unconditionally).
func New() *Estimator {
	c, ok := NewClient()
	if !ok {
		return &Estimator{client: nil}
	}
	return &Estimator{client: c}
}

func (e *Estimator) Name() string { return "sports-sharp-odds" }

func (e *Estimator) Estimate(c *types.Candidate) estimator.Estimate {
	if e.client == nil {
		return estimator.Estimate{OK: false}
	}
	sportKey, ok := MapCompetition(c.Question)
	if !ok {
		return estimator.Estimate{OK: false}
	}
	team := ExtractTeam(c.Question)
	if team == "" {
		return estimator.Estimate{OK: false}
	}
	outcomes, err := e.client.Outrights(sportKey)
	if err != nil || len(outcomes) < 2 {
		return estimator.Estimate{OK: false}
	}
	names := make([]string, len(outcomes))
	odds := make([]float64, len(outcomes))
	for i, o := range outcomes {
		names[i] = o.Name
		odds[i] = o.Price
	}
	idx := MatchOutcome(team, names)
	if idx < 0 {
		return estimator.Estimate{OK: false}
	}
	p, ok := TrueProb(odds, idx)
	if !ok {
		return estimator.Estimate{OK: false}
	}
	return estimator.Estimate{
		EstimatedProb: p,
		EdgeType:      "model-sharp-odds",
		EdgeDescription: fmt.Sprintf("sharp %s: %q devig P=%.3f (campo %d) vs implied=%.3f",
			sportKey, team, p, len(outcomes), c.CurrentPriceYes),
		Confidence: 0.8,
		OK:         true,
	}
}
