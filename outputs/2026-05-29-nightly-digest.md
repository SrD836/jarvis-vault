---
title: "Nightly digest — 2026-05-29"
type: nightly-digest
date: 2026-05-29
generated_at: 2026-05-29T09:16:18Z
model: deepseek-chat
related:
  - "[[01-briefs/2026-05-29]]"
  - "[[agents/index]]"
  - "[[agents/runs-today]]"
  - "[[outputs/2026-05-28-nightly-digest]]"
  - "[[outputs/memory/cron_nightly_digest]]"
tags: [nightly-digest, jarvis, autonomous]
# auto-linked 2026-05-29
---


# Digest nocturno — 2026-05-29

> Generado autónomamente por `hermes/nightly_digest.py` (DeepSeek). FASE 3.

## Stats

- Decisiones: **102** · outcomes `{'pending': 102}`
- Reglas top: `{'M1': 78, 'P9': 11, 'M2': 2}`
- Sesiones: **0**
- Runs: **18** · modelos `{'none': 10, 'deepseek-chat': 4, 'ollama:qwen2.5-coder:7b-instruct': 3, 'claude(blocked)': 1}` · results `{'ok': 14, 'error': 4}`

## Resumen del día
Día sin sesiones de trading. 102 decisiones todas en estado "pending" (sin resolución). 18 ejecuciones de agentes: 14 ok, 4 error. Dominio aplastante de la regla M1 (memoria, 78/102 decisiones) y P9 (geopolítica, 11). Los errores vienen de `test:fase2-failloud` (2) más un `claude(blocked)` y otro sin especificar.

## Patrones y decisiones
- **M1 satura el pipeline**: 76% de decisiones son veto por coincidencia exacta de slug en memoria. Temas recurrentes: Irán (18 tesis), deportes (fútbol, tenis, NBA), Bitcoin, política internacional.
- **P9 activo en geopolítica**: 11 vetos por "geopolitics pump cluster", todos sobre Irán/Israel/Hormuz con precios bajos (0.05-0.21) y plazos cortos (0-8 días).
- **M2 marginal**: solo 2 vetos (Suez, NATO-Rusia) con 0% win rate histórico en 7 intentos.
- **P4_pre_event**: 8 vetos sobre FDV de lanzamientos (Metamask, Opensea, Predict.fun, etc.) con plazos de +200 días.

## Insights (no obvios)
- **M1 está matando la diversidad**: 78 decisiones idénticas por "exact slug match" sugiere que el sistema está vetando en masa tesis que ya existen, sin filtrar por calidad o novedad. Esto puede estar ocultando señales reales.
- **P9 es la única regla no-memoria con tracción**: pero sus precios son bajísimos (<0.21) y plazos cortos. Si alguna de estas tesis se resuelve "Sí", el sistema perdió una apuesta de alto retorno.
- **4 errores en 18 runs (22%)**: alto para un sistema autónomo. `fase2-failloud` falla consistentemente; `claude(blocked)` sugiere problemas de API.

## Qué vigilar mañana
- **Resolución de tesis P9 del 29-30 mayo**: varias vencen hoy/mañana (US-Iran ceasefire, Hormuz). Si alguna se resuelve "Sí", revisar si el veto fue correcto.
- **Los 4 errores**: si `fase2-failloud` y `claude(blocked)` persisten, escalar a revisión de infraestructura.
- **Saturación M1**: si mañana vuelve a haber >70% de decisiones por memoria, considerar ajustar el threshold de score o añadir un filtro de diversidad temática.
