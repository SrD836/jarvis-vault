"""
memory_lib.py — single source of truth for reading/writing per-agent memory.md.

Each agent has a memory file at:
    vault/agents/<agent_id>/memory.md

Sections (canonical order):
  ## Vetos pattern (cap 500, append-only, tail rotation)
  ## Losses pattern (cap 500)
  ## Soft rules (auto-generadas tras 3+ ocurrencias)
  ## Anti-patterns identificados

API:
  load_memory(agent_id)         -> MemoryView
  append_loss(agent_id, ...)
  append_veto(agent_id, ...)
  append_antipattern(agent_id, tag, description, run_id)
  generate_soft_rules(agent_id, min_freq=3)
  summarize_for_prompt(agent_id, max_chars=4000)  -> str injected into agent prompts

CLI:
  python3 memory_lib.py init <agent_id>       # bootstrap empty memory file
  python3 memory_lib.py soft-rules <agent_id> # regenerate soft rules
  python3 memory_lib.py summary <agent_id>    # print prompt-ready summary

NEVER call api.anthropic.com from here. Pure local I/O.
"""
from __future__ import annotations

import argparse
import dataclasses
import datetime as dt
import json
import os
import pathlib
import re
import sys
from collections import Counter
from typing import Any

VAULT_ROOT = pathlib.Path(os.environ.get("JARVIS_VAULT", "/home/agent/agent-stack/vault")).resolve()
CAP_VETOS = 500
CAP_LOSSES = 500
CAP_ANTIPATTERNS = 200


def memory_path(agent_id: str) -> pathlib.Path:
    return VAULT_ROOT / "agents" / agent_id / "memory.md"


def now_iso() -> str:
    return dt.datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%SZ")


# ----------------------------- Bootstrap -----------------------------

TEMPLATE = """---
title: "{agent_id} — memory"
type: agent-memory
agent: "[[agents/{agent_id}]]"
created: {date}
tags: [agent-memory, {agent_id}]
schema_version: 1
---

# {agent_id} — memoria persistente

> Léeme al arrancar cada turn. Escribe vía `memory_lib.py` (NO edición manual).

## Vetos pattern (cap {cap_v}, append-only, tail rotation)

| timestamp | task_type | input_pattern | rule | reason |
|---|---|---|---|---|

## Losses pattern (cap {cap_l})

| timestamp | task_id | input | output | mistake | lesson |
|---|---|---|---|---|---|

## Soft rules (auto-generadas tras 3+ ocurrencias de pattern)

| condition | action | confidence | uses |
|---|---|---|---|

## Anti-patterns identificados

"""


def init(agent_id: str) -> pathlib.Path:
    path = memory_path(agent_id)
    if path.exists():
        return path
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(TEMPLATE.format(
        agent_id=agent_id,
        date=dt.date.today().isoformat(),
        cap_v=CAP_VETOS,
        cap_l=CAP_LOSSES,
    ), encoding="utf-8")
    return path


# ----------------------------- Parser -----------------------------

SECTION_RE = re.compile(r"^## (.+)$", re.M)


@dataclasses.dataclass
class MemoryView:
    agent_id: str
    path: pathlib.Path
    raw: str
    sections: dict[str, str]  # section_name -> raw markdown body

    def section_row_count(self, name: str) -> int:
        body = self.sections.get(name, "")
        # Count markdown table data rows (skip header + separator).
        lines = [l for l in body.splitlines() if l.startswith("|") and not re.match(r"^\|[\s\-:|]+\|$", l)]
        return max(0, len(lines) - 1)


def load_memory(agent_id: str) -> MemoryView:
    path = memory_path(agent_id)
    if not path.exists():
        init(agent_id)
    raw = path.read_text(encoding="utf-8")
    sections: dict[str, str] = {}
    matches = list(SECTION_RE.finditer(raw))
    for i, m in enumerate(matches):
        name = m.group(1).strip()
        start = m.end()
        end = matches[i + 1].start() if i + 1 < len(matches) else len(raw)
        sections[name] = raw[start:end]
    return MemoryView(agent_id=agent_id, path=path, raw=raw, sections=sections)


# ----------------------------- Writers -----------------------------

def _append_table_row(agent_id: str, section_prefix: str, row_cells: list[str], cap: int) -> None:
    """Append a row to the matching section's table. Rotate (drop oldest) if over cap."""
    path = memory_path(agent_id)
    if not path.exists():
        init(agent_id)
    raw = path.read_text(encoding="utf-8")
    # Find section by prefix match (e.g., "Vetos pattern")
    matches = list(SECTION_RE.finditer(raw))
    sect_idx = None
    for i, m in enumerate(matches):
        if m.group(1).strip().startswith(section_prefix):
            sect_idx = i
            break
    if sect_idx is None:
        raise ValueError(f"section starting with '{section_prefix}' not found in {path}")

    start = matches[sect_idx].end()
    end = matches[sect_idx + 1].start() if sect_idx + 1 < len(matches) else len(raw)
    section_body = raw[start:end]

    lines = section_body.splitlines()
    table_lines = [i for i, l in enumerate(lines) if l.startswith("|")]
    if not table_lines:
        # No table yet — leave header missing, just write below
        new_row = "| " + " | ".join(_esc(c) for c in row_cells) + " |"
        new_body = section_body.rstrip() + "\n" + new_row + "\n\n"
    else:
        new_row = "| " + " | ".join(_esc(c) for c in row_cells) + " |"
        # Insert after last data row (i.e., append at end of table block)
        last_table_idx = table_lines[-1]
        lines.insert(last_table_idx + 1, new_row)
        # Rotate: if rows beyond header+separator exceed cap, drop oldest data rows
        header_idx = table_lines[0]
        sep_idx = table_lines[1] if len(table_lines) > 1 else None
        data_start = (sep_idx + 1) if sep_idx is not None else (header_idx + 1)
        data_indices = [i for i in range(data_start, len(lines)) if lines[i].startswith("|")]
        excess = len(data_indices) - cap
        if excess > 0:
            for drop_i in data_indices[:excess][::-1]:
                lines.pop(drop_i)
        new_body = "\n".join(lines)
        if not new_body.endswith("\n"):
            new_body += "\n"

    raw_new = raw[:start] + new_body + raw[end:]
    path.write_text(raw_new, encoding="utf-8")


def _esc(s: Any) -> str:
    return str(s).replace("|", "\\|").replace("\n", " ").strip()


def append_veto(agent_id: str, task_type: str, input_pattern: str, rule: str, reason: str) -> None:
    _append_table_row(agent_id, "Vetos pattern", [now_iso(), task_type, input_pattern, rule, reason], CAP_VETOS)


def append_loss(agent_id: str, task_id: str, input_: str, output: str, mistake: str, lesson: str) -> None:
    _append_table_row(agent_id, "Losses pattern", [now_iso(), task_id, input_, output, mistake, lesson], CAP_LOSSES)


def append_antipattern(agent_id: str, tag: str, description: str, run_id: str) -> None:
    path = memory_path(agent_id)
    if not path.exists():
        init(agent_id)
    raw = path.read_text(encoding="utf-8")
    matches = list(SECTION_RE.finditer(raw))
    target = None
    for i, m in enumerate(matches):
        if m.group(1).strip().startswith("Anti-patterns"):
            target = i
            break
    if target is None:
        raise ValueError(f"Anti-patterns section not found in {path}")
    start = matches[target].end()
    end = matches[target + 1].start() if target + 1 < len(matches) else len(raw)
    body = raw[start:end].rstrip()
    bullet = f"- **{tag}** ({now_iso()}, run {run_id}): {description}"
    new_body = body + "\n" + bullet + "\n\n"
    # Cap entries
    bullets = [l for l in new_body.splitlines() if l.startswith("- ")]
    if len(bullets) > CAP_ANTIPATTERNS:
        keep = bullets[-CAP_ANTIPATTERNS:]
        new_body = "\n".join(keep) + "\n\n"
    raw_new = raw[:start] + "\n" + new_body + raw[end:]
    path.write_text(raw_new, encoding="utf-8")


# ----------------------------- Soft rules -----------------------------

def generate_soft_rules(agent_id: str, min_freq: int = 3) -> int:
    """Scan losses + vetos, find recurring (task_type|rule) pairs, write to Soft rules table.
    Returns number of new soft rules added (idempotent)."""
    mv = load_memory(agent_id)

    def parse_table(section_name: str) -> list[list[str]]:
        body = mv.sections.get(section_name, "")
        rows: list[list[str]] = []
        for line in body.splitlines():
            if not line.startswith("|"):
                continue
            if re.match(r"^\|[\s\-:|]+\|$", line):
                continue
            cells = [c.strip() for c in line.strip("|").split("|")]
            rows.append(cells)
        # First row is header
        return rows[1:] if rows else []

    veto_rows = next((parse_table(k) for k in mv.sections if k.startswith("Vetos pattern")), [])
    loss_rows = next((parse_table(k) for k in mv.sections if k.startswith("Losses pattern")), [])

    counter: Counter = Counter()
    # Vetos: cols = [ts, task_type, input_pattern, rule, reason]
    for r in veto_rows:
        if len(r) >= 4:
            counter[("veto", r[1], r[3])] += 1
    # Losses: cols = [ts, task_id, input, output, mistake, lesson]
    for r in loss_rows:
        if len(r) >= 5:
            counter[("loss", r[1].split("/")[0] if r[1] else "", r[4])] += 1

    new_rules: list[tuple[str, str, str, int]] = []
    for (kind, key, sig), freq in counter.items():
        if freq < min_freq:
            continue
        if kind == "veto":
            condition = f"task_type={key}, signal matches {sig}"
            action = f"apply veto rule {sig} preemptively"
        else:
            condition = f"task_class={key}, mistake matches '{sig}'"
            action = f"avoid mistake '{sig}'; cite this rule in plan"
        confidence = "medium" if freq < 6 else "high"
        new_rules.append((condition, action, confidence, freq))

    # Read existing soft rules to avoid duplicates
    existing_rows = next((parse_table(k) for k in mv.sections if k.startswith("Soft rules")), [])
    existing_keys = {(r[0], r[1]) for r in existing_rows if len(r) >= 2}

    added = 0
    for cond, act, conf, freq in new_rules:
        if (cond, act) in existing_keys:
            continue
        _append_table_row(agent_id, "Soft rules", [cond, act, conf, str(freq)], cap=200)
        added += 1
    return added


# ----------------------------- Prompt summarizer -----------------------------

def summarize_for_prompt(agent_id: str, max_chars: int = 4000) -> str:
    """Returns a compact summary suitable for injecting into the agent's system prompt."""
    path = memory_path(agent_id)
    if not path.exists():
        return ""
    mv = load_memory(agent_id)

    parts: list[str] = [f"### Memoria persistente ({agent_id})"]

    def tail_rows(section_prefix: str, n: int) -> list[str]:
        body = next((mv.sections[k] for k in mv.sections if k.startswith(section_prefix)), "")
        rows = [l for l in body.splitlines() if l.startswith("|") and not re.match(r"^\|[\s\-:|]+\|$", l)]
        # rows[0] is header
        return rows[-n:] if len(rows) > 1 else []

    # Soft rules first (most actionable)
    sr = tail_rows("Soft rules", 15)
    if sr:
        parts.append("**Soft rules activas:**")
        parts.extend(sr)

    # Recent losses (last 5)
    lo = tail_rows("Losses pattern", 5)
    if lo:
        parts.append("**Últimos errores (no repetir):**")
        parts.extend(lo)

    # Anti-patterns (last 10 bullets)
    ap_body = next((mv.sections[k] for k in mv.sections if k.startswith("Anti-patterns")), "")
    bullets = [l for l in ap_body.splitlines() if l.startswith("- ")]
    if bullets:
        parts.append("**Anti-patterns vivos:**")
        parts.extend(bullets[-10:])

    out = "\n".join(parts)
    if len(out) > max_chars:
        out = out[: max_chars - 10] + "\n[...]"
    return out


# ----------------------------- CLI -----------------------------

def main(argv: list[str] | None = None) -> int:
    p = argparse.ArgumentParser(prog="memory_lib")
    sub = p.add_subparsers(dest="cmd", required=True)
    sp = sub.add_parser("init"); sp.add_argument("agent_id")
    sp = sub.add_parser("soft-rules"); sp.add_argument("agent_id"); sp.add_argument("--min-freq", type=int, default=3)
    sp = sub.add_parser("summary"); sp.add_argument("agent_id"); sp.add_argument("--max-chars", type=int, default=4000)
    sp = sub.add_parser("init-all")
    args = p.parse_args(argv)

    if args.cmd == "init":
        path = init(args.agent_id)
        print(f"initialized {path}")
    elif args.cmd == "soft-rules":
        n = generate_soft_rules(args.agent_id, min_freq=args.min_freq)
        print(f"added {n} new soft rules to {args.agent_id}")
    elif args.cmd == "summary":
        print(summarize_for_prompt(args.agent_id, max_chars=args.max_chars))
    elif args.cmd == "init-all":
        agents = ["main","planner","apier","archivist","auditor","code-reviewer","debugger",
                  "designer","documenter","job-hunter","monitor","researcher","skill-reviewer",
                  "tester","polymarket-handler"]
        for a in agents:
            path = init(a)
            print(f"  {a}: {path}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
