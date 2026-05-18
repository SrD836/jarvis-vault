---
title: "gsd-ingest-docs (external skill)"
type: skill-mirror
external: true
category: general
date: 2026-05-18T03:15:01+00:00
source_path: ~/.claude/skills/gsd-ingest-docs
source_file: SKILL.md
related:
  - "[[00-MOC]]"
  - "[[agents/external/gsd-doc-classifier]]"
  - "[[skills/external/gsd-docs-update]]"
  - "[[skills/external/gsd-graphify]]"
  - "[[skills/external/gsd-new-project]]"
  - "[[skills/external/gsd-scan]]"
  - "[[skills/index]]"
tags: [skill, external, general, mirror]
# auto-linked 2026-05-18
---


# `gsd-ingest-docs` (external skill)

**Categoría**: `general`  
**Source**: `~/.claude/skills/gsd-ingest-docs/`  
**Mirrored at**: 2026-05-18T03:15:01+00:00

## Descripción

Scan a repo for mixed ADRs, PRDs, SPECs, and DOCs and bootstrap or merge the full .planning/ setup from them. Classifies each doc in parallel, synthesizes a consolidated context with a conflicts report, and routes to new-project or merge-milestone depending on whether .planning/ already exists.

## Cómo usar

Skill local del usuario. Para invocarla: `/gsd-ingest-docs` desde Claude Code, o referenciar el path desde un agent runtime que sepa leerlo.

## Human notes

_(no se sobrescribe por automatización)_
