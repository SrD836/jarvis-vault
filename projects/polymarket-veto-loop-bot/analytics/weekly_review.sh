#!/usr/bin/env bash
# Weekly performance review for polymarket-veto-loop-bot.
# Cron: domingo 19:00 UTC. Synth via dashboard /api/llm/call provider=claudemax.
#
# Hard rule (jarvis_never_anthropic_api): NEVER reach api.anthropic.com with
# ANTHROPIC_API_KEY. Always go through the in-cluster dashboard which shells
# out to `docker exec openclaw-fork-openclaw-gateway-1 claude -p` under the
# user's paid Claude Max OAuth.
set -euo pipefail

ROOT="${POLYBOT_ROOT:-/home/agent/agent-stack}"
DASHBOARD_URL="${DASHBOARD_URL:-http://jarvis-dashboard:3000}"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROMPT_FILE="$SCRIPT_DIR/weekly_review_prompt.md"

cd "$ROOT"

WEEK=$(date -u +%G-W%V)
OUT="vault/analysis/performance/weekly/${WEEK}.md"
mkdir -p "$(dirname "$OUT")"

if [[ ! -f vault/inbox/trading/closed.jsonl ]]; then
  echo "weekly_review: closed.jsonl not found at $ROOT/vault/inbox/trading/closed.jsonl" >&2
  exit 1
fi

SINCE=$(date -u -d '7 days ago' +%Y-%m-%dT%H:%M:%SZ)
CTX=$(mktemp)
trap 'rm -f "$CTX"' EXIT

{
  echo "## Aggregate stats últimos 7d (closed.jsonl)"
  echo ""
  jq -s --arg s "$SINCE" '
    [.[] | select(.exit_timestamp >= $s)] |
    {
      total: length,
      wins: ([.[] | select(.pnl > 0)] | length),
      losses: ([.[] | select(.pnl <= 0)] | length),
      total_pnl: ([.[] | .pnl] | add),
      by_horizon: (group_by(.horizon) | map({h: .[0].horizon, n: length, pnl: ([.[] | .pnl] | add), winrate: (([.[] | select(.pnl>0)] | length) / length)})),
      by_category_worst: (group_by(.category) | map({c: .[0].category, n: length, pnl: ([.[] | .pnl] | add), winrate: (([.[] | select(.pnl>0)] | length) / length)}) | sort_by(.pnl) | .[0:10])
    }
  ' vault/inbox/trading/closed.jsonl || true
  echo ""
  echo "## Últimos 30 cierres (más recientes)"
  echo ""
  jq -c --arg s "$SINCE" 'select(.exit_timestamp >= $s)' vault/inbox/trading/closed.jsonl | tail -30 || true
  echo ""
  echo "## Portfolio snapshot (portfolio.json)"
  echo ""
  if [[ -f vault/inbox/trading/portfolio.json ]]; then
    jq '{bankroll, used_exposure, n_positions: (.positions | length), by_horizon: (.positions | group_by(.horizon) | map({h: .[0].horizon, n: length}))}' vault/inbox/trading/portfolio.json
  fi
} > "$CTX"

REQ_BODY=$(jq -n \
  --arg model "claudemax/claude-sonnet-4-6" \
  --rawfile sys "$PROMPT_FILE" \
  --rawfile usr "$CTX" \
  '{model:$model, system:$sys, messages:[{role:"user",content:$usr}], max_tokens:2500}')

RESP=$(curl -sS -X POST "$DASHBOARD_URL/api/llm/call" \
  -H 'Content-Type: application/json' \
  --data-binary @<(printf %s "$REQ_BODY"))

# Dashboard's /api/llm/call returns either {"text":"..."} or {"content":[{"text":"..."}]}.
# Try both shapes.
TEXT=$(echo "$RESP" | jq -r '
  if .text then .text
  elif .content and (.content | type == "array") then (.content[0].text // "")
  elif .message and .message.content then .message.content
  else .
  end' 2>/dev/null || true)

if [[ -z "$TEXT" || "$TEXT" == "null" ]]; then
  echo "weekly_review: empty response from $DASHBOARD_URL/api/llm/call" >&2
  echo "raw response:" >&2
  echo "$RESP" >&2
  exit 1
fi

printf '%s\n' "$TEXT" > "$OUT"
echo "weekly_review: wrote $OUT ($(wc -c < "$OUT") bytes)"
