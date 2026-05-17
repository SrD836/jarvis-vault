---
title: "hermes-upgrade — puntero"
type: project-pointer
created: 2026-05-14
moved: 2026-05-16
tags: [project, pointer, hermes, scripts]
related:
  - "[[index]]"
  - "[[../00-MOC]]"
  - "[[../TASKS]]"
location: "C:\\Users\\david\\agent-stack\\hermes-upgrade"
---

# 🐍 hermes-upgrade

> Código movido fuera del vault el 2026-05-16. Esta nota mantiene la referencia.

## Ubicación real del código

`C:\Users\david\agent-stack\hermes-upgrade\` — abrir con file://C:/Users/david/agent-stack/hermes-upgrade

## Contenido

Scripts Python desarrollados durante Antigravity (14 May 2026) para parchear el stack Hermes:

| Script | Propósito | Estado deploy |
|---|---|---|
| `consolidate.py` | Consolidación de memoria — sustituye el modo `active` no soportado por el schema | F1.5.1 `[x]` reportado |
| `curator_patch.py` | Parche puntual de deduplicación skill-logs | F1.5.2 `[x]` reportado |
| `curator_v3.py` | Curator reescrito con dedupe formal integrado | F1.5.3 `[~]` fallido |
| `daily_insights.py` | Insights diarios cross-session | pendiente |
| `dashboard_generator.py` | Genera `Dashboard.md` del vault | activo |
| `goal_tracker.py` | Tracker de metas a largo plazo | pendiente |
| `weekly_synthesis.py` | Síntesis semanal cross-brief | pendiente |

## Skills derivadas

Dentro de `hermes-upgrade/skills/`:
- `cron-status.py` — diagnóstico cron host
- `health-check.py` — health VPS
- `ralph-loop.py` — implementación del Ralph Loop
- `vault-integrity.py` — comprueba wikilinks rotos y frontmatter inválido (ver [[../04-skills-log/2026-05-14-vault-integrity-new]])
- `skill-refiner/` — refinador automático de skills

## Verificación pendiente

- SSH al VPS: `ls ~/.openclaw/scripts/` para confirmar qué scripts están realmente en producción.
- `crontab -l` para confirmar `consolidate.py 04:30` desplegado.

← [[index|Volver a projects]] | [[../TASKS|TASKS]]
