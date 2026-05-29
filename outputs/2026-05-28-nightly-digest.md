---
title: "Nightly digest — 2026-05-28"
type: nightly-digest
date: 2026-05-28
generated_at: 2026-05-29T09:16:09Z
model: deepseek-chat
tags: [nightly-digest, jarvis, autonomous]
---

# Digest nocturno — 2026-05-28

> Generado autónomamente por `hermes/nightly_digest.py` (DeepSeek). FASE 3.

## Stats

- Decisiones: **155** · outcomes `{'confirmed_bad': 12, 'pending': 143}`
- Reglas top: `{'M1': 110, 'P9': 13, 'M2': 5}`
- Sesiones: **0**
- Runs: **0** · modelos `{}` · results `{}`

## Resumen del día
155 decisiones, todas con resultado pendiente excepto 12 pérdidas confirmadas. Sin sesiones ni ejecuciones de agentes. Las pérdidas suman ~-145 USD en trades de cripto, geopolítica, deportes y política colombiana. Regla M1 domina (110 usos), seguida de P9 (13) y M2 (5).

## Patrones y decisiones
- **M1 (memoria)**: 71% de las decisiones. Bloquea réplicas exactas de tesis previas (score 1.00). Cobertura amplia: cripto, geopolítica, deportes, política.
- **P9 (geopolítica)**: 8%. Cluster de tesis Irán-EEUU con precios bajos (0.06-0.24) y horizontes cortos (0-9 días). Señal de evento inminente o ruido de baja probabilidad.
- **M2 (soft-learned)**: 3%. Patrones de derrota histórica: categorías cortas con >0.70 de precio (19L/0W) y rangos 0.10-0.30 (7L/0W). Baja muestra pero 100% pérdidas.

## Insights (no obvios)
- **M1 como filtro pasivo**: No evalúa mérito, solo duplicidad. Está bloqueando tesis que ya perdieron (ej: Bitcoin >$74k, Abelardo de la Espriella). Sin mecanismo de "no repetir errores", solo "no repetir preguntas".
- **P9 y la paradoja del corto plazo**: Tesis sobre acuerdo Irán-EEUU con resolución en días tienen precios 0.06-0.24. Si el mercado es eficiente, son apuestas casi perdidas. Pero el cluster sugiere que el sistema las trata como "oportunidad de pump" cuando quizás son ruido de baja probabilidad.
- **M2 confirma sesgo de favoritismo**: Las categorías donde el sistema apostó fuerte (>0.70) y perdió 19 de 19 veces indican sobreconfianza en eventos populares (deportes, política).

## Qué vigilar mañana
- **Resolución de tesis P9 del 28-30 mayo**: Si alguna se resuelve como "sí", validaría el cluster. Si todas pierden, P9 necesita recalibración.
- **Nuevas tesis sobre Irán/EEUU**: Si aparecen con precios >0.30, podría indicar fuga de información real.
- **M2 en categorías >0.70**: Ver si el sistema sigue vetando o si aprende a evitarlas.
