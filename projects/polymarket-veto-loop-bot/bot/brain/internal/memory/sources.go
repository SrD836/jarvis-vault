package memory

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// SourceEvent is one append-only attribution row in source-stats.jsonl.
type SourceEvent struct {
	TS       string  `json:"ts"`
	TradeID  string  `json:"trade_id"`
	Domain   string  `json:"domain"`
	Outcome  string  `json:"outcome"` // "win" | "loss" | "breakeven"
	PnlUSD   float64 `json:"pnl_usd"`
	DaysOpen float64 `json:"days_open,omitempty"`
}

// SourceStats are the rolling aggregates computed from SourceEvent stream.
type SourceStats struct {
	Domain    string
	NTrades   int
	NWins     int
	NLosses   int
	NBreakEven int
	WinRate   float64 // wins / (wins + losses), excluding break-even from denominator
	TotalPnl  float64
	LastTS    time.Time
}

// LoadSourceStats reads source-stats.jsonl, aggregates by domain, considers only
// events within the rolling window (e.g. 30 days). Returns empty map if file missing.
func LoadSourceStats(path string, window time.Duration) (map[string]*SourceStats, error) {
	out := map[string]*SourceStats{}
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return out, nil
		}
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	cutoff := time.Now().UTC().Add(-window)
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var ev SourceEvent
		if err := json.Unmarshal(line, &ev); err != nil {
			continue
		}
		ts, err := time.Parse(time.RFC3339, ev.TS)
		if err != nil {
			continue
		}
		if window > 0 && ts.Before(cutoff) {
			continue
		}
		dom := strings.ToLower(strings.TrimSpace(ev.Domain))
		if dom == "" || dom == "legacy" {
			continue
		}
		st, ok := out[dom]
		if !ok {
			st = &SourceStats{Domain: dom}
			out[dom] = st
		}
		st.NTrades++
		st.TotalPnl += ev.PnlUSD
		switch ev.Outcome {
		case "win":
			st.NWins++
		case "loss":
			st.NLosses++
		case "breakeven":
			st.NBreakEven++
		}
		if ts.After(st.LastTS) {
			st.LastTS = ts
		}
	}
	for _, st := range out {
		denom := st.NWins + st.NLosses
		if denom > 0 {
			st.WinRate = float64(st.NWins) / float64(denom)
		}
	}
	return out, nil
}

// Blacklisted returns the domains that fail the quality bar:
//   NTrades >= minTrades  AND  WinRate < minWinRate
// The TTL is implicit in the window passed to LoadSourceStats.
func Blacklisted(stats map[string]*SourceStats, minTrades int, minWinRate float64) []string {
	var out []string
	for dom, st := range stats {
		if st.NTrades >= minTrades && st.WinRate < minWinRate {
			out = append(out, dom)
		}
	}
	sort.Strings(out)
	return out
}

// SortedByVolume returns the stats sorted by NTrades desc (for dashboard display).
func SortedByVolume(stats map[string]*SourceStats) []*SourceStats {
	out := make([]*SourceStats, 0, len(stats))
	for _, st := range stats {
		out = append(out, st)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].NTrades > out[j].NTrades })
	return out
}

// WriteSourceStatsSection rewrites the "## Sources stats" H2 in memory.md atomically.
// Leaves other H2 sections (Vetos pattern, Losses, Reglas, Human notes) untouched.
// blacklist is the precomputed set of blacklisted domains (for the YES/no column).
func WriteSourceStatsSection(path string, stats map[string]*SourceStats, blacklist []string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	bSet := map[string]bool{}
	for _, d := range blacklist {
		bSet[d] = true
	}

	var b strings.Builder
	b.WriteString("## Sources stats (rolling 30d, computed by brain)\n\n")
	b.WriteString("| domain | trades | wins | losses | win_rate | total_pnl | blacklisted |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, st := range SortedByVolume(stats) {
		bl := "no"
		if bSet[st.Domain] {
			bl = "YES"
		}
		b.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %.1f%% | %+.2f | %s |\n",
			st.Domain, st.NTrades, st.NWins, st.NLosses, st.WinRate*100, st.TotalPnl, bl))
	}

	text := string(data)
	header := "## Sources stats"
	startIdx := strings.Index(text, header)
	var out string
	if startIdx < 0 {
		// Append at end (preceded by newline if needed).
		sep := ""
		if !strings.HasSuffix(text, "\n\n") {
			if strings.HasSuffix(text, "\n") {
				sep = "\n"
			} else {
				sep = "\n\n"
			}
		}
		out = text + sep + b.String()
	} else {
		// Replace from startIdx until next H2 (## ) or EOF.
		rest := text[startIdx+len(header):]
		nextIdx := strings.Index(rest, "\n## ")
		var tail string
		if nextIdx < 0 {
			tail = ""
		} else {
			tail = rest[nextIdx:]
		}
		out = text[:startIdx] + b.String() + tail
		out = strings.TrimRight(out, "\n") + "\n"
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(out), 0644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
