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
| 2026-05-20T15:30:05Z | vanta-ipo-before-2027 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-there-be-between-8-and-10-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | uncategorized | 0.0980 | P0_floor | price floor: 0.0980 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | fannie-mae-ipo-before-2027 | uncategorized | 0.1200 | P3_low_absolute_liquidity | liquidity $2510 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-kendrick-lamar-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-elon-musk-win-his-case-against-sam-altman | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-the-doge-1-lunar-mission-launch-before-2027 | uncategorized | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | deel-ipo-before-2027 | uncategorized | 0.2100 | P3_low_absolute_liquidity | liquidity $4413 < absolute min $5000 |
| 2026-05-20T15:30:05Z | anysphere-cursor-ipo-before-2027 | uncategorized | 0.0560 | P0_floor | price floor: 0.0560 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-there-be-between-14-and-16-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | uncategorized | 0.3100 | P3_low_absolute_liquidity | liquidity $2015 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-austria-recognize-palestine-before-2027 | uncategorized | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-the-feds-lower-bound-reach-2pt75-or-lower-before-2027-448-727-854 | uncategorized | 0.1000 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-20T15:30:05Z | us-enacts-ai-safety-bill-before-2027 | uncategorized | 0.2900 | P3_low_absolute_liquidity | liquidity $1715 < absolute min $5000 |
| 2026-05-20T15:30:05Z | kadyrov-out-as-head-of-the-chechen-republic-by-june-30-2026 | uncategorized | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | rippling-ipo-before-2027 | uncategorized | 0.1500 | P3_low_absolute_liquidity | liquidity $2279 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-7-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-clavicular-be-named-peoples-sexiest-man-alive-in-2026-399 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-spacex-have-200-or-more-launches-in-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | uncategorized | 0.1600 | P3_low_absolute_liquidity | liquidity $4329 < absolute min $5000 |
| 2026-05-20T15:30:05Z | milei-out-as-president-of-argentina-before-2027 | uncategorized | 0.0610 | P0_floor | price floor: 0.0610 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-israel-strike-8-countries-in-2026 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-toy-story-5-have-the-best-domestic-opening-weekend-in-2026 | uncategorized | 0.0320 | P0_floor | price floor: 0.0320 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-snapchat-be-acquired-before-2027-286-465 | uncategorized | 0.2390 | P3_low_absolute_liquidity | liquidity $2772 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-israel-strike-7-countries-in-2026 | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | openai-1t-ipo-before-2027 | uncategorized | 0.2000 | P3_low_absolute_liquidity | liquidity $2256 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-macky-sall-be-the-next-secretary-general-of-the-united-nations | uncategorized | 0.1170 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-20T15:30:05Z | will-any-country-leave-nato-by-december-31-2026 | uncategorized | 0.0760 | P0_floor | price floor: 0.0760 < 0.100 (horizon 224 d) |
| 2026-05-20T15:30:05Z | will-ole-gunnar-solskjr-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-fomo-launch-a-token-by-december-31-2026 | uncategorized | 0.3200 | P3_low_absolute_liquidity | liquidity $772 < absolute min $5000 |
| 2026-05-20T15:30:05Z | extended-fdv-above-2b-one-day-after-launch-692-597-187 | uncategorized | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | puffpaw-fdv-above-400m-one-day-after-launch | uncategorized | 0.0660 | P0_floor | price floor: 0.0660 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-perena-launch-a-token-by-june-30-2026 | uncategorized | 0.1700 | P3_low_absolute_liquidity | liquidity $1214 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-ethereum-reach-6000-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-lighter-reach-8-before-2027 | uncategorized | 0.0390 | P0_floor | price floor: 0.0390 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | consensys-ipo-closing-market-cap-above-1b | uncategorized | 0.4100 | P3_low_absolute_liquidity | liquidity $1080 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-stablecoins-hit-500b-before-2027 | uncategorized | 0.1300 | P3_low_absolute_liquidity | liquidity $2674 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | uncategorized | 0.0410 | P0_floor | price floor: 0.0410 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | over-1b-raised-on-coinbase-in-2026 | uncategorized | 0.2500 | P3_low_absolute_liquidity | liquidity $1465 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-ethereum-reach-7000-by-december-31-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | usdc-depeg-by-december-31-968 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.3200 | P3_low_absolute_liquidity | liquidity $1365 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-loopscale-launch-a-token-by-december-31-2026 | uncategorized | 0.4200 | P3_low_absolute_liquidity | liquidity $647 < absolute min $5000 |
| 2026-05-20T15:30:05Z | metamask-fdv-above-3b-one-day-after-launch-363-663-664-569-222 | uncategorized | 0.0780 | P0_floor | price floor: 0.0780 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | uncategorized | 0.1800 | P3_low_absolute_liquidity | liquidity $4894 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.6800 | P3_low_absolute_liquidity | liquidity $3099 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-tempo-launch-a-token-by-june-30-2026 | uncategorized | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-monero-hit-1000-in-2026 | uncategorized | 0.1800 | P3_low_absolute_liquidity | liquidity $3315 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-ethereum-reach-6500-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 225 d) |
| 2026-05-20T15:30:05Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | uncategorized | 0.0530 | P0_floor | price floor: 0.0530 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.0810 | P0_floor | price floor: 0.0810 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-pittsburgh-steelers-win-the-2027-nfl-afc-championship-898 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-miami-dolphins-win-the-2027-nfl-afc-championship-357 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 249 d) |
| 2026-05-20T15:30:05Z | will-the-atlanta-falcons-win-the-2027-nfl-league-championship | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 315 d) |
| 2026-05-20T15:30:05Z | openai-ipo-closing-market-cap-above-1pt2t | uncategorized | 0.5800 | P3_low_absolute_liquidity | liquidity $1235 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-spacexs-market-cap-be-between-500b-and-600b-at-market-close-on-ipo-day | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 589 d) |
| 2026-05-20T15:30:05Z | maduro-guilty-of-all-counts | uncategorized | 0.1500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-20T15:30:05Z | openai-ipo-closing-market-cap-above-1t | uncategorized | 0.6300 | P3_low_absolute_liquidity | liquidity $2633 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-perplexity-not-ipo-by-december-31-2027 | uncategorized | 0.5400 | P3_low_absolute_liquidity | liquidity $979 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-ubs-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offering-884 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 589 d) |
| 2026-05-20T15:30:05Z | abstract-fdv-above-3b-one-day-after-launch | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 590 d) |
| 2026-05-20T15:30:05Z | will-joe-kent-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-donald-trump-jr-win-the-2028-us-presidential-election | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-nikki-haley-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-kristi-noem-win-the-2028-republican-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-chris-murphy-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-josh-hawley-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-michelle-obama-win-the-2028-democratic-presidential-nomination-777 | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-james-talarico-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-thomas-massie-win-the-2028-us-presidential-election | uncategorized | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-ron-desantis-win-the-2028-republican-presidential-nomination | uncategorized | 0.0450 | P0_floor | price floor: 0.0450 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-tucker-carlson-win-the-2028-us-presidential-election | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T15:30:05Z | will-the-republicans-win-the-michigan-governor-race-in-2026 | uncategorized | 0.1700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-20T15:30:05Z | will-netflix-close-warner-bros-acquisition | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon -1 d) |
| 2026-05-20T15:30:05Z | will-the-iran-ceasefire-continue-through-may-20 | uncategorized | 0.9840 | P0_ceiling | price ceiling: 0.9840 > 0.950 |
| 2026-05-20T15:30:05Z | will-axiom-launch-a-token-by-december-31-2026 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $370 < absolute min $5000 |
| 2026-05-20T15:30:05Z | will-the-iran-ceasefire-continue-through-may-21 | uncategorized | 0.9600 | P0_ceiling | price ceiling: 0.9600 > 0.950 |
| 2026-05-20T15:30:05Z | will-venezuela-recognize-israel-by-june-30-115 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon -1 d) |
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
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-20T15:50:03Z | T-2241872-1779289205006 | uncategorized | 0.0500 | 0.0700 | 24.62 | take_profit |  | 61.55 | short | 0.03 |
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


## Human notes
_(no se toca por automatización)_
