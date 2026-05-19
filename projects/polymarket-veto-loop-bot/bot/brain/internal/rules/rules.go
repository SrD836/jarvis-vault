package rules

import (
	"fmt"
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

// VetoV2 checks catalyst window: Veto if market resolution is <72h from now.
// EndDate is the market resolution date. If empty, V2 passes (no known catalyst → let LLM decide V6).
// If end_date is within 72h from now → veto.
func VetoV2(c *types.Candidate) *types.VetoResult {
	if c.EndDate == "" {
		return nil // no known resolution date → not V2 domain
	}

	end, err := time.Parse(time.RFC3339, c.EndDate)
	if err != nil {
		// try common ISO format
		end, err = time.Parse("2006-01-02T15:04:05Z", c.EndDate)
		if err != nil {
			return nil // can't parse → not blocking on ambiguous date
		}
	}

	now := time.Now().UTC()
	remaining := end.Sub(now)

	if remaining > 0 && remaining < 72*time.Hour {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("catalyst dentro de ventana 72h pre-resolución: resolución en %.0f horas", remaining.Hours()),
			VetoedBy:    "V2",
		}
	}

	return nil
}

// VetoV4 checks momentum chasing: Veto if price moved ≥8% in the last 4h.
// Uses the midpoint of current_price_yes as reference. In a real setup this
// would compare against a 4h-old snapshot. Here we derive from the fact that
// the scanner's price is current; for V4 to be meaningful we'd need a 4h-old
// price. If unavailable, we pass (no data → no chase detected conservatively).
// For v1 we implement a version that works with the candidate data: we assume
// if the market has no price history in the candidate, we log and pass.
func VetoV4(c *types.Candidate) *types.VetoResult {
	// V4 requires historical price data (4h ago) which scanner doesn't provide.
	// In production, brain would fetch 4h candles from CLOB API.
	// Here: pass by default and log that V4 was skipped due to missing data.
	// The actual implementation should call gamma-api /markets/:id to get
	// recent prices. For now we mark as pass (conservative: don't block on
	// insufficient data).
	return nil
}

// EvaluateNumeric runs V1, V2, V4 sequentially. Returns first veto or nil.
func EvaluateNumeric(c *types.Candidate) *types.VetoResult {
	for _, fn := range []func(*types.Candidate) *types.VetoResult{
		VetoV1,
		VetoV2,
		VetoV4,
	} {
		if r := fn(c); r != nil {
			return r
		}
	}
	return nil
}
