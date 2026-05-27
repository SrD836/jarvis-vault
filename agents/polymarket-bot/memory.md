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
| 2026-05-27T11:01:27Z | us-announces-new-iran-agreementceasefire-extension-by-may-26 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -1 d) |
| 2026-05-27T11:01:27Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon -1 d) |
| 2026-05-27T11:01:27Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon -1 d) |
| 2026-05-27T11:01:27Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $1517 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon -147 d) |
| 2026-05-27T11:30:04Z | ukraine-election-held-by-june-30-2026-465-757 | elections | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon -146 d) |
| 2026-05-27T11:30:04Z | new-stranger-things-episode-released-by-may-31 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -140 d) |
| 2026-05-27T11:30:04Z | will-grigor-dimitrov-win-the-2026-australian-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -115 d) |
| 2026-05-27T11:30:04Z | will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -108 d) |
| 2026-05-27T11:30:04Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -93 d) |
| 2026-05-27T11:30:04Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -93 d) |
| 2026-05-27T11:30:04Z | internet-access-restored-in-iran-by-may-31-2026 | geopolitics | 0.5780 | P2 | mercado ya expiró (endDate=2026-03-14T00:00:00Z, hace 74 días) |
| 2026-05-27T11:30:04Z | internet-access-restored-in-iran-by-june-30-2026 | geopolitics | 0.7900 | P2 | mercado ya expiró (endDate=2026-03-14T00:00:00Z, hace 74 días) |
| 2026-05-27T11:30:04Z | will-matej-tonin-be-the-next-prime-minister-of-slovenia | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon -66 d) |
| 2026-05-27T11:30:04Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon -66 d) |
| 2026-05-27T11:30:04Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -66 d) |
| 2026-05-27T11:30:04Z | will-lars-boje-mathiesen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon -64 d) |
| 2026-05-27T11:30:04Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -64 d) |
| 2026-05-27T11:30:04Z | tim-walz-charged-by-december-31-2026 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4246 < absolute min $5000 |
| 2026-05-27T11:30:04Z | weed-rescheduled-by-december-31 | other | 0.3030 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 57 días) |
| 2026-05-27T11:30:04Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -45 d) |
| 2026-05-27T11:30:04Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -44 d) |
| 2026-05-27T11:30:04Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.1900 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 12 días) |
| 2026-05-27T11:30:04Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon -10 d) |
| 2026-05-27T11:30:04Z | fl1-nan-tou-2026-05-17-nan | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -9 d) |
| 2026-05-27T11:30:04Z | will-0-world-records-be-broken-at-the-2026-enhanced-games | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-19-may-26-220-239 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-27 | geopolitics | 0.0800 | P9 | P9: geopolitics pump cluster (price 0.08, 0d) |
| 2026-05-27T11:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | highest-temperature-in-hong-kong-on-may-27-2026-34corhigher | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | bitcoin-above-76k-on-may-27-2026 | market | 0.5100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | bitcoin-above-82k-on-may-27-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-25-may-27-40-64 | other | 0.9690 | P0_ceiling | price ceiling: 0.9690 > 0.950 |
| 2026-05-27T11:30:04Z | bitcoin-above-80k-on-may-27-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-25-may-27-90-114 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-25-may-27-65-89 | other | 0.0400 | P0_floor | price floor: 0.0400 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-25-may-27-0-39 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | bitcoin-above-74k-on-may-27-2026 | market | 0.9900 | P0_ceiling | price ceiling: 0.9900 > 0.950 |
| 2026-05-27T11:30:04Z | bitcoin-above-78k-on-may-27-2026 | market | 0.0100 | P0_floor | price floor: 0.0100 < 0.050 (horizon 0 d) |
| 2026-05-27T11:30:04Z | col-cry-ray-2026-05-27-cry | other | 0.4700 | M2 | M2 soft-learned: uncategorized·?·0.30-0.70 = 6L/0W (wr 0%) |
| 2026-05-27T11:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-28 | geopolitics | 0.1700 | P9 | P9: geopolitics pump cluster (price 0.17, 0d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-260-279 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-380-399 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-360-379 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-120-139 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-340-359 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-400-419 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-220-239 | other | 0.1300 | M1 | memoria: slug prefix match; same category; same price bucket low (score 0.90) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-180-199 | other | 0.4190 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-280-299 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-140-159 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-320-339 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-160-179 | other | 0.1540 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-300-319 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-240-259 | other | 0.0340 | P0_floor | price floor: 0.0340 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-22-may-29-480-499 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 2 d) |
| 2026-05-27T11:30:04Z | ucl-psg-ars-2026-05-30-psg | other | 0.4200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0690 | P9 | P9: geopolitics pump cluster (price 0.07, 3d) |
| 2026-05-27T11:30:04Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.2110 | P9 | P9: geopolitics pump cluster (price 0.21, 3d) |
| 2026-05-27T11:30:04Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0360 | P0_floor | price floor: 0.0360 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.3000 | P9 | P9: geopolitics pump cluster (price 0.30, 3d) |
| 2026-05-27T11:30:04Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-trump-restart-project-freedom-by-may-31 | other | 0.0890 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0450 | P0_floor | price floor: 0.0450 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.5000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-japan-send-warships-through-the-strait-of-hormuz-by-may-31-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-gujarat-titans-win-the-2026-indian-premier-league | sports-season | 0.2310 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.2000 | P9 | P9: geopolitics pump cluster (price 0.20, 3d) |
| 2026-05-27T11:30:04Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0630 | P9 | P9: geopolitics pump cluster (price 0.06, 3d) |
| 2026-05-27T11:30:04Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0800 | P9 | P9: geopolitics pump cluster (price 0.08, 3d) |
| 2026-05-27T11:30:04Z | trump-out-as-president-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0160 | P0_floor | price floor: 0.0160 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-psg-win-the-202526-champions-league | sports-season | 0.5800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-deepseek-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0880 | P9 | P9: geopolitics pump cluster (price 0.09, 3d) |
| 2026-05-27T11:30:04Z | will-nvidia-be-the-largest-company-in-the-world-by-market-cap-on-may-31-971 | market | 0.9970 | P0_ceiling | price ceiling: 0.9970 > 0.950 |
| 2026-05-27T11:30:04Z | will-apple-be-the-largest-company-in-the-world-by-market-cap-on-may-31-332 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-may-31-111 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 3 d) |
| 2026-05-27T11:30:04Z | gta-vi-released-before-june-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-ivn-cepeda-castro-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.6700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-abelardo-de-la-espriella-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.3480 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-paloma-valencia-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.4100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-wti-crude-oil-wti-hit-high-200-in-may | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-wti-reach-120-in-may-2026-592-217 | market | 0.0110 | P0_floor | price floor: 0.0110 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-wti-reach-130-in-may-2026-733 | market | 0.0070 | P0_floor | price floor: 0.0070 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-wti-reach-140-in-may-2026-916 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-wti-reach-110-in-may-2026-116-472 | market | 0.0310 | P0_floor | price floor: 0.0310 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-wti-crude-oil-wti-hit-high-105-in-may | market | 0.0600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0130 | P0_floor | price floor: 0.0130 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-bitcoin-reach-100k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-bitcoin-reach-95k-in-may-2026 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-2026-1520-1559 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-2026-1400-1439 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-2026-2000plus | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-ethereum-reach-2600-in-may-2026 | market | 0.0040 | P0_floor | price floor: 0.0040 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-bitcoin-reach-150k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | over-40m-committed-to-the-printr-public-sale | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0530 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-bitcoin-dip-to-65k-in-may-2026-183-857-425 | market | 0.0080 | P0_floor | price floor: 0.0080 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-2026-1320-1359 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-ethereum-dip-to-1400-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-solana-dip-to-40-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0220 | P0_floor | price floor: 0.0220 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | over-15m-committed-to-the-printr-public-sale | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-solana-reach-100-in-may-2026 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.050 (horizon 4 d) |
| 2026-05-27T11:30:04Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.050 (horizon 5 d) |
| 2026-05-27T11:30:04Z | will-ian-calderon-get-the-first-or-second-most-votes-in-the-2026-california-governor-primary-election | elections | 0.0130 | P0_floor | price floor: 0.0130 < 0.050 (horizon 5 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-26-june-2-500plus | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | elon-musk-of-tweets-may-26-june-2-60-79 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-the-rebuilding-korea-party-rkp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-chong-won-oh-win-the-2026-seoul-mayoral-election | elections | 0.7800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-oh-se-hoon-win-the-2026-seoul-mayoral-election | elections | 0.2400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-yoon-sang-hyun-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-park-hong-keun-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-park-yong-jin-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.050 (horizon 6 d) |
| 2026-05-27T11:30:04Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 9 d) |
| 2026-05-27T11:30:04Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | will-rafael-lpez-aliaga-win-the-2026-peruvian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.3300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-roberto-snchez-palomino-win-the-2026-peruvian-presidential-election | elections | 0.2870 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 10 d) |
| 2026-05-27T11:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-june-15 | geopolitics | 0.1900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | iran-closes-its-airspace-by-june-15 | geopolitics | 0.1810 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.4100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-al-mina-be-the-republican-nominee-for-senate-in-virginia | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 19 d) |
| 2026-05-27T11:30:04Z | will-the-oklahoma-city-thunder-win-the-nba-western-conference-finals | sports-season | 0.7900 | M1 | memoria: slug prefix match; same category; same price bucket mid (score 0.90) |
| 2026-05-27T11:30:04Z | will-the-san-antonio-spurs-win-the-nba-western-conference-finals | sports-season | 0.2170 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9740 | P0_ceiling | price ceiling: 0.9740 > 0.950 |
| 2026-05-27T11:30:04Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 20 d) |
| 2026-05-27T11:30:04Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 20 d) |
| 2026-05-27T11:30:04Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 20 d) |
| 2026-05-27T11:30:04Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 20 d) |
| 2026-05-27T11:30:04Z | will-juan-daniel-oviedo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-carlos-felipe-crdoba-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-ivan-cepeda-castro-win-the-2026-colombian-presidential-election | elections | 0.3100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 25 d) |
| 2026-05-27T11:30:04Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 30 d) |
| 2026-05-27T11:30:04Z | will-tynan-lawrence-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 30 d) |
| 2026-05-27T11:30:04Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | iran-closes-its-airspace-by-june-30-432-786-462-866 | geopolitics | 0.2370 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-discords-market-cap-be-between-25b-and-30b-at-market-close-on-ipo-day | crypto-launch | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-june | geopolitics | 0.4200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | trump-out-as-president-by-june-30 | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-mohammed-bin-salman-cease-to-be-the-de-facto-leader-of-saudi-arabia-by-june-30-2026 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-russia-enter-kherson-by-june-30-645 | geopolitics | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-michael-jackson-be-confirmed-to-have-visited-epsteins-island | other | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-base-launch-a-token-by-june-30-2026 | crypto-launch | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | israel-x-iran-permanent-peace-deal-by-june-30-2026-262 | geopolitics | 0.1890 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | timothy-chalamet-confirmed-to-be-esdeekid-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | us-iran-nuclear-deal-by-june-30 | geopolitics | 0.4900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-june-30-2026 | geopolitics | 0.2000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | anthropic-ceo-arrested | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-trump-pardon-tiger-woods-by-june-30 | executive-action | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0440 | P0_floor | price floor: 0.0440 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-june-30-612 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-the-colorado-avalanche-win-the-western-conference-993 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-anthropics-market-cap-be-between-400b-and-600b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | will-france-uk-or-germany-strike-iran-by-june-30-259 | geopolitics | 0.0360 | P0_floor | price floor: 0.0360 < 0.100 (horizon 33 d) |
| 2026-05-27T11:30:04Z | starmer-out-by-june-30-2026-862-594-548-219-739 | executive-action | 0.2000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-silver-si-hit-high-130-by-end-of-june | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | will-silver-si-hit-low-55-by-end-of-june-185-484 | other | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | will-crude-oil-cl-hit-high-200-by-end-of-june-677 | market | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | kharg-island-no-longer-under-iranian-control-by-june-30-561-854 | other | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | will-the-new-york-knicks-win-the-2026-nba-finals | sports-season | 0.2850 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.1440 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-the-oklahoma-city-thunder-win-the-2026-nba-finals | sports-season | 0.5900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-bitcoin-hit-150k-by-june-30-2026 | market | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | microstrategy-sells-any-bitcoin-by-june-30-2026 | market | 0.3670 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | microstrategy-sells-any-bitcoin-by-may-31-2026 | market | 0.0560 | P0_floor | price floor: 0.0560 < 0.100 (horizon 34 d) |
| 2026-05-27T11:30:04Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-victoria-mboko-be-the-2026-womens-wimbledon-winner | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-andrey-rublev-be-the-2026-mens-wimbledon-winner | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0440 | P0_floor | price floor: 0.0440 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-taylor-fritz-be-the-2026-mens-wimbledon-winner | other | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-lorenzo-musetti-be-the-2026-mens-wimbledon-winner | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-alejandro-davidovich-fokina-be-the-2026-mens-wimbledon-winner | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-arthur-fils-be-the-2026-mens-wimbledon-winner | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-madison-keys-be-the-2026-womens-wimbledon-winner | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-clara-tauson-be-the-2026-womens-wimbledon-winner | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-ons-jabeur-be-the-2026-womens-wimbledon-winner | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 45 d) |
| 2026-05-27T11:30:04Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-portugal-win-the-2026-fifa-world-cup-912 | sports-season | 0.1070 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-qatar-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-spain-win-the-2026-fifa-world-cup-963 | sports-season | 0.1740 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-france-win-the-2026-fifa-world-cup-924 | sports-season | 0.1720 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0370 | P0_floor | price floor: 0.0370 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0830 | P0_floor | price floor: 0.0830 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | P0_floor | price floor: 0.0520 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-england-win-the-2026-fifa-world-cup-937 | sports-season | 0.1130 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 53 d) |
| 2026-05-27T11:30:04Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 64 d) |
| 2026-05-27T11:30:04Z | strait-of-hormuz-traffic-returns-to-normal-by-july-31 | geopolitics | 0.6200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | us-x-iran-permanent-peace-deal-by-july-31-2026-831-252 | geopolitics | 0.6600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-paul-reevs-be-the-republican-nominee-for-az-01 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 68 d) |
| 2026-05-27T11:30:04Z | will-kristen-mcdonald-rivet-win-the-2026-michigan-democratic-primary | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 68 d) |
| 2026-05-27T11:30:04Z | will-scott-jensen-win-the-2026-minnesota-governor-republican-primary-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 75 d) |
| 2026-05-27T11:30:04Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 108 d) |
| 2026-05-27T11:30:04Z | will-ebba-busch-be-the-next-prime-minister-of-sweden | elections | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 108 d) |
| 2026-05-27T11:30:04Z | will-tereza-valentova-win-the-2026-womens-us-open | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 108 d) |
| 2026-05-27T11:30:04Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 108 d) |
| 2026-05-27T11:30:04Z | will-andrey-rublev-win-the-2026-mens-us-open | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 108 d) |
| 2026-05-27T11:30:04Z | will-the-moderate-party-m-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0360 | P0_floor | price floor: 0.0360 < 0.100 (horizon 108 d) |
| 2026-05-27T11:30:04Z | will-china-invade-taiwan-by-september-30-2026 | geopolitics | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 125 d) |
| 2026-05-27T11:30:04Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5900 | P8 | P8: election 125d out, price 0.59 en banda ruido [0.30, 0.70] |
| 2026-05-27T11:30:04Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 126 d) |
| 2026-05-27T11:30:04Z | will-romeu-zema-win-the-2026-brazilian-presidential-election | elections | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-partido-liberal-pl-win-the-most-seats-in-the-next-brazilian-senate-election | elections | 0.8300 | P3_low_absolute_liquidity | liquidity $1790 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-tereza-cristina-win-the-2026-brazilian-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-ronaldo-caiado-win-the-2026-brazilian-presidential-election | elections | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-camilo-santana-win-the-2026-brazilian-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-renan-santos-win-the-2026-brazilian-presidential-election | elections | 0.1460 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-eduardo-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-jair-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 129 d) |
| 2026-05-27T11:30:04Z | will-parti-vert-du-qubec-win-the-most-seats-in-the-2026-quebec-general-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 130 d) |
| 2026-05-27T11:30:04Z | will-parti-conservateur-du-qubec-win-the-most-seats-in-the-2026-quebec-general-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 130 d) |
| 2026-05-27T11:30:04Z | will-milwaukee-brewers-win-the-2026-national-league-championship-series | sports-season | 0.0810 | P0_floor | price floor: 0.0810 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-colorado-rockies-win-the-2026-national-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-st-louis-cardinals-win-the-2026-national-league-championship-series | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-cincinnati-reds-win-the-2026-national-league-championship-series | sports-season | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-san-diego-padres-win-the-2026-national-league-championship-series | sports-season | 0.0650 | P0_floor | price floor: 0.0650 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-cleveland-guardians-win-the-2026-american-league-championship-series | sports-season | 0.0860 | P0_floor | price floor: 0.0860 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-san-francisco-giants-win-the-2026-national-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-toronto-blue-jays-win-the-2026-american-league-championship-series | sports-season | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 157 d) |
| 2026-05-27T11:30:04Z | will-elaine-culotti-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 159 d) |
| 2026-05-27T11:30:04Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 159 d) |
| 2026-05-27T11:30:04Z | will-michael-younger-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 159 d) |
| 2026-05-27T11:30:04Z | will-nancy-dahlstrom-win-the-2026-alaska-governor-election | elections | 0.0560 | P0_floor | price floor: 0.0560 < 0.100 (horizon 159 d) |
| 2026-05-27T11:30:04Z | will-antonio-villaraigosa-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 159 d) |
| 2026-05-27T11:30:04Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 159 d) |
| 2026-05-27T11:30:04Z | will-lew-evans-win-the-bachelorette-season-22 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 186 d) |
| 2026-05-27T11:30:04Z | will-richard-van-de-water-win-the-bachelorette-season-22 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 186 d) |
| 2026-05-27T11:30:04Z | will-oliver-bearman-be-the-2026-f1-drivers-champion | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 192 d) |
| 2026-05-27T11:30:04Z | will-lance-stroll-be-the-2026-f1-drivers-champion | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 192 d) |
| 2026-05-27T11:30:04Z | will-michael-harris-ii-win-the-2026-nl-comeback-player-of-the-year-award | other | 0.8800 | P3_low_absolute_liquidity | liquidity $26 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-leo-de-vries-win-the-2026-al-rookie-of-the-year-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 205 d) |
| 2026-05-27T11:30:04Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | us-x-iran-permanent-peace-deal-by-december-31-2026-961-587-341-574-555 | geopolitics | 0.8100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-reza-pahlavi-lead-iran-in-2026 | geopolitics | 0.0690 | P0_floor | price floor: 0.0690 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-macky-sall-be-the-next-secretary-general-of-the-united-nations | other | 0.1180 | P3_low_absolute_liquidity | liquidity $3747 < absolute min $5000 |
| 2026-05-27T11:30:04Z | poilievre-out-as-leader-of-conservatives-before-2027 | other | 0.1600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | spacex-starship-fully-reusable-before-2027 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $1891 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-snapchat-be-acquired-before-2027-286-465 | other | 0.2720 | P3_low_absolute_liquidity | liquidity $2542 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-israel-strike-9-countries-in-2026 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-trump-sell-10k-25k-gold-cards-in-2026 | other | 0.0470 | P0_floor | price floor: 0.0470 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | nato-x-russia-military-clash-by-december-31-2026-244 | geopolitics | 0.2700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-lovable-be-acquired-before-2027-423-881 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4220 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-magomed-ankalaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | brex-ipo-before-2027 | crypto-launch | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | kadyrov-out-as-head-of-the-chechen-republic-by-june-30-2026 | other | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | prince-andrew-sentenced-to-prison | other | 0.1000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-bytedance-have-the-highest-ipo-market-cap-2026-727 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3765 < absolute min $5000 |
| 2026-05-27T11:30:04Z | ebola-pandemic-in-2026 | weather-natural | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-the-us-strike-15-or-more-countries-in-2026 | other | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-there-be-between-17-and-19-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $2174 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-anthropic-be-acquired-before-2027-142-515 | other | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-qatar-uae | other | 0.0530 | P0_floor | price floor: 0.0530 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-bridget-phillipson-be-the-next-prime-minister-of-the-united-kingdom-in-2026-539 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-the-feds-lower-bound-reach-0pt25-or-lower-before-2027-173-921-916-764-754-593-935 | other | 0.0510 | P0_floor | price floor: 0.0510 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-the-us-invade-iran-before-2027 | geopolitics | 0.2000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | 9pt0-or-above-earthquake-before-2027 | weather-natural | 0.0780 | P0_floor | price floor: 0.0780 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-israel-strike-11-countries-in-2026 | geopolitics | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | jerome-powell-out-from-fed-board-by-december-31 | other | 0.4400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-israel-strike-10-countries-in-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-the-us-invade-mexico-in-2026 | geopolitics | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | vanta-ipo-before-2027 | crypto-launch | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-russia-capture-all-of-kostyantynivka-by-june-30-2026 | geopolitics | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026 | entertainment | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-any-category-5-hurricane-make-landfall-in-the-us-in-before-2027 | weather-natural | 0.1300 | P3_low_absolute_liquidity | liquidity $1938 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | us-obtains-iranian-enriched-uranium-by-december-31-725-733 | other | 0.2200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | us-strike-on-cuba-by-december-31 | other | 0.4900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-china-invade-taiwan-before-2027 | geopolitics | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | russia-x-ukraine-ceasefire-agreement-by-december-31-2026 | geopolitics | 0.4500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4992 < absolute min $5000 |
| 2026-05-27T11:30:04Z | save-act-signed-into-law-in-2026 | other | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-us | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0550 | P0_floor | price floor: 0.0550 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-2026-rank-as-the-sixth-hottest-year-on-record-or-lower | other | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $2180 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-viking-therapeutics-be-acquired-before-2027-542-549 | other | 0.3900 | P3_low_absolute_liquidity | liquidity $3687 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-alireza-arafi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-larry-page-be-richest-person-on-december-31 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.6410 | P3_low_absolute_liquidity | liquidity $311 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-the-feds-upper-bound-reach-5pt25-or-higher-before-2027-679 | other | 0.0340 | P0_floor | price floor: 0.0340 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1768 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-beyonc-be-the-top-spotify-artist-for-2026 | entertainment | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-scream-7-be-the-top-grossing-movie-of-2026 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-mamdani-freeze-nyc-rents-before-2027 | other | 0.3000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | major-meteor-strike-10kt-in-2026 | weather-natural | 0.1300 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-27T11:30:04Z | will-israel-strike-6-countries-in-2026 | geopolitics | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-israel-strike-8-countries-in-2026 | geopolitics | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | delcy-rodrguez-out-as-leader-of-venezuela-by-december-31-2026-148-853 | other | 0.1300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-openai-be-acquired-before-2027-859 | other | 0.0730 | P0_floor | price floor: 0.0730 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-alberta-join-the-us | other | 0.0470 | P0_floor | price floor: 0.0470 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-israel-strike-7-countries-in-2026 | geopolitics | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1190 | P3_low_absolute_liquidity | liquidity $3106 < absolute min $5000 |
| 2026-05-27T11:30:04Z | 10pt0-or-above-earthquake-before-2027 | weather-natural | 0.0490 | P0_floor | price floor: 0.0490 < 0.100 (horizon 217 d) |
| 2026-05-27T11:30:04Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3193 < absolute min $5000 |
| 2026-05-27T11:30:04Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 218 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | metamask-fdv-above-1b-one-day-after-launch-978-851-628-634 | crypto-launch | 0.2100 | P4_pre_event | pre-event slug + 218 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | okx-ipo-in-2026 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 218 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | will-opensea-launch-a-token-by-june-30-2026 | crypto-launch | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-pumpfun-perform-an-airdrop-by-december-31-2026 | crypto-launch | 0.2800 | P3_low_absolute_liquidity | liquidity $1740 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-hyperliquid-dip-to-12-by-december-31-2026 | other | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | puffpaw-fdv-above-400m-one-day-after-launch | crypto-launch | 0.1610 | P4_pre_event | pre-event slug + 218 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | will-ethereum-reach-6000-by-december-31-2026 | market | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | hyperbeat-fdv-above-300m-one-day-after-launch-156 | crypto-launch | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | microstrategy-sells-any-bitcoin-by-december-31-2026 | market | 0.7800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4200 | P3_low_absolute_liquidity | liquidity $1084 < absolute min $5000 |
| 2026-05-27T11:30:04Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | crypto-launch | 0.0480 | P0_floor | price floor: 0.0480 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | altcoin-market-cap-dip-to-150b-before-2027 | other | 0.7800 | P3_low_absolute_liquidity | liquidity $705 < absolute min $5000 |
| 2026-05-27T11:30:04Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 218 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-fomo-launch-a-token-by-december-31-2026 | crypto-launch | 0.3600 | P3_low_absolute_liquidity | liquidity $1170 < absolute min $5000 |
| 2026-05-27T11:30:04Z | usdc-depeg-by-december-31-968 | other | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2600 | P3_low_absolute_liquidity | liquidity $1625 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $1296 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-stablecoins-hit-500b-before-2027 | other | 0.1500 | P3_low_absolute_liquidity | liquidity $3159 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0480 | P0_floor | price floor: 0.0480 < 0.100 (horizon 218 d) |
| 2026-05-27T11:30:04Z | will-anthropics-valuation-hit-high-2pt0t-by-december-31-922 | other | 0.4700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-any-month-of-2026-be-the-hottest-on-record | other | 0.8800 | P3_low_absolute_liquidity | liquidity $1443 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | sports-season | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | sports-season | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-atlanta-falcons-win-the-2027-nfl-nfc-championship-312 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-minnesota-vikings-win-the-2027-nfl-nfc-championship-884 | sports-season | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | sports-season | 0.0580 | P0_floor | price floor: 0.0580 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | sports-season | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | sports-season | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-pittsburgh-steelers-win-the-2027-nfl-afc-championship-898 | sports-season | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | sports-season | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-tampa-bay-buccaneers-win-the-2027-nfl-nfc-championship-312 | sports-season | 0.0370 | P0_floor | price floor: 0.0370 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | sports-season | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 242 d) |
| 2026-05-27T11:30:04Z | will-saquon-barkley-win-the-2026-nfl-mvp | sports-season | 0.0300 | P0_floor | price floor: 0.0300 < 0.100 (horizon 263 d) |
| 2026-05-27T11:30:04Z | will-raphal-glucksmann-win-the-2027-french-presidential-election | elections | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 337 d) |
| 2026-05-27T11:30:04Z | will-utah-jazz-win-the-2027-nba-finals | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 399 d) |
| 2026-05-27T11:30:04Z | will-sacramento-kings-win-the-2027-nba-finals | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 399 d) |
| 2026-05-27T11:30:04Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 582 d) |
| 2026-05-27T11:30:04Z | will-spacexs-market-cap-be-between-1pt4t-and-1pt6t-at-market-close-on-ipo-day-168 | crypto-launch | 0.0460 | P0_floor | price floor: 0.0460 < 0.100 (horizon 582 d) |
| 2026-05-27T11:30:04Z | will-spacexs-market-cap-be-between-1pt8t-and-2pt0t-at-market-close-on-ipo-day-516 | crypto-launch | 0.1600 | P4_pre_event | pre-event slug + 582 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | will-spacexs-market-cap-be-between-800b-and-900b-at-market-close-on-ipo-day | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 582 d) |
| 2026-05-27T11:30:04Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7100 | P3_low_absolute_liquidity | liquidity $3066 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 582 d) |
| 2026-05-27T11:30:04Z | will-spacex-not-ipo-by-december-31-2027-481-338 | crypto-launch | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 582 d) |
| 2026-05-27T11:30:04Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 582 d) |
| 2026-05-27T11:30:04Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P4_pre_event | pre-event slug + 582 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.3200 | P4_pre_event | pre-event slug + 583 d to resolution (>=7 threshold) |
| 2026-05-27T11:30:04Z | will-greg-abbott-win-the-2028-republican-presidential-nomination | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-thomas-massie-win-the-2028-us-presidential-election | elections | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-oprah-winfrey-win-the-2028-democratic-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-chris-murphy-win-the-2028-democratic-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-eric-trump-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-jamie-dimon-win-the-2028-us-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-mark-kelly-win-the-2028-democratic-presidential-nomination-479 | elections | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-chelsea-clinton-win-the-2028-democratic-presidential-nomination | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-pete-hegseth-win-the-2028-republican-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | elections | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 894 d) |
| 2026-05-27T11:30:04Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9980 | P0_ceiling | price ceiling: 0.9980 > 0.950 |
| 2026-05-27T11:30:04Z | will-rebecca-shepherd-win-the-2026-makerfield-by-election | elections | 0.0550 | P0_floor | price floor: 0.0550 < 0.100 (horizon -1 d) |
| 2026-05-27T11:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-26 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon -1 d) |
| 2026-05-27T11:30:04Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $1565 < absolute min $5000 |
| 2026-05-27T11:30:04Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon -1 d) |
| 2026-05-27T11:30:04Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon -1 d) |
| 2026-05-27T11:30:04Z | will-hyperliquid-perform-an-airdrop-by-december-31-2026 | crypto-launch | 0.2000 | P3_low_absolute_liquidity | liquidity $4989 < absolute min $5000 |
| 2026-05-27T11:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-31-665-831-238 | geopolitics | 0.4100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-27T11:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-june-7-265-824-655 | geopolitics | 0.6800 | M1 | memoria: exact slug match (score 1.00) |
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
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-27T08:25:01Z | T-562186-1779210004342 | uncategorized | 0.9250 | 0.9990 | 6.38 | market_closed |  | 79.73 | medium | 7.64 |
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
| claudemax-websearch | 11 | 2 | 9 | 18.2% | -5.51 | YES |

## Anti-patterns identificados- same-day-price-target · intraday-decay · no-momentum-confirm · short-horizon-chalk · high-entry-low-exit — visto en  (2026-05-26, pnl $-57.55)


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
