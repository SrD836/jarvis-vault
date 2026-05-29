package brain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/calibration"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/config"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/decisionlog"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/predictions"
	commontypes "github.com/davidgn/polymarket-veto-loop-bot/bot/common/types"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/antipatterns"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/llmclient"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/marketcheck"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/memory"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/softrules"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/research"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/rules"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/tradingview"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

var (
	candidatesPath    = envOr("BRAIN_CANDIDATES", "vault/inbox/trading/candidates.jsonl")
	approvedPath      = envOr("BRAIN_APPROVED", "vault/inbox/trading/approved.jsonl")
	blockedPath       = envOr("BRAIN_BLOCKED", "vault/inbox/trading/blocked.jsonl")
	memoryPath        = envOr("BRAIN_MEMORY", "vault/agents/polymarket-bot/memory.md")
	decisionsDir      = envOr("BRAIN_DECISIONS_DIR", "vault/03-decisions")
	researchCachePath = envOr("BRAIN_RESEARCH_CACHE", "vault/inbox/trading/research-cache.jsonl")
	tavilyConnPath    = envOr("BRAIN_TAVILY_CONNECTOR", "/openclaw-workspace/connectors/tavily.json")
	dashboardURL      = envOr("DASHBOARD_URL", "http://jarvis-dashboard:3000")
	sourceStatsPath   = envOr("BRAIN_SOURCE_STATS", "vault/agents/polymarket-bot/source-stats.jsonl")
	predictionsPath   = envOr("BRAIN_PREDICTIONS_PATH", predictions.DefaultPath)
	suspendedPath     = envOr("BRAIN_SUSPENDED_CATEGORIES", "vault/inbox/trading/suspended_categories.json")
	disableResearch   = envOr("BRAIN_DISABLE_RESEARCH", "") != ""
)

const (
	rulePrice          = "P0"
	rulePastEnd        = "P2"
	ruleWindow         = "P1"
	ruleMem            = "M1"
	ruleSoftLearned    = "M2"
	ruleAntiPattern    = "M3"
	ruleMarketCheck    = "P6"
	ruleTradingView    = "P11"
	ruleNews1          = "N1"
	ruleNews2          = "N2"
	ruleEdgeMissing    = "E1" // v7: LLM declared edge_type=none or empty
	ruleEdgeBelowMin   = "E2" // v7: |estimated_prob - implied| < MinEdgePoints
	ruleCategorySusp   = "S1" // v7: category Brier > 0.25 × 3w
	memoryHitThreshold = 0.7
	// maxResearchPerCycle caps Claude Max research calls to keep brain runtime
	// inside the 30-min cron window. Cache TTL is 6h so cumulative coverage
	// across cycles still reaches most candidates.
	maxResearchPerCycle = 30
	// v8 BOT-CAL: cap LLM (DeepSeek) evals per cycle. Mirrors maxResearchPerCycle.
	// Bounds DeepSeek cost/latency; candidates beyond the cap are deferred to the
	// next cycle (NOT blocked) so calibration data accrues steadily, not in a burst.
	maxLLMPerCycle = 80
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	log.Printf("Brain v3: config loaded mode=%s v2_under=%dh horizon=[short<=%dd, medium<=%dd, long>%dd] quota=%.0f/%.0f/%.0f",
		cfg.Mode, cfg.V2VetoUnderHours, cfg.HorizonShortMaxDays, cfg.HorizonMediumMaxDays, cfg.HorizonMediumMaxDays,
		cfg.HorizonQuotaShort*100, cfg.HorizonQuotaMedium*100, cfg.HorizonQuotaLong*100)

	// v8: adaptive layer written daily by analytics/daily_calibration.py. Absent
	// file → zero value → no edge override here (and shrink=1.0 in the executor).
	calib := calibration.Load()
	minEdge := cfg.MinEdgePoints
	if calib.MinEdgePointsOverride > minEdge {
		minEdge = calib.MinEdgePointsOverride
		log.Printf("E2: calibration raises min_edge_points %.3f -> %.3f (%s)", cfg.MinEdgePoints, minEdge, calib.Rationale)
	}

	if err := ensureDir(approvedPath); err != nil {
		return fmt.Errorf("ensure output dir: %w", err)
	}

	candidates, err := readCandidates(candidatesPath)
	if err != nil {
		return fmt.Errorf("read candidates: %w", err)
	}
	log.Printf("Brain v2: processing %d candidates", len(candidates))

	mem, err := memory.Load(memoryPath)
	if err != nil {
		log.Printf("WARN: memory load failed: %v (continuing without)", err)
		mem = nil
	} else if mem != nil {
		vetos, losses := mem.CountsSummary()
		log.Printf("Memory loaded: %d vetos, %d losses", vetos, losses)
	}

	var researcher *research.Client
	if !disableResearch {
		researcher, err = research.NewClient(tavilyConnPath, dashboardURL, researchCachePath)
		if err != nil {
			log.Printf("WARN: Tavily connector missing (%v) — falling back to Claude Max research (no external API key)", err)
			researcher = research.NewClaudemaxClient(dashboardURL, researchCachePath)
		}

	// v4: load per-source quality stats and feed the blacklist to the researcher.
	if researcher != nil {
		stats, sErr := memory.LoadSourceStats(sourceStatsPath, 30*24*time.Hour)
		if sErr != nil {
			log.Printf("WARN: load source stats: %v (continuing without blacklist)", sErr)
		} else {
			bl := memory.Blacklisted(stats, 5, 0.30)
			researcher.BlacklistedDomains = bl
			if len(bl) > 0 {
				log.Printf("research: blacklisted domains this cycle: %s", strings.Join(bl, ", "))
			}
			if wErr := memory.WriteSourceStatsSection(memoryPath, stats, bl); wErr != nil {
				log.Printf("WARN: write source stats section: %v", wErr)
			}
		}
	}
	}

	approved := make([]types.Approved, 0)
	blocked := make([]types.VetoResult, 0)
	now := time.Now().UTC().Format(time.RFC3339)
	_ = now

	stats := struct {
		P0, P1, P2, P3, P4, P6, P7, P8, P9, P10, P11, V1, V2, V4, M1, M2, M3, N1, N2, LLM, E1, E2, S1, LLMNoEdge, LLMParseFail int
	}{}

	// v7: load suspended categories. Brier-driven auto-suspend writes this file
	// when a category posts Brier > 0.25 for 3 consecutive weeks. Brain blocks
	// candidates in those categories with rule S1.
	suspendedCats := loadSuspendedCategories(suspendedPath)
	if len(suspendedCats) > 0 {
		log.Printf("S1: %d categories suspended by Brier review: %v", len(suspendedCats), suspendedCats)
	}

	// M3 prep: anti-patterns extracted from LLM post-mortems on past losses.
	// Tags appearing ≥3 times are treated as active veto signals.
	activeAntiPatterns, apErr := antipatterns.LoadActive(memoryPath, 3)
	if apErr != nil {
		log.Printf("WARN: load anti-patterns: %v (M3 disabled this cycle)", apErr)
	}
	log.Printf("M3: loaded %d active anti-patterns from memory.md", len(activeAntiPatterns))

	researchCalls := 0
	_ = researchCalls // referenced in loop below
	llmCalls := 0

	// M2 prep: load soft-learned veto rules from memory.md. These are the
	// auto-generated cluster rules (category·horizon·priceBand → win rate < 30%).
	// Without this consumer the rules were generated but never enforced.
	activeSoftRules, sErr := softrules.LoadActiveRules(memoryPath)
	if sErr != nil {
		log.Printf("WARN: load soft rules failed: %v (M2 disabled this cycle)", sErr)
	}
	log.Printf("M2: loaded %d active soft rules from memory.md", len(activeSoftRules))

	for i, c := range candidates {
		// S1: suspended-category gate (v7). Comes first because the rest of the
		// pipeline is pointless work for vetoed categories.
		if len(suspendedCats) > 0 {
			catKey := predictions.NormalizeCategory(c.Category)
			if _, susp := suspendedCats[catKey]; susp {
				vr := types.VetoResult{
					CandidateID: c.ID, Slug: c.Slug, Blocked: true,
					Reason:   fmt.Sprintf("categoría suspendida por Brier (cat=%s)", catKey),
					VetoedBy: ruleCategorySusp,
				}
				blocked = append(blocked, vr)
				recordVeto(mem, &c, vr, decisionsDir, true)
				appendSkipPrediction(&c, cfg, vr.Reason, ruleCategorySusp)
				stats.S1++
				continue
			}
		}

		// P0: price band sanity (v7 P4 — floor 0.05 default; 0.03 if horizon≤7d AND liquidity≥$20k).
		if vr := rules.EvalPriceBandV7(&c, cfg.PriceFloor, cfg.PriceCeiling, cfg.PriceFloorShortLiqRelax, cfg.PriceFloorShortLiqRelaxUSD); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false)
			stats.P0++
			continue
		}

		// P3: absolute liquidity floor (v6 — orderbook depth in USD, independent of trade size)
		if vr := rules.EvalAbsoluteLiquidity(&c, cfg.MinAbsoluteLiquidityUSD); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false)
			stats.P3++
			continue
		}

		// P4: pre-event veto (v6 — pre-IPO / pre-launch markets with >= N days to resolution)
		if vr := rules.EvalPreEventVeto(&c, cfg.PreEventVetoMinDays); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, true)
			stats.P4++
			continue
		}

		// P2: market already resolved (endDate in the past)
		if d := daysUntil(c.EndDate); d >= 0 && d <= 0 && c.EndDate != "" {
			// daysUntil returns -1 for empty/unknown — keep those; only veto when we KNOW endDate is past
		}
		if d := daysUntil(c.EndDate); d < 0 && c.EndDate != "" {
			vr := &types.VetoResult{
				CandidateID: c.ID,
				Slug:        c.Slug,
				Blocked:     true,
				Reason:      fmt.Sprintf("mercado ya expiró (endDate=%s, hace %d días)", c.EndDate, -d),
				VetoedBy:    rulePastEnd,
			}
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false)
			stats.P2++
			continue
		}

		// P7-P10: per-category hard rules (cheap pre-research filters).
		dToEnd := daysUntil(c.EndDate)
		if vr := rules.EvalWeatherNatural(&c); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false)
			stats.P7++
			continue
		}
		if vr := rules.EvalElectionFar(&c, dToEnd); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false)
			stats.P8++
			continue
		}
		if vr := rules.EvalGeopoliticsPump(&c, dToEnd); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, true)
			stats.P9++
			continue
		}
		if vr := rules.EvalSportsChalk(&c, dToEnd); vr != nil {
			blocked = append(blocked, *vr)
			recordVeto(mem, &c, *vr, decisionsDir, false)
			stats.P10++
			continue
		}

		// V1/V2/V4: numeric rules
		if vr := rules.EvaluateNumeric(&c, cfg.V2VetoUnderHours); vr != nil {
			blocked = append(blocked, *vr)
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

		// M1: memory pattern match
		if mem != nil {
			cl := memory.CandidateLike{
				Slug:            c.Slug,
				Category:        c.Category,
				CurrentPriceYes: c.CurrentPriceYes,
				Question:        c.Question,
			}
			// v8 BOT-CAL: in simulation/calibration mode, consult ONLY real loss
			// outcomes (MatchLosses) — matching the circular veto history starves
			// the pipeline. Other modes keep full veto+loss matching.
			var hits []memory.Match
			if cfg.Mode == "simulation" {
				hits = mem.MatchLosses(cl)
			} else {
				hits = mem.Match(cl)
			}
			if len(hits) > 0 && hits[0].Score >= memoryHitThreshold {
				vr := types.VetoResult{
					CandidateID: c.ID,
					Slug:        c.Slug,
					Blocked:     true,
					Reason:      fmt.Sprintf("memoria: %s (score %.2f)", hits[0].Why, hits[0].Score),
					VetoedBy:    ruleMem,
				}
				blocked = append(blocked, vr)
				memHits := make([]string, 0, len(hits))
				for _, h := range hits {
					memHits = append(memHits, fmt.Sprintf("%s `%s` score=%.2f (%s)", h.Pattern.Source, h.Pattern.Slug, h.Score, h.Why))
				}
				recordVetoWithExtras(mem, &c, vr, decisionsDir, true, "", memHits)
				stats.M1++
				continue
			}
		}

		// P6: market-asset reality check via Yahoo Finance. Vetoes BTC/ETH/WTI/SPX/etc.
		// candidates when the live spot price contradicts the implied yes-probability.
		if mr := marketcheck.Evaluate(c.Slug, c.Question, c.CurrentPriceYes); mr != nil && mr.Block {
			vr := types.VetoResult{
				CandidateID: c.ID, Slug: c.Slug, Blocked: true,
				Reason: mr.Reason, VetoedBy: ruleMarketCheck,
			}
			blocked = append(blocked, vr)
			recordVeto(mem, &c, vr, decisionsDir, true)
			stats.P6++
			continue
		}

		// P11: TradingView MCP cross-check for crypto candidates. Calls
		// tradingview-mcp coin_analysis and vetoes when the directional sentiment
		// strongly contradicts the implied yes-probability of the market.
		// Best-effort (network errors return no opinion).
		if tr := tradingview.Evaluate(c.Slug, c.Question, c.Category, c.CurrentPriceYes); tr != nil && tr.Block {
			vr := types.VetoResult{
				CandidateID: c.ID, Slug: c.Slug, Blocked: true,
				Reason: tr.Reason, VetoedBy: ruleTradingView,
			}
			blocked = append(blocked, vr)
			recordVeto(mem, &c, vr, decisionsDir, true)
			stats.P11++
			continue
		}

		// M2: soft-learned cluster veto. If the candidate's (category, horizon, price band)
		// has been losing systematically (< 30% win rate, >= 5 losses) the rule is enforced
		// here. Without this brain ignored its own learned patterns — see post-mortem in
		// agents/polymarket-bot/memory.md "## Reglas blandas aprendidas".
		if len(activeSoftRules) > 0 {
			d := daysUntil(c.EndDate)
			horizonGuess := cfg.Classify(d)
			if ar := softrules.MatchVeto(activeSoftRules, c.Category, horizonGuess, c.CurrentPriceYes); ar != nil {
				vr := types.VetoResult{
					CandidateID: c.ID,
					Slug:        c.Slug,
					Blocked:     true,
					Reason: fmt.Sprintf("M2 soft-learned: %s·%s·%s = %dL/%dW (wr %.0f%%)",
						ar.Category, ar.Horizon, ar.Band, ar.Losses, ar.Wins, ar.WinRate*100),
					VetoedBy: ruleSoftLearned,
				}
				blocked = append(blocked, vr)
				recordVeto(mem, &c, vr, decisionsDir, true)
				stats.M2++
				continue
			}
		}

		// M3: anti-pattern textual match (from LLM-extracted lessons in past losses).
		if len(activeAntiPatterns) > 0 {
			if ap := antipatterns.MatchVeto(activeAntiPatterns, c.Slug, c.Question); ap != nil {
				vr := types.VetoResult{
					CandidateID: c.ID, Slug: c.Slug, Blocked: true,
					Reason:   fmt.Sprintf("M3 anti-pattern: %q (visto %d veces)", ap.Text, ap.Count),
					VetoedBy: ruleAntiPattern,
				}
				blocked = append(blocked, vr)
				recordVeto(mem, &c, vr, decisionsDir, true)
				stats.M3++
				continue
			}
		}

		// N1/N2: news research. Rate-limited: candidates from the price floor up to
		// 0.80, capped per cycle to keep the cron tick under budget. v8 widened the
		// lower bound from 0.20 to cfg.PriceFloor so the longshots we actually trade
		// (<0.20) also get news attribution; above 0.80 the YES is too lopsided for
		// a single headline to matter. (Dormant while BRAIN_DISABLE_RESEARCH is set.)
		var researchSummary string
		var sourceCites []commontypes.SourceCite
		if researcher != nil && c.CurrentPriceYes >= cfg.PriceFloor && c.CurrentPriceYes <= 0.80 && researchCalls < maxResearchPerCycle {
			researchCalls++
			verdict, _, cites := researcher.EvaluateFull(c.Question, "yes")
			sourceCites = cites
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

		// v8 BOT-CAL: cap LLM (DeepSeek) evals/cycle. Candidates beyond the cap are
		// deferred to the next cycle (left as candidates, NOT blocked) so DeepSeek
		// cost/latency stay bounded while calibration data accrues steadily.
		if llmCalls >= maxLLMPerCycle {
			continue
		}
		llmCalls++

		// LLM: semantic V3/V5/V6 + v7 edge declaration.
		llmResult := llmclient.EvaluateV3V5V6(&c)
		if llmResult == nil {
			log.Printf("WARN: LLM returned nil for %s, passing", c.Slug)
			llmResult = &types.LLMBlockResult{}
		}
		if llmResult.Block {
			vr := types.VetoResult{
				CandidateID: c.ID,
				Slug:        c.Slug,
				Blocked:     true,
				Reason:      fmt.Sprintf("%s: %s", llmResult.Rule, llmResult.Reason),
				VetoedBy:    llmResult.Rule,
			}
			blocked = append(blocked, vr)
			recordVetoWithExtras(mem, &c, vr, decisionsDir, true, researchSummary, nil)
			appendDecisionPrediction(&c, cfg, llmResult, "skip", vr.Reason, 0)
			stats.LLM++
			continue
		}

		// E1: LLM did not declare a real edge. Without an edge we are not
		// allowed to enter — see prompt maestro principle #2.
		// v7 P3: track LLM parse failures (rule=PARSE) and edge=none separately
		// from generic E1 so we can tell apart "LLM broken" vs "LLM declined".
		edgeType := strings.ToLower(strings.TrimSpace(llmResult.EdgeType))
		if llmResult.Rule == "PARSE" {
			stats.LLMParseFail++
		} else if edgeType == "" || edgeType == "none" {
			stats.LLMNoEdge++
		}
		if edgeType == "" || edgeType == "none" {
			vr := types.VetoResult{
				CandidateID: c.ID, Slug: c.Slug, Blocked: true,
				Reason:   "edge no declarado por LLM (edge_type=none)",
				VetoedBy: ruleEdgeMissing,
			}
			blocked = append(blocked, vr)
			recordVetoWithExtras(mem, &c, vr, decisionsDir, true, researchSummary, nil)
			appendDecisionPrediction(&c, cfg, llmResult, "skip", vr.Reason, 0)
			stats.E1++
			continue
		}

		// E2: edge below configured minimum (default 5pp). Even if the LLM
		// declared an edge type, we require the magnitude to clear fees + noise.
		edgePts := math.Abs(llmResult.EstimatedProb - c.CurrentPriceYes)
		if edgePts < minEdge {
			vr := types.VetoResult{
				CandidateID: c.ID, Slug: c.Slug, Blocked: true,
				Reason:   fmt.Sprintf("edge %.3f < mín %.3f (p̂=%.3f, implied=%.3f)", edgePts, minEdge, llmResult.EstimatedProb, c.CurrentPriceYes),
				VetoedBy: ruleEdgeBelowMin,
			}
			blocked = append(blocked, vr)
			recordVetoWithExtras(mem, &c, vr, decisionsDir, true, researchSummary, nil)
			appendDecisionPrediction(&c, cfg, llmResult, "skip", vr.Reason, 0)
			stats.E2++
			continue
		}

		d := daysUntil(c.EndDate)
		horizon := cfg.Classify(d)
		side := "buy_yes"
		if llmResult.EstimatedProb < c.CurrentPriceYes {
			side = "buy_no"
		}
		approved = append(approved, types.Approved{
			CandidateID:        c.ID,
			Slug:               c.Slug,
			Question:           c.Question,
			Category:           c.Category,
			CurrentPriceYes:    c.CurrentPriceYes,
			Volume24h:          c.Volume24h,
			LiquidityUSD:       c.LiquidityUSD,
			EndDate:            c.EndDate,
			ScannedAt:          c.ScannedAt,
			ApprovedAt:         time.Now().UTC().Format(time.RFC3339),
			ApprovedPriceYes:   c.CurrentPriceYes,
			DaysToResolution:   d,
			Horizon:            horizon,
			SourcesUsed:        sourceCites,
			EstimatedProb:      llmResult.EstimatedProb,
			EdgeType:           edgeType,
			EdgeDescription:    llmResult.EdgeDescription,
			ThesisInvalidation: llmResult.ThesisInvalidation,
		})
		appendDecisionPrediction(&c, cfg, llmResult, side, "", 0)

		// Persistence of the trade .md happens in executor after actual entry,
		// not here — brain may re-approve the same slug across cycles until the
		// executor takes it, and a per-approve .md created one duplicate per day.
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

	softrules.GenerateAndAppend(memoryPath)
	log.Printf("Brain v7 done: %d approved, %d blocked (P0=%d P2=%d P3=%d P4=%d P6=%d P7=%d P8=%d P9=%d P10=%d P11=%d V1=%d V2=%d V4=%d M1=%d M2=%d M3=%d N1=%d N2=%d LLM=%d E1=%d E2=%d S1=%d LLM_NoEdge=%d LLM_ParseFail=%d) horizons in approved: short=%d medium=%d long=%d",
		len(approved), len(blocked),
		stats.P0, stats.P2, stats.P3, stats.P4, stats.P6, stats.P7, stats.P8, stats.P9, stats.P10, stats.P11, stats.V1, stats.V2, stats.V4, stats.M1, stats.M2, stats.M3, stats.N1, stats.N2, stats.LLM, stats.E1, stats.E2, stats.S1, stats.LLMNoEdge, stats.LLMParseFail,
		countHorizon(approved, "short"), countHorizon(approved, "medium"), countHorizon(approved, "long"))
	return nil
}

// loadSuspendedCategories reads the auto-suspend list emitted by analytics/brier.py.
// Schema: {"suspended": [{"category": "geopolitics", "since": "...", "until": "..."}]}
// or a flat list of strings. Missing file → empty map.
func loadSuspendedCategories(path string) map[string]struct{} {
	out := map[string]struct{}{}
	raw, err := os.ReadFile(path)
	if err != nil {
		return out
	}
	var rich struct {
		Suspended []struct {
			Category string `json:"category"`
		} `json:"suspended"`
	}
	if err := json.Unmarshal(raw, &rich); err == nil && len(rich.Suspended) > 0 {
		for _, s := range rich.Suspended {
			c := predictions.NormalizeCategory(s.Category)
			if c != "" {
				out[c] = struct{}{}
			}
		}
		return out
	}
	var flat []string
	if err := json.Unmarshal(raw, &flat); err == nil {
		for _, s := range flat {
			out[predictions.NormalizeCategory(s)] = struct{}{}
		}
	}
	return out
}

// appendSkipPrediction logs an S1-style veto where the LLM never ran.
func appendSkipPrediction(c *types.Candidate, cfg *config.BotConfig, reason, rule string) {
	d := daysUntil(c.EndDate)
	p := predictions.Prediction{
		MarketID:     c.ID,
		Slug:         c.Slug,
		Question:     c.Question,
		Category:     predictions.NormalizeCategory(c.Category),
		HorizonDays:  d,
		Horizon:      cfg.Classify(d),
		ImpliedPrice: c.CurrentPriceYes,
		Decision:     "skip",
		SkipReason:   rule + ": " + reason,
	}
	if err := predictions.Append(predictionsPath, p); err != nil {
		log.Printf("WARN: predictions append (skip %s): %v", rule, err)
	}
}

// appendDecisionPrediction logs a post-LLM decision (veto or approve). The
// estimated_prob / edge_type fields are carried from the LLM result so the
// Brier loop has data to score once the market resolves.
func appendDecisionPrediction(c *types.Candidate, cfg *config.BotConfig, lr *types.LLMBlockResult, decision, skipReason string, sizeUSD float64) {
	d := daysUntil(c.EndDate)
	p := predictions.Prediction{
		MarketID:           c.ID,
		Slug:               c.Slug,
		Question:           c.Question,
		Category:           predictions.NormalizeCategory(c.Category),
		HorizonDays:        d,
		Horizon:            cfg.Classify(d),
		ImpliedPrice:       c.CurrentPriceYes,
		EstimatedProb:      lr.EstimatedProb,
		EdgePoints:         math.Abs(lr.EstimatedProb - c.CurrentPriceYes),
		EdgeType:           strings.ToLower(strings.TrimSpace(lr.EdgeType)),
		EdgeDescription:    lr.EdgeDescription,
		ChecklistPassed:    decision != "skip",
		Decision:           decision,
		SkipReason:         skipReason,
		SizeUSD:            sizeUSD,
		ThesisInvalidation: lr.ThesisInvalidation,
	}
	if err := predictions.Append(predictionsPath, p); err != nil {
		log.Printf("WARN: predictions append (%s): %v", decision, err)
	}
}

func evalWindowGate(c *types.Candidate, minDays, maxDays int) *types.VetoResult {
	days := daysUntil(c.EndDate)
	if days < minDays || days > maxDays {
		return &types.VetoResult{
			CandidateID: c.ID,
			Slug:        c.Slug,
			Blocked:     true,
			Reason:      fmt.Sprintf("P1: fuera de ventana sim corto plazo (%d días, ventana [%d, %d])", days, minDays, maxDays),
			VetoedBy:    ruleWindow,
		}
	}
	return nil
}

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
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
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
	return out, sc.Err()
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

func countHorizon(approved []types.Approved, h string) int {
	n := 0
	for _, a := range approved {
		if a.Horizon == h {
			n++
		}
	}
	return n
}

func envOr(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
