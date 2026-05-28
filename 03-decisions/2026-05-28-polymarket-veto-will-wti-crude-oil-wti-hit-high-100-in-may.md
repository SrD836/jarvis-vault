---
title: "Polymarket veto — will-wti-crude-oil-wti-hit-high-100-in-may"
type: decision
date: 2026-05-28
decision: "Veto de tesis 'Will WTI Crude Oil (WTI) hit (HIGH) $100 in May?' (rule M1): memoria: slug prefix match; same category (score 0.70)"
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

# Veto: Will WTI Crude Oil (WTI) hit (HIGH) $100 in May?

## Mercado

- **Slug**: `will-wti-crude-oil-wti-hit-high-100-in-may`
- **Categoría**: market
- **Precio YES**: 0.0800
- **Volumen 24h**: 123308.59 USD
- **End date**: 2026-06-01T03:59:59.999Z

## Razón del veto

**Regla aplicada**: `M1`

memoria: slug prefix match; same category (score 0.70)

## Patterns en memoria que contribuyeron

- veto `will-wti-crude-oil-wti-hit-high-105-in-may` score=0.70 (slug prefix match; same category)
- veto `will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678` score=0.70 (slug prefix match; same category)
- veto `microstrategy-sells-any-bitcoin-by-may-31-2026` score=0.40 (same category; same price bucket low)

## Human notes

_(no se toca por automatización)_
