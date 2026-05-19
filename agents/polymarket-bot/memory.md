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
| 2026-05-19T14:30:03Z | will-jensen-huang-be-richest-person-on-december-31-229 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-steve-ballmer-be-richest-person-on-december-31 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-israel-strike-9-countries-in-2026 | uncategorized | 0.0125 | P0 | precio fuera de [0.05, 0.95]: 0.0125 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | uncategorized | 0.0190 | P0 | precio fuera de [0.05, 0.95]: 0.0190 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | fannie-mae-ipo-before-2027 | uncategorized | 0.1150 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0205 | P0 | precio fuera de [0.05, 0.95]: 0.0205 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | uncategorized | 0.0195 | P0 | precio fuera de [0.05, 0.95]: 0.0195 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | kadyrov-out-as-head-of-the-chechen-republic-by-june-30-2026 | uncategorized | 0.0395 | P0 | precio fuera de [0.05, 0.95]: 0.0395 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | israel-and-lebanon-normalize-relations-before-2027 | uncategorized | 0.1800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | major-meteor-strike-10kt-in-2026 | uncategorized | 0.1450 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-9-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-petr-yan-fight-dominick-cruz-next | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-ubisoft-be-acquired-before-2027-175-488 | uncategorized | 0.2350 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-israel-strike-11-countries-in-2026 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-russia-capture-all-of-lyman-by-december-31-2026 | uncategorized | 0.2500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | poilievre-out-as-leader-of-conservatives-before-2027 | uncategorized | 0.1100 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-syria-join-the-abraham-accords-before-2027 | uncategorized | 0.1050 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026 | uncategorized | 0.0100 | P0 | precio fuera de [0.05, 0.95]: 0.0100 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | insurrection-act-invoked-by-december-31-184-556 | uncategorized | 0.2400 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-oman-join-the-abraham-accords-before-2027 | uncategorized | 0.1150 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | ramp-ipo-before-2027 | uncategorized | 0.1050 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-10-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-beyonc-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-israel-strike-8-countries-in-2026 | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-larry-page-be-richest-person-on-december-31 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | ledger-ipo-before-2027 | uncategorized | 0.1050 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | starmer-out-by-december-31-2026-936-416-977-234-134-475 | uncategorized | 0.7350 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | hyperliquid-listed-on-binance-in-2026 | uncategorized | 0.2850 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-tempo-launch-a-token-by-december-31-2026 | uncategorized | 0.2500 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-exponent-launch-a-token-by-december-31-2026 | uncategorized | 0.4850 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | metamask-fdv-above-700m-one-day-after-launch-696-977-652-246-632 | uncategorized | 0.3310 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-ethereum-reach-7000-by-december-31-2026 | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-tempo-launch-a-token-by-september-30-2026 | uncategorized | 0.0805 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0415 | P0 | precio fuera de [0.05, 0.95]: 0.0415 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-ethereum-reach-6500-by-december-31-2026 | uncategorized | 0.0550 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | major-cex-insolvent-in-2026 | uncategorized | 0.0800 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-consensys-ipo-by-december-31-2026 | uncategorized | 0.2350 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-ostium-launch-a-token-by-december-31-2026 | uncategorized | 0.6550 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | puffpaw-fdv-above-200m-one-day-after-launch | uncategorized | 0.2600 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-fomo-launch-a-token-by-june-30-2026 | uncategorized | 0.0245 | P0 | precio fuera de [0.05, 0.95]: 0.0245 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0175 | P0 | precio fuera de [0.05, 0.95]: 0.0175 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-bitcoin-outperform-gold-in-2026 | uncategorized | 0.3450 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0155 | P0 | precio fuera de [0.05, 0.95]: 0.0155 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0190 | P0 | precio fuera de [0.05, 0.95]: 0.0190 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0255 | P0 | precio fuera de [0.05, 0.95]: 0.0255 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0225 | P0 | precio fuera de [0.05, 0.95]: 0.0225 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-dallas-cowboys-win-the-2027-nfl-nfc-championship-584 | uncategorized | 0.0725 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-minnesota-vikings-win-the-2027-nfl-nfc-championship-884 | uncategorized | 0.0395 | P0 | precio fuera de [0.05, 0.95]: 0.0395 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0235 | P0 | precio fuera de [0.05, 0.95]: 0.0235 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-the-new-orleans-saints-win-the-2027-nfl-league-championship | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-lisabeth-borne-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-franois-bayrou-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-manuel-bompard-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-olivier-faure-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-bernard-cazeneuve-win-the-2027-french-presidential-election | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-carole-delga-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-michel-barnier-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-xavier-bertrand-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-nicols-maduro-be-sentenced-to-no-prison-time-974 | uncategorized | 0.2350 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | openai-ipo-closing-market-cap-above-1pt2t | uncategorized | 0.5700 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-perplexity-not-ipo-by-december-31-2027 | uncategorized | 0.4950 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-erika-kirk-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0125 | P0 | precio fuera de [0.05, 0.95]: 0.0125 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-john-thune-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-tom-brady-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-wes-moore-win-the-2028-us-presidential-election | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0285 | P0 | precio fuera de [0.05, 0.95]: 0.0285 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-tucker-carlson-win-the-2028-us-presidential-election | uncategorized | 0.0325 | P0 | precio fuera de [0.05, 0.95]: 0.0325 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-thomas-massie-win-the-2028-us-presidential-election | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-nikki-haley-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-elon-musk-win-the-2028-republican-presidential-nomination | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-joe-kent-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0145 | P0 | precio fuera de [0.05, 0.95]: 0.0145 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-katie-britt-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-jamie-dimon-win-the-2028-us-presidential-election | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-tim-walz-win-the-2028-us-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-kamala-harris-win-the-2028-us-presidential-election | uncategorized | 0.0770 | M1 | memoria: exact slug match (score 1.00) |
| 2026-05-19T14:30:03Z | will-steve-bannon-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-comcast-close-warner-bros-acquisition | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | reya-fdv-above-1b-one-day-after-launch-348-347 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T14:30:03Z | will-netflix-close-warner-bros-acquisition | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
| 2026-05-19T11:30:00Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0065 | 0.0065 | 0.00 | manual_close_v1_audit · ejemplo long-tail absurdo aprobado por v1 |
| 2026-05-19T11:30:00Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | uncategorized | 0.0075 | 0.0075 | 0.00 | manual_close_v1_audit · ejemplo long-tail absurdo aprobado por v1 |

## Reglas blandas aprendidas (autogenerada por brain cuando un pattern se repite 3+ veces)

_vacío — primera ejecución v2_

**Reglas iniciales seed** (basadas en post-mortem v1, ver [[../../03-decisions/2026-05-19-polymarket-v1-baseline-failure]]):

- **Evitar long-tail extremo**: `current_price_yes < 0.05` en categorías políticas a largo plazo (>12 meses). Razón: take-profit +60% sobre 0.005 = 0.008 → ganancia $0.45 sobre $150 size, no merece el riesgo.
- **Categoría `uncategorized` con resolución >12 meses**: tratar con escepticismo elevado. La API no expone categoría → no podemos verificar correlación. Veto si volumen 24h real (cuando disponible) <$1k.
- **Volumen `volumeNum` cumulativo NO es proxy fiable**. Para v2 preferir spread (`bestAsk - bestBid`) si está disponible. Si `bestBid` o `bestAsk` faltan/cero, considerar mercado ilíquido y vetar.

## Human notes
_(no se toca por automatización)_
