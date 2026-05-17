---
title: "TOOLS — Catálogo de Herramientas y Stack"
type: tools
date: 2026-05-14
tags: [tools, stack, jarvis, core]
related:
  - "[[00-MOC]]"
  - "[[IDENTITY]]"
  - "[[USER]]"
  - "[[seeds/sistema-arquitectura]]"
---

# 🛠️ TOOLS — Herramientas de JARVIS

> Catálogo de las herramientas principales y componentes del stack técnico disponibles para la operación autónoma.

---

## Ecosistema de Agentes y Asistentes

### 1. OpenClaw (Orquestador principal)
- **Rol:** Gateway de agentes, ejecución de crons, gestión de plugins.
- **Capacidades:**
  - Delegación multi-agente (`planner` -> `researcher`, `code-reviewer`, etc.)
  - Sandboxing de ejecución.
  - Hooks para endpoints (UI, Telegram).

### 2. Claude Code
- **Rol:** Asistente especializado en código y tareas de infraestructura.
- **Conexión:** Totalmente integrado con el vault de Obsidian. Puede leer y escribir directamente en el vault, analizar logs y ejecutar modificaciones estructurales en los proyectos.
- **Uso estratégico:** Ideal para refactorizaciones profundas, scripting, y análisis del vault que requieran un contexto extenso y herramientas CLI.

## Modelos de Lenguaje

- **Claude Sonnet/Opus 4.6+ (vía Claude Max):** Para tareas de alto nivel de razonamiento, generación de planes y procesamiento profundo de contexto.
- **Ollama (Qwen2.5-Coder 7B):** Modelo local primario para fallback y tareas de baja latencia o triviales.
- **Nomic Embed Text (Ollama):** Generación de embeddings para búsquedas vectoriales.

## Infraestructura y Persistencia

- **Syncthing:** Sincronización bidireccional entre el VPS (Hetzner) y el portátil de David. Esencial para la integridad del Vault.
- **Docker & Docker Compose:** Contenedorización de todos los servicios (Traefik, OpenClaw, Ollama).
- **Obsidian (Vault):** Interfaz humana y almacenamiento del "Segundo Cerebro". Las modificaciones aquí son la memoria persistente lejana compartida por todos los agentes.

## Herramientas de Interfaz (Tools ejecutables por agentes)

- `read_file` / `write_file` / `edit_file`: Manejo del vault y del workspace de OpenClaw.
- `bash` / `run_command`: Ejecución de scripts en el host (sujeto a sandboxing/cap-drop según el agente).
- `telegram_send`: Envío de notificaciones out-of-band a David.
- `web_fetch` / `browser`: Scraping e investigación en vivo.
- `skill-workshop`: Auto-generación de skills basadas en patrones de uso.

## Reglas de uso de Herramientas

1. **Prioridad de la herramienta específica:** Usa la herramienta nativa más especializada para la tarea antes de recurrir a scripts bash genéricos.
2. **Respetar el Vault:** Cualquier cambio persistente importante debe quedar reflejado en Obsidian. Si Claude Code o OpenClaw detectan un procedimiento nuevo, debe documentarse.
3. **No pisarse:** Asegurar que los cronjobs de OpenClaw y las ejecuciones de Claude Code no escriban en el mismo archivo simultáneamente sin validación (e.g. logs duplicados).
