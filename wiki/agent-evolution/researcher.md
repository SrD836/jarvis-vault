---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-24T04:00:27+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-24T04:00:27+00:00. Ventana: 7 días._

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

**Archivo:** `agents/researcher/briefing.md` (sección de política)

**Snippet a añadir al final de "Política de delegación":**

```markdown
## Límite de iteraciones

- **max_iterations_per_task:** 6
- Si alcanzas 6 iteraciones sin respuesta final, envía un resumen parcial con `[PARCIAL]` en el subject y aborta explícitamente.
- NO continúes iterando tras el límite aunque tengas capacidad restante.
```

**Razón:** El agente agota las 8 iteraciones permitidas por el sistema sin producir respuesta final, lo que causa abortos por cap. Un límite interno más bajo fuerza un cierre controlado.

---

### Cambio 2: Reducir `max_iterations` en la configuración del agente

**Archivo:** `openclaw.json` (sección del agente `researcher`)

**Diff:**

```diff
- "max_iterations": 8,
+ "max_iterations": 5,
```

**Razón:** Bajar el límite global de 8 a 5 fuerza al agente a ser más directo y evita que entre en bucles de búsqueda excesivos.

---

### Cambio 3: Añadir directiva de "respuesta mínima viable"

**Archivo:** `agents/researcher/briefing.md` (sección "Directiva de búsqueda")

**Snippet a reemplazar:**

```markdown
## Directiva de búsqueda

- **Objetivo:** respuesta mínima viable en ≤5 iteraciones.
- Si tras 3 tool_calls no tienes un hallazgo concreto, cambia de estrategia (reduce scope, usa grep más específico, o pregunta directamente).
- Prohibido: leer más de 5 archivos sin sintetizar un hallazgo intermedio.
- Si la tarea requiere >5 fuentes, delega subtareas a otro agente o reporta `[SCOPE_EXCEDIDO]`.
```

**Razón:** Los runs muestran 11-37 tool_calls sin respuesta final. Una directiva explícita de "respuesta mínima viable" rompe el bucle de búsqueda excesiva.

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

### Cambio 1: Limitar herramientas de búsqueda para reducir tokens de entrada

**Archivo:** `briefing del agente researcher` (sección de herramientas o directivas)

**Snippet a añadir (al final del briefing, antes de los comentarios):**

```markdown
## Límites de búsqueda

- Para investigaciones que requieran >5 búsquedas `grep_search` o `read_file`, usar primero `list_dir` para acotar.
- Si tras 3 tool_calls consecutivas de lectura/búsqueda no hay avance → reportar hallazgo parcial y abortar.
- **Prohibido**: leer archivos completos >50KB sin necesidad justificada. Usar `head -n 100` vía `run_bash`.
```

**Razón:** Los runs abortados muestran 21-37 tool_calls con `grep_search`, `read_file`, `list_dir` que inflan input_tokens hasta 434K. Acotar búsquedas reduce drásticamente el consumo.

---

### Cambio 2: Configurar `max_input_tokens` en caps/env para researcher

**Archivo:** `openclaw.json` (sección `caps` o `env` del agente researcher)

**Snippet:**

```json
{
  "agent_id": "researcher",
  "caps": {
    "max_input_tokens_per_turn": 250000,
    "max_tool_calls_per_turn": 15,
    "max_iterations": 6
  }
}
```

**Razón:** El límite global de 500K se excede porque researcher no tiene límite propio. Poner 250K fuerza al agente a ser más conciso antes de abortar por el global, y 15 tool_calls evita loops de búsqueda.

---

### Cambio 3: Añadir directiva de "primer tool_call informativo"

**Archivo:** `briefing del agente researcher` (sección de directivas, antes de "Directiva de búsqueda")

**Snippet:**

```markdown
### Regla de primer tool_call

Antes de cualquier búsqueda, ejecuta `run_bash` con:
```bash
echo "Plan: [1-3 pasos concretos]. Archivos clave: [paths]. Tamaño estimado: [KB]"
```
Esto fuerza planificación antes de consumir tokens. Si el plan >3 pasos → abortar y reportar.
```

**Razón:** Los runs abortados muestran 11-37 tool_calls sin un plan visible. Forzar un tool_call inicial de planificación reduce iteraciones innecesarias y tokens de contexto.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
