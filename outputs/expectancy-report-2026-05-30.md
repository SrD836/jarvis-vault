# Expectancy report — v7 policy cohort — 2026-05-30

> Honesto: netea el realizado v7 (ganadores cerrados) con el mark-to-market del libro abierto
> (perdedores latentes). Sin esto, el 98% wr es puro sesgo de supervivencia.

| métrica | valor |
|---|---|
| realizado v7 (cerrados, policy) | +500.57 USD · n=53 · wr=98.1% · exp/trade=+9.44 |
| histórico contaminado (excluido) | -1932.30 USD · n=333 (stop_loss/fallback régimen muerto) |
| mark-to-market abiertas | -401.80 USD · n=17 priced / 18 |
| **expectancy COMBINADA (a mercado)** | **+98.77 USD · n=70 · exp/trade=+1.41** |
| peor caso (abiertas→0) | -784.35 USD · n=71 · exp/trade=-11.05 · (stake abierto $1285) |
| Brier (predicciones resueltas) | 0.1199 · n=89 |

### Reglas nuevas firing (blocked.jsonl)
- R1_longshot: 12
- R5_precedents: 5
- R3_catalyst_24h: 2

### Libro abierto por categoría (riesgo latente, mark-to-market)
- other: -137.94 USD (n=8)
- geopolitics: -121.50 USD (n=5)
- market: -80.14 USD (n=2)
- executive-action: -52.39 USD (n=1)
- sports-season: -9.03 USD (n=1)
- elections: -0.79 USD (n=1)

### Veredicto (criterio de David)
- expectancy combinada > 0 con n>=50: SÍ (exp=+1.41, n=70)
- robusto al peor caso (abiertas→0) > 0: NO (exp=-11.05)
- Brier <= 0.25: SÍ (Brier=0.1199)
- reglas nuevas bloqueando: SÍ

## FAIL — NO avanzar a dinero real
Falla: peor caso -11.05 (libro longshot abierto $1285).
Sangra (abiertas): other (-138), geopolitics (-122), market (-80).
La estrategia v7 cierra ganadores rápido y deja correr perdedores (sin stop) → el realizado SIEMPRE parece bueno hasta que el libro abierto resuelve. Prueba de sostenibilidad = ventana forward con R1/R3 activos desde el inicio (ya bloquean longshots).
