#!/usr/bin/env python3
"""
Backfill slug field on rows in closed.jsonl that ended up with slug=null.
Looks up the slug in candidates.jsonl (current snapshot) by market_id. If
the candidates snapshot doesn't have it, optionally scans rotated archives
under vault/inbox/trading/candidates.*.jsonl (none today but future-proof).

Idempotent. Writes backup closed.jsonl.bak-YYYY-MM-DD before mutating.

Usage:
    python3 backfill_slugs.py [--apply]
"""
from __future__ import annotations

import argparse
import datetime as dt
import glob
import json
import pathlib
import sys

CLOSED_PATH = "vault/inbox/trading/closed.jsonl"
CANDIDATES_PATH = "vault/inbox/trading/candidates.jsonl"
ARCHIVE_GLOB = "vault/inbox/trading/candidates.*.jsonl"


def load_slug_lookup(paths: list[pathlib.Path]) -> dict[str, str]:
    """Build market_id → slug map from one or more candidate files. Last
    seen wins (newer archives override older)."""
    lookup: dict[str, str] = {}
    for p in paths:
        if not p.exists():
            continue
        with p.open("r", encoding="utf-8") as f:
            for line in f:
                line = line.strip()
                if not line:
                    continue
                try:
                    o = json.loads(line)
                except json.JSONDecodeError:
                    continue
                mid = str(o.get("id") or o.get("market_id") or "").strip()
                slug = (o.get("slug") or "").strip()
                if mid and slug:
                    lookup[mid] = slug
    return lookup


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--closed", default=CLOSED_PATH)
    ap.add_argument("--candidates", default=CANDIDATES_PATH)
    ap.add_argument("--apply", action="store_true", help="write changes; default dry-run")
    args = ap.parse_args()

    closed_path = pathlib.Path(args.closed)
    if not closed_path.exists():
        print(f"closed.jsonl not found at {closed_path}", file=sys.stderr)
        return 1
    candidate_paths = [pathlib.Path(args.candidates)] + [
        pathlib.Path(p) for p in glob.glob(ARCHIVE_GLOB)
    ]
    lookup = load_slug_lookup(candidate_paths)
    print(f"loaded {len(lookup)} market_id->slug mappings from {sum(1 for p in candidate_paths if p.exists())} candidate file(s)")

    rows: list[dict] = []
    with closed_path.open("r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                rows.append(json.loads(line))
            except json.JSONDecodeError:
                # preserve unparseable rows verbatim by stashing back as text — but
                # closed.jsonl is JSONL so this should not happen; skip.
                continue

    fixed = 0
    missing = 0
    for r in rows:
        slug = r.get("slug")
        if slug:
            continue
        mid = str(r.get("market_id") or "").strip()
        if not mid:
            missing += 1
            continue
        if mid in lookup:
            r["slug"] = lookup[mid]
            fixed += 1
        else:
            missing += 1

    print(f"backfilled {fixed} slugs; {missing} still missing (market_id not in candidates)")
    if not args.apply:
        print("dry-run (pass --apply to write)")
        return 0
    if fixed == 0:
        print("nothing to write")
        return 0

    backup = closed_path.with_suffix(closed_path.suffix + ".bak-" + dt.date.today().isoformat())
    backup.write_text(closed_path.read_text(encoding="utf-8"), encoding="utf-8")
    print(f"backup -> {backup}")

    tmp = closed_path.with_suffix(closed_path.suffix + ".tmp")
    with tmp.open("w", encoding="utf-8") as f:
        for r in rows:
            f.write(json.dumps(r, ensure_ascii=False) + "\n")
    tmp.replace(closed_path)
    print(f"wrote {closed_path} ({fixed} slugs filled)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
