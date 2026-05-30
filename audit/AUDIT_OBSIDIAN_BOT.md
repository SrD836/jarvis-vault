---
title: "Auditoria Obsidian — nodos del bot Polymarket: residuo vs senal"
date: 2026-05-29
audit_scope: "Clasificacion de TODOS los nodos del vault relacionados con polymarket-veto-loop-bot; % de ruido; estructura limpia; plan de limpieza reversible"
auditor: "Claude Opus 4.8 (1M context)"
status: EJECUTADO_2026-05-29
related:
  - "[[agents/polymarket-bot]]"
  - "[[agents/polymarket-bot/memory]]"
  - "[[projects/polymarket-veto-loop-bot]]"
  - "[[audit/2026-05-26-polymarket-bot-audit]]"
tags: [audit, polymarket, bot, obsidian, vault-hygiene]
---

# Auditoria Obsidian — nodos del bot Polymarket (2026-05-29)

## TL;DR

1. **La premisa del dashboard es enganosa.** "126 reglas blandas" y "+24734 rechazos por cupo" NO son explosiones de nodos: son filas/lineas en UN solo fichero cada uno (arquitectura correcta).
2. **El ruido real esta en `03-decisions/`**, que el dashboard ni muestra: **1639 nodos auto-generados** por el bot (968 veto + 530 trade + 141 loss) de 1655 totales en esa carpeta.
3. **El ruido es el ~49.7% de TODO el markdown del vault** (1639 de 3296 `.md`). Dentro de `03-decisions/`, el **99%** es residuo del bot.
4. **Los nodos no estan "sueltos"**, estan **sobre-enlazados**: los 1639 apuntan a los mismos 3 hubs -> el grafo es una bola de pelo y los hubs pierden significado.
5. **0 de 300 nodos muestreados** tienen `## Human notes` no-vacias -> cero senal humana -> archivado seguro.
6. **Borrar sin frenar el writer es inutil:** el dedupe del 2026-05-26 bajo a 504 nodos; 3 dias despues hay 1655 (+~380/dia). Hay que tapar la fuente en codigo.

---

## Seccion 1 — Conteo por tipo (evidencia)

| Nodo / fichero | Conteo | Donde | Generador | Veredicto |
|---|---|---|---|---|
| `*-polymarket-veto-*.md` | **968** | `03-decisions/` | `brain.go:657` `decisionlog.WriteVeto` | **RESIDUO** — auto, `outcome: pending`, repetitivo, 1/mercado |
| `*-polymarket-trade-*.md` | **530** | `03-decisions/` | `executor.go:314` `decisionlog.WriteTrade` | **RESIDUO** — auto, 1/trade |
| `*-polymarket-loss-*.md` | **141** | `03-decisions/` | `loglosses.go:24` `WriteLoss` | **RESIDUO** — post-mortem auto, tambien en memory.md `## Losses` |
| decisiones NO-bot | **14** | `03-decisions/` | manual | **SENAL** — conservar |
| `index.md` | 1 | `03-decisions/` | manual | SENAL — conservar |
| `memory.md` (vetos<=500, losses 88, soft-rules 126) | 1 fichero (130 KB) | `agents/polymarket-bot/` | bot, **capped** | **SENAL** — el brain re-lee (M1/M2/M3); conservar |
| `rejections.jsonl` | 24944 lineas (4.2 MB) | `inbox/trading/` | scanner/brain | LOG correcto (no son nodos) — conservar; rotar opc. |
| `closed.jsonl` (381) `predictions.jsonl` (968) `candidates.jsonl` (500) `blocked.jsonl` (394) `approved.jsonl` (46) | logs | `inbox/trading/` | bot | FUENTE-DE-VERDAD — conservar |

**Distribucion temporal de `03-decisions/` (nodos/dia, por fecha en el nombre):** 280 (05-19), 207 (05-29), 159 (05-26), 157 (05-25), 155 (05-28), 143 (05-27)... -> ~130-280 nodos/dia.

---

## Seccion 2 — Residuo vs Senal (criterio)

**SENAL** = decision con valor de aprendizaje para un humano (notas, contexto, razonamiento unico).
**RESIDUO** = un nodo por evento trivial, auto-generado, sin notas, cuyo dato ya vive integro en el log o en memory.md.

- **968 veto + 530 trade + 141 loss = RESIDUO.** Auto-generados, frontmatter identico, cuerpo plantillado ("memoria: exact slug match score 1.00"), `outcome: pending`, `## Human notes` siempre vacia (0/300 muestreados). Cada dato ya esta en `inbox/trading/*.jsonl` (replay) y en `memory.md` (aprendizaje que el brain re-lee). El nodo individual no anade nada — solo infla el grafo.
- **14 decisiones no-bot + index.md = SENAL.** Decisiones manuales del sistema (budget-guardian, memory-wiki-status, plan-v6, etc.). Conservar y reconectar.
- **memory.md, rejections.jsonl, demas jsonl = capa correcta**, no son nodos Obsidian. El "ruido" que ve el dashboard (126 soft-rules, 24734 rechazos) es solo el conteo de filas/lineas de esos ficheros unicos. Estan bien.

---

## Seccion 3 — Estructura limpia propuesta

**Principio: Obsidian = solo lo legible/util para un humano. El dato bruto vive en el log; el aprendizaje en memory.md.**

| Capa | Contenido | Fichero | Estado |
|---|---|---|---|
| **Fuente de verdad / replay** | cada veto, trade, rechazo, prediccion, candidato | `inbox/trading/*.jsonl` | ya existe |
| **Aprendizaje (la maquina re-lee)** | patrones veto/loss + soft-rules, capped | `agents/polymarket-bot/memory.md` | ya existe |
| **Obsidian legible (humano)** | **1 rollup/dia** + nodo proyecto + nodo memoria + post-mortems CON notas humanas | `03-decisions/<date>-polymarket-daily.md` | NUEVO (1/dia) |

-> **NUNCA un nodo por veto/trade/loss.** Pasa de ~130-280 nodos/dia a ~1/dia.

> Nota: el bot no usa SQLite como fuente de verdad (eso es la memoria del AGENTE JARVIS, `hermes/memory_sqlite.py`, otra capa). La fuente de verdad del bot es el conjunto `inbox/trading/*.jsonl`. La distincion "Obsidian vs log" es la correcta aqui.

---

## Seccion 4 — Plan de limpieza (GATED en OK de David; reversible)

### Parte A — Frenar el firehose (codigo, raiz)
Gate los 3 writers detras de env flag `DECISION_NODES`, **default OFF**. memory.md + jsonl ya capturan todo.
- `bot/common/decisionlog/decisionlog.go`: helper `NodesEnabled()`; `WriteVeto/WriteLoss/WriteTrade` early-return si off. TDD off/on.
- Sin tocar el append a memory.md (`brain.go:638`) — la capa de aprendizaje se mantiene intacta.
- Build/test en VPS (`docker run golang:1.22` desde `bot/`, 5 modulos; Docker no corre en PC).

### Parte B — Rollup diario (el nuevo nodo senal)
- `analytics/daily_rollup.py` (determinista, sin LLM; reusa `brier.load_rows`). Escribe `03-decisions/<date>-polymarket-daily.md`: #vetos por regla, #trades, #losses, P&L dia, expectancy (de `calibration.json`), top-3 cierres, links a jsonl. Idempotente. Sin retroactivo.
- Cron `45 1 * * *` (tras daily_calibration), envuelto en `run_logger.py`.

### Parte C — Archivar el residuo existente (DESTRUCTIVO -> OK explicito)
- **Mover** (git mv, no borrar) 1639 nodos `*-polymarket-{veto,trade,loss}-*.md` -> `03-decisions/.archive-polymarket-md/cleanup-2026-05-29/` (reusa el patron `.archive-polymarket-md/` del dedupe 05-26).
- **Conservar:** 14 decisiones no-bot + index.md.
- **Reconectar:** verificar `related:` de los 15 supervivientes; enlazar desde `00-MOC.md`/`index.md` si falta.
- **NO tocar:** memory.md, todos los `inbox/trading/*.jsonl`, backups `*.bak-*`.

### Plan de limpieza — numeros
- **Se archivan: 1639 nodos** (968 + 530 + 141).
- **Se conservan: 15 nodos** (14 + index).
- **Se agregan: 1 nodo/dia** (rollup) reemplaza ~130-280/dia.
- **Se reconectan: 15** (verificacion de enlaces de supervivientes).
- **Reduccion de markdown del vault: ~49.7%** (3296 -> ~1657 `.md`).

---

## Seccion 5 — Verificacion (post-OK)
1. `go test ./...` verde (gate: off->0 escrituras+sin error; on->escribe).
2. `daily_rollup.py` sobre data viva -> 1 `*-polymarket-daily.md` con conteos + P&L + expectancy.
3. Tras archivar: `03-decisions/*.md` top-level = **15**; `.archive-polymarket-md/cleanup-2026-05-29/` = **1639**. `git status`: moves, 0 deletes.
4. Re-correr brain+executor en sim con `DECISION_NODES` unset -> **0 nodos nuevos** en top-level (solo el rollup).

---

## Anexo — Reproduccion (read-only)

```bash
ROOT="<vault>"
# Conteo por tipo
find "$ROOT/03-decisions" -maxdepth 1 -name '*.md' | wc -l                       # 1655
find "$ROOT/03-decisions" -maxdepth 1 -name '*polymarket-veto*.md'  | wc -l       # 968
find "$ROOT/03-decisions" -maxdepth 1 -name '*polymarket-trade*.md' | wc -l       # 530
find "$ROOT/03-decisions" -maxdepth 1 -name '*polymarket-loss*.md'  | wc -l       # 141
find "$ROOT" -name '*.md' -not -path '*/.stversions/*' | wc -l                    # 3296 -> ruido 1639/3296 = 49.7%
# rejections es UN log, no nodos
wc -l "$ROOT/inbox/trading/rejections.jsonl"                                      # 24944
# 0 notas humanas en el residuo
for f in $(find "$ROOT/03-decisions" -maxdepth 1 -name '*polymarket-veto*' | head -300); do
  awk '/## Human notes/{f=1;next} f' "$f" | grep -v '_(no se toca' | tr -d '[:space:]'; done | grep . | wc -l   # 0
```

**Conclusion:** el vault es un buen archivo de datos (jsonl) pero un mal grafo de conocimiento (1639 stubs auto). El arreglo es de raiz (frenar el writer) + agregar (rollup diario) + archivar reversible. Capital real = 0 (simulacion); toda la limpieza es reversible (git mv a `.archive-polymarket-md/`).
