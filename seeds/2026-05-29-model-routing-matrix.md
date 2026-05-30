---
title: "Matriz de routing de modelos — agente x contexto"
type: seed
date: 2026-05-29
status: vigente
related:
  - "[[03-decisions/2026-05-29-polymarket-daily]]"
  - "[[CLAUDE]]"
  - "[[agents/CATALOG]]"
  - "[[agents/index]]"
  - "[[agents/planner]]"
  - "[[agents/polymarket-bot]]"
  - "[[agents/skill-dispatcher]]"
tags: [seed, arquitectura, llm-routing, claude, max, anti-baneo]
# auto-linked 2026-05-30
---


# Matriz de routing de modelos (2026-05-29)

**Objetivo:** usar Claude (incl. Opus 4.8) en los agentes interactivos SIN uso automatizado de Max (riesgo de baneo). Eje de decision: **interactivo (David presente) = Claude OK; cualquier cron/unattended = Claude IMPOSIBLE**.

## Regla de oro
- **Interactivo** (David dispara via gateway/Telegram/dashboard, o `planner_route --allow-claude` en vivo) -> **Claude permitido**.
- **Automatizado** (cualquier cron, dia o noche) -> **Claude bloqueado por diseno**: bloqueo duro en `llm_call` + crons del gateway desactivados.

## Matriz agente -> modelo -> contexto

| Agente | Interactivo (gateway/live) | Automatizado (cron) |
|---|---|---|
| main | **claude-opus-4-8** (fb opus-4-7) | no corre en cron |
| planner | **claude-opus-4-8** (fb opus-4-7) | DeepSeek (hermes; cron-block) |
| polymarket-handler | **claude-opus-4-8** | DeepSeek |
| code-reviewer | **claude-opus-4-8** | DeepSeek |
| debugger | **claude-opus-4-8** | DeepSeek |
| researcher, apier, documenter, designer, tester, auditor, skill-reviewer, job-hunter, skill-dispatcher | claude-sonnet-4-6 (default) | DeepSeek / Ollama |
| archivist, monitor | claude-haiku-4-5 | DeepSeek / Ollama |
| polymarket-bot (scanner/brain/executor/exit_monitor) | — | deepseek-chat / Go (nunca Claude) |
| crons gateway (curator-weekly, daily-brief, decision-log, skill-reviewer-fork) | — | **DESACTIVADOS** (eran Claude) -> equivalentes hermes DeepSeek |

## Mecanismo (3 capas)

1. **Path 1 — gateway runtime** (`~/.openclaw/openclaw.json`): `agents.list[].model.primary`. main/planner/polymarket-handler/code-reviewer/debugger = `anthropic/claude-opus-4-8` (+fallback opus-4-7). Resto hereda default `sonnet-4-6`. Solo afecta al runtime INTERACTIVO.
2. **Path 2 — hermes scripts/crons** (`hermes/llm_call.py`): bloqueo duro — `if provider=="claude" and os.environ["JARVIS_EXEC_CONTEXT"]=="cron": return claude(cron-blocked)`, **ignora allow_claude**. El contexto lo setean el crontab (env global `JARVIS_EXEC_CONTEXT=cron`) y `run_logger.py`. Segunda capa: el guard `allow_claude=False` por defecto.
3. **Path 3 — gateway jobs.json** (`hermes/cron_claude_guard.py`, cron 04:20): desactiva cualquier job `agentTurn` (=Claude programado) y alerta. Mantiene la invariante "cero Claude programado en el gateway".

## Brief migrado a DeepSeek
El `daily-brief` (gateway, Claude 04:00) se reemplazo por `hermes/daily_brief.py` (DeepSeek, cron 06:00) -> `vault/01-briefs/<date>.md` con el mismo formato chief-of-staff. **Caveat:** pierde calendar-MCP y memory_search (sintetiza desde ficheros del vault). Coste 0, sin riesgo de baneo.

## Para promover/degradar un agente
- Interactivo: editar `agents.list[].model.primary` en `openclaw.json` + restart gateway (validacion solo en startup). Reversible (1 linea/agente).
- Worker en vivo a Opus: `export LLM_CLAUDE_MODEL=claude-opus-4-8` antes de `planner_route --allow-claude`.

## Verificado (2026-05-29)
- Traza runtime: `openclaw agent --agent main/planner --json` -> `"model":"claude-opus-4-8"`, `"winnerModel":"claude-opus-4-8"`.
- Bloqueo: `JARVIS_EXEC_CONTEXT=cron ... allow_claude=True` -> `claude(cron-blocked)` (loggeado en runs.jsonl).
- jobs.json: 0 jobs agentTurn habilitados tras el guard.
- Brief: `daily_brief.py` escribe via `deepseek-chat` (result ok).
