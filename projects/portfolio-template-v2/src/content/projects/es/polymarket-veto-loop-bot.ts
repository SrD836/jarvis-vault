import type { ProjectContent } from "../../types";

export default {
  title: "Polymarket Veto-Loop Bot",
  theme: "dark",
  tags: ["typescript", "node", "vue"],
  
  description: "Bot de trading 24/7 sobre Polymarket. Pipeline scanner → brain → executor → exit_monitor. Memoria persistente de wins/losses, auto-reglas blandas a partir de clustering de patrones.",
  components: [],
} as const satisfies ProjectContent;
