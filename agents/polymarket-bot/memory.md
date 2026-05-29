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
| 2026-05-29T15:03:27Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-massoud-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $48184) |
| 2026-05-29T15:03:27Z | will-muhammad-mirbaqiri-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $45514) |
| 2026-05-29T15:03:27Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $560) |
| 2026-05-29T15:03:27Z | will-mahmoud-ahmadinejad-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $51422) |
| 2026-05-29T15:03:27Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $786 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-sadegh-larijani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $42715) |
| 2026-05-29T15:03:27Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $1192 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $7696) |
| 2026-05-29T15:03:27Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0650 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T15:03:27Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $65413) |
| 2026-05-29T15:03:27Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4895 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | sports-season | 0.0880 | P3_low_absolute_liquidity | liquidity $807 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6440) |
| 2026-05-29T15:03:27Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4209 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $233) |
| 2026-05-29T15:03:27Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $323) |
| 2026-05-29T15:03:27Z | will-maryam-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $50126) |
| 2026-05-29T15:03:27Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1220 | P3_low_absolute_liquidity | liquidity $2916 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $11082) |
| 2026-05-29T15:03:27Z | 10pt0-or-above-earthquake-before-2027 | weather-natural | 0.0500 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T15:03:27Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $5244) |
| 2026-05-29T15:03:27Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $1146) |
| 2026-05-29T15:03:27Z | natural-disaster-in-2026 | other | 0.2800 | P3_low_absolute_liquidity | liquidity $1482 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $3108 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $176) |
| 2026-05-29T15:03:27Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $462) |
| 2026-05-29T15:03:27Z | anduril-ipo-before-2027 | crypto-launch | 0.1200 | P3_low_absolute_liquidity | liquidity $3122 < absolute min $5000 |
| 2026-05-29T15:03:27Z | sbf-released-from-custody-in-2026 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $4566 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $5556) |
| 2026-05-29T15:03:27Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $9147) |
| 2026-05-29T15:03:27Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3154 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2312 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $100091) |
| 2026-05-29T15:03:27Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $11131) |
| 2026-05-29T15:03:27Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $142744) |
| 2026-05-29T15:03:27Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $15430) |
| 2026-05-29T15:03:27Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $561) |
| 2026-05-29T15:03:27Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $2078) |
| 2026-05-29T15:03:27Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $1440 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-usdt-market-cap-hit-200b-before-2027 | other | 0.9780 | P3_low_absolute_liquidity | liquidity $919 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $472 < absolute min $5000 |
| 2026-05-29T15:03:27Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1634 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.07) |
| 2026-05-29T15:03:27Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6900 | P3_low_absolute_liquidity | liquidity $2877 < absolute min $5000 |
| 2026-05-29T15:03:27Z | okx-ipo-in-2026 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $806 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-microstrategy-announce-holding-1m-btc-by-december-31-2026-bv81 | market | 0.4700 | P11 | tradingview BTCUSDT sentiment=bearish (conf 0.80) contradicts bull thesis (price_yes=0.94) |
| 2026-05-29T15:03:27Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.07) |
| 2026-05-29T15:03:27Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $932 < absolute min $5000 |
| 2026-05-29T15:03:27Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $2947 < absolute min $5000 |
| 2026-05-29T15:03:27Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3700 | P3_low_absolute_liquidity | liquidity $1093 < absolute min $5000 |
| 2026-05-29T15:03:27Z | puffpaw-fdv-above-100m-one-day-after-launch | crypto-launch | 0.5400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1462 < absolute min $5000 |
| 2026-05-29T15:03:27Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.2150 | P3_low_absolute_liquidity | liquidity $2352 < absolute min $5000 |
| 2026-05-29T15:03:27Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12369) |
| 2026-05-29T15:03:27Z | puffpaw-fdv-above-200m-one-day-after-launch | crypto-launch | 0.2600 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | bitcoin-all-time-high-by-december-31-2026 | market | 0.1400 | P11 | tradingview BTCUSDT sentiment=bearish (conf 0.80) contradicts bull thesis (price_yes=0.94) |
| 2026-05-29T15:03:27Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52754) |
| 2026-05-29T15:03:27Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45511) |
| 2026-05-29T15:03:27Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $18910) |
| 2026-05-29T15:03:27Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $113801) |
| 2026-05-29T15:03:27Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $261611) |
| 2026-05-29T15:03:27Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $22737) |
| 2026-05-29T15:03:27Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P3_low_absolute_liquidity | liquidity $4509 < absolute min $5000 |
| 2026-05-29T15:03:27Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7800 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $44789) |
| 2026-05-29T15:03:27Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3095 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 580 d, liq $12313) |
| 2026-05-29T15:03:27Z | predictfun-fdv-above-600m-one-day-after-launch | crypto-launch | 0.3200 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T15:03:27Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1223794) |
| 2026-05-29T15:03:27Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2377664) |
| 2026-05-29T15:03:27Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1935242) |
| 2026-05-29T15:03:27Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1918095) |
| 2026-05-29T15:03:27Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $768619) |
| 2026-05-29T15:03:27Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $291661) |
| 2026-05-29T15:03:27Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1086113) |
| 2026-05-29T15:03:27Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1939806) |
| 2026-05-29T15:03:27Z | will-ro-khanna-win-the-2028-democratic-presidential-nomination | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $264811) |
| 2026-05-29T15:03:27Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1766828) |
| 2026-05-29T15:03:27Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2148735) |
| 2026-05-29T15:03:27Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T15:03:27Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2277 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2600 | P3_low_absolute_liquidity | liquidity $547 < absolute min $5000 |
| 2026-05-29T15:03:27Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $11111) |
| 2026-05-29T15:03:27Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $16038) |
| 2026-05-29T15:03:27Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $17070) |
| 2026-05-29T15:30:04Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -149 d, liq $44721) |
| 2026-05-29T15:30:04Z | will-us-unemployment-reach-at-least-5pt5-in-2026 | other | 0.1970 | P3_low_absolute_liquidity | liquidity $811 < absolute min $5000 |
| 2026-05-29T15:30:04Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 118 días) |
| 2026-05-29T15:30:04Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $766685) |
| 2026-05-29T15:30:04Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -68 d, liq $32356) |
| 2026-05-29T15:30:04Z | will-jernej-vrtovec-be-the-next-prime-minister-of-slovenia | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -68 d, liq $25550) |
| 2026-05-29T15:30:04Z | will-matej-tonin-be-the-next-prime-minister-of-slovenia | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -68 d, liq $13409) |
| 2026-05-29T15:30:04Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $41741) |
| 2026-05-29T15:30:04Z | will-robert-golob-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $20479) |
| 2026-05-29T15:30:04Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -66 d, liq $29848) |
| 2026-05-29T15:30:04Z | will-lars-boje-mathiesen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon -66 d, liq $24975) |
| 2026-05-29T15:30:04Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -66 d, liq $27814) |
| 2026-05-29T15:30:04Z | will-the-us-invade-venezuela-by-december-31-2026 | geopolitics | 0.1100 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T15:30:04Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $199314) |
| 2026-05-29T15:30:04Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -46 d, liq $344883) |
| 2026-05-29T15:30:04Z | will-a-claude-mythos-model-be-released-by-june-30-2026-684-583 | other | 0.3700 | P2 | mercado ya expiró (endDate=2026-04-30T00:00:00Z, hace 29 días) |
| 2026-05-29T15:30:04Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.7800 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T15:30:04Z | us-x-iran-diplomatic-meeting-by-june-15-2026-671 | geopolitics | 0.5200 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T15:30:04Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0550 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T15:30:04Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -12 d, liq $5902) |
| 2026-05-29T15:30:04Z | will-brad-raffensperger-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -10 d, liq $42386) |
| 2026-05-29T15:30:04Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -10 d, liq $5614) |
| 2026-05-29T15:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -2 d, liq $462335) |
| 2026-05-29T15:30:04Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-28-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -1 d, liq $263737) |
| 2026-05-29T15:30:09Z | highest-temperature-in-shanghai-on-may-29-2026-26c | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $6035) |
| 2026-05-29T15:30:09Z | elon-musk-of-tweets-may-22-may-29-240-259 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $75602) |
| 2026-05-29T15:30:09Z | elon-musk-of-tweets-may-22-may-29-200-219 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $81358) |
| 2026-05-29T15:30:09Z | elon-musk-of-tweets-may-22-may-29-160-179 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $172505) |
| 2026-05-29T15:30:09Z | elon-musk-of-tweets-may-22-may-29-220-239 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $108467) |
| 2026-05-29T15:30:09Z | bitcoin-above-78k-on-may-29-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $192803) |
| 2026-05-29T15:30:09Z | bitcoin-above-72k-on-may-29-2026 | market | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T15:30:11Z | bitcoin-above-74k-on-may-29-2026 | market | 0.0990 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.10) |
| 2026-05-29T15:30:11Z | bitcoin-above-76k-on-may-29-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $82093) |
| 2026-05-29T15:30:14Z | nor-rbk-bog-2026-05-29-bog | other | 0.6000 | E2 | edge 0.020 < mín 0.030 (p̂=0.620, implied=0.600) |
| 2026-05-29T15:30:14Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $16947) |
| 2026-05-29T15:30:17Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.4800 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T15:30:23Z | ucl-psg-ars-2026-05-30-ars | other | 0.3100 | E2 | edge 0.010 < mín 0.030 (p̂=0.320, implied=0.310) |
| 2026-05-29T15:30:23Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $354446) |
| 2026-05-29T15:30:23Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0520 | P9 | P9: geopolitics pump cluster (price 0.05, 1d) |
| 2026-05-29T15:30:27Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0310 | E2 | edge 0.009 < mín 0.030 (p̂=0.040, implied=0.031) |
| 2026-05-29T15:30:27Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $198403) |
| 2026-05-29T15:30:27Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $187662) |
| 2026-05-29T15:30:31Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.1490 | P9 | P9: geopolitics pump cluster (price 0.15, 1d) |
| 2026-05-29T15:30:35Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0420 | V3: Trigger vago sin fecha concreta o sin evento verificable. V6: Sin catalyst identificable en los próximos 7 días. | V3: Trigger vago sin fecha concreta o sin evento verificable. V6: Sin catalyst identificable en los próximos 7 días... |
| 2026-05-29T15:30:35Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0650 | P9 | P9: geopolitics pump cluster (price 0.07, 1d) |
| 2026-05-29T15:30:35Z | will-60-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0500 | P9 | P9: geopolitics pump cluster (price 0.05, 1d) |
| 2026-05-29T15:30:35Z | will-donald-trump-publicly-insult-someone-on-may-28-2026 | other | 0.9940 | P0_ceiling | price ceiling: 0.9940 > 0.980 |
| 2026-05-29T15:30:35Z | trump-out-as-president-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $870202) |
| 2026-05-29T15:30:35Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 1 d, liq $58591) |
| 2026-05-29T15:30:35Z | uefa-champions-league-unbeaten-team | sports-season | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T15:30:35Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0610 | P9 | P9: geopolitics pump cluster (price 0.06, 1d) |
| 2026-05-29T15:30:35Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.1310 | P9 | P9: geopolitics pump cluster (price 0.13, 1d) |
| 2026-05-29T15:30:35Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 1 d, liq $245254) |
| 2026-05-29T15:30:38Z | will-trump-restart-project-freedom-by-may-31 | other | 0.0330 | V3, V5, V6 | V3, V5, V6: V3 Trigger vago: 'Project Freedom' no es un evento verificable concreto con fecha definida ni existe un c... |
| 2026-05-29T15:30:38Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 1 d, liq $277902) |
| 2026-05-29T15:30:38Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.2790 | P9 | P9: geopolitics pump cluster (price 0.28, 1d) |
| 2026-05-29T15:30:38Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $167097) |
| 2026-05-29T15:30:38Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $3608288) |
| 2026-05-29T15:30:44Z | will-trump-agree-to-iranian-oil-sanction-relief-by-may-31 | executive-action | 0.3000 | E2 | edge 0.020 < mín 0.030 (p̂=0.280, implied=0.300) |
| 2026-05-29T15:30:49Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0520 | P9 | P9: geopolitics pump cluster (price 0.05, 1d) |
| 2026-05-29T15:30:49Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.1600 | P9 | P9: geopolitics pump cluster (price 0.16, 1d) |
| 2026-05-29T15:30:56Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $200857) |
| 2026-05-29T15:30:56Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $221663) |
| 2026-05-29T15:31:00Z | will-psg-win-the-202526-champions-league | sports-season | 0.5800 | V3 | V3: V3 Trigger vago: la pregunta se refiere a un evento que ocurriría en mayo de 2026, demasiado lejano en el futuro... |
| 2026-05-29T15:31:00Z | will-declan-rice-record-the-most-yellow-cards-in-2025-26-uefa-champions-league | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 1 d, liq $5893) |
| 2026-05-29T15:31:00Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $111486) |
| 2026-05-29T15:31:00Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $69261) |
| 2026-05-29T15:31:00Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $101638) |
| 2026-05-29T15:31:00Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $91421) |
| 2026-05-29T15:31:00Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $114358) |
| 2026-05-29T15:31:00Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $116310) |
| 2026-05-29T15:31:00Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $123744) |
| 2026-05-29T15:31:00Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $118603) |
| 2026-05-29T15:31:00Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $83222) |
| 2026-05-29T15:31:00Z | will-paloma-valencia-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $49780) |
| 2026-05-29T15:31:00Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $125406) |
| 2026-05-29T15:31:00Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $67279) |
| 2026-05-29T15:31:08Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $83958) |
| 2026-05-29T15:31:08Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $65493) |
| 2026-05-29T15:31:08Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $93872) |
| 2026-05-29T15:31:08Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 2 d, liq $314618) |
| 2026-05-29T15:31:11Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.5400 | E2 | edge 0.010 < mín 0.030 (p̂=0.550, implied=0.540) |
| 2026-05-29T15:31:11Z | will-wti-reach-130-in-may-2026-733 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $146614) |
| 2026-05-29T15:31:11Z | will-wti-crude-oil-wti-hit-high-105-in-may | market | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 2 d, liq $77311) |
| 2026-05-29T15:31:11Z | will-wti-reach-110-in-may-2026-116-472 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $40506) |
| 2026-05-29T15:31:11Z | will-wti-reach-120-in-may-2026-592-217 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $118398) |
| 2026-05-29T15:31:11Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0850 | P6 | P6 market: CL=F spot $87.20 already > target $80.00 but yes=0.09 (confused book) |
| 2026-05-29T15:31:11Z | will-wti-reach-150-in-may-2026-196-364 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $113893) |
| 2026-05-29T15:31:11Z | will-wti-crude-oil-wti-hit-high-115-in-may-221 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $57566) |
| 2026-05-29T15:31:11Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 2 d, liq $86782) |
| 2026-05-29T15:31:11Z | will-wti-reach-140-in-may-2026-916 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $147695) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-72500-in-may-2026-from-may-27 | market | 0.5000 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.10) |
| 2026-05-29T15:31:11Z | will-bitcoin-reach-150k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $1134704) |
| 2026-05-29T15:31:11Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $208105) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $208765) |
| 2026-05-29T15:31:11Z | will-ethereum-reach-3600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $84199) |
| 2026-05-29T15:31:11Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $10337) |
| 2026-05-29T15:31:11Z | will-ethereum-reach-4000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $76161) |
| 2026-05-29T15:31:11Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $61022) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-50k-in-may-2026-896 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $167717) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $157616) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-35k-in-may-2026-217-769-834 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $124386) |
| 2026-05-29T15:31:11Z | will-solana-dip-to-50-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $53506) |
| 2026-05-29T15:31:11Z | will-ethereum-dip-to-400-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $74585) |
| 2026-05-29T15:31:11Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $242324) |
| 2026-05-29T15:31:11Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0230 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.02) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0590 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.10) |
| 2026-05-29T15:31:11Z | will-bitcoin-dip-to-65k-in-may-2026-183-857-425 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $191014) |
| 2026-05-29T15:31:15Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2000 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: El evento carece de un catalyst veri... |
| 2026-05-29T15:31:15Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $37140) |
| 2026-05-29T15:31:15Z | will-rob-sand-win-the-2026-iowa-governor-democratic-primary-election | elections | 0.9960 | P0_ceiling | price ceiling: 0.9960 > 0.980 |
| 2026-05-29T15:31:15Z | will-robert-lebovics-be-the-republican-nominee-for-senate-in-new-jersey-744 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 3 d, liq $6345) |
| 2026-05-29T15:31:15Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $17497) |
| 2026-05-29T15:31:15Z | elon-musk-of-tweets-may-26-june-2-60-79 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $63127) |
| 2026-05-29T15:31:15Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 4 d, liq $43131) |
| 2026-05-29T15:31:20Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.6800 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto en los próximos 7 días. La fecha del market es junio 3, 20... |
| 2026-05-29T15:31:20Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $29770) |
| 2026-05-29T15:31:20Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $43126) |
| 2026-05-29T15:31:20Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $29966) |
| 2026-05-29T15:31:23Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $94721) |
| 2026-05-29T15:31:23Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30723) |
| 2026-05-29T15:31:27Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $93662) |
| 2026-05-29T15:31:27Z | will-the-rebuilding-korea-party-rkp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $82650) |
| 2026-05-29T15:31:27Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $23021) |
| 2026-05-29T15:31:27Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $40495) |
| 2026-05-29T15:31:27Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $44957) |
| 2026-05-29T15:31:27Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $21780) |
| 2026-05-29T15:31:27Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $35022) |
| 2026-05-29T15:31:27Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $244614) |
| 2026-05-29T15:31:27Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $110080) |
| 2026-05-29T15:31:27Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 8 d, liq $69360) |
| 2026-05-29T15:31:27Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $225737) |
| 2026-05-29T15:31:31Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 8 d, liq $31682) |
| 2026-05-29T15:31:31Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $204403) |
| 2026-05-29T15:31:33Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0470 | V3 Trigger vago | V3 Trigger vago: V3 Trigger vago: Falta fecha concreta o evento verificable para un torneo que finaliza en 2026, sin ... |
| 2026-05-29T15:31:33Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 8 d, liq $34727) |
| 2026-05-29T15:31:33Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $211547) |
| 2026-05-29T15:31:33Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $193854) |
| 2026-05-29T15:31:37Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0210 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda afectar sig... |
| 2026-05-29T15:31:37Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $243128) |
| 2026-05-29T15:31:37Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $239306) |
| 2026-05-29T15:31:41Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0600 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: La pregunta no especifica un evento c... |
| 2026-05-29T15:31:44Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0370 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda modificar s... |
| 2026-05-29T15:31:48Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1070 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda afectar materialmente la... |
| 2026-05-29T15:31:48Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $208144) |
| 2026-05-29T15:31:48Z | will-tommy-paul-win-the-2026-mens-french-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 8 d, liq $51423) |
| 2026-05-29T15:31:48Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $220448) |
| 2026-05-29T15:31:53Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0220 | V3, V5 | V3, V5: V3: El evento no tiene un catalizador verificable concreto en los próximos 7 días. La pregunta es sobre un ... |
| 2026-05-29T15:32:01Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $188331) |
| 2026-05-29T15:32:01Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $196362) |
| 2026-05-29T15:32:01Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152157) |
| 2026-05-29T15:32:04Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0280 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: No se puede verificar un evento disc... |
| 2026-05-29T15:32:04Z | will-alex-michelsen-win-the-2026-mens-french-open | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 8 d, liq $38997) |
| 2026-05-29T15:32:04Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $121928) |
| 2026-05-29T15:32:04Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 8 d, liq $22090) |
| 2026-05-29T15:32:17Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.2600 | P9 | P9: geopolitics pump cluster (price 0.26, 8d) |
| 2026-05-29T15:32:17Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $239763) |
| 2026-05-29T15:32:17Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $158272) |
| 2026-05-29T15:32:32Z | will-the-san-antonio-spurs-win-the-nba-western-conference-finals | sports-season | 0.4200 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: El evento es una predicción sobre s... |
| 2026-05-29T15:32:32Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 18 d, liq $727461) |
| 2026-05-29T15:32:36Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 18 d, liq $2561436) |
| 2026-05-29T15:32:36Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 18 d, liq $285667) |
| 2026-05-29T15:32:36Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 18 d, liq $639886) |
| 2026-05-29T15:32:36Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $172522) |
| 2026-05-29T15:32:36Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $135258) |
| 2026-05-29T15:32:36Z | will-claudia-lpez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $129669) |
| 2026-05-29T15:32:36Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $191505) |
| 2026-05-29T15:32:36Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $210863) |
| 2026-05-29T15:32:43Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6400 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: Veto V3: Trigger vago. La pregunta no especifica una f... |
| 2026-05-29T15:32:43Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $150714) |
| 2026-05-29T15:32:43Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $133415) |
| 2026-05-29T15:32:43Z | will-luis-gilberto-murillo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $148401) |
| 2026-05-29T15:32:43Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $180055) |
| 2026-05-29T15:32:47Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0350 | E2 | edge 0.015 < mín 0.030 (p̂=0.020, implied=0.035) |
| 2026-05-29T15:32:51Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | V3 | V3: V3 Trigger vago: El evento es una primaria para gobernador en 2026, sin fecha concreta de primaria (más allá de... |
| 2026-05-29T15:32:51Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $2062) |
| 2026-05-29T15:32:51Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 25 d, liq $1471) |
| 2026-05-29T15:32:51Z | will-viggo-bjrck-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 28 d, liq $565) |
| 2026-05-29T15:32:51Z | will-tynan-lawrence-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $2638) |
| 2026-05-29T15:32:51Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $1258) |
| 2026-05-29T15:32:58Z | will-russia-enter-kramatorsk-by-june-30-821-192 | geopolitics | 0.0280 | E2 | edge 0.002 < mín 0.030 (p̂=0.030, implied=0.028) |
| 2026-05-29T15:33:05Z | aaron-taylor-johnson-announced-as-next-james-bond | other | 0.0420 | P3_low_absolute_liquidity | liquidity $3894 < absolute min $5000 |
| 2026-05-29T15:33:05Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $16911) |
| 2026-05-29T15:33:05Z | will-zai-have-the-top-ai-model-at-the-end-of-june-2026-987 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 31 d, liq $13360) |
| 2026-05-29T15:33:05Z | ukraine-peace-referendum-scheduled-by-june-30 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $23930) |
| 2026-05-29T15:33:05Z | harris-dickinson-announced-as-next-james-bond-784 | other | 0.0300 | P3_low_absolute_liquidity | liquidity $3941 < absolute min $5000 |
| 2026-05-29T15:33:05Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $322913) |
| 2026-05-29T15:33:05Z | paul-mescal-announced-as-next-james-bond | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 31 d, liq $4650) |
| 2026-05-29T15:33:05Z | hamad-bin-isa-al-khalifa-out-as-bahrain-king | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 31 d, liq $8408) |
| 2026-05-29T15:33:09Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0210 | V3 | V3: V3 Trigger vago: 'fall' del régimen iraní es un concepto ambiguo sin una definición verificable concreta (¿co... |
| 2026-05-29T15:33:12Z | will-deepseek-have-the-top-ai-model-at-the-end-of-june-2026-514 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 31 d, liq $8593) |
| 2026-05-29T15:33:16Z | will-donald-trump-be-confirmed-to-have-visited-epsteins-island | other | 0.0200 | E2 | edge 0.010 < mín 0.030 (p̂=0.030, implied=0.020) |
| 2026-05-29T15:33:16Z | jack-lowdon-announced-as-next-james-bond | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 31 d, liq $4191) |
| 2026-05-29T15:33:19Z | us-announces-new-iran-agreementceasefire-extension-by-june-30-848-925 | geopolitics | 0.8400 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-29T15:33:19Z | henry-cavill-announced-as-next-james-bond-857-642 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 31 d, liq $4561) |
| 2026-05-29T15:33:19Z | will-trump-pardon-tiger-woods-by-june-30 | executive-action | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $25789) |
| 2026-05-29T15:33:19Z | pierce-brosnan-announced-as-next-james-bond-557 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $6930) |
| 2026-05-29T15:33:27Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V3 | V3: V3 Trigger vago: no se especifica un evento concreto ni una fecha verificable para la confirmación, lo que hace ... |
| 2026-05-29T15:33:27Z | gustavo-petro-out-as-leader-of-colombia-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $13980) |
| 2026-05-29T15:33:27Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 31 d, liq $16467) |
| 2026-05-29T15:33:27Z | james-norton-announced-as-next-james-bond-575 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $6579) |
| 2026-05-29T15:33:27Z | forsen-to-get-signed-by-an-lck-prog-org-by-jun-30 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 31 d, liq $48599) |
| 2026-05-29T15:33:27Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 31 d, liq $112131) |
| 2026-05-29T15:33:31Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-30-2026-15... | geopolitics | 0.9100 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-29T15:33:31Z | will-the-montreal-canadiens-win-the-2026-nhl-stanley-cup | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 31 d, liq $239864) |
| 2026-05-29T15:33:31Z | will-us-withdraw-from-nato-by-june-30 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $27171) |
| 2026-05-29T15:33:31Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $343290) |
| 2026-05-29T15:33:35Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0380 | E2 | edge 0.022 < mín 0.030 (p̂=0.060, implied=0.038) |
| 2026-05-29T15:33:42Z | nato-x-russia-military-clash-by-june-30-2026 | geopolitics | 0.0580 | E2 | edge 0.012 < mín 0.030 (p̂=0.070, implied=0.058) |
| 2026-05-29T15:33:42Z | will-amazon-be-the-largest-company-in-the-world-by-market-cap-on-june-30-123 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $338306) |
| 2026-05-29T15:33:50Z | will-silver-si-hit-low-45-by-end-of-june-972-272 | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 32 d, liq $14174) |
| 2026-05-29T15:33:50Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 32 d, liq $31473) |
| 2026-05-29T15:33:55Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0350 | V3 Trigger vago y V6 Sin catalyst | V3 Trigger vago y V6 Sin catalyst: Evento vago y sin catalyst concreto: 'hit (HIGH) $140' es ambiguo (posiblemente un... |
| 2026-05-29T15:33:59Z | will-crude-oil-cl-hit-high-150-by-end-of-june-788-691 | market | 0.0250 | V5 | V5: V5 Patrón débil: sin 3 fuentes independientes ni precedente análogo documentado que soporte una probabilidad s... |
| 2026-05-29T15:34:14Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.2740 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-29T15:34:21Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P3_low_absolute_liquidity | liquidity $3105 < absolute min $5000 |
| 2026-05-29T15:34:21Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 43 d, liq $3247) |
| 2026-05-29T15:34:21Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 43 d, liq $2274) |
| 2026-05-29T15:34:21Z | will-iga-witek-be-the-2026-womens-wimbledon-winner | other | 0.1890 | P3_low_absolute_liquidity | liquidity $1979 < absolute min $5000 |
| 2026-05-29T15:34:21Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0430 | P3_low_absolute_liquidity | liquidity $3318 < absolute min $5000 |
| 2026-05-29T15:34:26Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $600 < absolute min $5000 |
| 2026-05-29T15:34:26Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $7014894) |
| 2026-05-29T15:34:26Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $5555206) |
| 2026-05-29T15:34:26Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8234830) |
| 2026-05-29T15:34:26Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9493038) |
| 2026-05-29T15:34:26Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $6324314) |
| 2026-05-29T15:34:26Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $10514759) |
| 2026-05-29T15:34:33Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4528042) |
| 2026-05-29T15:34:33Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9226625) |
| 2026-05-29T15:34:33Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $8263369) |
| 2026-05-29T15:34:33Z | will-iraq-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10690356) |
| 2026-05-29T15:34:33Z | will-team-z-win-the-2026-fifa-world-cup-316 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10728021) |
| 2026-05-29T15:34:33Z | will-jordan-win-the-2026-fifa-world-cup-233 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10313367) |
| 2026-05-29T15:34:33Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 51 d, liq $3514157) |
| 2026-05-29T15:34:33Z | will-haiti-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10535927) |
| 2026-05-29T15:34:33Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $8149244) |
| 2026-05-29T15:34:33Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8706380) |
| 2026-05-29T15:34:37Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $4034749) |
| 2026-05-29T15:34:37Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8091323) |
| 2026-05-29T15:34:37Z | will-new-zealand-win-the-2026-fifa-world-cup-635 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9889756) |
| 2026-05-29T15:34:40Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.018 < mín 0.030 (p̂=0.070, implied=0.052) |
| 2026-05-29T15:34:40Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $5810349) |
| 2026-05-29T15:34:40Z | will-qatar-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10402249) |
| 2026-05-29T15:34:40Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $3023379) |
| 2026-05-29T15:34:40Z | will-curaao-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10032809) |
| 2026-05-29T15:34:43Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La Copa del Mundo 20... |
| 2026-05-29T15:34:47Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | E2 | edge 0.026 < mín 0.030 (p̂=0.120, implied=0.094) |
| 2026-05-29T15:34:47Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 51 d, liq $9262169) |
| 2026-05-29T15:34:51Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0300 | E2 | edge 0.020 < mín 0.030 (p̂=0.010, implied=0.030) |
| 2026-05-29T15:34:51Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $4770078) |
| 2026-05-29T15:34:51Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9366204) |
| 2026-05-29T15:34:51Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $2690210) |
| 2026-05-29T15:34:51Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 51 d, liq $406680) |
| 2026-05-29T15:34:51Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $1556138) |
| 2026-05-29T15:34:51Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $9667089) |
| 2026-05-29T15:34:51Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 51 d, liq $5391783) |
| 2026-05-29T15:34:55Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0390 | E2 | edge 0.001 < mín 0.030 (p̂=0.040, implied=0.039) |
| 2026-05-29T15:34:55Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $10683001) |
| 2026-05-29T15:34:59Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $7317431) |
| 2026-05-29T15:34:59Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 51 d, liq $2367247) |
| 2026-05-29T15:34:59Z | will-cape-verde-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10381794) |
| 2026-05-29T15:34:59Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4639762) |
| 2026-05-29T15:34:59Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $4662938) |
| 2026-05-29T15:34:59Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $2679885) |
| 2026-05-29T15:34:59Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9803750) |
| 2026-05-29T15:35:05Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $32901) |
| 2026-05-29T15:35:05Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $15560) |
| 2026-05-29T15:35:05Z | will-paul-reevs-be-the-republican-nominee-for-az-01 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 66 d, liq $10866) |
| 2026-05-29T15:35:05Z | will-jay-collins-be-the-republican-nominee-for-florida-governor | elections | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 80 d, liq $35645) |
| 2026-05-29T15:35:05Z | will-tyreek-hill-play-for-philadelphia-eagles-next | other | 0.1380 | P3_low_absolute_liquidity | liquidity $284 < absolute min $5000 |
| 2026-05-29T15:35:05Z | will-david-njoku-play-for-tennessee-titans-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 94 d, liq $2472) |
| 2026-05-29T15:35:09Z | tush-push-banned-for-2026-nfl-season | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 103 d, liq $1820) |
| 2026-05-29T15:35:09Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 106 d, liq $2373) |
| 2026-05-29T15:35:09Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2455) |
| 2026-05-29T15:35:09Z | will-anastasia-potapova-win-the-2026-womens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2002) |
| 2026-05-29T15:35:09Z | will-barbora-krejcikova-win-the-2026-womens-us-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 106 d, liq $1982) |
| 2026-05-29T15:35:09Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 106 d, liq $3040) |
| 2026-05-29T15:35:09Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $3104) |
| 2026-05-29T15:35:09Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $1459) |
| 2026-05-29T15:35:09Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 106 d, liq $2134) |
| 2026-05-29T15:35:09Z | will-tereza-valentova-win-the-2026-womens-us-open | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 106 d, liq $2366) |
| 2026-05-29T15:35:09Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 121 d, liq $1920) |
| 2026-05-29T15:35:09Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 121 d, liq $1061) |
| 2026-05-29T15:35:09Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $863 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5700 | P8 | P8: election 123d out, price 0.57 en banda ruido [0.30, 0.70] |
| 2026-05-29T15:35:09Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $763 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-michelle-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 127 d, liq $221252) |
| 2026-05-29T15:35:09Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 127 d, liq $1078364) |
| 2026-05-29T15:35:09Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 155 d, liq $31917) |
| 2026-05-29T15:35:09Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 155 d, liq $27988) |
| 2026-05-29T15:35:09Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 155 d, liq $27640) |
| 2026-05-29T15:35:09Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 155 d, liq $29266) |
| 2026-05-29T15:35:09Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $13258) |
| 2026-05-29T15:35:09Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $164608) |
| 2026-05-29T15:35:09Z | will-corey-seager-win-the-2026-american-league-hank-aaron-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 167 d, liq $2590) |
| 2026-05-29T15:35:09Z | will-sergio-prez-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 190 d, liq $753390) |
| 2026-05-29T15:35:09Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 193 d, liq $11702) |
| 2026-05-29T15:35:09Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $11544) |
| 2026-05-29T15:35:09Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $8758) |
| 2026-05-29T15:35:09Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 203 d, liq $5761) |
| 2026-05-29T15:35:09Z | will-payton-tolle-win-the-2026-al-rookie-of-the-year-award | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 203 d, liq $7822) |
| 2026-05-29T15:35:09Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $173) |
| 2026-05-29T15:35:09Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $2043 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $6961) |
| 2026-05-29T15:35:09Z | will-ahmad-hosseini-khorasani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $49184) |
| 2026-05-29T15:35:09Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4485 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-gholam-ali-haddad-adel-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $48677) |
| 2026-05-29T15:35:09Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $280) |
| 2026-05-29T15:35:09Z | ramp-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $2738 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $150 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-mohammad-khatami-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $44566) |
| 2026-05-29T15:35:09Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1669 < absolute min $5000 |
| 2026-05-29T15:35:09Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0650 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T15:35:09Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 215 d, liq $14134) |
| 2026-05-29T15:35:09Z | ripple-labs-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $2042 < absolute min $5000 |
| 2026-05-29T15:35:09Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4039 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $399) |
| 2026-05-29T15:35:09Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2885 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $228 < absolute min $5000 |
| 2026-05-29T15:35:09Z | 10pt0-or-above-earthquake-before-2027 | weather-natural | 0.0500 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T15:35:09Z | will-mahmoud-ahmadinejad-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $46841) |
| 2026-05-29T15:35:09Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1011 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-sadegh-larijani-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $42728) |
| 2026-05-29T15:35:09Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $155822) |
| 2026-05-29T15:35:09Z | will-muhammad-mirbaqiri-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $46302) |
| 2026-05-29T15:35:09Z | will-massoud-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $48540) |
| 2026-05-29T15:35:09Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $7646) |
| 2026-05-29T15:35:09Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $10998) |
| 2026-05-29T15:35:09Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $133020) |
| 2026-05-29T15:35:09Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | sports-season | 0.0890 | P3_low_absolute_liquidity | liquidity $808 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $14315) |
| 2026-05-29T15:35:09Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $205) |
| 2026-05-29T15:35:09Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $552) |
| 2026-05-29T15:35:09Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $63825) |
| 2026-05-29T15:35:09Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1220 | P3_low_absolute_liquidity | liquidity $3173 < absolute min $5000 |
| 2026-05-29T15:35:09Z | natural-disaster-in-2026 | other | 0.2800 | P3_low_absolute_liquidity | liquidity $1542 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $18515) |
| 2026-05-29T15:35:09Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $5508) |
| 2026-05-29T15:35:09Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $3247 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4114 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $447) |
| 2026-05-29T15:35:09Z | will-maryam-rajavi-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $46757) |
| 2026-05-29T15:35:09Z | anduril-ipo-before-2027 | crypto-launch | 0.1200 | P3_low_absolute_liquidity | liquidity $3498 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-petr-yan-fight-dominick-cruz-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $1142) |
| 2026-05-29T15:35:09Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3641 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $5855) |
| 2026-05-29T15:35:09Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4894 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $91397) |
| 2026-05-29T15:35:09Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $557) |
| 2026-05-29T15:35:09Z | will-hassan-khomeini-be-head-of-state-in-iran-end-of-2026 | geopolitics | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 215 d, liq $47276) |
| 2026-05-29T15:35:09Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $15608) |
| 2026-05-29T15:35:09Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $19234) |
| 2026-05-29T15:35:09Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $16907) |
| 2026-05-29T15:35:09Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $3105) |
| 2026-05-29T15:35:09Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $1700 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-usdt-market-cap-hit-200b-before-2027 | other | 0.9780 | P3_low_absolute_liquidity | liquidity $948 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.02) |
| 2026-05-29T15:35:09Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.02) |
| 2026-05-29T15:35:09Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12376) |
| 2026-05-29T15:35:09Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $454 < absolute min $5000 |
| 2026-05-29T15:35:09Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3700 | P3_low_absolute_liquidity | liquidity $1096 < absolute min $5000 |
| 2026-05-29T15:35:09Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3381 < absolute min $5000 |
| 2026-05-29T15:35:09Z | bitcoin-all-time-high-by-december-31-2026 | market | 0.1500 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.10) |
| 2026-05-29T15:35:09Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | okx-ipo-in-2026 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $987 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $789 < absolute min $5000 |
| 2026-05-29T15:35:09Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | puffpaw-fdv-above-100m-one-day-after-launch | crypto-launch | 0.5400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1616 < absolute min $5000 |
| 2026-05-29T15:35:09Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.1980 | P3_low_absolute_liquidity | liquidity $2346 < absolute min $5000 |
| 2026-05-29T15:35:09Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1395 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2840 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52755) |
| 2026-05-29T15:35:09Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $18813) |
| 2026-05-29T15:35:09Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45416) |
| 2026-05-29T15:35:09Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $113639) |
| 2026-05-29T15:35:09Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $260104) |
| 2026-05-29T15:35:09Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7700 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $23728) |
| 2026-05-29T15:35:09Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45747) |
| 2026-05-29T15:35:09Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $4487 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 580 d, liq $12667) |
| 2026-05-29T15:35:09Z | predictfun-fdv-above-600m-one-day-after-launch | crypto-launch | 0.3200 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T15:35:09Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $777117) |
| 2026-05-29T15:35:09Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $286680) |
| 2026-05-29T15:35:09Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1774814) |
| 2026-05-29T15:35:09Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1923001) |
| 2026-05-29T15:35:09Z | will-ro-khanna-win-the-2028-democratic-presidential-nomination | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $252545) |
| 2026-05-29T15:35:09Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2375242) |
| 2026-05-29T15:35:09Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1087836) |
| 2026-05-29T15:35:09Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1948430) |
| 2026-05-29T15:35:09Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1937056) |
| 2026-05-29T15:35:09Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1229248) |
| 2026-05-29T15:35:09Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2156231) |
| 2026-05-29T15:35:09Z | will-comcast-close-warner-bros-acquisition | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -1 d, liq $23080) |
| 2026-05-29T15:35:09Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $507 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22029) |
| 2026-05-29T15:35:09Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T15:35:09Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2260 < absolute min $5000 |
| 2026-05-29T15:35:09Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $17156) |
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
