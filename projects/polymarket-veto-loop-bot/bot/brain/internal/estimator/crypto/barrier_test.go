package crypto

import (
	"math"
	"testing"
)

func TestProb(t *testing.T) {
	// At-the-money terminal: roughly a coin flip.
	if p := Prob(100, 100, 0.5, 0.1, "above", "terminal"); p < 0.40 || p > 0.55 {
		t.Errorf("ATM terminal above = %.3f, want ~0.5", p)
	}

	// Deep in-the-money terminal above: near-certain → clamped to 0.99.
	if p := Prob(200, 100, 0.5, 0.1, "above", "terminal"); p < 0.95 {
		t.Errorf("deep-ITM terminal above = %.3f, want ~0.99", p)
	}
	// Deep out-of-the-money terminal above: near-zero → clamped to 0.01.
	if p := Prob(50, 100, 0.5, 0.1, "above", "terminal"); p > 0.05 {
		t.Errorf("deep-OTM terminal above = %.3f, want ~0.01", p)
	}

	// Touch probability >= terminal probability for the same out-of-money K>S
	// (you can hit the barrier intraday and come back).
	term := Prob(100, 110, 0.5, 0.1, "above", "terminal")
	touch := Prob(100, 110, 0.5, 0.1, "above", "touch")
	if touch < term {
		t.Errorf("touch %.3f should be >= terminal %.3f", touch, term)
	}

	// below is the symmetric complement of above (terminal).
	above := Prob(100, 90, 0.5, 0.1, "above", "terminal")
	below := Prob(100, 90, 0.5, 0.1, "below", "terminal")
	if math.Abs((above+below)-1.0) > 1e-9 {
		t.Errorf("terminal above+below = %.4f, want 1.0", above+below)
	}

	// Already through the barrier (spot past target) on a touch market → ~certain.
	if p := Prob(120, 100, 0.5, 0.1, "above", "touch"); p < 0.95 {
		t.Errorf("already-touched above = %.3f, want 0.99", p)
	}
	if p := Prob(80, 100, 0.5, 0.1, "below", "touch"); p < 0.95 {
		t.Errorf("already-touched below = %.3f, want 0.99", p)
	}

	// Degenerate inputs → no opinion (0.5), never NaN/panic.
	if p := Prob(100, 100, 0, 0.1, "above", "terminal"); p != 0.5 {
		t.Errorf("zero sigma = %.3f, want 0.5", p)
	}
	if p := Prob(100, 100, 0.5, 0, "above", "terminal"); p != 0.5 {
		t.Errorf("zero tau = %.3f, want 0.5", p)
	}
}
