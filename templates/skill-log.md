---
title: "Template: Skill log entry"
type: template
tags: [template, skill, hermes]
---

# Template para skill-log

```markdown
---
title: "{skill-name} — {action}"
type: skill-log
date: {YYYY-MM-DDTHH:MM:SSZ}
skill: {skill-name}
action: created|updated|stale|archived
version_before: {x.y.z}
version_after: {x.y.z}
tags: [skill-log, {category}]
related:
  - "[[02-sessions/{date}-{topic}]]"  (sesión que originó el cambio, si aplica)
---

## Cambio

{Describe qué cambió en una frase.}

## Diff frontmatter

```diff
- last_used_at: 2026-04-15
+ last_used_at: 2026-05-13
- use_count: 3
+ use_count: 4
- version: 1.0.0
+ version: 1.0.1
```

## Razón

{Por qué curator.py emitió este log. Auto-detectado o causado por sesión específica.}
```
