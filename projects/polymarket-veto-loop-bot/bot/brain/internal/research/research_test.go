package research

import "testing"

// Normalized-title match: LLM paraphrases the headline ("...!" vs plain) but it
// still resolves to the real Tavily headline's source host.
func TestCitesFromVerdictNormalizedTitle(t *testing.T) {
	hls := []Headline{
		{Title: "Trump signs Iran nuclear deal", URL: "https://www.reuters.com/world/x", Source: "reuters.com", PublishedDate: "2026-05-28"},
	}
	v := Verdict{Score: 0.7, CitedDates: []CitedDate{{HeadlineTitle: "Trump signs Iran nuclear deal!", Date: "2026-05-28"}}}
	got := citesFromVerdict(v, hls)
	if len(got) != 1 || got[0].Domain != "reuters.com" {
		t.Fatalf("want reuters.com via normalized match, got %+v", got)
	}
}

// URL fallback: no headline matches the (very different) cited title, but the
// cite carries a real URL — we still attribute the real domain.
func TestCitesFromVerdictURLFallback(t *testing.T) {
	v := Verdict{Score: 0.6, CitedDates: []CitedDate{
		{HeadlineTitle: "totally unrelated phrasing about something else entirely", Date: "2026-05-27", URL: "https://apnews.com/article/abc123"},
	}}
	got := citesFromVerdict(v, nil)
	if len(got) != 1 || got[0].Domain != "apnews.com" {
		t.Fatalf("want apnews.com via URL fallback, got %+v", got)
	}
}

// Fuzzy contains: cited title is a substring of the real headline.
func TestCitesFromVerdictFuzzyContains(t *testing.T) {
	hls := []Headline{
		{Title: "Breaking: Strait of Hormuz traffic returns to normal levels", URL: "https://bbc.com/n/1", Source: "bbc.com"},
	}
	v := Verdict{Score: 0.6, CitedDates: []CitedDate{{HeadlineTitle: "Strait of Hormuz traffic returns to normal"}}}
	got := citesFromVerdict(v, hls)
	if len(got) != 1 || got[0].Domain != "bbc.com" {
		t.Fatalf("want bbc.com via fuzzy contains, got %+v", got)
	}
}

// Silent verdict yields no attribution.
func TestCitesFromVerdictSilentNil(t *testing.T) {
	if got := citesFromVerdict(Verdict{Silent: true}, nil); got != nil {
		t.Fatalf("silent must yield nil, got %+v", got)
	}
}

// Distinct domains only, capped at 3, www stripped → reuters dedups.
func TestCitesFromVerdictDedupDomain(t *testing.T) {
	v := Verdict{Score: 0.7, CitedDates: []CitedDate{
		{HeadlineTitle: "a", URL: "https://reuters.com/1"},
		{HeadlineTitle: "b", URL: "https://www.reuters.com/2"},
		{HeadlineTitle: "c", URL: "https://bbc.com/3"},
	}}
	got := citesFromVerdict(v, nil)
	if len(got) != 2 {
		t.Fatalf("want 2 distinct domains (reuters,bbc), got %d: %+v", len(got), got)
	}
}

func TestNormalizeTitle(t *testing.T) {
	cases := map[string]string{
		"Trump's deal, signed!": "trump s deal signed",
		"  Hello   World  ":     "hello world",
		"!!!":                   "",
	}
	for in, want := range cases {
		if got := normalizeTitle(in); got != want {
			t.Errorf("normalizeTitle(%q)=%q want %q", in, got, want)
		}
	}
}
