package decisionlog

import (
	"os"
	"path/filepath"
	"testing"
)

// Gate behaviour: with DECISION_NODES off the writers are a pure no-op
// (return "", nil and touch no files); with it on they write as before.

func TestWriteVetoGated(t *testing.T) {
	dir := t.TempDir()

	t.Setenv("DECISION_NODES", "")
	p, err := WriteVeto(dir, VetoDecision{Slug: "x", Question: "Q", Rule: "V1"})
	if err != nil || p != "" {
		t.Fatalf("gated off: want \"\",nil got %q,%v", p, err)
	}
	if m, _ := filepath.Glob(filepath.Join(dir, "*.md")); len(m) != 0 {
		t.Fatalf("gated off: expected 0 files, got %d", len(m))
	}

	t.Setenv("DECISION_NODES", "1")
	p, err = WriteVeto(dir, VetoDecision{Slug: "x", Question: "Q", Rule: "V1"})
	if err != nil || p == "" {
		t.Fatalf("gated on: want path,nil got %q,%v", p, err)
	}
	if _, e := os.Stat(p); e != nil {
		t.Fatalf("gated on: file missing: %v", e)
	}
}

func TestWriteTradeGated(t *testing.T) {
	dir := t.TempDir()

	t.Setenv("DECISION_NODES", "")
	p, err := WriteTrade(dir, TradeDecision{Slug: "y", Question: "Q"})
	if err != nil || p != "" {
		t.Fatalf("gated off: want \"\",nil got %q,%v", p, err)
	}

	t.Setenv("DECISION_NODES", "1")
	p, err = WriteTrade(dir, TradeDecision{Slug: "y", Question: "Q"})
	if err != nil || p == "" {
		t.Fatalf("gated on: want path,nil got %q,%v", p, err)
	}
	if _, e := os.Stat(p); e != nil {
		t.Fatalf("gated on: file missing: %v", e)
	}
}
