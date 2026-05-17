---
title: "Arquitectura del stack JARVIS"
type: seed
date: 2026-05-13
tags: [seed, arquitectura, infra]
related:
  - "[[00-MOC]]"
  - "[[seeds/decisiones-clave]]"
  - "[[seeds/operacion-diaria]]"
---

# 🏗️ Arquitectura del stack JARVIS

## Hardware

- **VPS Hetzner CPX32** (Falkenstein): 4 vCPU AMD, 8GB RAM, 80GB SSD, sin swap. IP `88.198.168.61`, hostname `ubuntu-4gb-fsn1-1`.
- Usuario operativo: `agent` (root deshabilitado, ufw + fail2ban activos).
- **Portátil David** (Windows 11): cliente Telegram + navegador para JARVIS UI + vault Obsidian + Syncthing.

## Containers (Docker Compose)

| Container | Imagen | Rol |
|---|---|---|
| `openclaw-fork-openclaw-gateway-1` | `openclaw:local-claude` | Gateway principal: chat agente, control-ui, cron, plugins |
| `openclaw-fork-openclaw-cli-1` | idem | Bridge CLI (no usado activamente) |
| `openclaw-fork-ollama-1` | `ollama/ollama` | Modelos locales: qwen2.5-coder:7b, nomic-embed-text |
| `openclaw-fork-traefik-1` | `traefik:v3.1` | Reverse proxy HTTPS, Let's Encrypt, basic auth en /syncthing |
| `openclaw-fork-syncthing-1` | `syncthing/syncthing` | Sync bidireccional VPS↔portátil |

## Endpoints

- `https://jarvss.duckdns.org/` — [[seeds/decisiones-clave#jarvis-control-ui\|JARVIS Control UI]] (auth: gateway token).
- `https://jarvss.duckdns.org/syncthing/` — Syncthing GUI (auth: basic, david/PxMxhs9rt6/mxLOk).
- Telegram bot `@Jarvvssbot` — interfaz móvil (allowlist user `1478080070`).

## Volumes y mounts críticos

| Path host | Mount container | Quién escribe |
|---|---|---|
| `~/.openclaw/` | `/home/node/.openclaw/` (gateway) | OpenClaw: skills, cron, agents, sessions |
| `~/agent-stack/vault/` | mismo path (DooD parity) | memory-wiki, daily-brief, session-export, este vault |
| `~/.local/state/syncthing/` | `/var/syncthing/config/` (syncthing) | Syncthing config, device IDs |
| `~/agent-stack/jarvis-ui-overrides/` | `/app/dist/control-ui/:ro` | Rebrand JARVIS UI |
| `~/agent-stack/traefik/` | `/etc/traefik/:ro` | Config + acme.json |

## Routing modelos

1. Tareas triviales → Ollama Qwen2.5-Coder 7B local.
2. Tareas medias → claude-sonnet-4-6 vía claude-cli (cuota Max).
3. Tareas pesadas → claude-opus-4-7 vía claude-cli.
4. Fallback automático: ver [[seeds/decisiones-clave#failover-chain]].

## Crons activos

Host (no OpenClaw):
- `*/15 * * * *` alerts.sh — container/disk watchdog.
- `30 3 * * *` curator.py — lifecycle determinístico skills.
- `0 4 * * *` backup.sh — tar.gz rotativo 7 días.

OpenClaw (`openclaw cron list`):
- `curator-weekly` — lunes 03:00, audita catálogo skills semánticamente.
- `daily-brief` — diario 06:00, escribe en [[01-briefs/index]].
- `oauth-health` — cada hora, pattern HEARTBEAT_OK.

## Ver también

- [[seeds/decisiones-clave]] — por qué cada componente es como es.
- [[seeds/operacion-diaria]] — qué pasa cada día sin intervención.
