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
_(no se toca por automatización)_
