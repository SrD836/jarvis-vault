---
title: "Skill log — Índice"
type: moc
tags: [moc, skills, hermes]
related: ["[[00-MOC]]"]
---

# 🛠️ Skill log

Cambios al catálogo de skills. Emitido por `curator.py` (cron host 03:30) cuando detecta:
- Skill nueva creada por `skill-workshop` modo `auto`.
- Skill marcada `stale` (>30 días sin uso).
- Skill archivada (>90 días sin uso).
- Cambios en frontmatter (use_count++, version bump).

## Formato

Archivo: `YYYY-MM-DD-{skill-name}-{action}.md`

Frontmatter: `date, skill, action, version_before, version_after, tags, related`.

Cuerpo:
- Diff del frontmatter.
- Razón del cambio (auto-detectada por curator).
- Si es `archive`: descripción de para qué servía.

## Skills activas

Ver [[../00-MOC#estado-actual]] o consultar `~/.openclaw/skills/SKILL_INDEX.md` (auto-generado). Ver también [[../skills/index]].

## Entradas recientes

### 2026-05-14

| Skill | Acción | Versión |
|---|---|---|
| [[2026-05-14-crear-skills-new\|meta/crear-skills]] | new | 1.0.0 |
| [[2026-05-14-ralph-loop-new\|meta/ralph-loop]] | new | 1.0.0 |
| [[2026-05-14-skill-refiner-new\|meta/skill-refiner]] | new | 1.0.0 |
| [[2026-05-14-curator-new\|meta/curator]] | new (v1→v2) | 2.0.0 |
| [[2026-05-14-vault-integrity-new\|meta/vault-integrity]] | new | 1.0.0 |
| [[2026-05-14-cron-status-new\|meta/cron-status]] | new | 1.0.0 |

## Volver: [[../00-MOC]]

