// Package antipatterns — M3 rule consumer.
//
// Reads the "## Anti-patterns identificados" section from memory.md (populated
// asynchronously by exit_monitor's claudemax post-mortem). Each line lists
// tags extracted from a loss. When the same tag appears across enough losses
// (≥3) it becomes an active anti-pattern; new candidates whose slug or
// question contain that tag (case-insensitive substring) are vetoed by M3.
package antipatterns

import (
	"os"
	"regexp"
	"strings"
)

// Tag represents an anti-pattern tag with its accumulated frequency.
type Tag struct {
	Text  string
	Count int
}

// lineRE captures lines like:
//
//	- ceasefire-rumor · short-horizon · low-liquidity — visto en ... (2026-05-21, pnl $-40)
var lineRE = regexp.MustCompile(`^-\s+(.+?)\s+—\s+visto en`)

// LoadActive parses memory.md and returns tags with frequency >= minFreq.
// Returns nil on file-not-found (no anti-patterns yet is fine).
func LoadActive(memoryPath string, minFreq int) ([]Tag, error) {
	data, err := os.ReadFile(memoryPath)
	if err != nil {
		return nil, err
	}
	text := string(data)
	header := "## Anti-patterns identificados"
	idx := strings.Index(text, header)
	if idx < 0 {
		return nil, nil
	}
	rest := text[idx+len(header):]
	if j := strings.Index(rest, "\n## "); j >= 0 {
		rest = rest[:j]
	}
	counts := map[string]int{}
	for _, ln := range strings.Split(rest, "\n") {
		m := lineRE.FindStringSubmatch(strings.TrimSpace(ln))
		if m == nil {
			continue
		}
		for _, tag := range strings.Split(m[1], "·") {
			t := strings.ToLower(strings.TrimSpace(tag))
			if t == "" {
				continue
			}
			counts[t]++
		}
	}
	out := make([]Tag, 0, len(counts))
	for t, c := range counts {
		if c >= minFreq {
			out = append(out, Tag{Text: t, Count: c})
		}
	}
	return out, nil
}

// MatchVeto returns the first tag that matches the candidate text. Case-
// insensitive substring match against slug+question concatenated.
func MatchVeto(tags []Tag, slug, question string) *Tag {
	if len(tags) == 0 {
		return nil
	}
	hay := strings.ToLower(slug + " " + question)
	for i := range tags {
		// Tags are multi-word phrases; split on hyphen/space and require that
		// AT LEAST half of the constituent words appear in hay. This avoids
		// trivial single-common-word matches polluting M3.
		words := strings.FieldsFunc(tags[i].Text, func(r rune) bool { return r == '-' || r == ' ' })
		if len(words) == 0 {
			continue
		}
		hit := 0
		for _, w := range words {
			if len(w) >= 4 && strings.Contains(hay, w) {
				hit++
			}
		}
		if hit*2 >= len(words) {
			return &tags[i]
		}
	}
	return nil
}
