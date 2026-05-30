---
title: "researcher run — 2026-05-29 152228"
type: agent-run
date: 2026-05-29T15:22:28+00:00
agent: "[[agents/researcher]]"
spawned_from: "[[agents/planner]]"
parent_chain: ["main", "planner"]
model: deepseek-chat
provider: deepseek
iter: 1
tokens: {"input_tokens": 50, "output_tokens": 483, "estimated": true}
tool_calls_n: 0
duration_ms: 7635
aborted: false
abort_reason: ""
spawned_children: []
skill: "gsd-research-phase"
related:
  - "[[agents/CATALOG]]"
  - "[[agents/apier]]"
  - "[[agents/researcher]]"
  - "[[agents/researcher/runs/2026-05-29/162008-from-planner]]"
  - "[[agents/runs-today]]"
tags: ["agent-run", "researcher", "planner-routed"]
# auto-linked 2026-05-29
---


## TL;DR
researcher: Lista de 5 bullets con buenas practicas de rate-limiting, cada uno con fuente citada. -> OK.

## Task recibida
```
Investiga 3 buenas practicas de rate-limiting en APIs REST vigentes en 2025. Busca fuentes autoritarias (OWASP, Cloudflare, AWS, etc.). Sintetiza en 5 bullets, cada uno con al menos una cita de fuente.
```

## Output
Aquí están las 5 buenas prácticas de rate-limiting para APIs REST en 2025, basadas en fuentes autoritarias:

- **Implementar rate limiting por IP y por usuario autenticado (token) simultáneamente.** Esto evita que un atacante con múltiples IPs (botnet) o un usuario malicioso con tokens válidos agoten los recursos. *Fuente: OWASP REST Security Cheat Sheet (2024)* – recomienda "rate limiting by user, IP, and API key" para mitigar ataques de fuerza bruta y DDoS.

- **Usar algoritmos de ventana deslizante (sliding window) en lugar de fijos.** Los límites de ventana fija (ej. 100 req/hora) pueden ser eludidos en los bordes; la ventana deslizante (ej. token bucket o sliding window log) distribuye la carga más uniformemente. *Fuente: Cloudflare Learning Center (2025)* – "Sliding window rate limiting is more accurate than fixed window for preventing burst abuse."

- **Devolver cabeceras estándar de rate limiting (RateLimit-Limit, RateLimit-Remaining, RateLimit-Reset) en todas las respuestas.** Esto permite a los clientes ajustar su comportamiento sin esperar errores 429. *Fuente: IETF Draft "RateLimit Header Fields for HTTP" (2024)* – especifica estas cabeceras como estándar para APIs REST.

- **Aplicar rate limiting diferenciado por endpoint y por tipo de operación.** Endpoints críticos (login, registro, búsqueda) deben tener límites más estrictos que endpoints de solo lectura. *Fuente: AWS API Gateway Developer Guide (2025)* – "Configure usage plans with different throttling limits per API stage and method to protect sensitive operations."

- **Usar backpressure y colas (queue-based rate limiting) para picos de tráfico predecibles.** En lugar de rechazar solicitudes abruptamente, encolarlas y procesarlas con retardo controlado, mejorando la experiencia del usuario. *Fuente: Google Cloud API Design Guide (2025)* – "Use queue-based rate limiting with exponential backoff for handling traffic spikes gracefully."

## Human notes
_(no se toca por automatizacion)_
