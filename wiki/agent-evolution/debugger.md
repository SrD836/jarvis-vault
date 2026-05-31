---
title: "Evolution proposal — debugger"
type: agent-evolution
date: 2026-05-31T04:00:37+00:00
agent: "[[agents/debugger]]"
pattern: abort-recurrent
confidence: medium
proposed: true
applied: false
tags: [agent-evolution, debugger, abort-recurrent]
related:
  - "[[agents/debugger]]"
  - "[[00-MOC]]"
---

# Evolución propuesta — `debugger`

_Generado por `hermes/learnings.py` el 2026-05-31T04:00:37+00:00. Ventana: 7 días._

## Patrón: `abort-recurrent`

- Confianza: **medium**
- Ocurrencias: **3** runs
- `abort_reason`: `success_criterion no cumplido`

### Evidencia

- [[agents/debugger/runs/2026-05-29/152813-from-planner]]
- [[agents/debugger/runs/2026-05-29/152826-from-planner]]
- [[agents/debugger/runs/2026-05-29/152800-from-planner]]

### Propuesta

## Diagnóstico

Señal dominante en los 3 runs: **`tool_calls_n: 0`** + abort en **`iter: 1`**. El debugger aborta antes de investigar nada — produce solo texto y no cumple el `success_criterion`. Además, **`model: deepseek-chat`** en runtime contradice el `model_primary: anthropic/claude-opus-4-8` del briefing: la skill `gsd-debug` corre en un modelo barato que devuelve conclusión sin tool use. Causa raíz probable: el agente no sabe qué criterio debe cumplir ni que debe usar herramientas.

## Cambios

**1. Briefing — exigir investigación antes de cualquier verdict/abort**

(a) `agents/debugger/index.md` (o el `.md` del briefing), en "Human notes".

(b) snippet a añadir:
```markdown
### Contrato de ejecución (anti-abort)

- PROHIBIDO abortar en `iter: 1` con `tool_calls_n: 0`. Un debugger que no
  inspecciona nada no ha debugueado. Mínimo 1 tool call (read/grep/run) antes
  de concluir.
- Si falta contexto para reproducir el bug, NO abortes: emite hipótesis +
  qué dato/log falta y devuélvelo como output parcial (no como abort).
- `aborted: true` solo es válido si: tarea fuera de scope de debug, o
  bloqueo externo irrecuperable. Documenta cuál en `abort_reason`.
```
(c) Razón: 0 tool calls es el factor común de los 3 fallos; convierte el abort prematuro en output parcial accionable.

**2. Briefing — explicitar el `success_criterion`**

(a) mismo archivo.

(b) snippet:
```markdown
### Definition of Done (debugger)

Un run cumple `success_criterion` si entrega AL MENOS:
1. Causa raíz identificada **o** hipótesis rankeadas con evidencia citada
   (path:línea / log).
2. Fix propuesto (diff) **o** siguiente paso de instrumentación concreto.
3. Reproducción confirmada o, si no, el dato exacto que falta para reproducir.
```
(c) Razón: el abort recurrente es "criterio no cumplido" pero el agente no tiene el criterio escrito; sin Done explícito no puede satisfacerlo.

**3. openclaw.json — pinnear modelo del debugger (no degradar a deepseek)**

(a) `openclaw.json` (entrada del agente `debugger`).

(b) diff orientativo:
```diff
   "debugger": {
     "model_primary": "anthropic/claude-opus-4-8",
+    "model_override_allowed": false,
+    "skills": {
+      "gsd-debug": { "model": "anthropic/claude-opus-4-8" }
+    },
     "allow_agents": []
   }
```
(c) Razón: el runtime corre `deepseek-chat` pese al briefing; un modelo débil tiende a responder sin tool use y a fallar el criterio.

## Confianza y datos que faltan

Confianza **media**. El cambio 3 asume que la clave que fuerza `deepseek-chat` vive en `openclaw.json` o en la config de la skill `gsd-debug`; **no tengo ese archivo**. Para confirmar necesitaría: (a) `openclaw.json` real (sección `debugger` y routing de skills), y (b) el cuerpo (no solo frontmatter) de un run fallido, para ver si el agente recibió tarea válida o devolvió texto vacío. Si el cuerpo muestra tareas mal formadas desde `planner`, el fix correcto sería en el briefing del **planner**, no del debugger.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
