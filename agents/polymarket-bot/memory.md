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
| 2026-05-30T11:03:18Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $504) |
| 2026-05-30T11:03:18Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $582) |
| 2026-05-30T11:03:18Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $189880) |
| 2026-05-30T11:03:18Z | will-any-category-5-hurricane-make-landfall-in-the-us-in-before-2027 | weather-natural | 0.1300 | P3_low_absolute_liquidity | liquidity $1663 < absolute min $5000 |
| 2026-05-30T11:03:18Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1586 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $320) |
| 2026-05-30T11:03:18Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $13230) |
| 2026-05-30T11:03:18Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1219 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $11288) |
| 2026-05-30T11:03:18Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $202) |
| 2026-05-30T11:03:18Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $255 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4999 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $638) |
| 2026-05-30T11:03:18Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $268 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $410) |
| 2026-05-30T11:03:18Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $694) |
| 2026-05-30T11:03:18Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $195 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $18684) |
| 2026-05-30T11:03:18Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $570) |
| 2026-05-30T11:03:18Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4656 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $7290) |
| 2026-05-30T11:03:18Z | will-there-be-exactly-1-confirmed-vei-4-or-higher-volcanic-eruption-worldwide-in-2026-937 | weather-natural | 0.2700 | P3_low_absolute_liquidity | liquidity $1362 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 214 d, liq $9115) |
| 2026-05-30T11:03:18Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2653 < absolute min $5000 |
| 2026-05-30T11:03:18Z | jerome-powell-out-from-fed-board-by-may-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $45291) |
| 2026-05-30T11:03:18Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $221) |
| 2026-05-30T11:03:18Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4635 < absolute min $5000 |
| 2026-05-30T11:03:18Z | 5kt-meteor-strike-in-2026 | weather-natural | 0.3200 | P3_low_absolute_liquidity | liquidity $3063 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $7923) |
| 2026-05-30T11:03:18Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $828) |
| 2026-05-30T11:03:18Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $69268) |
| 2026-05-30T11:03:18Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $12535) |
| 2026-05-30T11:03:18Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $6069) |
| 2026-05-30T11:03:18Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1230 | P3_low_absolute_liquidity | liquidity $3124 < absolute min $5000 |
| 2026-05-30T11:03:18Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T11:03:18Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 214 d, liq $9617) |
| 2026-05-30T11:03:18Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 214 d, liq $14170) |
| 2026-05-30T11:03:18Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $439) |
| 2026-05-30T11:03:18Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $283) |
| 2026-05-30T11:03:18Z | major-meteor-strike-10kt-in-2026 | weather-natural | 0.1300 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T11:03:18Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $17259) |
| 2026-05-30T11:03:18Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $5918) |
| 2026-05-30T11:03:18Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3972 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $6244) |
| 2026-05-30T11:03:18Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 215 d, liq $3643) |
| 2026-05-30T11:03:18Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $654 < absolute min $5000 |
| 2026-05-30T11:03:18Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1300 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3800 | P3_low_absolute_liquidity | liquidity $1127 < absolute min $5000 |
| 2026-05-30T11:03:18Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3298 < absolute min $5000 |
| 2026-05-30T11:03:18Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3200 | P3_low_absolute_liquidity | liquidity $572 < absolute min $5000 |
| 2026-05-30T11:03:18Z | opensea-fdv-above-100m-one-day-after-launch-172-151-588-987 | crypto-launch | 0.6000 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | opensea-fdv-above-5b-one-day-after-launch-919-666-825-992-543-333 | crypto-launch | 0.0500 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | altcoin-market-cap-dip-to-150b-before-2027 | other | 0.7200 | P3_low_absolute_liquidity | liquidity $858 < absolute min $5000 |
| 2026-05-30T11:03:18Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0910 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2384 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $1450) |
| 2026-05-30T11:03:18Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P3_low_absolute_liquidity | liquidity $770 < absolute min $5000 |
| 2026-05-30T11:03:18Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $12639) |
| 2026-05-30T11:03:18Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $52849) |
| 2026-05-30T11:03:18Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 239 d, liq $45928) |
| 2026-05-30T11:03:18Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 239 d, liq $19247) |
| 2026-05-30T11:03:18Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $269499) |
| 2026-05-30T11:03:18Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $46209) |
| 2026-05-30T11:03:18Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $43593) |
| 2026-05-30T11:03:18Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3988 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $24242) |
| 2026-05-30T11:03:18Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8200 | P4_pre_event | pre-event slug + 579 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2210 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.2400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T11:03:18Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1625518) |
| 2026-05-30T11:03:18Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1075178) |
| 2026-05-30T11:03:18Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1447299) |
| 2026-05-30T11:03:18Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $946759) |
| 2026-05-30T11:03:18Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2437835) |
| 2026-05-30T11:03:18Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1167865) |
| 2026-05-30T11:03:18Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 891 d, liq $1916256) |
| 2026-05-30T11:03:18Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1724462) |
| 2026-05-30T11:03:18Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1311097) |
| 2026-05-30T11:03:18Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $936817) |
| 2026-05-30T11:03:18Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $474818) |
| 2026-05-30T11:03:18Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1480392) |
| 2026-05-30T11:03:18Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1151886) |
| 2026-05-30T11:03:18Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2365119) |
| 2026-05-30T11:03:18Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1059391) |
| 2026-05-30T11:03:18Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2333830) |
| 2026-05-30T11:03:18Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1457190) |
| 2026-05-30T11:03:18Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $5063) |
| 2026-05-30T11:03:18Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2600 | P3_low_absolute_liquidity | liquidity $569 < absolute min $5000 |
| 2026-05-30T11:03:18Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $1208) |
| 2026-05-30T11:03:18Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $13107) |
| 2026-05-30T11:03:18Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0400 | P3_low_absolute_liquidity | liquidity $2944 < absolute min $5000 |
| 2026-05-30T11:03:18Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $23046) |
| 2026-05-30T11:30:03Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon -150 d, liq $49118) |
| 2026-05-30T11:30:03Z | ukraine-recognizes-russian-sovereignty-over-its-territory-by-june-30-2026 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -149 d, liq $13494) |
| 2026-05-30T11:30:03Z | new-stranger-things-episode-released-by-december-31-587-657-824-246 | other | 0.0700 | P2 | mercado ya expiró (endDate=2026-01-07T00:00:00Z, hace 143 días) |
| 2026-05-30T11:30:03Z | will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -111 d, liq $93209) |
| 2026-05-30T11:30:03Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $766688) |
| 2026-05-30T11:30:03Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $776104) |
| 2026-05-30T11:30:03Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -69 d, liq $32355) |
| 2026-05-30T11:30:03Z | will-troels-lund-poulsen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon -67 d, liq $20239) |
| 2026-05-30T11:30:03Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -67 d, liq $28498) |
| 2026-05-30T11:30:03Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $20598) |
| 2026-05-30T11:30:03Z | will-alex-vanopslagh-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $30147) |
| 2026-05-30T11:30:03Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -67 d, liq $28715) |
| 2026-05-30T11:30:03Z | weed-rescheduled-by-june-30 | other | 0.0210 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T11:30:03Z | weed-rescheduled-by-december-31 | other | 0.3020 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T11:30:03Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -48 d, liq $201270) |
| 2026-05-30T11:30:03Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $345155) |
| 2026-05-30T11:30:03Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon -15 d, liq $105917) |
| 2026-05-30T11:30:03Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.6400 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 15 días) |
| 2026-05-30T11:30:03Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -13 d, liq $5944) |
| 2026-05-30T11:30:04Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -11 d, liq $8892) |
| 2026-05-30T11:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -3 d, liq $1699778) |
| 2026-05-30T11:30:04Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $14833) |
| 2026-05-30T11:30:04Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $14599) |
| 2026-05-30T11:30:04Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.0600 | P9 | P9: geopolitics pump cluster (price 0.06, 0d) |
| 2026-05-30T11:30:04Z | bitcoin-above-70k-on-may-30-2026 | market | 0.9980 | P0_ceiling | price ceiling: 0.9980 > 0.980 |
| 2026-05-30T11:30:06Z | elon-musk-of-tweets-may-28-may-30-0-39 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $78211) |
| 2026-05-30T11:30:06Z | bitcoin-above-76k-on-may-30-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $41465) |
| 2026-05-30T11:30:06Z | bitcoin-above-72k-on-may-30-2026 | market | 0.9900 | P0_ceiling | price ceiling: 0.9900 > 0.980 |
| 2026-05-30T11:30:06Z | elon-musk-of-tweets-may-28-may-30-90-114 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $101028) |
| 2026-05-30T11:30:09Z | elon-musk-of-tweets-may-28-may-30-65-89 | other | 0.1200 | E2 | edge 0.020 < mín 0.030 (p̂=0.100, implied=0.120) |
| 2026-05-30T11:30:12Z | ucl-psg-ars-2026-05-30-psg | other | 0.4000 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T11:30:23Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0460 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T11:30:25Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.0410 | E2 | edge 0.021 < mín 0.030 (p̂=0.020, implied=0.041) |
| 2026-05-30T11:30:25Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.0510 | P9 | P9: geopolitics pump cluster (price 0.05, 0d) |
| 2026-05-30T11:30:25Z | uefa-champions-league-unbeaten-team | sports-season | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-30T11:30:25Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $162294) |
| 2026-05-30T11:30:25Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 0 d, liq $61485) |
| 2026-05-30T11:30:25Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $3895831) |
| 2026-05-30T11:30:27Z | will-40-ships-transit-the-strait-of-hormuz-on-any-day-by-may-31 | geopolitics | 0.0580 | P9 | P9: geopolitics pump cluster (price 0.06, 0d) |
| 2026-05-30T11:30:27Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 0 d, liq $602489) |
| 2026-05-30T11:30:27Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 0 d, liq $107319) |
| 2026-05-30T11:30:27Z | will-reza-pahlavi-enter-iran-by-may-31 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $186462) |
| 2026-05-30T11:30:27Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $220568) |
| 2026-05-30T11:30:27Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $186921) |
| 2026-05-30T11:30:28Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $198008) |
| 2026-05-30T11:30:28Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.1900 | P9 | P9: geopolitics pump cluster (price 0.19, 0d) |
| 2026-05-30T11:30:29Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.0330 | E2 | edge 0.007 < mín 0.030 (p̂=0.040, implied=0.033) |
| 2026-05-30T11:30:29Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $200461) |
| 2026-05-30T11:30:32Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0250 | E2 | edge 0.020 < mín 0.030 (p̂=0.005, implied=0.025) |
| 2026-05-30T11:30:32Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 0 d, liq $62560) |
| 2026-05-30T11:30:32Z | mojtaba-khamenei-leaves-iran-by-may-31-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $252376) |
| 2026-05-30T11:30:34Z | will-trump-restart-project-freedom-by-may-31 | other | 0.0350 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. 'Project Freedom' no es un térmi... |
| 2026-05-30T11:30:36Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0360 | V3 Trigger vago & V6 Sin catalyst | V3 Trigger vago & V6 Sin catalyst: V3: Trigger vago. 'Permanent peace deal' carece de definición concreta: no hay un... |
| 2026-05-30T11:30:36Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 0 d, liq $110982) |
| 2026-05-30T11:30:36Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 0 d, liq $205100) |
| 2026-05-30T11:30:36Z | will-alibaba-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $136102) |
| 2026-05-30T11:30:39Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0250 | E2 | edge 0.015 < mín 0.030 (p̂=0.010, implied=0.025) |
| 2026-05-30T11:30:46Z | will-arsenal-win-the-202526-champions-league | sports-season | 0.4400 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: El mercado carece de un catalizador ... |
| 2026-05-30T11:30:48Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0280 | E2 | edge 0.018 < mín 0.030 (p̂=0.010, implied=0.028) |
| 2026-05-30T11:30:48Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $97270) |
| 2026-05-30T11:30:48Z | will-nvidia-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $38451) |
| 2026-05-30T11:30:48Z | will-alphabet-be-the-third-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $18677) |
| 2026-05-30T11:30:48Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116472) |
| 2026-05-30T11:30:48Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $97055) |
| 2026-05-30T11:30:48Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $157276) |
| 2026-05-30T11:30:48Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $149905) |
| 2026-05-30T11:30:48Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $147666) |
| 2026-05-30T11:30:48Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $100949) |
| 2026-05-30T11:30:48Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $125447) |
| 2026-05-30T11:30:48Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $157125) |
| 2026-05-30T11:30:48Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $147542) |
| 2026-05-30T11:30:48Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $142461) |
| 2026-05-30T11:30:48Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $115798) |
| 2026-05-30T11:30:48Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $98843) |
| 2026-05-30T11:30:48Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 1 d, liq $307684) |
| 2026-05-30T11:30:48Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $744225) |
| 2026-05-30T11:30:48Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $585908) |
| 2026-05-30T11:30:48Z | will-wti-dip-to-70-in-may-2026-155-395-333-182-889-988-959-341-883-234-375-347-135 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $210846) |
| 2026-05-30T11:30:48Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $1295044) |
| 2026-05-30T11:30:48Z | over-30m-committed-to-the-printr-public-sale | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $13413) |
| 2026-05-30T11:30:48Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $263926) |
| 2026-05-30T11:30:52Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $165769) |
| 2026-05-30T11:30:52Z | will-ethereum-dip-to-1600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $43085) |
| 2026-05-30T11:30:52Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $121169) |
| 2026-05-30T11:30:52Z | will-ethereum-reach-2600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $59916) |
| 2026-05-30T11:30:52Z | over-60m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $13086) |
| 2026-05-30T11:30:52Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 1 d, liq $34981) |
| 2026-05-30T11:30:52Z | will-bitcoin-reach-77500-in-may-2026-from-may-27 | market | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 1 d, liq $27353) |
| 2026-05-30T11:30:52Z | will-bitcoin-reach-80000-in-may-2026-from-may-27 | market | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 1 d, liq $37890) |
| 2026-05-30T11:30:52Z | will-ethereum-dip-to-400-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $43043) |
| 2026-05-30T11:30:52Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $27296) |
| 2026-05-30T11:30:52Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $237503) |
| 2026-05-30T11:30:55Z | will-xrp-reach-1pt6-in-may-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $14706) |
| 2026-05-30T11:30:55Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $12265) |
| 2026-05-30T11:30:57Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0350 | V3: Si la pregunta no tiene un evento binario verificable con fecha concreta, se bloquea. | V3: Si la pregunta no tiene un evento binario verificable con fecha concreta, se bloquea.: V3 Trigger vago: la pregun... |
| 2026-05-30T11:30:59Z | will-backrooms-opening-weekend-box-office-be-greater-than-79m | entertainment | 0.8560 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T11:31:02Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2100 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: No hay evento discreto identificable en los próximos 7 días que pueda cambiar sig... |
| 2026-05-30T11:31:02Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $18196) |
| 2026-05-30T11:31:02Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $38568) |
| 2026-05-30T11:31:02Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 3 d, liq $44824) |
| 2026-05-30T11:31:02Z | elon-musk-of-tweets-may-26-june-2-320-339 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 3 d, liq $40517) |
| 2026-05-30T11:31:02Z | elon-musk-of-tweets-may-26-june-2-100-119 | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 3 d, liq $57439) |
| 2026-05-30T11:31:02Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $23247) |
| 2026-05-30T11:31:02Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $20237) |
| 2026-05-30T11:31:02Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $20841) |
| 2026-05-30T11:31:02Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $26382) |
| 2026-05-30T11:31:02Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $43526) |
| 2026-05-30T11:31:02Z | will-kang-seung-kyu-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $24030) |
| 2026-05-30T11:31:02Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $31243) |
| 2026-05-30T11:31:02Z | will-yoon-sang-hyun-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $38472) |
| 2026-05-30T11:31:02Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $45833) |
| 2026-05-30T11:31:02Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $44360) |
| 2026-05-30T11:31:04Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $32240) |
| 2026-05-30T11:31:04Z | will-park-hong-keun-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $476681) |
| 2026-05-30T11:31:04Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $32483) |
| 2026-05-30T11:31:04Z | will-na-kyung-won-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $475130) |
| 2026-05-30T11:31:04Z | will-han-dong-hoon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $33329) |
| 2026-05-30T11:31:04Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $97189) |
| 2026-05-30T11:31:04Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $44570) |
| 2026-05-30T11:31:04Z | will-park-yong-jin-win-the-2026-seoul-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $457264) |
| 2026-05-30T11:31:07Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $40016) |
| 2026-05-30T11:31:07Z | will-chung-jin-suk-win-the-2026-chungcheongnam-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $19886) |
| 2026-05-30T11:31:07Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $19766) |
| 2026-05-30T11:31:07Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $95264) |
| 2026-05-30T11:31:07Z | us-announces-new-iran-agreementceasefire-extension-by-june-3-569 | geopolitics | 0.2900 | P9 | P9: geopolitics pump cluster (price 0.29, 3d) |
| 2026-05-30T11:31:07Z | elon-musk-of-tweets-may-29-june-5-500plus | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 6 d, liq $18365) |
| 2026-05-30T11:31:07Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $230324) |
| 2026-05-30T11:31:07Z | will-moise-kouame-win-the-2026-mens-french-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 7 d, liq $32971) |
| 2026-05-30T11:31:07Z | will-jakub-mensik-win-the-2026-mens-french-open | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 7 d, liq $56616) |
| 2026-05-30T11:31:09Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0200 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-30T11:31:09Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $201461) |
| 2026-05-30T11:31:09Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $123151) |
| 2026-05-30T11:31:09Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $210076) |
| 2026-05-30T11:31:09Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 7 d, liq $40992) |
| 2026-05-30T11:31:11Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0700 | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... |
| 2026-05-30T11:31:13Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0640 | E2 | edge 0.004 < mín 0.030 (p̂=0.060, implied=0.064) |
| 2026-05-30T11:31:13Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $220493) |
| 2026-05-30T11:31:16Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0290 | E2 | edge 0.014 < mín 0.030 (p̂=0.015, implied=0.029) |
| 2026-05-30T11:31:16Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $198367) |
| 2026-05-30T11:31:18Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0750 | V3 | V3: V3 Trigger vago: la pregunta carece de un evento discreto y verificable con fecha concreta; se trata de un mercad... |
| 2026-05-30T11:31:18Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $189914) |
| 2026-05-30T11:31:21Z | will-raphael-collignon-win-the-2026-mens-french-open | other | 0.0370 | V3 Trigger vago / V6 Sin catalyst | V3 Trigger vago / V6 Sin catalyst: V3 Trigger vago: sin fecha concreta o sin evento verificable. La pregunta no espec... |
| 2026-05-30T11:31:21Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.1500 | P9 | P9: geopolitics pump cluster (price 0.15, 7d) |
| 2026-05-30T11:31:21Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $159949) |
| 2026-05-30T11:31:21Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $227532) |
| 2026-05-30T11:31:21Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $247636) |
| 2026-05-30T11:31:33Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1040 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-30T11:31:36Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0560 | V3 | V3: V3 Trigger vago: no hay un evento discreto y verificable asociado a la pregunta. La probabilidad de que Cerundolo... |
| 2026-05-30T11:31:38Z | will-martin-landaluce-win-the-2026-mens-french-open | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 7 d, liq $43947) |
| 2026-05-30T11:31:40Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0240 | E2 | edge 0.014 < mín 0.030 (p̂=0.010, implied=0.024) |
| 2026-05-30T11:31:40Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $160562) |
| 2026-05-30T11:31:40Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $222209) |
| 2026-05-30T11:31:40Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 7 d, liq $189619) |
| 2026-05-30T11:31:40Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $213656) |
| 2026-05-30T11:31:40Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $112222) |
| 2026-05-30T11:31:40Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $212379) |
| 2026-05-30T11:31:48Z | will-al-mina-be-the-republican-nominee-for-senate-in-virginia | elections | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 16 d, liq $22030) |
| 2026-05-30T11:31:50Z | will-the-san-antonio-spurs-win-the-nba-western-conference-finals | sports-season | 0.4240 | V3 Trigger vago | V3 Trigger vago: V3 Trigger vago: La pregunta no especifica una temporada concreta, ni hay un evento verificable con ... |
| 2026-05-30T11:31:50Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 17 d, liq $816766) |
| 2026-05-30T11:31:50Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 17 d, liq $2640635) |
| 2026-05-30T11:31:50Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 17 d, liq $336281) |
| 2026-05-30T11:31:50Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9820 | P0_ceiling | price ceiling: 0.9820 > 0.980 |
| 2026-05-30T11:31:50Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 17 d, liq $796105) |
| 2026-05-30T11:31:50Z | will-m80-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $46668) |
| 2026-05-30T11:31:50Z | will-sinners-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $44312) |
| 2026-05-30T11:31:50Z | will-thunder-downunder-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $76773) |
| 2026-05-30T11:31:50Z | will-lynn-vision-win-iem-cologne-major-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 21 d, liq $19599) |
| 2026-05-30T11:31:50Z | will-flyquest-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $73334) |
| 2026-05-30T11:31:50Z | will-nrg-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $77189) |
| 2026-05-30T11:31:50Z | will-gaimin-gladiators-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $82425) |
| 2026-05-30T11:31:50Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $156920) |
| 2026-05-30T11:31:50Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $203207) |
| 2026-05-30T11:31:50Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 22 d, liq $72014) |
| 2026-05-30T11:31:53Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0280 | E2 | edge 0.002 < mín 0.030 (p̂=0.030, implied=0.028) |
| 2026-05-30T11:31:58Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $199672) |
| 2026-05-30T11:31:58Z | will-juan-daniel-oviedo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $257741) |
| 2026-05-30T11:31:58Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $226034) |
| 2026-05-30T11:31:58Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $184773) |
| 2026-05-30T11:31:58Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $195679) |
| 2026-05-30T11:31:58Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $233983) |
| 2026-05-30T11:31:58Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $215202) |
| 2026-05-30T11:32:01Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E2 | edge 0.010 < mín 0.030 (p̂=0.030, implied=0.020) |
| 2026-05-30T11:32:01Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 24 d, liq $1124) |
| 2026-05-30T11:32:01Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 24 d, liq $684) |
| 2026-05-30T11:32:01Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 24 d, liq $529) |
| 2026-05-30T11:32:01Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 27 d, liq $1683) |
| 2026-05-30T11:32:01Z | will-trump-and-putin-meet-next-in-south-korea-582-297-268 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $22287) |
| 2026-05-30T11:32:04Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0340 | E2 | edge 0.014 < mín 0.030 (p̂=0.020, implied=0.034) |
| 2026-05-30T11:32:04Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 30 d, liq $13062) |
| 2026-05-30T11:32:06Z | will-deepseek-have-the-second-best-ai-model-at-the-end-of-june-2026 | other | 0.0230 | E2 | edge 0.027 < mín 0.030 (p̂=0.050, implied=0.023) |
| 2026-05-30T11:32:11Z | mojtaba-khamenei-leaves-iran-by-june-30-2026 | geopolitics | 0.0270 | E2 | edge 0.013 < mín 0.030 (p̂=0.040, implied=0.027) |
| 2026-05-30T11:32:11Z | israel-withdraws-from-lebanon-by-may-31-2026 | geopolitics | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 30 d, liq $43770) |
| 2026-05-30T11:32:13Z | will-trump-and-putin-meet-next-in-ukraine-328-474 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $17336) |
| 2026-05-30T11:32:19Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-30-2026-15... | geopolitics | 0.7100 | V6 | V6: V6 Sin catalyst: No hay evento discreto identificable en los próximos 7 días que condicione el anuncio. La fech... |
| 2026-05-30T11:32:20Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda confirmar la visita de J... |
| 2026-05-30T11:32:22Z | european-country-agrees-to-give-ukraine-security-guarantee-by-june-30 | geopolitics | 0.0400 | E2 | edge 0.010 < mín 0.030 (p̂=0.050, implied=0.040) |
| 2026-05-30T11:32:22Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $348153) |
| 2026-05-30T11:32:24Z | will-moonshot-have-the-top-ai-model-at-the-end-of-june-2026 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $11280) |
| 2026-05-30T11:32:24Z | hantavirus-lab-leak-confirmed-by-june-30-1 | weather-natural | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 30 d, liq $111067) |
| 2026-05-30T11:32:26Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0220 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: 'fall of the regime' es un concepto a... |
| 2026-05-30T11:32:28Z | will-the-montreal-canadiens-win-the-eastern-conference | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $387278) |
| 2026-05-30T11:32:28Z | will-trump-and-putin-meet-next-in-belarus-572 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $18218) |
| 2026-05-30T11:32:28Z | anthropic-ceo-arrested | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $8192) |
| 2026-05-30T11:32:28Z | will-microsoft-be-the-largest-company-in-the-world-by-market-cap-on-june-30-612 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $244523) |
| 2026-05-30T11:32:33Z | israel-x-iran-permanent-peace-deal-by-june-30-2026-262 | geopolitics | 0.0980 | V3: Trigger vago - falta de definición concreta y verificable. V6: Sin catalyst en ventana de 7 días. | V3: Trigger vago - falta de definición concreta y verificable. V6: Sin catalyst en ventana de 7 días.: V3 Trigger v... |
| 2026-05-30T11:32:33Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $16255) |
| 2026-05-30T11:32:38Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0320 | E2 | edge 0.018 < mín 0.030 (p̂=0.050, implied=0.032) |
| 2026-05-30T11:32:40Z | will-meituan-have-the-top-ai-model-at-the-end-of-june-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 30 d, liq $11022) |
| 2026-05-30T11:32:40Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $387376) |
| 2026-05-30T11:32:42Z | jerome-powell-federally-charged-by-june-30 | other | 0.0380 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T11:32:42Z | will-tim-walz-resign-by-june-30 | executive-action | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 30 d, liq $26410) |
| 2026-05-30T11:32:42Z | will-openais-market-cap-be-less-than-500b-at-market-close-on-ipo-day | crypto-launch | 0.0200 | P4_pre_event | pre-event slug + 30 d to resolution (>=7 threshold) |
| 2026-05-30T11:32:44Z | iran-closes-its-airspace-by-june-30-432-786-462-866 | geopolitics | 0.2210 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda desencadena... |
| 2026-05-30T11:32:44Z | will-any-country-leave-nato-by-june-30-2026 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $23620) |
| 2026-05-30T11:32:44Z | will-silver-si-hit-high-150-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $37242) |
| 2026-05-30T11:32:44Z | will-gold-gc-hit-high-6500-by-end-of-june | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 31 d, liq $23274) |
| 2026-05-30T11:32:44Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $31184) |
| 2026-05-30T11:32:44Z | will-silver-si-hit-low-35-by-end-of-june-325-231 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 31 d, liq $20190) |
| 2026-05-30T11:32:44Z | will-silver-si-hit-high-130-by-end-of-june | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $34628) |
| 2026-05-30T11:32:44Z | will-crude-oil-cl-hit-low-50-by-end-of-june | market | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 31 d, liq $28982) |
| 2026-05-30T11:32:56Z | ethereum-all-time-high-by-june-30-2026 | market | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 31 d, liq $18788) |
| 2026-05-30T11:32:56Z | will-clara-tauson-be-the-2026-womens-wimbledon-winner | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 42 d, liq $2689) |
| 2026-05-30T11:32:56Z | will-ons-jabeur-be-the-2026-womens-wimbledon-winner | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 42 d, liq $2846) |
| 2026-05-30T11:32:56Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P3_low_absolute_liquidity | liquidity $3220 < absolute min $5000 |
| 2026-05-30T11:32:56Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0450 | P3_low_absolute_liquidity | liquidity $4082 < absolute min $5000 |
| 2026-05-30T11:32:56Z | will-victoria-mboko-be-the-2026-womens-wimbledon-winner | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 42 d, liq $3229) |
| 2026-05-30T11:32:56Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0200 | P3_low_absolute_liquidity | liquidity $3998 < absolute min $5000 |
| 2026-05-30T11:32:56Z | will-madison-keys-be-the-2026-womens-wimbledon-winner | other | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 42 d, liq $3397) |
| 2026-05-30T11:32:56Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $586 < absolute min $5000 |
| 2026-05-30T11:32:56Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $4228744) |
| 2026-05-30T11:32:56Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 50 d, liq $5001853) |
| 2026-05-30T11:32:59Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0270 | V3 \| V5 | V3 \| V5: V3 Trigger vago: sin fecha concreta ni evento verificable próximo. V5 Patrón débil: no hay fuentes indep... |
| 2026-05-30T11:32:59Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $7560794) |
| 2026-05-30T11:32:59Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 50 d, liq $5935233) |
| 2026-05-30T11:32:59Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $7408264) |
| 2026-05-30T11:32:59Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $9709902) |
| 2026-05-30T11:32:59Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8492556) |
| 2026-05-30T11:32:59Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $4505286) |
| 2026-05-30T11:32:59Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $10765017) |
| 2026-05-30T11:32:59Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $6634681) |
| 2026-05-30T11:32:59Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 50 d, liq $8415720) |
| 2026-05-30T11:32:59Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 50 d, liq $3743366) |
| 2026-05-30T11:33:01Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $5270091) |
| 2026-05-30T11:33:01Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $2781684) |
| 2026-05-30T11:33:04Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8620162) |
| 2026-05-30T11:33:04Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8066747) |
| 2026-05-30T11:33:04Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $7795009) |
| 2026-05-30T11:33:04Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 50 d, liq $3176287) |
| 2026-05-30T11:33:04Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 50 d, liq $9648971) |
| 2026-05-30T11:33:04Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 50 d, liq $2545975) |
| 2026-05-30T11:33:09Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $9624964) |
| 2026-05-30T11:33:09Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $5763339) |
| 2026-05-30T11:33:09Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 50 d, liq $405099) |
| 2026-05-30T11:33:11Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.008 < mín 0.030 (p̂=0.060, implied=0.052) |
| 2026-05-30T11:33:11Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $7252170) |
| 2026-05-30T11:33:15Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T11:33:18Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0920 | E2 | edge 0.028 < mín 0.030 (p̂=0.120, implied=0.092) |
| 2026-05-30T11:33:18Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $1153687) |
| 2026-05-30T11:33:18Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 50 d, liq $2490676) |
| 2026-05-30T11:33:18Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $9642885) |
| 2026-05-30T11:33:18Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $10948110) |
| 2026-05-30T11:33:18Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 50 d, liq $5459003) |
| 2026-05-30T11:33:18Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $2460775) |
| 2026-05-30T11:33:18Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $10047287) |
| 2026-05-30T11:33:20Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0370 | E2 | edge 0.003 < mín 0.030 (p̂=0.040, implied=0.037) |
| 2026-05-30T11:33:23Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $16960) |
| 2026-05-30T11:33:23Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $33639) |
| 2026-05-30T11:33:23Z | will-tyreek-hill-play-for-philadelphia-eagles-next | other | 0.1390 | P3_low_absolute_liquidity | liquidity $248 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0240 | P3_low_absolute_liquidity | liquidity $1016 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-david-njoku-play-for-miami-dolphins-in-2026-27 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 93 d, liq $2783) |
| 2026-05-30T11:33:23Z | will-david-njoku-play-for-denver-broncos-in-2026-27 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 93 d, liq $998) |
| 2026-05-30T11:33:23Z | will-david-njoku-play-for-indianapolis-colts-in-2026-27 | other | 0.0200 | P3_low_absolute_liquidity | liquidity $483 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-filipe-luis-be-the-next-manager-of-ssc-napoli | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 93 d, liq $1609) |
| 2026-05-30T11:33:23Z | will-david-njoku-play-for-kansas-city-chiefs-in-2026-27 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 93 d, liq $3877) |
| 2026-05-30T11:33:23Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $1982) |
| 2026-05-30T11:33:23Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $1871) |
| 2026-05-30T11:33:23Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 105 d, liq $2295) |
| 2026-05-30T11:33:23Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 105 d, liq $2950) |
| 2026-05-30T11:33:23Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $779 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 120 d, liq $1095) |
| 2026-05-30T11:33:23Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 120 d, liq $752) |
| 2026-05-30T11:33:23Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $281 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5700 | P8 | P8: election 122d out, price 0.57 en banda ruido [0.30, 0.70] |
| 2026-05-30T11:33:23Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $485 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-partido-liberal-pl-win-the-most-seats-in-the-next-brazilian-senate-election | elections | 0.7200 | P3_low_absolute_liquidity | liquidity $2299 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-carlos-roberto-massa-jnior-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 126 d, liq $28021) |
| 2026-05-30T11:33:23Z | will-aldo-rebelo-win-the-2026-brazilian-presidential-election | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 126 d, liq $1080097) |
| 2026-05-30T11:33:23Z | will-ronaldo-caiado-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 126 d, liq $28094) |
| 2026-05-30T11:33:23Z | will-the-miami-marlins-win-the-2026-world-series | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 154 d, liq $109520) |
| 2026-05-30T11:33:23Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 154 d, liq $30073) |
| 2026-05-30T11:33:23Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 154 d, liq $28084) |
| 2026-05-30T11:33:23Z | will-st-louis-cardinals-win-the-2026-national-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 154 d, liq $23604) |
| 2026-05-30T11:33:23Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 154 d, liq $27476) |
| 2026-05-30T11:33:23Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 154 d, liq $31299) |
| 2026-05-30T11:33:23Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 156 d, liq $7712) |
| 2026-05-30T11:33:23Z | will-jakob-glesnes-win-2026-mls-defender-of-the-year | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 165 d, liq $282) |
| 2026-05-30T11:33:23Z | will-david-brekalo-win-2026-mls-defender-of-the-year | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 165 d, liq $296) |
| 2026-05-30T11:33:23Z | will-richard-van-de-water-win-the-bachelorette-season-22 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 183 d, liq $3830) |
| 2026-05-30T11:33:23Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 192 d, liq $12258) |
| 2026-05-30T11:33:23Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 192 d, liq $10051) |
| 2026-05-30T11:33:23Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9295) |
| 2026-05-30T11:33:23Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $10112) |
| 2026-05-30T11:33:23Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 192 d, liq $12074) |
| 2026-05-30T11:33:23Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $8916) |
| 2026-05-30T11:33:23Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 202 d, liq $4780) |
| 2026-05-30T11:33:23Z | 5kt-meteor-strike-in-2026 | weather-natural | 0.3200 | P3_low_absolute_liquidity | liquidity $2992 < absolute min $5000 |
| 2026-05-30T11:33:23Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1224 < absolute min $5000 |
| 2026-05-30T11:33:23Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4221 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $199 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $69451) |
| 2026-05-30T11:33:23Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $832) |
| 2026-05-30T11:33:23Z | will-any-category-5-hurricane-make-landfall-in-the-us-in-before-2027 | weather-natural | 0.1300 | P3_low_absolute_liquidity | liquidity $1659 < absolute min $5000 |
| 2026-05-30T11:33:23Z | brex-ipo-before-2027 | crypto-launch | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $5003) |
| 2026-05-30T11:33:23Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $6246) |
| 2026-05-30T11:33:23Z | ledger-ipo-before-2027 | crypto-launch | 0.1400 | P3_low_absolute_liquidity | liquidity $1467 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $321) |
| 2026-05-30T11:33:23Z | major-meteor-strike-10kt-in-2026 | weather-natural | 0.1300 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T11:33:23Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T11:33:23Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4984 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $425) |
| 2026-05-30T11:33:23Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $12866) |
| 2026-05-30T11:33:23Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $259 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $10926) |
| 2026-05-30T11:33:23Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $690) |
| 2026-05-30T11:33:23Z | databricks-ipo-before-2027 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $2547 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4656 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $147) |
| 2026-05-30T11:33:23Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $226 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $7241) |
| 2026-05-30T11:33:23Z | will-there-be-exactly-1-confirmed-vei-4-or-higher-volcanic-eruption-worldwide-in-2026-937 | weather-natural | 0.2700 | P3_low_absolute_liquidity | liquidity $1407 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-a-team-from-lcp-asia-pacific-win-lol-worlds-2026 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $5927) |
| 2026-05-30T11:33:23Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 214 d, liq $9109) |
| 2026-05-30T11:33:23Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $641) |
| 2026-05-30T11:33:23Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3032 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 214 d, liq $14258) |
| 2026-05-30T11:33:23Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $6076) |
| 2026-05-30T11:33:23Z | will-there-be-between-14-and-16-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2454 < absolute min $5000 |
| 2026-05-30T11:33:23Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3368 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $195) |
| 2026-05-30T11:33:23Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $275) |
| 2026-05-30T11:33:23Z | will-there-be-between-11-and-13-earthquakes-of-magnitude-7pt0-or-higher-worldwide-in-2026 | other | 0.3600 | P3_low_absolute_liquidity | liquidity $1626 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $416) |
| 2026-05-30T11:33:23Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $573) |
| 2026-05-30T11:33:23Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3500 | P3_low_absolute_liquidity | liquidity $2496 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $13162) |
| 2026-05-30T11:33:23Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 214 d, liq $9620) |
| 2026-05-30T11:33:23Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1230 | P3_low_absolute_liquidity | liquidity $3176 < absolute min $5000 |
| 2026-05-30T11:33:23Z | jerome-powell-out-from-fed-board-by-may-30 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $55271) |
| 2026-05-30T11:33:23Z | will-2026-be-the-fifth-hottest-year-on-record | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $18681) |
| 2026-05-30T11:33:23Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $7923) |
| 2026-05-30T11:33:23Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $16780) |
| 2026-05-30T11:33:23Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $188670) |
| 2026-05-30T11:33:23Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 215 d, liq $3971) |
| 2026-05-30T11:33:23Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $659 < absolute min $5000 |
| 2026-05-30T11:33:23Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3298 < absolute min $5000 |
| 2026-05-30T11:33:23Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1300 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2458 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P3_low_absolute_liquidity | liquidity $768 < absolute min $5000 |
| 2026-05-30T11:33:23Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0910 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | altcoin-market-cap-dip-to-150b-before-2027 | other | 0.7200 | P3_low_absolute_liquidity | liquidity $740 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3200 | P3_low_absolute_liquidity | liquidity $548 < absolute min $5000 |
| 2026-05-30T11:33:23Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.3900 | P3_low_absolute_liquidity | liquidity $1145 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 215 d, liq $1451) |
| 2026-05-30T11:33:23Z | opensea-fdv-above-5b-one-day-after-launch-919-666-825-992-543-333 | crypto-launch | 0.0500 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $12642) |
| 2026-05-30T11:33:23Z | opensea-fdv-above-100m-one-day-after-launch-172-151-588-987 | crypto-launch | 0.6000 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 239 d, liq $45932) |
| 2026-05-30T11:33:23Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $52849) |
| 2026-05-30T11:33:23Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 239 d, liq $19256) |
| 2026-05-30T11:33:23Z | will-the-atlanta-falcons-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 305 d, liq $194388) |
| 2026-05-30T11:33:23Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $267087) |
| 2026-05-30T11:33:23Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $43594) |
| 2026-05-30T11:33:23Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8200 | P4_pre_event | pre-event slug + 579 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3985 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $46210) |
| 2026-05-30T11:33:23Z | predictfun-fdv-above-800m-one-day-after-launch | crypto-launch | 0.2400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2180 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-30T11:33:23Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1457424) |
| 2026-05-30T11:33:23Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 891 d, liq $1916329) |
| 2026-05-30T11:33:23Z | will-nikki-haley-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1061962) |
| 2026-05-30T11:33:23Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $936777) |
| 2026-05-30T11:33:23Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1075522) |
| 2026-05-30T11:33:23Z | will-phil-murphy-win-the-2028-democratic-presidential-nomination-611 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2333689) |
| 2026-05-30T11:33:23Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1170445) |
| 2026-05-30T11:33:23Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1626736) |
| 2026-05-30T11:33:23Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1151228) |
| 2026-05-30T11:33:23Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $475340) |
| 2026-05-30T11:33:23Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $946912) |
| 2026-05-30T11:33:23Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1311487) |
| 2026-05-30T11:33:23Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1447090) |
| 2026-05-30T11:33:23Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2365201) |
| 2026-05-30T11:33:23Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1479159) |
| 2026-05-30T11:33:23Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2438046) |
| 2026-05-30T11:33:23Z | will-pete-hegseth-win-the-2028-us-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1723572) |
| 2026-05-30T11:33:23Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $1208) |
| 2026-05-30T11:33:23Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $4991) |
| 2026-05-30T11:33:23Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $13103) |
| 2026-05-30T11:33:23Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2500 | P3_low_absolute_liquidity | liquidity $624 < absolute min $5000 |
| 2026-05-30T11:33:23Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $23046) |
| 2026-05-30T11:33:23Z | will-valve-remove-overpass-from-the-map-pool | other | 0.0400 | P3_low_absolute_liquidity | liquidity $2951 < absolute min $5000 |
## Losses pattern (últimos 100, append-only, rotación a tail)

| timestamp | slug | category | entry | exit | pnl | reason_thesis_failed |
|---|---|---|---|---|---|---|
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
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T03:25:01Z | T-2097884-1779678004140 | uncategorized | 0.3100 | 0.9700 | 112.00 | take_profit |  | 52.60 | short | 0.02 |
| 2026-05-25T05:15:02Z | T-2002531-1779652804112 | uncategorized | 0.0770 | 0.1110 | 22.90 | take_profit |  | 51.86 | medium | 0.39 |
| 2026-05-25T05:25:02Z | T-2313547-1779604204419 | uncategorized | 0.1640 | 0.2350 | 28.36 | take_profit |  | 65.51 | short | 0.95 |
| 2026-05-25T06:05:02Z | T-2321929-1779679804173 | uncategorized | 0.2200 | 0.3300 | 28.24 | take_profit |  | 56.47 | short | 0.11 |
| 2026-05-25T06:15:01Z | T-2334096-1779517803519 | uncategorized | 0.9200 | 0.9990 | 4.94 | market_closed |  | 57.50 | short | 1.99 |
| 2026-05-25T12:30:03Z | T-2002531-1779690604559 | uncategorized | 0.1180 | 0.2300 | 53.23 | take_profit |  | 56.08 | medium | 0.25 |
| 2026-05-25T16:35:13Z | T-2133405-1779487204257 | uncategorized | 0.8400 | 0.1010 | -54.32 | stop_loss |  | 61.74 | medium | 2.77 |
| 2026-05-25T16:35:29Z | T-2313550-1779571804012 | uncategorized | 0.1960 | 0.0280 | -53.13 | stop_loss |  | 61.99 | short | 1.80 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T16:35:29Z | T-1731345-1779643804205 | uncategorized | 0.2200 | 0.3300 | 24.30 | take_profit |  | 48.59 | long | 0.96 |
| 2026-05-25T16:35:29Z | T-2321926-1779665404197 | uncategorized | 0.7000 | 0.9990 | 21.65 | take_profit |  | 50.70 | short | 0.71 |
| 2026-05-25T16:35:29Z | T-2313548-1779674403929 | uncategorized | 0.4300 | 0.6300 | 25.39 | take_profit |  | 54.60 | short | 0.61 |
| 2026-05-25T16:35:38Z | T-2293518-1779687003842 | uncategorized | 0.2400 | 0.0010 | -56.29 | stop_loss |  | 56.52 | short | 0.46 |
| 2026-05-25T16:35:46Z | T-2321929-1779690604559 | uncategorized | 0.2400 | 0.0010 | -56.99 | stop_loss |  | 57.23 | short | 0.42 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-25T18:55:01Z | T-2111605-1779647403942 | uncategorized | 0.1650 | 0.2480 | 25.67 | take_profit |  | 51.03 | medium | 1.02 |
| 2026-05-25T19:00:02Z | T-2293512-1779660004159 | uncategorized | 0.9450 | 0.9990 | 2.94 | market_closed |  | 51.51 | short | 0.87 |
| 2026-05-25T19:50:02Z | T-2324252-1779642003777 | uncategorized | 0.1300 | 0.2700 | 64.40 | take_profit |  | 59.80 | short | 1.12 |
| 2026-05-25T20:00:02Z | T-2111564-1779643804205 | uncategorized | 0.2700 | 0.5300 | 50.73 | take_profit |  | 52.68 | medium | 1.10 |
| 2026-05-25T20:05:02Z | T-2230536-1779679804173 | uncategorized | 0.9000 | 0.9990 | 6.09 | market_closed |  | 55.34 | medium | 0.69 |
| 2026-05-25T20:20:01Z | T-2126515-1779647403942 | uncategorized | 0.0810 | 0.1200 | 24.08 | take_profit |  | 50.01 | medium | 1.08 |
| 2026-05-25T21:35:09Z | T-2313547-1779687003842 | uncategorized | 0.2360 | 0.0400 | -46.01 | stop_loss |  | 55.39 | short | 0.67 |
| 2026-05-25T23:00:26Z | T-2119073-1779654603617 | uncategorized | 0.1230 | 0.0230 | -42.79 | stop_loss |  | 52.64 | medium | 1.10 |
| 2026-05-25T23:05:08Z | T-2334097-1779570004545 | uncategorized | 0.9500 | 0.0800 | -57.23 | stop_loss |  | 62.49 | short | 2.09 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-26T00:10:01Z | T-2159726-1779597003852 | uncategorized | 0.1200 | 0.1800 | 33.26 | take_profit |  | 66.51 | medium | 1.82 |
| 2026-05-26T00:30:20Z | T-1808970-1779211803817 | uncategorized | 0.0530 | 0.0080 | -44.01 | stop_loss |  | 51.83 | medium | 6.29 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-26T01:15:02Z | T-2169995-1779730346526 | market | 0.1000 | 0.1400 | 22.45 | take_profit |  | 56.12 | long | 0.32 |
| 2026-05-26T01:50:02Z | T-2132799-1779726947128 | market | 0.2500 | 0.3700 | 29.21 | take_profit | claudemax-websearch | 60.85 | medium | 0.38 |
| 2026-05-26T04:05:02Z | T-2111561-1779643804205 | uncategorized | 0.0520 | 0.0750 | 22.84 | take_profit |  | 51.63 | medium | 1.44 |
| 2026-05-26T05:50:02Z | T-2303427-1779766204447 | market | 0.0700 | 0.1000 | 26.78 | take_profit |  | 62.48 | short | 0.10 |
| 2026-05-26T07:10:08Z | T-2313549-1779652804112 | uncategorized | 0.3400 | 0.0600 | -43.58 | stop_loss |  | 52.92 | short | 1.47 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-26T07:40:01Z | T-957019-1779643804205 | uncategorized | 0.3300 | 0.4800 | 22.54 | take_profit |  | 49.58 | long | 1.59 |
| 2026-05-26T10:50:02Z | T-2303427-1779775204890 | market | 0.1000 | 0.2260 | 81.28 | take_profit |  | 64.51 | short | 0.20 |
| 2026-05-26T12:20:26Z | T-2169995-1779759026423 | market | 0.1500 | 0.0200 | -54.79 | stop_loss |  | 63.22 | long | 0.45 |
| 2026-05-26T12:20:51Z | T-2303427-1779793283201 | market | 0.2900 | 0.0570 | -52.38 | stop_loss |  | 65.20 | short | 0.05 |
| 2026-05-26T15:30:08Z | T-2303427-1779798604693 | market | 0.0760 | 0.0080 | -57.55 | stop_loss |  | 64.32 | short | 0.12 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-27T08:25:01Z | T-562186-1779210004342 | uncategorized | 0.9250 | 0.9990 | 6.38 | market_closed |  | 79.73 | medium | 7.64 |
| 2026-05-27T13:50:23Z | T-2313280-1779814832575 | market | 0.6800 | 0.0600 | -70.91 | stop_loss | claudemax-websearch | 77.77 | short | 0.87 |
| 2026-05-27T14:55:09Z | T-2322404-1779712204481 | market | 0.1500 | 0.0300 | -45.72 | stop_loss |  | 57.15 | medium | 2.10 |
| 2026-05-28T05:25:18Z | T-2322008-1779899405064 | market | 0.8070 | 0.1600 | -60.30 | stop_loss |  | 75.21 | short | 0.54 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-28T07:48:17Z | T-648378-1779206389814 | uncategorized | 0.8800 | 0.8800 | 0.00 | manual_reset_v7 | legacy | 150.00 | long | 8.66 |
| 2026-05-28T07:48:17Z | T-839358-1779210004342 | uncategorized | 0.0545 | 0.0570 | 3.31 | manual_reset_v7 |  | 72.07 | long | 8.62 |
| 2026-05-28T07:48:17Z | T-566136-1779211803817 | uncategorized | 0.5900 | 0.5700 | -1.83 | manual_reset_v7 |  | 53.97 | medium | 8.60 |
| 2026-05-28T07:48:17Z | T-566140-1779233403338 | uncategorized | 0.4300 | 0.4200 | -1.26 | manual_reset_v7 |  | 54.25 | medium | 8.35 |
| 2026-05-28T07:48:17Z | T-2316216-1779642003777 | uncategorized | 0.4200 | 0.4100 | -1.37 | manual_reset_v7 |  | 57.43 | medium | 3.62 |
| 2026-05-28T07:48:17Z | T-2111640-1779643804205 | uncategorized | 0.2080 | 0.0830 | -30.41 | manual_reset_v7 |  | 50.60 | medium | 3.60 |
| 2026-05-28T07:48:17Z | T-569344-1779667204230 | uncategorized | 0.3200 | 0.2920 | -4.47 | manual_reset_v7 |  | 51.13 | medium | 3.33 |
| 2026-05-28T07:48:17Z | T-2155052-1779726947128 | geopolitics | 0.8400 | 0.7800 | -4.17 | manual_reset_v7 |  | 58.44 | long | 2.63 |
| 2026-05-28T07:48:17Z | T-2270338-1779732049801 | geopolitics | 0.7200 | 0.6000 | -9.17 | manual_reset_v7 | claudemax-websearch | 55.00 | long | 2.57 |
| 2026-05-28T07:48:17Z | T-2155023-1779733851679 | geopolitics | 0.7200 | 0.6000 | -8.98 | manual_reset_v7 | claudemax-websearch | 53.90 | long | 2.55 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-28T07:48:17Z | T-1809418-1779760804552 | other | 0.8770 | 0.8950 | 1.31 | manual_reset_v7 |  | 63.76 | long | 2.24 |
| 2026-05-28T07:48:17Z | T-824952-1779784234712 | market | 0.7300 | 0.8500 | 10.66 | manual_reset_v7 | claudemax-websearch | 64.85 | long | 1.97 |
| 2026-05-28T07:48:17Z | T-701547-1779791404097 | market | 0.1200 | 0.0900 | -15.89 | manual_reset_v7 |  | 63.55 | long | 1.89 |
| 2026-05-28T07:48:17Z | T-569366-1779809574675 | elections | 0.7200 | 0.6500 | -6.14 | manual_reset_v7 | claudemax-websearch | 63.17 | medium | 1.68 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-28T07:48:17Z | T-947269-1779870679596 | elections | 0.7100 | 0.7100 | 0.00 | manual_reset_v7 | claudemax-websearch | 77.94 | medium | 0.97 |
| 2026-05-28T07:48:17Z | T-1087513-1779894004820 | other | 0.7400 | 0.7300 | -1.04 | manual_reset_v7 | claudemax-websearch | 76.74 | medium | 0.70 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T19:45:01Z | T-2325745-1780068909718 | geopolitics | 0.9100 | 0.9500 | 8.12 | target_hit |  | 184.71 | short | 0.17 |
| 2026-05-29T20:35:41Z | T-678929-1780067007527 | elections | 0.2800 | 0.2700 | -8.09 | horizon-shortonly-2026-05-29 |  | 226.48 | medium | 0.23 |
| 2026-05-29T20:35:41Z | T-2354976-1780067007527 | geopolitics | 0.0900 | 0.0800 | -7.88 | horizon-shortonly-2026-05-29 |  | 70.93 | medium | 0.23 |
| 2026-05-29T20:35:41Z | T-564210-1780067007527 | sports-season | 0.5900 | 0.5700 | -2.62 | horizon-shortonly-2026-05-29 |  | 77.42 | medium | 0.23 |
| 2026-05-29T20:35:41Z | T-907355-1780067007527 | elections | 0.0200 | 0.0190 | -3.18 | horizon-shortonly-2026-05-29 |  | 63.59 | medium | 0.23 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T20:35:41Z | T-2183424-1780067007527 | other | 0.0500 | 0.0500 | 0.00 | horizon-shortonly-2026-05-29 |  | 32.30 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-1971905-1780067007527 | geopolitics | 0.3800 | 0.3400 | -21.36 | horizon-shortonly-2026-05-29 |  | 202.95 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-1494702-1780067007527 | market | 0.0300 | 0.0230 | -27.83 | horizon-shortonly-2026-05-29 |  | 119.26 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-1363963-1780067007527 | other | 0.0570 | 0.0560 | -0.75 | horizon-shortonly-2026-05-29 |  | 42.60 | long | 0.23 |
| 2026-05-29T20:35:41Z | T-597967-1780068909718 | executive-action | 0.1400 | 0.1300 | -4.37 | horizon-shortonly-2026-05-29 |  | 61.21 | long | 0.21 |
| 2026-05-29T20:35:41Z | T-569373-1780085069997 | elections | 0.0370 | 0.0350 | -2.05 | horizon-shortonly-2026-05-29 |  | 37.84 | medium | 0.02 |
| 2026-05-29T20:35:41Z | T-677314-1780085069997 | geopolitics | 0.0300 | 0.0230 | -8.06 | horizon-shortonly-2026-05-29 |  | 34.56 | long | 0.02 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T20:40:01Z | T-2340268-1780087006486 | market | 0.2900 | 0.3100 | 3.06 | no_remaining_edge |  | 44.37 | short | 0.00 |
| 2026-05-30T03:30:03Z | T-1999689-1780087006486 | other | 0.6390 | 0.0120 | -169.63 | market_closed |  | 172.88 | short | 0.29 |
| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-30T04:20:01Z | T-2365528-1780111994961 | market | 0.2700 | 0.3900 | 41.63 | target_hit |  | 93.68 | short | 0.03 |
| 2026-05-30T06:10:02Z | T-2139900-1780085069997 | other | 0.4400 | 0.0010 | -94.51 | market_closed |  | 94.73 | short | 0.42 |
| 2026-05-30T06:10:04Z | T-2364498-1780092236400 | geopolitics | 0.0290 | 0.0010 | -28.46 | market_closed |  | 29.47 | short | 0.34 |
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
<!-- auto-generated 2026-05-25T20:00:04Z -->
- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 15 losses, 1 wins (win rate 6%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-25T23:30:03Z -->
- En categoría `uncategorized` · horizonte `short` · precio `>0.70`: 19 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-26T01:02:03Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `<0.10`: 5 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-26T04:30:33Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `<0.10`: 5 losses, 1 wins (win rate 17%). **Veto soft o reducir size 50%.**
<!-- auto-generated 2026-05-26T07:31:30Z -->
- En categoría `uncategorized` · horizonte `medium` · precio `<0.10`: 6 losses, 0 wins (win rate 0%). **Veto soft o reducir size 50%.**


## Human notes
_(no se toca por automatización)_

## Sources stats (rolling 30d, computed by brain)

| domain | trades | wins | losses | win_rate | total_pnl | blacklisted |
|---|---|---|---|---|---|---|
| claudemax-websearch | 11 | 2 | 8 | 20.0% | -77.68 | YES |

## Anti-patterns identificados- same-day-price-target · intraday-decay · no-momentum-confirm · short-horizon-chalk · high-entry-low-exit — visto en  (2026-05-26, pnl $-57.55)

- ceasefire-rumor · short-horizon-geopolitics · market-closed-capitulation · low-liquidity-pump — visto en us-announces-new-iran-agreementceasefire-extension-by-may-29 (2026-05-30, pnl $-28.46)
- short-horizon-binary · no-sources-checked · market-closed-bet · low-conviction-chalk — visto en will-apple-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 (2026-05-30, pnl $-169.63)
- same-day-price-level · no-source-approval · overconfident-high-entry · crypto-short-horizon — visto en bitcoin-above-74k-on-may-28-2026 (2026-05-28, pnl $-60.30)
- extreme-price-target · short-horizon-commodity · no-catalyst-confirmed · low-probability-moonshot — visto en  (2026-05-27, pnl $-45.72)
- same-day-price-target · price-near-threshold · recency-bias-entry · high-vol-short-horizon · geopolitics-short-horizon — visto en bitcoin-above-76k-on-may-27-2026 (2026-05-27, pnl $-70.91)

- zero-research-entry · same-day-expiry · crypto-price-level · intraday-short-horizon · below-50-underdog — visto en  (2026-05-26, pnl $-52.38)
- no-sources-consulted · against-stated-hodl-policy · horizon-mismatch · high-entry-low-prob-event · stop-loss-under-hours — visto en  (2026-05-26, pnl $-54.79)
- exact-count-range-bet · celebrity-behavior-noise · short-horizon-volatility · uncategorized-low-edge · no-source-validation — visto en  (2026-05-26, pnl $-43.58)
- zero-sources-entry · geopolitics-short-horizon · treaty-deadline-bet · low-prob-long-shot · uncategorized-no-research — visto en  (2026-05-26, pnl $-44.01)
- ceasefire-fragile · geopolitics-short-horizon · overpriced-high-prob · low-margin-safety · stop-loss-cliff — visto en  (2026-05-25, pnl $-57.23)
- zero-research-entry · geopolitics-numeric-threshold · uncategorized-low-liquidity · low-prob-no-edge — visto en  (2026-05-25, pnl $-42.79)
- narrow-activity-range · social-behavior-bet · unpredictable-subject · low-probability-band · stop-loss-short-horizon — visto en  (2026-05-25, pnl $-46.01)
- tweet-count-range · behavior-prediction · no-sources-consulted · low-edge-bet · short-horizon-noise — visto en  (2026-05-25, pnl $-56.99)
- short-horizon-crypto · same-day-chalk-bet · no-source-validation · low-timeframe-bet · stop-loss-wipeout — visto en  (2026-05-25, pnl $-56.29)
- narrow-count-range · social-activity-bet · no-sources-entry · uncategorized-market · behavior-prediction — visto en  (2026-05-25, pnl $-53.13)
- geopolitical-action-bet · no-source-validation · high-entry-no-catalyst · short-horizon-collapse · stop-loss-whipsaw — visto en  (2026-05-25, pnl $-54.32)
_(autogenerada por exit_monitor tras cada pérdida — brain consume vía M3)_

## Brier score (semanal) — 2026-05-29

resolved=2 overall_brier=0.5000

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 2 | 0.5000 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 1 | 0.0000 |
| long | 1 | 1.0000 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| unknown | 2 | 0.5000 |

## Calibracion diaria — 2026-05-29

- Ventana: 7d · short thesis-closes n=26 (W26/L0)
- Win rate 100.0% · avg win $12.78 · avg loss $0.0 · payoff 0.0 (break-even wr 0.0%)
- Expectancy/trade $12.7792 · total $332.26
- Brier short: n=4 score=0.0925
- **kelly_shrink=1.0** · min_edge_points_override=ninguno
- short n=26 wr=1.000 expectancy=+12.7792/trade; shrink=1.0 (no losers in window, shrink_exp=1.0); min_edge 0.030->0.030

## Brier score (semanal) — 2026-05-30

resolved=12 overall_brier=0.2034

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 11 | 0.1944 |
| market | 1 | 0.3025 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 1 | 0.0000 |
| short | 10 | 0.1441 |
| long | 1 | 1.0000 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| info | 2 | 0.0137 |
| other | 1 | 0.1225 |
| none | 4 | 0.1671 |
| calibration | 3 | 0.2076 |
| unknown | 2 | 0.5000 |


## Calibracion diaria — 2026-05-30

- Ventana: 7d · short thesis-closes n=25 (W25/L0)
- Win rate 100.0% · avg win $12.51 · avg loss $0.0 · payoff 0.0 (break-even wr 0.0%)
- Expectancy/trade $12.5072 · total $312.68
- Brier short: n=33 score=0.1743
- **kelly_shrink=1.0** · min_edge_points_override=ninguno
- short n=25 wr=1.000 expectancy=+12.5072/trade; shrink=1.0 (no losers in window, shrink_exp=1.0 * brier_factor=1.000 (brier=0.174 n=33)); min_edge 0.030->0.030

## Brier score (semanal) — 2026-05-30

resolved=62 overall_brier=0.1608

### Por categoría

| group | n | brier |
|-------|---|-------|
| other | 38 | 0.1410 |
| market | 2 | 0.1531 |
| geopolitics | 22 | 0.1957 |

### Por horizonte

| group | n | brier |
|-------|---|-------|
| medium | 6 | 0.0008 |
| short | 55 | 0.1630 |
| long | 1 | 1.0000 |

### Por edge_type

| group | n | brier |
|-------|---|-------|
| calibration | 14 | 0.1028 |
| info | 20 | 0.1068 |
| other | 1 | 0.1225 |
| none | 22 | 0.1942 |
| unknown | 5 | 0.4000 |

