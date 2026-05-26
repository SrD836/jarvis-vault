---
title: "cuantos-agentes-hay-registrados-en-el-sistema — 2026-05-26"
type: session
date: 2026-05-26T19:33:28+02:00
topic: cuantos-agentes-hay-registrados-en-el-sistema
agent: "[[agents/main]]"
model_used: "claudemax/claude-opus-4-7"
tools_used: []
tools_calls_n: 0
tokens: {"input_tokens":10,"output_tokens":1047}
duration_ms: 28475
iterations: 1
aborted: false
spawned_children:[]
related:[]
tags: ["sesion", "main"]
---
## TL;DR
¿cuántos agentes hay registrados en el sistema? → 15 agentes registrados (`vault/agents/index.md`, snapshot 2026-05-26 19:30): **main** · **planner** · **job-hunter** · **polymarket-handler*…
## Prompt usuario
```
¿cuántos agentes hay registrados en el sistema?
```
## Respuesta main
15 agentes registrados (`vault/agents/index.md`, snapshot 2026-05-26 19:30):

**main** · **planner** · **job-hunter** · **polymarket-handler** (orquestadores)
**researcher** · **code-reviewer** · **debugger** · **tester** · **documenter** · **apier** · **archivist** · **auditor** · **monitor** · **designer** · **skill-reviewer** (workers)

Existen además dos `.md` extra en la carpeta (`c1-doctor`, `dashboard-dev`, `polymarket-bot`) que no están en el index canónico — son orfandades o agentes de host/runtime que `orgchart.py` no incluye.
## Tool calls
_(sin tool calls)_
## Human notes
_(añade observaciones manuales aquí; este bloque no se sobreescribe por automatización)_