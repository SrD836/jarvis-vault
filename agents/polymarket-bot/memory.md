---
title: "polymarket-bot — memory"
type: agent-memory
agent: "[[agents/polymarket-bot]]"
date: 2026-05-19
tags: [agent-memory, polymarket, bot, lessons]
related:
  - "[[../polymarket-bot]]"
  - "[[../../03-decisions]]"
---

# polymarket-bot — memoria persistente

> Brain consulta este archivo ANTES de aprobar una tesis. Si encuentra patterns similares con outcome negativo, aplica veto preventivo (M1).

## Vetos pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | price_yes | rule | reason |
|---|---|---|---|---|---|
| 2026-05-28T17:30:04Z | israel-and-syria-normalize-relations-by-june-30-2026 | geopolitics | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon -148 d) |
| 2026-05-28T17:30:04Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon -148 d) |
| 2026-05-28T17:30:04Z | will-us-unemployment-reach-at-least-5pt5-in-2026 | other | 0.1960 | P3_low_absolute_liquidity | liquidity $809 < absolute min $5000 |
| 2026-05-28T17:30:04Z | nicols-maduro-released-from-custody-by-december-31-2026 | other | 0.1300 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 117 días) |
| 2026-05-28T17:30:04Z | will-grigor-dimitrov-win-the-2026-australian-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -116 d) |
| 2026-05-28T17:30:04Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -94 d) |
| 2026-05-28T17:30:04Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -94 d) |
| 2026-05-28T17:30:04Z | internet-access-restored-in-iran-by-may-31-2026 | geopolitics | 0.7880 | P2 | mercado ya expiró (endDate=2026-03-14T00:00:00Z, hace 75 días) |
| 2026-05-28T17:30:04Z | internet-access-restored-in-iran-by-june-30-2026 | geopolitics | 0.8400 | P2 | mercado ya expiró (endDate=2026-03-14T00:00:00Z, hace 75 días) |
| 2026-05-28T17:30:04Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon -67 d) |
| 2026-05-28T17:30:04Z | will-matej-tonin-be-the-next-prime-minister-of-slovenia | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon -67 d) |
| 2026-05-28T17:30:04Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -67 d) |
| 2026-05-28T17:30:04Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -65 d) |
| 2026-05-28T17:30:04Z | will-alex-vanopslagh-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -65 d) |
| 2026-05-28T17:30:04Z | will-lars-boje-mathiesen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon -65 d) |
| 2026-05-28T17:30:04Z | will-lars-lkke-rasmussen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0410 | P0_floor | price floor: 0.0410 < 0.100 (horizon -65 d) |
| 2026-05-28T17:30:04Z | weed-rescheduled-by-june-30 | other | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon -58 d) |
| 2026-05-28T17:30:04Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -46 d) |
| 2026-05-28T17:30:04Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -45 d) |
| 2026-05-28T17:30:04Z | us-x-iran-diplomatic-meeting-by-june-15-2026-671 | geopolitics | 0.4800 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 13 días) |
| 2026-05-28T17:30:04Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.7500 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 13 días) |
| 2026-05-28T17:30:04Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.1000 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 13 días) |
| 2026-05-28T17:30:04Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon -9 d) |
| 2026-05-28T17:30:04Z | will-gregg-kirkpatrick-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -9 d) |
| 2026-05-28T17:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon -1 d) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-28 | geopolitics | 0.0920 | P9 | P9: geopolitics pump cluster (price 0.09, 0d) |
| 2026-05-28T17:30:04Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-28-2026 | geopolitics | 0.0790 | P9 | P9: geopolitics pump cluster (price 0.08, 0d) |
| 2026-05-28T17:30:04Z | bitcoin-above-76k-on-may-28-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | bitcoin-above-74k-on-may-28-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | will-the-price-of-bitcoin-be-between-76000-78000-on-may-28-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | ethereum-above-2000-on-may-28-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | bitcoin-above-78k-on-may-28-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | bitcoin-above-82k-on-may-28-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | bitcoin-above-80k-on-may-28-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-29 | geopolitics | 0.1900 | P9 | P9: geopolitics pump cluster (price 0.19, 0d) |
| 2026-05-28T17:30:04Z | bitcoin-above-66k-on-may-29-2026 | market | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.950 |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-240-259 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-180-199 | other | 0.2890 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-200-219 | other | 0.0270 | P0_floor | price floor: 0.0270 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-160-179 | other | 0.6500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-140-159 | other | 0.0240 | P0_floor | price floor: 0.0240 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-220-239 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-260-279 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-22-may-29-280-299 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-28T17:30:04Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 1 d) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.2900 | P9 | P9: geopolitics pump cluster (price 0.29, 1d) |
| 2026-05-28T17:30:04Z | ucl-psg-ars-2026-05-30-psg | other | 0.4200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | ucl-psg-ars-2026-05-30-ars | other | 0.3100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0600 | P9 | P9: geopolitics pump cluster (price 0.06, 2d) |
| 2026-05-28T17:30:04Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.1600 | P9 | P9: geopolitics pump cluster (price 0.16, 2d) |
| 2026-05-28T17:30:04Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0510 | P9 | P9: geopolitics pump cluster (price 0.05, 2d) |
| 2026-05-28T17:30:04Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0600 | P9 | P9: geopolitics pump cluster (price 0.06, 2d) |
| 2026-05-28T17:30:04Z | will-alibaba-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0130 | P0_floor | price floor: 0.0130 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.4700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0230 | P0_floor | price floor: 0.0230 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | trump-out-as-president-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-trump-agree-to-iranian-oil-sanction-relief-by-may-31 | executive-action | 0.2400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.0850 | P9 | P9: geopolitics pump cluster (price 0.09, 2d) |
| 2026-05-28T17:30:04Z | will-trump-agree-to-unfreeze-iranian-assets-by-may-31 | other | 0.1700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-arsenal-win-the-202526-champions-league | sports-season | 0.4300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0640 | P9 | P9: geopolitics pump cluster (price 0.06, 2d) |
| 2026-05-28T17:30:04Z | will-trump-restart-project-freedom-by-may-31 | other | 0.0630 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.2100 | P9 | P9: geopolitics pump cluster (price 0.21, 2d) |
| 2026-05-28T17:30:04Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.1620 | P9 | P9: geopolitics pump cluster (price 0.16, 2d) |
| 2026-05-28T17:30:04Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.2700 | P9 | P9: geopolitics pump cluster (price 0.27, 2d) |
| 2026-05-28T17:30:04Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0830 | P9 | P9: geopolitics pump cluster (price 0.08, 2d) |
| 2026-05-28T17:30:04Z | will-psg-win-the-202526-champions-league | sports-season | 0.5800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-may-31-111 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | starmer-out-by-may-31-2026 | executive-action | 0.0120 | P0_floor | price floor: 0.0120 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-abelardo-de-la-espriella-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.2770 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-28T17:30:04Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-reach-130-in-may-2026-733 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-crude-oil-wti-hit-high-105-in-may | market | 0.0110 | P0_floor | price floor: 0.0110 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0390 | P0_floor | price floor: 0.0390 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-reach-150-in-may-2026-196-364 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0760 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-wti-reach-120-in-may-2026-592-217 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-reach-140-in-may-2026-916 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.3000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-wti-reach-110-in-may-2026-116-472 | market | 0.0080 | P0_floor | price floor: 0.0080 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-xrp-reach-2-in-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-reach-95k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-ethereum-dip-to-1400-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-40k-in-may-2026-277-395 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-reach-100k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0080 | P0_floor | price floor: 0.0080 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | over-4m-committed-to-the-printr-public-sale | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-45k-in-may-2026-669-185-973 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-50k-in-may-2026-896 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-65k-in-may-2026-183-857-425 | market | 0.0160 | P0_floor | price floor: 0.0160 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.1590 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-solana-reach-170-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-reach-80k-may-25-31-2026 | market | 0.0180 | P0_floor | price floor: 0.0180 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-ethereum-reach-2600-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0290 | P0_floor | price floor: 0.0290 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-solana-dip-to-40-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-28T17:30:04Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-28T17:30:04Z | elon-musk-of-tweets-may-26-june-2-500plus | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-28T17:30:04Z | will-yoon-sang-hyun-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-chong-won-oh-win-the-2026-seoul-mayoral-election | elections | 0.7800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-lee-kwang-jae-win-the-2026-gangwon-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-han-dong-hoon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-kim-boo-kyum-win-the-2026-daegu-mayoral-election | elections | 0.1400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-oh-se-hoon-win-the-2026-seoul-mayoral-election | elections | 0.2300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.5800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-won-hee-ryong-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 5 d) |
| 2026-05-28T17:30:04Z | will-emma-navarro-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 8 d) |
| 2026-05-28T17:30:04Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 8 d) |
| 2026-05-28T17:30:04Z | will-rafael-lpez-aliaga-win-the-2026-peruvian-presidential-election | elections | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0370 | P0_floor | price floor: 0.0370 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-novak-djokovic-win-the-2026-mens-french-open | other | 0.1990 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-alexander-zverev-win-the-2026-mens-french-open | other | 0.3290 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.0840 | P0_floor | price floor: 0.0840 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-tommy-paul-win-the-2026-mens-french-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-karen-khachanov-win-the-2026-mens-french-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-jan-lennard-struff-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-alex-michelsen-win-the-2026-mens-french-open | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-casper-ruud-win-the-2026-mens-french-open | other | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-stefanos-tsitsipas-win-the-2026-mens-french-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.2800 | P9 | P9: geopolitics pump cluster (price 0.28, 9d) |
| 2026-05-28T17:30:04Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0350 | P0_floor | price floor: 0.0350 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-jannik-sinner-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-ben-shelton-win-the-2026-mens-french-open | other | 0.0550 | P0_floor | price floor: 0.0550 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 9 d) |
| 2026-05-28T17:30:04Z | will-a-world-cup-game-in-mexico-be-relocated-abroad | sports-season | 0.0390 | P0_floor | price floor: 0.0390 < 0.100 (horizon 12 d) |
| 2026-05-28T17:30:04Z | iran-closes-its-airspace-by-june-15 | geopolitics | 0.1940 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-june-15 | geopolitics | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.3400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-kim-farington-be-the-republican-nominee-for-senate-in-virginia | elections | 0.3110 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-al-mina-be-the-republican-nominee-for-senate-in-virginia | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 18 d) |
| 2026-05-28T17:30:04Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 19 d) |
| 2026-05-28T17:30:04Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 19 d) |
| 2026-05-28T17:30:04Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 19 d) |
| 2026-05-28T17:30:04Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9770 | P0_ceiling | price ceiling: 0.9770 > 0.950 |
| 2026-05-28T17:30:04Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 19 d) |
| 2026-05-28T17:30:04Z | will-vitality-win-iem-cologne-major-2026 | other | 0.5100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-ivan-cepeda-castro-win-the-2026-colombian-presidential-election | elections | 0.3400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 23 d) |
| 2026-05-28T17:30:04Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 25 d) |
| 2026-05-28T17:30:04Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 26 d) |
| 2026-05-28T17:30:04Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 26 d) |
| 2026-05-28T17:30:04Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 26 d) |
| 2026-05-28T17:30:04Z | will-tynan-lawrence-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 29 d) |
| 2026-05-28T17:30:04Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 29 d) |
| 2026-05-28T17:30:04Z | gustavo-petro-out-as-leader-of-colombia-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | greece-x-turkey-military-engagement-by-june-30 | other | 0.0340 | P0_floor | price floor: 0.0340 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-jpmorgan-chase-fail-by-june-30-2026 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-30-2026-15... | geopolitics | 0.6900 | M1 | memoria: slug prefix match; same category; same price bucket mid (score 0.90) |
| 2026-05-28T17:30:04Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0520 | P0_floor | price floor: 0.0520 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | iran-closes-its-airspace-by-june-30-432-786-462-866 | geopolitics | 0.2440 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | masoud-pezeshkian-out-by-june-30-242-112 | executive-action | 0.0580 | P0_floor | price floor: 0.0580 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-the-vegas-golden-knights-win-the-2026-nhl-stanley-cup | other | 0.4280 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-michael-jackson-be-confirmed-to-have-visited-epsteins-island | other | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | hamad-bin-isa-al-khalifa-out-as-bahrain-king | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | israel-x-iran-permanent-peace-deal-by-june-30-2026-262 | geopolitics | 0.1580 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-deutsche-bank-fail-by-june-30-2026 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0420 | P0_floor | price floor: 0.0420 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-russia-enter-sloviansk-by-june-30 | geopolitics | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-june-30-2026 | geopolitics | 0.2500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-trump-pardon-tiger-woods-by-june-30 | executive-action | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-the-montreal-canadiens-win-the-2026-nhl-stanley-cup | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-the-carolina-hurricanes-win-the-2026-nhl-stanley-cup | other | 0.5600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | us-iran-nuclear-deal-by-june-30 | geopolitics | 0.4200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-mohammed-bin-salman-cease-to-be-the-de-facto-leader-of-saudi-arabia-by-june-30-2026 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-ukraine-recapture-crimean-territory-by-june-30-2026 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-june-30-848-925 | geopolitics | 0.7600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | bill-clinton-divorce-by-june-30 | other | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-june | geopolitics | 0.3900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-discords-market-cap-be-less-than-15b-at-market-close-on-ipo-day | crypto-launch | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-reza-pahlavi-enter-iran-by-june-30-555-947-362-379-213 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-silver-si-hit-low-35-by-end-of-june-325-231 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-silver-si-hit-low-45-by-end-of-june-972-272 | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 32 d) |
| 2026-05-28T17:30:04Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 33 d) |
| 2026-05-28T17:30:04Z | will-the-oklahoma-city-thunder-win-the-2026-nba-finals | sports-season | 0.5800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.1520 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-new-york-knicks-win-the-2026-nba-finals | sports-season | 0.2850 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | microstrategy-sells-any-bitcoin-by-june-30-2026 | market | 0.5340 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | microstrategy-sells-any-bitcoin-by-may-31-2026 | market | 0.1340 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-bitcoin-hit-150k-by-june-30-2026 | market | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 33 d) |
| 2026-05-28T17:30:04Z | will-jasmine-paolini-be-the-2026-womens-wimbledon-winner | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0440 | P0_floor | price floor: 0.0440 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-clara-tauson-be-the-2026-womens-wimbledon-winner | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-iga-witek-be-the-2026-womens-wimbledon-winner | other | 0.1710 | P3_low_absolute_liquidity | liquidity $1747 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-ons-jabeur-be-the-2026-womens-wimbledon-winner | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 44 d) |
| 2026-05-28T17:30:04Z | will-neymar-play-in-the-2026-fifa-world-cup-for-brazil | sports-season | 0.8000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $613 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-iraq-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-france-win-the-2026-fifa-world-cup-924 | sports-season | 0.1670 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-jordan-win-the-2026-fifa-world-cup-233 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-england-win-the-2026-fifa-world-cup-937 | sports-season | 0.1120 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-qatar-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | P0_floor | price floor: 0.0520 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0850 | P0_floor | price floor: 0.0850 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0380 | P0_floor | price floor: 0.0380 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-portugal-win-the-2026-fifa-world-cup-912 | sports-season | 0.1020 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 52 d) |
| 2026-05-28T17:30:04Z | will-spain-win-the-2026-fifa-world-cup-963 | sports-season | 0.1700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | us-x-iran-permanent-peace-deal-by-july-31-2026-831-252 | geopolitics | 0.6300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 63 d) |
| 2026-05-28T17:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-july-31 | geopolitics | 0.6000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 63 d) |
| 2026-05-28T17:30:04Z | will-paul-reevs-be-the-republican-nominee-for-az-01 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 67 d) |
| 2026-05-28T17:30:04Z | will-kristen-mcdonald-rivet-win-the-2026-michigan-democratic-primary | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 67 d) |
| 2026-05-28T17:30:04Z | will-scott-jensen-win-the-2026-minnesota-governor-republican-primary-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 74 d) |
| 2026-05-28T17:30:04Z | will-jay-collins-be-the-republican-nominee-for-florida-governor | elections | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 81 d) |
| 2026-05-28T17:30:04Z | will-tyreek-hill-play-for-philadelphia-eagles-next | other | 0.6940 | P3_low_absolute_liquidity | liquidity $85 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-david-njoku-play-for-tennessee-titans-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 95 d) |
| 2026-05-28T17:30:04Z | will-cdu-win-the-most-seats-in-the-2026-sachsen-anhalt-parliamentary-elections | other | 0.0570 | P0_floor | price floor: 0.0570 < 0.100 (horizon 100 d) |
| 2026-05-28T17:30:04Z | tush-push-banned-for-2026-nfl-season | other | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 104 d) |
| 2026-05-28T17:30:04Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 107 d) |
| 2026-05-28T17:30:04Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 107 d) |
| 2026-05-28T17:30:04Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 107 d) |
| 2026-05-28T17:30:04Z | will-the-moderate-party-m-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0360 | P0_floor | price floor: 0.0360 < 0.100 (horizon 107 d) |
| 2026-05-28T17:30:04Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 107 d) |
| 2026-05-28T17:30:04Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 122 d) |
| 2026-05-28T17:30:04Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0300 | P0_floor | price floor: 0.0300 < 0.100 (horizon 122 d) |
| 2026-05-28T17:30:04Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P0_floor | price floor: 0.0380 < 0.100 (horizon 122 d) |
| 2026-05-28T17:30:04Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 122 d) |
| 2026-05-28T17:30:04Z | will-civic-platform-gp-win-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 124 d) |
| 2026-05-28T17:30:04Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5800 | P8 | P8: election 124d out, price 0.58 en banda ruido [0.30, 0.70] |
| 2026-05-28T17:30:04Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $700 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 128 d) |
| 2026-05-28T17:30:04Z | will-eduardo-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 128 d) |
| 2026-05-28T17:30:04Z | will-jair-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 128 d) |
| 2026-05-28T17:30:04Z | will-renan-santos-win-the-2026-brazilian-presidential-election | elections | 0.1540 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-partido-liberal-pl-win-the-most-seats-in-the-next-brazilian-senate-election | elections | 0.7200 | P3_low_absolute_liquidity | liquidity $2373 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-luiz-incio-lula-da-silva-win-the-2026-brazilian-presidential-election | elections | 0.4200 | P8 | P8: election 128d out, price 0.42 en banda ruido [0.30, 0.70] |
| 2026-05-28T17:30:04Z | will-tarcisio-de-frietas-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 128 d) |
| 2026-05-28T17:30:04Z | will-romeu-zema-win-the-2026-brazilian-presidential-election | elections | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon 128 d) |
| 2026-05-28T17:30:04Z | will-the-chicago-white-sox-win-the-2026-world-series | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-the-miami-marlins-win-the-2026-world-series | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-cleveland-guardians-win-the-2026-american-league-championship-series | sports-season | 0.0860 | P0_floor | price floor: 0.0860 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-colorado-rockies-win-the-2026-national-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-houston-astros-win-the-2026-american-league-championship-series | sports-season | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-pittsburgh-pirates-win-the-2026-national-league-championship-series | sports-season | 0.0410 | P0_floor | price floor: 0.0410 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-san-francisco-giants-win-the-2026-national-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-san-diego-padres-win-the-2026-national-league-championship-series | sports-season | 0.0650 | P0_floor | price floor: 0.0650 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-st-louis-cardinals-win-the-2026-national-league-championship-series | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 156 d) |
| 2026-05-28T17:30:04Z | will-toni-atkins-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-tony-thurmond-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-betty-yee-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-tom-steyer-win-the-california-governor-election-in-2026 | elections | 0.1710 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-republican-party-control-the-house-after-the-2026-midterm-elections | elections | 0.2000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-eric-swalwell-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-elaine-culotti-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-rick-caruso-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 158 d) |
| 2026-05-28T17:30:04Z | will-sergio-prez-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 191 d) |
| 2026-05-28T17:30:04Z | will-lance-stroll-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 191 d) |
| 2026-05-28T17:30:04Z | will-oliver-bearman-be-the-2026-f1-drivers-champion | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 191 d) |
| 2026-05-28T17:30:04Z | will-isack-hadjar-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 191 d) |
| 2026-05-28T17:30:04Z | will-pierre-gasly-be-the-2026-f1-drivers-champion | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 191 d) |
| 2026-05-28T17:30:04Z | will-liam-lawson-be-the-2026-f1-drivers-champion | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 191 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-4pt5-at-the-end-of-2026-139 | other | 0.0420 | P0_floor | price floor: 0.0420 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-4pt0-at-the-end-of-2026-386 | other | 0.2330 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-3pt75-at-the-end-of-2026-166 | other | 0.4610 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-4pt25-at-the-end-of-2026-984 | other | 0.1010 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-3pt0-at-the-end-of-2026-593 | other | 0.0380 | P0_floor | price floor: 0.0380 < 0.100 (horizon 194 d) |
| 2026-05-28T17:30:04Z | will-payton-tolle-win-the-2026-al-rookie-of-the-year-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-brice-matthews-win-the-2026-al-rookie-of-the-year-award | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-columbus-crew-win-the-2026-mls-cup | other | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-new-york-city-fc-win-the-2026-mls-cup | other | 0.0370 | P0_floor | price floor: 0.0370 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-nashville-sc-win-the-2026-mls-cup | other | 0.0580 | P0_floor | price floor: 0.0580 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-san-jose-earthquakes-win-the-2026-mls-cup | other | 0.0460 | P0_floor | price floor: 0.0460 < 0.100 (horizon 204 d) |
| 2026-05-28T17:30:04Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-larry-page-be-richest-person-on-december-31 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3191 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-alberta-join-the-us | other | 0.0420 | P0_floor | price floor: 0.0420 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-jorge-rodrguez-be-the-leader-of-venezuela-end-of-2026 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | jerome-powell-out-from-fed-board-by-december-31 | other | 0.4800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | new-covid-variant-of-concern-before-2027 | other | 0.1700 | P3_low_absolute_liquidity | liquidity $2231 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-anyone-be-charged-over-epstein-disclosures | other | 0.1600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | vanta-ipo-before-2027 | crypto-launch | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | us-x-iran-permanent-peace-deal-by-december-31-2026-961-587-341-574-555 | geopolitics | 0.8100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-zelenskyy-and-putin-meet-next-in-qatar-uae | other | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-israel-reopen-its-embassy-in-iran-in-2026 | geopolitics | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-china-invade-taiwan-before-2027 | geopolitics | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3644 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-alireza-arafi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | zhang-youxia-sentenced-to-prison-before-2027 | other | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4497 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-the-us-invade-iran-before-2027 | geopolitics | 0.1900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-feds-lower-bound-reach-2pt5-or-lower-before-2027-289-849-151-768 | other | 0.0570 | P0_floor | price floor: 0.0570 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0620 | P0_floor | price floor: 0.0620 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-bitcoin-replace-sha-256-before-2027 | other | 0.0520 | P0_floor | price floor: 0.0520 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-the-feds-upper-bound-reach-5pt25-or-higher-before-2027-679 | other | 0.0340 | P0_floor | price floor: 0.0340 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-9-fed-rate-cuts-happen-in-2026 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | russia-x-ukraine-ceasefire-agreement-by-december-31-2026 | geopolitics | 0.4400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-beyonc-be-the-top-spotify-artist-for-2026 | entertainment | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | masoud-pezeshkian-out-by-december-31-793-289 | executive-action | 0.2500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-donald-trump-be-the-leader-of-venezuela-end-of-2026 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | ledger-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1721 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-there-be-at-least-7500-measles-cases-in-the-us-in-2026-287-181 | other | 0.1600 | P3_low_absolute_liquidity | liquidity $3564 < absolute min $5000 |
| 2026-05-28T17:30:04Z | 5kt-meteor-strike-in-2026 | weather-natural | 0.3400 | P3_low_absolute_liquidity | liquidity $3573 < absolute min $5000 |
| 2026-05-28T17:30:04Z | ripple-labs-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $3176 < absolute min $5000 |
| 2026-05-28T17:30:04Z | us-strike-on-cuba-by-december-31 | other | 0.5400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-jensen-huang-be-richest-person-on-december-31-229 | other | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1320 | P3_low_absolute_liquidity | liquidity $3934 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-the-iranian-regime-fall-by-the-end-of-2026 | other | 0.1400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | eu-dissolves-before-2027 | other | 0.0350 | P0_floor | price floor: 0.0350 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-zelenskyy-and-putin-meet-next-in-us | other | 0.0430 | P0_floor | price floor: 0.0430 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-russia-capture-all-of-lyman-by-december-31-2026 | geopolitics | 0.2200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | save-act-signed-into-law-in-2026 | other | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-openai-be-acquired-before-2027-859 | other | 0.0730 | P0_floor | price floor: 0.0730 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | ukraine-signs-peace-deal-with-russia-before-2027 | geopolitics | 0.3000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $1678 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | sports-season | 0.0960 | P0_floor | price floor: 0.0960 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 216 d) |
| 2026-05-28T17:30:04Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 217 d to resolution (>=7 threshold) |
| 2026-05-28T17:30:04Z | will-stablecoins-hit-500b-before-2027 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $2147 < absolute min $5000 |
| 2026-05-28T17:30:04Z | usdc-depeg-by-december-31-968 | other | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | crypto-launch | 0.0470 | P0_floor | price floor: 0.0470 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $750 < absolute min $5000 |
| 2026-05-28T17:30:04Z | metamask-fdv-above-3b-one-day-after-launch-363-663-664-569-222 | crypto-launch | 0.0970 | P0_floor | price floor: 0.0970 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2600 | P3_low_absolute_liquidity | liquidity $1615 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1353 < absolute min $5000 |
| 2026-05-28T17:30:04Z | opensea-fdv-above-3b-one-day-after-launch-786-419-967-433-241-436 | crypto-launch | 0.0870 | P0_floor | price floor: 0.0870 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | metamask-fdv-above-1b-one-day-after-launch-978-851-628-634 | crypto-launch | 0.2100 | P4_pre_event | pre-event slug + 217 d to resolution (>=7 threshold) |
| 2026-05-28T17:30:04Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4200 | P3_low_absolute_liquidity | liquidity $1071 < absolute min $5000 |
| 2026-05-28T17:30:04Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | crypto-launch | 0.3170 | P4_pre_event | pre-event slug + 217 d to resolution (>=7 threshold) |
| 2026-05-28T17:30:04Z | microstrategy-sells-any-bitcoin-by-december-31-2026 | market | 0.8500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-fomo-launch-a-token-by-december-31-2026 | crypto-launch | 0.3600 | P3_low_absolute_liquidity | liquidity $1149 < absolute min $5000 |
| 2026-05-28T17:30:04Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.1990 | P3_low_absolute_liquidity | liquidity $2804 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | puffpaw-fdv-above-400m-one-day-after-launch | crypto-launch | 0.0880 | P0_floor | price floor: 0.0880 < 0.100 (horizon 217 d) |
| 2026-05-28T17:30:04Z | will-any-month-of-2026-be-the-hottest-on-record | other | 0.8800 | P3_low_absolute_liquidity | liquidity $1507 < absolute min $5000 |
| 2026-05-28T17:30:04Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | sports-season | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | sports-season | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | sports-season | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | sports-season | 0.0570 | P0_floor | price floor: 0.0570 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-atlanta-falcons-win-the-2027-nfl-nfc-championship-312 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-minnesota-vikings-win-the-2027-nfl-nfc-championship-884 | sports-season | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | sports-season | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | sports-season | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-pittsburgh-steelers-win-the-2027-nfl-afc-championship-898 | sports-season | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 241 d) |
| 2026-05-28T17:30:04Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 307 d) |
| 2026-05-28T17:30:04Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 336 d) |
| 2026-05-28T17:30:04Z | will-fabien-roussel-win-the-2027-french-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 336 d) |
| 2026-05-28T17:30:04Z | will-sacramento-kings-win-the-2027-nba-finals | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 398 d) |
| 2026-05-28T17:30:04Z | spacex-ipo-closing-market-cap-above-3t-337-436 | crypto-launch | 0.1500 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-28T17:30:04Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 581 d) |
| 2026-05-28T17:30:04Z | will-spacexs-public-ticker-be-spax-713 | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 581 d) |
| 2026-05-28T17:30:04Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 581 d) |
| 2026-05-28T17:30:04Z | will-spacex-not-ipo-by-december-31-2027 | crypto-launch | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 581 d) |
| 2026-05-28T17:30:04Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2230 | P4_pre_event | pre-event slug + 582 d to resolution (>=7 threshold) |
| 2026-05-28T17:30:04Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-tulsi-gabbard-win-the-2028-republican-presidential-nomination | elections | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-tom-brady-win-the-2028-republican-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-pete-hegseth-win-the-2028-republican-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-rand-paul-win-the-2028-republican-presidential-nomination | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-oprah-winfrey-win-the-2028-democratic-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | elections | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-michelle-obama-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-tim-walz-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-james-talarico-win-the-2028-democratic-presidential-nomination | elections | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-greg-abbott-win-the-2028-republican-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-jon-stewart-win-the-2028-democratic-presidential-nomination-518 | elections | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 893 d) |
| 2026-05-28T17:30:04Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon -1 d) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-june-7-265-824-655 | geopolitics | 0.6700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | kevin-warsh-cuts-rates-at-first-fed-meeting | other | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon -1 d) |
| 2026-05-28T17:30:04Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon -1 d) |
| 2026-05-28T17:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-31-665-831-238 | geopolitics | 0.4300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-28T17:30:04Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9980 | P0_ceiling | price ceiling: 0.9980 > 0.950 |
| 2026-05-28T17:30:04Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0330 | P0_floor | price floor: 0.0330 < 0.100 (horizon -1 d) |
| 2026-05-28T17:30:04Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon -1 d) |
| 2026-05-28T17:30:04Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon -1 d) |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T16:50:02Z | T-2287649-1779634803475 | uncategorized | 0.9300 | 0.9990 | 4.40 | market_closed |  | 59.27 | short | 0.08 |
| 2026-05-24T16:50:02Z | T-2227286-1779640204759 | uncategorized | 0.4430 | 0.7140 | 36.80 | take_profit |  | 60.15 | short | 0.01 |
| 2026-05-24T17:05:02Z | T-2227370-1779638403705 | uncategorized | 0.6600 | 0.1000 | -54.67 | stop_loss |  | 64.43 | short | 0.05 |
| 2026-05-24T17:05:02Z | T-2227281-1779642003777 | uncategorized | 0.0500 | 0.0100 | -50.83 | stop_loss |  | 63.54 | short | 0.00 |
| 2026-05-24T17:15:02Z | T-2227341-1779570004545 | uncategorized | 0.5200 | 0.0100 | -63.82 | stop_loss |  | 65.07 | short | 0.84 |
| 2026-05-24T17:15:02Z | T-2227235-1779642003777 | uncategorized | 0.1400 | 0.0100 | -56.66 | stop_loss |  | 61.02 | short | 0.01 |
| 2026-05-24T17:20:02Z | T-2227243-1779620403996 | uncategorized | 0.4900 | 0.0100 | -62.47 | stop_loss |  | 63.77 | short | 0.26 |
| 2026-05-24T17:20:02Z | T-2227163-1779638403705 | uncategorized | 0.0500 | 0.0020 | -60.62 | stop_loss |  | 63.14 | short | 0.06 |
| 2026-05-24T17:20:02Z | T-2227160-1779638403705 | uncategorized | 0.1900 | 0.0100 | -57.45 | stop_loss |  | 60.64 | short | 0.06 |
| 2026-05-24T17:20:02Z | T-2227284-1779640204759 | uncategorized | 0.3100 | 0.0100 | -59.40 | stop_loss |  | 61.38 | short | 0.03 |
| 2026-05-24T17:20:02Z | T-2227345-1779642003777 | uncategorized | 0.1400 | 0.0100 | -57.82 | stop_loss |  | 62.27 | short | 0.01 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T17:50:01Z | T-2227158-1779638403705 | uncategorized | 0.7700 | 0.9990 | 17.67 | market_closed |  | 59.43 | short | 0.08 |
| 2026-05-24T18:10:01Z | T-2285040-1779489003932 | uncategorized | 0.9380 | 0.9990 | 3.87 | market_closed |  | 59.53 | short | 1.82 |
| 2026-05-24T18:10:02Z | T-2285043-1779579003772 | uncategorized | 0.7600 | 0.9990 | 18.90 | market_closed |  | 60.09 | short | 0.78 |
| 2026-05-24T19:15:02Z | T-2126542-1779643804205 | uncategorized | 0.0700 | 0.1080 | 30.39 | take_profit |  | 55.97 | medium | 0.07 |
| 2026-05-24T19:20:01Z | T-2313549-1779571804012 | uncategorized | 0.2300 | 0.3300 | 28.64 | take_profit |  | 65.86 | short | 0.91 |
| 2026-05-24T19:50:02Z | T-2227309-1779640204759 | uncategorized | 0.8400 | 0.9990 | 11.86 | market_closed |  | 62.63 | short | 0.14 |
| 2026-05-24T19:50:02Z | T-2227286-1779642003777 | uncategorized | 0.7200 | 0.9990 | 25.12 | market_closed |  | 64.84 | short | 0.12 |
| 2026-05-24T19:55:02Z | T-2238716-1779611403231 | uncategorized | 0.3800 | 0.0500 | -56.74 | stop_loss |  | 65.34 | short | 0.48 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T20:05:01Z | T-2287695-1779607803227 | uncategorized | 0.5400 | 0.7600 | 26.87 | take_profit |  | 65.94 | short | 0.52 |
| 2026-05-24T20:05:02Z | T-2238716-1779652804112 | uncategorized | 0.0600 | 0.0100 | -45.00 | stop_loss |  | 54.00 | short | 0.00 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T20:15:02Z | T-2287581-1779651003578 | uncategorized | 0.6000 | 0.8800 | 24.56 | take_profit |  | 52.62 | short | 0.03 |
| 2026-05-24T20:25:02Z | T-2119073-1779643804205 | uncategorized | 0.0750 | 0.1060 | 22.22 | take_profit |  | 53.76 | medium | 0.12 |
| 2026-05-24T20:45:02Z | T-2287672-1779651003578 | uncategorized | 0.6700 | 0.0600 | -46.95 | stop_loss |  | 51.57 | short | 0.05 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T20:55:02Z | T-2287674-1779654603617 | uncategorized | 0.5420 | 0.9150 | 36.96 | take_profit |  | 53.71 | short | 0.02 |
| 2026-05-24T21:00:02Z | T-2287625-1779645603833 | uncategorized | 0.9300 | 0.9990 | 3.65 | market_closed |  | 49.16 | short | 0.12 |
| 2026-05-24T21:15:02Z | T-2287791-1779618603939 | uncategorized | 0.6500 | 0.9500 | 28.59 | take_profit |  | 61.93 | short | 0.45 |
| 2026-05-24T21:40:02Z | T-2287674-1779656404369 | uncategorized | 0.9190 | 0.9990 | 4.75 | market_closed |  | 54.54 | short | 0.03 |
| 2026-05-24T21:50:02Z | T-2287791-1779658203711 | uncategorized | 0.8200 | 0.1300 | -44.72 | stop_loss |  | 53.15 | short | 0.01 |
| 2026-05-24T21:55:02Z | T-2321923-1779633003824 | uncategorized | 0.1180 | 0.0190 | -50.69 | stop_loss |  | 60.42 | short | 0.31 |
| 2026-05-24T23:25:02Z | T-2313551-1779499803774 | uncategorized | 0.2040 | 0.0360 | -50.52 | stop_loss |  | 61.34 | medium | 1.91 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T23:45:02Z | T-2287695-1779654603617 | uncategorized | 0.8400 | 0.9990 | 10.37 | market_closed |  | 54.81 | short | 0.14 |
| 2026-05-24T23:55:02Z | T-2287581-1779654603617 | uncategorized | 0.9400 | 0.9990 | 3.51 | market_closed |  | 55.92 | short | 0.14 |
| 2026-05-25T00:05:02Z | T-2287768-1779643804205 | uncategorized | 0.7300 | 0.9990 | 21.05 | market_closed |  | 57.12 | short | 0.27 |
| 2026-05-25T00:50:02Z | T-2097867-1779669003789 | uncategorized | 0.1800 | 0.2600 | 22.97 | take_profit |  | 51.67 | short | 0.01 |
| 2026-05-25T01:50:02Z | T-2097857-1779660004159 | uncategorized | 0.8000 | 0.9990 | 13.34 | market_closed |  | 53.63 | short | 0.16 |
| 2026-05-25T01:55:02Z | T-2313548-1779642003777 | uncategorized | 0.3000 | 0.4200 | 23.44 | take_profit |  | 58.61 | short | 0.37 |
| 2026-05-25T02:00:02Z | T-2294008-1779670804814 | uncategorized | 0.3300 | 0.5200 | 30.01 | take_profit |  | 52.13 | short | 0.04 |
| 2026-05-25T03:00:02Z | T-2241873-1779562803685 | uncategorized | 0.1480 | 0.0210 | -51.25 | stop_loss |  | 59.72 | short | 1.33 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T03:10:02Z | T-2097863-1779660004159 | uncategorized | 0.7300 | 0.9990 | 19.37 | market_closed |  | 52.56 | short | 0.22 |
| 2026-05-25T03:15:02Z | T-2294008-1779674403929 | uncategorized | 0.5300 | 0.0900 | -46.25 | stop_loss |  | 55.71 | short | 0.05 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T03:25:01Z | T-2097884-1779678004140 | uncategorized | 0.3100 | 0.9700 | 112.00 | take_profit |  | 52.60 | short | 0.02 |
| 2026-05-25T05:15:02Z | T-2002531-1779652804112 | uncategorized | 0.0770 | 0.1110 | 22.90 | take_profit |  | 51.86 | medium | 0.39 |
| 2026-05-25T05:25:02Z | T-2313547-1779604204419 | uncategorized | 0.1640 | 0.2350 | 28.36 | take_profit |  | 65.51 | short | 0.95 |
| 2026-05-25T06:05:02Z | T-2321929-1779679804173 | uncategorized | 0.2200 | 0.3300 | 28.24 | take_profit |  | 56.47 | short | 0.11 |
| 2026-05-25T06:15:01Z | T-2334096-1779517803519 | uncategorized | 0.9200 | 0.9990 | 4.94 | market_closed |  | 57.50 | short | 1.99 |
| 2026-05-25T12:30:03Z | T-2002531-1779690604559 | uncategorized | 0.1180 | 0.2300 | 53.23 | take_profit |  | 56.08 | medium | 0.25 |
| 2026-05-25T16:35:13Z | T-2133405-1779487204257 | uncategorized | 0.8400 | 0.1010 | -54.32 | stop_loss |  | 61.74 | medium | 2.77 |
| 2026-05-25T16:35:29Z | T-2313550-1779571804012 | uncategorized | 0.1960 | 0.0280 | -53.13 | stop_loss |  | 61.99 | short | 1.80 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T16:35:29Z | T-1731345-1779643804205 | uncategorized | 0.2200 | 0.3300 | 24.30 | take_profit |  | 48.59 | long | 0.96 |
| 2026-05-25T16:35:29Z | T-2321926-1779665404197 | uncategorized | 0.7000 | 0.9990 | 21.65 | take_profit |  | 50.70 | short | 0.71 |
| 2026-05-25T16:35:29Z | T-2313548-1779674403929 | uncategorized | 0.4300 | 0.6300 | 25.39 | take_profit |  | 54.60 | short | 0.61 |
| 2026-05-25T16:35:38Z | T-2293518-1779687003842 | uncategorized | 0.2400 | 0.0010 | -56.29 | stop_loss |  | 56.52 | short | 0.46 |
| 2026-05-25T16:35:46Z | T-2321929-1779690604559 | uncategorized | 0.2400 | 0.0010 | -56.99 | stop_loss |  | 57.23 | short | 0.42 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T18:55:01Z | T-2111605-1779647403942 | uncategorized | 0.1650 | 0.2480 | 25.67 | take_profit |  | 51.03 | medium | 1.02 |
| 2026-05-25T19:00:02Z | T-2293512-1779660004159 | uncategorized | 0.9450 | 0.9990 | 2.94 | market_closed |  | 51.51 | short | 0.87 |
| 2026-05-25T19:50:02Z | T-2324252-1779642003777 | uncategorized | 0.1300 | 0.2700 | 64.40 | take_profit |  | 59.80 | short | 1.12 |
| 2026-05-25T20:00:02Z | T-2111564-1779643804205 | uncategorized | 0.2700 | 0.5300 | 50.73 | take_profit |  | 52.68 | medium | 1.10 |
| 2026-05-25T20:05:02Z | T-2230536-1779679804173 | uncategorized | 0.9000 | 0.9990 | 6.09 | market_closed |  | 55.34 | medium | 0.69 |
| 2026-05-25T20:20:01Z | T-2126515-1779647403942 | uncategorized | 0.0810 | 0.1200 | 24.08 | take_profit |  | 50.01 | medium | 1.08 |
| 2026-05-25T21:35:09Z | T-2313547-1779687003842 | uncategorized | 0.2360 | 0.0400 | -46.01 | stop_loss |  | 55.39 | short | 0.67 |
| 2026-05-25T23:00:26Z | T-2119073-1779654603617 | uncategorized | 0.1230 | 0.0230 | -42.79 | stop_loss |  | 52.64 | medium | 1.10 |
| 2026-05-25T23:05:08Z | T-2334097-1779570004545 | uncategorized | 0.9500 | 0.0800 | -57.23 | stop_loss |  | 62.49 | short | 2.09 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-26T00:10:01Z | T-2159726-1779597003852 | uncategorized | 0.1200 | 0.1800 | 33.26 | take_profit |  | 66.51 | medium | 1.82 |
| 2026-05-26T00:30:20Z | T-1808970-1779211803817 | uncategorized | 0.0530 | 0.0080 | -44.01 | stop_loss |  | 51.83 | medium | 6.29 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-26T01:15:02Z | T-2169995-1779730346526 | market | 0.1000 | 0.1400 | 22.45 | take_profit |  | 56.12 | long | 0.32 |
| 2026-05-26T01:50:02Z | T-2132799-1779726947128 | market | 0.2500 | 0.3700 | 29.21 | take_profit | claudemax-websearch | 60.85 | medium | 0.38 |
| 2026-05-26T04:05:02Z | T-2111561-1779643804205 | uncategorized | 0.0520 | 0.0750 | 22.84 | take_profit |  | 51.63 | medium | 1.44 |
| 2026-05-26T05:50:02Z | T-2303427-1779766204447 | market | 0.0700 | 0.1000 | 26.78 | take_profit |  | 62.48 | short | 0.10 |
| 2026-05-26T07:10:08Z | T-2313549-1779652804112 | uncategorized | 0.3400 | 0.0600 | -43.58 | stop_loss |  | 52.92 | short | 1.47 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-26T07:40:01Z | T-957019-1779643804205 | uncategorized | 0.3300 | 0.4800 | 22.54 | take_profit |  | 49.58 | long | 1.59 |
| 2026-05-26T10:50:02Z | T-2303427-1779775204890 | market | 0.1000 | 0.2260 | 81.28 | take_profit |  | 64.51 | short | 0.20 |
| 2026-05-26T12:20:26Z | T-2169995-1779759026423 | market | 0.1500 | 0.0200 | -54.79 | stop_loss |  | 63.22 | long | 0.45 |
| 2026-05-26T12:20:51Z | T-2303427-1779793283201 | market | 0.2900 | 0.0570 | -52.38 | stop_loss |  | 65.20 | short | 0.05 |
| 2026-05-26T15:30:08Z | T-2303427-1779798604693 | market | 0.0760 | 0.0080 | -57.55 | stop_loss |  | 64.32 | short | 0.12 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-27T08:25:01Z | T-562186-1779210004342 | uncategorized | 0.9250 | 0.9990 | 6.38 | market_closed |  | 79.73 | medium | 7.64 |
| 2026-05-27T13:50:23Z | T-2313280-1779814832575 | market | 0.6800 | 0.0600 | -70.91 | stop_loss | claudemax-websearch | 77.77 | short | 0.87 |
| 2026-05-27T14:55:09Z | T-2322404-1779712204481 | market | 0.1500 | 0.0300 | -45.72 | stop_loss |  | 57.15 | medium | 2.10 |
| 2026-05-28T05:25:18Z | T-2322008-1779899405064 | market | 0.8070 | 0.1600 | -60.30 | stop_loss |  | 75.21 | short | 0.54 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-28T07:48:17Z | T-648378-1779206389814 | uncategorized | 0.8800 | 0.8800 | 0.00 | manual_reset_v7 | legacy | 150.00 | long | 8.66 |
| 2026-05-28T07:48:17Z | T-839358-1779210004342 | uncategorized | 0.0545 | 0.0570 | 3.31 | manual_reset_v7 |  | 72.07 | long | 8.62 |
| 2026-05-28T07:48:17Z | T-566136-1779211803817 | uncategorized | 0.5900 | 0.5700 | -1.83 | manual_reset_v7 |  | 53.97 | medium | 8.60 |
| 2026-05-28T07:48:17Z | T-566140-1779233403338 | uncategorized | 0.4300 | 0.4200 | -1.26 | manual_reset_v7 |  | 54.25 | medium | 8.35 |
| 2026-05-28T07:48:17Z | T-2316216-1779642003777 | uncategorized | 0.4200 | 0.4100 | -1.37 | manual_reset_v7 |  | 57.43 | medium | 3.62 |
| 2026-05-28T07:48:17Z | T-2111640-1779643804205 | uncategorized | 0.2080 | 0.0830 | -30.41 | manual_reset_v7 |  | 50.60 | medium | 3.60 |
| 2026-05-28T07:48:17Z | T-569344-1779667204230 | uncategorized | 0.3200 | 0.2920 | -4.47 | manual_reset_v7 |  | 51.13 | medium | 3.33 |
| 2026-05-28T07:48:17Z | T-2155052-1779726947128 | geopolitics | 0.8400 | 0.7800 | -4.17 | manual_reset_v7 |  | 58.44 | long | 2.63 |
| 2026-05-28T07:48:17Z | T-2270338-1779732049801 | geopolitics | 0.7200 | 0.6000 | -9.17 | manual_reset_v7 | claudemax-websearch | 55.00 | long | 2.57 |
| 2026-05-28T07:48:17Z | T-2155023-1779733851679 | geopolitics | 0.7200 | 0.6000 | -8.98 | manual_reset_v7 | claudemax-websearch | 53.90 | long | 2.55 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-28T07:48:17Z | T-1809418-1779760804552 | other | 0.8770 | 0.8950 | 1.31 | manual_reset_v7 |  | 63.76 | long | 2.24 |
| 2026-05-28T07:48:17Z | T-824952-1779784234712 | market | 0.7300 | 0.8500 | 10.66 | manual_reset_v7 | claudemax-websearch | 64.85 | long | 1.97 |
| 2026-05-28T07:48:17Z | T-701547-1779791404097 | market | 0.1200 | 0.0900 | -15.89 | manual_reset_v7 |  | 63.55 | long | 1.89 |
| 2026-05-28T07:48:17Z | T-569366-1779809574675 | elections | 0.7200 | 0.6500 | -6.14 | manual_reset_v7 | claudemax-websearch | 63.17 | medium | 1.68 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-28T07:48:17Z | T-947269-1779870679596 | elections | 0.7100 | 0.7100 | 0.00 | manual_reset_v7 | claudemax-websearch | 77.94 | medium | 0.97 |
| 2026-05-28T07:48:17Z | T-1087513-1779894004820 | other | 0.7400 | 0.7300 | -1.04 | manual_reset_v7 | claudemax-websearch | 76.74 | medium | 0.70 |
## Reglas blandas aprendidas (autogenerada por brain cuando un pattern se repite 3+ veces)



**Reglas iniciales seed** (basadas en post-mortem v1, ver [[../../03-decisions/2026-05-19-polymarket-v1-baseline-failure]]):

- **Evitar long-tail extremo**: `current_price_yes < 0.05` en categorías políticas a largo plazo (>12 meses). Razón: take-profit +60% sobre 0.005 = 0.008 → ganancia $0.45 sobre $150 size, no merece el riesgo.
- **Categoría `uncategorized` con resolución >12 meses**: tratar con escepticismo elevado. La API no expone categoría → no podemos verificar correlación. Veto si volumen 24h real (cuando disponible) <$1k.
- **Volumen `volumeNum` cumulativo NO es proxy fiable**. Para v2 preferir spread (`bestAsk - bestBid`) si está disponible. Si `bestBid` o `bestAsk` faltan/cero, considerar mercado ilíquido y vetar.
- **Long-tail floor SUBIDO de 0.05 -> 0.10** (2026-05-19, post-mortem trade T-608399 Discord): entry 0.066, exit 0.007 = -89.4% en 10 min. Cerca del floor, stop_loss_pct 80% se dispara con el spread bid/ask normal sin noticias. Veto: current_price_yes < 0.10 si horizon > 7 dias.
- **Liquidez absoluta minima**: vetar mercados con liquidity_usd < $5000 en bruto, ademas del gate Size x 4. Mercados con $2-3k de profundidad son trampas: el bestBid colapsa al primer market-sell. Discord tenia $2916.
- **Pre-IPO / pre-launch markets**: si slug contiene ipo-day | launch | tge | mainnet | listing y faltan >=7 dias hasta endDate -> veto. No hay forma honesta de modelar el dia-cero de un asset que aun no existe.
<!-- auto-generated 2026-05-20T08:55:57Z -->
- En categoría `uncategorized` · horizonte `?` · precio `0.10-0.30`: 13 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `?` · precio `0.30-0.70`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-20T15:00:05Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-20T15:30:05Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-20T16:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 7 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-20T19:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 7 losses, 1 wins (win rate 12%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-20T20:30:04Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 10 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-20T21:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 7 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T01:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 7 losses, 1 wins (win rate 12%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T02:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 11 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 8 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T02:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 12 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 9 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T03:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 10 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T04:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 12 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 10 losses, 1 wins (win rate 9%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T07:00:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 13 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 11 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T07:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 11 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T11:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 11 losses, 2 wins (win rate 15%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T13:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 7 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 13 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T14:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 14 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T15:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 14 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T17:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 14 losses, 2 wins (win rate 12%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T19:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 17 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T20:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 6 losses, 1 wins (win rate 14%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-21T20:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 6 losses, 2 wins (win rate 25%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T00:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 8 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 11 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T02:30:03Z -->
- En categoría `uncategorized` · horizonte `?` · precio `0.10-0.30`: 12 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 12 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T03:00:03Z -->
- En categoría `uncategorized` · horizonte `?` · precio `0.10-0.30`: 7 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 16 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 13 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T03:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 8 losses, 1 wins (win rate 11%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T04:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 17 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 10 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T06:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 13 losses, 1 wins (win rate 7%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T07:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 18 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 14 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T13:00:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 6 losses, 1 wins (win rate 14%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T14:00:04Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 7 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 11 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 15 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T16:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 17 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T17:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 12 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 18 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T19:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 18 losses, 1 wins (win rate 5%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T19:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 15 losses, 2 wins (win rate 12%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T20:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 12 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T20:30:05Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 15 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 19 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 17 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T21:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 17 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T21:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 17 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T22:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 10 losses, 2 wins (win rate 17%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T22:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 16 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 16 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-22T23:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 15 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T01:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 11 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T02:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 15 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T06:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 9 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T08:30:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 5 losses, 1 wins (win rate 17%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T13:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 16 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T16:00:04Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 5 losses, 2 wins (win rate 29%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T16:30:04Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 8 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T17:30:04Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 9 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T18:30:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 10 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 12 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T19:00:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 11 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `medium` · precio `0.30-0.70`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 9 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T19:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 15 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T20:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 8 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T20:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 20 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 7 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T21:00:04Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 12 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 22 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T21:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 15 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-23T23:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 21 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T00:00:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 11 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T01:00:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `0.10-0.30`: 11 losses, 2 wins (win rate 15%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T02:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 21 losses, 1 wins (win rate 5%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T02:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 21 losses, 2 wins (win rate 9%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T03:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 21 losses, 3 wins (win rate 12%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T04:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 23 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 9 losses, 1 wins (win rate 10%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T05:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 9 losses, 2 wins (win rate 18%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T06:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 13 losses, 1 wins (win rate 7%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T16:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 10 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T16:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `<0.10`: 12 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T17:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 11 losses, 1 wins (win rate 8%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T17:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 18 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T18:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 12 losses, 3 wins (win rate 20%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T19:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 14 losses, 1 wins (win rate 7%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T20:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 16 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T21:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 18 losses, 1 wins (win rate 5%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T21:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 18 losses, 2 wins (win rate 10%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-24T22:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 19 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T00:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 17 losses, 2 wins (win rate 11%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T00:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 17 losses, 3 wins (win rate 15%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T02:00:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 19 losses, 2 wins (win rate 10%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 17 losses, 4 wins (win rate 19%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T03:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 18 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T03:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.30-0.70`: 19 losses, 1 wins (win rate 5%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T06:30:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 13 losses, 2 wins (win rate 13%). **Veto soft o reducir size 50%.**
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 17 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T20:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 15 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T23:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 19 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-26T01:02:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `<0.10`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-26T04:30:33Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `<0.10`: 5 losses, 1 wins (win rate 17%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-26T07:31:30Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `<0.10`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**


## Human notes
_(no se toca por automatización)_

## Sources stats (rolling 30d, computed by brain)

| domain | trades | wins | losses | win_rate | total_pnl | blacklisted |
|---|---|---|---|---|---|---|
| claudemax-websearch | 31 | 5 | 24 | 17.2% | -262.25 | YES |

## Anti-patterns identificados- same-day-price-target · intraday-decay · no-momentum-confirm · short-horizon-chalk · high-entry-low-exit — visto en  (2026-05-26, pnl $-57.55)

- same-day-price-level · no-source-approval · overconfident-high-entry · crypto-short-horizon — visto en bitcoin-above-74k-on-may-28-2026 (2026-05-28, pnl $-60.30)
- extreme-price-target · short-horizon-commodity · no-catalyst-confirmed · low-probability-moonshot — visto en  (2026-05-27, pnl $-45.72)
- same-day-price-target · price-near-threshold · recency-bias-entry · high-vol-short-horizon · geopolitics-short-horizon — visto en bitcoin-above-76k-on-may-27-2026 (2026-05-27, pnl $-70.91)

- zero-research-entry · same-day-expiry · crypto-price-level · intraday-short-horizon · below-50-underdog — visto en  (2026-05-26, pnl $-52.38)
- no-sources-consulted · against-stated-hodl-policy · horizon-mismatch · high-entry-low-prob-event · stop-loss-under-hours — visto en  (2026-05-26, pnl $-54.79)
- exact-count-range-bet · celebrity-behavior-noise · short-horizon-volatility · uncategorized-low-edge · no-source-validation — visto en  (2026-05-26, pnl $-43.58)
- zero-sources-entry · geopolitics-short-horizon · treaty-deadline-bet · low-prob-long-shot · uncategorized-no-research — visto en  (2026-05-26, pnl $-44.01)
- ceasefire-fragile · geopolitics-short-horizon · overpriced-high-prob · low-margin-safety · stop-loss-cliff — visto en  (2026-05-25, pnl $-57.23)
- zero-research-entry · geopolitics-numeric-threshold · uncategorized-low-liquidity · low-prob-no-edge — visto en  (2026-05-25, pnl $-42.79)
- narrow-activity-range · social-behavior-bet · unpredictable-subject · low-probability-band · stop-loss-short-horizon — visto en  (2026-05-25, pnl $-46.01)
- tweet-count-range · behavior-prediction · no-sources-consulted · low-edge-bet · short-horizon-noise — visto en  (2026-05-25, pnl $-56.99)
- short-horizon-crypto · same-day-chalk-bet · no-source-validation · low-timeframe-bet · stop-loss-wipeout — visto en  (2026-05-25, pnl $-56.29)
- narrow-count-range · social-activity-bet · no-sources-entry · uncategorized-market · behavior-prediction — visto en  (2026-05-25, pnl $-53.13)
- geopolitical-action-bet · no-source-validation · high-entry-no-catalyst · short-horizon-collapse · stop-loss-whipsaw — visto en  (2026-05-25, pnl $-54.32)
_(autogenerada por exit_monitor tras cada pérdida — brain consume vía M3)_
