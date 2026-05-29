---
title: "gsd-doc-classifier (external agent)"
type: agent-mirror
external: true
date: 2026-05-29T03:15:01+00:00
source_path: ~/.claude/agents/gsd-doc-classifier.md
tags: [agent, external, mirror]
related:
  - "[[agents/external/index]]"
  - "[[00-MOC]]"
---

# `gsd-doc-classifier` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-doc-classifier.md`  
**Mirrored at**: 2026-05-29T03:15:01+00:00

## Descripción

Classifies a single planning document as ADR, PRD, SPEC, DOC, or UNKNOWN. Extracts title, scope summary, and cross-references. Spawned in parallel by /gsd-ingest-docs. Writes a JSON classification file and returns a one-line confirmation.

## Human notes

_(no se sobrescribe por automatización)_
