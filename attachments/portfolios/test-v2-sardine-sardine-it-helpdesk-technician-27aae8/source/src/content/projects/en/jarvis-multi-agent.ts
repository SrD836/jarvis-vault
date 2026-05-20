import type { ProjectContent } from "../../types";

export default {
  "title": "JARVIS · Autonomous Production Infrastructure",
  "theme": "dark",
  "tags": [
    "typescript",
    "docker",
    "node"
  ],
  "description": "System deployed and operated 24/7 on a Hetzner VPS under Linux, managed and monitored entirely remotely.<br/>Demonstrates real systems administration: production service configuration, incident resolution without manual intervention, and background process management — skills directly applicable to L1/L2 support and endpoint management.",
  "components": [],
  "live": "https://jarvss.duckdns.org"
} as const satisfies ProjectContent;
