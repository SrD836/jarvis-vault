---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-24T04:00:14+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-24T04:00:14+00:00. Ventana: 7 días._

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

## Cambios propuestos

### Cambio 1: Limitar `allow_agents` en `openclaw.json`

**Archivo:** `openclaw.json` (config del agente planner)

**Snippet:**
```json
{
  "agent_id": "planner",
  "allow_agents": ["code-reviewer", "researcher", "documenter", "apier", "debugger", "tester", "auditor", "archivist", "monitor"],
  "max_iterations": 6
}
```

**Razón:** Eliminar `skill-reviewer` y `job-hunter` (no usados en runs fallidos) reduce el espacio de búsqueda en `list_agents`/`get_agent_briefing` que causó loops de tool calls duplicadas.

### Cambio 2: Añadir directiva anti-loop en briefing

**Archivo:** `briefing.md` del planner (sección "Política de delegación")

**Diff:**
```diff
+ ## Límite de iteraciones y anti-loop
+ - Si alcanzas iter 6 sin respuesta final → fuerza una decisión con `abort_reason: "forced_decision_at_iter_limit"` y reporta lo que tengas.
+ - Prohibido llamar `list_agents` o `get_agent_briefing` más de 1 vez por run. Si necesitas info de otro agente, usa `delegate` directamente.
+ - Si detectas que estás repitiendo el mismo tool+args (misma herramienta, mismos parámetros) → aborta con `abort_reason: "duplicate_tool_loop"` y reporta.
```

**Razón:** Los 3 runs fallidos muestran `iter: 8-17` con loops de `list_agents`/`get_agent_briefing` repetidos. La directiva fuerza abort temprano y evita saturación.

### Cambio 3: Configurar `max_iterations` en caps env

**Archivo:** `.env` o `caps.env` del workspace planner

**Snippet:**
```env
OPENCLAW_AGENT_MAX_ITERATIONS=6
OPENCLAW_AGENT_MAX_TOOL_CALLS=15
```

**Razón:** El patrón `iter-cap-saturated` con `cap_value: 8` muestra que el límite actual es 8. Bajarlo a 6 fuerza decisión temprana y evita loops de 10-17 iteraciones que consumen tokens sin progreso.

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

Basado en el patrón `tokens-budget-tight` con 4 ocurrencias en 7 días y los runs fallidos, propongo 3 cambios concretos:

## Cambio 1: Limitar contexto inicial del planner

**Archivo:** `briefings/planner.md` (sección de política de delegación)

**Snippet a añadir:**
```markdown
## Control de contexto

- **Límite estricto de input_tokens:** 120000 tokens por run.
- Si el contexto acumulado supera 120K tokens antes de delegar, debes delegar inmediatamente al worker más afín sin leer más archivos.
- Prohibido leer archivos >50KB completos. Usa `read_file` con `maxLines: 50` o `grep_search` para extraer solo lo necesario.
```

**Razón:** Los runs muestran inputs de 190K, 227K, 299K tokens. Poner un límite duro evita que el planner consuma todo el budget antes de delegar.

## Cambio 2: Forzar delegación temprana en tareas complejas

**Archivo:** `briefings/planner.md` (sección "Delegación real obligatoria")

**Snippet a modificar (añadir al final):**
```markdown
- **Regla de 3 pasos:** Si tras 3 tool_calls no has delegado a un worker, aborta y delega forzosamente al worker más afín. No sigas acumulando contexto.
- **Prohibido** ejecutar más de 5 tool_calls sin haber delegado al menos una subtarea.
```

**Razón:** Los runs fallidos muestran 12-27 tool_calls sin hijos (run 121033: 27 calls, 0 hijos). Esto fuerza delegación temprana y reduce consumo de tokens.

## Cambio 3: Configurar límite de iteraciones en openclaw.json

**Archivo:** `openclaw.json` (config del planner)

**Diff:**
```json
{
  "agents": {
    "planner": {
      "max_iterations": 8,
      "max_input_tokens": 120000,
      "early_delegation_threshold": 3
    }
  }
}
```

**Razón:** El run 095210 abortó por "cap iterations (8) sin respuesta final". Poner `max_iterations: 8` como límite duro (ya se alcanzó) y añadir `max_input_tokens: 120000` para cortar antes de que el contexto explote.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
