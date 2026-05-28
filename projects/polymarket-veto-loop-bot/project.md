---
name: Polymarket Veto-Loop Bot
description: Bot de trading simulado en Polymarket con auditor LLM de reglas duras y bankroll virtual
status: active
created: "2026-05-19T09:33:02.685Z"
updated: "2026-05-19T09:33:02.685Z"
related:
  - "[[03-decisions/2026-05-19-polymarket-v1-baseline-failure]]"
  - "[[agents/auditor]]"
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot/design]]"
tags: []
# auto-linked 2026-05-19
---


# Polymarket Veto-Loop Bot

## Estado (2026-05-28)
Bot pausado en **`mode: shadow`** tras post-mortem v6 (semanas perdiendo capital — sin calibración, sin edge declarado, cuotas por nº de trades). Re-operación condicional a `analytics/brier.py` confirmando Brier < 0.20 en ≥1 categoría con ≥20 muestras resueltas. Toggle `mode → simulation` manual del user. Detalle: `design.md` sección "Cambios duros v7".

Acciones requeridas en VPS tras este merge:
1. `git pull && (cd bot && go build ./...)` + rebuild containers brain/executor/exit_monitor.
2. Quitar `force_close_horizon_excess` del crontab.
3. `go run ./bot/exit_monitor/cmd/force_close_all/ --apply` (cierra todo, reason=manual_reset_v7).
4. Esperar 7 días. Leer `vault/inbox/trading/shadow_ready.json`. Si `ready=true` y Brier < 0.20 → `config.json: "mode": "simulation"` + deploy.

## Objetivo
Construir un bot de trading sobre mercados de Polymarket que opere inicialmente en modo **simulación** (sin dinero real) con un bankroll virtual de **10.000 USD**. La tesis central, tomada del artículo de referencia, es que un LLM es **mal predictor pero buen revisor de reglas**: por eso el sistema separa la generación de tesis del filtrado, dejando al LLM el rol de auditor estricto que veta cualquier idea que no cumpla criterios duros.

## Alcance
Este proyecto cubre **únicamente la fase simulada**. El paso a dinero real está condicionado a 30 días consecutivos de operación sin incidentes y queda **explícitamente fuera** del alcance actual.

### Universo de mercados
- **Incluidos:** Politics, Crypto, Sports, Tech
- **Excluidos:** Pop Culture (drama mediático atrae counterparty informado y degrada el edge)

## Arquitectura (4 procesos)
1. **Scanner** — Escanea mercados del universo permitido y genera tesis candidatas.
2. **Brain (veto-loop)** — Auditor LLM que aplica reglas duras y bloquea tesis sospechosas. Solo deja pasar las limpias.
3. **Executor virtual** — Aplica sizing Kelly conservador al **1,5% del bankroll** por trade. No mueve dinero real.
4. **Exit monitor** — Cierra posiciones con **stop-loss −30%** y **take-profit +60%**.

## Reglas duras del auditor (veto)
El Brain rechaza cualquier tesis que incumpla alguna de:
- Volumen del mercado **< 50.000 USD**
- Catalizador previsto en las próximas **24h** (no operar en ventana de catálisis)
- Triggers vagos tipo "cuando suba el sentimiento" o equivalentes no medibles
- **Chasing**: el precio ya se movió **≥ 8% en las últimas 4h**
- Patrones históricos respaldados por **menos de 3 casos**

## Stack / Recursos
- **Logging:** JSONL estructurado en vault/inbox/trading/
- **Dashboard:** nueva página **Bot** en el sidebar con P&L acumulado y listado de trades simulados
- **Bankroll virtual inicial:** 10.000 USD
- **Sizing:** Kelly conservador 1,5% por trade
- **Stop / Take-profit:** −30% / +60%

## Fuera de alcance
- Ejecución con dinero real
- Integración con cuentas o wallets de producción
- Estrategias fuera del universo permitido