#!/bin/bash
# Re-evaluación programada del veredicto de expectancy del polymarket-bot.
# Agendada 2026-05-30 (autorizada por David) para el 2026-06-13, cuando el libro
# longshot pre-R1 ya haya resuelto y la ventana forward con R1/R3 tenga datos limpios.
# Guarda de fecha: solo actúa el día objetivo (crontab no fija el año).
set -uo pipefail
TARGET="2026-06-13"
[ "$(date -u +%Y-%m-%d)" = "$TARGET" ] || exit 0

cd /home/agent/agent-stack || exit 1
BOT=vault/projects/polymarket-veto-loop-bot
./$BOT/bot/bin/mark_to_market >/dev/null 2>&1
REPORT=$(/usr/bin/python3 $BOT/analytics/expectancy_report.py --no-write 2>/dev/null)

/usr/bin/python3 - "$REPORT" <<'PY'
import sys, os
sys.path.insert(0, "hermes")
import llm_call
env = llm_call.load_env()
chat = env.get("TELEGRAM_ALLOWED_USER_ID") or os.environ.get("TELEGRAM_ALLOWED_USER_ID")
body = sys.argv[1][:3500] if len(sys.argv) > 1 else "(sin reporte)"
msg = ("\U0001F514 RE-EVALUACION PROGRAMADA (2026-06-13) — expectancy v7\n\n"
       "El libro longshot pre-R1 ya deberia haber resuelto. Veredicto actual:\n\n" + body)
try:
    llm_call.send_message(chat, msg, env)
except Exception as e:
    sys.stderr.write(f"reeval telegram fallo: {e}\n")
PY
