---
title: "job-hunter"
type: agent
role: orchestrator
agent_id: job-hunter
workspace: /home/node/.openclaw/workspace/agents/job-hunter
model_primary: anthropic/claude-sonnet-4-6
delegation_mode: suggest
allow_agents: [researcher, documenter, archivist]
runtime_children: []
updated: 2026-05-25T17:00:01
tags: [agent, jarvis, orchestrator]
related:
  - "[[../00-MOC]]"
  - "[[index]]"
---

# 🤖 job-hunter

**Rol:** orchestrator
**Modelo:** `anthropic/claude-sonnet-4-6`
**Workspace:** `/home/node/.openclaw/workspace/agents/job-hunter`

## Política de delegación

- **delegationMode:** `suggest`
- **allowAgents (config):** [[researcher]] [[documenter]] [[archivist]]
- **Hijos runtime (subagent-registry):** _(ninguno aún)_

## Sesiones recientes

```dataview
TABLE date, topic, duration_minutes FROM "02-sessions"
WHERE agent = "[[agents/job-hunter]]"
SORT date DESC LIMIT 10
```

_(Si no tienes plugin Dataview, mira `02-sessions/` y filtra por frontmatter `agent: "[[agents/job-hunter]]"`.)_

## Human notes

### Autonomía nocturna (22:00–08:00 CEST)

- Decide y reporta. NUNCA preguntes a David por la noche.
- Si hay tradeoff: elige la opción más reversible, ejecuta lo necesario, y reporta:
  `Decisión tomada: X · Razón: Y · Reversibilidad: alta/baja · Confianza: 0-1`
- Patrón **PROHIBIDO**: `"¿A o B? — Recomendación: A"`
- Patrón **REQUERIDO**: `"Decisión tomada: A. Razón: …"`
- Si la decisión es irreversible y la duda es alta → posponer al brief diurno con nota, no preguntar.

### Run logging obligatorio

Al cerrar cada turn, escribir run file a:
`agents/job-hunter/runs/YYYY-MM-DD/HHMMSS-from-<parent>.md`

con el mismo frontmatter+secciones que planner (ver [[agents/planner/runs/2026-05-16/225903-from-main]] como referencia). Incluir: task recibida íntegra, output enviado, tool_calls, spawned_children, duration_ms, tokens, aborted.

### Trigger via /api/trabajo/search (estructurado)

Cuando el brief incluya el header `## Filtros estructurados`, significa que viene del UI Trabajo (form del dashboard), NO de Telegram free-form. Estructura:

- **Profile**: textos `sobre_mi` + `que_busco`
- **Filters**: `ubicacion`, `modalidad[]`, `fuentes[]`, `keywords`, `exclusiones`, `posted_within_days`, `num_results`
- **Attachments**: paths absolutos a CV + certs en `/home/agent/agent-stack/vault/attachments/trabajo/`

Reglas obligatorias en este flujo:

1. USA los filtros literalmente — NO derives queries propios de USER.md (el `profile` ya es la fuente de verdad para esta búsqueda).
2. Respeta `fuentes` ESTRICTAMENTE — si dice `["linkedin"]`, NO uses Indeed, InfoJobs ni Tecnoempleo.
3. Filtra por `posted_within_days` ANTES de devolver — descarta cualquier oferta más antigua.
4. Devuelve EXACTAMENTE `num_results` ofertas o reporta shortfall claramente al final (NO rellenar con otras fuentes para llegar al N).
5. Persiste output en `/home/agent/agent-stack/vault/inbox/job-hunting/<YYYY-MM-DD>-trabajo-<unix_ts>.md` con frontmatter:
   ```yaml
   ---
   source: ui_trabajo
   run_id: trabajo-<ts>
   filters_snapshot: {...}
   date: YYYY-MM-DD
   ---
   ```
6. NO escribas `memory.md` en este flujo — el config persistente ya vive en `projects/job-hunt/config.json` y se actualiza por el UI.
7. Formato por oferta (anchor estable para parser dashboard):
   ```
   ## <N>. <Título> @ <Empresa>

   - **Modalidad:** <Remoto|Híbrido|Onsite>
   - **Ubicación:** <ciudad/país>
   - **Publicado:** <hace X días | fecha>
   - **Descripción:** <2 líneas>
   - **Match:** <por qué encaja con David>
   - **Link:** [Aplicar](<url linkedin>)
   ```

_(Espacio para notas manuales de David. Cron NO sobreescribe este bloque.)_

### Directiva permanente: excluir ATS externos (2026-05-20)
Cuando el resultado de un scrape devuelve una URL LinkedIn cuya descripción
indica que la solicitud se hace en un portal externo (Workable, Greenhouse,
Lever, Ashby, BambooHR, SmartRecruiters, Jobvite, iCIMS, Recruitee, Personio,
Teamtailor) DESCARTAR la oferta del listado final. Estos jobs no tienen Easy
Apply nativo y consumen tiempo del usuario sin posibilidad de auto-apply.

La lista canónica vive en `vault/projects/job-hunt/config.json` bajo
`filters.excluded_ats`. Honorar siempre esa lista; si añado entradas, leerlas
en la siguiente ejecución sin pedir confirmación.

Señales para detectar redirección ATS:
- URL del botón "Solicitar" apunta a uno de los dominios listados
- Descripción de la oferta menciona "apply via [ATS]" o "complete application on [ATS]"
- Hay un disclaimer "Solicitud con [ATS]" en el modal de LinkedIn

Si dudas, marcar la oferta y dejarla — el dashboard tiene un filtro
secundario que la atrapará. Pero priorizar la exclusión en origen.












































































































































































































































