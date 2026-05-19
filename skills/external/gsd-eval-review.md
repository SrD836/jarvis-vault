---
title: "gsd-eval-review (external skill)"
type: skill-mirror
external: true
category: general
date: 2026-05-19T03:15:01+00:00
source_path: ~/.claude/skills/gsd-eval-review
source_file: SKILL.md
related:
  - "[[00-MOC]]"
  - "[[agents/external/gsd-eval-auditor]]"
  - "[[agents/external/gsd-eval-planner]]"
  - "[[skills/external/gsd-audit-milestone]]"
  - "[[skills/external/gsd-review]]"
  - "[[skills/external/gsd-verify-work]]"
  - "[[skills/index]]"
tags: [skill, external, general, mirror]
# auto-linked 2026-05-19
---


# `gsd-eval-review` (external skill)

**Categoría**: `general`  
**Source**: `~/.claude/skills/gsd-eval-review/`  
**Mirrored at**: 2026-05-19T03:15:01+00:00

## Descripción

Retroactively audit an executed AI phase's evaluation coverage — scores each eval dimension as COVERED/PARTIAL/MISSING and produces an actionable EVAL-REVIEW.md with remediation plan

## Cómo usar

Skill local del usuario. Para invocarla: `/gsd-eval-review` desde Claude Code, o referenciar el path desde un agent runtime que sepa leerlo.

## Human notes

_(no se sobrescribe por automatización)_
