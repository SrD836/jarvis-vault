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
| 2026-05-23T05:30:03Z | will-zoom-video-communications-be-acquired-before-2027-486-866 | uncategorized | 0.2110 | P3_low_absolute_liquidity | liquidity $2181 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-fearx-win-the-lck-2026-season-playoffs | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | ripple-labs-ipo-before-2027 | uncategorized | 0.1500 | P3_low_absolute_liquidity | liquidity $2611 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-gitlab-be-acquired-before-2027-944-667 | uncategorized | 0.2100 | P3_low_absolute_liquidity | liquidity $4660 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-paramount-close-warner-bros-acquisition-by-end-of-2026 | uncategorized | 0.7500 | P3_low_absolute_liquidity | liquidity $2999 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-trump-sell-10k-25k-gold-cards-in-2026 | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-ali-asghar-hejazi-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-bernard-arnault-be-richest-person-on-december-31-747 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | sbf-released-from-custody-in-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-beyonc-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | uncategorized | 0.0590 | P0_floor | price floor: 0.0590 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-zelenskyy-and-putin-meet-next-in-russia | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | rippling-ipo-before-2027 | uncategorized | 0.1700 | P3_low_absolute_liquidity | liquidity $3034 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-nongshim-redforce-win-the-lck-2026-season-playoffs | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-china-invade-taiwan-before-2027 | uncategorized | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-jeff-bezos-be-richest-person-on-december-31-243 | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-ubisoft-be-acquired-before-2027-175-488 | uncategorized | 0.2000 | P3_low_absolute_liquidity | liquidity $2931 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-zelenskyy-and-putin-meet-next-in-belarus | uncategorized | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-eoghan-ceannabhin-win-the-2026-dublin-central-by-election | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | uncategorized | 0.0020 | P0_floor | price floor: 0.0020 < 0.100 (horizon 221 d) |
| 2026-05-23T05:30:03Z | will-ethereum-reach-8000-by-december-31-2026 | uncategorized | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon 222 d) |
| 2026-05-23T05:30:03Z | usdc-depeg-by-december-31-968 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 222 d) |
| 2026-05-23T05:30:03Z | consensys-ipo-closing-market-cap-above-1b | uncategorized | 0.3900 | P3_low_absolute_liquidity | liquidity $1258 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-exponent-launch-a-token-by-december-31-2026 | uncategorized | 0.2700 | P3_low_absolute_liquidity | liquidity $990 < absolute min $5000 |
| 2026-05-23T05:30:03Z | standx-fdv-above-2b-one-day-after-launch-143-135-174-596-974 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 222 d) |
| 2026-05-23T05:30:03Z | puffpaw-fdv-above-200m-one-day-after-launch | uncategorized | 0.2700 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-23T05:30:03Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | uncategorized | 0.3300 | P4_pre_event | pre-event slug + 222 d to resolution (>=7 threshold) |
| 2026-05-23T05:30:03Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 222 d) |
| 2026-05-23T05:30:03Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $1523 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-consensys-ipo-by-december-31-2026 | uncategorized | 0.2900 | P3_low_absolute_liquidity | liquidity $1928 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 222 d) |
| 2026-05-23T05:30:03Z | major-cex-insolvent-in-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 222 d) |
| 2026-05-23T05:30:03Z | will-usdc-hit-50-of-usdt-market-cap-by-december-31-2026 | uncategorized | 0.4500 | P3_low_absolute_liquidity | liquidity $600 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-fomo-launch-a-token-by-december-31-2026 | uncategorized | 0.3800 | P3_low_absolute_liquidity | liquidity $786 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-indianapolis-colts-win-the-2027-nfl-afc-championship-168 | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-23T05:30:03Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-atlanta-falcons-win-the-2027-nfl-nfc-championship-312 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-tampa-bay-buccaneers-win-the-2027-nfl-nfc-championship-312 | uncategorized | 0.0380 | P0_floor | price floor: 0.0380 < 0.100 (horizon 246 d) |
| 2026-05-23T05:30:03Z | will-saquon-barkley-win-the-2026-nfl-mvp | uncategorized | 0.0550 | P0_floor | price floor: 0.0550 < 0.100 (horizon 267 d) |
| 2026-05-23T05:30:03Z | will-baker-mayfield-win-the-2026-nfl-mvp | uncategorized | 0.0590 | P0_floor | price floor: 0.0590 < 0.100 (horizon 267 d) |
| 2026-05-23T05:30:03Z | will-maurcio-ruffy-fight-charles-oliveira-next-853-769 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 281 d) |
| 2026-05-23T05:30:03Z | will-mathilde-panot-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 341 d) |
| 2026-05-23T05:30:03Z | will-olivier-faure-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 341 d) |
| 2026-05-23T05:30:03Z | will-sgolne-royal-win-the-2027-french-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 341 d) |
| 2026-05-23T05:30:03Z | will-deutsche-bank-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offe... | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 586 d) |
| 2026-05-23T05:30:03Z | openai-ipo-closing-market-cap-above-800b | uncategorized | 0.8800 | P3_low_absolute_liquidity | liquidity $3075 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-ubs-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offering-884 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 586 d) |
| 2026-05-23T05:30:03Z | will-perplexity-not-ipo-by-december-31-2027 | uncategorized | 0.4700 | P3_low_absolute_liquidity | liquidity $2073 < absolute min $5000 |
| 2026-05-23T05:30:03Z | openai-ipo-closing-market-cap-above-1t | uncategorized | 0.7800 | P3_low_absolute_liquidity | liquidity $2650 < absolute min $5000 |
| 2026-05-23T05:30:03Z | abstract-fdv-above-3b-one-day-after-launch | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 587 d) |
| 2026-05-23T05:30:03Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-tulsi-gabbard-win-the-2028-republican-presidential-nomination | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-thomas-massie-win-the-2028-us-presidential-election | uncategorized | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-pete-hegseth-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-ruben-gallego-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-gina-raimondo-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-kristi-noem-win-the-2028-republican-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-mark-kelly-win-the-2028-democratic-presidential-nomination-479 | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-donald-trump-win-the-2028-us-presidential-election | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-glenn-youngkin-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-michelle-obama-win-the-2028-democratic-presidential-nomination-777 | uncategorized | 0.0140 | P0_floor | price floor: 0.0140 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-mrbeast-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-donald-trump-win-the-2028-republican-presidential-nomination | uncategorized | 0.0300 | P0_floor | price floor: 0.0300 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-tim-walz-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-tim-walz-win-the-2028-democratic-presidential-nomination-475 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-donald-trump-jr-win-the-2028-us-presidential-election | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 898 d) |
| 2026-05-23T05:30:03Z | will-variational-launch-a-token-by-december-31-2026-577-621 | uncategorized | 0.8200 | P3_low_absolute_liquidity | liquidity $1894 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-rebecca-shepherd-win-the-2026-makerfield-by-election | uncategorized | 0.0780 | P0_floor | price floor: 0.0780 < 0.100 (horizon -1 d) |
| 2026-05-23T05:30:03Z | will-the-iran-ceasefire-continue-through-may-24-733 | uncategorized | 0.8200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-23T05:30:03Z | will-netflix-close-warner-bros-acquisition | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon -1 d) |
| 2026-05-23T05:30:03Z | will-axiom-launch-a-token-by-december-31-2026 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $478 < absolute min $5000 |
| 2026-05-23T05:30:03Z | will-the-iran-ceasefire-continue-through-may-22 | uncategorized | 0.9980 | P0_ceiling | price ceiling: 0.9980 > 0.950 |
| 2026-05-23T05:30:03Z | will-the-iran-ceasefire-continue-through-may-27-496 | uncategorized | 0.7200 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-23T05:30:03Z | will-the-iran-ceasefire-continue-through-may-31-654-633 | uncategorized | 0.6500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-23T05:30:03Z | will-comcast-close-warner-bros-acquisition | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon -1 d) |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
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
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T10:25:03Z | T-2324252-1779429603987 | uncategorized | 0.0900 | 0.1300 | 23.52 | take_profit |  | 52.92 | medium | 0.18 |
| 2026-05-22T12:55:02Z | T-2324252-1779445803654 | uncategorized | 0.1400 | 0.2100 | 25.86 | take_profit |  | 51.71 | medium | 0.10 |
| 2026-05-22T13:35:03Z | T-2237156-1779424203950 | uncategorized | 0.7280 | 0.1420 | -43.29 | stop_loss |  | 53.78 | short | 0.38 |
| 2026-05-22T13:50:02Z | T-2269345-1779395404150 | uncategorized | 0.3900 | 0.0600 | -48.48 | stop_loss |  | 57.29 | short | 0.72 |
| 2026-05-22T14:45:02Z | T-2269345-1779458404300 | uncategorized | 0.0700 | 0.0050 | -46.93 | stop_loss |  | 50.54 | short | 0.03 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T15:50:02Z | T-2297854-1779458404300 | uncategorized | 0.0500 | 0.0730 | 23.72 | take_profit |  | 51.57 | short | 0.08 |
| 2026-05-22T17:20:02Z | T-1691780-1779211803817 | uncategorized | 0.2000 | 0.9800 | 198.11 | take_profit |  | 50.80 | long | 2.99 |
| 2026-05-22T17:30:02Z | T-2169751-1779435004151 | uncategorized | 0.4590 | 0.0160 | -50.27 | stop_loss |  | 52.09 | short | 0.42 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T18:40:02Z | T-2241873-1779445803654 | uncategorized | 0.0800 | 0.1300 | 32.98 | take_profit |  | 52.77 | short | 0.34 |
| 2026-05-22T19:05:02Z | T-2269343-1779406203931 | uncategorized | 0.9500 | 0.9990 | 2.98 | market_closed |  | 57.68 | short | 0.82 |
| 2026-05-22T19:05:02Z | T-2269308-1779417003389 | uncategorized | 0.8900 | 0.9990 | 6.90 | market_closed |  | 56.31 | short | 0.69 |
| 2026-05-22T20:00:02Z | T-2078702-1779471003629 | uncategorized | 0.6500 | 0.9300 | 23.32 | take_profit |  | 54.13 | short | 0.10 |
| 2026-05-22T20:05:02Z | T-2287742-1779424203950 | uncategorized | 0.3600 | 0.5300 | 24.89 | take_profit |  | 52.70 | short | 0.65 |
| 2026-05-22T20:30:03Z | T-2287744-1779418803725 | uncategorized | 0.3700 | 0.0600 | -45.43 | stop_loss |  | 54.22 | short | 0.73 |
| 2026-05-22T20:30:03Z | T-2119227-1779454803784 | uncategorized | 0.0940 | 0.0130 | -45.01 | stop_loss |  | 52.23 | short | 0.31 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T20:40:02Z | T-2287743-1779478204560 | uncategorized | 0.3100 | 0.7100 | 69.68 | take_profit |  | 54.00 | short | 0.05 |
| 2026-05-22T20:40:02Z | T-2287744-1779481805168 | uncategorized | 0.0700 | 0.1700 | 78.93 | take_profit |  | 55.25 | short | 0.01 |
| 2026-05-22T20:40:02Z | T-2287742-1779481805168 | uncategorized | 0.6900 | 0.1100 | -45.52 | stop_loss |  | 54.15 | short | 0.01 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T20:40:02Z | T-2290413-1779481805168 | uncategorized | 0.4000 | 0.9900 | 78.27 | take_profit |  | 53.07 | short | 0.01 |
| 2026-05-22T21:05:01Z | T-2241873-1779476404632 | uncategorized | 0.1400 | 0.2000 | 23.01 | take_profit |  | 53.70 | short | 0.09 |
| 2026-05-22T21:10:02Z | T-2297854-1779465603905 | uncategorized | 0.0720 | 0.0120 | -41.76 | stop_loss |  | 50.12 | short | 0.22 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-22T21:30:02Z | T-2296135-1779310803840 | uncategorized | 0.2400 | 0.4100 | 42.89 | take_profit |  | 60.55 | medium | 2.02 |
| 2026-05-22T21:35:02Z | T-2241873-1779485404133 | uncategorized | 0.2100 | 0.3100 | 28.13 | take_profit |  | 59.08 | short | 0.00 |
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


## Human notes
_(no se toca por automatización)_
