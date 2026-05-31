---
title: "Polymarket — resumen diario 2026-05-30"
type: decision
date: 2026-05-30
decision_type: daily-rollup
tags: [polymarket, bot, daily, rollup]
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
---

# Polymarket — resumen diario 2026-05-30

> Nodo unico del dia (reemplaza el firehose de nodos por evento; el dato bruto vive en `inbox/trading/*.jsonl`, el aprendizaje en `memory.md`).

## Decisiones del dia

- Evaluadas (predictions): **4355**
- Abiertas hoy (approved): **0**
- Cerradas hoy (closed): **20**  ·  P&L dia: **-589.96 USD**  ·  wins 4 / losses 16

### Por tipo de decision

| decision | n |
|---|---|
| skip | 2972 |
| buy_no | 1078 |
| buy_yes | 305 |

### Top motivos skip/veto

| motivo | n |
|---|---|
| edge no declarado por LLM (edge_type=none) | 209 |
| edge 0.001 < mín 0.030 (p̂=0.002, implied=0.003) | 36 |
| R5 precedentes: 0 < 3 casos análogos | 34 |
| R5 precedentes: 2 < 3 casos análogos | 31 |
| edge 0.009 < mín 0.030 (p̂=0.010, implied=0.001) | 27 |

### Trades abiertos por horizonte

_ninguno_

### Top 3 cierres (por |P&L|)

| slug | pnl | exit_reason |
|---|---|---|
| will-apple-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | -169.63 | market_closed |
| ucl-psg-ars-2026-05-30-psg | -99.77 | market_closed |
| fif-bih-mac-2026-05-29-bih | -94.51 | market_closed |

## Calibracion actual (`calibration.json`)

- short: n=40 wr=0.600 expectancy=-9.3088/trade total_pnl=-372.35
- kelly_shrink=0.8108 min_edge_override=0.04 brier=0.1944 (n=406)

## Human notes

_(no se toca por automatizacion)_
