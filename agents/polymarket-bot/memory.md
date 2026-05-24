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
| 2026-05-24T21:30:03Z | will-the-us-invade-iran-before-2027 | uncategorized | 0.1900 | M1 | memoria: slug prefix match; same price bucket low (score 0.70) |
| 2026-05-24T21:30:03Z | will-alireza-arafi-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-petr-yan-fight-dominick-cruz-next | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-zelenskyy-and-putin-meet-next-in-russia | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-fearx-win-the-lck-2026-season-playoffs | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-openai-be-acquired-before-2027-859 | uncategorized | 0.0770 | P0_floor | price floor: 0.0770 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-caio-borralho-be-the-ufc-middleweight-champion-on-december-31-2026 | uncategorized | 0.0300 | P0_floor | price floor: 0.0300 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-steve-ballmer-be-richest-person-on-december-31 | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-freecs-win-the-lck-2026-season-playoffs | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-gitlab-be-acquired-before-2027-944-667 | uncategorized | 0.2100 | P3_low_absolute_liquidity | liquidity $4401 < absolute min $5000 |
| 2026-05-24T21:30:03Z | will-alberta-join-the-us | uncategorized | 0.0410 | P0_floor | price floor: 0.0410 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-china-invade-taiwan-before-2027 | uncategorized | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | deel-ipo-before-2027 | uncategorized | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-jeff-bezos-be-richest-person-on-december-31-243 | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | ramp-ipo-before-2027 | uncategorized | 0.1100 | P3_low_absolute_liquidity | liquidity $3039 < absolute min $5000 |
| 2026-05-24T21:30:03Z | will-the-feds-lower-bound-reach-0pt25-or-lower-before-2027-173-921-916-764-754-593-935 | uncategorized | 0.0510 | P0_floor | price floor: 0.0510 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | mistral-ai-ipo-before-2027 | uncategorized | 0.1600 | P3_low_absolute_liquidity | liquidity $3879 < absolute min $5000 |
| 2026-05-24T21:30:03Z | anduril-ipo-before-2027 | uncategorized | 0.1500 | P3_low_absolute_liquidity | liquidity $2994 < absolute min $5000 |
| 2026-05-24T21:30:03Z | sbf-released-from-custody-in-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | will-bernard-arnault-be-richest-person-on-december-31-747 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 220 d) |
| 2026-05-24T21:30:03Z | opensea-fdv-above-1b-one-day-after-launch-426-422-754-756-879-328-494 | uncategorized | 0.2400 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | puffpaw-fdv-above-50m-one-day-after-launch | uncategorized | 0.7300 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | uncategorized | 0.1600 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | ethereum-all-time-high-by-december-31-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | puffpaw-fdv-above-400m-one-day-after-launch | uncategorized | 0.1260 | P3_low_absolute_liquidity | liquidity $3678 < absolute min $5000 |
| 2026-05-24T21:30:03Z | okx-ipo-in-2026 | uncategorized | 0.1300 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | will-tempo-launch-a-token-by-june-30-2026 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | will-usdt-market-cap-hit-200b-before-2027 | uncategorized | 0.9760 | P0_ceiling | price ceiling: 0.9760 > 0.950 |
| 2026-05-24T21:30:03Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | uncategorized | 0.1800 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | uncategorized | 0.1500 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | will-opensea-launch-a-token-by-december-31-2026 | uncategorized | 0.6780 | P3_low_absolute_liquidity | liquidity $3717 < absolute min $5000 |
| 2026-05-24T21:30:03Z | will-consensys-ipo-by-december-31-2026 | uncategorized | 0.2900 | P3_low_absolute_liquidity | liquidity $1722 < absolute min $5000 |
| 2026-05-24T21:30:03Z | hyperbeat-fdv-above-300m-one-day-after-launch | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.2300 | P3_low_absolute_liquidity | liquidity $1281 < absolute min $5000 |
| 2026-05-24T21:30:03Z | extended-fdv-above-150m-one-day-after-launch-452 | uncategorized | 0.6500 | P4_pre_event | pre-event slug + 221 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0620 | P0_floor | price floor: 0.0620 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | will-pumpfun-perform-an-airdrop-by-december-31-2026 | uncategorized | 0.2500 | P3_low_absolute_liquidity | liquidity $1351 < absolute min $5000 |
| 2026-05-24T21:30:03Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 221 d) |
| 2026-05-24T21:30:03Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-atlanta-falcons-win-the-2027-nfl-nfc-championship-312 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T21:30:03Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 245 d) |
| 2026-05-24T21:30:03Z | will-saquon-barkley-win-the-2026-nfl-mvp | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 266 d) |
| 2026-05-24T21:30:03Z | will-the-atlanta-falcons-win-the-2027-nfl-league-championship | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 311 d) |
| 2026-05-24T21:30:03Z | will-utah-jazz-win-the-2027-nba-finals | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 402 d) |
| 2026-05-24T21:30:03Z | will-sacramento-kings-win-the-2027-nba-finals | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 402 d) |
| 2026-05-24T21:30:03Z | will-spacexs-public-ticker-be-spc-937 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 585 d) |
| 2026-05-24T21:30:03Z | will-spacexs-public-ticker-be-star-796 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 585 d) |
| 2026-05-24T21:30:03Z | abstract-fdv-above-3b-one-day-after-launch | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 586 d) |
| 2026-05-24T21:30:03Z | base-fdv-above-2b-one-day-after-launch | uncategorized | 0.8000 | P4_pre_event | pre-event slug + 586 d to resolution (>=7 threshold) |
| 2026-05-24T21:30:03Z | will-john-fetterman-win-the-2028-democratic-presidential-nomination-941 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-beto-orourke-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-matt-gaetz-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-elon-musk-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-dwayne-the-rock-johnson-win-the-2028-us-presidential-election | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-glenn-youngkin-win-the-2028-republican-presidential-nomination | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-george-clooney-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-ro-khanna-win-the-2028-us-presidential-election | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-eric-trump-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-mark-kelly-win-the-2028-democratic-presidential-nomination-479 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-donald-trump-jr-win-the-2028-us-presidential-election | uncategorized | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 897 d) |
| 2026-05-24T21:30:03Z | will-variational-launch-a-token-by-december-31-2026-577-621 | uncategorized | 0.8200 | P3_low_absolute_liquidity | liquidity $1344 < absolute min $5000 |
| 2026-05-24T21:30:03Z | reya-fdv-above-1b-one-day-after-launch-348-347 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon -1 d) |
| 2026-05-24T21:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-25 | uncategorized | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon -1 d) |
| 2026-05-24T21:30:03Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon -1 d) |
| 2026-05-24T21:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-26 | uncategorized | 0.1600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T21:30:03Z | will-the-iran-ceasefire-continue-through-may-31-654-633 | uncategorized | 0.8200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T21:30:03Z | will-the-iran-ceasefire-continue-through-june-30-529-427 | uncategorized | 0.7200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T21:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-24 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon -1 d) |
| 2026-05-24T21:30:03Z | will-the-iran-ceasefire-continue-through-may-27-496 | uncategorized | 0.8900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T21:30:03Z | will-rebecca-shepherd-win-the-2026-makerfield-by-election | uncategorized | 0.0580 | P0_floor | price floor: 0.0580 < 0.100 (horizon -1 d) |
| 2026-05-24T21:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-31 | uncategorized | 0.3400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-24T21:30:03Z | will-the-iran-ceasefire-continue-through-may-24-733 | uncategorized | 0.9810 | P0_ceiling | price ceiling: 0.9810 > 0.950 |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
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
| 2026-05-24T04:10:02Z | T-2086878-1779588003845 | uncategorized | 0.8200 | 0.9990 | 14.35 | market_closed |  | 65.74 | short | 0.09 |
| 2026-05-24T04:15:01Z | T-1708132-1779211529037 | uncategorized | 0.0600 | 0.0120 | -48.74 | stop_loss |  | 60.93 | medium | 4.45 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T04:15:02Z | T-2086885-1779589803605 | uncategorized | 0.8700 | 0.9990 | 9.62 | market_closed |  | 64.85 | short | 0.07 |
| 2026-05-24T05:05:02Z | T-2086922-1779591603222 | uncategorized | 0.8000 | 0.9990 | 16.72 | market_closed |  | 67.20 | short | 0.09 |
| 2026-05-24T05:35:02Z | T-2285045-1779571804012 | uncategorized | 0.1800 | 0.0350 | -56.37 | stop_loss |  | 69.98 | short | 0.34 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T06:10:02Z | T-2313547-1779582603386 | uncategorized | 0.1100 | 0.1590 | 27.50 | take_profit |  | 61.74 | short | 0.24 |
| 2026-05-24T06:10:02Z | T-2111605-1779597003852 | uncategorized | 0.2050 | 0.2890 | 26.71 | take_profit |  | 65.18 | medium | 0.07 |
| 2026-05-24T07:30:02Z | T-2111563-1779577204027 | uncategorized | 0.0730 | 0.1050 | 26.53 | take_profit |  | 60.52 | medium | 0.35 |
| 2026-05-24T07:50:02Z | T-2289407-1779570004545 | uncategorized | 0.0760 | 0.0150 | -49.16 | stop_loss |  | 61.24 | short | 0.45 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T08:30:02Z | T-2313548-1779584403287 | uncategorized | 0.1900 | 0.2700 | 26.23 | take_profit |  | 62.31 | short | 0.31 |
| 2026-05-24T09:05:01Z | T-2132638-1779487204257 | uncategorized | 0.1900 | 0.0300 | -50.95 | stop_loss |  | 60.51 | medium | 1.46 |
| 2026-05-24T09:35:01Z | T-2289402-1779570004545 | uncategorized | 0.0710 | 0.0080 | -53.26 | stop_loss |  | 60.02 | short | 0.52 |
| 2026-05-24T10:20:02Z | T-2296151-1779597003852 | uncategorized | 0.0900 | 0.0080 | -61.84 | stop_loss |  | 67.87 | short | 0.24 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T10:45:02Z | T-2285045-1779615004015 | uncategorized | 0.0500 | 0.1110 | 78.35 | take_profit |  | 64.23 | short | 0.05 |
| 2026-05-24T10:50:01Z | T-2324516-1779492603815 | uncategorized | 0.0550 | 0.0100 | -50.21 | stop_loss |  | 61.37 | short | 1.47 |
| 2026-05-24T12:05:01Z | T-2313552-1779571804012 | uncategorized | 0.0920 | 0.0180 | -51.92 | stop_loss |  | 64.55 | short | 0.61 |
| 2026-05-24T14:30:02Z | T-2243584-1779571804012 | uncategorized | 0.0970 | 0.0110 | -59.59 | stop_loss |  | 67.21 | short | 0.71 |
| 2026-05-24T14:50:02Z | T-2287651-1779625804088 | uncategorized | 0.4000 | 0.0200 | -58.42 | stop_loss |  | 61.50 | short | 0.10 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T15:50:01Z | T-2227129-1779224404592 | uncategorized | 0.2900 | 0.4800 | 40.38 | take_profit |  | 61.63 | medium | 4.78 |
| 2026-05-24T15:55:02Z | T-1809560-1779577204027 | uncategorized | 0.1680 | 0.0310 | -48.36 | stop_loss |  | 59.31 | medium | 0.70 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T15:55:02Z | T-2227158-1779620403996 | uncategorized | 0.5000 | 0.7600 | 32.50 | take_profit |  | 62.50 | short | 0.20 |
| 2026-05-24T16:00:01Z | T-2227198-1779226204013 | uncategorized | 0.5300 | 0.7600 | 26.47 | take_profit |  | 60.99 | medium | 4.77 |
| 2026-05-24T16:00:01Z | T-2227126-1779309004485 | uncategorized | 0.5100 | 0.0700 | -55.02 | stop_loss |  | 63.77 | medium | 3.81 |
| 2026-05-24T16:00:02Z | T-2227193-1779602403485 | uncategorized | 0.2500 | 0.0500 | -51.61 | stop_loss |  | 64.52 | short | 0.42 |
| 2026-05-24T16:20:02Z | T-2227313-1779616804049 | uncategorized | 0.4800 | 0.0200 | -60.45 | stop_loss |  | 63.08 | short | 0.26 |
| 2026-05-24T16:20:02Z | T-2227193-1779638403705 | uncategorized | 0.0600 | 0.0100 | -48.53 | stop_loss |  | 58.24 | short | 0.01 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T16:30:03Z | T-2227286-1779604204419 | uncategorized | 0.1100 | 0.4370 | 198.73 | take_profit |  | 66.85 | short | 0.42 |
| 2026-05-24T16:35:02Z | T-2227098-1779600603581 | uncategorized | 0.4000 | 0.0800 | -52.45 | stop_loss |  | 65.56 | short | 0.46 |
| 2026-05-24T16:40:01Z | T-2324252-1779571804012 | uncategorized | 0.6200 | 0.1200 | -55.31 | stop_loss |  | 68.58 | short | 0.80 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-24T16:45:01Z | T-2227393-1779310803840 | uncategorized | 0.5300 | 0.7600 | 26.81 | take_profit |  | 61.79 | medium | 3.82 |
| 2026-05-24T16:45:02Z | T-2227281-1779609604098 | uncategorized | 0.7500 | 0.1100 | -55.35 | stop_loss |  | 64.87 | short | 0.36 |
| 2026-05-24T16:45:02Z | T-2227397-1779638403705 | uncategorized | 0.2700 | 0.0400 | -52.71 | stop_loss |  | 61.88 | short | 0.03 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
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


## Human notes
_(no se toca por automatización)_
