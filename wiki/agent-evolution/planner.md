---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-25T04:00:15+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-25T04:00:15+00:00. Ventana: 7 días._

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

### Cambio 1: Reducir `allow_agents` y forzar delegación temprana

**Archivo:** `openclaw.json` → `agents.planner.allow_agents`

**Snippet:**
```json
"allow_agents": ["researcher", "code-reviewer", "debugger", "tester", "auditor", "archivist"]
```

**Razón:** El planner tiene 11 workers permitidos. En runs con `iter: 16-17` y `spawned_children: []` o solo 1 hijo, el agente pierde iteraciones explorando workers irrelevantes (job-hunter, monitor, skill-reviewer, apier, documenter). Reducir a 6 workers nucleares fuerza delegación real y reduce el espacio de búsqueda.

---

### Cambio 2: Añadir límite de iteraciones con delegación obligatoria

**Archivo:** `briefing/planner.md` → sección "Delegación real obligatoria"

**Snippet:**
```markdown
### Regla de iteraciones

- Si `iter >= 5` y `spawned_children` está vacío → ABORTA inmediatamente con razón "autoejecución detectada (iter N sin delegar)".
- Si `iter >= 8` y solo has delegado 1 hijo → ABORTA con razón "delegación insuficiente (1 hijo en 8 iter)".
- Si `iter >= 12` → ABORTA siempre con razón "iter-cap-saturated (12 iter)".
```

**Razón:** Los runs fallidos muestran `iter: 16` y `iter: 17` sin aborto automático. El agente necesita reglas duras de aborto basadas en iteraciones para no saturar el límite de 8 caps sin respuesta.

---

### Cambio 3: Añadir `max_iterations` en configuración del agente

**Archivo:** `openclaw.json` → `agents.planner`

**Snippet:**
```json
"max_iterations": 8,
"abort_on_iter_cap": true
```

**Razón:** El patrón `iter-cap-saturated` con `cap_value: 8` indica que el sistema aborta por límite de caps, pero el planner no tiene un límite interno propio. Forzar `max_iterations: 8` alinea el agente con el límite del sistema y evita que gaste iteraciones en loops de herramientas duplicadas (como el run con `same tool+args 3x`).

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
{
  "agent_id": "planner",
  "allow_agents": ["code-reviewer", "researcher", "documenter", "debugger", "tester", "auditor"]
}
```

**Razón:** Eliminar 5 agentes (apier, skill-reviewer, archivist, monitor, job-hunter) que no aparecen en los runs fallidos pero inflan el contexto de planificación con briefings innecesarios.

---

### Cambio 2: Añadir límite de tokens de entrada en briefing

**Archivo:** `briefing.md` del agente planner (sección de políticas)

**Snippet:**
```markdown
## Límite de contexto

- **Input tokens máx por run:** 120000
- Si al cargar briefings de workers superas este límite → delega sin leer briefings completos, usa solo `get_agent_briefing` con resumen.
- Si tras 3 tool calls seguidas superas 100k tokens input → aborta y delega la subtarea más grande al worker más afín.
```

**Razón:** Los runs muestran inputs de 120k-299k tokens; un límite explícito fuerza delegación temprana antes de saturar el contexto.

---

### Cambio 3: Forzar delegación en tareas >1 paso

**Archivo:** `briefing.md` (sección "Delegación real obligatoria")

**Snippet:**
```markdown
### Regla de delegación forzosa

- Si la task requiere >3 tool calls o >2 archivos → DELEGAR obligatoriamente a un worker.
- Si `spawned_children: []` y task no es trivial (1 archivo, 1 paso) → error de autoejecución.
- En cada run, antes de tool call #5, si no has delegado aún → delega la parte más pesada.
```

**Razón:** El run abortado por loop (17 iter, 0 hijos) y el de 27 calls sin hijos muestran que planner se queda tareas complejas en vez de delegar, quemando tokens.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
