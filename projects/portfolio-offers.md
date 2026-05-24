---
title: "Portfolios de ofertas — Índice"
type: project
created: 2026-05-24
tags: [project, portfolio, trabajo, job-hunting]
related:
  - "[[00-MOC]]"
  - "[[inbox/job-hunting/index]]"
  - "[[projects/jarvis-dashboard]]"
---

# 💼 Portfolios generados por oferta

HTML sites estáticos adaptados a cada oferta, generados por el pipeline `apply-batch.js` en jarvis-dashboard. Carpeta raíz: `attachments/portfolios/`.

Cada portfolio es servido por nginx VPS en `https://jarvss.duckdns.org/portfolios/<slug>/` y el link va inyectado en el header del CV adaptado.

## Aplicaciones reales

| Oferta | Empresa | Estado | Portfolio | Fecha |
|---|---|---|---|---|
| Soporte Técnico Junior Genesys Cloud CX | IOON | ✅ **Enviada** | [offer-3-ioon-...-1e8241](../attachments/portfolios/offer-3-ioon-ioon-soporte-t-cnico-junior-genesys-cloud-cx-1e8241/) | 2026-05-24 |

## Adaptaciones generadas (no enviadas)

| Oferta | Empresa | Carpeta |
|---|---|---|
| Fully Remote Executive Assistant | Donna Pro | [offer-10-donna-pro-...-9fe80d](../attachments/portfolios/offer-10-donna-pro-donna-pro-fully-remote-executive-assistan-9fe80d/) |
| Administrative Assistant (WFH) | Persona | [offer-1-persona-...-d564f6](../attachments/portfolios/offer-1-persona-persona-administrative-assistant-work-from-h-d564f6/) |
| Virtual Assistant (WFH) | Persona | [offer-2-persona-...-40da07](../attachments/portfolios/offer-2-persona-persona-virtual-assistant-work-from-home-40da07/) |
| Customer Support Specialist | Holafly | [offer-3-holafly-...-7d90ff](../attachments/portfolios/offer-3-holafly-holafly-customer-support-specialist-english-7d90ff/) |
| Technical Customer Success Manager | Hubtype | [offer-4-hubtype-...-e1ef4f](../attachments/portfolios/offer-4-hubtype-hubtype-technical-customer-success-manager-e1ef4f/) |
| Admin Billing & Compliance Ops | Hubtype | [offer-5-hubtype-...-ef79aa](../attachments/portfolios/offer-5-hubtype-hubtype-admin-billing-compliance-operations--ef79aa/) |
| Remote Event Operations Assistant | Your Startup Operations | [offer-7-your-startup-ops-...-bac456](../attachments/portfolios/offer-7-your-startup-operations-your-startup-operations-remo-bac456/) |
| Backend Systems AI Model Trainer | ChatGPT Jobs | [offer-9-chatgpt-jobs-...-2abb6c](../attachments/portfolios/offer-9-chatgpt-jobs-chatgpt-jobs-backend-systems-ai-model-t-2abb6c/) |
| IT Helpdesk Technician | Sardine | [offer-1-sardine-...-e015d0](../attachments/portfolios/offer-1-sardine-sardine-it-helpdesk-technician-e015d0/) |

## Tests / sandbox

| Slug | Notas |
|---|---|
| `test-helpdesk-sardine-...-3c8a9c` | Test pipeline inicial |
| `test-v2-sardine-...-27aae8` | Test iteración v2 |

## Pipeline

Generadas por `/home/agent/jarvis-dashboard/src/lib/portfolio-builder.js` desde `apply-batch.js`. Cada carpeta contiene `index.html`, `style.css`, `robots.txt` + assets adaptados al perfil de la oferta. El CV adaptado las referencia en su línea de contacto: `🌐 portfolio`.

## TODO

- [ ] Auto-mantener este índice cuando se generen nuevos portfolios (script `update-portfolio-index.py` cron post-batch).
- [ ] Limpiar carpetas `test-*` cuando dejen de ser útiles.
- [ ] Marcar portfolios cuyas ofertas se hayan cerrado / aplicado / rechazado.
