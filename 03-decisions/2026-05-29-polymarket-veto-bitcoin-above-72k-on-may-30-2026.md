---
title: "Polymarket veto — bitcoin-above-72k-on-may-30-2026"
type: decision
date: 2026-05-29
decision: "Veto de tesis 'Will the price of Bitcoin be above $72,000 on May 30?' (rule V3 Trigger vago, V5 Patrón débil, V6 Sin catalyst): V3 Trigger vago, V5 Patrón débil, V6 Sin catalyst: V3 Trigger vago: la fech..."
alternatives:
  - "Aprobar tesis y entrar trade simulado"
  - "Vetar y mantener bankroll"
outcome: pending
outcome_observed_after_days: 30
tags: [decision, polymarket, bot, veto, v3 trigger vago, v5 patrón débil, v6 sin catalyst]
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
---

# Veto: Will the price of Bitcoin be above $72,000 on May 30?

## Mercado

- **Slug**: `bitcoin-above-72k-on-may-30-2026`
- **Categoría**: market
- **Precio YES**: 0.9700
- **Volumen 24h**: 77435.70 USD
- **End date**: 2026-05-30T16:00:00Z

## Razón del veto

**Regla aplicada**: `V3 Trigger vago, V5 Patrón débil, V6 Sin catalyst`

V3 Trigger vago, V5 Patrón débil, V6 Sin catalyst: V3 Trigger vago: la fecha es concreta pero el evento 'el precio de Bitcoin esté por encima de $72,000' no es verificable de forma discreta como un evento determinista; depende de fluctuaciones continuas del mercado. Además V5 Patrón débil: no hay fuentes independientes ni precedente análogo para predecir un precio específico en una fecha concreta con alta fiabilidad. V6 Sin catalyst: no hay un evento discreto identificable en los próximos 7 días que mueva el precio de manera predecible.

## Human notes

_(no se toca por automatización)_
