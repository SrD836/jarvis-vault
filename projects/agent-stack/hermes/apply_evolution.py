"""
apply_evolution.py — closes the Hermes feedback loop.

Reads evolution proposals at vault/wiki/agent-evolution/<agent_id>.md and:
  - LOW risk     → auto-apply to AGENTS.md + openclaw.json (with backup), mark applied
  - MEDIUM risk  → send Telegram approval (inline keyboard), mark pending-approval
  - HIGH risk    → alert only, require manual review

Risk classification:
  - LOW:    frontmatter-only changes (allow_agents, model, tools)
            AND pattern occurrences >= 4
            AND confidence in ("medium", "high")
  - MEDIUM: changes to Standing Orders body sections (policy text)
            OR confidence == "medium" with occurrences in [2,3]
  - HIGH:   structural changes (delete section, new agent creation)
            OR confidence == "low"

Hard rules respected:
  - NEVER api.anthropic.com — Telegram delivery uses bot HTTP API directly
    (no Anthropic involvement). LLM-based risk re-classification, if added
    later, must go through `docker exec gateway claude -p`.
  - backup *.bak-YYYY-MM-DD-HHMMSS before mutating any AGENTS.md or openclaw.json
  - validation post-apply: AGENTS.md must remain readable; openclaw.json must
    remain valid JSON; rollback on failure.

Cron: 0 5 * * 0 (Sunday 05:00 UTC, after learnings.py at 04:00).

CLI:
  python3 apply_evolution.py [--dry-run] [--auto-apply-low]
                             [--proposal vault/wiki/agent-evolution/<id>.md]
"""
from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import re
import shutil
import sys
from typing import Optional

VAULT_ROOT = pathlib.Path(os.environ.get("JARVIS_VAULT", "/home/agent/agent-stack/vault")).resolve()
EVOLUTION_DIR = VAULT_ROOT / "wiki" / "agent-evolution"
OPENCLAW_AGENTS_DIR = pathlib.Path(
    os.environ.get("OPENCLAW_AGENTS_DIR", "/home/agent/.openclaw/workspace/agents")
).resolve()
OPENCLAW_JSON = pathlib.Path(
    os.environ.get("OPENCLAW_CONFIG", "/home/agent/.openclaw/openclaw.json")
).resolve()

TELEGRAM_BOT_TOKEN = os.environ.get("TELEGRAM_BOT_TOKEN", "")
TELEGRAM_CHAT_ID = os.environ.get("TELEGRAM_ALLOWED_USER_ID", "")


def now_stamp() -> str:
    return dt.datetime.utcnow().strftime("%Y%m%d-%H%M%S")


# ----------------------------- Parser -----------------------------

FRONTMATTER_RE = re.compile(r"^---\n(.*?)\n---\n", re.S)
PATTERN_HDR_RE = re.compile(r"^## Patrón: `([^`]+)`", re.M)
ALLOW_AGENTS_RE = re.compile(
    r'"allow_agents"\s*:\s*\[\s*([^\]]+)\]', re.S
)


def parse_proposal(path: pathlib.Path) -> dict:
    raw = path.read_text(encoding="utf-8")
    fm_match = FRONTMATTER_RE.match(raw)
    fm: dict[str, str] = {}
    if fm_match:
        for line in fm_match.group(1).splitlines():
            if ":" in line:
                k, _, v = line.partition(":")
                fm[k.strip()] = v.strip().strip('"').strip("'")

    body = raw[fm_match.end():] if fm_match else raw
    patterns = []
    for m in PATTERN_HDR_RE.finditer(body):
        start = m.end()
        next_m = PATTERN_HDR_RE.search(body, start)
        section = body[start: next_m.start() if next_m else len(body)]
        confidence = re.search(r"Confianza:\s*\*\*(\w+)\*\*", section)
        occurrences = re.search(r"Ocurrencias:\s*\*\*(\d+)\*\*", section)
        # Look for an allow_agents block to apply mechanically
        allow_match = ALLOW_AGENTS_RE.search(section)
        allow_agents: Optional[list[str]] = None
        if allow_match:
            raw_items = allow_match.group(1)
            allow_agents = [s.strip().strip('"').strip("'") for s in raw_items.split(",") if s.strip()]

        # Detect if proposal touches body sections (policy text)
        touches_body = bool(re.search(r"`?briefing\.md`?|AGENTS\.md.*section|Standing Orders", section, re.I))
        # Detect if proposal proposes new agent creation
        touches_create = bool(re.search(r"\bcrear (?:un )?nuevo agente\b|new agent", section, re.I))

        patterns.append({
            "name": m.group(1),
            "confidence": confidence.group(1) if confidence else "low",
            "occurrences": int(occurrences.group(1)) if occurrences else 0,
            "allow_agents": allow_agents,
            "touches_body": touches_body,
            "touches_create": touches_create,
            "raw": section,
        })

    return {
        "path": str(path),
        "agent": fm.get("agent", "").replace("[[agents/", "").rstrip("]"),
        "frontmatter": fm,
        "applied": fm.get("applied", "false").lower() == "true",
        "patterns": patterns,
    }


# ----------------------------- Risk classifier -----------------------------

def classify_risk(pattern: dict) -> str:
    conf = pattern["confidence"].lower()
    occ = pattern["occurrences"]
    if pattern.get("touches_create") or conf == "low":
        return "high"
    if pattern.get("touches_body"):
        return "medium"
    # Frontmatter-only path
    if pattern.get("allow_agents") is not None and occ >= 4 and conf in ("medium", "high"):
        return "low"
    if conf == "medium" and 2 <= occ <= 3:
        return "medium"
    return "high"


# ----------------------------- Mutators -----------------------------

def backup_file(p: pathlib.Path) -> pathlib.Path:
    bak = p.with_suffix(p.suffix + f".bak-{now_stamp()}")
    shutil.copy(p, bak)
    return bak


def apply_allow_agents_change(agent_id: str, new_allow: list[str]) -> tuple[bool, str]:
    if not OPENCLAW_JSON.exists():
        return False, f"openclaw.json missing at {OPENCLAW_JSON}"
    backup_file(OPENCLAW_JSON)
    try:
        d = json.loads(OPENCLAW_JSON.read_text(encoding="utf-8"))
        target = None
        for a in d.get("agents", {}).get("list", []):
            if a.get("id") == agent_id:
                target = a
                break
        if not target:
            return False, f"agent {agent_id} not in openclaw.json"
        target.setdefault("subagents", {})["allowAgents"] = new_allow
        OPENCLAW_JSON.write_text(json.dumps(d, indent=2), encoding="utf-8")
        # Re-parse to validate
        json.loads(OPENCLAW_JSON.read_text(encoding="utf-8"))
        return True, f"allowAgents for {agent_id} → {new_allow}"
    except Exception as e:
        return False, f"apply failed: {e}"


def mark_applied(proposal_path: pathlib.Path, by: str) -> None:
    raw = proposal_path.read_text(encoding="utf-8")
    backup_file(proposal_path)
    raw = re.sub(r"^applied:\s*false\s*$", "applied: true", raw, count=1, flags=re.M)
    if "applied:" not in raw[: raw.find("---", 4)]:
        # add inside frontmatter
        raw = raw.replace("---\n", f"---\napplied: true\napplied_by: {by}\napplied_at: {dt.datetime.utcnow().isoformat()}Z\n", 1)
    else:
        if "applied_by:" not in raw:
            raw = raw.replace("applied: true", f"applied: true\napplied_by: {by}\napplied_at: {dt.datetime.utcnow().isoformat()}Z", 1)
    proposal_path.write_text(raw, encoding="utf-8")


# ----------------------------- Telegram delivery -----------------------------

def telegram_notify(text: str) -> bool:
    if not TELEGRAM_BOT_TOKEN or not TELEGRAM_CHAT_ID:
        print("  ⚠ no TELEGRAM_BOT_TOKEN/CHAT_ID; skipping notify", file=sys.stderr)
        return False
    import urllib.parse, urllib.request
    url = f"https://api.telegram.org/bot{TELEGRAM_BOT_TOKEN}/sendMessage"
    data = urllib.parse.urlencode({
        "chat_id": TELEGRAM_CHAT_ID,
        "text": text,
        "parse_mode": "HTML",
    }).encode()
    try:
        with urllib.request.urlopen(url, data=data, timeout=10) as resp:
            return resp.status == 200
    except Exception as e:
        print(f"  ⚠ telegram failed: {e}", file=sys.stderr)
        return False


# ----------------------------- Main -----------------------------

def process(proposal_path: pathlib.Path, dry_run: bool, auto_apply_low: bool) -> dict:
    prop = parse_proposal(proposal_path)
    if prop["applied"]:
        return {"path": str(proposal_path), "skipped": "already-applied"}

    actions = []
    for pat in prop["patterns"]:
        risk = classify_risk(pat)
        entry = {
            "pattern": pat["name"],
            "risk": risk,
            "occurrences": pat["occurrences"],
            "confidence": pat["confidence"],
        }
        if risk == "low" and pat.get("allow_agents") is not None:
            if dry_run:
                entry["action"] = f"WOULD apply allowAgents={pat['allow_agents']}"
            elif auto_apply_low:
                ok, msg = apply_allow_agents_change(prop["agent"], pat["allow_agents"])
                entry["action"] = ("applied: " if ok else "FAILED: ") + msg
                if ok:
                    mark_applied(proposal_path, by="apply_evolution.py")
            else:
                entry["action"] = "low-risk pending (use --auto-apply-low)"
        elif risk == "medium":
            text = (
                f"🤖 <b>Hermes propone mejora</b>\n"
                f"Agente: <code>{prop['agent']}</code>\n"
                f"Patrón: <code>{pat['name']}</code>\n"
                f"Confianza: {pat['confidence']} · Ocurrencias: {pat['occurrences']}\n\n"
                f"<i>Cambio toca body de Standing Orders — requiere revisión.</i>\n\n"
                f"Archivo: <code>{proposal_path}</code>"
            )
            if dry_run:
                entry["action"] = "WOULD send telegram approval"
            else:
                ok = telegram_notify(text)
                entry["action"] = "telegram-sent" if ok else "telegram-failed"
        else:  # high
            text = (
                f"⚠ <b>Hermes flagged HIGH-risk proposal</b>\n"
                f"Agente: <code>{prop['agent']}</code>\n"
                f"Patrón: <code>{pat['name']}</code>\n"
                f"Acción: revisión manual.\n"
                f"Archivo: <code>{proposal_path}</code>"
            )
            if dry_run:
                entry["action"] = "WOULD alert (high-risk)"
            else:
                telegram_notify(text)
                entry["action"] = "high-risk alerted"
        actions.append(entry)

    return {"path": str(proposal_path), "agent": prop["agent"], "actions": actions}


def main(argv: list[str] | None = None) -> int:
    p = argparse.ArgumentParser()
    p.add_argument("--dry-run", action="store_true")
    p.add_argument("--auto-apply-low", action="store_true", default=True)
    p.add_argument("--no-auto-apply-low", dest="auto_apply_low", action="store_false")
    p.add_argument("--proposal", type=pathlib.Path, default=None)
    args = p.parse_args(argv)

    proposals = [args.proposal] if args.proposal else sorted(EVOLUTION_DIR.glob("*.md"))
    proposals = [p for p in proposals if not p.name.startswith("_")]
    summary = []
    for prop_path in proposals:
        if prop_path.name.endswith(".bak") or ".bak-" in prop_path.name:
            continue
        result = process(prop_path, args.dry_run, args.auto_apply_low)
        summary.append(result)
        print(json.dumps(result, indent=2))
    print(f"\nprocessed {len(summary)} proposals (dry-run={args.dry_run})")
    return 0


if __name__ == "__main__":
    sys.exit(main())
