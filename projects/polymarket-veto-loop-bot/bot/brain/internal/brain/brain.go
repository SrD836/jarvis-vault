package brain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/decisionlog"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/llmclient"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/memory"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/research"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/rules"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

// Paths are resolved relative to working directory (typically /home/agent/agent-stack).
// All vault paths are relative — caller sets cwd accordingly.
var (
	candidatesPath     = envOr("BRAIN_CANDIDATES", "vault/inbox/trading/candidates.jsonl")
	approvedPath       = envOr("BRAIN_APPROVED", "vault/inbox/trading/approved.jsonl")
	blockedPath        = envOr("BRAIN_BLOCKED", "vault/inbox/trading/blocked.jsonl")
	memoryPath         = envOr("BRAIN_MEMORY", "vault/agents/polymarket-bot/memory.md")
	decisionsDir       = envOr("BRAIN_DECISIONS_DIR", "vault/03-decisions")
	researchCachePath  = envOr("BRAIN_RESEARCH_CACHE", "vault/inbox/trading/research-cache.jsonl")
	tavilyConnPath     = envOr("BRAIN_TAVILY_CONNECTOR", "/openclaw-workspace/connectors/tavily.json")
	dashboardURL       = envOr("DASHBOARD_URL", "http://jarvis-dashboard:3000")
	disableResearch    = envOr("BRAIN_DISABLE_RESEARCH", "") != ""
)

// Veto rule constants for the new v2 stages.
const (
	rulePrice = "P0" // price out of [0.05, 0.95]
	ruleMem   = "M1" // memory pattern match
	ruleNews1 = "N1" // news contradicts
	ruleNews2 = "N2" // news silent on imminent catalyst
	memoryHitThreshold = 0.7
)

// Run reads candidates from JSONL, vets them, writes approved/blocked, and side-effects
// to vault: memory.md (append) + 03-decisions/ (one note per significant veto).
func Run() error {
	if err := ensureDir(approvedPath); err != nil {
		return fmt.Errorf("ensure output dir: %w", err)
	}

	candidates, err := readCandidates(candidatesPath)
	if err != nil {
		return fmt.Errorf("read candidates: %w", err)
	}
	log.Printf("Brain v2: processing %d candidates", len(candidates))

	// Load memory (best-effort: missing file means no patterns yet)
	mem, err := memory.Load(memoryPath)
	if err != nil {
		log.Printf("WARN: memory load failed: %v (continuing without)", err)
		mem = nil
	} else if mem != nil {
		vetos, losses := mem.CountsSummary()
		log.Printf("Memory loaded: %d vetos, %d losses", vetos, losses)
	}

	// Load research client (best-effort: missing connector means no news veto)
	var researcher *research.Client
	if !disableResearch {
		researcher, err = research.NewClient(tavilyConnPath, dashboardURL, researchCachePath)
		if err != nil {
			log.Printf("WARN: research client failed (%v) — news vetoes disabled", err)
			researcher = nil
		} else {
			log.Printf("Research client ready (Tavily + %s)", dashboardURL)
		}
	}

	approved := make([]types.Approved, 0)
	blocked := make([]types.VetoResult, 0)
	now := time.Now().UTC().Format(time.RFC3339)
	_ = now

	stats := struct {
		P0, V1, V2, V4, M1, N1, N2, LLM int
	}{}

	for i, c := range candidates {
		// stage 0: price band sanity (long-tail extremo) — trivial veto, NO decision log
		if vr := evalPriceBand(&c); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false) // memory yes, decision no
			stats.P0++
			continue
		}

		// stage 1: numeric V1/V2/V4
		if vr := rules.EvaluateNumeric(&c); vr != nil {
			blocked = append(blocked, *vr)
			// V1 (volumen) y V4 (chasing 8%) son triviales — no decision log
			// V2 (catalyst inminente) sí es significativo
			significant := vr.VetoedBy == "V2"
			recordVeto(mem, &c, *vr, decisionsDir, significant)
			switch vr.VetoedBy {
			case "V1":
				stats.V1++
			case "V2":
				stats.V2++
			case "V4":
				stats.V4++
			}
			continue
		}

		// stage 2: memory pattern match
		if mem != nil {
			cl := memory.CandidateLike{
				Slug:            c.Slug,
				Category:        c.Category,
				CurrentPriceYes: c.CurrentPriceYes,
				Question:        c.Question,
			}
			hits := mem.Match(cl)
			if len(hits) > 0 && hits[0].Score >= memoryHitThreshold {
				vr := types.VetoResult{
					CandidateID: c.ID,
					Slug:        c.Slug,
					Blocked:     true,
					Reason:      fmt.Sprintf("memoria: %s (score %.2f)", hits[0].Why, hits[0].Score),
					VetoedBy:    ruleMem,
				}
				blocked = append(blocked, vr)
				// build memHits strings for decision log
				memHits := make([]string, 0, len(hits))
				for _, h := range hits {
					memHits = append(memHits, fmt.Sprintf("%s `%s` score=%.2f (%s)", h.Pattern.Source, h.Pattern.Slug, h.Score, h.Why))
				}
				recordVetoWithExtras(mem, &c, vr, decisionsDir, true, "", memHits)
				stats.M1++
				continue
			}
		}

		// stage 3: news research
		var researchSummary string
		if researcher != nil {
			verdict := researcher.Evaluate(c.Question, "yes")
			researchSummary = fmt.Sprintf("Tavily+DeepSeek: confirms=%v contradicts=%v silent=%v score=%.2f — %s",
				verdict.Confirms, verdict.Contradicts, verdict.Silent, verdict.Score, verdict.Summary)
			daysToEnd := daysUntil(c.EndDate)
			if verdict.Contradicts {
				vr := types.VetoResult{
					CandidateID: c.ID,
					Slug:        c.Slug,
					Blocked:     true,
					Reason:      fmt.Sprintf("noticias contradicen tesis: %s", verdict.Summary),
					VetoedBy:    ruleNews1,
				}
				blocked = append(blocked, vr)
				recordVetoWithExtras(mem, &c, vr, decisionsDir, true, researchSummary, nil)
				stats.N1++
				continue
			}
			if verdict.Silent && daysToEnd > 0 && daysToEnd < 30 {
				vr := types.VetoResult{
					CandidateID: c.ID,
					Slug:        c.Slug,
					Blocked:     true,
					Reason:      fmt.Sprintf("silencio mediático sobre catalyst inminente (%d días al cierre)", daysToEnd),
					VetoedBy:    ruleNews2,
				}
				blocked = append(blocked, vr)
				recordVetoWithExtras(mem, &c, vr, decisionsDir, true, researchSummary, nil)
				stats.N2++
				continue
			}
		}

		// stage 4: LLM semantic V3/V5/V6 with research context
		llmResult := llmclient.EvaluateV3V5V6(&c)
		if llmResult == nil {
			log.Printf("WARN: LLM returned nil for %s, passing", c.Slug)
		} else if llmResult.Block {
			vr := types.VetoResult{
				CandidateID: c.ID,
				Slug:        c.Slug,
				Blocked:     true,
				Reason:      fmt.Sprintf("%s: %s", llmResult.Rule, llmResult.Reason),
				VetoedBy:    llmResult.Rule,
			}
			blocked = append(blocked, vr)
			recordVetoWithExtras(mem, &c, vr, decisionsDir, true, researchSummary, nil)
			stats.LLM++
			continue
		}

		// passed all
		approved = append(approved, types.Approved{
			CandidateID:      c.ID,
			Slug:             c.Slug,
			Question:         c.Question,
			Category:         c.Category,
			CurrentPriceYes:  c.CurrentPriceYes,
			Volume24h:        c.Volume24h,
			EndDate:          c.EndDate,
			ScannedAt:        c.ScannedAt,
			ApprovedAt:       time.Now().UTC().Format(time.RFC3339),
			ApprovedPriceYes: c.CurrentPriceYes,
		})
		if i%50 == 0 {
			log.Printf("Brain progress: %d/%d processed, %d approved, %d blocked",
				i+1, len(candidates), len(approved), len(blocked))
		}
	}

	if err := writeJSONL(approvedPath, approved); err != nil {
		return fmt.Errorf("write approved: %w", err)
	}
	if err := writeJSONL(blockedPath, blocked); err != nil {
		return fmt.Errorf("write blocked: %w", err)
	}

	log.Printf("Brain v2 done: %d approved, %d blocked (P0=%d V1=%d V2=%d V4=%d M1=%d N1=%d N2=%d LLM=%d)",
		len(approved), len(blocked),
		stats.P0, stats.V1, stats.V2, stats.V4, stats.M1, stats.N1, stats.N2, stats.LLM)
	return nil
}

func evalPriceBand(c *types.Candidate) *types.VetoResult {
	if c.CurrentPriceYes < 0.05 || c.CurrentPriceYes > 0.95 {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("precio fuera de [0.05, 0.95]: %.4f (long-tail no rentable en sim)", c.CurrentPriceYes),
			VetoedBy:    rulePrice,
		}
	}
	return nil
}

// recordVeto appends to memory.md (always if mem!=nil) and writes 03-decisions only if significant.
func recordVeto(mem *memory.Memory, c *types.Candidate, vr types.VetoResult, decisionsDir string, significant bool) {
	recordVetoWithExtras(mem, c, vr, decisionsDir, significant, "", nil)
}

func recordVetoWithExtras(mem *memory.Memory, c *types.Candidate, vr types.VetoResult, decisionsDir string, significant bool, researchSummary string, memHits []string) {
	if mem != nil {
		if err := mem.AppendVeto(c.Slug, c.Category, vr.VetoedBy, vr.Reason, c.CurrentPriceYes); err != nil {
			log.Printf("WARN: memory append veto failed: %v", err)
		}
	}
	if !significant {
		return
	}
	d := decisionlog.VetoDecision{
		Slug:            c.Slug,
		Question:        c.Question,
		Category:        c.Category,
		PriceYes:        c.CurrentPriceYes,
		Volume24h:       c.Volume24h,
		EndDate:         c.EndDate,
		Rule:            vr.VetoedBy,
		Reason:          vr.Reason,
		ResearchSummary: researchSummary,
		MemoryHits:      memHits,
	}
	path, err := decisionlog.WriteVeto(decisionsDir, d)
	if err != nil {
		log.Printf("WARN: decision log write failed for %s: %v", c.Slug, err)
		return
	}
	if !strings.Contains(path, "/") {
		path = filepath.Join(decisionsDir, path)
	}
}

func daysUntil(endDate string) int {
	if endDate == "" {
		return -1
	}
	t, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		// try other formats
		t, err = time.Parse("2006-01-02T15:04:05Z", endDate)
		if err != nil {
			return -1
		}
	}
	return int(time.Until(t).Hours() / 24)
}

func readCandidates(path string) ([]types.Candidate, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()
	var out []types.Candidate
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var c types.Candidate
		if err := json.Unmarshal([]byte(line), &c); err != nil {
			log.Printf("WARN: skip unparseable candidate line: %v", err)
			continue
		}
		out = append(out, c)
	}
	return out, scanner.Err()
}

func writeJSONL(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	switch v := data.(type) {
	case []types.Approved:
		for _, item := range v {
			if err := enc.Encode(item); err != nil {
				return err
			}
		}
	case []types.VetoResult:
		for _, item := range v {
			if err := enc.Encode(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func ensureDir(p string) error {
	return os.MkdirAll(filepath.Dir(p), 0755)
}

func envOr(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
