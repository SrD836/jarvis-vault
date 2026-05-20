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
| 2026-05-20T08:55:57Z | rippling-ipo-before-2027 | uncategorized | 0.1500 | P3_low_absolute_liquidity | liquidity $3165 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-sabrina-carpenter-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0060 | P0_floor | price floor: 0.0060 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-north-korea-invade-south-korea-before-2027 | uncategorized | 0.0650 | P0_floor | price floor: 0.0650 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | ramp-ipo-before-2027 | uncategorized | 0.1100 | P3_low_absolute_liquidity | liquidity $3643 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-china-invade-taiwan-before-2027 | uncategorized | 0.0750 | P0_floor | price floor: 0.0750 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-bp-be-acquired-before-2027-549-844 | uncategorized | 0.1900 | P3_low_absolute_liquidity | liquidity $4641 < absolute min $5000 |
| 2026-05-20T08:55:57Z | us-enacts-ai-safety-bill-before-2027 | uncategorized | 0.2800 | P3_low_absolute_liquidity | liquidity $3005 < absolute min $5000 |
| 2026-05-20T08:55:57Z | nara-smith-confirmed-pregnant-in-2026 | uncategorized | 0.7100 | P3_low_absolute_liquidity | liquidity $214 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-snapchat-be-acquired-before-2027-286-465 | uncategorized | 0.2400 | P3_low_absolute_liquidity | liquidity $3520 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-10-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-macky-sall-be-the-next-secretary-general-of-the-united-nations | uncategorized | 0.1140 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-20T08:55:57Z | celonis-ipo-before-2027 | uncategorized | 0.0800 | P0_floor | price floor: 0.0800 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-clavicular-be-named-peoples-sexiest-man-alive-in-2026-399 | uncategorized | 0.0190 | P0_floor | price floor: 0.0190 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | milei-out-as-president-of-argentina-before-2027 | uncategorized | 0.0620 | P0_floor | price floor: 0.0620 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | brian-armstrong-out-as-coinbase-ceo-before-2027 | uncategorized | 0.1200 | P3_low_absolute_liquidity | liquidity $2934 < absolute min $5000 |
| 2026-05-20T08:55:57Z | uae-x-qatar-sever-diplomatic-relations-in-2026 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-spacex-have-200-or-more-launches-in-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | us-agrees-to-give-ukraine-security-guarantee-by-june-30 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-israel-strike-12-countries-in-2026 | uncategorized | 0.0050 | P0_floor | price floor: 0.0050 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-israel-strike-9-countries-in-2026 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 224 d) |
| 2026-05-20T08:55:57Z | will-any-country-leave-nato-by-december-31-2026 | uncategorized | 0.0760 | P0_floor | price floor: 0.0760 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-ole-gunnar-solskjr-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | reya-fdv-above-150m-one-day-after-launch-514-746 | uncategorized | 0.3300 | P4_pre_event | pre-event slug + 225 d to resolution (>=7 threshold) |
| 2026-05-20T08:55:57Z | will-stablecoins-hit-500b-before-2027 | uncategorized | 0.1300 | P3_low_absolute_liquidity | liquidity $3246 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-tempo-launch-a-token-by-june-30-2026 | uncategorized | 0.0210 | P0_floor | price floor: 0.0210 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | uncategorized | 0.0410 | P0_floor | price floor: 0.0410 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.6800 | P3_low_absolute_liquidity | liquidity $3146 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-trump-launch-a-coin-by-january-31-637-684 | uncategorized | 0.3000 | P3_low_absolute_liquidity | liquidity $1268 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-nansen-launch-a-token-by-september-30-2026 | uncategorized | 0.1700 | P3_low_absolute_liquidity | liquidity $2087 < absolute min $5000 |
| 2026-05-20T08:55:57Z | hyperbeat-fdv-above-300m-one-day-after-launch | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-solana-reach-320-by-december-31-2026 | uncategorized | 0.0700 | P0_floor | price floor: 0.0700 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | over-1b-raised-on-coinbase-in-2026 | uncategorized | 0.2500 | P3_low_absolute_liquidity | liquidity $1838 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-ethereum-reach-6000-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-perena-launch-a-token-by-june-30-2026 | uncategorized | 0.1700 | P3_low_absolute_liquidity | liquidity $1234 < absolute min $5000 |
| 2026-05-20T08:55:57Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | uncategorized | 0.3530 | P3_low_absolute_liquidity | liquidity $4860 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0940 | P0_floor | price floor: 0.0940 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-fomo-launch-a-token-by-december-31-2026 | uncategorized | 0.3300 | P3_low_absolute_liquidity | liquidity $787 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-ethereum-reach-6500-by-december-31-2026 | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | puffpaw-fdv-above-400m-one-day-after-launch | uncategorized | 0.0660 | P0_floor | price floor: 0.0660 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | uncategorized | 0.1100 | P4_pre_event | pre-event slug + 225 d to resolution (>=7 threshold) |
| 2026-05-20T08:55:57Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.3100 | P3_low_absolute_liquidity | liquidity $1361 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-lighter-reach-8-before-2027 | uncategorized | 0.0390 | P0_floor | price floor: 0.0390 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-consensys-ipo-by-december-31-2026 | uncategorized | 0.3100 | P3_low_absolute_liquidity | liquidity $1826 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-ethereum-reach-7000-by-december-31-2026 | uncategorized | 0.0500 | P0_floor | price floor: 0.0500 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0180 | P0_floor | price floor: 0.0180 < 0.100 (horizon 225 d) |
| 2026-05-20T08:55:57Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0230 | P0_floor | price floor: 0.0230 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.0810 | P0_floor | price floor: 0.0810 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0160 | P0_floor | price floor: 0.0160 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0240 | P0_floor | price floor: 0.0240 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-cincinnati-bengals-win-the-2027-nfl-afc-championship-267 | uncategorized | 0.0530 | P0_floor | price floor: 0.0530 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0110 | P0_floor | price floor: 0.0110 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0260 | P0_floor | price floor: 0.0260 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-miami-dolphins-win-the-2027-nfl-afc-championship-357 | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 249 d) |
| 2026-05-20T08:55:57Z | will-ubs-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offering-884 | uncategorized | 0.0010 | P0_floor | price floor: 0.0010 < 0.100 (horizon 589 d) |
| 2026-05-20T08:55:57Z | openai-ipo-closing-market-cap-above-1pt2t | uncategorized | 0.5800 | P3_low_absolute_liquidity | liquidity $1523 < absolute min $5000 |
| 2026-05-20T08:55:57Z | will-perplexity-not-ipo-by-december-31-2027 | uncategorized | 0.5600 | P3_low_absolute_liquidity | liquidity $993 < absolute min $5000 |
| 2026-05-20T08:55:57Z | abstract-fdv-above-3b-one-day-after-launch | uncategorized | 0.0280 | P0_floor | price floor: 0.0280 < 0.100 (horizon 590 d) |
| 2026-05-20T08:55:57Z | will-josh-hawley-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-kamala-harris-win-the-2028-us-presidential-election | uncategorized | 0.0630 | P0_floor | price floor: 0.0630 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-ron-desantis-win-the-2028-republican-presidential-nomination | uncategorized | 0.0470 | P0_floor | price floor: 0.0470 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-john-thune-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-kristi-noem-win-the-2028-republican-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-tucker-carlson-win-the-2028-us-presidential-election | uncategorized | 0.0340 | P0_floor | price floor: 0.0340 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0170 | P0_floor | price floor: 0.0170 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-donald-trump-jr-win-the-2028-us-presidential-election | uncategorized | 0.0130 | P0_floor | price floor: 0.0130 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-thomas-massie-win-the-2028-us-presidential-election | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-wes-moore-win-the-2028-us-presidential-election | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-joe-kent-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-michelle-obama-win-the-2028-democratic-presidential-nomination-777 | uncategorized | 0.0150 | P0_floor | price floor: 0.0150 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-james-talarico-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0200 | P0_floor | price floor: 0.0200 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-tim-walz-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-chris-murphy-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-donald-trump-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0250 | P0_floor | price floor: 0.0250 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0070 | P0_floor | price floor: 0.0070 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0120 | P0_floor | price floor: 0.0120 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | uncategorized | 0.0100 | P0_floor | price floor: 0.0100 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0090 | P0_floor | price floor: 0.0090 < 0.100 (horizon 901 d) |
| 2026-05-20T08:55:57Z | will-valve-remove-inferno-from-the-map-pool | uncategorized | 0.0600 | P0_floor | price floor: 0.0600 < 0.100 (horizon -1 d) |
| 2026-05-20T08:55:57Z | will-valve-remove-dust-2-from-the-map-pool | uncategorized | 0.0400 | P0_floor | price floor: 0.0400 < 0.100 (horizon -1 d) |
| 2026-05-20T08:55:57Z | will-oceania-win-the-2026-fifa-world-cup | uncategorized | 0.0030 | P0_floor | price floor: 0.0030 < 0.100 (horizon -1 d) |
| 2026-05-20T08:55:57Z | will-asia-win-the-2026-fifa-world-cup | uncategorized | 0.0290 | P0_floor | price floor: 0.0290 < 0.100 (horizon -1 d) |
| 2026-05-20T08:55:57Z | will-venezuela-recognize-israel-by-june-30-115 | uncategorized | 0.0080 | P0_floor | price floor: 0.0080 < 0.100 (horizon -1 d) |
| 2026-05-20T08:55:57Z | will-netflix-close-warner-bros-acquisition | uncategorized | 0.0040 | P0_floor | price floor: 0.0040 < 0.100 (horizon -1 d) |
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


## Human notes
_(no se toca por automatización)_
