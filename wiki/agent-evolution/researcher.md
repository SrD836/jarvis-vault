---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-21T04:00:30+00:00
agent: "[[agents/researcher]]"
pattern: iter-cap-saturated
confidence: medium
proposed: true
applied: false
tags: [agent-evolution, researcher, iter-cap-saturated]
related:
  - "[[agents/researcher]]"
  - "[[00-MOC]]"
---

# Evolución propuesta — `researcher`

_Generado por `hermes/learnings.py` el 2026-05-21T04:00:30+00:00. Ventana: 7 días._

## Patrón: `iter-cap-saturated`

- Confianza: **medium**
- Ocurrencias: **3** runs
- `cap_value`: `8`

### Evidencia

- [[agents/researcher/runs/2026-05-19/104458-from-planner]]
- [[agents/researcher/runs/2026-05-19/095210-from-planner]]
- [[agents/researcher/runs/2026-05-19/101232-from-planner]]

### Propuesta

## Cambios propuestos

### Cambio 1: Limitar iteraciones por tarea en el briefing

**Archivo:** `briefing del agente researcher` (sección de política de delegación)

**Snippet a añadir:**

```markdown
## Límites operativos

- **MAX_ITERATIONS_PER_TASK:** 6
- Si alcanzas 6 iteraciones sin respuesta final → ejecuta `run_bash` con `echo "ITER_LIMIT:researcher:$(date -Iseconds)" >> /tmp/iter_overflow.log` y devuelve el mejor resultado parcial con nota explícita de límite alcanzado.
- **Prohibido:** seguir iterando tras el límite aunque queden herramientas por probar.
```

**Razón:** El patrón `iter-cap-saturated` (cap=8) se activa porque el agente no tiene un límite interno y consume todas las iteraciones permitidas. Un límite de 6 deja margen para el overhead del sistema.

---

### Cambio 2: Reducir `max_iterations` en la configuración del agente

**Archivo:** `openclaw.json` (config del agente researcher)

**Diff:**

```json
{
  "agent_id": "researcher",
  "max_iterations": 6,
  "max_tokens_per_turn": 300000
}
```

**Razón:** El valor actual (implícito 8) permite que el agente consuma el tope del sistema. Bajarlo a 6 fuerza al agente a ser más directo o delegar subtareas.

---

### Cambio 3: Añadir tool de auto-límite con abort temprano

**Archivo:** `caps/env` o configuración de herramientas del agente

**Snippet:**

```yaml
tools:
  - name: check_iteration_budget
    command: |
      if [ -f /tmp/researcher_iter_count ]; then
        COUNT=$(cat /tmp/researcher_iter_count)
        if [ "$COUNT" -ge 5 ]; then
          echo "ABORT: iteration budget exhausted"
          exit 1
        fi
      fi
    run_on_every_turn: true
```

**Razón:** El agente no tiene conciencia de cuántas iteraciones lleva. Un contador externo que aborta en la iteración 5 (dejando 1 de margen) evita llegar al cap del sistema.

## Patrón: `tokens-budget-tight`

- Confianza: **medium**
- Ocurrencias: **3** runs
- `threshold`: `133333`

### Evidencia

- [[agents/researcher/runs/2026-05-19/104458-from-planner]]
- [[agents/researcher/runs/2026-05-19/095210-from-planner]]
- [[agents/researcher/runs/2026-05-19/101232-from-planner]]

### Propuesta

## Cambios propuestos

### Cambio 1: Limitar iteraciones máximas en briefing

**Archivo:** `briefing researcher` (sección de configuración)

**Snippet a añadir (al inicio, tras el frontmatter):**

```markdown
## Límites operativos

- **Máximo de iteraciones por turno:** 6 (actual: 8-9, causa abortos por cap)
- **Si alcanzas 6 iteraciones sin respuesta final:** ejecuta `run_bash` con `echo "RESULT: <resumen>"` y finaliza.
- **No uses tool_calls para exploración excesiva:** prioriza `grep_search` sobre `read_file` para archivos >50KB.
```

**Razón:** Los abortos por `cap iterations (8)` y `MAX_TOKENS_PER_USER_TURN` ocurren porque el agente explora demasiado. Fijar un límite inferior y una salida forzada evita el timeout.

---

### Cambio 2: Reducir tokens de entrada con poda de contexto

**Archivo:** `openclaw.json` (config del agente researcher)

**Diff:**

```json
{
  "agent_id": "researcher",
  "model_primary": "anthropic/claude-sonnet-4-6",
  "max_input_tokens": 250000,
  "max_output_tokens": 4000,
  "max_iterations": 6,
  "context_window": 300000,
  "prune_strategy": "aggressive"
}
```

**Razón:** El run abortado por `MAX_TOKENS_PER_USER_TURN (500000)` tenía 434K input tokens. Reducir el límite a 250K fuerza al modelo a ser más conciso y evita el hard cap global.

---

### Cambio 3: Añadir directiva de salida temprana en briefing

**Archivo:** `briefing researcher` (sección "Directiva de búsqueda")

**Snippet a añadir al final de esa sección:**

```markdown
- **Si tras 4 tool_calls no has encontrado respuesta clara:** ejecuta `run_bash echo "RESULT: No encontrado - <razón>"` y finaliza.
- **No acumules contexto:** cada tool_call debe tener un propósito único. Si un `read_file` devuelve >100 líneas, usa `grep_search` en su lugar.
```

**Razón:** Los runs fallidos muestran 21-37 tool_calls antes del aborto. Forzar una salida temprana reduce drásticamente el consumo de tokens y evita el patrón `tokens-budget-tight`.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
