package types

import commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"

// Approved is a candidate that passed all veto rules, ready for execution.
//
// v7: edge fields carry the LLM's probability estimate so the executor can
// Kelly-size and so the monitor can exit on thesis invalidation / target hit.
type Approved struct {
	CandidateID      string  `json:"candidate_id"`
	Slug             string  `json:"slug"`
	Question         string  `json:"question"`
	Category         string  `json:"category"`
	CurrentPriceYes  float64 `json:"current_price_yes"`
	Volume24h        float64 `json:"volume_24h"`
	EndDate          string  `json:"end_date"`
	ScannedAt        string  `json:"scanned_at"`
	ApprovedAt       string  `json:"approved_at"`
	ApprovedPriceYes float64                  `json:"approved_price_yes"`
	DaysToResolution int                      `json:"days_to_resolution"`
	Horizon          string                   `json:"horizon"`
	LiquidityUSD     float64                  `json:"liquidity_usd,omitempty"`
	SourcesUsed      []commontypes.SourceCite `json:"sources_used,omitempty"`
	// v7 edge fields.
	EstimatedProb      float64 `json:"estimated_prob,omitempty"`
	EdgeType           string  `json:"edge_type,omitempty"`
	EdgeDescription    string  `json:"edge_description,omitempty"`
	ThesisInvalidation string  `json:"thesis_invalidation,omitempty"`
}

// ActiveTrade stored in active.jsonl after successful entry.
//
// v7 adds EstimatedProb / TargetProb / ThesisInvalidation / EdgeType so the
// exit monitor can decide on tesis-invalidated / target-hit / no-remaining-
// edge instead of SL/TP percent triggers.
type ActiveTrade struct {
	ID             string  `json:"id"`
	MarketID       string  `json:"market_id"`
	Slug           string  `json:"slug,omitempty"`
	Question       string  `json:"question"`
	Side           string  `json:"side"` // "yes" | "no"
	EntryPrice     float64 `json:"entry_price"`
	Size           float64 `json:"size"` // 150 USD
	Category       string  `json:"category"`
	EntryTimestamp  string  `json:"entry_timestamp"`
	ApprovedPrice   float64 `json:"approved_price"`
	DaysToResolution int    `json:"days_to_resolution"`
	Horizon          string                   `json:"horizon"`
	EndDate          string                   `json:"end_date,omitempty"`
	SizeUSD          float64                  `json:"size_usd"`
	LiquidityUSD     float64                  `json:"liquidity_usd,omitempty"`
	SourcesUsed      []commontypes.SourceCite `json:"sources_used,omitempty"`
	// v7 fields.
	EstimatedProb      float64 `json:"estimated_prob,omitempty"`
	TargetProb         float64 `json:"target_prob,omitempty"`
	ThesisInvalidation string  `json:"thesis_invalidation,omitempty"`
	EdgeType           string  `json:"edge_type,omitempty"`
}

// Portfolio state, single JSON object.
type Portfolio struct {
	Bankroll      float64        `json:"bankroll"`
	UsedExposure  float64        `json:"used_exposure"`
	MaxExposure   float64        `json:"max_exposure"`    // 5000 (10%)
	MaxPerTrade   float64        `json:"max_per_trade"`   // 150
	MaxSameCat    int            `json:"max_same_category"` // 2
	Positions     []ActiveTrade  `json:"positions"`
}

// Rejection logged when a trade cannot enter.
type Rejection struct {
	Timestamp  string `json:"timestamp"`
	MarketID   string `json:"market_id"`
	Question   string `json:"question"`
	Reason     string `json:"reason"`
}
