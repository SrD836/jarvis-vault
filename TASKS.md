---
title: "TASKS — Checklist ejecutable"
type: tasks
date: 2026-05-14
tags: [tasks, checklist, jarvis, ralph-loop]
related:
  - "[[PRD]]"
  - "[[AGENT_RULES]]"
  - "[[00-MOC]]"
---

# ✅ TASKS — Checklist ejecutable JARVIS

> Derivado del [[PRD]]. Actualizado por agentes autónomos en cada iteración.
> Formato: `[ ]` pendiente | `[x]` completado | `[~]` en progreso

---

## Fase 0: Bootstrap entorno autónomo

- [x] Analizar estructura completa del proyecto
- [x] Detectar stack, frameworks, dependencias y arquitectura
- [x] Crear PRD.md
- [x] Crear progress.txt
- [x] Crear TASKS.md (este archivo)
- [x] Crear AGENT_RULES.md
- [x] Crear LOOP_STATE.json
- [x] Crear .antigravity/config.json
- [x] Crear .antigravity/memory.md
- [x] Configurar Ralph Loop
- [x] Configurar Supersmooth
- [x] Validar configuración completa (smoke test)

---

## Fase 1: Estabilización 🔴

### F1.1 — Verificar session-export
- [ ] Revisar logs del cron session-export en VPS
- [ ] Verificar formato esperado de sesiones
- [ ] Probar exportar una sesión manualmente
- [x] Documentar resultado en 03-decisions/

### F1.2 — Completar archivos de identidad
- [x] Crear IDENTITY.md
- [x] Crear USER.md
- [x] Crear TOOLS.md
- [x] Verificar que agente main los referencia

### F1.3 — Investigar duplicación skill-logs
- [ ] Revisar curator.py lógica de deduplicación
- [ ] Comparar hashes 2026-05-13 vs 2026-05-14
- [x] Implementar guard de duplicación si es bug (Análisis completado en 03-decisions/)

### F1.4 — Validar memory-wiki bridge
- [ ] Verificar configuración memory-wiki en OpenClaw
- [ ] Confirmar modo del bridge
- [ ] Probar inyección manual de concepto
- [ ] Verificar aparición en wiki/concepts/
- [x] Documentar estado en 03-decisions/

### F1.5 — Confirmar failover chain
- [ ] Verificar configuración failover en OpenClaw
- [ ] Probar fallback a ollama-qwen
- [x] Documentar tiempos de recovery (Análisis en 03-decisions/)

---

## Fase 1.5: Despliegue Hermes Upgrade (Acciones VPS) 🚀

### F1.5.1 — Desplegar Consolidación de Memoria
- [x] Mover `consolidate.py` a `~/.openclaw/scripts/` en el VPS
- [x] Dar permisos de ejecución (`chmod +x`)
- [x] Configurar cron job diario a las 04:30

### F1.5.2 — Parchear Curator
- [x] Integrar deduplicación de `curator_patch.py` en el `curator.py` real
- [x] Añadir lógica de refinamiento de skills

### F1.5.3 — Reparar Memory-Wiki
- [~] Modificar `openclaw.json` para modo `active` *(Fallido: valor no soportado)*
- [x] Mitigación: `consolidate.py` asume el rol de la wiki de manera autónoma.

---

## Fase 2: Memoria y Contexto 🟡

### F2.1 — Activar memory-wiki bridge
- [ ] Configurar bridge en modo active
- [ ] Verificar auto-generación de entries wiki
- [ ] Monitorear 24h

### F2.2 — Configurar embeddings para vault
- [ ] Verificar nomic-embed-text en Ollama
- [ ] Conectar vault como fuente de embeddings
- [ ] Probar búsqueda semántica

### F2.3 — Consolidación semanal
- [ ] Crear cron weekly-synthesis
- [ ] Crear template de síntesis
- [ ] Agregar a operacion-diaria.md

### F2.4 — Pipeline concept→entity→synthesis
- [ ] Definir reglas de extracción
- [ ] Implementar pipeline
- [ ] Probar con sesión real

---

## Fase 3: Automatización Avanzada 🟢

### F3.1 — Ralph Loop operativo
- [ ] Probar loop completo con tarea simple
- [ ] Verificar checkpoints se generan
- [ ] Verificar recovery ante error simulado
- [ ] Ejecutar 5+ iteraciones sin drift

### F3.2 — Delegación multi-agente
- [ ] Probar planner→code-reviewer
- [ ] Probar planner→researcher
- [ ] Probar planner→documenter
- [ ] Documentar patrones exitosos

### F3.3 — Goal tracking multi-sesión
- [ ] Diseñar esquema de goals
- [ ] Implementar tracking cross-session
- [ ] Integrar con vault

### F3.4 — Skills de auto-diagnóstico
- [ ] Skill: health-check sistema
- [ ] Skill: vault-integrity
- [ ] Skill: cron-status

---

## Fase 4: Inteligencia 🔵

### F4.1 — Skill auto-refinement
- [ ] Analizar patrones de uso de skills
- [ ] Implementar refinement basado en feedback
- [ ] Probar con skill existente

### F4.2 — Detección proactiva de tareas
- [ ] Implementar scanner de oportunidades
- [ ] Integrar con daily-brief

### F4.3 — Resúmenes ejecutivos con insights
- [ ] Mejorar template daily-brief con insights
- [ ] Agregar métricas de tendencia

### F4.4 — Dashboard de salud
- [ ] Diseñar métricas clave
- [ ] Implementar generación automática
- [ ] Agregar a MOC

---

## Registro de cambios

| Fecha | Cambio |
|---|---|
| 2026-05-14 | Creación inicial del checklist desde PRD |
