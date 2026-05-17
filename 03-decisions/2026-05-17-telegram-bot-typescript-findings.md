---
title: "TS Bot Telegram — diagnóstico inicial bug formato + plan multi-sesión"
type: decision
date: 2026-05-17
tags: [decision, telegram, openclaw, typescript, bug, multi-sesion]
related:
  - "[[PRD]]"
  - "[[agents/main]]"
  - "[[TOOLS.md]]"
status: open
---

# TS Bot Telegram — diagnóstico inicial

## Contexto

Sesión 2026-05-17 ~10:00–11:15 CET. Tarea: mejorar formato de Telegram (asteriscos literales `**foo**` visibles + falta copy-button en code blocks) y añadir multi-sesión con comandos `/new /list /resume`.

**Discovery clave** (antes asumido erróneamente):

El bot productivo de JARVIS Telegram NO es `gateway/platforms/telegram.py` (Python, en `C:\Users\david\agent-stack\temphermes-agent\`). Ese es un proyecto paralelo/staging **no deployado**.

El bot real es **TypeScript** dentro del contenedor `openclaw-fork-openclaw-gateway-1` (imagen `openclaw:local`). Código fuente en VPS:

```
/home/agent/agent-stack/openclaw-fork/extensions/telegram/src/
├── send.ts                       (1740 líneas, parse_mode: "HTML" en líneas 692/853/1392)
├── format.ts                     (684 líneas, contiene markdownToTelegramHtml)
├── outbound-adapter.ts           (textMode: "html" si parseMode==="HTML")
├── bot/delivery.send.ts          (línea 114: const htmlText = textMode === "html" ? text : markdownToTelegramHtml(text))
└── bot/delivery.replies.ts       (líneas 202, 243, 298: pasa chunk.html con textMode: "html")
```

Snapshot de los 4 archivos clave copiado a `C:\Users\david\agent-stack\openclaw-fork-snapshot\` para análisis offline en próxima sesión.

## Pipeline actual del bot

1. Modelo (Claude/DeepSeek) emite respuesta — formato esperado: **markdown estándar**.
2. `delivery.replies.ts` chunkea el texto vía `markdownToTelegramChunks(markdown, limit)` → array `{html, text}[]`.
3. `format.ts:131 markdownToTelegramHtml()`:
   - parsea con `markdownToIR()` (parser openclaw)
   - renderiza con `renderTelegramHtml(ir)`
   - normaliza con `preserveSupportedTelegramHtmlTags()`
   - envuelve refs a archivos en `<code>` con `wrapFileReferencesInHtml()`
4. `delivery.replies.ts:202` envía cada `chunk.html` con `textMode: "html"` a `sendTelegramText()`.
5. `delivery.send.ts:114`: `const htmlText = textMode === "html" ? text : markdownToTelegramHtml(text)`. **Si textMode === "html", asume que text ya es HTML válido → NO convierte.**
6. `delivery.send.ts:152`: `bot.api.sendMessage(chatId, htmlText, { parse_mode: "HTML", ... })`.

## Hipótesis del bug "asteriscos literales"

Tres candidatas, no descartadas, ordenadas por probabilidad:

### H1 — Markdown roto del modelo cae a passthrough
Cuando el LLM emite markdown con marcado mal balanceado (`**foo` sin cerrar, `*foo**bar*`, etc.), `markdownToIR()` lo trata como texto raw (sin emit de `<b>`). El render preserva los `**` literales en el HTML output. Telegram HTML mode los muestra tal cual.

**Evidencia**: el output del modelo en la captura — `**Para confirmar...**` justo seguido de un guión y otra línea — es exactamente el tipo de markdown que un parser estricto rechaza si el siguiente token no es un word boundary limpio.

**Fix candidato**: añadir un sanitizer post-render que detecte `\*\*([^\n*]+)\*\*` residual y lo convierta a `<b>$1</b>`, o lo escape como `&#42;&#42;...`.

### H2 — Caller con textMode="html" mete markdown crudo
Algún caller del bot llama `sendTelegramText(..., text_con_asteriscos, ..., textMode: "html")` pensando que es HTML. Como `delivery.send.ts:114` no convierte cuando `textMode === "html"`, los asteriscos sobreviven.

**Candidatos a auditar**:
- `delivery.replies.ts` (líneas 202, 243, 298) — pasan `chunk.html`; verificar que `chunk.html` viene del chunker correcto y no de un branch que lo deja como markdown raw.
- `outbound-adapter.ts:74-93` — `textMode: "html"` se establece si `parseMode === "HTML"`. El caller del adapter podría estar marcando como HTML un texto markdown.

**Fix candidato**: hacer `delivery.send.ts:114` defensive — incluso con `textMode: "html"`, pasar `text` por un sanitizer que convierta `**...**` y `*...*` residuales a tags HTML, o que los escape.

### H3 — Tabla / lista con markdown anidado
`markdownToTelegramHtml` con `tableMode` puede generar HTML donde celdas internas conserven `**` si el parser de IR no profundiza en cells. Menor probabilidad — más fácil de descartar testeando.

## Estrategia de fix recomendada (próxima sesión)

**Cambio mínimo (~30 líneas TS) en `format.ts`**:

Después de `preserveSupportedTelegramHtmlTags(html)` en `markdownToTelegramHtml`, añadir paso de "residual markdown sanitization":

```typescript
function sanitizeResidualMarkdown(html: string): string {
  // Skip content inside <code>, <pre>, <a> to avoid double-conversion.
  // Convert leftover **bold** and *italic* to HTML tags.
  // Escape lonely asterisks that can't form a pair.
}
```

**Y** edit defensive en `bot/delivery.send.ts:114`:
```typescript
const htmlText = textMode === "html" 
  ? sanitizeResidualMarkdown(text)  // <-- nuevo
  : markdownToTelegramHtml(text);
```

**Tests TS** a añadir en `format.test.ts` (debe existir, no verificado):
- input `"**foo**bar*baz unclosed"` → output sin asteriscos raw, con `<b>foo</b>` o equivalente
- input con fenced code y bold mezclado → fenced intacto, bold convertido

## Multi-sesión (`/new` `/list` `/resume`)

Toda la lógica debe ir en TypeScript también. Archivos candidatos a tocar:

- `bot-native-commands.ts` (línea 567+, 871+, 1094+, 1267+, 1330+) — patrón de command handler ya existente.
- Nuevo archivo `sessions.ts` con SQLite manager (better-sqlite3, ya disponible en el openclaw image).
- El vault Obsidian ya está montado en el container — escritura directa a `/vault/02-sessions/{YYYY-MM-DD}/{HHMMSS}-{slug}.md`.

## Deployment (cuándo se aplique el fix)

```bash
ssh agent@88.198.168.61
cd /home/agent/agent-stack/openclaw-fork
# Editar los archivos TS
docker compose build openclaw-gateway
docker compose up -d openclaw-gateway
sleep 20 && docker ps | grep openclaw-gateway   # → debe estar Up healthy
docker logs openclaw-fork-openclaw-gateway-1 --tail 50  # vigilar errores TS al arrancar
```

**Riesgo**: build de imagen openclaw:local toma 5–10 min, si falla rollback requiere `docker tag <prev-sha> openclaw:local && docker compose up -d`. Tener tag de la imagen actual antes del rebuild.

## Estado tras esta sesión

- ✅ Heurística de delegación añadida a `agents/main.md` dentro de `## Human notes` (cron-safe).
- ✅ Gateway restart healthy.
- ⚠️ Python `gateway/platforms/telegram.py` edits **revertidos** — eran en codebase no productivo.
- ✅ TS Telegram fix **aplicado**: `sanitizeResidualMarkdown()` añadido a `format.ts` (idempotente, escapa lonely `**` a `&#42;&#42;`, convierte pares residuales a `<b>`). Defensive call también en `delivery.send.ts:114` para path con `textMode: "html"`. 36/36 tests format+repro pass; suite extension-telegram pasa 1583/1585 (los 2 fallos son flake pre-existente confirmado revertiendo a HEAD). Test nuevo: `format.residual-asterisk.test.ts` con 9 casos del bug. Imagen `openclaw:local-pre-mdfix-1779013723` tagged como rollback. **Swap 8GB activado por David (sudo fallocate)**, rebuild de imagen en progreso (~10 min con swap).
- ⏳ Multi-sesión **código TS escrito en local snapshot** (`session-manager.ts`, `session-export.ts`, `session-commands.ts`, integración en `bot-core.ts`). Pendiente: subir al VPS y rebuild en un segundo ciclo después de validar formato fix.

## Comandos rollback fix formato (si necesario)

```bash
ssh agent@88.198.168.61
docker tag openclaw:local-pre-mdfix-1779013723 openclaw:local
cd /home/agent/agent-stack/openclaw-fork && docker compose up -d openclaw-gateway
```

## Próxima sesión — orden recomendado

1. Abrir `C:\Users\david\agent-stack\openclaw-fork-snapshot\` y leer `format.ts`, `delivery.send.ts` con calma.
2. Decidir entre H1, H2, H3 (instrumentando un log para capturar `text` justo antes del `bot.api.sendMessage`).
3. Implementar `sanitizeResidualMarkdown` y los defensive edits.
4. Validar con tests TS locales (probablemente `npm test` dentro de `extensions/telegram/`).
5. Deploy iterativo.
6. Tras formato OK, abordar multi-sesión.

Memorias relacionadas:
- `[[jarvis_telegram_bot_is_typescript]]` — confirmación del pivote
- `[[jarvis_main_md_regenerated_by_cron]]` — por qué la directiva fue a Human notes
