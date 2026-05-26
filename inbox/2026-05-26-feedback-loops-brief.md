---
type: brief
tags: [brief, feedback-loops, jarvis, agentic-patterns]
created: 2026-05-26T22:05:00+02:00
---

# Brief — Feedback loops para agentes LLM (2026) aplicables a JARVIS

## TL;DR
Tres patrones maduros en 2026 que JARVIS puede adoptar sin reescritura arquitectónica:
1. **Reflexion con memoria verbal auditable** — el agente critica su propio intento, guarda la crítica, reintenta.
2. **Process Reward Models (PRM) sobre trayectorias** — verifica cada paso, no solo el output final. Evita el "coherence trap".
3. **Insights→Evals→Traces loop** — pipeline de mejora compuesta: traces de producción alimentan datasets, datasets alimentan evals, evals dirigen cambios.

---

## Patrón 1 — Reflexion con memoria verbal

**Qué es:** tras cada tarea, el agente escribe en lenguaje natural qué falló y por qué; en el siguiente intento, esa crítica entra como contexto adicional. Bajo coste de ingeniería, ganancias incrementales grandes.

**Aplicación a JARVIS:**
- Tras cada ejecución de sub-agente (job-hunter, documenter, researcher), pedir un bloque `## Reflection` en su output: *qué salió bien, qué intentaría distinto*.
- Persistirlo en `vault/04-skills-log/reflections/YYYY-MM-DD-{agent}.md`.
- Antes de invocar al mismo agente para tarea similar, cargar las últimas 3 reflexiones como context inicial.
- Coste: ~200 tokens/turno. ROI alto si el agente se repite (job-hunter sí, one-shots no).

**Riesgo conocido (2026):** "coherence trap" — si generator y evaluator son el mismo modelo con sesgos correlacionados, el agente se auto-convence sin información nueva. Mitigación: usar modelo distinto para la reflexión (Sonnet refleja sobre output de Opus, o viceversa).

---

## Patrón 2 — Process Reward Models sobre trayectorias

**Qué es:** en lugar de evaluar solo el resultado final ("¿encontró ofertas?"), evalúas cada paso intermedio ("¿la query a Indeed fue razonable? ¿el filtrado por salario fue correcto?"). Detecta fallos antes de que se compongan.

**Aplicación a JARVIS:**
- Cada agente loguea su trayectoria en JSON estructurado: `{step, tool, args, observation, reasoning}`.
- Job nocturno (Ollama Qwen local, coste 0) puntúa cada paso 1-5 sobre rubrica fija: *necesario, correcto, eficiente*.
- Trayectorias con paso ≤2 entran a `vault/audit/low-quality-trajectories/` para review manual o auto-rewrite del prompt.
- Encaja con tu audit log JSONL ya existente (`vault/audit/YYYY-MM-DD-actions.jsonl`).

**Por qué importa más que evaluar el output:** en agentes de varios pasos, un output decente puede ocultar pasos ineficientes que queman tokens. PRM lo expone.

---

## Patrón 3 — Insights → Datasets → Evals → Traces (loop compuesto)

**Qué es:** el patrón productivo de 2026. Traces reales → Insights Agent extrae patrones → genera datasets de eval → los evals corren contra cambios futuros → producen nuevos traces. Cada iteración expande cobertura.

**Aplicación a JARVIS:**
- **Insights Agent**: cron semanal (domingo 04:00) que lee `vault/02-sessions/` última semana y extrae: *intents recurrentes, fallos repetidos, herramientas más usadas*. Output → `vault/insights/YYYY-WW.md`.
- **Dataset builder**: los intents recurrentes se convierten en casos de test (`vault/evals/cases/`). Ejemplo: si David pide "ofertas Madrid Go remoto" 3 veces en una semana, eso es un eval case.
- **Eval runner**: antes de cada cambio en un sub-agente (job-hunter prompt, herramienta nueva), correr los eval cases. Si baja la calidad, no se mergea.
- **Pinza con skills**: las skills existentes son la memoria procedimental; este loop genera la memoria evaluativa que actualmente no tienes.

---

## Recomendación de orden

Si solo hay budget para uno este mes, empieza con **Patrón 1 (Reflexion)** — es el de menor coste de implementación y el que más alimenta a los otros dos después. Patrón 3 requiere los datos que Patrón 1+2 generan.

**Antipatrones evitados:**
- No implementar self-correction sin evaluador independiente (coherence trap).
- No alimentar el loop con datos sintéticos al principio — los traces reales de David valen más que 1000 casos generados.
- No optimizar prematuramente la frecuencia: weekly insights bastan; daily es overkill para 1 operador.

## Fuentes
- [Yohei Nakajima — Better Ways to Build Self-Improving AI Agents](https://yoheinakajima.com/better-ways-to-build-self-improving-ai-agents/)
- [Zylos Research — Agent Self-Correction: Reflexion to PRM (May 2026)](https://zylos.ai/research/2026-05-12-agent-self-correction-reflexion-to-prm)
- [SitePoint — Definitive Guide to Agentic Design Patterns 2026](https://www.sitepoint.com/the-definitive-guide-to-agentic-design-patterns-in-2026/)
- [OpenAI Cookbook — Self-Evolving Agents](https://developers.openai.com/cookbook/examples/partners/self_evolving_agents/autonomous_agent_retraining)
- [LangChain — LLM Evaluation Framework: Trajectories vs Outputs](https://www.langchain.com/articles/llm-evaluation-framework)
- [VentureBeat — Designing LLM feedback loops that get smarter over time](https://venturebeat.com/ai/teaching-the-model-designing-llm-feedback-loops-that-get-smarter-over-time)
