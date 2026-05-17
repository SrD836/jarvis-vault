---
title: Memory Wiki Agent Guide
created: 2026-05-14
type: guide
tags: [wiki, agent, openclaw]
related:
  - "[[index]]"
  - "[[WIKI]]"
  - "[[../../agents/index]]"
  - "[[../../00-MOC]]"
---

# Memory Wiki Agent Guide

- Treat generated blocks as plugin-owned.
- Preserve human notes outside managed markers.
- Prefer source-backed claims over wiki-to-wiki citation loops.
- Prefer structured `claims` with evidence over burying key beliefs only in prose.
- Use `.openclaw-wiki/cache/agent-digest.json` and `claims.jsonl` for machine reads; markdown pages are the human view.
