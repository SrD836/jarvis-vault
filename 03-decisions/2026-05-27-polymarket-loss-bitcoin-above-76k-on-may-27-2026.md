---
title: "Polymarket loss — bitcoin-above-76k-on-may-27-2026"
type: decision
date: 2026-05-27
decision: "Trade cerrado en pérdida: Will the price of Bitcoin be above $76,000 on May 27? (PnL -70.91 USD)"
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
sources_used: [claudemax-websearch, claudemax-websearch, claudemax-websearch]
horizon: short
days_held: 0.87
---

# Loss: Will the price of Bitcoin be above $76,000 on May 27?

## Trade

- **Trade ID**: `T-2313280-1779814832575`
- **Slug**: `bitcoin-above-76k-on-may-27-2026`
- **Categoría**: market
- **Entry price**: 0.6800
- **Exit price**: 0.0600
- **Size**: $77.77
- **PnL**: $-70.91 (-91.18%)
- **Days held**: 0.87
- **Horizon**: short
- **Entry timestamp**: 2026-05-26T17:00:32Z
- **Exit timestamp**: 2026-05-27T13:50:01Z
- **Exit reason**: stop_loss

## Fuentes consultadas al aprobar

- **claudemax-websearch**: 
- **claudemax-websearch**: 
- **claudemax-websearch**: 

## Análisis

Evitar apostar a que BTC supere un umbral de precio en horizonte <24h cuando cotiza cerca de ese umbral: la volatilidad intradiaria puede colapsar la probabilidad aunque las fuentes actuales lo muestren por encima.

**Anti-patterns**: same-day-price-target · price-near-threshold · recency-bias-entry · high-vol-short-horizon · geopolitics-short-horizon

## Human notes

_(no se toca por automatización)_
