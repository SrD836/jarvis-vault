---
title: "Operación diaria JARVIS"
type: seed
date: 2026-05-13
tags: [seed, operacion, cron, automation]
related:
  - "[[00-MOC]]"
  - "[[seeds/sistema-arquitectura]]"
---

# 🔁 Operación diaria

Lo que JARVIS hace cada día sin que tú toques nada.

## Timeline 24h

| Hora (Madrid) | Job | Qué hace |
|---|---|---|
| 03:30 | `curator.py` host | Reconstruye `~/.openclaw/skills/SKILL_INDEX.md`, marca skills stale (>30d), archive (>90d). |
| 04:00 | `backup.sh` host | tar.gz rotativo 7 días con skills/workspace/agents/cron/vault/configs. |
| 06:00 | `daily-brief` OpenClaw | Escribe `[[01-briefs/{hoy}]]` con resumen 24h, decisiones, ideas. Envía resumen 3 líneas a Telegram. |
| Cada hora :00 | `oauth-health` OpenClaw | Verifica `claude auth status`. Si OK responde `HEARTBEAT_OK` (silencioso). Si falla manda alerta Telegram. |
| Cada hora | `session-export` OpenClaw (Plan v4) | Detecta sesiones cerradas >30 min sin actividad, exporta a `[[02-sessions/...]]`. |
| Cada hora | `decision-detect` host (Plan v4) | Escanea git log + cambios config y emite entries en `[[03-decisions/...]]` si aplica. |
| Cada 15 min | `alerts.sh` host | Watchdog: container UP, healthy, disco <85%. Alertas Telegram out-of-band. |
| Lunes 03:00 | `curator-weekly` OpenClaw | Audit semántico catálogo skills (LLM). Genera reporte en `~/.openclaw/skills/.curator-report-*.md`. |

## Flujos reactivos

- **Mensaje a `@Jarvvssbot`** → gateway → claude-cli → respuesta. Tools disponibles: read, write, edit, bash, docker (con cap-drop sandboxes), telegram_send, web fetch.
- **Sesión chat UI** → idem, vía WebSocket directo al gateway.
- **Skill emergente** (≥5 tool calls no obvios + corrección o procedimiento novel) → skill-workshop captura automáticamente (modo `auto`), escribe en `~/.openclaw/skills/<cat>/<name>/SKILL.md`. curator.py lo recoge en próxima pasada.
- **Rate limit 429 Anthropic** → failover chain rota a ollama-qwen automáticamente. Sin intervención.
- **OAuth Claude Code expira** (cada 8-12h normalmente) → oauth-health detecta → alerta Telegram → David ejecuta `docker exec -it openclaw-fork-openclaw-gateway-1 claude /login`.

## Memoria

- **Volátil**: contexto de sesión (1M tokens Opus, 200K Sonnet).
- **Persistente cercana**: `~/.openclaw/agents/main/` + memory-core embeddings (vector search via ollama nomic-embed-text).
- **Persistente lejana / segundo cerebro**: este vault (`~/agent-stack/vault/`). Bridge bidireccional con memory-wiki bundled (modo `bridge`).
- **Memoria humana**: la edición manual de notas en el vault desde Obsidian. AGENTS.md instruye preservar bloques de "Human notes" cuando regenere algo.

## Ver también

- [[seeds/sistema-arquitectura]] — qué corre dónde.
- [[seeds/decisiones-clave]] — por qué cada cron es como es.
- [[01-briefs/index]] — briefs históricos.
