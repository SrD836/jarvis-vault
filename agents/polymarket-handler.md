---
title: "polymarket-handler"
type: agent
role: orchestrator
agent_id: polymarket-handler
workspace: /home/node/.openclaw/workspace/agents/polymarket-handler
model_primary: anthropic/claude-opus-4-8
delegation_mode: suggest
allow_agents: [researcher]
runtime_children: []
updated: 2026-05-31T14:30:02
tags: [agent, jarvis, orchestrator]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 polymarket-handler

**Rol:** orchestrator
**Modelo:** `anthropic/claude-opus-4-8`
**Workspace:** `/home/node/.openclaw/workspace/agents/polymarket-handler`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** [[researcher]]
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/polymarket-handler]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/polymarket-handler]]"`.)_

## Human notes

_(Espacio para notas manuales de David. Cron NO sobreescribe este bloque.)_



































































































































































































































