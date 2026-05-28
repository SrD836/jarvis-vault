package memory

import "testing"

func TestSlugTokenJaccard_NearVariant(t *testing.T) {
	// Two slugs differing only on the date suffix should match strongly.
	j := slugTokenJaccard(
		"iran-closes-its-airspace-by-may-29",
		"iran-closes-its-airspace-by-may-31",
	)
	if j < 0.70 {
		t.Fatalf("expected jaccard >= 0.70 for near-identical slugs, got %.3f", j)
	}
}

func TestSlugTokenJaccard_Unrelated(t *testing.T) {
	j := slugTokenJaccard(
		"iran-closes-its-airspace-by-may-29",
		"will-jannik-sinner-win-the-2026-mens-french-open",
	)
	if j > 0.10 {
		t.Fatalf("expected jaccard <= 0.10 for unrelated slugs, got %.3f", j)
	}
}

func TestSlugTokenJaccard_StopwordsOnly(t *testing.T) {
	// Only stopwords + digits → 0 tokens → 0 jaccard.
	j := slugTokenJaccard("will-be-on-2026", "is-the-2027")
	if j != 0 {
		t.Fatalf("expected 0 jaccard for stopword-only slugs, got %.3f", j)
	}
}

func TestMatch_FuzzyHitsAboveBrainThreshold(t *testing.T) {
	// A loss memorized for one airspace market should hit the fuzzy threshold
	// when a sibling market (different date) shows up as a candidate.
	m := &Memory{
		Losses: []Pattern{{
			Slug:     "iran-closes-its-airspace-by-may-29",
			Category: "geopolitics",
			PriceYes: 0.12,
			Source:   "loss",
		}},
	}
	hits := m.Match(CandidateLike{
		Slug:            "iran-closes-its-airspace-by-may-31",
		Category:        "geopolitics",
		CurrentPriceYes: 0.13,
	})
	if len(hits) == 0 {
		t.Fatal("expected at least one fuzzy hit")
	}
	if hits[0].Score < 0.7 {
		t.Fatalf("expected fuzzy hit score >= 0.7 (brain veto threshold), got %.3f (why=%s)", hits[0].Score, hits[0].Why)
	}
}

func TestMatch_ExactSlugStillWins(t *testing.T) {
	m := &Memory{
		Vetos: []Pattern{{
			Slug:     "bitcoin-above-74k-on-may-28-2026",
			Category: "market",
			PriceYes: 0.26,
			Source:   "veto",
		}},
	}
	hits := m.Match(CandidateLike{
		Slug:            "bitcoin-above-74k-on-may-28-2026",
		Category:        "market",
		CurrentPriceYes: 0.27,
	})
	if len(hits) == 0 || hits[0].Score != 1.0 {
		t.Fatalf("expected exact-slug score=1.0, got hits=%+v", hits)
	}
}
