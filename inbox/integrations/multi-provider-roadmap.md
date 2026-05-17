---
date: 2026-05-16
topic: multi-provider runtime
status: planned
priority: medium
related:
  - "[[03-decisions/2026-05-13-budget-guardian]]"
  - "[[03-decisions/2026-05-14-plan-v6-evolucion-jarvis]]"
  - "[[agents/index]]"
  - "[[inbox/integrations/indeed-mcp-status]]"
  - "[[projects/guia-fases-pendientes]]"
# auto-linked 2026-05-17
---


# Multi-provider runtime — roadmap

## Estado actual (cosmético)

Frontmatter de los 14 agentes declara `model_primary` por agente:

| Provider | Agentes (10/14) | Tipo trabajo |
|---|---|---|
| anthropic (10) | main, planner, researcher, debugger, tester, apier, designer, skill-reviewer, code-reviewer, auditor | Razonamiento, orquestación, decisiones críticas |
| deepseek (3) | documenter, archivist, job-hunter | Prosa estructurada, persistencia, orquestación tools |
| ollama-local (1) | monitor | Health checks repetitivos |

`openclaw.json` ahora tiene los 3 providers configurados con API key vía env vars.

## Lo que NO funciona todavía

Los crons existentes (`curator.py`, `orgchart.py`, `consolidate.py`) **NO llaman a LLMs**. Solo persisten estado. Por tanto:

- Cambiar `model_primary` en frontmatter **no reduce uso real de tokens hoy**.
- El consumo actual viene de sesiones Claude Code (chat con David) — un solo modelo, no por agente.

## Lo que falta para activar el ahorro

1. **API keys**:
   - `DEEPSEEK_API_KEY` en `/etc/environment` o `~/.openclaw/.env` (https://platform.deepseek.com → API keys)
   - `ANTHROPIC_API_KEY` si se quiere uso desde scripts (no necesario si tu suscripción Max es solo via claude.ai)

2. **Worker runtime** (proyecto serio, ~2-4h):
   - Daemon Python o Node que poll-ee `~/.openclaw/workspace/agents/<id>/inbox/*.md`
   - Por cada task encontrada:
     a. Leer frontmatter del agente target (model_primary)
     b. Mapear a provider correcto (anthropic / deepseek / ollama-local)
     c. Llamar API correspondiente con briefing + task
     d. Escribir reply a `vault/agents/<id>/replies/<task>-reply.md`
   - Manejar timeouts, errores, retries
   - Lock-file para evitar dobles procesamientos

3. **MCP integration en el runtime** — cuando el worker tenga acceso a tools, exponerlos vía el connector layer del dashboard (`/api/connectors/:id/tool/:tool`).

## Coste comparativo (precios aproximados 2026-Q2)

| Modelo | Input $/1M | Output $/1M | vs Sonnet output |
|---|---|---|---|
| Claude Opus 4.7 | $15 | $75 | 5x más caro |
| Claude Sonnet 4.6 | $3 | $15 | baseline |
| Claude Haiku 4.5 | $0.80 | $4 | 4x más barato |
| DeepSeek Chat | $0.27 | $1.10 | ~14x más barato |
| DeepSeek Reasoner | $0.55 | $2.19 | ~7x más barato |
| Ollama qwen 7b CPU | $0 | $0 | gratis pero ~30-60s/respuesta |

**Ahorro estimado** si runtime existe: ~30% de turnos shifted a DeepSeek + 1 worker a Ollama → ~25% de coste agregado de tokens, asumiendo carga uniforme.

## Caveat técnico

- Ollama qwen 7b en CPU **será lento** (sin GPU dedicada en este VPS). Para tareas no-time-critical (monitor health checks cada 30min) está bien. Para tareas interactivas, no.
- DeepSeek **soporta function calling** pero con menos pulido que Anthropic. Mantener tools simples (1-3 params) para evitar edge cases.
- Si DeepSeek devuelve garbage en algún edge case → fallback automático a Sonnet (implementar en runtime).
