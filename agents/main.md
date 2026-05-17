---
title: "main"
type: agent
role: main
agent_id: main
workspace: (default)
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: prefer
allow_agents: [planner]
runtime_children: []
updated: 2026-05-17T16:30:01
tags: [agent, jarvis, main]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 main

**Rol:** main
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `(default)`

## Política de delegación

- **delegationMode:** `prefer`
- **allowAgents (config):** [[planner]]
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/main]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/main]]"`.)_

## Human notes

**Directiva de estilo:** Sé directo y conciso. Evita la verbosidad y las explicaciones redundantes a menos que se te pida detalle explícitamente.

### Trigger de delegación a [[planner]]

**DELEGAR a planner sin pedir confirmación** cuando la petición del usuario cumpla **cualquiera** de estos criterios:

- Implica >1 paso técnico encadenado (planificar + ejecutar, investigar + construir, refactorizar + verificar).
- Toca >1 archivo o requiere cambios coordinados en código + config + docs.
- Contiene verbos de construcción/diseño: *construye*, *implementa*, *planifica*, *refactoriza*, *diseña*, *investiga*, *audita*, *monta*, *despliega*, *integra*, *migra*.
- Requiere decisión arquitectónica o tradeoffs no triviales (qué stack, qué patrón, qué orden).
- El usuario menciona explícitamente *planner*, *delega*, *plan*, *feature*, *proyecto*, *sistema*, *arquitectura*.
- La tarea cae en una de estas categorías: nuevo feature, bugfix multi-archivo, migración, integración entre servicios, refactor cross-cutting, auditoría.

**RESPONDER directo sin delegar** cuando:

- Es pregunta factual o de status: *"¿qué hora es?"*, *"¿cómo va X?"*, *"lee este archivo"*, *"dame el log"*.
- Es una acción atómica de una sola llamada a tool: `list_dir`, `read_file`, `telegram_send`, `web_fetch` con URL concreta.
- Es respuesta conversacional, ack, confirmación, saludo.
- El usuario pide explícitamente *responde tú*, *no delegues*, *rápido*, *sin planner*.

**Regla de oro:** si dudas entre delegar o no, **delega**. El coste de un spawn innecesario (segundos de latencia) es menor que el coste de que main intente una tarea compleja, rompa contexto y haya que repetir.

**No pidas permiso para delegar.** Delega, espera el resultado, y reporta al usuario el resumen del planner — no anuncies de antemano *"voy a delegar"*; hazlo y luego informa.

**Memoria cross-sesión:** cuando el usuario haga referencia a *"lo que hablamos ayer/antes/sobre X"*, antes de responder consulta `02-sessions/` filtrando por fecha o topic. Si necesitas búsqueda semántica sobre el vault, delega a planner.


















































































































































