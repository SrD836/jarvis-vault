---
title: "Expansión del roster de agentes especializados"
type: decision
date: 2026-05-18
decision: Crear workspaces individuales para 14 agentes especializados bajo /workspace/agents/, cada uno con su propio AGENTS.md
alternatives:
  - "Mantener agente único monolítico (main) que cambia de rol según contexto"
  - "Agentes especializados con workspaces aislados y standing orders propios"
outcome: pending
outcome_observed_after_days: 14
tags: [decision, arquitectura, agentes, multi-agent]
related:
  - "[[agents/main]]"
  - "[[agents/job-hunter]]"
  - "[[agents/archivist]]"
  - "[[agents/planner]]"
---

## Contexto

El sistema JARVIS hasta ahora operaba principalmente con un agente principal (`main`) que asumía distintos roles según la tarea. Para tasks especializadas (búsqueda de empleo, revisión de código, debugging) el contexto de rol se inyectaba en el turno. Esto generaba prompt overhead y pérdida de memoria procedimental entre sesiones del mismo tipo de tarea.

## Alternativas

### Opción A: Agente único monolítico
- Pros: Menor complejidad de routing; un solo punto de configuración
- Cons: Contexto de rol se pierde entre sesiones; AGENTS.md global se hincha con instrucciones de múltiples roles; no hay memoria especializada por dominio

### Opción B: Agentes especializados con workspaces aislados
- Pros: Cada agente mantiene memoria propia (`memory.md`), standing orders enfocados, puede tener herramientas diferentes habilitadas; job-hunter probado en producción el mismo día con éxito (10 ofertas LinkedIn en primer run)
- Cons: Mayor superficie de mantenimiento; cambios en política global (AGENTS.md) deben propagarse a todos los hijos

## Criterios

- Reutilización de contexto entre sesiones del mismo tipo de tarea
- Capacidad de especialización sin contaminar el agente principal
- Evidencia de funcionamiento real antes de comprometer la arquitectura

## Decisión

Se crean 14 agentes especializados: `apier`, `archivist`, `auditor`, `code-reviewer`, `debugger`, `designer`, `documenter`, `job-hunter`, `main`, `monitor`, `planner`, `researcher`, `skill-reviewer`, `tester`. Cada uno con workspace propio bajo `/workspace/agents/<nombre>/` y AGENTS.md con standing orders del rol. job-hunter validado en producción el mismo día (run `2026-05-18T22:12:00` → 10 ofertas LinkedIn).

## Outcome

pending — evaluar tras 2 semanas si la especialización reduce tokens por tarea y mejora calidad de outputs frente al enfoque monolítico.
