// Command categorize_backfill fills the `category` field of legacy closed.jsonl
// rows that predate the scanner's categorize.Infer wiring (commit 55a0170c).
//
// Rows whose category is "" or "uncategorized" get categorize.Infer(slug,
// question) — the SAME canonical classifier the scanner uses, so there is zero
// regex drift. Every other row is copied byte-for-byte (json.RawMessage), so
// numeric fields (pnl, prices) keep their exact original formatting. Atomic
// write (tmp + rename), idempotent, one-shot.
//
// Usage:
//
//	categorize_backfill --path vault/inbox/trading/closed.jsonl [--dry-run]
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/common/categorize"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	path := flag.String("path", "vault/inbox/trading/closed.jsonl", "closed.jsonl path")
	dryRun := flag.Bool("dry-run", false, "report counts, do not write")
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("categorize_backfill: open %s: %v", *path, err)
	}
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var out bytes.Buffer
	total, changed := 0, 0
	for sc.Scan() {
		line := sc.Bytes()
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		total++

		// RawMessage preserves every untouched field byte-exact (no float
		// round-trip). Only the category value is replaced when missing.
		var row map[string]json.RawMessage
		if err := json.Unmarshal(line, &row); err != nil {
			out.Write(line) // unparseable: copy verbatim, never drop data
			out.WriteByte('\n')
			continue
		}

		var cat string
		if rc, ok := row["category"]; ok {
			_ = json.Unmarshal(rc, &cat)
		}
		if cat != "" && cat != "uncategorized" {
			out.Write(line) // already categorised: verbatim
			out.WriteByte('\n')
			continue
		}

		var slug, question string
		if rs, ok := row["slug"]; ok {
			_ = json.Unmarshal(rs, &slug)
		}
		if rq, ok := row["question"]; ok {
			_ = json.Unmarshal(rq, &question)
		}
		nc, err := json.Marshal(categorize.Infer(slug, question))
		if err != nil {
			log.Fatalf("categorize_backfill: marshal category: %v", err)
		}
		row["category"] = nc
		changed++
		b, err := json.Marshal(row)
		if err != nil {
			log.Fatalf("categorize_backfill: marshal row: %v", err)
		}
		out.Write(b)
		out.WriteByte('\n')
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("categorize_backfill: scan: %v", err)
	}
	f.Close()

	log.Printf("categorize_backfill: %d rows, %d backfilled", total, changed)
	if *dryRun {
		log.Printf("[dry-run] no write")
		return
	}

	tmp := filepath.Join(filepath.Dir(*path), ".categorize_backfill.tmp")
	if err := os.WriteFile(tmp, out.Bytes(), 0o644); err != nil {
		log.Fatalf("categorize_backfill: write tmp: %v", err)
	}
	if err := os.Rename(tmp, *path); err != nil {
		log.Fatalf("categorize_backfill: rename: %v", err)
	}
	fmt.Printf("OK backfilled=%d total=%d\n", changed, total)
}
