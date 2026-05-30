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
| 2026-05-30T12:03:35Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $10935) |
| 2026-05-30T12:03:35Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $127) |
| 2026-05-30T12:03:35Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $289 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $639) |
| 2026-05-30T12:03:35Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2496 < absolute min $5000 |
| 2026-05-30T12:03:35Z | jerome-powell-out-from-fed-board-by-may-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $55271) |
| 2026-05-30T12:03:35Z | 5kt-meteor-strike-in-2026 | weather-natural | 0.3200 | P3_low_absolute_liquidity | liquidity $3067 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $7927) |
| 2026-05-30T12:03:35Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $505) |
| 2026-05-30T12:03:35Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $544) |
| 2026-05-30T12:03:35Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T12:03:35Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $6077) |
| 2026-05-30T12:03:35Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $390) |
| 2026-05-30T12:03:35Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $6246) |
| 2026-05-30T12:03:35Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1226 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $700) |
| 2026-05-30T12:03:35Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4384 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $16466) |
| 2026-05-30T12:03:35Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 214 d, liq $9619) |
| 2026-05-30T12:03:35Z | will-there-be-between-11-and-13-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $1720 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4902 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $61735) |
| 2026-05-30T12:03:35Z | will-there-be-between-14-and-16-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2386 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $839) |
| 2026-05-30T12:03:35Z | major-meteor-strike-10kt-in-2026 | weather-natural | 0.1300 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T12:03:35Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $196) |
| 2026-05-30T12:03:35Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $13344) |
| 2026-05-30T12:03:35Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $6498) |
| 2026-05-30T12:03:35Z | databricks-ipo-before-2027 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $2529 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $210 < absolute min $5000 |
| 2026-05-30T12:03:35Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3350 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $12592) |
| 2026-05-30T12:03:35Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $18681) |
| 2026-05-30T12:03:35Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1230 | P3_low_absolute_liquidity | liquidity $3179 < absolute min $5000 |
| 2026-05-30T12:03:35Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1425 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $273) |
| 2026-05-30T12:03:35Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $174378) |
| 2026-05-30T12:03:35Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2833 < absolute min $5000 |
| 2026-05-30T12:03:35Z | altcoin-market-cap-dip-to-150b-before-2027 | other | 0.7300 | P3_low_absolute_liquidity | liquidity $695 < absolute min $5000 |
| 2026-05-30T12:03:35Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3282 < absolute min $5000 |
| 2026-05-30T12:03:35Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $12642) |
| 2026-05-30T12:03:35Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3300 | P3_low_absolute_liquidity | liquidity $583 < absolute min $5000 |
| 2026-05-30T12:03:35Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1300 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3800 | P3_low_absolute_liquidity | liquidity $1114 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $654 < absolute min $5000 |
| 2026-05-30T12:03:35Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0910 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P3_low_absolute_liquidity | liquidity $770 < absolute min $5000 |
| 2026-05-30T12:03:35Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | opensea-fdv-above-5b-one-day-after-launch-919-666-825-992-543-333 | crypto-launch | 0.0500 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | will-stablecoins-hit-500b-before-2027 | other | 0.1200 | P3_low_absolute_liquidity | liquidity $3353 < absolute min $5000 |
| 2026-05-30T12:03:35Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | opensea-fdv-above-100m-one-day-after-launch-172-151-588-987 | crypto-launch | 0.6000 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $1454) |
| 2026-05-30T12:03:35Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2388 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 239 d, liq $45937) |
| 2026-05-30T12:03:35Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 239 d, liq $19260) |
| 2026-05-30T12:03:35Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $52849) |
| 2026-05-30T12:03:35Z | will-the-atlanta-falcons-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 305 d, liq $193320) |
| 2026-05-30T12:03:35Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $261373) |
| 2026-05-30T12:03:35Z | will-spacex-not-ipo-by-december-31-2027-481-338 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 579 d, liq $11937) |
| 2026-05-30T12:03:35Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $46112) |
| 2026-05-30T12:03:35Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3981 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $43492) |
| 2026-05-30T12:03:35Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8200 | P4_pre_event | pre-event slug + 579 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2180 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.2400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T12:03:35Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1061831) |
| 2026-05-30T12:03:35Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1483146) |
| 2026-05-30T12:03:35Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1443228) |
| 2026-05-30T12:03:35Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1723916) |
| 2026-05-30T12:03:35Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2333807) |
| 2026-05-30T12:03:35Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $468059) |
| 2026-05-30T12:03:35Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1456922) |
| 2026-05-30T12:03:35Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1626557) |
| 2026-05-30T12:03:35Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1170599) |
| 2026-05-30T12:03:35Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 891 d, liq $1914772) |
| 2026-05-30T12:03:35Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $936888) |
| 2026-05-30T12:03:35Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1301708) |
| 2026-05-30T12:03:35Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2437882) |
| 2026-05-30T12:03:35Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2365093) |
| 2026-05-30T12:03:35Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $940080) |
| 2026-05-30T12:03:35Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1151058) |
| 2026-05-30T12:03:35Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1075037) |
| 2026-05-30T12:03:35Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $1208) |
| 2026-05-30T12:03:35Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $545 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0400 | P3_low_absolute_liquidity | liquidity $2729 < absolute min $5000 |
| 2026-05-30T12:03:35Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12539) |
| 2026-05-30T12:03:35Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $5017) |
| 2026-05-30T12:03:35Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22950) |
| 2026-05-30T12:30:04Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon -150 d, liq $48552) |
| 2026-05-30T12:30:04Z | ukraine-recognizes-russian-sovereignty-over-its-territory-by-june-30-2026 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -150 d, liq $13490) |
| 2026-05-30T12:30:04Z | new-stranger-things-episode-released-by-december-31-587-657-824-246 | other | 0.0700 | P2 | mercado ya expiró (endDate=2026-01-07T00:00:00Z, hace 143 días) |
| 2026-05-30T12:30:04Z | will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -111 d, liq $92995) |
| 2026-05-30T12:30:04Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $776100) |
| 2026-05-30T12:30:04Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $766685) |
| 2026-05-30T12:30:04Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -69 d, liq $32360) |
| 2026-05-30T12:30:04Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $20147) |
| 2026-05-30T12:30:04Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -67 d, liq $28210) |
| 2026-05-30T12:30:04Z | will-alex-vanopslagh-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $29615) |
| 2026-05-30T12:30:04Z | will-troels-lund-poulsen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon -67 d, liq $20038) |
| 2026-05-30T12:30:04Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -67 d, liq $28025) |
| 2026-05-30T12:30:04Z | weed-rescheduled-by-june-30 | other | 0.0210 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T12:30:04Z | weed-rescheduled-by-december-31 | other | 0.3020 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T12:30:04Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -48 d, liq $201183) |
| 2026-05-30T12:30:04Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $345156) |
| 2026-05-30T12:30:04Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.6500 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 15 días) |
| 2026-05-30T12:30:04Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon -15 d, liq $104564) |
| 2026-05-30T12:30:04Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -13 d, liq $5971) |
| 2026-05-30T12:30:04Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -11 d, liq $8561) |
| 2026-05-30T12:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -3 d, liq $1750120) |
| 2026-05-30T12:30:04Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $14597) |
| 2026-05-30T12:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.0600 | P9 | P9: geopolitics pump cluster (price 0.06, 0d) |
| 2026-05-30T12:30:04Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $14553) |
| 2026-05-30T12:30:04Z | elon-musk-of-tweets-may-28-may-30-90-114 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $106172) |
| 2026-05-30T12:30:04Z | bitcoin-above-70k-on-may-30-2026 | market | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-30T12:30:04Z | elon-musk-of-tweets-may-28-may-30-0-39 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $75108) |
| 2026-05-30T12:30:04Z | bitcoin-above-72k-on-may-30-2026 | market | 0.9940 | P0_ceiling | price ceiling: 0.9940 > 0.980 |
| 2026-05-30T12:30:06Z | ucl-psg-ars-2026-05-30-btts | other | 0.5200 | V3 Trigger vago | V3 Trigger vago: V3 Trigger vago: el evento 'Both Teams to Score' en un partido de fútbol a 2 años vista no tiene u... |
| 2026-05-30T12:30:06Z | bitcoin-above-74k-on-may-30-2026 | market | 0.0600 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:30:11Z | elon-musk-of-tweets-may-28-may-30-65-89 | other | 0.1100 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda desencadena... |
| 2026-05-30T12:30:14Z | elon-musk-of-tweets-may-28-may-30-40-64 | other | 0.9100 | E2 | edge 0.020 < mín 0.030 (p̂=0.930, implied=0.910) |
| 2026-05-30T12:30:17Z | bitcoin-above-76k-on-may-30-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $41388) |
| 2026-05-30T12:30:17Z | will-80-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 0 d, liq $71766) |
| 2026-05-30T12:30:17Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 0 d, liq $106429) |
| 2026-05-30T12:30:18Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0460 | E2 | edge 0.026 < mín 0.030 (p̂=0.020, implied=0.046) |
| 2026-05-30T12:30:18Z | will-alibaba-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $136072) |
| 2026-05-30T12:30:18Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $220568) |
| 2026-05-30T12:30:18Z | uefa-champions-league-unbeaten-team | sports-season | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-30T12:30:21Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $200461) |
| 2026-05-30T12:30:25Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0240 | E2 | edge 0.014 < mín 0.030 (p̂=0.010, implied=0.024) |
| 2026-05-30T12:30:25Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 0 d, liq $63949) |
| 2026-05-30T12:30:25Z | will-60-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 0 d, liq $66848) |
| 2026-05-30T12:30:28Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.1500 | P9 | P9: geopolitics pump cluster (price 0.15, 0d) |
| 2026-05-30T12:30:30Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.0470 | E2 | edge 0.027 < mín 0.030 (p̂=0.020, implied=0.047) |
| 2026-05-30T12:30:30Z | mojtaba-khamenei-leaves-iran-by-may-31-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $252365) |
| 2026-05-30T12:30:30Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $162294) |
| 2026-05-30T12:30:30Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $198007) |
| 2026-05-30T12:30:30Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 0 d, liq $857562) |
| 2026-05-30T12:30:30Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $3895115) |
| 2026-05-30T12:30:34Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0270 | E2 | edge 0.012 < mín 0.030 (p̂=0.015, implied=0.027) |
| 2026-05-30T12:30:37Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.0370 | E2 | edge 0.013 < mín 0.030 (p̂=0.050, implied=0.037) |
| 2026-05-30T12:30:37Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $186921) |
| 2026-05-30T12:30:37Z | will-reza-pahlavi-enter-iran-by-may-31 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $185589) |
| 2026-05-30T12:30:37Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 0 d, liq $204350) |
| 2026-05-30T12:30:39Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0340 | V3 Trigger vago / V6 Sin catalyst | V3 Trigger vago / V6 Sin catalyst: V3 Trigger vago: 'permanent peace deal' es una noción sin fecha concreta ni event... |
| 2026-05-30T12:30:40Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0240 | E2 | edge 0.006 < mín 0.030 (p̂=0.030, implied=0.024) |
| 2026-05-30T12:30:42Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 0 d, liq $96776) |
| 2026-05-30T12:30:42Z | will-crude-oil-reach-a-new-all-time-high-by-may-31 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 0 d, liq $33415) |
| 2026-05-30T12:30:42Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 0 d, liq $66077) |
| 2026-05-30T12:30:45Z | will-40-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0350 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T12:30:47Z | will-psg-win-the-202526-champions-league | sports-season | 0.5700 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: The market end date is 2026-05-31, m... |
| 2026-05-30T12:30:49Z | will-arsenal-win-the-202526-champions-league | sports-season | 0.4500 | V3 - Trigger vago | V3 - Trigger vago: V3 Trigger vago: falta una fecha concreta o evento verificable cercano. La pregunta es a largo pla... |
| 2026-05-30T12:30:49Z | will-nvidia-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $37954) |
| 2026-05-30T12:30:49Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $97272) |
| 2026-05-30T12:30:49Z | will-alphabet-be-the-third-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $18499) |
| 2026-05-30T12:30:49Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $155258) |
| 2026-05-30T12:30:49Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $145774) |
| 2026-05-30T12:30:49Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116098) |
| 2026-05-30T12:30:49Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $98849) |
| 2026-05-30T12:30:49Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $143020) |
| 2026-05-30T12:30:49Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $145293) |
| 2026-05-30T12:30:49Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $100931) |
| 2026-05-30T12:30:49Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $95244) |
| 2026-05-30T12:30:49Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116773) |
| 2026-05-30T12:30:49Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $123803) |
| 2026-05-30T12:30:49Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $148187) |
| 2026-05-30T12:30:49Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $155465) |
| 2026-05-30T12:30:49Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 1 d, liq $306249) |
| 2026-05-30T12:30:49Z | will-wti-dip-to-70-in-may-2026-155-395-333-182-889-988-959-341-883-234-375-347-135 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $211340) |
| 2026-05-30T12:30:49Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $555743) |
| 2026-05-30T12:30:49Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $739049) |
| 2026-05-30T12:30:49Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $1269700) |
| 2026-05-30T12:30:49Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $263930) |
| 2026-05-30T12:30:49Z | over-60m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $13085) |
| 2026-05-30T12:30:49Z | will-ethereum-reach-2600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $59447) |
| 2026-05-30T12:30:49Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $27296) |
| 2026-05-30T12:30:49Z | will-bitcoin-dip-to-72k-may-25-31-2026 | market | 0.1200 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:30:49Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 1 d, liq $34929) |
| 2026-05-30T12:30:49Z | over-30m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $13415) |
| 2026-05-30T12:30:49Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $164685) |
| 2026-05-30T12:30:49Z | over-20m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $10012) |
| 2026-05-30T12:30:49Z | will-bitcoin-reach-77500-in-may-2026-from-may-27 | market | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 1 d, liq $26240) |
| 2026-05-30T12:30:49Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $235051) |
| 2026-05-30T12:30:49Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $135690) |
| 2026-05-30T12:30:49Z | will-ethereum-dip-to-1600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $42819) |
| 2026-05-30T12:30:49Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0330 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:30:49Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $12263) |
| 2026-05-30T12:30:49Z | will-bitcoin-reach-80000-in-may-2026-from-may-27 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 1 d, liq $35948) |
| 2026-05-30T12:30:49Z | will-ethereum-dip-to-400-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $43043) |
| 2026-05-30T12:30:49Z | will-bitcoin-dip-to-72500-in-may-2026-from-may-27 | market | 0.2600 | P11 | tradingview BTCUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:30:49Z | will-xrp-reach-1pt6-in-may-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $13306) |
| 2026-05-30T12:30:51Z | will-backrooms-opening-weekend-box-office-be-greater-than-79m | entertainment | 0.8520 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: 'Backrooms' no es una película con f... |
| 2026-05-30T12:30:53Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2200 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: La pregunta usa el nombre 'Spencer P... |
| 2026-05-30T12:30:53Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $17702) |
| 2026-05-30T12:30:53Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $38572) |
| 2026-05-30T12:30:53Z | elon-musk-of-tweets-may-26-june-2-320-339 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 3 d, liq $41108) |
| 2026-05-30T12:30:53Z | elon-musk-of-tweets-may-26-june-2-100-119 | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 3 d, liq $55921) |
| 2026-05-30T12:30:53Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $41998) |
| 2026-05-30T12:30:53Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $95387) |
| 2026-05-30T12:30:53Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $45449) |
| 2026-05-30T12:30:53Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $31772) |
| 2026-05-30T12:30:53Z | will-na-kyung-won-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $479449) |
| 2026-05-30T12:30:53Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $44816) |
| 2026-05-30T12:30:53Z | will-park-hong-keun-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $481750) |
| 2026-05-30T12:30:58Z | will-chung-jin-suk-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $19901) |
| 2026-05-30T12:30:58Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $39811) |
| 2026-05-30T12:30:58Z | will-park-yong-jin-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $462606) |
| 2026-05-30T12:30:58Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $32924) |
| 2026-05-30T12:30:58Z | will-kang-seung-kyu-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $23926) |
| 2026-05-30T12:30:58Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $42682) |
| 2026-05-30T12:30:58Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $31325) |
| 2026-05-30T12:30:58Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.2700 | P9 | P9: geopolitics pump cluster (price 0.27, 3d) |
| 2026-05-30T12:30:58Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $26502) |
| 2026-05-30T12:30:58Z | will-yoon-sang-hyun-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $38394) |
| 2026-05-30T12:30:58Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $22008) |
| 2026-05-30T12:30:58Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $97587) |
| 2026-05-30T12:30:58Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $41068) |
| 2026-05-30T12:30:58Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $23724) |
| 2026-05-30T12:30:58Z | will-han-dong-hoon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $33994) |
| 2026-05-30T12:30:58Z | elon-musk-of-tweets-may-29-june-5-500plus | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 6 d, liq $17494) |
| 2026-05-30T12:30:58Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $230489) |
| 2026-05-30T12:30:58Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $226259) |
| 2026-05-30T12:30:58Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $210073) |
| 2026-05-30T12:30:58Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $123727) |
| 2026-05-30T12:30:58Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $213342) |
| 2026-05-30T12:30:58Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $112241) |
| 2026-05-30T12:30:58Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $220490) |
| 2026-05-30T12:31:00Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $247565) |
| 2026-05-30T12:31:00Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $189915) |
| 2026-05-30T12:31:02Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0700 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: sin fecha concreta o sin evento verif... |
| 2026-05-30T12:31:02Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.1400 | P9 | P9: geopolitics pump cluster (price 0.14, 7d) |
| 2026-05-30T12:31:05Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0290 | E2 | edge 0.024 < mín 0.030 (p̂=0.005, implied=0.029) |
| 2026-05-30T12:31:05Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $160639) |
| 2026-05-30T12:31:05Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $201461) |
| 2026-05-30T12:31:08Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $222170) |
| 2026-05-30T12:31:10Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0330 | E2 | edge 0.028 < mín 0.030 (p̂=0.005, implied=0.033) |
| 2026-05-30T12:31:13Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0210 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda cambiar sig... |
| 2026-05-30T12:31:20Z | will-moise-kouame-win-the-2026-mens-french-open | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 7 d, liq $25304) |
| 2026-05-30T12:31:22Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0900 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: No hay evento discreto identificable en los próximos 7 días que pueda afectar sig... |
| 2026-05-30T12:31:22Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $212377) |
| 2026-05-30T12:31:22Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $91119) |
| 2026-05-30T12:31:22Z | will-martin-landaluce-win-the-2026-mens-french-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 7 d, liq $39330) |
| 2026-05-30T12:31:22Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $198361) |
| 2026-05-30T12:31:22Z | will-jakub-mensik-win-the-2026-mens-french-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 7 d, liq $54975) |
| 2026-05-30T12:31:25Z | will-raphael-collignon-win-the-2026-mens-french-open | other | 0.0360 | V3 Trigger vago | V3 Trigger vago: V3 - Trigger vago: el evento 'ganar el Abierto de Francia 2026' es extremadamente específico y leja... |
| 2026-05-30T12:31:28Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $159949) |
| 2026-05-30T12:31:30Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0210 | V3 | V3: V3 Trigger vago: sin fecha concreta o evento verificable. No hay ningún catalizador identificable cercano para u... |
| 2026-05-30T12:31:32Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1120 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: El mercado pregunta por un tenista e... |
| 2026-05-30T12:31:32Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $227148) |
| 2026-05-30T12:31:37Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.2200 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T12:31:39Z | will-al-mina-be-the-republican-nominee-for-senate-in-virginia | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 16 d, liq $21821) |
| 2026-05-30T12:31:41Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 17 d, liq $801744) |
| 2026-05-30T12:31:41Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 17 d, liq $330835) |
| 2026-05-30T12:31:41Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 17 d, liq $794589) |
| 2026-05-30T12:31:41Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9820 | P0_ceiling | price ceiling: 0.9820 > 0.980 |
| 2026-05-30T12:31:41Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 17 d, liq $2640133) |
| 2026-05-30T12:31:41Z | will-lynn-vision-win-iem-cologne-major-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 21 d, liq $19500) |
| 2026-05-30T12:31:41Z | will-gaimin-gladiators-win-iem-cologne-major-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $82239) |
| 2026-05-30T12:31:41Z | will-sinners-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $42560) |
| 2026-05-30T12:31:41Z | will-nrg-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $74513) |
| 2026-05-30T12:31:41Z | will-flyquest-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $70530) |
| 2026-05-30T12:31:41Z | will-m80-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $45241) |
| 2026-05-30T12:31:41Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $200049) |
| 2026-05-30T12:31:41Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $226038) |
| 2026-05-30T12:31:41Z | will-juan-daniel-oviedo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $257810) |
| 2026-05-30T12:31:41Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 22 d, liq $62605) |
| 2026-05-30T12:31:41Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $215212) |
| 2026-05-30T12:31:41Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $156925) |
| 2026-05-30T12:31:41Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $203579) |
| 2026-05-30T12:31:41Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $184780) |
| 2026-05-30T12:31:43Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $195675) |
| 2026-05-30T12:31:43Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $233864) |
| 2026-05-30T12:31:49Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0260 | E2 | edge 0.024 < mín 0.030 (p̂=0.050, implied=0.026) |
| 2026-05-30T12:31:51Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E2 | edge 0.010 < mín 0.030 (p̂=0.010, implied=0.020) |
| 2026-05-30T12:31:51Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 24 d, liq $687) |
| 2026-05-30T12:31:51Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 24 d, liq $510) |
| 2026-05-30T12:31:51Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 24 d, liq $1115) |
| 2026-05-30T12:31:51Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 27 d, liq $1706) |
| 2026-05-30T12:31:51Z | hantavirus-lab-leak-confirmed-by-june-30-1 | weather-natural | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 30 d, liq $121459) |
| 2026-05-30T12:31:55Z | mojtaba-khamenei-leaves-iran-by-june-30-2026 | geopolitics | 0.0270 | E2 | edge 0.023 < mín 0.030 (p̂=0.050, implied=0.027) |
| 2026-05-30T12:31:55Z | will-moonshot-have-the-top-ai-model-at-the-end-of-june-2026 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $11299) |
| 2026-05-30T12:31:55Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 30 d, liq $12669) |
| 2026-05-30T12:31:57Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0220 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: 'Iranian regime fall' no está defini... |
| 2026-05-30T12:32:00Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-june-30-2026 | geopolitics | 0.1800 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: No cumple con el catalizador: el eve... |
| 2026-05-30T12:32:00Z | anthropic-ceo-arrested | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $8190) |
| 2026-05-30T12:32:00Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $349096) |
| 2026-05-30T12:32:00Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-june-30-612 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $247446) |
| 2026-05-30T12:32:02Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0340 | E2 | edge 0.016 < mín 0.030 (p̂=0.050, implied=0.034) |
| 2026-05-30T12:32:02Z | israel-withdraws-from-lebanon-by-may-31-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 30 d, liq $42555) |
| 2026-05-30T12:32:04Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V5 Patrón débil: <3 fuentes independientes o sin precedente análogo. V6 Sin catalyst: no hay evento discreto ident... | V5 Patrón débil: <3 fuentes independientes o sin precedente análogo. V6 Sin catalyst: no hay evento discreto ident... |
| 2026-05-30T12:32:09Z | masoud-pezeshkian-out-by-june-30-242-112 | executive-action | 0.0530 | E2 | edge 0.013 < mín 0.030 (p̂=0.040, implied=0.053) |
| 2026-05-30T12:32:09Z | will-trump-and-putin-meet-next-in-ukraine-328-474 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $17379) |
| 2026-05-30T12:32:15Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0320 | E2 | edge 0.018 < mín 0.030 (p̂=0.050, implied=0.032) |
| 2026-05-30T12:32:15Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $373420) |
| 2026-05-30T12:32:20Z | will-meituan-have-the-top-ai-model-at-the-end-of-june-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 30 d, liq $11042) |
| 2026-05-30T12:32:20Z | will-trump-and-putin-meet-next-in-belarus-572 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $18280) |
| 2026-05-30T12:32:22Z | will-tim-walz-resign-by-june-30 | executive-action | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $26412) |
| 2026-05-30T12:32:22Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $16253) |
| 2026-05-30T12:32:25Z | will-russia-enter-kramatorsk-by-june-30-821-192 | geopolitics | 0.0280 | E2 | edge 0.022 < mín 0.030 (p̂=0.050, implied=0.028) |
| 2026-05-30T12:32:25Z | will-trump-and-putin-meet-next-in-south-korea-582-297-268 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $22206) |
| 2026-05-30T12:32:27Z | will-openais-market-cap-be-less-than-500b-at-market-close-on-ipo-day | crypto-launch | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 30 d, liq $6239) |
| 2026-05-30T12:32:27Z | will-the-montreal-canadiens-win-the-eastern-conference | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $384825) |
| 2026-05-30T12:32:30Z | will-deepseek-have-the-second-best-ai-model-at-the-end-of-june-2026 | other | 0.0230 | E2 | edge 0.021 < mín 0.030 (p̂=0.002, implied=0.023) |
| 2026-05-30T12:32:32Z | jerome-powell-federally-charged-by-june-30 | other | 0.0380 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: la pregunta es extremadamente vaga ('... |
| 2026-05-30T12:32:34Z | trump-renames-ice-to-nice-by-june-30 | other | 0.0410 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. |
| 2026-05-30T12:32:36Z | european-country-agrees-to-give-ukraine-security-guarantee-by-june-30 | geopolitics | 0.0400 | E2 | edge 0.020 < mín 0.030 (p̂=0.060, implied=0.040) |
| 2026-05-30T12:32:36Z | will-any-country-leave-nato-by-june-30-2026 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 30 d, liq $23620) |
| 2026-05-30T12:32:36Z | gc-hit-7000-high-jun-2026-433-244-291 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 31 d, liq $28531) |
| 2026-05-30T12:32:36Z | will-gold-gc-hit-high-6500-by-end-of-june | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 31 d, liq $23278) |
| 2026-05-30T12:32:36Z | will-silver-si-hit-low-35-by-end-of-june-325-231 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 31 d, liq $20172) |
| 2026-05-30T12:32:36Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $31190) |
| 2026-05-30T12:32:36Z | will-silver-si-hit-high-150-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $37250) |
| 2026-05-30T12:32:36Z | will-silver-si-hit-high-130-by-end-of-june | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $34612) |
| 2026-05-30T12:32:36Z | will-crude-oil-cl-hit-low-50-by-end-of-june | market | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 31 d, liq $29000) |
| 2026-05-30T12:32:49Z | ethereum-all-time-high-by-june-30-2026 | market | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 31 d, liq $18812) |
| 2026-05-30T12:32:49Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $626 < absolute min $5000 |
| 2026-05-30T12:32:49Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8069382) |
| 2026-05-30T12:32:49Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8494569) |
| 2026-05-30T12:32:51Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0270 | E2 | edge 0.022 < mín 0.030 (p̂=0.005, implied=0.027) |
| 2026-05-30T12:32:51Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 50 d, liq $3178613) |
| 2026-05-30T12:32:51Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8617575) |
| 2026-05-30T12:32:51Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $9643903) |
| 2026-05-30T12:32:53Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0370 | E2 | edge 0.004 < mín 0.030 (p̂=0.033, implied=0.037) |
| 2026-05-30T12:32:53Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $1156220) |
| 2026-05-30T12:32:56Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 50 d, liq $9651435) |
| 2026-05-30T12:32:56Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $7401574) |
| 2026-05-30T12:32:56Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 50 d, liq $2526377) |
| 2026-05-30T12:32:59Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 50 d, liq $3745267) |
| 2026-05-30T12:32:59Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 50 d, liq $5004190) |
| 2026-05-30T12:32:59Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $10766795) |
| 2026-05-30T12:32:59Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 50 d, liq $5937744) |
| 2026-05-30T12:32:59Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $9628197) |
| 2026-05-30T12:32:59Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 50 d, liq $408204) |
| 2026-05-30T12:32:59Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 50 d, liq $2492520) |
| 2026-05-30T12:32:59Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $10950738) |
| 2026-05-30T12:32:59Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $4230827) |
| 2026-05-30T12:32:59Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $2458924) |
| 2026-05-30T12:32:59Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 50 d, liq $8418518) |
| 2026-05-30T12:33:01Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0910 | E2 | edge 0.009 < mín 0.030 (p̂=0.100, implied=0.091) |
| 2026-05-30T12:33:01Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 50 d, liq $5460686) |
| 2026-05-30T12:33:01Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $5272983) |
| 2026-05-30T12:33:04Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E2 | edge 0.004 < mín 0.030 (p̂=0.090, implied=0.086) |
| 2026-05-30T12:33:06Z | will-spain-win-the-2026-fifa-world-cup-963 | sports-season | 0.1680 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T12:33:06Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $7253734) |
| 2026-05-30T12:33:09Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $10050060) |
| 2026-05-30T12:33:12Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.018 < mín 0.030 (p̂=0.070, implied=0.052) |
| 2026-05-30T12:33:12Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $2783644) |
| 2026-05-30T12:33:12Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $7562858) |
| 2026-05-30T12:33:12Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $5765505) |
| 2026-05-30T12:33:12Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $6637632) |
| 2026-05-30T12:33:12Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $4507757) |
| 2026-05-30T12:33:12Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $9611469) |
| 2026-05-30T12:33:12Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $7797026) |
| 2026-05-30T12:33:12Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $16030) |
| 2026-05-30T12:33:12Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $35594) |
| 2026-05-30T12:33:16Z | will-tyreek-hill-play-for-philadelphia-eagles-next | other | 0.1480 | P3_low_absolute_liquidity | liquidity $247 < absolute min $5000 |
| 2026-05-30T12:33:16Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0240 | P3_low_absolute_liquidity | liquidity $1014 < absolute min $5000 |
| 2026-05-30T12:33:16Z | will-david-njoku-play-for-kansas-city-chiefs-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $2607) |
| 2026-05-30T12:33:16Z | will-david-njoku-play-for-indianapolis-colts-in-2026-27 | other | 0.0200 | P3_low_absolute_liquidity | liquidity $522 < absolute min $5000 |
| 2026-05-30T12:33:16Z | will-david-njoku-play-for-denver-broncos-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $939) |
| 2026-05-30T12:33:16Z | will-filipe-luis-be-the-next-manager-of-ssc-napoli | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 93 d, liq $1487) |
| 2026-05-30T12:33:16Z | will-david-njoku-play-for-miami-dolphins-in-2026-27 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 93 d, liq $2826) |
| 2026-05-30T12:33:19Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $2467) |
| 2026-05-30T12:33:19Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 105 d, liq $2317) |
| 2026-05-30T12:33:19Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $1981) |
| 2026-05-30T12:33:19Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 105 d, liq $3177) |
| 2026-05-30T12:33:19Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 120 d, liq $1118) |
| 2026-05-30T12:33:19Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 120 d, liq $761) |
| 2026-05-30T12:33:19Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $787 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $290 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5800 | P8 | P8: election 122d out, price 0.58 en banda ruido [0.30, 0.70] |
| 2026-05-30T12:33:19Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $484 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-ronaldo-caiado-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 126 d, liq $28135) |
| 2026-05-30T12:33:19Z | will-carlos-roberto-massa-jnior-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 126 d, liq $28346) |
| 2026-05-30T12:33:19Z | will-partido-liberal-pl-win-the-most-seats-in-the-next-brazilian-senate-election | elections | 0.7200 | P3_low_absolute_liquidity | liquidity $2286 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 126 d, liq $1080793) |
| 2026-05-30T12:33:19Z | will-the-miami-marlins-win-the-2026-world-series | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 154 d, liq $108375) |
| 2026-05-30T12:33:19Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 154 d, liq $27428) |
| 2026-05-30T12:33:19Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 154 d, liq $29980) |
| 2026-05-30T12:33:19Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 154 d, liq $31209) |
| 2026-05-30T12:33:19Z | will-st-louis-cardinals-win-the-2026-national-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 154 d, liq $23612) |
| 2026-05-30T12:33:19Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 154 d, liq $28035) |
| 2026-05-30T12:33:19Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 156 d, liq $7601) |
| 2026-05-30T12:33:19Z | will-jakob-glesnes-win-2026-mls-defender-of-the-year | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 165 d, liq $283) |
| 2026-05-30T12:33:19Z | will-david-brekalo-win-2026-mls-defender-of-the-year | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 165 d, liq $270) |
| 2026-05-30T12:33:19Z | will-richard-van-de-water-win-the-bachelorette-season-22 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 183 d, liq $3822) |
| 2026-05-30T12:33:19Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 192 d, liq $12077) |
| 2026-05-30T12:33:19Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9267) |
| 2026-05-30T12:33:19Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9970) |
| 2026-05-30T12:33:19Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $8899) |
| 2026-05-30T12:33:19Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 192 d, liq $12400) |
| 2026-05-30T12:33:19Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 202 d, liq $6442) |
| 2026-05-30T12:33:19Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4298 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $11360) |
| 2026-05-30T12:33:19Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $1048) |
| 2026-05-30T12:33:19Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1230 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $69805) |
| 2026-05-30T12:33:19Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $605) |
| 2026-05-30T12:33:19Z | will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026 | entertainment | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 214 d, liq $18022) |
| 2026-05-30T12:33:19Z | databricks-ipo-before-2027 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $2545 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $205 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4581 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $422) |
| 2026-05-30T12:33:19Z | will-there-be-between-11-and-13-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3600 | P3_low_absolute_liquidity | liquidity $1556 < absolute min $5000 |
| 2026-05-30T12:33:19Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1564 < absolute min $5000 |
| 2026-05-30T12:33:19Z | brex-ipo-before-2027 | crypto-launch | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $4766) |
| 2026-05-30T12:33:19Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T12:33:19Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $638) |
| 2026-05-30T12:33:19Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $218 < absolute min $5000 |
| 2026-05-30T12:33:19Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 214 d, liq $9625) |
| 2026-05-30T12:33:19Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $249 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3106 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4653 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $134) |
| 2026-05-30T12:33:19Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $7282) |
| 2026-05-30T12:33:19Z | major-meteor-strike-10kt-in-2026 | weather-natural | 0.1300 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T12:33:19Z | will-there-be-exactly-1-confirmed-vei-4-or-higher-volcanic-eruption-worldwide-in-2026-937 | weather-natural | 0.2700 | P3_low_absolute_liquidity | liquidity $1442 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 214 d, liq $7972) |
| 2026-05-30T12:33:19Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $183397) |
| 2026-05-30T12:33:19Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $4789) |
| 2026-05-30T12:33:19Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $688) |
| 2026-05-30T12:33:19Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $16334) |
| 2026-05-30T12:33:19Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $18684) |
| 2026-05-30T12:33:19Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $12633) |
| 2026-05-30T12:33:19Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3963 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $6242) |
| 2026-05-30T12:33:19Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $322) |
| 2026-05-30T12:33:19Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $14194) |
| 2026-05-30T12:33:19Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $279) |
| 2026-05-30T12:33:19Z | will-there-be-between-14-and-16-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2406 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $427) |
| 2026-05-30T12:33:19Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2640 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $194) |
| 2026-05-30T12:33:19Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $6070) |
| 2026-05-30T12:33:19Z | jerome-powell-out-from-fed-board-by-may-30 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 214 d, liq $56398) |
| 2026-05-30T12:33:19Z | will-any-category-5-hurricane-make-landfall-in-the-us-in-before-2027 | weather-natural | 0.1300 | P3_low_absolute_liquidity | liquidity $1626 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $7927) |
| 2026-05-30T12:33:19Z | 5kt-meteor-strike-in-2026 | weather-natural | 0.3200 | P3_low_absolute_liquidity | liquidity $3062 < absolute min $5000 |
| 2026-05-30T12:33:19Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $12625) |
| 2026-05-30T12:33:19Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1230 | P3_low_absolute_liquidity | liquidity $3177 < absolute min $5000 |
| 2026-05-30T12:33:19Z | opensea-fdv-above-5b-one-day-after-launch-919-666-825-992-543-333 | crypto-launch | 0.0500 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:19Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $12642) |
| 2026-05-30T12:33:19Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0910 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:19Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:19Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3200 | P3_low_absolute_liquidity | liquidity $574 < absolute min $5000 |
| 2026-05-30T12:33:19Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1300 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | will-ethereum-reach-6000-by-december-31-2026 | market | 0.0600 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:33:20Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:33:20Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P3_low_absolute_liquidity | liquidity $774 < absolute min $5000 |
| 2026-05-30T12:33:20Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:33:20Z | altcoin-market-cap-dip-to-150b-before-2027 | other | 0.7300 | P3_low_absolute_liquidity | liquidity $642 < absolute min $5000 |
| 2026-05-30T12:33:20Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $663 < absolute min $5000 |
| 2026-05-30T12:33:20Z | will-ethereum-reach-7000-by-december-31-2026 | market | 0.0320 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.06) |
| 2026-05-30T12:33:20Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3900 | P3_low_absolute_liquidity | liquidity $1146 < absolute min $5000 |
| 2026-05-30T12:33:20Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3268 < absolute min $5000 |
| 2026-05-30T12:33:20Z | will-stablecoins-hit-500b-before-2027 | other | 0.1200 | P3_low_absolute_liquidity | liquidity $2287 < absolute min $5000 |
| 2026-05-30T12:33:20Z | opensea-fdv-above-100m-one-day-after-launch-172-151-588-987 | crypto-launch | 0.6000 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | will-pumpfun-perform-an-airdrop-by-december-31-2026 | crypto-launch | 0.2800 | P3_low_absolute_liquidity | liquidity $1431 < absolute min $5000 |
| 2026-05-30T12:33:20Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2406 < absolute min $5000 |
| 2026-05-30T12:33:20Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $1446) |
| 2026-05-30T12:33:20Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 239 d, liq $19269) |
| 2026-05-30T12:33:20Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $52851) |
| 2026-05-30T12:33:20Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 239 d, liq $45943) |
| 2026-05-30T12:33:20Z | will-the-atlanta-falcons-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 305 d, liq $192221) |
| 2026-05-30T12:33:20Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $265014) |
| 2026-05-30T12:33:20Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $43592) |
| 2026-05-30T12:33:20Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3983 < absolute min $5000 |
| 2026-05-30T12:33:20Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8200 | P4_pre_event | pre-event slug + 579 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | will-spacex-not-ipo-by-december-31-2027-481-338 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 579 d, liq $11467) |
| 2026-05-30T12:33:20Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $46212) |
| 2026-05-30T12:33:20Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.2400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2180 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T12:33:20Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2364615) |
| 2026-05-30T12:33:20Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 891 d, liq $1916229) |
| 2026-05-30T12:33:20Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1457300) |
| 2026-05-30T12:33:20Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1311194) |
| 2026-05-30T12:33:20Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1450544) |
| 2026-05-30T12:33:20Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1157955) |
| 2026-05-30T12:33:20Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $475235) |
| 2026-05-30T12:33:20Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1075568) |
| 2026-05-30T12:33:20Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2333837) |
| 2026-05-30T12:33:20Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1170540) |
| 2026-05-30T12:33:20Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1722563) |
| 2026-05-30T12:33:20Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $937039) |
| 2026-05-30T12:33:20Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2438153) |
| 2026-05-30T12:33:20Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1483587) |
| 2026-05-30T12:33:20Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $946172) |
| 2026-05-30T12:33:20Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1060939) |
| 2026-05-30T12:33:20Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1626709) |
| 2026-05-30T12:33:20Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0400 | P3_low_absolute_liquidity | liquidity $2930 < absolute min $5000 |
| 2026-05-30T12:33:20Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $23048) |
| 2026-05-30T12:33:20Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $4608) |
| 2026-05-30T12:33:20Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12018) |
| 2026-05-30T12:33:20Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2400 | P3_low_absolute_liquidity | liquidity $624 < absolute min $5000 |
| 2026-05-30T12:33:20Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $1211) |
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
| claudemax-websearch | 11 | 2 | 8 | 20.0% | -77.68 | YES |

## Anti-patterns identificados- same-day-price-target · intraday-decay · no-momentum-confirm · short-horizon-chalk · high-entry-low-exit — visto en  (2026-05-26, pnl $-57.55)

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

