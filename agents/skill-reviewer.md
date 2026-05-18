---
title: "skill-reviewer"
type: agent
role: worker
agent_id: skill-reviewer
workspace: /home/node/.openclaw/workspace/agents/skill-reviewer
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
updated: 2026-05-18T05:30:02
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 skill-reviewer

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/skill-reviewer`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/skill-reviewer]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/skill-reviewer]]"`.)_

## Human notes

**Directiva de auditorÃ­a:** MantÃ©n un catÃ¡logo de skills limpio. Sugiere agresivamente la deduplicaciÃ³n, refactorizaciÃ³n y consolidaciÃ³n de skills redundantes.












































































































































































