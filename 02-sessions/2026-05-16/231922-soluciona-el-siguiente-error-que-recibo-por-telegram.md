---
title: "soluciona-el-siguiente-error-que-recibo-por-telegram — 2026-05-16"
type: session
date: 2026-05-16T23:19:22+02:00
topic: soluciona-el-siguiente-error-que-recibo-por-telegram
agent: "[[agents/main]]"
model_used: "deepseek/deepseek-chat"
tools_used: ["grep_search", "run_bash_networked", "run_bash", "read_file"]
tools_calls_n: 14
tokens: {"input_tokens":142074,"output_tokens":1789}
duration_ms: 22215
iterations: 7
aborted: false
spawned_children:[]
related:
  - "[[02-sessions/2026-05-16/232120-dime-simplemente-como-solucionarlo-puedo-acceder-a-la]]"
  - "[[02-sessions/2026-05-16/233746-es-decir-como-eliminamos-el-error-recibido-en]]"
  - "[[03-decisions/2026-05-14-session-export-analysis]]"
  - "[[agents/debugger]]"
  - "[[skills/external/gsd-debug]]"
tags: ["sesion", "main"]
# auto-linked 2026-05-17
---

## TL;DR
Soluciona el siguiente error que recibo por telegram:⚠️ Cron job "session-export" failed: Error: Sandbox image not found: openclaw-sandbox:b… → Veo que docker no está disponible desde dentro del sandbox (no responde `docker ps`). Y el script `sandbox-setup.sh` no existe en el sistema…
## Prompt usuario
```
Soluciona el siguiente error que recibo por telegram:⚠️ Cron job "session-export" failed: Error: Sandbox image not found: openclaw-sandbox:bookworm-slim. Build it with scripts/sandbox-setup.sh before enabling Docker sandboxing. The default image includes python3 for sandbox write/edit helpers; OpenClaw will not substitute plain debian:bookworm-slim.
```
## Respuesta main
Veo que docker no está disponible desde dentro del sandbox (no responde `docker ps`). Y el script `sandbox-setup.sh` no existe en el sistema de archivos local. Esto es porque el script es parte del código fuente de OpenClaw — está en su repositorio.

El fix es simple: necesito que tú, David, ejecutes un comando en el host para construir la imagen Docker. Pero antes, déjame confirmar si el VPS host tiene docker y si la imagen no está presente.

**El error**: `session-export` (y cualquier cron job que use sandbox aislado) necesita la imagen Docker `openclaw-sandbox:bookworm-slim` que no está construida en el host. El script `scripts/sandbox-setup.sh` viene con el código fuente de OpenClaw — no está en el container del gateway.

**Solución**: necesito que executes esto desde el host del VPS:
…
## Tool calls
| # | tool | target/args | ok | ms |
|---|---|---|---|---|
| 1 | `grep_search` | {"pattern":"session-export","glob":"*.md","max_results":20} | ✓ | 39 |
| 2 | `run_bash_networked` | {"command":"docker images \| grep -i sandbox 2>/dev/null; echo \"---\"; docker im… | ✓ | 656 |
| 3 | `grep_search` | {"pattern":"session.export\|session-export\|sandbox-setup","max_results":20} | ✓ | 34 |
| 4 | `run_bash` | {"command":"find /openclaw-workspace -name \"sandbox-setup.sh\" 2>/dev/null; fin… | ✓ | 549 |
| 5 | `read_file` | openclaw/openclaw.json | ✓ | 5 |
| 6 | `grep_search` | /openclaw | ✓ | 69 |
| 7 | `grep_search` | /openclaw | ✓ | 605 |
| 8 | `read_file` | openclaw/cron/jobs.json | ✓ | 9 |
| 9 | `run_bash` | {"command":"find /openclaw -name \"sandbox-setup.sh\" -o -name \"*.sh\" 2>/dev/n… | ✓ | 574 |
| 10 | `run_bash` | {"command":"cat /openclaw/cron/runs/c10549cd-7362-4a9c-b3dc-c4f599aea8df.jsonl 2… | ✓ | 555 |
| 11 | `run_bash` | {"command":"find /home/node/.openclaw -name \"sandbox-setup*\" 2>/dev/null; find… | ✓ | 644 |
| 12 | `run_bash` | {"command":"ls -la /openclaw/scripts/ 2>/dev/null \|\| echo \"no /openclaw/scripts… | ✓ | 542 |
| 13 | `run_bash_networked` | {"command":"# Buscar Dockerfile y scripts dentro del container openclaw-gateway\… | ✓ | 705 |
| 14 | `grep_search` | /openclaw | ✓ | 480 |
## Human notes
_(añade observaciones manuales aquí; este bloque no se sobreescribe por automatización)_