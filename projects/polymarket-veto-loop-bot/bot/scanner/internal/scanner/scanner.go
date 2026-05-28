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

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/categorize"
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

	// v7 P5 mixed sort:
	//   - 30% of slots reserved for short-horizon candidates (endDate ≤ 7d),
	//     ordered endDate asc so executor consumes them under the short quota.
	//   - Remaining slots ordered by volume desc, regardless of endDate.
	//
	// Rationale: legacy sort (pure endDate asc) collapsed the universe to
	// markets resolving today/tomorrow, starving medium/long horizon quotas
	// and giving the brain almost nothing to evaluate beyond same-day noise.
	allCandidates = mixedSort(allCandidates, maxMarkets, 30)

	fmt.Printf("[scanner] scan complete: %d candidates (mixed sort: 30%% short-horizon + 70%% by volume)\n", len(allCandidates))
	if len(allCandidates) == 0 {
		fmt.Println("[scanner] no candidates; clearing output file")
	}
	return writeCandidates(allCandidates)
}

// mixedSort returns up to capN candidates split between two priorities:
//   - shortPct % of slots: candidates with endDate within 7 days, ordered endDate asc
//   - remainder: by Volume24h desc
//
// Short bucket is bounded; excess short candidates spill into the volume bucket.
// Empty/unparsable endDate sorts as far-future and lands in the volume bucket.
func mixedSort(cands []types.Candidate, capN, shortPct int) []types.Candidate {
	cutoff := time.Now().UTC().AddDate(0, 0, 7)
	short := make([]types.Candidate, 0, len(cands))
	rest := make([]types.Candidate, 0, len(cands))
	for _, c := range cands {
		t := parseEndDate(c.EndDate)
		if t.Before(cutoff) {
			short = append(short, c)
		} else {
			rest = append(rest, c)
		}
	}
	sort.Slice(short, func(i, j int) bool {
		return parseEndDate(short[i].EndDate).Before(parseEndDate(short[j].EndDate))
	})
	sort.Slice(rest, func(i, j int) bool {
		return rest[i].Volume24h > rest[j].Volume24h
	})

	shortSlots := capN * shortPct / 100
	if len(short) > shortSlots {
		// Spill excess short candidates into the volume pool so they aren't dropped.
		spill := short[shortSlots:]
		short = short[:shortSlots]
		rest = append(rest, spill...)
		sort.Slice(rest, func(i, j int) bool {
			return rest[i].Volume24h > rest[j].Volume24h
		})
	}

	out := make([]types.Candidate, 0, capN)
	out = append(out, short...)
	out = append(out, rest...)
	if len(out) > capN {
		out = out[:capN]
	}
	return out
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

	// v5: realistic entry pricing — use bestAsk (what we'd actually pay to buy YES).
	// Fallback: outcomePrices[0] if bestAsk missing. Reject if both unavailable.
	var prices []string
	_ = json.Unmarshal([]byte(m.OutcomePrices), &prices)
	priceYes := m.BestAsk
	if priceYes <= 0 && len(prices) >= 1 {
		priceYes = parseFloat(prices[0])
	}
	if priceYes <= 0 || priceYes >= 1 {
		return types.Candidate{}, false
	}
	// v5: drop markets paused or with zero liquidity (cannot trade).
	if !m.AcceptingOrders {
		return types.Candidate{}, false
	}
	if m.LiquidityNum <= 0 {
		return types.Candidate{}, false
	}

	category := ""
	if len(m.Events) > 0 && m.Events[0].Category != "" {
		category = m.Events[0].Category
	}
	if category == "" || category == "uncategorized" {
		// Gamma API returns empty for most markets; infer locally so that
		// softrules clusters and per-category hard rules (P7-P10) work.
		category = categorize.Infer(m.Slug, m.Question)
	}

	return types.Candidate{
		ID:              m.ID,
		Slug:            m.Slug,
		Question:        m.Question,
		Category:        category,
		Volume24h:       volume,
		CurrentPriceYes: priceYes,
		LiquidityUSD:    m.LiquidityNum,
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
