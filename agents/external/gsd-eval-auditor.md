---
title: "gsd-eval-auditor (external agent)"
type: agent-mirror
external: true
date: 2026-05-19T03:15:01+00:00
source_path: ~/.claude/agents/gsd-eval-auditor.md
related:
  - "[[00-MOC]]"
  - "[[agents/auditor]]"
  - "[[agents/external/gsd-eval-planner]]"
  - "[[agents/external/gsd-security-auditor]]"
  - "[[agents/external/gsd-verifier]]"
  - "[[agents/external/index]]"
  - "[[skills/external/gsd-eval-review]]"
tags: [agent, external, mirror]
# auto-linked 2026-05-19
---


# `gsd-eval-auditor` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-eval-auditor.md`  
**Mirrored at**: 2026-05-19T03:15:01+00:00

## Descripción

Retroactive audit of an implemented AI phase's evaluation coverage. Checks implementation against the AI-SPEC.md evaluation plan. Scores each eval dimension as COVERED/PARTIAL/MISSING. Produces a scored EVAL-REVIEW.md with findings, gaps, and remediation guidance. Spawned by /gsd-eval-review orchestrator.

## Human notes

_(no se sobrescribe por automatización)_
