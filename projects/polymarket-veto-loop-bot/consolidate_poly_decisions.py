#!/usr/bin/env python3
"""Consolidate vault/03-decisions/<date>-polymarket-*.md into monthly JSONL.

Behavior:
- Reads every YYYY-MM-DD-polymarket-(trade|veto|loss)-*.md in 03-decisions/
- Parses YAML frontmatter (simple line-by-line, no yaml dep needed)
- Appends one JSON line per file to 03-decisions/polymarket/<YYYY-MM>.jsonl
- Files with outcome != 'pending' are MOVED to 03-decisions/.archive-polymarket-md/<YYYY-MM>/
  Files with outcome == 'pending' stay in place (exit_monitor still needs to patch them)

Run once from the vault root:
    python3 projects/polymarket-veto-loop-bot/consolidate_poly_decisions.py
"""
from __future__ import annotations
import json
import re
import shutil
import sys
from pathlib import Path

VAULT = Path(__file__).resolve().parents[2]  # vault root
DEC = VAULT / "03-decisions"
OUT_DIR = DEC / "polymarket"
ARCHIVE = DEC / ".archive-polymarket-md"

FN_RE = re.compile(r"^(\d{4})-(\d{2})-(\d{2})-polymarket-(trade|veto|loss)-(.+)\.md$")


def parse_frontmatter(text: str) -> dict:
    if not text.startswith("---"):
        return {}
    end = text.find("\n---", 3)
    if end < 0:
        return {}
    block = text[3:end].strip("\n")
    out = {}
    for line in block.splitlines():
        if ":" not in line or line.startswith(" ") or line.startswith("-"):
            continue
        k, _, v = line.partition(":")
        v = v.strip().strip('"').strip("'")
        out[k.strip()] = v
    return out


def extract_record(path: Path) -> dict | None:
    m = FN_RE.match(path.name)
    if not m:
        return None
    yyyy, mm, dd, kind, slug = m.groups()
    text = path.read_text(encoding="utf-8", errors="replace")
    fm = parse_frontmatter(text)
    rec = {
        "ts": f"{yyyy}-{mm}-{dd}T00:00:00Z",
        "kind": kind,
        "slug": slug,
        "title": fm.get("title", "").strip(),
        "outcome": fm.get("outcome", "pending"),
        "source_file": path.name,
    }
    for key in ("market_id", "entry_price", "size_usd", "horizon",
                "days_to_resolution", "outcome_observed_after_days",
                "decision_type", "decision", "exit_price", "pnl"):
        if key in fm:
            rec[key] = fm[key]
    body_start = text.find("\n---", 3)
    if body_start > 0:
        body = text[body_start + 4 :].strip()
        rec["body_excerpt"] = body[:600]
    return rec


def main(dry_run: bool = False) -> int:
    if not DEC.is_dir():
        print(f"ERROR: {DEC} not found", file=sys.stderr)
        return 1
    OUT_DIR.mkdir(exist_ok=True)
    by_month: dict[str, list[dict]] = {}
    archive_paths: list[tuple[Path, Path]] = []
    skipped_pending = 0
    parse_fail = 0

    for md in sorted(DEC.glob("*-polymarket-*.md")):
        rec = extract_record(md)
        if rec is None:
            parse_fail += 1
            continue
        month = rec["ts"][:7]
        by_month.setdefault(month, []).append(rec)
        if rec["outcome"] == "pending":
            skipped_pending += 1
        else:
            arch_dir = ARCHIVE / month
            archive_paths.append((md, arch_dir / md.name))

    print(f"Found {sum(len(v) for v in by_month.values())} polymarket .md files across {len(by_month)} months")
    print(f"  pending (stay in place): {skipped_pending}")
    print(f"  closed (will archive):   {len(archive_paths)}")
    print(f"  parse failures:          {parse_fail}")

    if dry_run:
        print("DRY-RUN — no files written")
        return 0

    for month, recs in by_month.items():
        recs.sort(key=lambda r: (r["ts"], r["slug"]))
        out_file = OUT_DIR / f"{month}.jsonl"
        with out_file.open("w", encoding="utf-8") as f:
            for r in recs:
                f.write(json.dumps(r, ensure_ascii=False) + "\n")
        print(f"  wrote {out_file} ({len(recs)} lines)")

    for src, dst in archive_paths:
        dst.parent.mkdir(parents=True, exist_ok=True)
        shutil.move(str(src), str(dst))
    print(f"  archived {len(archive_paths)} closed .md files to {ARCHIVE}/")
    return 0


if __name__ == "__main__":
    sys.exit(main(dry_run="--dry-run" in sys.argv))
