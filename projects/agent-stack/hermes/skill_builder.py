"""
skill_builder.py — autonomous skill creation from agent run outputs.

Mechanic:
  Agents emit `<skill-proposal>` blocks inside their responses when they detect
  a reusable pattern. Example block (free-form key:value lines):

    <skill-proposal>
    name: linkedin-easy-apply-flow
    agent: job-hunter
    when_to_use: aplicar a oferta con botón "Easy Apply" en LinkedIn
    steps:
      1. browser_evaluate to find selectors
      2. setInputViaJS for tel inputs
      3. fuzzy match dropdowns
    examples:
      - run: 2026-05-24/123456-from-planner
        result: submitted
    </skill-proposal>

  This script:
    1. Scans agent runs at vault/agents/<id>/runs/ for unprocessed proposals.
    2. Validates schema (required fields: name, when_to_use, steps).
    3. Writes a skill file to ~/.openclaw/skills/<agent>/<skill-name>.md
       with proper frontmatter so curator.py picks it up next cycle.
    4. Tracks processed runs in vault/agents/<id>/.skills-built.json to avoid
       re-creating duplicates.

  Repeated-pattern trigger:
    If skill_builder finds the SAME proposal name across >=3 distinct runs and
    the skill doesn't exist yet → auto-create it (signal: the agent itself has
    re-discovered the pattern, so it is genuinely reusable).
    Single-occurrence proposals are recorded as candidates but not built.

Cron: 45 3 * * * (daily 03:45 UTC, just before curator at 04:00... actually
                  curator runs at 03:30 Madrid which is 02:30 UTC, so before)

Hard rules respected:
  - No api.anthropic.com here — pure file I/O.
  - Backup any overwritten skill file as *.bak-YYYY-MM-DD-HHMMSS (we only
    create new, never overwrite an existing skill).
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import re
import sys
from collections import defaultdict

VAULT_ROOT = pathlib.Path(os.environ.get("JARVIS_VAULT", "/home/agent/agent-stack/vault")).resolve()
SKILLS_ROOT = pathlib.Path(os.environ.get("OPENCLAW_SKILLS_DIR", "/home/agent/.openclaw/skills")).resolve()

PROPOSAL_RE = re.compile(r"<skill-proposal>(.*?)</skill-proposal>", re.S)
TRACK_FILE_NAME = ".skills-built.json"
MIN_OCCURRENCES = 3


def now_iso() -> str:
    return dt.datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%SZ")


def parse_proposal_block(block: str) -> dict:
    """Parse a free-form key:value block (YAML-lite). Multiline values OK
    under indentation."""
    out: dict = {}
    current_key = None
    current_lines: list[str] = []
    for line in block.splitlines():
        m = re.match(r"^([a-zA-Z_][\w\-]*):\s*(.*)$", line)
        if m and (not line.startswith(" ") and not line.startswith("\t")):
            if current_key is not None:
                out[current_key] = "\n".join(current_lines).strip()
            current_key = m.group(1).strip()
            current_lines = [m.group(2)] if m.group(2) else []
        elif current_key is not None:
            current_lines.append(line.rstrip())
    if current_key is not None:
        out[current_key] = "\n".join(current_lines).strip()
    return out


def validate(p: dict) -> tuple[bool, str]:
    for field in ("name", "when_to_use", "steps"):
        if not p.get(field):
            return False, f"missing required field: {field}"
    if not re.match(r"^[a-z0-9][a-z0-9\-]+$", p["name"]):
        return False, f"invalid skill name (kebab-case lowercase): {p['name']!r}"
    return True, "ok"


def scan_runs(agent_id: str) -> list[tuple[pathlib.Path, dict]]:
    """Returns list of (run_file_path, parsed_proposal) from this agent's runs."""
    runs_dir = VAULT_ROOT / "agents" / agent_id / "runs"
    if not runs_dir.exists():
        return []
    found: list[tuple[pathlib.Path, dict]] = []
    for f in runs_dir.rglob("*.md"):
        try:
            text = f.read_text(encoding="utf-8", errors="replace")
        except Exception:
            continue
        for m in PROPOSAL_RE.finditer(text):
            p = parse_proposal_block(m.group(1))
            ok, _ = validate(p)
            if ok:
                p.setdefault("agent", agent_id)
                found.append((f, p))
    return found


def load_tracking(agent_id: str) -> dict:
    f = VAULT_ROOT / "agents" / agent_id / TRACK_FILE_NAME
    if not f.exists():
        return {}
    try:
        return json.loads(f.read_text(encoding="utf-8"))
    except Exception:
        return {}


def save_tracking(agent_id: str, data: dict) -> None:
    f = VAULT_ROOT / "agents" / agent_id / TRACK_FILE_NAME
    f.parent.mkdir(parents=True, exist_ok=True)
    f.write_text(json.dumps(data, indent=2), encoding="utf-8")


def skill_path(agent_id: str, name: str) -> pathlib.Path:
    return SKILLS_ROOT / agent_id / f"{name}.md"


def render_skill_md(p: dict) -> str:
    now = now_iso()
    fm_lines = [
        "---",
        f'name: {p["name"]}',
        f'agent: {p.get("agent","unknown")}',
        f'state: active',
        f'pinned: false',
        f'source: skill_builder.py',
        f'created: {now}',
        "---",
        "",
        f"# {p['name']}",
        "",
        "## When to use",
        "",
        p["when_to_use"].strip(),
        "",
        "## Steps",
        "",
        p["steps"].strip(),
        "",
    ]
    if p.get("examples"):
        fm_lines.append("## Examples")
        fm_lines.append("")
        fm_lines.append(p["examples"].strip())
        fm_lines.append("")
    return "\n".join(fm_lines)


def process_agent(agent_id: str, dry_run: bool) -> dict:
    proposals = scan_runs(agent_id)
    if not proposals:
        return {"agent": agent_id, "found": 0, "created": 0}

    # Group by skill name
    by_name: dict[str, list[tuple[pathlib.Path, dict]]] = defaultdict(list)
    for run_path, p in proposals:
        by_name[p["name"]].append((run_path, p))

    tracking = load_tracking(agent_id)
    created = 0
    candidates_logged = 0

    for name, items in by_name.items():
        # Already built? Skip.
        target = skill_path(agent_id, name)
        if target.exists() or tracking.get(name, {}).get("created"):
            continue
        # Count distinct runs
        distinct_runs = {str(r) for r, _ in items}
        if len(distinct_runs) < MIN_OCCURRENCES:
            tracking.setdefault(name, {})["candidate_runs"] = sorted(distinct_runs)
            tracking[name]["last_seen"] = now_iso()
            candidates_logged += 1
            continue
        # Build
        if dry_run:
            print(f"  [dry] would create {target} from {len(distinct_runs)} runs")
        else:
            target.parent.mkdir(parents=True, exist_ok=True)
            # Use the most recent proposal as the source of truth
            items.sort(key=lambda t: t[0].stat().st_mtime, reverse=True)
            _, latest_proposal = items[0]
            target.write_text(render_skill_md(latest_proposal), encoding="utf-8")
            tracking[name] = {
                "created": now_iso(),
                "source_runs": sorted(distinct_runs),
                "path": str(target),
            }
            created += 1

    if not dry_run and tracking:
        save_tracking(agent_id, tracking)

    return {
        "agent": agent_id,
        "found_proposals": sum(len(v) for v in by_name.values()),
        "distinct_skills": len(by_name),
        "created": created,
        "candidates_logged": candidates_logged,
    }


def main(argv: list[str] | None = None) -> int:
    p = argparse.ArgumentParser()
    p.add_argument("--dry-run", action="store_true")
    p.add_argument("agents", nargs="*",
                   help="agent ids to scan (default: all subdirs of vault/agents/)")
    args = p.parse_args(argv)

    if args.agents:
        agent_ids = args.agents
    else:
        agents_dir = VAULT_ROOT / "agents"
        agent_ids = [d.name for d in agents_dir.iterdir() if d.is_dir() and not d.name.startswith(".")]

    summary = []
    for aid in agent_ids:
        r = process_agent(aid, args.dry_run)
        summary.append(r)
        print(json.dumps(r))
    total_created = sum(r.get("created", 0) for r in summary)
    print(f"\nskill_builder: {total_created} skills created across {len(summary)} agents (dry-run={args.dry_run})")
    return 0


if __name__ == "__main__":
    sys.exit(main())
