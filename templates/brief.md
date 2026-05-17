---
title: "Template: Daily brief"
type: template
tags: [template, brief]
related:
  - "[[index]]"
  - "[[../00-MOC]]"
---

# Template para daily-brief

```markdown
---
title: "Daily Brief — {YYYY-MM-DD}"
type: brief
date: {YYYY-MM-DD}
source: daily-brief
tags: [brief]
related:
  - "[[01-briefs/{YYYY-MM-DD-1}]]"  (brief anterior, si existe)
---

## 1. Resumen 24h (ordenado por importancia)

- {sesión 1 con [[wikilink]] al export}
- {sesión 2}

## 2. Decisiones tomadas

- [[03-decisions/{date}-{decision}]] — {breve}

## 3. Ideas nuevas pendientes

- {idea 1, capturada en sesión X}

## 4. Acciones pendientes para hoy

- [ ] {acción}
- [ ] {acción}

## 5. Estado del sistema

- Skills activas: {N} (ver `~/.openclaw/skills/SKILL_INDEX.md`)
- Cron próximos: {N} (próximo curator-weekly: {fecha})
- Disco usado: {%}

## 6. 💡 Insights y Oportunidades (Auto-generado)

> Análisis proactivo basado en la actividad reciente.

- {insight_1}
- {insight_2}

---
_Generado por daily-brief cron a las {HH:MM}._
```
