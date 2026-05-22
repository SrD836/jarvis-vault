---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-22T04:00:35+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-22T04:00:35+00:00. Ventana: 7 días._

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

### 1. Añadir `max_iterations` en el frontmatter del agente

**Archivo:** `agents/researcher/README.md` (o el archivo que contenga el frontmatter YAML)

**Snippet:**
```yaml
---
title: "researcher"
type: agent
role: worker
agent_id: researcher
workspace: /home/node/.openclaw/workspace/agents/researcher
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
max_iterations: 6  # ← NUEVO: límite duro antes de abortar y reportar
updated: 2026-05-22T04:00:01
tags: [agent, jarvis, worker]
---
```

**Razón:** El patrón `iter-cap-saturated` (cap=8) muestra que el agente llega a 8 iteraciones sin respuesta. Fijar `max_iterations: 6` fuerza un aborto temprano con `abort_reason: "max_iterations (6) excedido sin respuesta final"`, evitando consumir tokens hasta el límite global de 500K.

### 2. Añadir directiva de "auto-aborto por iteraciones" en el briefing

**Archivo:** `agents/researcher/README.md` (sección de política de delegación o human notes)

**Snippet:**
```markdown
### Auto-aborto por iteraciones

- Si tras 6 iteraciones de tool_calls no has producido una respuesta final → **ABORTA** automáticamente.
- Escribe run file con `aborted: true` y `abort_reason: "max_iterations (6) sin respuesta final"`.
- NO continúes iterando. NO intentes "una última búsqueda".
- Reporta al llamante: `"Abortado tras 6 iteraciones. Resumen parcial: ..."`
```

**Razón:** El agente no tiene instrucción explícita de abortar por iteraciones; al llegar a 8 (cap) aborta el sistema, pero ya consumió ~435K tokens. Con 6 iteraciones forzadas se aborta antes, ahorrando ~150K tokens por ocurrencia.

### 3. Reducir `allow_agents` a `[]` y añadir `delegation_mode: force`

**Archivo:** `openclaw.json` (config del agente researcher)

**Snippet:**
```json
{
  "agent_id": "researcher",
  "allow_agents": [],
  "delegation_mode": "force",
  "max_iterations": 6,
  "max_tokens_per_turn": 150000
}
```

**Razón:** El agente no delega (allow_agents vacío), pero `delegation_mode: suggest` le permite sugerir delegaciones que nunca se ejecutan, perdiendo iteraciones. Forzar `force` con `allow_agents: []` elimina ese gasto de tokens en sugerencias de delegación imposibles.

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

### 1. Limitar `max_iterations` en briefing del agente

**Archivo:** briefing del agente `researcher` (sección de configuración)

**Snippet a añadir (al inicio del briefing, tras el frontmatter):**

```markdown
## Límites operativos

- **max_iterations:** 6 (si no hay respuesta final tras 6 tool calls, abortar y reportar resumen parcial)
- **max_input_tokens_per_run:** 250000 (evitar que el contexto crezca hasta el límite global de 500k)
```

**Razón:** Los runs abortados muestran 9, 12 y 19 iteraciones. Fijar un tope de 6 evita que el agente consuma tokens en loops de búsqueda sin progreso.

---

### 2. Configurar `allow_agents` para delegar subtareas

**Archivo:** `openclaw.json` (config del agente researcher)

**Diff:**

```json
{
  "agent_id": "researcher",
  "allow_agents": [
    "web-scraper",
    "code-executor",
    "data-analyzer"
  ]
}
```

**Razón:** El researcher actualmente no delega (allow_agents vacío). Al permitir agentes especializados, puede partir tareas grandes en subtareas más pequeñas, reduciendo el contexto por turno y evitando el límite de 500k tokens.

---

### 3. Añadir `max_tokens_per_tool_call` en caps env

**Archivo:** `caps.env` o configuración de entorno del agente

**Snippet:**

```bash
# Límite de tokens por tool_call para evitar acumulación
MAX_TOKENS_PER_TOOL_CALL=50000
# Si una tool_call devuelve más de 50k tokens, truncar automáticamente
TOOL_CALL_TRUNCATE=true
```

**Razón:** Los runs muestran 37 y 21 tool_calls con inputs de 434k y 143k tokens. Limitar cada tool_call a 50k tokens evita que una sola llamada (ej. `read_file` de un archivo grande) dispare el límite global.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
