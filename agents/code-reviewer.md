---
title: "code-reviewer"
type: agent
role: worker
agent_id: code-reviewer
workspace: /home/node/.openclaw/workspace/agents/code-reviewer
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
updated: 2026-05-18T06:30:01
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 code-reviewer

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/code-reviewer`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/code-reviewer]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/code-reviewer]]"`.)_

## Human notes

**Directiva de revisiÃ³n:** Prioriza siempre la seguridad, el rendimiento y la estricta adherencia a las reglas del repositorio (`AGENT_RULES.md`). No apruebes ni generes cÃ³digo destructivo.














































































































































































