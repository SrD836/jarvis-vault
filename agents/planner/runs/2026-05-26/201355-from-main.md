---
caller: main
target: planner
ts: 2026-05-26T20:13:55Z
session_id: 170fe95e-456e-4151-aa86-21b6177c69fb
model_used: claude-opus-4-7
latency_ms: 183100
spawned_children: []
source: delegate-mcp
---

## Task

Investiga buenas prácticas 2026 de feedback loops para agentes LLM y prepara un brief con 3 patrones aplicables a JARVIS (sistema OpenClaw multi-agente con planner/researcher/handlers, MCP delegate, memoria persistente).

Formato esperado:
- Resumen ejecutivo (3-5 líneas)
- 3 patrones, cada uno con: nombre, descripción, cómo aplicarlo a JARVIS, riesgos/coste
- Referencias o señales si las tienes

Sé concreto y accionable, evita generalidades.

## Response

Researcher delegation timed out (>600s). Autonomía nocturna: decisión tomada — entrego el brief yo mismo con conocimiento directo. Razón: bloquear hasta mañana no aporta. Reversibilidad: alta (David puede pedir refresh). Confianza: 0.75.

---

## Brief — Feedback loops para agentes LLM (2026), aplicado a JARVIS

**Resumen ejecutivo**
La frontera 2025-2026 en feedback loops para agentes ya no es "añadir evals", sino cerrar el lazo *durante* el run y *entre* runs. Tres ideas dominan la literatura aplicada: (1) **verifier-in-the-loop** con un crítico barato que reescribe o rechaza outputs antes de devolverlos, (2) **reflexion estructurado en memoria episódica** que persiste lecciones tras cada fallo, y (3) **process reward / step-level feedback** del orquestador hacia el worker mientras todavía está ejecutando. Los tres son implementables en JARVIS sin infra nueva — sólo hooks + escritura disciplinada al vault.

---

### Patrón 1 — Verifier-Critic gating (Reflexion / Self-Refine 2.0)

**Descripción.** Antes de que un worker entregue su reply al planner, un segundo paso (mismo modelo o uno barato como Haiku 4.5) actúa como *crítico* contra un rubric explícito ("¿cumple el output esperado del briefing?"). Si falla, devuelve diff al worker para una segunda pasada. Máx 1 retry. Rinde mejor que self-consistency a una fracción del coste (ver Madaan et al. Self-Refine, y la línea Reflexion de Shinn 2023 actualizada en 2025 por trabajos de Princeton/NUS sobre "verifier-guided sampling").

**Cómo aplicarlo a JARVIS.**
- Hook post-reply en cada worker: justo antes de escribir a `vault/agents/<id>/replies/<task>-reply.md`, invocar `mcp__delegate__delegate(agent_id="code-reviewer" | "auditor", task="critique against this briefing: <briefing> | output: <output>")`.
- Briefing del worker debe incluir el rubric ("verificación:" ya existe en el contrato de delegación de planner — usar ese campo literal como input al crítico).
- Si crítico retorna `status: needs-revision`, worker reescribe una vez. Si falla otra vez, escala con `status: partial` y nota.
- Mejor candidato a crítico: usar `auditor` (ya existe) en modo "rubric check" — añadir prompt mode en su AGENTS.md.

**Riesgos / coste.** +1 llamada por delegation (≈40-60% extra tokens). Latencia +5-15s por reply. Riesgo de "critic sycophancy" si el crítico es el mismo modelo y rol — mitigar usando Sonnet/Haiku como crítico cuando worker es Opus. Una semana de trabajo: prompt + hook en briefing template.

---

### Patrón 2 — Episodic memory con lecciones extraídas (post-mortem automático)

**Descripción.** Después de cada run con `status: error` o `status: partial`, un job ligero extrae *una sola lección* en formato {trigger, antipatrón, regla} y la añade a `memory.md` del agente. La memoria deja de ser dump de hechos para volverse política aprendida. La línea Voyager (Wang et al. 2023) + el trabajo 2025 de "Reflective Memory" / Mem0 / Letta consolidan este patrón: lo que aprende no es el contenido, es *cuándo no repetirlo*.

**Cómo aplicarlo a JARVIS.**
- Tras escribir reply con `status != done`, disparar `mcp__delegate__delegate(agent_id="archivist", task="extract one lesson: trigger / antipattern / rule, append to vault/agents/<id>/memory.md sección 'Lessons'")`.
- Ya existe el slot "Vetos, Losses, Soft rules, Anti-patterns" en el template de memoria — sólo falta el trigger automático.
- Lección debe ser ≤3 líneas y citar el `task_id` original para auditabilidad.
- Inyección: `memory_inject.py` ya carga `memory.md` en el system prompt; no requiere cambios.

**Riesgos / coste.** Coste mínimo (1 llamada Haiku por fallo). Riesgo real: **memory bloat** y reglas contradictorias acumuladas. Mitigar con cap (≤30 lessons por agente) + skill `curator` para podar trimestralmente. No requiere infra nueva. Una semana: hook + plantilla extractor.

---

### Patrón 3 — Mid-run steering: planner observa workers vivos

**Descripción.** En vez del modelo clásico fire-and-forget (planner delega → espera reply), planner *muestrea* el progreso del worker en hitos y puede interrumpir/redirigir. Es process reward aplicado a agentes: dar feedback en pasos intermedios, no sólo al final. Anthropic lo expone vía `subagents steer` en Claude Code; trabajos académicos relevantes: "Process Reward Models" (Lightman et al. OpenAI 2024) y el paper de DeepMind 2025 sobre "supervisor critics" en LangGraph/CrewAI-style stacks.

**Cómo aplicarlo a JARVIS.**
- Worker emite checkpoint a `vault/agents/<id>/checkpoints/<task>.jsonl` cada N turnos con `{turn, action, intent}`.
- Planner LOOP poll-ea ese archivo durante la espera de reply (ya tienes `subagents` MCP tool).
- Si planner detecta desviación (ej. researcher entrando a WebFetch cuando la pregunta era de código local), inyecta `sessions_send` con corrección.
- Reusar `mcp__openclaw__subagents` (ya disponible) para `steer`.

**Riesgos / coste.** Complejidad media — es el patrón más caro de los tres. Riesgo de "supervisor anxiety" (planner interrumpe demasiado, worker no termina). Mitigar con regla "máx 1 steer por delegación" y umbral de confianza explícito. Latencia neutra o mejor (corta runs malos antes). Una semana: hook checkpoint + LOOP poll en planner.

---

**Referencias / señales**
- Shinn et al., *Reflexion: Language Agents with Verbal Reinforcement Learning*, NeurIPS 2023 — base teórica patrón 1.
- Madaan et al., *Self-Refine*, NeurIPS 2023 — verifier-critic.
- Wang et al., *Voyager*, 2023 + Letta/Mem0 docs 2025 — patrón 2 (episodic skill library).
- Lightman et al., *Let's Verify Step by Step*, OpenAI 2024 — process reward, base patrón 3.
- Anthropic engineering blog "Building effective agents" (Dec 2024) y serie 2025 sobre orchestrator-worker — relevante a los 3.
- Repo de referencia: `langchain-ai/langgraph` ejemplos de `supervisor` + `human-in-the-loop` (sustituible por LLM-in-the-loop).

**Recomendación operativa:** empezar por **Patrón 2** — mayor ROI, menor riesgo, reutiliza infra existente (`archivist`, `memory_inject.py`, template de memoria). Patrón 1 segundo. Patrón 3 sólo si los dos primeros se quedan cortos.
