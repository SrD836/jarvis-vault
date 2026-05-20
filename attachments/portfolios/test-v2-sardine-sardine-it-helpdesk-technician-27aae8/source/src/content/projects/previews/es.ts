import thumb_jarvis_multi_agent from "../../../assets/thumbnails/jarvis-multi-agent.webp";
import thumb_trabajo_auto_apply from "../../../assets/thumbnails/trabajo-auto-apply.webp";
import thumb_jarvis_dashboard from "../../../assets/thumbnails/jarvis-dashboard.webp";
import thumb_polymarket_veto_loop_bot from "../../../assets/thumbnails/polymarket-veto-loop-bot.webp";

import type { ProjectPreview } from "../../types";

export default [
  {
    title: "JARVIS · Infraestructura autónoma en producción",
    slug: "jarvis-multi-agent",
    thumbnail: thumb_jarvis_multi_agent,
    description: "Sistema desplegado y operado 24/7 en VPS Hetzner bajo Linux, gestionado y monito",
  },
  {
    title: "Trabajo · Pipeline de automatización de procesos",
    slug: "trabajo-auto-apply",
    thumbnail: thumb_trabajo_auto_apply,
    description: "Pipeline end-to-end con gates de aprobación que automatiza búsqueda, adaptación ",
  },
  {
    title: "Dashboard JARVIS · Panel centralizado de sistemas",
    slug: "jarvis-dashboard",
    thumbnail: thumb_jarvis_dashboard,
    description: "Panel web centralizado para gestión de agentes, monitorización de servicios y co",
  },
  {
    title: "Polymarket Bot · Scripting con lógica autónoma",
    slug: "polymarket-veto-loop-bot",
    thumbnail: thumb_polymarket_veto_loop_bot,
    description: "Bot autónomo con pipeline multi-etapa, tareas programadas vía cron y memoria per",
  },
] as const satisfies ProjectPreview[];
