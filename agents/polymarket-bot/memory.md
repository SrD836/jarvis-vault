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
| 2026-05-22T12:30:03Z | will-spacex-have-200-or-more-launches-in-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-sadegh-mahsouli-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | ledger-ipo-before-2027 | uncategorized | 0.1300 | P3_low_absolute_liquidity | liquidity $1007 < absolute min $5000 |
| 2026-05-22T12:30:03Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-israel-strike-9-countries-in-2026 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-ali-motahari-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-kendrick-lamar-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | databricks-ipo-before-2027 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $2513 < absolute min $5000 |
| 2026-05-22T12:30:03Z | will-gideon-saar-be-the-next-prime-minister-of-israel | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-gillian-sherratt-win-the-2026-dublin-central-by-election | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-jeff-bezos-be-richest-person-on-december-31-243 | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | freddie-mac-ipo-before-2027 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | vanta-ipo-before-2027 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | deel-ipo-before-2027 | uncategorized | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-israel-strike-11-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-2026-rank-as-the-sixth-hottest-year-on-record-or-lower | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-gholam-ali-haddad-adel-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | anysphere-cursor-ipo-before-2027 | uncategorized | 0.0560 | P0_floor | price floor: 0.0560 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-germany-recognize-palestine-before-2027 | uncategorized | 0.0880 | P0_floor | price floor: 0.0880 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-reza-pahlavi-lead-iran-in-2026 | uncategorized | 0.0460 | P0_floor | price floor: 0.0460 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-zelenskyy-and-putin-meet-next-in-ukraine | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-fearx-win-the-lck-2026-season-playoffs | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 222 d) |
| 2026-05-22T12:30:03Z | will-ole-gunnar-solskjr-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | uncategorized | 0.1600 | P4_pre_event | pre-event slug + 223 d to resolution (>=7 threshold) |
| 2026-05-22T12:30:03Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | uncategorized | 0.3460 | P3_low_absolute_liquidity | liquidity $2184 < absolute min $5000 |
| 2026-05-22T12:30:03Z | over-1b-raised-on-coinbase-in-2026 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $1177 < absolute min $5000 |
| 2026-05-22T12:30:03Z | puffpaw-fdv-above-200m-one-day-after-launch | uncategorized | 0.2700 | P4_pre_event | pre-event slug + 223 d to resolution (>=7 threshold) |
| 2026-05-22T12:30:03Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | uncategorized | 0.1800 | P4_pre_event | pre-event slug + 223 d to resolution (>=7 threshold) |
| 2026-05-22T12:30:03Z | usdc-depeg-by-december-31-968 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | major-cex-insolvent-in-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.7000 | P3_low_absolute_liquidity | liquidity $3279 < absolute min $5000 |
| 2026-05-22T12:30:03Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | puffpaw-fdv-above-300m-one-day-after-launch | uncategorized | 0.0970 | P0_floor | price floor: 0.0970 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | will-fomo-launch-a-token-by-december-31-2026 | uncategorized | 0.3800 | P3_low_absolute_liquidity | liquidity $947 < absolute min $5000 |
| 2026-05-22T12:30:03Z | will-ethereum-reach-6500-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | will-ethereum-reach-6000-by-december-31-2026 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 223 d) |
| 2026-05-22T12:30:03Z | altcoin-market-cap-dip-to-150b-before-2027 | uncategorized | 0.7600 | P3_low_absolute_liquidity | liquidity $877 < absolute min $5000 |
| 2026-05-22T12:30:03Z | will-miami-dolphins-win-the-2027-nfl-afc-championship-357 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-pittsburgh-steelers-win-the-2027-nfl-afc-championship-898 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.0820 | P0_floor | price floor: 0.0820 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-houston-texans-win-the-2027-nfl-afc-championship-334 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 247 d) |
| 2026-05-22T12:30:03Z | will-maurcio-ruffy-fight-charles-oliveira-next-853-769 | uncategorized | 0.0390 | P0_floor | price floor: 0.0390 < 0.100 (horizon 282 d) |
| 2026-05-22T12:30:03Z | will-clmentine-autain-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 342 d) |
| 2026-05-22T12:30:03Z | will-olivier-faure-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 342 d) |
| 2026-05-22T12:30:03Z | will-mathilde-panot-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 342 d) |
| 2026-05-22T12:30:03Z | will-sacramento-kings-win-the-2027-nba-finals | uncategorized | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 404 d) |
| 2026-05-22T12:30:03Z | will-utah-jazz-win-the-2027-nba-finals | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 404 d) |
| 2026-05-22T12:30:03Z | will-jpmorgan-chase-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-off... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 587 d) |
| 2026-05-22T12:30:03Z | will-deutsche-bank-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offe... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 587 d) |
| 2026-05-22T12:30:03Z | will-anthropics-market-cap-be-less-than-100b-at-market-close-on-ipo-day-by-december-31-2027 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 587 d) |
| 2026-05-22T12:30:03Z | will-spacexs-public-ticker-be-spc-937 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 587 d) |
| 2026-05-22T12:30:03Z | will-ubs-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offering-884 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 587 d) |
| 2026-05-22T12:30:03Z | spacex-ipo-closing-market-cap-above-3t-337-436 | uncategorized | 0.1600 | P4_pre_event | pre-event slug + 587 d to resolution (>=7 threshold) |
| 2026-05-22T12:30:03Z | will-spacexs-market-cap-be-between-500b-and-600b-at-market-close-on-ipo-day | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 587 d) |
| 2026-05-22T12:30:03Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-john-fetterman-win-the-2028-democratic-presidential-nomination-941 | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-tim-walz-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-kristi-noem-win-the-2028-republican-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-zohran-mamdani-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-donald-trump-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-rahm-emanuel-win-the-2028-democratic-presidential-nomination-299 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-ruben-gallego-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-joe-kent-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-pete-hegseth-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-thomas-massie-win-the-2028-us-presidential-election | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-mrbeast-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 899 d) |
| 2026-05-22T12:30:03Z | will-the-republican-party-hold-exactly-55-senate-seats-after-the-2026-midterm-elections | uncategorized | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon -1 d) |
| 2026-05-22T12:30:03Z | will-rebecca-shepherd-win-the-2026-makerfield-by-election | uncategorized | 0.0680 | P0_floor | price floor: 0.0680 < 0.100 (horizon -1 d) |
| 2026-05-22T12:30:03Z | will-variational-launch-a-token-by-december-31-2026-577-621 | uncategorized | 0.8200 | P3_low_absolute_liquidity | liquidity $2027 < absolute min $5000 |
| 2026-05-22T12:30:03Z | will-the-iran-ceasefire-continue-through-may-22 | uncategorized | 0.9780 | P0_ceiling | price ceiling: 0.9780 > 0.950 |
| 2026-05-22T12:30:03Z | will-the-iran-ceasefire-continue-through-may-27 | uncategorized | 0.8200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-22T12:30:03Z | will-the-iran-ceasefire-continue-through-may-31 | uncategorized | 0.7400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-22T12:30:03Z | will-venezuela-recognize-israel-by-june-30-115 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon -1 d) |
| 2026-05-22T12:30:03Z | will-the-iran-ceasefire-continue-through-may-24 | uncategorized | 0.9100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-22T12:30:03Z | will-the-iran-ceasefire-continue-through-may-21 | uncategorized | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.950 |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
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
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-21T19:35:01Z | T-2261230-1779310803840 | uncategorized | 0.9500 | 0.9990 | 3.25 | market_closed |  | 63.05 | short | 0.94 |
| 2026-05-21T20:15:02Z | T-2302775-1779384603953 | uncategorized | 0.7500 | 0.9990 | 19.94 | market_closed |  | 60.07 | short | 0.11 |
| 2026-05-21T22:30:02Z | T-2068130-1779390003884 | uncategorized | 0.8900 | 0.9990 | 6.94 | market_closed |  | 56.63 | short | 0.15 |
| 2026-05-21T23:05:02Z | T-2068134-1779393603920 | uncategorized | 0.9300 | 0.9990 | 4.22 | market_closed |  | 56.83 | short | 0.13 |
| 2026-05-21T23:10:01Z | T-2068137-1779384603953 | uncategorized | 0.8600 | 0.9990 | 9.51 | market_closed |  | 58.87 | short | 0.24 |
| 2026-05-22T00:05:03Z | T-2069862-1779406203931 | uncategorized | 0.3300 | 0.0300 | -53.51 | stop_loss |  | 58.86 | short | 0.02 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T00:55:01Z | T-553849-1779211529037 | uncategorized | 0.1080 | 0.1590 | 26.01 | take_profit |  | 55.07 | long | 2.31 |
| 2026-05-22T02:15:02Z | T-2069865-1779411603425 | uncategorized | 0.9400 | 0.0100 | -56.51 | stop_loss |  | 57.12 | short | 0.05 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T02:25:02Z | T-2068316-1779402604186 | uncategorized | 0.7800 | 0.9990 | 16.12 | market_closed |  | 57.42 | short | 0.16 |
| 2026-05-22T02:35:02Z | T-2068327-1779409804281 | uncategorized | 0.2800 | 0.0100 | -54.61 | stop_loss |  | 56.63 | short | 0.09 |
| 2026-05-22T02:35:02Z | T-2068323-1779417003389 | uncategorized | 0.1200 | 0.0100 | -52.67 | stop_loss |  | 57.46 | short | 0.00 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T03:30:02Z | T-2237156-1779343204726 | uncategorized | 0.4520 | 0.6360 | 23.94 | take_profit |  | 58.80 | short | 0.90 |
| 2026-05-22T04:20:02Z | T-2293987-1779418803725 | uncategorized | 0.1700 | 0.0100 | -52.07 | stop_loss |  | 55.32 | short | 0.06 |
| 2026-05-22T04:25:02Z | T-2293985-1779420604199 | uncategorized | 0.4800 | 0.0100 | -53.65 | stop_loss |  | 54.79 | short | 0.04 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T05:55:02Z | T-2261157-1779372004234 | uncategorized | 0.8800 | 0.9990 | 7.58 | market_closed |  | 56.02 | short | 0.66 |
| 2026-05-22T07:30:02Z | T-2297854-1779384603953 | uncategorized | 0.1320 | 0.0260 | -46.33 | stop_loss |  | 57.69 | short | 0.58 |
| 2026-05-22T10:25:02Z | T-2241873-1779220804587 | uncategorized | 0.3800 | 0.0700 | -42.55 | stop_loss |  | 52.16 | medium | 2.60 |
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T10:25:03Z | T-2324252-1779429603987 | uncategorized | 0.0900 | 0.1300 | 23.52 | take_profit |  | 52.92 | medium | 0.18 |
| 2026-05-22T12:55:02Z | T-2324252-1779445803654 | uncategorized | 0.1400 | 0.2100 | 25.86 | take_profit |  | 51.71 | medium | 0.10 |
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


## Human notes
_(no se toca por automatización)_
