---
caller: main
target: planner
ts: 2026-05-26T20:22:25Z
session_id: 170fe95e-456e-4151-aa86-21b6177c69fb
model_used: claude-opus-4-7
latency_ms: 79867
spawned_children: []
source: delegate-mcp
---

## Task

Investiga buenas prácticas 2026 de feedback loops para agentes LLM y prepara un brief con 3 patrones aplicables a JARVIS.

Contexto JARVIS: sistema OpenClaw multi-agente. Arquitectura: agente main delega vía MCP (mcp__delegate__delegate) a especialistas (planner, researcher, polymarket-handler, etc.). Memoria persistente en /home/node/.claude/projects/.../memory/ (tipos: user, feedback, project, reference). Wiki vault Obsidian-friendly. Skills cargables. Modelos Claude 4.X (Opus 4.7, Sonnet 4.6, Haiku 4.5). Hooks configurables en settings.json. TaskFlow para jobs durables. Subagentes con whitelist en openclaw.json.

Formato esperado del brief:
1. Resumen ejecutivo (3-5 líneas)
2. Tres patrones, cada uno con:
   - Nombre
   - Descripción breve (2-3 líneas)
   - Aplicación concreta a JARVIS (qué componente tocar: memoria, hooks, delegate, skills, TaskFlow)
   - Riesgos / coste en tokens o latencia
3. Referencias o señales 2026 si las conoces (papers, frameworks: LangGraph, Letta, DSPy, Mem0, etc.)

Sé concreto y accionable. Evita generalidades tipo "usa evaluación humana". Máximo 400 palabras.

## Response

## Feedback loops para agentes LLM — Brief JARVIS

**Resumen ejecutivo.** En 2026 el consenso aplicado es: cerrar el lazo *dentro* del run (verifier-critic), *entre* runs (memoria episódica con lecciones) y *durante* la delegación (steering en checkpoints). Los tres patrones son implementables en JARVIS reutilizando MCP delegate, el sistema de memory files y hooks de settings.json. ROI más alto: memoria episódica. Mayor complejidad: steering.

---

**1. Verifier-Critic gating (Reflexion / Self-Refine)**
*Descripción.* Antes de emitir output, un crítico barato (Haiku 4.5) evalúa contra rubric explícito; si falla, una sola revisión. Supera a self-consistency a una fracción del coste.
*Aplicación JARVIS.* Hook `PostToolUse` sobre `mcp__delegate__delegate`: capturar reply del especialista y re-delegar a `auditor` con prompt-mode "rubric check" leyendo el campo `verificación:` del briefing. Si `needs-revision`, segundo intento; si vuelve a fallar, devolver con `status: partial`. Whitelist en `openclaw.json` para que cualquier worker pueda invocar `auditor`.
*Riesgos / coste.* +1 llamada Haiku por delegate (~30% tokens extra). Latencia +5-10s. Sycophancy si crítico = mismo rol; mitigar usando modelo distinto al worker.

**2. Episodic memory con lecciones automáticas (Voyager / Letta / Mem0)**
*Descripción.* Tras cada run con error o partial, extraer **una** lección {trigger, antipatrón, regla} y persistirla como memoria tipo `feedback`. La memoria deja de ser hechos y se vuelve política aprendida.
*Aplicación JARVIS.* Hook `Stop` o `SubagentStop`: si el último reply marca fallo, disparar tool/skill que escriba a `/home/node/.claude/projects/.../memory/feedback_<slug>.md` con frontmatter `type: feedback` y entrada en `MEMORY.md`. Reusa el sistema de memory que ya cargas en cada conversación. Cap ≤30 lessons/agente; podar con skill `curator`.
*Riesgos / coste.* Una llamada Haiku por fallo (~1k tokens). Riesgo real: memory bloat y reglas contradictorias. Mitigar con cap + revisión trimestral.

**3. Mid-run steering vía TaskFlow checkpoints (Process Reward)**
*Descripción.* En vez de fire-and-forget, el orquestador observa hitos del worker y puede redirigir antes del final. Equivale a process reward aplicado a delegación.
*Aplicación JARVIS.* Modelar delegaciones largas como TaskFlow job con child tasks; worker emite checkpoints (`turn, action, intent`); `main` poll-ea con `mcp__openclaw__session_status` y, ante desviación, envía corrección con `mcp__openclaw__sessions_send`. Regla dura: máx 1 steer por delegación.
*Riesgos / coste.* Complejidad media-alta. Riesgo "supervisor anxiety" → interrupciones excesivas. Sólo recomendable para delegaciones >5 min.

---

**Referencias / señales 2026.** Shinn et al. *Reflexion* (NeurIPS 23); Madaan et al. *Self-Refine*; Wang et al. *Voyager*; Lightman et al. *Let's Verify Step by Step* (OpenAI); Letta (ex MemGPT) y Mem0 docs 2025; LangGraph `supervisor` + human-in-the-loop; Anthropic engineering *Building effective agents* (2024) y serie orchestrator-worker 2025; DSPy `Refine`/`BestOfN` modules.

**Recomendación:** empezar por **Patrón 2** — mayor ROI, reutiliza memory system existente, sin infra nueva.
