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
| 2026-05-25T10:41:42Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-freecs-win-the-lck-2026-season-playoffs | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-the-us-invade-iran-before-2027 | geopolitics | 0.2000 | M1 | memoria: slug prefix match; same category (score 0.70) |
| 2026-05-25T10:41:42Z | insurrection-act-invoked-by-december-31-184-556 | other | 0.3000 | P3_low_absolute_liquidity | liquidity $4608 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-alberta-join-the-us | other | 0.0480 | P0_floor | price floor: 0.0480 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | sbf-released-from-custody-in-2026 | other | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-there-be-between-8-and-10-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.0720 | P0_floor | price floor: 0.0720 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-the-us-invade-mexico-in-2026 | geopolitics | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-israel-strike-11-countries-in-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-macky-sall-be-the-next-secretary-general-of-the-united-nations | other | 0.1180 | P3_low_absolute_liquidity | liquidity $3093 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-germany-recognize-palestine-before-2027 | executive-action | 0.0880 | P0_floor | price floor: 0.0880 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | anduril-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $3058 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-there-be-between-17-and-19-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3479 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-zelenskyy-and-putin-meet-next-in-ukraine | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-there-be-exactly-0-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026 | weather-natural | 0.7100 | P3_low_absolute_liquidity | liquidity $1741 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-brion-win-the-lck-2026-season-playoffs | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-israel-strike-6-countries-in-2026 | geopolitics | 0.0790 | P0_floor | price floor: 0.0790 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0530 | P0_floor | price floor: 0.0530 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-magomed-ankalaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-nongshim-redforce-win-the-lck-2026-season-playoffs | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-the-feds-lower-bound-reach-0pt25-or-lower-before-2027-173-921-916-764-754-593-935 | other | 0.0510 | P0_floor | price floor: 0.0510 < 0.100 (horizon 219 d) |
| 2026-05-25T10:41:42Z | will-lerone-murphy-fight-alexander-volkanovski-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 220 d) |
| 2026-05-25T10:41:42Z | hyperbeat-fdv-above-300m-one-day-after-launch | crypto-launch | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 220 d) |
| 2026-05-25T10:41:42Z | solstice-fdv-above-100m-one-day-after-launch | crypto-launch | 0.9300 | P4_pre_event | pre-event slug + 220 d to resolution (>=7 threshold) |
| 2026-05-25T10:41:42Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.8700 | P3_low_absolute_liquidity | liquidity $787 < absolute min $5000 |
| 2026-05-25T10:41:42Z | over-1b-raised-on-coinbase-in-2026 | other | 0.1700 | P3_low_absolute_liquidity | liquidity $1203 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 220 d) |
| 2026-05-25T10:41:42Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 220 d) |
| 2026-05-25T10:41:42Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1363 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-usdc-hit-50-of-usdt-market-cap-by-december-31-2026 | other | 0.4500 | P3_low_absolute_liquidity | liquidity $218 < absolute min $5000 |
| 2026-05-25T10:41:42Z | extended-fdv-above-800m-one-day-after-launch-299-255-178 | crypto-launch | 0.1000 | P4_pre_event | pre-event slug + 220 d to resolution (>=7 threshold) |
| 2026-05-25T10:41:42Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6900 | P3_low_absolute_liquidity | liquidity $3044 < absolute min $5000 |
| 2026-05-25T10:41:42Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | crypto-launch | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 220 d) |
| 2026-05-25T10:41:42Z | will-lighter-reach-8-before-2027 | other | 0.1130 | P3_low_absolute_liquidity | liquidity $1479 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-opensea-launch-a-token-by-december-31-2026 | crypto-launch | 0.6780 | P3_low_absolute_liquidity | liquidity $3780 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-anthropics-valuation-hit-high-1pt75t-by-december-31-328 | other | 0.6000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-anthropics-valuation-hit-high-2pt0t-by-december-31-922 | other | 0.4500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-mitch-mcconnell-resign-from-the-senate-before-his-term-ends | executive-action | 0.2300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | sports-season | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | sports-season | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-houston-texans-win-the-2027-nfl-afc-championship-334 | sports-season | 0.1000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-miami-dolphins-win-the-2027-nfl-afc-championship-357 | sports-season | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | sports-season | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-atlanta-falcons-win-the-2027-nfl-nfc-championship-312 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 244 d) |
| 2026-05-25T10:41:42Z | will-saquon-barkley-win-the-2026-nfl-mvp | sports-season | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon 265 d) |
| 2026-05-25T10:41:42Z | will-baker-mayfield-win-the-2026-nfl-mvp | sports-season | 0.0570 | P0_floor | price floor: 0.0570 < 0.100 (horizon 265 d) |
| 2026-05-25T10:41:42Z | will-maurcio-ruffy-fight-charles-oliveira-next-853-769 | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 279 d) |
| 2026-05-25T10:41:42Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 310 d) |
| 2026-05-25T10:41:42Z | will-the-cleveland-browns-win-the-2027-nfl-league-championship | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 310 d) |
| 2026-05-25T10:41:42Z | will-mathilde-panot-win-the-2027-french-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 339 d) |
| 2026-05-25T10:41:42Z | will-sacramento-kings-win-the-2027-nba-finals | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 401 d) |
| 2026-05-25T10:41:42Z | will-utah-jazz-win-the-2027-nba-finals | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 401 d) |
| 2026-05-25T10:41:42Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 584 d) |
| 2026-05-25T10:41:42Z | spacex-ipo-closing-market-cap-above-3pt2t | crypto-launch | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 584 d) |
| 2026-05-25T10:41:42Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.3100 | P4_pre_event | pre-event slug + 585 d to resolution (>=7 threshold) |
| 2026-05-25T10:41:42Z | base-fdv-above-2b-one-day-after-launch | crypto-launch | 0.8000 | P4_pre_event | pre-event slug + 585 d to resolution (>=7 threshold) |
| 2026-05-25T10:41:42Z | will-gavin-newsom-win-the-2028-us-presidential-election | elections | 0.1640 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-elon-musk-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-mark-kelly-win-the-2028-democratic-presidential-nomination-479 | elections | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-eric-trump-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | elections | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-joe-kent-win-the-2028-republican-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-jd-vance-win-the-2028-us-presidential-election | elections | 0.1900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | elections | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-katie-britt-win-the-2028-republican-presidential-nomination | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-wes-moore-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-matt-gaetz-win-the-2028-republican-presidential-nomination | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-john-fetterman-win-the-2028-democratic-presidential-nomination-941 | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 896 d) |
| 2026-05-25T10:41:42Z | us-announces-new-iran-agreementceasefire-extension-by-june-7 | geopolitics | 0.5600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-rebecca-shepherd-win-the-2026-makerfield-by-election | elections | 0.0510 | P0_floor | price floor: 0.0510 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | will-comcast-close-warner-bros-acquisition | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | will-the-iran-ceasefire-continue-through-june-7-849 | geopolitics | 0.9000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-netflix-close-warner-bros-acquisition | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | will-variational-launch-a-token-by-december-31-2026-577-621 | crypto-launch | 0.8200 | P3_low_absolute_liquidity | liquidity $1338 < absolute min $5000 |
| 2026-05-25T10:41:42Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.950 |
| 2026-05-25T10:41:42Z | will-the-iran-ceasefire-continue-through-june-15-136-565 | geopolitics | 0.8700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-the-iran-ceasefire-continue-through-may-27-496 | geopolitics | 0.9680 | P0_ceiling | price ceiling: 0.9680 > 0.950 |
| 2026-05-25T10:41:42Z | us-announces-new-iran-agreementceasefire-extension-by-may-31 | geopolitics | 0.3500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | us-announces-new-iran-agreementceasefire-extension-by-may-25 | geopolitics | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | us-announces-new-iran-agreementceasefire-extension-by-may-26 | geopolitics | 0.1300 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-25T10:41:42Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon -1 d) |
| 2026-05-25T10:41:42Z | will-the-iran-ceasefire-continue-through-may-31-654-633 | geopolitics | 0.9400 | M1 | memoria: exact slug match (score 1.00) |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
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
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T03:25:01Z | T-2097884-1779678004140 | uncategorized | 0.3100 | 0.9700 | 112.00 | take_profit |  | 52.60 | short | 0.02 |
| 2026-05-25T05:15:02Z | T-2002531-1779652804112 | uncategorized | 0.0770 | 0.1110 | 22.90 | take_profit |  | 51.86 | medium | 0.39 |
| 2026-05-25T05:25:02Z | T-2313547-1779604204419 | uncategorized | 0.1640 | 0.2350 | 28.36 | take_profit |  | 65.51 | short | 0.95 |
| 2026-05-25T06:05:02Z | T-2321929-1779679804173 | uncategorized | 0.2200 | 0.3300 | 28.24 | take_profit |  | 56.47 | short | 0.11 |
| 2026-05-25T06:15:01Z | T-2334096-1779517803519 | uncategorized | 0.9200 | 0.9990 | 4.94 | market_closed |  | 57.50 | short | 1.99 |
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


## Human notes
_(no se toca por automatización)_
