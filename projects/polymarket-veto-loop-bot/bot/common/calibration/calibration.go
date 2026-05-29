// Package calibration loads the adaptive layer written daily by
// analytics/daily_calibration.py. It is an OVERRIDE on top of the static
// config.json defaults: config.json stays the documented baseline, while this
// file carries learned adjustments (Kelly shrink, edge-gate bump) recomputed
// from realized expectancy + Brier.
//
// A missing or unreadable file yields the zero value, which both consumers
// (brain edge gate, executor sizing) treat as "no adjustment" — shrink 1.0 and
// no edge override. The bot therefore behaves exactly like the config defaults
// until the daily cycle has enough data to justify a change. The Python writer
// never edits config.json (design rule); it only writes this file.
package calibration

import (
	"encoding/json"
	"os"
)

// DefaultPath mirrors the trading data dir used by the rest of the bot.
const DefaultPath = "vault/inbox/trading/calibration.json"

// Calibration is the on-disk schema of calibration.json.
type Calibration struct {
	GeneratedAt           string  `json:"generated_at"`
	WindowDays            int     `json:"window_days"`
	KellyShrink           float64 `json:"kelly_shrink"`
	MinEdgePointsOverride float64 `json:"min_edge_points_override"`
	Rationale             string  `json:"rationale"`
}

// Load reads calibration.json from CALIBRATION_PATH (env) or DefaultPath. Any
// error (missing file, bad JSON) returns a zero-value Calibration: the adaptive
// layer is best-effort and never blocks the pipeline.
func Load() Calibration {
	path := os.Getenv("CALIBRATION_PATH")
	if path == "" {
		path = DefaultPath
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Calibration{}
	}
	var c Calibration
	if err := json.Unmarshal(data, &c); err != nil {
		return Calibration{}
	}
	return c
}

// Shrink returns the Kelly shrink factor clamped to (0,1]. A missing or
// out-of-range value means "no shrink" (1.0), so an empty or stale file never
// zeroes out position sizing.
func (c Calibration) Shrink() float64 {
	if c.KellyShrink <= 0 || c.KellyShrink > 1 {
		return 1.0
	}
	return c.KellyShrink
}

// ShrinkProb pulls the model probability toward the market-implied price by the
// shrink factor: pAdj = implied + shrink*(p - implied). shrink=1 is a no-op;
// shrink<1 trusts the LLM edge less, which shrinks Kelly size accordingly.
func (c Calibration) ShrinkProb(p, implied float64) float64 {
	s := c.Shrink()
	return implied + s*(p-implied)
}
