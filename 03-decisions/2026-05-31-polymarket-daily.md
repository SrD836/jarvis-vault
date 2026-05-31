---
title: "Polymarket — resumen diario 2026-05-31"
type: decision
date: 2026-05-31
decision_type: daily-rollup
tags: [polymarket, bot, daily, rollup]
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
---

# Polymarket — resumen diario 2026-05-31

> Nodo unico del dia (reemplaza el firehose de nodos por evento; el dato bruto vive en `inbox/trading/*.jsonl`, el aprendizaje en `memory.md`).

## Decisiones del dia

- Evaluadas (predictions): **798**
- Abiertas hoy (approved): **8**
- Cerradas hoy (closed): **0**  ·  P&L dia: **+0.00 USD**  ·  wins 0 / losses 0

### Por tipo de decision

| decision | n |
|---|---|
| skip | 751 |
| buy_no | 34 |
| buy_yes | 13 |

### Top motivos skip/veto

| motivo | n |
|---|---|
| S1: categoría suspendida por Brier (cat=sports-season) | 339 |
| edge no declarado por LLM (edge_type=none) | 20 |
| edge 0.009 < mín 0.030 (p̂=0.010, implied=0.001) | 17 |
| R5 precedentes: 2 < 3 casos análogos | 10 |
| R5 precedentes: 0 < 3 casos análogos | 10 |

### Trades abiertos por horizonte

| horizonte | n |
|---|---|
| long | 3 |
| medium | 3 |
| short | 2 |

### Top 3 cierres (por |P&L|)

_ninguno_

## Calibracion actual (`calibration.json`)

- short: n=40 wr=0.600 expectancy=-9.3088/trade total_pnl=-372.35
- kelly_shrink=0.8108 min_edge_override=0.04 brier=0.1944 (n=406)

## Human notes

_(no se toca por automatizacion)_
