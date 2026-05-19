package types

// Market mirrors the relevant fields of the Polymarket gamma-api response.
// API: https://gamma-api.polymarket.com/markets
type Market struct {
	ID              string   `json:"id"`
	Slug            string   `json:"slug"`
	Question        string   `json:"question"`
	VolumeNum       float64  `json:"volumeNum"`
	Volume24Hr      *float64 `json:"volume24Hr"`
	LastTradePrice  float64  `json:"lastTradePrice"`
	BestBid         float64  `json:"bestBid"`
	BestAsk         float64  `json:"bestAsk"`
	LiquidityNum    float64  `json:"liquidityNum"`
	AcceptingOrders bool     `json:"acceptingOrders"`
	OutcomePrices   string   `json:"outcomePrices"`
	Outcomes        string   `json:"outcomes"`
	EndDate         string   `json:"endDate"`
	StartDate       string   `json:"startDate"`
	Closed          bool     `json:"closed"`
	Active          bool     `json:"active"`
	Events          []struct {
		Title    string `json:"title"`
		Category string `json:"category"`
	} `json:"events"`
}

// Candidate written as JSONL to vault/inbox/trading/candidates.jsonl.
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
