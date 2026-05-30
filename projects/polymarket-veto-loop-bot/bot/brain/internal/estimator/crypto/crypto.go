package crypto

import (
	"fmt"
	"strings"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/marketdata"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/estimator"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/marketcheck"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

// Estimator prices asset price-level markets (BTC/ETH/SOL/WTI/SPX/TSLA/NVDA...)
// from live spot + realized volatility via the barrier model. It declines
// (OK=false) for anything it can't parse or price, so the brain falls back to
// the generic LLM.
type Estimator struct {
	spot func(string) (float64, error) // injectable for tests
	vol  func(string) (float64, error)
}

// New builds the estimator wired to the live Yahoo marketdata client.
func New() *Estimator {
	return &Estimator{spot: marketdata.Spot, vol: marketdata.DailyVol}
}

func (e *Estimator) Name() string { return "crypto-barrier" }

// touchWords mark barrier-style resolution (hit any time before T);
// terminalWords mark at-expiry resolution (state of the price at the close).
var touchWords = []string{"hit", "reach", "by ", "before", "any point", "anytime", "ever"}
var terminalWords = []string{" on ", "at close", "closing", "end of", "at the end"}

func classifyMode(question string) string {
	q := strings.ToLower(question)
	for _, w := range terminalWords {
		if strings.Contains(q, w) {
			return "terminal"
		}
	}
	for _, w := range touchWords {
		if strings.Contains(q, w) {
			return "touch"
		}
	}
	return "touch" // default: most "<asset> above $K by <date>" markets are touch-style
}

// Estimate prices the candidate, or returns OK=false to defer to the LLM.
func (e *Estimator) Estimate(c *types.Candidate) estimator.Estimate {
	symbol, target, dir, ok := marketcheck.Parse(c.Slug, c.Question)
	if !ok {
		return estimator.Estimate{OK: false}
	}
	days := daysUntil(c.EndDate)
	if days < 0 {
		return estimator.Estimate{OK: false} // unknown horizon → can't price
	}
	spot, err := e.spot(symbol)
	if err != nil || spot <= 0 {
		return estimator.Estimate{OK: false}
	}
	sigma, err := e.vol(symbol)
	if err != nil || sigma <= 0 {
		return estimator.Estimate{OK: false}
	}
	tau := float64(days) / 365.0
	if tau <= 0 {
		tau = 0.5 / 365.0 // resolves within a day → half-day floor so the model still prices
	}
	mode := classifyMode(c.Question)
	p := Prob(spot, target, sigma, tau, dir, mode)
	return estimator.Estimate{
		EstimatedProb: p,
		EdgeType:      "model-barrier",
		EdgeDescription: fmt.Sprintf("barrier %s/%s: S=$%.2f K=$%.2f sigma=%.0f%% tau=%dd -> P=%.3f vs implied=%.3f",
			symbol, mode, spot, target, sigma*100, days, p, c.CurrentPriceYes),
		Confidence: 0.7,
		OK:         true,
	}
}

func daysUntil(endDate string) int {
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
	return int(time.Until(t).Hours() / 24)
}
