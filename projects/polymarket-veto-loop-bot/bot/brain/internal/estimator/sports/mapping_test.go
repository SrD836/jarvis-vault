package sports

import "testing"

func TestMapCompetition(t *testing.T) {
	cases := map[string]string{
		"Will the Oklahoma City Thunder win the 2026 NBA Finals?": "basketball_nba_championship_winner",
		"Will the Chiefs win the Super Bowl?":                     "americanfootball_nfl_super_bowl_winner",
		"Will the Dodgers win the World Series?":                  "baseball_mlb_world_series_winner",
		"Will the Panthers win the Stanley Cup?":                  "icehockey_nhl_championship_winner",
	}
	for q, want := range cases {
		got, ok := MapCompetition(q)
		if !ok || got != want {
			t.Errorf("MapCompetition(%q) = %q,%v want %q", q, got, ok, want)
		}
	}
	// Uncovered competitions decline.
	for _, q := range []string{
		"Will PSG win the 2025–26 Champions League?",
		"Will the Oklahoma City Thunder win the NBA Western Conference Finals?",
	} {
		if _, ok := MapCompetition(q); ok {
			t.Errorf("MapCompetition(%q) should decline (uncovered)", q)
		}
	}
}

func TestExtractTeamAndMatch(t *testing.T) {
	q := "Will the Oklahoma City Thunder win the 2026 NBA Finals?"
	team := ExtractTeam(q)
	if team != "Oklahoma City Thunder" {
		t.Errorf("ExtractTeam = %q, want 'Oklahoma City Thunder'", team)
	}
	outcomes := []string{"San Antonio Spurs", "Oklahoma City Thunder", "Boston Celtics"}
	if i := MatchOutcome(team, outcomes); i != 1 {
		t.Errorf("MatchOutcome exact = %d, want 1", i)
	}
	// Nickname / partial: Polymarket "Thunder", odds-api full name.
	if i := MatchOutcome("Thunder", outcomes); i != 1 {
		t.Errorf("MatchOutcome nickname = %d, want 1", i)
	}
	// No match → -1.
	if i := MatchOutcome("Lakers", outcomes); i != -1 {
		t.Errorf("MatchOutcome no-match = %d, want -1", i)
	}
	// Non-"will X win" question → empty team.
	if ExtractTeam("Will Bitcoin reach $90,000 in May?") != "" {
		t.Error("ExtractTeam should be empty for non-team question")
	}
}
