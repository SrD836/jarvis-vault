---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-22T04:00:17+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-22T04:00:17+00:00. Ventana: 7 días._

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

### Cambio 1: Reducir `allow_agents` para evitar saturación de contexto

**Archivo:** `openclaw.json` (config global del agente planner)

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
    "archivist"
  ]
}
```

**Razón:** Eliminar `apier`, `skill-reviewer`, `auditor`, `monitor`, `job-hunter` reduce el espacio de decisión del planner, evitando loops de `list_agents`/`get_agent_briefing` que consumen iteraciones sin delegar.

---

### Cambio 2: Añadir límite de herramientas por run en briefing

**Archivo:** `briefing.md` del agente planner (sección "Política de delegación")

**Snippet a añadir al final:**
```markdown
## Límite operativo

- **Máximo de tool_calls por run:** 15
- Si alcanzas 15 tool_calls sin haber delegado → fuerza delegación inmediata al agente más afín.
- Si alcanzas 15 tool_calls y ya delegaste → finaliza con resumen.
- **Prohibido:** llamar a `list_agents` o `get_agent_briefing` más de 2 veces por run.
```

**Razón:** Los runs abortados muestran 20-27 tool_calls con loops de herramientas de descubrimiento; un límite duro fuerza cierre antes de saturar el cap de iteraciones.

---

### Cambio 3: Añadir directiva anti-loop en briefing

**Archivo:** `briefing.md` del agente planner (sección "Delegación real obligatoria")

**Snippet a añadir:**
```markdown
### Anti-loop

- Si detectas que estás leyendo el mismo archivo o llamando a la misma herramienta con los mismos argumentos por segunda vez → **detente inmediatamente** y delega o finaliza.
- Patrón prohibido: `read_file` del mismo path >1 vez en同一 run.
- Si no sabes qué hacer tras 3 tool_calls distintas → delega a `researcher` o `code-reviewer` según la task.
```

**Razón:** El run abortado por "duplicate tool loop (same tool+args 3x)" muestra que el agente entra en ciclos de lectura sin progreso; esta regla corta el patrón de raíz.

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

**Snippet:**
```json
"allow_agents": ["code-reviewer", "researcher", "documenter", "debugger", "tester", "auditor"]
```

**Razón:** Eliminar 5 agentes poco usados (apier, skill-reviewer, archivist, monitor, job-hunter) reduce el árbol de decisión y el contexto que el planner carga en cada run, evitando que el input_token se dispare >200k.

---

### Cambio 2: Añadir límite de contexto en briefing del planner

**Archivo:** `briefing.md` del planner (sección de política de delegación)

**Diff:**
```diff
+ ## Límite de contexto
+ - Si input_tokens > 100k en el primer tool_call, aborta y delega directamente al agente más específico sin leer más archivos.
+ - No cargues más de 3 archivos de contexto por run. Si necesitas más, delega a researcher para que los resuma.
```

**Razón:** El patrón muestra runs con 120k-299k input_tokens. Forzar aborto temprano y delegación evita el bucle de tool calls y el agotamiento de tokens.

---

### Cambio 3: Añadir cap de tool_calls en `caps.env`

**Archivo:** `caps.env` (o `openclaw.json` si es donde se definen caps)

**Snippet:**
```env
MAX_TOOL_CALLS_PER_RUN=15
MAX_INPUT_TOKENS=150000
```

**Razón:** Los runs abortados muestran 16-27 tool_calls. Un cap duro de 15 tool_calls fuerza al planner a delegar antes o abortar rápido, evitando loops duplicados (como el del run 105317 con mismo tool+args 3x).

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
