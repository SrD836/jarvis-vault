---
title: "Polymarket Bot — despliegue autónomo 24/7 en modo simulation"
type: decision
date: 2026-05-20
decision: "Arrancar el pipeline completo (scanner 30min, brain, executor, exit_monitor 5min) en modo simulation de forma autónoma continua sin intervención manual"
alternatives:
  - "Operación manual: David lanza el scanner a demanda y revisa outputs antes de que executor entre trades"
  - "Autonomía plena 24/7 con cron (elegida): pipeline corre solo, David solo monitoriza resultados y P&L"
outcome: pending
outcome_observed_after_days: 7
tags: [decision, polymarket, bot, autonomia, arquitectura]
driven_by_sessions:
  - "[[02-sessions/2026-05-19/121118-configura-el-bot-polymarket-para-operacion-autonoma-247]]"
  - "[[02-sessions/2026-05-19/110241-has-creado-el-agente-dashboard-devmd-perfecto-eso]]"
related:
  - "[[03-decisions/2026-05-19-polymarket-v1-baseline-failure]]"
  - "[[agents/polymarket-bot]]"
  - "[[projects/polymarket-veto-loop-bot/project]]"
  - "[[projects/polymarket-veto-loop-bot/bot/config.json]]"
---

## Contexto

El 2026-05-19 se completó el build del pipeline Go de 4 microservicios (scanner, brain, executor, exit_monitor). Todos compilaron limpio (verificado con `go build` en golang:1.22-alpine). La pregunta era si arrancar en modo manual o dejar correr de forma autónoma.

El fallo v1 (LeBron 2028 + Bernie Sanders, ver [[03-decisions/2026-05-19-polymarket-v1-baseline-failure]]) ocurrió en el primer ciclo manual; David decidió continuar y dejar que el pipeline corriera autónomamente para acumular datos de simulación reales.

## Alternativas

### Opción A: Operación manual por demanda
- Pros: control total, David revisa cada batch antes de ejecutar.
- Cons: sin datos de largo plazo; el objetivo era 24/7 simulation para validar EV real; David tendría que estar disponible constantemente.

### Opción B: Autonomía 24/7 con cron (elegida)
- Pros: acumula historial real de trades simulados sin fricción; permite medir drift de bankroll y calidad del veto-loop; condición necesaria para los 30 días de operación limpia antes de considerar dinero real.
- Cons: si el brain falla silenciosamente, puede acumular pérdidas simuladas sin alerta inmediata; categoría `uncategorized` (null en API) limita el cap de correlación.

## Criterios de evaluación

1. Acumular suficiente historial simulado para evaluar EV real del sistema de reglas.
2. No requerir intervención manual continua de David.
3. Mantener exposure virtual ≤ $5.000 (cap configurado).

## Decisión

**Elegimos autonomía 24/7** porque el objetivo primario es validar el sistema sobre 30 días de operación, lo cual es imposible sin ejecución continua. El modo simulation elimina el riesgo financiero real y permite absorber los fallos del brain sin consecuencias.

Config activa en `vault/projects/polymarket-veto-loop-bot/bot/config.json`:
- `scanner_interval_min: 30`
- `exit_monitor_interval_min: 5`
- `mode: simulation`
- `max_exposure_usd: 5000`
- `size_fraction: 0.02` (2% bankroll por trade, subido desde 1.5% del diseño original)
- `take_profit_pct: 40` / `stop_loss_pct: 80`

Estado al cierre de sesión 2026-05-20 05:00: bankroll virtual $2.633, used_exposure $3.441 (~50 posiciones abiertas).

## Outcome observado

(Rellenar tras 7 días — evaluar si el sistema de reglas genera P&L positivo en simulation y si hay fallos sistémicos en el brain.)
