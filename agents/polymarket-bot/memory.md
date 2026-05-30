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
| 2026-05-30T21:05:50Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $655 < absolute min $5000 |
| 2026-05-30T21:05:50Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $1452) |
| 2026-05-30T21:05:50Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $49296) |
| 2026-05-30T21:05:50Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $279072) |
| 2026-05-30T21:05:50Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $43521) |
| 2026-05-30T21:05:50Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $46212) |
| 2026-05-30T21:05:50Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.2400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T21:05:50Z | variational-fdv-above-800m-one-day-after-launch-842-835-884 | crypto-launch | 0.3800 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T21:05:50Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1445731) |
| 2026-05-30T21:05:50Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $783765) |
| 2026-05-30T21:05:50Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1165004) |
| 2026-05-30T21:05:50Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2435004) |
| 2026-05-30T21:05:50Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2157046) |
| 2026-05-30T21:05:50Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1720289) |
| 2026-05-30T21:05:50Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2325259) |
| 2026-05-30T21:05:50Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1482836) |
| 2026-05-30T21:05:50Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1072017) |
| 2026-05-30T21:05:50Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1619491) |
| 2026-05-30T21:05:50Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2749233) |
| 2026-05-30T21:05:50Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1299747) |
| 2026-05-30T21:05:50Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $934929) |
| 2026-05-30T21:05:50Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1449649) |
| 2026-05-30T21:05:50Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1058455) |
| 2026-05-30T21:05:50Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $945775) |
| 2026-05-30T21:05:50Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $5009) |
| 2026-05-30T21:05:50Z | us-announces-new-iran-agreementceasefire-extension-by-may-31-665-831-238 | geopolitics | 0.1100 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 4 veces) |
| 2026-05-30T21:05:50Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0400 | P3_low_absolute_liquidity | liquidity $2901 < absolute min $5000 |
| 2026-05-30T21:05:50Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $17406) |
| 2026-05-30T21:05:50Z | us-announces-new-iran-agreementceasefire-extension-by-june-7-265-824-655 | geopolitics | 0.4100 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 4 veces) |
| 2026-05-30T21:05:50Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $1770 < absolute min $5000 |
| 2026-05-30T21:05:50Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12517) |
| 2026-05-30T21:05:50Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $1211) |
| 2026-05-30T21:30:03Z | jeffrey-epstein-foul-play-confirmed-by-december-31-2026 | other | 0.0700 | P2 | mercado ya expiró (endDate=2025-12-31T00:00:00Z, hace 150 días) |
| 2026-05-30T21:30:03Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon -150 d, liq $41507) |
| 2026-05-30T21:30:03Z | new-stranger-things-episode-released-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -143 d, liq $79954) |
| 2026-05-30T21:30:03Z | will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -111 d, liq $93226) |
| 2026-05-30T21:30:03Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $766685) |
| 2026-05-30T21:30:03Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $776099) |
| 2026-05-30T21:30:03Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -67 d, liq $28018) |
| 2026-05-30T21:30:03Z | will-alex-vanopslagh-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $29859) |
| 2026-05-30T21:30:03Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -67 d, liq $27837) |
| 2026-05-30T21:30:03Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $19647) |
| 2026-05-30T21:30:03Z | israel-x-hamas-ceasefire-phase-ii-by-june-30 | geopolitics | 0.1000 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T21:30:03Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $345220) |
| 2026-05-30T21:30:03Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon -15 d, liq $83690) |
| 2026-05-30T21:30:03Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -13 d, liq $5950) |
| 2026-05-30T21:30:03Z | will-gregg-kirkpatrick-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -11 d, liq $10202) |
| 2026-05-30T21:30:03Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -11 d, liq $6309) |
| 2026-05-30T21:30:03Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -3 d, liq $1706150) |
| 2026-05-30T21:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 0 d, liq $47946) |
| 2026-05-30T21:30:03Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $228820) |
| 2026-05-30T21:30:03Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $228229) |
| 2026-05-30T21:30:03Z | chi-qin-sgr-2026-05-30-qin | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $296925) |
| 2026-05-30T21:30:03Z | will-the-price-of-bitcoin-be-between-68000-70000-on-may-30-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $459266) |
| 2026-05-30T21:30:03Z | ucl-psg-ars-2026-05-30-psg | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $7691933) |
| 2026-05-30T21:30:03Z | elon-musk-of-tweets-may-28-may-30-90-114 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $270065) |
| 2026-05-30T21:30:05Z | bra-bah-bot-2026-05-30-bot | other | 0.5500 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T21:30:05Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $197651) |
| 2026-05-30T21:30:05Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $220241) |
| 2026-05-30T21:30:05Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.1200 | P9 | P9: geopolitics pump cluster (price 0.12, 0d) |
| 2026-05-30T21:30:05Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0430 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-30T21:30:05Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.0370 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-30T21:30:05Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.2800 | P9 | P9: geopolitics pump cluster (price 0.28, 0d) |
| 2026-05-30T21:30:05Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0200 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-30T21:30:05Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $193458) |
| 2026-05-30T21:30:05Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $3796641) |
| 2026-05-30T21:30:05Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 0 d, liq $150541) |
| 2026-05-30T21:30:05Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 0 d, liq $114488) |
| 2026-05-30T21:30:05Z | will-arsenal-win-the-202526-champions-league | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $3258477) |
| 2026-05-30T21:30:05Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $161978) |
| 2026-05-30T21:30:05Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0260 | M2 | M2 soft-learned: geopolitics·short·<0.10 = 6L/0W (wr 0%) |
| 2026-05-30T21:30:05Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 0 d, liq $63877) |
| 2026-05-30T21:30:05Z | will-alibaba-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $135745) |
| 2026-05-30T21:30:05Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $186599) |
| 2026-05-30T21:30:05Z | will-80-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 0 d, liq $76223) |
| 2026-05-30T21:30:05Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 0 d, liq $903635) |
| 2026-05-30T21:30:05Z | netanyahu-out-by-may-31 | executive-action | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $69035) |
| 2026-05-30T21:30:06Z | will-bitcoin-reach-80k-on-may-30 | market | 0.0060 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:06Z | will-alphabet-be-the-third-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $15965) |
| 2026-05-30T21:30:06Z | will-nvidia-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $38077) |
| 2026-05-30T21:30:06Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $91383) |
| 2026-05-30T21:30:06Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $93044) |
| 2026-05-30T21:30:08Z | will-abelardo-de-la-espriella-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.2180 | N1 | noticias contradicen tesis: Encuestas sitúan a De la Espriella 2º (~21-28%) muy lejos del 50%+1 necesario para gana... |
| 2026-05-30T21:30:08Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $108687) |
| 2026-05-30T21:30:08Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $108251) |
| 2026-05-30T21:30:08Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $118244) |
| 2026-05-30T21:30:08Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $89568) |
| 2026-05-30T21:30:08Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $139939) |
| 2026-05-30T21:30:08Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $149634) |
| 2026-05-30T21:30:08Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $135266) |
| 2026-05-30T21:30:08Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $141202) |
| 2026-05-30T21:30:08Z | bitcoin-above-72k-on-may-31-2026 | market | 0.9910 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:08Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $149889) |
| 2026-05-30T21:30:08Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $139819) |
| 2026-05-30T21:30:08Z | es2-rso-leo-2026-05-31-rso | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $308170) |
| 2026-05-30T21:30:08Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 1 d, liq $350751) |
| 2026-05-30T21:30:08Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.0010 | R1_longshot | R1 longshot: precio 0.0010 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.619 < 0.15) |
| 2026-05-30T21:30:08Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0010 | R1_longshot | R1 longshot: precio 0.0010 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.110 < 0.15) |
| 2026-05-30T21:30:08Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0010 | E2 | edge 0.013 < mín 0.030 (p̂=0.014, implied=0.001) |
| 2026-05-30T21:30:08Z | will-wti-dip-to-70-in-may-2026-155-395-333-182-889-988-959-341-883-234-375-347-135 | market | 0.0010 | E2 | edge 0.009 < mín 0.030 (p̂=0.010, implied=0.001) |
| 2026-05-30T21:30:08Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $108920) |
| 2026-05-30T21:30:11Z | will-xrp-reach-1pt6-in-may-2026 | other | 0.0300 | R1_longshot | R1 longshot: precio 0.0300 < 0.10 sin edge fuerte (edge_type=info, edge=0.120 < 0.15) |
| 2026-05-30T21:30:11Z | will-ethereum-reach-2600-in-may-2026 | market | 0.0010 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.00) |
| 2026-05-30T21:30:11Z | will-bitcoin-dip-to-35k-in-may-2026-217-769-834 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $78029) |
| 2026-05-30T21:30:11Z | will-bitcoin-dip-to-72k-may-25-31-2026 | market | 0.0600 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:11Z | will-bitcoin-dip-to-72500-in-may-2026-from-may-27 | market | 0.1800 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:11Z | will-bitcoin-dip-to-65k-in-may-2026-183-857-425 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $92977) |
| 2026-05-30T21:30:11Z | will-bitcoin-reach-77500-in-may-2026-from-may-27 | market | 0.0190 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:11Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $175751) |
| 2026-05-30T21:30:11Z | over-30m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $13413) |
| 2026-05-30T21:30:11Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 1 d, liq $83532) |
| 2026-05-30T21:30:11Z | over-40m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $14263) |
| 2026-05-30T21:30:11Z | will-ethereum-dip-to-400-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $33303) |
| 2026-05-30T21:30:11Z | will-ethereum-dip-to-1600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $41392) |
| 2026-05-30T21:30:11Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $28704) |
| 2026-05-30T21:30:11Z | over-60m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $16084) |
| 2026-05-30T21:30:11Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0010 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:11Z | will-bitcoin-reach-80000-in-may-2026-from-may-27 | market | 0.0060 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:30:17Z | will-the-breadwinner-opening-weekend-box-office-be-greater-than-7m | entertainment | 0.9000 | N1 | noticias contradicen tesis: Proyecciones sitúan apertura en ~$8M desde 3.300 cines; supera umbral de $7M pero con pr... |
| 2026-05-30T21:30:19Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2500 | N1 | noticias contradicen tesis: Encuestas al 28-may sitúan a Pratt 3º (22%) tras Bass (26%) y Raman (25%); primaria el ... |
| 2026-05-30T21:30:19Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $17941) |
| 2026-05-30T21:30:19Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $38049) |
| 2026-05-30T21:30:20Z | elon-musk-of-tweets-may-26-june-2-360-379 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $66301) |
| 2026-05-30T21:30:20Z | elon-musk-of-tweets-may-26-june-2-300-319 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $60893) |
| 2026-05-30T21:30:20Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $66314) |
| 2026-05-30T21:30:20Z | elon-musk-of-tweets-may-26-june-2-500plus | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $90952) |
| 2026-05-30T21:30:20Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $26915) |
| 2026-05-30T21:30:20Z | will-na-kyung-won-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $478919) |
| 2026-05-30T21:30:20Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $31742) |
| 2026-05-30T21:30:20Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $39411) |
| 2026-05-30T21:30:20Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $41608) |
| 2026-05-30T21:30:20Z | will-lee-chul-gyu-win-the-2026-gangwon-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $16141) |
| 2026-05-30T21:30:20Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $32720) |
| 2026-05-30T21:30:20Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $20926) |
| 2026-05-30T21:30:20Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $43380) |
| 2026-05-30T21:30:20Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $45362) |
| 2026-05-30T21:30:20Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.2500 | P9 | P9: geopolitics pump cluster (price 0.25, 3d) |
| 2026-05-30T21:30:20Z | will-yoon-sang-hyun-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $38250) |
| 2026-05-30T21:30:20Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $44684) |
| 2026-05-30T21:30:20Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $44027) |
| 2026-05-30T21:30:20Z | will-han-dong-hoon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $33033) |
| 2026-05-30T21:30:23Z | will-oh-se-hoon-win-the-2026-seoul-mayoral-election | elections | 0.2000 | N1 | noticias contradicen tesis: Encuestas de mayo muestran empate técnico: Chong 43% vs Oh 42.6%; Chong lidera levemente... |
| 2026-05-30T21:30:23Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $97247) |
| 2026-05-30T21:30:25Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $94955) |
| 2026-05-30T21:30:25Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $32173) |
| 2026-05-30T21:30:25Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $45030) |
| 2026-05-30T21:30:25Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $22757) |
| 2026-05-30T21:30:25Z | will-chung-jin-suk-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $19745) |
| 2026-05-30T21:30:25Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $230113) |
| 2026-05-30T21:30:25Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $209438) |
| 2026-05-30T21:30:25Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $108213) |
| 2026-05-30T21:30:27Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $201634) |
| 2026-05-30T21:30:27Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $159740) |
| 2026-05-30T21:30:27Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $213266) |
| 2026-05-30T21:30:31Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0270 | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... |
| 2026-05-30T21:30:31Z | will-raphael-collignon-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $57412) |
| 2026-05-30T21:30:31Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $220012) |
| 2026-05-30T21:30:33Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-7-2026 | geopolitics | 0.3500 | V3, V5, V6 | V3, V5, V6: V3 Trigger vago: no hay una fecha concreta de anuncio ni evento verificable (la fecha de resolución es f... |
| 2026-05-30T21:30:33Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $221788) |
| 2026-05-30T21:30:33Z | will-ricardo-belmont-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $94448) |
| 2026-05-30T21:30:33Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $197657) |
| 2026-05-30T21:30:36Z | will-casper-ruud-win-the-2026-mens-french-open | other | 0.1380 | N2 | silencio mediático sobre catalyst inminente (7 días al cierre) |
| 2026-05-30T21:30:36Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.1300 | P9 | P9: geopolitics pump cluster (price 0.13, 7d) |
| 2026-05-30T21:30:38Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.1140 | V3 | V3: V3 Trigger vago: la pregunta no tiene una fecha concreta ni un evento verificable próximo. Depende del resultado... |
| 2026-05-30T21:30:38Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 7 d, liq $38163) |
| 2026-05-30T21:30:38Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $189662) |
| 2026-05-30T21:30:41Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0560 | R1_longshot | R1 longshot: precio 0.0560 < 0.10 sin edge fuerte (edge_type=info, edge=0.036 < 0.15) |
| 2026-05-30T21:30:41Z | will-martin-landaluce-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $331522) |
| 2026-05-30T21:30:41Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $244775) |
| 2026-05-30T21:30:41Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $211926) |
| 2026-05-30T21:30:41Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $240567) |
| 2026-05-30T21:30:43Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0900 | R1_longshot | R1 longshot: precio 0.0900 < 0.10 sin edge fuerte (edge_type=info, edge=0.040 < 0.15) |
| 2026-05-30T21:30:43Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $247187) |
| 2026-05-30T21:30:43Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $160007) |
| 2026-05-30T21:30:43Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $119824) |
| 2026-05-30T21:30:43Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $227067) |
| 2026-05-30T21:30:45Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.1900 | R5_precedents | R5 precedentes: 0 < 3 casos análogos |
| 2026-05-30T21:30:47Z | strait-of-hormuz-traffic-returns-to-normal-by-june-15 | geopolitics | 0.0900 | R1_longshot | R1 longshot: precio 0.0900 < 0.10 sin edge fuerte (edge_type=info, edge=0.060 < 0.15) |
| 2026-05-30T21:30:50Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-15-2026 | geopolitics | 0.5000 | R5_precedents | R5 precedentes: 0 < 3 casos análogos |
| 2026-05-30T21:30:50Z | will-al-mina-be-the-republican-nominee-for-senate-in-virginia | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 16 d, liq $21205) |
| 2026-05-30T21:30:52Z | will-the-san-antonio-spurs-win-the-nba-western-conference-finals | sports-season | 0.4280 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. V3 Trigger vago: sin fecha concreta o... | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. V3 Trigger vago: sin fecha concreta o... |
| 2026-05-30T21:30:52Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 17 d, liq $2589918) |
| 2026-05-30T21:30:52Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 17 d, liq $814386) |
| 2026-05-30T21:30:52Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 17 d, liq $272803) |
| 2026-05-30T21:30:52Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9810 | P0_ceiling | price ceiling: 0.9810 > 0.980 |
| 2026-05-30T21:30:52Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 17 d, liq $601299) |
| 2026-05-30T21:30:52Z | will-thunder-downunder-win-iem-cologne-major-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $78699) |
| 2026-05-30T21:30:52Z | will-sinners-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $44390) |
| 2026-05-30T21:30:52Z | will-gaimin-gladiators-win-iem-cologne-major-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $85556) |
| 2026-05-30T21:30:52Z | will-flyquest-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $77127) |
| 2026-05-30T21:30:52Z | will-nrg-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $80825) |
| 2026-05-30T21:30:54Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0290 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T21:30:54Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $228855) |
| 2026-05-30T21:30:54Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $182392) |
| 2026-05-30T21:30:54Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $65310) |
| 2026-05-30T21:30:54Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $154918) |
| 2026-05-30T21:30:54Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $212836) |
| 2026-05-30T21:30:57Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6300 | V3 (Trigger vago), V6 (Sin catalyst) | V3 (Trigger vago), V6 (Sin catalyst): Falta fecha concreta y evento discreto verificable para un mercado que resuelve... |
| 2026-05-30T21:31:00Z | will-ivan-cepeda-castro-win-the-2026-colombian-presidential-election | elections | 0.3600 | N1 | noticias contradicen tesis: Cepeda lidera primera vuelta (~44%) pero no alcanza 50%; segunda vuelta es inevitable y e... |
| 2026-05-30T21:31:00Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $223284) |
| 2026-05-30T21:31:00Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $197496) |
| 2026-05-30T21:31:00Z | will-juan-daniel-oviedo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $255784) |
| 2026-05-30T21:31:00Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $200376) |
| 2026-05-30T21:31:00Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $193656) |
| 2026-05-30T21:31:02Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E2 | edge 0.010 < mín 0.030 (p̂=0.030, implied=0.020) |
| 2026-05-30T21:31:02Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 24 d, liq $534) |
| 2026-05-30T21:31:02Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 24 d, liq $382) |
| 2026-05-30T21:31:02Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 24 d, liq $1075) |
| 2026-05-30T21:31:02Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 27 d, liq $1558) |
| 2026-05-30T21:31:05Z | mojtaba-khamenei-leaves-iran-by-june-30-2026 | geopolitics | 0.0270 | E2 | edge 0.003 < mín 0.030 (p̂=0.030, implied=0.027) |
| 2026-05-30T21:31:05Z | will-trump-and-putin-meet-next-in-turkey-182-161-789 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 30 d, liq $18538) |
| 2026-05-30T21:31:05Z | will-discords-market-cap-be-less-than-15b-at-market-close-on-ipo-day | crypto-launch | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 30 d, liq $3628) |
| 2026-05-30T21:31:05Z | will-michael-jackson-be-confirmed-to-have-visited-epsteins-island | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 30 d, liq $21564) |
| 2026-05-30T21:31:07Z | us-obtains-iranian-enriched-uranium-by-june-30 | other | 0.0500 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T21:31:07Z | anthropic-ceo-arrested | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $7983) |
| 2026-05-30T21:31:07Z | tom-holland-announced-as-next-james-bond-679 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 30 d, liq $4468) |
| 2026-05-30T21:31:07Z | forsen-to-get-signed-by-an-lck-prog-org-by-jun-30 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 30 d, liq $48644) |
| 2026-05-30T21:31:09Z | will-donald-trump-attend-the-next-us-x-iran-diplomatic-meeting | geopolitics | 0.0670 | R1_longshot | R1 longshot: precio 0.0670 < 0.10 sin edge fuerte (edge_type=info, edge=0.053 < 0.15) |
| 2026-05-30T21:31:09Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 30 d, liq $84829) |
| 2026-05-30T21:31:13Z | will-apple-be-the-largest-company-in-the-world-by-market-cap-on-june-30-416 | other | 0.0580 | R1_longshot | R1 longshot: precio 0.0580 < 0.10 sin edge fuerte (edge_type=info, edge=0.092 < 0.15) |
| 2026-05-30T21:31:16Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0230 | E2 | edge 0.003 < mín 0.030 (p̂=0.020, implied=0.023) |
| 2026-05-30T21:31:18Z | will-putin-meet-with-zelenskyy-by-june-30-2026 | other | 0.0260 | V3 Trigger vago | V3 Trigger vago: V3 Trigger vago: el evento 'Putin meet with Zelenskyy' no tiene fecha concreta ni evento diplomátic... |
| 2026-05-30T21:31:23Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0340 | R1_longshot | R1 longshot: precio 0.0340 < 0.10 sin edge fuerte (edge_type=info, edge=0.066 < 0.15) |
| 2026-05-30T21:31:23Z | will-us-withdraw-from-nato-by-june-30 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 30 d, liq $27674) |
| 2026-05-30T21:31:25Z | us-iran-nuclear-deal-by-june-30 | geopolitics | 0.3700 | N1 | noticias contradicen tesis: Solo hay un borrador de MOU de 60 días pendiente de Trump; un acuerdo nuclear completo a... |
| 2026-05-30T21:31:25Z | will-trump-and-putin-meet-next-in-south-korea-582-297-268 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $22265) |
| 2026-05-30T21:31:28Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-30-2026-15... | geopolitics | 0.6800 | R5_precedents | R5 precedentes: 2 < 3 casos análogos |
| 2026-05-30T21:31:31Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $16576) |
| 2026-05-30T21:31:31Z | israel-withdraws-from-lebanon-by-may-31-2026 | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 30 d, liq $52083) |
| 2026-05-30T21:31:31Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $373098) |
| 2026-05-30T21:31:33Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0320 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T21:31:33Z | harris-dickinson-announced-as-next-james-bond-784 | other | 0.0460 | P3_low_absolute_liquidity | liquidity $2755 < absolute min $5000 |
| 2026-05-30T21:31:33Z | will-trump-and-putin-meet-next-in-belarus-572 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $18800) |
| 2026-05-30T21:31:33Z | will-amazon-be-the-largest-company-in-the-world-by-market-cap-on-june-30-123 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $387386) |
| 2026-05-30T21:31:33Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $336146) |
| 2026-05-30T21:31:33Z | will-trump-and-putin-meet-next-in-finland-772-412 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 30 d, liq $17013) |
| 2026-05-30T21:31:36Z | trump-renames-ice-to-nice-by-june-30 | other | 0.0410 | E2 | edge 0.021 < mín 0.030 (p̂=0.020, implied=0.041) |
| 2026-05-30T21:31:36Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-june-30-612 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $247135) |
| 2026-05-30T21:31:38Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La pregunta depende ... |
| 2026-05-30T21:31:38Z | lai-ching-te-impeached-by-june-30 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 30 d, liq $23391) |
| 2026-05-30T21:31:38Z | henry-cavill-announced-as-next-james-bond-857-642 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 30 d, liq $4645) |
| 2026-05-30T21:31:38Z | will-openais-market-cap-be-between-1t-and-1pt25t-at-market-close-on-ipo-day | crypto-launch | 0.1450 | P4_pre_event | pre-event slug + 30 d to resolution (>=7 threshold) |
| 2026-05-30T21:31:41Z | european-country-agrees-to-give-ukraine-security-guarantee-by-june-30 | geopolitics | 0.0400 | R1_longshot | R1 longshot: precio 0.0400 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.040 < 0.15) |
| 2026-05-30T21:31:41Z | will-trump-and-putin-meet-next-in-another-country-313-781-734-447 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 30 d, liq $21013) |
| 2026-05-30T21:31:43Z | iran-closes-its-airspace-by-june-30-432-786-462-866 | geopolitics | 0.2180 | N1 | noticias contradicen tesis: Tregua violada el 26 mayo con nuevos ataques EEUU e Irán; escalada reciente eleva riesgo... |
| 2026-05-30T21:31:43Z | will-openais-market-cap-be-less-than-500b-at-market-close-on-ipo-day | crypto-launch | 0.0210 | P4_pre_event | pre-event slug + 30 d to resolution (>=7 threshold) |
| 2026-05-30T21:31:43Z | will-trump-and-putin-meet-next-in-united-states-224-875-469 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 30 d, liq $18724) |
| 2026-05-30T21:31:43Z | us-x-russia-military-clash-by-june-30-2026-249 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 30 d, liq $24310) |
| 2026-05-30T21:31:43Z | will-any-country-leave-nato-by-june-30-2026 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 30 d, liq $23620) |
| 2026-05-30T21:31:43Z | will-silver-si-hit-high-130-by-end-of-june | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $34625) |
| 2026-05-30T21:31:43Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 30 d, liq $31870) |
| 2026-05-30T21:31:43Z | will-gold-gc-hit-high-6500-by-end-of-june | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 30 d, liq $23342) |
| 2026-05-30T21:31:43Z | will-silver-si-hit-high-150-by-end-of-june | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $37065) |
| 2026-05-30T21:31:43Z | will-silver-si-hit-low-45-by-end-of-june-972-272 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 30 d, liq $14355) |
| 2026-05-30T21:31:43Z | gc-hit-8000-high-jun-2026-342-647-753 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 30 d, liq $31016) |
| 2026-05-30T21:31:43Z | will-crude-oil-cl-hit-low-60-by-end-of-june-529 | market | 0.0270 | R1_longshot | R1 longshot: precio 0.0270 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.111 < 0.15) |
| 2026-05-30T21:31:44Z | will-crude-oil-cl-hit-low-50-by-end-of-june | market | 0.0150 | R1_longshot | R1 longshot: precio 0.0150 < 0.10 sin edge fuerte (edge_type=model-barrier, edge=0.030 < 0.15) |
| 2026-05-30T21:31:46Z | kharg-island-no-longer-under-iranian-control-by-june-30-561-854 | other | 0.0500 | R1_longshot | R1 longshot: precio 0.0500 < 0.10 sin edge fuerte (edge_type=info, edge=0.030 < 0.15) |
| 2026-05-30T21:31:46Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.2590 | E2 | edge 0.003 < mín 0.030 (p̂=0.256, implied=0.259) |
| 2026-05-30T21:31:46Z | will-the-new-york-knicks-win-the-2026-nba-finals | sports-season | 0.3230 | E2 | edge 0.015 < mín 0.030 (p̂=0.308, implied=0.323) |
| 2026-05-30T21:31:46Z | will-the-oklahoma-city-thunder-win-the-2026-nba-finals | sports-season | 0.4200 | E2 | edge 0.016 < mín 0.030 (p̂=0.436, implied=0.420) |
| 2026-05-30T21:31:48Z | microstrategy-sells-any-bitcoin-by-may-31-2026 | market | 0.1350 | N1 | noticias contradicen tesis: Strategy ha declarado intención de vender pero no ha ejecutado ninguna venta; sigue comp... |
| 2026-05-30T21:31:49Z | will-bitcoin-hit-150k-by-june-30-2026 | market | 0.0070 | P6 | P6 market: BTC-USD spot $73875.17 already > target $150.00 but yes=0.01 (confused book) |
| 2026-05-30T21:31:51Z | will-a-team-from-lcp-asia-pacific-win-msi-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 42 d, liq $6637) |
| 2026-05-30T21:31:53Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $606 < absolute min $5000 |
| 2026-05-30T21:31:53Z | will-france-win-the-2026-fifa-world-cup-924 | sports-season | 0.1640 | E2 | edge 0.010 < mín 0.030 (p̂=0.154, implied=0.164) |
| 2026-05-30T21:31:53Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.002 < mín 0.030 (p̂=0.054, implied=0.052) |
| 2026-05-30T21:31:53Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | E2 | edge 0.000 < mín 0.030 (p̂=0.002, implied=0.002) |
| 2026-05-30T21:31:53Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | E2 | edge 0.001 < mín 0.030 (p̂=0.019, implied=0.018) |
| 2026-05-30T21:31:53Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | E2 | edge 0.000 < mín 0.030 (p̂=0.001, implied=0.001) |
| 2026-05-30T21:31:53Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | E2 | edge 0.003 < mín 0.030 (p̂=0.006, implied=0.009) |
| 2026-05-30T21:31:53Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | E2 | edge 0.001 < mín 0.030 (p̂=0.005, implied=0.006) |
| 2026-05-30T21:31:53Z | will-new-zealand-win-the-2026-fifa-world-cup-635 | sports-season | 0.0010 | E2 | edge 0.000 < mín 0.030 (p̂=0.001, implied=0.001) |
| 2026-05-30T21:31:53Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | E2 | edge 0.001 < mín 0.030 (p̂=0.001, implied=0.002) |
| 2026-05-30T21:31:53Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | E2 | edge 0.000 < mín 0.030 (p̂=0.003, implied=0.003) |
| 2026-05-30T21:31:53Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | E2 | edge 0.000 < mín 0.030 (p̂=0.002, implied=0.002) |
| 2026-05-30T21:31:53Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | E2 | edge 0.001 < mín 0.030 (p̂=0.002, implied=0.003) |
| 2026-05-30T21:31:53Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | E2 | edge 0.001 < mín 0.030 (p̂=0.004, implied=0.003) |
| 2026-05-30T21:31:53Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | E2 | edge 0.000 < mín 0.030 (p̂=0.001, implied=0.001) |
| 2026-05-30T21:31:53Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8521727) |
| 2026-05-30T21:31:53Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | E2 | edge 0.000 < mín 0.030 (p̂=0.011, implied=0.011) |
| 2026-05-30T21:31:53Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | E2 | edge 0.001 < mín 0.030 (p̂=0.006, implied=0.007) |
| 2026-05-30T21:31:53Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | E2 | edge 0.001 < mín 0.030 (p̂=0.002, implied=0.003) |
| 2026-05-30T21:31:53Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0200 | E2 | edge 0.005 < mín 0.030 (p̂=0.015, implied=0.020) |
| 2026-05-30T21:31:53Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | E2 | edge 0.002 < mín 0.030 (p̂=0.020, implied=0.018) |
| 2026-05-30T21:31:53Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | E2 | edge 0.000 < mín 0.030 (p̂=0.006, implied=0.006) |
| 2026-05-30T21:31:53Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | E2 | edge 0.001 < mín 0.030 (p̂=0.001, implied=0.002) |
| 2026-05-30T21:31:53Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | E2 | edge 0.001 < mín 0.030 (p̂=0.002, implied=0.003) |
| 2026-05-30T21:31:53Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0260 | E2 | edge 0.001 < mín 0.030 (p̂=0.025, implied=0.026) |
| 2026-05-30T21:31:53Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0400 | E2 | edge 0.003 < mín 0.030 (p̂=0.037, implied=0.040) |
| 2026-05-30T21:31:53Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0900 | E2 | edge 0.001 < mín 0.030 (p̂=0.091, implied=0.090) |
| 2026-05-30T21:31:53Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | E2 | edge 0.001 < mín 0.030 (p̂=0.011, implied=0.012) |
| 2026-05-30T21:31:53Z | will-england-win-the-2026-fifa-world-cup-937 | sports-season | 0.1130 | E2 | edge 0.006 < mín 0.030 (p̂=0.119, implied=0.113) |
| 2026-05-30T21:31:53Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | E2 | edge 0.001 < mín 0.030 (p̂=0.003, implied=0.004) |
| 2026-05-30T21:31:53Z | will-portugal-win-the-2026-fifa-world-cup-912 | sports-season | 0.1010 | E2 | edge 0.022 < mín 0.030 (p̂=0.079, implied=0.101) |
| 2026-05-30T21:31:53Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | E2 | edge 0.001 < mín 0.030 (p̂=0.002, implied=0.003) |
| 2026-05-30T21:31:53Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | E2 | edge 0.002 < mín 0.030 (p̂=0.001, implied=0.003) |
| 2026-05-30T21:31:53Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | E2 | edge 0.001 < mín 0.030 (p̂=0.010, implied=0.011) |
| 2026-05-30T21:31:53Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E2 | edge 0.005 < mín 0.030 (p̂=0.091, implied=0.086) |
| 2026-05-30T21:31:53Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | E2 | edge 0.001 < mín 0.030 (p̂=0.010, implied=0.009) |
| 2026-05-30T21:31:53Z | will-spain-win-the-2026-fifa-world-cup-963 | sports-season | 0.1680 | E2 | edge 0.009 < mín 0.030 (p̂=0.159, implied=0.168) |
| 2026-05-30T21:31:53Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 50 d, liq $3646069) |
| 2026-05-30T21:31:53Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | E2 | edge 0.001 < mín 0.030 (p̂=0.014, implied=0.015) |
| 2026-05-30T21:31:53Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | E2 | edge 0.002 < mín 0.030 (p̂=0.010, implied=0.012) |
| 2026-05-30T21:31:53Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0020 | E2 | edge 0.001 < mín 0.030 (p̂=0.001, implied=0.002) |
| 2026-05-30T21:31:55Z | us-x-iran-permanent-peace-deal-by-july-31-2026-831-252 | geopolitics | 0.4400 | R5_precedents | R5 precedentes: 0 < 3 casos análogos |
| 2026-05-30T21:31:57Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $35424) |
| 2026-05-30T21:31:57Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $16441) |
| 2026-05-30T21:31:57Z | will-scott-jensen-win-the-2026-minnesota-governor-republican-primary-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 72 d, liq $11527) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-miami-dolphins-in-2026-27 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 93 d, liq $2941) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-denver-broncos-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $1493) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-new-orleans-saints-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $2457) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-tennessee-titans-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $1305) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0230 | P3_low_absolute_liquidity | liquidity $1003 < absolute min $5000 |
| 2026-05-30T21:31:57Z | will-filipe-luis-be-the-next-manager-of-ssc-napoli | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 93 d, liq $998) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-kansas-city-chiefs-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $3063) |
| 2026-05-30T21:31:57Z | will-david-njoku-play-for-indianapolis-colts-in-2026-27 | other | 0.0200 | P3_low_absolute_liquidity | liquidity $500 < absolute min $5000 |
| 2026-05-30T21:31:59Z | will-cdu-win-the-most-seats-in-the-2026-sachsen-anhalt-parliamentary-elections | other | 0.0570 | R1_longshot | R1 longshot: precio 0.0570 < 0.10 sin edge fuerte (edge_type=info, edge=0.063 < 0.15) |
| 2026-05-30T21:31:59Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 105 d, liq $3168) |
| 2026-05-30T21:31:59Z | will-ebba-busch-be-the-next-prime-minister-of-sweden | elections | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 105 d, liq $23782) |
| 2026-05-30T21:31:59Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $1928) |
| 2026-05-30T21:32:02Z | will-the-moderate-party-m-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0360 | E2 | edge 0.014 < mín 0.030 (p̂=0.050, implied=0.036) |
| 2026-05-30T21:32:02Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 105 d, liq $2320) |
| 2026-05-30T21:32:04Z | will-jimmie-kesson-be-the-next-prime-minister-of-sweden | elections | 0.0280 | E2 | edge 0.012 < mín 0.030 (p̂=0.040, implied=0.028) |
| 2026-05-30T21:32:04Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $2823) |
| 2026-05-30T21:32:04Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 120 d, liq $971) |
| 2026-05-30T21:32:04Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $116 < absolute min $5000 |
| 2026-05-30T21:32:04Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $603 < absolute min $5000 |
| 2026-05-30T21:32:04Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 120 d, liq $609) |
| 2026-05-30T21:32:04Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5600 | P8 | P8: election 122d out, price 0.56 en banda ruido [0.30, 0.70] |
| 2026-05-30T21:32:04Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $440 < absolute min $5000 |
| 2026-05-30T21:32:04Z | will-geraldo-alckmin-win-the-2026-brazilian-presidential-election | elections | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 126 d, liq $226450) |
| 2026-05-30T21:32:07Z | will-renan-santos-win-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0910 | E2 | edge 0.029 < mín 0.030 (p̂=0.120, implied=0.091) |
| 2026-05-30T21:32:07Z | will-michelle-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 126 d, liq $228556) |
| 2026-05-30T21:32:07Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 126 d, liq $1099049) |
| 2026-05-30T21:32:07Z | will-carlos-roberto-massa-jnior-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 126 d, liq $27170) |
| 2026-05-30T21:32:09Z | will-renan-santos-win-the-2026-brazilian-presidential-election | elections | 0.1720 | V3 - Trigger vago | V3 - Trigger vago: V3 Trigger vago: el evento 'Renan Santos gane las elecciones presidenciales brasileñas de 2026' d... |
| 2026-05-30T21:32:12Z | will-fernando-haddad-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0600 | E2 | edge 0.020 < mín 0.030 (p̂=0.080, implied=0.060) |
| 2026-05-30T21:32:12Z | will-parti-conservateur-du-qubec-win-the-most-seats-in-the-2026-quebec-general-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 127 d, liq $15225) |
| 2026-05-30T21:32:14Z | will-ousmane-dembl-win-the-2026-ballon-dor | other | 0.2960 | N1 | noticias contradicen tesis: Kane lidera Polymarket (25%) y apuestas (5/2); Dembélé 2º (16-9/2) con Rice, Yamal y O... |
| 2026-05-30T21:32:14Z | will-the-los-angeles-dodgers-win-the-2026-world-series | other | 0.2600 | E2 | edge 0.013 < mín 0.030 (p̂=0.247, implied=0.260) |
| 2026-05-30T21:32:14Z | will-the-miami-marlins-win-the-2026-world-series | other | 0.0080 | E2 | edge 0.005 < mín 0.030 (p̂=0.003, implied=0.008) |
| 2026-05-30T21:32:19Z | will-san-diego-padres-win-the-2026-national-league-championship-series | sports-season | 0.0560 | E2 | edge 0.024 < mín 0.030 (p̂=0.080, implied=0.056) |
| 2026-05-30T21:32:19Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 154 d, liq $26107) |
| 2026-05-30T21:32:21Z | will-cleveland-guardians-win-the-2026-american-league-championship-series | sports-season | 0.0860 | E2 | edge 0.014 < mín 0.030 (p̂=0.100, implied=0.086) |
| 2026-05-30T21:32:21Z | will-st-louis-cardinals-win-the-2026-national-league-championship-series | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 154 d, liq $23500) |
| 2026-05-30T21:32:21Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 154 d, liq $23919) |
| 2026-05-30T21:32:24Z | will-milwaukee-brewers-win-the-2026-national-league-championship-series | sports-season | 0.0810 | R1_longshot | R1 longshot: precio 0.0810 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.031 < 0.15) |
| 2026-05-30T21:32:24Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 154 d, liq $23795) |
| 2026-05-30T21:32:26Z | will-pittsburgh-pirates-win-the-2026-national-league-championship-series | sports-season | 0.0410 | E2 | edge 0.004 < mín 0.030 (p̂=0.045, implied=0.041) |
| 2026-05-30T21:32:26Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 154 d, liq $28990) |
| 2026-05-30T21:32:29Z | will-cincinnati-reds-win-the-2026-national-league-championship-series | sports-season | 0.0260 | E2 | edge 0.014 < mín 0.030 (p̂=0.040, implied=0.026) |
| 2026-05-30T21:32:32Z | will-the-republican-party-hold-exactly-28-or-29-governorships-after-the-2026-midterm-elections | elections | 0.0970 | E2 | edge 0.027 < mín 0.030 (p̂=0.070, implied=0.097) |
| 2026-05-30T21:32:32Z | will-betty-yee-win-the-california-governor-election-in-2026 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 156 d, liq $224763) |
| 2026-05-30T21:32:35Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 156 d, liq $6735) |
| 2026-05-30T21:32:35Z | will-david-brekalo-win-2026-mls-defender-of-the-year | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 165 d, liq $325) |
| 2026-05-30T21:32:35Z | will-jakob-glesnes-win-2026-mls-defender-of-the-year | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 165 d, liq $290) |
| 2026-05-30T21:32:35Z | will-richard-van-de-water-win-the-bachelorette-season-22 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 183 d, liq $3860) |
| 2026-05-30T21:32:35Z | will-lew-evans-win-the-bachelorette-season-22 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 183 d, liq $6959) |
| 2026-05-30T21:32:35Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 192 d, liq $12161) |
| 2026-05-30T21:32:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-4pt0-at-the-end-of-2026-386 | other | 0.2340 | N1 | noticias contradicen tesis: Upper bound actual 3,75%; Fed en ciclo de recortes tras reunión de abril, improbable reg... |
| 2026-05-30T21:32:41Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-3pt5-at-the-end-of-2026-352 | other | 0.0880 | R1_longshot | R1 longshot: precio 0.0880 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.032 < 0.15) |
| 2026-05-30T21:32:43Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-3pt0-at-the-end-of-2026-593 | other | 0.0380 | R1_longshot | R1 longshot: precio 0.0380 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.042 < 0.15) |
| 2026-05-30T21:32:43Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9280) |
| 2026-05-30T21:32:43Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9065) |
| 2026-05-30T21:32:43Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 192 d, liq $12245) |
| 2026-05-30T21:32:43Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9876) |
| 2026-05-30T21:32:43Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 192 d, liq $10013) |
| 2026-05-30T21:32:45Z | will-san-jose-earthquakes-win-the-2026-mls-cup | other | 0.0460 | E2 | edge 0.024 < mín 0.030 (p̂=0.070, implied=0.046) |
| 2026-05-30T21:32:50Z | will-nashville-sc-win-the-2026-mls-cup | other | 0.0580 | R1_longshot | R1 longshot: precio 0.0580 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.033 < 0.15) |
| 2026-05-30T21:32:50Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 202 d, liq $4029) |
| 2026-05-30T21:32:50Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $262 < absolute min $5000 |
| 2026-05-30T21:32:53Z | will-2026-rank-as-the-sixth-hottest-year-on-record-or-lower | other | 0.0230 | R1_longshot | R1 longshot: precio 0.0230 < 0.10 sin edge fuerte (edge_type=info, edge=0.077 < 0.15) |
| 2026-05-30T21:32:53Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $601) |
| 2026-05-30T21:32:55Z | will-mohammad-bagher-ghalibaf-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0400 | R1_longshot | R1 longshot: precio 0.0400 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.060 < 0.15) |
| 2026-05-30T21:32:58Z | will-china-invade-taiwan-before-2027 | geopolitics | 0.0670 | V3+V6 | V3+V6: V3 Trigger vago: la pregunta 'Will China invade Taiwan by end of 2026?' no especifica un evento verificable co... |
| 2026-05-30T21:32:58Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $427) |
| 2026-05-30T21:32:58Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $495) |
| 2026-05-30T21:32:58Z | will-massoud-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $59447) |
| 2026-05-30T21:33:54Z | insurrection-act-invoked-by-december-31-184-556 | other | 0.2100 | N2 | edge sin soporte de noticias 0.020<0.030 tras shrink ×0.50 |
| 2026-05-30T21:33:54Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $6500) |
| 2026-05-30T21:33:54Z | will-gholam-ali-haddad-adel-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 214 d, liq $61035) |
| 2026-05-30T21:33:57Z | us-x-iran-permanent-peace-deal-by-december-31-2026-961-587-341-574-555 | geopolitics | 0.7100 | V5 Patrón débil: <3 fuentes independientes o sin precedente análogo. V6 Sin catalyst: no hay evento discreto ident... | V5 Patrón débil: <3 fuentes independientes o sin precedente análogo. V6 Sin catalyst: no hay evento discreto ident... |
| 2026-05-30T21:33:59Z | will-north-korea-invade-south-korea-before-2027 | geopolitics | 0.0590 | E2 | edge 0.029 < mín 0.030 (p̂=0.030, implied=0.059) |
| 2026-05-30T21:33:59Z | will-mahmoud-ahmadinejad-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $56983) |
| 2026-05-30T21:33:59Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0600 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T21:34:01Z | will-abbas-araghchi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0310 | E2 | edge 0.011 < mín 0.030 (p̂=0.020, implied=0.031) |
| 2026-05-30T21:34:01Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1211 < absolute min $5000 |
| 2026-05-30T21:34:01Z | will-any-category-5-hurricane-make-landfall-in-the-us-in-before-2027 | weather-natural | 0.1300 | P3_low_absolute_liquidity | liquidity $1623 < absolute min $5000 |
| 2026-05-30T21:34:05Z | will-zelenskyy-and-putin-meet-next-in-qatar-uae | other | 0.0500 | R1_longshot | R1 longshot: precio 0.0500 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.030 < 0.15) |
| 2026-05-30T21:34:05Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $14563) |
| 2026-05-30T21:34:05Z | brex-ipo-before-2027 | crypto-launch | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $4157) |
| 2026-05-30T21:34:05Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $323) |
| 2026-05-30T21:34:05Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $298) |
| 2026-05-30T21:34:05Z | will-there-be-exactly-1-confirmed-vei-4-or-higher-volcanic-eruption-worldwide-in-2026-937 | weather-natural | 0.2700 | P3_low_absolute_liquidity | liquidity $1328 < absolute min $5000 |
| 2026-05-30T21:34:05Z | will-ali-asghar-hejazi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 214 d, liq $55288) |
| 2026-05-30T21:34:05Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $8846) |
| 2026-05-30T21:34:05Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 214 d, liq $9637) |
| 2026-05-30T21:34:05Z | jerome-powell-out-from-fed-board-by-may-30 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 214 d, liq $69179) |
| 2026-05-30T21:34:05Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $1189) |
| 2026-05-30T21:34:05Z | will-there-be-between-11-and-13-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3600 | P3_low_absolute_liquidity | liquidity $1443 < absolute min $5000 |
| 2026-05-30T21:34:05Z | will-there-be-between-17-and-19-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $3497 < absolute min $5000 |
| 2026-05-30T21:34:05Z | will-ahmad-vahidi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 214 d, liq $58889) |
| 2026-05-30T21:34:05Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $10980) |
| 2026-05-30T21:34:05Z | will-meta-acquire-tiktok-745-612-641 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 214 d, liq $11084) |
| 2026-05-30T21:34:05Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $739) |
| 2026-05-30T21:34:05Z | new-covid-variant-of-concern-before-2027 | other | 0.1700 | P3_low_absolute_liquidity | liquidity $2169 < absolute min $5000 |
| 2026-05-30T21:34:05Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $135) |
| 2026-05-30T21:34:05Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $635) |
| 2026-05-30T21:34:07Z | will-2026-be-the-third-hottest-year-on-record | other | 0.0240 | E2 | edge 0.016 < mín 0.030 (p̂=0.040, implied=0.024) |
| 2026-05-30T21:34:09Z | will-zelenskyy-and-putin-meet-next-in-us | other | 0.0430 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda mover el resultado; la p... |
| 2026-05-30T21:34:09Z | discord-ipo-before-2027 | crypto-launch | 0.6000 | P3_low_absolute_liquidity | liquidity $2629 < absolute min $5000 |
| 2026-05-30T21:34:11Z | iran-nuclear-test-before-2027 | geopolitics | 0.0800 | R1_longshot | R1 longshot: precio 0.0800 < 0.10 sin edge fuerte (edge_type=calibration, edge=0.050 < 0.15) |
| 2026-05-30T21:34:11Z | will-larry-page-be-richest-person-on-december-31 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $6594) |
| 2026-05-30T21:34:11Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $5147) |
| 2026-05-30T21:34:13Z | will-mamdani-freeze-nyc-rents-before-2027 | other | 0.3000 | N1 | noticias contradicen tesis: RGB votó rango 0-2% preliminar; votación final en junio 2026 apunta a congelación ante... |
| 2026-05-30T21:34:13Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $6137) |
| 2026-05-30T21:34:15Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 214 d, liq $13110) |
| 2026-05-30T21:34:15Z | will-beyonc-be-the-top-spotify-artist-for-2026 | entertainment | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $15233) |
| 2026-05-30T21:34:15Z | databricks-ipo-before-2027 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $2263 < absolute min $5000 |
| 2026-05-30T21:34:18Z | will-bitcoin-replace-sha-256-before-2027 | other | 0.0520 | V3 (vago), V5 (patrón débil), V6 (sin catalyst) | V3 (vago), V5 (patrón débil), V6 (sin catalyst): V5 Patrón débil: sin precedentes históricos análogos (<3 fuent... |
| 2026-05-30T21:34:18Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $221 < absolute min $5000 |
| 2026-05-30T21:34:18Z | will-maryam-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 214 d, liq $61887) |
| 2026-05-30T21:34:18Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $307) |
| 2026-05-30T21:34:20Z | will-toy-story-5-have-the-best-domestic-opening-weekend-in-2026 | entertainment | 0.0670 | R1_longshot | R1 longshot: precio 0.0670 < 0.10 sin edge fuerte (edge_type=info, edge=0.083 < 0.15) |
| 2026-05-30T21:34:20Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $7121) |
| 2026-05-30T21:34:20Z | will-there-be-between-14-and-16-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2393 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-there-be-exactly-0-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026 | weather-natural | 0.6700 | P3_low_absolute_liquidity | liquidity $1956 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $558) |
| 2026-05-30T21:34:20Z | will-sadegh-larijani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $59567) |
| 2026-05-30T21:34:20Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $201) |
| 2026-05-30T21:34:20Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1567 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4916 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-nir-barkat-be-the-next-prime-minister-of-israel | elections | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $57806) |
| 2026-05-30T21:34:20Z | will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026 | entertainment | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 214 d, liq $18033) |
| 2026-05-30T21:34:20Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $12817) |
| 2026-05-30T21:34:20Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3234 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1390 | P3_low_absolute_liquidity | liquidity $3722 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-mohammad-khatami-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 214 d, liq $54908) |
| 2026-05-30T21:34:20Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4891 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4787 < absolute min $5000 |
| 2026-05-30T21:34:20Z | russia-x-ukraine-ceasefire-agreement-by-december-31-2026 | geopolitics | 0.4300 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 4 veces) |
| 2026-05-30T21:34:20Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4163 < absolute min $5000 |
| 2026-05-30T21:34:20Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $179735) |
| 2026-05-30T21:34:20Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $18561) |
| 2026-05-30T21:34:20Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $235) |
| 2026-05-30T21:34:20Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $4175 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $287 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-yair-golan-be-the-next-prime-minister-of-israel | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $70098) |
| 2026-05-30T21:34:20Z | will-metamask-launch-a-token-by-december-31-2026 | crypto-launch | 0.3300 | P3_low_absolute_liquidity | liquidity $3713 < absolute min $5000 |
| 2026-05-30T21:34:20Z | opensea-fdv-above-5b-one-day-after-launch-919-666-825-992-543-333 | crypto-launch | 0.0500 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.00) |
| 2026-05-30T21:34:20Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3900 | P3_low_absolute_liquidity | liquidity $1089 < absolute min $5000 |
| 2026-05-30T21:34:20Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3268 < absolute min $5000 |
| 2026-05-30T21:34:20Z | puffpaw-fdv-above-400m-one-day-after-launch | crypto-launch | 0.0880 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-ethereum-reach-6000-by-december-31-2026 | market | 0.0600 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.00) |
| 2026-05-30T21:34:20Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2568 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $1450) |
| 2026-05-30T21:34:20Z | will-bitcoin-reach-160000-by-december-31-2026-934-934-164 | market | 0.0700 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.01) |
| 2026-05-30T21:34:20Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1313 < absolute min $5000 |
| 2026-05-30T21:34:20Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $655 < absolute min $5000 |
| 2026-05-30T21:34:20Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | crypto-launch | 0.3130 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | metamask-fdv-above-3b-one-day-after-launch-363-663-664-569-222 | crypto-launch | 0.0910 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-stablecoins-hit-500b-before-2027 | other | 0.1200 | P3_low_absolute_liquidity | liquidity $2137 < absolute min $5000 |
| 2026-05-30T21:34:20Z | okx-ipo-in-2026 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3200 | P3_low_absolute_liquidity | liquidity $586 < absolute min $5000 |
| 2026-05-30T21:34:20Z | altcoin-market-cap-dip-to-150b-before-2027 | other | 0.7200 | P3_low_absolute_liquidity | liquidity $579 < absolute min $5000 |
| 2026-05-30T21:34:20Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0890 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.00) |
| 2026-05-30T21:34:20Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $49307) |
| 2026-05-30T21:34:20Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $279618) |
| 2026-05-30T21:34:20Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $43521) |
| 2026-05-30T21:34:20Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $46139) |
| 2026-05-30T21:34:20Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.2400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | variational-fdv-above-800m-one-day-after-launch-842-835-884 | crypto-launch | 0.3800 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T21:34:20Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1614721) |
| 2026-05-30T21:34:20Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2322482) |
| 2026-05-30T21:34:20Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2431004) |
| 2026-05-30T21:34:20Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1447784) |
| 2026-05-30T21:34:20Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1482373) |
| 2026-05-30T21:34:20Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1056796) |
| 2026-05-30T21:34:20Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $945880) |
| 2026-05-30T21:34:20Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2157375) |
| 2026-05-30T21:34:20Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1163320) |
| 2026-05-30T21:34:20Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $934075) |
| 2026-05-30T21:34:20Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1447661) |
| 2026-05-30T21:34:20Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2752990) |
| 2026-05-30T21:34:20Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $780043) |
| 2026-05-30T21:34:20Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1309029) |
| 2026-05-30T21:34:20Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1720073) |
| 2026-05-30T21:34:20Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1071316) |
| 2026-05-30T21:34:20Z | us-announces-new-iran-agreementceasefire-extension-by-june-7-265-824-655 | geopolitics | 0.4000 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 4 veces) |
| 2026-05-30T21:34:20Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $17260) |
| 2026-05-30T21:34:20Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12012) |
| 2026-05-30T21:34:20Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $4992) |
| 2026-05-30T21:34:20Z | us-announces-new-iran-agreementceasefire-extension-by-may-31-665-831-238 | geopolitics | 0.1200 | M3 | M3 anti-pattern: "ceasefire-rumor" (visto 4 veces) |
| 2026-05-30T21:34:20Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $1770 < absolute min $5000 |
| 2026-05-30T21:34:20Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0400 | P3_low_absolute_liquidity | liquidity $2901 < absolute min $5000 |
| 2026-05-30T21:34:20Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $1211) |
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


## Human notes
_(no se toca por automatización)_

## Sources stats (rolling 30d, computed by brain)

| domain | trades | wins | losses | win_rate | total_pnl | blacklisted |
|---|---|---|---|---|---|---|
| claudemax-websearch | 11 | 2 | 8 | 20.0% | -77.68 | YES |

## Anti-patterns identificados- same-day-price-target · intraday-decay · no-momentum-confirm · short-horizon-chalk · high-entry-low-exit — visto en  (2026-05-26, pnl $-57.55)

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
