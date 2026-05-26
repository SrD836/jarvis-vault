#!/usr/bin/env python3
"""One-shot: collapse duplicate <date>-polymarket-trade-<slug>.md files.

Before WriteTrade was moved from brain to executor (2026-05-26), brain wrote
one .md per re-approval day. The same slug could end up with 5+ .md files
(one entry day, four re-approvals) all flagged outcome: pending — only the
oldest one corresponds to a real trade.

Keeps the OLDEST .md per slug (closest to real entry). Moves the rest to
03-decisions/.archive-polymarket-md/dedupe-YYYY-MM-DD/.

Usage from vault root:
    python3 projects/polymarket-veto-loop-bot/scripts/dedupe_trade_md.py --dry-run
    python3 projects/polymarket-veto-loop-bot/scripts/dedupe_trade_md.py --apply
"""
from __future__ import annotations
import argparse
import re
import shutil
import sys
from collections import defaultdict
from datetime import datetime, timezone
from pathlib import Path

VAULT = Path(__file__).resolve().parents[3]
DEC = VAULT / "03-decisions"
ARCHIVE = DEC / ".archive-polymarket-md" / f"dedupe-{datetime.now(timezone.utc):%Y-%m-%d}"

FN_RE = re.compile(r"^(\d{4})-(\d{2})-(\d{2})-polymarket-trade-(.+)\.md$")


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--apply", action="store_true", help="actually move files; default is dry-run")
    ap.add_argument("--dry-run", action="store_true", help="(default) just report")
    args = ap.parse_args()
    apply = args.apply and not args.dry_run

    bySlug: dict[str, list[tuple[str, Path]]] = defaultdict(list)
    for md in DEC.glob("*-polymarket-trade-*.md"):
        m = FN_RE.match(md.name)
        if not m:
            continue
        date = f"{m.group(1)}-{m.group(2)}-{m.group(3)}"
        slug = m.group(4)
        bySlug[slug].append((date, md))

    total = sum(len(v) for v in bySlug.values())
    dup_groups = {s: v for s, v in bySlug.items() if len(v) > 1}
    dup_count = sum(len(v) - 1 for v in dup_groups.values())
    print(f"Found {total} trade .md across {len(bySlug)} slugs.")
    print(f"Duplicates: {dup_count} files in {len(dup_groups)} slugs.")

    if not apply:
        # Show 10 worst offenders.
        worst = sorted(dup_groups.items(), key=lambda kv: -len(kv[1]))[:10]
        for slug, items in worst:
            print(f"  {len(items)}× {slug}")
        print("\nDRY-RUN — pass --apply to actually move the duplicates.")
        return 0

    ARCHIVE.mkdir(parents=True, exist_ok=True)
    moved = 0
    for slug, items in dup_groups.items():
        items.sort()  # ascending by date string → oldest first
        keep_date, keep_path = items[0]
        for date, path in items[1:]:
            dst = ARCHIVE / path.name
            shutil.move(str(path), str(dst))
            moved += 1
        # leave a marker noting the keep file
    print(f"Moved {moved} duplicate trade .md to {ARCHIVE}")
    print(f"Kept {len(bySlug)} canonical .md (one per slug, oldest date).")
    return 0


if __name__ == "__main__":
    sys.exit(main())
