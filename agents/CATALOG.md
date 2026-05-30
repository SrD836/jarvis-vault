---
title: "Agent CATALOG — fuente unica de routing del Planner"
type: agent-catalog
date: 2026-05-29
generated_from: ["role-bindings.yaml", "openclaw.json", "AUDIT.md s1"]
related:
  - "[[00-MOC]]"
  - "[[agents/index]]"
  - "[[agents/main]]"
  - "[[agents/planner]]"
  - "[[agents/runs-today]]"
  - "[[agents/skill-dispatcher]]"
tags: [agents, catalog, routing, jarvis]
# auto-linked 2026-05-29
---


# Agent CATALOG

> Fuente UNICA que el Planner (`hermes/planner_route.py`) lee para decidir a quien delegar.
> Cada agente: especialidad en una frase, skills, inputs, outputs, y **model_tier**.
> `model_tier` mapea 1:1 a `hermes/llm_call.py`: `reasoning`=DeepSeek, `mechanical`=Ollama, `frontier`=Claude Max (solo orquestadores / justificado).
> Los **workers SIEMPRE** corren `reasoning` o `mechanical` (0 Claude). `frontier` queda reservado a main/planner.
> Editar el bloque JSON canonico de abajo es lo que cambia el routing; la tabla es la vista humana.

## Tabla (vista humana)

| Agente | Especialidad (1 frase) | Tier | Skills clave | Inputs | Outputs | Estado |
|---|---|---|---|---|---|---|
| **main** | Router Telegram: orquesta, delega agresivo, resuelve solo lo trivial. | frontier | gsd-do, gsd-next | peticion de David | delega + responde | orchestrator |
| **planner** | Recibe del main, descompone en tareas y delega a workers sin ambiguedad. | frontier | writing-plans, brainstorming, gsd-plan-phase | request + CATALOG | plan de routing | orchestrator |
| **researcher** | Investigacion factual y web research con sintesis razonada. | reasoning | gsd-research-phase, gsd-scan | tema/pregunta + contexto | brief con hallazgos | worker |
| **code-reviewer** | Review adversarial de codigo: bugs, security, malas practicas. | reasoning | gsd-code-review, requesting-code-review | diff/archivo | hallazgos por severidad | worker |
| **debugger** | Diagnostica bugs por metodo cientifico (causa raiz, no parche). | reasoning | gsd-debug, gsd-forensics | error/stack/sintoma | hipotesis + fix | worker |
| **tester** | Genera tests y red/green TDD sobre interfaces publicas. | reasoning | test-driven-development, gsd-add-tests | funcion/modulo | casos de test | worker |
| **auditor** | Auditorias retroactivas de seguridad, compliance y UAT. | reasoning | gsd-audit-uat, gsd-secure-phase | archivo/sistema | informe de auditoria | worker |
| **documenter** | Docs tecnicos, API refs y guias de usuario. | reasoning | gsd-docs-update | tema + outputs de otros | markdown de doc | worker |
| **designer** | UI/UX, identidad visual, banners, slides (drafts/spec). | reasoning | design-system, ui-ux-pro-max | brief de UI | spec/copy de diseno | worker |
| **apier** | Diseno de APIs, integraciones MCP y tooling. | reasoning | gsd-graphify | requisitos de API | contrato/diseno | worker |
| **skill-reviewer** | Revisa y cura skills; propone nuevas. | reasoning | writing-skills, gsd-extract_learnings | catalogo de skills | review + propuestas | worker |
| **monitor** | Observabilidad: resume logs, metricas y alertas. | mechanical | gsd-health, gsd-stats | logs/metricas | resumen/clasificacion | worker |
| **archivist** | Longevidad del repo: clasifica backlogs, todos, cleanups. | mechanical | gsd-cleanup, gsd-add-todo | listado/estado del repo | clasificacion/listas | worker |
| **job-hunter** | Automatiza LinkedIn apply + adapta portfolio (pipeline propio). | reasoning | triage-issue, gsd-inbox | oferta/CV | aplicacion/portfolio | worker (cron propio) |
| **polymarket-handler** | (Concepto) bot polymarket scanner/brain/executor. | review | gsd-debug, gsd-forensics | — | — | CANDIDATO_A_RETIRAR (dup de binarios Go */30) |
| **skill-dispatcher** | NL -> skill router: lee SKILL_INDEX e invoca. | mechanical | gsd-help | frase NL | skill a invocar | review (solapa con main/planner routing) |

## Bloque JSON canonico (lo parsea planner_route.py y agent_exec.py)

```json
{
  "main":               {"tier": "frontier",   "role": "orchestrator", "skills": ["gsd-do","gsd-next","gsd-progress"], "specialty": "Router Telegram que orquesta y delega.", "inputs": "peticion de David", "outputs": "delegacion + respuesta"},
  "planner":            {"tier": "frontier",   "role": "orchestrator", "skills": ["writing-plans","brainstorming","gsd-plan-phase","gsd-discuss-phase","gsd-spec-phase","grill-me"], "specialty": "Descompone la solicitud y delega a workers sin ambiguedad.", "inputs": "request + CATALOG", "outputs": "plan de routing"},
  "researcher":         {"tier": "reasoning",  "role": "worker", "skills": ["gsd-research-phase","gsd-scan"], "specialty": "Investigacion factual y web research con sintesis razonada.", "inputs": "tema o pregunta + contexto", "outputs": "brief con hallazgos y fuentes"},
  "code-reviewer":      {"tier": "reasoning",  "role": "worker", "skills": ["gsd-code-review","gsd-code-review-fix","requesting-code-review","receiving-code-review"], "specialty": "Review adversarial de codigo: bugs, security, malas practicas.", "inputs": "diff o archivo de codigo", "outputs": "hallazgos clasificados por severidad"},
  "debugger":           {"tier": "reasoning",  "role": "worker", "skills": ["gsd-debug","gsd-forensics","gsd-spike"], "specialty": "Diagnostica bugs por metodo cientifico, causa raiz no parche.", "inputs": "error, stack trace o sintoma", "outputs": "hipotesis de causa raiz + fix propuesto"},
  "tester":             {"tier": "reasoning",  "role": "worker", "skills": ["test-driven-development","gsd-add-tests","tdd"], "specialty": "Genera tests y red/green TDD sobre interfaces publicas.", "inputs": "funcion o modulo", "outputs": "casos de test"},
  "auditor":            {"tier": "reasoning",  "role": "worker", "skills": ["gsd-audit-uat","gsd-audit-fix","gsd-secure-phase","gsd-validate-phase"], "specialty": "Auditorias retroactivas de seguridad, compliance y UAT.", "inputs": "archivo o sistema", "outputs": "informe de auditoria"},
  "documenter":         {"tier": "reasoning",  "role": "worker", "skills": ["gsd-docs-update"], "specialty": "Docs tecnicos, API refs y guias de usuario.", "inputs": "tema + outputs de otros agentes", "outputs": "markdown de documentacion"},
  "designer":           {"tier": "reasoning",  "role": "worker", "skills": ["design-system","ui-ux-pro-max","banner-design","slides"], "specialty": "UI/UX, identidad visual, banners y slides en draft/spec.", "inputs": "brief de UI", "outputs": "spec o copy de diseno"},
  "apier":              {"tier": "reasoning",  "role": "worker", "skills": ["gsd-graphify"], "specialty": "Diseno de APIs, integraciones MCP y tooling.", "inputs": "requisitos de API", "outputs": "contrato o diseno de API"},
  "skill-reviewer":     {"tier": "reasoning",  "role": "worker", "skills": ["writing-skills","gsd-extract_learnings"], "specialty": "Revisa y cura skills; propone nuevas.", "inputs": "catalogo de skills", "outputs": "review + propuestas"},
  "monitor":            {"tier": "mechanical", "role": "worker", "skills": ["gsd-health","gsd-stats","gsd-intel"], "specialty": "Observabilidad: resume logs, metricas y alertas.", "inputs": "logs o metricas", "outputs": "resumen o clasificacion"},
  "archivist":          {"tier": "mechanical", "role": "worker", "skills": ["gsd-cleanup","gsd-add-todo","gsd-add-backlog","gsd-note"], "specialty": "Longevidad del repo: clasifica backlogs, todos y cleanups.", "inputs": "listado o estado del repo", "outputs": "clasificacion o listas"},
  "job-hunter":         {"tier": "reasoning",  "role": "worker", "skills": ["triage-issue","gsd-inbox"], "specialty": "Automatiza LinkedIn apply y adapta portfolio.", "inputs": "oferta o CV", "outputs": "aplicacion o portfolio", "note": "tiene pipeline cron propio (09:00)"},
  "polymarket-handler": {"tier": "review",     "role": "retire_candidate", "skills": ["gsd-debug","gsd-forensics"], "specialty": "Concepto solapado: el bot polymarket real son binarios Go */30, no este agente.", "inputs": "-", "outputs": "-", "retire_reason": "Duplica los binarios Go autonomos; sin dir de sesion propio."},
  "skill-dispatcher":   {"tier": "mechanical", "role": "review", "skills": ["gsd-help"], "specialty": "NL -> skill router: lee SKILL_INDEX e invoca.", "inputs": "frase en lenguaje natural", "outputs": "skill a invocar", "retire_reason": "Solapa con el routing de main/planner; evaluar si aporta."}
}
```

## Candidatos a retirar (registrado, NO se retira sin OK de David)

- `polymarket-handler` — concepto duplicado: el bot real son binarios Go autonomos (cron */30), no un agente conversacional. Sin dir de sesion propio.
- `skill-dispatcher` — su routing NL->skill solapa con lo que ya hacen main y planner. Evaluar si justifica existir aparte.
- **No-agentes contados por la metrica** (ruido en `learnings.py`, NO son agentes): `runs-today` (reporte del dashboard), `dashboard-dev`, `c1-doctor`, `polymarket-bot` (vistas/legacy, no estan en role-bindings ni en openclaw.json agents.list).

> Decision final sobre retiros: pendiente de David. Este catalogo solo registra el diagnostico.
