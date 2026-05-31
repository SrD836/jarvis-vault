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
| 2026-05-31T08:04:27Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 890 d, liq $1104549) |
| 2026-05-31T08:04:27Z | will-zohran-mamdani-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1615787) |
| 2026-05-31T08:04:27Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1188099) |
| 2026-05-31T08:04:27Z | will-tim-walz-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1756886) |
| 2026-05-31T08:04:27Z | will-steve-bannon-win-the-2028-republican-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 890 d, liq $1272711) |
| 2026-05-31T08:04:27Z | will-glenn-youngkin-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1087143) |
| 2026-05-31T08:04:27Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1704692) |
| 2026-05-31T08:04:27Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1119369) |
| 2026-05-31T08:04:27Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $2271270) |
| 2026-05-31T08:04:27Z | will-eric-trump-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1961058) |
| 2026-05-31T08:04:27Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1227176) |
| 2026-05-31T08:04:27Z | will-jamie-dimon-win-the-2028-us-presidential-election | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 890 d, liq $766655) |
| 2026-05-31T08:04:27Z | will-lebron-james-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $2195108) |
| 2026-05-31T08:04:27Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $957) |
| 2026-05-31T08:04:27Z | us-announces-new-iran-agreementceasefire-extension-by-june-7-265-824-655 | geopolitics | 0.3500 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 5 veces) |
| 2026-05-31T08:04:27Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $1766 < absolute min $5000 |
| 2026-05-31T08:04:27Z | us-announces-new-iran-agreementceasefire-extension-by-may-31-665-831-238 | geopolitics | 0.0430 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 5 veces) |
| 2026-05-31T08:30:03Z | jeffrey-epstein-foul-play-confirmed-by-december-31-2026 | other | 0.0700 | P2 | mercado ya expiró (endDate=2025-12-31T00:00:00Z, hace 151 días) |
| 2026-05-31T08:30:03Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon -151 d, liq $40093) |
| 2026-05-31T08:30:03Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 120 días) |
| 2026-05-31T08:30:03Z | will-chinas-2026-annual-gdp-growth-yy-be-between-4pt0-and-5pt0 | other | 0.7600 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 120 días) |
| 2026-05-31T08:30:03Z | will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -112 d, liq $93218) |
| 2026-05-31T08:30:03Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -97 d, liq $766688) |
| 2026-05-31T08:30:03Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -97 d, liq $775854) |
| 2026-05-31T08:30:03Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -70 d, liq $344379) |
| 2026-05-31T08:30:03Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $28575) |
| 2026-05-31T08:30:03Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -68 d, liq $20147) |
| 2026-05-31T08:30:03Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -68 d, liq $28084) |
| 2026-05-31T08:30:03Z | will-alex-vanopslagh-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -68 d, liq $30124) |
| 2026-05-31T08:30:03Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -48 d, liq $345248) |
| 2026-05-31T08:30:03Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon -16 d, liq $115569) |
| 2026-05-31T08:30:03Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -14 d, liq $5487) |
| 2026-05-31T08:30:03Z | will-gregg-kirkpatrick-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -12 d, liq $10504) |
| 2026-05-31T08:30:03Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -12 d, liq $8484) |
| 2026-05-31T08:30:03Z | bra-bah-bot-2026-05-30-bot | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $501792) |
| 2026-05-31T08:30:03Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.3000 | P9 | P9: geopolitics pump cluster (price 0.30, 0d) |
| 2026-05-31T08:30:03Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $220527) |
| 2026-05-31T08:30:03Z | will-alibaba-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $136031) |
| 2026-05-31T08:30:03Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $166662) |
| 2026-05-31T08:30:03Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $290498) |
| 2026-05-31T08:30:03Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $162252) |
| 2026-05-31T08:30:03Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.0240 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-31T08:30:03Z | will-80-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 0 d, liq $88761) |
| 2026-05-31T08:30:03Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0470 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-31T08:30:03Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 0 d, liq $896941) |
| 2026-05-31T08:30:03Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $197966) |
| 2026-05-31T08:30:03Z | will-finland-win-the-2026-iihf-world-championship | sports-season | 0.4700 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:30:03Z | netanyahu-out-by-may-31 | executive-action | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $68013) |
| 2026-05-31T08:30:03Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.0380 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-31T08:30:03Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $3455925) |
| 2026-05-31T08:30:03Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 0 d, liq $58283) |
| 2026-05-31T08:30:03Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 0 d, liq $1133299) |
| 2026-05-31T08:30:03Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 0 d, liq $109557) |
| 2026-05-31T08:30:03Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 0 d, liq $119382) |
| 2026-05-31T08:30:03Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $127239) |
| 2026-05-31T08:30:03Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $91327) |
| 2026-05-31T08:30:04Z | bitcoin-above-78k-on-may-31-2026 | market | 0.0020 | E2 | edge 0.008 < mín 0.040 (p̂=0.010, implied=0.002) |
| 2026-05-31T08:30:04Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $120049) |
| 2026-05-31T08:30:04Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $95302) |
| 2026-05-31T08:30:04Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $111027) |
| 2026-05-31T08:30:06Z | will-ivn-cepeda-castro-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.7600 | N1 | noticias contradicen tesis: Cepeda lidera encuestas con ~38%, muy por debajo del 50%+1 requerido para ganar en primer... |
| 2026-05-31T08:30:06Z | bitcoin-above-72k-on-may-31-2026 | market | 0.9900 | E2 | edge 0.027 < mín 0.040 (p̂=0.963, implied=0.990) |
| 2026-05-31T08:30:06Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $142007) |
| 2026-05-31T08:30:06Z | bitcoin-above-70k-on-may-31-2026 | market | 0.9980 | E2 | edge 0.008 < mín 0.040 (p̂=0.990, implied=0.998) |
| 2026-05-31T08:30:06Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $110138) |
| 2026-05-31T08:30:06Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $93225) |
| 2026-05-31T08:30:06Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $151498) |
| 2026-05-31T08:30:06Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $151681) |
| 2026-05-31T08:30:06Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $137226) |
| 2026-05-31T08:30:06Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $141881) |
| 2026-05-31T08:30:08Z | will-abelardo-de-la-espriella-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.2450 | N1 | noticias contradicen tesis: Encuestas recientes sitúan a De la Espriella 2º con ~28-31%, lejos del 50%+ necesario p... |
| 2026-05-31T08:30:08Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $406452) |
| 2026-05-31T08:30:08Z | bitcoin-above-76k-on-may-31-2026 | market | 0.0100 | E2 | edge 0.012 < mín 0.040 (p̂=0.022, implied=0.010) |
| 2026-05-31T08:30:08Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $352207) |
| 2026-05-31T08:30:08Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.0010 | R1_longshot | R1 longshot: precio 0.0010 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.482 < 0.15) |
| 2026-05-31T08:30:08Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0010 | E2 | edge 0.023 < mín 0.040 (p̂=0.024, implied=0.001) |
| 2026-05-31T08:30:08Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | will-ethereum-reach-3600-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | will-solana-reach-140-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | over-60m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $14023) |
| 2026-05-31T08:30:09Z | over-40m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17262) |
| 2026-05-31T08:30:09Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | will-xrp-reach-2pt8-in-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17263) |
| 2026-05-31T08:30:09Z | will-solana-reach-110-in-may-2026 | market | 0.0020 | E2 | edge 0.008 < mín 0.040 (p̂=0.010, implied=0.002) |
| 2026-05-31T08:30:09Z | will-ethereum-reach-3000-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | will-bitcoin-reach-80000-in-may-2026-from-may-27 | market | 0.0020 | E2 | edge 0.008 < mín 0.040 (p̂=0.010, implied=0.002) |
| 2026-05-31T08:30:09Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $128777) |
| 2026-05-31T08:30:09Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 0 d, liq $92401) |
| 2026-05-31T08:30:09Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | over-30m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $11419) |
| 2026-05-31T08:30:09Z | will-bitcoin-reach-80k-may-25-31-2026 | market | 0.0060 | E2 | edge 0.004 < mín 0.040 (p̂=0.010, implied=0.006) |
| 2026-05-31T08:30:09Z | will-solana-reach-130-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:09Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $27212) |
| 2026-05-31T08:30:11Z | elon-musk-of-tweets-may-2026-800-839 | other | 0.2100 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-31T08:30:11Z | elon-musk-of-tweets-may-2026-880-919 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 0 d, liq $9173) |
| 2026-05-31T08:30:11Z | will-ethereum-reach-3800-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:11Z | will-bitcoin-dip-to-65k-in-may-2026-183-857-425 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $88492) |
| 2026-05-31T08:30:11Z | will-bitcoin-dip-to-35k-in-may-2026-217-769-834 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $90133) |
| 2026-05-31T08:30:11Z | will-ethereum-reach-2200-may-25-31-2026 | market | 0.0020 | E2 | edge 0.008 < mín 0.040 (p̂=0.010, implied=0.002) |
| 2026-05-31T08:30:11Z | elon-musk-of-tweets-may-2026-1080-1119 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $6065) |
| 2026-05-31T08:30:11Z | will-ethereum-reach-3200-in-may-2026 | market | 0.0020 | E2 | edge 0.008 < mín 0.040 (p̂=0.010, implied=0.002) |
| 2026-05-31T08:30:11Z | will-bitcoin-dip-to-72500-in-may-2026-from-may-27 | market | 0.1700 | M2 | M2 soft-learned: market·short·0.10-0.30 = 5L/0W (wr 0%) |
| 2026-05-31T08:30:11Z | will-solana-reach-120-in-may-2026 | market | 0.0010 | E2 | edge 0.009 < mín 0.040 (p̂=0.010, implied=0.001) |
| 2026-05-31T08:30:11Z | will-ethereum-dip-to-1200-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17154) |
| 2026-05-31T08:30:11Z | will-the-ethereum-volatility-index-dip-to-40-by-may-31 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $714) |
| 2026-05-31T08:30:11Z | will-the-breadwinner-opening-weekend-box-office-be-greater-than-7m | entertainment | 0.9810 | P0_ceiling | price ceiling: 0.9810 > 0.980 |
| 2026-05-31T08:30:13Z | will-backrooms-opening-weekend-box-office-be-greater-than-79m | entertainment | 0.7640 | N1 | noticias contradicen tesis: Backrooms abrió con $85M-$89M según Deadline/THR, superando el umbral de $79M; lado 'no... |
| 2026-05-31T08:30:13Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $37614) |
| 2026-05-31T08:30:15Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2400 | R5_precedents | R5 precedentes: 1 < 3 casos análogos |
| 2026-05-31T08:30:15Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $18907) |
| 2026-05-31T08:30:15Z | elon-musk-of-tweets-may-26-june-2-300-319 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $72231) |
| 2026-05-31T08:30:15Z | elon-musk-of-tweets-may-26-june-2-100-119 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 2 d, liq $42590) |
| 2026-05-31T08:30:15Z | elon-musk-of-tweets-may-26-june-2-500plus | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $118298) |
| 2026-05-31T08:30:15Z | elon-musk-of-tweets-may-26-june-2-260-279 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $50097) |
| 2026-05-31T08:30:15Z | elon-musk-of-tweets-may-26-june-2-380-399 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $65508) |
| 2026-05-31T08:30:18Z | elon-musk-of-tweets-may-26-june-2-180-199 | other | 0.1100 | E2 | edge 0.040 < mín 0.040 (p̂=0.150, implied=0.110) |
| 2026-05-31T08:30:18Z | elon-musk-of-tweets-may-26-june-2-360-379 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $63824) |
| 2026-05-31T08:30:18Z | elon-musk-of-tweets-may-26-june-2-280-299 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $41923) |
| 2026-05-31T08:30:18Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $30795) |
| 2026-05-31T08:30:20Z | will-won-hee-ryong-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $35274) |
| 2026-05-31T08:30:20Z | will-han-dong-hoon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $33824) |
| 2026-05-31T08:30:20Z | will-lee-un-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $42337) |
| 2026-05-31T08:30:20Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $35723) |
| 2026-05-31T08:30:20Z | will-yoon-sang-hyun-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $38647) |
| 2026-05-31T08:30:20Z | will-na-kyung-won-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $477400) |
| 2026-05-31T08:30:20Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $33278) |
| 2026-05-31T08:30:20Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $44295) |
| 2026-05-31T08:30:20Z | will-chung-jin-suk-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $20247) |
| 2026-05-31T08:30:20Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $21834) |
| 2026-05-31T08:30:20Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $44795) |
| 2026-05-31T08:30:20Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $44928) |
| 2026-05-31T08:30:20Z | will-lee-chul-gyu-win-the-2026-gangwon-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $15648) |
| 2026-05-31T08:30:20Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $46041) |
| 2026-05-31T08:30:20Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $97197) |
| 2026-05-31T08:30:20Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.1700 | P9 | P9: geopolitics pump cluster (price 0.17, 2d) |
| 2026-05-31T08:30:20Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $25972) |
| 2026-05-31T08:30:20Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $31888) |
| 2026-05-31T08:30:20Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $94863) |
| 2026-05-31T08:30:20Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $45450) |
| 2026-05-31T08:30:22Z | will-oh-se-hoon-win-the-2026-seoul-mayoral-election | elections | 0.1900 | N1 | noticias contradicen tesis: Encuestas de 26-28 mayo muestran a Jung Won-oh liderando (41-49,6%) frente a Oh Se-hoon (... |
| 2026-05-31T08:30:22Z | elon-musk-of-tweets-may-29-june-5-500plus | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 5 d, liq $43688) |
| 2026-05-31T08:30:22Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $209356) |
| 2026-05-31T08:30:22Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $119669) |
| 2026-05-31T08:30:25Z | will-rafael-lpez-aliaga-win-the-2026-peruvian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 6 d, liq $136272) |
| 2026-05-31T08:30:25Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $213196) |
| 2026-05-31T08:30:27Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0280 | E2 | edge 0.008 < mín 0.040 (p̂=0.020, implied=0.028) |
| 2026-05-31T08:30:27Z | will-ricardo-belmont-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $94397) |
| 2026-05-31T08:30:27Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $220011) |
| 2026-05-31T08:30:27Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $211847) |
| 2026-05-31T08:30:29Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0850 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que afecte significat... |
| 2026-05-31T08:30:29Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $189583) |
| 2026-05-31T08:30:29Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $159686) |
| 2026-05-31T08:30:29Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $240495) |
| 2026-05-31T08:30:29Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.1300 | P9 | P9: geopolitics pump cluster (price 0.13, 6d) |
| 2026-05-31T08:30:29Z | will-jorge-nieto-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $141602) |
| 2026-05-31T08:30:29Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $226997) |
| 2026-05-31T08:30:29Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $107877) |
| 2026-05-31T08:30:31Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.1080 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda influir en la probabilid... |
| 2026-05-31T08:30:31Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $247123) |
| 2026-05-31T08:30:31Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $201302) |
| 2026-05-31T08:30:31Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $159866) |
| 2026-05-31T08:30:31Z | will-raphael-collignon-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $356058) |
| 2026-05-31T08:30:31Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-7-2026 | geopolitics | 0.3000 | P9 | P9: geopolitics pump cluster (price 0.30, 6d) |
| 2026-05-31T08:30:31Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $197648) |
| 2026-05-31T08:30:36Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $244383) |
| 2026-05-31T08:30:36Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $221717) |
| 2026-05-31T08:30:36Z | will-a-world-cup-game-in-mexico-be-relocated-abroad | sports-season | 0.0390 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:30:36Z | strait-of-hormuz-traffic-returns-to-normal-by-june-15 | geopolitics | 0.0700 | P9 | P9: geopolitics pump cluster (price 0.07, 14d) |
| 2026-05-31T08:30:36Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.1900 | P9 | P9: geopolitics pump cluster (price 0.19, 14d) |
| 2026-05-31T08:30:36Z | will-the-oklahoma-city-thunder-win-the-nba-western-conference-finals | sports-season | 0.0010 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:30:36Z | will-al-mina-be-the-republican-nominee-for-senate-in-virginia | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 15 d, liq $13950) |
| 2026-05-31T08:30:36Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9820 | P0_ceiling | price ceiling: 0.9820 > 0.980 |
| 2026-05-31T08:30:36Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 16 d, liq $839854) |
| 2026-05-31T08:30:36Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 16 d, liq $297037) |
| 2026-05-31T08:30:36Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 16 d, liq $2623211) |
| 2026-05-31T08:30:36Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 16 d, liq $1011480) |
| 2026-05-31T08:30:36Z | will-flyquest-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 20 d, liq $85584) |
| 2026-05-31T08:30:36Z | will-nrg-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 20 d, liq $88396) |
| 2026-05-31T08:30:36Z | will-gaimin-gladiators-win-iem-cologne-major-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 20 d, liq $87845) |
| 2026-05-31T08:30:36Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $219847) |
| 2026-05-31T08:30:38Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $189407) |
| 2026-05-31T08:30:38Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $207212) |
| 2026-05-31T08:30:38Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $229978) |
| 2026-05-31T08:30:41Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6500 | N1 | noticias contradicen tesis: Encuestas lo sitúan en 21.5% vs 44.3% de Cepeda; muy improbable ganar en primera vuelta ... |
| 2026-05-31T08:30:41Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $200808) |
| 2026-05-31T08:30:41Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 21 d, liq $58443) |
| 2026-05-31T08:30:41Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 21 d, liq $111336) |
| 2026-05-31T08:30:43Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-31T08:30:43Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 23 d, liq $655) |
| 2026-05-31T08:30:43Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 23 d, liq $809) |
| 2026-05-31T08:30:43Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 23 d, liq $873) |
| 2026-05-31T08:30:43Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 26 d, liq $1574) |
| 2026-05-31T08:30:43Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-june-30-612 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 29 d, liq $268949) |
| 2026-05-31T08:30:43Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 29 d, liq $378671) |
| 2026-05-31T08:30:43Z | tom-holland-announced-as-next-james-bond-679 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 29 d, liq $5281) |
| 2026-05-31T08:30:44Z | trump-renames-ice-to-nice-by-june-30 | other | 0.0410 | V6 | V6: No hay ningún evento discreto verificable en los próximos 7 días que pueda catalizar un cambio en el precio. L... |
| 2026-05-31T08:30:46Z | us-embassy-in-damaskus-reopend-by-june-30-2026-815 | other | 0.0200 | E2 | edge 0.030 < mín 0.040 (p̂=0.050, implied=0.020) |
| 2026-05-31T08:30:46Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 29 d, liq $16708) |
| 2026-05-31T08:30:46Z | will-trump-and-putin-meet-next-in-belarus-572 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 29 d, liq $18815) |
| 2026-05-31T08:30:46Z | israel-withdraws-from-lebanon-by-may-31-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 29 d, liq $99726) |
| 2026-05-31T08:30:46Z | will-trump-and-putin-meet-next-in-united-states-224-875-469 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 29 d, liq $19398) |
| 2026-05-31T08:30:48Z | mojtaba-khamenei-leaves-iran-by-june-30-2026 | geopolitics | 0.0280 | E2 | edge 0.018 < mín 0.040 (p̂=0.010, implied=0.028) |
| 2026-05-31T08:30:48Z | jack-lowdon-announced-as-next-james-bond | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 29 d, liq $9485) |
| 2026-05-31T08:30:50Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0300 | R1_longshot | R1 longshot: precio 0.0300 < 0.10 sin edge fuerte (edge_type=info, edge=0.050 < 0.15) |
| 2026-05-31T08:30:50Z | will-trump-and-putin-meet-next-in-turkey-182-161-789 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 29 d, liq $18343) |
| 2026-05-31T08:30:50Z | anthropic-ceo-arrested | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 29 d, liq $8560) |
| 2026-05-31T08:30:52Z | iran-closes-its-airspace-by-june-30-432-786-462-866 | geopolitics | 0.2380 | N1 | noticias contradicen tesis: Tregua EEUU-Irán vigente y MOU de 60 días acordado el 28-may; Iran en reapertura parcia... |
| 2026-05-31T08:30:52Z | tom-hardy-announced-as-next-james-bond-996-479 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 29 d, liq $8066) |
| 2026-05-31T08:30:52Z | will-trump-and-putin-meet-next-in-south-korea-582-297-268 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 29 d, liq $22286) |
| 2026-05-31T08:30:54Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0330 | E2 | edge 0.017 < mín 0.040 (p̂=0.050, implied=0.033) |
| 2026-05-31T08:30:54Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 29 d, liq $354933) |
| 2026-05-31T08:30:56Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La confirmación de una visita de... |
| 2026-05-31T08:30:58Z | us-obtains-iranian-enriched-uranium-by-june-30 | other | 0.0380 | E2 | edge 0.002 < mín 0.040 (p̂=0.040, implied=0.038) |
| 2026-05-31T08:30:58Z | will-trump-and-putin-meet-next-in-finland-772-412 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 29 d, liq $17185) |
| 2026-05-31T08:30:58Z | will-amazon-be-the-largest-company-in-the-world-by-market-cap-on-june-30-123 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 29 d, liq $395329) |
| 2026-05-31T08:30:58Z | will-bill-gates-be-confirmed-to-have-visited-epsteins-island | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 29 d, liq $23104) |
| 2026-05-31T08:30:58Z | will-trump-and-putin-meet-next-in-another-country-313-781-734-447 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 29 d, liq $20662) |
| 2026-05-31T08:31:01Z | will-apple-be-the-largest-company-in-the-world-by-market-cap-on-june-30-416 | other | 0.0380 | R1_longshot | R1 longshot: precio 0.0380 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.082 < 0.15) |
| 2026-05-31T08:31:01Z | will-trump-and-putin-meet-next-in-another-european-country-954-837-364 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 29 d, liq $21824) |
| 2026-05-31T08:31:03Z | will-discords-market-cap-be-less-than-15b-at-market-close-on-ipo-day | crypto-launch | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 29 d, liq $3478) |
| 2026-05-31T08:31:03Z | will-michael-jackson-be-confirmed-to-have-visited-epsteins-island | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 29 d, liq $21651) |
| 2026-05-31T08:31:03Z | will-the-next-diplomatic-us-iran-meeting-be-in-italy-332 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 29 d, liq $14932) |
| 2026-05-31T08:31:04Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-june | geopolitics | 0.2900 | N1 | noticias contradicen tesis: Tráfico al 5% pre-guerra; ataques EEUU 28-may reaviván tensiones; DHL estima 4-6 meses ... |
| 2026-05-31T08:31:04Z | us-announces-new-iran-agreementceasefire-extension-by-june-30-848-925 | geopolitics | 0.6400 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 5 veces) |
| 2026-05-31T08:31:04Z | will-trump-and-putin-meet-next-in-a-gulf-country-733-116 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 29 d, liq $22200) |
| 2026-05-31T08:31:06Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0220 | E2 | edge 0.019 < mín 0.040 (p̂=0.003, implied=0.022) |
| 2026-05-31T08:31:08Z | us-x-cuba-diplomatic-meeting-by-may-31-265 | other | 0.9640 | V3 Trigger vago, V6 Sin catalyst | V3 Trigger vago, V6 Sin catalyst: V3 Trigger vago: la pregunta es genérica sin especificar detalles concretos (como ... |
| 2026-05-31T08:31:08Z | will-jpmorgan-chase-fail-by-june-30-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 29 d, liq $4129) |
| 2026-05-31T08:31:11Z | pete-hegseth-impeached-by-june-30 | other | 0.0360 | E2 | edge 0.016 < mín 0.040 (p̂=0.020, implied=0.036) |
| 2026-05-31T08:31:11Z | lai-ching-te-impeached-by-june-30 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 29 d, liq $22499) |
| 2026-05-31T08:31:11Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 29 d, liq $89487) |
| 2026-05-31T08:31:11Z | will-any-country-leave-nato-by-june-30-2026 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 30 d, liq $23506) |
| 2026-05-31T08:31:11Z | will-silver-si-hit-low-45-by-end-of-june-972-272 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 30 d, liq $14850) |
| 2026-05-31T08:31:11Z | will-crude-oil-cl-hit-low-50-by-end-of-june | market | 0.0150 | E2 | edge 0.030 < mín 0.040 (p̂=0.045, implied=0.015) |
| 2026-05-31T08:31:14Z | kharg-island-no-longer-under-iranian-control-by-june-30-561-854 | other | 0.0480 | E2 | edge 0.018 < mín 0.040 (p̂=0.030, implied=0.048) |
| 2026-05-31T08:31:14Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.6440 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:14Z | will-the-new-york-knicks-win-the-2026-nba-finals | sports-season | 0.3530 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:15Z | microstrategy-sells-any-bitcoin-by-may-31-2026 | market | 0.1550 | N1 | noticias contradicen tesis: Sin ventas confirmadas de BTC en 2026; Strategy anunció condiciones para vender el 5-may... |
| 2026-05-31T08:31:17Z | microstrategy-sells-any-bitcoin-by-june-30-2026 | market | 0.7250 | R5_precedents | R5 precedentes: 0 < 3 casos análogos |
| 2026-05-31T08:31:17Z | will-bitcoin-hit-150k-by-june-30-2026 | market | 0.0070 | P6 | P6 market: BTC-USD spot $73836.87 already > target $150.00 but yes=0.01 (confused book) |
| 2026-05-31T08:31:17Z | will-a-team-from-lcp-asia-pacific-win-msi-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 41 d, liq $8005) |
| 2026-05-31T08:31:17Z | will-neymar-play-in-the-world-cup | sports-season | 0.9000 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-haiti-win-the-2026-fifa-world-cup | sports-season | 0.0010 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-england-win-the-2026-fifa-world-cup-937 | sports-season | 0.1130 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-new-zealand-win-the-2026-fifa-world-cup-635 | sports-season | 0.0010 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0900 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-france-win-the-2026-fifa-world-cup-924 | sports-season | 0.1710 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0870 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0110 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0420 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-spain-win-the-2026-fifa-world-cup-963 | sports-season | 0.1680 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0280 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-portugal-win-the-2026-fifa-world-cup-912 | sports-season | 0.1000 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0020 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:17Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:19Z | us-x-iran-permanent-peace-deal-by-july-31-2026-831-252 | geopolitics | 0.4300 | V5 | V5: V5 Patrón débil: menos de 3 fuentes independientes y sin precedente análogo de acuerdo de paz permanente entre... |
| 2026-05-31T08:31:19Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 60 d, liq $32906) |
| 2026-05-31T08:31:19Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 60 d, liq $13777) |
| 2026-05-31T08:31:21Z | strait-of-hormuz-traffic-returns-to-normal-by-july-31 | geopolitics | 0.4800 | N1 | noticias contradicen tesis: Acuerdo EE.UU.-Irán 'casi negociado' que reabrirá el estrecho en 30 días; podría norm... |
| 2026-05-31T08:31:21Z | will-filipe-luis-be-the-next-manager-of-ssc-napoli | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 92 d, liq $376) |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-indianapolis-colts-in-2026-27 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 92 d, liq $517) |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-jacksonville-jaguars-in-2026-27 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 92 d, liq $991) |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-miami-dolphins-in-2026-27 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 92 d, liq $3247) |
| 2026-05-31T08:31:21Z | will-george-pickens-play-for-cleveland-browns-in-2026-27 | other | 0.8740 | P3_low_absolute_liquidity | liquidity $15 < absolute min $5000 |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-new-orleans-saints-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 92 d, liq $2566) |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-denver-broncos-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 92 d, liq $1508) |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-kansas-city-chiefs-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 92 d, liq $3047) |
| 2026-05-31T08:31:21Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 92 d, liq $1081) |
| 2026-05-31T08:31:23Z | will-cdu-win-the-most-seats-in-the-2026-sachsen-anhalt-parliamentary-elections | other | 0.0570 | E2 | edge 0.023 < mín 0.040 (p̂=0.080, implied=0.057) |
| 2026-05-31T08:31:23Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 104 d, liq $2463) |
| 2026-05-31T08:31:23Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 104 d, liq $1959) |
| 2026-05-31T08:31:23Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 104 d, liq $2303) |
| 2026-05-31T08:31:23Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 104 d, liq $3153) |
| 2026-05-31T08:31:23Z | will-ebba-busch-be-the-next-prime-minister-of-sweden | elections | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 104 d, liq $25265) |
| 2026-05-31T08:31:23Z | will-alexandra-eala-win-the-2026-womens-us-open | other | 0.0340 | P3_low_absolute_liquidity | liquidity $1278 < absolute min $5000 |
| 2026-05-31T08:31:23Z | will-anastasia-potapova-win-the-2026-womens-us-open | other | 0.0610 | P3_low_absolute_liquidity | liquidity $2318 < absolute min $5000 |
| 2026-05-31T08:31:26Z | will-the-sweden-democrats-sd-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0450 | E2 | edge 0.035 < mín 0.040 (p̂=0.080, implied=0.045) |
| 2026-05-31T08:31:26Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 104 d, liq $3201) |
| 2026-05-31T08:31:26Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 104 d, liq $2556) |
| 2026-05-31T08:31:26Z | will-barbora-krejcikova-win-the-2026-womens-us-open | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 104 d, liq $2198) |
| 2026-05-31T08:31:28Z | will-jimmie-kesson-be-the-next-prime-minister-of-sweden | elections | 0.0280 | E2 | edge 0.022 < mín 0.040 (p̂=0.050, implied=0.028) |
| 2026-05-31T08:31:30Z | will-the-moderate-party-m-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0360 | R1_longshot | R1 longshot: precio 0.0360 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.044 < 0.15) |
| 2026-05-31T08:31:30Z | will-tereza-valentova-win-the-2026-womens-us-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 104 d, liq $2293) |
| 2026-05-31T08:31:30Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 119 d, liq $1384) |
| 2026-05-31T08:31:30Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $248 < absolute min $5000 |
| 2026-05-31T08:31:30Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0290 | P3_low_absolute_liquidity | liquidity $976 < absolute min $5000 |
| 2026-05-31T08:31:30Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $973 < absolute min $5000 |
| 2026-05-31T08:31:30Z | will-carlos-roberto-massa-jnior-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 125 d, liq $28149) |
| 2026-05-31T08:31:30Z | will-geraldo-alckmin-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 125 d, liq $18332) |
| 2026-05-31T08:31:30Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 125 d, liq $1025962) |
| 2026-05-31T08:31:32Z | will-romeu-zema-win-the-2026-brazilian-presidential-election | elections | 0.0270 | E2 | edge 0.007 < mín 0.040 (p̂=0.020, implied=0.027) |
| 2026-05-31T08:31:34Z | will-renan-santos-win-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0910 | R1_longshot | R1 longshot: precio 0.0910 < 0.10 sin edge fuerte (edge_type=info, edge=0.071 < 0.15) |
| 2026-05-31T08:31:34Z | will-michelle-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 125 d, liq $232479) |
| 2026-05-31T08:31:36Z | will-renan-santos-win-the-2026-brazilian-presidential-election | elections | 0.1700 | E2 | edge 0.020 < mín 0.040 (p̂=0.150, implied=0.170) |
| 2026-05-31T08:31:36Z | will-jair-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 125 d, liq $310516) |
| 2026-05-31T08:31:38Z | will-romeu-zema-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0370 | E2 | edge 0.017 < mín 0.040 (p̂=0.020, implied=0.037) |
| 2026-05-31T08:31:38Z | will-parti-conservateur-du-qubec-win-the-most-seats-in-the-2026-quebec-general-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 126 d, liq $15289) |
| 2026-05-31T08:31:42Z | will-ousmane-dembl-win-the-2026-ballon-dor | other | 0.2800 | N1 | noticias contradicen tesis: Kane lidera apuestas (5/2); Dembélé solo al 24% en Kalshi. Carrera abierta con Rice, Ya... |
| 2026-05-31T08:31:42Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-st-louis-cardinals-win-the-2026-national-league-championship-series | sports-season | 0.0140 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-pittsburgh-pirates-win-the-2026-national-league-championship-series | sports-season | 0.0410 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-colorado-rockies-win-the-2026-national-league-championship-series | sports-season | 0.0030 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-san-diego-padres-win-the-2026-national-league-championship-series | sports-season | 0.0550 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-cincinnati-reds-win-the-2026-national-league-championship-series | sports-season | 0.0260 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:31:42Z | will-betty-yee-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 155 d, liq $228852) |
| 2026-05-31T08:31:44Z | will-bernadette-wilson-win-the-2026-alaska-governor-election | elections | 0.2100 | N1 | noticias contradicen tesis: Wilson segunda en sondeos primarios (15.8%) tras Begich (20.8%); RCP y Polymarket no la s... |
| 2026-05-31T08:31:44Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 155 d, liq $6304) |
| 2026-05-31T08:31:44Z | will-corey-seager-win-the-2026-american-league-hank-aaron-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 165 d, liq $934) |
| 2026-05-31T08:31:44Z | will-richard-van-de-water-win-the-bachelorette-season-22 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 182 d, liq $3729) |
| 2026-05-31T08:31:44Z | will-lew-evans-win-the-bachelorette-season-22 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 182 d, liq $6964) |
| 2026-05-31T08:31:46Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-4pt0-at-the-end-of-2026-386 | other | 0.2530 | N1 | noticias contradicen tesis: El tipo actual está en 3,75% (techo) y el mercado espera recortes o pausa; subir a 4,0% ... |
| 2026-05-31T08:31:46Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 191 d, liq $12478) |
| 2026-05-31T08:31:46Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 191 d, liq $12289) |
| 2026-05-31T08:31:48Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-3pt5-at-the-end-of-2026-352 | other | 0.0960 | E2 | edge 0.024 < mín 0.040 (p̂=0.120, implied=0.096) |
| 2026-05-31T08:31:48Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 191 d, liq $9151) |
| 2026-05-31T08:31:51Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-3pt0-at-the-end-of-2026-593 | other | 0.0380 | E2 | edge 0.012 < mín 0.040 (p̂=0.050, implied=0.038) |
| 2026-05-31T08:31:51Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 191 d, liq $9083) |
| 2026-05-31T08:31:51Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 191 d, liq $10087) |
| 2026-05-31T08:31:51Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 191 d, liq $9994) |
| 2026-05-31T08:31:51Z | will-brice-matthews-win-the-2026-al-rookie-of-the-year-award | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 201 d, liq $4418) |
| 2026-05-31T08:31:51Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 201 d, liq $6425) |
| 2026-05-31T08:31:54Z | will-nashville-sc-win-the-2026-mls-cup | other | 0.0580 | E2 | edge 0.022 < mín 0.040 (p̂=0.080, implied=0.058) |
| 2026-05-31T08:31:54Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 213 d, liq $7974) |
| 2026-05-31T08:31:54Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 213 d, liq $147) |
| 2026-05-31T08:31:54Z | will-gholam-ali-haddad-adel-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 213 d, liq $57392) |
| 2026-05-31T08:31:56Z | will-israel-strike-10-countries-in-2026 | geopolitics | 0.0260 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La pregunta es extremadamente amp... |
| 2026-05-31T08:31:56Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 213 d, liq $9637) |
| 2026-05-31T08:31:56Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 213 d, liq $16407) |
| 2026-05-31T08:31:58Z | will-zelenskyy-and-putin-meet-next-in-qatar-uae | other | 0.0470 | E2 | edge 0.017 < mín 0.040 (p̂=0.030, implied=0.047) |
| 2026-05-31T08:31:58Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 213 d, liq $1308) |
| 2026-05-31T08:31:58Z | will-there-be-between-11-and-13-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $1503 < absolute min $5000 |
| 2026-05-31T08:31:58Z | will-ahmad-vahidi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 213 d, liq $58498) |
| 2026-05-31T08:31:58Z | will-israel-strike-9-countries-in-2026 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 213 d, liq $30981) |
| 2026-05-31T08:31:58Z | jerome-powell-out-from-fed-board-by-may-30 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 213 d, liq $62629) |
| 2026-05-31T08:31:58Z | will-ali-asghar-hejazi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 213 d, liq $54965) |
| 2026-05-31T08:31:59Z | will-abbas-araghchi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0300 | E2 | edge 0.020 < mín 0.040 (p̂=0.050, implied=0.030) |
| 2026-05-31T08:31:59Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 213 d, liq $687) |
| 2026-05-31T08:32:02Z | will-the-us-invade-iran-before-2027 | geopolitics | 0.1900 | R5_precedents | R5 precedentes: 0 < 3 casos análogos |
| 2026-05-31T08:32:04Z | will-zelenskyy-and-putin-meet-next-in-us | other | 0.0430 | E2 | edge 0.023 < mín 0.040 (p̂=0.020, implied=0.043) |
| 2026-05-31T08:32:04Z | will-david-lammy-be-the-next-prime-minister-of-the-united-kingdom-in-2026-654 | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 213 d, liq $49541) |
| 2026-05-31T08:32:04Z | will-there-be-between-17-and-19-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $3384 < absolute min $5000 |
| 2026-05-31T08:32:04Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 213 d, liq $6452) |
| 2026-05-31T08:32:04Z | will-nir-barkat-be-the-next-prime-minister-of-israel | elections | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 213 d, liq $67081) |
| 2026-05-31T08:32:04Z | will-germany-recognize-palestine-before-2027 | executive-action | 0.0900 | P3_low_absolute_liquidity | liquidity $4079 < absolute min $5000 |
| 2026-05-31T08:32:07Z | will-2026-rank-as-the-sixth-hottest-year-on-record-or-lower | other | 0.0230 | R1_longshot | R1 longshot: precio 0.0230 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.077 < 0.15) |
| 2026-05-31T08:32:07Z | databricks-ipo-before-2027 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $2493 < absolute min $5000 |
| 2026-05-31T08:32:07Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 213 d, liq $442) |
| 2026-05-31T08:32:07Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:32:07Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $211 < absolute min $5000 |
| 2026-05-31T08:32:07Z | will-maryam-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 213 d, liq $60548) |
| 2026-05-31T08:32:07Z | will-there-be-exactly-0-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026 | weather-natural | 0.6800 | P3_low_absolute_liquidity | liquidity $1805 < absolute min $5000 |
| 2026-05-31T08:32:09Z | iran-nuclear-test-before-2027 | geopolitics | 0.0800 | E2 | edge 0.020 < mín 0.040 (p̂=0.060, implied=0.080) |
| 2026-05-31T08:32:09Z | vanta-ipo-before-2027 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $3782 < absolute min $5000 |
| 2026-05-31T08:32:09Z | will-meta-acquire-tiktok-745-612-641 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 213 d, liq $11103) |
| 2026-05-31T08:32:09Z | new-covid-variant-of-concern-before-2027 | other | 0.1700 | P3_low_absolute_liquidity | liquidity $2171 < absolute min $5000 |
| 2026-05-31T08:32:09Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1540 < absolute min $5000 |
| 2026-05-31T08:32:09Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 213 d, liq $12323) |
| 2026-05-31T08:32:11Z | will-israel-strike-8-countries-in-2026 | geopolitics | 0.0390 | V6 - Sin catalyst | V6 - Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La resolución dep... |
| 2026-05-31T08:32:11Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 213 d, liq $698) |
| 2026-05-31T08:32:14Z | will-toy-story-5-have-the-best-domestic-opening-weekend-in-2026 | entertainment | 0.0680 | R1_longshot | R1 longshot: precio 0.0680 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.182 < 0.15) |
| 2026-05-31T08:32:14Z | will-there-be-between-14-and-16-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2148 < absolute min $5000 |
| 2026-05-31T08:32:15Z | will-israel-strike-11-countries-in-2026 | geopolitics | 0.0250 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La pregunta es vaga ... |
| 2026-05-31T08:32:18Z | will-china-invade-taiwan-before-2027 | geopolitics | 0.0660 | R1_longshot | R1 longshot: precio 0.0660 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.046 < 0.15) |
| 2026-05-31T08:32:18Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 213 d, liq $200) |
| 2026-05-31T08:32:18Z | will-mohammad-khatami-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 213 d, liq $53992) |
| 2026-05-31T08:32:18Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 213 d, liq $638) |
| 2026-05-31T08:32:20Z | will-israel-strike-7-countries-in-2026 | geopolitics | 0.0430 | E2 | edge 0.023 < mín 0.040 (p̂=0.020, implied=0.043) |
| 2026-05-31T08:32:20Z | will-massoud-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 213 d, liq $58482) |
| 2026-05-31T08:32:22Z | will-zelenskyy-and-putin-meet-next-in-turkey | other | 0.0270 | E2 | edge 0.003 < mín 0.040 (p̂=0.030, implied=0.027) |
| 2026-05-31T08:32:22Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0800 | P3_low_absolute_liquidity | liquidity $3267 < absolute min $5000 |
| 2026-05-31T08:32:25Z | will-jensen-huang-be-richest-person-on-december-31-229 | other | 0.0380 | E2 | edge 0.002 < mín 0.040 (p̂=0.040, implied=0.038) |
| 2026-05-31T08:32:25Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 213 d, liq $509) |
| 2026-05-31T08:32:25Z | will-vladimir-padrino-lpez-be-the-leader-of-venezuela-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 213 d, liq $65052) |
| 2026-05-31T08:32:25Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 213 d, liq $11667) |
| 2026-05-31T08:32:25Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 213 d, liq $19058) |
| 2026-05-31T08:32:25Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1160 | P3_low_absolute_liquidity | liquidity $3874 < absolute min $5000 |
| 2026-05-31T08:32:27Z | insurrection-act-invoked-by-december-31-184-556 | other | 0.2100 | N1 | noticias contradicen tesis: Trump amenazó con invocarla en enero 2026 (Minnesota/ICE) pero no lo hizo; sin evidencia... |
| 2026-05-31T08:32:27Z | russia-x-ukraine-ceasefire-agreement-by-december-31-2026 | geopolitics | 0.4300 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 5 veces) |
| 2026-05-31T08:32:29Z | save-act-signed-into-law-in-2026 | other | 0.0800 | E2 | edge 0.030 < mín 0.040 (p̂=0.050, implied=0.080) |
| 2026-05-31T08:32:29Z | will-there-be-exactly-1-confirmed-vei-4-or-higher-volcanic-eruption-worldwide-in-2026-937 | weather-natural | 0.2700 | P3_low_absolute_liquidity | liquidity $1293 < absolute min $5000 |
| 2026-05-31T08:32:29Z | will-any-category-5-hurricane-make-landfall-in-the-us-in-before-2027 | weather-natural | 0.1300 | P3_low_absolute_liquidity | liquidity $1631 < absolute min $5000 |
| 2026-05-31T08:32:31Z | will-mohammad-bagher-ghalibaf-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0400 | E2 | edge 0.020 < mín 0.040 (p̂=0.020, implied=0.040) |
| 2026-05-31T08:32:31Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1374 < absolute min $5000 |
| 2026-05-31T08:32:34Z | will-the-iranian-regime-fall-by-the-end-of-2026 | other | 0.1300 | N1 | noticias contradicen tesis: Régimen iraní sobrevive al asesinato de Jamenei; su hijo Mojtaba designado nuevo líder... |
| 2026-05-31T08:32:34Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:32:34Z | will-yair-golan-be-the-next-prime-minister-of-israel | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 213 d, liq $81638) |
| 2026-05-31T08:32:34Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:32:34Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0590 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-31T08:32:36Z | uae-x-qatar-sever-diplomatic-relations-in-2026 | other | 0.0700 | R1_longshot | R1 longshot: precio 0.0700 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.050 < 0.15) |
| 2026-05-31T08:32:36Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3393 < absolute min $5000 |
| 2026-05-31T08:32:39Z | will-alberta-join-the-us | other | 0.0450 | E2 | edge 0.035 < mín 0.040 (p̂=0.010, implied=0.045) |
| 2026-05-31T08:32:42Z | will-satoshis-identity-be-revealed-by-december-31 | other | 0.0700 | R1_longshot | R1 longshot: precio 0.0700 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.050 < 0.15) |
| 2026-05-31T08:32:42Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 213 d to resolution (>=7 threshold) |
| 2026-05-31T08:32:42Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 213 d, liq $3958) |
| 2026-05-31T08:32:42Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 213 d, liq $9857) |
| 2026-05-31T08:32:42Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 213 d, liq $8338) |
| 2026-05-31T08:32:44Z | will-north-korea-invade-south-korea-before-2027 | geopolitics | 0.0590 | E2 | edge 0.039 < mín 0.040 (p̂=0.020, implied=0.059) |
| 2026-05-31T08:32:44Z | will-larry-page-be-richest-person-on-december-31 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 213 d, liq $6581) |
| 2026-05-31T08:32:44Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 213 d, liq $18930) |
| 2026-05-31T08:32:44Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $254 < absolute min $5000 |
| 2026-05-31T08:32:44Z | will-sadegh-larijani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 213 d, liq $59460) |
| 2026-05-31T08:32:46Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | N2 | edge sin soporte de noticias 0.030<0.040 tras shrink ×0.50 |
| 2026-05-31T08:32:48Z | us-agrees-to-give-ukraine-security-guarantee-by-june-30 | geopolitics | 0.0230 | E2 | edge 0.027 < mín 0.040 (p̂=0.050, implied=0.023) |
| 2026-05-31T08:32:50Z | will-the-us-reopen-its-embassy-in-iran-in-2026 | geopolitics | 0.1200 | R5_precedents | R5 precedentes: 0 < 3 casos análogos |
| 2026-05-31T08:32:52Z | will-russia-capture-all-of-lyman-by-december-31-2026 | geopolitics | 0.2200 | E2 | edge 0.030 < mín 0.040 (p̂=0.250, implied=0.220) |
| 2026-05-31T08:32:52Z | will-al-carns-be-the-next-prime-minister-of-the-united-kingdom-in-2026-126 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 213 d, liq $56581) |
| 2026-05-31T08:32:52Z | will-7-8-spacex-starship-launches-successfully-reach-space-in-2026 | other | 0.0680 | P3_low_absolute_liquidity | liquidity $3722 < absolute min $5000 |
| 2026-05-31T08:32:54Z | us-x-iran-permanent-peace-deal-by-december-31-2026-961-587-341-574-555 | geopolitics | 0.7000 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: Evento sin catalyst verificable en l... |
| 2026-05-31T08:32:54Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0100 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:32:56Z | ukraine-signs-peace-deal-with-russia-before-2027 | geopolitics | 0.3100 | N1 | noticias contradicen tesis: Conversaciones estancadas; solo alto al fuego temporal sin acuerdo de paz formal firmado ... |
| 2026-05-31T08:32:56Z | romanian-pm-bolojan-out-by-december-31-832-595 | executive-action | 0.9860 | P0_ceiling | price ceiling: 0.9860 > 0.980 |
| 2026-05-31T08:32:56Z | will-mahmoud-ahmadinejad-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 213 d, liq $54463) |
| 2026-05-31T08:32:56Z | will-beyonc-be-the-top-spotify-artist-for-2026 | entertainment | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 213 d, liq $16227) |
| 2026-05-31T08:32:56Z | will-alireza-arafi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 213 d, liq $64493) |
| 2026-05-31T08:32:59Z | poilievre-out-as-leader-of-conservatives-before-2027 | other | 0.1300 | E2 | edge 0.020 < mín 0.040 (p̂=0.150, implied=0.130) |
| 2026-05-31T08:32:59Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4079 < absolute min $5000 |
| 2026-05-31T08:32:59Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:32:59Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 213 d, liq $491) |
| 2026-05-31T08:33:01Z | will-syria-join-the-abraham-accords-before-2027 | other | 0.1000 | E2 | edge 0.020 < mín 0.040 (p̂=0.080, implied=0.100) |
| 2026-05-31T08:33:01Z | will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026 | entertainment | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 213 d, liq $18031) |
| 2026-05-31T08:33:01Z | ramp-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $2904 < absolute min $5000 |
| 2026-05-31T08:33:01Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-metamask-launch-a-token-by-december-31-2026 | crypto-launch | 0.3300 | P3_low_absolute_liquidity | liquidity $4021 < absolute min $5000 |
| 2026-05-31T08:33:01Z | will-opensea-launch-a-token-by-june-30-2026 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-31T08:33:01Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2472 < absolute min $5000 |
| 2026-05-31T08:33:01Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | R1_longshot | R1 longshot: precio 0.0500 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.040 < 0.15) |
| 2026-05-31T08:33:01Z | will-stablecoins-hit-500b-before-2027 | other | 0.1200 | P3_low_absolute_liquidity | liquidity $2130 < absolute min $5000 |
| 2026-05-31T08:33:01Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $2406 < absolute min $5000 |
| 2026-05-31T08:33:01Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3900 | P3_low_absolute_liquidity | liquidity $1155 < absolute min $5000 |
| 2026-05-31T08:33:01Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-31T08:33:01Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | crypto-launch | 0.0470 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-31T08:33:01Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 214 d, liq $1449) |
| 2026-05-31T08:33:01Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1335 < absolute min $5000 |
| 2026-05-31T08:33:01Z | will-ethereum-reach-6000-by-december-31-2026 | market | 0.0600 | R1_longshot | R1 longshot: precio 0.0600 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.050 < 0.15) |
| 2026-05-31T08:33:01Z | metamask-fdv-above-3b-one-day-after-launch-363-663-664-569-222 | crypto-launch | 0.0910 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-31T08:33:01Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-31T08:33:01Z | puffpaw-fdv-above-400m-one-day-after-launch | crypto-launch | 0.0880 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-31T08:33:01Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $520 < absolute min $5000 |
| 2026-05-31T08:33:01Z | will-pittsburgh-steelers-win-the-2027-nfl-afc-championship-898 | sports-season | 0.0230 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | sports-season | 0.0570 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-atlanta-falcons-win-the-2027-nfl-nfc-championship-312 | sports-season | 0.0090 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0090 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | sports-season | 0.1100 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | sports-season | 0.0240 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | sports-season | 0.0280 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-minnesota-vikings-win-the-2027-nfl-nfc-championship-884 | sports-season | 0.0310 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-the-new-orleans-saints-win-the-2027-nfl-league-championship | sports-season | 0.0110 | S1 | categoría suspendida por Brier (cat=sports-season) |
| 2026-05-31T08:33:01Z | will-yal-braun-pivet-win-the-2027-french-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 333 d, liq $287785) |
| 2026-05-31T08:33:01Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 578 d, liq $43433) |
| 2026-05-31T08:33:01Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1707435) |
| 2026-05-31T08:33:01Z | will-zohran-mamdani-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1615528) |
| 2026-05-31T08:33:01Z | will-glenn-youngkin-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1087635) |
| 2026-05-31T08:33:01Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1226599) |
| 2026-05-31T08:33:01Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $2270512) |
| 2026-05-31T08:33:01Z | will-eric-trump-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1960381) |
| 2026-05-31T08:33:01Z | will-lebron-james-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $2194165) |
| 2026-05-31T08:33:01Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1190820) |
| 2026-05-31T08:33:01Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 890 d, liq $1104209) |
| 2026-05-31T08:33:01Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $2145718) |
| 2026-05-31T08:33:01Z | will-jamie-dimon-win-the-2028-us-presidential-election | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 890 d, liq $766578) |
| 2026-05-31T08:33:01Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $953759) |
| 2026-05-31T08:33:01Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1685856) |
| 2026-05-31T08:33:01Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 890 d, liq $1119125) |
| 2026-05-31T08:33:01Z | will-steve-bannon-win-the-2028-republican-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 890 d, liq $1271923) |
| 2026-05-31T08:33:01Z | will-tim-walz-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 890 d, liq $1756399) |
| 2026-05-31T08:33:01Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $957) |
| 2026-05-31T08:33:01Z | us-announces-new-iran-agreementceasefire-extension-by-june-7-265-824-655 | geopolitics | 0.3500 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 5 veces) |
| 2026-05-31T08:33:01Z | us-announces-new-iran-agreementceasefire-extension-by-may-31-665-831-238 | geopolitics | 0.0440 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 5 veces) |
| 2026-05-31T08:33:01Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2282 < absolute min $5000 |
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
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T19:45:01Z | T-2325745-1780068909718 | geopolitics | 0.9100 | 0.9500 | 8.12 | target_hit |  | 184.71 | short | 0.17 |
| 2026-05-29T20:35:41Z | T-678929-1780067007527 | elections | 0.2800 | 0.2700 | -8.09 | horizon-shortonly-2026-05-29 |  | 226.48 | medium | 0.23 |
| 2026-05-29T20:35:41Z | T-2354976-1780067007527 | geopolitics | 0.0900 | 0.0800 | -7.88 | horizon-shortonly-2026-05-29 |  | 70.93 | medium | 0.23 |
| 2026-05-29T20:35:41Z | T-564210-1780067007527 | sports-season | 0.5900 | 0.5700 | -2.62 | horizon-shortonly-2026-05-29 |  | 77.42 | medium | 0.23 |
| 2026-05-29T20:35:41Z | T-907355-1780067007527 | elections | 0.0200 | 0.0190 | -3.18 | horizon-shortonly-2026-05-29 |  | 63.59 | medium | 0.23 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T20:35:41Z | T-2183424-1780067007527 | other | 0.0500 | 0.0500 | 0.00 | horizon-shortonly-2026-05-29 |  | 32.30 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-1971905-1780067007527 | geopolitics | 0.3800 | 0.3400 | -21.36 | horizon-shortonly-2026-05-29 |  | 202.95 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-1494702-1780067007527 | market | 0.0300 | 0.0230 | -27.83 | horizon-shortonly-2026-05-29 |  | 119.26 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-1363963-1780067007527 | other | 0.0570 | 0.0560 | -0.75 | horizon-shortonly-2026-05-29 |  | 42.60 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-597967-1780068909718 | executive-action | 0.1400 | 0.1300 | -4.37 | horizon-shortonly-2026-05-29 |  | 61.21 | long | 0.21 |
| 2026-05-29T20:35:41Z | T-569373-1780085069997 | elections | 0.0370 | 0.0350 | -2.05 | horizon-shortonly-2026-05-29 |  | 37.84 | medium | 0.02 |
| 2026-05-29T20:35:41Z | T-677314-1780085069997 | geopolitics | 0.0300 | 0.0230 | -8.06 | horizon-shortonly-2026-05-29 |  | 34.56 | long | 0.02 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T20:40:01Z | T-2340268-1780087006486 | market | 0.2900 | 0.3100 | 3.06 | no_remaining_edge |  | 44.37 | short | 0.00 |
| 2026-05-30T03:30:03Z | T-1999689-1780087006486 | other | 0.6390 | 0.0120 | -169.63 | market_closed |  | 172.88 | short | 0.29 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-30T04:20:01Z | T-2365528-1780111994961 | market | 0.2700 | 0.3900 | 41.63 | target_hit |  | 93.68 | short | 0.03 |
| 2026-05-30T06:10:02Z | T-2139900-1780085069997 | other | 0.4400 | 0.0010 | -94.51 | market_closed |  | 94.73 | short | 0.42 |
| 2026-05-30T06:10:04Z | T-2364498-1780092236400 | geopolitics | 0.0290 | 0.0010 | -28.46 | market_closed |  | 29.47 | short | 0.34 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-30T14:30:01Z | T-2340268-1780088637838 | market | 0.3300 | 0.4400 | 58.47 | no_remaining_edge |  | 175.41 | short | 0.73 |
| 2026-05-30T16:20:01Z | T-2324493-1780131807761 | other | 0.5200 | 0.7400 | 57.81 | target_hit |  | 136.65 | short | 0.30 |
| 2026-05-30T16:25:01Z | T-2316218-1780067007527 | other | 0.3100 | 0.4900 | 39.53 | target_hit |  | 68.08 | short | 1.06 |
| 2026-05-30T16:54:14Z | T-2111561-1780070680630 | other | 0.0250 | 0.0090 | -31.12 | thesis_stale |  | 48.63 | short | 1.03 |
| 2026-05-30T16:54:16Z | T-2244268-1780070680630 | other | 0.0580 | 0.0280 | -28.93 | thesis_stale |  | 55.94 | short | 1.03 |
| 2026-05-30T16:54:18Z | T-2111563-1780087006486 | geopolitics | 0.0200 | 0.0070 | -41.35 | thesis_stale |  | 63.61 | short | 0.85 |
| 2026-05-30T16:54:19Z | T-2002531-1780090425061 | geopolitics | 0.0290 | 0.0110 | -30.50 | thesis_stale |  | 49.14 | short | 0.81 |
| 2026-05-30T16:54:21Z | T-2111562-1780092236400 | executive-action | 0.1100 | 0.0500 | -78.58 | thesis_stale |  | 144.07 | short | 0.78 |
| 2026-05-30T16:54:23Z | T-2111605-1780108388595 | geopolitics | 0.0490 | 0.0370 | -16.09 | thesis_stale |  | 65.69 | short | 0.60 |
| 2026-05-30T16:54:25Z | T-569344-1780115594609 | elections | 0.2460 | 0.2410 | -0.79 | thesis_stale |  | 39.02 | short | 0.51 |
| 2026-05-30T16:54:27Z | T-2159726-1780122799218 | geopolitics | 0.0240 | 0.0200 | -18.23 | thesis_stale |  | 109.38 | short | 0.43 |
| 2026-05-30T16:54:28Z | T-2119071-1780151620830 | geopolitics | 0.0240 | 0.0140 | -16.45 | thesis_stale |  | 39.47 | short | 0.10 |
| 2026-05-30T16:54:30Z | T-566136-1780151620830 | sports-season | 0.5500 | 0.3200 | -69.22 | thesis_stale |  | 165.53 | short | 0.10 |
| 2026-05-30T18:30:03Z | T-2354184-1780135407266 | other | 0.1200 | 0.0000 | -26.79 | market_closed |  | 26.79 | short | 0.35 |
| 2026-05-30T20:40:03Z | T-1999661-1780108388595 | other | 0.0590 | 0.0010 | -36.98 | thesis_stale |  | 37.62 | short | 0.75 |
| 2026-05-30T23:45:02Z | T-2316216-1780067007527 | other | 0.4200 | 0.0000 | -99.77 | market_closed |  | 99.77 | short | 1.36 |
| 2026-05-31T04:00:03Z | T-2132800-1780085069997 | market | 0.0210 | 0.0030 | -44.45 | thesis_stale |  | 51.85 | short | 1.33 |
| 2026-05-31T04:00:04Z | T-2132779-1780087006486 | market | 0.0520 | 0.0050 | -44.84 | thesis_stale |  | 49.61 | short | 1.31 |
| 2026-05-31T04:00:06Z | T-2132832-1780115594609 | other | 0.0210 | 0.0010 | -32.46 | thesis_stale |  | 34.08 | short | 0.98 |
| 2026-05-31T04:00:07Z | T-2365528-1780173320643 | market | 0.2100 | 0.0900 | -93.71 | thesis_stale | fortune.com,finance.yahoo.com,coindesk.com | 163.99 | short | 0.31 |
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-31T04:20:01Z | T-2361313-1780167889157 | entertainment | 0.8800 | 0.9200 | 8.86 | target_hit | comingsoon.net,the-numbers.com,popculture.com | 194.82 | short | 0.39 |
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


<!-- auto-generated 2026-05-30T17:04:18Z -->
- En categoría `geopolitics` · horizonte `short` · precio `<0.10`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-31T00:11:10Z -->
- En categoría `other` · horizonte `short` · precio `0.30-0.70`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-31T04:05:34Z -->
- En categoría `market` · horizonte `short` · precio `0.10-0.30`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**


## Human notes
_(no se toca por automatización)_

## Sources stats (rolling 30d, computed by brain)

| domain | trades | wins | losses | win_rate | total_pnl | blacklisted |
|---|---|---|---|---|---|---|
| claudemax-websearch | 11 | 2 | 8 | 20.0% | -77.68 | YES |
| fortune.com | 1 | 0 | 1 | 0.0% | -93.71 | no |
| finance.yahoo.com | 1 | 0 | 1 | 0.0% | -93.71 | no |
| coindesk.com | 1 | 0 | 1 | 0.0% | -93.71 | no |
| comingsoon.net | 1 | 1 | 0 | 100.0% | +8.86 | no |
| the-numbers.com | 1 | 1 | 0 | 100.0% | +8.86 | no |
| popculture.com | 1 | 1 | 0 | 100.0% | +8.86 | no |

## Anti-patterns identificados- same-day-price-target · intraday-decay · no-momentum-confirm · short-horizon-chalk · high-entry-low-exit — visto en  (2026-05-26, pnl $-57.55)

- daily-news-position · short-term-fading · price-action-thesis — visto en will-bitcoin-dip-to-72500-in-may-2026-from-may-27 (2026-05-31, pnl $-93.71)
- no-catalyst-before-bet · short-horizon-hype · price-target-no-edge · unverified-thesis · ceasefire-rumor — visto en will-xrp-reach-1pt6-in-may-2026 (2026-05-31, pnl $-32.46)
- thesis-stale · short-term-chase · no-catalyst · oversold-trap — visto en will-bitcoin-dip-to-70k-in-may-2026-438-356-919 (2026-05-31, pnl $-44.84)
- thesis-stale · short-horizon-weakness · missing-sources · crypto-momentum-bet — visto en will-ethereum-dip-to-1800-in-may-2026 (2026-05-31, pnl $-44.45)
- market-cap-ranking · short-horizon-mega-cap · no-edge-no-catalyst · same-day-chalk-bet · stale-thesis-fast-exit — visto en will-alphabet-be-the-third-largest-company-in-the-world-by-market-cap-on-may-31 (2026-05-30, pnl $-36.98)
- short-horizon-closed · no-exit-liquidity · event-catalyst-blind · low-volume-trap — visto en elon-musk-of-tweets-may-28-may-30-65-89 (2026-05-30, pnl $-26.79)
- short-horizon-sports · same-day-chalk-bet · season-stale-thesis — visto en will-psg-win-the-202526-champions-league (2026-05-30, pnl $-69.22)
- intraday-geopolitics · no-recent-source-check · short-thesis-stale · low-sample-data-bet — visto en will-40-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 (2026-05-30, pnl $-16.45)
- ceasefire-rumor · geopolitics-short-horizon · low-liquidity-microcap · thesis-stale — visto en israel-closes-its-airspace-by-may-31 (2026-05-30, pnl $-18.23)
- early-election-bet · no-recent-polls · short-horizon-mispricing — visto en will-abelardo-de-la-espriella-win-the-1st-round-of-the-2026-colombian-presidential-election (2026-05-30, pnl $-0.79)
- short-horizon-unverified · geopolitical-rumor · stale-thesis-entry · no-source-trade — visto en iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 (2026-05-30, pnl $-16.09)
- geopolitics-short-horizon · ceasefire-rumor · executive-action-vacuum · thesis-stale-ignored · low-liquidity-chalk — visto en will-trump-agree-to-iranian-oil-sanction-relief-by-may-31 (2026-05-30, pnl $-78.58)
- ceasefire-rumor · short-horizon-geopolitics · permanent-peace-due-date · high-noise-low-liquidity · thesis-stale-rapidly — visto en israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 (2026-05-30, pnl $-30.50)
- unconfirmed-rumor · arbitrary-deadline · geopolitics-short-horizon · low-liquidity-trigger — visto en will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 (2026-05-30, pnl $-41.35)
- no-catalyst-entry · stale-thesis-hold · short-horizon-unknown — visto en will-trump-restart-project-freedom-by-may-31 (2026-05-30, pnl $-28.93)
- geopolitics-short-horizon · no-catalyst-entry · stale-thesis-trap — visto en will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 (2026-05-30, pnl $-31.12)
- ceasefire-rumor · short-horizon-geopolitics · market-closed-capitulation · low-liquidity-pump — visto en us-announces-new-iran-agreementceasefire-extension-by-may-29 (2026-05-30, pnl $-28.46)
- short-horizon-binary · no-sources-checked · market-closed-bet · low-conviction-chalk — visto en will-apple-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 (2026-05-30, pnl $-169.63)
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

## Brier score (semanal) — 2026-05-29

resolved=2 overall_brier=0.5000

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 2 | 0.5000 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 1 | 0.0000 |
| long | 1 | 1.0000 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| unknown | 2 | 0.5000 |

## Calibracion diaria — 2026-05-29

- Ventana: 7d · short thesis-closes n=26 (W26/L0)
- Win rate 100.0% · avg win $12.78 · avg loss $0.0 · payoff 0.0 (break-even wr 0.0%)
- Expectancy/trade $12.7792 · total $332.26
- Brier short: n=4 score=0.0925
- **kelly_shrink=1.0** · min_edge_points_override=ninguno
- short n=26 wr=1.000 expectancy=+12.7792/trade; shrink=1.0 (no losers in window, shrink_exp=1.0); min_edge 0.030->0.030

## Brier score (semanal) — 2026-05-30

resolved=12 overall_brier=0.2034

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 11 | 0.1944 |
| market | 1 | 0.3025 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 1 | 0.0000 |
| short | 10 | 0.1441 |
| long | 1 | 1.0000 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| info | 2 | 0.0137 |
| other | 1 | 0.1225 |
| none | 4 | 0.1671 |
| calibration | 3 | 0.2076 |
| unknown | 2 | 0.5000 |


## Calibracion diaria — 2026-05-30

- Ventana: 7d · short thesis-closes n=25 (W25/L0)
- Win rate 100.0% · avg win $12.51 · avg loss $0.0 · payoff 0.0 (break-even wr 0.0%)
- Expectancy/trade $12.5072 · total $312.68
- Brier short: n=33 score=0.1743
- **kelly_shrink=1.0** · min_edge_points_override=ninguno
- short n=25 wr=1.000 expectancy=+12.5072/trade; shrink=1.0 (no losers in window, shrink_exp=1.0 * brier_factor=1.000 (brier=0.174 n=33)); min_edge 0.030->0.030

## Brier score (semanal) — 2026-05-30

resolved=62 overall_brier=0.1608

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 38 | 0.1410 |
| market | 2 | 0.1531 |
| geopolitics | 22 | 0.1957 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 6 | 0.0008 |
| short | 55 | 0.1630 |
| long | 1 | 1.0000 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| calibration | 14 | 0.1028 |
| info | 20 | 0.1068 |
| other | 1 | 0.1225 |
| none | 22 | 0.1942 |
| unknown | 5 | 0.4000 |


## Brier score (semanal) — 2026-05-30

resolved=89 overall_brier=0.1199

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 58 | 0.1023 |
| geopolitics | 29 | 0.1530 |
| market | 2 | 0.1531 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 22 | 0.0029 |
| short | 65 | 0.1478 |
| long | 2 | 0.5002 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| calibration | 21 | 0.0751 |
| info | 28 | 0.0779 |
| none | 28 | 0.1528 |
| other | 4 | 0.1594 |
| unknown | 8 | 0.2500 |

## Brier score (semanal) — 2026-05-30

resolved=135 overall_brier=0.0816

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 104 | 0.0603 |
| geopolitics | 29 | 0.1530 |
| market | 2 | 0.1531 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 64 | 0.0016 |
| short | 69 | 0.1436 |
| long | 2 | 0.5002 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| none | 68 | 0.0667 |
| calibration | 24 | 0.0689 |
| info | 28 | 0.0779 |
| other | 5 | 0.1275 |
| unknown | 10 | 0.2000 |

### P&L realizado por categoría

| categoría | n | pnl | win% | avg/trade |
|---|---|---|---|---|
| other | 233 | -585.62 | 53% | -2.51 |
| market | 57 | -537.10 | 49% | -9.42 |
| geopolitics | 82 | -233.24 | 48% | -2.84 |
| elections | 17 | -148.92 | 18% | -8.76 |
| executive-action | 3 | -94.50 | 0% | -31.50 |
| sports-season | 4 | -45.26 | 25% | -11.32 |
| crypto-launch | 2 | -21.01 | 50% | -10.50 |

## Brier score (semanal) — 2026-05-31

resolved=511 overall_brier=0.1561

### Por categoría

| group | n | brier |
|-------|---|-------|
| market | 56 | 0.0932 |
| other | 309 | 0.1074 |
| geopolitics | 29 | 0.1530 |
| sports-season | 117 | 0.3157 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 104 | 0.0010 |
| short | 405 | 0.1943 |
| long | 2 | 0.5002 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| none | 170 | 0.1253 |
| info | 130 | 0.1279 |
| other | 18 | 0.1318 |
| unknown | 15 | 0.2000 |
| calibration | 177 | 0.2049 |
| liquidity | 1 | 0.2209 |

### P&L realizado por categoría

| categoría | n | pnl | win% | avg/trade |
|---|---|---|---|---|
| other | 236 | -749.16 | 52% | -3.17 |
| market | 57 | -537.10 | 49% | -9.42 |
| geopolitics | 82 | -233.24 | 48% | -2.84 |
| elections | 17 | -148.92 | 18% | -8.76 |
| executive-action | 3 | -94.50 | 0% | -31.50 |
| sports-season | 4 | -45.26 | 25% | -11.32 |
| crypto-launch | 2 | -21.01 | 50% | -10.50 |

## Calibracion diaria — 2026-05-31

- Ventana: 7d · short thesis-closes n=40 (W24/L16)
- Win rate 60.0% · avg win $17.29 · avg loss $-49.21 · payoff 0.3514 (break-even wr 74.0%)
- Expectancy/trade $-9.3088 · total $-372.35
- Brier short: n=406 score=0.1944
- **kelly_shrink=0.8108** · min_edge_points_override=0.04
- short n=40 wr=0.600 expectancy=-9.3088/trade; shrink=0.8108 (wr/breakeven=0.600/0.740=0.811 * brier_factor=1.000 (brier=0.194 n=406)); min_edge 0.030->0.040

## Brier score (semanal) — 2026-05-31

resolved=525 overall_brier=0.1542

### Por categoría

| group | n | brier |
|-------|---|-------|
| market | 63 | 0.0828 |
| other | 316 | 0.1087 |
| geopolitics | 29 | 0.1530 |
| sports-season | 117 | 0.3157 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 104 | 0.0010 |
| short | 419 | 0.1905 |
| long | 2 | 0.5002 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| model-barrier | 7 | 0.0001 |
| none | 174 | 0.1271 |
| info | 130 | 0.1279 |
| other | 19 | 0.1355 |
| unknown | 15 | 0.2000 |
| calibration | 179 | 0.2033 |
| liquidity | 1 | 0.2209 |

### P&L realizado por categoría

| categoría | n | pnl | win% | avg/trade |
|---|---|---|---|---|
| other | 237 | -781.62 | 52% | -3.30 |
| market | 60 | -720.10 | 47% | -12.00 |
| geopolitics | 82 | -233.24 | 48% | -2.84 |
| elections | 17 | -148.92 | 18% | -8.76 |
| executive-action | 3 | -94.50 | 0% | -31.50 |
| sports-season | 4 | -45.26 | 25% | -11.32 |
| crypto-launch | 2 | -21.01 | 50% | -10.50 |
| entertainment | 1 | +8.86 | 100% | +8.86 |
