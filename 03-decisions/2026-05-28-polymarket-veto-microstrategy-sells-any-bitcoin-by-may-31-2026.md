---
title: "Polymarket veto — microstrategy-sells-any-bitcoin-by-may-31-2026"
type: decision
date: 2026-05-28
decision: "Veto de tesis 'MicroStrategy sells any Bitcoin by May 31, 2026?' (rule M1): memoria: exact slug match (score 1.00)"
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

# Veto: MicroStrategy sells any Bitcoin by May 31, 2026?

## Mercado

- **Slug**: `microstrategy-sells-any-bitcoin-by-may-31-2026`
- **Categoría**: market
- **Precio YES**: 0.1340
- **Volumen 24h**: 250031.40 USD
- **End date**: 2026-07-01T04:00:00Z

## Razón del veto

**Regla aplicada**: `M1`

memoria: exact slug match (score 1.00)

## Patterns en memoria que contribuyeron

- veto `microstrategy-sells-any-bitcoin-by-may-31-2026` score=1.00 (exact slug match)
- veto `microstrategy-sells-any-bitcoin-by-june-30-2026` score=0.70 (slug prefix match; same category)
- veto `microstrategy-sells-any-bitcoin-by-december-31-2026` score=0.70 (slug prefix match; same category)
- veto `will-wti-crude-oil-wti-hit-high-100-in-may` score=0.40 (same category; same price bucket low)
- veto `will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832` score=0.40 (same category; same price bucket low)

## Human notes

_(no se toca por automatización)_
