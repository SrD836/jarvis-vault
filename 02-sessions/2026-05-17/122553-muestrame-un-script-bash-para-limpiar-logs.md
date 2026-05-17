---
title: "muestrame-un-script-bash-para-limpiar-logs — 2026-05-17"
type: session
date: 2026-05-17T12:25:53+02:00
topic: muestrame-un-script-bash-para-limpiar-logs
agent: "[[agents/main]]"
model_used: "deepseek/deepseek-chat"
tools_used: []
tools_calls_n: 0
tokens: {"input_tokens":11661,"output_tokens":146}
duration_ms: 2759
iterations: 1
aborted: false
spawned_children:[]
related:
  - "[[02-sessions/2026-05-16/233227-ejecute-ssh-root8819816861-me-da-el-error-de]]"
  - "[[03-decisions/2026-05-14-session-export-analysis]]"
  - "[[agents/debugger]]"
  - "[[skills/external/gsd-cleanup]]"
  - "[[wiki/reports/stale-pages]]"
tags: ["sesion", "main"]
# auto-linked 2026-05-17
---

## TL;DR
muéstrame un script bash para limpiar logs → ¿Qué logs exactamente? Necesito contexto para que el script sea útil y no peligroso: - **¿Qué directorio?** `/var/log/`, `/home/node/.opencl…
## Prompt usuario
```
muéstrame un script bash para limpiar logs
```
## Respuesta main
¿Qué logs exactamente? Necesito contexto para que el script sea útil y no peligroso:

- **¿Qué directorio?** `/var/log/`, `/home/node/.openclaw/logs/`, docker container logs?
- **¿Qué antigüedad?** Rotar/días, o borrar todos?
- **¿Rotar o borrar?** Comprimir + archivar, o `rm -rf` directo?
- **¿Extensión/patrón?** `.log`, `*.jsonl`, o todo lo que tenga dentro?

Sin eso el script sería genérico e inaplicable. Dame dos líneas y lo escribo.
## Tool calls
_(sin tool calls)_
## Human notes
_(añade observaciones manuales aquí; este bloque no se sobreescribe por automatización)_