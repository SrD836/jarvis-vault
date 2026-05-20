---
title: "AGENT_RULES — Reglas obligatorias"
type: rules
date: 2026-05-14
tags: [rules, agent, safety, jarvis]
related:
  - "[[PRD]]"
  - "[[TASKS]]"
  - "[[00-MOC]]"
---

# 🛡️ AGENT_RULES — Reglas obligatorias para agentes autónomos

> Complementan las reglas de seguridad en `.agent/Rules.md`.

---

## 1. Seguridad y estabilidad

- **Nunca romper funcionalidades existentes.** Verificar antes de modificar.
- **No eliminar archivos críticos.** NUNCA borrar `seeds/`, `templates/`, `00-MOC.md`, `.agent/Rules.md`.
- **NUNCA modificar** bloques `## Human notes`.
- **Crear backups antes de refactors** que afecten >3 archivos.
- **Mantener vault navegable.** MOC actualizado, wikilinks válidos.

## 2. Calidad de ejecución

- **Cambios pequeños y verificables.** Máximo 1 concepto por iteración.
- **Validar antes de continuar:** frontmatter YAML, wikilinks, consistencia de estilo.
- **Corregir errores antes de añadir features.**
- **Prioridades:** 1) estabilidad 2) precisión 3) mantenibilidad 4) claridad 5) velocidad.

## 3. Estilo y convenciones

- Frontmatter YAML obligatorio: `title`, `type`, `date`, `tags`, `related`.
- Wikilinks `[[así]]` para referencias internas.
- Idioma: español para contenido, inglés para código/config.
- Documentar decisiones importantes en `03-decisions/`.

## 4. Flujo de trabajo autónomo

```
1. Leer LOOP_STATE.json → identificar tarea activa
2. Leer TASKS.md → verificar tarea pendiente
3. Planificar cambio mínimo
4. Ejecutar cambio
5. Validar resultado
6. Actualizar progress.txt + LOOP_STATE.json
7. Marcar tarea en TASKS.md
8. Siguiente tarea o finalizar iteración
```

### Límites
- Máximo **10 iteraciones** por sesión Ralph Loop.
- Tarea no completada en 3 intentos → marcar bloqueada.
- >3 errores consecutivos → pausar y reportar.

### Recuperación
- Error recuperable → 1 intento de fix automático.
- Fix falla → documentar, marcar bloqueada.
- Error de estabilidad → revertir, pausar loop.

### Anti-drift
- Cada 3 iteraciones, releer `PRD.md` y `TASKS.md`.
- No crear ramas sin registrar en TASKS.md.

## 5. Protección del proyecto

**Archivos protegidos:** `seeds/**`, `templates/**`, `.agent/Rules.md`, `00-MOC.md`.

**Zonas de escritura segura:** `01-briefs/`, `02-sessions/`, `03-decisions/`, `04-skills-log/`, `wiki/`, `projects/`, `agents/`, `skills/`, `memories/`, ``.

## 6. Comunicación

- Minimizar preguntas innecesarias.
- Explicar brevemente decisiones importantes.
- Proponer soluciones, no solo problemas.

---

Ver también: [[PRD]] | [[TASKS]] | [[.agent/Rules]]
