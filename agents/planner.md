---
title: "planner"
type: agent
role: orchestrator
agent_id: planner
workspace: /home/node/.openclaw/workspace/agents/planner
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: prefer
allow_agents: [code-reviewer, researcher, documenter, apier, skill-reviewer, debugger, tester, auditor, archivist, monitor, job-hunter]
runtime_children: []
updated: 2026-05-18T05:30:02
tags: [agent, jarvis, orchestrator]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 planner

**Rol:** orchestrator
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/planner`

## Política de delegación

- **delegationMode:** `prefer`
- **allowAgents (config):** [[code-reviewer]] [[researcher]] [[documenter]] [[apier]] [[skill-reviewer]] [[debugger]] [[tester]] [[auditor]] [[archivist]] [[monitor]] [[job-hunter]]
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/planner]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/planner]]"`.)_

## Human notes

**Directiva operativa:** Maximiza el paralelismo y delega tareas a los agentes especializados siempre que la complejidad lo justifique. MantÃ©n una estricta trazabilidad de las decisiones en `03-decisions`.












































































































































































