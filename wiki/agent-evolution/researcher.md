---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-25T04:00:28+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-25T04:00:28+00:00. Ventana: 7 días._

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

### 1. Limitar `max_iterations` en briefing del agente

**Archivo:** `agents/researcher/briefing.md` (sección de política de delegación)

**Snippet a añadir:**
```markdown
## Límites operativos

- **max_iterations:** 6
- Si alcanzas iter 6 sin respuesta final → ejecuta `run_bash` con `echo "ITER_CAP_REACHED: resumen parcial"` y finaliza con un output estructurado que incluya `status: partial` y `findings: [...]`.
- Prohibido seguir iterando tras el límite.
```

**Razón:** El agente llega a iter 8-19 porque no tiene un toque de queda interno. Fijar 6 iters fuerza un cierre temprano con resumen parcial en vez de abortar por cap global.

### 2. Añadir `max_input_tokens` en configuración del agente

**Archivo:** `openclaw.json` (sección del agente `researcher`)

**Diff:**
```json
{
  "agent_id": "researcher",
  "model_primary": "anthropic/claude-sonnet-4-6",
  "max_input_tokens": 250000,
  "max_iterations": 6
}
```

**Razón:** El run de iter 19 explotó por 434K tokens input. Poner 250K evita que acumule contexto excesivo y fuerza a resumir o delegar antes de saturar.

### 3. Añadir `allow_agents` para delegar subtareas de búsqueda pesada

**Archivo:** `openclaw.json` (sección `allow_agents`)

**Diff:**
```json
{
  "allow_agents": ["researcher-sub"]
}
```

Y crear agente auxiliar `researcher-sub` con briefing mínimo: solo `grep_search`, `read_file`, `run_bash` sin capacidad de escribir runs ni delegar. El researcher principal delega búsquedas específicas a este subagente, manteniendo su propio contexto ligero.

**Razón:** Las búsquedas con 37 tool_calls saturan el contexto. Delegar a un subagente especializado en búsqueda permite al principal mantener contexto pequeño y no exceder límites.

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

### Cambio 1: Limitar contexto de entrada en briefing

**Archivo:** `briefing del agente researcher` (sección de directivas)

**Snippet a añadir al final de las directivas:**

```markdown
### Límite de contexto

- Si el contexto de entrada supera 100k tokens, debes:
  1. Ignorar archivos de sesión y runs anteriores a 48h
  2. No leer archivos >50k tokens salvo que sean estrictamente necesarios
  3. Usar `grep_search` con patrones específicos en vez de `read_file` completo
  4. Priorizar fuentes oficiales sobre documentación interna extensa
```

**Razón:** Los runs fallidos muestran inputs de 143k-434k tokens. Limitar explícitamente el consumo evita que el agente arrastre contexto innecesario.

---

### Cambio 2: Configurar `max_input_tokens` en OpenClaw

**Archivo:** `openclaw.json` (config del agente researcher)

**Diff:**
```json
{
  "agents": {
    "researcher": {
      "model": "anthropic/claude-sonnet-4-6",
      "max_input_tokens": 100000,
      "max_output_tokens": 4000,
      "max_iterations": 15
    }
  }
}
```

**Razón:** El aborto por `MAX_TOKENS_PER_USER_TURN (500000)` y el patrón `tokens-budget-tight` indican que el agente no tiene cota propia. Fijar `max_input_tokens: 100000` fuerza al runtime a truncar antes de que el modelo consuma todo el presupuesto global.

---

### Cambio 3: Añadir directiva de aborto temprano por iteraciones

**Archivo:** `briefing del agente researcher` (sección de directivas)

**Snippet a añadir tras el punto "Run logging obligatorio":**

```markdown
### Control de iteraciones

- Si llevas más de 8 iteraciones sin output final, aborta automáticamente con:
  `ABORT: max_iteraciones sin progreso`
- No uses más de 3 `read_file` consecutivos sin producir un resultado parcial.
- Si el contexto supera 80k tokens y no has generado respuesta, prioriza cerrar con un resumen parcial.
```

**Razón:** El run `095210-from-planner` abortó por "cap iterations (8) sin respuesta final". Dar instrucción explícita de aborto temprano evita consumir tokens en loops improductivos.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
