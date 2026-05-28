// Package softrules — consumer side (M2 veto rule).
//
// LoadActiveRules parses the "## Reglas blandas aprendidas" section in
// memory.md and returns the most recent set of auto-generated cluster rules
// (deduped by category|horizon|band, last occurrence wins since softrules
// appends time-ordered).
//
// MatchVeto checks if a candidate (category, horizon, price) matches any
// active rule. If yes → brain blocks the candidate with rule "M2".
//
// Why: softrules.GenerateAndAppend WRITES rules to memory.md but nothing
// READ them back into the decision loop. Patterns like
// "uncategorized · short · 0.10-0.30 → 12L/1W = 8% win rate" were generated
// repeatedly but candidates matching kept being approved. M2 closes the loop.
package softrules

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ActiveRule is a parsed soft rule that vetoes candidates matching
// (category, horizon, band).
type ActiveRule struct {
	Category string
	Horizon  string
	Band     string // "<0.10", "0.10-0.30", "0.30-0.70", ">0.70"
	Losses   int
	Wins     int
	WinRate  float64 // 0-1
}

// Matches reports whether a candidate matches this rule.
//
// v7 P1 (2026-05-28): removed the legacy wildcard fallback whereby an
// "uncategorized" rule matched any non-market category. Rationale: 126
// uncategorized rules accumulated from pre-categorize losses and the
// wildcard caused them to soft-veto candidates whose real categories were
// nothing alike. New uncategorized rules are no longer generated
// (GenerateAndAppend filters them out); legacy ones become inert here and
// will be purged from memory.md by analytics/cleanup_legacy_soft_rules.py.
//
// Horizon "?" in the rule = wildcard (still honored — pre-categorize horizon
// was sometimes unresolvable).
func (r ActiveRule) Matches(category, horizon string, price float64) bool {
	cat := strings.TrimSpace(strings.ToLower(category))
	ruleCat := strings.TrimSpace(strings.ToLower(r.Category))
	if ruleCat == "" || ruleCat == "uncategorized" {
		// inert rule from legacy data — no implicit wildcard.
		return false
	}
	if ruleCat != cat {
		return false
	}
	if r.Horizon != "?" && !strings.EqualFold(r.Horizon, horizon) {
		return false
	}
	return priceBand(price) == r.Band
}

// ruleRE captures the auto-generated line format produced by GenerateAndAppend.
// Example:
//
//	- En categoría `uncategorized` · horizonte `short` · precio `0.10-0.30`: 12 losses, 1 wins (win rate 8%).
var ruleRE = regexp.MustCompile(
	"^- En categoría `([^`]+)` · horizonte `([^`]+)` · precio `([^`]+)`: (\\d+) losses, (\\d+) wins \\(win rate (\\d+)%\\)",
)

// LoadActiveRules reads memory.md and returns the deduped latest soft rules.
// Returns (nil, nil) if the section is absent (no error).
func LoadActiveRules(memoryPath string) ([]ActiveRule, error) {
	data, err := os.ReadFile(memoryPath)
	if err != nil {
		return nil, err
	}
	text := string(data)

	header := "## Reglas blandas"
	idx := strings.Index(text, header)
	if idx < 0 {
		return nil, nil
	}
	rest := text[idx+len(header):]
	if j := strings.Index(rest, "\n## "); j >= 0 {
		rest = rest[:j]
	}

	// Dedupe by (cat|hor|band). Last occurrence wins (softrules appends in
	// chronological order, so the latest counts are the most accurate).
	seen := map[string]ActiveRule{}
	for _, ln := range strings.Split(rest, "\n") {
		t := strings.TrimSpace(ln)
		m := ruleRE.FindStringSubmatch(t)
		if m == nil {
			continue
		}
		ar := ActiveRule{Category: m[1], Horizon: m[2], Band: m[3]}
		ar.Losses, _ = strconv.Atoi(m[4])
		ar.Wins, _ = strconv.Atoi(m[5])
		wr, _ := strconv.ParseFloat(m[6], 64)
		ar.WinRate = wr / 100.0
		key := ar.Category + "|" + ar.Horizon + "|" + ar.Band
		seen[key] = ar
	}
	out := make([]ActiveRule, 0, len(seen))
	for _, r := range seen {
		out = append(out, r)
	}
	return out, nil
}

// MatchVeto returns the first matching rule for a candidate, or nil.
func MatchVeto(rules []ActiveRule, category, horizon string, price float64) *ActiveRule {
	for i := range rules {
		if rules[i].Matches(category, horizon, price) {
			return &rules[i]
		}
	}
	return nil
}
