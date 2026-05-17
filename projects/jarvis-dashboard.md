---
title: "jarvis-dashboard — puntero"
type: project-pointer
created: 2026-05-14
moved: 2026-05-16
tags: [project, pointer, dashboard]
related:
  - "[[index]]"
  - "[[../00-MOC]]"
  - "[[../Dashboard]]"
location: "C:\\Users\\david\\agent-stack\\jarvis-dashboard"
---

# 📦 jarvis-dashboard

> Código movido fuera del vault el 2026-05-16. Esta nota mantiene la referencia para que los wikilinks del segundo cerebro no se rompan.

## Ubicación real del código

`C:\Users\david\agent-stack\jarvis-dashboard\` — abrir con file://C:/Users/david/agent-stack/jarvis-dashboard

## Qué es

Nuevo dashboard web (Vite SPA + Node `server.js` + Dockerfile + Traefik dynamic.yml) que reemplaza el OpenClaw Control UI bundled en `jarvss.duckdns.org`. Creado durante la sesión Antigravity del 14-15 May 2026.

## Stack

- Vite 5 (frontend SPA)
- Express en `server.js` (backend HTTP)
- Docker Compose + Traefik
- Artefactos de deploy: `agent-stack/dist/jarvis-dashboard.tar.gz` (26 MB con node_modules) y `jarvis-dashboard-clean.tar.gz` (70 KB)

## Por qué se movió

El código + `node_modules` ocupaba 119 MB dentro del vault. Obsidian indexaba todo y Syncthing lo replicaba en cada PC. Mover a `agent-stack/` (repo git existente) resuelve ambos problemas y permite versionar el código.

## Verificación pendiente

- Confirmar vía SSH que el VPS sirve este dashboard en `jarvss.duckdns.org` y no el rebrand antiguo.
- Decidir si `agent-stack/jarvis-ui-overrides/` (overlay Plan v3) sigue activo o queda obsoleto.

← [[index|Volver a projects]] | [[../Dashboard|Dashboard métricas]]
