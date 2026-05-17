---
title: "Análisis: Fallo de session-export"
type: decision
date: 2026-05-14
decision: Análisis local completado. Se requiere revisión de logs de cron en VPS.
alternatives:
  - "Revisar logs de Docker para el container openclaw-fork-openclaw-gateway-1"
  - "Provocar export manual para aislar el problema"
outcome: pending
outcome_observed_after_days: 1
tags: [decision, bug, sessions, cron]
related:
  - "[[../02-sessions/index]]"
---

## Contexto

El directorio `02-sessions/` se encuentra vacío de exportaciones automáticas. La tarea F1.1 indica que "session-export retorna 0 sesiones". 

## Análisis

El job `session-export` se ejecuta cada hora para exportar sesiones cerradas >30 min. Si no hay exportaciones, puede deberse a:
1. No ha habido sesiones con suficiente actividad que hayan superado el tiempo de inactividad.
2. El cron job falla silenciosamente o carece de permisos de escritura en el volumen montado del vault.
3. Desajuste en la configuración del plugin en OpenClaw.

## Siguiente paso (Acción VPS requerida)

1. Revisar los logs del cron: `docker exec openclaw-fork-openclaw-gateway-1 openclaw cron logs session-export`.
2. Realizar un export manual de una sesión conocida para verificar que el pipeline de escritura funciona.
3. Si el problema es de permisos, corregir UID/GID del volume mount en docker-compose.yml.

## Outcome observado

(Rellenar cuando se verifiquen los logs en el VPS.)
