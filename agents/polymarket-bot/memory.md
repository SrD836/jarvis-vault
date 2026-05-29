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
| 2026-05-29T16:04:40Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $149151) |
| 2026-05-29T16:04:40Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $7305) |
| 2026-05-29T16:04:40Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $222) |
| 2026-05-29T16:04:40Z | will-gholam-ali-haddad-adel-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $52520) |
| 2026-05-29T16:04:40Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $5511) |
| 2026-05-29T16:04:40Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3641 < absolute min $5000 |
| 2026-05-29T16:04:40Z | ledger-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $1659 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-ahmad-hosseini-khorasani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $53017) |
| 2026-05-29T16:04:40Z | will-mohammad-khatami-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $49255) |
| 2026-05-29T16:04:40Z | will-beyonc-be-the-top-spotify-artist-for-2026 | entertainment | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $15381) |
| 2026-05-29T16:04:40Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $447) |
| 2026-05-29T16:04:40Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $10639) |
| 2026-05-29T16:04:40Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $559) |
| 2026-05-29T16:04:40Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $872) |
| 2026-05-29T16:04:40Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 215 d, liq $5433) |
| 2026-05-29T16:04:40Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $3223 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $10925) |
| 2026-05-29T16:04:40Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $398) |
| 2026-05-29T16:04:40Z | 10pt0-or-above-earthquake-before-2027 | weather-natural | 0.0500 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T16:04:40Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $320 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1240 | P3_low_absolute_liquidity | liquidity $3204 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $173) |
| 2026-05-29T16:04:40Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $3120) |
| 2026-05-29T16:04:40Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $1570 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2859 < absolute min $5000 |
| 2026-05-29T16:04:40Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3310 < absolute min $5000 |
| 2026-05-29T16:04:40Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.05) |
| 2026-05-29T16:04:40Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $817 < absolute min $5000 |
| 2026-05-29T16:04:40Z | puffpaw-fdv-above-100m-one-day-after-launch | crypto-launch | 0.5400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $1019 < absolute min $5000 |
| 2026-05-29T16:04:40Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.05) |
| 2026-05-29T16:04:40Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4100 | P3_low_absolute_liquidity | liquidity $1058 < absolute min $5000 |
| 2026-05-29T16:04:40Z | okx-ipo-in-2026 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.1810 | P3_low_absolute_liquidity | liquidity $3093 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1397 < absolute min $5000 |
| 2026-05-29T16:04:40Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-usdt-market-cap-hit-200b-before-2027 | other | 0.9780 | P3_low_absolute_liquidity | liquidity $953 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $454 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1601 < absolute min $5000 |
| 2026-05-29T16:04:40Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12377) |
| 2026-05-29T16:04:40Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52747) |
| 2026-05-29T16:04:40Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45535) |
| 2026-05-29T16:04:40Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 240 d, liq $52722) |
| 2026-05-29T16:04:40Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $18664) |
| 2026-05-29T16:04:40Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $114090) |
| 2026-05-29T16:04:40Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $252880) |
| 2026-05-29T16:04:40Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7700 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $4660 < absolute min $5000 |
| 2026-05-29T16:04:40Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $24053) |
| 2026-05-29T16:04:40Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45855) |
| 2026-05-29T16:04:40Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 580 d, liq $12757) |
| 2026-05-29T16:04:40Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T16:04:40Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1913255) |
| 2026-05-29T16:04:40Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $281651) |
| 2026-05-29T16:04:40Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1761891) |
| 2026-05-29T16:04:40Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2361524) |
| 2026-05-29T16:04:40Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $762196) |
| 2026-05-29T16:04:40Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1218113) |
| 2026-05-29T16:04:40Z | will-ro-khanna-win-the-2028-democratic-presidential-nomination | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $246980) |
| 2026-05-29T16:04:40Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1934246) |
| 2026-05-29T16:04:40Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1928405) |
| 2026-05-29T16:04:40Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1074431) |
| 2026-05-29T16:04:40Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2142944) |
| 2026-05-29T16:04:40Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2260 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2600 | P3_low_absolute_liquidity | liquidity $486 < absolute min $5000 |
| 2026-05-29T16:04:40Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $12065) |
| 2026-05-29T16:04:40Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22029) |
| 2026-05-29T16:04:40Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $17033) |
| 2026-05-29T16:04:40Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T16:30:03Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -149 d, liq $36800) |
| 2026-05-29T16:30:03Z | will-us-unemployment-reach-at-least-5pt5-in-2026 | other | 0.1960 | P3_low_absolute_liquidity | liquidity $791 < absolute min $5000 |
| 2026-05-29T16:30:03Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 118 días) |
| 2026-05-29T16:30:03Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $766688) |
| 2026-05-29T16:30:03Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -68 d, liq $32044) |
| 2026-05-29T16:30:03Z | will-jernej-vrtovec-be-the-next-prime-minister-of-slovenia | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -68 d, liq $19571) |
| 2026-05-29T16:30:03Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $35565) |
| 2026-05-29T16:30:03Z | will-matej-tonin-be-the-next-prime-minister-of-slovenia | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -68 d, liq $7699) |
| 2026-05-29T16:30:03Z | will-robert-golob-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $14203) |
| 2026-05-29T16:30:03Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -66 d, liq $27604) |
| 2026-05-29T16:30:03Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -66 d, liq $19846) |
| 2026-05-29T16:30:03Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -66 d, liq $28449) |
| 2026-05-29T16:30:03Z | will-lars-boje-mathiesen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon -66 d, liq $23542) |
| 2026-05-29T16:30:03Z | will-the-us-invade-venezuela-by-december-31-2026 | geopolitics | 0.1100 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T16:30:03Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $199099) |
| 2026-05-29T16:30:03Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -46 d, liq $344882) |
| 2026-05-29T16:30:03Z | will-a-claude-mythos-model-be-released-by-june-30-2026-684-583 | other | 0.3800 | P2 | mercado ya expiró (endDate=2026-04-30T00:00:00Z, hace 29 días) |
| 2026-05-29T16:30:03Z | us-x-iran-diplomatic-meeting-by-june-15-2026-671 | geopolitics | 0.5400 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T16:30:03Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.8000 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T16:30:03Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0690 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T16:30:03Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -12 d, liq $5905) |
| 2026-05-29T16:30:03Z | will-brad-raffensperger-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -10 d, liq $43083) |
| 2026-05-29T16:30:03Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -10 d, liq $6291) |
| 2026-05-29T16:30:03Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -2 d, liq $427718) |
| 2026-05-29T16:30:03Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-28-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $230430) |
| 2026-05-29T16:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-29 | geopolitics | 0.3000 | P9 | P9: geopolitics pump cluster (price 0.30, 0d) |
| 2026-05-29T16:30:03Z | highest-temperature-in-shanghai-on-may-29-2026-26c | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $32615) |
| 2026-05-29T16:30:03Z | elon-musk-of-tweets-may-22-may-29-200-219 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $157665) |
| 2026-05-29T16:30:03Z | elon-musk-of-tweets-may-22-may-29-220-239 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $110642) |
| 2026-05-29T16:30:03Z | bitcoin-above-74k-on-may-29-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $321019) |
| 2026-05-29T16:30:03Z | bitcoin-above-76k-on-may-29-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $277757) |
| 2026-05-29T16:30:03Z | bitcoin-above-78k-on-may-29-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $221249) |
| 2026-05-29T16:30:03Z | elon-musk-of-tweets-may-22-may-29-160-179 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $437424) |
| 2026-05-29T16:30:03Z | elon-musk-of-tweets-may-22-may-29-240-259 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $89577) |
| 2026-05-29T16:30:09Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $16724) |
| 2026-05-29T16:30:15Z | ucl-psg-ars-2026-05-30-psg | other | 0.4100 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:24Z | will-psg-win-the-202526-champions-league | sports-season | 0.5800 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6: Sin catalyst concreto. No hay un... |
| 2026-05-29T16:30:24Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $167195) |
| 2026-05-29T16:30:24Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $187662) |
| 2026-05-29T16:30:27Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0540 | P9 | P9: geopolitics pump cluster (price 0.05, 1d) |
| 2026-05-29T16:30:30Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0380 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-29T16:30:30Z | will-donald-trump-publicly-insult-someone-on-may-28-2026 | other | 0.9920 | P0_ceiling | price ceiling: 0.9920 > 0.980 |
| 2026-05-29T16:30:30Z | will-declan-rice-record-the-most-yellow-cards-in-2025-26-uefa-champions-league | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 1 d, liq $5895) |
| 2026-05-29T16:30:30Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.1580 | P9 | P9: geopolitics pump cluster (price 0.16, 1d) |
| 2026-05-29T16:30:30Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $200957) |
| 2026-05-29T16:30:30Z | trump-out-as-president-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $871696) |
| 2026-05-29T16:30:36Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0300 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0860 | P9 | P9: geopolitics pump cluster (price 0.09, 1d) |
| 2026-05-29T16:30:36Z | will-40-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.1700 | P9 | P9: geopolitics pump cluster (price 0.17, 1d) |
| 2026-05-29T16:30:36Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 1 d, liq $259736) |
| 2026-05-29T16:30:36Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $3266887) |
| 2026-05-29T16:30:36Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0570 | P9 | P9: geopolitics pump cluster (price 0.06, 1d) |
| 2026-05-29T16:30:36Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $221662) |
| 2026-05-29T16:30:36Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.1400 | P9 | P9: geopolitics pump cluster (price 0.14, 1d) |
| 2026-05-29T16:30:36Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $198504) |
| 2026-05-29T16:30:36Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.3400 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0470 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0230 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $350367) |
| 2026-05-29T16:30:36Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.1770 | P9 | P9: geopolitics pump cluster (price 0.18, 1d) |
| 2026-05-29T16:30:36Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.2070 | P9 | P9: geopolitics pump cluster (price 0.21, 1d) |
| 2026-05-29T16:30:36Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 1 d, liq $258247) |
| 2026-05-29T16:30:36Z | will-trump-agree-to-iranian-oil-sanction-relief-by-may-31 | executive-action | 0.2300 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | will-60-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0650 | P9 | P9: geopolitics pump cluster (price 0.07, 1d) |
| 2026-05-29T16:30:36Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $68246) |
| 2026-05-29T16:30:36Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $110209) |
| 2026-05-29T16:30:36Z | will-paloma-valencia-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $46389) |
| 2026-05-29T16:30:36Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $100519) |
| 2026-05-29T16:30:36Z | will-abelardo-de-la-espriella-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.2530 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $124571) |
| 2026-05-29T16:30:36Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $83000) |
| 2026-05-29T16:30:36Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $66111) |
| 2026-05-29T16:30:36Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $117260) |
| 2026-05-29T16:30:36Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116965) |
| 2026-05-29T16:30:36Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $90310) |
| 2026-05-29T16:30:36Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $114893) |
| 2026-05-29T16:30:36Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $124444) |
| 2026-05-29T16:30:36Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $66343) |
| 2026-05-29T16:30:36Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $83969) |
| 2026-05-29T16:30:36Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $92747) |
| 2026-05-29T16:30:36Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 2 d, liq $310596) |
| 2026-05-29T16:30:36Z | will-wti-reach-110-in-may-2026-116-472 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $37584) |
| 2026-05-29T16:30:36Z | will-wti-reach-150-in-may-2026-196-364 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $112951) |
| 2026-05-29T16:30:36Z | will-wti-reach-140-in-may-2026-916 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $147771) |
| 2026-05-29T16:30:36Z | will-wti-crude-oil-wti-hit-high-115-in-may-221 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $58745) |
| 2026-05-29T16:30:36Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 2 d, liq $107309) |
| 2026-05-29T16:30:36Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.3100 | P6 | P6 market: CL=F spot $87.70 already > target $85.00 but yes=0.31 (confused book) |
| 2026-05-29T16:30:36Z | will-wti-reach-120-in-may-2026-592-217 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $119038) |
| 2026-05-29T16:30:36Z | will-wti-reach-130-in-may-2026-733 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $147005) |
| 2026-05-29T16:30:36Z | will-wti-crude-oil-wti-hit-high-105-in-may | market | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 2 d, liq $78908) |
| 2026-05-29T16:30:36Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0610 | P6 | P6 market: CL=F spot $87.70 already > target $80.00 but yes=0.06 (confused book) |
| 2026-05-29T16:30:36Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $163722) |
| 2026-05-29T16:30:36Z | will-bitcoin-reach-150k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $1134222) |
| 2026-05-29T16:30:36Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $49344) |
| 2026-05-29T16:30:36Z | will-bitcoin-dip-to-72500-in-may-2026-from-may-27 | market | 0.4200 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.42) |
| 2026-05-29T16:30:36Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $10337) |
| 2026-05-29T16:30:36Z | will-ethereum-reach-3600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $66656) |
| 2026-05-29T16:30:36Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $55224) |
| 2026-05-29T16:30:36Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 2 d, liq $35264) |
| 2026-05-29T16:30:36Z | will-bitcoin-dip-to-35k-in-may-2026-217-769-834 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $131687) |
| 2026-05-29T16:30:36Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $215968) |
| 2026-05-29T16:30:36Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $201725) |
| 2026-05-29T16:30:36Z | will-ethereum-dip-to-400-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $56543) |
| 2026-05-29T16:30:36Z | will-ethereum-reach-4000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $65904) |
| 2026-05-29T16:30:36Z | will-solana-dip-to-50-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $47746) |
| 2026-05-29T16:30:36Z | will-bitcoin-dip-to-50k-in-may-2026-896 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $175155) |
| 2026-05-29T16:30:36Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $242666) |
| 2026-05-29T16:30:36Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0590 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.42) |
| 2026-05-29T16:30:36Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2000 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:36Z | will-robert-lebovics-be-the-republican-nominee-for-senate-in-new-jersey-744 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 3 d, liq $6743) |
| 2026-05-29T16:30:36Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $36630) |
| 2026-05-29T16:30:36Z | will-rob-sand-win-the-2026-iowa-governor-democratic-primary-election | elections | 0.9960 | P0_ceiling | price ceiling: 0.9960 > 0.980 |
| 2026-05-29T16:30:36Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $17733) |
| 2026-05-29T16:30:36Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 3 d, liq $65359) |
| 2026-05-29T16:30:37Z | elon-musk-of-tweets-may-26-june-2-60-79 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $78678) |
| 2026-05-29T16:30:37Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $93900) |
| 2026-05-29T16:30:37Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $45273) |
| 2026-05-29T16:30:37Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $40702) |
| 2026-05-29T16:30:37Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30216) |
| 2026-05-29T16:30:37Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30121) |
| 2026-05-29T16:30:37Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $94779) |
| 2026-05-29T16:30:37Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $43321) |
| 2026-05-29T16:30:37Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $23365) |
| 2026-05-29T16:30:37Z | will-the-rebuilding-korea-party-rkp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $82812) |
| 2026-05-29T16:30:37Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $21910) |
| 2026-05-29T16:30:37Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $35120) |
| 2026-05-29T16:30:37Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30733) |
| 2026-05-29T16:30:37Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.6400 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-oh-se-hoon-win-the-2026-seoul-mayoral-election | elections | 0.2800 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $244664) |
| 2026-05-29T16:30:37Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $236748) |
| 2026-05-29T16:30:37Z | will-tommy-paul-win-the-2026-mens-french-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 8 d, liq $55105) |
| 2026-05-29T16:30:37Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 8 d, liq $30683) |
| 2026-05-29T16:30:37Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $239449) |
| 2026-05-29T16:30:37Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 8 d, liq $68128) |
| 2026-05-29T16:30:37Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152018) |
| 2026-05-29T16:30:37Z | will-keiko-fujimori-win-the-2026-peruvian-presidential-election | elections | 0.7900 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0530 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $219340) |
| 2026-05-29T16:30:37Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152077) |
| 2026-05-29T16:30:37Z | will-casper-ruud-win-the-2026-mens-french-open | other | 0.0880 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $204192) |
| 2026-05-29T16:30:37Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $193646) |
| 2026-05-29T16:30:37Z | will-alexander-zverev-win-the-2026-mens-french-open | other | 0.3060 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-roberto-snchez-palomino-win-the-2026-peruvian-presidential-election | elections | 0.2130 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-novak-djokovic-win-the-2026-mens-french-open | other | 0.2410 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0400 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $109941) |
| 2026-05-29T16:30:37Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0600 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 8 d, liq $35766) |
| 2026-05-29T16:30:37Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-7-2026 | geopolitics | 0.8200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0360 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $201702) |
| 2026-05-29T16:30:37Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $232830) |
| 2026-05-29T16:30:37Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $120843) |
| 2026-05-29T16:30:37Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 8 d, liq $29700) |
| 2026-05-29T16:30:37Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $205541) |
| 2026-05-29T16:30:37Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-alex-michelsen-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $100809) |
| 2026-05-29T16:30:37Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1010 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-karen-khachanov-win-the-2026-mens-french-open | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 8 d, liq $56329) |
| 2026-05-29T16:30:37Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0220 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $181697) |
| 2026-05-29T16:30:37Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $189994) |
| 2026-05-29T16:30:37Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.2200 | P9 | P9: geopolitics pump cluster (price 0.22, 8d) |
| 2026-05-29T16:30:37Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $214062) |
| 2026-05-29T16:30:37Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.2600 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | strait-of-hormuz-traffic-returns-to-normal-by-june-15 | geopolitics | 0.1200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-the-oklahoma-city-thunder-win-the-nba-western-conference-finals | sports-season | 0.5900 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-the-san-antonio-spurs-win-the-nba-western-conference-finals | sports-season | 0.4260 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 18 d, liq $728164) |
| 2026-05-29T16:30:37Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 18 d, liq $203505) |
| 2026-05-29T16:30:37Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 18 d, liq $2557292) |
| 2026-05-29T16:30:37Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9780 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 18 d, liq $639083) |
| 2026-05-29T16:30:37Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6400 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-claudia-lpez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $106960) |
| 2026-05-29T16:30:37Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 22 d, liq $44724) |
| 2026-05-29T16:30:37Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $153668) |
| 2026-05-29T16:30:37Z | will-ivan-cepeda-castro-win-the-2026-colombian-presidential-election | elections | 0.3400 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $138302) |
| 2026-05-29T16:30:37Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $157405) |
| 2026-05-29T16:30:37Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $149833) |
| 2026-05-29T16:30:37Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0380 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $110923) |
| 2026-05-29T16:30:37Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $169000) |
| 2026-05-29T16:30:37Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $188037) |
| 2026-05-29T16:30:37Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $1967) |
| 2026-05-29T16:30:37Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 25 d, liq $1374) |
| 2026-05-29T16:30:37Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $1539) |
| 2026-05-29T16:30:37Z | will-viggo-bjrck-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 28 d, liq $1024) |
| 2026-05-29T16:30:37Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $1275) |
| 2026-05-29T16:30:37Z | will-tynan-lawrence-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $2766) |
| 2026-05-29T16:30:37Z | us-iran-nuclear-deal-by-june-30 | geopolitics | 0.5300 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | iran-closes-its-airspace-by-june-30-432-786-462-866 | geopolitics | 0.2490 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 31 d, liq $101900) |
| 2026-05-29T16:30:37Z | forsen-to-get-signed-by-an-lck-prog-org-by-jun-30 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 31 d, liq $48608) |
| 2026-05-29T16:30:37Z | hamad-bin-isa-al-khalifa-out-as-bahrain-king | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 31 d, liq $13741) |
| 2026-05-29T16:30:37Z | nato-x-russia-military-clash-by-june-30-2026 | geopolitics | 0.0560 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0370 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-june | geopolitics | 0.4300 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 31 d, liq $16605) |
| 2026-05-29T16:30:37Z | aaron-taylor-johnson-announced-as-next-james-bond | other | 0.0400 | P3_low_absolute_liquidity | liquidity $4833 < absolute min $5000 |
| 2026-05-29T16:30:37Z | will-donald-trump-be-confirmed-to-have-visited-epsteins-island | other | 0.0200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:30:37Z | pierce-brosnan-announced-as-next-james-bond-557 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $7038) |
| 2026-05-29T16:30:37Z | gustavo-petro-out-as-leader-of-colombia-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $15547) |
| 2026-05-29T16:30:37Z | paul-mescal-announced-as-next-james-bond | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 31 d, liq $5021) |
| 2026-05-29T16:30:45Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0230 | V3 | V3: V3 Trigger vago: 'the Iranian regime fall' no es un evento verificable con criterios objetivos y fecha concreta. ... |
| 2026-05-29T16:30:45Z | jack-lowdon-announced-as-next-james-bond | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 31 d, liq $4560) |
| 2026-05-29T16:30:45Z | will-deepseek-have-the-top-ai-model-at-the-end-of-june-2026-514 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 31 d, liq $8711) |
| 2026-05-29T16:30:45Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $345178) |
| 2026-05-29T16:30:45Z | james-norton-announced-as-next-james-bond-575 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $6865) |
| 2026-05-29T16:30:50Z | henry-cavill-announced-as-next-james-bond-857-642 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 31 d, liq $4861) |
| 2026-05-29T16:30:50Z | will-trump-pardon-tiger-woods-by-june-30 | executive-action | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 31 d, liq $25871) |
| 2026-05-29T16:30:50Z | will-us-withdraw-from-nato-by-june-30 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $27755) |
| 2026-05-29T16:30:53Z | will-russia-enter-kramatorsk-by-june-30-821-192 | geopolitics | 0.0280 | E2 | edge 0.022 < mín 0.030 (p̂=0.050, implied=0.028) |
| 2026-05-29T16:30:53Z | ukraine-peace-referendum-scheduled-by-june-30 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $18350) |
| 2026-05-29T16:30:53Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $329302) |
| 2026-05-29T16:30:53Z | will-alibaba-have-the-top-ai-model-at-the-end-of-june-2026-248 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 31 d, liq $8219) |
| 2026-05-29T16:30:56Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $11442) |
| 2026-05-29T16:30:56Z | will-amazon-be-the-largest-company-in-the-world-by-market-cap-on-june-30-123 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $342697) |
| 2026-05-29T16:30:59Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V3 Trigger vago: requiere fecha concreta o evento verificable. V6 Sin catalyst: no hay evento discreto identificable ... | V3 Trigger vago: requiere fecha concreta o evento verificable. V6 Sin catalyst: no hay evento discreto identificable ... |
| 2026-05-29T16:31:01Z | harris-dickinson-announced-as-next-james-bond-784 | other | 0.0330 | P3_low_absolute_liquidity | liquidity $4125 < absolute min $5000 |
| 2026-05-29T16:31:01Z | will-the-montreal-canadiens-win-the-2026-nhl-stanley-cup | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $231598) |
| 2026-05-29T16:31:04Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 32 d, liq $30905) |
| 2026-05-29T16:31:04Z | will-silver-si-hit-low-45-by-end-of-june-972-272 | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 32 d, liq $13707) |
| 2026-05-29T16:31:08Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0460 | E2 | edge 0.016 < mín 0.030 (p̂=0.030, implied=0.046) |
| 2026-05-29T16:31:11Z | will-crude-oil-cl-hit-high-150-by-end-of-june-788-691 | market | 0.0240 | E2 | edge 0.014 < mín 0.030 (p̂=0.010, implied=0.024) |
| 2026-05-29T16:31:28Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P3_low_absolute_liquidity | liquidity $2437 < absolute min $5000 |
| 2026-05-29T16:31:28Z | will-iga-witek-be-the-2026-womens-wimbledon-winner | other | 0.1890 | P3_low_absolute_liquidity | liquidity $2086 < absolute min $5000 |
| 2026-05-29T16:31:28Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 43 d, liq $2492) |
| 2026-05-29T16:31:28Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0430 | P3_low_absolute_liquidity | liquidity $3103 < absolute min $5000 |
| 2026-05-29T16:31:28Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 43 d, liq $3141) |
| 2026-05-29T16:31:28Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $599 < absolute min $5000 |
| 2026-05-29T16:31:31Z | will-haiti-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10535000) |
| 2026-05-29T16:31:31Z | will-cape-verde-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10380348) |
| 2026-05-29T16:31:31Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $7013480) |
| 2026-05-29T16:31:35Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E2 | edge 0.004 < mín 0.030 (p̂=0.090, implied=0.086) |
| 2026-05-29T16:31:39Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.008 < mín 0.030 (p̂=0.060, implied=0.052) |
| 2026-05-29T16:31:39Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $8093350) |
| 2026-05-29T16:31:39Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $5554369) |
| 2026-05-29T16:31:39Z | will-curaao-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10022613) |
| 2026-05-29T16:31:39Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $5809705) |
| 2026-05-29T16:31:42Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | E2 | edge 0.026 < mín 0.030 (p̂=0.120, implied=0.094) |
| 2026-05-29T16:31:45Z | will-new-zealand-win-the-2026-fifa-world-cup-635 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9888757) |
| 2026-05-29T16:31:45Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $4033414) |
| 2026-05-29T16:31:45Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9494043) |
| 2026-05-29T16:31:45Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $6324204) |
| 2026-05-29T16:31:45Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $4660522) |
| 2026-05-29T16:31:45Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $1137413) |
| 2026-05-29T16:31:45Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4569814) |
| 2026-05-29T16:31:45Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 51 d, liq $3513285) |
| 2026-05-29T16:31:48Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4525474) |
| 2026-05-29T16:31:48Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9224221) |
| 2026-05-29T16:31:48Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $9666510) |
| 2026-05-29T16:31:48Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 51 d, liq $5360762) |
| 2026-05-29T16:31:50Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0390 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T16:31:53Z | will-jordan-win-the-2026-fifa-world-cup-233 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10313162) |
| 2026-05-29T16:31:53Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8700841) |
| 2026-05-29T16:31:53Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8090565) |
| 2026-05-29T16:31:53Z | will-qatar-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10402597) |
| 2026-05-29T16:31:53Z | will-team-z-win-the-2026-fifa-world-cup-316 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10726023) |
| 2026-05-29T16:31:53Z | will-iraq-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10684942) |
| 2026-05-29T16:31:53Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $3023207) |
| 2026-05-29T16:31:53Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 51 d, liq $9260371) |
| 2026-05-29T16:31:53Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $7317115) |
| 2026-05-29T16:31:57Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $4759421) |
| 2026-05-29T16:31:57Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $8263155) |
| 2026-05-29T16:31:57Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9362188) |
| 2026-05-29T16:32:00Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0290 | E2 | edge 0.019 < mín 0.030 (p̂=0.010, implied=0.029) |
| 2026-05-29T16:32:00Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $2686835) |
| 2026-05-29T16:32:00Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 51 d, liq $401581) |
| 2026-05-29T16:32:00Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $10682150) |
| 2026-05-29T16:32:00Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8234831) |
| 2026-05-29T16:32:00Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $2657444) |
| 2026-05-29T16:32:00Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 51 d, liq $2362488) |
| 2026-05-29T16:32:00Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $10512944) |
| 2026-05-29T16:32:00Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9803549) |
| 2026-05-29T16:32:00Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $15815) |
| 2026-05-29T16:32:02Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $33068) |
| 2026-05-29T16:32:05Z | us-x-iran-permanent-peace-deal-by-july-31-2026-831-252 | geopolitics | 0.5700 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: 'peace deal' es un concepto ambiguo s... |
| 2026-05-29T16:32:05Z | will-paul-reevs-be-the-republican-nominee-for-az-01 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 66 d, liq $5773) |
| 2026-05-29T16:32:05Z | will-jay-collins-be-the-republican-nominee-for-florida-governor | elections | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 80 d, liq $37349) |
| 2026-05-29T16:32:05Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0400 | P3_low_absolute_liquidity | liquidity $808 < absolute min $5000 |
| 2026-05-29T16:32:05Z | will-david-njoku-play-for-tennessee-titans-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 94 d, liq $2102) |
| 2026-05-29T16:32:09Z | tush-push-banned-for-2026-nfl-season | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 103 d, liq $1875) |
| 2026-05-29T16:32:09Z | will-tereza-valentova-win-the-2026-womens-us-open | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 106 d, liq $2707) |
| 2026-05-29T16:32:09Z | will-anastasia-potapova-win-the-2026-womens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2425) |
| 2026-05-29T16:32:09Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $1796) |
| 2026-05-29T16:32:09Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 106 d, liq $2729) |
| 2026-05-29T16:32:12Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 106 d, liq $3503) |
| 2026-05-29T16:32:12Z | will-barbora-krejcikova-win-the-2026-womens-us-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 106 d, liq $2580) |
| 2026-05-29T16:32:12Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 106 d, liq $2591) |
| 2026-05-29T16:32:12Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2948) |
| 2026-05-29T16:32:12Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $3644) |
| 2026-05-29T16:32:12Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 121 d, liq $1362) |
| 2026-05-29T16:32:12Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $909 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $440 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 121 d, liq $1917) |
| 2026-05-29T16:32:12Z | will-civic-platform-gp-win-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 123 d, liq $22962) |
| 2026-05-29T16:32:12Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5700 | P8 | P8: election 123d out, price 0.57 en banda ruido [0.30, 0.70] |
| 2026-05-29T16:32:12Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $801 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-michelle-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 127 d, liq $215150) |
| 2026-05-29T16:32:12Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 155 d, liq $30641) |
| 2026-05-29T16:32:12Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 155 d, liq $31936) |
| 2026-05-29T16:32:12Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 155 d, liq $28207) |
| 2026-05-29T16:32:12Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 155 d, liq $28556) |
| 2026-05-29T16:32:12Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $7118) |
| 2026-05-29T16:32:12Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $160706) |
| 2026-05-29T16:32:12Z | will-corey-seager-win-the-2026-american-league-hank-aaron-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 167 d, liq $2651) |
| 2026-05-29T16:32:12Z | will-sergio-prez-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 190 d, liq $751985) |
| 2026-05-29T16:32:12Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 193 d, liq $12491) |
| 2026-05-29T16:32:12Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $9320) |
| 2026-05-29T16:32:12Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $12590) |
| 2026-05-29T16:32:12Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 203 d, liq $5761) |
| 2026-05-29T16:32:12Z | will-payton-tolle-win-the-2026-al-rookie-of-the-year-award | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 203 d, liq $7824) |
| 2026-05-29T16:32:12Z | will-mahmoud-ahmadinejad-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $49556) |
| 2026-05-29T16:32:12Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $320 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $118 < absolute min $5000 |
| 2026-05-29T16:32:12Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0640 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T16:32:12Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4641 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $7396) |
| 2026-05-29T16:32:12Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $3965 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 215 d, liq $6062) |
| 2026-05-29T16:32:12Z | will-muhammad-mirbaqiri-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $49494) |
| 2026-05-29T16:32:12Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $8132) |
| 2026-05-29T16:32:12Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $8982) |
| 2026-05-29T16:32:12Z | ramp-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $2577 < absolute min $5000 |
| 2026-05-29T16:32:12Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1315 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-sadegh-larijani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $47024) |
| 2026-05-29T16:32:12Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $156598) |
| 2026-05-29T16:32:12Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $5700) |
| 2026-05-29T16:32:12Z | will-massoud-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $51562) |
| 2026-05-29T16:32:12Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $3234 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $421) |
| 2026-05-29T16:32:12Z | anduril-ipo-before-2027 | crypto-launch | 0.1200 | P3_low_absolute_liquidity | liquidity $3531 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $5917) |
| 2026-05-29T16:32:12Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $86) |
| 2026-05-29T16:32:12Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4163 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $95986) |
| 2026-05-29T16:32:12Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $1517 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $14554) |
| 2026-05-29T16:32:12Z | natural-disaster-in-2026 | other | 0.2800 | P3_low_absolute_liquidity | liquidity $2301 < absolute min $5000 |
| 2026-05-29T16:32:12Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $18247) |
| 2026-05-29T16:32:12Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | sports-season | 0.0970 | P3_low_absolute_liquidity | liquidity $873 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $1076) |
| 2026-05-29T16:32:12Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6850) |
| 2026-05-29T16:32:12Z | will-larry-page-be-richest-person-on-december-31 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $6798) |
| 2026-05-29T16:32:12Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $61735) |
| 2026-05-29T16:32:12Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $223) |
| 2026-05-29T16:32:12Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3642 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $545) |
| 2026-05-29T16:32:12Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $225 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $3895 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $280) |
| 2026-05-29T16:32:12Z | ledger-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1484 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-mohammad-khatami-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $47909) |
| 2026-05-29T16:32:12Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $539) |
| 2026-05-29T16:32:12Z | will-gholam-ali-haddad-adel-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $51207) |
| 2026-05-29T16:32:12Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1220 | P3_low_absolute_liquidity | liquidity $3273 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $399) |
| 2026-05-29T16:32:12Z | ripple-labs-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $2071 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $11594) |
| 2026-05-29T16:32:12Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2532 < absolute min $5000 |
| 2026-05-29T16:32:12Z | 10pt0-or-above-earthquake-before-2027 | weather-natural | 0.0500 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T16:32:12Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $11140) |
| 2026-05-29T16:32:12Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $137123) |
| 2026-05-29T16:32:12Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $174) |
| 2026-05-29T16:32:12Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $12264) |
| 2026-05-29T16:32:12Z | will-ahmad-hosseini-khorasani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $52043) |
| 2026-05-29T16:32:12Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $874) |
| 2026-05-29T16:32:12Z | will-beyonc-be-the-top-spotify-artist-for-2026 | entertainment | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $15520) |
| 2026-05-29T16:32:12Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $2957) |
| 2026-05-29T16:32:12Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $1565 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1608 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $454 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.07) |
| 2026-05-29T16:32:12Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2856 < absolute min $5000 |
| 2026-05-29T16:32:12Z | okx-ipo-in-2026 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3800 | P3_low_absolute_liquidity | liquidity $1082 < absolute min $5000 |
| 2026-05-29T16:32:12Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.1760 | P3_low_absolute_liquidity | liquidity $2335 < absolute min $5000 |
| 2026-05-29T16:32:12Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $1013 < absolute min $5000 |
| 2026-05-29T16:32:12Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12378) |
| 2026-05-29T16:32:12Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $815 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1432 < absolute min $5000 |
| 2026-05-29T16:32:12Z | bitcoin-all-time-high-by-december-31-2026 | market | 0.1500 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.42) |
| 2026-05-29T16:32:12Z | puffpaw-fdv-above-100m-one-day-after-launch | crypto-launch | 0.5400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | will-usdt-market-cap-hit-200b-before-2027 | other | 0.9780 | P3_low_absolute_liquidity | liquidity $964 < absolute min $5000 |
| 2026-05-29T16:32:12Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3462 < absolute min $5000 |
| 2026-05-29T16:32:12Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.07) |
| 2026-05-29T16:32:12Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $18865) |
| 2026-05-29T16:32:12Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 240 d, liq $52792) |
| 2026-05-29T16:32:12Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45547) |
| 2026-05-29T16:32:12Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52852) |
| 2026-05-29T16:32:12Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $114090) |
| 2026-05-29T16:32:12Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $255103) |
| 2026-05-29T16:32:12Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7700 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7400 | P3_low_absolute_liquidity | liquidity $4202 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 580 d, liq $12859) |
| 2026-05-29T16:32:12Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45536) |
| 2026-05-29T16:32:12Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $23519) |
| 2026-05-29T16:32:12Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T16:32:12Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $769566) |
| 2026-05-29T16:32:12Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $289722) |
| 2026-05-29T16:32:12Z | will-ro-khanna-win-the-2028-democratic-presidential-nomination | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $253757) |
| 2026-05-29T16:32:12Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1766422) |
| 2026-05-29T16:32:12Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1917948) |
| 2026-05-29T16:32:12Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2147378) |
| 2026-05-29T16:32:12Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1081887) |
| 2026-05-29T16:32:12Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2368334) |
| 2026-05-29T16:32:12Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1939760) |
| 2026-05-29T16:32:12Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1933456) |
| 2026-05-29T16:32:12Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22027) |
| 2026-05-29T16:32:12Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $11181) |
| 2026-05-29T16:32:12Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $544 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $16823) |
| 2026-05-29T16:32:12Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2259 < absolute min $5000 |
| 2026-05-29T16:32:12Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
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
