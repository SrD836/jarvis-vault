---
title: "skill-dispatcher (external agent)"
type: agent-mirror
external: true
date: 2026-05-28T03:15:01+00:00
source_path: ~/.claude/agents/skill-dispatcher.md
related:
  - "[[00-MOC]]"
  - "[[agents/external/index]]"
  - "[[agents/main]]"
  - "[[agents/skill-dispatcher]]"
  - "[[skills/external/dispatching-parallel-agents]]"
  - "[[skills/index]]"
tags: [agent, external, mirror]
# auto-linked 2026-05-28
---


# `skill-dispatcher` (external Claude Code agent)

**Source**: `~/.claude/agents/skill-dispatcher.md`  
**Mirrored at**: 2026-05-28T03:15:01+00:00

## Descripción

Translate natural language requests in Spanish into invocations of the correct skill or plugin from ~/.claude/skills/ and the registered plugin catalog. Use when the user describes WHAT they want to do without naming the skill/command. Returns either an invoked Skill call or a short list of candidates if ambiguous.

## Human notes

_(no se sobrescribe por automatización)_
