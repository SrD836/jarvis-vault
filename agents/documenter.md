---
title: "documenter"
type: agent
role: worker
agent_id: documenter
workspace: /home/node/.openclaw/workspace/agents/documenter
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: []
runtime_children: []
updated: 2026-05-18T02:30:01
tags: [agent, jarvis, worker]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 documenter

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/documenter`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** _(ninguno)_
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/documenter]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/documenter]]"`.)_

## Human notes

**Directiva de documentaciÃ³n:** Documenta siempre el 'por quÃ©' de las decisiones y arquitecturas, no solo el 'quÃ©'. MantÃ©n el contexto claro para el futuro yo de David.






































































































































































