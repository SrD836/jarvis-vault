---
title: JARVIS System Audit — multi-agent + memory + delegation + Obsidian
date: 2026-05-26
auditor: Claude Opus 4.7 (1M context)
scope: 13 agentes orgchart + Hermes feedback loop + pipeline delegación + nodos huérfanos Obsidian
status: COMPLETO
related:
  - "[[audit/2026-05-26-polymarket-bot-audit]]"
  - "[[audit/2026-05-24-vault-audit]]"
tags: [audit, jarvis, system, multi-agent, hermes, obsidian]
---

# JARVIS System Audit — 2026-05-26

## TL;DR

- **Memoria per-agente**: 13 agentes del orgchart, **solo 2 tienen memory.md** (`job-hunter` snapshot estático 3.5KB sin actualización desde 19-may; `polymarket-bot` único con loop activo 117KB). Los otros 11 son **vistas estáticas sin aprendizaje persistente**.
- **Hermes feedback loop**: detector funcionando (`learnings.py` corrió 04:00 UTC hoy y produjo 2 propuestas), **aplicación rota** — propuestas languish con `applied:false`, ningún hook auto-parchea AGENTS.md.
- **Pipeline delegación**: planner registró **0 runs en 7 días**. Última actividad real fue 2026-05-19 (8 runs), 4 abortaron por iter-cap saturado. **Pipeline está dormido** post-19-may.
- **Obsidian huérfanos**: 1,399/1,589 (88%). 1,140 son polymarket decisions sin indexar — fix existe (`hermes/backfill_links.py`).

---

## Sección 1 — Memoria de los 13 agentes

**Verdict: NO LEARNING per-agent (excepto polymarket-bot)**

### Tabla estado por agente

| Agente | memory.md | runs/ | Última mtime memory | Estado |
|---|---|---|---|---|
| main | ❌ | 0 (cero) | — | STATIC (view) |
| planner | ❌ | 8 (todos 19-may) | — | STATIC (view) |
| apier | ❌ | 0 | — | STATIC |
| archivist | ❌ | 0 | — | STATIC |
| auditor | ❌ | 0 | — | STATIC |
| code-reviewer | ❌ | 0 | — | STATIC |
| debugger | ❌ | 0 | — | STATIC |
| designer | ❌ | **41** (cron renders) | — | STATIC + render activity |
| documenter | ❌ | 0 | — | STATIC |
| job-hunter | ✅ 3.5KB | 44 | 2026-05-19 00:30 | SNAPSHOT (preferences cache, no learning) |
| monitor | ❌ | 0 | — | STATIC |
| researcher | ❌ | 3 | — | STATIC |
| skill-reviewer | ❌ | 0 | — | STATIC |
| tester | ❌ | 0 | — | STATIC |
| **polymarket-bot** *(referencia)* | ✅ **117KB** | 0 | 2026-05-26 16:30 | **LEARNING activo** (M1/M2/M3) |

### Hallazgos

- **11 de 13 agentes son STATIC**: `agents/<id>.md` regenerado cada 30min por orgchart cron desde AGENTS.md (memoria `jarvis_main_md_regenerated_by_cron`). No tienen estado dinámico per-sesión que persista.
- **job-hunter es SNAPSHOT, no LEARNING**: su memory.md es perfil de David estable (edad, ubicación, idiomas, skills). No registra "intenté X estrategia y falló". Sin update desde 19-may.
- **polymarket-bot es el ÚNICO ejemplo funcional** del patrón propuesto por el usuario (estilo Hermes): brain.go consulta `memory.md` en CADA decisión y veta si match histórico. Es lo que el resto debería hacer.

### Brecha vs intent del usuario

Usuario pidió: *"con cada operación errónea aprende para no volver a cometer el mismo error"*.

Estado real para los 13: **ningún agente fuera de polymarket-bot tiene mecanismo de aprendizaje por error**. Las trazas en `runs/<agent>/YYYY-MM-DD/*.md` existen pero no se condensan en lecciones reutilizables.

---

## Sección 2 — Hermes feedback loop

**Verdict: PARCIAL — detector OK, aplicación NO existe**

### Flujo `hermes/learnings.py` (mtime 2026-05-16)

| Step | Función | Estado |
|---|---|---|
| 1 | `collect_runs()` — escanea `vault/agents/*/runs/**` últimos 7d | ✅ funciona |
| 2 | `detect_patterns()` — agrupa por agent: aborts, iter-cap, tokens>80% | ✅ funciona |
| 3 | `deepseek_propose()` — LLM auditor propone mejora | ✅ funciona |
| 4 | `write_evolution()` — escribe `vault/wiki/agent-evolution/{agent_id}.md` con `proposed:true / applied:false` | ✅ funciona |
| 5 | **Aplicar propuesta → parchear AGENTS.md / agente específico** | ❌ **NO EXISTE** |

### Cita literal del código

`learnings.py:349`:
> *"Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio."*

### Estado runtime (2026-05-26 16:30 UTC)

- Archivos en `vault/wiki/agent-evolution/`: **2** (`planner.md`, `researcher.md`)
- Ambos con timestamp **2026-05-26 04:00** UTC (cron Hermes corrió esta madrugada)
- Frontmatter de `planner.md`:
  ```yaml
  pattern: iter-cap-saturated
  confidence: medium
  proposed: true
  applied: false
  ```
- Propuesta concreta: reducir `allow_agents` de 11 a 5 nucleares (researcher, code-reviewer, debugger, tester, documenter) para bajar iter-cap saturation
- **Ningún archivo tiene `applied: true`**. La propuesta lleva 12+ horas sin aplicar.

### Brecha vs Hermes-style auto-improvement

El patrón Hermes en su forma completa requiere:
1. ✅ Detección (hecha)
2. ✅ Propuesta (hecha vía LLM)
3. ❌ **Aplicación automática o semi-automática con review humano** (ausente)
4. ❌ **Validación post-aplicación** (¿la mejora redujo el patrón observado en el siguiente ciclo de 7d?)

El sistema actual es **observability + recommender**, no **self-improvement**. La diferencia clave: las propuestas nunca llegan al system prompt real.

---

## Sección 3 — Pipeline delegación (main → planner → subagentes)

**Verdict: DORMIDO desde 19-may + 4 bugs estructurales**

### Tabla de actividad real

| Métrica | Valor | Evidencia |
|---|---|---|
| Planner runs últimos 7d | **0** | `find vault/agents/planner/runs -mtime -7` → vacío |
| Total planner runs históricos | 8 | Todos del 2026-05-19 |
| Runs con `spawned_children` no vacío | 1/8 (12.5%) | Solo `095210-from-main.md` delegó a researcher |
| Runs aborted con `iter-cap-saturated` | 4/8 (50%) | Detectado por Hermes, propuesta sin aplicar |
| Runs hoy con `parent ≠ cron` | **0/49** | runs-today.md líneas 17-19 confirman |

### Trace de 5 planner runs (todos del 19-may)

| Run | Parent | Spawned | Status | Veredicto |
|---|---|---|---|---|
| `095210-from-main` | main | `[researcher]` | aborted | ✅ Delegó correctamente, abortó después |
| `101321-from-main` | main | `[]` | token-limit | ❌ Auto-ejecución sin delegar |
| `104519-from-main` | main | `[]` | iter-cap | ❌ Saturó loop sin delegar |
| `105317-from-main` | main | `[]` | iter-cap | ❌ Mismo patrón |
| `105919-from-main` | main | `[]` | iter-cap | ❌ Mismo patrón |

**90% de runs auto-ejecutaron contra la regla `planner.md:54`**: *"si tu run termina con `spawned_children: []` pero la task era no-trivial, eso es autoejecución prohibida"*.

### Bugs estructurales identificados

| # | Bug | Severidad | Reproducción |
|---|---|---|---|
| **D-1** | Pipeline dormido — 7 días sin runs planner | **HIGH** | `find vault/agents/planner/runs -mtime -7 \| wc -l` → 0 |
| **D-2** | `spawned_from` etiqueta incorrecta | MEDIUM | researcher runs reportan `parent: main` cuando vinieron de planner. `parent_chain` real es `["main","planner","planner"]` |
| **D-3** | `runs_dashboard.py` no detecta delegación intra-sesión | MEDIUM | runs-today.md:17 dice "(sin runs de planner hoy en vault — el runtime aún no escribe per-turn runs)". Dashboard solo lee cron runs |
| **D-4** | Formato "autonomía nocturna" sin enforcement de código | LOW | Política en `AGENTS.md` `## Human notes` pero ningún validator runtime obliga el formato `Decisión · Razón · Reversibilidad · Confianza` |

### Causa raíz probable de D-1

Hipótesis: tras el incidente del 19-may con iter-cap saturation, el bot Telegram dejó de invocar planner por cualquiera de:
1. Reglas en main.md desactivaron triggers de delegación
2. El usuario dejó de mandar tareas que requieran planner (todos sus mensajes recientes son polymarket/audit, ejecutados directamente)
3. Bug silencioso en main → planner spawn que no loguea

No verificable sin trazas Telegram inbound o logs `claude -p` del gateway.

---

## Sección 4 — Nodos huérfanos en Obsidian

**Verdict: 88% huérfanos pero MAYORÍA REPARABLE con tool existente**

### Conteo por carpeta (1,589 .md totales)

| Carpeta | Huérfanos | % vault | Naturaleza | Acción |
|---|---|---|---|---|
| `03-decisions/` | **1,140** | 71.7% | Polymarket decisions sin indexar | `hermes/backfill_links.py` (existe, no se ejecuta) |
| `projects/` | 216 | 13.6% | 171 = `node_modules`/licenses (ruido), 45 = orphan reales | Limpiar deps + crear `projects/index.md` |
| `inbox/` | 27 | 1.7% | 24 en `.processed/` (legítimo), 3 WIP stale | Revisar 3 WIP |
| `wiki/` | 11 | 0.7% | agent-evolution dormido (planner+researcher) | **Critical — son los que paran Hermes** |
| Otros (root, .agent, analysis, audit, skills) | 5 | 0.3% | 1-2 por carpeta, marginal | Cosmético |
| **Total** | **1,399** | **88.0%** | | |

### Hallazgos críticos por carpeta

- **03-decisions/index.md tiene 7 entradas, decisions reales 1,157**: 99.4% sin indexar. `hermes/backfill_links.py` existe pero no se ejecuta como cron. Una sola ejecución resuelve el bulk del problema.
- **wiki/agent-evolution/ con 2 propuestas dormidas**: esto NO es ruido — es exactamente el síntoma de Sección 2 (loop Hermes roto). Los huérfanos aquí indican que las propuestas NUNCA se enlazaron desde un workflow de revisión.
- **projects/ con 171 node_modules huérfanos**: artefacto de `npm install` en repos clonados. Debería estar en `.gitignore` + excluido de Obsidian (`.obsidianignore` o equivalente).

### Carpetas dormidas verificadas

- `seeds/` — 3 archivos (decisiones-clave, operacion-diaria, sistema-arquitectura) → **están en MOC** ✓
- `memories/` — system-memory + temporal-log → sin backlinks pero `intentional` (dormido por diseño)
- `Research/`, `secrets/` — vacíos por diseño ✓

---

## Sección 5 — Bugs encontrados

| # | Severidad | Bug | Archivo | Reproducción |
|---|---|---|---|---|
| **H-1** | HIGH | Hermes propuestas languish sin aplicar | `hermes/learnings.py:349` (manual merge required) | `grep "applied: false" vault/wiki/agent-evolution/*.md \| wc -l` → 2 archivos sin aplicar 12h+ |
| **M-1** | HIGH | 11 de 13 agentes orgchart sin `memory.md` | `vault/agents/<id>/` | Ver tabla Sección 1. Solo polymarket-bot + job-hunter tienen memory file |
| **D-1** | HIGH | Pipeline planner dormido 7 días | `vault/agents/planner/runs/` | `find vault/agents/planner/runs -mtime -7 \| wc -l` → 0 |
| **D-2** | MEDIUM | `spawned_from` mal etiquetado | `agents/researcher/runs/*-from-main.md` cuando deberían ser `-from-planner` | grep cruzado: `parent: main` en researcher runs cuyo `parent_chain` incluye planner |
| **D-3** | MEDIUM | `runs_dashboard.py` no detecta delegación intra-sesión | `hermes/runs_dashboard.py` | runs-today.md:17 "(sin runs de planner hoy en vault)" pese a actividad real esperada |
| **O-1** | MEDIUM | 1140 decisions huérfanas | `vault/03-decisions/index.md` lista 7/1157 | `wc -l vault/03-decisions/index.md` vs `ls vault/03-decisions/*.md \| wc -l` |
| **O-2** | LOW | 171 node_modules .md en projects/ | `vault/projects/*/node_modules/` | Esos archivos NO deberían estar en Obsidian vault |
| **D-4** | LOW | Formato autonomía nocturna sin enforce | Política en AGENTS.md | No hay validator runtime |

---

## Sección 6 — Acciones recomendadas (priorizadas)

| Prio | Acción | Esfuerzo | Impacto | Archivo |
|---|---|---|---|---|
| **P0** | Implementar auto-apply de `agent-evolution/*.md` (semi-supervisado, requiere review en Telegram) | 4-6h dev | Cierra loop Hermes — único cambio que activa auto-mejora real | `hermes/apply_evolution.py` (nuevo) + parcheo `AGENTS.md` |
| **P0** | Aplicar manualmente la propuesta planner.md ahora (limitar `allow_agents` a 5) | 10 min manual | Reduce iter-cap saturation, desbloquea pipeline | `agents/planner/openclaw.json` |
| **P1** | Diagnosticar D-1 (pipeline dormido 7d) | 1h investigación | Si planner está roto, todo el sistema multi-agente está roto | `openclaw-fork/extensions/telegram/src/` + gateway logs |
| **P1** | Ejecutar `hermes/backfill_links.py` para indexar 1140 decisions | 5 min | -88% huérfanos en una corrida | `vault/03-decisions/index.md` |
| **P1** | Implementar memory.md para los 11 agentes restantes — copy pattern polymarket-bot | 1-2 días | Habilita Hermes a producir propuestas más específicas (más context) | `agents/<id>/memory.md` (×11) |
| **P2** | Fix D-2: corregir label `spawned_from` en cron runs | 30 min | Trazas correctas, debugging más fácil | Script que escribe los runs (encontrar) |
| **P2** | Fix D-3: ampliar `runs_dashboard.py` para leer `~/.claude/projects/.../subagents/*.meta.json` | 1h | Dashboard refleja delegación real | `hermes/runs_dashboard.py` |
| **P2** | Excluir `node_modules/` del vault Obsidian | 5 min | -171 huérfanos cosméticos | `.gitignore` vault + `.obsidianignore` |
| **P3** | Crear `projects/index.md` con 5 proyectos activos | 15 min | Cosmético, ayuda navegación | `vault/projects/index.md` |
| **P3** | Reactivar wiki/agent-evolution con enlace desde MOC | 5 min | Visibilidad de propuestas pendientes | `00-MOC.md` añadir sección |

### Top-3 si solo hay tiempo para 3 acciones

1. **P0** — Ejecutar `backfill_links.py` (5 min, mata el 71% de huérfanos)
2. **P0** — Aplicar propuesta planner.md manualmente (10 min, desbloquea iter-cap)
3. **P1** — Investigar por qué planner está dormido 7d (1h) → si está roto técnicamente, todo lo demás importa menos

---

## Conclusión global

**El sistema JARVIS tiene esqueleto correcto pero carne incompleta:**

- **Polymarket-bot funciona como debe** — feedback loop activo, memoria que se consulta en cada decisión, learnings concretos.
- **Hermes detecta correctamente** pero no cierra el círculo. Genera 2 propuestas/semana que nadie aplica.
- **Pipeline delegación está documentado en main.md/planner.md** con reglas anti-fraude correctas, pero **lleva 7 días sin actividad**. Bien diseñado, no usado.
- **Obsidian tiene tooling existente** (`backfill_links.py`) para resolver el 71% de huérfanos en 5 minutos. No se ejecuta.

**El gap principal NO es de diseño** — es de **operacionalización**: scripts existen, propuestas se generan, pero nada cierra el ciclo automáticamente. Cada vez que aparece una mejora hay que pedir a David que la aplique manualmente.

**Recomendación strategic**: priorizar P0-P1 esta semana. Sin cerrar el loop Hermes (auto-apply) y sin reanimar el pipeline planner, el resto de mejoras son cosméticas.

---

## Anexo — Comandos para re-auditar

```bash
# Memoria per-agente
ssh agent@88.198.168.61 "for a in main planner apier archivist auditor code-reviewer debugger designer documenter job-hunter monitor researcher skill-reviewer tester polymarket-bot; do echo \$a: \$([ -f /home/agent/agent-stack/vault/agents/\$a/memory.md ] && echo MEMORY || echo NO); done"

# Estado Hermes feedback
ssh agent@88.198.168.61 "ls -la /home/agent/agent-stack/vault/wiki/agent-evolution/; grep -l 'applied: true' /home/agent/agent-stack/vault/wiki/agent-evolution/*.md 2>/dev/null | wc -l"

# Pipeline planner activity
ssh agent@88.198.168.61 "find /home/agent/agent-stack/vault/agents/planner/runs -mtime -7 -name '*.md' | wc -l"

# Huérfanos count
ssh agent@88.198.168.61 "find /home/agent/agent-stack/vault -name '*.md' | wc -l; cat /home/agent/agent-stack/vault/03-decisions/index.md | grep -c '\[\['"
```
