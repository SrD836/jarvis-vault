---
title: "Polymarket Bot v1 — Baseline aprobó 2 long-tail absurdas"
type: decision
date: 2026-05-19
decision: "Veto retrospectivo: v1 brain aprobó LeBron James 2028 President YES @ 0.0065 y Bernie Sanders 2028 nom YES @ 0.0075 — ambos long-tail con probabilidad realista <1%"
alternatives:
  - "Dejar v1 funcionando con filtros laxos (V1 vol, V2 catalyst 72h) — aprueba 477/500 candidates"
  - "v2: añadir news research (Tavily + DeepSeek) + memoria de errores antes de aprobar"
outcome: confirmed_bad
outcome_observed_after_days: 0
tags: [decision, polymarket, bot, baseline-failure, post-mortem]
related:
  - "[[agents/polymarket-bot]]"
  - "[[projects/polymarket-veto-loop-bot]]"
---

## Contexto

V1 del bot (deployed 2026-05-19 ~11:00) corrió scanner sobre Polymarket gamma-api público y obtuvo 500 candidates. Brain v1 con 6 reglas duras (V1-V6: volumen, catalyst 72h, triggers vagos, chasing 8%, patterns <3 casos, sin trigger claro) aprobó 477/500. Executor entró 2 (LeBron James, Bernie Sanders) y descartó 475 por cap correlación (todos `category=uncategorized`).

## Por qué v1 falló

1. **V1 (volumen)**: usaba `volumeNum` cumulativo en lugar de `volume24Hr` (que la API devuelve null). LeBron tenía $231k cumulativo pero 24h real era ~0. Filtro inútil contra long-tail.
2. **V2 (catalyst 72h pre-resolución)**: LeBron y Bernie resuelven en 2028 (~30 meses fuera). Catalyst no inminente → pasa.
3. **V3/V5/V6 (LLM auditor semántico)**: ejecutaron pero el system prompt no tenía contexto noticioso. DeepSeek aprobó por defecto porque "no hay evidencia de violación de reglas".
4. **Sin memoria**: ningún feedback de "mercados 2028 presidential con price_yes<0.05 son trampas".
5. **Sin news research**: una búsqueda en Tavily ("LeBron James 2028 president") devolvería 0 headlines sustantivos → debería disparar veto preventivo.

## Lecciones extraídas

- `volumeNum` cumulativo NO es proxy fiable de liquidez actual. Para v2 considerar bestBid/bestAsk + spread.
- Filtro `current_price_yes < 0.05` veta long-tails (probabilidad <5% en sim no merece la pena: take-profit +60% sobre 0.005 = 0.008, ganancia ridícula).
- Categoría `uncategorized` (por null en API) hace inútil el cap de correlación 2/category. v2 debe usar `events[0].title` como pseudo-categoría.
- Veto basado en reglas numéricas SOLO es insuficiente cuando los números son trampa (volume cumulativo, catalyst lejano). Necesita capa semántica con news.

## Posiciones cerradas como parte de esta auditoría

| Trade ID | Market | Side | Entry | Size | Cerrado por |
|---|---|---|---|---|---|
| T-561251-1779189157480 | Will LeBron James win the 2028 US Presidential Election? | YES | 0.0065 | $150 | manual_close_v1_audit |
| T-100001-1779188896893 | Fed cuts rate in June 2026? (mock data del primer run) | YES | 0.62 | $150 | manual_close_v1_audit |
| T-540817-1779188896893 | New Rihanna Album before GTA VI? (mock data del primer run) | YES | 0.35 | $150 | manual_close_v1_audit |
| T-559679-1779189157480 | Will Bernie Sanders win the 2028 Democratic presidential nomination? | YES | 0.0075 | $150 | manual_close_v1_audit |

Bankroll reset a 10,000 USD. v2 arranca con state limpio.

## Propuesta v2 (aplicada inmediatamente tras este post-mortem)

- Filtro adicional `current_price_yes ∈ [0.05, 0.95]` en Brain (descartar deep long-tail).
- Categoría = `events[0].title` parseado por scanner, no `category` directo de la API.
- Memory check: Brain consulta `agents/polymarket-bot/memory.md` antes de aprobar. Si pattern previo similar (slug-prefix, categoría, price-bucket) tiene outcome negativo → veto M1.
- News research (Tavily) sobre la pregunta del mercado. DeepSeek clasifica `{confirms, contradicts, silent, score}`. Si `contradicts` o `silent && catalyst_due_days<30` → veto N1/N2.
- Cada veto significativo (≠ V1 trivial) escribe `03-decisions/<date>-polymarket-veto-<slug>.md`. Cada cierre en pérdida escribe `loss-<slug>.md`.

## Human notes

_(no se toca por automatización)_
