---
title: "debugger"
type: agent
role: worker
agent_id: debugger
workspace: /home/node/.openclaw/workspace/agents/debugger
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
updated: 2026-05-17T16:30:01
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 debugger

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/debugger`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/debugger]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/debugger]]"`.)_

## Human notes

_(Espacio para notas manuales de David. Cron NO sobreescribe este bloque.)_























































