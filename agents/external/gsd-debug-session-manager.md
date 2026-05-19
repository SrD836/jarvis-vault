---
title: "gsd-debug-session-manager (external agent)"
type: agent-mirror
external: true
date: 2026-05-19T03:15:01+00:00
source_path: ~/.claude/agents/gsd-debug-session-manager.md
tags: [agent, external, mirror]
related:
  - "[[agents/external/index]]"
  - "[[00-MOC]]"
---

# `gsd-debug-session-manager` (external Claude Code agent)

**Source**: `~/.claude/agents/gsd-debug-session-manager.md`  
**Mirrored at**: 2026-05-19T03:15:01+00:00

## Descripción

Manages multi-cycle /gsd-debug checkpoint and continuation loop in isolated context. Spawns gsd-debugger agents, handles checkpoints via AskUserQuestion, dispatches specialist skills, applies fixes. Returns compact summary to main context. Spawned by /gsd-debug command.

## Human notes

_(no se sobrescribe por automatización)_
