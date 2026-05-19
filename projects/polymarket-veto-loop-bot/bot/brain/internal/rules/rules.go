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
