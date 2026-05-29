---
title: Research brief 2026-05-29
type: research-brief
date: 2026-05-29
source: cron/overnight-research-parallel
model: deepseek-chat
scope: vault-local (sin web-search)
related:
  - "[[agents/researcher]]"
tags: [research, overnight, parallel]
---

# Research brief — 2026-05-29

> Research PARALELO autonomo (`hermes/overnight_research.py`, DeepSeek, sin web). FASE 5. Scope: sintesis sobre el vault local; lo no verificable localmente queda marcado.

## Resumen

El análisis de hoy revela dos áreas críticas: fallos recurrentes en la infraestructura nocturna (operaciones no atómicas, falta de timeouts, inyección FTS5 y errores de sintaxis) y un posible sobre-filtrado en las reglas del polymarket-bot, especialmente P0_floor. Ambos problemas tienen mitigaciones identificadas pero requieren verificación de implementación y datos históricos adicionales.

## Por pregunta

### ¿Qué patrones de fallo recurrentes muestran los runs.jsonl y los logs de los crons nocturnos, y qué mitigaciones concretas aplican?

Se identifican cuatro patrones de fallo de severidad ALTA en el vault: (1) `_maybe_rotate` en `run_logger.py` no es atómica entre `os.path.getsize` y `os.replace`, causando pérdida de datos con procesos concurrentes; (2) `llm_call.py` carece de timeout en `_call_ollama` y `_call_claude`, bloqueando el cron indefinidamente; (3) `memory_sqlite.py` no escapa comillas dobles en `_fts_query`, permitiendo inyección FTS5; (4) `curator.py` tiene `check_probation` truncado que genera `SyntaxError`. Las mitigaciones propuestas son bloqueos de archivos, timeouts configurables, escape de comillas dobles y validación de sintaxis previa. Además, `session-export` muestra un fallo recurrente con 81663 ms y estado "error", sin causa detallada en el vault.

### ¿Qué reglas del polymarket-bot (P0–P11) generan más vetos y cuáles podrían estar sobre-filtrando oportunidades?

Según la tabla de patrones en `agents/polymarket-bot/memory.md`, P0_floor genera 2 de 4 vetos totales (doble que cualquier otra regla), seguida de P2 (1 veto por mercado expirado) y P3_low_absolute_liquidity (1 veto). P0_floor usa un umbral fijo de 0.100 sin ajuste por horizonte temporal, lo que sobre-filtra mercados con horizontes largos (ej. -149 días) donde precios bajos son esperables. P2 podría evitarse con mejor pre-filtrado upstream. El umbral exacto de P3_low_absolute_liquidity no está documentado.

## Pendiente de verificar online

- Frecuencia exacta de cada patrón de fallo en runs.jsonl histórico (no solo los 4 casos del vault) y si las mitigaciones ya están desplegadas en producción.
- Causa detallada del error recurrente en `session-export` (81663 ms, status error): si es timeout, corrupción de datos u otro.
- Datos de runs.jsonl y logs de crons nocturnos para validar frecuencias absolutas de vetos del polymarket-bot sobre más muestras.
- Umbral exacto de P3_low_absolute_liquidity y si P0_floor tiene ajuste dinámico por horizonte temporal.

## Respuestas por pregunta (crudo)

### ¿Qué patrones de fallo recurrentes muestran los runs.jsonl y los logs de los crons nocturnos, y qué mitigaciones concretas aplican?
Los patrones de fallo recurrentes en los runs.jsonl y logs de crons nocturnos son: **operaciones no atómicas** que causan pérdida de datos, **falta de timeouts** que bloquean indefinidamente el cron, **inyección FTS5** por falta de escape, y **errores de sintaxis** que impiden ejecutar módulos completos. Las mitigaciones concretas incluyen atomicidad con bloqueos de archivos, timeouts configurables en llamadas LLM, escape de comillas dobles en consultas FTS5, y validación de sintaxis antes de ejecución.

**Evidencia del vault:**
- `Code-Audits/2026-05-29.md`: `run_logger.py` tiene `_maybe_rotate` no atómica entre `os.path.getsize` y `os.replace` → pérdida de datos con procesos concurrentes (severidad ALTA).
- Mismo archivo: `llm_call.py` sin timeout en `_call_ollama` y `_call_claude` → bloqueo indefinido del cron (ALTA).
- Mismo archivo: `memory_sqlite.py` no escapa comillas dobles en `_fts_query` → inyección FTS5 (ALTA).
- Mismo archivo: `curator.py` con `check_probation` truncado → `SyntaxError`, módulo no ejecutable (ALTA).
- `agents/runs-today.md`: `session-export` muestra estado "error" con 81663 ms, indicando fallo recurrente en exportación nocturna.

**Huecos / por verificar:**
- No se especifica si las mitigaciones ya están implementadas o son propuestas; el vault solo lista hallazgos sin plan de acción.
- El error concreto de `session-export` (81663 ms, status error) no tiene causa detallada en el vault; podría ser timeout o corrupción de datos.
- REQUIERE VERIFICAR ONLINE: frecuencia exacta de cada patrón en runs.jsonl histórico y si hay mitigaciones ya desplegadas en producción.

### ¿Qué reglas del polymarket-bot (P0–P11) generan más vetos y cuáles podrían estar sobre-filtrando oportunidades?
Basado en el contexto local, las reglas que generan más vetos son **P0_floor** y **P2**, seguidas de **P3_low_absolute_liquidity**. La regla **P0_floor** podría estar sobre-filtrando oportunidades.

**Evidencia del vault:**

- En `agents/polymarket-bot/memory.md` se registran 4 vetos en la tabla de patrones:
  - **P2**: 1 veto (mercado ya expiró, endDate=2025-12-31, hace 149 días).
  - **P0_floor**: 2 vetos (price floor < 0.100 con horizon -149 días).
  - **P3_low_absolute_liquidity**: 1 veto (liquidity $772 < umbral no especificado).

- **P0_floor** genera el doble de vetos que cualquier otra regla (2 de 4). Además, el umbral de 0.100 parece fijo y no ajustado al horizonte temporal: mercados con horizonte largo (ej. -149 días) pueden tener precios bajos por lejanía, no por falta de valor. Esto sugiere sobre-filtrado.

- **P2** (mercados expirados) es un veto correcto pero podría evitarse con mejor pre-filtrado upstream.

**Huecos / por verificar:**
- No hay datos de runs.jsonl ni logs de crons nocturnos para validar frecuencias absolutas sobre más muestras.
- Umbral exacto de P3_low_absolute_liquidity no está documentado en el contexto.
- No se especifica si P0_floor tiene ajuste dinámico por horizonte; si no lo tiene, es candidata a sobre-filtrado.
