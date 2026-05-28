package types

import (
	"time"

	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
)

// Candidate from scanner output (candidates.jsonl)
type Candidate struct {
	ID              string  `json:"id"`
	Slug            string  `json:"slug"`
	Question        string  `json:"question"`
	Category        string  `json:"category"`
	Volume24h       float64 `json:"volume_24h"`
	CurrentPriceYes float64 `json:"current_price_yes"`
	LiquidityUSD    float64 `json:"liquidity_usd,omitempty"`
	EndDate         string  `json:"end_date"`
	ScannedAt       string  `json:"scanned_at"`
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
//
// v7: carries the LLM's declared edge so the executor can Kelly-size and
// so the monitor can exit on tesis invalidation / target hit.
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
	LiquidityUSD     float64                  `json:"liquidity_usd,omitempty"`
	SourcesUsed      []commontypes.SourceCite `json:"sources_used,omitempty"`
	// v7 edge-gate fields. EstimatedProb is the LLM's probability that YES
	// resolves true. TargetProb is the price at which the executor closes
	// the position (defaults to EstimatedProb).
	EstimatedProb      float64 `json:"estimated_prob,omitempty"`
	EdgeType           string  `json:"edge_type,omitempty"`
	EdgeDescription    string  `json:"edge_description,omitempty"`
	ThesisInvalidation string  `json:"thesis_invalidation,omitempty"`
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
//
// v7: extended with the edge-gate fields. EstimatedProb / EdgeType /
// EdgeDescription / ThesisInvalidation are populated even when Block=false,
// so the brain can route the candidate through E1/E2 checks before approving.
type LLMBlockResult struct {
	Block              bool    `json:"block"`
	Reason             string  `json:"reason"`
	Rule               string  `json:"rule"` // "V3", "V5", "V6"
	EstimatedProb      float64 `json:"estimated_prob,omitempty"`
	EdgeType           string  `json:"edge_type,omitempty"` // info|arb|calibration|liquidity|other|none
	EdgeDescription    string  `json:"edge_description,omitempty"`
	ThesisInvalidation string  `json:"thesis_invalidation,omitempty"`
}

// PriceHistory for V4 chasing check (tracking price movement in last 4h).
type PriceHistory struct {
	Prices []float64 `json:"prices"`
	Timestamps []time.Time `json:"timestamps"`
}
