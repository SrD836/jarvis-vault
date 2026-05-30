---
title: "Memoria — cron:overnight_research"
type: memory-export
agent: cron:overnight_research
generated_at: 2026-05-30T05:00:03Z
source: memory/jarvis_memory.db
tags: [memory-export, jarvis, read-only]
---

# Memoria — `cron:overnight_research`

> ⚠️ GENERADO desde la DB (`memory/jarvis_memory.db`) por `hermes/memory_export.py`. **NO editar** — la fuente de verdad es SQLite (FASE 4). Los cambios manuales se sobrescriben.

Total: **6** memorias.

## research (2)

- `2026-05-30T03:00:57` ## Resumen El análisis de hoy (2026-05-30) confirma dos frentes ya detectados ayer pero los precisa con evidencia del vault. En infraestructura nocturna aparecen **cuatro patrones de fallo sistémicos** (cron de Claude bloqueado, telemetría descuadrada entre fuentes, 1 error sobre 77 runs sin identif…
- `2026-05-29T10:38:28` ## Resumen El análisis de hoy revela dos áreas críticas: fallos recurrentes en la infraestructura nocturna (operaciones no atómicas, falta de timeouts, inyección FTS5 y errores de sintaxis) y un posible sobre-filtrado en las reglas del polymarket-bot, especialmente P0_floor. Ambos problemas tienen m…

## research_finding (4)

- `2026-05-30T03:00:35` ## Respuesta directa Con el contexto disponible se observan **cuatro patrones de fallo recurrentes**, ninguno catastrófico pero todos sistémicos: 1. **Cron de Claude bloqueado** (`claude(cron-blocked)`): aparece como modelo en 1 run del digest. Es un cron que no llega a ejecutar el modelo Claude — p…
- `2026-05-30T03:00:35` ## Respuesta directa El contexto local que se me dio **no contiene el desglose de vetos por regla P0–P11**, ni las definiciones de esas reglas. Por tanto **no puedo afirmar cuáles generan más vetos ni cuáles sobre-filtran** sin inventar. Lo que sí está documentado es el **volumen agregado** de vetos…
- `2026-05-29T10:38:23` Los patrones de fallo recurrentes en los runs.jsonl y logs de crons nocturnos son: **operaciones no atómicas** que causan pérdida de datos, **falta de timeouts** que bloquean indefinidamente el cron, **inyección FTS5** por falta de escape, y **errores de sintaxis** que impiden ejecutar módulos compl…
- `2026-05-29T10:38:23` Basado en el contexto local, las reglas que generan más vetos son **P0_floor** y **P2**, seguidas de **P3_low_absolute_liquidity**. La regla **P0_floor** podría estar sobre-filtrando oportunidades. **Evidencia del vault:** - En `agents/polymarket-bot/memory.md` se registran 4 vetos en la tabla de pa…
