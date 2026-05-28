---
title: "Polymarket veto — microstrategy-sells-any-bitcoin-by-june-30-2026"
type: decision
date: 2026-05-28
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
- **Precio YES**: 0.5550
- **Volumen 24h**: 177971.76 USD
- **End date**: 2026-07-01T04:00:00Z

## Razón del veto

**Regla aplicada**: `M1`

memoria: exact slug match (score 1.00)

## Patterns en memoria que contribuyeron

- veto `microstrategy-sells-any-bitcoin-by-june-30-2026` score=1.00 (exact slug match)
- veto `microstrategy-sells-any-bitcoin-by-may-31-2026` score=0.70 (slug prefix match; same category)
- veto `microstrategy-sells-any-bitcoin-by-december-31-2026` score=0.70 (slug prefix match; same category)
- veto `will-bitcoin-dip-to-45000-by-december-31-2026-674-923-755-971-998-525-926-245-316-517-544-589-965-923-986-841-815-224` score=0.40 (same category; same price bucket mid)
- veto `will-bitcoin-dip-to-70k-in-may-2026-438-356-919` score=0.40 (same category; same price bucket mid)

## Human notes

_(no se toca por automatización)_
