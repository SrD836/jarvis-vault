package calibration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShrinkClamp(t *testing.T) {
	cases := map[float64]float64{0: 1.0, -1: 1.0, 2: 1.0, 0.5: 0.5, 1.0: 1.0}
	for in, want := range cases {
		if got := (Calibration{KellyShrink: in}).Shrink(); got != want {
			t.Errorf("Shrink(%v)=%v want %v", in, got, want)
		}
	}
}

func TestShrinkProb(t *testing.T) {
	// implied 0.30, p 0.50, shrink 0.5 -> 0.30 + 0.5*0.20 = 0.40
	if got := (Calibration{KellyShrink: 0.5}).ShrinkProb(0.50, 0.30); got < 0.3999 || got > 0.4001 {
		t.Fatalf("shrink want 0.40 got %v", got)
	}
	// no calibration (zero value) is a no-op
	if got := (Calibration{}).ShrinkProb(0.50, 0.30); got != 0.50 {
		t.Fatalf("no-op want 0.50 got %v", got)
	}
}

func TestLoadFallbackAndParse(t *testing.T) {
	// missing file -> zero value, never errors
	t.Setenv("CALIBRATION_PATH", filepath.Join(t.TempDir(), "nope.json"))
	if c := Load(); c.KellyShrink != 0 || c.MinEdgePointsOverride != 0 {
		t.Fatalf("missing file must be zero value, got %+v", c)
	}
	// valid file -> parsed
	p := filepath.Join(t.TempDir(), "calibration.json")
	if err := os.WriteFile(p, []byte(`{"kelly_shrink":0.7,"min_edge_points_override":0.04}`), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CALIBRATION_PATH", p)
	c := Load()
	if c.Shrink() != 0.7 || c.MinEdgePointsOverride != 0.04 {
		t.Fatalf("parsed wrong: %+v", c)
	}
}
