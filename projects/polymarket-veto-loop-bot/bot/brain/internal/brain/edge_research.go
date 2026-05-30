package brain

import (
	"fmt"
	"math"
)

// EdgeResearchDecision is the outcome of the G2 post-veto research gate.
//
// The gate runs ONLY on near-approved candidates (those that already passed
// every cheap veto + the LLM edge declaration + R1/R3/R5). It turns fresh news
// into either a block or an informed edge, so the bot stops trading DeepSeek's
// uninformed deviation from the market price (structurally -EV) and only enters
// when news supports the side it would take.
type EdgeResearchDecision struct {
	Block        bool
	Rule         string  // "N1" (news against) | "N2" (silent) when Block
	Reason       string  // human-readable veto reason; empty when approved
	AdjustedProb float64 // estimated_prob after the silent-shrink (== input when not shrunk)
}

// DecideEdgeResearch applies the G2 gate to one near-approved candidate.
//
// The verdict is pre-normalized to the SIDE the bot would take:
//   - against: news points AGAINST our side (yes-bet + contradicts, or
//     no-bet + confirms). The thesis is freshly refuted -> block N1.
//   - silent: no fresh media support for our side. If a public catalyst is
//     imminent (0 < daysToEnd < 30) the trade is a coin-flip -> block N2.
//     Otherwise the LLM's prior has no fresh backing, so we shrink the edge
//     toward the market price by shrinkSilent and re-apply the E2 minimum; an
//     edge that only existed in the model's head dies here.
//   - neither against nor silent (news supports our side) -> keep the estimate.
//
// shrinkSilent in (0,1]: 1.0 keeps the full edge on silence (no shrink),
// smaller values are more conservative. The shrink moves the probability toward
// impliedPrice and never crosses it, so the trade side is preserved.
func DecideEdgeResearch(against, silent bool, estimatedProb, impliedPrice, minEdge, shrinkSilent float64, daysToEnd int, summary string) EdgeResearchDecision {
	if against {
		return EdgeResearchDecision{
			Block: true, Rule: "N1",
			Reason:       "noticias contradicen tesis: " + summary,
			AdjustedProb: estimatedProb,
		}
	}
	if silent {
		if daysToEnd > 0 && daysToEnd < 30 {
			return EdgeResearchDecision{
				Block: true, Rule: "N2",
				Reason:       fmt.Sprintf("silencio mediático sobre catalyst inminente (%d días al cierre)", daysToEnd),
				AdjustedProb: estimatedProb,
			}
		}
		adj := impliedPrice + shrinkSilent*(estimatedProb-impliedPrice)
		if math.Abs(adj-impliedPrice) < minEdge {
			return EdgeResearchDecision{
				Block: true, Rule: "N2",
				Reason:       fmt.Sprintf("edge sin soporte de noticias %.3f<%.3f tras shrink ×%.2f", math.Abs(adj-impliedPrice), minEdge, shrinkSilent),
				AdjustedProb: adj,
			}
		}
		return EdgeResearchDecision{Block: false, AdjustedProb: adj}
	}
	// News supports our side: keep the LLM estimate untouched.
	return EdgeResearchDecision{Block: false, AdjustedProb: estimatedProb}
}
