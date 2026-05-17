---
title: "Template: Session"
type: template
tags: [template, sesion]
---

# Template para session-export

```markdown
---
title: "{topic} — {date}"
type: session
date: {YYYY-MM-DDTHH:MM:SS+02:00}
topic: {topic-kebab}
model_used: {model-id}
tools_used: [bash, read, edit, ...]
duration_minutes: {N}
# F5.3 JARVIS: campos opcionales para grafo multi-agente
agent: "[[agents/{agent_id}]]"          # quién ejecutó la sesión (default: [[agents/main]])
spawned_from: "[[02-sessions/...]]"     # si subagent: parent session
spawned_children: ["[[02-sessions/...]]"]  # si planner: hijos spawneados
related_skills: ["[[skills/.../...]]"]
related_decisions: ["[[03-decisions/...]]"]
tags: [sesion, {area}]
related:
  - "[[03-decisions/{date}-{decision}]]"
  - "[[04-skills-log/{date}-{skill}]]"
---

## TL;DR

{3 frases máximo. Qué se hizo, por qué, outcome.}

## Highlights

- **Decisión**: {si aplica, link a decision}.
- **Fix**: {si se resolvió bug}.
- **Skill**: {si se creó/usó skill}.

## Transcript resumido

{Resumen estructurado por turnos, no literal. Omite ack/saludos.}

## Mencionados

- Skills: `[[../skills/<nombre>]]`
- Decisions: `[[../03-decisions/<nombre>]]`
- Archivos tocados: `~/.openclaw/openclaw.json`, etc.
```
