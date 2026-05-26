---
caller: main
target: planner
ts: 2026-05-26T20:22:58Z
session_id: 170fe95e-456e-4151-aa86-21b6177c69fb
model_used: claude-opus-4-7
latency_ms: 20468
spawned_children: []
source: delegate-mcp
---

## Task

Brief corto (máx 350 palabras): 3 patrones 2026 de feedback loops para agentes LLM aplicables a JARVIS.

JARVIS = OpenClaw multi-agente. Componentes clave: delegate MCP (main→planner/researcher/handlers), memoria persistente (user/feedback/project/reference), hooks en settings.json, skills cargables, TaskFlow para jobs durables, wiki Obsidian. Modelos Claude 4.X.

Entrega:
1. Resumen ejecutivo (3 líneas)
2. Para cada patrón: Nombre | Qué es (2 líneas) | Cómo aplicarlo a JARVIS (componente concreto: memoria/hooks/delegate/skills/TaskFlow) | Riesgo/coste
3. Señales 2026 si las conoces (Reflexion, LangGraph checkpoints, DSPy GEPA, Mem0, Letta, etc.)

Sé concreto. No generalidades.

## Response

## Feedback loops 2026 → JARVIS

**Resumen.** Tres lazos cubren el espectro: crítico antes de emitir (intra-run), lección persistida tras fallo (inter-run), y steering en hitos (durante delegación). Mayor ROI inmediato: memoria episódica reusando el sistema `feedback` existente. Complejidad creciente en ese orden.

---

**1. Verifier-Critic gating** *(Reflexion / Self-Refine / DSPy Refine)*
*Qué es.* Un crítico barato evalúa el output contra un rubric explícito antes de devolverlo; si falla, una revisión y stop.
*JARVIS.* Hook `PostToolUse` sobre `mcp__delegate__delegate`: re-delegar el reply al agente `auditor` (modelo Haiku 4.5) con el campo `verificación:` del briefing como rubric. Si `needs-revision` → segundo intento del worker original; si vuelve a fallar → `status: partial`. Whitelist `auditor` en `openclaw.json` para todos los workers.
*Coste.* +1 llamada Haiku por delegación (~30% tokens), +5-10s latencia. Riesgo: sycophancy si crítico = modelo del worker; usar tier distinto.

**2. Episodic lesson memory** *(Voyager / Letta / Mem0)*
*Qué es.* Tras cada run fallido o parcial, extraer **una** lección {trigger, antipatrón, regla} y persistirla como política, no como hecho.
*JARVIS.* Hook `SubagentStop`: si reply marca `error|partial`, invocar skill que escriba `memory/feedback_<slug>.md` con frontmatter `type: feedback` y línea en `MEMORY.md`. Se inyecta automáticamente en futuras conversaciones del agente. Cap ≤30 lessons/agente, podar trimestral con `curator`.
*Coste.* ~1k tokens Haiku por fallo. Riesgo real: memory bloat y reglas contradictorias acumuladas.

**3. Mid-run steering vía TaskFlow** *(Process Reward / LangGraph checkpoints)*
*Qué es.* Workers largos emiten checkpoints `{turn, intent, action}`; el orquestador puede redirigir antes del final en vez de fire-and-forget.
*JARVIS.* Modelar delegaciones >5min como TaskFlow job con child tasks; worker append a checkpoint file; `main` poll-ea con `mcp__openclaw__session_status` y corrige con `mcp__openclaw__sessions_send`. Regla dura: máx 1 steer por delegación.
*Coste.* Complejidad media-alta, infra ya existe. Riesgo: "supervisor anxiety" → interrupciones excesivas; sólo activar en jobs largos.

---

**Señales 2026.** Shinn *Reflexion* (NeurIPS 23) · Madaan *Self-Refine* · Wang *Voyager* · Lightman *Let's Verify Step by Step* (OpenAI) · Letta (ex MemGPT) + Mem0 docs 2025 · LangGraph `supervisor` + checkpointers · DSPy `Refine`/`BestOfN` y línea GEPA · Anthropic *Building effective agents* (2024) y serie orchestrator-worker 2025.

**Empezar por Patrón 2** — reusa memoria existente, sin infra nueva, una semana de trabajo.
