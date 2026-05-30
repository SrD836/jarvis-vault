package rules

import (
	"testing"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

func timeFromNowHours(h int) string {
	return time.Now().UTC().Add(time.Duration(h) * time.Hour).Format(time.RFC3339)
}

// R1 longshot: bloquea precio bajo salvo edge info/arb fuerte.
func TestEvalLongshotEdgeGate(t *testing.T) {
	cases := []struct {
		name      string
		price     float64
		edgeType  string
		estProb   float64
		wantBlock bool
	}{
		{"longshot sin edge -> block", 0.03, "none", 0.04, true},
		{"longshot edge especulativo -> block", 0.03, "calibration", 0.30, true},
		{"longshot edge info debil -> block", 0.05, "info", 0.12, true},        // edge 0.07 < 0.15
		{"longshot edge info fuerte -> pasa", 0.05, "info", 0.25, false},       // edge 0.20 >= 0.15
		{"longshot edge arb fuerte -> pasa", 0.04, "arb", 0.22, false},         // edge 0.18 >= 0.15
		{"precio normal -> pasa (regla no aplica)", 0.20, "none", 0.05, false}, // >= 0.10
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cand := &types.Candidate{ID: "x", Slug: "s", CurrentPriceYes: c.price}
			vr := EvalLongshotEdgeGate(cand, c.edgeType, c.estProb, 0.10, 0.15)
			if (vr != nil) != c.wantBlock {
				t.Fatalf("price=%.2f edge=%s est=%.2f: block=%v, want %v (vr=%v)", c.price, c.edgeType, c.estProb, vr != nil, c.wantBlock, vr)
			}
			if vr != nil && vr.VetoedBy != "R1_longshot" {
				t.Fatalf("VetoedBy=%q, want R1_longshot", vr.VetoedBy)
			}
		})
	}
}

// R3 catalyst: bloquea evento <Nh salvo edge info contemplado.
func TestEvalCatalyst24h(t *testing.T) {
	soon := timeFromNowHours(12)
	far := timeFromNowHours(72)
	cases := []struct {
		name      string
		end       string
		edgeType  string
		thesisInv string
		wantBlock bool
	}{
		{"catalyst 12h sin edge -> block", soon, "none", "", true},
		{"catalyst 12h edge especulativo -> block", soon, "calibration", "algo", true},
		{"catalyst 12h info sin invalidacion -> block", soon, "info", "", true},
		{"catalyst 12h info contemplado -> pasa", soon, "info", "si X pasa, cerrar", false},
		{"catalyst 72h -> pasa (fuera de ventana)", far, "none", "", false},
		{"sin endDate -> pasa", "", "none", "", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cand := &types.Candidate{ID: "x", Slug: "s", EndDate: c.end}
			vr := EvalCatalyst24h(cand, c.edgeType, c.thesisInv, 24)
			if (vr != nil) != c.wantBlock {
				t.Fatalf("end=%s edge=%s inv=%q: block=%v, want %v", c.end, c.edgeType, c.thesisInv, vr != nil, c.wantBlock)
			}
		})
	}
}

// R4: VetoV1 con umbral configurable.
func TestVetoV1Configurable(t *testing.T) {
	cand := &types.Candidate{ID: "x", Slug: "s", Volume24h: 30000}
	if vr := VetoV1(cand, 50000); vr == nil {
		t.Fatal("vol 30k < 50k debe bloquear")
	}
	if vr := VetoV1(cand, 20000); vr != nil {
		t.Fatal("vol 30k >= 20k no debe bloquear")
	}
}
