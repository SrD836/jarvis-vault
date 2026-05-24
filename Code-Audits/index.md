---
title: "Code-Audits — Índice"
type: moc
created: 2026-05-24
tags: [moc, code-audit, qa, overnight]
related:
  - "[[00-MOC]]"
  - "[[agents/auditor]]"
  - "[[agents/code-reviewer]]"
---

# 🔍 Code-Audits — Auditorías overnight

Salida del cron `overnight-code` (01:00 daily). Cada nota = auditoría completa de un repo. Linkea a `[[agents/auditor]]` y `[[agents/code-reviewer]]`.

## Auditorías

| Fecha | Archivo |
|---|---|
| 2026-05-22 | [[Code-Audits/2026-05-22]] |
| 2026-05-21 | [[Code-Audits/2026-05-21]] |
| 2026-05-20 | [[Code-Audits/2026-05-20]] |
| 2026-05-19 | [[Code-Audits/2026-05-19]] |
| 2026-05-18 | [[Code-Audits/2026-05-18]] |
| 2026-05-17 | [[Code-Audits/2026-05-17]] |

## Cobertura

- Repos auditados: `praktor`, `jarvis-dashboard`, `openclaw-fork`, `agent-stack/hermes`.
- Triggers: cron 01:00 + bajo demanda (`/gsd-code-review`).
- Output: severidad-classified findings (critical/high/medium/low) + remediación sugerida.

## Patrones recurrentes

_(rellenar cuando weekly synthesis cron P2.1 esté activo — detectará issues que reaparecen semana tras semana)_
