export type Locale = 'es' | 'en';

export interface ProjectContent {
  slug: string;
  title: string;
  tech: string[];
  summary_es: string;
  summary_en: string;
  live?: string | null;
  source?: string | null;
  thumbnail?: string;
  media?: string[];
  highlighted_tech?: string[];
}

export interface Hero {
  name: string;
  tagline_es: string;
  tagline_en: string;
  intro_es: string;
  intro_en: string;
}

export interface PortfolioData {
  hero: Hero;
  projects: ProjectContent[];
  generated_for?: {
    offer_title?: string;
    company?: string;
    generated_at?: string;
  };
}
