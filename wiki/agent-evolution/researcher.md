---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-26T04:00:30+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-26T04:00:30+00:00. Ventana: 7 días._

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

### Cambio 1: Limitar iteraciones máximas en briefing

**Archivo:** `briefing del agente researcher` (sección de políticas)

**Snippet a añadir (al final de la sección "Política de delegación"):**

```markdown
## Límites operativos

- **MAX_ITERATIONS**: 6 (abortar si se supera, reportar "Iteraciones agotadas - tarea incompleta")
- **MAX_TOOLS_PER_TURN**: 15 (si se alcanza, forzar síntesis parcial y devolver resultado)
- **Si se alcanza cualquiera de estos límites**: escribir run file con estado `aborted` y razón clara, NO reintentar automáticamente.
```

**Razón:** El agente llega a iter 8-19 sin entregar respuesta. Un límite duro de 6 fuerza a sintetizar antes.

---

### Cambio 2: Configurar `max_iterations` en OpenClaw

**Archivo:** `openclaw.json` (config del agente researcher)

**Diff:**

```json
{
  "agent_id": "researcher",
  "max_iterations": 6,
  "max_tokens_per_turn": 150000,
  "abort_on_iteration_cap": true
}
```

**Razón:** El fallo actual aborta por cap de sistema (8) o tokens (500k). Reduciendo a 6 iteraciones y 150k tokens se fuerza a ser conciso.

---

### Cambio 3: Añadir directiva de síntesis temprana

**Archivo:** `briefing del agente researcher` (sección "Directiva de búsqueda")

**Snippet a modificar (reemplazar la directiva actual):**

```markdown
**Directiva de búsqueda:** 
1. Lee máximo 3 archivos antes de sintetizar.
2. Si después de 3 tool calls no tienes respuesta clara, devuelve lo que tengas con nota "parcial - requiere más investigación".
3. Prohibido hacer más de 5 búsquedas/greps sin producir un output intermedio.
4. Basa tus investigaciones en documentación oficial y fuentes recientes. Verifica meticulosamente la información antes de integrarla.
```

**Razón:** Los runs muestran 11-37 tool calls sin output. Forzar síntesis cada 3-5 calls evita bucles de búsqueda.

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

### 1. Limitar `max_iterations` en briefing del researcher

**Archivo:** briefing del agente `researcher` (sección de configuración)

**Snippet a añadir al final del briefing:**
```markdown
## Límites de ejecución

- **max_iterations:** 8
- **max_input_tokens_per_turn:** 120000
- Si alcanzas cualquiera de estos límites, aborta inmediatamente con un resumen de lo encontrado hasta el momento. No intentes continuar ni reiniciar.
```

**Razón:** Los runs abortados muestran 9, 19 y 12 iteraciones con input_tokens de 143k, 434k y 146k. Fijar un límite inferior (8 iteraciones, 120k tokens) fuerza al agente a ser más conciso o delegar antes.

---

### 2. Configurar `allow_agents` en openclaw.json para delegación temprana

**Archivo:** `openclaw.json` (config global de agentes)

**Diff:**
```json
{
  "agents": {
    "researcher": {
      "allow_agents": ["planner", "dashboard-dev", "polymarket-bot"],
      "delegation_mode": "auto",
      "max_iterations": 8,
      "max_input_tokens": 120000
    }
  }
}
```

**Razón:** `delegation_mode: suggest` + `allow_agents: []` fuerza al researcher a hacer todo él mismo. Cambiar a `auto` con agentes específicos permite delegar subtareas (ej. consultar dashboard-dev para datos, polymarket-bot para estado del bot) antes de saturar tokens.

---

### 3. Añadir guarda de `max_input_tokens` en el sistema de caps/env

**Archivo:** `caps.env` o configuración de entorno del agente

**Snippet:**
```bash
# Límite por turno para researcher (más restrictivo que el global)
AGENT_RESEARCHER_MAX_INPUT_TOKENS=120000
AGENT_RESEARCHER_MAX_ITERATIONS=8
```

**Razón:** El aborto por `global MAX_TOKENS_PER_USER_TURN (500000)` ocurre porque no hay un límite por agente. Este cap específico para researcher corta antes (120k) y evita llegar al límite global, forzando un aborto limpio con lo investigado hasta el momento.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
