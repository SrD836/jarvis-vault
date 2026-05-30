---
title: "Polymarket — resumen diario 2026-05-29"
type: decision
date: 2026-05-29
decision_type: daily-rollup
tags: [polymarket, bot, daily, rollup]
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
---

# Polymarket — resumen diario 2026-05-29

> Nodo unico del dia (reemplaza el firehose de nodos por evento; el dato bruto vive en `inbox/trading/*.jsonl`, el aprendizaje en `memory.md`).

## Decisiones del dia

- Evaluadas (predictions): **1447**
- Abiertas hoy (approved): **0**
- Cerradas hoy (closed): **13**  ·  P&L dia: **-75.01 USD**  ·  wins 2 / losses 10

### Por tipo de decision

| decision | n |
|---|---|
| skip | 725 |
| buy_no | 567 |
| buy_yes | 155 |

### Top motivos skip/veto

| motivo | n |
|---|---|
| edge no declarado por LLM (edge_type=none) | 97 |
| edge 0.017 < mín 0.030 (p̂=0.020, implied=0.037) | 8 |
| edge 0.020 < mín 0.030 (p̂=0.030, implied=0.050) | 7 |
| edge 0.010 < mín 0.030 (p̂=0.030, implied=0.020) | 6 |
| edge 0.014 < mín 0.030 (p̂=0.080, implied=0.094) | 6 |

### Trades abiertos por horizonte

_ninguno_

### Top 3 cierres (por |P&L|)

| slug | pnl | exit_reason |
|---|---|---|
| will-crude-oil-cl-hit-high-150-by-end-of-june-788-691 | -27.83 | horizon-shortonly-2026-05-29 |
| strait-of-hormuz-traffic-returns-to-normal-by-end-of-june | -21.36 | horizon-shortonly-2026-05-29 |
| israel-x-lebanon-diplomatic-meeting-by-may-31-177 | +8.12 | target_hit |

## Calibracion actual (`calibration.json`)

- short: n=25 wr=1.000 expectancy=+12.5072/trade total_pnl=+312.68
- kelly_shrink=1.0 min_edge_override=0.0 brier=0.1743 (n=33)

## Human notes

_(no se toca por automatizacion)_
