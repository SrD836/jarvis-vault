---
date: 2026-05-24
type: audit
scope: vault-obsidian
status: complete
related:
  - "[[00-MOC]]"
  - "[[CLAUDE]]"
  - "[[agents/archivist]]"
  - "[[agents/auditor]]"
  - "[[audit/2026-05-24-system-e2e-audit]]"
tags: [audit, vault, second-brain, obsidian]
# auto-linked 2026-05-25
---


# Auditoría Vault Obsidian — 2026-05-24

Disparada por: petición user tras leer "Your Obsidian Vault Is Probably Dead". Objetivo = pasar de **almacén de logs** a **segundo cerebro vivo** (capture · automation · memory · intelligence).

Métodos: 3 Explore agents read-only en paralelo · grep · glob · cron inspection.

---

## 1. Inventario

| Métrica | Valor |
|---|---|
| Tamaño | 477 MB (sin .stversions) · 1.2 GB con histórico Syncthing |
| Notas `.md` | **2.043** |
| MOCs (índices) | 12 |
| Cron generators activos | 6 |
| `.obsidian/community-plugins.json` | **NO existe** (dataview/templater no instalados) |

### Carpetas top-level

| Carpeta | # notas | Generador |
|---|---:|---|
| `00-MOC.md` | 1 | Manual (hub raíz) |
| `01-briefs/` | 13 | `daily-brief.py` cron 06:00 |
| `02-sessions/` | 88 | `hermes/session_export.py` (Telegram multi-session) |
| `03-decisions/` | 1.321 | `backfill_links.py` + manual |
| `04-skills-log/` | 7 | `curator.py` cron 03:30 |
| `agents/` | 55 | `hermes/orgchart.py` cron 30 min |
| `skills/` | 105 | `mirror_external.py` cron |
| `Code-Audits/` | 6 | `overnight-code` cron 01:00 |
| `inbox/` | 33 | Chat + job-hunting manual/export |
| `wiki/` | 21 | `memory-wiki` (dormido) |
| `memories/` | 3 | `memory-wiki` (auto, vacío en concepts/entities) |
| `projects/` | 8 | Manual |
| `templates/, goals/, seeds/, Research/, secrets/, audit/` | 18 | Plantillas/manual |

---

## 2. Generadores auto — salud

```
03:30  curator.py         → 04-skills-log/*.md      ✅ linked
06:00  daily-brief.py     → 01-briefs/YYYY-MM-DD.md ⚠️  sin weekly rollup
real   session_export.py  → 02-sessions/YYYY-MM-DD/ ✅ linked via index
30min  orgchart.py        → agents/index.md         ✅ linked
?      mirror_external.py → skills/                 ✅ linked
?      memory-wiki        → wiki/, memories/        ❌ DORMIDO
01:00  overnight-code     → Code-Audits/*.md        ⚠️  no link a 00-MOC
real   backfill_links.py  → 03-decisions/index.md   ✅ linked
```

---

## 3. Nodos huérfanos (sin links entrantes ni salientes)

| Path | # | Severidad |
|---|---:|---|
| `inbox/job-hunting/*.md` | 6 | 🔴 alta — búsquedas sin síntesis ni MOC |
| `Code-Audits/*.md` | 6 | 🟡 media — linkean agents, no MOC raíz |
| `attachments/portfolios/offer-*/` | ~40 dirs HTML | 🔴 alta — sin nota índice |
| `02-sessions/*.md` (archivos sueltos) | 88 | 🟡 media — dependen 100% del index |
| `Research/`, `secrets/`, `audit/` | 0 | ➖ carpetas vacías |

---

## 4. Gaps detectados

1. **No `.obsidian/community-plugins.json`** → no dataview, no templater, no periodic-notes. Toda relación depende de wikilinks manuales.
2. **No `CLAUDE.md` en raíz del vault** → cualquier IA que recorra el vault arranca sin contexto de David. Sólo existe el global `~/.claude/CLAUDE.md`.
3. **No weekly synthesis** → template `templates/weekly-synthesis.md` existe pero ningún cron lo ejecuta. No hay rollup que detecte patrones cross-week.
4. **No tag taxonomy** → tags ad-hoc, sin semántica documentada.
5. **`memory-wiki` dormido** → carpetas `memories/concepts/` y `memories/entities/` vacías. No hay extracción dialéctica (thesis/antithesis/synthesis) desde sesiones.
6. **No daily push de insights** → el vault no "habla de vuelta".
7. **`inbox/job-hunting/` desconectado** → 6 búsquedas sin ningún MOC ni síntesis. Refleja exactamente la queja del user.

---

## 5. Diagnóstico — caveman

Vault funciona como **buen archivo bibliográfico**: cron escribe limpio, indexes auto-gen ok, 00-MOC raíz sólido.

Falla en **cognición**:
- Acumula sin sintetizar.
- Sin Claude conectado al vault como copiloto.
- Sin loop "vault habla de vuelta" (weekly synthesis, daily insights).
- Sin memory dialéctica (memory-wiki dormido).

= **second-brain sin segundo paso**. Tiene memoria. Falta inteligencia.

---

## 6. Propuestas Fase 3 — esperan decisión del user

Cada una es trabajo independiente. Marcar prioridad (P0/P1/P2) tras leer.

### P1.1 — Conectar huérfanos (10 min · 0 código)
- `inbox/job-hunting/index.md` con tabla de búsquedas + estado (aplicada/no).
- `projects/portfolio-offers.md` mapeando `offer-ID → attachments/portfolios/...`
- Linkear ambos + `Code-Audits/index.md` desde `00-MOC.md`.

### P1.2 — `CLAUDE.md` raíz del vault (15 min)
Enseña a Claude:
- Quién es David (rol, formación, B1 inglés, intereses).
- Qué construye (JARVIS multi-agente, Polymarket bot, /trabajo).
- Cómo piensa (caveman, anti-complacencia, modelos a usar).
- Memoria a leer (`~/.claude/projects/.../memory/MEMORY.md`).
- Qué carpetas del vault son qué.

Cualquier sesión IA que entre al vault arranca ya cargada de contexto → outputs dramáticamente más afilados.

### P2.1 — Weekly synthesis cron (1-2h)
`scripts/weekly-synthesis.py` lunes 06:00 Europe/Madrid:
- Lee 7d previos de `02-sessions/`, `03-decisions/`, `Code-Audits/`, `01-briefs/`.
- Llama Claude Max via `docker exec gateway claude -p` (NUNCA api.anthropic.com).
- Escribe `01-briefs/synthesis-week-NN.md` con: patrones recurrentes, contradicciones, preguntas que resurgen, próximas acciones.
- Linkea desde `00-MOC.md` y cada brief diario.

### P2.2 — Plugins Obsidian (30 min)
Instalar en `.obsidian/community-plugins.json`: `dataview`, `templater`, `periodic-notes`. Permite queries dinámicas `FROM #decision WHERE...`.

### P2.3 — Tag taxonomy (`tags-index.md`, 30 min)
Documentar semántica: `#decision`, `#session`, `#brief`, `#skill`, `#bug`, `#fix`, `#job`, `#audit`. Reglas para crear tags nuevos.

### P3.1 — Reactivar memory-wiki (2-3h)
Cron extrae:
- `memories/concepts/<concept>.md` — ideas recurrentes con thesis/antithesis/synthesis.
- `memories/entities/<entity>.md` — sistemas, personas, proyectos con estado bitemporal.
Fuente: sesiones + decisiones. Output linkado desde memories/index.md.

### P3.2 — Daily Telegram insight push (1h)
Cron 07:00: Claude lee captures del día anterior + memorias relevantes, genera 3 insights + 1 pregunta-que-vale-la-pena-pensar. Push a Telegram bot. Vault "habla de vuelta".

---

## 7. Pregunta para el user

Prioridad sugerida: **P1.1 + P1.2 primero** (1h total, cero coste, alto impacto). Después decidir P2/P3 según uso real.

¿Vamos con P1.1 + P1.2 ya?
