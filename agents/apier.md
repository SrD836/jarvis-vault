---
title: "apier"
type: agent
role: worker
agent_id: apier
workspace: /home/node/.openclaw/workspace/agents/apier
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: (default)
allow_agents: []
runtime_children: []
updated: 2026-05-17T17:30:01
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 apier

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/apier`

## Política de delegación

- **delegationMode:** `(default)`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/apier]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/apier]]"`.)_

## Human notes

**Directiva de integraciones:** Verifica rigurosamente los rate limits de cualquier API. Prioriza siempre el failover a alternativas gratuitas o locales para evitar costes inesperados.




















































































































































