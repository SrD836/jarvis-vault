import type { ProjectContent } from "../../types";

export default {
  title: "Polymarket Veto-Loop Bot",
  theme: "dark",
  tags: ["typescript", "node", "vue"],
  
  description: "24/7 trading bot on Polymarket. Pipeline scanner → brain → executor → exit_monitor. Persistent win/loss memory, auto-generated soft rules from pattern clustering.",
  components: [],
} as const satisfies ProjectContent;
