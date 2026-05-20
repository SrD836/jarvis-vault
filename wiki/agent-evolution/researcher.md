---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-20T04:00:37+00:00
agent: "[[agents/researcher]]"
pattern: iter-cap-saturated
confidence: medium
proposed: true
applied: false
tags: [agent-evolution, researcher, iter-cap-saturated]
related:
  - "[[agents/researcher]]"
  - "[[00-MOC]]"
---

# Evolución propuesta — `researcher`

_Generado por `hermes/learnings.py` el 2026-05-20T04:00:37+00:00. Ventana: 7 días._

## Patrón: `iter-cap-saturated`

- Confianza: **medium**
- Ocurrencias: **3** runs
- `cap_value`: `8`

### Evidencia

- [[agents/researcher/runs/2026-05-19/104458-from-planner]]
- [[agents/researcher/runs/2026-05-19/095210-from-planner]]
- [[agents/researcher/runs/2026-05-19/101232-from-planner]]

### Propuesta

## Cambios propuestos

### Cambio 1: Limitar iteraciones por tarea en el briefing

**Archivo:** `agents/researcher/briefing.md` (sección de políticas)

**Snippet a añadir (al inicio de la sección de políticas):**

```markdown
## Límites de ejecución

- **MAX_ITERATIONS_PER_TASK:** 6
- Si alcanzas 6 iteraciones sin respuesta final, ejecuta `run_bash echo "ITER_CAP_REACHED: resumen parcial"` y devuelve lo investigado hasta ahora como respuesta final. NO continúes iterando.
- **MAX_TOKENS_INPUT_PER_TASK:** 300000
- Si el contexto de entrada supera 300K tokens, detén la tarea actual y reporta con `run_bash echo "TOKEN_CAP_REACHED: resumen parcial"`.
```

**Razón:** Los 3 fallos muestran iter 9, 12 y 19 con tokens de entrada de 143K, 146K y 434K. Un límite explícito de 6 iteraciones fuerza al agente a cerrar antes de saturar.

---

### Cambio 2: Configurar `max_iterations` en OpenClaw

**Archivo:** `openclaw.json` (sección del agente researcher)

**Diff:**

```json
{
  "agent_id": "researcher",
  "max_iterations": 6,
  "max_input_tokens": 300000
}
```

**Razón:** El sistema aborta por `MAX_TOKENS_PER_USER_TURN` (500K) pero no por iteraciones. Este límite hardcodea el corte antes de que el agente pueda seguir iterando sin control.

---

### Cambio 3: Añadir directiva de "respuesta parcial" en briefing

**Archivo:** `agents/researcher/briefing.md` (al final, antes de los espacios en blanco)

**Snippet:**

```markdown
## Política de respuesta parcial

Si no puedes completar la investigación dentro de los límites (6 iteraciones o 300K tokens de entrada):
1. Ejecuta `run_bash echo "RESUMEN_PARCIAL: <hallazgos clave>"`
2. Devuelve ese resumen como respuesta final al agente que te llamó.
3. No intentes "una iteración más" ni pidas permiso para continuar.
```

**Razón:** Los runs abortados muestran que el agente sigue iterando sin producir respuesta final. Esta directiva le da un protocolo de salida concreto en vez de quedarse en bucle.

## Patrón: `tokens-budget-tight`

- Confianza: **medium**
- Ocurrencias: **3** runs
- `threshold`: `133333`

### Evidencia

- [[agents/researcher/runs/2026-05-19/104458-from-planner]]
- [[agents/researcher/runs/2026-05-19/095210-from-planner]]
- [[agents/researcher/runs/2026-05-19/101232-from-planner]]

### Propuesta

## Cambios propuestos

### Cambio 1: Limitar iteraciones máximas en briefing del agente

**Archivo:** `agents/researcher/briefing.md` (sección de directivas)

**Snippet a añadir al final de las directivas:**

```markdown
### Límite de iteraciones

- Si tras 5 tool_calls no has producido una respuesta final sintetizada, fuerza un resumen y termina.
- No realices más de 3 búsquedas consecutivas sin sintetizar resultados parciales.
- Prioriza `read_file` sobre `grep_search` cuando el archivo sea <50KB.
```

**Razón:** Los runs abortados muestran 21, 11 y 37 tool_calls sin respuesta final. Forzar síntesis temprana evita acumular tokens de contexto.

---

### Cambio 2: Reducir `max_iterations` en configuración del agente

**Archivo:** `openclaw.json` (sección del agente researcher)

**Diff:**

```json
// Antes
"researcher": {
  "model": "anthropic/claude-sonnet-4-6",
  "max_iterations": 20,
  ...
}

// Después
"researcher": {
  "model": "anthropic/claude-sonnet-4-6",
  "max_iterations": 8,
  "max_input_tokens": 150000,
  ...
}
```

**Razón:** El run de 434981 input_tokens excede cualquier límite razonable. 8 iteraciones forzarán al agente a ser más directo.

---

### Cambio 3: Añadir `max_tokens_per_turn` en caps env

**Archivo:** `.env` o `caps.env`

**Snippet:**

```bash
# Límite estricto de tokens por turno para researcher
OPENCLAW_RESEARCHER_MAX_INPUT_TOKENS=150000
OPENCLAW_RESEARCHER_MAX_OUTPUT_TOKENS=8000
```

**Razón:** El aborto por `MAX_TOKENS_PER_USER_TURN (500000)` muestra que el límite global es demasiado alto. Un límite específico de 150K input + 8K output para researcher evita que consuma todo el presupuesto.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
