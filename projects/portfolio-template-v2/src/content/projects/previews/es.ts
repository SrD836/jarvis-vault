import thumbJarvis from "../../../assets/thumbnails/jarvis-multi-agent.webp";
import thumbPoly from "../../../assets/thumbnails/polymarket-veto-loop-bot.webp";
import thumbTrabajo from "../../../assets/thumbnails/trabajo-auto-apply.webp";
import thumbDashboard from "../../../assets/thumbnails/jarvis-dashboard.webp";

import type { ProjectPreview } from "../../types";

export default [
  {
    title: "JARVIS · Sistema multi-agente",
    slug: "jarvis-multi-agent",
    thumbnail: thumbJarvis,
    description: "Asistente personal autónomo en VPS",
  },
  {
    title: "Polymarket Veto-Loop Bot",
    slug: "polymarket-veto-loop-bot",
    thumbnail: thumbPoly,
    description: "Bot de trading 24/7 con memoria",
  },
  {
    title: "Trabajo · LinkedIn Easy Apply",
    slug: "trabajo-auto-apply",
    thumbnail: thumbTrabajo,
    description: "Pipeline búsqueda + CV + auto-apply",
  },
  {
    title: "Dashboard JARVIS",
    slug: "jarvis-dashboard",
    thumbnail: thumbDashboard,
    description: "Panel de control web del sistema",
  },
] as const satisfies ProjectPreview[];
