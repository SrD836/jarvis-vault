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
| 2026-05-19T11:52:49Z | will-dominic-solanke-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-graham-platner-win-the-2028-democratic-presidential-nomination | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-trump-and-putin-meet-next-in-south-korea-582-297-268 | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | elon-musk-of-tweets-may-15-may-22-340-359 | uncategorized | 0.0040 | P0 | precio fuera de [0.05, 0.95]: 0.0040 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-sebastian-korda-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | bitcoin-above-82k-on-may-19 | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-zelenskyy-and-putin-meet-next-in-ukraine | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-darren-fletcher-be-appointed-as-manager-of-manchester-united | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-hassan-shariatmadari-be-head-of-state-in-iran-end-of-2026 | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | will-vladimir-prebili-be-the-next-prime-minister-of-slovenia | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | uncategorized | 0.0415 | P0 | precio fuera de [0.05, 0.95]: 0.0415 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | ukraine-agrees-not-to-join-nato-by-june-30 | uncategorized | 0.0395 | P0 | precio fuera de [0.05, 0.95]: 0.0395 (long-tail no rentable en sim) |
| 2026-05-19T11:52:49Z | epl-che-tot-2026-05-19-che | uncategorized | 0.4950 | V2 | catalyst dentro de ventana 72h pre-resolución: resolución en 7 horas |
| 2026-05-19T11:52:50Z | will-alex-michelsen-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-jan-lennard-struff-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-austria-recognize-palestine-before-2027 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-jensen-huang-be-richest-person-on-december-31-229 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | bitcoin-above-78k-on-may-19 | uncategorized | 0.0850 | V2 | catalyst dentro de ventana 72h pre-resolución: resolución en 4 horas |
| 2026-05-19T11:52:50Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | uncategorized | 0.0185 | P0 | precio fuera de [0.05, 0.95]: 0.0185 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-parti-conservateur-du-qubec-win-the-most-seats-in-the-2026-quebec-general-election | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-fiorella-molinelli-win-the-2026-peruvian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-michael-jackson-be-confirmed-to-have-visited-epsteins-island | uncategorized | 0.0205 | P0 | precio fuera de [0.05, 0.95]: 0.0205 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | uncategorized | 0.0135 | P0 | precio fuera de [0.05, 0.95]: 0.0135 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | reya-fdv-above-1b-one-day-after-launch-348-347 | uncategorized | 0.0090 | P0 | precio fuera de [0.05, 0.95]: 0.0090 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | pete-hegseth-impeached-by-june-30 | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | donald-trump-of-truth-social-posts-may-12-may-19-100-119 | uncategorized | 0.0020 | P0 | precio fuera de [0.05, 0.95]: 0.0020 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-kang-seung-kyu-win-the-2026-chungcheongnam-province-gubernatorial-election | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-sung-jae-im-win-the-2026-masters-tournament | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-taylor-fritz-be-the-2026-mens-wimbledon-winner | uncategorized | 0.0235 | P0 | precio fuera de [0.05, 0.95]: 0.0235 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-washington-commanders-win-the-2027-nfl-nfc-championship-484 | uncategorized | 0.0235 | P0 | precio fuera de [0.05, 0.95]: 0.0235 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-bukayo-saka-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026 | uncategorized | 0.0105 | P0 | precio fuera de [0.05, 0.95]: 0.0105 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-jos-williams-win-the-2026-peruvian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-silver-si-hit-low-45-by-end-of-june-972-272 | uncategorized | 0.0160 | P0 | precio fuera de [0.05, 0.95]: 0.0160 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-beyonc-be-the-top-spotify-artist-for-2026 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | kurds-declare-independence-from-iran | uncategorized | 0.0460 | P0 | precio fuera de [0.05, 0.95]: 0.0460 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-openais-market-cap-be-between-750b-and-1t-at-market-close-on-ipo-day | uncategorized | 0.0230 | P0 | precio fuera de [0.05, 0.95]: 0.0230 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-larry-page-be-richest-person-on-december-31 | uncategorized | 0.0065 | P0 | precio fuera de [0.05, 0.95]: 0.0065 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-ethereum-reach-3800-in-may-2026 | uncategorized | 0.0020 | P0 | precio fuera de [0.05, 0.95]: 0.0020 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-zelenskyy-and-putin-meet-next-in-russia | uncategorized | 0.0185 | P0 | precio fuera de [0.05, 0.95]: 0.0185 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | uncategorized | 0.0075 | P0 | precio fuera de [0.05, 0.95]: 0.0075 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-san-francisco-giants-win-the-2026-national-league-championship-series | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-grigor-dimitrov-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-scott-jensen-win-the-2026-minnesota-governor-republican-primary-election | uncategorized | 0.0090 | P0 | precio fuera de [0.05, 0.95]: 0.0090 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-lew-evans-win-the-bachelorette-season-22 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | us-x-iran-permanent-peace-deal-by-may-22-2026 | uncategorized | 0.0365 | P0 | precio fuera de [0.05, 0.95]: 0.0365 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-lorenzo-musetti-be-the-2026-mens-wimbledon-winner | uncategorized | 0.0085 | P0 | precio fuera de [0.05, 0.95]: 0.0085 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-israel-strike-11-countries-in-2026 | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-hubert-hurkacz-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-thierno-barry-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-joao-pedro-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-colorado-rockies-win-the-2026-national-league-championship-series | uncategorized | 0.0035 | P0 | precio fuera de [0.05, 0.95]: 0.0035 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-cleveland-browns-win-the-2027-nfl-afc-championship-776 | uncategorized | 0.0225 | P0 | precio fuera de [0.05, 0.95]: 0.0225 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-omar-marmoush-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | ethereum-above-2000-on-may-19 | uncategorized | 0.9875 | P0 | precio fuera de [0.05, 0.95]: 0.9875 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-mark-zuckerberg-be-richest-person-on-december-31-366 | uncategorized | 0.0120 | P0 | precio fuera de [0.05, 0.95]: 0.0120 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-alex-baena-be-the-top-goal-scorer-in-the-202526-la-liga-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-the-us-strike-15-or-more-countries-in-2026 | uncategorized | 0.0220 | P0 | precio fuera de [0.05, 0.95]: 0.0220 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-kai-havertz-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | taiwanese-premier-cho-jung-tai-out-by-june-30-2026 | uncategorized | 0.0440 | P0 | precio fuera de [0.05, 0.95]: 0.0440 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | standx-fdv-above-5b-one-day-after-launch-536-723-718-827-974 | uncategorized | 0.0175 | P0 | precio fuera de [0.05, 0.95]: 0.0175 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | timothy-chalamet-confirmed-to-be-esdeekid-by-june-30 | uncategorized | 0.0115 | P0 | precio fuera de [0.05, 0.95]: 0.0115 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | mh370-underwater-wreckage-found-by-june-30-2026 | uncategorized | 0.0195 | P0 | precio fuera de [0.05, 0.95]: 0.0195 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | donald-trump-of-truth-social-posts-may-12-may-19-120-139 | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | kadyrov-out-as-head-of-the-chechen-republic-by-june-30-2026 | uncategorized | 0.0395 | P0 | precio fuera de [0.05, 0.95]: 0.0395 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-chris-wood-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-ollie-watkins-be-the-top-goal-scorer-in-the-202526-english-premier-league-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-reilly-opelka-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | bitcoin-above-74k-on-may-20 | uncategorized | 0.9685 | P0 | precio fuera de [0.05, 0.95]: 0.9685 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-denis-shapovalov-win-the-2026-mens-french-open | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-ian-calderon-get-the-first-or-second-most-votes-in-the-2026-california-governor-primary-election | uncategorized | 0.0145 | P0 | precio fuera de [0.05, 0.95]: 0.0145 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | uncategorized | 0.0025 | P0 | precio fuera de [0.05, 0.95]: 0.0025 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-discords-market-cap-be-between-15b-and-20b-at-market-close-on-ipo-day | uncategorized | 0.0415 | P0 | precio fuera de [0.05, 0.95]: 0.0415 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-robert-lewandowski-be-the-top-goal-scorer-in-the-202526-la-liga-season | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-jared-golden-be-the-democratic-nominee-for-senate-in-maine | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-jasmine-crockett-and-ken-paxton-be-the-candidates-for-the-texas-senate-election-897 | uncategorized | 0.0005 | P0 | precio fuera de [0.05, 0.95]: 0.0005 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-coco-gauff-be-the-2026-womens-wimbledon-winner | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-colorado-rapids-win-the-2026-mls-cup | uncategorized | 0.0095 | P0 | precio fuera de [0.05, 0.95]: 0.0095 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-gold-gc-hit-high-6200-by-end-of-june | uncategorized | 0.0175 | P0 | precio fuera de [0.05, 0.95]: 0.0175 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | will-zelenskyy-and-putin-meet-next-in-saudi-arabia | uncategorized | 0.0145 | P0 | precio fuera de [0.05, 0.95]: 0.0145 (long-tail no rentable en sim) |
| 2026-05-19T11:52:50Z | openai-receives-federal-backstop-for-infrastructure-before-july | uncategorized | 0.0275 | P0 | precio fuera de [0.05, 0.95]: 0.0275 (long-tail no rentable en sim) |
| 2026-05-19T11:54:22Z | will-the-iranian-regime-fall-by-june-30 | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T11:54:35Z | will-the-iranian-regime-fall-by-june-30 | uncategorized | 0.0450 | P0 | precio fuera de [0.05, 0.95]: 0.0450 (long-tail no rentable en sim) |
| 2026-05-19T11:54:42Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449 | uncategorized | 0.1150 | N1 | noticias contradicen tesis: Trump rechazó propuesta de paz de Irán, indicando oposición a un acuerdo permanente |
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
