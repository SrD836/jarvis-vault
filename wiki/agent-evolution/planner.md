---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-23T04:00:17+00:00
agent: "[[agents/planner]]"
pattern: iter-cap-saturated
confidence: medium
proposed: true
applied: false
tags: [agent-evolution, planner, iter-cap-saturated]
related:
  - "[[agents/planner]]"
  - "[[00-MOC]]"
---

# Evolución propuesta — `planner`

_Generado por `hermes/learnings.py` el 2026-05-23T04:00:17+00:00. Ventana: 7 días._

## Patrón: `iter-cap-saturated`

- Confianza: **medium**
- Ocurrencias: **4** runs
- `cap_value`: `8`

### Evidencia

- [[agents/planner/runs/2026-05-19/095210-from-main]]
- [[agents/planner/runs/2026-05-19/105919-from-main]]
- [[agents/planner/runs/2026-05-19/105317-from-main]]
- [[agents/planner/runs/2026-05-19/104519-from-main]]

### Propuesta

Basado en el patrón detectado, propongo los siguientes cambios:

## Cambio 1: Limitar `allow_agents` en `openclaw.json`

**Archivo:** `openclaw.json` (config del agente planner)

**Snippet:**
```json
{
  "agent_id": "planner",
  "allow_agents": ["code-reviewer", "researcher", "documenter", "apier", "debugger", "tester", "auditor", "archivist", "monitor"],
  "max_children": 3,
  "max_iterations": 6
}
```

**Razón:** Reducir de 11 a 9 agentes permitidos (eliminar `skill-reviewer` y `job-hunter` que no aparecen en runs) y añadir `max_iterations: 6` para abortar antes del límite de 8 iteraciones.

## Cambio 2: Añadir directiva de aborto temprano en briefing

**Archivo:** `agents/planner/briefing.md` (sección Política de delegación)

**Snippet:**
```markdown
### Límite de iteraciones

- Si tras 5 iteraciones no has delegado al menos una subtarea → aborta con `abort_reason: "no_delegation_in_5_iters"`.
- Si detectas que estás re-leyendo el mismo archivo >2 veces seguidas → aborta con `abort_reason: "duplicate_tool_loop"`.
- Máximo 6 iteraciones totales por run.
```

**Razón:** Los runs fallidos muestran 7-17 iteraciones sin delegar o con loops duplicados; un aborto temprano evita saturación de capacidad.

## Cambio 3: Forzar delegación en tareas no-triviales via `caps.env`

**Archivo:** `caps.env` (variables de entorno del agente)

**Snippet:**
```env
PLANNER_FORCE_DELEGATE=true
PLANNER_MIN_DELEGATIONS_PER_RUN=1
PLANNER_MAX_ITER_BEFORE_DELEGATE=3
```

**Razón:** El run `121033-from-main` tuvo 0 hijos pero 27 tool_calls; forzar al menos 1 delegación por run y abortar si no delega antes de 3 iteraciones evita autoejecución.

## Patrón: `tokens-budget-tight`

- Confianza: **medium**
- Ocurrencias: **4** runs
- `threshold`: `133333`

### Evidencia

- [[agents/planner/runs/2026-05-19/121033-from-main]]
- [[agents/planner/runs/2026-05-19/105919-from-main]]
- [[agents/planner/runs/2026-05-19/105317-from-main]]
- [[agents/planner/runs/2026-05-19/104519-from-main]]

### Propuesta

## Cambios propuestos

### Cambio 1: Limitar `allow_agents` para reducir contexto de planificación

**Archivo:** `openclaw.json` (config del agente planner)

**Diff:**
```diff
- "allow_agents": ["code-reviewer", "researcher", "documenter", "apier", "skill-reviewer", "debugger", "tester", "auditor", "archivist", "monitor", "job-hunter"],
+ "allow_agents": ["researcher", "code-reviewer", "debugger", "tester", "auditor"],
```

**Razón:** 11 agentes disponibles fuerzan al planner a cargar briefings de todos ellos en cada run (120-300k tokens input). Reducir a 5 agentes esenciales baja el contexto un 50-60%.

---

### Cambio 2: Añadir límite de tokens de entrada en briefing

**Archivo:** `briefings/agents/planner.md` (sección de política de delegación)

**Diff:**
```diff
## Política de delegación

- **delegationMode:** `prefer`
+ **inputTokenBudget:** 80000
+ **Si input_tokens > 80000:** abortar y delegar inmediatamente al agente más afín sin más planificación.
- **allowAgents (config):** [[code-reviewer]] [[researcher]] ...
```

**Razón:** El planner consume 120k-300k tokens en leer contexto antes de delegar. Con un hard limit de 80k, se fuerza a delegar rápido sin saturar el contexto.

---

### Cambio 3: Añadir directiva anti-bloqueo por contexto grande

**Archivo:** `briefings/agents/planner.md` (sección Human notes)

**Diff:**
```diff
### Delegación real obligatoria

+ ### Límite de contexto (anti-bloqueo)
+ - Si input_tokens > 80000: NO leas más archivos. Delegar inmediatamente al worker más afín.
+ - Si llevas >5 iteraciones sin delegar: aborta y delega al researcher con la pregunta original.
+ - Prohibido: leer briefings de workers que no vas a usar en esta run.

- Cuando recibas una task que encaja con un worker ...
```

**Razón:** Los runs fallidos muestran 7-17 iteraciones sin delegar, consumiendo tokens en leer contexto de workers no relevantes. Esta directiva corta el loop temprano.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
