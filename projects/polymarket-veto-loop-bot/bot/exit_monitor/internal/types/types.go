package types

import commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"

// ActiveTrade from the executor's active.jsonl.
//
// v7: EstimatedProb / TargetProb / ThesisInvalidation drive thesis-based
// exits in monitor.go (target hit, no remaining edge, future: invalidation).
type ActiveTrade struct {
	ID               string                   `json:"id"`
	MarketID         string                   `json:"market_id"`
	Question         string                   `json:"question"`
	Slug             string                   `json:"slug,omitempty"`
	Side             string                   `json:"side"`
	EntryPrice       float64                  `json:"entry_price"`
	Size             float64                  `json:"size"`
	SizeUSD          float64                  `json:"size_usd,omitempty"`
	Category         string                   `json:"category"`
	EntryTimestamp   string                   `json:"entry_timestamp"`
	ApprovedPrice    float64                  `json:"approved_price"`
	DaysToResolution int                      `json:"days_to_resolution,omitempty"`
	Horizon          string                   `json:"horizon,omitempty"`
	EndDate          string                   `json:"end_date,omitempty"`
	SourcesUsed      []commontypes.SourceCite `json:"sources_used,omitempty"`
	EstimatedProb      float64 `json:"estimated_prob,omitempty"`
	TargetProb         float64 `json:"target_prob,omitempty"`
	ThesisInvalidation string  `json:"thesis_invalidation,omitempty"`
	EdgeType           string  `json:"edge_type,omitempty"`
}

// ClosedTrade written to closed.jsonl on exit.
// Includes both the original `pnl_usd` (legacy) and `pnl` (consumed by dashboard /api/bot/pnl).
type ClosedTrade struct {
	ID         string  `json:"id"`
	MarketID   string  `json:"market_id"`
	Slug       string  `json:"slug,omitempty"`
	Question   string  `json:"question,omitempty"`
	Category   string  `json:"category,omitempty"`
	Side       string  `json:"side,omitempty"`
	Size       float64 `json:"size,omitempty"`
	EntryPrice float64 `json:"entry_price"`
	ExitPrice  float64 `json:"exit_price"`
	EntryTime  string  `json:"entry_time"`
	ExitTime   string  `json:"exit_timestamp"`
	Pnl        float64 `json:"pnl"`
	PnlUSD     float64 `json:"pnl_usd"`
	PnlPct          float64                  `json:"pnl_pct"`
	Reason          string                   `json:"exit_reason"`
	SourcesUsed     []commontypes.SourceCite `json:"sources_used,omitempty"`
	DaysOpen        float64                  `json:"days_open,omitempty"`
	EarlyExit       bool                     `json:"early_exit,omitempty"`
	EndDate         string                   `json:"end_date,omitempty"`
	ExitPriceSource string                   `json:"exit_price_source,omitempty"`
	LiquidityUSD    float64                  `json:"liquidity_usd,omitempty"`
	Horizon         string                   `json:"horizon,omitempty"`
}

// MarketPrice from Gamma API GET /markets/{id}. Gamma a veces deja currentPrice vacío,
// hay que caer a lastTradePrice → outcomePrices[0] → mid(bestBid,bestAsk).
type MarketPrice struct {
	ID              string  `json:"id"`
	OutcomeType     string  `json:"outcomeType"`
	Closed          bool    `json:"closed"`
	Active          bool    `json:"active"`
	AcceptingOrders bool    `json:"acceptingOrders"`
	EndDate         string  `json:"endDate"`
	CurrentPrice    string  `json:"currentPrice"`
	LastTradePrice  float64 `json:"lastTradePrice"`
	BestBid         float64 `json:"bestBid"`
	BestAsk         float64 `json:"bestAsk"`
	OutcomePrices   string  `json:"outcomePrices"`
	Volume24h       string  `json:"volume24h"`
	LiquidityNum    float64 `json:"liquidityNum"`
}
