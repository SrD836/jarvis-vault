---
title: "Evolution proposal — researcher"
type: agent-evolution
date: 2026-05-23T04:00:34+00:00
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

_Generado por `hermes/learnings.py` el 2026-05-23T04:00:34+00:00. Ventana: 7 días._

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

**Archivo:** `agents/researcher.md` (briefing del agente)

**Snippet a añadir** (justo antes de "Política de delegación"):

```markdown
## Límites operativos

- **Máximo de iteraciones por tarea:** 6
- Si tras 6 iteraciones no hay respuesta final → abortar y reportar `"Límite de iteraciones alcanzado. Pendiente: [resumen de lo investigado]"`.
- No iniciar nuevas búsquedas si ya se han usado 4 iteraciones; priorizar sintetizar lo encontrado.
```

**Razón:** El agente llega a 8 iteraciones (cap) porque no tiene un límite interno. Forzar aborto temprano con reporte evita saturación y da visibilidad.

---

### Cambio 2: Reducir `max_iterations` en configuración del agente

**Archivo:** `openclaw.json` (config global de agentes)

**Diff:**

```json
{
  "agents": {
    "researcher": {
      // ... resto de props
      "max_iterations": 6,
      "max_tokens_per_turn": 300000
    }
  }
}
```

**Razón:** Bajar de 8 a 6 iteraciones máximas evita que el agente tope el cap de iteraciones (8) y fuerza a sintetizar antes. El aborto por tokens (500k) también se mitiga indirectamente.

---

### Cambio 3: Añadir directiva de síntesis temprana en el briefing

**Archivo:** `agents/researcher.md`

**Snippet a añadir** al final de "Human notes":

```markdown
### Síntesis obligatoria

- A partir de la iteración 4, **prohíbo** nuevas tool_calls de búsqueda/lectura.
- En iteración 4-6: solo sintetizar lo recolectado y generar respuesta final.
- Si la tarea requiere más datos de los obtenidos en 4 iteraciones → incluir en el reporte: `"Datos insuficientes para conclusión definitiva. Áreas no cubiertas: [lista]"`.
```

**Razón:** Los runs fallidos muestran 11-37 tool_calls. Cortar búsquedas en iter 4 fuerza a decidir con lo que hay, evitando bucles de recolección.

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

### Cambio 1: Limitar contexto inicial del researcher vía `max_input_tokens` en caps env

**Archivo:** `openclaw.json` (sección `caps.env` del agente researcher)

**Snippet:**
```json
{
  "agent_id": "researcher",
  "caps": {
    "env": {
      "MAX_INPUT_TOKENS_PER_TURN": 150000,
      "MAX_TOKENS_PER_USER_TURN": 300000,
      "MAX_ITERATIONS": 12
    }
  }
}
```

**Razón:** El run 101232 consumió 434k input tokens (87% del límite global de 500k). Forzar un tope de 150k por turno obliga al agente a delegar o resumir antes de saturar el contexto.

### Cambio 2: Añadir directiva de chunking y delegación en briefing

**Archivo:** `briefings/researcher.md` (sección "Directiva de búsqueda")

**Diff:**
```diff
- **Directiva de búsqueda:** Basa tus investigaciones en documentación oficial y fuentes recientes. Verifica meticulosamente la información antes de integrarla o sugerirla para el vault.
+ **Directiva de búsqueda:** 
+ - Si una tarea requiere >5 archivos o >100k tokens de contexto, divide en subtareas y delega a ti mismo (spawn child researcher) con `max_iterations: 6`.
+ - Antes de cada tool_call, verifica tokens acumulados. Si input_tokens > 120k, fuerza un resumen intermedio y reinicia contexto.
+ - Basa tus investigaciones en documentación oficial y fuentes recientes. Verifica meticulosamente la información antes de integrarla o sugerirla para el vault.
```

**Razón:** Los 3 fallos comparten tool_calls excesivos (11-37) sin delegación. La directiva fuerza chunking explícito.

### Cambio 3: Configurar `allow_agents` para auto-delegación

**Archivo:** `openclaw.json` (sección `allow_agents` del researcher)

**Snippet:**
```json
{
  "agent_id": "researcher",
  "allow_agents": ["researcher", "planner"]
}
```

**Razón:** Permitir que researcher se spawn a sí mismo permite ejecutar el chunking propuesto en cambio 2 sin depender de planner.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
