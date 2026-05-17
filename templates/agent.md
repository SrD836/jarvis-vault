---
title: "Template: Agent"
type: template
tags: [template, agent]
---

# Template para `vault/agents/{id}.md` (auto-fill por `hermes/orgchart.py`)

```markdown
---
title: "{agent_id}"
type: agent
role: orchestrator|worker|main
agent_id: {agent_id}
workspace: {path}
model_primary: {model_id}
delegation_mode: prefer|suggest|(default)
allow_agents: [...]
runtime_children: [...]
updated: {ISO}
tags: [agent, jarvis, {role}]
related:
  - "[[../00-MOC]]"
  - "[[agents/index]]"
---

# ðŸ¤– {agent_id}

**Rol:** {role}
**Modelo:** `{model_id}`
**Workspace:** `{workspace}`

## PolÃ­tica de delegaciÃ³n

- **delegationMode:** `prefer`
- **allowAgents (config):** [[worker-a]] [[worker-b]]
- **Hijos runtime (subagent-registry):** [[worker-a]]

## Sesiones recientes

(Dataview block; ver index.md.)

## Human notes

_(David puede escribir aquÃ­. Cron NO sobreescribe este bloque.)_
```
