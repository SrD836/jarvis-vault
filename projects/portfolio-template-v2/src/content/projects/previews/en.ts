import thumbJarvis from "../../../assets/thumbnails/jarvis-multi-agent.webp";
import thumbPoly from "../../../assets/thumbnails/polymarket-veto-loop-bot.webp";
import thumbTrabajo from "../../../assets/thumbnails/trabajo-auto-apply.webp";
import thumbDashboard from "../../../assets/thumbnails/jarvis-dashboard.webp";

import type { ProjectPreview } from "../../types";

export default [
  {
    title: "JARVIS · Multi-agent system",
    slug: "jarvis-multi-agent",
    thumbnail: thumbJarvis,
    description: "24/7 autonomous personal assistant on VPS",
  },
  {
    title: "Polymarket Veto-Loop Bot",
    slug: "polymarket-veto-loop-bot",
    thumbnail: thumbPoly,
    description: "24/7 trading bot with persistent memory",
  },
  {
    title: "Trabajo · LinkedIn Easy Apply",
    slug: "trabajo-auto-apply",
    thumbnail: thumbTrabajo,
    description: "Search + CV adapt + auto-apply pipeline",
  },
  {
    title: "Dashboard JARVIS",
    slug: "jarvis-dashboard",
    thumbnail: thumbDashboard,
    description: "Web control panel for the system",
  },
] as const satisfies ProjectPreview[];
