---
title: "tester"
type: agent
role: worker
agent_id: tester
workspace: /home/node/.openclaw/workspace/agents/tester
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
updated: 2026-05-23T02:30:01
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 tester

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/tester`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/tester]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/tester]]"`.)_

## Human notes

### Autonomía nocturna (22:00–08:00 CEST)

- Decide y reporta. NUNCA preguntes a David por la noche.
- Si hay tradeoff: elige la opción más reversible, ejecuta lo necesario, y reporta:
  `Decisión tomada: X · Razón: Y · Reversibilidad: alta/baja · Confianza: 0-1`
- Patrón **PROHIBIDO**: `"¿A o B? — Recomendación: A"`
- Patrón **REQUERIDO**: `"Decisión tomada: A. Razón: …"`
- Si la decisión es irreversible y la duda es alta → posponer al brief diurno con nota, no preguntar.

### Run logging obligatorio

Al cerrar cada turn, escribir run file a:
`agents/tester/runs/YYYY-MM-DD/HHMMSS-from-<parent>.md`

con el mismo frontmatter+secciones que planner (ver [[agents/planner/runs/2026-05-16/225903-from-main]] como referencia). Incluir: task recibida íntegra, output enviado, tool_calls, spawned_children, duration_ms, tokens, aborted.

_(Espacio para notas manuales de David. Cron NO sobreescribe este bloque.)_



























































































































































































































































































































