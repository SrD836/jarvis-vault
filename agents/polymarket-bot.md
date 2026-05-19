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

### Limitaciones conocidas
- Tavily quota: search_depth=advanced consume ~2× créditos. Free tier (1000/mes) se agota en horas con cron */30 sobre ~50 candidates/ciclo. Cache TTL 6h mitiga parcialmente. Considerar plan pagado o NewsAPI/RSS scraping como alternativas.
- gamma-api: `currentPrice` viene vacío frecuentemente. Fallback chain instalado: lastTradePrice → outcomePrices[0] → mid(bestBid, bestAsk).
- Categoría: scanner devuelve "uncategorized" para >95% mercados. max_same_category=10 + skip para uncategorized para no matar diversidad.
