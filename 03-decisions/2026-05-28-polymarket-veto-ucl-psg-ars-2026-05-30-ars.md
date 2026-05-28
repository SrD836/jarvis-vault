---
title: "Polymarket veto — ucl-psg-ars-2026-05-30-ars"
type: decision
date: 2026-05-28
decision: "Veto de tesis 'Will Arsenal FC win on 2026-05-30?' (rule M1): memoria: slug prefix match; same category; same price bucket mid (score 0.90)"
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

# Veto: Will Arsenal FC win on 2026-05-30?

## Mercado

- **Slug**: `ucl-psg-ars-2026-05-30-ars`
- **Categoría**: other
- **Precio YES**: 0.3100
- **Volumen 24h**: 55367.92 USD
- **End date**: 2026-05-30T16:00:00Z

## Razón del veto

**Regla aplicada**: `M1`

memoria: slug prefix match; same category; same price bucket mid (score 0.90)

## Patterns en memoria que contribuyeron

- veto `ucl-psg-ars-2026-05-30-psg` score=0.90 (slug prefix match; same category; same price bucket mid)
- veto `elon-musk-of-tweets-may-22-may-29-180-199` score=0.40 (same category; same price bucket mid)
- veto `elon-musk-of-tweets-may-22-may-29-160-179` score=0.40 (same category; same price bucket mid)
- veto `will-jannik-sinner-win-the-2026-mens-french-open` score=0.40 (same category; same price bucket mid)
- veto `will-vitality-win-iem-cologne-major-2026` score=0.40 (same category; same price bucket mid)

## Human notes

_(no se toca por automatización)_
