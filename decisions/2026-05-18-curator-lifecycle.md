---
title: "Curator v4: lifecycle gaps + dry-run + update_last_activity"
type: decision
date: 2026-05-18
decision: curator.py.new listo para revisión. No ejecutar hasta code-review + test pass.
outcome: pending
tags: [decision, curator, lifecycle, skills, dry-run]
related:
  - "[[../04-skills-log/index]]"
  - "[[../02-sessions/index]]"
---

## Contexto

El usuario reportó que `curator.py` "carece de lógica de lifecycle". El análisis del script v3 (en
`/home/node/.openclaw/scripts/curator.py`) revela que la lógica stale/archive **ya existe** en la
función `check_lifecycle()`, pero tiene tres gaps reales que la hacen inefectiva en producción.

---

## Gaps identificados en v3

### Gap 1 — `last_activity_at` nunca se actualiza (BUG CRÍTICO)

`last_activity_at` se escribe **solo una vez**, en `is_skill_duplicate()` al registrar la skill por
primera vez. Después nunca vuelve a actualizarse. Resultado: todas las skills eventualmente
llegan a stale/archived basándose en la fecha de primera vista, no en el último uso real.

**Evidencia:** no existe ninguna función pública que el runtime pueda llamar para actualizar la
actividad de una skill tras ejecutarla.

### Gap 2 — Sin modo dry-run

El script solo puede ejecutarse en modo destructivo (escribe `.usage.json` y logs de vault). No
hay forma de simular cambios antes de aplicarlos, lo que hace arriesgado el testing.

### Gap 3 — Skills en PROBATION/REWRITE se archivan automáticamente

Una skill que falla frecuentemente (estado REWRITE o PROBATION) podría ser archivada por
inactividad antes de que los agentes la revisen. El v3 no tiene guard para este caso.

---

## Diff conceptual v3 → v4

| Aspecto | v3 | v4 |
|---------|----|----|
| `--dry-run` flag | ❌ no existe | ✅ imprime todo, no escribe nada |
| `update_last_activity()` | ❌ no existe | ✅ función pública para el runtime |
| `--update-activity PATH` | ❌ no existe | ✅ CLI flag para hooks post-ejecución |
| Guard PROBATION/REWRITE | ❌ se archivan igual | ✅ máximo stale, no archived |
| `is_skill_duplicate()` | confuso (retorna bool invertido) | renombrado a `is_skill_registered()` |
| Resumen final | ❌ sin conteos | ✅ `active=N stale=N archived=N` |
| argparse | ❌ sin CLI | ✅ argparse completo |

### Lógica lifecycle: sin cambio funcional

Las reglas stale/archive ya eran correctas en v3. v4 no las modifica:
- `last_activity_at > 30d` → `stale`
- `last_activity_at > 90d` → `archived` (salvo guard PROBATION/REWRITE → queda en stale)
- Regeneración de `SKILL_INDEX.md` al final: idéntica, con dry-run support añadido

---

## Archivos generados

| Archivo | Ruta | Propósito |
|---------|------|-----------|
| Backup | `/home/node/.openclaw/scripts/curator.py.bak-pre-lifecycle-1779095848` | Restore point |
| Nuevo script | `/home/node/.openclaw/scripts/curator.py.new` | Candidato para reemplazar v3 |

`curator.py` original NO fue modificado.

---

## Plan de rollout

### Fase 1 — Revisión (code-reviewer)
- Verificar que la lógica de lifecycle no tiene regresiones vs v3
- Validar que el guard PROBATION/REWRITE no bloquea transiciones legítimas
- Revisar que dry-run no tiene side effects (ninguna escritura en disco)
- Confirmar que `update_last_activity()` es idempotente

### Fase 2 — Testing (tester)
Casos mínimos a cubrir:

```
1. dry-run con skill sin .usage.json       → registra en memoria, no escribe archivo
2. dry-run con skill con 45d de inactividad → imprime "active → stale", no muta .usage.json
3. dry-run con skill con 95d de inactividad → imprime "active → archived", no muta
4. dry-run con skill PROBATION + 95d       → imprime "active → stale" (guard activo)
5. update_last_activity en skill stale     → state vuelve a "active"
6. generate_index dry-run                  → no sobreescribe SKILL_INDEX.md
```

### Fase 3 — Promoción (requiere aprobación manual)
```bash
# Solo cuando code-review + tests pasen:
cp /home/node/.openclaw/scripts/curator.py.new \
   /home/node/.openclaw/scripts/curator.py
```

### Fase 4 — Wiring del runtime
Para que `last_activity_at` se actualice en producción, añadir al hook post-ejecución de skill:
```bash
python3 ~/.openclaw/scripts/curator.py --update-activity "$SKILL_PATH"
```

---

## Rollback

```bash
cp /home/node/.openclaw/scripts/curator.py.bak-pre-lifecycle-1779095848 \
   /home/node/.openclaw/scripts/curator.py
```
