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
| 2026-05-19T16:30:04Z | will-colorado-rockies-win-the-2026-national-league-championship-series | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-cincinnati-reds-win-the-2026-national-league-championship-series | uncategorized | 0.0255 | P0 | precio fuera de [0.05, 0.95]: 0.0255 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-richard-van-de-water-win-the-bachelorette-season-22 | uncategorized | 0.0125 | P0 | precio fuera de [0.05, 0.95]: 0.0125 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-lew-evans-win-the-bachelorette-season-22 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-liam-lawson-be-the-2026-f1-drivers-champion | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-pierre-gasly-be-the-2026-f1-drivers-champion | uncategorized | 0.0015 | P0 | precio fuera de [0.05, 0.95]: 0.0015 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-alexander-albon-be-the-2026-f1-drivers-champion | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-racing-bulls-be-the-2026-f1-constructors-champion | uncategorized | 0.0055 | P0 | precio fuera de [0.05, 0.95]: 0.0055 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-fernando-alonso-be-the-2026-f1-drivers-champion | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | uncategorized | 0.0045 | P0 | precio fuera de [0.05, 0.95]: 0.0045 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | uncategorized | 0.0045 | P0 | precio fuera de [0.05, 0.95]: 0.0045 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-4pt5-at-the-end-of-2026-139 | uncategorized | 0.0135 | P0 | precio fuera de [0.05, 0.95]: 0.0135 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-fc-cincinnati-win-the-2026-mls-cup | uncategorized | 0.0445 | P0 | precio fuera de [0.05, 0.95]: 0.0445 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-steve-ballmer-be-richest-person-on-december-31 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-petr-yan-fight-dominick-cruz-next | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | anysphere-cursor-ipo-before-2027 | uncategorized | 0.0485 | P0 | precio fuera de [0.05, 0.95]: 0.0485 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-israel-strike-9-countries-in-2026 | uncategorized | 0.0125 | P0 | precio fuera de [0.05, 0.95]: 0.0125 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-clavicular-be-named-peoples-sexiest-man-alive-in-2026-399 | uncategorized | 0.0170 | P0 | precio fuera de [0.05, 0.95]: 0.0170 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-spacex-have-200-or-more-launches-in-2026 | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | us-x-iran-permanent-peace-deal-by-december-31-2026-961-587-341-574 | uncategorized | 0.6850 | M1 | memoria: slug prefix match; same price bucket mid (score 0.70) |
| 2026-05-19T16:30:04Z | will-zelenskyy-and-putin-meet-next-in-saudi-arabia | uncategorized | 0.0145 | P0 | precio fuera de [0.05, 0.95]: 0.0145 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-zelenskyy-and-putin-meet-next-in-ukraine | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-hassan-shariatmadari-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | uncategorized | 0.0195 | P0 | precio fuera de [0.05, 0.95]: 0.0195 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-israel-strike-10-countries-in-2026 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-elon-musk-win-his-case-against-sam-altman | uncategorized | 0.0220 | P0 | precio fuera de [0.05, 0.95]: 0.0220 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-7-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-austria-recognize-palestine-before-2027 | uncategorized | 0.0110 | P0 | precio fuera de [0.05, 0.95]: 0.0110 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-jensen-huang-be-richest-person-on-december-31-229 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | uncategorized | 0.0180 | P0 | precio fuera de [0.05, 0.95]: 0.0180 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0205 | P0 | precio fuera de [0.05, 0.95]: 0.0205 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | kadyrov-out-as-head-of-the-chechen-republic-by-june-30-2026 | uncategorized | 0.0395 | P0 | precio fuera de [0.05, 0.95]: 0.0395 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-israel-strike-8-countries-in-2026 | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-bernard-arnault-be-richest-person-on-december-31-747 | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | uncategorized | 0.0055 | P0 | precio fuera de [0.05, 0.95]: 0.0055 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-israel-strike-11-countries-in-2026 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-beyonc-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-10-fed-rate-cuts-happen-in-2026 | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-jeff-bezos-be-richest-person-on-december-31-243 | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-larry-page-be-richest-person-on-december-31 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-zelenskyy-and-putin-meet-next-in-russia | uncategorized | 0.0185 | P0 | precio fuera de [0.05, 0.95]: 0.0185 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-ethereum-reach-7000-by-december-31-2026 | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0415 | P0 | precio fuera de [0.05, 0.95]: 0.0415 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0175 | P0 | precio fuera de [0.05, 0.95]: 0.0175 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0235 | P0 | precio fuera de [0.05, 0.95]: 0.0235 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | uncategorized | 0.0155 | P0 | precio fuera de [0.05, 0.95]: 0.0155 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-las-vegas-raiders-win-the-2027-nfl-afc-championship-699 | uncategorized | 0.0255 | P0 | precio fuera de [0.05, 0.95]: 0.0255 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-minnesota-vikings-win-the-2027-nfl-nfc-championship-884 | uncategorized | 0.0395 | P0 | precio fuera de [0.05, 0.95]: 0.0395 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0225 | P0 | precio fuera de [0.05, 0.95]: 0.0225 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-new-orleans-saints-win-the-2027-nfl-nfc-championship-398 | uncategorized | 0.0190 | P0 | precio fuera de [0.05, 0.95]: 0.0190 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-the-new-orleans-saints-win-the-2027-nfl-league-championship | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-olivier-faure-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-bernard-cazeneuve-win-the-2027-french-presidential-election | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-franois-bayrou-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-michel-barnier-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-carole-delga-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-manuel-bompard-win-the-2027-french-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-deutsche-bank-or-any-of-its-underwriting-affiliates-serve-as-the-lead-underwriter-in-spacexs-initial-public-offe... | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-eric-trump-win-the-2028-us-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-ron-desantis-win-the-2028-us-presidential-election | uncategorized | 0.0285 | P0 | precio fuera de [0.05, 0.95]: 0.0285 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-tucker-carlson-win-the-2028-us-presidential-election | uncategorized | 0.0335 | P0 | precio fuera de [0.05, 0.95]: 0.0335 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-joe-kent-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-andrew-yang-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-steve-bannon-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-lebron-james-win-the-2028-us-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-michelle-obama-win-the-2028-us-presidential-election | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-gretchen-whitmer-win-the-2028-democratic-presidential-nomination-676 | uncategorized | 0.0125 | P0 | precio fuera de [0.05, 0.95]: 0.0125 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-katie-britt-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-tom-brady-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-vivek-ramaswamy-win-the-2028-republican-presidential-nomination | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-thomas-massie-win-the-2028-republican-presidential-nomination | uncategorized | 0.0155 | P0 | precio fuera de [0.05, 0.95]: 0.0155 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-john-thune-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-erika-kirk-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-tim-walz-win-the-2028-us-presidential-election | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-nikki-haley-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-gretchen-whitmer-win-the-2028-us-presidential-election | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-hillary-clinton-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-wes-moore-win-the-2028-us-presidential-election | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-robert-f-kennedy-jr-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-josh-shapiro-win-the-2028-us-presidential-election | uncategorized | 0.0310 | P0 | precio fuera de [0.05, 0.95]: 0.0310 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-byron-donalds-win-the-2028-republican-presidential-nomination | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-elise-stefanik-win-the-2028-republican-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-oceania-win-the-2026-fifa-world-cup | uncategorized | 0.0020 | P0 | precio fuera de [0.05, 0.95]: 0.0020 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-netflix-close-warner-bros-acquisition | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | reya-fdv-above-1b-one-day-after-launch-348-347 | uncategorized | 0.0090 | P0 | precio fuera de [0.05, 0.95]: 0.0090 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-comcast-close-warner-bros-acquisition | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T16:30:04Z | will-asia-win-the-2026-fifa-world-cup | uncategorized | 0.0225 | P0 | precio fuera de [0.05, 0.95]: 0.0225 (long-tail no rentable en sim) |
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
