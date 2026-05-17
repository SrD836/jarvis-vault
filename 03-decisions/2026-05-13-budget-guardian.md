---
title: "Decisión: Descartar 'budget guardian' manual en favor de Claude Max y Failover"
type: decision
date: 2026-05-13
decision: Utilizar Claude Max ($100/mes) en vez de API de Anthropic y descartar el budget guardian.
alternatives:
  - "Mantener API key de Anthropic y crear un script 'budget guardian' para monitorear tokens"
  - "Utilizar la suscripción de Claude Max y depender de un Failover Chain hacia Ollama local al agotar la cuota"
outcome: positive
outcome_observed_after_days: 1
tags: [decision, budget, infra, claude]
driven_by_sessions: []
involves_agents: ["[[agents/main]]", "[[agents/planner]]"]
related:
  - "[[../seeds/decisiones-clave]]"
---

## Contexto

En el Plan v2 se había propuesto la implementación de un "budget guardian" para evitar un gasto excesivo en la API de Anthropic, llevando el control preciso de los tokens consumidos por los agentes. Sin embargo, en el rediseño de arquitectura surgió la posibilidad de utilizar la suscripción "Claude Max" ($100/mes) mediante `claude-cli`, lo cual proporciona una cuota generosa por un coste fijo mensual.

## Alternativas

### Opción A: API key + Budget Guardian
- Pros: Control granular, se paga exactamente por lo consumido. Posibilidad de usar varios modelos libremente.
- Cons: Claude Max no expone un billing transparente por tokens a través de la CLI de la misma manera que la API. Requiere construir un componente complejo (`budget guardian`) que podría fallar o desincronizarse. 

### Opción B: Claude Max + Failover a Ollama
- Pros: Gasto mensual topado ($100). No hay riesgo de sorpresas a fin de mes. Simplifica la arquitectura al no necesitar el componente del guardian.
- Cons: Estamos atados a los límites de rate (HTTP 429) cuando se alcanza la cuota intensiva, en lugar de un límite de dinero.

## Criterios de evaluación

1. **Riesgo financiero:** Evitar sustos en la tarjeta de crédito a fin de mes.
2. **Complejidad del sistema:** Menor número de piezas móviles es preferible para mantener el sistema autónomo (Fase de Estabilización).

## Decisión

**Elegimos la Opción B (Claude Max + Failover)** porque el riesgo financiero desaparece y se elimina la necesidad de mantener código personalizado para un *budget guardian*. Para mitigar los "límites de rate" (HTTP 429), se configuró una *failover chain* en OpenClaw para que, si Anthropic bloquea por cuota, el orquestador derive las consultas al modelo local `ollama-local/qwen2.5-coder:7b-instruct`.

## Outcome observado

**positive** (tras 1 día). El costo está limitado. La configuración de failover ya está implementada y a la espera de validación real (tarea F1.5).
