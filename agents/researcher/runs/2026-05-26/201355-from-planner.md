---
caller: planner
target: researcher
ts: 2026-05-26T20:13:55Z
session_id: fa38362a-f817-4911-9a2b-87c64b8e560d
model_used: claude-sonnet-4-6
latency_ms: 161081
spawned_children: []
source: delegate-mcp
---

## Task

Investiga buenas prácticas 2026 sobre **feedback loops para agentes LLM** y entrega un brief aplicable a JARVIS.

Contexto JARVIS: sistema multi-agente sobre OpenClaw. Topología: `main` orquestador → `planner` descompone → workers (`researcher`, `code-reviewer`, `debugger`, `tester`, `documenter`, `auditor`, `apier`, `designer`, etc.). Delegación vía MCP tool `mcp__delegate__delegate`. Memoria persistente por agente en `vault/agents/<id>/memory.md`. Replies en `vault/agents/<id>/replies/`. Modelo principal Claude Opus 4.7 / Sonnet 4.6.

Objetivo: identificar 3 patrones de feedback loop concretos (no genéricos tipo "añade evals") aplicables hoy a JARVIS, con foco en: (a) self-correction durante un run, (b) cross-agent feedback (worker→planner→main), (c) memoria que mejora con la experiencia.

Formato exacto del output:
- **Resumen ejecutivo** (3-5 líneas)
- **Patrón 1 / 2 / 3** cada uno con:
  - Nombre
  - Descripción (2-3 frases)
  - Cómo aplicarlo a JARVIS (concreto: qué archivo, qué hook, qué worker)
  - Riesgos / coste (tokens, latencia, complejidad)
- **Referencias / señales** (papers 2025-2026, repos, blogs, posts; si no tienes URL exacta cita autor+venue)

Restricciones:
- Concreto y accionable. Evita "considera usar evals" o "añade observabilidad".
- Cada patrón debe poder implementarse en <1 semana por una persona.
- Si un patrón depende de infra que JARVIS no tiene, dilo explícito.

Verificación: el brief debe permitir a David elegir UN patrón y empezar a implementarlo sin más research.

## Response

Tengo suficiente material. Sintetizando el brief.

---

## Brief: Feedback Loops para Agentes LLM — JARVIS 2026

### Resumen ejecutivo

El estado del arte en 2026 converge en tres mecanismos complementarios: **autocorrección acotada dentro del run** (critic-in-the-loop con budget de intentos), **señal de calidad ascendente worker→planner** (structured critic reply), y **destilación episódica diferida** (reflexiones de run que se promueven a memoria long-term). Los tres son implementables sin infra nueva sobre OpenClaw: solo archivos markdown, frontmatter, y prompt engineering. El mayor riesgo transversal es el **garbage-in amplification**: si el mecanismo de corrección o memoria no tiene gate de calidad, refuerza errores en lugar de corregirlos.

---

### Patrón 1 — Bounded Critic-in-the-Loop (self-correction durante un run)

**Descripción**
Cada worker, antes de escribir su reply final, ejecuta un pase de autocrítica contra un rubric fijo (máx. 2 reintentos). El reintento solo se activa si el output falla criterios concretos (schema, longitud, presencia de citas, ausencia de `TODO`). El tercer intento —si falla— escala al planner con flag `needs_review`.

**Cómo aplicarlo a JARVIS**

1. Añadir al `SOUL.md` de cada worker (o a un `EVAL.md` compartido en `vault/agents/shared/`) un **rubric de 4-5 criterios** específicos por rol. Ej. para `researcher`: *(a) ≥1 URL por claim, (b) síntesis presente antes de raw data, (c) campo `confidence` en frontmatter, (d) sin frases aspiracionales sin cita)*.
2. En el AGENTS.md del worker, añadir al **CICLO DE TRABAJO** un paso `4b. AUTOCRÍTICA`: antes de escribir la reply, pasar el draft contra el rubric. Si ≥1 criterio falla → reescribir (intento 2). Si vuelve a fallar → flag en frontmatter `status: needs_review` + campo `failed_criteria: [...]`.
3. El `planner` ya lee las replies; solo necesita reaccionar al campo `needs_review: true` para decidir si reintenta con instrucciones adicionales o escala a `main`.

**Riesgos / coste**
- **Tokens**: +20-40% por run del worker (1 pase extra de evaluación). En `researcher` con 80k cap, manejable.
- **Latencia**: +1 LLM call por run. Asumible si los criterios son binarios (fácil de evaluar).
- **Riesgo**: Si el rubric es vago, el worker pasa siempre. Los criterios deben ser **verificables sin LLM** cuando sea posible (regex, conteo de URLs).
- **No requiere infra nueva**: solo cambios en SOUL.md y AGENTS.md de cada worker.

---

### Patrón 2 — Structured Quality Signal (cross-agent: worker→planner→main)

**Descripción**
Las replies actuales de workers son opacas para el planner en cuanto a calidad: hay un `status: done|partial|error` pero no hay señal sobre *qué tan bien* se ejecutó la tarea ni *qué debería hacerse diferente*. Este patrón añade un bloque `## Retrospectiva` al final de cada reply con 3 campos estructurados: `what_worked`, `what_failed`, `suggest_for_next`. El planner los agrega y los incluye en el briefing del siguiente run paralelo o en el resumen a `main`.

**Cómo aplicarlo a JARVIS**

1. **Ampliar el schema de reply** en `vault/agents/<id>/replies/<task-slug>-reply.md`. Añadir al frontmatter:
   ```yaml
   quality_score: 0-3   # 0=failed, 1=partial, 2=ok, 3=exceeded
   confidence: 0.0-1.0
   ```
   Y al final del body, sección obligatoria:
   ```markdown
   ## Retrospectiva
   - what_worked: <1 frase>
   - what_failed: <1 frase o "nothing">
   - suggest_for_next: <1 frase accionable o "nothing">
   ```
2. **El `planner`**, al integrar replies con `parent_task` común, lee los bloques `Retrospectiva` y los incluye en el **run log** (`vault/agents/planner/runs/…`) como sección `## Señales de calidad`.
3. **El `main`**, al recibir el resumen del planner, tiene visibilidad de qué workers tuvieron friction y puede ajustar el routing futuro (ej. añadir más contexto al briefing de `researcher` en tareas similares).
4. Hermes ya tiene `curator.py` y el vault Obsidian — los `quality_score` acumulados por worker son material directo para el curador semanal.

**Riesgos / coste**
- **Tokens**: ~50-100 tokens adicionales por reply (negligible).
- **Complejidad**: Requiere que el planner implemente lógica de agregación. Si el planner no lee activamente `suggest_for_next`, el patrón es cosmético.
- **Riesgo principal**: **Co-evolving critics problem** (paper arxiv:2601.06794): si el planner evalúa con criterios fijos mientras los workers mejoran, las señales se vuelven stale. Mitigación: revisar el rubric del planner cada 30 runs (puede ser manual).
- **Dependencia**: Planner debe implementar la sección `## Señales de calidad` en sus run logs. Cambio en `vault/agents/planner/AGENTS.md`.

---

### Patrón 3 — Episodic Distillation con Promotion Gate (memoria que mejora con experiencia)

**Descripción**
JARVIS ya tiene `memory/YYYY-MM-DD.md` (logs crudos) y `MEMORY.md` (memoria curada) por agente. El problema: la promoción de episódico a long-term es manual y ad-hoc. Este patrón añade una **promotion gate automática**: al final de cada run exitoso, el worker genera 1-3 "reflexiones candidatas" en formato structured. Esas reflexiones van a un buffer `memory/pending/` con metadata de confianza. Un cron (o el propio worker al inicio del siguiente run) evalúa si la reflexión es válida antes de promoverla a `MEMORY.md`.

**Cómo aplicarlo a JARVIS**

1. **Al cerrar cada run exitoso** (ya está en el CICLO DE TRABAJO de researcher), el worker escribe en `vault/agents/<id>/memory/pending/<date>-<slug>.md`:
   ```yaml
   ---
   reflection: "Cuando la tarea pide comparativa técnica, buscar primero arxiv + GitHub antes que blogs — reduce hallucinations en 2/3 searches"
   evidence_run: vault/agents/researcher/runs/2026-05-26/...
   confidence: 0.7
   type: heuristic|anti-pattern|context-rule
   status: pending   # pending → promoted | rejected
   created: 2026-05-26
   ---
   ```
2. **Gate de promoción** (implementable como paso en el CICLO DE TRABAJO del worker al iniciar el siguiente run):
   - Leer todos los `.md` en `memory/pending/` con `status: pending`.
   - Para cada uno: si `confidence ≥ 0.8` O aparece en ≥2 runs distintos → **promover** a `MEMORY.md` + marcar `status: promoted`.
   - Si `confidence < 0.5` Y older than 14 días → `status: rejected`.
3. **`MEMORY.md`** recibe entradas estructuradas con el campo `evidence_runs: [...]` para trazabilidad. Hermes curator ya puede leer esto para su limpieza semanal.
4. **Reflexiones negativas** (anti-patterns de EVOLUTION.md) siguen el mismo flujo pero con `type: anti-pattern` — el gate es más permisivo (se promueven con confidence ≥ 0.6 porque un error que ya pasó es evidencia directa).

**Riesgos / coste**
- **Tokens**: Mínimos — la generación de reflexiones es un bloque small al final del run (~200 tokens). La gate es lógica de archivo, no LLM.
- **Riesgo #1**: **Garbage-in amplification** — si el worker genera reflexiones incorrectas con alta confianza, se promueven y envenenan la memoria. Mitigación: el campo `evidence_run` es obligatorio; reflexiones sin evidencia específica se rechazan automáticamente.
- **Riesgo #2**: **Forgetting gap** — el sistema no tiene aún lógica de decay. Mitigación: añadir un campo `last_reinforced` y en el gate, reflexiones promovidas hace >90 días sin refuerzo vuelven a `pending`. Hermes curator puede hacer esto.
- **Dependencia de infra**: Ninguna nueva. Solo convención de carpetas + 10 líneas de lógica de gate en el CICLO DE TRABAJO. El cron de Hermes ya existe.
- **Complejidad**: Baja-media. El mayor riesgo es olvidar ejecutar la gate — mitigación: añadirla como **primer paso** del CICLO (ANALIZAR incluye "leer memory/pending/").

---

### Recomendación de orden de implementación

| Prioridad | Patrón | Tiempo est. | Impacto inmediato |
|---|---|---|---|
| 1 | Patrón 3 (Episodic Distillation) | 2-3h | JARVIS deja de "olvidar" lo que aprende entre sesiones |
| 2 | Patrón 2 (Quality Signal) | 3-4h | Planner y main obtienen visibilidad real de friction por worker |
| 3 | Patrón 1 (Bounded Critic) | 4-6h | Outputs de workers mejoran antes de llegar al planner |

El Patrón 3 primero porque modifica solo el worker existente y el CICLO DE TRABAJO ya prevé el punto de escritura — el cambio es mínimo y el beneficio es acumulativo desde el día 1.

---

### Referencias / señales

- [Agent Self-Correction: From Reflexion to Process Reward Models (Zylos Research, mayo 2026)](https://zylos.ai/research/2026-05-12-agent-self-correction-reflexion-to-prm)
- [No More Stale Feedback: Co-Evolving Critics for Open-World Agent Learning (arxiv:2601.06794)](https://arxiv.org/pdf/2601.06794)
- [Memory for Autonomous LLM Agents: Mechanisms, Evaluation, and Emerging Frontiers (arxiv:2603.07670, marzo 2026)](https://arxiv.org/html/2603.07670v1)
- [Memori: A Persistent Memory Layer for Efficient, Context-Aware LLM Agents (arxiv:2603.19935)](https://arxiv.org/pdf/2603.19935)
- [MemRL: Self-Evolving Agents via Runtime Reinforcement Learning on Episodic Memory (arxiv:2601.03192)](https://arxiv.org/html/2601.03192v1)
- [Better Ways to Build Self-Improving AI Agents — Yohei Nakajima (blog, 2025)](https://yoheinakajima.com/better-ways-to-build-self-improving-ai-agents/)
- [Reflexion: Language Agents with Verbal Reinforcement Learning — Shinn et al. (NeurIPS 2023, base conceptual vigente)](https://arxiv.org/abs/2303.11366)
- [Externalization in LLM Agents: Memory, Skills, Protocols and Harness Engineering (arxiv:2604.08224, abril 2026)](https://arxiv.org/html/2604.08224v1)
