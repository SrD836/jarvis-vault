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
