---
title: "gsd-code-review-fix (external skill)"
type: skill-mirror
external: true
category: general
date: 2026-05-28T03:15:01+00:00
source_path: ~/.claude/skills/gsd-code-review-fix
source_file: SKILL.md
related:
  - "[[00-MOC]]"
  - "[[skills/external/gsd-audit-fix]]"
  - "[[skills/external/gsd-code-review]]"
  - "[[skills/external/gsd-review]]"
  - "[[skills/external/receiving-code-review]]"
  - "[[skills/external/requesting-code-review]]"
  - "[[skills/index]]"
tags: [skill, external, general, mirror]
# auto-linked 2026-05-28
---


# `gsd-code-review-fix` (external skill)

**Categoría**: `general`  
**Source**: `~/.claude/skills/gsd-code-review-fix/`  
**Mirrored at**: 2026-05-28T03:15:01+00:00

## Descripción

Auto-fix issues found by code review in REVIEW.md. Spawns fixer agent, commits each fix atomically, produces REVIEW-FIX.md summary.

## Cómo usar

Skill local del usuario. Para invocarla: `/gsd-code-review-fix` desde Claude Code, o referenciar el path desde un agent runtime que sepa leerlo.

## Human notes

_(no se sobrescribe por automatización)_
