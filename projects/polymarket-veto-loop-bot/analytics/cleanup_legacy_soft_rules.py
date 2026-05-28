#!/usr/bin/env python3
"""
One-shot cleanup: purge auto-generated "Reglas blandas aprendidas" lines that
target the legacy `uncategorized` bucket. Those rules accumulated to 126
entries over the v6 era and stopped matching once the scanner started
emitting real categories. Phase 1 of v7 remediation.

Idempotent. Writes backup memory.md.bak-YYYY-MM-DD next to the file before
mutating. Preserves the "Reglas iniciales seed" block (human-curated) by
only deleting lines that start with the auto-generated marker pattern:

    - En categoría `uncategorized` · horizonte ...

The HTML comments (`<!-- auto-generated ... -->`) directly above purged
rules are also dropped when they would orphan as adjacent siblings.

Usage:
    python3 cleanup_legacy_soft_rules.py [--memory PATH] [--apply]
"""
from __future__ import annotations

import argparse
import datetime as dt
import pathlib
import re
import sys

DEFAULT_MEMORY = "vault/agents/polymarket-bot/memory.md"
RULE_RE = re.compile(r"^- En categoría `uncategorized` · ")


def cleanup(text: str) -> tuple[str, int]:
    """Return (new_text, n_purged). Walks lines, drops uncategorized rule
    rows, also drops the immediately-preceding auto-generated comment when
    the next non-comment line is purged."""
    lines = text.split("\n")
    out: list[str] = []
    purged = 0
    skip_pending_comment: int | None = None  # index in `out` to potentially drop later
    for ln in lines:
        if RULE_RE.match(ln.strip()):
            purged += 1
            # If the previous emitted line was the auto-generated comment AND
            # no other content sits between → drop that comment too so we
            # don't leave dangling timestamps.
            if skip_pending_comment is not None and skip_pending_comment == len(out) - 1:
                out.pop()
            skip_pending_comment = None
            continue
        stripped = ln.strip()
        if stripped.startswith("<!-- auto-generated ") and stripped.endswith("-->"):
            skip_pending_comment = len(out)
        elif stripped:
            skip_pending_comment = None
        out.append(ln)
    return "\n".join(out), purged


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--memory", default=DEFAULT_MEMORY)
    ap.add_argument("--apply", action="store_true", help="write changes; default is dry-run")
    args = ap.parse_args()

    p = pathlib.Path(args.memory)
    if not p.exists():
        print(f"memory.md not found at {p}", file=sys.stderr)
        return 1
    text = p.read_text(encoding="utf-8")
    new_text, purged = cleanup(text)
    print(f"purged {purged} uncategorized soft-rule lines")
    if not args.apply:
        print("dry-run (pass --apply to write)")
        return 0
    if purged == 0:
        print("nothing to write")
        return 0
    backup = p.with_suffix(p.suffix + ".bak-" + dt.date.today().isoformat())
    backup.write_text(text, encoding="utf-8")
    print(f"backup -> {backup}")
    p.write_text(new_text, encoding="utf-8")
    print(f"wrote {p} ({purged} lines removed)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
