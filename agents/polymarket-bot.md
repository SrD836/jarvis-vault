---
title: "polymarket-bot"
type: agent
role: worker
agent_id: polymarket-bot
workspace: /home/agent/agent-stack/vault/projects/polymarket-veto-loop-bot
model_primary: deepseek/deepseek-chat
delegation_mode: standalone
allow_agents: []
runtime_children: []
updated: 2026-05-19T11:30:00+02:00
tags: [agent, jarvis, worker, polymarket, trading-bot]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
  - "[[projects/polymarket-veto-loop-bot]]"
  - "[[polymarket-bot/memory]]"
---

# 🤖 polymarket-bot

**Rol:** worker (no-LLM, ejecuta como binarios Go en cron)
**Modelo LLM auditor:** `deepseek/deepseek-chat` via `/api/llm/call` del dashboard
**Workspace:** `/home/agent/agent-stack/vault/projects/polymarket-veto-loop-bot/`

## Identidad

Bot de trading simulado sobre Polymarket. Estrategia veto-loop: LLM como auditor estricto de reglas, no como predictor. Arranca en modo simulación, bankroll virtual 10,000 USD. La fase real (dinero) está explícitamente fuera de alcance hasta 30 días sin dramas.

## Componentes

1. **scanner** — Polymarket gamma-api → `vault/inbox/trading/candidates.jsonl`
2. **brain** — Aplica 6 reglas duras (V1–V6) + memory check + news research (Tavily + DeepSeek) → `approved.jsonl` / `blocked.jsonl`
3. **executor** — Sizing Kelly 1.5%, exposure cap 10%, FIFO → `active.jsonl` + `portfolio.json`
4. **exit_monitor** — Stop USD-at-risk + take +60% + cierre de mercado → `closed.jsonl`

## Memoria persistente

Ver [[polymarket-bot/memory]]. Brain consulta esa memoria ANTES de aprobar una tesis. Si pattern previo (slug/category/price-bucket) tuvo outcome negativo → veto preventivo (M1).

## Decision log

Cada veto significativo (≠ V1 trivial) y cada cierre en pérdida escribe un markdown bajo `03-decisions/` con prefix `polymarket-`. Navegable desde Obsidian.

## Política de delegación

- **delegationMode:** `standalone` (binarios Go, no LLM en el loop principal — solo el auditor V3/V5/V6 vía /api/llm/call)
- **allowAgents:** ninguno (no delega a otros agentes del sistema)

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/polymarket-bot]]"
SORT date DESC LIMIT 10
```

## Human notes
- Modo: simulation
- Operación: autónoma 24/7 vía cron del host (usuario agent)
- Cuota horizon: 70% corto (≤2 d) / 15% medio (3-30 d) / 15% largo (>30 d) — aplicada por executor v3 con Q1 cap por bucket
- Config: `bot/config.json` (BOT_CONFIG_PATH env var)
- Fecha activación: 2026-05-19
- v3.1 (2026-05-19): gamma-api fix (lastTradePrice fallback), tab "Auditoría" en dashboard, Tavily advanced + recencia 7 días
- Cron crontab del usuario `agent` debe ser `0600 agent:crontab` — si Claude lo edita desde gateway-as-root rompe perms y cron silencioso

### Fuentes de research (Tavily search_depth=advanced, days=7, topic=news)
Dominios filtrados (include_domains):
- reuters.com, apnews.com, bloomberg.com, ft.com, wsj.com
- nytimes.com, washingtonpost.com, theguardian.com, bbc.com, aljazeera.com
- cnn.com, axios.com, politico.com
- x.com, twitter.com (búsqueda web indexada por Tavily, no API directa)

El LLM auditor (DeepSeek vía /api/llm/call) exige campo `cited_dates` con fechas de últimos 7 días o devuelve `silent: true`. Si Tavily no encuentra cobertura reciente en estos dominios → silent.

X API directa NO está integrada. Si hace falta tweets en tiempo real (no indexados por Tavily), integrar X API en sprint posterior.



### v4 (2026-05-19): dynamic sizing + memoria de fuentes + cap 50 posiciones

- **Sizing**: `size_mode=dynamic_equal` → tamaño = bankroll × 0.02, clamp [$25, $500]. Auto-escala con ganancias/pérdidas.
- **Cap máximo posiciones abiertas**: 50 (`max_open_positions`). Independiente de `max_new_trades_per_day=50`.
- **Memoria de fuentes**: cada trade aprobado guarda en `approved.jsonl` los dominios citados por el LLM en `sources_used`. Al cerrar, exit_monitor escribe per-source attribution a `vault/agents/polymarket-bot/source-stats.jsonl` y patchea el frontmatter de `vault/03-decisions/<date>-polymarket-trade-<slug>.md` con `outcome`/`pnl_usd`/`closed_at`.
- **Blacklist automática de dominios**: si una fuente acumula ≥5 trades atribuidos con win-rate <30% → brain.go la excluye del `include_domains` Tavily durante 30 días (rolling). Visible en tab "Memoria" del dashboard.
- **Polymarket URL en audit**: ahora usa `/market/<slug>` (con fallback `/event/<eventSlug>` si existe) en vez de `/event/<slug>` que devolvía 404 para mercados individuales.
- **Bankroll fix**: exit_monitor v4 devuelve el principal (`Size`) además del PnL al bankroll al cerrar (antes solo añadía el delta, lo que sangraba bankroll progresivamente).

### v5 (2026-05-19): pricing bid/ask correcto + early-exit etiquetado + Q3 liquidez

- **Pricing realista**: entry usa `bestAsk` (lo que pagaríamos al comprar YES), exit usa `bestBid` (lo que recibimos al vender YES). Antes ambos usaban mid/lastTradePrice y sobrestimaban PnL por ~spread/2 en cada lado.
- **Q3 liquidity gate** en executor: rechaza entry si `liquidity_usd < tradeSize × 4`. Threshold configurable vía `liquidity_min_ratio` (default 4 = nuestra posición ≤25% del orderbook visible).
- **Exit liquidity gate** en exit_monitor: si liquidity < Size×4 o `acceptingOrders=false`, deferir cierre al próximo ciclo (evita slippage al vender lote completo).
- **EarlyExit flag**: si el mercado sigue abierto al cerrar (take_profit/stop_loss antes de endDate) → `early_exit: true` + badge naranja "Anticipado · resuelve en N d" en UI. Diferencia clara entre cierre por trigger vs cierre por resolución del mercado.
- **stop_loss_pct** añadido (default 80%): segundo cap relativo al tamaño del trade, complementa stop_loss_usd (relevante con sizing dinámico que va desde $25).
- **Audit retro 2026-05-19**: trade Iran +$110 (exit 0.13) recalculado a bestBid=0.12 → PnL real $90, bankroll corregido -$20. Registro en `vault/agents/polymarket-bot/audit-log.jsonl`.
- **Dashboard**: tabla "Trades cerrados" añade columnas Tipo (Anticipado/Resuelto), Fuente precio (bestBid / fallback / audit), Liquidez al cierre (warning si <Size×4).

### v6 (2026-05-19): brain rules hardened code-side (post-mortem Discord)

- **Disparador**: trade Discord T-608399 perdio -$40.23 (-89.4%) en 10 min, entry 0.066, exit 0.007. Causa: cerca del floor, stop_loss_pct 80% se dispara con el spread bid/ask normal sin ningun catalyst.
- **Cambio**: las 3 reglas que vivian en memory.md (prompt-side, frias y dependientes de la fidelidad del LLM) se hardcodean en brain v6:
  - **P0_floor**: precio YES debe ser >= cfg.PriceFloor (default 0.10). Floor relajado a 0.05 si horizon <= 7d (binary events imminentes son OK a precios bajos).
  - **P0_ceiling**: precio YES <= cfg.PriceCeiling (default 0.95).
  - **P3_low_absolute_liquidity**: liquidity_usd >= cfg.MinAbsoluteLiquidityUSD (default $5000). Complementa el Q3 del executor (Size x ratio).
  - **P4_pre_event**: regex `(ipo-day|ipo|tge|mainnet|listing|release-date|launch|launches|launching)` con horizon >= cfg.PreEventVetoMinDays (default 7d) -> veto.
- **Politica grandfather**: las 12 posiciones activas con entry<0.10 o liquidity<$5k siguen vivas hasta su cierre natural (TP/SL/resolucion). Las reglas v6 son veto de **entrada**, no de **mantenimiento**.
- **Files**: bot/brain/internal/rules/rules.go (+3 funcs), bot/brain/internal/brain/brain.go (drop inline evalPriceBand, wire 3 calls + stats), bot/common/config/config.go (+4 fields with defaults).
- **Smoke test**: 3 candidatos sinteticos (Discord-like 0.066, thin book $2.9k, slug 'ipo-day' 29d) -> los 3 bloqueados con el VetoedBy correcto.
- **Binarios reconstruidos**: 2026-05-19 18:26 UTC (golang:1.23-alpine container). Cron `*/30` y `*/5` apuntan a bot/<m>/bin/<m> -> pick up automatico siguiente ciclo.

### Limitaciones conocidas
- Tavily quota: search_depth=advanced consume ~2× créditos. Free tier (1000/mes) se agota en horas con cron */30 sobre ~50 candidates/ciclo. Cache TTL 6h mitiga parcialmente. Considerar plan pagado o NewsAPI/RSS scraping como alternativas.
- gamma-api: `currentPrice` viene vacío frecuentemente. Fallback chain instalado: lastTradePrice → outcomePrices[0] → mid(bestBid, bestAsk).
- Categoría: scanner devuelve "uncategorized" para >95% mercados. max_same_category=10 + skip para uncategorized para no matar diversidad.
