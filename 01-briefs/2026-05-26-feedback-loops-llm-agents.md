---
title: "Brief — Feedback loops 2026 aplicables a JARVIS"
type: brief
date: 2026-05-26
source: main + researcher
tags: [brief, agentes, feedback-loops, arquitectura, self-improvement]
related:
  - "[[00-MOC]]"
  - "[[seeds/sistema-arquitectura]]"
  - "[[seeds/decisiones-clave]]"
  - "[[seeds/operacion-diaria]]"
---

# Feedback loops 2026 — 3 patrones que JARVIS no tiene aún

## TL;DR

Lo que JARVIS ya tiene cubre **reflexión-sobre-agregado** (curator diario, skill catalog con `use_count`, session-export) y **control de recursos** (budget-guardian, caps por goal). Los tres patrones siguientes son los de mejor evidencia 2025–2026 que **no están implícitos** en la arquitectura actual. Cada uno se monta sobre el vault y los crons existentes; ninguno requiere infra nueva.

**Pushback de entrada:** no añadir Reflexion ni Self-Refine. Hay evidencia 2025 sólida de que self-critique prompted **degrada** en modelos por debajo de ~70B y empate-o-peor por encima en tareas knowledge-heavy (d = −0.14 a −0.33; [arXiv 2601.00513](https://arxiv.org/pdf/2601.00513)). El sustituto correcto es verificación externa (patrón #2).

---

## Patrón 1 — Trajectory→Regression Flywheel

**Definición.** Clusterizar los traces reales por modo de fallo y promover el fallo canónico de cada cluster a un set de regresión replay-able que actúa como **gate** antes de aplicar cualquier cambio (skill nueva, prompt de planner editado, modelo enrutado distinto).

**Failure mode que ataca.** Bugs de coordinación recurrentes que **parecen novedosos cada vez**. Estudios de despliegues reales reportan tasas de fallo 41–86.7% y la causa dominante son ~2–3 clusters que concentran la masa de errores (en uno publicado: routing 5.25% + rephrasal 3.2% explicaban la mayoría — solo arreglables tras *nombrarlos*). Tienes los traces (session-export cron), te falta la capa de clustering + regression-set.

**Mecanismo concreto sobre JARVIS:**

- **Mide:** cada run del session-export ya emite outcome, worker, errores de tool, retries, gasto. Suficiente substrate.
- **Escribe:** nuevo agente `trace-analyst` (cron OpenClaw semanal, p.ej. domingo 04:00). Embed con `nomic-embed-text` local + HDBSCAN. Emite `vault/regressions/<cluster-id>.md` con: prompt mínimo de replay + invariantes esperados + worker(s) responsables.
- **Consume:** harness CI-style que re-corre el set antes de:
  - promover skill `pending` → `active` (hoy el meta/curator semanal hace solo análisis semántico, sin gate de regresión).
  - aplicar cualquier edit al briefing de `planner` o `main`.
- **Segundo loop hacia curator:** clusters donde la skill responsable tiene `use_count` alto pero success-rate bajo → la skill se **demota**, no se preserva por uso.
- **Cierra cuando** el cluster pasa N corridas seguidas → se archiva.

**Coste.** Embeddings + clustering semanal sobre cientos de traces es sub-dólar (Ollama local). El coste real es el harness de replay: ~5–10k tokens × |regression set| por gate. Empezar con set ≤ 10 fallos canónicos.

**Por qué encaja en JARVIS ya:** el slot está vacío. `meta/curator` (lunes 03:00) propone merge/quarantine semántico sin gate empírico. Este patrón cierra ese hueco.

**Referencias:**
- [LangChain — The Agent Improvement Loop Starts with a Trace](https://www.langchain.com/blog/traces-start-agent-improvement-loop)
- [Augment Code — Agent Learning Flywheel](https://www.augmentcode.com/guides/agent-learning-flywheel) (números reales de clusters dominantes)
- [Adaptive Data Flywheel — arXiv 2510.27051](https://arxiv.org/html/2510.27051v1) (MAPE control loop formalizado)

---

## Patrón 2 — Step-level Verification Gate en handoffs worker→worker

**Definición.** Un `verifier` ligero **externo** valida cada artefacto de handoff contra el contrato del planner **antes** de que arranque el siguiente worker. **No** es self-critique del worker.

**Failure mode que ataca.** Hallucinación compuesta a lo largo de la cadena: cada worker es localmente correcto, el output integrado del run no cumple el goal. Reflexion/Self-Refine atacan el mismo problema **mal**: convergencia ~75% en rondas 1–2, deeper-loops queman tokens, y en modelos pequeño/medio **degradan**. AgentPRM (process reward model externo) es >8× más compute-efficient que el baseline de reflexión ([arXiv 2511.08325](https://arxiv.org/abs/2511.08325)).

**Mecanismo concreto sobre JARVIS:**

- **Mide:** cada `vault/agents/<worker>/replies/<task>-reply.md` pasa por verifier inmediatamente al ser escrito.
- **Escribe:** verifier emite hermano `.verdict.md` con `verdict: pass | retry | escalate` + razón en 1 línea. El verifier corre en modelo **más barato** que el worker (Haiku-class, o Qwen2.5-Coder 7B local para verificaciones de forma; Sonnet 4.6 solo cuando el verdict requiere juicio semántico).
- **Rubric fijo y corto** (no inventes uno nuevo por dominio — lo mata):
  1. ¿Satisface los success-criteria que el planner escribió en el task?
  2. ¿Contradice estado previo del vault?
  3. ¿Outputs de tool citados?
- **Consume:** `main` lee `.verdict.md` antes de schedular el siguiente worker.
  - `retry`: re-run del **mismo** worker con la razón del verifier inyectada. Máx **1 retry**. Más allá → escalate.
  - `escalate`: bounce a planner para re-decomposición.
- **Worker no se auto-critica.** Importante para que la responsabilidad sea atribuible y el coste, acotado.

**Coste.** ~500–2000 tokens por handoff. Con ~12 workers × ~3 handoffs medio por run = ~25k tokens extra/run. <5% overhead si el verifier prompt se mantiene estático. Si crece el prompt → se vuelve un segundo worker, anti-patrón.

**Pushback ready:** si alguien sugiere "que el worker se revise a sí mismo" → no. Misma inferencia, mismos blind spots. La externalidad es la propiedad clave.

**Referencias:**
- [AgentPRM — arXiv 2511.08325](https://arxiv.org/abs/2511.08325) (canónico 2025 sobre process reward para agentes)
- [VERIMAP — arXiv 2510.17109](https://arxiv.org/pdf/2510.17109) (verificación en boundaries de un planner-DAG, isomórfico a nuestros handoffs)
- [Reflect, Retry, Reward — arXiv 2505.24726](https://arxiv.org/html/2505.24726v1) (calibra cuándo retry-once funciona: baja accuracy inicial + señal externa)

---

## Patrón 3 — Goal-Anchor Loop (PLAN.md externalizado + drift monitor)

**Definición.** El output de `planner` se commitea a un artefacto estable `vault/runs/<id>/PLAN.md` tratado como **ground truth**. Cada worker re-lee el goal al entrar; un `drift-monitor` cheap chequea divergencia cada K tool-calls.

**Failure mode que ataca.** Goal drift en runs largos. Performance agéntica cae >50% a ~100K tokens incluso en modelos de 1–2M ([Drift No More? arXiv 2510.07777](https://arxiv.org/html/2510.07777v1)). JARVIS tiene vault pero **el plan no es un artefacto first-class**: vive en el thread de chat y en handoffs sueltos. Con 12 workers especializados es el killer silencioso: cada uno localmente correcto, el run no resuelve la tarea original.

**Mecanismo concreto sobre JARVIS:**

- **Mide:** similaridad coseno (con `nomic-embed-text`) entre el subgoal declarado por el worker activo y el objetivo raíz de `PLAN.md` + contador acumulado de tool-calls desde el último anchor read.
- **Escribe:**
  - `planner` siempre estampa `vault/runs/<id>/PLAN.md` con frontmatter (objetivo, success-criteria, sub-tasks). Es **inmutable** salvo re-plan explícito.
  - Cada worker prepende 1 línea "mi subgoal vs PLAN.md root" en su reply (auto-anchor).
  - Nuevo agente `drift-monitor` cheap (Qwen2.5 local) corre cada K=10 tool-calls *o* cada M=10 min sobre el run activo.
- **Consume:**
  - Si drift > umbral → `main` fuerza **reconciliación**: invoca `planner` con state actual + PLAN.md original. Planner emite sello "still-valid" *o* un replan diff explícito.
  - Workers leen siempre la versión **estampada más reciente** de `PLAN.md`, nunca el thread de chat.
- **Cierra cuando** el run termina con `outcome` y goal-vs-output score loggeado — esto alimenta el **patrón #1** (los clusters de drift se vuelven regresiones).

**Coste.** PLAN.md externalizado es estructural, ~0 tokens. Drift monitor: ~300–500 tokens por check. Replans son raros pero caros — gateados por budget-guardian.

**Sinergia importante:** loops #1 y #3 se realimentan. Drift detectado → regression cluster → próximo run con `planner` editado se gatea contra ese cluster.

**Referencias:**
- [Anthropic — Multi-agent coordination patterns](https://claude.com/blog/multi-agent-coordination-patterns) (endorsa `progress.txt`-style state externo, mismo principio)
- [CORAL — Don't Lose the Thread (OpenReview)](https://openreview.net/forum?id=NBGlItueYE) (asignación proactiva de contexto contra drift)
- [Agent Drift — arXiv 2601.04170](https://arxiv.org/html/2601.04170v1) (drift-aware routing + behavioral anchoring para multi-agent)

---

## Priorización propuesta

1. **#2 (verification gate) primero.** Atajable en 1–2 semanas: rubric fijo + un agente `verifier` + hook en `main` antes de schedular siguiente worker. ROI inmediato, sin esperar volumen de traces.
2. **#3 (PLAN.md + drift) segundo.** Refactor pequeño: que `planner` estampe a `vault/runs/<id>/PLAN.md`; el drift-monitor se puede iniciar con umbral conservador.
3. **#1 (regression flywheel) tercero.** Requiere acumular meses de traces de calidad — ahora session-export es reciente. Empezar a coleccionar ya, montar el `trace-analyst` cuando haya ≥200 runs.

## Decisiones que esto exige tomar

- ¿Verifier en modelo local (Qwen2.5) o Haiku-class vía API? Recomendación: empezar con Sonnet 4.6 para tener señal limpia, mover a local una vez la rubric esté estabilizada.
- ¿`drift-monitor` corre por tiempo o por tool-calls? Recomendación: tool-calls (K=10). Tiempo es ruidoso con tareas heterogéneas.
- ¿Skill promotion automática gateada por regression set, o solo informativa? Recomendación: hard-gate cuando el set tenga ≥5 fallos canónicos. Antes, informativo.

## Anti-patrones que NO añadir (descartados en la investigación)

- Reflexion / Self-Refine clásicos — ver pushback arriba.
- Deep verifier loops (3+ rondas iterando misma check) — wasted tokens, plateau en ronda 1–2.
- Human-in-the-loop como feedback signal: David es operador único, el canal no escala. Sustituir por verifier programático + escalation a David solo en `escalate`.

## Próximo paso sugerido

Crear `vault/03-decisions/2026-05-26-adopcion-verifier-gate.md` si decides aplicar #2. Yo no he tocado nada — este brief es solo input para tu decisión.
