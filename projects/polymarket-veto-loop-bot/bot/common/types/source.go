package types

// SourceCite identifies a single news headline that informed a trading decision.
// Shared between brain (producer), executor (passthrough) and exit_monitor (consumer).
//
// Synthetic=true means the LLM produced this citation without surfacing a real
// URL (legacy Claudemax-websearch fallback). The closer excludes synthetic
// cites from source-stats.jsonl so the domain blacklist does not collapse
// every outlet into one pseudo-domain. v7 P2 (2026-05-28).
type SourceCite struct {
	Domain        string `json:"domain"`         // host extracted from URL, e.g. "apnews.com"
	URL           string `json:"url"`
	HeadlineTitle string `json:"headline_title"`
	PublishedDate string `json:"published_date"` // ISO date string from Tavily
	Synthetic     bool   `json:"synthetic,omitempty"`
}
