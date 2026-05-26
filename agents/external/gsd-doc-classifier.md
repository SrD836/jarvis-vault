---
title: "gsd-doc-classifier (external agent)"
type: agent-mirror
external: true
date: 2026-05-26T03:15:01+00:00
source_path: ~/.claude/agents/gsd-doc-classifier.md
related:
  - "[[00-MOC]]"
  - "[[agents/external/gsd-doc-synthesizer]]"
  - "[[agents/external/gsd-doc-verifier]]"
  - "[[agents/external/gsd-doc-writer]]"
  - "[[agents/external/gsd-pattern-mapper]]"
  - "[[agents/external/index]]"
  - "[[skills/external/gsd-ingest-docs]]"
tags: [agent, external, mirror]
# auto-linked 2026-05-26
---


# `gsd-doc-classifier` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-doc-classifier.md`  
**Mirrored at**: 2026-05-26T03:15:01+00:00

## Descripción

Classifies a single planning document as ADR, PRD, SPEC, DOC, or UNKNOWN. Extracts title, scope summary, and cross-references. Spawned in parallel by /gsd-ingest-docs. Writes a JSON classification file and returns a one-line confirmation.

## Human notes

_(no se sobrescribe por automatización)_
