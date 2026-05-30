---
title: "🤖 Agentes JARVIS"
type: moc
updated: 2026-05-30T10:30:01
tags: [moc, agent, jarvis]
related:
  - "[[../00-MOC]]"
---

# 🤖 Agentes JARVIS

> Mapa del equipo multi-agente. Actualizado automáticamente por `hermes/orgchart.py` cada 30 min.

## Tabla

| Agente | Rol | Modelo | allowAgents | Children runtime |
|---|---|---|---|---|
| [[main]] | main | `anthropic/claude-opus-4-8` | planner, researcher, archivist, polymarket-handler, skill-dispatcher, job-hunter | — |
| [[planner]] | orchestrator | `anthropic/claude-opus-4-8` | researcher, code-reviewer, debugger, tester, documenter, designer, apier, archivist, auditor, monitor, skill-reviewer, job-hunter, polymarket-handler, skill-dispatcher | — |
| [[code-reviewer]] | worker | `anthropic/claude-opus-4-8` | — | — |
| [[researcher]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[documenter]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[apier]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[skill-reviewer]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[debugger]] | worker | `anthropic/claude-opus-4-8` | — | — |
| [[tester]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[auditor]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[archivist]] | worker | `anthropic/claude-haiku-4-5` | — | — |
| [[monitor]] | worker | `anthropic/claude-haiku-4-5` | — | — |
| [[designer]] | worker | `anthropic/claude-sonnet-4-6` | — | — |
| [[job-hunter]] | orchestrator | `anthropic/claude-sonnet-4-6` | researcher, documenter, archivist | — |
| [[polymarket-handler]] | orchestrator | `anthropic/claude-opus-4-8` | researcher | — |
| [[skill-dispatcher]] | worker | `anthropic/claude-sonnet-4-6` | — | — |

## Org chart (Mermaid)

```mermaid
graph TD
  main -.allowed.-> planner
  main -.allowed.-> researcher
  main -.allowed.-> archivist
  main -.allowed.-> polymarket-handler
  main -.allowed.-> skill-dispatcher
  main -.allowed.-> job-hunter
  planner -.allowed.-> researcher
  planner -.allowed.-> code-reviewer
  planner -.allowed.-> debugger
  planner -.allowed.-> tester
  planner -.allowed.-> documenter
  planner -.allowed.-> designer
  planner -.allowed.-> apier
  planner -.allowed.-> archivist
  planner -.allowed.-> auditor
  planner -.allowed.-> monitor
  planner -.allowed.-> skill-reviewer
  planner -.allowed.-> job-hunter
  planner -.allowed.-> polymarket-handler
  planner -.allowed.-> skill-dispatcher
  job-hunter -.allowed.-> researcher
  job-hunter -.allowed.-> documenter
  job-hunter -.allowed.-> archivist
  polymarket-handler -.allowed.-> researcher
```

Leyenda:
- `..>` línea punteada: target permitido en config (`allowAgents`).
- `==>` línea sólida: spawn real registrado en subagent-registry.
