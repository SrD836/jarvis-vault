package types

// ActiveTrade from the executor's active.jsonl.
type ActiveTrade struct {
	ID             string  `json:"id"`
	MarketID       string  `json:"market_id"`
	Question       string  `json:"question"`
	Side           string  `json:"side"`
	EntryPrice     float64 `json:"entry_price"`
	Size           float64 `json:"size"`
	Category       string  `json:"category"`
	EntryTimestamp  string `json:"entry_timestamp"`
	ApprovedPrice   float64 `json:"approved_price"`
}

// ClosedTrade written to closed.jsonl on exit.
type ClosedTrade struct {
	ID            string  `json:"id"`
	MarketID      string  `json:"market_id"`
	EntryPrice    float64 `json:"entry_price"`
	ExitPrice     float64 `json:"exit_price"`
	EntryTime     string  `json:"entry_time"`
	ExitTime      string  `json:"exit_time"`
	PnlUSD        float64 `json:"pnl_usd"`
	PnlPct        float64 `json:"pnl_pct"`
	Reason        string  `json:"reason"` // "stop_loss" | "take_profit" | "market_closed"
}

// MarketPrice from Gamma API GET /markets/{id}.
type MarketPrice struct {
	ID            string `json:"id"`
	OutcomeType   string `json:"outcomeType"`
	Closed        bool   `json:"closed"`
	CurrentPrice  string `json:"currentPrice"`
	Volume24h     string `json:"volume24h"`
}
