---
title: "Análisis: Estado inactivo memory-wiki"
type: decision
date: 2026-05-14
decision: Análisis local completado. Configurar bridge en modo active en VPS.
alternatives:
  - "Cambiar modo del plugin memory-wiki a 'active'"
  - "Forzar resincronización si depende de session-export"
outcome: pending
outcome_observed_after_days: 1
tags: [decision, bug, wiki, memory]
related:
  - "[[../wiki/index]]"
---

## Contexto

El `memory-wiki bridge` está inactivo. La carpeta `wiki/` (concepts, entities, syntheses) tiene 0 artifacts generados.

## Análisis

De acuerdo al PRD y a las notas del sistema, el plugin memory-wiki en OpenClaw probablemente se encuentre en un modo pasivo o inactivo (quizás configurado solo como `bridge` o desactivado temporalmente), lo cual previene la auto-generación de conocimiento a partir de las sesiones. Además, como el export de sesiones está fallando (F1.1), el wiki carece de material primario (sesiones) para analizar.

## Siguiente paso (Acción VPS requerida)

1. Verificar que el error no sea arrastrado por `session-export` (solucionar F1.1 primero).
2. Verificar la configuración del plugin memory-wiki en `~/.openclaw/config.json`.
3. Activar el modo de ingesta (`active`) y enviar un concepto de prueba.

## Outcome observado

(Rellenar cuando se reconfigure el plugin en el VPS.)
