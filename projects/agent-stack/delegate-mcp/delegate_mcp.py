"""
delegate-mcp — minimal MCP HTTP server exposing one tool `delegate(agent_id, task)`.

Closes the gap identified in audit 2026-05-26: the openclaw bundle does not
expose a delegation tool, so the configured multi-agent hierarchy
(main→planner→specialists) was non-functional. This server wraps
`openclaw agent --agent <id>` so any allowed agent can be invoked by another.

Hard rules respected:
  - NEVER api.anthropic.com — invokes the openclaw CLI which goes through
    Claude Max OAuth in the gateway container.
  - Authority gate: caller (from x-openclaw-agent-id header) must have target
    in its subagents.allowAgents per openclaw.json.
  - Run logging: every delegation writes a markdown trace to
    vault/agents/<target>/runs/<YYYY-MM-DD>/<HHMMSS>-from-<caller>.md mirroring
    the historical run-logger schema (jarvis_run_logger_schema memory).

MCP protocol: 2024-11-05 streamable-http transport.
"""
from __future__ import annotations

import asyncio
import datetime as dt
import json
import logging
import os
import pathlib
import subprocess
import time
from typing import Any

from aiohttp import web

OPENCLAW_JSON = pathlib.Path(os.environ.get("OPENCLAW_JSON", "/home/agent/.openclaw/openclaw.json"))
VAULT_AGENTS = pathlib.Path(os.environ.get("VAULT_AGENTS", "/home/agent/agent-stack/vault/agents"))
GATEWAY_CONTAINER = os.environ.get("GATEWAY_CONTAINER", "openclaw-fork-openclaw-gateway-1")
DOCKER_BIN = os.environ.get("DOCKER_BIN", "/usr/bin/docker")
PORT = int(os.environ.get("PORT", "38900"))
MAX_TIMEOUT = int(os.environ.get("MAX_TIMEOUT", "900"))

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")
log = logging.getLogger("delegate-mcp")


# ----------------------------- helpers -----------------------------

def _utc() -> str:
    return dt.datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%SZ")


def _err(rid: Any, message: str, code: int = -32000) -> web.Response:
    return web.json_response({
        "jsonrpc": "2.0",
        "id": rid,
        "error": {"code": code, "message": message},
    })


def _ok(rid: Any, result: dict) -> web.Response:
    return web.json_response({"jsonrpc": "2.0", "id": rid, "result": result})


def _extract_text(out: dict) -> str:
    """Extract final assistant text from `openclaw agent --json` output.

    Known shape (openclaw 2026.5.12-beta.1):
        {runId, status, summary, result: {payloads: [{text, mediaUrl}], meta: {...}}}
    Fallbacks added defensively in case the shape varies.
    """
    # 1. Canonical path
    result = out.get("result")
    if isinstance(result, dict):
        payloads = result.get("payloads")
        if isinstance(payloads, list) and payloads:
            texts = [p.get("text", "") for p in payloads if isinstance(p, dict) and p.get("text")]
            if texts:
                return "\n".join(t for t in texts if t)
        meta = result.get("meta")
        if isinstance(meta, dict):
            for k in ("finalAssistantVisibleText", "finalAssistantRawText"):
                v = meta.get(k)
                if isinstance(v, str) and v.strip():
                    return v

    # 2. Legacy/alternative shapes
    data = out.get("data") or out
    for k in ("response", "text"):
        v = data.get(k)
        if isinstance(v, str):
            return v
    msg = data.get("message")
    if isinstance(msg, dict):
        for kk in ("text", "content"):
            vv = msg.get(kk)
            if isinstance(vv, str):
                return vv
            if isinstance(vv, list):
                parts = [p.get("text", "") for p in vv if isinstance(p, dict)]
                if parts:
                    return "\n".join(parts)

    # 3. Last resort: short stringified preview
    return f"[delegate-mcp: could not extract text; raw keys={list(out.keys())}]"


def _extract_meta(out: dict) -> tuple[str, str]:
    """Return (model_used, session_id)."""
    result = out.get("result")
    if isinstance(result, dict):
        meta = result.get("meta") or {}
        spr = meta.get("systemPromptReport") or {}
        return spr.get("model", "unknown"), spr.get("sessionId", "")
    # Legacy
    am = (out.get("data") or {}).get("agentMeta") or {}
    return am.get("model", "unknown"), am.get("sessionId", "")


def _log_delegation(caller: str, target: str, task: str, response: str,
                    latency_ms: int, model_used: str, session_id: str) -> pathlib.Path:
    now = dt.datetime.utcnow()
    date_dir = VAULT_AGENTS / target / "runs" / now.strftime("%Y-%m-%d")
    date_dir.mkdir(parents=True, exist_ok=True)
    fname = f"{now.strftime('%H%M%S')}-from-{caller}.md"
    p = date_dir / fname
    body = (
        "---\n"
        f"caller: {caller}\n"
        f"target: {target}\n"
        f"ts: {_utc()}\n"
        f"session_id: {session_id}\n"
        f"model_used: {model_used}\n"
        f"latency_ms: {latency_ms}\n"
        f"spawned_children: []\n"
        f"source: delegate-mcp\n"
        "---\n\n"
        "## Task\n\n"
        f"{task}\n\n"
        "## Response\n\n"
        f"{response}\n"
    )
    p.write_text(body, encoding="utf-8")
    return p


def _load_caller_allowlist(caller: str) -> list[str] | None:
    """Returns the allowAgents list for caller, or None if caller not configured."""
    try:
        cfg = json.loads(OPENCLAW_JSON.read_text(encoding="utf-8"))
    except Exception as e:
        log.error("could not read openclaw.json: %s", e)
        return None
    for a in cfg.get("agents", {}).get("list", []):
        if a.get("id") == caller:
            return a.get("subagents", {}).get("allowAgents", []) or []
    return None


async def _run_target_agent(target: str, task: str, timeout: int) -> dict:
    """Invoke `openclaw agent` in the gateway container. Async via subprocess."""
    cmd = [
        DOCKER_BIN, "exec", GATEWAY_CONTAINER,
        "openclaw", "agent",
        "--agent", target,
        "--message", task,
        "--thinking", "low",
        "--timeout", str(timeout),
        "--json",
    ]
    log.info("delegating: caller→%s timeout=%ds", target, timeout)
    start = time.time()
    proc = await asyncio.create_subprocess_exec(
        *cmd,
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )
    try:
        stdout, stderr = await asyncio.wait_for(proc.communicate(), timeout=timeout + 30)
    except asyncio.TimeoutError:
        proc.kill()
        return {"_error": f"timeout after {timeout}s"}
    latency_ms = int((time.time() - start) * 1000)
    if proc.returncode != 0:
        return {"_error": f"exit {proc.returncode}; stderr={stderr.decode()[:500]}"}
    # Parse last JSON line of stdout
    raw = stdout.decode().strip()
    last_brace = raw.rfind("\n{")
    json_payload = raw[last_brace + 1:] if last_brace >= 0 else raw
    try:
        out = json.loads(json_payload)
    except json.JSONDecodeError as e:
        return {"_error": f"parse failure: {e}; raw_tail={raw[-500:]}"}
    out["_latency_ms"] = latency_ms
    return out


# ----------------------------- handlers -----------------------------

async def handle_mcp(req: web.Request) -> web.Response:
    try:
        body = await req.json()
    except Exception:
        return _err(None, "invalid JSON body", -32700)

    method = body.get("method")
    rid = body.get("id")
    caller = req.headers.get("x-openclaw-agent-id") or req.headers.get("X-Openclaw-Agent-Id") or "unknown"

    if method == "initialize":
        return _ok(rid, {
            "protocolVersion": "2024-11-05",
            "capabilities": {"tools": {}},
            "serverInfo": {"name": "delegate", "version": "0.1.0"},
        })

    if method == "notifications/initialized":
        return web.Response(status=204)

    if method == "tools/list":
        return _ok(rid, {
            "tools": [{
                "name": "delegate",
                "description": (
                    "Delegate a task to another openclaw agent and wait for its response. "
                    "The caller (identified via x-openclaw-agent-id header) must have the target "
                    "in its subagents.allowAgents whitelist per openclaw.json. "
                    "Use this to invoke specialists: e.g. main→planner for complex tasks, "
                    "main→polymarket-handler for bot queries, planner→researcher for investigation."
                ),
                "inputSchema": {
                    "type": "object",
                    "required": ["agent_id", "task"],
                    "properties": {
                        "agent_id": {
                            "type": "string",
                            "description": "Target agent id (e.g. 'planner', 'researcher', 'polymarket-handler')",
                        },
                        "task": {
                            "type": "string",
                            "description": "Specific task instructions for the target agent. Include context: goal, constraints, expected output format.",
                        },
                        "timeout": {
                            "type": "integer",
                            "description": f"Max seconds (default 600, cap {MAX_TIMEOUT})",
                            "default": 600,
                        },
                    },
                },
            }]
        })

    if method == "tools/call":
        params = body.get("params") or {}
        name = params.get("name")
        args = params.get("arguments") or {}

        if name != "delegate":
            return _err(rid, f"unknown tool: {name}", -32601)

        target = (args.get("agent_id") or "").strip()
        task = (args.get("task") or "").strip()
        timeout = min(max(int(args.get("timeout", 600) or 600), 30), MAX_TIMEOUT)

        if not target or not task:
            return _err(rid, "agent_id and task are required")

        if target == caller:
            return _err(rid, f"refusing self-delegation ({caller} → {target})")

        allow = _load_caller_allowlist(caller)
        if allow is None:
            return _err(rid, f"caller {caller!r} not registered in openclaw.json agents.list")
        if target not in allow:
            return _err(rid, f"{caller} not authorized to delegate to {target}; allowAgents={allow}")

        out = await _run_target_agent(target, task, timeout)
        if "_error" in out:
            return _ok(rid, {
                "content": [{"type": "text", "text": f"[delegation failed] {out['_error']}"}],
                "isError": True,
            })

        text = _extract_text(out)
        model_used, session_id = _extract_meta(out)
        latency_ms = out.get("_latency_ms", 0)

        try:
            run_path = _log_delegation(caller, target, task, text, latency_ms, model_used, session_id)
            log.info("delegation logged → %s", run_path)
        except Exception as e:
            log.warning("run log failed: %s", e)
            run_path = None

        return _ok(rid, {
            "content": [{"type": "text", "text": text}],
            "isError": False,
            "_meta": {
                "agent": target,
                "model_used": model_used,
                "latency_ms": latency_ms,
                "session_id": session_id,
                "run_log": str(run_path) if run_path else None,
            },
        })

    return _err(rid, f"unknown method: {method}", -32601)


async def handle_health(req: web.Request) -> web.Response:
    return web.json_response({"status": "ok", "ts": _utc()})


def make_app() -> web.Application:
    app = web.Application()
    app.router.add_post("/mcp", handle_mcp)
    app.router.add_get("/health", handle_health)
    return app


if __name__ == "__main__":
    log.info("delegate-mcp starting on :%d  gateway=%s  openclaw=%s", PORT, GATEWAY_CONTAINER, OPENCLAW_JSON)
    web.run_app(make_app(), host="0.0.0.0", port=PORT, access_log=None)
