package main

import (
	"strings"
	"testing"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/types"
)

func TestSplitKeepShortOnly(t *testing.T) {
	actives := []types.ActiveTrade{
		{ID: "1", Horizon: "short"},
		{ID: "2", Horizon: "medium"},
		{ID: "3", Horizon: "long"},
		{ID: "4", Horizon: "Short"}, // case-insensitive
		{ID: "5", Horizon: " short "}, // trimmed
	}
	toClose, kept := splitKeep(actives, map[string]bool{"short": true})
	if len(kept) != 3 {
		t.Fatalf("want 3 short kept, got %d: %+v", len(kept), kept)
	}
	if len(toClose) != 2 {
		t.Fatalf("want 2 (medium,long) to close, got %d: %+v", len(toClose), toClose)
	}
	for _, k := range kept {
		if strings.TrimSpace(strings.ToLower(k.Horizon)) != "short" {
			t.Fatalf("kept a non-short position: %+v", k)
		}
	}
}

func TestSplitKeepEmptyClosesAll(t *testing.T) {
	actives := []types.ActiveTrade{{ID: "1", Horizon: "short"}, {ID: "2", Horizon: "long"}}
	toClose, kept := splitKeep(actives, map[string]bool{})
	if len(toClose) != 2 || len(kept) != 0 {
		t.Fatalf("empty keep must close all; got close=%d kept=%d", len(toClose), len(kept))
	}
}
