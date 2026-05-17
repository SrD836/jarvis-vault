---
title: "Plan v6 — Evolución de JARVIS: memoria viva, multi-agente afilado, segundo cerebro"
type: decision
date: 2026-05-14
decision: "Adoptar Plan v6 como roadmap de la próxima evolución de JARVIS (memoria semántica activa + caps duros + destilación de skills + consolidación episódica + middleware de seguridad)"
alternatives:
  - "Continuar solo con Plan v5 (estado actual + Ralph Loop)"
  - "Saltar directo a integración Antigravity sin estabilizar memoria primero"
outcome: pending
outcome_observed_after_days: 30
tags: [decision, roadmap, plan-v6, memoria, multi-agente, antigravity]
related:
  - "[[../PRD]]"
  - "[[../TASKS]]"
  - "[[../00-MOC]]"
  - "[[../projects/guia-fases-pendientes]]"
---

# Plan v6 — Evolución de JARVIS: memoria viva, multi-agente afilado y exploración del segundo cerebro

> Movido desde `Desktop\Plan v6 — Evolución de JARVIS memor.txt` el 2026-05-16. Texto original conservado íntegro abajo.

## Objetivo general

Convertir JARVIS en un agente que realmente aprende de su experiencia, consolida conocimiento en el vault Obsidian, utiliza el multi-agente con herramientas especializadas, y se controla con límites estrictos. Además, se integra la visualización del grafo de conocimiento con Antigravity para que David pueda navegar su segundo cerebro de manera inmersiva.

Las fases son acumulativas y respetan la infraestructura existente (Docker Compose, Traefik, Syncthing, contenedores actuales). Cada fase puede probarse de forma independiente.

---

## Fase 6.1 — Memoria semántica activa en cada interacción

**Objetivo:** Que el agente principal recupere automáticamente información relevante del vault al inicio de cada conversación o meta, inyectándola en el contexto para que "recuerde" sin que el usuario tenga que mencionarlo.

**Comportamiento esperado:**

- Antes de que el modelo genere una respuesta, el sistema busca en los últimos 30 días de contenido del vault (briefs, sesiones, decisiones, wiki) aquellos fragmentos semánticamente cercanos a la consulta del usuario o al objetivo actual.
- Se extraen los fragmentos más relevantes y se añaden al mensaje del sistema como "Recuerdos contextuales recientes".
- El índice vectorial se actualiza automáticamente cada día (cron) para incluir el nuevo contenido.

**Piezas involucradas:** memory-core (embeddings nomic), cron de indexación, hook en el gateway de OpenClaw, templates de prompt del agente main.

**Entregable:** El agente comenzará a mostrar memoria contextual en sus respuestas sin intervención manual.

---

## Fase 6.2 — Máquina de estados para metas y caps estrictos

**Objetivo:** Reforzar el control de ejecución de metas con una máquina de estados y límites duros (no solo textuales) que cancelen una meta si se exceden iteraciones, tokens o tiempo, notificando a David por Telegram.

**Comportamiento esperado:**

- Toda meta pasa por los estados: `planning` → `approved` (por David o automático para tareas seguras) → `executing` → `verifying` → `done`.
- Un monitor en el gateway cuenta iteraciones, tokens consumidos y tiempo transcurrido. Si se alcanza un límite, la meta se detiene, se guarda un registro en `vault/03-decisions/` y se envía una alerta a Telegram.
- El sistema preserva el trabajo parcial y permite reanudar si David lo autoriza.

**Piezas involucradas:** `openclaw.json` (nuevo esquema de metas), lógica de control en el gateway, cron de heartbeats y alertas.

**Entregable:** Metas más seguras y transparentes, con frenos automáticos.

---

## Fase 6.3 — Destilación y fusión semántica de skills

**Objetivo:** Mantener el catálogo de skills limpio y útil, fusionando automáticamente aquellas que son redundantes o muy similares, con aprobación final de David.

**Comportamiento esperado:**

- Una vez por semana, el sistema analiza todas las skills activas. Usando embeddings y un modelo local (Qwen o DeepSeek), detecta pares con alta similitud semántica.
- Para cada par candidato, se genera una skill unificada (propuesta) y se notifica a David por Telegram con un resumen y un comando para aprobar o rechazar.
- Las skills originales se marcan como `archived` con una referencia a la unificada una vez aprobada.

**Piezas involucradas:** curator-weekly (cron), `curator.py`, skill-workshop (modo propuesta), notificaciones Telegram.

**Entregable:** Catálogo de skills que no se degrada con el tiempo.

---

## Fase 6.4 — Consolidación episódica: el "Hermes completo"

**Objetivo:** Que el sistema recorra semanalmente sus sesiones y briefs, extraiga aprendizajes, actualice skills, enriquezca el wiki y genere resúmenes, todo reflejado en el vault Obsidian.

**Comportamiento esperado:**

- **Recolección:** cada domingo, un nuevo worker (`memory-consolidator`) examina las notas de la semana en `02-sessions/`, `01-briefs/` y las decisiones recientes.
- **Extracción:** identifica comandos exitosos, flujos nuevos, errores recurrentes y conceptos técnicos.
- **Actualización de skills:**
  - Si detecta patrones de éxito repetidos (mínimo 2 sesiones), genera una propuesta de skill en `state: proposed` y notifica a David.
  - Las skills con `failure_streak` alto se marcan `probation` y se sugiere revisión.
- **Enriquecimiento del wiki:**
  - Los conceptos extraídos se vinculan a notas existentes en `wiki/` usando embeddings. Si no existe nota, se crea un esbozo con enlaces a las sesiones fuente.
  - Las seeds (`sistema-arquitectura.md`, etc.) reciben una actualización en una sección "Lecciones recientes" automática, sin modificar el contenido humano.
- **Resumen semanal:** se genera `weekly-brief.md` en `01-briefs/` y se envía un resumen de 5 líneas a Telegram.

**Piezas involucradas:** nuevo worker `memory-consolidator`, scripts de acceso al vault, modelos locales para extracción, cron semanal.

**Entregable:** El vault de Obsidian comienza a crecer con conocimiento consolidado automáticamente, y las skills evolucionan a partir de la experiencia real.

---

## Fase 6.5 — Especialización de roles y herramientas del sistema multi-agente

**Objetivo:** Asignar herramientas específicas a cada tipo de agente (planner, researcher, code-reviewer, documenter) y formalizar el paso de mensajes entre ellos para mejorar la colaboración.

**Comportamiento esperado:**

- En `openclaw.json` se restringen las tools disponibles para cada rol. Ejemplo: `researcher` puede usar `web_search` y `read_file`; `code-reviewer` puede ejecutar shell en modo lectura; `documenter` puede escribir en el vault.
- Cuando el planner delega en un worker, le pasa un contexto estructurado con la tarea, restricciones y el historial relevante. El worker devuelve un resumen con los hallazgos o el código modificado.
- Se impide que un worker invoque herramientas no autorizadas; el gateway bloquea dichas llamadas y lo registra.

**Piezas involucradas:** configuración de agentes en `openclaw.json`, middleware de autorización de tools, plantillas de handoff.

**Entregable:** Agentes más seguros y efectivos, cada uno con responsabilidades claras.

---

## Fase 6.6 — Graph of Thoughts para decisiones complejas

**Objetivo:** Dotar al planner de la capacidad de razonar mediante un grafo de pensamientos cuando la tarea es muy abierta, y almacenar ese grafo en el vault como diagrama Mermaid.

**Comportamiento esperado:**

- Para metas marcadas como complejas (o a petición de David), el planner genera un grafo de razonamiento: nodos con ideas, preguntas, hipótesis y dependencias.
- Ese grafo se guarda en `vault/projects/` como una nota con un bloque Mermaid.
- A medida que el agente avanza, puede actualizar el grafo, marcando nodos como resueltos o añadiendo nuevos.
- David puede ver el grafo en Obsidian (con un plugin de Mermaid) o en Antigravity como una constelación de ideas.

**Piezas involucradas:** nuevo tipo de tool `graph_reason`, integración con el sistema de archivos del vault, visualización en Obsidian.

**Entregable:** Mayor transparencia y calidad en las tareas de planificación compleja.

---

## Fase 6.7 — Fallback DeepSeek y mejora de la cadena de modelos

**Objetivo:** Incorporar DeepSeek V3 como paso intermedio antes de Ollama cuando Claude está rate-limited, manteniendo la calidad de respuesta.

**Comportamiento esperado:**

- Se configura DeepSeek V3 con autenticación en el gateway.
- La cadena de fallback se ajusta a: Sonnet → (si rate-limit) DeepSeek → (si falla) Qwen local.
- Las tareas no urgentes (batch) pueden ser desviadas explícitamente a DeepSeek para ahorrar cuota de Claude.

**Piezas involucradas:** configuración de `auth-profiles` en OpenClaw, script de routing en el gateway.

**Entregable:** Robustez mejorada, menor dependencia del modelo local débil.

---

## Fase 6.8 — Middleware de seguridad en las herramientas

**Objetivo:** Añadir una capa de validación previa a la ejecución de ciertas herramientas (shell, http) para bloquear operaciones peligrosas aunque el modelo las intente.

**Comportamiento esperado:**

- Antes de ejecutar un comando shell, un filtro revisa que no contenga patrones prohibidos (`rm -rf /`, `git push --force`, etc.) y que el destino de las peticiones HTTP esté en una lista blanca.
- Cualquier intento bloqueado se registra en `vault/03-decisions/` y se notifica a Telegram.
- Este middleware actúa incluso en sesiones automáticas, como capa adicional a los guardarrailes textuales.

**Piezas involucradas:** middleware en el contenedor gateway, listas negras/blancas configurables, integración con alertas.

**Entregable:** Sistema mucho más resistente a alucinaciones peligrosas.

---

## Fase 6.9 — Auto-crítica de calidad (opcional)

**Objetivo:** Detectar respuestas incoherentes o de baja calidad mediante una revisión automática con un modelo local, y generar alertas si es necesario.

**Comportamiento esperado:**

- Tras ciertas interacciones importantes (por ejemplo, después de ejecutar una skill crítica), un modelo local evalúa la respuesta del agente según criterios de coherencia y exactitud.
- Si la puntuación de calidad es baja, se guarda un registro en `vault/03-decisions/` y se envía una alerta a Telegram para que David pueda revisar.

**Piezas involucradas:** worker de evaluación, modelo local (Qwen), conexión con el log de decisiones.

**Entregable:** Supervisión continua de la calidad sin intervención humana.

---

## Orden recomendado de ejecución

`6.1 → 6.2 → 6.3 → 6.4 → 6.5 → 6.7 → 6.8 → 6.6 → 6.9`

Las fases 6.4 y 6.6 dependen de que 6.1 esté operativa. La 6.8 puede implementarse en cualquier momento.

---

## Outcome observado

Pendiente. Rellenar tras 30 días una vez se haya iniciado al menos F6.1.

---

← [[../00-MOC|Volver al MOC]] | [[../PRD|PRD]] | [[../TASKS|TASKS]]
