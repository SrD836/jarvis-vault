// Command mark_to_market values every OPEN position in active.jsonl at the
// current market price, reusing the canonical polyclient.ExitPrice (Fase 7):
// resolved markets settle at 0/1, live markets use a sanity-banded bestBid, and
// anything without a reliable price is reported as "unpriceable" (never a
// garbage fallback). Read-only: it never closes a position. Emits a JSON
// snapshot consumed by analytics/expectancy_report.py so the survivorship-free
// expectancy can net unrealized P&L of the still-open (mostly longshot) book.
//
// Usage:
//
//	mark_to_market --active vault/inbox/trading/active.jsonl --out vault/inbox/trading/mark_to_market.json
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/polyclient"
	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

type posResult struct {
	MarketID       string   `json:"market_id"`
	Slug           string   `json:"slug,omitempty"`
	Category       string   `json:"category,omitempty"`
	Side           string   `json:"side"`
	EntryPrice     float64  `json:"entry_price"`
	SizeUSD        float64  `json:"size_usd"`
	ExitPrice      float64  `json:"exit_price"`
	Source         string   `json:"source"`
	Resolved       bool     `json:"resolved"`
	Priceable      bool     `json:"priceable"`
	UnrealizedPnl  *float64 `json:"unrealized_pnl"` // nil when unpriceable
}

type snapshot struct {
	GeneratedAt     string      `json:"generated_at"`
	N               int         `json:"n"`
	NPriced         int         `json:"n_priced"`
	NUnpriceable    int         `json:"n_unpriceable"`
	TotalUnrealized float64     `json:"total_unrealized"` // over priced positions only
	CostUnpriceable float64     `json:"cost_unpriceable"` // size_usd at risk we could not price
	Positions       []posResult `json:"positions"`
}

func sizeUSD(a types.ActiveTrade) float64 {
	if a.SizeUSD > 0 {
		return a.SizeUSD
	}
	return a.Size
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	active := flag.String("active", "vault/inbox/trading/active.jsonl", "active.jsonl path")
	out := flag.String("out", "vault/inbox/trading/mark_to_market.json", "output JSON path")
	flag.Parse()

	f, err := os.Open(*active)
	if err != nil {
		log.Fatalf("mark_to_market: open %s: %v", *active, err)
	}
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	snap := snapshot{GeneratedAt: time.Now().UTC().Format(time.RFC3339)}
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var a types.ActiveTrade
		if err := json.Unmarshal(line, &a); err != nil {
			continue
		}
		if a.MarketID == "" || a.EntryPrice <= 0 {
			continue
		}
		snap.N++
		q, err := polyclient.FetchQuote(a.MarketID)
		sz := sizeUSD(a)
		pr := posResult{
			MarketID: a.MarketID, Slug: a.Slug, Category: a.Category,
			Side: a.Side, EntryPrice: a.EntryPrice, SizeUSD: sz,
		}
		if err != nil {
			log.Printf("mark_to_market: fetch %s failed: %v", a.MarketID, err)
			pr.Source = "fetch_error"
			snap.NUnpriceable++
			snap.CostUnpriceable += sz
			snap.Positions = append(snap.Positions, pr)
			continue
		}
		d := polyclient.ExitPrice(q, a.Side, a.EntryPrice)
		pr.ExitPrice = d.Price
		pr.Source = d.Source
		pr.Resolved = d.Resolved
		pr.Priceable = d.OK
		if !d.OK {
			snap.NUnpriceable++
			snap.CostUnpriceable += sz
			snap.Positions = append(snap.Positions, pr)
			continue
		}
		// shares = sizeUSD/entry ; valor_now = shares*exit ; unrealized = sizeUSD*(exit/entry - 1)
		u := sz * (d.Price/a.EntryPrice - 1.0)
		pr.UnrealizedPnl = &u
		snap.NPriced++
		snap.TotalUnrealized += u
		snap.Positions = append(snap.Positions, pr)
		log.Printf("mark_to_market: %s side=%s entry=%.3f exit=%.3f (%s) unrealized=%.2f",
			a.Slug, a.Side, a.EntryPrice, d.Price, d.Source, u)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("mark_to_market: scan: %v", err)
	}
	f.Close()

	b, _ := json.MarshalIndent(snap, "", "  ")
	if err := os.WriteFile(*out, b, 0o644); err != nil {
		log.Fatalf("mark_to_market: write %s: %v", *out, err)
	}
	log.Printf("mark_to_market: n=%d priced=%d unpriceable=%d total_unrealized=%.2f cost_unpriceable=%.2f",
		snap.N, snap.NPriced, snap.NUnpriceable, snap.TotalUnrealized, snap.CostUnpriceable)
}
