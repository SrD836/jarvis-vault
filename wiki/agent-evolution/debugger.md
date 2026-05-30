---
title: "Evolution proposal — debugger"
type: agent-evolution
date: 2026-05-30T04:00:38+00:00
agent: "[[agents/debugger]]"
pattern: abort-recurrent
confidence: medium
proposed: true
applied: false
tags: [agent-evolution, debugger, abort-recurrent]
related:
  - "[[agents/debugger]]"
  - "[[00-MOC]]"
---

# Evolución propuesta — `debugger`

_Generado por `hermes/learnings.py` el 2026-05-30T04:00:38+00:00. Ventana: 7 días._

## Patrón: `abort-recurrent`

- Confianza: **medium**
- Ocurrencias: **3** runs
- `abort_reason`: `success_criterion no cumplido`

### Evidencia

- [[agents/debugger/runs/2026-05-29/152813-from-planner]]
- [[agents/debugger/runs/2026-05-29/152826-from-planner]]
- [[agents/debugger/runs/2026-05-29/152800-from-planner]]

### Propuesta

## Diagnóstico

Los tres runs comparten tres señales que apuntan al mismo fallo, no a tres distintos:

1. **`model: deepseek-chat` / `provider: deepseek`** — pero el briefing declara `model_primary: anthropic/claude-opus-4-8`. El agente NO corre en el modelo configurado; hay un override de routing que lo manda a deepseek-chat.
2. **`tool_calls_n: 0` en los tres** — el debugger nunca investiga: emite texto y aborta en `iter: 1`. Un debugger con cero tool_calls jamás puede cumplir un `success_criterion` que exige reproducir/verificar.
3. Tokens minúsculos (46–93 in) → tareas que ni se intentan.

La causa raíz más probable: modelo barato sin disciplina de herramientas + briefing sin "definition of done", así que ante la primera duda aborta en vez de reportar hallazgos.

---

## Cambios propuestos

### 1. Pinear el modelo (corregir el routing a deepseek)
**(a)** `openclaw.json` (entrada del agente `debugger`)
**(b)**
```diff
   "debugger": {
     "agent_id": "debugger",
-    "model": "deepseek-chat",
+    "model": "anthropic/claude-opus-4-8",
+    "model_fallback": "anthropic/claude-sonnet-4-6",
     "delegation_mode": "suggest"
   }
```
**(c)** Los runs ejecutan en deepseek-chat pese a `model_primary: opus`; alinear evita el modelo que aborta sin usar tools.

> Si el override no está en `openclaw.json` sino en una regla de routing/coste, indícame ese archivo: **sin acceso a la config de routing no puedo garantizar el path exacto** (ver "Datos que faltan").

### 2. Añadir contrato de ejecución al briefing (anti-abort vacío)
**(a)** `agents/debugger/index.md` (dentro de Human notes, fuera del bloque que el cron sobreescribe)
**(b)**
```markdown
### Contrato de debugging (obligatorio antes de abortar)

- PROHIBIDO cerrar con `tool_calls_n: 0`. Antes de cualquier conclusión:
  reproducir → localizar (grep/read) → hipótesis → fix → verificar.
- `success_criterion no cumplido` NO es razón de abort si no investigaste.
  Si tras ≥1 ciclo de tools no se cumple → `aborted: false` + reportar
  hallazgos parciales y siguiente paso, no abortar en seco.
- Abort solo permitido si: falta acceso/credenciales, o tarea ambigua
  irreversible. En ese caso, `abort_reason` debe nombrar el blocker concreto.
```
**(c)** El agente aborta en iter 1 sin tocar ninguna herramienta; el contrato fuerza investigación y convierte "no cumplido" en reporte útil.

### 3. Cap de iteraciones mínimas / fail-on-empty
**(a)** caps env del agente (`openclaw.json` → `env` o `caps`)
**(b)**
```diff
   "debugger": {
+    "min_tool_calls_before_abort": 1,
+    "max_iters": 6
   }
```
**(c)** Bloquea el patrón "0 tools → abort"; obliga al menos un intento real antes de rendirse.

---

## Si no hay fix obvio
Hay propuesta clara para (2) y (3). Para (1) la **confianza es media**: no veo `openclaw.json` ni dónde se origina el override a deepseek-chat.

**Datos que faltarían:** (a) `openclaw.json` completo del bloque `debugger` + reglas de routing/coste; (b) el cuerpo (no solo frontmatter) de un run abortado, para ver si el `success_criterion` venía mal definido desde `planner`; (c) si `min_tool_calls_before_abort` es una cap soportada por el runtime o requiere hook.

## Apply

Para aplicar: editar manualmente los paths sugeridos, marcar `applied: true` en el frontmatter, y commitear el cambio.

## Human notes

_(añade comentarios manuales aquí; este bloque se preserva)_
