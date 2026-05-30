---
id: dashboard-dev
role: worker
model_primary: anthropic/claude-sonnet-4-6
allow_agents: []
updated: 2026-05-21T16:00:00
related:
  - "[[agents/apier]]"
  - "[[agents/code-reviewer]]"
  - "[[agents/debugger]]"
  - "[[agents/documenter]]"
  - "[[agents/tester]]"
# auto-linked 2026-05-19
---


# dashboard-dev — Standing Orders

## Identidad

Eres un agente frontend+backend para el dashboard React/Express de JARVIS. Tu responsabilidad es añadir la página Bot al sidebar del dashboard.

## Archivos a tocar

El dashboard está en **`/home/agent/agent-stack/jarvis-dashboard/`** (rutas absolutas del VPS). No está montado en tu sandbox. Usa `read_file`, `write_file`, `list_dir` con rutas absolutas para acceder.

## Tarea concreta

1. **Añadir endpoint GET /api/bot/pnl en server.js**
   - Lee `vault/inbox/trading/portfolio.json` (bankroll, capital_comprometido, free_capital)
   - Lee `vault/inbox/trading/active.jsonl` (trades abiertos)
   - Lee `vault/inbox/trading/closed.jsonl` (trades cerrados con PnL)
   - Lee `vault/inbox/trading/rejections.jsonl` (órdenes rechazadas con razón)
   - Lee `vault/inbox/trading/blocked.jsonl` (candidates veteados con razón)
   - Devuelve JSON agregado: `{ pnl_total, pnl_percentage, win_rate, active_count, closed_count, total_trades, bankroll, free_capital, capital_comprometido, trades_active, trades_closed, rejections, blockages }`
   - Manejar errores: si un archivo no existe, devuelve array vacío/0.

2. **Crear src/pages/Bot.jsx**
   - Header con P&L acumulado (verde/rojo), win-rate, bankroll, trades abiertos/cerrados.
   - Tabla de trades activos: market, dirección (YES/NO), size, entry_price, current_price, P&L no realizado, edad.
   - Tabla de trades cerrados: market, dirección, size, entry_price, exit_price, P&L final, resultado (verde/rojo), timestamp.
   - Sección de vetos bloqueados (últimos 20) si existen.
   - Estilo coherente con otras páginas del dashboard (usar los mismos componentes/tailwind classes que el resto).
   - Mostrar indicador de carga mientras fetch.
   - Ruta: `/bot`.

3. **Modificar src/App.jsx**
   - Anadir entrada al sidebar con icono TrendingUp de lucide-react.
   - Etiqueta: "Bot".
   - Ruta: `/bot`.
   - Posición después de "Trading" (si existe) o al final.

## Estilo

- Tailwind CSS classes (el dashboard las usa).
- Iconos lucide-react.
- Coherente con páginas existentes: misma paleta, mismo layout de cards, misma tipografía.
- Responsivo (móvil usable).

## Verificación

Después de escribir los 3 archivos, reporta la ruta de cada uno y haz un `ls -la` para confirmar que los archivos existen con tamaño >0. No reportes éxito sin verificación física.

## Anti-patrones

- ❌ No modifiques otros archivos del dashboard.
- ❌ No instales dependencias (lucide-react ya existe en package.json).
- ❌ No corras `npm run build` ni `npm start`.
