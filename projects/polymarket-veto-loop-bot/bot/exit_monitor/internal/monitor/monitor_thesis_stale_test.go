package monitor

import "testing"

func TestIsThesisStale(t *testing.T) {
	const staleDays, margin = 1.0, 0.03
	cases := []struct {
		name                          string
		side                          string
		price, entry, est, daysLeft   float64
		want                          bool
	}{
		// longshot cerca de expiry, no convergió, sigue <= entrada -> stale
		{"yes longshot expiry no progreso", "yes", 0.01, 0.025, 0.08, 0.5, true},
		// con tiempo por delante -> NO (respeta no-whipsaw)
		{"yes tiempo restante", "yes", 0.01, 0.025, 0.08, 1.5, false},
		// progresó al alza (price > entry) -> NO
		{"yes progreso al alza", "yes", 0.06, 0.025, 0.08, 0.5, false},
		// bajó algo pero pegado al estimado (dentro del margen) -> NO (ruido)
		{"yes pegado al estimado", "yes", 0.06, 0.07, 0.08, 0.5, false},
		// lado NO simétrico: precio subió contra la tesis cerca de expiry -> stale
		{"no contra tesis expiry", "no", 0.95, 0.90, 0.85, 0.5, true},
		{"no progreso a la baja", "no", 0.80, 0.90, 0.85, 0.5, false},
		// est=0 (sin estimado) -> NO
		{"sin estimado", "yes", 0.01, 0.025, 0.0, 0.5, false},
		// daysLeft negativo (endDate no parseable) -> NO
		{"daysLeft negativo", "yes", 0.01, 0.025, 0.08, -1, false},
		// justo en el borde daysLeft == staleDays -> NO (estricto <)
		{"borde daysLeft==staleDays", "yes", 0.01, 0.025, 0.08, 1.0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := IsThesisStale(c.side, c.price, c.entry, c.est, c.daysLeft, staleDays, margin)
			if got != c.want {
				t.Fatalf("IsThesisStale(%s, price=%.3f entry=%.3f est=%.3f dl=%.2f)=%v, want %v",
					c.side, c.price, c.entry, c.est, c.daysLeft, got, c.want)
			}
		})
	}
}
