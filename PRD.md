---
title: "PRD — JARVIS Autonomous Agent System"
type: prd
date: 2026-05-14
tags: [prd, roadmap, backlog, jarvis, antigravity]
related:
  - "[[00-MOC]]"
  - "[[seeds/sistema-arquitectura]]"
  - "[[seeds/decisiones-clave]]"
  - "[[TASKS]]"
  - "[[AGENT_RULES]]"
---

# 📋 PRD — JARVIS Autonomous Agent System

> Documento de requisitos del producto para el sistema JARVIS.
> Optimizado para ejecución iterativa por agentes autónomos (Ralph Loop + Supersmooth).

---

## 1. Objetivo general

Construir y mantener un **asistente personal autónomo (JARVIS)** que opera 24/7 en un VPS Hetzner, orquestado por OpenClaw con modelos Claude (Max) y Ollama local, capaz de:

- Gestionar un segundo cerebro en Obsidian (este vault) con sincronización bidireccional
- Ejecutar tareas programadas (crons) sin intervención humana
- Aprender y mejorar automáticamente via skill-workshop
- Mantener memoria persistente (memory-core + memory-wiki + vault)
- Coordinar múltiples agentes especializados (planner, code-reviewer, researcher, documenter, apier, skill-reviewer)
- Proporcionar interfaz via Telegram bot y Control UI web

---

## 2. Roadmap

### Fase 1: Estabilización (actual) — Prioridad CRÍTICA
> Objetivo: Asegurar que toda la infraestructura existente funciona correctamente

| ID | Tarea | Estado | Dependencia |
|---|---|---|---|
| F1.1 | Verificar session-export captura sesiones | pendiente | — |
| F1.2 | Completar IDENTITY.md, USER.md, TOOLS.md | pendiente | — |
| F1.3 | Investigar duplicación skill-logs (curator) | pendiente | — |
| F1.4 | Validar memory-wiki bridge genera artifacts | pendiente | F1.1 |
| F1.5 | Confirmar failover chain Anthropic→Ollama | pendiente | — |

### Fase 2: Memoria y Contexto — Prioridad ALTA
> Objetivo: Memoria persistente funcional y auto-enriquecida

| ID | Tarea | Estado | Dependencia |
|---|---|---|---|
| F2.1 | Activar memory-wiki bridge modo real | pendiente | F1.4 |
| F2.2 | Configurar embeddings nomic para vault | pendiente | F2.1 |
| F2.3 | Implementar consolidación semanal de memorias | pendiente | F2.1 |
| F2.4 | Crear pipeline concept→entity→synthesis | pendiente | F2.1 |

### Fase 3: Automatización Avanzada — Prioridad MEDIA
> Objetivo: Agentes que trabajan sin supervisión en tareas complejas

| ID | Tarea | Estado | Dependencia |
|---|---|---|---|
| F3.1 | Configurar Ralph Loop en el sistema | pendiente | F1.* |
| F3.2 | Activar delegación planner→workers | pendiente | F1.2 |
| F3.3 | Implementar goal tracking multi-sesión | pendiente | F2.* |
| F3.4 | Crear skills de auto-diagnóstico | pendiente | F3.1 |

### Fase 4: Inteligencia — Prioridad BAJA
> Objetivo: Auto-mejora y aprendizaje continuo

| ID | Tarea | Estado | Dependencia |
|---|---|---|---|
| F4.1 | Skill auto-refinement basado en uso | pendiente | F3.4 |
| F4.2 | Detección proactiva de tareas | pendiente | F3.3 |
| F4.3 | Resúmenes ejecutivos con insights | pendiente | F2.3 |
| F4.4 | Dashboard de salud del sistema | pendiente | F3.4 |

---

## 3. Backlog detallado (tareas atómicas)

### 🔴 Prioridad Crítica

#### F1.1 — Verificar session-export
- [ ] Revisar logs del cron session-export en el VPS
- [ ] Verificar que el formato de sesiones coincide con lo esperado
- [ ] Probar exportar una sesión manualmente
- [ ] Documentar resultado en [[03-decisions/]]
- **Criterio de finalización:** Al menos 1 sesión aparece en `02-sessions/` tras interacción

#### F1.2 — Completar archivos de identidad
- [ ] Crear IDENTITY.md con personalidad y directivas del agente
- [ ] Crear USER.md con perfil de David (preferencias, timezone, idioma)
- [ ] Crear TOOLS.md con catálogo de herramientas disponibles
- **Criterio de finalización:** Los 3 archivos existen y son referenciados por el agente main

#### F1.3 — Investigar duplicación skill-logs
- [ ] Revisar curator.py lógica de deduplicación
- [ ] Comparar hashes de 2026-05-13 vs 2026-05-14
- [ ] Implementar guard de duplicación si es bug
- **Criterio de finalización:** No hay entradas duplicadas en `04-skills-log/`

#### F1.4 — Validar memory-wiki bridge
- [ ] Verificar configuración memory-wiki en OpenClaw
- [ ] Confirmar que el bridge está en modo correcto
- [ ] Probar inyección manual de un concepto
- [ ] Verificar que aparece en `wiki/concepts/`
- **Criterio de finalización:** Al menos 1 concepto, 1 entidad y 1 síntesis generados automáticamente

#### F1.5 — Confirmar failover chain
- [ ] Provocar rate limit intencional (si posible en test)
- [ ] Verificar que ollama-qwen responde como fallback
- [ ] Documentar tiempos de recovery
- **Criterio de finalización:** Failover documentado y funcional

### 🟡 Prioridad Alta

#### F2.1 — Activar memory-wiki bridge
- [ ] Configurar bridge en modo `active` (no solo `bridge`)
- [ ] Verificar que nuevas sesiones generan entries wiki
- [ ] Monitorear por 24h
- **Criterio de finalización:** Wiki se auto-enriquece con cada sesión relevante

#### F2.2 — Configurar embeddings para vault
- [ ] Verificar que nomic-embed-text está funcionando en Ollama
- [ ] Conectar vault como fuente de embeddings
- [ ] Probar búsqueda semántica contra vault
- **Criterio de finalización:** `memory search "algo del vault"` retorna resultados

#### F2.3 — Consolidación semanal
- [ ] Crear cron weekly-synthesis
- [ ] Template para síntesis semanal
- [ ] Agregar a operacion-diaria.md
- **Criterio de finalización:** Cada lunes aparece síntesis en `memories/`

### 🟢 Prioridad Media

#### F3.1 — Configurar Ralph Loop
- [ ] Implementar LOOP_STATE.json tracking
- [ ] Configurar checkpoints automáticos
- [ ] Definir límites de iteración
- [ ] Probar loop completo con tarea simple
- **Criterio de finalización:** Loop ejecuta 3+ iteraciones sin drift

#### F3.2 — Activar delegación
- [ ] Verificar que planner puede invocar workers
- [ ] Probar delegación code-reviewer
- [ ] Probar delegación researcher
- [ ] Documentar patrones de delegación exitosos
- **Criterio de finalización:** Al menos 1 tarea completada via delegación

---

## 4. Prioridades globales

1. **Estabilidad** — Todo lo que funciona debe seguir funcionando
2. **Observabilidad** — Saber qué pasa en todo momento (logs, briefs, decisions)
3. **Memoria** — El sistema no debe olvidar lo aprendido
4. **Autonomía** — Reducir dependencia de intervención humana
5. **Escalabilidad** — Preparar para más agentes y más skills

---

## 5. Criterios de finalización globales

- [ ] 0 cron jobs fallando
- [ ] session-export generando al menos 1 sesión por día de uso
- [ ] memory-wiki generando concepts/entities/syntheses
- [ ] Skill catalog sin duplicaciones
- [ ] Failover chain verificada
- [ ] 3 archivos de identidad completados
- [ ] Ralph Loop funcional con checkpoints
- [ ] Vault auto-enriquecido sin intervención manual

---

## Ver también

- [[TASKS]] — Checklist ejecutable derivado de este PRD
- [[AGENT_RULES]] — Reglas de comportamiento para agentes
- [[seeds/sistema-arquitectura]] — Stack técnico
- [[00-MOC]] — Mapa principal
