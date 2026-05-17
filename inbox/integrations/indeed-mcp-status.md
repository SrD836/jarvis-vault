---
date: 2026-05-16
integration: indeed-mcp
status: blocked-for-third-parties
investigated_by: claude-sonnet-session
related:
  - "[[03-decisions/2026-05-14-skill-logs-duplication]]"
  - "[[agents/external/gsd-integration-checker]]"
  - "[[inbox/integrations/multi-provider-roadmap]]"
  - "[[inbox/job-hunting/2026-05-16-ai-engineer-cross-source]]"
  - "[[projects/job-hunt/2026-05-16-search]]"
# auto-linked 2026-05-17
---


# Indeed MCP — estado de integración para JARVIS

## TL;DR

**No es viable** integrar el Indeed MCP server desde JARVIS sin partnership formal aprobado por Indeed. Dynamic Client Registration es público y concede tokens, pero todos los endpoints reales de search/data devuelven `FORBIDDEN` para client_ids fuera del allowlist de partners.

## Lo que funciona

- ✅ `https://secure.indeed.com/oauth/v2/register` — Dynamic Client Registration sin login. Concede `client_id` + `client_secret`.
- ✅ Authorization Code flow + PKCE — token válido emitido.
- ✅ `offline_access` + refresh_token — token renovable.
- ✅ `https://secure.indeed.com/v2/api/userinfo` — devuelve el `sub` del usuario autorizado.

## Lo que NO funciona (con DCR client)

- ❌ `https://mcp.indeed.com/claude/mcp` — 401 invalid_token. Endpoint gateado a tokens con `aud` específico de claude.ai.
- ❌ `https://apis.indeed.com/graphql` — `Client is not authorized` / FORBIDDEN incluso para introspection.
- ❌ `https://apis.indeed.com/v2/jobs/search` (y variantes) — 403 Forbidden.
- ❌ Cualquier scope distinto de `job_seeker.jobs.search` + `offline_access` rechaza el registration mismo.

## JWT decodificado de nuestro token

```
aud: "https://default.indeed.com"
azp: "<nuestro-client_id>"
scope: "offline_access job_seeker.jobs.search"
```

El `aud` que claude.ai usa probablemente es `https://mcp.indeed.com` o similar — reservado a su client_id como partner.

## Conclusión

Indeed mantiene un **walled garden**: el MCP server público existe (`https://mcp.indeed.com/claude/mcp`) pero solo claude.ai tiene un client_id con `aud` para acceder. Mismo patrón en su API REST y GraphQL.

**Para que JARVIS use Indeed via MCP necesitarías:**
1. Aplicar a Indeed Partner Program (https://docs.indeed.com)
2. Obtener approval (semanas, requiere caso de uso comercial)
3. Recibir client_id+secret con allowlist de scopes/audiences expandido

## Camino alternativo recomendado

El job-hunter ya tiene skill `job-search` documentando WebFetch como path nativo:
- `https://es.indeed.com/jobs?q=<q>&l=<loc>` — listings públicos (sin auth)
- `https://www.linkedin.com/jobs/search?keywords=<q>` — snippets públicos
- `https://www.tecnoempleo.com/busqueda-empleo.php?te=<q>` — España dev
- `https://www.infojobs.net/jobsearch/search-results/list.xhtml?keyword=<q>` — España general

Resultado: las mismas ofertas que el MCP devuelve, mediante WebFetch, sin auth, sin partnership.

## Decisión

- **NO** integrar Indeed MCP real (bloqueado).
- **SÍ** mantener el código OAuth helper (`C:\temp\indeed-oauth-helper.mjs`) por si David obtiene partnership futuro — entonces flip de tokens es trivial.
- **Job-hunter usa WebFetch** sobre Indeed/LinkedIn/Tecnoempleo/InfoJobs — flow ya documentado en skill `job-search`.
- Sobre "no quiero que expire nunca": N/A porque no vamos a usar tokens. Si en el futuro se usa, el código del backend hará auto-refresh con offline_access.

## Otros conectores claude.ai — ¿mismo gate?

Probable que SÍ para conectores oficiales de empresas grandes (Slack, Linear, Notion, Gmail, Indeed). Probable que NO para connectors community/open-source (GitHub, Sentry, Cloudflare, etc) — esos suelen aceptar DCR sin allowlist.

Estrategia para "decenas de conectores": probar caso por caso con DCR. Si DCR no funciona, asumir walled garden.
