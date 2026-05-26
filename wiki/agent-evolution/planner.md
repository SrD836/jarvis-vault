---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-26T04:00:16+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-26T04:00:16+00:00. Ventana: 7 días._

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

### 1. Limitar `allow_agents` para reducir saturación de iteraciones

**Archivo:** `agents/planner/openclaw.json`

**Cambio:**
```json
{
  "allow_agents": ["researcher", "code-reviewer", "debugger", "tester", "documenter"]
}
```

**Razón:** 11 agentes disponibles fuerza al planner a iterar evaluando opciones de delegación. Reducir a 5 especialistas nucleares elimina el loop de decisión y baja el ratio de iter/cap.

---

### 2. Añadir directiva de aborto temprano por iteraciones

**Archivo:** `agents/planner/briefing.md` (sección "Política de delegación")

**Cambio:**
```markdown
## Límite de iteraciones

- Si alcanzas iter 5 sin haber delegado al menos una subtarea → ABORTA inmediatamente con `abort_reason: "no delegation by iter 5"`.
- Si detectas que estás llamando al mismo tool con los mismos args >1 vez → ABORTA con `abort_reason: "duplicate tool loop"`.
```

**Razón:** Los 4 fallos muestran iter 7-17 sin delegación efectiva o con loops duplicados. Un aborto temprano evita consumir 100k+ tokens.

---

### 3. Forzar delegación obligatoria en tareas >1 paso

**Archivo:** `agents/planner/briefing.md` (sección "Delegación real obligatoria")

**Cambio:**
```markdown
### Regla de delegación forzosa

- Toda tarea que requiera >1 tool call o >1 archivo DEBE delegarse a un worker.
- Si no hay worker adecuado → usa `researcher` como fallback genérico.
- Prohibido: ejecutar más de 3 tool calls seguidas sin delegar.
```

**Razón:** Los runs abortados muestran `spawned_children: []` en tareas complejas (polymarket bot). Forzar delegación rompe el ciclo de autoejecución.

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
- "allow_agents": ["code-reviewer", "researcher", "documenter", "apier", "skill-reviewer", "debugger", "tester", "auditor", "archivist", "monitor", "job-hunter"]
+ "allow_agents": ["researcher", "code-reviewer", "debugger", "tester", "auditor"]
```

**Razón:** 11 agentes en `allow_agents` fuerza al planner a cargar briefings de todos ellos en cada run, inflando input_tokens. Reducir a los 5 más usados (basado en `spawned_children` reales) recorta ~50% del contexto base.

---

### Cambio 2: Añadir límite de iteraciones y detección de loops en briefing

**Archivo:** `briefing.md` del planner (sección de política de delegación)

**Snippet a añadir al final:**
```markdown
## Límites operativos

- **Máximo 8 iteraciones por run.** Si no has completado la task en 8 tool calls, aborta con `abort_reason: "max_iterations_exceeded"` y reporta progreso parcial.
- **Detección de loops:** Si invocas el mismo tool con los mismos argumentos 2 veces seguidas, es un loop. Aborta inmediatamente con `abort_reason: "duplicate_tool_loop"`.
- **Budget de tokens:** Si `input_tokens` supera 100,000 en cualquier iteración, delega la subtarea actual a un worker y finaliza tu parte.
```

**Razón:** Los runs fallidos muestran 17 y 10 iteraciones con loops de herramientas idénticas. Un límite explícito corta el ciclo antes de quemar tokens.

---

### Cambio 3: Forzar delegación temprana en briefing

**Archivo:** `briefing.md` (sección "Delegación real obligatoria")

**Snippet a añadir:**
```markdown
### Regla de delegación temprana

- **Si la task requiere >3 archivos o >2 pasos → delega al worker más afín en la primera iteración.**
- **Prohibido leer más de 5 archivos antes de delegar.** Si necesitas más contexto, delega al researcher para que lo recolecte.
- **Patrón de fallo conocido:** leer 10+ archivos, hacer 5+ tool calls, y luego delegar. Eso quema tokens sin valor. Delegar primero, luego refinar.
```

**Razón:** Los runs con `spawned_children: []` muestran planner ejecutando tareas complejas él mismo. Forzar delegación en iteración 1 reduce drásticamente el consumo de tokens del planner.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
