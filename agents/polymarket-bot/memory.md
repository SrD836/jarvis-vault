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
| 2026-05-30T03:03:05Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 214 d, liq $8871) |
| 2026-05-30T03:03:05Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $11127) |
| 2026-05-30T03:03:05Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $16496) |
| 2026-05-30T03:03:05Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $128) |
| 2026-05-30T03:03:05Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $10401) |
| 2026-05-30T03:03:05Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $326 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $626) |
| 2026-05-30T03:03:05Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $450) |
| 2026-05-30T03:03:05Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $683) |
| 2026-05-30T03:03:05Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4899 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $686) |
| 2026-05-30T03:03:05Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $8071) |
| 2026-05-30T03:03:05Z | openai-1t-ipo-before-2027 | crypto-launch | 0.6700 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $68923) |
| 2026-05-30T03:03:05Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4905 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $954) |
| 2026-05-30T03:03:05Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $157840) |
| 2026-05-30T03:03:05Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1197 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $176) |
| 2026-05-30T03:03:05Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $182 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $157 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-7-8-spacex-starship-launches-successfully-reach-space-in-2026 | other | 0.0580 | P3_low_absolute_liquidity | liquidity $3788 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-jensen-huang-be-richest-person-on-december-31-229 | other | 0.0300 | P3_low_absolute_liquidity | liquidity $4938 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $12994) |
| 2026-05-30T03:03:05Z | ledger-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $1459 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1220 | P3_low_absolute_liquidity | liquidity $3205 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3600 | P3_low_absolute_liquidity | liquidity $2705 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $145922) |
| 2026-05-30T03:03:05Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $501) |
| 2026-05-30T03:03:05Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $160601) |
| 2026-05-30T03:03:05Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 214 d, liq $14450) |
| 2026-05-30T03:03:05Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3079 < absolute min $5000 |
| 2026-05-30T03:03:05Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3276 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 214 d, liq $296) |
| 2026-05-30T03:03:05Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T03:03:05Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 215 d, liq $3004) |
| 2026-05-30T03:03:05Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $797 < absolute min $5000 |
| 2026-05-30T03:03:05Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1300 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1219 < absolute min $5000 |
| 2026-05-30T03:03:05Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3210 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-ethereum-dip-to-800-by-december-31-2026 | market | 0.2100 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.21) |
| 2026-05-30T03:03:05Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $659 < absolute min $5000 |
| 2026-05-30T03:03:05Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4000 | P3_low_absolute_liquidity | liquidity $1123 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.7000 | P3_low_absolute_liquidity | liquidity $2419 < absolute min $5000 |
| 2026-05-30T03:03:05Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 216 d, liq $1479) |
| 2026-05-30T03:03:05Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12376) |
| 2026-05-30T03:03:05Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.2620 | P3_low_absolute_liquidity | liquidity $2675 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-ethereum-reach-5500-by-december-31-2026 | market | 0.0700 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.21) |
| 2026-05-30T03:03:05Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P3_low_absolute_liquidity | liquidity $858 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-ethereum-reach-6500-by-december-31-2026 | market | 0.0500 | P11 | tradingview ETHUSDT sentiment=bullish (conf 0.65) contradicts bear thesis (price_yes=0.21) |
| 2026-05-30T03:03:05Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3200 | P3_low_absolute_liquidity | liquidity $577 < absolute min $5000 |
| 2026-05-30T03:03:05Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4500 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 239 d, liq $52906) |
| 2026-05-30T03:03:05Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 239 d, liq $46003) |
| 2026-05-30T03:03:05Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $52859) |
| 2026-05-30T03:03:05Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 239 d, liq $19105) |
| 2026-05-30T03:03:05Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 305 d, liq $117287) |
| 2026-05-30T03:03:05Z | will-the-cleveland-browns-win-the-2027-nfl-league-championship | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 305 d, liq $117273) |
| 2026-05-30T03:03:05Z | will-the-arizona-cardinals-win-the-2027-nfl-league-championship | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 305 d, liq $191834) |
| 2026-05-30T03:03:05Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $262582) |
| 2026-05-30T03:03:05Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $180937) |
| 2026-05-30T03:03:05Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $45454) |
| 2026-05-30T03:03:05Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3702 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $42795) |
| 2026-05-30T03:03:05Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $23513) |
| 2026-05-30T03:03:05Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8300 | P4_pre_event | pre-event slug + 579 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2230 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | predictfun-fdv-above-300m-one-day-after-launch | crypto-launch | 0.6700 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-30T03:03:05Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1463417) |
| 2026-05-30T03:03:05Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $478988) |
| 2026-05-30T03:03:05Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2444970) |
| 2026-05-30T03:03:05Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2373884) |
| 2026-05-30T03:03:05Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1480940) |
| 2026-05-30T03:03:05Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $878431) |
| 2026-05-30T03:03:05Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $937546) |
| 2026-05-30T03:03:05Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 891 d, liq $1916874) |
| 2026-05-30T03:03:05Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1154762) |
| 2026-05-30T03:03:05Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1629542) |
| 2026-05-30T03:03:05Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1452616) |
| 2026-05-30T03:03:05Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1081982) |
| 2026-05-30T03:03:05Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 891 d, liq $303072) |
| 2026-05-30T03:03:05Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1175058) |
| 2026-05-30T03:03:05Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $552 < absolute min $5000 |
| 2026-05-30T03:03:05Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22439) |
| 2026-05-30T03:03:05Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12758) |
| 2026-05-30T03:03:05Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-30T03:03:05Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $982) |
| 2026-05-30T03:30:03Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon -150 d, liq $45955) |
| 2026-05-30T03:30:03Z | israel-and-syria-normalize-relations-by-june-30-2026 | geopolitics | 0.0400 | P2 | mercado ya expiró (endDate=2025-12-31T00:00:00Z, hace 150 días) |
| 2026-05-30T03:30:03Z | ukraine-recognizes-russian-sovereignty-over-its-territory-by-june-30-2026 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -149 d, liq $13576) |
| 2026-05-30T03:30:03Z | new-stranger-things-episode-released-by-december-31-587-657-824-246 | other | 0.0700 | P2 | mercado ya expiró (endDate=2026-01-07T00:00:00Z, hace 143 días) |
| 2026-05-30T03:30:03Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 119 días) |
| 2026-05-30T03:30:03Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $776098) |
| 2026-05-30T03:30:03Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -96 d, liq $766685) |
| 2026-05-30T03:30:03Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -69 d, liq $34200) |
| 2026-05-30T03:30:03Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -67 d, liq $28840) |
| 2026-05-30T03:30:03Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -67 d, liq $29018) |
| 2026-05-30T03:30:03Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -67 d, liq $20251) |
| 2026-05-30T03:30:03Z | will-the-us-invade-venezuela-by-december-31-2026 | geopolitics | 0.1100 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T03:30:03Z | weed-rescheduled-by-december-31 | other | 0.3020 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 60 días) |
| 2026-05-30T03:30:03Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -48 d, liq $199582) |
| 2026-05-30T03:30:03Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $345358) |
| 2026-05-30T03:30:03Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.6800 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 15 días) |
| 2026-05-30T03:30:03Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon -15 d, liq $122538) |
| 2026-05-30T03:30:03Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -13 d, liq $5958) |
| 2026-05-30T03:30:03Z | will-gregg-kirkpatrick-win-the-2026-georgia-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -11 d, liq $10499) |
| 2026-05-30T03:30:03Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -11 d, liq $6726) |
| 2026-05-30T03:30:03Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -3 d, liq $1421149) |
| 2026-05-30T03:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-29 | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -1 d, liq $69502) |
| 2026-05-30T03:30:03Z | highest-temperature-in-london-on-may-29-2026-25c | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $260151) |
| 2026-05-30T03:30:03Z | fif-bih-mac-2026-05-29-bih | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $441752) |
| 2026-05-30T03:30:03Z | nor-rbk-bog-2026-05-29-bog | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $360241) |
| 2026-05-30T03:30:03Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $14805) |
| 2026-05-30T03:30:03Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.0900 | P9 | P9: geopolitics pump cluster (price 0.09, 0d) |
| 2026-05-30T03:30:03Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $15041) |
| 2026-05-30T03:30:03Z | will-bitcoin-reach-78k-on-may-29 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $16163) |
| 2026-05-30T03:30:03Z | bitcoin-above-70k-on-may-30-2026 | market | 0.9960 | P0_ceiling | price ceiling: 0.9960 > 0.980 |
| 2026-05-30T03:30:03Z | bitcoin-above-68k-on-may-30-2026 | market | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-30T03:30:06Z | bitcoin-above-72k-on-may-30-2026 | market | 0.9700 | E2 | edge 0.020 < mín 0.030 (p̂=0.950, implied=0.970) |
| 2026-05-30T03:30:06Z | bitcoin-above-76k-on-may-30-2026 | market | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 0 d, liq $36580) |
| 2026-05-30T03:30:06Z | elon-musk-of-tweets-may-28-may-30-0-39 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $62703) |
| 2026-05-30T03:30:08Z | ucl-psg-ars-2026-05-30-psg | other | 0.4000 | E2 | edge 0.020 < mín 0.030 (p̂=0.420, implied=0.400) |
| 2026-05-30T03:30:13Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 0 d, liq $589660) |
| 2026-05-30T03:30:13Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.2000 | P9 | P9: geopolitics pump cluster (price 0.20, 0d) |
| 2026-05-30T03:30:13Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $199860) |
| 2026-05-30T03:30:13Z | will-reza-pahlavi-enter-iran-by-may-31 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $187478) |
| 2026-05-30T03:30:15Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0240 | E2 | edge 0.026 < mín 0.030 (p̂=0.050, implied=0.024) |
| 2026-05-30T03:30:15Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0600 | P9 | P9: geopolitics pump cluster (price 0.06, 0d) |
| 2026-05-30T03:30:15Z | will-rajasthan-royals-win-the-2026-indian-premier-league | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $307553) |
| 2026-05-30T03:30:15Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $186647) |
| 2026-05-30T03:30:15Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 0 d, liq $138296) |
| 2026-05-30T03:30:15Z | will-spacex-ipo-by-may-31-2026-259 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $145968) |
| 2026-05-30T03:30:15Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.0720 | P9 | P9: geopolitics pump cluster (price 0.07, 0d) |
| 2026-05-30T03:30:17Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0250 | E2 | edge 0.020 < mín 0.030 (p̂=0.005, implied=0.025) |
| 2026-05-30T03:30:19Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.0620 | P9 | P9: geopolitics pump cluster (price 0.06, 0d) |
| 2026-05-30T03:30:19Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0630 | P9 | P9: geopolitics pump cluster (price 0.06, 0d) |
| 2026-05-30T03:30:19Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $842282) |
| 2026-05-30T03:30:19Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $4722443) |
| 2026-05-30T03:30:19Z | mojtaba-khamenei-leaves-iran-by-may-31-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $233086) |
| 2026-05-30T03:30:19Z | will-microsoft-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $161505) |
| 2026-05-30T03:30:19Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $166100) |
| 2026-05-30T03:30:22Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 0 d, liq $261136) |
| 2026-05-30T03:30:22Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 0 d, liq $58158) |
| 2026-05-30T03:30:24Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0280 | E2 | edge 0.016 < mín 0.030 (p̂=0.012, implied=0.028) |
| 2026-05-30T03:30:26Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.3400 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda desencadena... |
| 2026-05-30T03:30:26Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $220566) |
| 2026-05-30T03:30:26Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 0 d, liq $94537) |
| 2026-05-30T03:30:28Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.0360 | E2 | edge 0.016 < mín 0.030 (p̂=0.020, implied=0.036) |
| 2026-05-30T03:30:28Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $197406) |
| 2026-05-30T03:30:28Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 0 d, liq $50325) |
| 2026-05-30T03:30:30Z | trump-out-as-president-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $921695) |
| 2026-05-30T03:30:33Z | will-psg-win-the-202526-champions-league | sports-season | 0.5800 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. La pregunta se refiere a un torne... |
| 2026-05-30T03:30:33Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $97471) |
| 2026-05-30T03:30:33Z | will-alphabet-be-the-third-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 1 d, liq $12699) |
| 2026-05-30T03:30:33Z | will-nvidia-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $40148) |
| 2026-05-30T03:30:33Z | will-alphabet-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $401613) |
| 2026-05-30T03:30:33Z | will-alphabet-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-30T03:30:33Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $113304) |
| 2026-05-30T03:30:33Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $87119) |
| 2026-05-30T03:30:33Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $86183) |
| 2026-05-30T03:30:33Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $67477) |
| 2026-05-30T03:30:33Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $71103) |
| 2026-05-30T03:30:33Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $127432) |
| 2026-05-30T03:30:33Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $120389) |
| 2026-05-30T03:30:33Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $127668) |
| 2026-05-30T03:30:33Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $96091) |
| 2026-05-30T03:30:35Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $103168) |
| 2026-05-30T03:30:35Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $118007) |
| 2026-05-30T03:30:35Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $117702) |
| 2026-05-30T03:30:35Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $69274) |
| 2026-05-30T03:30:35Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $93110) |
| 2026-05-30T03:30:35Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 1 d, liq $325280) |
| 2026-05-30T03:30:35Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $985169) |
| 2026-05-30T03:30:35Z | will-wti-dip-to-40-in-may-2026-286-341-192-757-961-613-811-817-951-378-218-432-375 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $490195) |
| 2026-05-30T03:30:35Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $500141) |
| 2026-05-30T03:30:35Z | will-natural-gas-ng-hit-high-3pt40-in-may | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $288031) |
| 2026-05-30T03:30:35Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $465163) |
| 2026-05-30T03:30:35Z | will-wti-dip-to-70-in-may-2026-155-395-333-182-889-988-959-341-883-234-375-347-135 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $196103) |
| 2026-05-30T03:30:35Z | will-bitcoin-dip-to-50k-in-may-2026-896 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $135216) |
| 2026-05-30T03:30:35Z | will-ethereum-dip-to-1800-in-may-2026 | market | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 2 d, liq $31695) |
| 2026-05-30T03:30:35Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $172937) |
| 2026-05-30T03:30:35Z | will-ethereum-dip-to-1600-in-may-2026 | market | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 2 d, liq $40162) |
| 2026-05-30T03:30:35Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $180896) |
| 2026-05-30T03:30:35Z | will-bitcoin-reach-80000-in-may-2026-from-may-27 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $69012) |
| 2026-05-30T03:30:35Z | will-ethereum-reach-2600-in-may-2026 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $43395) |
| 2026-05-30T03:30:35Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $203798) |
| 2026-05-30T03:30:37Z | will-bitcoin-reach-150k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $1115312) |
| 2026-05-30T03:30:40Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $169589) |
| 2026-05-30T03:30:40Z | elon-musk-of-tweets-may-2026-2000plus | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $3929) |
| 2026-05-30T03:30:40Z | will-xrp-reach-1pt6-in-may-2026 | other | 0.0220 | P11 | tradingview XRPUSDT sentiment=bullish (conf 0.80) contradicts bear thesis (price_yes=0.02) |
| 2026-05-30T03:30:40Z | will-bitcoin-reach-77500-in-may-2026-from-may-27 | market | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 2 d, liq $65415) |
| 2026-05-30T03:30:40Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $29992) |
| 2026-05-30T03:30:40Z | will-ethereum-reach-3600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $62856) |
| 2026-05-30T03:30:40Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $10268) |
| 2026-05-30T03:30:43Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $38189) |
| 2026-05-30T03:30:43Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $18089) |
| 2026-05-30T03:30:45Z | will-spencer-pratt-win-the-2026-los-angeles-mayoral-election-983 | elections | 0.2100 | V3 | V3: V3 Trigger vago: el horizonte es lejano y Spencer Pratt es una figura mediática sin trayectoria política verifi... |
| 2026-05-30T03:30:45Z | elon-musk-of-tweets-may-26-june-2-320-339 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $34609) |
| 2026-05-30T03:30:45Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $52125) |
| 2026-05-30T03:30:45Z | elon-musk-of-tweets-may-26-june-2-100-119 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 3 d, liq $48707) |
| 2026-05-30T03:30:45Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $26319) |
| 2026-05-30T03:30:45Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $45571) |
| 2026-05-30T03:30:45Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $31920) |
| 2026-05-30T03:30:45Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $44712) |
| 2026-05-30T03:30:45Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $46096) |
| 2026-05-30T03:30:45Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $40102) |
| 2026-05-30T03:30:45Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $33146) |
| 2026-05-30T03:30:45Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $41067) |
| 2026-05-30T03:30:45Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $23958) |
| 2026-05-30T03:30:47Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $31065) |
| 2026-05-30T03:30:49Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $18949) |
| 2026-05-30T03:30:49Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $36219) |
| 2026-05-30T03:30:49Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $20480) |
| 2026-05-30T03:30:49Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $22476) |
| 2026-05-30T03:30:52Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $94839) |
| 2026-05-30T03:30:52Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $94951) |
| 2026-05-30T03:30:52Z | elon-musk-of-tweets-may-29-june-5-500plus | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 6 d, liq $16961) |
| 2026-05-30T03:30:52Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 6 d, liq $220836) |
| 2026-05-30T03:30:54Z | will-alexander-zverev-win-the-2026-mens-french-open | other | 0.3940 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días; la resolución está... |
| 2026-05-30T03:30:54Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $194354) |
| 2026-05-30T03:30:54Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $122729) |
| 2026-05-30T03:30:57Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $220048) |
| 2026-05-30T03:30:59Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0260 | V3 - Trigger vago: sin fecha concreta o sin evento verificable. | V3 - Trigger vago: sin fecha concreta o sin evento verificable.: V3 Trigger vago: sin fecha concreta o sin evento ver... |
| 2026-05-30T03:30:59Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $202581) |
| 2026-05-30T03:30:59Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $110947) |
| 2026-05-30T03:31:01Z | will-raphael-collignon-win-the-2026-mens-french-open | other | 0.0230 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda alterar sig... |
| 2026-05-30T03:31:04Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0630 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda modificar s... |
| 2026-05-30T03:31:04Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $190944) |
| 2026-05-30T03:31:04Z | will-karen-khachanov-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $382484) |
| 2026-05-30T03:31:04Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $313087) |
| 2026-05-30T03:31:04Z | will-novak-djokovic-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $665574) |
| 2026-05-30T03:31:06Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $152919) |
| 2026-05-30T03:31:06Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 7 d, liq $49285) |
| 2026-05-30T03:31:06Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $214686) |
| 2026-05-30T03:31:06Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 7 d, liq $45673) |
| 2026-05-30T03:31:08Z | will-jakub-mensik-win-the-2026-mens-french-open | other | 0.0230 | E2 | edge 0.013 < mín 0.030 (p̂=0.010, implied=0.023) |
| 2026-05-30T03:31:08Z | will-moise-kouame-win-the-2026-mens-french-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 7 d, liq $36668) |
| 2026-05-30T03:31:08Z | will-tommy-paul-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $370475) |
| 2026-05-30T03:31:11Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0330 | V6 Sin catalyst | V6 Sin catalyst: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. El French Open 2026 ... |
| 2026-05-30T03:31:11Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $182812) |
| 2026-05-30T03:31:14Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1080 | V3 Trigger vago + V6 Sin catalyst | V3 Trigger vago + V6 Sin catalyst: V3 Trigger vago: el evento 'Rafael Jodar gana el Abierto de Francia 2026' es extre... |
| 2026-05-30T03:31:14Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $240075) |
| 2026-05-30T03:31:16Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0240 | E2 | edge 0.014 < mín 0.030 (p̂=0.010, implied=0.024) |
| 2026-05-30T03:31:19Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0250 | E2 | edge 0.020 < mín 0.030 (p̂=0.005, implied=0.025) |
| 2026-05-30T03:31:22Z | will-joao-fonseca-win-the-2026-mens-french-open | other | 0.0840 | V3 | V3: V3 Trigger vago: no hay fecha concreta ni evento verificable que permita evaluar la probabilidad de que un tenist... |
| 2026-05-30T03:31:22Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $205068) |
| 2026-05-30T03:31:25Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $153148) |
| 2026-05-30T03:31:25Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $206167) |
| 2026-05-30T03:31:28Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.1700 | P9 | P9: geopolitics pump cluster (price 0.17, 7d) |
| 2026-05-30T03:31:28Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $213007) |
| 2026-05-30T03:31:28Z | will-martin-landaluce-win-the-2026-mens-french-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 7 d, liq $44650) |
| 2026-05-30T03:31:32Z | us-x-iran-permanent-peace-deal-by-june-15-2026-734-856-129 | geopolitics | 0.2300 | E2 | edge 0.020 < mín 0.030 (p̂=0.250, implied=0.230) |
| 2026-05-30T03:31:40Z | will-the-san-antonio-spurs-win-the-nba-western-conference-finals | sports-season | 0.4240 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días. | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.: V6 Sin catalyst: no hay evento discr... |
| 2026-05-30T03:31:40Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 17 d, liq $1010150) |
| 2026-05-30T03:31:40Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 17 d, liq $313326) |
| 2026-05-30T03:31:42Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 17 d, liq $2600091) |
| 2026-05-30T03:31:42Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 17 d, liq $800320) |
| 2026-05-30T03:31:42Z | will-lynn-vision-win-iem-cologne-major-2026 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 21 d, liq $18928) |
| 2026-05-30T03:31:42Z | will-sinners-win-iem-cologne-major-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $49738) |
| 2026-05-30T03:31:42Z | will-flyquest-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 21 d, liq $59199) |
| 2026-05-30T03:31:42Z | will-gaimin-gladiators-win-iem-cologne-major-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 21 d, liq $79667) |
| 2026-05-30T03:31:42Z | will-m80-win-iem-cologne-major-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 21 d, liq $37755) |
| 2026-05-30T03:31:42Z | will-nrg-win-iem-cologne-major-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 21 d, liq $70694) |
| 2026-05-30T03:31:43Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $184247) |
| 2026-05-30T03:31:43Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $156777) |
| 2026-05-30T03:31:43Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $199518) |
| 2026-05-30T03:31:43Z | will-juan-daniel-oviedo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $257209) |
| 2026-05-30T03:31:43Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $214599) |
| 2026-05-30T03:31:43Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $225894) |
| 2026-05-30T03:31:43Z | will-claudia-lpez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $152813) |
| 2026-05-30T03:31:43Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 22 d, liq $48920) |
| 2026-05-30T03:31:43Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $195151) |
| 2026-05-30T03:31:43Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $233688) |
| 2026-05-30T03:31:43Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $203051) |
| 2026-05-30T03:31:44Z | will-abelardo-de-la-espriella-win-the-2026-colombian-presidential-election | elections | 0.6300 | V5 Patrón débil | V5 Patrón débil: V5 Patrón débil: no se identifican al menos 3 fuentes independientes y fiables que respalden una... |
| 2026-05-30T03:31:47Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 22 d, liq $87711) |
| 2026-05-30T03:31:50Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 24 d, liq $1030) |
| 2026-05-30T03:31:50Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 24 d, liq $855) |
| 2026-05-30T03:31:50Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 24 d, liq $1457) |
| 2026-05-30T03:31:50Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 27 d, liq $1615) |
| 2026-05-30T03:31:50Z | will-moonshot-have-the-top-ai-model-at-the-end-of-june-2026 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $11144) |
| 2026-05-30T03:31:52Z | israel-and-indonesia-normalize-relations-by-june-30-2026 | geopolitics | 0.0470 | E2 | edge 0.027 < mín 0.030 (p̂=0.020, implied=0.047) |
| 2026-05-30T03:31:54Z | european-country-agrees-to-give-ukraine-security-guarantee-by-june-30 | geopolitics | 0.0400 | E2 | edge 0.020 < mín 0.030 (p̂=0.060, implied=0.040) |
| 2026-05-30T03:31:54Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 30 d, liq $106958) |
| 2026-05-30T03:31:56Z | will-the-montreal-canadiens-win-the-eastern-conference | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $113203) |
| 2026-05-30T03:31:56Z | gustavo-petro-out-as-leader-of-colombia-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 30 d, liq $15367) |
| 2026-05-30T03:32:00Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0230 | V3 | V3: V3 Trigger vago: 'regime fall' no tiene una definición verificable y unívoca (¿colapso total, cambio de líder... |
| 2026-05-30T03:32:02Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... | V3 Trigger vago: sin fecha concreta o sin evento verificable. V6 Sin catalyst: no hay evento discreto identificable e... |
| 2026-05-30T03:32:04Z | will-deutsche-bank-fail-by-june-30-2026 | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 30 d, liq $2562) |
| 2026-05-30T03:32:04Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 30 d, liq $11644) |
| 2026-05-30T03:32:07Z | will-russia-enter-kramatorsk-by-june-30-821-192 | geopolitics | 0.0280 | E2 | edge 0.007 < mín 0.030 (p̂=0.035, implied=0.028) |
| 2026-05-30T03:32:07Z | hantavirus-lab-leak-confirmed-by-june-30-1 | weather-natural | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 30 d, liq $105246) |
| 2026-05-30T03:32:10Z | mamdani-opens-city-owned-grocery-store-by-june-30 | other | 0.0210 | E2 | edge 0.029 < mín 0.030 (p̂=0.050, implied=0.021) |
| 2026-05-30T03:32:11Z | israel-x-iran-permanent-peace-deal-by-june-30-2026-262 | geopolitics | 0.1580 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: Pregunta vaga: 'permanent peace deal' no tiene una def... |
| 2026-05-30T03:32:11Z | hamad-bin-isa-al-khalifa-out-as-bahrain-king | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 30 d, liq $13739) |
| 2026-05-30T03:32:11Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $16631) |
| 2026-05-30T03:32:11Z | will-tim-walz-resign-by-june-30 | executive-action | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 30 d, liq $26597) |
| 2026-05-30T03:32:14Z | will-the-montreal-canadiens-win-the-2026-nhl-stanley-cup | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 30 d, liq $233936) |
| 2026-05-30T03:32:14Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 30 d, liq $369606) |
| 2026-05-30T03:32:17Z | will-russia-enter-kherson-by-june-30-645 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 30 d, liq $5447) |
| 2026-05-30T03:32:17Z | will-trump-and-putin-meet-next-in-belarus-572 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 30 d, liq $18302) |
| 2026-05-30T03:32:18Z | will-deepseek-have-the-top-ai-model-at-the-end-of-june-2026-514 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 30 d, liq $8723) |
| 2026-05-30T03:32:23Z | will-any-country-leave-nato-by-june-30-2026 | geopolitics | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $23648) |
| 2026-05-30T03:32:23Z | will-silver-si-hit-high-130-by-end-of-june | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $34624) |
| 2026-05-30T03:32:23Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $31184) |
| 2026-05-30T03:32:23Z | will-silver-si-hit-high-150-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $37337) |
| 2026-05-30T03:32:25Z | will-crude-oil-cl-hit-high-120-by-end-of-june | market | 0.1000 | E2 | edge 0.020 < mín 0.030 (p̂=0.120, implied=0.100) |
| 2026-05-30T03:32:27Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0370 | E2 | edge 0.017 < mín 0.030 (p̂=0.020, implied=0.037) |
| 2026-05-30T03:32:32Z | will-the-san-antonio-spurs-win-the-2026-nba-finals | sports-season | 0.2560 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: V3: Vago sin fecha concreta ni evento verificable; el ... |
| 2026-05-30T03:32:38Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0200 | P3_low_absolute_liquidity | liquidity $3103 < absolute min $5000 |
| 2026-05-30T03:32:38Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P3_low_absolute_liquidity | liquidity $2630 < absolute min $5000 |
| 2026-05-30T03:32:38Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0450 | P3_low_absolute_liquidity | liquidity $3633 < absolute min $5000 |
| 2026-05-30T03:32:38Z | will-clara-tauson-be-the-2026-womens-wimbledon-winner | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 42 d, liq $4060) |
| 2026-05-30T03:32:38Z | will-madison-keys-be-the-2026-womens-wimbledon-winner | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 42 d, liq $3027) |
| 2026-05-30T03:32:38Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 42 d, liq $2028) |
| 2026-05-30T03:32:38Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $539 < absolute min $5000 |
| 2026-05-30T03:32:38Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 50 d, liq $2577896) |
| 2026-05-30T03:32:38Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $7352874) |
| 2026-05-30T03:32:38Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 50 d, liq $4737816) |
| 2026-05-30T03:32:38Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 50 d, liq $2297570) |
| 2026-05-30T03:32:38Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $5542976) |
| 2026-05-30T03:32:38Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $3990771) |
| 2026-05-30T03:32:41Z | will-england-win-the-2026-fifa-world-cup-937 | sports-season | 0.1120 | E2 | edge 0.022 < mín 0.030 (p̂=0.090, implied=0.112) |
| 2026-05-30T03:32:41Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 50 d, liq $9235872) |
| 2026-05-30T03:32:44Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.012 < mín 0.030 (p̂=0.040, implied=0.052) |
| 2026-05-30T03:32:44Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 50 d, liq $5781386) |
| 2026-05-30T03:32:44Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 50 d, liq $3016289) |
| 2026-05-30T03:32:44Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $4572308) |
| 2026-05-30T03:32:44Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 50 d, liq $1165517) |
| 2026-05-30T03:32:46Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0290 | E2 | edge 0.024 < mín 0.030 (p̂=0.005, implied=0.029) |
| 2026-05-30T03:32:48Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | E2 | edge 0.006 < mín 0.030 (p̂=0.100, implied=0.094) |
| 2026-05-30T03:32:48Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $9656055) |
| 2026-05-30T03:32:48Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8224410) |
| 2026-05-30T03:32:48Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $9480809) |
| 2026-05-30T03:32:48Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 50 d, liq $497129) |
| 2026-05-30T03:32:48Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 50 d, liq $3467468) |
| 2026-05-30T03:32:48Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $6999653) |
| 2026-05-30T03:32:48Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $6297992) |
| 2026-05-30T03:32:51Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-30T03:32:51Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $9646561) |
| 2026-05-30T03:32:51Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $10676473) |
| 2026-05-30T03:32:51Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $7296915) |
| 2026-05-30T03:32:56Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8073704) |
| 2026-05-30T03:32:56Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 50 d, liq $5219227) |
| 2026-05-30T03:32:56Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $2465611) |
| 2026-05-30T03:32:56Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $8672642) |
| 2026-05-30T03:32:56Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $5284480) |
| 2026-05-30T03:32:56Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 50 d, liq $9784032) |
| 2026-05-30T03:32:56Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 50 d, liq $2588794) |
| 2026-05-30T03:32:56Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 50 d, liq $8250089) |
| 2026-05-30T03:32:56Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 50 d, liq $7739813) |
| 2026-05-30T03:32:59Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0390 | E2 | edge 0.024 < mín 0.030 (p̂=0.015, implied=0.039) |
| 2026-05-30T03:33:02Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 50 d, liq $10504776) |
| 2026-05-30T03:33:05Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $17012) |
| 2026-05-30T03:33:08Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 61 d, liq $35583) |
| 2026-05-30T03:33:08Z | will-tyreek-hill-play-for-philadelphia-eagles-next | other | 0.1480 | P3_low_absolute_liquidity | liquidity $206 < absolute min $5000 |
| 2026-05-30T03:33:08Z | will-filipe-luis-be-the-next-manager-of-ssc-napoli | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 93 d, liq $1879) |
| 2026-05-30T03:33:08Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0240 | P3_low_absolute_liquidity | liquidity $999 < absolute min $5000 |
| 2026-05-30T03:33:11Z | tush-push-banned-for-2026-nfl-season | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 102 d, liq $1998) |
| 2026-05-30T03:33:11Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $1954) |
| 2026-05-30T03:33:11Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 105 d, liq $3259) |
| 2026-05-30T03:33:14Z | will-the-moderate-party-m-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0360 | E2 | edge 0.016 < mín 0.030 (p̂=0.020, implied=0.036) |
| 2026-05-30T03:33:14Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 105 d, liq $2240) |
| 2026-05-30T03:33:14Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 105 d, liq $2798) |
| 2026-05-30T03:33:14Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 105 d, liq $2577) |
| 2026-05-30T03:33:14Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $598 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $1014 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 120 d, liq $1097) |
| 2026-05-30T03:33:14Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 120 d, liq $1496) |
| 2026-05-30T03:33:14Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5700 | P8 | P8: election 122d out, price 0.57 en banda ruido [0.30, 0.70] |
| 2026-05-30T03:33:14Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $462 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-carlos-roberto-massa-jnior-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 126 d, liq $27838) |
| 2026-05-30T03:33:14Z | will-ronaldo-caiado-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 126 d, liq $28019) |
| 2026-05-30T03:33:14Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 154 d, liq $31021) |
| 2026-05-30T03:33:14Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 154 d, liq $32024) |
| 2026-05-30T03:33:14Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 154 d, liq $28556) |
| 2026-05-30T03:33:14Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 154 d, liq $28008) |
| 2026-05-30T03:33:14Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 156 d, liq $7330) |
| 2026-05-30T03:33:14Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 156 d, liq $155788) |
| 2026-05-30T03:33:14Z | will-david-brekalo-win-2026-mls-defender-of-the-year | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 165 d, liq $182) |
| 2026-05-30T03:33:14Z | will-corey-seager-win-the-2026-american-league-hank-aaron-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 166 d, liq $3076) |
| 2026-05-30T03:33:14Z | will-sergio-prez-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 189 d, liq $751665) |
| 2026-05-30T03:33:14Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 192 d, liq $12460) |
| 2026-05-30T03:33:14Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9971) |
| 2026-05-30T03:33:14Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9014) |
| 2026-05-30T03:33:14Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 192 d, liq $12172) |
| 2026-05-30T03:33:14Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 192 d, liq $9953) |
| 2026-05-30T03:33:14Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 192 d, liq $9277) |
| 2026-05-30T03:33:14Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 202 d, liq $5948) |
| 2026-05-30T03:33:14Z | will-brice-matthews-win-the-2026-al-rookie-of-the-year-award | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 202 d, liq $4142) |
| 2026-05-30T03:33:14Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $960) |
| 2026-05-30T03:33:14Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $6165) |
| 2026-05-30T03:33:14Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1170 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $613) |
| 2026-05-30T03:33:14Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $618) |
| 2026-05-30T03:33:14Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $301) |
| 2026-05-30T03:33:14Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $290 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $12531) |
| 2026-05-30T03:33:14Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4645 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $7270) |
| 2026-05-30T03:33:14Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 214 d, liq $7518) |
| 2026-05-30T03:33:14Z | openai-1t-ipo-before-2027 | crypto-launch | 0.6700 | P4_pre_event | pre-event slug + 214 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 214 d, liq $11196) |
| 2026-05-30T03:33:14Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 214 d, liq $126) |
| 2026-05-30T03:33:14Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $302 < absolute min $5000 |
| 2026-05-30T03:33:14Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3376 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3600 | P3_low_absolute_liquidity | liquidity $2707 < absolute min $5000 |
| 2026-05-30T03:33:14Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-30T03:33:14Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $662) |
| 2026-05-30T03:33:14Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $390) |
| 2026-05-30T03:33:14Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $8077) |
| 2026-05-30T03:33:14Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 214 d, liq $670) |
| 2026-05-30T03:33:14Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $158016) |
| 2026-05-30T03:33:14Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $146719) |
| 2026-05-30T03:33:14Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $685) |
| 2026-05-30T03:33:14Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $160856) |
| 2026-05-30T03:33:14Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4815 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 214 d, liq $14560) |
| 2026-05-30T03:33:14Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4587 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 214 d, liq $413) |
| 2026-05-30T03:33:14Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $4167 < absolute min $5000 |
| 2026-05-30T03:33:14Z | anduril-ipo-before-2027 | crypto-launch | 0.1100 | P3_low_absolute_liquidity | liquidity $3158 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $6345) |
| 2026-05-30T03:33:14Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 214 d, liq $104137) |
| 2026-05-30T03:33:14Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $179) |
| 2026-05-30T03:33:14Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $3001 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 214 d, liq $493) |
| 2026-05-30T03:33:14Z | will-7-8-spacex-starship-launches-successfully-reach-space-in-2026 | other | 0.0580 | P3_low_absolute_liquidity | liquidity $3785 < absolute min $5000 |
| 2026-05-30T03:33:14Z | ledger-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $1481 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $16499) |
| 2026-05-30T03:33:14Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 214 d, liq $9857) |
| 2026-05-30T03:33:14Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 214 d, liq $12537) |
| 2026-05-30T03:33:14Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1350 | P3_low_absolute_liquidity | liquidity $3219 < absolute min $5000 |
| 2026-05-30T03:33:14Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3721 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 214 d, liq $70745) |
| 2026-05-30T03:33:14Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $184 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 214 d, liq $9042) |
| 2026-05-30T03:33:14Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 214 d, liq $297) |
| 2026-05-30T03:33:14Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 215 d, liq $3002) |
| 2026-05-30T03:33:14Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $801 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1208 < absolute min $5000 |
| 2026-05-30T03:33:14Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3256 < absolute min $5000 |
| 2026-05-30T03:33:14Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4200 | P3_low_absolute_liquidity | liquidity $1048 < absolute min $5000 |
| 2026-05-30T03:33:14Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4500 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | standx-fdv-above-800m-one-day-after-launch-732-932-813-941-568 | crypto-launch | 0.1900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3200 | P3_low_absolute_liquidity | liquidity $582 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-tempo-launch-a-token-by-september-30-2026 | crypto-launch | 0.0940 | P3_low_absolute_liquidity | liquidity $837 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $650 < absolute min $5000 |
| 2026-05-30T03:33:14Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2437 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-tempo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 216 d, liq $1472) |
| 2026-05-30T03:33:14Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1300 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12379) |
| 2026-05-30T03:33:14Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 239 d, liq $52859) |
| 2026-05-30T03:33:14Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 239 d, liq $19113) |
| 2026-05-30T03:33:14Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 239 d, liq $53465) |
| 2026-05-30T03:33:14Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 239 d, liq $46024) |
| 2026-05-30T03:33:14Z | will-the-arizona-cardinals-win-the-2027-nfl-league-championship | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 305 d, liq $190912) |
| 2026-05-30T03:33:14Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 305 d, liq $117477) |
| 2026-05-30T03:33:14Z | will-the-cleveland-browns-win-the-2027-nfl-league-championship | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 305 d, liq $118392) |
| 2026-05-30T03:33:14Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $181274) |
| 2026-05-30T03:33:14Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 334 d, liq $263738) |
| 2026-05-30T03:33:14Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $23369) |
| 2026-05-30T03:33:14Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8300 | P4_pre_event | pre-event slug + 579 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7200 | P3_low_absolute_liquidity | liquidity $3629 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $45457) |
| 2026-05-30T03:33:14Z | will-spacexs-public-ticker-be-star-796 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 579 d, liq $42733) |
| 2026-05-30T03:33:14Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | predictfun-fdv-above-300m-one-day-after-launch | crypto-launch | 0.6700 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | predictfun-fdv-above-1b-one-day-after-launch | crypto-launch | 0.2210 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-30T03:33:14Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 891 d, liq $937969) |
| 2026-05-30T03:33:14Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 891 d, liq $305276) |
| 2026-05-30T03:33:14Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1082071) |
| 2026-05-30T03:33:14Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $878633) |
| 2026-05-30T03:33:14Z | will-tulsi-gabbard-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1308257) |
| 2026-05-30T03:33:14Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $2374339) |
| 2026-05-30T03:33:14Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $2444839) |
| 2026-05-30T03:33:14Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1479291) |
| 2026-05-30T03:33:14Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 891 d, liq $1625166) |
| 2026-05-30T03:33:14Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1453142) |
| 2026-05-30T03:33:14Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1461590) |
| 2026-05-30T03:33:14Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 891 d, liq $1154736) |
| 2026-05-30T03:33:14Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $479392) |
| 2026-05-30T03:33:14Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 891 d, liq $1917065) |
| 2026-05-30T03:33:14Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 891 d, liq $1174307) |
| 2026-05-30T03:33:14Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12758) |
| 2026-05-30T03:33:14Z | will-pepe-be-included-in-portugals-official-2026-world-cup-squad-list | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $4980) |
| 2026-05-30T03:33:14Z | reya-fdv-above-1b-one-day-after-launch-348-347 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon -1 d, liq $982) |
| 2026-05-30T03:33:14Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $23169) |
| 2026-05-30T03:33:14Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $538 < absolute min $5000 |
| 2026-05-30T03:33:14Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
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
