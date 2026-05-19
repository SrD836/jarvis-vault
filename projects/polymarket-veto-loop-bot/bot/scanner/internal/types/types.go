package types

// Market from the Gamma API.
type Market struct {
	ID           string `json:"id"`
	Slug         string `json:"slug"`
	Question     string `json:"question"`
	OutcomeType  string `json:"outcomeType"`
	Tag          string `json:"tag"`
	Volume24h    string `json:"volume24h"`
	CurrentPrice string `json:"currentPrice"`
	EndDate      string `json:"endDate"`
	Closed       bool   `json:"closed"`
}

// GammaAPIResponse with cursor-based pagination.
type GammaAPIResponse struct {
	Data []Market `json:"data"`
	Next *string  `json:"next"`
}

// Candidate written as JSONL.
type Candidate struct {
	ID              string  `json:"id"`
	Slug            string  `json:"slug"`
	Question        string  `json:"question"`
	Category        string  `json:"category"`
	Volume24h       float64 `json:"volume_24h"`
	CurrentPriceYes float64 `json:"current_price_yes"`
	EndDate         string  `json:"end_date"`
	ScannedAt       string  `json:"scanned_at"`
}
