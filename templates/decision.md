---
title: "Template: Decision"
type: template
tags: [template, decision]
---

# Template para decision-log

```markdown
---
title: "{decision-name}"
type: decision
date: {YYYY-MM-DD}
decision: {qué se decidió}
alternatives:
  - "{opción A}"
  - "{opción B}"
outcome: pending|positive|negative|mixed
outcome_observed_after_days: {N}
tags: [decision, {area}]
# F5.3 JARVIS: campos opcionales multi-agente
driven_by_sessions: ["[[02-sessions/...]]"]
involves_agents: ["[[agents/planner]]", "[[agents/code-reviewer]]"]
related:
  - "[[02-sessions/{date}-{topic}]]"
---

## Contexto

{Qué problema/situación llevó a tener que decidir.}

## Alternativas

### Opción A: {nombre}
- Pros: ...
- Cons: ...

### Opción B: {nombre}
- Pros: ...
- Cons: ...

## Criterios de evaluación

1. {criterio principal}
2. {criterio secundario}

## Decisión

**Elegimos {opción}** porque {razón principal}.

## Outcome observado

(Rellenar tras {N} días. Inicialmente `pending`.)
```
