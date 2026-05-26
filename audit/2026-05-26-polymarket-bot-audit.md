---
title: Polymarket Veto-Loop Bot — Post-Deploy Audit
date: 2026-05-26
audit_scope: PR #1 (06dd6af) — horizon enforcement, brain learning loop, vault consolidation
auditor: Claude Opus 4.7 (1M context)
status: COMPLETO
tags: [audit, polymarket, bot, post-deploy]
---

# Polymarket Bot Audit — 2026-05-26

## TL;DR

1. **Feedback loop CERRADO pero con UN AGUJERO crítico**: las 21 posiciones cerradas por `force_close_horizon_excess` el 26-05 a las 15:36 UTC NO se escribieron a `memory.md`. El bot no aprenderá de ese cluster de pérdidas. Fix de 1 línea en el script de cierre forzado.
2. **Memoria Obsidian funcionando** (Vetos=501/500 saturado, Losses=88, M2 active=126). M1 ya filtra 7 vetos/cycle (era 6 pre-PR).
3. **Cuota 70/15/15 NO converge sin cambios**: el universo Polymarket solo ofrece ~8 short candidates por ciclo de 500. Imposible llenar 35 slots short con esa oferta.
4. **Contabilidad y outcomes verificados externamente**: 5/5 trades random tienen PnL matemáticamente correcto y sides alineados con realidad polymarket (gamma-api). **NO hay bug de cuenta ni misreporting**.

---

## Sección 1 — Feedback loop end-to-end

**Veredicto: LOOP_PARCIAL** ⚠

### Tabla del loop

| Step | Implementado | Evidencia (file:line) | Runtime check |
|---|---|---|---|
| Trade cierra (exit_monitor cron */5) | ✅ | `bot/exit_monitor/internal/monitor/monitor.go:129` llama `loglosses.LogClose` | `closed.jsonl` crece, último cierre 15:30:08Z |
| LogClose escribe wins+losses a memory.md | ✅ | `bot/exit_monitor/internal/loglosses/loglosses.go:22-27` (`appendMemoryRow`) | memory.md: 501 vetos + 88 losses, último 15:30:08Z |
| **force_close_horizon_excess también debería llamar LogClose** | ❌ | `bot/exit_monitor/cmd/force_close_horizon_excess/main.go` — grep `loglosses\|LogClose` → 0 matches | 21 cierres a 15:36 NO están en memory.md losses (último row es 15:30:08, no hay rows con timestamp 15:36-15:50) |
| Brain carga memory.md en cycle siguiente | ✅ | `bot/brain/internal/brain/brain.go:75-82`, `mem.Load(memoryPath)` cap 500 | Log 16:00: "Memory loaded: 500 vetos, 75 losses" |
| M1 (veto exact pattern) aplica | ✅ | `brain.go:225-250`, `mem.Match(cl)` score≥0.7 | Log 16:00: `M1=7` (subió desde 6 pre-PR, +17%) |
| M2 (soft-rules) ahora SÍ se leen | ✅ POST-PR | `softrules.LoadActiveRules` invocado en `brain.go:130-137`, `MatchVeto` en `brain.go:270-287` | Log 16:00: "M2: loaded 9 active soft rules", `M2=43` vetos aplicados |
| M3 (anti-patterns) aplica | ✅ | `antipatterns.LoadActive(memoryPath, minFreq=3)` `brain.go:119-125` | Log 16:00: `M3=0` esta corrida — sin matches en candidatos |
| GenerateAndAppend escribe nuevas reglas | ✅ | `bot/brain/internal/softrules/generate.go` | 126 reglas blandas en dashboard (sección "## Reglas blandas") |

### Hallazgos

- **BUG-1 (Crítico):** `force_close_horizon_excess` cierra trades en `closed.jsonl` + `portfolio.json` pero salta `loglosses.LogClose`. Las 21 force-closures del 26-05 son data huérfana — el bot no aprende del cluster.
- **OK:** El resto del loop funciona. Comprobado: M1 subió +17% post-PR confirmando que más vetos pasados se aplican; M2 ya no es escritura ciega (consumer.go agregado).

### Fix recomendado para BUG-1

En `bot/exit_monitor/cmd/force_close_horizon_excess/main.go`, después de `closer.AppendClosed(...)` y antes de `closer.RewriteActive(...)`, añadir:

```go
import "github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/loglosses"

memoryPath := os.Getenv("EXIT_MEMORY_PATH")
if memoryPath == "" {
    memoryPath = "vault/agents/polymarket-bot/memory.md"
}
for _, ct := range closedTrades {
    loglosses.LogClose(closer.DecisionsDir, memoryPath, ct)
}
```

3 líneas de import + 5 líneas de loop. Tras este fix, futuras rebalances forzadas también enriquecerán memoria.

---

## Sección 2 — Memoria Obsidian

**Veredicto: OK con saturación incipiente**

| Métrica | Esperado | Observado | Desviación |
|---|---|---|---|
| `## Vetos pattern` rows | 100-500 | **501** | Cap alcanzado (1 fila excedente) |
| `## Losses pattern` rows | 50-500 | 88 | Saludable, espacio libre |
| `## Reglas blandas` rows | 9-126 | 126 | Coincide con dashboard |
| `## Anti-patterns` section | 1+ línea | 1 line + tags inline | OK pero header mal formateado |
| Cap 100→500 aplicado | sí | sí | `memory.go:46` confirmado |

### Hallazgos

- **WARN-1:** Sección Vetos tiene 501 filas (cap 500). La rotación append-only con tail debería estar funcionando, pero el extra row sugiere off-by-one. No crítico, lo manejará la próxima append.
- **WARN-2:** Header "## Anti-patterns identificados" tiene contenido inline en la propia línea del header (mira `memory.md:854`). Funciona pero es feo y puede confundir parsers futuros.
- **OK:** Vetos M1 muestran campos completos: `slug | category | price | rule | reason | ts`. Sample 3 vetos last → todos campos presentes.
- **OK:** Losses muestran `id | category | entry | exit | pnl | reason | empty | size | horizon | days`. Format consistent.

### Conteo de Obsidian decision nodes

- Tras dedupe Fase 3: 504 .md canónicos en `vault/03-decisions/` (era 1.720)
- 627 duplicados archivados en `.archive-polymarket-md/dedupe-2026-05-26/`
- Grafo de obsidian debería cargar sin lag (no verificado en GUI desde aquí)

---

## Sección 3 — Proyección distribución 70/15/15

**Veredicto: NO CONVERGE SIN CAMBIOS** ❌

### Datos crudos del ciclo 16:00

- Universo de candidates: **short=8 medium=121 long=317** (expired=46 descartados)
- Slot quota: short=35, medium=8, long=8 (en max_open=50)
- Aprobaciones brain últimos 5 runs: total **6 short / 13 medium / 16 long** = 17% / 37% / 46%
- Lifecycle promedio: no calculable (entry_timestamp es null en closed.jsonl — bug menor de tracking)

### Modelo simplificado

Para llenar 35 slots short en steady-state necesitas:
- **Oferta**: solo 8 short candidates/cycle. Si brain aprueba 100% de ellos = 8 short approvals/cycle.
- **Lifecycle**: short ≤2d teóricamente. Asume 1.5d promedio.
- **Llenado**: 48 cycles/día × 0.5 short approved/cycle (P0-P10 filtra ~95%) = 24 short/day apertura.
- **Cierre**: 35 / 1.5d = 23/day necesarias para cerrar al ritmo de apertura.
- **Steady-state esperado a 7 días**: ~35 short en cartera **SI el brain aprueba todos los short que pasan filtros y SI no hay drought de candidatos short**.

Pero hoy ofrecen 8 válidos, ayer ofrecían 48 (con expired contados). Volatilidad alta del universo. Si en una semana caen a 3-5 válidos diarios, no se llenarán 35 slots.

### Acciones recomendadas (priorizadas)

1. **Bajar filtro P3 liquidez para short**: actualmente $5k min absoluto. Markets short de baja liquidez pero high-conviction quedan fuera. Posible $2k para `horizon=short` solamente.
2. **Subir cap medium+long en cuota**: si solo hay 8 short candidates real, perseguir 70% es contraproducente. Realista 30/35/35 dado el feed.
3. **Scanner reordenado**: priorizar `endDate asc` PRIMERO antes que `volumeNum`. Aumentaría short candidates en cola desde 8 hacia 30+.

**Predicción honesta:** sin alguno de los 3 cambios, distribución se estabilizará en **15-25% short / 35-45% medium / 35-45% long** dentro de 7 días.

---

## Sección 4 — Verificación externa de 5 trades

**Veredicto: 5/5 OK — no hay bug contable ni de outcome**

| ID | Pregunta | Side | Entry → Exit | PnL bot | PnL math | Outcome real (gamma) | Verdict |
|---|---|---|---|---|---|---|---|
| T-2324252 | US x Iran peace deal by May 26? | YES | 0.13 → 0.27 | +$64.40 | +$64.42 | Unresolved (gamma: NO 96.6%) — bot vendió en pico antes de reversión | ✅ trading correcto |
| T-2227309 | Sunderland AFC win 2026-05-24? | YES | 0.84 → 0.999 | +$11.86 | +$11.85 | YES ganó (gamma: 1) | ✅ alineado |
| T-2241873 | Iran closes airspace by May 24? | YES | 0.38 → 0.07 | -$42.55 | -$42.55 | NO ganó (gamma: 0.9925) | ✅ alineado |
| T-789542 | Ed Gallrein KY-04 nominee? | YES | 0.5835 → 0.871 | +$73.91 | +$73.93 | YES ganó (gamma: 1) | ✅ alineado |
| T-2169751 | Trump tariff reduction on China? | YES | 0.459 → 0.016 | -$50.27 | -$50.27 | NO ganó (gamma: 1) | ✅ alineado |

### Observaciones

- **Trade 1 (Iran peace)**: caso interesante. Bot apostó YES a 0.13, vendió @0.27 con +$64. Mercado luego revirtió a NO (gamma muestra 0.034/0.966). El bot **no se equivocó** — capturó volatilidad legítima. Tomó beneficio en tope técnico, no esperó resolución.
- **Math accounting**: error máximo $0.02 en 5 trades (rounding floating-point). Sin desviación sistemática.
- **Side vs outcome**: en 4 de 5, bot acertó dirección (YES correcto → ganó, YES erróneo → perdió). En trade 1 hizo timing arbitrage. **Cero casos de "bot gana cuando debió perder"**.

### Fuente primaria utilizada

- `https://gamma-api.polymarket.com/markets/<id>` — sin auth, read-only, retorna `outcomePrices` con probabilidad final implícita

---

## Sección 5 — Bugs encontrados

| # | Severidad | Bug | Archivo | Evidencia |
|---|---|---|---|---|
| BUG-1 | CRÍTICO | `force_close_horizon_excess` no escribe a memory.md | `bot/exit_monitor/cmd/force_close_horizon_excess/main.go` | grep `loglosses\|LogClose` → 0 matches; 21 cierres @ 15:36 ausentes de memory.md |
| BUG-2 | MENOR | `entry_timestamp` null en closed.jsonl | desconocido (escritura monitor o closer) | Dificulta cálculo de lifecycle. `closed.jsonl` muestra `"entry_timestamp":null` |
| WARN-1 | INFO | Sección Vetos tiene 501 filas (cap 500 off-by-one) | `bot/brain/internal/memory/memory.go:46` | `awk` count returned 501 |
| WARN-2 | INFO | Header "## Anti-patterns identificados" mezcla contenido inline | `memory.md:854` | Visual inspection |

### Reproducción de BUG-1

```bash
ssh agent@88.198.168.61 'cd /home/agent/agent-stack && \
  # Trades que se cerraron por force_close a 15:36 UTC:
  ~/bin/jq -r "select(.exit_reason==\"horizon-quota-rebalance-2026-05-26\") | .id" \
    vault/inbox/trading/closed.jsonl | head -5; \
  echo "---"; \
  # ¿Aparecen en memory.md losses?
  for id in $(~/bin/jq -r "select(.exit_reason==\"horizon-quota-rebalance-2026-05-26\") | .id" vault/inbox/trading/closed.jsonl | head -5); do \
    grep -c "$id" vault/agents/polymarket-bot/memory.md; \
  done'
# Esperado tras fix: ≥1 match cada uno. Actualmente: 0 matches.
```

---

## Sección 6 — Acciones recomendadas (priorizadas)

| Prioridad | Acción | Esfuerzo | Impacto |
|---|---|---|---|
| **P0** | Fix BUG-1: añadir `loglosses.LogClose` a `force_close_horizon_excess` | 10 min (8 líneas + rebuild) | Loop completo, ya no se pierde aprendizaje en rebalances futuros |
| **P1** | Decidir cuota realista (30/35/35 o reordenar scanner) | 30 min análisis + 5 min config | Distribución estabiliza, no se queda en frustración 0/50/50 |
| **P1** | Fix BUG-2: propagar `entry_timestamp` a closed records | 15 min | Habilita cálculo de lifecycle, mejor weekly review |
| **P2** | Próximo weekly_review @ Dom 19:00: confirmar M1 sigue subiendo | 0 min (auto) | Validar que memoria sigue creciendo en utilidad |
| **P3** | WARN-2: limpiar header Anti-patterns en memory.md | 2 min manual | Cosmético |

### Conclusión global

El sistema **funciona como diseñado** en sus componentes principales (loop M1+M2+M3, contabilidad, outcomes). Tiene UN bug crítico aislado (force_close skip de memory write) y UN problema de diseño (cuota imposible dado el universo real). Ambos arreglables en <1 hora.

**No hay evidencia de comportamiento engañoso, alucinación, o cálculo erróneo**. El bot ha perdido $924 porque las apuestas perdieron — no porque el contador esté roto.

El próximo `weekly_review.sh` Sun 19:00 ya generará la siguiente capa de análisis. Si M1 sigue subiendo a 10+ y P&L deja de sangrar, el sistema está sano.

---

## Anexo — Comandos para re-auditar

```bash
# Loop integrity
ssh agent@88.198.168.61 "tail -5 /home/agent/agent-stack/vault/agents/polymarket-bot/memory.md"
ssh agent@88.198.168.61 "tail -10 /home/agent/agent-stack/vault/inbox/trading/cron-pipeline.log | grep 'Brain v3 done'"

# Distribución actual
ssh agent@88.198.168.61 "python3 -c 'import json; from collections import Counter; print(Counter(json.loads(l)[\"horizon\"] for l in open(\"/home/agent/agent-stack/vault/inbox/trading/active.jsonl\")))'"

# Sample 5 trades random + verify
ssh agent@88.198.168.61 "~/bin/jq -c \"select(.exit_reason != \\\"horizon-quota-rebalance-2026-05-26\\\")\" /home/agent/agent-stack/vault/inbox/trading/closed.jsonl | shuf -n 5"
# Luego: curl https://gamma-api.polymarket.com/markets/<market_id> para cada uno
```
