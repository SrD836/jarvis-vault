# Polymarket Veto-Loop Bot — Design v7

> v7 (2026-05-28): edge-gate + Kelly fraccional + thesis-based exits + calibración Brier.
> El bot arranca en `mode=shadow` (no abre posiciones, solo registra predicciones)
> hasta que `analytics/brier.py` confirme Brier < 0.20 en ≥1 categoría con ≥20 muestras
> resueltas. Toggle `mode: shadow → simulation` manual del user (1 línea en config.json).

## Cambios duros v7 (sobre el viejo Design v3.1 abajo)

### Reglas duras añadidas
| # | Regla | Significado |
|---|---|---|
| E1 | **Edge no declarado** | LLM devuelve `edge_type=none` o vacío → veto. Sin edge no se opera (prompt maestro principio #2). |
| E2 | **Edge bajo mínimo** | `|estimated_prob − implied_price| < min_edge_points` (default 0.05). Tras fees + ruido no hay margen. |
| S1 | **Categoría suspendida** | `analytics/brier.py` detecta Brier > 0.25 durante 3 semanas consecutivas en la categoría → escribe `suspended_categories.json`. Brain bloquea nuevos candidatos. |

### Sizing — Kelly fraccional ¼
- Reemplaza Kelly 1.5% fijo.
- `b = (1 - market_price) / market_price`; `f* = (p̂·b − (1−p̂))/b`; `size = bankroll · 0.25 · f*`.
- Caps duros: 5% por trade (`max_per_trade_pct`), 40% exposición total (`max_total_exposure_pct`).
- Si `f* ≤ 0` → reject `kelly_size_zero`, sin consumir cuota.

### Cuotas por horizonte — USD, no número de trades
- `horizon_quota_short × max_exposure_usd` = USD cap del bucket short, etc.
- Default v7: 75/15/10 (sobre 0.7/0.15/0.15 del v6).
- Justificación: un trade $25 ya no cuenta como un trade $500.

### Salidas — solo thesis-based (no SL/TP)
1. `market_closed` — Polymarket resolvió. Backfillea `outcome` en `predictions.jsonl`.
2. `no_remaining_edge` — `daysLeft < 1 ∧ |price − estimated_prob| < 0.02`.
3. `target_hit` — yes: `price ≥ target_prob`; no: `price ≤ target_prob`.
- **Eliminados** `take_profit`/`stop_loss` (% y USD). Prompt maestro principio #4.
- `bot/exit_monitor/cmd/force_close_horizon_excess/` marcado **DEPRECATED**. NO debe estar en cron.

### Modo shadow (calibración warm-up)
- `mode=shadow` → executor itera approved, calcula Kelly, log "would have opened …" y escribe `predictions.jsonl` con `decision="skip_shadow"`. No muta `portfolio.json`/`active.jsonl`.
- Release: cuando `brier.py` ve ≥1 categoría con Brier < 0.20 y ≥20 muestras resueltas → escribe `shadow_ready.json` + log `READY_TO_RESUME`.
- Toggle a `simulation` lo hace el user; el script nunca edita `config.json`.

### Calibración — `predictions.jsonl`
- Brain escribe una fila por cada decisión post-LLM (skip por E1/E2/S1/LLM, o approve).
- Executor en shadow escribe una fila por candidato aprobado.
- `exit_monitor` backfillea `outcome`/`resolution`/`resolved_at` cuando cierra por `market_closed`.
- `analytics/brier.py` agrega por `category × horizon × edge_type` → tabla markdown a memory.md.

### TradingView scope (P11) acotado
- Solo se evalúa si `dependsOnUnderlying(slug, question)` matches (price/above/below/$N/k|m|b/ATH...).
- Mercados como "will-coinbase-list-token-x" o "ethereum-merge-completes" ya no consumen llamadas MCP.

### Hard rules vigentes
- Regla 1 CLAUDE.md (jarvis_never_anthropic_api): LLM siempre vía `dashboard /api/llm/call` (Claude Max OAuth).
- Regla 6 (commit & push protocol): `git add -A && commit && push` al cerrar. Nunca `--no-verify`/`--force`/`--amend`.

---

# Polymarket Veto-Loop Bot — Design v3.1 (histórico)

> Contrato de referencia del sistema. Alcance cerrado. Sin prosa.

## Stack
- **Backend**: Go (bot CLI+loop), JSONL logging rotado diario
- **Frontend**: página Bot en sidebar del dashboard React (JARVIS)
- **DB**: solo JSONL (sin SQLite ni Redis)
- **Deploy**: mismo VPS Hetzner, container Docker en red JARVIS

## Universo de mercados
- **Solo Polymarket**, solo mercados **binarios** (Sí/No)
- **Volumen 24h ≥ 50,000 USD**
- Categorías explícitas: **Politics, Crypto, Sports, Tech**
- **Excluye Pop Culture** (memecoins, celebrities, eventos virales no replicables)
- Cobertura: solo las 4 categorías arriba, con filtros adicionales de liquidez

## Cola FIFO (única)
- **Input**: mercados vivos que pasan scanner
- **Orden**: timestamp de aprobación (no de creación del mercado)
- **Descarte terminal** si falla exposure/correlación/trigger — **no reencola**
- Máximo 3 mercados activos simultáneos (por size + colchón)

## Sizing & stops (Kelly + caps duros)
- **Kelly fraccional 1.5%** = 150 USD/trade sobre bankroll 10,000 USD
- **Exposure cap 10%** = 5,000 USD total en riesgo simultáneo
- **Stop-loss**: USD-at-risk = 150 USD (salir si esa posición pierde 150)
- **Take-profit**: +60% sobre precio de entrada (no +60% USDC)
- **No trailing stop**

## Las 8 críticas cerradas (errores que no cometer)

1. **Front-running mal implementado**: no se ejecuta contra taker — se pide límite a +2% del mid (slippage + protección)
2. **Cola sin fuerza**: orden FIFO por timestamp de aprobación, sin reencolado, sin prioridades
3. **Sin correlación**: máximo 2 mercados de la misma categoría, y solo si correlación <0.3 por precio histórico 7d
4. **Trigger único**: solo 1 trigger = movimiento de mid-price ±2% en zona de ejecución. No hay triggers combinados ni scoring compuesto
5. **Kelly agresivo**: Kelly 1.5% fijo. Ni 2.5% ni variable. Esto se controla en dashboard, no en código duro
6. **Ventana catalyst 72h**: si catalyst expira (fecha de resolución conocida), el mercado se descarta aunque cumpliera antes
7. **Sin chasing de momentum ni reversión a media**: solo trigger ±2%, fin
8. **Sin cobertura forex/acciones/cripto**: solo Polymarket, solo binarios

## Veto-loop: 6 reglas duras (filtro terminal, sin iteración)

Aplica en Brain, antes de sizing. Si alguna TRUE → mercado descartado sin appellatio.

| # | Regla | Significado |
|---|---|---|
| V1 | **Volumen insuficiente** | Vol 24h < 50,000 USD en el momento de evaluación |
| V2 | **Catalyst dentro de 72h de cierre del mercado** | Catalyst con resolución del MERCADO <72h en el futuro: ventana 72h-0h pre-resolución del mercado activa el veto |
| V3 | **Triggers vagos** | Trigger textual sin fecha concreta o sin referéndum asociado |
| V4 | **Chasing de momentum** | Precio se movió ≥8% en las últimas 4h previas al trigger formal |
| V5 | **Patrón débil** | Evento con <3 fuentes independientes o sin precedente análogo |
| V6 | **Sin trigger claro** | No hay un catalyst identificable en los próximos 7d con fecha de resolución |

## 4 procesos secuenciales

```
Scanner → Brain (veto) → Executor → Exit monitor
```

### 1. Scanner
- Poll 15 min a API Polymarket: `/markets?tag=...&volume_24h_gt=50000&closed=no`
- Filtro estático: binarios, volumen, solo categorías acordadas (Politics, Crypto, Sports, Tech)
- Sin filtro de edad de creación
- Escribe a cola FIFO (archivo JSONL)

### 2. Brain (veto-loop)
- Lee cola FIFO
- Aplica 6 reglas duras V1–V6 (determinista, datos on-chain + calendario)
- Si pasa: calcula Kelly 1.5%, verifica exposure cap + correlación + trigger ±2%
- Trigger ±2%: precio actual dentro del ±2% del precio en el momento de aprobación de la tesis (zona de ejecución)
- Si todo OK: escribe orden a pending (cola de ejecución)
- Si falla: descarta, log con razón

### 3. Executor
- Lee pending orders
- Envía límite a Polymarket a +2% del mid-price
- Si se llena: escribe posición activa a JSONL activo
- Timeout 30 min por orden

### 4. Exit monitor
- Lee posiciones activas cada 5 min
- Evalúa take-profit (+60%) y stop-loss (150 USD at risk)
- Si cualquiera se alcanza → envía orden de mercado para salir
- Log completo de cada salida

## Triggers y caps (resumen)

| Parámetro | Valor |
|---|---|
| Trigger único | Precio actual dentro del ±2% del momento de aprobación de tesis |
| Kelly | 1.5% fijo (150 USD/trade) |
| Exposure cap | 10% (5,000 USD) |
| Stop por posición | 150 USD at risk |
| Take-profit | +60% sobre precio entrada |
| Máx activos simultáneos | 3 |
| Máx misma categoría | 2, correlación <0.3 |
| Poll scanner | 15 min |
| Poll exit monitor | 5 min |
| Timeout executor | 30 min |

## Future work (fuera del alcance v1)
- **Max drawdown automático**: pausar el bot si drawdown alcanza 30% del bankroll inicial. No implementado en v1.
