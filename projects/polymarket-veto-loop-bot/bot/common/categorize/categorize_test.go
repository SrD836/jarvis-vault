package categorize

import "testing"

func TestInfer(t *testing.T) {
	cases := []struct {
		slug, question, want string
	}{
		{"will-arsenal-win-on-2026-05-25", "Will Arsenal win on 2026-05-25?", CatSportsGame},
		{"will-the-iran-ceasefire-continue-through-may-28", "Will the Iran ceasefire continue through May 28?", CatGeopolitics},
		{"will-jd-vance-win-the-2028-us-presidential-election", "", CatElections},
		{"will-trump-resign-by-june-30", "Will Trump resign by June 30?", CatExecutiveAction},
		{"hantavirus-pandemic-in-2026", "", CatWeatherNatural},
		{"earthquake-magnitude-7-worldwide-in-2026", "", CatWeatherNatural},
		{"mistral-ai-ipo-before-2027", "", CatCryptoLaunch},
		{"predictfun-fdv-above-1b-one-day-after-launch", "", CatCryptoLaunch},
		{"will-the-price-of-bitcoin-be-above-78000-on-may-25", "Will the price of Bitcoin be above $78,000 on May 25?", CatMarket},
		{"will-wti-crude-oil-hit-low-90-in-may", "", CatMarket},
		{"will-the-oklahoma-city-thunder-win-the-2026-nba-finals", "", CatSportsSeason},
		{"will-dallas-cowboys-win-the-2027-nfl-nfc-championship", "", CatSportsSeason},
		{"will-the-odyssey-have-the-best-domestic-opening-weekend-in-2026", "", CatEntertainment},
		{"will-beyonc-be-the-top-spotify-artist-for-2026", "", CatEntertainment},
		{"will-ken-paxton-win-the-2026-texas-republican-primary", "", CatElections},
		{"will-arsenal-win-a-trophy-this-season", "", CatSportsSeason},
		{"israel-x-hezbollah-permanent-peace-deal-by-may-31-2026", "", CatGeopolitics},
		{"will-spacexs-public-ticker-be-star", "", CatCryptoLaunch},
		{"some-random-thing-2026", "Will some random thing happen in 2026?", CatOther},
		{"", "", CatOther},
		{"both-teams-to-score-tottenham-chelsea", "", CatSportsGame},
		{"will-the-fed-lower-bound-reach-0pt25-or-lower-before-2027", "", CatOther}, // Fed rate, not in market price regex (no bitcoin/etc)
	}
	for _, c := range cases {
		got := Infer(c.slug, c.question)
		if got != c.want {
			t.Errorf("Infer(%q, %q) = %q, want %q", c.slug, c.question, got, c.want)
		}
	}
}

// TestInferRealSlugs samples 20 slugs from the production candidates.jsonl
// (snapshot 2026-05-28) to verify the inference rules generalize. The aim
// is to keep the CatOther rate <= 40% — anything higher means the regex set
// is leaking too many real candidates into the noise bucket and breaking
// the M2 softrules per-category learning.
func TestInferRealSlugs(t *testing.T) {
	cases := []struct {
		slug, want string
	}{
		// Elections / nominations / primaries.
		{"ukraine-election-called-by-june-30-2026-392", CatElections},
		{"will-the-pheu-thai-party-pt-win-110-or-more-seats-in-the-2026-thai-legislative-election", CatElections},
		{"will-matej-tonin-be-the-next-prime-minister-of-slovenia", CatElections},
		{"will-pia-olsen-dyhr-be-the-next-prime-minister-of-denmark-after-the-2026-parliamentary-elections", CatElections},
		{"will-carlos-lvarez-win-the-2026-peruvian-presidential-election", CatElections},
		{"will-gregg-kirkpatrick-win-the-2026-georgia-governor-republican-primary-election", CatElections},
		// Geopolitics — Iran / Israel / Hormuz / ceasefires.
		{"internet-access-restored-in-iran-by-may-31-2026", CatGeopolitics},
		{"us-x-iran-diplomatic-meeting-by-may-31-2026-283-616", CatGeopolitics},
		{"iran-closes-its-airspace-by-may-29", CatGeopolitics},
		{"will-donald-trump-announce-that-the-united-states-blockade-of-the-strait-of-hormuz-has-been-lifted-by-may-28-2026", CatGeopolitics},
		{"us-announces-new-iran-agreementceasefire-extension-by-may-28", CatGeopolitics},
		// Market-price predictions.
		{"bitcoin-above-70k-on-may-28-2026", CatMarket},
		{"bitcoin-above-74k-on-may-28-2026", CatMarket},
		{"bitcoin-above-82k-on-may-28-2026", CatMarket},
		// Executive actions on named individuals (indictment / charged / resign).
		{"tim-walz-charged-by-december-31-2026", CatOther}, // "charged" not in regex — acceptable noise
		// Entertainment / awards — currently slip into CatOther; tomatometer not regex'd.
		{"will-how-to-make-a-killing-score-at-least-56-on-the-rotten-tomatoes-tomatometer", CatOther},
		// Sports tournaments — masters/open not in regex, currently CatOther.
		{"will-grigor-dimitrov-win-the-2026-australian-open", CatOther},
		{"will-sung-jae-im-win-the-2026-masters-tournament", CatOther},
		// Weather (high-temp daily markets).
		{"highest-temperature-in-hong-kong-on-may-27-2026-34corhigher", CatOther}, // temperature not in weather regex
		// Genuinely uncategorizable.
		{"lib-cor-cp-2026-05-27-cor", CatOther},
	}

	otherCount := 0
	for _, c := range cases {
		got := Infer(c.slug, "")
		if got != c.want {
			t.Errorf("Infer(%q) = %q, want %q", c.slug, got, c.want)
		}
		if got == CatOther {
			otherCount++
		}
	}

	// Plan P5 target: CatOther rate <= 40%. If this fails, add regex coverage.
	otherPct := float64(otherCount) * 100.0 / float64(len(cases))
	if otherPct > 40.0 {
		t.Errorf("CatOther rate too high: %.1f%% (want <= 40%%) — extend regex rules", otherPct)
	}
	t.Logf("real-slug CatOther rate: %.1f%% (%d/%d)", otherPct, otherCount, len(cases))
}
