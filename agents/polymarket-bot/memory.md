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
| 2026-05-24T03:30:03Z | romanian-pm-bolojan-out-by-december-31-832-595 | uncategorized | 0.9860 | P0_ceiling | price ceiling: 0.9860 > 0.950 |
| 2026-05-24T03:30:03Z | applied-intuition-ipo-before-2027 | uncategorized | 0.1500 | P3_low_absolute_liquidity | liquidity $2621 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-zelenskyy-and-putin-meet-next-in-ukraine | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-hassan-shariatmadari-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-zelenskyy-and-putin-meet-next-in-us | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-china-invade-taiwan-before-2027 | uncategorized | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-alberta-join-the-us | uncategorized | 0.0410 | P0_floor | price floor: 0.0410 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-there-be-at-least-7500-measles-cases-in-the-us-in-2026-287-181 | uncategorized | 0.1500 | P3_low_absolute_liquidity | liquidity $3130 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | uncategorized | 0.5110 | P3_low_absolute_liquidity | liquidity $130 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-zelenskyy-and-putin-meet-next-in-qatar-uae | uncategorized | 0.0470 | P0_floor | price floor: 0.0470 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-jensen-huang-be-richest-person-on-december-31-229 | uncategorized | 0.0530 | P0_floor | price floor: 0.0530 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | uncategorized | 0.1000 | P3_low_absolute_liquidity | liquidity $2375 < absolute min $5000 |
| 2026-05-24T03:30:03Z | vanta-ipo-before-2027 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-jeff-bezos-be-richest-person-on-december-31-243 | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-any-category-4-hurricane-make-landfall-in-the-us-in-before-2027 | uncategorized | 0.3600 | P3_low_absolute_liquidity | liquidity $1675 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | uncategorized | 0.0480 | P0_floor | price floor: 0.0480 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-freecs-win-the-lck-2026-season-playoffs | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-magomed-ankalaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-steve-ballmer-be-richest-person-on-december-31 | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 220 d) |
| 2026-05-24T03:30:03Z | will-no-player-win-a-calendar-grand-slam-in-2026 | uncategorized | 0.9920 | P0_ceiling | price ceiling: 0.9920 > 0.950 |
| 2026-05-24T03:30:03Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0620 | P0_floor | price floor: 0.0620 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | uncategorized | 0.1800 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | will-fomo-launch-a-token-by-june-30-2026 | uncategorized | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | will-consensys-ipo-by-december-31-2026 | uncategorized | 0.2900 | P3_low_absolute_liquidity | liquidity $1735 < absolute min $5000 |
| 2026-05-24T03:30:03Z | hyperbeat-fdv-above-300m-one-day-after-launch | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | consensys-ipo-closing-market-cap-above-1b | uncategorized | 0.3700 | P3_low_absolute_liquidity | liquidity $1060 < absolute min $5000 |
| 2026-05-24T03:30:03Z | opensea-fdv-above-5b-one-day-after-launch-919-666-825-992-543-333 | uncategorized | 0.0550 | P0_floor | price floor: 0.0550 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | extended-fdv-above-800m-one-day-after-launch-299-255-178 | uncategorized | 0.1200 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | uncategorized | 0.1400 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | puffpaw-fdv-above-50m-one-day-after-launch | uncategorized | 0.7300 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | will-fomo-launch-a-token-by-december-31-2026 | uncategorized | 0.3800 | P3_low_absolute_liquidity | liquidity $740 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | uncategorized | 0.3460 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | will-exponent-launch-a-token-by-december-31-2026 | uncategorized | 0.3800 | P3_low_absolute_liquidity | liquidity $927 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-pumpfun-perform-an-airdrop-by-december-31-2026 | uncategorized | 0.2600 | P3_low_absolute_liquidity | liquidity $1134 < absolute min $5000 |
| 2026-05-24T03:30:03Z | hyperliquid-listed-on-binance-in-2026-485 | uncategorized | 0.2700 | P3_low_absolute_liquidity | liquidity $2680 < absolute min $5000 |
| 2026-05-24T03:30:03Z | metamask-fdv-above-1b-one-day-after-launch-978-851-628-634 | uncategorized | 0.2300 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | metamask-fdv-above-3b-one-day-after-launch-363-663-664-569-222 | uncategorized | 0.1310 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | puffpaw-fdv-above-300m-one-day-after-launch | uncategorized | 0.0990 | P0_floor | price floor: 0.0990 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.6900 | P3_low_absolute_liquidity | liquidity $2780 < absolute min $5000 |
| 2026-05-24T03:30:03Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 222 d) |
| 2026-05-24T03:30:03Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-tampa-bay-buccaneers-win-the-2027-nfl-nfc-championship-312 | uncategorized | 0.0380 | P0_floor | price floor: 0.0380 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T03:30:03Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 245 d) |
| 2026-05-24T03:30:03Z | will-mathilde-panot-win-the-2027-french-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 340 d) |
| 2026-05-24T03:30:03Z | will-lisabeth-borne-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 340 d) |
| 2026-05-24T03:30:03Z | will-sgolne-royal-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 340 d) |
| 2026-05-24T03:30:03Z | will-utah-jazz-win-the-2027-nba-finals | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 402 d) |
| 2026-05-24T03:30:03Z | will-sacramento-kings-win-the-2027-nba-finals | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 402 d) |
| 2026-05-24T03:30:03Z | will-jpmorgan-chase-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-off... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 585 d) |
| 2026-05-24T03:30:03Z | will-morgan-stanley-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-off... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 585 d) |
| 2026-05-24T03:30:03Z | will-wells-fargo-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offeri... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 585 d) |
| 2026-05-24T03:30:03Z | will-deutsche-bank-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offe... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 585 d) |
| 2026-05-24T03:30:03Z | openai-ipo-closing-market-cap-above-1pt2t | uncategorized | 0.6600 | P3_low_absolute_liquidity | liquidity $3016 < absolute min $5000 |
| 2026-05-24T03:30:03Z | spacex-ipo-closing-market-cap-above-3t-337-436 | uncategorized | 0.1300 | P4_pre_event | pre-event slug + 585 d to resolution (>=7 threshold) |
| 2026-05-24T03:30:03Z | abstract-fdv-above-3b-one-day-after-launch | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 587 d) |
| 2026-05-24T03:30:03Z | will-kristi-noem-win-the-2028-republican-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-tulsi-gabbard-win-the-2028-republican-presidential-nomination | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-wes-moore-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-glenn-youngkin-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-person-a-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-eric-trump-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-ro-khanna-win-the-2028-us-presidential-election | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-glenn-youngkin-win-the-2028-republican-presidential-nomination | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T03:30:03Z | will-the-iran-ceasefire-continue-through-june-15-136-565 | uncategorized | 0.8600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T03:30:03Z | will-the-iran-ceasefire-continue-through-june-30-529-427 | uncategorized | 0.8600 | M1 | memoria: slug prefix match; same price bucket high (score 0.70) |
| 2026-05-24T03:30:03Z | will-the-iran-ceasefire-continue-through-may-31-654-633 | uncategorized | 0.8800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T03:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-23 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon -1 d) |
| 2026-05-24T03:30:03Z | will-the-iran-ceasefire-continue-through-december-31-395-943 | uncategorized | 0.6600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T03:30:03Z | will-the-iran-ceasefire-continue-through-may-27-496 | uncategorized | 0.9300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T03:30:03Z | will-axiom-launch-a-token-by-december-31-2026 | uncategorized | 0.3100 | P3_low_absolute_liquidity | liquidity $387 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-the-iran-ceasefire-continue-through-may-24-733 | uncategorized | 0.9780 | P0_ceiling | price ceiling: 0.9780 > 0.950 |
| 2026-05-24T03:30:03Z | will-variational-launch-a-token-by-december-31-2026-577-621 | uncategorized | 0.8200 | P3_low_absolute_liquidity | liquidity $1347 < absolute min $5000 |
| 2026-05-24T03:30:03Z | will-valve-remove-overpass-from-the-map-pool | uncategorized | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon -1 d) |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
| 2026-05-22T21:35:02Z | T-2296135-1779485404133 | uncategorized | 0.2700 | 0.4100 | 30.02 | take_profit |  | 57.89 | medium | 0.00 |
| 2026-05-22T21:50:01Z | T-2133405-1779211529037 | uncategorized | 0.4600 | 0.8200 | 48.65 | take_profit |  | 62.17 | medium | 3.18 |
| 2026-05-22T21:55:02Z | T-2296151-1779363004639 | uncategorized | 0.1500 | 0.2900 | 54.05 | take_profit |  | 57.91 | short | 1.43 |
| 2026-05-22T22:00:01Z | T-2241742-1779210004342 | uncategorized | 0.6200 | 0.9700 | 36.77 | take_profit |  | 65.14 | long | 3.21 |
| 2026-05-22T22:10:02Z | T-2111640-1779310803840 | uncategorized | 0.1500 | 0.0300 | -47.47 | stop_loss |  | 59.34 | medium | 2.05 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T23:25:02Z | T-2078702-1779480004213 | uncategorized | 0.9400 | 0.9990 | 3.42 | market_closed |  | 54.47 | short | 0.14 |
| 2026-05-22T23:30:01Z | T-2183424-1779210004342 | uncategorized | 0.1350 | 0.1900 | 27.08 | take_profit |  | 66.47 | long | 3.27 |
| 2026-05-23T00:30:01Z | T-2132778-1779255003919 | uncategorized | 0.6300 | 0.9000 | 23.12 | take_profit |  | 53.95 | medium | 2.79 |
| 2026-05-23T00:45:02Z | T-2327930-1779496203488 | uncategorized | 0.2110 | 0.3100 | 28.38 | take_profit |  | 60.48 | short | 0.01 |
| 2026-05-23T01:30:01Z | T-2261347-1779251403658 | uncategorized | 0.7640 | 0.9990 | 16.32 | market_closed |  | 53.05 | medium | 2.87 |
| 2026-05-23T01:30:01Z | T-2159735-1779462003955 | uncategorized | 0.7200 | 0.9990 | 19.22 | market_closed |  | 49.60 | short | 0.44 |
| 2026-05-23T02:10:02Z | T-2327930-1779498003709 | uncategorized | 0.5040 | 0.0720 | -52.33 | stop_loss |  | 61.05 | short | 0.05 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T03:05:02Z | T-553856-1779233403338 | uncategorized | 0.4200 | 0.6300 | 26.58 | take_profit |  | 53.16 | long | 3.15 |
| 2026-05-23T03:20:02Z | T-2313547-1779483603887 | uncategorized | 0.0900 | 0.0130 | -50.27 | stop_loss |  | 58.76 | medium | 0.26 |
| 2026-05-23T05:50:02Z | T-2277981-1779478204560 | uncategorized | 0.6800 | 0.1300 | -44.57 | stop_loss |  | 55.10 | short | 0.43 |
| 2026-05-23T06:05:02Z | T-2080149-1779516004165 | uncategorized | 0.0500 | 0.0010 | -57.48 | stop_loss |  | 58.65 | short | 0.00 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T08:05:01Z | T-2132779-1779260404871 | uncategorized | 0.1400 | 0.2600 | 45.62 | take_profit |  | 53.22 | medium | 3.05 |
| 2026-05-23T13:30:02Z | T-2088261-1779525003640 | uncategorized | 0.5300 | 0.9500 | 46.22 | take_profit |  | 58.33 | short | 0.21 |
| 2026-05-23T13:40:02Z | T-2079836-1779543004239 | uncategorized | 0.2800 | 0.0100 | -57.13 | stop_loss |  | 59.25 | short | 0.01 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T15:10:02Z | T-1919425-1779483603887 | uncategorized | 0.2000 | 0.2800 | 22.57 | take_profit |  | 56.43 | medium | 0.76 |
| 2026-05-23T15:40:02Z | T-1972137-1779485404133 | uncategorized | 0.2500 | 0.3600 | 24.96 | take_profit |  | 56.74 | medium | 0.76 |
| 2026-05-23T16:30:02Z | T-2324252-1779471003629 | uncategorized | 0.1900 | 0.2700 | 22.34 | take_profit |  | 53.05 | medium | 0.96 |
| 2026-05-23T16:30:02Z | T-2333646-1779544804084 | uncategorized | 0.5100 | 0.0900 | -47.85 | stop_loss |  | 58.11 | short | 0.10 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T16:40:02Z | T-2333644-1779550203541 | uncategorized | 0.1700 | 0.9600 | 271.97 | take_profit |  | 58.53 | short | 0.05 |
| 2026-05-23T16:40:02Z | T-2333646-1779553804069 | uncategorized | 0.2400 | 0.0100 | -57.04 | stop_loss |  | 59.52 | short | 0.01 |
| 2026-05-23T17:15:02Z | T-2126542-1779337803770 | uncategorized | 0.1170 | 0.0220 | -47.13 | stop_loss |  | 58.04 | medium | 2.53 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T18:10:02Z | T-2313549-1779503403697 | uncategorized | 0.1500 | 0.2100 | 24.12 | take_profit |  | 60.29 | medium | 0.65 |
| 2026-05-23T18:25:01Z | T-1962237-1779314404475 | uncategorized | 0.4300 | 0.6400 | 29.10 | take_profit |  | 59.59 | medium | 2.85 |
| 2026-05-23T18:25:01Z | T-2296151-1779487204257 | uncategorized | 0.3200 | 0.0400 | -55.13 | stop_loss |  | 63.00 | short | 0.85 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T18:35:02Z | T-2132637-1779492603815 | uncategorized | 0.2400 | 0.3400 | 25.06 | take_profit |  | 60.14 | medium | 0.80 |
| 2026-05-23T18:50:02Z | T-2241873-1779487204257 | uncategorized | 0.8200 | 0.1390 | -53.39 | stop_loss |  | 64.29 | short | 0.87 |
| 2026-05-23T19:00:01Z | T-2159726-1779211529037 | uncategorized | 0.4300 | 0.0400 | -57.54 | stop_loss |  | 63.44 | medium | 4.07 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T19:30:02Z | T-2238460-1779562803685 | uncategorized | 0.4100 | 0.5900 | 27.30 | take_profit |  | 62.19 | short | 0.02 |
| 2026-05-23T19:35:02Z | T-2238414-1779555603718 | uncategorized | 0.3900 | 0.6600 | 44.19 | take_profit |  | 63.82 | short | 0.11 |
| 2026-05-23T19:35:02Z | T-2087618-1779561003587 | uncategorized | 0.1600 | 0.0300 | -51.92 | stop_loss |  | 63.90 | short | 0.05 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T19:55:02Z | T-2238258-1779552004213 | uncategorized | 0.6500 | 0.9400 | 26.32 | take_profit |  | 58.99 | short | 0.16 |
| 2026-05-23T20:05:02Z | T-2238462-1779564603418 | uncategorized | 0.2400 | 0.3700 | 32.67 | take_profit |  | 60.32 | short | 0.02 |
| 2026-05-23T20:25:02Z | T-2238307-1779553804069 | uncategorized | 0.4900 | 0.7600 | 32.14 | take_profit |  | 58.33 | short | 0.16 |
| 2026-05-23T20:30:02Z | T-2287602-1779483603887 | uncategorized | 0.6000 | 0.8700 | 26.98 | take_profit |  | 59.96 | short | 0.98 |
| 2026-05-23T20:30:03Z | T-2238511-1779557404276 | uncategorized | 0.3100 | 0.0500 | -51.59 | stop_loss |  | 61.52 | short | 0.12 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T20:35:02Z | T-2238370-1779499803774 | uncategorized | 0.5100 | 0.7200 | 25.77 | take_profit |  | 62.59 | short | 0.80 |
| 2026-05-23T20:35:02Z | T-2238507-1779507004281 | uncategorized | 0.3400 | 0.7000 | 64.43 | take_profit |  | 60.85 | short | 0.71 |
| 2026-05-23T20:35:02Z | T-2238416-1779562803685 | uncategorized | 0.2700 | 0.0400 | -51.91 | stop_loss |  | 60.94 | short | 0.07 |
| 2026-05-23T20:45:02Z | T-2074237-1779229804249 | uncategorized | 0.6200 | 0.1200 | -47.13 | stop_loss |  | 58.44 | medium | 3.93 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T20:45:02Z | T-2238359-1779561003587 | uncategorized | 0.2900 | 0.6800 | 84.22 | take_profit |  | 62.63 | short | 0.09 |
| 2026-05-23T20:45:02Z | T-2238460-1779566403548 | uncategorized | 0.3300 | 0.7000 | 68.07 | take_profit |  | 60.72 | short | 0.03 |
| 2026-05-23T20:55:01Z | T-1961536-1779210004342 | uncategorized | 0.4050 | 0.6050 | 34.88 | take_profit |  | 70.62 | long | 4.16 |
| 2026-05-23T20:55:02Z | T-2238649-1779561003587 | uncategorized | 0.6700 | 0.9900 | 29.31 | take_profit |  | 61.37 | short | 0.10 |
| 2026-05-23T20:55:02Z | T-2238614-1779566403548 | uncategorized | 0.5100 | 0.1000 | -49.81 | stop_loss |  | 61.95 | short | 0.04 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T21:00:01Z | T-957019-1779211529037 | uncategorized | 0.2900 | 0.4200 | 25.19 | take_profit |  | 56.20 | long | 4.15 |
| 2026-05-23T21:00:02Z | T-2074236-1779215403605 | uncategorized | 0.2300 | 0.0400 | -38.34 | stop_loss |  | 46.41 | medium | 4.10 |
| 2026-05-23T21:00:02Z | T-2074235-1779228007314 | uncategorized | 0.1100 | 0.0130 | -49.76 | stop_loss |  | 56.42 | medium | 3.96 |
| 2026-05-23T21:00:02Z | T-2238511-1779568204835 | uncategorized | 0.0600 | 0.0070 | -57.51 | stop_loss |  | 65.11 | short | 0.02 |
| 2026-05-23T21:05:02Z | T-2238614-1779570004545 | uncategorized | 0.1100 | 0.0200 | -60.10 | stop_loss |  | 73.46 | short | 0.00 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T21:05:02Z | T-2285045-1779570004545 | uncategorized | 0.1280 | 0.2100 | 40.85 | take_profit |  | 63.77 | short | 0.00 |
| 2026-05-23T21:10:02Z | T-2238602-1779507004281 | uncategorized | 0.5700 | 0.0600 | -53.35 | stop_loss |  | 59.63 | short | 0.74 |
| 2026-05-23T21:10:02Z | T-2238414-1779566403548 | uncategorized | 0.7100 | 0.0100 | -62.33 | stop_loss |  | 63.22 | short | 0.05 |
| 2026-05-23T21:10:02Z | T-2238462-1779568204835 | uncategorized | 0.4600 | 0.0600 | -55.48 | stop_loss |  | 63.81 | short | 0.03 |
| 2026-05-23T21:10:02Z | T-2238558-1779568204835 | uncategorized | 0.5700 | 0.0100 | -61.43 | stop_loss |  | 62.53 | short | 0.03 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T21:10:02Z | T-2238416-1779570004545 | uncategorized | 0.1400 | 0.9900 | 419.76 | take_profit |  | 69.14 | short | 0.01 |
| 2026-05-23T21:15:02Z | T-2238554-1779555603718 | uncategorized | 0.4500 | 0.0100 | -61.16 | stop_loss |  | 62.55 | short | 0.18 |
| 2026-05-23T21:20:02Z | T-2238509-1779570004545 | uncategorized | 0.1700 | 0.0100 | -71.99 | stop_loss |  | 76.49 | short | 0.01 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T21:40:02Z | T-2238507-1779570004545 | uncategorized | 0.8300 | 0.9990 | 14.66 | market_closed |  | 71.99 | short | 0.03 |
| 2026-05-23T22:30:02Z | T-2086805-1779571804012 | uncategorized | 0.1900 | 0.0300 | -60.13 | stop_loss |  | 71.41 | short | 0.04 |
| 2026-05-23T22:40:02Z | T-2296151-1779570004545 | uncategorized | 0.0780 | 0.0110 | -58.20 | stop_loss |  | 67.75 | short | 0.07 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T22:40:02Z | T-2086811-1779573603629 | uncategorized | 0.5800 | 0.8200 | 25.35 | take_profit |  | 61.27 | short | 0.03 |
| 2026-05-23T22:50:02Z | T-2238460-1779570004545 | uncategorized | 0.8100 | 0.9990 | 16.46 | market_closed |  | 70.55 | short | 0.08 |
| 2026-05-23T22:50:02Z | T-2086805-1779575403333 | uncategorized | 0.0500 | 0.0010 | -59.06 | stop_loss |  | 60.27 | short | 0.01 |
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-23T23:00:02Z | T-2238307-1779568204835 | uncategorized | 0.7900 | 0.9990 | 16.21 | market_closed |  | 61.28 | short | 0.10 |
| 2026-05-23T23:25:02Z | T-2238359-1779570004545 | uncategorized | 0.7600 | 0.9990 | 23.57 | market_closed |  | 74.96 | short | 0.10 |
| 2026-05-23T23:50:01Z | T-2313548-1779483603887 | uncategorized | 0.1500 | 0.2100 | 23.03 | take_profit |  | 57.59 | medium | 1.12 |
| 2026-05-24T00:10:02Z | T-2087616-1779487204257 | uncategorized | 0.7300 | 0.9990 | 24.17 | market_closed |  | 65.60 | short | 1.09 |
| 2026-05-24T00:25:03Z | T-2313547-1779571804012 | uncategorized | 0.0800 | 0.1210 | 32.42 | take_profit |  | 63.26 | short | 0.12 |
| 2026-05-24T00:40:02Z | T-2111640-1779577204027 | uncategorized | 0.1450 | 0.2030 | 25.72 | take_profit |  | 64.30 | medium | 0.07 |
| 2026-05-24T01:40:01Z | T-2270164-1779210004342 | uncategorized | 0.2250 | 0.3800 | 47.68 | take_profit |  | 69.21 | long | 4.36 |
| 2026-05-24T01:40:02Z | T-2086878-1779570004545 | uncategorized | 0.4700 | 0.8300 | 50.86 | take_profit |  | 66.40 | short | 0.19 |
| 2026-05-24T02:05:02Z | T-2086885-1779580803643 | uncategorized | 0.5300 | 0.7500 | 25.11 | take_profit |  | 60.50 | short | 0.09 |
| 2026-05-24T02:35:02Z | T-2086922-1779588003845 | uncategorized | 0.4900 | 0.7100 | 28.93 | take_profit |  | 64.43 | short | 0.02 |
| 2026-05-24T02:50:02Z | T-2086901-1779582603386 | uncategorized | 0.7100 | 0.9990 | 25.64 | take_profit |  | 63.00 | short | 0.10 |
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


## Human notes
_(no se toca por automatización)_
