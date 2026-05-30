---
title: Auditoría Polymarket-bot — estado y roadmap a dinero real
date: 2026-05-30
tags: [polymarket, bot, audit, decision]
---

# Auditoría Polymarket-bot + roadmap a beneficios y dinero real

Auditoría read-only tras 5 fases (fallback fix, veto rules Fase 9, categorización, expectancy proof).
**Headline honesto:** el bot es un buen motor de CALIBRACIÓN sin fuente de edge y sin gestión de
pérdidas; la capa de dinero real NO existe. La rentabilidad está bloqueada por estrategia, no por tooling.

## Lo que FUNCIONA (no tocar)
- Arquitectura sólida: scanner→brain→executor (*/30) + exit_monitor (*/5) + passive_resolver (1h) +
  analytics (brier */6h, daily_calibration 1:30, weekly, expectancy diario 8:20). go.work 5 módulos.
- Calibración buena: **Brier 0.12** (<0.25) sobre resueltas → las probabilidades son sanas.
- Capa de veto rica: S1 + P0-P11 + V1-V6 + M1-M3 + N1-N2 + E1-E2 + R1/R3/R5 + suspensión por categoría.
- Memoria consultada pre-trade: M1 (match vs pérdidas), M2 (soft-rules), M3 (anti-patterns), blacklist fuentes.
- Infra de medición honesta (esta sesión): mark_to_market + expectancy_report + cron diario.

## Lo que SANGRA — gaps a rentabilidad (por impacto)
- **G1 · Sin gestión de pérdidas (EL killer).** `monitor.go:87` "nunca cierres por % stop": v7 solo sale
  por target_hit / no_remaining_edge / resolución. Longshots perdedores cabalgan hasta 0. Asimetría
  (ganadores pronto, perdedores se mantienen) → realizado siempre parece bueno = sesgo supervivencia.
  Causa del peor-caso −$784. `stop_loss_*`/`take_profit_pct` del config están MUERTOS.
- **G2 · Sin fuente de edge. [HECHO 2026-05-30, vault VPS `296a4a9a`]** Raíz real: edge y research
  estaban DESCONECTADOS — el LLM declaraba `estimated_prob` sin web (desviación no informada vs precio =
  -EV), y el research solo vetaba aparte. Fix: research pasa a GATE POST-VETO (claudemax websearch solo
  sobre casi-aprobados, sobre el lado real; `DecideEdgeResearch` testeada: against→N1, silent+catalyst→N2,
  silent lejano→shrink edge + re-E2, support→opera+sources). `edge_research_enabled` config-gated.
  Verificado live: 9 approved (8 c/sources), N1=15 N2=3 → 67% de casi-aprobados muertos por noticia.
  ~27 claudemax/ciclo (tope 30, caché 6h). NO prueba expectancy+ aún → se mide en forward (06-13).
- **G3 · Spread/fees no modelados.** entry=bestAsk, exit=bestBid; libros finos → coste round-trip alto
  que NO entra en `min_edge_points` (0.03).
- **G4 · Kelly sobre-dimensiona longshots.** Kelly sobre edges sobreconfiados de baja prob → $50-165 en
  longshots de 2¢. shrink de calibración = 1.0 (inactivo).
- **G5 · Short-only + sesgo longshot.** quota 100% short (≤2d) → binarios "by may-31" = los perdedores.
- **G6 · Datos resueltos finos (2.2%):** 89/4013 → Brier/expectancy poco robustos aún.
- **G7 · Caps de riesgo flojos:** max_open 200, max_per_day 150, max_same_cat 10.
- **G8 · Atribución de fuentes rota** (0 dominios reales) → no aprende qué fuentes sirven.

## Lo que falta para DINERO REAL (build, no flip)
- **R1 · No existe capa de ejecución real.** `executor.go:46` `case "live": panic("live trading not
  authorized")`. Falta TODO: cliente CLOB, wallet + firma eip712, órdenes/cancelaciones, fills parciales,
  slippage, reconciliación on-chain. Proyecto entero.
- **R2 · Sin controles live:** kill-switch, tope pérdida diaria, reconciliación posición vs cadena.
- **R3 · Expectancy forward no probada** (FAIL actual). Precondición dura antes de R1.

## Roadmap
- **A — Prueba (ahora → 2026-06-13):** ya montada. R1/R3 activos, NO tocar estrategia (contamina datos).
  Cron diario acumula serie. Re-eval agendada 13-jun (cron VPS → Telegram).
- **B — Arreglar expectancy (tras 06-13):** ~~G1 gestión de pérdidas~~ [HECHO] → ~~G2 reactivar edge~~
  [HECHO] → G3 spread/fees en el gate → G4 sizing (cap Kelly low-price + shrink) → G7 caps estrictos.
- **C — Re-probar:** ventana forward con fixes B; criterio: expectancy combinada >0 Y robusta peor-caso
  Y Brier ≤0.25 Y sostenida (serie diaria), N≥50.
- **D — Build dinero real (SOLO si C pasa):** capa CLOB + controles live + canary capital pequeño.

### Especialización por nicho v1 (2026-05-30, vault VPS `04093b68`) — track aparte
Construido: router de estimadores + modelo barrera cripto (live spot + vol realizada, P0-skip para edges
de modelo). **HALLAZGO HONESTO: cripto líquido de 0-1 día es EFICIENTE** — el modelo de barrera coincide
con el precio (Model fires pero ~0 aprobados; E2 bloquea edge<0.03). El "+15.97/t market" histórico era
n=14 (vetos P6/P11 + supervivencia), no edge repetible del modelo. DECISIÓN: NO seguir cripto (corto
eficiente; medio especulativo + scanner no lo surte). **El dinero real está en DEPORTES vs casas sharp**
(arbitraje vs líneas Pinnacle = información asimétrica real, no apostar contra mercado eficiente) =
spec #2, BLOQUEADO en una odds API key que solo David puede dar de alta. El router ya soporta enchufarlo.

## Veredicto de criterio
NO ir a dinero real con el estado actual. El bloqueo no es de herramientas — **falta edge real (research
off) y falta cortar pérdidas (no stop)**. Calibración buena ≠ rentabilidad. La prueba debe venir de la
ventana forward LIMPIA, no del +$500 sesgado por supervivencia.

Ver [[jarvis_polymarket_expectancy_proof]].
