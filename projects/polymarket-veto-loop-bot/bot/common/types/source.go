package types

// SourceCite identifies a single news headline that informed a trading decision.
// Shared between brain (producer), executor (passthrough) and exit_monitor (consumer).
type SourceCite struct {
	Domain        string `json:"domain"`         // host extracted from URL, e.g. "apnews.com"
	URL           string `json:"url"`
	HeadlineTitle string `json:"headline_title"`
	PublishedDate string `json:"published_date"` // ISO date string from Tavily
}
