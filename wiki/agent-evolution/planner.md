---
title: "Evolution proposal — planner"
type: agent-evolution
date: 2026-05-20T04:00:18+00:00
agent: "[[agents/planner]]"
pattern: iter-cap-saturated
confidence: medium
proposed: true
applied: false
tags: [agent-evolution, planner, iter-cap-saturated]
related:
  - "[[agents/planner]]"
  - "[[00-MOC]]"
---

# Evolución propuesta — `planner`

_Generado por `hermes/learnings.py` el 2026-05-20T04:00:18+00:00. Ventana: 7 días._

## Patrón: `iter-cap-saturated`

- Confianza: **medium**
- Ocurrencias: **4** runs
- `cap_value`: `8`

### Evidencia

- [[agents/planner/runs/2026-05-19/095210-from-main]]
- [[agents/planner/runs/2026-05-19/105919-from-main]]
- [[agents/planner/runs/2026-05-19/105317-from-main]]
- [[agents/planner/runs/2026-05-19/104519-from-main]]

### Propuesta

## Cambios propuestos

### Cambio 1: Limitar `allow_agents` para reducir carga cognitiva

**Archivo:** `openclaw.json` (config del agente planner)

**Snippet:**
```json
{
  "agent_id": "planner",
  "allow_agents": [
    "code-reviewer",
    "researcher",
    "documenter",
    "debugger",
    "tester",
    "auditor"
  ]
}
```

**Razón:** De 11 agentes permitidos → 6. Los 5 eliminados (apier, skill-reviewer, archivist, monitor, job-hunter) son casos de borde que fuerzan al planner a leer briefings extra y tomar decisiones de delegación innecesarias, saturando el límite de 8 iteraciones.

---

### Cambio 2: Añadir límite de herramientas por run en briefing

**Archivo:** `briefing.md` del planner (sección "Política de delegación")

**Snippet a añadir:**
```markdown
## Límites operativos

- **Máximo de tool_calls por run:** 15
- **Si alcanzas iter 6 sin respuesta final → fuerza delegación inmediata** al worker más afín con `delegate(task_summary, target_agent)`. No intentes resolver tú mismo.
- **Prohibido leer briefings de workers** durante la ejecución. Usa solo `list_agents` para confirmar existencia.
```

**Razón:** El patrón muestra runs con 16-27 tool_calls y lecturas repetidas de briefings de workers (list_agents + get_agent_briefing). Esto quema iteraciones en overhead en lugar de delegar.

---

### Cambio 3: Reducir `max_iterations` del agente a 6

**Archivo:** `openclaw.json` (config del planner)

**Snippet:**
```json
{
  "agent_id": "planner",
  "max_iterations": 6
}
```

**Razón:** El fallo ocurre en iter 8 (cap). Bajando a 6, el agente aborta antes y fuerza delegación temprana. El run `095210` murió exactamente en iter 9 con "cap iterations (8) sin respuesta fin". Con límite 6, abortaría en iter 6 y se podría reintentar con delegación forzada.

## Patrón: `tokens-budget-tight`

- Confianza: **medium**
- Ocurrencias: **4** runs
- `threshold`: `133333`

### Evidencia

- [[agents/planner/runs/2026-05-19/121033-from-main]]
- [[agents/planner/runs/2026-05-19/105919-from-main]]
- [[agents/planner/runs/2026-05-19/105317-from-main]]
- [[agents/planner/runs/2026-05-19/104519-from-main]]

### Propuesta

## Cambios propuestos

### 1. Limitar `allow_agents` a los estrictamente necesarios

**Archivo:** `openclaw.json` (config del agente planner)

**Cambio:**
```json
"allow_agents": ["researcher", "code-reviewer", "debugger", "tester", "auditor", "documenter"]
```

**Razón:** Los 5 agentes eliminados (`apier`, `skill-reviewer`, `archivist`, `monitor`, `job-hunter`) rara vez son delegados realmente, pero su presencia en el contexto infla el briefing del planner con descripciones innecesarias, contribuyendo al exceso de tokens de entrada.

---

### 2. Añadir límite de tokens de entrada en el briefing

**Archivo:** `briefing.md` del agente planner (sección de políticas)

**Cambio:**
```markdown
## Límite de contexto

- **Max input tokens por run:** 120,000
- Si al cargar el briefing + sesiones recientes + estado del workspace superas 120K tokens de entrada, **debes delegar inmediatamente** al agente más especializado sin procesar más contexto.
- No intentes leer archivos grandes (>50KB) tú mismo; usa `delegate` a researcher o code-reviewer.
```

**Razón:** Los runs fallidos muestran inputs de 133K-299K tokens. Un límite explícito fuerza delegación temprana antes de saturar el contexto.

---

### 3. Añadir regla de aborto temprano por tokens

**Archivo:** `caps.env` o configuración de capacidades del planner

**Cambio:**
```env
MAX_INPUT_TOKENS_PER_RUN=120000
ABORT_ON_INPUT_TOKENS_EXCEEDED=true
```

**Razón:** Corta automáticamente runs que ya superan el umbral antes de que consuman iteraciones y fallen por timeout o loops.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
