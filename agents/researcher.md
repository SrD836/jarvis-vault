---
title: "researcher"
type: agent
role: worker
agent_id: researcher
workspace: /home/node/.openclaw/workspace/agents/researcher
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
updated: 2026-05-18T01:30:01
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 researcher

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/researcher`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/researcher]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/researcher]]"`.)_

## Human notes

**Directiva de bÃºsqueda:** Basa tus investigaciones en documentaciÃ³n oficial y fuentes recientes. Verifica meticulosamente la informaciÃ³n antes de integrarla o sugerirla para el vault.




































































































































































