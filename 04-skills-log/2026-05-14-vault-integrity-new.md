---
title: "Nueva Skill — meta/vault-integrity"
type: skill-log
date: 2026-05-14
skill: meta/vault-integrity
action: new
version_before: ~
version_after: 1.0.0
tags: [skill, new, meta]
related:
  - "[[index]]"
  - "[[../skills/index]]"
  - "[[../00-MOC]]"
---

# 🛠️ Skill nueva: `meta/vault-integrity`

**Fecha:** 2026-05-14
**Acción:** `new` — registrada por el curator.
**Categoría:** `meta`
**Versión:** `1.0.0`

## Descripción

Comprueba la integridad del vault Obsidian: wikilinks rotos, frontmatter inválido (falta `title`, `type`, `tags`), archivos huérfanos sin `related`, y notas sin ningún enlace entrante ni saliente. Genera un reporte de salud del grafo.

## Razón del registro

Planificada en F3.4 del roadmap (ver [[../projects/guia-fases-pendientes]]). Implementación local en `agent-stack/hermes-upgrade/skills/vault-integrity.py` (ver puntero [[../projects/hermes-upgrade]]).

---

← [[index|Volver al skill log]] | [[../projects/guia-fases-pendientes|Roadmap de fases]]
