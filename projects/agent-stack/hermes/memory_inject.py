"""
memory_inject.py — periodically refreshes a `## Memoria viva` section in each
agent's AGENTS.md (at /home/agent/.openclaw/workspace/agents/<id>/AGENTS.md)
with the compact memory summary produced by memory_lib.summarize_for_prompt.

The section is delimited by HTML comments so we can safely replace it on every
run without touching surrounding human-authored Standing Orders.

Cron: every 15 min (low cost; only writes when content changes).

Usage:
    python3 memory_inject.py            # all known agents
    python3 memory_inject.py main planner   # subset
"""
from __future__ import annotations

import os
import pathlib
import sys

sys.path.insert(0, str(pathlib.Path(__file__).parent))
import memory_lib  # noqa: E402

AGENTS_WORKSPACE = pathlib.Path(
    os.environ.get("OPENCLAW_AGENTS_DIR", "/home/agent/.openclaw/workspace/agents")
).resolve()

BEGIN = "<!-- BEGIN AUTO-MEMORY (managed by memory_inject.py) -->"
END = "<!-- END AUTO-MEMORY -->"

ALL_AGENTS = [
    "main", "planner", "apier", "archivist", "auditor", "code-reviewer",
    "debugger", "designer", "documenter", "job-hunter", "monitor",
    "researcher", "skill-reviewer", "tester", "polymarket-handler",
]


def inject(agent_id: str) -> tuple[bool, str]:
    agents_md = AGENTS_WORKSPACE / agent_id / "AGENTS.md"
    if not agents_md.exists():
        return False, f"missing {agents_md}"
    summary = memory_lib.summarize_for_prompt(agent_id, max_chars=3500)
    if not summary.strip():
        return False, "empty summary"

    block = f"\n\n{BEGIN}\n## Memoria viva\n\n{summary}\n{END}\n"

    raw = agents_md.read_text(encoding="utf-8")
    if BEGIN in raw and END in raw:
        pre, _, rest = raw.partition(BEGIN)
        _, _, post = rest.partition(END)
        new = pre.rstrip() + block + post.lstrip()
    else:
        new = raw.rstrip() + block

    if new == raw:
        return False, "unchanged"
    agents_md.write_text(new, encoding="utf-8")
    return True, f"updated {agents_md} ({len(summary)} chars)"


def main(argv: list[str]) -> int:
    agents = argv[1:] or ALL_AGENTS
    changed = 0
    for a in agents:
        ok, msg = inject(a)
        marker = "✓" if ok else "·"
        print(f"  {marker} {a}: {msg}")
        if ok:
            changed += 1
    print(f"memory_inject: {changed}/{len(agents)} updated")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
