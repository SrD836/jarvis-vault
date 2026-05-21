---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-21T04:00:15+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-21T04:00:15+00:00. Ventana: 7 días._

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

**Archivo:** `openclaw.json` (config raíz del agente planner)

**Snippet:**
```json
{
  "agent_id": "planner",
  "allow_agents": [
    "code-reviewer",
    "researcher",
    "documenter",
    "debugger",
    "tester",
    "auditor"
  ]
}
```

**Razón:** Eliminar `apier`, `skill-reviewer`, `archivist`, `monitor`, `job-hunter` reduce opciones de delegación que causan loops de decisión. 11 agentes permitidos → 6. El planner se satura iterando sobre qué worker elegir.

---

### Cambio 2: Añadir límite de iteraciones por tool call

**Archivo:** Briefing del agente `planner.md` (sección "Política de delegación")

**Snippet a añadir:**
```markdown
## Límite de iteraciones

- Si tras 3 tool_calls consecutivas con `delegate` no has recibido respuesta del worker → aborta y reporta error.
- Si `iter > 6` y no has producido output final → fuerza respuesta con `"ERROR: iter-cap-saturated, abortando"`.
- Prohibido llamar `list_agents` o `get_agent_briefing` más de 2 veces por run.
```

**Razón:** Los runs abortados muestran `iter: 9-17` con loops de `list_agents`/`get_agent_briefing`. Cortocircuito evita saturación.

---

### Cambio 3: Añadir `runtime_children` predefinidos y cap en caps.env

**Archivo:** `caps.env` (o equivalente de configuración de capacidades)

**Snippet:**
```env
PLANNER_MAX_ITER=6
PLANNER_MAX_DELEGATE_CALLS=3
PLANNER_FORCE_CHILDREN=researcher,code-reviewer,documenter
```

**Razón:** Forzar hijos runtime predefinidos elimina la fase de descubrimiento de agentes que causa loops. El cap `8` actual se excede porque el planner gasta iteraciones decidiendo a quién delegar.

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

### Cambio 1: Limitar `allow_agents` en `openclaw.json` para reducir carga cognitiva

**Archivo:** `openclaw.json` (config del agente planner)

**Snippet:**
```json
{
  "agent_id": "planner",
  "allow_agents": ["researcher", "code-reviewer", "debugger", "tester", "auditor"],
  "max_input_tokens": 80000,
  "max_iterations": 12
}
```

**Razón:** Tiene 11 workers permitidos. Cada uno añade contexto de briefing en cada run. Reducir a 5 esenciales baja el input tokens ~40% y evita saturación.

---

### Cambio 2: Añadir directiva de resumen temprano en el briefing

**Archivo:** `briefing/planner.md` (sección de política de delegación)

**Snippet:**
```markdown
## Control de tokens

- Si tras 3 iteraciones el input acumulado supera 100k tokens → ejecuta `compress_context()` o genera un resumen de lo hecho hasta ahora y continúa.
- Prohibido leer más de 3 archivos de briefing de workers por run. Si necesitas más info, usa `delegate` con instrucción explícita de resumen.
```

**Razón:** Los runs abortados muestran input tokens de 120k-299k. Un checkpoint de compresión a las 3 iteraciones evita el loop de lectura excesiva.

---

### Cambio 3: Configurar `caps.env` con límite de tokens por tool call

**Archivo:** `caps.env` (o `openclaw.json` si soporta caps)

**Snippet:**
```env
MAX_INPUT_TOKENS_PER_RUN=100000
MAX_TOOL_CALLS_PER_RUN=15
COMPRESS_AFTER_ITERATIONS=3
```

**Razón:** El run de 299k tokens con 20 tool calls y loop duplicado se evita con límites duros. El aborto por iteraciones (8) llega tarde; comprimir en iteración 3 corta antes.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
