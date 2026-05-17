---
title: "Análisis: Duplicación de skill-logs"
type: decision
date: 2026-05-14
decision: Análisis local completado. Se requiere revisión de curator.py en VPS.
alternatives:
  - "Modificar script curator.py para almacenar estado de skills ya procesadas"
  - "Comprobar hashes de archivos existentes en 04-skills-log antes de crear nuevos"
outcome: pending
outcome_observed_after_days: 1
tags: [decision, bug, skills, curator]
related:
  - "[[../04-skills-log/2026-05-13-crear-skills-new]]"
  - "[[../04-skills-log/2026-05-14-crear-skills-new]]"
---

## Contexto

Se ha detectado que `curator.py` está generando entradas duplicadas en la carpeta `04-skills-log` en pasadas sucesivas (ej. `2026-05-13-crear-skills-new.md` y `2026-05-14-crear-skills-new.md`). La información interna del log es idéntica excepto la fecha de ejecución.

## Análisis

El script determinístico `curator.py` parece no estar guardando o leyendo correctamente el estado previo de las skills. Al correr diariamente, marca las mismas skills pre-existentes como "new" de nuevo y emite un log en el vault.

## Siguiente paso (Acción VPS requerida)

1. Leer el código de `curator.py` en `~/.openclaw/scripts/` o donde resida.
2. Añadir lógica de deduplicación: comprobar si la skill ya existe en el index o si ya tiene un log de tipo "new".
3. **Decisión propuesta:** Implementar comprobación de existencia en la carpeta `04-skills-log` antes de emitir un nuevo archivo "new" para una misma skill.

## Outcome observado

(Rellenar cuando se despliegue el fix en el VPS.)
