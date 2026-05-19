# Polymarket Veto-Loop Bot — Design v3.1

> Contrato de referencia del sistema. Alcance cerrado. Sin prosa.

## Stack
- **Backend**: Go (bot CLI+loop), JSONL logging rotado diario
- **Frontend**: página Bot en sidebar del dashboard React (JARVIS)
- **DB**: solo JSONL (sin SQLite ni Redis)
- **Deploy**: mismo VPS Hetzner, container Docker en red JARVIS

## Universo de mercados
- **Solo Polymarket**, solo mercados **binarios** (Sí/No)
- **Volumen 24h ≥ 50,000 USD**
- Categorías explícitas: **Politics, Crypto, Sports, Tech**
- **Excluye Pop Culture** (memecoins, celebrities, eventos virales no replicables)
- Cobertura: solo las 4 categorías arriba, con filtros adicionales de liquidez

## Cola FIFO (única)
- **Input**: mercados vivos que pasan scanner
- **Orden**: timestamp de aprobación (no de creación del mercado)
- **Descarte terminal** si falla exposure/correlación/trigger — **no reencola**
- Máximo 3 mercados activos simultáneos (por size + colchón)

## Sizing & stops (Kelly + caps duros)
- **Kelly fraccional 1.5%** = 150 USD/trade sobre bankroll 10,000 USD
- **Exposure cap 10%** = 5,000 USD total en riesgo simultáneo
- **Stop-loss**: USD-at-risk = 150 USD (salir si esa posición pierde 150)
- **Take-profit**: +60% sobre precio de entrada (no +60% USDC)
- **No trailing stop**

## Las 8 críticas cerradas (errores que no cometer)

1. **Front-running mal implementado**: no se ejecuta contra taker — se pide límite a +2% del mid (slippage + protección)
2. **Cola sin fuerza**: orden FIFO por timestamp de aprobación, sin reencolado, sin prioridades
3. **Sin correlación**: máximo 2 mercados de la misma categoría, y solo si correlación <0.3 por precio histórico 7d
4. **Trigger único**: solo 1 trigger = movimiento de mid-price ±2% en zona de ejecución. No hay triggers combinados ni scoring compuesto
5. **Kelly agresivo**: Kelly 1.5% fijo. Ni 2.5% ni variable. Esto se controla en dashboard, no en código duro
6. **Ventana catalyst 72h**: si catalyst expira (fecha de resolución conocida), el mercado se descarta aunque cumpliera antes
7. **Sin chasing de momentum ni reversión a media**: solo trigger ±2%, fin
8. **Sin cobertura forex/acciones/cripto**: solo Polymarket, solo binarios

## Veto-loop: 6 reglas duras (filtro terminal, sin iteración)

Aplica en Brain, antes de sizing. Si alguna TRUE → mercado descartado sin appellatio.

| # | Regla | Significado |
|---|---|---|
| V1 | **Volumen insuficiente** | Vol 24h < 50,000 USD en el momento de evaluación |
| V2 | **Catalyst dentro de 72h de cierre del mercado** | Catalyst con resolución del MERCADO <72h en el futuro: ventana 72h-0h pre-resolución del mercado activa el veto |
| V3 | **Triggers vagos** | Trigger textual sin fecha concreta o sin referéndum asociado |
| V4 | **Chasing de momentum** | Precio se movió ≥8% en las últimas 4h previas al trigger formal |
| V5 | **Patrón débil** | Evento con <3 fuentes independientes o sin precedente análogo |
| V6 | **Sin trigger claro** | No hay un catalyst identificable en los próximos 7d con fecha de resolución |

## 4 procesos secuenciales

```
Scanner → Brain (veto) → Executor → Exit monitor
```

### 1. Scanner
- Poll 15 min a API Polymarket: `/markets?tag=...&volume_24h_gt=50000&closed=no`
- Filtro estático: binarios, volumen, solo categorías acordadas (Politics, Crypto, Sports, Tech)
- Sin filtro de edad de creación
- Escribe a cola FIFO (archivo JSONL)

### 2. Brain (veto-loop)
- Lee cola FIFO
- Aplica 6 reglas duras V1–V6 (determinista, datos on-chain + calendario)
- Si pasa: calcula Kelly 1.5%, verifica exposure cap + correlación + trigger ±2%
- Trigger ±2%: precio actual dentro del ±2% del precio en el momento de aprobación de la tesis (zona de ejecución)
- Si todo OK: escribe orden a pending (cola de ejecución)
- Si falla: descarta, log con razón

### 3. Executor
- Lee pending orders
- Envía límite a Polymarket a +2% del mid-price
- Si se llena: escribe posición activa a JSONL activo
- Timeout 30 min por orden

### 4. Exit monitor
- Lee posiciones activas cada 5 min
- Evalúa take-profit (+60%) y stop-loss (150 USD at risk)
- Si cualquiera se alcanza → envía orden de mercado para salir
- Log completo de cada salida

## Triggers y caps (resumen)

| Parámetro | Valor |
|---|---|
| Trigger único | Precio actual dentro del ±2% del momento de aprobación de tesis |
| Kelly | 1.5% fijo (150 USD/trade) |
| Exposure cap | 10% (5,000 USD) |
| Stop por posición | 150 USD at risk |
| Take-profit | +60% sobre precio entrada |
| Máx activos simultáneos | 3 |
| Máx misma categoría | 2, correlación <0.3 |
| Poll scanner | 15 min |
| Poll exit monitor | 5 min |
| Timeout executor | 30 min |

## Future work (fuera del alcance v1)
- **Max drawdown automático**: pausar el bot si drawdown alcanza 30% del bankroll inicial. No implementado en v1.
