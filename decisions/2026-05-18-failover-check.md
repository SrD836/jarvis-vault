---
title: "Failover chain audit — 2026-05-18"
date: 2026-05-18
related:
  - "[[03-decisions/2026-05-13-budget-guardian]]"
  - "[[03-decisions/2026-05-14-failover-chain-check]]"
  - "[[agents/monitor]]"
  - "[[decisions/2026-05-18-config-edits]]"
  - "[[inbox/integrations/multi-provider-roadmap]]"
tags: [infra, runtime, ollama, failover, audit]
status: reviewed
# auto-linked 2026-05-18
---


# Failover chain audit — 2026-05-18

## 1. Config actual (`agents.defaults.model`)

Extraído de `~/.openclaw/openclaw.json` (líneas 42–47):

```json
"agents": {
  "defaults": {
    "model": {
      "primary": "anthropic/claude-sonnet-4-6",
      "fallbacks": []
    }
```

**Estado:** `fallbacks` está **vacío**. No hay cadena de failover configurada. Si el modelo primario (claude-sonnet-4-6 vía claude-cli) no está disponible, el runtime no tiene alternativa declarada.

Nota adicional: en `agents.defaults.models` están registrados `anthropic/claude-sonnet-4-6` y `anthropic/claude-opus-4-7`, ambos usando `agentRuntime.id = "claude-cli"`. El proveedor ollama-local aparece solo en `models.providers` y en `tools.byProvider` (para restricciones de herramientas), pero **no está en la lista de fallbacks ni en `agents.defaults.models`**.

---

## 2. Test de Ollama — qwen2.5-coder:7b-instruct

**Comando ejecutado:**
```
docker exec openclaw-fork-ollama-1 ollama run qwen2.5-coder:7b-instruct "responde solo: pong"
```

**Output literal (sin escapes ANSI):**
```
Pong
```

**Resultado:** ✅ El modelo responde. El contenedor `openclaw-fork-ollama-1` está activo y el modelo `qwen2.5-coder:7b-instruct` carga y contesta correctamente.

---

## 3. Propuesta de fallback (NO aplicada — pendiente aprobación de David)

Para activar el fallover a ollama-local cuando claude-cli no esté disponible, añadir en `~/.openclaw/openclaw.json` dentro de `agents.defaults.model`:

```json
"model": {
  "primary": "anthropic/claude-sonnet-4-6",
  "fallbacks": [
    "ollama-local/qwen2.5-coder:7b-instruct"
  ]
}
```

Y registrar el modelo en `agents.defaults.models` para que el runtime sepa qué runtime usar:

```json
"models": {
  "anthropic/claude-sonnet-4-6": {
    "agentRuntime": { "id": "claude-cli" }
  },
  "anthropic/claude-opus-4-7": {
    "agentRuntime": { "id": "claude-cli" }
  },
  "ollama-local/qwen2.5-coder:7b-instruct": {
    "agentRuntime": { "id": "ollama-local" }
  }
}
```

**Restricciones ya en vigor:** `tools.byProvider` ya deniega `group:web` y `browser` para este modelo. Esa restricción se mantendría tras el failover — correcto.

**Consideraciones antes de aplicar:**
- Confirmar que `agentRuntime.id = "ollama-local"` es el identificador correcto para el runtime ollama en esta versión de OpenClaw (verificar docs o `openclaw gateway config.schema.lookup`).
- El modelo local es 7B — rendimiento y ventana de contexto muy inferiores a Sonnet. Útil como fallback para tareas triviales; no adecuado para planning complejo.
- Si se activa, actualizar `AGENTS.md` (sección "Routing de modelos") para reflejar que el fallback automático ya está declarado en config.

---

## Resumen ejecutivo

| Check | Estado |
|---|---|
| primary configurado | ✅ `anthropic/claude-sonnet-4-6` |
| fallbacks configurados | ❌ lista vacía |
| ollama container activo | ✅ `openclaw-fork-ollama-1` up |
| qwen2.5-coder:7b responde | ✅ output: "Pong" |
| propuesta redactada | ✅ (pendiente aprobación) |
