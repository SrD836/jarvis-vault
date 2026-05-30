package sports

import (
	"math"
	"testing"
)

func TestDevig(t *testing.T) {
	// Two-outcome h2h: Pinnacle 1.69 / 2.31 (≈2.5% overround).
	p := Devig([]float64{1.69, 2.31})
	if p == nil {
		t.Fatal("Devig returned nil for valid odds")
	}
	if math.Abs((p[0]+p[1])-1.0) > 1e-9 {
		t.Errorf("devigged probs sum to %.6f, want 1.0", p[0]+p[1])
	}
	// 1/1.69=0.5917, 1/2.31=0.4329, sum=1.0246 → p0=0.5775.
	if math.Abs(p[0]-0.5775) > 0.001 {
		t.Errorf("p0 = %.4f, want ~0.5775", p[0])
	}

	// Multi-outcome futures field sums to 1.0 after devig.
	field := []float64{2.5, 3.0, 4.0, 8.0, 15.0}
	dp := Devig(field)
	var s float64
	for _, x := range dp {
		s += x
	}
	if math.Abs(s-1.0) > 1e-9 {
		t.Errorf("futures field devig sums to %.6f, want 1.0", s)
	}
	// Favorite (lowest odds) has the highest probability.
	if dp[0] <= dp[1] || dp[1] <= dp[2] {
		t.Errorf("devigged probs not monotonic with odds: %v", dp)
	}

	// Invalid inputs → nil.
	if Devig(nil) != nil {
		t.Error("empty field should return nil")
	}
	if Devig([]float64{1.0, 2.0}) != nil {
		t.Error("odds <= 1.0 should return nil")
	}

	// TrueProb indexing.
	if v, ok := TrueProb([]float64{1.69, 2.31}, 0); !ok || math.Abs(v-0.5775) > 0.001 {
		t.Errorf("TrueProb[0] = %.4f ok=%v, want ~0.5775 true", v, ok)
	}
	if _, ok := TrueProb([]float64{1.69, 2.31}, 5); ok {
		t.Error("out-of-range target should return ok=false")
	}
}
