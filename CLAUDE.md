# CLAUDE.md — Vault context para Claude

> Este archivo lo lee toda IA que entre al vault. Carga contexto en 30 segundos para que el output sea afilado en vez de genérico.

---

## Quién es el user

**David González Nuez** · davidgnjunior@gmail.com · Teror, Gran Canaria, España.

- Estudia: Ciclo Formativo de Grado Superior en **Administración y Finanzas** (IES Felo Monzón Grau-Bassas). **NO** ULPGC (sí cursó 2 años de Grado en Economía allí antes).
- Inglés **B1** (canon). Cualquier nota o CV que diga B2/C1 es alucinación → corregir.
- Formación adicional: Google IT Support (Coursera), Inteligencia Artificial con Microsoft Azure.
- Experiencia laboral declarada: camarero/cajero/host en The Fallow (Irlanda, verano 2021).
- Busca: soporte técnico junior, atención al cliente, AI engineering trainee, VA remoto. Cuota auto-apply 10/día.

---

## Qué construye — JARVIS

Sistema multi-agente personal corriendo 24/7 en VPS Hetzner (`agent@88.198.168.61`, key `id_ed25519`). Interfaz primaria Telegram. Dashboard web `jarvss.duckdns.org` (puerto 3000 internamente).

**Stack**:
- Backend: Go + Docker.
- Frontend: React (jarvis-dashboard).
- Sync vault PC↔VPS: Syncthing.
- IA: Claude Max OAuth via `docker exec openclaw-fork-openclaw-gateway-1 claude -p`. **NUNCA** `api.anthropic.com` con `ANTHROPIC_API_KEY` (regla dura: David paga Max, doble pago si usas API).
- Computer Use: PC Playwright + Desktop Commander, túnel autossh PC↔VPS, Chrome CDP attach `C:\jarvis\chrome-cdp-profile` :9222.

**Agentes principales** (ver `[[agents/index]]`):
- `main` — orquestador Telegram.
- `planner` · `apier` · `archivist` · `auditor` · `code-reviewer` · `debugger` · `designer` · `documenter` · `job-hunter` · `monitor` · `researcher` · `skill-reviewer` · `tester`.
- `polymarket-bot` — autónomo 24/7, opera vía cron host (scanner/brain/executor */30, exit_monitor */5).

**Proyectos vivos**:
- `/trabajo` (LinkedIn Easy Apply auto + portfolio + CV adapt).
- `polymarket-bot` v6 con reglas hardcoded P0/P1/P2/P3/P4.
- `jarvis-dashboard` con tabs Brainstorm/Manual, chat-processor (claudemax provider).

---

## Cómo piensa David

- **Caveman style** (`~/.claude/skills/caveman/`): terse, fragments OK, drop filler/hedging/pleasantries.
- **Critical thinking mode** (default ON): cuestionar premisas, pushback con razones técnicas, una pregunta concreta o ejecutar — **nunca** 3 preguntas vagas.
- **Anti-complacencia**: "tienes razón" SOLO cuando user tiene razón. Si está equivocado → contradecir con evidencia.
- **Autonomía nocturna 22:00-08:00**: agentes deciden sin preguntar. Formato "Decisión tomada · Razón · Reversibilidad · Confianza".
- **Token efficiency**: prompt cache TTL 5 min. Lock de modelo, no `/model` mid-task. `/compact` al 80%, no al 95%. `/clear` sólo entre temas distintos.

---

## Reglas duras (do NOT violate)

1. **Nunca API Anthropic con key** → solo Claude Max OAuth via `docker exec`.
2. **jarvis-dashboard NO bind mount** → cambios al código host no afectan runtime sin `docker compose build` + restart.
3. **applications.jsonl anti-fraude** → marcar `applied/submitted` requiere `linkedin_confirmation:true` + `submitted_at` no-null (validation en `applications.js`).
4. **main.md regenerado por cron** → directivas persistentes van en `## Human notes` dentro de AGENTS.md (`~/.openclaw/workspace/AGENTS.md`), NO en `vault/agents/<id>.md` (eso es view).
5. **Sin emojis** en código/CV salvo que David pida (los emojis del vault son OK porque son notas humanas).

---

## Carpetas del vault — qué es qué

| Carpeta | Contenido | Auto/Manual |
|---|---|---|
| `00-MOC.md` | Hub raíz, único punto de entrada — leer SIEMPRE primero | Manual |
| `CLAUDE.md` (este archivo) | Contexto para IA que entra al vault | Manual |
| `01-briefs/` | Briefs diarios 06:00 | `daily-brief.py` cron |
| `02-sessions/` | Exports Telegram multi-session | `session_export.py` real-time |
| `03-decisions/` | Razonamiento + outcome de decisiones | `backfill_links.py` + manual |
| `04-skills-log/` | Diffs catálogo skills (new/stale/archive) | `curator.py` cron 03:30 |
| `agents/` | Mirror del equipo multi-agente (orgchart Mermaid) | `orgchart.py` cron 30min — NO editar manual, lo borra |
| `skills/` | Mirror `~/.openclaw/skills/` | `mirror_external.py` cron |
| `Code-Audits/` | Auditorías overnight de repos | `overnight-code` cron 01:00 |
| `inbox/job-hunting/` | Búsquedas + aplicaciones reales | Manual + dashboard |
| `attachments/portfolios/` | HTML sites adaptados por oferta | `apply-batch.js` (pipeline) |
| `projects/` | Roll-ups manuales cross-session | Manual |
| `seeds/` | Decisiones arquitectónicas duraderas | Manual |
| `wiki/` | Knowledge layer compilado (memory-wiki) | DORMIDO (P3.1 reactivará) |
| `memories/` | Concepts/entities dialécticos (thesis/antithesis/synthesis) | DORMIDO |
| `templates/` | Plantillas | Referencia |
| `audit/` | Auditorías del propio vault | Manual / on-demand |
| `secrets/`, `Research/` | Vacías hoy | — |
| `.stversions/` | Histórico Syncthing | Ignorar |

---

## Memorias persistentes claves

Cargadas auto al iniciar sesión desde `~/.claude/projects/.../memory/MEMORY.md`. Si alguna parece stale (>24h, conflicto con código actual), **verificar en el codebase** antes de actuar y actualizar la memoria.

Top hits:
- `jarvis_vps_access` — credenciales y paths VPS.
- `jarvis_never_anthropic_api` — regla dura Claude Max only.
- `jarvis_dashboard_no_bindmount` — recordar siempre.
- `jarvis_autonomy_pattern` — horario nocturno.
- `polymarket_v6_2026_05_19` — reglas hardcoded del bot.
- `jarvis_trabajo_turn8_complete` — pipeline LinkedIn Easy Apply.

---

## Cómo trabajar con David — protocolo

1. **Leer primero**: este archivo + `MEMORY.md` + `00-MOC.md`. Cargado en <60s.
2. **Repomix**: si existe `repomix-output.xml` <500 KB → leer ANTES de Read/Grep/Glob individuales. Si >500 KB → ignorar.
3. **Caveman output**: fragments OK, sin filler. Drop "buena pregunta", "por supuesto".
4. **Si user da ruta `@path/file`** → leer directo, no buscar.
5. **No mocks**: integration tests sobre interfaces públicas (`~/.claude/skills/tdd/`).
6. **Cambios destructivos**: backup antes (`*.bak-YYYY-MM-DD`).
7. **Dashboard cambios**: edit en VPS → `docker compose build dashboard && docker compose up -d dashboard`.
8. **Tareas largas**: TaskCreate al inicio. Mark in_progress al empezar, completed al terminar (no batch).
9. **Memorias**: guarda decisiones surprising/non-obvious. NO guardes git history ni patterns derivables del código.

---

## Estado actual del vault (snapshot 2026-05-24)

- 477 MB · 2.043 notas · 12 MOCs · 6 cron generators activos.
- Auditoría completa: [[audit/2026-05-24-vault-audit]].
- **Gap principal**: vault es buen archivo, falla en cognición. memory-wiki dormido, no weekly synthesis, no daily insight push, no dataview.
- **Roadmap segundo cerebro** (esperan aprobación David):
  - ✅ P1.1 conectar huérfanos · ✅ P1.2 este CLAUDE.md
  - ⏳ P2.1 weekly synthesis cron · P2.2 plugins Obsidian · P2.3 tag taxonomy
  - ⏳ P3.1 reactivar memory-wiki · P3.2 daily Telegram insight push

---

## Cuando entres al vault como IA

Pregúntate antes de responder:
- ¿He leído `MEMORY.md` + este archivo + `00-MOC.md`?
- ¿La acción es reversible? Si no, ¿confirmé con David?
- ¿Estoy duplicando trabajo que un subagente puede hacer?
- ¿Estoy a punto de cometer alguna de las reglas duras (API key, bind mount, anti-fraude)?

Si todo OK → ejecuta. No preguntes para validar.
