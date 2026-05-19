package scanner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/davidgnuez/polymarket-veto-bot/scanner/internal/types"
)

const (
	gammaBaseURL = "https://gamma-api.polymarket.com/markets"
	outputPath   = "vault/inbox/trading/candidates.jsonl"
	minVolume    = 50000.0
	httpTimeout  = 30 * time.Second
	userAgent    = "polymarket-veto-bot/1.0"
	maxMarkets   = 500
)

func Run() error {
	fmt.Println("[scanner] starting scan")

	var allCandidates []types.Candidate
	var cursor *string
	page := 0

	for {
		resp, err := fetchMarkets(cursor)
		if err != nil {
			return fmt.Errorf("fetch page %d: %w", page, err)
		}

		for _, m := range resp.Data {
			if !filterMarket(m) {
				continue
			}
			volume, _ := parseVolume(m.Volume24h)
			price, _ := parseVolume(m.CurrentPrice)

			c := types.Candidate{
				ID:              m.ID,
				Slug:            m.Slug,
				Question:        m.Question,
				Category:        m.Tag,
				Volume24h:       volume,
				CurrentPriceYes: price,
				EndDate:         m.EndDate,
				ScannedAt:       time.Now().UTC().Format(time.RFC3339),
			}
			allCandidates = append(allCandidates, c)
		}

		fmt.Printf("[scanner] page %d: %d raw, %d candidates so far\n",
			page, len(resp.Data), len(allCandidates))

		// Stop if no more pages or cap reached
		if resp.Next == nil || *resp.Next == "" || len(allCandidates) >= maxMarkets {
			break
		}
		cursor = resp.Next
		page++
	}

	fmt.Printf("[scanner] scan complete: %d total candidates\n", len(allCandidates))
	if len(allCandidates) == 0 {
		fmt.Println("[scanner] no candidates found, skipping write")
		return nil
	}
	return writeCandidates(allCandidates)
}

// filterMarket returns true if the market passes all filters.
func filterMarket(m types.Market) bool {
	// Only binary (Yes/No) markets
	if m.OutcomeType != "UNARY" {
		return false
	}

	// Volume check: parse error or below threshold → reject
	volume, err := parseVolume(m.Volume24h)
	if err != nil || volume < minVolume {
		return false
	}

	// Closed check (redundant with API param but defensive)
	if m.Closed {
		return false
	}

	return true
}

func fetchMarkets(cursor *string) (*types.GammaAPIResponse, error) {
	client := &http.Client{Timeout: httpTimeout}

	req, err := http.NewRequest("GET", gammaBaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := req.URL.Query()
	q.Set("closed", "no")
	q.Set("tag", "Politics")
	q.Add("tag", "Crypto")
	q.Add("tag", "Sports")
	q.Add("tag", "Tech")
	q.Set("volume_24h_gt", "50000")
	q.Set("sort", "volume_24h")
	q.Set("order", "desc")
	q.Set("limit", "100")
	if cursor != nil && *cursor != "" {
		q.Set("next", *cursor)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", userAgent)

	fmt.Printf("[scanner] fetching %s\n", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var apiResp types.GammaAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &apiResp, nil
}

func parseVolume(raw string) (float64, error) {
	return strconv.ParseFloat(raw, 64)
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
