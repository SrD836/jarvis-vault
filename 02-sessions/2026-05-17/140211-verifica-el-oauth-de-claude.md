---
title: "Verifica El Oauth De Claude"
date: 2026-05-17
session_id: 4abb297c-40a5-438b-bfdc-3158faae97ae
source: trajectory-recovery
trigger: cron
related:
  - "[[02-sessions/2026-05-17/105131-heartbeat-poll]]"
  - "[[03-decisions/2026-05-14-failover-chain-check]]"
  - "[[agents/index]]"
  - "[[agents/monitor]]"
  - "[[skills/external/gsd-health]]"
tags: [session, recovered]
# auto-linked 2026-05-18
---


**User:** [cron:e4e7d923-5ad4-4f87-9082-7d78062888d0 oauth-health] Verifica el OAuth de Claude Code. Ejecuta en bash: docker exec openclaw-fork-openclaw-gateway-1 claude auth status (si tu sandbox ya esta dentro del container, solo: claude auth status). Parsea el JSON. Si loggedIn=true y subscriptionType=max, responde EXACTAMENTE 'HEARTBEAT_OK' (sin comillas, sin nada mas) para que no se entregue mensaje. Si loggedIn=false o el comando falla, manda este mensaje literal a David: 'ALERTA OAUTH JARVIS — Claude Code OAuth caido. Re-login: docker exec -it openclaw-fork-openclaw-gateway-1 claude /login' y termina.
Current time: Sunday, May 17th, 2026 - 4:02 PM (Europe/Madrid)
Reference UTC: 2026-05-17 14:02 UTC

Return your response as plain text; it will be delivered automatically. If the task explicitly calls for messaging a specific external recipient, note who/where it should go instead of sending it yourself.

*[Sesión terminada sin respuesta del modelo]*