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
| 2026-05-29T18:03:53Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $485) |
| 2026-05-29T18:03:53Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $5963) |
| 2026-05-29T18:03:53Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $3417 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $453) |
| 2026-05-29T18:03:53Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $438) |
| 2026-05-29T18:03:53Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 215 d, liq $6128) |
| 2026-05-29T18:03:53Z | anduril-ipo-before-2027 | crypto-launch | 0.1200 | P3_low_absolute_liquidity | liquidity $3144 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6240) |
| 2026-05-29T18:03:53Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1345 < absolute min $5000 |
| 2026-05-29T18:03:53Z | natural-disaster-in-2026 | other | 0.2800 | P3_low_absolute_liquidity | liquidity $2329 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $7418) |
| 2026-05-29T18:03:53Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $13537) |
| 2026-05-29T18:03:53Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $301) |
| 2026-05-29T18:03:53Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $95244) |
| 2026-05-29T18:03:53Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $12169) |
| 2026-05-29T18:03:53Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3317 < absolute min $5000 |
| 2026-05-29T18:03:53Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0640 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T18:03:53Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $16060) |
| 2026-05-29T18:03:53Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $10121) |
| 2026-05-29T18:03:53Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $9063) |
| 2026-05-29T18:03:53Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $127) |
| 2026-05-29T18:03:53Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $555) |
| 2026-05-29T18:03:53Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $290 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $158983) |
| 2026-05-29T18:03:53Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $530) |
| 2026-05-29T18:03:53Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $61116) |
| 2026-05-29T18:03:53Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $280) |
| 2026-05-29T18:03:53Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $11591) |
| 2026-05-29T18:03:53Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $887) |
| 2026-05-29T18:03:53Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $7925) |
| 2026-05-29T18:03:53Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2408 < absolute min $5000 |
| 2026-05-29T18:03:53Z | ramp-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $2674 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $345) |
| 2026-05-29T18:03:53Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $3765 < absolute min $5000 |
| 2026-05-29T18:03:53Z | ledger-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $1591 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $177 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $118 < absolute min $5000 |
| 2026-05-29T18:03:53Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $2713 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $10948) |
| 2026-05-29T18:03:53Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1220 | P3_low_absolute_liquidity | liquidity $3139 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4751 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $139226) |
| 2026-05-29T18:03:53Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | sports-season | 0.0910 | P3_low_absolute_liquidity | liquidity $849 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4658 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6726) |
| 2026-05-29T18:03:53Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $225) |
| 2026-05-29T18:03:53Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $2998) |
| 2026-05-29T18:03:53Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $1352 < absolute min $5000 |
| 2026-05-29T18:03:53Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12376) |
| 2026-05-29T18:03:53Z | will-ethereum-dip-to-800-by-december-31-2026 | market | 0.2100 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.02) |
| 2026-05-29T18:03:53Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $2373 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $192 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $498 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $1071 < absolute min $5000 |
| 2026-05-29T18:03:53Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.5000 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3325 < absolute min $5000 |
| 2026-05-29T18:03:53Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2320 < absolute min $5000 |
| 2026-05-29T18:03:53Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $643 < absolute min $5000 |
| 2026-05-29T18:03:53Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4000 | P3_low_absolute_liquidity | liquidity $1083 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.02) |
| 2026-05-29T18:03:53Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | will-usdt-market-cap-hit-200b-before-2027 | other | 0.9780 | P3_low_absolute_liquidity | liquidity $942 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.02) |
| 2026-05-29T18:03:53Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1296 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52835) |
| 2026-05-29T18:03:53Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $18866) |
| 2026-05-29T18:03:53Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 240 d, liq $51900) |
| 2026-05-29T18:03:53Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45606) |
| 2026-05-29T18:03:53Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $113863) |
| 2026-05-29T18:03:53Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $254101) |
| 2026-05-29T18:03:53Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 580 d, liq $12952) |
| 2026-05-29T18:03:53Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45530) |
| 2026-05-29T18:03:53Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7300 | P3_low_absolute_liquidity | liquidity $4972 < absolute min $5000 |
| 2026-05-29T18:03:53Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8300 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7700 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $23516) |
| 2026-05-29T18:03:53Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T18:03:53Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $289984) |
| 2026-05-29T18:03:53Z | will-ro-khanna-win-the-2028-democratic-presidential-nomination | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $253687) |
| 2026-05-29T18:03:53Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1926941) |
| 2026-05-29T18:03:53Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1917804) |
| 2026-05-29T18:03:53Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1941583) |
| 2026-05-29T18:03:53Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1769603) |
| 2026-05-29T18:03:53Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $771436) |
| 2026-05-29T18:03:53Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1082693) |
| 2026-05-29T18:03:53Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2368388) |
| 2026-05-29T18:03:53Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2150087) |
| 2026-05-29T18:03:53Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2250 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2500 | P3_low_absolute_liquidity | liquidity $557 < absolute min $5000 |
| 2026-05-29T18:03:53Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T18:03:53Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22288) |
| 2026-05-29T18:03:53Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $12086) |
| 2026-05-29T18:30:04Z | israel-and-syria-normalize-relations-by-june-30-2026 | geopolitics | 0.0400 | P2 | mercado ya expiró (endDate=2025-12-31T00:00:00Z, hace 149 días) |
| 2026-05-29T18:30:04Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon -149 d, liq $38491) |
| 2026-05-29T18:30:04Z | will-us-unemployment-reach-at-least-5pt5-in-2026 | other | 0.1950 | P3_low_absolute_liquidity | liquidity $824 < absolute min $5000 |
| 2026-05-29T18:30:04Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 118 días) |
| 2026-05-29T18:30:04Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $766685) |
| 2026-05-29T18:30:04Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -68 d, liq $32367) |
| 2026-05-29T18:30:04Z | will-robert-golob-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $14550) |
| 2026-05-29T18:30:04Z | will-jernej-vrtovec-be-the-next-prime-minister-of-slovenia | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -68 d, liq $19155) |
| 2026-05-29T18:30:04Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -68 d, liq $35380) |
| 2026-05-29T18:30:04Z | will-matej-tonin-be-the-next-prime-minister-of-slovenia | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -68 d, liq $7477) |
| 2026-05-29T18:30:04Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -66 d, liq $28035) |
| 2026-05-29T18:30:04Z | will-lars-boje-mathiesen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon -66 d, liq $23928) |
| 2026-05-29T18:30:04Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -66 d, liq $28713) |
| 2026-05-29T18:30:04Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -66 d, liq $19850) |
| 2026-05-29T18:30:04Z | will-the-us-invade-venezuela-by-december-31-2026 | geopolitics | 0.1100 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T18:30:04Z | weed-rescheduled-by-december-31 | other | 0.3030 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T18:30:04Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $61585) |
| 2026-05-29T18:30:04Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -46 d, liq $344892) |
| 2026-05-29T18:30:04Z | us-x-iran-diplomatic-meeting-by-june-15-2026-671 | geopolitics | 0.5500 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T18:30:04Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0650 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T18:30:04Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.7600 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T18:30:04Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -12 d, liq $5909) |
| 2026-05-29T18:30:04Z | will-brad-raffensperger-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -10 d, liq $43166) |
| 2026-05-29T18:30:04Z | will-gregg-kirkpatrick-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -10 d, liq $10239) |
| 2026-05-29T18:30:04Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -10 d, liq $5939) |
| 2026-05-29T18:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -2 d, liq $410762) |
| 2026-05-29T18:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-29 | geopolitics | 0.1900 | P9 | P9: geopolitics pump cluster (price 0.19, 0d) |
| 2026-05-29T18:30:04Z | highest-temperature-in-shanghai-on-may-29-2026-26c | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $345547) |
| 2026-05-29T18:30:04Z | elon-musk-of-tweets-may-22-may-29-240-259 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $394212) |
| 2026-05-29T18:30:04Z | elon-musk-of-tweets-may-22-may-29-200-219 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $450633) |
| 2026-05-29T18:30:06Z | nor-rbk-bog-2026-05-29-bog | other | 0.4800 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T18:30:09Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17605) |
| 2026-05-29T18:30:09Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.2700 | P9 | P9: geopolitics pump cluster (price 0.27, 0d) |
| 2026-05-29T18:30:09Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17292) |
| 2026-05-29T18:30:16Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.1490 | P9 | P9: geopolitics pump cluster (price 0.15, 1d) |
| 2026-05-29T18:30:19Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0240 | E2 | edge 0.004 < mín 0.030 (p̂=0.020, implied=0.024) |
| 2026-05-29T18:30:19Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0580 | P9 | P9: geopolitics pump cluster (price 0.06, 1d) |
| 2026-05-29T18:30:19Z | israeli-forces-cross-the-litani-river-again-by-may-31 | other | 0.9900 | P0_ceiling | price ceiling: 0.9900 > 0.980 |
| 2026-05-29T18:30:19Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $187931) |
| 2026-05-29T18:30:19Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0530 | P9 | P9: geopolitics pump cluster (price 0.05, 1d) |
| 2026-05-29T18:30:22Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.1620 | P9 | P9: geopolitics pump cluster (price 0.16, 1d) |
| 2026-05-29T18:30:22Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $350280) |
| 2026-05-29T18:30:25Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0250 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: la pregunta no especifica un evento v... |
| 2026-05-29T18:30:27Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 1 d, liq $290025) |
| 2026-05-29T18:30:27Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.1300 | P9 | P9: geopolitics pump cluster (price 0.13, 1d) |
| 2026-05-29T18:30:27Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0920 | P9 | P9: geopolitics pump cluster (price 0.09, 1d) |
| 2026-05-29T18:30:27Z | will-40-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.1700 | P9 | P9: geopolitics pump cluster (price 0.17, 1d) |
| 2026-05-29T18:30:27Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $200956) |
| 2026-05-29T18:30:27Z | trump-out-as-president-by-may-31 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $869582) |
| 2026-05-29T18:30:27Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $167196) |
| 2026-05-29T18:30:30Z | will-trump-restart-project-freedom-by-may-31 | other | 0.0490 | V3 | V3: V3 Trigger vago: 'restart Project Freedom' no está definido como un evento verificable concreto, sin fecha o doc... |
| 2026-05-29T18:30:30Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 1 d, liq $195427) |
| 2026-05-29T18:30:32Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $198503) |
| 2026-05-29T18:30:34Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0380 | E2 | edge 0.018 < mín 0.030 (p̂=0.020, implied=0.038) |
| 2026-05-29T18:30:34Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $3217737) |
| 2026-05-29T18:30:37Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.1930 | P9 | P9: geopolitics pump cluster (price 0.19, 1d) |
| 2026-05-29T18:30:37Z | will-rajasthan-royals-win-the-2026-indian-premier-league | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $38296) |
| 2026-05-29T18:30:37Z | will-declan-rice-record-the-most-yellow-cards-in-2025-26-uefa-champions-league | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 1 d, liq $5893) |
| 2026-05-29T18:30:37Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.1200 | P9 | P9: geopolitics pump cluster (price 0.12, 1d) |
| 2026-05-29T18:30:37Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $221662) |
| 2026-05-29T18:30:46Z | will-nvidia-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $41323) |
| 2026-05-29T18:30:46Z | will-nvidia-be-the-largest-company-in-the-world-by-market-cap-on-may-31-971 | market | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T18:30:46Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $85416) |
| 2026-05-29T18:30:46Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $66006) |
| 2026-05-29T18:30:46Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $125981) |
| 2026-05-29T18:30:46Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $125882) |
| 2026-05-29T18:30:46Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $101885) |
| 2026-05-29T18:30:46Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116516) |
| 2026-05-29T18:30:46Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $67581) |
| 2026-05-29T18:30:46Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $84606) |
| 2026-05-29T18:30:46Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $69689) |
| 2026-05-29T18:30:46Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $94339) |
| 2026-05-29T18:30:46Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $118767) |
| 2026-05-29T18:30:46Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $111739) |
| 2026-05-29T18:30:46Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $91664) |
| 2026-05-29T18:30:46Z | will-paloma-valencia-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 1 d, liq $53859) |
| 2026-05-29T18:30:46Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116302) |
| 2026-05-29T18:30:46Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 2 d, liq $287373) |
| 2026-05-29T18:30:46Z | will-wti-reach-120-in-may-2026-592-217 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $70175) |
| 2026-05-29T18:30:46Z | will-wti-reach-140-in-may-2026-916 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $76736) |
| 2026-05-29T18:30:46Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 2 d, liq $83188) |
| 2026-05-29T18:30:46Z | will-wti-dip-to-30-in-may-2026-668-273-232-424-957-651-726-496-699-125-667-922-317 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $29654) |
| 2026-05-29T18:30:46Z | will-wti-reach-130-in-may-2026-733 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $65399) |
| 2026-05-29T18:30:46Z | will-wti-crude-oil-wti-hit-high-105-in-may | market | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 2 d, liq $112787) |
| 2026-05-29T18:30:46Z | will-wti-reach-150-in-may-2026-196-364 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $87340) |
| 2026-05-29T18:30:46Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.2400 | P6 | P6 market: CL=F spot $87.25 already > target $85.00 but yes=0.24 (confused book) |
| 2026-05-29T18:30:46Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0570 | P6 | P6 market: CL=F spot $87.25 already > target $80.00 but yes=0.06 (confused book) |
| 2026-05-29T18:30:46Z | will-wti-crude-oil-wti-hit-high-115-in-may-221 | market | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 2 d, liq $34243) |
| 2026-05-29T18:30:46Z | will-wti-dip-to-20-in-may-2026-865-174-115-863-197-663-536-962-219-163-329-638 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $50954) |
| 2026-05-29T18:30:46Z | will-wti-reach-110-in-may-2026-116-472 | market | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 2 d, liq $52074) |
| 2026-05-29T18:30:46Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $47529) |
| 2026-05-29T18:30:46Z | will-solana-dip-to-50-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $53506) |
| 2026-05-29T18:30:46Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $176841) |
| 2026-05-29T18:30:46Z | will-bitcoin-dip-to-50k-in-may-2026-896 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $166489) |
| 2026-05-29T18:30:46Z | will-ethereum-reach-3600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $78011) |
| 2026-05-29T18:30:46Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $61022) |
| 2026-05-29T18:30:46Z | will-ethereum-dip-to-1200-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $32760) |
| 2026-05-29T18:30:48Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $152669) |
| 2026-05-29T18:30:48Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $10337) |
| 2026-05-29T18:30:48Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $179221) |
| 2026-05-29T18:30:51Z | will-bitcoin-dip-to-35k-in-may-2026-217-769-834 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $115372) |
| 2026-05-29T18:30:51Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $187476) |
| 2026-05-29T18:30:51Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 2 d, liq $48321) |
| 2026-05-29T18:30:51Z | will-bitcoin-reach-150k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $1124734) |
| 2026-05-29T18:30:51Z | will-robert-lebovics-be-the-republican-nominee-for-senate-in-new-jersey-744 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 3 d, liq $6743) |
| 2026-05-29T18:30:55Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2000 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T18:30:55Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $37267) |
| 2026-05-29T18:30:55Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $17770) |
| 2026-05-29T18:30:55Z | will-rob-sand-win-the-2026-iowa-governor-democratic-primary-election | elections | 0.9960 | P0_ceiling | price ceiling: 0.9960 > 0.980 |
| 2026-05-29T18:30:55Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 3 d, liq $53411) |
| 2026-05-29T18:30:55Z | elon-musk-of-tweets-may-26-june-2-60-79 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $75178) |
| 2026-05-29T18:30:55Z | elon-musk-of-tweets-may-26-june-2-360-379 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $45836) |
| 2026-05-29T18:30:58Z | elon-musk-of-tweets-may-26-june-2-100-119 | other | 0.0250 | E2 | edge 0.010 < mín 0.030 (p̂=0.015, implied=0.025) |
| 2026-05-29T18:30:58Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $21642) |
| 2026-05-29T18:30:58Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $44463) |
| 2026-05-29T18:31:01Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $40667) |
| 2026-05-29T18:31:01Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $35430) |
| 2026-05-29T18:31:01Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $25599) |
| 2026-05-29T18:31:01Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $94079) |
| 2026-05-29T18:31:01Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $23189) |
| 2026-05-29T18:31:03Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $45729) |
| 2026-05-29T18:31:03Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30385) |
| 2026-05-29T18:31:03Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30425) |
| 2026-05-29T18:31:03Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $94466) |
| 2026-05-29T18:31:03Z | will-lee-un-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $42262) |
| 2026-05-29T18:31:03Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $43856) |
| 2026-05-29T18:31:03Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $40092) |
| 2026-05-29T18:31:03Z | will-the-rebuilding-korea-party-rkp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $82839) |
| 2026-05-29T18:31:03Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30631) |
| 2026-05-29T18:31:03Z | elon-musk-of-tweets-may-29-june-5-500plus | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 6 d, liq $27113) |
| 2026-05-29T18:31:03Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $244702) |
| 2026-05-29T18:31:06Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $214193) |
| 2026-05-29T18:31:06Z | will-tommy-paul-win-the-2026-mens-french-open | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 8 d, liq $29273) |
| 2026-05-29T18:31:06Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $193778) |
| 2026-05-29T18:31:09Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 8 d, liq $27812) |
| 2026-05-29T18:31:12Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-7-2026 | geopolitics | 0.6700 | V3 - Trigger vago | V3 - Trigger vago: V3 Trigger vago: La pregunta se refiere a un anuncio de Donald Trump sobre el levantamiento de un ... |
| 2026-05-29T18:31:12Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $56203) |
| 2026-05-29T18:31:14Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0220 | V3, V5 | V3, V5: V3 Trigger vago: 'win the 2026 Men's French Open' es un evento futuro sin un catalizador concreto ni fecha l�... |
| 2026-05-29T18:31:14Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $182060) |
| 2026-05-29T18:31:14Z | will-alex-michelsen-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $394464) |
| 2026-05-29T18:31:14Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $201834) |
| 2026-05-29T18:31:14Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $51914) |
| 2026-05-29T18:31:16Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1160 | E2 | edge 0.004 < mín 0.030 (p̂=0.120, implied=0.116) |
| 2026-05-29T18:31:16Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $236880) |
| 2026-05-29T18:31:16Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $204323) |
| 2026-05-29T18:31:18Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0500 | V3 Trigger vago + V6 Sin catalyst en 7 días | V3 Trigger vago + V6 Sin catalyst en 7 días: V3 Trigger vago: el evento 'Joao Fonseca gana el Abierto de Francia 202... |
| 2026-05-29T18:31:21Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0600 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: sin fecha concreta o sin evento verif... |
| 2026-05-29T18:31:26Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0300 | V6 Sin catalyst | V6 Sin catalyst: El evento no tiene un catalizador concreto identificable en los próximos 7 días; es una pregunta s... |
| 2026-05-29T18:31:26Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $205672) |
| 2026-05-29T18:31:26Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $219474) |
| 2026-05-29T18:31:26Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.2100 | P9 | P9: geopolitics pump cluster (price 0.21, 8d) |
| 2026-05-29T18:31:26Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $212428) |
| 2026-05-29T18:31:26Z | will-enrique-valderrama-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $232962) |
| 2026-05-29T18:31:28Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0420 | V3: No menciona fecha concreta ni evento verificable próximo. | V3: No menciona fecha concreta ni evento verificable próximo.: V3 Trigger vago: el evento 'ganar el French Open 2026... |
| 2026-05-29T18:31:28Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $190125) |
| 2026-05-29T18:31:34Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 8 d, liq $31238) |
| 2026-05-29T18:31:36Z | will-jakub-mensik-win-the-2026-mens-french-open | other | 0.0310 | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... |
| 2026-05-29T18:31:36Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152205) |
| 2026-05-29T18:31:36Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $75850) |
| 2026-05-29T18:31:36Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $239580) |
| 2026-05-29T18:31:39Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152149) |
| 2026-05-29T18:31:39Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 8 d, liq $29113) |
| 2026-05-29T18:31:39Z | will-karen-khachanov-win-the-2026-mens-french-open | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 8 d, liq $42778) |
| 2026-05-29T18:31:54Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 18 d, liq $574062) |
| 2026-05-29T18:31:54Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 18 d, liq $201701) |
| 2026-05-29T18:31:56Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 18 d, liq $2556818) |
| 2026-05-29T18:31:56Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 18 d, liq $737489) |
| 2026-05-29T18:31:56Z | will-lynn-vision-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 22 d, liq $25814) |
| 2026-05-29T18:31:56Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $45670) |
| 2026-05-29T18:31:56Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $155966) |
| 2026-05-29T18:31:56Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $152163) |
| 2026-05-29T18:31:56Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $140682) |
| 2026-05-29T18:31:59Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $171344) |
| 2026-05-29T18:32:01Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0350 | E2 | edge 0.005 < mín 0.030 (p̂=0.030, implied=0.035) |
| 2026-05-29T18:32:01Z | will-claudia-lpez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $109278) |
| 2026-05-29T18:32:03Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $113210) |
| 2026-05-29T18:32:03Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $190467) |
| 2026-05-29T18:32:03Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $160122) |
| 2026-05-29T18:32:06Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E2 | edge 0.010 < mín 0.030 (p̂=0.010, implied=0.020) |
| 2026-05-29T18:32:06Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $1925) |
| 2026-05-29T18:32:06Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $1502) |
| 2026-05-29T18:32:06Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 25 d, liq $1334) |
| 2026-05-29T18:32:06Z | will-viggo-bjrck-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 28 d, liq $975) |
| 2026-05-29T18:32:06Z | will-tynan-lawrence-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $2741) |
| 2026-05-29T18:32:06Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $1290) |
| 2026-05-29T18:32:06Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $328905) |
| 2026-05-29T18:32:06Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 31 d, liq $99471) |
| 2026-05-29T18:32:06Z | ukraine-peace-referendum-scheduled-by-june-30 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $18235) |
| 2026-05-29T18:32:09Z | will-the-montreal-canadiens-win-the-2026-nhl-stanley-cup | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $236477) |
| 2026-05-29T18:32:09Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $347537) |
| 2026-05-29T18:32:09Z | will-amazon-be-the-largest-company-in-the-world-by-market-cap-on-june-30-123 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $342418) |
| 2026-05-29T18:32:11Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $11326) |
| 2026-05-29T18:32:14Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0250 | V3 Trigger vago + V6 Sin catalyst | V3 Trigger vago + V6 Sin catalyst: V3: El evento 'the Iranian regime fall' es vago: no hay una definición verificabl... |
| 2026-05-29T18:32:14Z | gustavo-petro-out-as-leader-of-colombia-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $15687) |
| 2026-05-29T18:32:16Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 31 d, liq $16710) |
| 2026-05-29T18:32:19Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-30-2026-15... | geopolitics | 0.8800 | V3 Trigger vago: sin fecha concreta o sin evento verificable | V3 Trigger vago: sin fecha concreta o sin evento verificable: evento especulativo sin fecha concreta ni fuente verifi... |
| 2026-05-29T18:32:27Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | Si no hay un evento discreto identificable en los próximos 7 días (como una filtración, reportaje, declaración of... | Si no hay un evento discreto identificable en los próximos 7 días (como una filtración, reportaje, declaración of... |
| 2026-05-29T18:32:29Z | will-deepseek-have-the-top-ai-model-at-the-end-of-june-2026-514 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 31 d, liq $8177) |
| 2026-05-29T18:32:32Z | will-russia-enter-kramatorsk-by-june-30-821-192 | geopolitics | 0.0280 | E2 | edge 0.002 < mín 0.030 (p̂=0.030, implied=0.028) |
| 2026-05-29T18:32:32Z | will-trump-pardon-tiger-woods-by-june-30 | executive-action | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $26020) |
| 2026-05-29T18:32:32Z | hantavirus-lab-leak-confirmed-by-june-30-1 | weather-natural | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 31 d, liq $71577) |
| 2026-05-29T18:32:34Z | nato-x-russia-military-clash-by-june-30-2026 | geopolitics | 0.0630 | E2 | edge 0.023 < mín 0.030 (p̂=0.040, implied=0.063) |
| 2026-05-29T18:32:36Z | will-us-withdraw-from-nato-by-june-30 | geopolitics | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $27821) |
| 2026-05-29T18:32:39Z | israel-x-iran-permanent-peace-deal-by-june-30-2026-262 | geopolitics | 0.1820 | V3 Trigger vago / V6 Sin catalyst | V3 Trigger vago / V6 Sin catalyst: V3: La pregunta es vaga - 'permanent peace deal' es ambiguo (¿un tratado formal, ... |
| 2026-05-29T18:32:41Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0340 | E2 | edge 0.016 < mín 0.030 (p̂=0.050, implied=0.034) |
| 2026-05-29T18:32:41Z | hamad-bin-isa-al-khalifa-out-as-bahrain-king | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 31 d, liq $13612) |
| 2026-05-29T18:32:46Z | starmer-out-by-june-30-2026-862-594-548-219-739 | executive-action | 0.1400 | E2 | edge 0.020 < mín 0.030 (p̂=0.120, implied=0.140) |
| 2026-05-29T18:32:46Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $31377) |
| 2026-05-29T18:32:50Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0370 | E2 | edge 0.027 < mín 0.030 (p̂=0.010, implied=0.037) |
| 2026-05-29T18:32:52Z | will-crude-oil-cl-hit-high-150-by-end-of-june-788-691 | market | 0.0240 | E2 | edge 0.019 < mín 0.030 (p̂=0.005, implied=0.024) |
| 2026-05-29T18:32:54Z | kharg-island-no-longer-under-iranian-control-by-june-30-561-854 | other | 0.0500 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda desencadena... |
| 2026-05-29T18:32:57Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.2710 | V3 | V3: V3 Trigger vago: el evento 'ganar las Finales NBA 2026' depende de múltiples factores impredecibles a más de un... |
| 2026-05-29T18:33:09Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P3_low_absolute_liquidity | liquidity $2450 < absolute min $5000 |
| 2026-05-29T18:33:09Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 43 d, liq $2303) |
| 2026-05-29T18:33:09Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0450 | P3_low_absolute_liquidity | liquidity $3393 < absolute min $5000 |
| 2026-05-29T18:33:09Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 43 d, liq $3176) |
| 2026-05-29T18:33:09Z | will-iga-witek-be-the-2026-womens-wimbledon-winner | other | 0.1890 | P3_low_absolute_liquidity | liquidity $2563 < absolute min $5000 |
| 2026-05-29T18:33:11Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $554 < absolute min $5000 |
| 2026-05-29T18:33:11Z | will-iraq-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10687104) |
| 2026-05-29T18:33:11Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $7528663) |
| 2026-05-29T18:33:14Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | E2 | edge 0.014 < mín 0.030 (p̂=0.080, implied=0.094) |
| 2026-05-29T18:33:14Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $10509994) |
| 2026-05-29T18:33:14Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8694377) |
| 2026-05-29T18:33:14Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $1128799) |
| 2026-05-29T18:33:14Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8235928) |
| 2026-05-29T18:33:14Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $6304129) |
| 2026-05-29T18:33:16Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E2 | edge 0.026 < mín 0.030 (p̂=0.060, implied=0.086) |
| 2026-05-29T18:33:16Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8089615) |
| 2026-05-29T18:33:16Z | will-qatar-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10404405) |
| 2026-05-29T18:33:16Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $9659927) |
| 2026-05-29T18:33:16Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $3035495) |
| 2026-05-29T18:33:19Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0290 | V6 - Sin catalyst | V6 - Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda alterar s... |
| 2026-05-29T18:33:19Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 51 d, liq $9244647) |
| 2026-05-29T18:33:19Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4511674) |
| 2026-05-29T18:33:19Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $8260112) |
| 2026-05-29T18:33:19Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $4760534) |
| 2026-05-29T18:33:19Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 51 d, liq $5358811) |
| 2026-05-29T18:33:19Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4568509) |
| 2026-05-29T18:33:26Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $7317980) |
| 2026-05-29T18:33:26Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $6313955) |
| 2026-05-29T18:33:26Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9357786) |
| 2026-05-29T18:33:26Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $2687381) |
| 2026-05-29T18:33:26Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 51 d, liq $423039) |
| 2026-05-29T18:33:28Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0390 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: el evento finaliza ... |
| 2026-05-29T18:33:28Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9467026) |
| 2026-05-29T18:33:28Z | will-jordan-win-the-2026-fifa-world-cup-233 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10312801) |
| 2026-05-29T18:33:28Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $7012961) |
| 2026-05-29T18:33:28Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $2655552) |
| 2026-05-29T18:33:28Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $4030655) |
| 2026-05-29T18:33:34Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.028 < mín 0.030 (p̂=0.080, implied=0.052) |
| 2026-05-29T18:33:34Z | will-new-zealand-win-the-2026-fifa-world-cup-635 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9897588) |
| 2026-05-29T18:33:34Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $5802903) |
| 2026-05-29T18:33:34Z | will-team-z-win-the-2026-fifa-world-cup-316 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10721936) |
| 2026-05-29T18:33:34Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $10681233) |
| 2026-05-29T18:33:34Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 51 d, liq $3507701) |
| 2026-05-29T18:33:34Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 51 d, liq $2357107) |
| 2026-05-29T18:33:34Z | will-cape-verde-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10381523) |
| 2026-05-29T18:33:34Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9802353) |
| 2026-05-29T18:33:34Z | will-curaao-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10025224) |
| 2026-05-29T18:33:34Z | will-haiti-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10535053) |
| 2026-05-29T18:33:37Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $1675738) |
| 2026-05-29T18:33:37Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $5554466) |
| 2026-05-29T18:33:37Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $15809) |
| 2026-05-29T18:33:37Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $32955) |
| 2026-05-29T18:33:37Z | will-paul-reevs-be-the-republican-nominee-for-az-01 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 66 d, liq $5971) |
| 2026-05-29T18:33:37Z | will-scott-jensen-win-the-2026-minnesota-governor-republican-primary-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 73 d, liq $11045) |
| 2026-05-29T18:33:37Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0330 | P3_low_absolute_liquidity | liquidity $892 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-david-njoku-play-for-denver-broncos-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 94 d, liq $1067) |
| 2026-05-29T18:33:37Z | tush-push-banned-for-2026-nfl-season | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 103 d, liq $1886) |
| 2026-05-29T18:33:37Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 106 d, liq $3497) |
| 2026-05-29T18:33:37Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2944) |
| 2026-05-29T18:33:37Z | will-barbora-krejcikova-win-the-2026-womens-us-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 106 d, liq $2400) |
| 2026-05-29T18:33:37Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $1888) |
| 2026-05-29T18:33:37Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $3354) |
| 2026-05-29T18:33:37Z | will-anastasia-potapova-win-the-2026-womens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2240) |
| 2026-05-29T18:33:37Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 106 d, liq $2740) |
| 2026-05-29T18:33:37Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 106 d, liq $2818) |
| 2026-05-29T18:33:37Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $975 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 121 d, liq $1914) |
| 2026-05-29T18:33:37Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 121 d, liq $1325) |
| 2026-05-29T18:33:37Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $468 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5600 | P8 | P8: election 123d out, price 0.56 en banda ruido [0.30, 0.70] |
| 2026-05-29T18:33:37Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $491 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-michelle-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 127 d, liq $214246) |
| 2026-05-29T18:33:37Z | will-ronaldo-caiado-win-the-2026-brazilian-presidential-election | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 127 d, liq $155302) |
| 2026-05-29T18:33:37Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 155 d, liq $27872) |
| 2026-05-29T18:33:37Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 155 d, liq $30585) |
| 2026-05-29T18:33:37Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 155 d, liq $27817) |
| 2026-05-29T18:33:37Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 155 d, liq $31508) |
| 2026-05-29T18:33:37Z | will-san-francisco-giants-win-the-2026-national-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 155 d, liq $26436) |
| 2026-05-29T18:33:37Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $160819) |
| 2026-05-29T18:33:37Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $6625) |
| 2026-05-29T18:33:37Z | will-corey-seager-win-the-2026-american-league-hank-aaron-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 167 d, liq $2464) |
| 2026-05-29T18:33:37Z | will-lew-evans-win-the-bachelorette-season-22 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 184 d, liq $6936) |
| 2026-05-29T18:33:37Z | will-sergio-prez-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 190 d, liq $752862) |
| 2026-05-29T18:33:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 193 d, liq $12237) |
| 2026-05-29T18:33:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $9197) |
| 2026-05-29T18:33:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $9979) |
| 2026-05-29T18:33:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 193 d, liq $10160) |
| 2026-05-29T18:33:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $9083) |
| 2026-05-29T18:33:37Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 193 d, liq $12478) |
| 2026-05-29T18:33:37Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 203 d, liq $4323) |
| 2026-05-29T18:33:37Z | will-michael-harris-ii-win-the-2026-nl-comeback-player-of-the-year-award | other | 0.1500 | P3_low_absolute_liquidity | liquidity $74 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-brice-matthews-win-the-2026-al-rookie-of-the-year-award | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 203 d, liq $7656) |
| 2026-05-29T18:33:37Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $280) |
| 2026-05-29T18:33:37Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $13053) |
| 2026-05-29T18:33:37Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1431 < absolute min $5000 |
| 2026-05-29T18:33:37Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $474) |
| 2026-05-29T18:33:37Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $139784) |
| 2026-05-29T18:33:37Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $437) |
| 2026-05-29T18:33:37Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $10121) |
| 2026-05-29T18:33:37Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $120 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2415 < absolute min $5000 |
| 2026-05-29T18:33:37Z | natural-disaster-in-2026 | other | 0.2800 | P3_low_absolute_liquidity | liquidity $2326 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4621 < absolute min $5000 |
| 2026-05-29T18:33:37Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0640 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T18:33:37Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $301) |
| 2026-05-29T18:33:37Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $12120) |
| 2026-05-29T18:33:37Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $7417) |
| 2026-05-29T18:33:37Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 215 d, liq $6314) |
| 2026-05-29T18:33:37Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 215 d, liq $205) |
| 2026-05-29T18:33:37Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3317 < absolute min $5000 |
| 2026-05-29T18:33:37Z | ledger-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1644 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $9063) |
| 2026-05-29T18:33:37Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $128) |
| 2026-05-29T18:33:37Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $288 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $7922) |
| 2026-05-29T18:33:37Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $557) |
| 2026-05-29T18:33:37Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 215 d, liq $271) |
| 2026-05-29T18:33:37Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $16070) |
| 2026-05-29T18:33:37Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $2833 < absolute min $5000 |
| 2026-05-29T18:33:37Z | ramp-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $2808 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $180 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $11814) |
| 2026-05-29T18:33:37Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $5961) |
| 2026-05-29T18:33:37Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4818 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0900 | P3_low_absolute_liquidity | liquidity $3337 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $453) |
| 2026-05-29T18:33:37Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $177) |
| 2026-05-29T18:33:37Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $3844 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4647 < absolute min $5000 |
| 2026-05-29T18:33:37Z | anduril-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $3241 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $11127) |
| 2026-05-29T18:33:37Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1200 | P3_low_absolute_liquidity | liquidity $3253 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6229) |
| 2026-05-29T18:33:37Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $889) |
| 2026-05-29T18:33:37Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $95425) |
| 2026-05-29T18:33:37Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $62464) |
| 2026-05-29T18:33:37Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $159092) |
| 2026-05-29T18:33:37Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $226) |
| 2026-05-29T18:33:37Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6993) |
| 2026-05-29T18:33:37Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | sports-season | 0.0910 | P3_low_absolute_liquidity | liquidity $842 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $2956) |
| 2026-05-29T18:33:37Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $1350 < absolute min $5000 |
| 2026-05-29T18:33:37Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3266 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-consensys-ipo-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $1298 < absolute min $5000 |
| 2026-05-29T18:33:37Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.2560 | P3_low_absolute_liquidity | liquidity $2370 < absolute min $5000 |
| 2026-05-29T18:33:37Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4100 | P3_low_absolute_liquidity | liquidity $1033 < absolute min $5000 |
| 2026-05-29T18:33:37Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $1044 < absolute min $5000 |
| 2026-05-29T18:33:37Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12375) |
| 2026-05-29T18:33:37Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $642 < absolute min $5000 |
| 2026-05-29T18:33:37Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $463 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-ethereum-dip-to-800-by-december-31-2026 | market | 0.2100 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.21) |
| 2026-05-29T18:33:37Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2318 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $194 < absolute min $5000 |
| 2026-05-29T18:33:37Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.21) |
| 2026-05-29T18:33:37Z | will-usdt-market-cap-hit-200b-before-2027 | other | 0.9780 | P3_low_absolute_liquidity | liquidity $937 < absolute min $5000 |
| 2026-05-29T18:33:37Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.21) |
| 2026-05-29T18:33:37Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.5000 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45532) |
| 2026-05-29T18:33:37Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $18866) |
| 2026-05-29T18:33:37Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52833) |
| 2026-05-29T18:33:37Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 240 d, liq $51907) |
| 2026-05-29T18:33:37Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $113947) |
| 2026-05-29T18:33:37Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $251207) |
| 2026-05-29T18:33:37Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $23521) |
| 2026-05-29T18:33:37Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7300 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8300 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7700 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45530) |
| 2026-05-29T18:33:37Z | will-spacexs-market-cap-be-between-600b-and-700b-at-market-close-on-ipo-day | crypto-launch | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 580 d, liq $12949) |
| 2026-05-29T18:33:37Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T18:33:37Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2368718) |
| 2026-05-29T18:33:37Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1762172) |
| 2026-05-29T18:33:37Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1915398) |
| 2026-05-29T18:33:37Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1082709) |
| 2026-05-29T18:33:37Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $771483) |
| 2026-05-29T18:33:37Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1935881) |
| 2026-05-29T18:33:37Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1924328) |
| 2026-05-29T18:33:37Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $288081) |
| 2026-05-29T18:33:37Z | will-ro-khanna-win-the-2028-democratic-presidential-nomination | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $251522) |
| 2026-05-29T18:33:37Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2144787) |
| 2026-05-29T18:33:37Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $12383) |
| 2026-05-29T18:33:37Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $2250 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T18:33:37Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $499 < absolute min $5000 |
| 2026-05-29T18:33:37Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22631) |
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

