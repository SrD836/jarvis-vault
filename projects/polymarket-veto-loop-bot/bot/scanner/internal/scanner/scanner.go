package scanner

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgnuez/polymarket-veto-bot/scanner/internal/types"
)

const (
	gammaBaseURL = "https://gamma-api.polymarket.com/markets"
	outputPath   = "vault/inbox/trading/candidates.jsonl"
	minVolume    = 50000.0
	httpTimeout  = 30 * time.Second
	userAgent    = "polymarket-veto-bot/1.0"
	maxMarkets   = 500
	pageSize     = 100
)

var popCultureBlacklist = []string{
	"kardashian", "kanye", "drake", "taylor swift", "rihanna", "beyonce",
	"elon musk's", "trump tweet", "celebrity", "oscar", "grammy", "met gala",
	"jesus christ", "second coming", "alien", "bigfoot", "loch ness",
	"hot dog", "pizza", "emoji", "meme",
}

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	log.Printf("[scanner] config loaded: mode=%s sort=%s window=[%d,%d]d",
		cfg.Mode, "end_date_asc", cfg.HorizonShortMaxDays, cfg.HorizonMediumMaxDays)

	fmt.Println("[scanner] starting scan")

	allCandidates := []types.Candidate{}
	now := time.Now().UTC().Format(time.RFC3339)
	offset := 0
	page := 0

	for {
		markets, err := fetchMarkets(offset)
		if err != nil {
			return fmt.Errorf("fetch page %d: %w", page, err)
		}
		if len(markets) == 0 {
			break
		}

		for _, m := range markets {
			c, ok := convert(m, now)
			if !ok {
				continue
			}
			allCandidates = append(allCandidates, c)
			if len(allCandidates) >= maxMarkets {
				break
			}
		}

		fmt.Printf("[scanner] page %d (offset %d): %d raw, %d candidates so far\n",
			page, offset, len(markets), len(allCandidates))

		if len(markets) < pageSize || len(allCandidates) >= maxMarkets {
			break
		}
		offset += pageSize
		page++
	}

	// Sort by endDate ascending so brain sees short-term markets first
	sort.Slice(allCandidates, func(i, j int) bool {
		return parseEndDate(allCandidates[i].EndDate).Before(parseEndDate(allCandidates[j].EndDate))
	})

	fmt.Printf("[scanner] scan complete: %d candidates (sorted by endDate asc)\n", len(allCandidates))
	if len(allCandidates) == 0 {
		fmt.Println("[scanner] no candidates; clearing output file")
	}
	return writeCandidates(allCandidates)
}

func parseEndDate(s string) time.Time {
	if s == "" {
		return time.Date(2099, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t, _ = time.Parse("2006-01-02T15:04:05Z", s)
	}
	return t
}

func convert(m types.Market, now string) (types.Candidate, bool) {
	if m.Closed || !m.Active {
		return types.Candidate{}, false
	}

	var outcomes []string
	if err := json.Unmarshal([]byte(m.Outcomes), &outcomes); err != nil {
		return types.Candidate{}, false
	}
	if len(outcomes) != 2 {
		return types.Candidate{}, false
	}
	if !(strings.EqualFold(outcomes[0], "Yes") && strings.EqualFold(outcomes[1], "No")) {
		return types.Candidate{}, false
	}

	ql := strings.ToLower(m.Question)
	for _, kw := range popCultureBlacklist {
		if strings.Contains(ql, kw) {
			return types.Candidate{}, false
		}
	}

	volume := m.VolumeNum
	if m.Volume24Hr != nil {
		volume = *m.Volume24Hr
	}
	if volume < minVolume {
		return types.Candidate{}, false
	}

	var prices []string
	if err := json.Unmarshal([]byte(m.OutcomePrices), &prices); err != nil || len(prices) < 1 {
		return types.Candidate{}, false
	}
	priceYes := parseFloat(prices[0])
	if priceYes <= 0 || priceYes >= 1 {
		return types.Candidate{}, false
	}

	category := "uncategorized"
	if len(m.Events) > 0 && m.Events[0].Category != "" {
		category = m.Events[0].Category
	}

	return types.Candidate{
		ID:              m.ID,
		Slug:            m.Slug,
		Question:        m.Question,
		Category:        category,
		Volume24h:       volume,
		CurrentPriceYes: priceYes,
		EndDate:         m.EndDate,
		ScannedAt:       now,
	}, true
}

func parseFloat(s string) float64 {
	var v float64
	fmt.Sscanf(s, "%f", &v)
	return v
}

func fetchMarkets(offset int) ([]types.Market, error) {
	client := &http.Client{Timeout: httpTimeout}

	req, err := http.NewRequest("GET", gammaBaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := req.URL.Query()
	q.Set("closed", "false")
	q.Set("active", "true")
	q.Set("volume_num_min", fmt.Sprintf("%.0f", minVolume))
	q.Set("order", "volumeNum")
	q.Set("ascending", "false")
	q.Set("limit", fmt.Sprintf("%d", pageSize))
	q.Set("offset", fmt.Sprintf("%d", offset))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", userAgent)

	fmt.Printf("[scanner] GET %s\n", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var markets []types.Market
	if err := json.NewDecoder(resp.Body).Decode(&markets); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return markets, nil
}

func writeCandidates(candidates []types.Candidate) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	for _, c := range candidates {
		if err := enc.Encode(c); err != nil {
			return fmt.Errorf("encode candidate %s: %w", c.ID, err)
		}
	}
	fmt.Printf("[scanner] wrote %d candidates to %s\n", len(candidates), outputPath)
	return nil
}
