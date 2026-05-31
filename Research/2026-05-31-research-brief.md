---
title: Research brief 2026-05-31
type: research-brief
date: 2026-05-31
source: cron/overnight-research-parallel
model: deepseek-chat
scope: vault-local (sin web-search)
related:
  - "[[01-briefs/2026-05-31]]"
  - "[[Code-Audits/2026-05-31]]"
  - "[[Research/2026-05-30-research-brief]]"
  - "[[agents/polymarket-handler]]"
  - "[[agents/researcher]]"
  - "[[outputs/2026-05-30-nightly-digest]]"
tags: [research, overnight, parallel]
# auto-linked 2026-05-31
---


# Research brief — 2026-05-31

> Research PARALELO autonomo (`hermes/overnight_research.py`, DeepSeek, sin web). FASE 5. Scope: sintesis sobre el vault local; lo no verificable localmente queda marcado.

## Resumen
La infraestructura nocturna sigue mostrando los **cuatro patrones de fallo sistémicos** ya identificados (ya visto el 2026-05-29 y 2026-05-30): telemetría descuadrada, procesos sin timeout, crons que mueren por errores de import/sintaxis y operaciones no atómicas. Ninguno es catastrófico, pero todos derivan de código no defensivo y persisten sin resolver. En el polymarket-bot, el volumen de vetos es alto (968 nodos) pero **no hay desglose por regla** en el contexto local; solo `P0_floor` está señalado como sospechoso de sobre-filtrado. La rentabilidad del bot está bloqueada por estrategia, no por tooling.

## Por pregunta

### ¿Qué patrones de fallo recurrentes muestran los runs.jsonl y los logs de los crons nocturnos, y qué mitigaciones aplican?
Confirmados cuatro patrones (ya visto el 2026-05-29): **(1)** telemetría descuadrada — `run_logger.py` escribe en `~/agent-stack/logs/runs.jsonl` pero `runs_dashboard.py` lee otra ruta → 0 runs reportados; **(2)** procesos colgados sin `timeout=` en `cf.as_completed` y `subprocess.run`; **(3)** crons muertos por `dict | None` (Py3.9) y typo `r.erro` en `llm_call.py`; **(4)** race conditions y loops sin try/except. Mitigaciones: path canónico único, timeouts, fijar Python ≥3.10 + lint pre-deploy, escritura atómica (tmp+rename) y try/except por iteración. Síntoma nuevo persistente: `claude(cron-blocked)` y "1 error sobre 77 runs" aún sin identificar.

### ¿Qué reglas del polymarket-bot (P0–P11) generan más vetos y cuáles sobre-filtran?
No hay desglose numérico por regla en el contexto, así que no se puede rankear sin inventar. Documentado: **968 nodos veto** agregados (de 1639 residuo del bot) y un stack multinivel (S1+P0-P11+V1-V6+M1-M3+N1-N2+E1-E2+R1/R3/R5). Único señalado como sobre-filtrado: **`P0_floor`** (ya visto el 2026-05-30). El encadenamiento de capas hace el sobre-filtrado un riesgo estructural.

## Pendiente de verificar online
- El contenido real de `runs.jsonl` **no está en el contexto**: los patrones se infieren del audit, no de las líneas JSONL.
- El **"1 error sobre 77 runs" sigue sin identificar** en el vault.
- **Versión de Python en producción del VPS** no confirmada (clave para el patrón 3).
- **Ranking real de vetos P0-P11**: requiere parsear `03-decisions/*veto*` o `runs.jsonl` por código de regla (es trabajo sobre el vault local, no online).
- **Definición exacta de `P0_floor`** (umbral mínimo) para confirmar si su corte es demasiado agresivo.
- **Tasa de oportunidades que pasan todas las capas vs. vetadas**, para medir sobre-filtrado acumulado del stack completo.

## Respuestas por pregunta (crudo)

### ¿Qué patrones de fallo recurrentes muestran los runs.jsonl y los logs de los crons nocturnos, y qué mitigaciones concretas aplican?
## Respuesta directa
El vault documenta **cuatro patrones de fallo recurrentes y sistémicos** (ninguno catastrófico) en la infraestructura nocturna, todos derivados de código no defensivo:

1. **Telemetría descuadrada / runs invisibles.** `run_logger.py` escribe a `~/agent-stack/logs/runs.jsonl` pero `runs_dashboard.py` lee `~/.openclaw/cron/runs/*.jsonl`. Resultado visible: `agents/runs-today.md` reporta **0 runs** pese a haber actividad (`Code-Audits/2026-05-31.md`, ALTA). **Mitigación:** unificar la ruta de escritura/lectura a un único path canónico.

2. **Procesos colgados sin timeout.** `parallel_loop.py` usa `cf.as_completed` sin timeout → worker colgado bloquea el orquestador y acumula zombies en el VPS; `run_logger.py` llama `subprocess.run` sin `timeout=`. **Mitigación:** añadir `timeout=` a `as_completed` y a todo `subprocess.run`.

3. **Crons que mueren silenciosamente por errores de import/sintaxis.** `orgchart.py` usa `dict | None` que rompe en Python 3.9 (mata el cron sin traza); `llm_call.py` tiene typo `r.erro` → `AttributeError` en cada invocación CLI. **Mitigación:** fijar/validar versión de Python (≥3.10 o `from __future__ import annotations`), corregir typo, y CI/lint pre-deploy.

4. **Operaciones no atómicas y loops sin aislamiento.** Race condition read-modify-write en `memory_inject.py` sobre `AGENTS.md`; `memory_backfill.py` sin try/except → un agente fallido aborta los siguientes. **Mitigación:** escritura atómica (lock/tmp+rename) y `try/except` por iteración.

`cron_overnight_research.md` añade el síntoma observado `claude(cron-blocked)` (1 run del digest) y "1 error sobre 77 runs sin identificar".

## Evidencia del vault
- `Code-Audits/2026-05-31.md` — hallazgos ALTA/MEDIA por archivo.
- `outputs/memory/cron_overnight_research.md` — confirma 4 patrones, `cron-blocked`, 1/77.
- `agents/runs-today.md` — síntoma: 0 runs reportados.

## Huecos / por verificar
- **REQUIERE VERIFICAR ONLINE/LOCAL:** el contenido real de `runs.jsonl` no está en el contexto; los patrones se infieren del audit, no de las líneas JSONL.
- El "1 error sobre 77 runs" sigue **sin identificar** en el vault.
- Versión de Python en producción del VPS no confirmada.

### ¿Qué reglas del polymarket-bot (P0–P11) generan más vetos y cuáles podrían estar sobre-filtrando oportunidades?
## Respuesta directa
El contexto **no incluye el desglose numérico de vetos por regla individual** (P0, P1, … P11), así que no puedo rankear con cifras cuál genera más vetos sin inventar. Lo único señalado de forma explícita como sospechoso de **sobre-filtrado es `P0_floor`**.

Lo que sí está documentado:
- **Volumen agregado** de vetos: **968 nodos `veto`** auto-generados en `03-decisions/` (de 1639 nodos residuo del bot). Es el conteo de decisiones de veto, no su atribución por regla.
- La capa de veto es **rica y multinivel**: `S1 + P0-P11 + V1-V6 + M1-M3 + N1-N2 + E1-E2 + R1/R3/R5 + suspensión por categoría` (`audit/2026-05-30-polymarket-bot-audit.md`). Con tantas capas encadenadas, el riesgo de sobre-filtrado acumulado es estructural.

## Evidencia del vault
- `outputs/memory/cron_overnight_research.md`: apunta a **"posible sobre-filtrado en las reglas del polymarket-bot, especialmente P0_floor"**.
- `audit/AUDIT_OBSIDIAN_BOT.md`: **968 nodos veto** (parte del ~49.7% de ruido del vault); confirma el volumen de vetos pero no su origen por regla.
- `audit/2026-05-30-polymarket-bot-audit.md`: enumera el stack de reglas P0-P11 y concluye que la rentabilidad está **bloqueada por estrategia, no por tooling** — coherente con que un filtrado excesivo mate oportunidades sin que falle el código.
- `research_finding (2026-05-30)`: confirma que el desglose por regla **no está en el contexto local**.

## Huecos / por verificar
- **REQUIERE DATOS DEL VAULT (no online):** parsear `03-decisions/*veto*` o el `runs.jsonl` para contar vetos **por código de regla** y obtener el ranking real P0-P11.
- Definición exacta de `P0_floor` (umbral de probabilidad/precio mínimo) para confirmar si su corte es demasiado agresivo.
- Tasa de oportunidades que pasan todas las capas vs. vetadas, para medir sobre-filtrado acumulado del stack completo.
