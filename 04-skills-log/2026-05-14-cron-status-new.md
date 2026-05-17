---
title: "Nueva Skill — meta/cron-status"
type: skill-log
date: 2026-05-14
skill: meta/cron-status
action: new
version_before: ~
version_after: 1.0.0
tags: [skill, new, meta]
related:
  - "[[index]]"
  - "[[../skills/index]]"
  - "[[../seeds/operacion-diaria]]"
  - "[[../00-MOC]]"
---

# 🛠️ Skill nueva: `meta/cron-status`

**Fecha:** 2026-05-14
**Acción:** `new` — registrada por el curator.
**Categoría:** `meta`
**Versión:** `1.0.0`

## Descripción

Verifica que todos los cron jobs del sistema se ejecutaron sin errores en las últimas 24 horas. Comprueba: `oauth-health`, `session-export`, `daily-brief`, `decision-log`, `curator-weekly`. Genera alerta Telegram si algún job falló o no registró ejecución.

## Razón del registro

Planificada en F3.4 del roadmap (ver [[../projects/guia-fases-pendientes]]). Complementa `alerts.sh` con visibilidad a nivel de log de OpenClaw, no solo de containers Docker.

## Ver también

- [[../seeds/operacion-diaria]] — timeline de cron jobs
- [[../seeds/sistema-arquitectura]] — lista de containers

---

← [[index|Volver al skill log]] | [[../seeds/operacion-diaria|Operación diaria]]
