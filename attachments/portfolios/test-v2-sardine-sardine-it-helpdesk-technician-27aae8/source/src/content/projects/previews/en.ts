import thumb_jarvis_multi_agent from "../../../assets/thumbnails/jarvis-multi-agent.webp";
import thumb_trabajo_auto_apply from "../../../assets/thumbnails/trabajo-auto-apply.webp";
import thumb_jarvis_dashboard from "../../../assets/thumbnails/jarvis-dashboard.webp";
import thumb_polymarket_veto_loop_bot from "../../../assets/thumbnails/polymarket-veto-loop-bot.webp";

import type { ProjectPreview } from "../../types";

export default [
  {
    title: "JARVIS · Autonomous Production Infrastructure",
    slug: "jarvis-multi-agent",
    thumbnail: thumb_jarvis_multi_agent,
    description: "System deployed and operated 24/7 on a Hetzner VPS under Linux, managed and moni",
  },
  {
    title: "Trabajo · Process Automation Pipeline",
    slug: "trabajo-auto-apply",
    thumbnail: thumb_trabajo_auto_apply,
    description: "End-to-end pipeline with approval gates automating job search, CV adaptation, an",
  },
  {
    title: "JARVIS Dashboard · Centralized Systems Panel",
    slug: "jarvis-dashboard",
    thumbnail: thumb_jarvis_dashboard,
    description: "Centralized web panel for agent management, service monitoring, and real-time sy",
  },
  {
    title: "Polymarket Bot · Autonomous Scripting Logic",
    slug: "polymarket-veto-loop-bot",
    thumbnail: thumb_polymarket_veto_loop_bot,
    description: "Autonomous bot with a multi-stage pipeline, cron-scheduled tasks, and persistent",
  },
] as const satisfies ProjectPreview[];
