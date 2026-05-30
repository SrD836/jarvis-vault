---
title: "apier run — 2026-05-29 231641"
type: agent-run
date: 2026-05-29T23:16:41+00:00
agent: "[[agents/apier]]"
spawned_from: "[[agents/planner]]"
parent_chain: ["main", "planner"]
model: deepseek-chat
provider: deepseek
iter: 1
tokens: {"input_tokens": 50, "output_tokens": 359, "estimated": true}
tool_calls_n: 0
duration_ms: 3482
aborted: false
abort_reason: ""
spawned_children: []
skill: "gsd-graphify"
related:
  - "[[agents/CATALOG]]"
  - "[[agents/apier]]"
  - "[[agents/planner]]"
  - "[[seeds/2026-05-29-model-routing-matrix]]"
  - "[[skills/external/gsd-graphify]]"
tags: ["agent-run", "apier", "planner-routed"]
# auto-linked 2026-05-30
---


## TL;DR
apier: Descripción del endpoint con request, response y códigos de error. -> OK.

## Task recibida
```
Diseña el contrato REST de GET /agents/{id}/status. Especifica parámetros de request, estructura JSON de respuesta (con campos como agent_id, status, last_heartbeat), y códigos de error (200, 404, 500).
```

## Output
GSD > GRAPHIFY

## GET /agents/{id}/status

### Request

| Parámetro | Tipo | Ubicación | Requerido | Descripción |
|-----------|------|-----------|-----------|-------------|
| `id` | string | path | sí | UUID del agente |

No requiere query params ni body.

### Response (200 OK)

```json
{
  "agent_id": "uuid-string",
  "status": "online",
  "last_heartbeat": "2025-04-10T14:30:00Z",
  "version": "1.2.3",
  "uptime_seconds": 3600
}
```

Campos:
- `agent_id`: string (UUID) — identificador único del agente
- `status`: enum `"online" | "offline" | "error" | "unknown"`
- `last_heartbeat`: string (ISO 8601) — timestamp del último heartbeat recibido
- `version`: string (semver) — versión del agente
- `uptime_seconds`: integer — segundos desde último reinicio

### Códigos de error

| Código | Condición | Body |
|--------|-----------|------|
| **200** | Agente encontrado y status recuperado | Objeto status |
| **404** | `id` no corresponde a ningún agente registrado | `{ "error": "agent_not_found", "message": "Agent {id} not found" }` |
| **500** | Error interno (DB caída, timeout, etc.) | `{ "error": "internal_error", "message": "Failed to retrieve agent status" }` |

### Convenciones RESTful seguidas
- Path param para identificar recurso (`/agents/{id}/status`)
- GET es idempotente y seguro (solo lectura)
- Códigos HTTP estándar (200, 404, 500)
- Respuesta JSON con camelCase
- Error con `error` code + `message` legible

## Human notes
_(no se toca por automatizacion)_
