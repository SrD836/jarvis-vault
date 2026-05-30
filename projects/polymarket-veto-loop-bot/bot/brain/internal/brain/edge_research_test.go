package brain

import (
	"math"
	"testing"
)

func TestDecideEdgeResearch(t *testing.T) {
	const minEdge = 0.03
	const shrink = 0.5

	cases := []struct {
		name        string
		against     bool
		silent      bool
		est         float64
		implied     float64
		daysToEnd   int
		wantBlock   bool
		wantRule    string
		wantAdjProb float64
	}{
		{
			name:    "news against our side -> block N1",
			against: true, silent: false,
			est: 0.55, implied: 0.40, daysToEnd: 60,
			wantBlock: true, wantRule: "N1", wantAdjProb: 0.55,
		},
		{
			name:    "silent + imminent catalyst -> block N2",
			against: false, silent: true,
			est: 0.55, implied: 0.40, daysToEnd: 5,
			wantBlock: true, wantRule: "N2", wantAdjProb: 0.55,
		},
		{
			name:    "silent + far + edge survives shrink -> approve with shrunk prob",
			against: false, silent: true,
			est: 0.50, implied: 0.40, daysToEnd: 60,
			wantBlock: false, wantRule: "", wantAdjProb: 0.45,
		},
		{
			name:    "silent + far + edge dies after shrink -> block N2",
			against: false, silent: true,
			est: 0.45, implied: 0.40, daysToEnd: 60,
			wantBlock: true, wantRule: "N2", wantAdjProb: 0.425,
		},
		{
			name:    "silent + unknown end date -> shrink path (not imminent)",
			against: false, silent: true,
			est: 0.60, implied: 0.40, daysToEnd: -1,
			wantBlock: false, wantRule: "", wantAdjProb: 0.50,
		},
		{
			name:    "news supports our side -> approve, prob untouched",
			against: false, silent: false,
			est: 0.50, implied: 0.40, daysToEnd: 60,
			wantBlock: false, wantRule: "", wantAdjProb: 0.50,
		},
		{
			name:    "buy_no side: shrink moves prob up toward implied, stays below it",
			against: false, silent: true,
			est: 0.20, implied: 0.40, daysToEnd: 60,
			wantBlock: false, wantRule: "", wantAdjProb: 0.30,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := DecideEdgeResearch(tc.against, tc.silent, tc.est, tc.implied, minEdge, shrink, tc.daysToEnd, "summary")
			if got.Block != tc.wantBlock {
				t.Errorf("Block = %v, want %v", got.Block, tc.wantBlock)
			}
			if got.Block && got.Rule != tc.wantRule {
				t.Errorf("Rule = %q, want %q", got.Rule, tc.wantRule)
			}
			if math.Abs(got.AdjustedProb-tc.wantAdjProb) > 1e-9 {
				t.Errorf("AdjustedProb = %.4f, want %.4f", got.AdjustedProb, tc.wantAdjProb)
			}
			if got.Block && got.Reason == "" {
				t.Error("blocked decision must carry a reason")
			}
		})
	}
}
