---
title: "Run — job-hunter — 2026-05-18T22:12:00"
type: agent-run-log
agent: job-hunter
date: 2026-05-18
tags: [job-hunter, run, linkedin]
---

# Run Log — 2026-05-18T22:12:00

## Task recibida

Buscar 10 ofertas en LinkedIn (fuente OBLIGATORIA) que matcheen el perfil de David. Persistir resultados en vault/inbox/job-hunting/2026-05-18-trabajo-page-linkedin.md, escribir memory.md para sesiones futuras, y generar este run file.

## Output

- **Archivo informe:** `/home/agent/agent-stack/vault/inbox/job-hunting/2026-05-18-trabajo-page-linkedin.md`
- **Memory snapshot:** `/home/agent/agent-stack/vault/agents/job-hunter/memory.md`
- **Ofertas encontradas:** 10 (sin shortfall)
- **Fuente usada:** linkedin.com/jobs exclusivamente

### Resumen de las 10 ofertas

| # | Rol | Empresa | Ubicación | Bucket |
|---|---|---|---|---|
| 1 | IT Helpdesk Technician | Sardine | Spain (remote) | 3 |
| 2 | Virtual Assistant WFH | Persona | Spain (remote) | 2 |
| 3 | Customer Support Specialist EN | Holafly | Spain (remote) | 3 |
| 4 | Technical Customer Success Manager | Hubtype | Barcelona (remote) | 3 |
| 5 | Admin/Billing/Compliance Ops (AI-Native) | Hubtype | Barcelona (remote) | 3/4 |
| 6 | Data Science Consultant — AI Trainer | DataAnnotation | Remote (US listed) | 1 |
| 7 | Technical Consultant — AI Trainer | DataAnnotation | Remote (US listed) | 1 |
| 8 | English AI Trainer & Evaluator | CNTXT AI | Remote | 1 |
| 9 | Backend Systems AI Model Trainer | ChatGPT Jobs | Remote | 1 |
| 10 | Junior Account Executive – AI Agents | Process Street | Remote | 3/4 |

## Tool calls

| # | Tool | Descripción |
|---|---|---|
| 1 | Read | /home/agent/agent-stack/vault/USER.md |
| 2 | Read | SKILL_INDEX.md |
| 3 | Read | job-search/SKILL.md |
| 4 | WebFetch | LinkedIn: AI annotation RLHF Spain |
| 5 | WebFetch | LinkedIn: IT support helpdesk remote junior Spain |
| 6 | WebFetch | LinkedIn: prompt engineer AI assistant Spain |
| 7 | WebFetch | LinkedIn: customer success tech Spain |
| 8 | WebFetch | LinkedIn: data annotation AI trainer (global) |
| 9 | WebFetch | LinkedIn: IT support junior remote (global) |
| 10 | WebFetch | LinkedIn: Azure cloud support junior (global) |
| 11 | WebFetch | LinkedIn: AI ops automation junior remote |
| 12 | WebFetch | LinkedIn: AI trainer annotator Spanish bilingual |
| 13 | WebFetch | LinkedIn: virtual assistant remote AI Spain |
| 14 | WebFetch | LinkedIn: soporte técnico remoto junior Spain |
| 15 | WebFetch | LinkedIn: Excel automation remote admin |
| 16 | WebFetch | LinkedIn: customer support specialist Spain |
| 17 | WebFetch | LinkedIn: DataAnnotation AI trainer remote |
| 18 | WebFetch | LinkedIn: Outlier AI freelance remote |
| 19 | WebFetch | LinkedIn: tecnico soporte helpdesk remoto Spain |
| 20 | WebFetch | LinkedIn: helpdesk IT support remote Spain |
| 21 | Bash | mkdir dirs + ls checks |
| 22-24 | Write | 3 archivos de output |

## Métricas

- **tool_calls_count:** ~24
- **spawned_children:** 0
- **duration_ms:** estimado ~180,000 ms (múltiples WebFetch paralelos + procesamiento)
- **tokens:** estimado ~15,000–20,000 input + output (sonnet-4-6)
- **aborted:** false

## Observaciones

- LinkedIn anti-bot activo para queries muy específicas (RLHF, prompt engineer Spain) — búsquedas devolvieron 0 resultados con filtros rígidos. Solución: ampliar keywords + reducir geo-filtro.
- Buckets 1 (annotation/RLHF) tienen poca presencia en LinkedIn ES — plataformas directas (Outlier.ai, mercor.com) son mejor canal para ese bucket.
- Buckets 3 con sede Spain tienen buena disponibilidad: Sardine, Persona, Holafly, Hubtype — prioridad alta.
- DataAnnotation acepta globalmente pero lista US como HQ — David debe verificar geo-elegibilidad en cada JD antes de aplicar.
