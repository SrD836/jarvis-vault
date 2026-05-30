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

	// v7 (2026-05-28): edge-gate + Kelly sizing + shadow mode.
	// SizingMode: "kelly_fractional" | "dynamic_equal" | "fixed".
	// MinEdgePoints: required |estimated_prob - implied_price| for approval (E2).
	// Mode: "shadow" → log predictions, no trades. "simulation" → virtual fills.
	SizingMode          string  `json:"sizing_mode"`
	KellyFraction       float64 `json:"kelly_fraction"`
	MaxPerTradePct      float64 `json:"max_per_trade_pct"`
	MaxTotalExposurePct float64 `json:"max_total_exposure_pct"`
	MinEdgePoints       float64 `json:"min_edge_points"`

	// v7 P4: P0 floor relaxation. When a short-horizon candidate has
	// liquidity above PriceFloorShortLiqRelaxUSD, the floor drops to
	// PriceFloorShortLiqRelax. Defaults wired in applyDefaults.
	PriceFloorShortLiqRelax     float64 `json:"price_floor_short_liq_relax"`
	PriceFloorShortLiqRelaxUSD  float64 `json:"price_floor_short_liq_relax_usd"`

	Mode               string  `json:"mode"`

	// ----- Fase 9 (2026-05-30): reglas de veto desde pérdidas reales. -----
	// Todas reversibles: flag a false / umbral a 0 las desactiva sin rebuild.
	// R1 longshot: bloquear entradas por debajo de un precio salvo edge info/arb fuerte.
	LongshotBlockEnabled   bool    `json:"longshot_block_enabled"`
	LongshotPriceThreshold float64 `json:"longshot_price_threshold"` // default 0.10
	LongshotMinEdge        float64 `json:"longshot_min_edge"`        // default 0.15
	// R3 catalyst 24h: bloquear evento público inminente sin edge contemplado.
	CatalystBlockEnabled bool `json:"catalyst_block_enabled"`
	CatalystBlockHours   int  `json:"catalyst_block_hours"` // default 24
	// R4 volumen: expone el umbral antes hardcoded en VetoV1.
	MinVolume24hUSD float64 `json:"min_volume_24h_usd"` // default 50000
	// R2 tamaño/confianza: bloquear bet a confianza máxima con edge especulativo.
	SizeConfidenceGateEnabled bool `json:"size_confidence_gate_enabled"`
	// R5 precedentes: bloquear si la tesis se apoya en < N precedentes.
	MinPrecedents int `json:"min_precedents"` // default 3
	// Asimetría: fallo de infra del LLM => BLOQUEAR (preferir falso negativo).
	LLMFailClosed bool `json:"llm_fail_closed"`
	// G1 corte de pérdidas (Fase B): thesis_stale = contraparte simétrica de
	// target_hit. Cerca de resolución, si la tesis no se materializó (precio no
	// subió hacia EstimatedProb y sigue <= entrada) => cerrar. Respeta "no % stop"
	// (solo cerca de expiry, nunca con tiempo por delante) => cero whipsaw.
	ThesisStaleEnabled bool    `json:"thesis_stale_enabled"`
	ThesisStaleDays    float64 `json:"thesis_stale_days"`   // default 1.0
	ThesisStaleMargin  float64 `json:"thesis_stale_margin"` // default 0.03
}

// ShadowEnabled reports whether the bot must skip real fills and only log
// predictions for calibration. Source-of-truth is the Mode field.
func (c *BotConfig) ShadowEnabled() bool { return c.Mode == "shadow" }

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
	// v6 defaults — KEPT at 0.10 per v7 P4 audit (analytics/audit_p0_floor.py
	// 2026-05-28): trades in band 0.05-0.10 lost $-1363 net under v6 (44% win
	// rate). Removing the floor wholesale would re-expose to those losses.
	// v7 P4 instead introduces a narrow liquidity relaxation: floor drops to
	// 0.05 only for short-horizon candidates with liquidity≥$20k, where
	// deep-book absorption mitigates the spread-driven decay.
	if c.PriceFloor == 0 {
		c.PriceFloor = 0.10
	}
	if c.PriceFloorShortLiqRelax == 0 {
		c.PriceFloorShortLiqRelax = 0.05
	}
	if c.PriceFloorShortLiqRelaxUSD == 0 {
		c.PriceFloorShortLiqRelaxUSD = 20000
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
	// v7 defaults.
	if c.SizingMode == "" {
		c.SizingMode = "kelly_fractional"
	}
	if c.KellyFraction == 0 {
		c.KellyFraction = 0.25
	}
	if c.MaxPerTradePct == 0 {
		c.MaxPerTradePct = 0.05
	}
	if c.MaxTotalExposurePct == 0 {
		c.MaxTotalExposurePct = 0.40
	}
	if c.MinEdgePoints == 0 {
		c.MinEdgePoints = 0.05
	}
	if c.Mode == "" {
		c.Mode = "shadow"
	}
	// Fase 9 numeric defaults (los bool vienen explícitos de config.json;
	// ausencia => false => regla desactivada, fail-safe).
	if c.LongshotPriceThreshold == 0 {
		c.LongshotPriceThreshold = 0.10
	}
	if c.LongshotMinEdge == 0 {
		c.LongshotMinEdge = 0.15
	}
	if c.CatalystBlockHours == 0 {
		c.CatalystBlockHours = 24
	}
	if c.MinVolume24hUSD == 0 {
		c.MinVolume24hUSD = 50000
	}
	if c.MinPrecedents == 0 {
		c.MinPrecedents = 3
	}
	// G1: defaults numéricos (el bool thesis_stale_enabled viene explícito de config.json).
	if c.ThesisStaleDays == 0 {
		c.ThesisStaleDays = 1.0
	}
	if c.ThesisStaleMargin == 0 {
		c.ThesisStaleMargin = 0.03
	}
}

// ComputeTradeSize returns the size in USD for the next trade.
//
// v7: when SizingMode == "kelly_fractional", Kelly criterion is applied with
// the configured fractional multiplier (default ¼ Kelly):
//
//	b  = (1 - marketPrice) / marketPrice   (payoff odds)
//	f* = (estimatedProb * b - (1 - estimatedProb)) / b
//	size = bankroll * KellyFraction * f*
//
// Clamped to [SizeMin, bankroll * MaxPerTradePct]. Returns 0 when f* <= 0
// (no edge — caller must treat as reject, not as min size).
//
// Legacy modes preserved:
//   - "fixed": returns TradeSizeUSD.
//   - "dynamic_equal" (or any other value): bankroll * SizeFraction clamped to [SizeMin, SizeMax].
func (c *BotConfig) ComputeTradeSize(bankroll, estimatedProb, marketPrice float64) float64 {
	switch c.SizingMode {
	case "kelly_fractional":
		if marketPrice <= 0 || marketPrice >= 1 {
			return 0
		}
		b := (1.0 - marketPrice) / marketPrice
		q := 1.0 - estimatedProb
		fStar := (estimatedProb*b - q) / b
		if fStar <= 0 {
			return 0
		}
		size := bankroll * c.KellyFraction * fStar
		capPerTrade := bankroll * c.MaxPerTradePct
		if size > capPerTrade {
			size = capPerTrade
		}
		if size < c.SizeMin {
			size = c.SizeMin
		}
		return size
	case "fixed":
		return c.TradeSizeUSD
	default:
		s := bankroll * c.SizeFraction
		if s < c.SizeMin {
			s = c.SizeMin
		}
		if s > c.SizeMax {
			s = c.SizeMax
		}
		return s
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
