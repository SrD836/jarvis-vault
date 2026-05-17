---
title: "Decisiones clave Plan v2 + v3"
type: seed
date: 2026-05-13
tags: [seed, decision, historia]
related:
  - "[[00-MOC]]"
  - "[[seeds/sistema-arquitectura]]"
---

# 🗝️ Decisiones clave Plan v2 + v3

Las decisiones más importantes que dieron forma al sistema actual. Cada una tiene su entrada detallada en [[03-decisions/index]] (cuando aplique).

## Servidor: Hetzner CPX32 vs Oracle Free ARM

Plan v2 original proponía Oracle Cloud Free Tier ARM. Descartado por reportes de baneos y dificultad de provisión. **CPX32 amd64 €16.93/mes** prevaleció. Consecuencia: build images para `linux/amd64` siempre.

## OpenClaw vs sistema custom desde cero

Plan v2 evaluó construir agente custom. Optamos por **fork de OpenClaw** porque ya provee: gateway + chat + plugins (skill-workshop, memory-core, memory-wiki) + control-ui + sandboxing + cron nativo. Ahorro estimado: semanas de desarrollo.

## Claude Max vs API key Anthropic

**Claude Max ($100/mes)** via claude-cli backend de OpenClaw. Cuota generosa, suscripción fija. Trade-off: no hay tracking de tokens nativo (Max no expone billing), por eso "budget guardian" del plan v2 no aplicó literal — sustituido por **failover chain** + **OAuth health check** (ver [[03-decisions/2026-05-13-budget-guardian]]).

## Fallback chain

`anthropic/claude-sonnet-4-6` (primary) → `ollama-local/qwen2.5-coder:7b-instruct` (fallback). Auto-rotate cuando Anthropic rate-limita 429. DeepSeek queda pendiente (sintaxis `apiKeyEnv` rechazada por schema OpenClaw, requiere auth-profiles).

## Hermes auto-improvement: bundled vs plugin custom

Plan v2 original proponía plugin TS custom para user-model (~200 LoC). Pivot: **usar bundled** (`memory-core`, `memory-wiki`, `skill-workshop`). `skill-workshop` en modo `auto` desde Plan v3 — el agente crea skills SIN aprobación manual cuando detecta procedimiento reusable.

Ver también: [[../04-skills-log/index]] para historial de skills.

## Sandbox: non-main mode

Sesiones isolated (cron jobs, futuros `/goal`) corren en contenedores docker hijo via socket mount + cap-drop. Sesión `main` (Telegram + UI) sigue en el container gateway por simplicidad. Mitigación de risk "agente con docker socket = root effectivo" con AGENTS.md approval gates + sub-sandboxes restringidos.

## Backups: local rotativo vs Backblaze B2

Plan v2 sugirió Backblaze B2 ($0.005/GB/mes). Pivot a **backup local rotativo 7 días** (`~/agent-stack/backups/`) porque vault + skills + configs <50MB no justifican cuenta pagada. El vault además está syncthing'd al portátil = backup off-VPS gratis.

## JARVIS Control UI: overlay vs fork rebuild

OpenClaw incluye UI web nativa (`ui/openclaw-control-ui`, Vite+Lit) bundled en `/app/dist/control-ui/`. **Overlay via volume mount** (no rebuild image) — copia rebrand a `~/agent-stack/jarvis-ui-overrides/`, reemplazos "OpenClaw"→"JARVIS" en bundles, theme verde acid (Plan v3). Trade-off: shadow DOM Lit no acepta CSS injection externa para todos los componentes.

## Vault Obsidian: segundo cerebro real

Plan v2 confió en memory-wiki bridge para llenar el vault automáticamente. Falló: 0 exported artifacts después de Plan v2-v3. **Plan v4 (este)**: hooks explícitos para session-export, decision-log, skill-changelog + bootstrap rico de MOC y plantillas + AGENTS.md instruye al agente a usar el vault.

## Ver también

- [[seeds/sistema-arquitectura]] — qué corre dónde.
- [[seeds/operacion-diaria]] — flujos automáticos.
- [[00-MOC]] — vuelta al índice.
