---
title: "Tu Trabajo Pasar Overnight Por"
date: 2026-05-17
session_id: bd93bc5b-30cc-4a12-a5f9-b41b1010d1b3
source: trajectory-recovery
trigger: cron
tags: [session, recovered]
---

**User:** [cron:f418979f-53de-4764-9185-193021e7d6e0 overnight-code] Tu trabajo: pasar overnight por los repos del usuario y dejarlos verdes. Eres un senior engineer con un único objetivo: dejar el codebase más limpio que como lo encontraste. (1) Lee /home/agent/agent-stack/vault/Code-Audits/repos.txt para la lista de repos a auditar. Si no existe o está vacío, escribe a David "añade rutas de repos en Code-Audits/repos.txt para activar el cron de código" y termina. (2) Para cada repo en la lista, delegar a tester: ejecuta la suite de tests (npm test, pytest, go test según el lenguaje detectado), lista cada test fallido. (3) Para cada test fallido, distinguir entre: regresión real (toca lógica de aplicación), test flaky (pasa en re-run), env issue, dependency drift. (4) Delegar a code-reviewer para arreglar los flaky/env/dep-drift mecánicamente. NO tocar regresiones reales — flag them. (5) Linter: arregla violaciones de estilo que no cambien comportamiento. (6) Type checker: arregla type errors triviales con anotación obvia. (7) Dependencies: actualiza solo a patch releases no-breaking. Re-corre tests. Revierte cualquier update que rompa. (8) Escribe el audit en /home/agent/agent-stack/vault/Code-Audits/YYYY-MM-DD.md con frontmatter (title, type:code-audit, date, source:cron/overnight-code, repos:[lista], tags:[code, audit, overnight]). Secciones: ## Tests arreglados, ## Lint, ## Deps actualizadas, ## REGRESIONES REALES (para humano), ## Cambios omitidos. (9) Devuelve a David por Telegram solo el resumen: "N repos auditados. X tests arreglados. Y regresiones flaggeadas para revisión." + wikilink al audit. Reglas duras: (a) NUNCA pushear a main. Si haces commit, branch separada y mensaje claro. (b) NUNCA cambiar lógica de aplicación ni refactorizar código que funciona. Solo arreglos mecánicos. (c) NUNCA actualizar dependencies a major version. Solo patch y minor. (d) Si la test suite está fundamentalmente rota de forma no arreglable mecánicamente, no hagas nada, notifica a David.
Current time: Sunday, May 17th, 2026 - 3:14 PM (Europe/Madrid)
Reference UTC: 2026-05-17 13:14 UTC

Return your response as plain text; it will be delivered automatically. If the task explicitly calls for messaging a specific external recipient, note who/where it should go instead of sending it yourself.

**Assistant:** *(Error: 400 {"error":"\"qwen2.5-coder:7b-instruct\" does not support thinking"})*