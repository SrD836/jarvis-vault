---
title: "gsd-eval-auditor (external agent)"
type: agent-mirror
external: true
date: 2026-05-23T03:15:01+00:00
source_path: ~/.claude/agents/gsd-eval-auditor.md
tags: [agent, external, mirror]
related:
  - "[[agents/external/index]]"
  - "[[00-MOC]]"
---

# `gsd-eval-auditor` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-eval-auditor.md`  
**Mirrored at**: 2026-05-23T03:15:01+00:00

## Descripción

Retroactive audit of an implemented AI phase's evaluation coverage. Checks implementation against the AI-SPEC.md evaluation plan. Scores each eval dimension as COVERED/PARTIAL/MISSING. Produces a scored EVAL-REVIEW.md with findings, gaps, and remediation guidance. Spawned by /gsd-eval-review orchestrator.

## Human notes

_(no se sobrescribe por automatización)_
