---
name: Trabajo
title: Job Hunt & Auto-Apply
description: Búsqueda de empleo, adaptación de CV por oferta y aplicación automática vía LinkedIn Easy Apply
status: active
type: project
created: 2026-05-19
tags: [trabajo, job-hunt, cv-adapter, linkedin]
related:
  - "[[agents/job-hunter]]"
  - "[[agents/job-hunter/user-profile]]"
---

# Trabajo

Hub central de la búsqueda activa de empleo. Este chat es la conversación dedicada con el agente `job-hunter` — separado del chat global de JARVIS para no contaminar contextos.

## Qué se hace aquí

- Refinar filtros de búsqueda (ubicación, salario, modalidad, keywords)
- Validar y completar el `user-profile.json` (datos que el bot necesita para aplicar)
- Revisar adaptaciones de CV por oferta antes de enviar
- Responder preguntas específicas que LinkedIn pida durante el auto-apply (se persisten para no preguntar dos veces)

## Memoria

- Chat aquí: `vault/projects/trabajo/inbox/chat.jsonl` (este archivo)
- Decisiones cristalizadas: `vault/agents/job-hunter/memory.md`
- Datos estructurados del usuario: `vault/agents/job-hunter/user-profile.json`
