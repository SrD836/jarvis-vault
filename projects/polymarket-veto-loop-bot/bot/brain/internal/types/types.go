package types

import (
	"time"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
)

// Candidate from scanner output (candidates.jsonl)
type Candidate struct {
	ID             string  `json:"id"`
	Slug           string  `json:"slug"`
	Question       string  `json:"question"`
	Category       string  `json:"category"`
	Volume24h      float64 `json:"volume_24h"`
	CurrentPriceYes float64 `json:"current_price_yes"`
	EndDate        string  `json:"end_date"`
	ScannedAt      string  `json:"scanned_at"`
}

// VetoResult stores the outcome of evaluating a candidate against V1–V6.
type VetoResult struct {
	CandidateID string `json:"candidate_id"`
	Slug        string `json:"slug"`
	Blocked     bool   `json:"blocked"`
	Reason      string `json:"reason,omitempty"`
	VetoedBy    string `json:"vetoed_by,omitempty"` // e.g. "V1", "V4", "LLM_V3"
}

// Approved is a candidate that passed all veto rules, ready for sizing.
type Approved struct {
	CandidateID      string  `json:"candidate_id"`
	Slug             string  `json:"slug"`
	Question         string  `json:"question"`
	Category         string  `json:"category"`
	CurrentPriceYes  float64 `json:"current_price_yes"`
	Volume24h        float64 `json:"volume_24h"`
	EndDate          string  `json:"end_date"`
	ScannedAt        string  `json:"scanned_at"`
	ApprovedAt        string  `json:"approved_at"`
	ApprovedPriceYes float64                  `json:"approved_price_yes"`
	DaysToResolution int                      `json:"days_to_resolution"`
	Horizon          string                   `json:"horizon"` // short | medium | long
	SourcesUsed      []commontypes.SourceCite `json:"sources_used,omitempty"`
}

// LLMRequest payload to the JARVIS dashboard LLM bridge.
type LLMRequest struct {
	Model   string            `json:"model"`
	System  string            `json:"system"`
	Message string            `json:"message"`
	Context map[string]string `json:"context,omitempty"`
}

// LLMResponse from the dashboard LLM bridge.
type LLMResponse struct {
	Content string `json:"content"`
	Model   string `json:"model"`
	Usage   struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

// LLMBlockResult is the structured output we ask the LLM to return.
type LLMBlockResult struct {
	Block  bool   `json:"block"`
	Reason string `json:"reason"`
	Rule   string `json:"rule"` // "V3", "V5", "V6"
}

// PriceHistory for V4 chasing check (tracking price movement in last 4h).
type PriceHistory struct {
	Prices []float64 `json:"prices"`
	Timestamps []time.Time `json:"timestamps"`
}
