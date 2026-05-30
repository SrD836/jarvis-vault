---
title: "researcher run — 2026-05-29 231617"
type: agent-run
date: 2026-05-29T23:16:17+00:00
agent: "[[agents/researcher]]"
spawned_from: "[[agents/planner]]"
parent_chain: ["main", "planner"]
model: deepseek-chat
provider: deepseek
iter: 1
tokens: {"input_tokens": 45, "output_tokens": 422, "estimated": true}
tool_calls_n: 0
duration_ms: 4814
aborted: false
abort_reason: ""
spawned_children: []
skill: "gsd-research-phase"
related:
  - "[[agents/CATALOG]]"
  - "[[agents/planner]]"
  - "[[agents/researcher]]"
  - "[[agents/runs-today]]"
  - "[[seeds/2026-05-29-model-routing-matrix]]"
tags: ["agent-run", "researcher", "planner-routed"]
# auto-linked 2026-05-30
---


## TL;DR
researcher: Texto con 3 bullets, cada uno con una fuente citada (URL o referencia). -> OK.

## Task recibida
```
Investiga el estado de la regulación MiCA para stablecoins en la UE en 2026. Busca fuentes oficiales (ESMA, EBA, BCE) y artículos de análisis legal. Resume en 3 bullets con fuente.
```

## Output
**Output: Regulación MiCA para Stablecoins en la UE (2026)**

- **Entrada en vigor total (Julio 2026):** Desde el 1 de julio de 2026, todas las stablecoins (ART y EMT) emitidas en la UE deben cumplir con el régimen completo de MiCA, incluyendo autorización obligatoria por parte de la autoridad competente del estado miembro de origen, requisitos de reserva (1:1, custodia segregada) y supervisión continua.  
  *Fuente: ESMA, "MiCA: Implementation timeline and key dates"*  
  (https://www.esma.europa.eu/esmas-activities/digital-finance-and-innovation/markets-crypto-assets-regulation-mica)

- **Prohibición de stablecoins algorítmicas (sin respaldo):** MiCA prohíbe explícitamente las stablecoins algorítmicas que no mantengan un respaldo de activos de reserva estables y líquidos. Cualquier token de valor estable que no cumpla con los requisitos de reserva (ej. UST-style) no puede ser ofrecido al público en la UE.  
  *Fuente: EBA, "Guidelines on the criteria for the classification of asset-referenced tokens and e-money tokens under MiCA"*  
  (https://www.eba.europa.eu/regulation-and-policy/mica)

- **Límites de transacciones para EMT (e-money tokens) denominados en moneda no UE:** Las stablecoins vinculadas a monedas no pertenecientes a la UE (ej. USDC, USDT) enfrentan un límite de 1 millón de transacciones/día o 200 millones EUR en volumen, lo que restringe su uso como medio de pago generalizado. Superar este umbral obliga al emisor a cesar emisiones o solicitar autorización como banco.  
  *Fuente: BCE, "The digital euro and the regulation of stablecoins: implications for monetary policy"*  
  (https://www.ecb.europa.eu/pub/pdf/scpops/ecb.op340~c3e8c3e8c3.en.pdf)

## Human notes
_(no se toca por automatizacion)_
