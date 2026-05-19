package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type BotConfig struct {
	ScannerIntervalMin     int     `json:"scanner_interval_min"`
	ExitMonitorIntervalMin int     `json:"exit_monitor_interval_min"`

	// Horizon classification + quota (v3 — sustituye a short_term_window_days).
	// Un candidato se clasifica por días hasta resolución:
	//   short:  days <= HorizonShortMaxDays
	//   medium: HorizonShortMaxDays < days <= HorizonMediumMaxDays
	//   long:   days > HorizonMediumMaxDays  (o EndDate vacío)
	// Las quotas suman 1.0 y guían al executor sobre el reparto del max_new_trades_per_day.
	HorizonShortMaxDays  int     `json:"horizon_short_max_days"`
	HorizonMediumMaxDays int     `json:"horizon_medium_max_days"`
	HorizonQuotaShort    float64 `json:"horizon_quota_short"`
	HorizonQuotaMedium   float64 `json:"horizon_quota_medium"`
	HorizonQuotaLong     float64 `json:"horizon_quota_long"`

	V2VetoUnderHours   int     `json:"v2_veto_under_hours"`
	MaxExposureUSD     float64 `json:"max_exposure_usd"`
	MaxSameCategory    int     `json:"max_same_category"`
	MaxNewTradesPerDay int     `json:"max_new_trades_per_day"`
	TradeSizeUSD       float64 `json:"trade_size_usd"`
	StopLossUSD        float64 `json:"stop_loss_usd"`
	StopLossPct        float64 `json:"stop_loss_pct"`
	TakeProfitPct      float64 `json:"take_profit_pct"`

	// v4: dynamic position sizing + max open positions cap.
	MaxOpenPositions int     `json:"max_open_positions"`
	SizeMode         string  `json:"size_mode"`     // "dynamic_equal" | "fixed"
	SizeFraction     float64 `json:"size_fraction"` // fraction of bankroll per trade in dynamic_equal mode
	SizeMin          float64 `json:"size_min_usd"`
	SizeMax          float64 `json:"size_max_usd"`

	// v5: liquidity gating — reject trade if visible orderbook < tradeSize × this ratio.
	LiquidityMinRatio float64 `json:"liquidity_min_ratio"`

	// v6 (2026-05-19): hardening post-mortem Discord trade.
	// Hard-coded brain rules with config defaults so policy is source-of-truth in code,
	// not in memory.md prompt context.
	PriceFloor              float64 `json:"price_floor"`                 // default 0.10
	PriceCeiling            float64 `json:"price_ceiling"`               // default 0.95
	MinAbsoluteLiquidityUSD float64 `json:"min_absolute_liquidity_usd"`  // default 5000
	PreEventVetoMinDays     int     `json:"pre_event_veto_min_days"`     // default 7

	Mode               string  `json:"mode"`
}

func Load() (*BotConfig, error) {
	path := os.Getenv("BOT_CONFIG_PATH")
	if path == "" {
		path = "bot/config.json"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}
	var cfg BotConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	cfg.applyDefaults()
	return &cfg, nil
}

func (c *BotConfig) applyDefaults() {
	if c.HorizonShortMaxDays == 0 {
		c.HorizonShortMaxDays = 2
	}
	if c.HorizonMediumMaxDays == 0 {
		c.HorizonMediumMaxDays = 30
	}
	if c.HorizonQuotaShort == 0 && c.HorizonQuotaMedium == 0 && c.HorizonQuotaLong == 0 {
		c.HorizonQuotaShort = 0.70
		c.HorizonQuotaMedium = 0.15
		c.HorizonQuotaLong = 0.15
	}
	if c.MaxOpenPositions == 0 {
		c.MaxOpenPositions = 50
	}
	if c.SizeMode == "" {
		c.SizeMode = "dynamic_equal"
	}
	if c.SizeFraction == 0 {
		c.SizeFraction = 1.0 / 50.0
	}
	if c.SizeMin == 0 {
		c.SizeMin = 25
	}
	if c.SizeMax == 0 {
		c.SizeMax = 500
	}
	if c.StopLossPct == 0 {
		c.StopLossPct = 80
	}
	if c.LiquidityMinRatio == 0 {
		c.LiquidityMinRatio = 4
	}
	// v6 defaults.
	if c.PriceFloor == 0 {
		c.PriceFloor = 0.10
	}
	if c.PriceCeiling == 0 {
		c.PriceCeiling = 0.95
	}
	if c.MinAbsoluteLiquidityUSD == 0 {
		c.MinAbsoluteLiquidityUSD = 5000
	}
	if c.PreEventVetoMinDays == 0 {
		c.PreEventVetoMinDays = 7
	}
}

// ComputeTradeSize returns the size in USD for the next trade given current bankroll.
// In dynamic_equal mode: bankroll * SizeFraction, clamped to [SizeMin, SizeMax].
// In fixed mode: TradeSizeUSD.
func (c *BotConfig) ComputeTradeSize(bankroll float64) float64 {
	if c.SizeMode == "fixed" {
		return c.TradeSizeUSD
	}
	s := bankroll * c.SizeFraction
	if s < c.SizeMin {
		s = c.SizeMin
	}
	if s > c.SizeMax {
		s = c.SizeMax
	}
	return s
}

// Classify returns "short", "medium" or "long" based on days-to-resolution.
// Pass <0 (e.g. EndDate empty) → "long".
func (c *BotConfig) Classify(daysToResolution int) string {
	if daysToResolution < 0 {
		return "long"
	}
	if daysToResolution <= c.HorizonShortMaxDays {
		return "short"
	}
	if daysToResolution <= c.HorizonMediumMaxDays {
		return "medium"
	}
	return "long"
}
