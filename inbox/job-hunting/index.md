---
title: "Job hunting — Índice"
type: moc
created: 2026-05-16
updated: 2026-05-28
tags: [moc, job-hunting, trabajo, career]
related:
  - "[[00-MOC]]"
  - "[[projects/portfolio-offers]]"
  - "[[USER]]"
---

# 💼 Job hunting log — Índice

Tracking de búsquedas realizadas, ofertas presentadas, estado de aplicación. Dashboard live: jarvss.duckdns.org → sección /trabajo. Cuota auto-apply: 10/día (LinkedIn Easy Apply).

## Búsquedas

| Fecha | Query | Source | Total | Top picks | Status | Archivo |
|---|---|---|---|---|---|---|
| 2026-05-16 | AI Engineer + Backend + Remote DevOps Madrid | indeed-mcp | 30 | 5 | presented | [2026-05-16-madrid-ai.md](2026-05-16-madrid-ai.md) |
| 2026-05-16 | AI Engineer Madrid (cross-source) | indeed-web+linkedin+tecnoempleo | 30 | 8 | presented | [2026-05-16-ai-engineer-cross-source.md](2026-05-16-ai-engineer-cross-source.md) |
| 2026-05-17 | Multi-bucket David (IT support, AI training, VA) | multi-source | — | — | presented | [2026-05-17-david-multi.md](2026-05-17-david-multi.md) |
| 2026-05-18 | LinkedIn multi-bucket (IT support, AI training, CS, VA) | linkedin-only | 10 | 10 | presented | [2026-05-18-trabajo-page-linkedin.md](2026-05-18-trabajo-page-linkedin.md) |
| 2026-05-20 | auxiliar admin / helpdesk / soporte IT / data entry / atención cliente | linkedin+indeed+tecnoempleo | ~25 | 10 | presented | [2026-05-20-trabajo-1779305625.md](2026-05-20-trabajo-1779305625.md) |
| 2026-05-27 | auxiliar admin / back office / helpdesk / data entry / soporte IT (UI Trabajo) | linkedin+indeed+infojobs+tecnoempleo | ~30 | 10 | presented | [2026-05-27-trabajo-1779901440.md](2026-05-27-trabajo-1779901440.md) |
| 2026-05-28 | auxiliar admin / helpdesk / soporte IT / data entry / atención cliente / back office (UI Trabajo) | linkedin+indeed+tecnoempleo | ~35 | 10 | presented | [2026-05-28-trabajo-1748472600.md](2026-05-28-trabajo-1748472600.md) |
| 2026-05-28 | auxiliar admin / helpdesk / soporte IT / data entry / auxiliar contable / atención cliente (UI Trabajo #2) | linkedin+indeed+tecnoempleo | ~35 | 10 | presented | [2026-05-28-trabajo-1780009080.md](2026-05-28-trabajo-1780009080.md) |

## Empresas contactadas / aplicadas

Fuente de verdad: `inbox/job-hunting/applications.jsonl` + dashboard `/trabajo → Historial`. Sólo entries con `linkedin_confirmation:true` cuentan como envío real.

| Empresa | Puesto | Aplicado | Estado | Notas |
|---|---|---|---|---|
| IOON | Soporte Técnico Junior Genesys Cloud CX | 2026-05-24 17:12 | ✅ Enviada (LinkedIn confirma) | [[projects/portfolio-offers]] → offer-3-ioon · primera Easy Apply real |

## Blacklist

_(Empresas/listings a excluir en futuras búsquedas. Una por línea o tabla.)_

## Notas de fuentes

- **Indeed MCP**: funciona desde el harness de Claude Code. Ratio relevancia alto cuando query incluye stack específico ("AI Engineer Python LLM").
- **LinkedIn**: solo URLs públicas via WebFetch (auth-gated content prohibido). Easy Apply usa Chrome CDP attach + Playwright MCP.
- **Tecnoempleo / InfoJobs**: fallback España, no testeados aún.

## Sistema (cómo se conectan las piezas)

- **Pipeline**: `dashboard/api/trabajo/search` → job-hunter agent → `applications.jsonl` → `apply-batch.js` (portfolio → CV adapt → LinkedIn Easy Apply).
- **Anti-fraude**: tras 2026-05-24 sólo se marca `applied/submitted` si `linkedin_confirmation:true` + `submitted_at` no-null.
- **Cuota**: `applied-today.json` por día, cap 10.
- **CV templates**: executive · sidebar · compact-pro · jglovier (auto-pick por título oferta).
- **Memorias relevantes**:
  - `jarvis_trabajo_turn8_complete` — template jglovier + LLM auto-answer.
  - `jarvis_linkedin_apply_fase5c` — DOM scan, normalizeLabel, waitForSubmitConfirmation.
