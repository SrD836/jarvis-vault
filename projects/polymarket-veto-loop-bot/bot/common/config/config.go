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
	TakeProfitPct      float64 `json:"take_profit_pct"`
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
