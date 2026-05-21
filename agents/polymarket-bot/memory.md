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
| 2026-05-21T18:30:04Z | will-spacex-have-200-or-more-launches-in-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-petr-yan-fight-dominick-cruz-next | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-there-be-between-11-and-13-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | uncategorized | 0.2900 | P3_low_absolute_liquidity | liquidity $2338 < absolute min $5000 |
| 2026-05-21T18:30:04Z | will-clavicular-be-named-peoples-sexiest-man-alive-in-2026-399 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | trump-declares-election-interference-national-emergency | uncategorized | 0.2900 | P3_low_absolute_liquidity | liquidity $4880 < absolute min $5000 |
| 2026-05-21T18:30:04Z | will-hassan-khomeini-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-zelenskyy-and-putin-meet-next-in-russia | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-7-8-spacex-starship-launches-successfully-reach-space-in-2026 | uncategorized | 0.0740 | P0_floor | price floor: 0.0740 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-mohammad-khatami-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-kendrick-lamar-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-china-invade-taiwan-before-2027 | uncategorized | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-microstrategy-announce-bankruptcy-before-2027 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-ali-motahari-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-israel-strike-9-countries-in-2026 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | ramp-ipo-before-2027 | uncategorized | 0.1100 | P3_low_absolute_liquidity | liquidity $2100 < absolute min $5000 |
| 2026-05-21T18:30:04Z | will-mary-fitzpatrick-win-the-2026-dublin-central-by-election | uncategorized | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-beyonc-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-zelenskyy-and-putin-meet-next-in-kazakhstan-485-344 | uncategorized | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-zelenskyy-and-putin-meet-next-in-saudi-arabia | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-larry-page-be-richest-person-on-december-31 | uncategorized | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 223 d) |
| 2026-05-21T18:30:04Z | will-ole-gunnar-solskjr-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | uncategorized | 0.0540 | P0_floor | price floor: 0.0540 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | over-1b-raised-on-coinbase-in-2026 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $1817 < absolute min $5000 |
| 2026-05-21T18:30:04Z | puffpaw-fdv-above-50m-one-day-after-launch | uncategorized | 0.7300 | P4_pre_event | pre-event slug + 224 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | okx-ipo-in-2026 | uncategorized | 0.1300 | P4_pre_event | pre-event slug + 224 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | will-stablecoins-hit-500b-before-2027 | uncategorized | 0.1200 | P3_low_absolute_liquidity | liquidity $3080 < absolute min $5000 |
| 2026-05-21T18:30:04Z | puffpaw-fdv-above-300m-one-day-after-launch | uncategorized | 0.0970 | P0_floor | price floor: 0.0970 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.2800 | P3_low_absolute_liquidity | liquidity $1610 < absolute min $5000 |
| 2026-05-21T18:30:04Z | will-loopscale-launch-a-token-by-december-31-2026 | uncategorized | 0.4200 | P3_low_absolute_liquidity | liquidity $891 < absolute min $5000 |
| 2026-05-21T18:30:04Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0630 | P0_floor | price floor: 0.0630 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-ethereum-reach-6500-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | usdc-depeg-by-december-31-968 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.7000 | P3_low_absolute_liquidity | liquidity $3098 < absolute min $5000 |
| 2026-05-21T18:30:04Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | uncategorized | 0.1600 | P4_pre_event | pre-event slug + 224 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | uncategorized | 0.1800 | P4_pre_event | pre-event slug + 224 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | puffpaw-fdv-above-400m-one-day-after-launch | uncategorized | 0.0660 | P0_floor | price floor: 0.0660 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | major-cex-insolvent-in-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-ethereum-reach-7500-by-december-31-2026 | uncategorized | 0.0430 | P0_floor | price floor: 0.0430 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-tempo-launch-a-token-by-june-30-2026 | uncategorized | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-opensea-launch-a-token-by-december-31-2026 | uncategorized | 0.6590 | P3_low_absolute_liquidity | liquidity $4193 < absolute min $5000 |
| 2026-05-21T18:30:04Z | altcoin-market-cap-dip-to-150b-before-2027 | uncategorized | 0.7600 | P3_low_absolute_liquidity | liquidity $714 < absolute min $5000 |
| 2026-05-21T18:30:04Z | hyperbeat-fdv-above-300m-one-day-after-launch | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 224 d) |
| 2026-05-21T18:30:04Z | will-hibachi-launch-a-token-by-december-31-2026 | uncategorized | 0.6000 | P3_low_absolute_liquidity | liquidity $938 < absolute min $5000 |
| 2026-05-21T18:30:04Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-miami-dolphins-win-the-2027-nfl-afc-championship-357 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 248 d) |
| 2026-05-21T18:30:04Z | will-maurcio-ruffy-fight-charles-oliveira-next-853-769 | uncategorized | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 283 d) |
| 2026-05-21T18:30:04Z | will-marine-tondelier-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 343 d) |
| 2026-05-21T18:30:04Z | will-olivier-faure-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 343 d) |
| 2026-05-21T18:30:04Z | will-mathilde-panot-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 343 d) |
| 2026-05-21T18:30:04Z | will-ubs-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offering-884 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 588 d) |
| 2026-05-21T18:30:04Z | spacex-ipo-closing-market-cap-above-3t-337-436 | uncategorized | 0.2000 | P4_pre_event | pre-event slug + 588 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | will-spacexs-public-ticker-be-another-ticker-745 | uncategorized | 0.9800 | P0_ceiling | price ceiling: 0.9800 > 0.950 |
| 2026-05-21T18:30:04Z | will-spacex-not-ipo-by-december-31-2027-481-338 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 588 d) |
| 2026-05-21T18:30:04Z | spacex-ipo-closing-market-cap-above-2t-649-783 | uncategorized | 0.7100 | P4_pre_event | pre-event slug + 588 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | will-spacexs-public-ticker-be-x-333 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 588 d) |
| 2026-05-21T18:30:04Z | base-fdv-above-6b-one-day-after-launch | uncategorized | 0.4700 | P4_pre_event | pre-event slug + 589 d to resolution (>=7 threshold) |
| 2026-05-21T18:30:04Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-ruben-gallego-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-ivanka-trump-win-the-2028-republican-presidential-nomination | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-glenn-youngkin-win-the-2028-republican-presidential-nomination | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-dwayne-the-rock-johnson-win-the-2028-us-presidential-election | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-michelle-obama-win-the-2028-democratic-presidential-nomination-777 | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-jon-stewart-win-the-2028-democratic-presidential-nomination-518 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-pete-hegseth-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-glenn-youngkin-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-john-fetterman-win-the-2028-democratic-presidential-nomination-941 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-donald-trump-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 900 d) |
| 2026-05-21T18:30:04Z | will-the-iran-ceasefire-continue-through-may-22 | uncategorized | 0.9400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-21T18:30:04Z | will-the-iran-ceasefire-continue-through-may-24 | uncategorized | 0.8600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-21T18:30:04Z | will-axiom-launch-a-token-by-december-31-2026 | uncategorized | 0.2700 | P3_low_absolute_liquidity | liquidity $490 < absolute min $5000 |
| 2026-05-21T18:30:04Z | will-the-iran-ceasefire-continue-through-may-20 | uncategorized | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.950 |
| 2026-05-21T18:30:04Z | will-the-republican-party-hold-exactly-55-senate-seats-after-the-2026-midterm-elections | uncategorized | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon -1 d) |
| 2026-05-21T18:30:04Z | will-the-iran-ceasefire-continue-through-may-27 | uncategorized | 0.7900 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-21T18:30:04Z | will-the-iran-ceasefire-continue-through-may-21 | uncategorized | 0.9830 | P0_ceiling | price ceiling: 0.9830 > 0.950 |
| 2026-05-21T18:30:04Z | will-the-iran-ceasefire-continue-through-may-31 | uncategorized | 0.7400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-21T18:30:04Z | will-rebecca-shepherd-win-the-2026-makerfield-by-election | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon -1 d) |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
| 2026-05-19T11:30:00Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0065 | 0.0065 | 0.00 | manual_close_v1_audit · ejemplo long-tail absurdo aprobado por v1 |
| 2026-05-19T11:30:00Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | uncategorized | 0.0075 | 0.0075 | 0.00 | manual_close_v1_audit · ejemplo long-tail absurdo aprobado por v1 |
| 2026-05-19T18:10:02Z | T-608399-1779213603965 | uncategorized | 0.0660 | 0.0070 | -40.23 | stop_loss |
| 2026-05-19T19:55:02Z | T-566188-1779210004342 | uncategorized | 0.1350 | 0.0200 | -63.92 | stop_loss |
| 2026-05-19T20:10:01Z | T-2170973-1779206389814 | uncategorized | 0.6050 | 0.1200 | -120.25 | stop_loss |
| 2026-05-19T20:25:02Z | T-2172006-1779220804587 | uncategorized | 0.6300 | 0.0600 | -49.14 | stop_loss |
| 2026-05-19T20:35:02Z | T-2170967-1779222604158 | uncategorized | 0.9300 | 0.0010 | -53.23 | stop_loss |
| 2026-05-19T20:45:01Z | T-2170982-1779206389814 | uncategorized | 0.2650 | 0.0500 | -121.70 | stop_loss |
| 2026-05-19T20:45:02Z | T-2170981-1779215403605 | uncategorized | 0.2900 | 0.0400 | -40.82 | stop_loss |
| 2026-05-19T21:15:02Z | T-688127-1779224404592 | uncategorized | 0.1080 | 0.0150 | -52.01 | stop_loss |
| 2026-05-19T21:20:02Z | T-2170981-1779224404592 | uncategorized | 0.1600 | 0.0100 | -58.96 | stop_loss |
| 2026-05-19T22:00:02Z | T-688127-1779226204013 | uncategorized | 0.1080 | 0.0160 | -50.91 | stop_loss |
| 2026-05-19T22:05:02Z | T-688127-1779228007314 | uncategorized | 0.1070 | 0.0190 | -45.48 | stop_loss |
| 2026-05-19T22:35:01Z | T-789540-1779206389814 | uncategorized | 0.4150 | 0.0800 | -121.08 | stop_loss |
| 2026-05-19T22:35:02Z | T-2053878-1779228007314 | uncategorized | 0.3200 | 0.0120 | -56.55 | stop_loss |
| 2026-05-19T22:55:02Z | T-688127-1779229804249 | uncategorized | 0.1080 | 0.0160 | -48.78 | stop_loss |
| 2026-05-19T23:05:02Z | T-2036170-1779231604302 | uncategorized | 0.2000 | 0.0200 | -50.20 | stop_loss |
| 2026-05-19T23:15:02Z | T-789540-1779231604302 | uncategorized | 0.1700 | 0.0100 | -53.57 | stop_loss |
| 2026-05-20T01:30:02Z | T-2299145-1779237003760 | uncategorized | 0.1900 | 0.0300 | -45.36 | stop_loss |
| 2026-05-20T02:30:03Z | T-2053882-1779233403338 | uncategorized | 0.6100 | 0.0300 | -52.63 | stop_loss |
| 2026-05-20T02:45:02Z | T-2053012-1779244203980 | uncategorized | 0.1700 | 0.0100 | -52.60 | stop_loss |
| 2026-05-20T04:05:02Z | T-2053890-1779247803922 | uncategorized | 0.1500 | 0.0100 | -53.46 | stop_loss |
| 2026-05-20T04:05:02Z | T-2053888-1779247803922 | uncategorized | 0.3800 | 0.0500 | -48.75 | stop_loss |
| 2026-05-20T06:50:02Z | T-2237150-1779194451845 | uncategorized | 0.0605 | 0.0070 | -132.64 | stop_loss |
| 2026-05-20T13:50:02Z | T-2251419-1779267604555 | uncategorized | 0.2200 | 0.0400 | -44.16 | stop_loss |  | 53.97 | short | 0.20 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T14:25:01Z | T-2237155-1779206389814 | uncategorized | 0.2585 | 0.3660 | 62.38 | take_profit | legacy | 150.00 | short | 0.93 |
| 2026-05-20T14:25:02Z | T-2251419-1779285604469 | uncategorized | 0.0600 | 0.1520 | 81.40 | take_profit |  | 53.09 | short | 0.02 |
| 2026-05-20T14:40:01Z | T-2237156-1779206389814 | uncategorized | 0.1690 | 0.3090 | 124.26 | take_profit | legacy | 150.00 | short | 0.94 |
| 2026-05-20T14:40:02Z | T-2111640-1779211529037 | uncategorized | 0.1000 | 0.1400 | 23.88 | take_profit |  | 59.71 | medium | 0.89 |
| 2026-05-20T14:40:02Z | T-2237153-1779251403658 | uncategorized | 0.2190 | 0.0230 | -48.45 | stop_loss |  | 54.13 | short | 0.42 |
| 2026-05-20T14:50:02Z | T-2241872-1779219003841 | uncategorized | 0.2100 | 0.0400 | -39.82 | stop_loss |  | 49.19 | short | 0.81 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T15:10:01Z | T-2119227-1779206389814 | uncategorized | 0.0550 | 0.0770 | 60.00 | take_profit | legacy | 150.00 | short | 0.97 |
| 2026-05-20T15:15:02Z | T-2251419-1779287404273 | uncategorized | 0.1560 | 0.0300 | -46.67 | stop_loss |  | 57.78 | short | 0.03 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T15:25:01Z | T-1919425-1779213603965 | uncategorized | 0.1300 | 0.2100 | 28.83 | take_profit |  | 46.85 | medium | 0.89 |
| 2026-05-20T15:25:02Z | T-1962237-1779217203920 | uncategorized | 0.2800 | 0.4000 | 18.72 | take_profit |  | 43.68 | medium | 0.85 |
| 2026-05-20T15:45:02Z | T-2251419-1779291005492 | uncategorized | 0.1040 | 0.0070 | -61.83 | stop_loss |  | 66.30 | short | 0.01 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T15:50:03Z | T-2241872-1779289205006 | uncategorized | 0.0500 | 0.0700 | 24.62 | take_profit |  | 61.55 | short | 0.03 |
| 2026-05-20T17:25:01Z | T-2111564-1779213603965 | uncategorized | 0.1500 | 0.2200 | 21.43 | take_profit |  | 45.92 | medium | 0.98 |
| 2026-05-20T18:10:02Z | T-2251366-1779287404273 | uncategorized | 0.9500 | 0.9990 | 3.04 | market_closed |  | 58.96 | short | 0.15 |
| 2026-05-20T18:15:02Z | T-2251416-1779222604158 | uncategorized | 0.8000 | 0.9990 | 12.73 | market_closed |  | 51.18 | short | 0.91 |
| 2026-05-20T19:15:02Z | T-2237156-1779291005492 | uncategorized | 0.2100 | 0.2950 | 25.77 | take_profit |  | 63.67 | short | 0.16 |
| 2026-05-20T19:35:01Z | T-2237157-1779206389814 | uncategorized | 0.0645 | 0.0960 | 73.26 | take_profit | legacy | 150.00 | short | 1.15 |
| 2026-05-20T20:05:02Z | T-2192106-1779247803922 | uncategorized | 0.5900 | 0.9400 | 32.63 | take_profit |  | 55.01 | short | 0.69 |
| 2026-05-20T20:05:02Z | T-2192104-1779289205006 | uncategorized | 0.1700 | 0.0100 | -60.32 | stop_loss |  | 64.09 | short | 0.21 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T20:20:01Z | T-1972137-1779211803817 | uncategorized | 0.2100 | 0.3100 | 25.19 | take_profit |  | 52.89 | medium | 1.12 |
| 2026-05-20T20:30:02Z | T-2192105-1779307203835 | uncategorized | 0.1700 | 0.0100 | -61.75 | stop_loss |  | 65.60 | short | 0.02 |
| 2026-05-20T20:35:02Z | T-2192105-1779309004485 | uncategorized | 0.0500 | 0.0100 | -54.21 | stop_loss |  | 67.76 | short | 0.00 |
| 2026-05-20T20:40:02Z | T-2273922-1779301804949 | uncategorized | 0.0860 | 0.0030 | -60.90 | stop_loss |  | 63.10 | short | 0.09 |
| 2026-05-20T20:45:02Z | T-2273938-1779292804136 | uncategorized | 0.0770 | 0.0020 | -60.10 | stop_loss |  | 61.70 | short | 0.20 |
| 2026-05-20T20:55:02Z | T-2216585-1779289205006 | uncategorized | 0.5000 | 0.0800 | -52.76 | stop_loss |  | 62.81 | short | 0.25 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T21:40:02Z | T-2273931-1779298204126 | uncategorized | 0.8600 | 0.9990 | 9.99 | market_closed |  | 61.82 | short | 0.17 |
| 2026-05-21T00:05:02Z | T-2068164-1779305404625 | uncategorized | 0.9500 | 0.9990 | 3.22 | market_closed |  | 62.39 | short | 0.19 |
| 2026-05-21T00:40:02Z | T-2237157-1779309004485 | uncategorized | 0.0920 | 0.1310 | 28.15 | take_profit |  | 66.40 | short | 0.17 |
| 2026-05-21T01:40:01Z | T-2237154-1779210004342 | uncategorized | 0.2815 | 0.0510 | -69.36 | stop_loss |  | 84.71 | short | 1.36 |
| 2026-05-21T01:45:02Z | T-2061753-1779325203995 | uncategorized | 0.3200 | 0.0600 | -49.08 | stop_loss |  | 60.41 | short | 0.03 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-21T02:10:02Z | T-2060498-1779323403696 | uncategorized | 0.5200 | 0.8800 | 41.34 | take_profit |  | 59.71 | short | 0.07 |
| 2026-05-21T02:30:02Z | T-2060475-1779328804289 | uncategorized | 0.2300 | 0.0200 | -54.54 | stop_loss |  | 59.73 | short | 0.02 |
| 2026-05-21T02:30:02Z | T-2061753-1779328804289 | uncategorized | 0.0500 | 0.0010 | -57.37 | stop_loss |  | 58.54 | short | 0.02 |
| 2026-05-21T02:35:01Z | T-2241872-1779292804136 | uncategorized | 0.0700 | 0.0140 | -50.37 | stop_loss |  | 62.96 | short | 0.44 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-21T03:55:02Z | T-2237157-1779330604178 | uncategorized | 0.2300 | 0.3330 | 25.60 | take_profit |  | 57.16 | short | 0.06 |
| 2026-05-21T03:55:02Z | T-2237158-1779332403652 | uncategorized | 0.0590 | 0.1000 | 39.10 | take_profit |  | 56.27 | short | 0.04 |
| 2026-05-21T04:10:02Z | T-2060514-1779330604178 | uncategorized | 0.7000 | 0.9990 | 24.91 | take_profit |  | 58.32 | short | 0.07 |
| 2026-05-21T05:30:02Z | T-2060481-1779330604178 | uncategorized | 0.9500 | 0.9990 | 3.07 | market_closed |  | 59.51 | short | 0.12 |
| 2026-05-21T05:55:01Z | T-2237156-1779309004485 | uncategorized | 0.3240 | 0.4560 | 26.51 | take_profit |  | 65.08 | short | 0.39 |
| 2026-05-21T06:45:02Z | T-2237155-1779291005492 | uncategorized | 0.3600 | 0.0720 | -49.92 | stop_loss |  | 62.40 | short | 0.64 |
| 2026-05-21T07:00:02Z | T-2296151-1779301804949 | uncategorized | 0.1600 | 0.0300 | -50.24 | stop_loss |  | 61.84 | medium | 0.52 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-21T07:25:02Z | T-2296151-1779346803793 | uncategorized | 0.0500 | 0.0700 | 22.78 | take_profit |  | 56.95 | short | 0.02 |
| 2026-05-21T11:20:02Z | T-2296151-1779348603737 | uncategorized | 0.0900 | 0.1300 | 25.51 | take_profit |  | 57.40 | short | 0.16 |
| 2026-05-21T12:40:02Z | T-2261232-1779341403901 | uncategorized | 0.4600 | 0.0900 | -46.76 | stop_loss |  | 58.13 | short | 0.30 |
| 2026-05-21T13:50:02Z | T-2261232-1779368404392 | uncategorized | 0.1300 | 0.0200 | -48.22 | stop_loss |  | 56.98 | short | 0.03 |
| 2026-05-21T14:45:02Z | T-2237155-1779346803793 | uncategorized | 0.0680 | 0.0110 | -48.71 | stop_loss |  | 58.11 | short | 0.32 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-21T17:30:01Z | T-1971905-1779210004342 | uncategorized | 0.2950 | 0.4300 | 31.04 | take_profit |  | 67.83 | long | 2.02 |
| 2026-05-21T17:30:02Z | T-2297854-1779289205006 | uncategorized | 0.0690 | 0.1110 | 36.72 | take_profit |  | 60.32 | short | 1.10 |
| 2026-05-21T17:30:02Z | T-2119227-1779291005492 | uncategorized | 0.0690 | 0.1150 | 43.31 | take_profit |  | 64.97 | short | 1.08 |
| 2026-05-21T18:40:02Z | T-2068139-1779375604743 | uncategorized | 0.0550 | 0.0050 | -50.08 | stop_loss |  | 55.08 | short | 0.15 |
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


## Human notes
_(no se toca por automatización)_
