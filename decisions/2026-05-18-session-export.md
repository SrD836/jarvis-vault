---
title: "Diagnóstico: session-export insuficiente en vault/02-sessions"
type: decision
date: 2026-05-18
decision: Root cause identificado. Fix propuesto sin restart de gateway.
outcome: pending
related:
  - "[[../02-sessions/index]]"
  - "[[../03-decisions/2026-05-14-session-export-analysis]]"
  - "[[02-sessions/2026-05-16/234354-resumen-el-cron-session-export-sigue-activo-o-lo]]"
  - "[[02-sessions/index]]"
  - "[[03-decisions/2026-05-14-session-export-analysis]]"
  - "[[03-decisions/2026-05-17-telegram-bot-typescript-findings]]"
  - "[[agents/monitor]]"
tags: [decision, bug, sessions, cron, telegram, gateway]
# auto-linked 2026-05-18
---


## Síntoma reportado

Solo existe 1 archivo en `vault/02-sessions/` a nivel raíz (únicamente `index.md`) pese a múltiples sesiones de Telegram activas. Los subdirectorios de fecha tienen cobertura parcial e inconsistente.

## Estado real encontrado

```
vault/02-sessions/
├── .deleted/
│   └── sess-20260517-b9cb23.2026-05-18T09-00-00-721Z.112926-test.md  # movido hoy 09:00 UTC
├── 2026-05-16/   → 13 archivos (exportados en batch el 2026-05-17 08:30–08:31)
├── 2026-05-17/   → 4 archivos  (exportados individualmente a distintas horas)
├── 2026-05-18/   → NO EXISTE   (0 sesiones exportadas hoy)
└── index.md
```

Sesiones sin exportar en `.openclaw/agents/main/sessions/`:  
**9 archivos `.trajectory.jsonl`** (todos de 2026-05-17), tamaños entre 119 KB y 139 KB.

---

## Causa raíz

### 1. `crons: []` en openclaw.json — CAUSA PRIMARIA

```json
// /home/node/.openclaw/openclaw.json
"crons": []
```

El array de crons está vacío. No hay ningún job programado para exportar sesiones periódicamente. El análisis previo del 2026-05-14 (`03-decisions/2026-05-14-session-export-analysis.md`) ya identificó que el job `session-export` debía ejecutarse cada hora para sesiones cerradas >30 min. Ese job nunca se re-añadió tras el `openclaw doctor --fix` del 2026-05-14T19:21 UTC.

### 2. 16+ fallos de startup del gateway durante 2026-05-16/17 — CAUSA SECUNDARIA

```
/home/node/.openclaw/logs/stability/  →  16+ archivos "gateway.startup_failed"
Error: "Invalid config ... agents.list.0.subagents.delegationMode: Invalid option"
```

Cada crash del gateway termina sesiones activas de forma anormal. El hook de exportación al cierre de sesión nunca se dispara en sesiones terminadas por crash. Esto explica por qué las 9 trajectories de mayo-17 no fueron exportadas: el gateway falló durante o después de esas sesiones.

### 3. Export funciona solo en cierre limpio — CONSECUENCIA

Las 4 sesiones exportadas el 2026-05-17 (timestamps: 14:30, 20:30, 21:20×2) corresponden a sesiones que cerraron limpiamente. Las 9 trajectories no exportadas cerraron por crash o timeout.

---

## Evidencia

| Evidencia | Fuente | Relevancia |
|-----------|--------|------------|
| `"crons": []` | `/home/node/.openclaw/openclaw.json` | Sin job programado |
| 16 archivos `gateway.startup_failed` | `/home/node/.openclaw/logs/stability/` | Gateway inestable |
| Error: `delegationMode: Invalid option` | stability logs | Config inválida causaba restart loop |
| 9 `.trajectory.jsonl` sin exportar | `.openclaw/agents/main/sessions/` | Sesiones perdidas en vault |
| Batch-export 08:30–08:31 del 17-may | timestamps en `02-sessions/2026-05-16/` | El cron funcionó ese día, luego desapareció |
| Sesión test movida a `.deleted` hoy 09:00 | `.deleted/sess-20260517-*.md` | Hay mecanismo residual que aún corre |
| Análisis previo 2026-05-14 | `03-decisions/2026-05-14-session-export-analysis.md` | Ya se detectó este fallo, no se resolvió |

---

## Diagnóstico del mecanismo residual (09:00 UTC hoy)

Algo corrió hoy a las 09:00 UTC (`2026-05-18T09-00-00-721Z` en el nombre del archivo deletado). Solo encontró y procesó la sesión `test` de mayo-17, moviéndola a `.deleted` por tener 0 mensajes útiles. Este proceso no tiene acceso a las 9 trajectories sin exportar ni crea directorio `2026-05-18/`. Origen probable: systemd timer o docker cron externo residual, no registrado en `openclaw.json`.

---

## Fix propuesto

### Fix 1 — Re-añadir el cron de session-export a openclaw.json (prioritario)

Añadir al array `crons` en `openclaw.json`:

```json
{
  "id": "session-export",
  "label": "Session Export: Telegram → Vault",
  "schedule": "0 * * * *",
  "agent": "archivist",
  "prompt": "Exporta todas las sesiones de Telegram cerradas hace más de 30 minutos que no estén aún en vault/02-sessions/. Formato: HHMMSS-{primer-mensaje-slug}.md en el subdirectorio de fecha correspondiente.",
  "enabled": true
}
```

### Fix 2 — Export manual de sesiones pendientes

Una vez re-añadido el cron, invocar manualmente el archivist para recuperar las 9 sesiones huérfanas de `.openclaw/agents/main/sessions/`.

### Fix 3 — Verificar estabilidad del gateway

El error `delegationMode: Invalid option` fue corregido por `openclaw doctor --fix` (2026-05-14). Confirmar que no hay más startup failures desde entonces revisando `logs/stability/` con `ls -lt`.

---

## Lo que NO se debe hacer

- **No reiniciar el gateway** (instrucción explícita del usuario).
- No modificar las trajectory files directamente — son read-only del runtime.

---

## Historial relacionado

- **2026-05-14**: Primer análisis identificó session-export como fallido. Acción pendiente: revisar logs VPS + export manual. **No se completó.**
- **2026-05-17 08:30**: Batch export funcionó correctamente para sesiones de 2026-05-16 (evidencia de que el mecanismo es correcto cuando el cron existe).
- **2026-05-18 09:00**: Proceso residual corre pero solo encuentra sesiones triviales.
