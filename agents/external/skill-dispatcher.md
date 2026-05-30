---
title: "skill-dispatcher (external agent)"
type: agent-mirror
external: true
date: 2026-05-30T03:15:01+00:00
source_path: ~/.claude/agents/skill-dispatcher.md
tags: [agent, external, mirror]
related:
  - "[[agents/external/index]]"
  - "[[00-MOC]]"
---

# `skill-dispatcher` (external Claude Code agent)

**Source**: `~/.claude/agents/skill-dispatcher.md`  
**Mirrored at**: 2026-05-30T03:15:01+00:00

## Descripción

Translate natural language requests in Spanish into invocations of the correct skill or plugin from ~/.claude/skills/ and the registered plugin catalog. Use when the user describes WHAT they want to do without naming the skill/command. Returns either an invoked Skill call or a short list of candidates if ambiguous.

## Human notes

_(no se sobrescribe por automatización)_
