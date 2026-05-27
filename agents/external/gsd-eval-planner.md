---
title: "gsd-eval-planner (external agent)"
type: agent-mirror
external: true
date: 2026-05-27T03:15:01+00:00
source_path: ~/.claude/agents/gsd-eval-planner.md
related:
  - "[[00-MOC]]"
  - "[[agents/external/gsd-ai-integration-phase]]"
  - "[[agents/external/gsd-eval-auditor]]"
  - "[[agents/external/gsd-framework-selector]]"
  - "[[agents/external/gsd-plan-checker]]"
  - "[[agents/external/index]]"
  - "[[agents/tester]]"
tags: [agent, external, mirror]
# auto-linked 2026-05-27
---


# `gsd-eval-planner` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-eval-planner.md`  
**Mirrored at**: 2026-05-27T03:15:01+00:00

## Descripción

Designs a structured evaluation strategy for an AI phase. Identifies critical failure modes, selects eval dimensions with rubrics, recommends tooling, and specifies the reference dataset. Writes the Evaluation Strategy, Guardrails, and Production Monitoring sections of AI-SPEC.md. Spawned by /gsd-ai-integration-phase orchestrator.

## Human notes

_(no se sobrescribe por automatización)_
