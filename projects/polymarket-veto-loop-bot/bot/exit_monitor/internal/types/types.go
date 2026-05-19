package types

import commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"

// ActiveTrade from the executor's active.jsonl.
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
	PnlPct      float64                  `json:"pnl_pct"`
	Reason      string                   `json:"exit_reason"`
	SourcesUsed []commontypes.SourceCite `json:"sources_used,omitempty"`
	DaysOpen    float64                  `json:"days_open,omitempty"`
}

// MarketPrice from Gamma API GET /markets/{id}. Gamma a veces deja currentPrice vacío,
// hay que caer a lastTradePrice → outcomePrices[0] → mid(bestBid,bestAsk).
type MarketPrice struct {
	ID             string  `json:"id"`
	OutcomeType    string  `json:"outcomeType"`
	Closed         bool    `json:"closed"`
	CurrentPrice   string  `json:"currentPrice"`
	LastTradePrice float64 `json:"lastTradePrice"`
	BestBid        float64 `json:"bestBid"`
	BestAsk        float64 `json:"bestAsk"`
	OutcomePrices  string  `json:"outcomePrices"`
	Volume24h      string  `json:"volume24h"`
}
