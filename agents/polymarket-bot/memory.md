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
| 2026-05-19T21:30:03Z | ramp-ipo-before-2027 | uncategorized | 0.1100 | P3_low_absolute_liquidity | liquidity $3047 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-spacex-have-200-or-more-launches-in-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-beyonc-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-israel-strike-9-countries-in-2026 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-china-invade-taiwan-before-2027 | uncategorized | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | fed-emergency-rate-cut-before-2027 | uncategorized | 0.1000 | P3_low_absolute_liquidity | liquidity $4840 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-larry-page-be-richest-person-on-december-31 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-the-us-invade-mexico-in-2026 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-steve-ballmer-be-richest-person-on-december-31 | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | uae-x-qatar-sever-diplomatic-relations-in-2026 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | freddie-mac-ipo-before-2027 | uncategorized | 0.1400 | P3_low_absolute_liquidity | liquidity $1699 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-macky-sall-be-the-next-secretary-general-of-the-united-nations | uncategorized | 0.1160 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T21:30:03Z | spacex-starship-fully-reusable-before-2027 | uncategorized | 0.3400 | P3_low_absolute_liquidity | liquidity $2204 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-snapchat-be-acquired-before-2027-286-465 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $3577 < absolute min $5000 |
| 2026-05-19T21:30:03Z | openai-1t-ipo-before-2027 | uncategorized | 0.2000 | P3_low_absolute_liquidity | liquidity $2508 < absolute min $5000 |
| 2026-05-19T21:30:03Z | vanta-ipo-before-2027 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-israel-strike-11-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-apple-release-a-new-product-line-before-2027 | uncategorized | 0.3300 | P3_low_absolute_liquidity | liquidity $3237 < absolute min $5000 |
| 2026-05-19T21:30:03Z | rippling-ipo-before-2027 | uncategorized | 0.1400 | P3_low_absolute_liquidity | liquidity $3364 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-10-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | kadyrov-out-as-head-of-the-chechen-republic-by-june-30-2026 | uncategorized | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | 5kt-meteor-strike-in-2026 | uncategorized | 0.3400 | P3_low_absolute_liquidity | liquidity $3531 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-any-country-leave-nato-by-december-31-2026 | uncategorized | 0.0760 | P0_floor | price floor: 0.0760 < 0.100 (horizon 225 d) |
| 2026-05-19T21:30:03Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-ethereum-reach-6000-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-ethereum-reach-6500-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | uncategorized | 0.3540 | P3_low_absolute_liquidity | liquidity $4831 < absolute min $5000 |
| 2026-05-19T21:30:03Z | major-cex-insolvent-in-2026 | uncategorized | 0.0900 | P0_floor | price floor: 0.0900 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-lighter-reach-8-before-2027 | uncategorized | 0.0390 | P0_floor | price floor: 0.0390 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | over-1b-raised-on-coinbase-in-2026 | uncategorized | 0.2300 | P3_low_absolute_liquidity | liquidity $1818 < absolute min $5000 |
| 2026-05-19T21:30:03Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0540 | P0_floor | price floor: 0.0540 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.3400 | P3_low_absolute_liquidity | liquidity $1191 < absolute min $5000 |
| 2026-05-19T21:30:03Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-consensys-ipo-by-december-31-2026 | uncategorized | 0.3100 | P3_low_absolute_liquidity | liquidity $1785 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-stablecoins-hit-500b-before-2027 | uncategorized | 0.1300 | P3_low_absolute_liquidity | liquidity $3480 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-lighter-reach-10-before-2027 | uncategorized | 0.0430 | P0_floor | price floor: 0.0430 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-ethereum-reach-7000-by-december-31-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 226 d) |
| 2026-05-19T21:30:03Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.6800 | P3_low_absolute_liquidity | liquidity $3092 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-fomo-launch-a-token-by-december-31-2026 | uncategorized | 0.3800 | P3_low_absolute_liquidity | liquidity $887 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.0820 | P0_floor | price floor: 0.0820 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-tampa-bay-buccaneers-win-the-2027-nfl-nfc-championship-312 | uncategorized | 0.0370 | P0_floor | price floor: 0.0370 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-miami-dolphins-win-the-2027-nfl-afc-championship-357 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-minnesota-vikings-win-the-2027-nfl-nfc-championship-884 | uncategorized | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | uncategorized | 0.0530 | P0_floor | price floor: 0.0530 < 0.100 (horizon 250 d) |
| 2026-05-19T21:30:03Z | will-bernard-cazeneuve-win-the-2027-french-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 345 d) |
| 2026-05-19T21:30:03Z | openai-ipo-closing-market-cap-above-1pt2t | uncategorized | 0.5800 | P3_low_absolute_liquidity | liquidity $1549 < absolute min $5000 |
| 2026-05-19T21:30:03Z | will-perplexity-not-ipo-by-december-31-2027 | uncategorized | 0.5800 | P3_low_absolute_liquidity | liquidity $968 < absolute min $5000 |
| 2026-05-19T21:30:03Z | base-fdv-above-12b-one-day-after-launch | uncategorized | 0.3200 | P4_pre_event | pre-event slug + 591 d to resolution (>=7 threshold) |
| 2026-05-19T21:30:03Z | will-katie-britt-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-tom-brady-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-ron-desantis-win-the-2028-republican-presidential-nomination | uncategorized | 0.0370 | P0_floor | price floor: 0.0370 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-josh-hawley-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-erika-kirk-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-pete-hegseth-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-thomas-massie-win-the-2028-us-presidential-election | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-eric-trump-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-wes-moore-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-steve-bannon-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:03Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-joe-kent-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0220 | P0_floor | price floor: 0.0220 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-mike-pence-win-the-2028-republican-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-kamala-harris-win-the-2028-us-presidential-election | uncategorized | 0.0650 | P0_floor | price floor: 0.0650 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-tucker-carlson-win-the-2028-us-presidential-election | uncategorized | 0.0340 | P0_floor | price floor: 0.0340 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-john-thune-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-josh-shapiro-win-the-2028-us-presidential-election | uncategorized | 0.0310 | P0_floor | price floor: 0.0310 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0270 | P0_floor | price floor: 0.0270 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-pete-hegseth-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 902 d) |
| 2026-05-19T21:30:04Z | will-asia-win-the-2026-fifa-world-cup | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon -1 d) |
| 2026-05-19T21:30:04Z | reya-fdv-above-1b-one-day-after-launch-348-347 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon -1 d) |
| 2026-05-19T21:30:04Z | will-comcast-close-warner-bros-acquisition | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -1 d) |
| 2026-05-19T21:30:04Z | will-netflix-close-warner-bros-acquisition | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon -1 d) |
| 2026-05-19T21:30:04Z | will-oceania-win-the-2026-fifa-world-cup | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -1 d) |
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
## Reglas blandas aprendidas (autogenerada por brain cuando un pattern se repite 3+ veces)

_vacío — primera ejecución v2_

**Reglas iniciales seed** (basadas en post-mortem v1, ver [[../../03-decisions/2026-05-19-polymarket-v1-baseline-failure]]):

- **Evitar long-tail extremo**: `current_price_yes < 0.05` en categorías políticas a largo plazo (>12 meses). Razón: take-profit +60% sobre 0.005 = 0.008 → ganancia $0.45 sobre $150 size, no merece el riesgo.
- **Categoría `uncategorized` con resolución >12 meses**: tratar con escepticismo elevado. La API no expone categoría → no podemos verificar correlación. Veto si volumen 24h real (cuando disponible) <$1k.
- **Volumen `volumeNum` cumulativo NO es proxy fiable**. Para v2 preferir spread (`bestAsk - bestBid`) si está disponible. Si `bestBid` o `bestAsk` faltan/cero, considerar mercado ilíquido y vetar.
- **Long-tail floor SUBIDO de 0.05 -> 0.10** (2026-05-19, post-mortem trade T-608399 Discord): entry 0.066, exit 0.007 = -89.4% en 10 min. Cerca del floor, stop_loss_pct 80% se dispara con el spread bid/ask normal sin noticias. Veto: current_price_yes < 0.10 si horizon > 7 dias.
- **Liquidez absoluta minima**: vetar mercados con liquidity_usd < $5000 en bruto, ademas del gate Size x 4. Mercados con $2-3k de profundidad son trampas: el bestBid colapsa al primer market-sell. Discord tenia $2916.
- **Pre-IPO / pre-launch markets**: si slug contiene ipo-day | launch | tge | mainnet | listing y faltan >=7 dias hasta endDate -> veto. No hay forma honesta de modelar el dia-cero de un asset que aun no existe.


## Human notes
_(no se toca por automatización)_
