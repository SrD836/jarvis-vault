#!/usr/bin/env python3
"""Expire veto .md whose observation window has passed.

A veto file's frontmatter declares `date` and `outcome_observed_after_days`
(default 30). exit_monitor only patches *trade* .md, never vetos, so once a
veto is older than that window with outcome=pending it lingers forever and
clogs the Obsidian graph. This script:

  - Finds *-polymarket-veto-*.md with outcome: pending and date older than
    today - outcome_observed_after_days.
  - Appends an `outcome: expired` record to 03-decisions/polymarket/<YYYY-MM>.jsonl.
  - Moves the .md to 03-decisions/.archive-polymarket-md/<YYYY-MM>/.

Idempotent. Default dry-run.

Usage:
    python3 projects/polymarket-veto-loop-bot/scripts/expire_pending_vetos.py [--apply]
"""
from __future__ import annotations
import argparse
import json
import re
import shutil
import sys
from datetime import datetime, timezone, timedelta
from pathlib import Path

VAULT = Path(__file__).resolve().parents[3]
DEC = VAULT / "03-decisions"
OUT_DIR = DEC / "polymarket"
ARCHIVE_BASE = DEC / ".archive-polymarket-md"

FN_RE = re.compile(r"^(\d{4})-(\d{2})-(\d{2})-polymarket-veto-(.+)\.md$")


def parse_frontmatter(text: str) -> dict:
    if not text.startswith("---"):
        return {}
    end = text.find("\n---", 3)
    if end < 0:
        return {}
    block = text[3:end].strip("\n")
    out: dict = {}
    for line in block.splitlines():
        if ":" not in line or line.startswith(" ") or line.startswith("-"):
            continue
        k, _, v = line.partition(":")
        v = v.strip().strip('"').strip("'")
        out[k.strip()] = v
    return out


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--apply", action="store_true")
    args = ap.parse_args()
    today = datetime.now(timezone.utc).date()

    candidates: list[tuple[Path, dict, str]] = []  # (path, fm, month_key)
    for md in DEC.glob("*-polymarket-veto-*.md"):
        m = FN_RE.match(md.name)
        if not m:
            continue
        fname_date = f"{m.group(1)}-{m.group(2)}-{m.group(3)}"
        text = md.read_text(encoding="utf-8", errors="replace")
        fm = parse_frontmatter(text)
        if fm.get("outcome", "pending") != "pending":
            continue
        try:
            window = int(fm.get("outcome_observed_after_days", "30"))
        except ValueError:
            window = 30
        try:
            d = datetime.strptime(fm.get("date", fname_date), "%Y-%m-%d").date()
        except ValueError:
            d = datetime.strptime(fname_date, "%Y-%m-%d").date()
        if today - d <= timedelta(days=window):
            continue
        candidates.append((md, fm, f"{m.group(1)}-{m.group(2)}"))

    print(f"Pending vetos eligible for expiration: {len(candidates)}")
    if not candidates:
        return 0
    if not args.apply:
        for md, fm, _ in candidates[:5]:
            print(f"  {md.name}  (date={fm.get('date','?')}, rule={fm.get('tags','?')})")
        if len(candidates) > 5:
            print(f"  ... and {len(candidates)-5} more")
        print("\nDRY-RUN — pass --apply to expire and archive.")
        return 0

    OUT_DIR.mkdir(parents=True, exist_ok=True)
    by_month: dict[str, list[dict]] = {}
    for md, fm, month in candidates:
        rec = {
            "ts": f"{fm.get('date', md.name[:10])}T00:00:00Z",
            "kind": "veto",
            "slug": md.name.split("-polymarket-veto-", 1)[1].removesuffix(".md"),
            "title": fm.get("title", ""),
            "outcome": "expired",
            "source_file": md.name,
        }
        for key in ("market_id", "decision_type", "decision"):
            if key in fm:
                rec[key] = fm[key]
        by_month.setdefault(month, []).append(rec)

    for month, recs in by_month.items():
        out_file = OUT_DIR / f"{month}.jsonl"
        with out_file.open("a", encoding="utf-8") as f:
            for r in recs:
                f.write(json.dumps(r, ensure_ascii=False) + "\n")
        print(f"  appended {len(recs)} expired records → {out_file}")

    moved = 0
    for md, _, month in candidates:
        arch_dir = ARCHIVE_BASE / month
        arch_dir.mkdir(parents=True, exist_ok=True)
        shutil.move(str(md), str(arch_dir / md.name))
        moved += 1
    print(f"Moved {moved} expired veto .md to {ARCHIVE_BASE}/<month>/")
    return 0


if __name__ == "__main__":
    sys.exit(main())
