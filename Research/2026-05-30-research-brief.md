---
title: Research brief 2026-05-30
type: research-brief
date: 2026-05-30
source: cron/overnight-research-parallel
model: deepseek-chat
scope: vault-local (sin web-search)
related:
  - "[[01-briefs/2026-05-30]]"
  - "[[agents/researcher]]"
  - "[[outputs/memory/cron_overnight_research]]"
  - "[[research-queue]]"
  - "[[seeds/2026-05-29-model-routing-matrix]]"
tags: [research, overnight, parallel]
# auto-linked 2026-05-30
---


# Research brief — 2026-05-30

> Research PARALELO autonomo (`hermes/overnight_research.py`, DeepSeek, sin web). FASE 5. Scope: sintesis sobre el vault local; lo no verificable localmente queda marcado.

## Resumen

El análisis de hoy (2026-05-30) confirma dos frentes ya detectados ayer pero los precisa con evidencia del vault. En infraestructura nocturna aparecen **cuatro patrones de fallo sistémicos** (cron de Claude bloqueado, telemetría descuadrada entre fuentes, 1 error sobre 77 runs sin identificar, y generación masiva de nodos-residuo), todos confirmando los fallos de crons ya vistos el 2026-05-29. En el polymarket-bot se confirma un **filtro muy agresivo** (≈1.83 vetos por trade), coherente con la sospecha de ayer de sobre-filtrado, pero **falta la atribución por regla P0–P11**: el contexto entregado no incluye ni las definiciones ni el desglose, así que cualquier ranking sería invención.

## Por pregunta

### ¿Qué patrones de fallo recurrentes muestran runs.jsonl y los logs de crons nocturnos, y qué mitigaciones aplican?
Patrones: `claude(cron-blocked)`, discrepancia 0 vs 77 runs entre `runs-today.md` y el digest, 1 error/77 sin run-id, y ~130–280 nodos/día auto-generados con `## Human notes` vacías (0/300). Mitigaciones: loggear causa del bloqueo con alerta si reaparece, escribir per-turn runs al vault como fuente única, instrumentar el run-id del error, y dejar de crear nodos por evento trivial usando replay desde `inbox/trading/*.jsonl`. **Ya visto el 2026-05-29** (operaciones no atómicas, falta de timeouts) como patrón general; hoy se concreta con evidencia.

### ¿Qué reglas del polymarket-bot (P0–P11) generan más vetos y cuáles sobre-filtran?
No determinable con el contexto entregado: faltan las definiciones P0–P11 y el desglose por `rule_id`. Sí consta volumen agregado: **968 nodos veto** frente a 530 trades (ratio ≈1.83:1), más "+24734 rechazos por cupo" y "126 reglas blandas" (líneas de un fichero, no nodos). El diseño veto-loop con LLM como auditor estricto favorece estructuralmente el sobre-filtrado. **Ya visto el 2026-05-29**: P0_floor y P2 como hipótesis de mayores vetos, pero hoy no se puede confirmar sin los ficheros de reglas.

## Pendiente de verificar online

Nota: los workers marcaron todo lo siguiente como verificable **localmente, NO en web**:

- Abrir el `runs.jsonl` crudo y los logs de `overnight_audit`/`linker` para confirmar si el error 1/77 y el `cron-blocked` son el mismo evento (no hay stacktraces ni run-id en el contexto).
- Sin histórico multi-noche: no se puede afirmar recurrencia temporal más allá del 2026-05-30.
- Localizar definiciones de P0–P11 en `projects/polymarket-veto-loop-bot/` (`rules.go`, config o `RULES.md`).
- Agregar por `rule_id` el frontmatter de los 968 nodos veto en `03-decisions/` para obtener el ranking real por regla.
- Confirmar si "cupo" (24734) y "126 reglas blandas" mapean 1:1 con P0–P11 antes de concluir.
- Medir sobre-filtrado real requiere backtest (trades vetados vs. resultado de mercado), no presente en el vault.

## Respuestas por pregunta (crudo)

### ¿Qué patrones de fallo recurrentes muestran los runs.jsonl y los logs de los crons nocturnos, y qué mitigaciones concretas aplican?
## Respuesta directa

Con el contexto disponible se observan **cuatro patrones de fallo recurrentes**, ninguno catastrófico pero todos sistémicos:

1. **Cron de Claude bloqueado** (`claude(cron-blocked)`): aparece como modelo en 1 run del digest. Es un cron que no llega a ejecutar el modelo Claude — patrón de bloqueo, no de error de lógica.
2. **Discrepancia de telemetría entre fuentes**: `runs-today.md` reporta **0 runs** ("el runtime aún no escribe per-turn runs"), mientras el digest cuenta **77**. Las dos fuentes (cron+vault vs. dashboard) no están reconciliadas → ceguera del dashboard `planner`.
3. **1 run en error / 77** (`results {'ok':76,'error':1}`): tasa baja pero sin identificar cuál es ni si está ligado al `cron-blocked`.
4. **Generación de residuo en bucle**: crons (`cron:linker` 48 runs, `overnight_audit`, `memory_inject`) y el bot producen ~130–280 nodos/día auto-generados con frontmatter idéntico y `## Human notes` vacías (0/300 muestreados) — ruido que ya vive en los `.jsonl`.

## Evidencia del vault

- `claude(cron-blocked)` y `results {'ok':76,'error':1}` → `outputs/2026-05-30-nightly-digest.md`
- "el runtime aún no escribe per-turn runs" / "Total runs: 0" → `agents/runs-today.md`
- "RESIDUO… frontmatter idéntico… `## Human notes` siempre vacía (0/300)" y "~130–280 nodos/día" → `audit/AUDIT_OBSIDIAN_BOT.md`

## Mitigaciones concretas

- **cron-blocked**: añadir log de causa de bloqueo (lock/rate-limit/permiso) y alerta si reaparece >1×/noche.
- **Telemetría**: hacer que el runtime escriba per-turn runs al vault para que `runs_dashboard.py` cuadre con `nightly_digest.py` (fuente única).
- **Error 1/77**: el propio digest pide "identificar cuál de los dos es y si están relacionados" — instrumentar el run-id del error.
- **Residuo**: dejar de crear un nodo por evento trivial; usar replay desde `inbox/trading/*.jsonl` + rotar logs grandes (`rejections.jsonl`, 4.2 MB).

## Huecos / por verificar

- No hay `runs.jsonl` crudo en el contexto: no puedo ver stacktraces ni el run-id exacto del error/bloqueo.
- **REQUIERE VERIFICAR (local, no web)**: abrir el `runs.jsonl` real y los logs de `overnight_audit`/`linker` para confirmar si el error y el `cron-blocked` son el mismo evento.
- Sin histórico multi-noche aquí: no puedo afirmar recurrencia temporal más allá del 2026-05-30.

### ¿Qué reglas del polymarket-bot (P0–P11) generan más vetos y cuáles podrían estar sobre-filtrando oportunidades?
## Respuesta directa

El contexto local que se me dio **no contiene el desglose de vetos por regla P0–P11**, ni las definiciones de esas reglas. Por tanto **no puedo afirmar cuáles generan más vetos ni cuáles sobre-filtran** sin inventar. Lo que sí está documentado es el **volumen agregado** de vetos.

## Evidencia del vault

- `audit/AUDIT_OBSIDIAN_BOT.md`: hay **968 nodos veto** auto-generados (de 1639 nodos del bot en `03-decisions/`, junto a 530 trade y 141 loss). Esto da una proporción cruda **veto:trade ≈ 1.83:1** — casi dos vetos por cada trade ejecutado, señal *cuantitativa* de un filtro agresivo, pero **sin atribución por regla**.
- El dashboard menciona **"126 reglas blandas"** y **"+24734 rechazos por cupo"**, pero la auditoría aclara que son **líneas en un único fichero**, no nodos. El "rechazo por cupo" (límite de posiciones/bankroll) parece ser el mayor generador de descartes por volumen bruto (24.7k), aunque es un *cupo*, no necesariamente una regla P0–P11.
- `agents/polymarket-bot.md`: estrategia "veto-loop", LLM (`deepseek-chat`) como **auditor estricto de reglas, no predictor** — el diseño prioriza vetar ante la duda, lo que estructuralmente favorece sobre-filtrado.

## Huecos / por verificar

1. **Definiciones de P0–P11**: no están en este contexto. Buscar en `projects/polymarket-veto-loop-bot/` (probablemente `rules.go`, `config`, o un `RULES.md`).
2. **Conteo de vetos por regla**: los 968 nodos veto en `03-decisions/` deberían llevar en su frontmatter el `rule_id`. **REQUIERE inspeccionar** esos nodos y agregarlos por `P*` para responder con datos.
3. **Sobre-filtrado real**: medir falsos vetos requiere comparar trades vetados vs. resultado del mercado (backtest), no está en el vault provisto.
4. El "cupo" (24734) y las "126 reglas blandas" pueden no mapear 1:1 con P0–P11; confirmar taxonomía antes de concluir.

No tengo acceso a esos ficheros desde el contexto entregado; sin ellos, cualquier ranking P0–P11 sería invención.
