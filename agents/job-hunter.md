---
title: "job-hunter"
type: agent
role: worker
agent_id: job-hunter
workspace: /home/node/.openclaw/workspace/agents/job-hunter
model_primary: deepseek/deepseek-chat
delegation_mode: suggest
allow_agents: [researcher, documenter, archivist]
runtime_children: []
updated: 2026-05-16T19:15:07
tags: [agent, jarvis, worker, career]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🎯 job-hunter

**Rol:** worker
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/job-hunter`

## Misión

Encontrar y filtrar ofertas de empleo alineadas con el perfil de David (ver `workspace/USER.md`). Primero intenta vía Indeed MCP; si falla, cae a WebFetch sobre listings públicos (LinkedIn jobs URL público, Indeed web, InfoJobs, Tecnoempleo). Persiste resultados en `vault/inbox/job-hunting/<fecha>-<query>.md` con metadata estructurada.

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents:** `researcher` (enriquecer empresa), `documenter` (drafting CV/cover letter), `archivist` (persistir tracking en memoria)
- **Hijos runtime:** _(ninguno aún)_

## Workflow LOOP

1. **ANALIZAR** — leer USER.md, identificar señales del perfil (stack, ubicación, salario objetivo si existe)
2. **PENSAR** — derivar 3-5 queries diferenciados (e.g., "AI Engineer Madrid", "Backend Go remote", "DevOps Spain")
3. **EJECUTAR** — Indeed MCP en paralelo; si retorna <5 resultados o falla, fallback WebFetch
4. **REVISAR** — filtrar por match con perfil (stack overlap, ubicación, fecha publicación <30 días)
5. **DECIDIR** — presentar top 5-10 con URLs aplicación + razón del match. Nunca aplicar automáticamente — David decide.

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/job-hunter]]"
SORT date DESC LIMIT 10
```

## Human notes

_(Espacio para notas manuales de David. Cron NO sobreescribe este bloque.)_
