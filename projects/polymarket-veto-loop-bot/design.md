---
name: Polymarket Veto-Loop Bot — Design Consolidado
version: 2.0
date: 2026-05-19
status: approved (discusión David + main)
supersedes: project.md (v1)
---

# Polymarket Veto-Loop Bot — Diseño Definitivo

## Arquitectura (4 procesos secuenciales)

1. **Scanner** — Sondea mercados del universo permitido, genera tesis candidatas con precio actual y catalizador.
2. **Brain (veto-loop)** — Auditor LLM que aplica reglas duras (abajo). **Filtro binario terminal**: veta o pasa. No itera.
3. **Executor virtual** — Aplica sizing y cola FIFO. No mueve dinero real.
4. **Exit monitor** — Observa posiciones abiertas y dispara stop/take cuando corresponde.

## Reglas duras del Brain (veto si alguna falla)

- **Volumen** < 50,000 USD → veto
- **Catalyst 72h**: si hay catalizador previsto en las próximas 72h previas a resolución → veto. Ventana: 72h-0h antes de resolución.
- **Triggers vagos**: "cuando suba el sentimiento", "si el mercado reacciona", etc. no medibles → veto
- **Chasing**: precio ya se movió ≥ 8% en últimas 4h → veto
- **Patrones**: respaldados por < 3 casos → veto
- **Sin triggers claros**: no hay evento/catalizador concreto que mueva el precio → veto

## Sizing y stops (Executor)

- **Bankroll virtual**: 10,000 USD
- **Kelly**: 1.5% del bankroll por trade (= 150 USD)
- **Stop-loss**: USD-at-risk = Kelly = 150 USD. Se dispara cuando la pérdida en USD alcanza 150 USD. NO es % del precio del activo.
- **Take-profit**: +60% del precio de entrada (en % sobre precio, este sí va sobre precio)
- **Trigger entrada**: precio actual debe estar dentro de ±2% del precio al que se aprobó la tesis. Si no, descarte.

## Reglas de portfolio

- **Correlación dura**: máx 3 posiciones simultáneas en misma categoría + misma temática
- **Exposure cap**: 10% del bankroll total en posiciones abiertas (= 1,000 USD)
- **Universo**: Politics, Crypto, Sports, Tech
- **Excluye**: Pop Culture

## Cola de ejecución

- **FIFO** por timestamp de aprobación del Brain
- Si al llegar el turno el cap de correlación o exposure bloquea → se descarta (no se reencola)
- Trigger precio ±2% también se evalúa en este momento; si no está en rango → descarte

## Logging

- JSONL estructurado en `vault/inbox/trading/`
- Cada evento: scaneo, veto/aprobación, entrada, salida (stop/tp/manual)

## Dashboard

- Nueva página **Bot** en el sidebar del dashboard React
- Muestra: P&L acumulado, lista de trades simulados (estado, entry, current, stop, tp, P&L)

## Stack conocido del dashboard

- React (dashboard existente en jarvis-dashboard)
- Sidebar ya tiene páginas existentes
- La página Bot se añade al mismo routing

## Cron

- Ejecución periódica (cada N minutos/horas) para correr Scanner + Brain + Executor + Exit monitor

## Fuera de alcance

- Dinero real
- Wallets reales
- Estrategias fuera del universo
- Sharpe ranking (descartado en discusión)
- Stop como % del precio (descartado)
