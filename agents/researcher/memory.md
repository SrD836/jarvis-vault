---
title: "researcher — memory"
type: agent-memory
agent: "[[agents/researcher]]"
created: 2026-05-26
related:
  - "[[agents/archivist/memory]]"
  - "[[agents/debugger/memory]]"
  - "[[agents/main/memory]]"
  - "[[agents/planner/memory]]"
  - "[[agents/researcher]]"
tags: [agent-memory, researcher]
schema_version: 1
# auto-linked 2026-05-27
---


# researcher — memoria persistente

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

