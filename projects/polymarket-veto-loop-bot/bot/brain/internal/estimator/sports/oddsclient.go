package sports

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Outcome is one competitor's sharp decimal price.
type Outcome struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Client fetches outright (futures) odds from the-odds-api, preferring Pinnacle
// (fallback Betfair Exchange) and DISK-caching per sport key. The brain is a
// one-shot cron process (48 runs/day), so an in-memory cache would be lost every
// run and blow the 500-req/month free tier — the disk cache is what keeps the
// budget bounded (each sport fetched at most once per TTL across all runs).
type Client struct {
	apiKey    string
	cachePath string
	regions   string
	ttl       time.Duration
	mem       map[string][]Outcome // per-process memo (avoids re-reading disk per candidate)
}

type cacheFile struct {
	Entries map[string]cacheEntry `json:"entries"`
}
type cacheEntry struct {
	FetchedAt string    `json:"fetched_at"`
	Outcomes  []Outcome `json:"outcomes"`
}

const oddsBase = "https://api.the-odds-api.com/v4/sports/"

// NewClient reads the API key from $ODDS_API_KEY or the connector JSON
// ($ODDS_CONNECTOR, default the host gateway workspace). Returns ok=false when
// no key is available so the estimator declines cleanly (falls back to the LLM).
func NewClient() (*Client, bool) {
	key := os.Getenv("ODDS_API_KEY")
	if key == "" {
		path := envOr("ODDS_CONNECTOR", "/home/agent/.openclaw/workspace/connectors/odds.json")
		if data, err := os.ReadFile(path); err == nil {
			var conn struct {
				APIKey string `json:"api_key"`
			}
			if json.Unmarshal(data, &conn) == nil {
				key = conn.APIKey
			}
		}
	}
	if key == "" {
		return nil, false
	}
	ttlHours := 6
	if v := os.Getenv("ODDS_CACHE_TTL_HOURS"); v != "" {
		fmt.Sscanf(v, "%d", &ttlHours)
	}
	return &Client{
		apiKey:    key,
		cachePath: envOr("ODDS_CACHE", "vault/inbox/trading/odds-cache.json"),
		regions:   envOr("ODDS_REGIONS", "us,eu"),
		ttl:       time.Duration(ttlHours) * time.Hour,
		mem:       map[string][]Outcome{},
	}, true
}

// Outrights returns the sharp-book outright field for a futures sport key,
// served from disk cache when fresh. nil + error when unavailable.
func (c *Client) Outrights(sportKey string) ([]Outcome, error) {
	if o, ok := c.mem[sportKey]; ok {
		return o, nil
	}
	if o, ok := c.cacheLookup(sportKey); ok {
		c.mem[sportKey] = o
		return o, nil
	}
	o, err := c.fetchOutrights(sportKey)
	if err != nil {
		return nil, err
	}
	c.mem[sportKey] = o
	c.cacheStore(sportKey, o)
	return o, nil
}

func (c *Client) fetchOutrights(sportKey string) ([]Outcome, error) {
	url := fmt.Sprintf("%s%s/odds/?apiKey=%s&regions=%s&markets=outrights&oddsFormat=decimal",
		oddsBase, sportKey, c.apiKey, c.regions)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := (&http.Client{Timeout: 15 * time.Second}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("odds API HTTP %d for %s", resp.StatusCode, sportKey)
	}
	var events []struct {
		Bookmakers []struct {
			Key     string `json:"key"`
			Markets []struct {
				Key      string    `json:"key"`
				Outcomes []Outcome `json:"outcomes"`
			} `json:"markets"`
		} `json:"bookmakers"`
	}
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("decode odds for %s: %w", sportKey, err)
	}
	if len(events) == 0 {
		return nil, fmt.Errorf("no events for %s", sportKey)
	}
	// Prefer Pinnacle, fall back to Betfair Exchange (both sharp).
	out := pickBook(events[0].Bookmakers, "pinnacle")
	if out == nil {
		out = pickBook(events[0].Bookmakers, "betfair_ex_eu")
	}
	if out == nil {
		return nil, fmt.Errorf("no sharp bookmaker (pinnacle/betfair) for %s", sportKey)
	}
	return out, nil
}

func pickBook(bookmakers []struct {
	Key     string `json:"key"`
	Markets []struct {
		Key      string    `json:"key"`
		Outcomes []Outcome `json:"outcomes"`
	} `json:"markets"`
}, key string) []Outcome {
	for _, b := range bookmakers {
		if b.Key == key {
			for _, m := range b.Markets {
				if m.Key == "outrights" && len(m.Outcomes) > 0 {
					return m.Outcomes
				}
			}
		}
	}
	return nil
}

func (c *Client) cacheLookup(sportKey string) ([]Outcome, bool) {
	cf := c.loadCache()
	e, ok := cf.Entries[sportKey]
	if !ok {
		return nil, false
	}
	t, err := time.Parse(time.RFC3339, e.FetchedAt)
	if err != nil || time.Since(t) > c.ttl {
		return nil, false
	}
	return e.Outcomes, true
}

func (c *Client) loadCache() cacheFile {
	cf := cacheFile{Entries: map[string]cacheEntry{}}
	if data, err := os.ReadFile(c.cachePath); err == nil {
		_ = json.Unmarshal(data, &cf)
		if cf.Entries == nil {
			cf.Entries = map[string]cacheEntry{}
		}
	}
	return cf
}

func (c *Client) cacheStore(sportKey string, o []Outcome) {
	cf := c.loadCache()
	cf.Entries[sportKey] = cacheEntry{FetchedAt: time.Now().UTC().Format(time.RFC3339), Outcomes: o}
	data, err := json.MarshalIndent(cf, "", "  ")
	if err != nil {
		return
	}
	_ = os.MkdirAll(filepath.Dir(c.cachePath), 0755)
	_ = os.WriteFile(c.cachePath, data, 0644)
}

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
