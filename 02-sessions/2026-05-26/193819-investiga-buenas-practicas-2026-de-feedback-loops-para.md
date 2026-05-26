---
title: "investiga-buenas-practicas-2026-de-feedback-loops-para — 2026-05-26"
type: session
date: 2026-05-26T19:38:19+02:00
topic: investiga-buenas-practicas-2026-de-feedback-loops-para
agent: "[[agents/main]]"
model_used: "claudemax/claude-opus-4-7"
tools_used: []
tools_calls_n: 0
tokens: {"input_tokens":9,"output_tokens":8787}
duration_ms: 269365
iterations: 1
aborted: false
spawned_children:[]
related:[]
tags: ["sesion", "main"]
---
## TL;DR
Investiga buenas prácticas 2026 de feedback loops para agentes LLM y prepárame un brief con 3 patrones aplicables a JARVIS → Brief en `vault/01-briefs/2026-05-26-feedback-loops-llm-agents.md`. Tres patrones priorizados con mecanismo concreto sobre tu stack actual (…
## Prompt usuario
```
Investiga buenas prácticas 2026 de feedback loops para agentes LLM y prepárame un brief con 3 patrones aplicables a JARVIS
```
## Respuesta main
Brief en `vault/01-briefs/2026-05-26-feedback-loops-llm-agents.md`. Tres patrones priorizados con mecanismo concreto sobre tu stack actual (vault, crons, los 12 workers):

1. **Verification gate en handoffs** (NO Reflexion — evidencia 2025 de que degrada).
2. **PLAN.md externalizado + drift monitor** (anti goal-drift en runs largos).
3. **Regression flywheel sobre traces clusterizados** (gate empírico para skill promotion — cubre el hueco semántico de `meta/curator`).

Orden de implementación recomendado: #2 → #3 → #1 (este último necesita acumular traces antes). Si decides adoptar #2, el siguiente paso natural es un decision log en `vault/03-decisions/`.
## Tool calls
_(sin tool calls)_
## Human notes
_(añade observaciones manuales aquí; este bloque no se sobreescribe por automatización)_