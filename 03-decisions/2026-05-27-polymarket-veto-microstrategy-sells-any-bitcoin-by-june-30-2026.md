---
title: "Polymarket veto — microstrategy-sells-any-bitcoin-by-june-30-2026"
type: decision
date: 2026-05-27
decision: "Veto de tesis 'MicroStrategy sells any Bitcoin by June 30, 2026?' (rule M1): memoria: exact slug match (score 1.00)"
alternatives:
  - "Aprobar tesis y entrar trade simulado"
  - "Vetar y mantener bankroll"
outcome: pending
outcome_observed_after_days: 30
tags: [decision, polymarket, bot, veto, m1]
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
---

# Veto: MicroStrategy sells any Bitcoin by June 30, 2026?

## Mercado

- **Slug**: `microstrategy-sells-any-bitcoin-by-june-30-2026`
- **Categoría**: market
- **Precio YES**: 0.3290
- **Volumen 24h**: 160963.72 USD
- **End date**: 2026-07-01T04:00:00Z

## Razón del veto

**Regla aplicada**: `M1`

memoria: exact slug match (score 1.00)

## Patterns en memoria que contribuyeron

- veto `microstrategy-sells-any-bitcoin-by-june-30-2026` score=1.00 (exact slug match)
- veto `microstrategy-sells-any-bitcoin-by-december-31-2026` score=0.90 (slug prefix match; same category; same price bucket mid)
- veto `microstrategy-sells-any-bitcoin-by-may-31-2026` score=0.70 (slug prefix match; same category)
- veto `will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678` score=0.40 (same category; same price bucket mid)

## Human notes

_(no se toca por automatización)_
