package sports

import (
	"regexp"
	"strings"
)

type compRule struct {
	rx       *regexp.Regexp
	sportKey string
}

// compRules map a Polymarket "win the <competition>" question to the-odds-api
// futures (outright winner) sport key. Only competitions the-odds-api actually
// carries as futures are listed — Champions League / domestic-league titles are
// NOT covered by the provider, so those questions decline (fall back to the LLM).
var compRules = []compRule{
	{regexp.MustCompile(`(?i)\b(nba finals|nba championship|larry o.?brien)\b`), "basketball_nba_championship_winner"},
	{regexp.MustCompile(`(?i)\b(super bowl)\b`), "americanfootball_nfl_super_bowl_winner"},
	{regexp.MustCompile(`(?i)\b(world series)\b`), "baseball_mlb_world_series_winner"},
	{regexp.MustCompile(`(?i)\b(stanley cup|nhl championship)\b`), "icehockey_nhl_championship_winner"},
	{regexp.MustCompile(`(?i)\b(fifa world cup|world cup winner|win the .*world cup)\b`), "soccer_fifa_world_cup_winner"},
}

// MapCompetition returns the the-odds-api futures sport key for a Polymarket
// outright-winner question, or ok=false when no covered competition matches.
func MapCompetition(question string) (string, bool) {
	for _, r := range compRules {
		if r.rx.MatchString(question) {
			return r.sportKey, true
		}
	}
	return "", false
}

var teamRE = regexp.MustCompile(`(?i)^will\s+(?:the\s+)?(.+?)\s+win\b`)

// ExtractTeam pulls the competitor phrase from "Will [the] <team> win ...".
// Returns "" when the question doesn't fit that shape.
func ExtractTeam(question string) string {
	m := teamRE.FindStringSubmatch(strings.TrimSpace(question))
	if len(m) < 2 {
		return ""
	}
	return strings.TrimSpace(m[1])
}

// MatchOutcome returns the index of the outcome whose name best matches the team
// phrase, or -1. Tries exact-normalized, then bidirectional contains (handles
// "Thunder" vs "Oklahoma City Thunder"), then nickname (last word).
func MatchOutcome(team string, outcomes []string) int {
	t := normalize(team)
	if t == "" {
		return -1
	}
	for i, o := range outcomes {
		if normalize(o) == t {
			return i
		}
	}
	for i, o := range outcomes {
		no := normalize(o)
		if no != "" && (strings.Contains(no, t) || strings.Contains(t, no)) {
			return i
		}
	}
	if tw := lastWord(t); tw != "" {
		for i, o := range outcomes {
			if strings.Contains(normalize(o), tw) {
				return i
			}
		}
	}
	return -1
}

// normalize lowercases and collapses every non-alphanumeric run to a single
// space, trimmed. "Oklahoma City Thunder!" → "oklahoma city thunder".
func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	prevSpace := true
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevSpace = false
		} else if !prevSpace {
			b.WriteByte(' ')
			prevSpace = true
		}
	}
	return strings.TrimSpace(b.String())
}

func lastWord(s string) string {
	f := strings.Fields(s)
	if len(f) == 0 {
		return ""
	}
	return f[len(f)-1]
}
