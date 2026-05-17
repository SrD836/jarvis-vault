---
title: "Antigravity Memory — JARVIS"
type: memory
date: 2026-05-14
tags: [memory, antigravity, context]
---

# 🧠 Antigravity Memory — JARVIS

> Memoria persistente entre sesiones. Actualizada automáticamente por Ralph Loop.
> Última actualización: 2026-05-14T13:50:00+01:00

---

## Contexto del proyecto

**JARVIS** es un asistente personal autónomo de David que opera 24/7 en un VPS Hetzner CPX32.

### Stack técnico
- **Orquestador:** OpenClaw (fork) con gateway, CLI, plugins bundled
- **Modelos:** Claude Sonnet 4.6 / Opus (Claude Max $100/mes) + Ollama Qwen2.5-Coder 7B (fallback)
- **Infraestructura:** Docker Compose (5 containers), Traefik reverse proxy, Syncthing
- **Vault:** Obsidian vault sincronizado bidireccional VPS↔portátil
- **Interfaces:** Telegram bot (@Jarvvssbot) + Control UI web (jarvss.duckdns.org)
- **Cron jobs:** 5 activos (daily-brief, session-export, curator-weekly, oauth-health, alerts)

### Arquitectura de agentes
7 agentes configurados: main, planner (orchestrator), code-reviewer, researcher, documenter, apier, skill-reviewer. Planner puede delegar a code-reviewer, researcher, documenter.

### Memoria del sistema
- **Volátil:** contexto de sesión (1M tokens Opus, 200K Sonnet)
- **Persistente cercana:** memory-core embeddings (nomic-embed-text via Ollama)
- **Persistente lejana:** este vault (memory-wiki bridge, actualmente inactivo)

---

## Problemas conocidos

1. **session-export retorna 0 sesiones** — requiere revisión de logs del cron en VPS.
2. **Skill-logs duplicados** — confirmada duplicación por curator, requiere fix en VPS.
3. **memory-wiki inactivo** — requiere cambiar modo a 'active' en VPS.
4. ~~**Archivos de identidad pendientes**~~ — IDENTITY.md, USER.md, TOOLS.md creados y referenciados.

---

## Decisiones tomadas

| Fecha | Decisión | Justificación |
|---|---|---|
| 2026-05-14 | Bootstrap Ralph Loop + Supersmooth | Necesidad de entorno autónomo iterativo |
| 2026-05-14 | Crear PRD con 4 fases | Priorizar estabilización antes de features |
| 2026-05-14 | Max 10 iteraciones por sesión | Balance entre progreso y control |
| 2026-05-14 | Anti-drift cada 3 iteraciones | Prevenir desviación del objetivo |

| 2026-05-14 | Ejecución de 10 iteraciones del Loop | Validación completa de F0, completada F1.2 (identidades), e integración formal de Claude Code al ecosistema. |

---

## Patrones aprendidos
- [Auto-Consolidated] Se han consolidado los catálogos de habilidades: changed from 6 skills activas to 2 skills activas (meta/crear-skills, meta/curator)
- [Auto-Consolidated] David ha estado inactivo durante 4 días consecutivos: changed from None to 4 días consecutivos de inactividad

- El vault fue bootstrapped el 2026-05-13 (Plan v4) con estructura rica pero sin contenido auto-generado
- Los cron jobs están funcionando (daily-brief genera contenido) pero session-export no captura sesiones
- El sistema de skills tiene solo 2 skills activas (meta categoría)
- La delegación multi-agente está configurada pero no probada
- **Claude Code** opera de manera fluida conectado directamente al Obsidian Vault; está listado explícitamente en `TOOLS.md` y documentado para colaborar de la mano con las herramientas de OpenClaw.

---

## Contexto para próxima sesión

**Fase actual:** F1 — Estabilización (Continuación)
**Próxima tarea:** Conectarse al VPS para aplicar fixes.
**Prioridad:** Ejecutar las resoluciones físicas de F1.1 (session-export), F1.3 (deduplicación curator), F1.4 (activación memory-wiki), y F1.5 (pruebas failover) en el entorno del servidor.
**Precaución:** No romper nada en el VPS al modificar la configuración.
