---
title: "planner — memory"
type: agent-memory
agent: "[[agents/planner]]"
created: 2026-05-26
related:
  - "[[agents/auditor/memory]]"
  - "[[agents/main/memory]]"
  - "[[agents/planner]]"
  - "[[inbox/2026-05-26-feedback-loops-brief]]"
  - "[[wiki/agent-evolution/planner]]"
tags: [agent-memory, planner]
schema_version: 1
# auto-linked 2026-05-27
---


# planner — memoria persistente

> Léeme al arrancar cada turn. Escribe vía `memory_lib.py` (NO edición manual).

## Vetos pattern (cap 500, append-only, tail rotation)

| timestamp | task_type | input_pattern | rule | reason |
|---|---|---|---|---|

## Losses pattern (cap 500)

| timestamp | task_id | input | output | mistake | lesson |
|---|---|---|---|---|---|

## Soft rules (auto-generadas tras 3+ ocurrencias de pattern)

| condition | action | confidence | uses |
|---|---|---|---|

## Anti-patterns identificados

