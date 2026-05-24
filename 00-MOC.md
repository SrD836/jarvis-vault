---
title: "JARVIS — Mapa principal"
type: moc
created: 2026-05-13
updated: 2026-05-14
tags: [moc, root, jarvis]
related:
  - "[[PRD]]"
  - "[[TASKS]]"
  - "[[IDENTITY]]"
  - "[[USER]]"
  - "[[TOOLS]]"
  - "[[AGENT_RULES]]"
  - "[[Dashboard]]"
---

# 🧠 JARVIS — Mapa principal

> Segundo cerebro de David. Nodo raíz del vault. Todo lo demás se ramifica desde aquí.

## Áreas principales

| Categoría | Qué guarda | Frecuencia |
|---|---|---|
| [[01-briefs/index\|📰 Briefs diarios]] | Resumen 24h sesiones + decisiones + ideas pendientes | Diario 06:00 |
| [[02-sessions/index\|💬 Sesiones]] | Transcripts resumidos de conversaciones relevantes | Horario (auto-export) |
| [[03-decisions/index\|⚖️ Decisiones]] | Cadenas de razonamiento + outcome de cambios al sistema | Cuando aplique |
| [[04-skills-log/index\|🛠️ Skill log]] | Diffs y cambios al catálogo de skills | Cuando curator detecta |
| [[wiki/index\|📚 Wiki]] | Knowledge layer compilado por memory-wiki (provenance, dashboards) | Auto |
| [[agents/index\|🤖 Agentes]] | Espejo del equipo multi-agente: roles, allowAgents, runtime children | Cada 30 min (cron) |
| [[skills/index\|🛠️ Skills]] | Espejo Obsidian del catálogo `~/.openclaw/skills/` | Auto |
| [[memories/index\|🧠 Memorias]] | Memorias dialecticadas (concepts/entities/syntheses) | Auto memory-wiki |
| [[projects/index\|📁 Projects]] | Roll-ups manuales cross-session | Manual |
| [[goals/index\|🎯 Goals]] | Metas a largo plazo y proyectos activos | Manual / agente |
| [[templates/index\|🧩 Templates]] | Plantillas para briefs, sesiones, decisiones | Referencia |
| [[Dashboard\|🎛️ Dashboard]] | Métricas del sistema auto-generadas | Auto (dashboard_generator.py) |
| [[inbox/job-hunting/index\|💼 Job hunting]] | Búsquedas, aplicaciones, estado LinkedIn Easy Apply | Manual + cron |
| [[projects/portfolio-offers\|🗂️ Portfolios por oferta]] | HTML sites adaptados generados por `apply-batch.js` | Auto (pipeline trabajo) |
| [[Code-Audits/index\|🔍 Code audits]] | Auditorías overnight de repos (cron 01:00) | Diario |
| [[audit/2026-05-24-vault-audit\|📋 Auditoría vault]] | Salud del vault (cognición vs almacén) | Manual / on-demand |
| [[audit/2026-05-24-system-e2e-audit\|🩺 Auditoría sistema E2E]] | Foto fija stack JARVIS completo | Manual / on-demand |

## Arquitectura del sistema

- [[seeds/sistema-arquitectura\|🏗️ Arquitectura del stack]] — qué corre dónde, mounts, networks.
- [[seeds/decisiones-clave\|🗝️ Decisiones clave Plan v2/v3]] — por qué cada componente es como es.
- [[seeds/operacion-diaria\|🔁 Operación diaria]] — qué hace JARVIS cada día sin intervención.

- [[templates/brief|🗒️ Plantilla Brief diario]]
- [[templates/index|🧩 Todas las plantillas]]

## Documentos del sistema

| Documento | Propósito |
|---|---|
| [[PRD]] | Requisitos del producto y roadmap de fases |
| [[TASKS]] | Checklist operativo derivado del PRD |
| [[IDENTITY]] | Personalidad y directivas de JARVIS |
| [[USER]] | Perfil y preferencias de David |
| [[TOOLS]] | Catálogo de herramientas y stack técnico |
| [[AGENT_RULES]] | Reglas obligatorias de seguridad y ejecución |

## Cómo usar este vault

- **Leer**: la home (esta nota) tiene los puntos de entrada. Sigue wikilinks `[[así]]`.
- **Buscar**: `Ctrl+O` en Obsidian para abrir por nombre, `Ctrl+Shift+F` para full-text.
- **Grafo**: `Ctrl+G` para visualizar conexiones. Conforme el vault crece el grafo se densifica.
- **Tags**: filtros por `#sesion`, `#decision`, `#skill`, `#brief`, `#seed`, `#error`, `#fix`.

## Convenciones

- Frontmatter YAML obligatorio: `title, type, created, tags, related`.
- Wikilinks siempre que se mencione otra nota: `[[02-sessions/2026-05-13-foo]]`.
- Atomic notes: una idea por nota cuando sea posible.
- Plantillas en [[templates/index\|🧩 templates/]].

## Estado actual

- Bootstrap inicial 2026-05-13 (Plan v4).
- 7 fases Plan v2 + Plan v3 UI completadas, ver [[seeds/decisiones-clave]].
- Auto-fill: daily-brief + session-export + decision-log + skill-changelog activos vía cron.
- JARVIS Dashboard activo: [[Dashboard]] (auto-generado por `dashboard_generator.py`).
- Proyectos activos: [[projects/jarvis-dashboard|JARVIS Dashboard Web UI]], [[projects/guia-fases-pendientes|Roadmap de fases pendientes]].
- Upgrade Hermes en curso: [[projects/hermes-upgrade|hermes-upgrade scripts]] (consolidate, curator, dashboard_generator).

## Reports auto-generados

- [[wiki/reports/claim-health]]
- [[wiki/reports/contradictions]]
- [[wiki/reports/low-confidence]]
- [[wiki/reports/open-questions]]
- [[wiki/reports/person-agent-directory]]
- [[wiki/reports/privacy-review]]
- [[wiki/reports/provenance-coverage]]
- [[wiki/reports/relationship-graph]]
- [[wiki/reports/stale-pages]]
