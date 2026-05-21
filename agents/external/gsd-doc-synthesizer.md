---
title: "gsd-doc-synthesizer (external agent)"
type: agent-mirror
external: true
date: 2026-05-21T03:15:01+00:00
source_path: ~/.claude/agents/gsd-doc-synthesizer.md
related:
  - "[[00-MOC]]"
  - "[[agents/external/gsd-doc-verifier]]"
  - "[[agents/external/gsd-doc-writer]]"
  - "[[agents/external/gsd-plan-checker]]"
  - "[[agents/external/gsd-research-synthesizer]]"
  - "[[agents/external/index]]"
  - "[[skills/external/gsd-ingest-docs]]"
tags: [agent, external, mirror]
# auto-linked 2026-05-21
---


# `gsd-doc-synthesizer` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-doc-synthesizer.md`  
**Mirrored at**: 2026-05-21T03:15:01+00:00

## Descripción

Synthesizes classified planning docs into a single consolidated context. Applies precedence rules, detects cross-ref cycles, enforces LOCKED-vs-LOCKED hard-blocks, and writes INGEST-CONFLICTS.md with three buckets (auto-resolved, competing-variants, unresolved-blockers). Spawned by /gsd-ingest-docs.

## Human notes

_(no se sobrescribe por automatización)_
