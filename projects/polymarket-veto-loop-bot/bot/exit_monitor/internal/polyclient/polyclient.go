package polyclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

const gammaBase = "https://gamma-api.polymarket.com"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// GetMarketPrice fetches a single market from Gamma API and returns current yes price.
// Returns price, closed bool, error.
func GetMarketPrice(marketID string) (float64, bool, error) {
	url := fmt.Sprintf("%s/markets/%s", gammaBase, marketID)
	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, false, fmt.Errorf("gamma GET %s: %w", marketID, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, false, fmt.Errorf("gamma GET %s: HTTP %d: %s", marketID, resp.StatusCode, string(body))
	}

	var mkt types.MarketPrice
	if err := json.NewDecoder(resp.Body).Decode(&mkt); err != nil {
		return 0, false, fmt.Errorf("gamma decode %s: %w", marketID, err)
	}

	price, err := strconv.ParseFloat(mkt.CurrentPrice, 64)
	if err != nil {
		return 0, false, fmt.Errorf("gamma parse price %s: %w", mkt.CurrentPrice, err)
	}

	return price, mkt.Closed, nil
}
