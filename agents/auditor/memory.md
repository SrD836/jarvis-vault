---
title: "auditor — memory"
type: agent-memory
agent: "[[agents/auditor]]"
created: 2026-05-26
related:
  - "[[03-decisions/index]]"
  - "[[agents/auditor]]"
  - "[[agents/main/memory]]"
  - "[[agents/monitor/memory]]"
  - "[[agents/skill-reviewer/memory]]"
tags: [agent-memory, auditor]
schema_version: 1
# auto-linked 2026-05-27
---


# auditor — memoria persistente

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

