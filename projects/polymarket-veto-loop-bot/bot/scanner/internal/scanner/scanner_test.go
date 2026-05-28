package scanner

import (
	"fmt"
	"testing"
	"time"

	"github.com/davidgnuez/polymarket-veto-bot/scanner/internal/types"
)

// TestMixedSort_QuotaRespected: 30% short slots reserved, rest by volume desc.
func TestMixedSort_QuotaRespected(t *testing.T) {
	now := time.Now().UTC()
	short := now.Add(3 * 24 * time.Hour).Format(time.RFC3339)  // 3d, qualifies as short
	farLow := now.Add(60 * 24 * time.Hour).Format(time.RFC3339) // 60d, volume bucket

	cands := []types.Candidate{}
	// 5 short-horizon, low volume.
	for i := 0; i < 5; i++ {
		cands = append(cands, types.Candidate{
			ID: fmt.Sprintf("s%d", i), Slug: fmt.Sprintf("s%d", i),
			EndDate: short, Volume24h: 60000,
		})
	}
	// 15 long-horizon, high volume.
	for i := 0; i < 15; i++ {
		cands = append(cands, types.Candidate{
			ID: fmt.Sprintf("l%d", i), Slug: fmt.Sprintf("l%d", i),
			EndDate: farLow, Volume24h: float64(1_000_000 + i*1000),
		})
	}

	out := mixedSort(cands, 10, 30) // 30% of 10 = 3 short slots
	if len(out) != 10 {
		t.Fatalf("len = %d, want 10", len(out))
	}
	// First 3 must be short-horizon.
	for i := 0; i < 3; i++ {
		if out[i].EndDate != short {
			t.Errorf("slot %d should be short-horizon, got endDate=%s slug=%s", i, out[i].EndDate, out[i].Slug)
		}
	}
	// Slot 3 onwards: by volume desc. The remaining 2 short candidates spilled
	// into the volume pool but have lower volume (60k vs 1M), so they end up last.
	if out[3].Volume24h < out[4].Volume24h {
		t.Errorf("volume bucket not sorted desc: %v vs %v", out[3].Volume24h, out[4].Volume24h)
	}
}

// TestMixedSort_EmptyEndDateGoesToVolumeBucket: missing endDate (parsed as
// far-future 2099) must not occupy a short slot.
func TestMixedSort_EmptyEndDateGoesToVolumeBucket(t *testing.T) {
	cands := []types.Candidate{
		{ID: "a", EndDate: "", Volume24h: 100000},
		{ID: "b", EndDate: "", Volume24h: 200000},
	}
	out := mixedSort(cands, 10, 30)
	if len(out) != 2 {
		t.Fatalf("len = %d", len(out))
	}
	if out[0].ID != "b" || out[1].ID != "a" {
		t.Errorf("expected volume-desc order [b,a], got [%s,%s]", out[0].ID, out[1].ID)
	}
}
