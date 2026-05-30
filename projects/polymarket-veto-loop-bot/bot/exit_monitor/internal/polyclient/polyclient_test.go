package polyclient

import "testing"

// ExitPrice: el guardarrail del fix 2026-05-30. Un cierre nunca se ejecuta a un
// fallback no validado; distingue resolucion real (settlement 0/1) de fallo de fuente.
func TestExitPrice(t *testing.T) {
	cases := []struct {
		name      string
		q         MarketQuote
		side      string
		ref       float64
		wantPrice float64
		wantSrc   string
		wantOK    bool
		wantReslv bool
	}{
		{
			name:    "libro vacio, no resuelto -> diferir (no cierre a ~0)",
			q:       MarketQuote{BestBid: 0, BestAsk: 0, LastTradePrice: 0.001, Closed: false},
			side:    "yes", ref: 0.44,
			wantPrice: 0, wantSrc: "no_reliable_price", wantOK: false, wantReslv: false,
		},
		{
			name:    "libro vacio, resuelto NO -> settlement real 0",
			q:       MarketQuote{BestBid: 0, LastTradePrice: 0.001, OutcomePricesY: 0.0, OutcomePricesN: 1.0, Closed: true},
			side:    "yes", ref: 0.44,
			wantPrice: 0, wantSrc: "resolved_settlement", wantOK: true, wantReslv: true,
		},
		{
			name:    "libro vacio, resuelto YES -> settlement real 1",
			q:       MarketQuote{BestBid: 0, LastTradePrice: 0.999, OutcomePricesY: 1.0, OutcomePricesN: 0.0, Closed: true},
			side:    "yes", ref: 0.44,
			wantPrice: 1, wantSrc: "resolved_settlement", wantOK: true, wantReslv: true,
		},
		{
			name:    "bestBid real, sano -> bestBid",
			q:       MarketQuote{BestBid: 0.30, BestAsk: 0.32, Closed: false},
			side:    "yes", ref: 0.28,
			wantPrice: 0.30, wantSrc: "bestBid", wantOK: true, wantReslv: false,
		},
		{
			name:    "bestBid colapsado vs ref, no resuelto -> bloqueado",
			q:       MarketQuote{BestBid: 0.001, BestAsk: 0.01, Closed: false},
			side:    "yes", ref: 0.44,
			wantPrice: 0.001, wantSrc: "blocked_sanity_band", wantOK: false, wantReslv: false,
		},
		{
			name:    "side NO: bestAsk real -> bestBid_NO_derived",
			q:       MarketQuote{BestBid: 0.30, BestAsk: 0.35, Closed: false},
			side:    "no", ref: 0.60,
			wantPrice: 0.65, wantSrc: "bestBid_NO_derived", wantOK: true, wantReslv: false,
		},
		{
			name:    "ref bajo (longshot legitimo) NO aplica banda",
			q:       MarketQuote{BestBid: 0.005, BestAsk: 0.01, Closed: false},
			side:    "yes", ref: 0.03,
			wantPrice: 0.005, wantSrc: "bestBid", wantOK: true, wantReslv: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ExitPrice(c.q, c.side, c.ref)
			if got.OK != c.wantOK || got.Source != c.wantSrc || got.Resolved != c.wantReslv {
				t.Fatalf("ExitPrice = {price=%.4f src=%q ok=%v resolved=%v}, want {src=%q ok=%v resolved=%v}",
					got.Price, got.Source, got.OK, got.Resolved, c.wantSrc, c.wantOK, c.wantReslv)
			}
			if d := got.Price - c.wantPrice; d > 1e-9 || d < -1e-9 {
				t.Fatalf("ExitPrice price = %.4f, want %.4f", got.Price, c.wantPrice)
			}
		})
	}
}

// Replay del bug: las 32 filas que cerraron con lastTradePrice_fallback tenian el
// libro vacio (por eso disparo el fallback). 30 NO estaban resueltas (exit_reason
// stop_loss, regimen muerto) y 2 si (market_closed). Con el fix:
//   - no resueltas -> OK=false -> el caller DIFIERE, 0 liquidaciones a <0.01.
//   - resueltas    -> settlement real (resultado correcto, no basura).
func TestReplayFallbackRows_NoSubCentLiquidations(t *testing.T) {
	// (entry, resolved). Empty book reproduce el estado real del cierre fallback.
	type row struct {
		entry    float64
		resolved bool
	}
	rows := make([]row, 0, 32)
	// 30 stop_loss (regimen muerto), no resueltas, entradas dispares 0.05..0.94.
	for _, e := range []float64{0.17, 0.52, 0.49, 0.71, 0.17, 0.57, 0.45, 0.05, 0.31, 0.16,
		0.14, 0.19, 0.05, 0.28, 0.24, 0.24, 0.14, 0.94, 0.24, 0.46,
		0.05, 0.48, 0.17, 0.33, 0.15, 0.93, 0.17, 0.17, 0.38, 0.06} {
		rows = append(rows, row{e, false})
	}
	// 2 market_closed (resolucion genuina): 0.44 y 0.029.
	rows = append(rows, row{0.44, true}, row{0.029, true})

	subCent, deferred, settled := 0, 0, 0
	for _, r := range rows {
		q := MarketQuote{BestBid: 0, BestAsk: 0, LastTradePrice: 0.001, Closed: r.resolved}
		dec := ExitPrice(q, "yes", r.entry)
		if dec.OK && !dec.Resolved && dec.Price < 0.01 {
			subCent++ // exactamente el bug: liquidar vivo a <1 centimo
		}
		if !dec.OK {
			deferred++
		}
		if dec.Resolved {
			settled++
		}
	}
	if subCent != 0 {
		t.Fatalf("el fix permite %d liquidaciones a <0.01 con mercado no resuelto; debe ser 0", subCent)
	}
	if deferred != 30 {
		t.Fatalf("esperaba 30 diferidas (stop_loss regimen muerto), obtuve %d", deferred)
	}
	if settled != 2 {
		t.Fatalf("esperaba 2 settlements reales (market_closed), obtuve %d", settled)
	}
	t.Logf("replay 32 filas fallback: %d diferidas, %d settlement real, 0 liquidaciones a <0.01", deferred, settled)
}
