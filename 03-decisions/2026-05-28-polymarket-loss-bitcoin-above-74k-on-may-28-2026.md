---
title: "Polymarket loss — bitcoin-above-74k-on-may-28-2026"
type: decision
date: 2026-05-28
decision: "Trade cerrado en pérdida: Will the price of Bitcoin be above $74,000 on May 28? (PnL -60.30 USD)"
alternatives:
  - "Haber vetado la tesis antes de entrar"
  - "Cerrar antes (stop más conservador)"
outcome: confirmed_bad
outcome_observed_after_days: 0
tags: [decision, polymarket, bot, loss, post-mortem]
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
horizon: short
days_held: 0.54
---

# Loss: Will the price of Bitcoin be above $74,000 on May 28?

## Trade

- **Trade ID**: `T-2322008-1779899405064`
- **Slug**: `bitcoin-above-74k-on-may-28-2026`
- **Categoría**: market
- **Entry price**: 0.8070
- **Exit price**: 0.1600
- **Size**: $75.21
- **PnL**: $-60.30 (-80.17%)
- **Days held**: 0.54
- **Horizon**: short
- **Entry timestamp**: 2026-05-27T16:30:05Z
- **Exit timestamp**: 2026-05-28T05:25:01Z
- **Exit reason**: stop_loss

## Análisis

Evitar entradas >0.80 en predicciones de precio intraday de crypto sin fuentes de análisis técnico; la volatilidad BTC invalida confianzas altas en ventanas <1 día

**Anti-patterns**: same-day-price-level · no-source-approval · overconfident-high-entry · crypto-short-horizon

## Human notes

_(no se toca por automatización)_
