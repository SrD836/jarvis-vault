---
title: "Nueva Skill — meta/curator"
type: skill-log
date: 2026-05-14
skill: meta/curator
action: new
version_before: 1.0.0
version_after: 2.0.0
tags: [skill, new, meta]
related:
  - "[[index]]"
  - "[[../skills/index]]"
  - "[[../seeds/decisiones-clave]]"
  - "[[../00-MOC]]"
---

# 🛠️ Skill actualizada: `meta/curator`

**Fecha:** 2026-05-14
**Acción:** `new` (re-registro — v1→v2)
**Categoría:** `meta`
**Versión:** `2.0.0`

## Descripción

El curator gestiona el ciclo de vida de las skills: marca como `stale` (>30d sin uso), `archive` (>90d), regenera `SKILL_INDEX.md`, y ejecuta auditoría semántica semanal (lunes 03:00). La v2.0.0 incluye deduplicación y refinamiento LLM.

## Diff relevante

- **v1.0.0:** Lifecycle determinístico básico (stale/archive/SKILL_INDEX).
- **v2.0.0:** + deduplicación por overlap semántico + `curator_patch.py` para casos edge.

## Notas

Ver [[../seeds/decisiones-clave]] para el contexto del pivot de skill-workshop auto.

---

← [[index|Volver al skill log]] | [[../seeds/decisiones-clave|Decisiones clave]]
