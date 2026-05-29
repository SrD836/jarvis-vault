---
title: "apier run — 2026-05-29 152430"
type: agent-run
date: 2026-05-29T15:24:30+00:00
agent: "[[agents/apier]]"
spawned_from: "[[agents/planner]]"
parent_chain: ["main", "planner"]
model: deepseek-chat
provider: deepseek
iter: 1
tokens: {"input_tokens": 44, "output_tokens": 255, "estimated": true}
tool_calls_n: 0
duration_ms: 4339
aborted: false
abort_reason: ""
spawned_children: []
skill: "gsd-graphify"
tags: ["agent-run", "apier", "planner-routed"]
---

## TL;DR
apier: Tabla o lista con metodos, paths y payloads. -> OK.

## Task recibida
```
Disena el contrato REST minimo para una API de gestion de agentes: listar, ver detalle, invocar, y consultar runs. Especifica metodos HTTP, paths y payloads de request/response.
```

## Output
## Contrato REST Mínimo – API de Gestión de Agentes

| Método | Path | Request Body | Response Body |
|--------|------|-------------|---------------|
| **GET** | `/agents` | — | `{ "agents": [ { "id": "string", "name": "string", "status": "string" } ] }` |
| **GET** | `/agents/{id}` | — | `{ "id": "string", "name": "string", "status": "string", "config": {}, "created_at": "string" }` |
| **POST** | `/agents/{id}/invoke` | `{ "input": "string", "context": {} }` | `{ "run_id": "string", "status": "string", "output": "string" }` |
| **GET** | `/agents/{id}/runs` | — | `{ "runs": [ { "run_id": "string", "status": "string", "started_at": "string", "input": "string", "output": "string" } ] }` |

**Notas:**
- `{id}` es el identificador único del agente (UUID o slug).
- Todos los responses incluyen `Content-Type: application/json`.
- Errores estándar: `400 Bad Request`, `404 Not Found`, `500 Internal Server Error`.
- Paginación opcional en `GET /agents` y `GET /agents/{id}/runs` vía query params `?page=1&limit=20`.

## Human notes
_(no se toca por automatizacion)_
