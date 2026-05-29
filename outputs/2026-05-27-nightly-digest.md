---
title: "Nightly digest — 2026-05-27"
type: nightly-digest
date: 2026-05-27
generated_at: 2026-05-29T09:16:01Z
model: deepseek-chat
related:
  - "[[agents/archivist]]"
  - "[[agents/monitor]]"
  - "[[outputs/2026-05-28-nightly-digest]]"
  - "[[outputs/2026-05-29-nightly-digest]]"
  - "[[outputs/memory/cron_nightly_digest]]"
tags: [nightly-digest, jarvis, autonomous]
# auto-linked 2026-05-29
---


# Digest nocturno — 2026-05-27

> Generado autónomamente por `hermes/nightly_digest.py` (DeepSeek). FASE 3.

## Stats

- Decisiones: **143** · outcomes `{'confirmed_bad': 2, 'loss': 2, 'breakeven': 1, 'pending': 138}`
- Reglas top: `{'M1': 100, 'P9': 15, 'M2': 5}`
- Sesiones: **0**
- Runs: **0** · modelos `{}` · results `{}`

## Resumen del día
Día tranquilo en decisiones: 143 totales, 138 pendientes. Solo 2 trades confirmados como malos (Bitcoin y WTI Crudo, pérdidas de -70.91 y -45.72 USD). 2 pérdidas adicionales sin detalle, 1 breakeven. Sin sesiones ni runs ejecutados.

## Patrones y decisiones
- **Regla M1 domina**: 100 de 143 decisiones usaron memoria (exact slug match, score 1.00). Sistema vetando masivamente tesis duplicadas.
- **Regla P9 activa en geopolítica**: 15 decisiones con esta regla, todas sobre Irán/Israel/Estrecho de Hormuz. Precios bajos (0.05-0.29), plazos cortos (0-10 días).
- **Regla M2 (soft-learned)**: solo 3 casos, todos con win rate 0% (7L/0W). Tesis políticas y de cripto low-cap.

## Insights (no obvios)
- El sistema está hiperconservador con M1: veta cualquier tesis que ya exista en memoria, incluso variantes de fecha. Esto explica el 96.5% de decisiones pendientes.
- El clúster P9 de geopolítica sugiere que el sistema detectó un "pump" temprano en tesis de Irán, pero los precios siguen bajos (<0.30). Podría ser señal de ruido o de oportunidad si el mercado no ha priced in.
- Las 2 pérdidas confirmadas son trades direccionales (BTC y crudo) que fallaron. Sin datos de cuándo se abrieron ni el contexto de mercado.

## Qué vigilar mañana
- Resolución de las tesis P9 con fecha 0d (hoy): "Iran closes airspace by May 27" y "US announces Iran ceasefire by May 27". Si se resuelven como NO, el clúster P9 podría perder fuerza.
- Si aparecen nuevas tesis sobre BTC o crudo, ver si el sistema las veta con M1 o las deja pasar.
- Monitorear si las 2 pérdidas de hoy generan algún ajuste en reglas de trading.
