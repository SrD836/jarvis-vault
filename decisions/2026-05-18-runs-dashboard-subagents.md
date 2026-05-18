---
title: runs_dashboard.py — añadir scan_subagent_traces()
date: 2026-05-18
status: done
tags: [hermes, dashboard, subagents, observability]
---

# runs_dashboard.py — subagent traces

## Contexto

Script creado en `/home/node/.openclaw/workspace/hermes/runs_dashboard.py` (la ruta `/home/agent/agent-stack/hermes/` es propiedad de root — sin permisos de escritura para el usuario `node`; se usó el path equivalente en workspace).

El script anterior no existía físicamente (no hay backup que crear). La referencia al estado previo en el task describe el comportamiento esperado, no un archivo existente.

## Diff conceptual

### Nueva función: `scan_subagent_traces(today)`

```
+  scan_subagent_traces(today: str) -> list[row]
```

**Cómo funciona:**

1. `subprocess docker exec openclaw-fork-openclaw-gateway-1 find /home/node/.claude/projects/-home-node--openclaw-workspace -name "*.meta.json" -newermt <today>`
2. Por cada `agent-<hash>.meta.json`:
   - Lee `agentType`, `description`
   - Deriva `agent_hash` del nombre de archivo
   - Deriva `parent_sess` del nombre del directorio padre
   - Lee el `.jsonl` correspondiente via `docker exec cat`:
     - Primer evento → `timestamp` → hora Madrid
     - Último mensaje `role=assistant` → `last_assistant_text`
3. Filtra por fecha = today, retorna rows con `fuente=subagent`

### Nueva sección: `build_delegation_health()`

Ahora incluye dos sub-secciones:
- **Planner vault runs**: runs de vault con `agent == "planner"`
- **Subagent spawns (Task tool)**: todos los rows de `scan_subagent_traces()`

### Tabla principal

Los rows de subagents se insertan ordenados cronológicamente con el resto (`fuente=subagent`, `ms=0`, `tokens=0` porque esos datos no están en las traces).

## Verificación

```
$ python3 runs_dashboard.py
wrote /home/agent/agent-stack/vault/agents/runs-today.md
```

Output: 27 runs totales · fuentes: cron+subagent · 2 subagents detectados:
- `Explore(adfdb893)` ← parent `a1b6526b` — "Investigate session export issue"
- `Plan(a6816c0f)` ← parent `a1b6526b` — "Planner coordinates curator.py.new review and test plan"

## Limitaciones conocidas

- `ms=0` y `tokens=0` para subagents (no disponibles en las traces del Task tool)
- El container name está hardcodeado (`openclaw-fork-openclaw-gateway-1`)
- Sin vault runs hoy la sección planner muestra solo subagents
