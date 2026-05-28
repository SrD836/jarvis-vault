// Package predictions implements the v7 calibration registry.
//
// Every brain decision — approve, veto, or shadow-skip — appends one
// Prediction row to predictions.jsonl. When the underlying market resolves,
// the closer backfills Outcome / ResolvedAt so analytics/brier.py can score
// estimated_prob against the realized 0/1 outcome.
//
// The file is the source of truth for Brier scoring, auto-suspend rules
// (S1), and the shadow→simulation release gate.
package predictions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DefaultPath is the JSONL store used by every component unless overridden
// via BRAIN_PREDICTIONS_PATH or EXIT_PREDICTIONS_PATH env vars.
const DefaultPath = "vault/inbox/trading/predictions.jsonl"

// Prediction captures every decision the brain takes on a candidate so the
// Brier-score loop can learn from outcomes.
//
// Outcome is *float64 so that "not resolved yet" is distinguishable from
// "outcome zero". The closer fills Outcome + ResolvedAt + Resolution when
// the upstream market closes.
type Prediction struct {
	Timestamp          string   `json:"timestamp"`
	MarketID           string   `json:"market_id"`
	Slug               string   `json:"slug"`
	Question           string   `json:"question,omitempty"`
	Category           string   `json:"category,omitempty"`
	HorizonDays        int      `json:"horizon_days"`
	Horizon            string   `json:"horizon"`
	ImpliedPrice       float64  `json:"implied_price"`
	EstimatedProb      float64  `json:"estimated_prob"`
	EdgePoints         float64  `json:"edge_points"`
	EdgeType           string   `json:"edge_type,omitempty"`
	EdgeDescription    string   `json:"edge_description,omitempty"`
	ChecklistPassed    bool     `json:"checklist_passed"`
	Decision           string   `json:"decision"` // buy_yes | buy_no | skip | skip_shadow
	SkipReason         string   `json:"skip_reason,omitempty"`
	SizeUSD            float64  `json:"size_usd,omitempty"`
	ThesisInvalidation string   `json:"thesis_invalidation,omitempty"`
	Outcome            *float64 `json:"outcome,omitempty"`
	Resolution         string   `json:"resolution,omitempty"`
	ResolvedAt         string   `json:"resolved_at,omitempty"`
}

// Path resolves the predictions JSONL location from env vars with fallback to DefaultPath.
// Brain reads BRAIN_PREDICTIONS_PATH; exit_monitor reads EXIT_PREDICTIONS_PATH.
func Path(envKey string) string {
	if envKey != "" {
		if v := os.Getenv(envKey); v != "" {
			return v
		}
	}
	if v := os.Getenv("BRAIN_PREDICTIONS_PATH"); v != "" {
		return v
	}
	return DefaultPath
}

// Append serializes p as one JSON line and appends to path. The parent
// directory is created if missing. Errors are returned, not logged — caller
// decides logging policy (brain warns; closer fatals).
func Append(path string, p Prediction) error {
	if path == "" {
		path = DefaultPath
	}
	if p.Timestamp == "" {
		p.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("predictions mkdir: %w", err)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("predictions open: %w", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(p); err != nil {
		return fmt.Errorf("predictions encode: %w", err)
	}
	return nil
}

// BackfillOutcome rewrites every row in path whose MarketID matches, filling
// Outcome / Resolution / ResolvedAt for rows that don't already have an
// outcome. Returns the number of rows patched.
//
// Implementation reads the whole file, mutates matching rows in memory, and
// rewrites atomically via .tmp + rename. Volumes today are <10k rows so a
// streaming-rewrite is not yet necessary.
func BackfillOutcome(path, marketID string, outcome float64, resolution, resolvedAt string) (int, error) {
	if path == "" {
		path = DefaultPath
	}
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("backfill open: %w", err)
	}
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var rows []Prediction
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var p Prediction
		if err := json.Unmarshal(line, &p); err != nil {
			continue
		}
		rows = append(rows, p)
	}
	f.Close()
	if err := sc.Err(); err != nil {
		return 0, fmt.Errorf("backfill scan: %w", err)
	}

	patched := 0
	for i := range rows {
		if rows[i].MarketID != marketID {
			continue
		}
		if rows[i].Outcome != nil {
			continue
		}
		o := outcome
		rows[i].Outcome = &o
		rows[i].Resolution = resolution
		rows[i].ResolvedAt = resolvedAt
		patched++
	}
	if patched == 0 {
		return 0, nil
	}

	tmp := path + ".tmp"
	out, err := os.Create(tmp)
	if err != nil {
		return 0, fmt.Errorf("backfill create tmp: %w", err)
	}
	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)
	for _, row := range rows {
		if err := enc.Encode(row); err != nil {
			out.Close()
			os.Remove(tmp)
			return 0, fmt.Errorf("backfill encode: %w", err)
		}
	}
	if err := out.Close(); err != nil {
		os.Remove(tmp)
		return 0, fmt.Errorf("backfill close tmp: %w", err)
	}
	if err := os.Rename(tmp, path); err != nil {
		return 0, fmt.Errorf("backfill rename: %w", err)
	}
	return patched, nil
}

// Horizon classifies daysToResolution into short/medium/long, mirroring
// BotConfig.Classify but as a free function so the predictions package
// stays import-independent.
func Horizon(days, shortMax, mediumMax int) string {
	if days < 0 {
		return "long"
	}
	if days <= shortMax {
		return "short"
	}
	if days <= mediumMax {
		return "medium"
	}
	return "long"
}

// NormalizeCategory lowercases + trims so Brier groupings stay stable across
// scanner / brain category variants.
func NormalizeCategory(c string) string {
	c = strings.ToLower(strings.TrimSpace(c))
	if c == "" {
		return "uncategorized"
	}
	return c
}
