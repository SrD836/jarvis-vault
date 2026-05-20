import type { Locale } from "../../i18n/types";

export const projectIds = ["jarvis-multi-agent", "trabajo-auto-apply", "jarvis-dashboard", "polymarket-veto-loop-bot"] as const;

function simplifyModules(glob: Record<string, any>) {
  const result: Record<string, any> = {};
  for (const [path, mod] of Object.entries(glob)) {
    const match = path.match(/\/([a-z0-9_-]+)\.ts$/i);
    if (match) result[match[1] as string] = mod;
  }
  return result;
}

export const projectModules = {
  es: simplifyModules(import.meta.glob("./es/*.ts", { eager: true })),
  en: simplifyModules(import.meta.glob("./en/*.ts", { eager: true })),
} as const satisfies Record<Locale, Record<string, any>>;
