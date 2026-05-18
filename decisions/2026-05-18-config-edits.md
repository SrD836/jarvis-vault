---
title: Config edits openclaw.json — crons + ollama fallback
date: 2026-05-18
status: proposed
tags: [config, crons, ollama, fallback]
---

# Config edits 2026-05-18

Archivo propuesto: `/tmp/openclaw.json.proposed-2026-05-18` (container) → montar en host como `/home/agent/agent-stack/openclaw.json.proposed-2026-05-18`.

## Cambio 1 — `crons` (nuevo key top-level)

```diff
+ "crons": [
+   {
+     "id": "session-export",
+     "label": "Session Export Telegram to Vault",
+     "schedule": "0 * * * *",
+     "agent": "archivist",
+     "prompt": "Exporta a vault/02-sessions/YYYY-MM-DD/ todas las sesiones de ~/.openclaw/agents/main/sessions/*.jsonl cerradas hace mas de 30min que no existan ya en el vault. Formato: HHMMSS-{slug}.md.",
+     "enabled": true
+   }
+ ]
```

**Por qué:** automatizar archivado horario de sesiones Telegram→Vault sin intervención manual. Agente `archivist` ya existe en `agents.list`.

## Cambio 2 — `agents.defaults.model.fallbacks`

```diff
- "fallbacks": []
+ "fallbacks": ["ollama-local/qwen2.5-coder:7b-instruct"]
```

**Por qué:** cuando Claude Max alcanza cuota, el sistema necesita un modelo de fallback registrado para no bloquear tareas medias/triviales.

## Cambio 3 — `agents.defaults.models` (nueva entrada)

```diff
+ "ollama-local/qwen2.5-coder:7b-instruct": {
+   "agentRuntime": { "id": "ollama-local" }
+ }
```

**Por qué:** sin esta entrada el runtime no sabe qué agentRuntime usar para el modelo de fallback. El provider `ollama-local` ya está definido en `models.providers`.

## Validación

- JSON validado con `python3 -m json.tool` → OK (429 líneas).
- Original no modificado.

## Próximos pasos

1. David revisa `/tmp/openclaw.json.proposed-2026-05-18`.
2. Si aprueba: `cp /tmp/openclaw.json.proposed-2026-05-18 /home/node/.openclaw/openclaw.json` y restart gateway.
3. Verificar que cron `session-export` aparece en UI de gateway.
