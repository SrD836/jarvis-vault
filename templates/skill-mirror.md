---
title: "Template: Skill mirror"
type: template
tags: [template, skill]
---

# Template para `vault/skills/{category}/{name}.md` (espejo de `~/.openclaw/skills/...`)

```markdown
---
title: "{skill_name}"
type: skill-mirror
source: "~/.openclaw/skills/{category}/{name}/SKILL.md"
category: meta|coding|research|docs|...
state: active|stale|probation|archived
pinned: true|false
use_count: {N}
view_count: {N}
last_activity_at: {ISO}
owner_agent_ids: ["[[agents/...]]"]
created_in_session: "[[02-sessions/...]]"
evolved_from: "[[skills/.../v0]]"
tags: [skill, {category}, {state}]
related:
  - "[[../../00-MOC]]"
  - "[[../../04-skills-log/index]]"
---

# 🛠️ {skill_name}

(Resumen breve del SKILL.md canónico. La fuente de verdad está en `~/.openclaw/skills/`.)

## Used in
- `[[02-sessions/...]]`

## Lifecycle
- `[[04-skills-log/...-new]]`
- `[[04-skills-log/...-stale]]`
```
