---
title: "tester run — 2026-05-29 152245"
type: agent-run
date: 2026-05-29T15:22:45+00:00
agent: "[[agents/tester]]"
spawned_from: "[[agents/planner]]"
parent_chain: ["main", "planner"]
model: deepseek-chat
provider: deepseek
iter: 1
tokens: {"input_tokens": 71, "output_tokens": 127, "estimated": true}
tool_calls_n: 0
duration_ms: 2626
aborted: false
abort_reason: ""
spawned_children: []
skill: "test-driven-development"
related:
  - "[[agents/CATALOG]]"
  - "[[agents/code-reviewer/runs/2026-05-29/152234-from-planner]]"
  - "[[agents/debugger/runs/2026-05-29/152243-from-planner]]"
  - "[[agents/planner]]"
  - "[[agents/tester]]"
tags: ["agent-run", "tester", "planner-routed"]
# auto-linked 2026-05-29
---


## TL;DR
tester: 5 casos de test con entrada, salida esperada y tipo de caso. -> OK.

## Task recibida
```
Propon 5 casos de test para la funcion parse_price(s) que convierte un string tipo '1.234,56 EUR' a float 1234.56. Incluye edge cases como formato invalido, separadores mixtos, moneda faltante, etc. Describe cada caso con entrada, salida esperada y tipo de caso (normal, borde, error).
```

## Output
1. **Caso normal**: entrada `'1.234,56 EUR'`, salida esperada `1234.56`, tipo: normal.
2. **Caso borde**: entrada `'0,00 EUR'`, salida esperada `0.0`, tipo: borde.
3. **Caso error**: entrada `'abc'`, salida esperada `None` o excepción (según diseño), tipo: error.
4. **Caso separadores mixtos**: entrada `'1,234.56 EUR'` (punto como separador de miles y coma decimal), salida esperada `1234.56`, tipo: borde.
5. **Caso moneda faltante**: entrada `'1.234,56'`, salida esperada `None` o excepción, tipo: error.

## Human notes
_(no se toca por automatizacion)_
