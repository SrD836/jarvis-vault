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
| 2026-05-29T23:03:15Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $261246) |
| 2026-05-29T23:03:15Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T23:03:15Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7500 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T23:03:15Z | spacex-ipo-closing-market-cap-above-1pt8t-631-261 | crypto-launch | 0.9000 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T23:03:15Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45562) |
| 2026-05-29T23:03:15Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $23561) |
| 2026-05-29T23:03:15Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7700 | P3_low_absolute_liquidity | liquidity $4477 < absolute min $5000 |
| 2026-05-29T23:03:15Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T23:03:15Z | predictfun-fdv-above-300m-one-day-after-launch | crypto-launch | 0.6700 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T23:03:15Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $477367) |
| 2026-05-29T23:03:15Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1464022) |
| 2026-05-29T23:03:15Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1154139) |
| 2026-05-29T23:03:15Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1080172) |
| 2026-05-29T23:03:15Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1916308) |
| 2026-05-29T23:03:15Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1446488) |
| 2026-05-29T23:03:15Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1176736) |
| 2026-05-29T23:03:15Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $2445262) |
| 2026-05-29T23:03:15Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $300730) |
| 2026-05-29T23:03:15Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1476734) |
| 2026-05-29T23:03:15Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $548 < absolute min $5000 |
| 2026-05-29T23:03:15Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T23:03:15Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12751) |
| 2026-05-29T23:03:15Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $22953) |
| 2026-05-29T23:08:31Z | israel-and-syria-normalize-relations-by-june-30-2026 | geopolitics | 0.0400 | P2 | mercado ya expiró (endDate=2025-12-31T00:00:00Z, hace 149 días) |
| 2026-05-29T23:08:31Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon -149 d, liq $47062) |
| 2026-05-29T23:08:31Z | new-stranger-things-episode-released-by-december-31-587-657-824-246 | other | 0.0700 | P2 | mercado ya expiró (endDate=2026-01-07T00:00:00Z, hace 142 días) |
| 2026-05-29T23:08:31Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 118 días) |
| 2026-05-29T23:08:31Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $776087) |
| 2026-05-29T23:08:31Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $766688) |
| 2026-05-29T23:08:31Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -68 d, liq $33931) |
| 2026-05-29T23:08:31Z | will-lars-boje-mathiesen-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon -66 d, liq $24053) |
| 2026-05-29T23:08:31Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -66 d, liq $20025) |
| 2026-05-29T23:08:31Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -66 d, liq $27646) |
| 2026-05-29T23:08:31Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -66 d, liq $28883) |
| 2026-05-29T23:08:31Z | will-the-us-invade-venezuela-by-december-31-2026 | geopolitics | 0.1100 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T23:08:31Z | weed-rescheduled-by-december-31 | other | 0.2890 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T23:08:31Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $199459) |
| 2026-05-29T23:08:31Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -46 d, liq $344810) |
| 2026-05-29T23:08:31Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0500 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T23:08:31Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.7100 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T23:08:31Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -12 d, liq $5938) |
| 2026-05-29T23:08:31Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -10 d, liq $6748) |
| 2026-05-29T23:08:31Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -2 d, liq $1066489) |
| 2026-05-29T23:09:20Z | us-announces-new-iran-agreementceasefire-extension-by-may-29 | geopolitics | 0.0380 | N1 | noticias contradicen tesis: Acuerdo tentativo alcanzado pero Trump no dio aprobación final el 29-may; sin anuncio of... |
| 2026-05-29T23:09:20Z | highest-temperature-in-shanghai-on-may-29-2026-26c | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $331556) |
| 2026-05-29T23:09:20Z | elon-musk-of-tweets-may-22-may-29-240-259 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $381000) |
| 2026-05-29T23:09:20Z | fif-bih-mac-2026-05-29-bih | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $475486) |
| 2026-05-29T23:09:20Z | elon-musk-of-tweets-may-22-may-29-200-219 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $437717) |
| 2026-05-29T23:09:20Z | nor-rbk-bog-2026-05-29-bog | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $438816) |
| 2026-05-29T23:09:20Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.1400 | P9 | P9: geopolitics pump cluster (price 0.14, 0d) |
| 2026-05-29T23:09:20Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17526) |
| 2026-05-29T23:09:20Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17295) |
| 2026-05-29T23:09:20Z | will-bitcoin-reach-78k-on-may-29 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $8966) |
| 2026-05-29T23:10:01Z | bitcoin-above-74k-on-may-30-2026 | market | 0.1800 | N1 | noticias contradicen tesis: BTC cotiza ~$73,100–73,750 el 29-may, por debajo de $74k y con apertura bajista; necesi... |
| 2026-05-29T23:11:11Z | ucl-psg-ars-2026-05-30-psg | other | 0.4000 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T23:11:11Z | bitcoin-above-76k-on-may-30-2026 | market | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 0 d, liq $33919) |
| 2026-05-29T23:12:02Z | ucl-psg-ars-2026-05-30-ars | other | 0.3100 | N1 | noticias contradicen tesis: Ben White baja oficial y Timber dudoso para la final UCL vs PSG en Budapest, debilitando ... |
| 2026-05-29T23:12:02Z | bitcoin-above-70k-on-may-30-2026 | market | 0.9970 | P0_ceiling | price ceiling: 0.9970 > 0.980 |
| 2026-05-29T23:12:02Z | elon-musk-of-tweets-may-28-may-30-0-39 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $25470) |
| 2026-05-29T23:13:33Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0430 | E2 | edge 0.023 < mín 0.030 (p̂=0.020, implied=0.043) |
| 2026-05-29T23:13:33Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $221653) |
| 2026-05-29T23:13:33Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 1 d, liq $62863) |
| 2026-05-29T23:14:06Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0410 | N1 | noticias contradicen tesis: Israel excluida de negociaciones EE.UU.-Irán; no existe proceso bilateral Israel-Irán; ... |
| 2026-05-29T23:14:49Z | will-trump-agree-to-iranian-oil-sanction-relief-by-may-31 | executive-action | 0.1300 | N1 | noticias contradicen tesis: Trump declaró que Irán no recibirá alivio de sanciones; EE.UU. impuso nuevas sanciones... |
| 2026-05-29T23:14:49Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 1 d, liq $50007) |
| 2026-05-29T23:14:49Z | trump-out-as-president-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $931431) |
| 2026-05-29T23:14:49Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0700 | P9 | P9: geopolitics pump cluster (price 0.07, 1d) |
| 2026-05-29T23:14:49Z | will-spacex-ipo-by-may-31-2026-259 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $138065) |
| 2026-05-29T23:14:49Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 1 d, liq $530881) |
| 2026-05-29T23:14:49Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 1 d, liq $67321) |
| 2026-05-29T23:14:49Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $200947) |
| 2026-05-29T23:14:49Z | will-russia-capture-all-of-hryshyne-by-may-31 | geopolitics | 0.9910 | P0_ceiling | price ceiling: 0.9910 > 0.980 |
| 2026-05-29T23:14:49Z | israel-x-lebanon-diplomatic-meeting-by-may-31-177 | geopolitics | 0.9980 | P0_ceiling | price ceiling: 0.9980 > 0.980 |
| 2026-05-29T23:14:49Z | uefa-champions-league-unbeaten-team | sports-season | 0.9980 | P0_ceiling | price ceiling: 0.9980 > 0.980 |
| 2026-05-29T23:15:34Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.1050 | P9 | P9: geopolitics pump cluster (price 0.10, 1d) |
| 2026-05-29T23:16:45Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0350 | E2 | edge 0.025 < mín 0.030 (p̂=0.010, implied=0.035) |
| 2026-05-29T23:16:45Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 1 d, liq $228947) |
| 2026-05-29T23:16:45Z | will-reza-pahlavi-enter-iran-by-may-31 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $182104) |
| 2026-05-29T23:17:20Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0290 | N1 | noticias contradicen tesis: Sin acuerdo permanente; hay tregua temporal extendida 45 días desde 15 mayo. Negociacion... |
| 2026-05-29T23:17:59Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.3700 | N1 | noticias contradicen tesis: Solo hay borrador de MOU ceasefire 60 días pendiente de aprobación de Trump; ningún ac... |
| 2026-05-29T23:17:59Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $4754233) |
| 2026-05-29T23:17:59Z | mojtaba-khamenei-leaves-iran-by-may-31-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $241202) |
| 2026-05-29T23:17:59Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $409154) |
| 2026-05-29T23:17:59Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $187652) |
| 2026-05-29T23:17:59Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.0850 | P9 | P9: geopolitics pump cluster (price 0.09, 1d) |
| 2026-05-29T23:17:59Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $167095) |
| 2026-05-29T23:17:59Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.0770 | P9 | P9: geopolitics pump cluster (price 0.08, 1d) |
| 2026-05-29T23:30:04Z | israel-and-syria-normalize-relations-by-june-30-2026 | geopolitics | 0.0400 | P2 | mercado ya expiró (endDate=2025-12-31T00:00:00Z, hace 149 días) |
| 2026-05-29T23:30:04Z | will-the-us-officially-declare-war-on-venezuela-by-june-30-2026 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon -149 d, liq $46910) |
| 2026-05-29T23:30:04Z | new-stranger-things-episode-released-by-december-31-587-657-824-246 | other | 0.0700 | P2 | mercado ya expiró (endDate=2026-01-07T00:00:00Z, hace 142 días) |
| 2026-05-29T23:30:04Z | another-us-strike-on-venezuela-by-december-31 | other | 0.1400 | P2 | mercado ya expiró (endDate=2026-01-31T00:00:00Z, hace 118 días) |
| 2026-05-29T23:30:04Z | will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $776003) |
| 2026-05-29T23:30:04Z | will-how-to-make-a-killing-score-at-least-60-on-the-rotten-tomatoes-tomatometer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -95 d, liq $766685) |
| 2026-05-29T23:30:04Z | will-zoran-stevanovi-be-the-next-prime-minister-of-slovenia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -68 d, liq $34344) |
| 2026-05-29T23:30:04Z | will-mona-juul-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon -66 d, liq $28065) |
| 2026-05-29T23:30:04Z | will-morten-messerschmidt-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon -66 d, liq $29096) |
| 2026-05-29T23:30:04Z | will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -66 d, liq $20025) |
| 2026-05-29T23:30:04Z | weed-rescheduled-by-december-31 | other | 0.2990 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T23:30:04Z | will-the-us-invade-venezuela-by-december-31-2026 | geopolitics | 0.1100 | P2 | mercado ya expiró (endDate=2026-03-31T00:00:00Z, hace 59 días) |
| 2026-05-29T23:30:04Z | will-carlos-lvarez-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -47 d, liq $199664) |
| 2026-05-29T23:30:04Z | will-sung-jae-im-win-the-2026-masters-tournament | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -46 d, liq $344906) |
| 2026-05-29T23:30:04Z | us-x-iran-diplomatic-meeting-by-june-30-2026-983-259-948 | geopolitics | 0.7100 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T23:30:04Z | us-x-iran-diplomatic-meeting-by-may-31-2026-283-616 | geopolitics | 0.0510 | P2 | mercado ya expiró (endDate=2026-05-15T00:00:00Z, hace 14 días) |
| 2026-05-29T23:30:04Z | will-antonio-mallo-be-the-next-president-of-andalusia-following-the-regional-election | elections | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon -12 d, liq $5944) |
| 2026-05-29T23:30:04Z | will-earl-carter-be-the-republican-nominee-for-senate-in-georgia | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon -10 d, liq $6190) |
| 2026-05-29T23:30:04Z | iran-closes-its-airspace-by-may-29 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon -2 d, liq $1063770) |
| 2026-05-29T23:30:05Z | us-announces-new-iran-agreementceasefire-extension-by-may-29 | geopolitics | 0.0350 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda mover la probabilidad de... |
| 2026-05-29T23:30:05Z | highest-temperature-in-london-on-may-29-2026-25c | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $11790) |
| 2026-05-29T23:30:05Z | elon-musk-of-tweets-may-22-may-29-200-219 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $431309) |
| 2026-05-29T23:30:05Z | fif-bih-mac-2026-05-29-bih | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $467289) |
| 2026-05-29T23:30:05Z | nor-rbk-bog-2026-05-29-bog | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $431686) |
| 2026-05-29T23:30:05Z | will-bradley-barcola-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17525) |
| 2026-05-29T23:30:05Z | will-desire-doue-be-the-20252026-top-ucl-goal-scorer | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $17293) |
| 2026-05-29T23:30:05Z | us-announces-new-iran-agreementceasefire-extension-by-may-30 | geopolitics | 0.1500 | P9 | P9: geopolitics pump cluster (price 0.15, 0d) |
| 2026-05-29T23:30:05Z | will-bitcoin-reach-78k-on-may-29 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 0 d, liq $8958) |
| 2026-05-29T23:30:05Z | bitcoin-above-76k-on-may-30-2026 | market | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 0 d, liq $30449) |
| 2026-05-29T23:30:08Z | bitcoin-above-70k-on-may-30-2026 | market | 0.9950 | P0_ceiling | price ceiling: 0.9950 > 0.980 |
| 2026-05-29T23:30:09Z | bitcoin-above-72k-on-may-30-2026 | market | 0.9680 | P11 | tradingview BTCUSDT sentiment=bearish (conf 0.65) contradicts bull thesis (price_yes=0.97) |
| 2026-05-29T23:30:11Z | bitcoin-above-74k-on-may-30-2026 | market | 0.1600 | P11 | tradingview BTCUSDT sentiment=bearish (conf 0.65) contradicts bull thesis (price_yes=0.97) |
| 2026-05-29T23:30:11Z | elon-musk-of-tweets-may-28-may-30-0-39 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 0 d, liq $70033) |
| 2026-05-29T23:30:11Z | will-the-iranian-regime-fall-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $4723716) |
| 2026-05-29T23:30:11Z | iran-agrees-to-end-enrichment-of-uranium-by-may-31-945 | geopolitics | 0.0770 | P9 | P9: geopolitics pump cluster (price 0.08, 1d) |
| 2026-05-29T23:30:11Z | iran-agrees-to-unrestricted-shipping-through-hormuz-by-may-31 | geopolitics | 0.0960 | P9 | P9: geopolitics pump cluster (price 0.10, 1d) |
| 2026-05-29T23:30:11Z | will-rajasthan-royals-win-the-2026-indian-premier-league | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $318162) |
| 2026-05-29T23:30:11Z | will-trump-agree-to-iranian-enrichment-of-uranium-by-may-31 | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 1 d, liq $72941) |
| 2026-05-29T23:30:11Z | strait-of-hormuz-traffic-returns-to-normal-by-end-of-may | geopolitics | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 1 d, liq $536682) |
| 2026-05-29T23:30:11Z | us-iran-nuclear-deal-by-may-31-974 | geopolitics | 0.0740 | P9 | P9: geopolitics pump cluster (price 0.07, 1d) |
| 2026-05-29T23:30:11Z | will-baidu-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $200947) |
| 2026-05-29T23:30:11Z | us-x-iran-permanent-peace-deal-by-may-31-2026-333-871-241-192-799-449-125 | geopolitics | 0.0700 | P9 | P9: geopolitics pump cluster (price 0.07, 1d) |
| 2026-05-29T23:30:11Z | will-mistral-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $187652) |
| 2026-05-29T23:30:11Z | will-reza-pahlavi-enter-iran-by-may-31 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $181759) |
| 2026-05-29T23:30:11Z | trump-out-as-president-by-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $931031) |
| 2026-05-29T23:30:13Z | iran-agrees-to-surrender-enriched-uranium-stockpile-by-may-31-2026 | geopolitics | 0.0270 | E2 | edge 0.013 < mín 0.030 (p̂=0.040, implied=0.027) |
| 2026-05-29T23:30:13Z | will-trump-agree-to-iranian-transit-fees-in-the-strait-of-hormuz-by-may-31 | geopolitics | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 1 d, liq $61947) |
| 2026-05-29T23:30:13Z | israel-x-lebanon-diplomatic-meeting-by-may-31-177 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T23:30:16Z | israel-closes-its-airspace-by-may-31 | geopolitics | 0.0450 | E2 | edge 0.025 < mín 0.030 (p̂=0.020, implied=0.045) |
| 2026-05-29T23:30:18Z | israel-x-hezbollah-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0290 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T23:30:18Z | iran-closes-its-airspace-by-may-31-434-443-672-526-188-756 | geopolitics | 0.0630 | P9 | P9: geopolitics pump cluster (price 0.06, 1d) |
| 2026-05-29T23:30:18Z | mojtaba-khamenei-leaves-iran-by-may-31-2026 | geopolitics | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $241054) |
| 2026-05-29T23:30:18Z | will-donald-trump-dance-on-may-25-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $368262) |
| 2026-05-29T23:30:20Z | israel-x-iran-permanent-peace-deal-by-may-31-2026 | geopolitics | 0.0400 | V6 | V6: V6 - Sin catalyst: no hay evento discreto verificable en los próximos 7 días que mueva sustancialmente la proba... |
| 2026-05-29T23:30:20Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-31-2026-313... | geopolitics | 0.2900 | P9 | P9: geopolitics pump cluster (price 0.29, 1d) |
| 2026-05-29T23:30:20Z | will-zai-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $221653) |
| 2026-05-29T23:30:20Z | us-obtains-iranian-enriched-uranium-by-may-31-396 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 1 d, liq $223908) |
| 2026-05-29T23:30:20Z | will-amazon-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $167187) |
| 2026-05-29T23:30:20Z | iran-leadership-change-by-may-31-593-194-829 | geopolitics | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 1 d, liq $49065) |
| 2026-05-29T23:30:23Z | us-x-iran-permanent-peace-deal-by-june-30-2026-837-641-896-877-363-892-537-597 | geopolitics | 0.3700 | V3, V5, V6 | V3, V5, V6: V3 Trigger vago: 'permanent peace deal' es un término ambiguo sin una definición verificable y objetiva... |
| 2026-05-29T23:30:25Z | will-arsenal-win-the-202526-champions-league | sports-season | 0.4400 | V3 | V3: V3 Trigger vago: sin fecha concreta o evento verificable. La pregunta es sobre una competición que finaliza en m... |
| 2026-05-29T23:30:28Z | will-spacex-ipo-by-may-31-2026-259 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $138065) |
| 2026-05-29T23:30:28Z | will-meituan-have-the-best-ai-model-at-the-end-of-may-2026 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $198493) |
| 2026-05-29T23:30:28Z | will-broadcom-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $547917) |
| 2026-05-29T23:30:28Z | will-alphabet-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $171926) |
| 2026-05-29T23:30:28Z | will-nvidia-be-the-second-largest-company-in-the-world-by-market-cap-on-may-31 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 1 d, liq $27196) |
| 2026-05-29T23:30:28Z | will-saudi-aramco-be-the-largest-company-in-the-world-by-market-cap-on-may-31 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $97434) |
| 2026-05-29T23:30:34Z | will-luis-gilberto-murillo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116593) |
| 2026-05-29T23:30:34Z | will-sergio-fajardo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $91968) |
| 2026-05-29T23:30:34Z | will-juan-carlos-pinzn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $70284) |
| 2026-05-29T23:30:34Z | will-mauricio-crdenas-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $116733) |
| 2026-05-29T23:30:34Z | will-claudia-lpez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $102265) |
| 2026-05-29T23:30:34Z | will-juan-manuel-galn-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $85698) |
| 2026-05-29T23:30:34Z | will-juan-daniel-oviedo-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $67949) |
| 2026-05-29T23:30:34Z | will-gustavo-bolvar-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $94533) |
| 2026-05-29T23:30:34Z | will-germn-vargas-lleras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $119098) |
| 2026-05-29T23:30:34Z | will-roy-barreras-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $111930) |
| 2026-05-29T23:30:34Z | will-david-luna-snchez-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $84896) |
| 2026-05-29T23:30:34Z | will-candidate-h-win-the-1st-round-of-the-2026-colombian-presidential-election-146-339 | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $126543) |
| 2026-05-29T23:30:34Z | will-vicky-dvila-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $66328) |
| 2026-05-29T23:30:34Z | will-daniel-quintero-win-the-1st-round-of-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 1 d, liq $126308) |
| 2026-05-29T23:30:34Z | kharg-island-no-longer-under-iranian-control-by-may-31-689 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $294152) |
| 2026-05-29T23:30:34Z | will-wti-dip-to-70-in-may-2026-155-395-333-182-889-988-959-341-883-234-375-347-135 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $155169) |
| 2026-05-29T23:30:34Z | will-wti-reach-140-in-may-2026-916 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $108472) |
| 2026-05-29T23:30:34Z | will-wti-dip-to-80-in-may-2026-734-629-766-626-157-663-256-712-397-816-113-876-832 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $124052) |
| 2026-05-29T23:30:34Z | will-wti-crude-oil-wti-hit-low-85-in-may-266-388-493-155-935-263-225-869-296-224-342-748-678 | market | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 2 d, liq $177821) |
| 2026-05-29T23:30:34Z | will-wti-crude-oil-wti-hit-high-105-in-may | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $113062) |
| 2026-05-29T23:30:34Z | will-wti-reach-130-in-may-2026-733 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $144540) |
| 2026-05-29T23:30:34Z | will-natural-gas-ng-hit-high-3pt40-in-may | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $8950) |
| 2026-05-29T23:30:34Z | will-wti-reach-110-in-may-2026-116-472 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $20321) |
| 2026-05-29T23:30:34Z | will-nvda-dip-to-192-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $10452) |
| 2026-05-29T23:30:34Z | will-wti-reach-120-in-may-2026-592-217 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $131667) |
| 2026-05-29T23:30:34Z | will-wti-crude-oil-wti-hit-high-115-in-may-221 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $8373) |
| 2026-05-29T23:30:34Z | will-wti-dip-to-40-in-may-2026-286-341-192-757-961-613-811-817-951-378-218-432-375 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $115735) |
| 2026-05-29T23:30:34Z | will-wti-reach-150-in-may-2026-196-364 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $106306) |
| 2026-05-29T23:30:34Z | will-wti-crude-oil-wti-hit-high-100-in-may | market | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 2 d, liq $176783) |
| 2026-05-29T23:30:34Z | will-bitcoin-dip-to-70k-in-may-2026-438-356-919 | market | 0.0470 | P11 | tradingview BTCUSDT sentiment=bearish (conf 0.65) contradicts bull thesis (price_yes=0.97) |
| 2026-05-29T23:30:34Z | will-ethereum-dip-to-1600-in-may-2026 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $42409) |
| 2026-05-29T23:30:34Z | will-bitcoin-reach-90k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $166210) |
| 2026-05-29T23:30:37Z | will-bitcoin-reach-85k-in-may-2026 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $214250) |
| 2026-05-29T23:30:37Z | will-bitcoin-dip-to-60k-in-may-2026-973-269 | market | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 2 d, liq $128981) |
| 2026-05-29T23:30:37Z | elon-musk-of-tweets-may-2026-780-799 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $426159) |
| 2026-05-29T23:30:37Z | will-bitcoin-dip-to-30k-in-may-2026-696-525-724 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $113357) |
| 2026-05-29T23:30:37Z | will-bitcoin-reach-150k-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $1081681) |
| 2026-05-29T23:30:37Z | will-ethereum-reach-3600-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $55075) |
| 2026-05-29T23:30:37Z | will-ethereum-dip-to-1000-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $29982) |
| 2026-05-29T23:30:37Z | will-bitcoin-reach-77500-in-may-2026-from-may-27 | market | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 2 d, liq $16114) |
| 2026-05-29T23:30:37Z | will-solana-reach-150-in-may-2026 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $55224) |
| 2026-05-29T23:30:37Z | will-bitcoin-dip-to-50k-in-may-2026-896 | market | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 2 d, liq $103369) |
| 2026-05-29T23:30:37Z | over-40m-committed-to-the-printr-public-sale | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $10337) |
| 2026-05-29T23:30:37Z | will-bitcoin-dip-to-72500-in-may-2026-from-may-27 | market | 0.4600 | P11 | tradingview BTCUSDT sentiment=bearish (conf 0.65) contradicts bull thesis (price_yes=0.97) |
| 2026-05-29T23:30:37Z | will-bitcoin-reach-80000-in-may-2026-from-may-27 | market | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 2 d, liq $17965) |
| 2026-05-29T23:30:37Z | will-bitcoin-dip-to-55k-in-may-2026-941 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 2 d, liq $108339) |
| 2026-05-29T23:30:37Z | will-backrooms-opening-weekend-box-office-be-greater-than-61m | entertainment | 0.9930 | P0_ceiling | price ceiling: 0.9930 > 0.980 |
| 2026-05-29T23:30:37Z | will-rob-sand-win-the-2026-iowa-governor-democratic-primary-election | elections | 0.9960 | P0_ceiling | price ceiling: 0.9960 > 0.980 |
| 2026-05-29T23:30:38Z | will-steve-lanier-win-the-2026-new-mexico-governor-republican-primary-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $17750) |
| 2026-05-29T23:30:38Z | will-rick-caruso-win-the-2026-los-angeles-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $32002) |
| 2026-05-29T23:30:38Z | elon-musk-of-tweets-may-26-june-2-80-99 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $52113) |
| 2026-05-29T23:30:38Z | elon-musk-of-tweets-may-26-june-2-60-79 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 3 d, liq $92887) |
| 2026-05-29T23:30:38Z | elon-musk-of-tweets-may-26-june-2-100-119 | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 3 d, liq $45479) |
| 2026-05-29T23:30:38Z | will-yoon-jae-ok-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $21573) |
| 2026-05-29T23:30:38Z | will-kim-byeong-ju-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $44553) |
| 2026-05-29T23:30:38Z | will-suh-byung-soo-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $31305) |
| 2026-05-29T23:30:41Z | will-oh-se-hoon-win-the-2026-seoul-mayoral-election | elections | 0.2700 | E2 | edge 0.030 < mín 0.030 (p̂=0.300, implied=0.270) |
| 2026-05-29T23:30:41Z | will-yoon-sang-hyun-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $31026) |
| 2026-05-29T23:30:41Z | will-cho-kuk-win-the-2026-busan-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $30296) |
| 2026-05-29T23:30:41Z | will-na-kyung-won-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $45340) |
| 2026-05-29T23:30:41Z | will-reform-party-rp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $94721) |
| 2026-05-29T23:30:41Z | will-kim-dong-yeon-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $22260) |
| 2026-05-29T23:30:41Z | will-the-rebuilding-korea-party-rkp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $83019) |
| 2026-05-29T23:30:43Z | will-ahn-cheol-soo-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $41093) |
| 2026-05-29T23:30:43Z | will-yoo-dong-soo-win-the-2026-incheon-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $26200) |
| 2026-05-29T23:30:43Z | will-seo-jae-heon-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $21054) |
| 2026-05-29T23:30:43Z | will-han-jun-ho-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $23977) |
| 2026-05-29T23:30:43Z | will-the-progressive-party-pp-win-the-2026-south-korean-local-elections | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $94716) |
| 2026-05-29T23:30:43Z | will-yoo-seong-min-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $36158) |
| 2026-05-29T23:30:43Z | will-lee-jae-man-win-the-2026-daegu-mayoral-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $41089) |
| 2026-05-29T23:30:43Z | will-lee-jun-seok-win-the-2026-gyeonggi-province-gubernatorial-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 4 d, liq $46452) |
| 2026-05-29T23:30:43Z | elon-musk-of-tweets-may-29-june-5-500plus | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 6 d, liq $16463) |
| 2026-05-29T23:30:43Z | will-maya-joint-win-the-2026-womens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 7 d, liq $230490) |
| 2026-05-29T23:30:43Z | us-x-iran-permanent-peace-deal-by-june-7-2026 | geopolitics | 0.2000 | P9 | P9: geopolitics pump cluster (price 0.20, 8d) |
| 2026-05-29T23:30:43Z | will-marisol-prez-tello-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $121985) |
| 2026-05-29T23:30:46Z | will-george-forsyth-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $193962) |
| 2026-05-29T23:30:46Z | will-roberto-chiabra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $212586) |
| 2026-05-29T23:30:46Z | will-mario-vizcarra-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $206235) |
| 2026-05-29T23:30:48Z | will-martin-landaluce-win-the-2026-mens-french-open | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 8 d, liq $54296) |
| 2026-05-29T23:30:51Z | will-mesas-guevara-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $240140) |
| 2026-05-29T23:30:51Z | will-moise-kouame-win-the-2026-mens-french-open | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 8 d, liq $37307) |
| 2026-05-29T23:30:51Z | will-matteo-berrettini-win-the-2026-mens-french-open | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 8 d, liq $52301) |
| 2026-05-29T23:30:51Z | will-wolfgang-grozo-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152389) |
| 2026-05-29T23:30:51Z | will-jos-luna-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $214376) |
| 2026-05-29T23:30:54Z | will-learner-tien-win-the-2026-mens-french-open | other | 0.0240 | V3 | V3: V3 Trigger vago: sin fecha concreta ni evento verificable que asegure la resolución; el evento es extremadamente... |
| 2026-05-29T23:30:55Z | will-alejandro-tabilo-win-the-2026-mens-french-open | other | 0.0200 | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda alterar significativamente l... | V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda alterar significativamente l... |
| 2026-05-29T23:30:55Z | will-alex-michelsen-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $356222) |
| 2026-05-29T23:30:55Z | will-vladimir-cerrn-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $219656) |
| 2026-05-29T23:30:58Z | will-rafael-jodar-win-the-2026-mens-french-open | other | 0.1120 | V3 Trigger vago: sin fecha concreta o sin evento verificable. | V3 Trigger vago: sin fecha concreta o sin evento verificable.: Veto V3: El evento 'Rafael Jodar gana Roland Garros 20... |
| 2026-05-29T23:30:58Z | will-csar-acua-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $182786) |
| 2026-05-29T23:31:00Z | will-andrey-rublev-win-the-2026-mens-french-open | other | 0.0320 | E2 | edge 0.008 < mín 0.030 (p̂=0.040, implied=0.032) |
| 2026-05-29T23:31:03Z | will-yonhy-lescano-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $202189) |
| 2026-05-29T23:31:05Z | will-carlos-esp-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $110143) |
| 2026-05-29T23:31:05Z | will-rafael-belande-llosa-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $190557) |
| 2026-05-29T23:31:05Z | will-novak-djokovic-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $686230) |
| 2026-05-29T23:31:06Z | will-francisco-cerundolo-win-the-2026-mens-french-open | other | 0.0620 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda cambiar significativamen... |
| 2026-05-29T23:31:06Z | will-fernando-olivera-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $152501) |
| 2026-05-29T23:31:08Z | will-flavio-cobolli-win-the-2026-mens-french-open | other | 0.0400 | El mercado especula sobre un resultado deportivo a más de un año, sin ningún evento discreto verificable en los pr... | El mercado especula sobre un resultado deportivo a más de un año, sin ningún evento discreto verificable en los pr... |
| 2026-05-29T23:31:08Z | will-jakub-mensik-win-the-2026-mens-french-open | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 8 d, liq $64756) |
| 2026-05-29T23:31:08Z | will-tommy-paul-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $290713) |
| 2026-05-29T23:31:10Z | will-felix-auger-aliassime-win-the-2026-mens-french-open | other | 0.0490 | E2 | edge 0.019 < mín 0.030 (p̂=0.030, implied=0.049) |
| 2026-05-29T23:31:10Z | will-karen-khachanov-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $392359) |
| 2026-05-29T23:31:10Z | will-alex-de-minaur-win-the-2026-mens-french-open | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $316242) |
| 2026-05-29T23:31:12Z | will-frances-tiafoe-win-the-2026-mens-french-open | other | 0.0200 | E2 | edge 0.019 < mín 0.030 (p̂=0.001, implied=0.020) |
| 2026-05-29T23:31:15Z | will-jos-williams-win-the-2026-peruvian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 8 d, liq $204757) |
| 2026-05-29T23:31:21Z | strait-of-hormuz-traffic-returns-to-normal-by-june-15 | geopolitics | 0.0900 | E2 | edge 0.010 < mín 0.030 (p̂=0.080, implied=0.090) |
| 2026-05-29T23:31:25Z | will-the-fed-increase-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 18 d, liq $795060) |
| 2026-05-29T23:31:27Z | will-there-be-no-change-in-fed-interest-rates-after-the-june-2026-meeting | other | 0.9780 | E2 | edge 0.028 < mín 0.030 (p̂=0.950, implied=0.978) |
| 2026-05-29T23:31:27Z | will-the-fed-decrease-interest-rates-by-25-bps-after-the-june-2026-meeting | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 18 d, liq $245879) |
| 2026-05-29T23:31:27Z | will-the-fed-increase-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 18 d, liq $2619417) |
| 2026-05-29T23:31:27Z | will-the-fed-decrease-interest-rates-by-50-bps-after-the-june-2026-meeting | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 18 d, liq $617661) |
| 2026-05-29T23:31:27Z | will-m80-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 22 d, liq $38888) |
| 2026-05-29T23:31:27Z | will-lynn-vision-win-iem-cologne-major-2026 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 22 d, liq $20850) |
| 2026-05-29T23:31:27Z | will-juan-carlos-pinzn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $149465) |
| 2026-05-29T23:31:27Z | will-sergio-fajardo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $52608) |
| 2026-05-29T23:31:27Z | will-germn-vargas-lleras-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $226915) |
| 2026-05-29T23:31:27Z | will-mauricio-cardenas-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $207565) |
| 2026-05-29T23:31:27Z | will-gustavo-bolvar-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $195581) |
| 2026-05-29T23:31:27Z | will-claudia-lpez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $145782) |
| 2026-05-29T23:31:29Z | will-juan-manuel-galn-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $192213) |
| 2026-05-29T23:31:29Z | will-juan-daniel-oviedo-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $250532) |
| 2026-05-29T23:31:33Z | will-paloma-valencia-win-the-2026-colombian-presidential-election | elections | 0.0260 | E2 | edge 0.014 < mín 0.030 (p̂=0.040, implied=0.026) |
| 2026-05-29T23:31:33Z | will-enrique-pealosa-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $176937) |
| 2026-05-29T23:31:33Z | will-daniel-quintero-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $188041) |
| 2026-05-29T23:31:33Z | will-david-luna-snchez-win-the-2026-colombian-presidential-election | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 22 d, liq $219016) |
| 2026-05-29T23:31:35Z | will-steve-hershey-win-the-2026-maryland-governor-republican-primary-election | elections | 0.0200 | E2 | edge 0.010 < mín 0.030 (p̂=0.010, implied=0.020) |
| 2026-05-29T23:31:35Z | will-aj-dybantsa-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 25 d, liq $1039) |
| 2026-05-29T23:31:35Z | will-yaxel-lendeborg-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $1351) |
| 2026-05-29T23:31:35Z | will-darryn-peterson-be-the-5th-overall-pick-in-the-2026-nba-draft | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 25 d, liq $1190) |
| 2026-05-29T23:31:35Z | will-tynan-lawrence-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $2951) |
| 2026-05-29T23:31:35Z | will-nikita-shcherbakov-be-drafted-1st-overall-in-the-2026-nhl-draft | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 28 d, liq $1605) |
| 2026-05-29T23:31:37Z | will-the-iranian-regime-fall-by-june-30 | other | 0.0250 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T23:31:41Z | us-obtains-iranian-enriched-uranium-by-june-30 | other | 0.0600 | E2 | edge 0.020 < mín 0.030 (p̂=0.080, implied=0.060) |
| 2026-05-29T23:31:41Z | will-tesla-be-the-largest-company-in-the-world-by-market-cap-on-june-30 | market | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $385579) |
| 2026-05-29T23:31:44Z | will-russia-enter-kramatorsk-by-june-30-821-192 | geopolitics | 0.0280 | E2 | edge 0.008 < mín 0.030 (p̂=0.020, implied=0.028) |
| 2026-05-29T23:31:44Z | gustavo-petro-out-as-leader-of-colombia-by-june-30 | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 31 d, liq $15190) |
| 2026-05-29T23:31:46Z | ukraine-agrees-not-to-join-nato-by-june-30 | geopolitics | 0.0330 | E2 | edge 0.017 < mín 0.030 (p̂=0.050, implied=0.033) |
| 2026-05-29T23:31:46Z | hamad-bin-isa-al-khalifa-out-as-bahrain-king | other | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 31 d, liq $14476) |
| 2026-05-29T23:31:46Z | will-trump-pardon-tiger-woods-by-june-30 | executive-action | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 31 d, liq $25894) |
| 2026-05-29T23:31:50Z | will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-june-30-2026-15... | geopolitics | 0.7300 | V3: Trigger vago | V3: Trigger vago: V3 Trigger vago: El evento no especifica una fecha concreta ni una fuente verificable para el anunc... |
| 2026-05-29T23:31:51Z | will-the-montreal-canadiens-win-the-2026-nhl-stanley-cup | other | 0.0300 | E2 | edge 0.010 < mín 0.030 (p̂=0.020, implied=0.030) |
| 2026-05-29T23:31:53Z | will-jay-z-be-confirmed-to-have-visited-epsteins-island | other | 0.0210 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T23:31:53Z | will-trump-and-putin-meet-next-in-japan-711-288-527 | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 31 d, liq $16515) |
| 2026-05-29T23:31:53Z | claudia-sheinbaum-out-as-president-of-mexico-by-june-30-791 | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 31 d, liq $11642) |
| 2026-05-29T23:31:56Z | israel-x-iran-permanent-peace-deal-by-june-30-2026-262 | geopolitics | 0.1690 | V3 - Trigger vago | V3 - Trigger vago: V3 Trigger vago: 'permanent peace deal' entre Israel e Irán para junio de 2026 no tiene fecha con... |
| 2026-05-29T23:31:58Z | will-deepseek-have-the-top-ai-model-at-the-end-of-june-2026-514 | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 31 d, liq $8768) |
| 2026-05-29T23:32:00Z | hantavirus-lab-leak-confirmed-by-june-30-1 | weather-natural | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 31 d, liq $97451) |
| 2026-05-29T23:32:03Z | will-russia-enter-kherson-by-june-30-645 | geopolitics | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 31 d, liq $5476) |
| 2026-05-29T23:32:03Z | will-china-invade-taiwan-by-june-30-2026 | geopolitics | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 31 d, liq $101771) |
| 2026-05-29T23:32:04Z | will-amazon-be-the-largest-company-in-the-world-by-market-cap-on-june-30-123 | other | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 31 d, liq $340847) |
| 2026-05-29T23:32:07Z | nato-x-russia-military-clash-by-june-30-2026 | geopolitics | 0.0560 | E2 | edge 0.026 < mín 0.030 (p̂=0.030, implied=0.056) |
| 2026-05-29T23:32:09Z | european-country-agrees-to-give-ukraine-security-guarantee-by-june-30 | geopolitics | 0.0420 | E2 | edge 0.012 < mín 0.030 (p̂=0.030, implied=0.042) |
| 2026-05-29T23:32:09Z | will-silver-si-hit-high-200-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $31192) |
| 2026-05-29T23:32:09Z | will-silver-si-hit-high-150-by-end-of-june | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 31 d, liq $37467) |
| 2026-05-29T23:32:11Z | will-crude-oil-cl-hit-high-120-by-end-of-june | market | 0.1200 | E2 | edge 0.020 < mín 0.030 (p̂=0.100, implied=0.120) |
| 2026-05-29T23:32:12Z | will-crude-oil-cl-hit-high-140-by-end-of-june-828-295-574-155 | market | 0.0390 | E2 | edge 0.019 < mín 0.030 (p̂=0.020, implied=0.039) |
| 2026-05-29T23:32:14Z | kharg-island-no-longer-under-iranian-control-by-june-30-561-854 | other | 0.0500 | E2 | edge 0.020 < mín 0.030 (p̂=0.030, implied=0.050) |
| 2026-05-29T23:32:25Z | will-madison-keys-be-the-2026-womens-wimbledon-winner | other | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 43 d, liq $2360) |
| 2026-05-29T23:32:25Z | will-karolna-muchov-be-the-2026-womens-wimbledon-winner | other | 0.0200 | P3_low_absolute_liquidity | liquidity $2436 < absolute min $5000 |
| 2026-05-29T23:32:25Z | will-mirra-andreeva-be-the-2026-womens-wimbledon-winner | other | 0.0450 | P3_low_absolute_liquidity | liquidity $2902 < absolute min $5000 |
| 2026-05-29T23:32:25Z | will-iga-witek-be-the-2026-womens-wimbledon-winner | other | 0.1820 | P3_low_absolute_liquidity | liquidity $1718 < absolute min $5000 |
| 2026-05-29T23:32:25Z | will-clara-tauson-be-the-2026-womens-wimbledon-winner | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 43 d, liq $2457) |
| 2026-05-29T23:32:25Z | will-emma-raducanu-be-the-2026-womens-wimbledon-winner | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 43 d, liq $2104) |
| 2026-05-29T23:32:25Z | will-jessica-pegula-be-the-2026-womens-wimbledon-winner | other | 0.0320 | P3_low_absolute_liquidity | liquidity $2611 < absolute min $5000 |
| 2026-05-29T23:32:25Z | claudio-tapia-out-as-afa-president-by-july-19-2026 | other | 0.4000 | P3_low_absolute_liquidity | liquidity $635 < absolute min $5000 |
| 2026-05-29T23:32:27Z | will-cape-verde-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $10466143) |
| 2026-05-29T23:32:30Z | will-ivory-coast-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4485250) |
| 2026-05-29T23:32:30Z | will-senegal-win-the-2026-fifa-world-cup | sports-season | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 51 d, liq $5218656) |
| 2026-05-29T23:32:32Z | will-netherlands-win-the-2026-fifa-world-cup-739 | sports-season | 0.0390 | E2 | edge 0.019 < mín 0.030 (p̂=0.020, implied=0.039) |
| 2026-05-29T23:32:32Z | will-croatia-win-the-2026-fifa-world-cup | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $2617686) |
| 2026-05-29T23:32:34Z | will-usa-win-the-2026-fifa-world-cup-467 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $5794791) |
| 2026-05-29T23:32:34Z | will-ghana-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $9645805) |
| 2026-05-29T23:32:34Z | will-paraguay-win-the-2026-fifa-world-cup-967 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $7304954) |
| 2026-05-29T23:32:34Z | will-saudi-arabia-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $9525558) |
| 2026-05-29T23:32:34Z | will-switzerland-win-the-2026-fifa-world-cup | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 51 d, liq $1149017) |
| 2026-05-29T23:32:34Z | will-uzbekistan-win-the-2026-fifa-world-cup-773 | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $7231139) |
| 2026-05-29T23:32:34Z | will-ecuador-win-the-2026-fifa-world-cup-986 | sports-season | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 51 d, liq $4740040) |
| 2026-05-29T23:32:34Z | will-south-africa-win-the-2026-fifa-world-cup | sports-season | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 51 d, liq $7611313) |
| 2026-05-29T23:32:36Z | will-brazil-win-the-2026-fifa-world-cup-183 | sports-season | 0.0940 | E2 | edge 0.014 < mín 0.030 (p̂=0.080, implied=0.094) |
| 2026-05-29T23:32:38Z | will-argentina-win-the-2026-fifa-world-cup-245 | sports-season | 0.0860 | E2 | edge 0.004 < mín 0.030 (p̂=0.090, implied=0.086) |
| 2026-05-29T23:32:38Z | will-uruguay-win-the-2026-fifa-world-cup-932 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $6997868) |
| 2026-05-29T23:32:41Z | will-germany-win-the-2026-fifa-world-cup-467 | sports-season | 0.0520 | E2 | edge 0.028 < mín 0.030 (p̂=0.080, implied=0.052) |
| 2026-05-29T23:32:41Z | will-colombia-win-the-2026-fifa-world-cup-734 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $5542928) |
| 2026-05-29T23:32:41Z | will-belgium-win-the-2026-fifa-world-cup-358 | sports-season | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 51 d, liq $4011813) |
| 2026-05-29T23:32:43Z | will-england-win-the-2026-fifa-world-cup-937 | sports-season | 0.1120 | E1 | edge no declarado por LLM (edge_type=none) |
| 2026-05-29T23:32:43Z | will-turkiye-win-the-2026-fifa-world-cup | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 51 d, liq $3458415) |
| 2026-05-29T23:32:43Z | will-sweden-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $8243827) |
| 2026-05-29T23:32:43Z | will-australia-win-the-2026-fifa-world-cup-816 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9348854) |
| 2026-05-29T23:32:43Z | will-bosnia-herzegovina-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $10506170) |
| 2026-05-29T23:32:43Z | will-canada-win-the-2026-fifa-world-cup-755 | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 51 d, liq $9233769) |
| 2026-05-29T23:32:45Z | will-norway-win-the-2026-fifa-world-cup-893 | sports-season | 0.0290 | V6 | V6: V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días que pueda cambiar significativamen... |
| 2026-05-29T23:32:45Z | will-mexico-win-the-2026-fifa-world-cup-529 | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 51 d, liq $2627957) |
| 2026-05-29T23:32:45Z | will-algeria-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $9795411) |
| 2026-05-29T23:32:45Z | will-austria-win-the-2026-fifa-world-cup | sports-season | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 51 d, liq $3033180) |
| 2026-05-29T23:32:45Z | will-japan-win-the-2026-fifa-world-cup-112 | sports-season | 0.0190 | P0_floor | price floor: 0.0190 < 0.020 (horizon 51 d, liq $442613) |
| 2026-05-29T23:32:45Z | will-tunisia-win-the-2026-fifa-world-cup-165 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $10672901) |
| 2026-05-29T23:32:45Z | will-scotland-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8681948) |
| 2026-05-29T23:32:45Z | will-south-korea-win-the-2026-fifa-world-cup-485 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $4517598) |
| 2026-05-29T23:32:45Z | will-iran-win-the-2026-fifa-world-cup-788 | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $6297724) |
| 2026-05-29T23:32:45Z | will-morocco-win-the-2026-fifa-world-cup-464 | sports-season | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 51 d, liq $2296675) |
| 2026-05-29T23:32:45Z | will-egypt-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8063909) |
| 2026-05-29T23:32:45Z | will-czechia-win-the-2026-fifa-world-cup | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 51 d, liq $8224236) |
| 2026-05-29T23:32:45Z | will-congo-dr-win-the-2026-fifa-world-cup | sports-season | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 51 d, liq $2463122) |
| 2026-05-29T23:32:47Z | will-jordan-wood-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $31701) |
| 2026-05-29T23:32:47Z | will-troy-jackson-be-the-democratic-nominee-for-senate-in-maine | elections | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 62 d, liq $14586) |
| 2026-05-29T23:32:48Z | will-paul-reevs-be-the-republican-nominee-for-az-01 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 66 d, liq $5891) |
| 2026-05-29T23:32:48Z | will-tyreek-hill-play-for-philadelphia-eagles-next | other | 0.1410 | P3_low_absolute_liquidity | liquidity $203 < absolute min $5000 |
| 2026-05-29T23:32:48Z | will-david-njoku-play-for-philadelphia-eagles-in-2026-27 | other | 0.0240 | P3_low_absolute_liquidity | liquidity $995 < absolute min $5000 |
| 2026-05-29T23:32:51Z | tush-push-banned-for-2026-nfl-season | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 103 d, liq $2563) |
| 2026-05-29T23:32:51Z | will-lorenzo-musetti-win-the-2026-mens-us-open | other | 0.0170 | P0_floor | price floor: 0.0170 < 0.020 (horizon 106 d, liq $2196) |
| 2026-05-29T23:32:51Z | will-hubert-hurkacz-win-the-2026-mens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $2957) |
| 2026-05-29T23:32:51Z | will-anastasia-potapova-win-the-2026-womens-us-open | other | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 106 d, liq $1161) |
| 2026-05-29T23:32:51Z | will-qinwen-zheng-win-the-2026-womens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $1846) |
| 2026-05-29T23:32:53Z | will-the-moderate-party-m-win-the-most-seats-in-the-2026-swedish-parliamentary-election | elections | 0.0360 | E2 | edge 0.004 < mín 0.030 (p̂=0.040, implied=0.036) |
| 2026-05-29T23:32:53Z | will-barbora-krejcikova-win-the-2026-womens-us-open | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 106 d, liq $1237) |
| 2026-05-29T23:32:53Z | will-felix-auger-aliassime-win-the-2026-mens-us-open | other | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 106 d, liq $2640) |
| 2026-05-29T23:32:53Z | will-jiri-lehecka-win-the-2026-mens-us-open | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 106 d, liq $1895) |
| 2026-05-29T23:32:53Z | will-matteo-berrettini-win-the-2026-mens-us-open | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 106 d, liq $3577) |
| 2026-05-29T23:32:53Z | will-cal-raleigh-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 121 d, liq $1378) |
| 2026-05-29T23:32:53Z | will-andy-pages-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0640 | P3_low_absolute_liquidity | liquidity $516 < absolute min $5000 |
| 2026-05-29T23:32:54Z | will-taylor-ward-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 121 d, liq $1838) |
| 2026-05-29T23:32:54Z | will-james-wood-lead-the-mlb-in-rbis-for-the-2026-regular-season | other | 0.0380 | P3_low_absolute_liquidity | liquidity $1018 < absolute min $5000 |
| 2026-05-29T23:32:54Z | will-united-russia-er-gain-the-most-seats-in-the-next-russian-parliamentary-election | elections | 0.5700 | P8 | P8: election 123d out, price 0.57 en banda ruido [0.30, 0.70] |
| 2026-05-29T23:32:54Z | will-consensys-ipo-by-september-30-2026 | crypto-launch | 0.0900 | P3_low_absolute_liquidity | liquidity $481 < absolute min $5000 |
| 2026-05-29T23:32:55Z | will-renan-santos-win-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0900 | E2 | edge 0.010 < mín 0.030 (p̂=0.100, implied=0.090) |
| 2026-05-29T23:32:55Z | will-michelle-bolsonaro-win-the-2026-brazilian-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 127 d, liq $224863) |
| 2026-05-29T23:32:55Z | will-ronaldo-caiado-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 127 d, liq $28439) |
| 2026-05-29T23:32:55Z | will-tereza-cristina-win-the-2026-brazilian-presidential-election | elections | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 127 d, liq $271691) |
| 2026-05-29T23:32:55Z | will-carlos-roberto-massa-jnior-finish-in-second-place-in-the-first-round-of-the-2026-brazilian-presidential-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 127 d, liq $28280) |
| 2026-05-29T23:32:58Z | will-athletics-win-the-2026-american-league-championship-series | sports-season | 0.0260 | E2 | edge 0.011 < mín 0.030 (p̂=0.015, implied=0.026) |
| 2026-05-29T23:32:58Z | will-chicago-white-sox-win-the-2026-american-league-championship-series | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 155 d, liq $27914) |
| 2026-05-29T23:32:58Z | will-kansas-city-royals-win-the-2026-american-league-championship-series | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 155 d, liq $31731) |
| 2026-05-29T23:32:58Z | will-los-angeles-angels-win-the-2026-american-league-championship-series | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 155 d, liq $31034) |
| 2026-05-29T23:32:58Z | will-minnesota-twins-win-the-2026-american-league-championship-series | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 155 d, liq $28438) |
| 2026-05-29T23:32:58Z | will-nicole-shanahan-win-the-california-governor-election-in-2026 | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $155053) |
| 2026-05-29T23:32:58Z | will-mary-peltola-win-the-2026-alaska-governor-election | elections | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 157 d, liq $7041) |
| 2026-05-29T23:32:58Z | will-corey-seager-win-the-2026-american-league-hank-aaron-award | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 167 d, liq $2979) |
| 2026-05-29T23:32:58Z | will-sergio-prez-be-the-2026-f1-drivers-champion | sports-season | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 190 d, liq $753707) |
| 2026-05-29T23:32:58Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt5-at-the-end-of-2026-a8ms-587 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $9309) |
| 2026-05-29T23:32:58Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt0-at-the-end-of-2026-434 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $9345) |
| 2026-05-29T23:32:58Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt0-at-the-end-of-2026-736 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 193 d, liq $12486) |
| 2026-05-29T23:32:58Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt75-at-the-end-of-2026-739 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 193 d, liq $10143) |
| 2026-05-29T23:32:58Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-1pt25-at-the-end-of-2026-681 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 193 d, liq $10563) |
| 2026-05-29T23:32:58Z | will-the-upper-bound-of-the-target-federal-funds-rate-be-2pt5-at-the-end-of-2026-876 | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 193 d, liq $12611) |
| 2026-05-29T23:32:58Z | will-carson-williams-win-the-2026-al-rookie-of-the-year-award | other | 0.0050 | P0_floor | price floor: 0.0050 < 0.020 (horizon 203 d, liq $6114) |
| 2026-05-29T23:32:58Z | will-brice-matthews-win-the-2026-al-rookie-of-the-year-award | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 203 d, liq $4183) |
| 2026-05-29T23:32:58Z | will-alexander-volkanovski-be-the-next-fighter-to-be-ranked-first-in-the-ufc-pound-for-pound-rankings-in-2026-457 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $435) |
| 2026-05-29T23:32:58Z | will-merab-dvalishvili-fight-rob-font-next | other | 0.0220 | P3_low_absolute_liquidity | liquidity $188 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-alexander-volkanovski-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $602) |
| 2026-05-29T23:32:58Z | will-khamzat-chimaev-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $535) |
| 2026-05-29T23:32:58Z | freddie-mac-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1179 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-7-fed-rate-cuts-happen-in-2026 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $158107) |
| 2026-05-29T23:32:58Z | will-youssef-zalal-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $301) |
| 2026-05-29T23:32:58Z | will-the-weeknd-be-the-top-spotify-artist-for-2026 | entertainment | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $12169) |
| 2026-05-29T23:32:58Z | will-bogdan-guskov-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0660 | P3_low_absolute_liquidity | liquidity $382 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-gitlab-be-acquired-before-2027-944-667 | other | 0.2100 | P3_low_absolute_liquidity | liquidity $4977 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-bp-be-acquired-before-2027-549-844 | other | 0.1400 | P3_low_absolute_liquidity | liquidity $4726 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-the-feds-lower-bound-reach-3pt0-or-lower-before-2027-676-111 | other | 0.1300 | P3_low_absolute_liquidity | liquidity $4639 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-zelenskyy-and-putin-meet-next-in-russia | geopolitics | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $9573) |
| 2026-05-29T23:32:58Z | will-there-be-5-or-more-confirmed-vei-4-or-higher-volcanic-eruptions-worldwide-in-2026-655 | weather-natural | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $7562) |
| 2026-05-29T23:32:58Z | will-zelenskyy-and-putin-meet-next-in-belarus | other | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon 215 d, liq $6572) |
| 2026-05-29T23:32:58Z | openai-1t-ipo-before-2027 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $4909 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-apple-release-a-new-product-line-before-2027 | other | 0.3600 | P3_low_absolute_liquidity | liquidity $2713 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-merab-dvalishvili-fight-alexandre-pantoja-next | other | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 215 d, liq $87) |
| 2026-05-29T23:32:58Z | will-ilia-topuria-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.5220 | P3_low_absolute_liquidity | liquidity $317 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-steve-ballmer-be-richest-person-on-december-31 | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $8025) |
| 2026-05-29T23:32:58Z | will-aljamain-sterling-be-the-ufc-featherweight-champion-on-december-31-2026 | sports-season | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $610) |
| 2026-05-29T23:32:58Z | celonis-ipo-before-2027 | crypto-launch | 0.1000 | P3_low_absolute_liquidity | liquidity $3391 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-khalil-rountree-jr-be-the-ufc-light-heavyweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $591) |
| 2026-05-29T23:32:58Z | hantavirus-pandemic-in-2026 | weather-natural | 0.0630 | P7 | P7: weather/natural-disaster cluster — no edge |
| 2026-05-29T23:32:58Z | will-alexandre-pantoja-be-ranked-first-in-the-ufc-pound-for-pound-rankings-at-the-end-of-2026 | other | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $598) |
| 2026-05-29T23:32:58Z | will-jeff-bezos-be-richest-person-on-december-31-243 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $6089) |
| 2026-05-29T23:32:58Z | will-there-be-at-least-12500-measles-cases-in-the-us-in-2026-512-499 | other | 0.0800 | P3_low_absolute_liquidity | liquidity $3344 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-shavkat-rakhmonov-be-the-ufc-welterweight-champion-on-december-31-2026 | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $692) |
| 2026-05-29T23:32:58Z | will-merab-dvalishvili-fight-ricky-simn-next | other | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 215 d, liq $393) |
| 2026-05-29T23:32:58Z | will-the-hunger-games-sunrise-on-the-reaping-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $143498) |
| 2026-05-29T23:32:58Z | anduril-ipo-before-2027 | crypto-launch | 0.1200 | P3_low_absolute_liquidity | liquidity $3179 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-bernard-arnault-be-richest-person-on-december-31-747 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $6301) |
| 2026-05-29T23:32:58Z | will-jumanji-3-be-the-top-grossing-movie-of-2026 | other | 0.0060 | P0_floor | price floor: 0.0060 < 0.020 (horizon 215 d, liq $100652) |
| 2026-05-29T23:32:58Z | mistral-ai-ipo-before-2027 | crypto-launch | 0.1600 | P3_low_absolute_liquidity | liquidity $3613 < absolute min $5000 |
| 2026-05-29T23:32:58Z | russia-x-ukraine-ceasefire-agreement-by-may-31-2026 | geopolitics | 0.0030 | P0_floor | price floor: 0.0030 < 0.020 (horizon 215 d, liq $150834) |
| 2026-05-29T23:32:58Z | will-bruno-mars-be-the-top-spotify-artist-for-2026 | entertainment | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $16130) |
| 2026-05-29T23:32:58Z | will-merab-dvalishvili-fight-song-yadong-next | other | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 215 d, liq $171) |
| 2026-05-29T23:32:58Z | will-satoshis-identity-be-revealed-by-june-30 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $13960) |
| 2026-05-29T23:32:58Z | will-ubisoft-be-acquired-before-2027-175-488 | other | 0.2000 | P3_low_absolute_liquidity | liquidity $2920 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-7-8-spacex-starship-launches-successfully-reach-space-in-2026 | other | 0.0590 | P3_low_absolute_liquidity | liquidity $3832 < absolute min $5000 |
| 2026-05-29T23:32:58Z | eunato-country-announces-peacekeeping-force-in-ukraine-before-2027 | geopolitics | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 215 d, liq $9853) |
| 2026-05-29T23:32:58Z | ledger-ipo-before-2027 | crypto-launch | 0.1500 | P3_low_absolute_liquidity | liquidity $1765 < absolute min $5000 |
| 2026-05-29T23:32:58Z | fannie-mae-ipo-before-2027 | crypto-launch | 0.1300 | P3_low_absolute_liquidity | liquidity $3583 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-zelenskyy-and-putin-meet-next-in-india-416-286 | other | 0.0040 | P0_floor | price floor: 0.0040 < 0.020 (horizon 215 d, liq $11604) |
| 2026-05-29T23:32:58Z | will-austria-recognize-palestine-before-2027 | executive-action | 0.1230 | P3_low_absolute_liquidity | liquidity $3149 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-lucy-powell-be-the-next-prime-minister-of-the-united-kingdom-in-2026-663 | elections | 0.0020 | P0_floor | price floor: 0.0020 < 0.020 (horizon 215 d, liq $67302) |
| 2026-05-29T23:32:58Z | will-petr-yan-fight-alexander-volkanovski-next | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $752) |
| 2026-05-29T23:32:58Z | will-zelenskyy-and-putin-meet-next-in-switzerland-848-193 | other | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 215 d, liq $7872) |
| 2026-05-29T23:32:58Z | stripe-ipo-before-2027 | crypto-launch | 0.0900 | P4_pre_event | pre-event slug + 215 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | will-marcin-tybura-be-the-ufc-heavyweight-champion-on-december-31-2026 | sports-season | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon 215 d, liq $287) |
| 2026-05-29T23:32:58Z | will-andoni-iraola-be-appointed-as-manager-of-real-madrid | other | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 216 d, liq $3024) |
| 2026-05-29T23:32:58Z | will-ryan-flaherty-be-the-next-permanent-manager-of-the-philadelphia-phillies | other | 0.0780 | P3_low_absolute_liquidity | liquidity $797 < absolute min $5000 |
| 2026-05-29T23:32:58Z | standx-fdv-above-3b-one-day-after-launch-573-228-889-488-762 | crypto-launch | 0.0330 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | will-fomo-launch-a-token-by-june-30-2026 | crypto-launch | 0.0750 | P3_low_absolute_liquidity | liquidity $638 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-ostium-launch-a-token-by-december-31-2026 | crypto-launch | 0.6700 | P3_low_absolute_liquidity | liquidity $2433 < absolute min $5000 |
| 2026-05-29T23:32:58Z | puffpaw-fdv-above-300m-one-day-after-launch | crypto-launch | 0.2490 | P3_low_absolute_liquidity | liquidity $2388 < absolute min $5000 |
| 2026-05-29T23:32:58Z | consensys-ipo-closing-market-cap-above-1b | crypto-launch | 0.4000 | P3_low_absolute_liquidity | liquidity $1116 < absolute min $5000 |
| 2026-05-29T23:32:58Z | standx-fdv-above-1b-one-day-after-launch-758-887-458-572-867 | crypto-launch | 0.1400 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | metamask-fdv-above-2b-one-day-after-launch-222-955-573-228 | crypto-launch | 0.1200 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | standx-fdv-above-10b-one-day-after-launch-467-712-296 | crypto-launch | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 216 d, liq $12376) |
| 2026-05-29T23:32:58Z | will-tempo-launch-a-token-by-december-31-2026 | crypto-launch | 0.2200 | P3_low_absolute_liquidity | liquidity $1179 < absolute min $5000 |
| 2026-05-29T23:32:58Z | puffpaw-fdv-above-50m-one-day-after-launch | crypto-launch | 0.6900 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | will-hibachi-launch-a-token-by-december-31-2026 | crypto-launch | 0.3500 | P3_low_absolute_liquidity | liquidity $179 < absolute min $5000 |
| 2026-05-29T23:32:58Z | hyperliquid-listed-on-binance-in-2026-485 | other | 0.3400 | P3_low_absolute_liquidity | liquidity $3210 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-exponent-launch-a-token-by-december-31-2026 | crypto-launch | 0.3400 | P3_low_absolute_liquidity | liquidity $523 < absolute min $5000 |
| 2026-05-29T23:32:58Z | metamask-fdv-above-4b-one-day-after-launch-851-782-477-258 | crypto-launch | 0.0930 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | extended-fdv-above-300m-one-day-after-launch-351-333 | crypto-launch | 0.4500 | P4_pre_event | pre-event slug + 216 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | will-carolina-panthers-win-the-2027-nfl-nfc-championship-793 | sports-season | 0.0120 | P0_floor | price floor: 0.0120 < 0.020 (horizon 240 d, liq $52870) |
| 2026-05-29T23:32:58Z | will-new-york-jets-win-the-2027-nfl-afc-championship-268 | sports-season | 0.0160 | P0_floor | price floor: 0.0160 < 0.020 (horizon 240 d, liq $19000) |
| 2026-05-29T23:32:58Z | will-tennessee-titans-win-the-2027-nfl-afc-championship-594 | sports-season | 0.0130 | P0_floor | price floor: 0.0130 < 0.020 (horizon 240 d, liq $45813) |
| 2026-05-29T23:32:58Z | will-arizona-cardinals-win-the-2027-nfl-nfc-championship-199 | sports-season | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 240 d, liq $53122) |
| 2026-05-29T23:32:58Z | will-the-las-vegas-raiders-win-the-2027-nfl-league-championship | sports-season | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 306 d, liq $116787) |
| 2026-05-29T23:32:58Z | will-the-cleveland-browns-win-the-2027-nfl-league-championship | sports-season | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 306 d, liq $117845) |
| 2026-05-29T23:32:58Z | will-nicolas-dupont-aignan-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $177990) |
| 2026-05-29T23:32:58Z | will-franois-asselineau-win-the-2027-french-presidential-election | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 335 d, liq $261510) |
| 2026-05-29T23:32:58Z | openai-ipo-closing-market-cap-above-1t | crypto-launch | 0.8400 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | spacex-ipo-closing-market-cap-above-2t-649-783 | crypto-launch | 0.7500 | P4_pre_event | pre-event slug + 580 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | will-spacexs-public-ticker-be-mars-627 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $45560) |
| 2026-05-29T23:32:58Z | will-spacexs-public-ticker-be-spc-937 | crypto-launch | 0.0010 | P0_floor | price floor: 0.0010 < 0.020 (horizon 580 d, liq $23558) |
| 2026-05-29T23:32:58Z | openai-ipo-closing-market-cap-above-1pt2t | crypto-launch | 0.7500 | P3_low_absolute_liquidity | liquidity $4899 < absolute min $5000 |
| 2026-05-29T23:32:58Z | predictfun-fdv-above-300m-one-day-after-launch | crypto-launch | 0.6700 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | abstract-fdv-above-3b-one-day-after-launch | crypto-launch | 0.0480 | P4_pre_event | pre-event slug + 581 d to resolution (>=7 threshold) |
| 2026-05-29T23:32:58Z | will-liz-cheney-win-the-2028-democratic-presidential-nomination-551 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1173355) |
| 2026-05-29T23:32:58Z | will-barack-obama-win-the-2028-democratic-presidential-nomination-265 | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $1153816) |
| 2026-05-29T23:32:58Z | will-person-a-win-the-2028-democratic-presidential-nomination | elections | 0.0070 | P0_floor | price floor: 0.0070 < 0.020 (horizon 892 d, liq $2372541) |
| 2026-05-29T23:32:58Z | will-greg-abbott-win-the-2028-us-presidential-election | elections | 0.0090 | P0_floor | price floor: 0.0090 < 0.020 (horizon 892 d, liq $1916050) |
| 2026-05-29T23:32:58Z | will-zohran-mamdani-win-the-2028-democratic-presidential-nomination-445 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $2444817) |
| 2026-05-29T23:32:58Z | will-raphael-warnock-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1461022) |
| 2026-05-29T23:32:58Z | will-stephen-a-smith-win-the-2028-democratic-presidential-nomination-914 | elections | 0.0110 | P0_floor | price floor: 0.0110 < 0.020 (horizon 892 d, liq $783788) |
| 2026-05-29T23:32:58Z | will-ron-desantis-win-the-2028-us-presidential-election | elections | 0.0180 | P0_floor | price floor: 0.0180 < 0.020 (horizon 892 d, liq $300024) |
| 2026-05-29T23:32:58Z | will-bernie-sanders-win-the-2028-democratic-presidential-nomination-879 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $1446178) |
| 2026-05-29T23:32:58Z | will-jared-polis-win-the-2028-democratic-presidential-nomination-837 | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $476809) |
| 2026-05-29T23:32:58Z | will-vivek-ramaswamy-win-the-2028-us-presidential-election | elections | 0.0080 | P0_floor | price floor: 0.0080 < 0.020 (horizon 892 d, liq $867352) |
| 2026-05-29T23:32:58Z | will-roy-cooper-win-the-2028-democratic-presidential-nomination-286 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1476200) |
| 2026-05-29T23:32:58Z | will-stephen-smith-win-the-2028-us-presidential-election | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1081090) |
| 2026-05-29T23:32:58Z | will-corey-booker-win-the-2028-democratic-presidential-nomination-125 | elections | 0.0100 | P0_floor | price floor: 0.0100 < 0.020 (horizon 892 d, liq $1628231) |
| 2026-05-29T23:32:58Z | will-axiom-launch-a-token-by-december-31-2026 | crypto-launch | 0.2300 | P3_low_absolute_liquidity | liquidity $536 < absolute min $5000 |
| 2026-05-29T23:32:58Z | will-the-iran-ceasefire-continue-through-may-24-733 | geopolitics | 0.9990 | P0_ceiling | price ceiling: 0.9990 > 0.980 |
| 2026-05-29T23:32:58Z | will-the-republican-party-hold-exactly-56-senate-seats-after-the-2026-midterm-elections | elections | 0.0140 | P0_floor | price floor: 0.0140 < 0.020 (horizon -1 d, liq $12813) |
| 2026-05-29T23:32:58Z | will-the-republican-party-hold-exactly-54-senate-seats-after-the-2026-midterm-elections | elections | 0.0150 | P0_floor | price floor: 0.0150 < 0.020 (horizon -1 d, liq $23015) |
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
## Wins pattern (últimos 100, append-only, rotación a tail)

| timestamp | trade_id | category | entry | exit | pnl | reason | sources | size | horizon | days_held |
|---|---|---|---|---|---|---|---|---|---|---|
| 2026-05-29T20:40:01Z | T-2340268-1780087006486 | market | 0.2900 | 0.3100 | 3.06 | no_remaining_edge |  | 44.37 | short | 0.00 |
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
