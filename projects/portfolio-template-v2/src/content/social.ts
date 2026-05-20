export const social = [
  { url: "mailto:davidgnjunior@gmail.com", name: "mail" },
  { url: "https://www.linkedin.com/in/david-gonz%C3%A1lez-a8267032b/", name: "linkedin" },
] as const satisfies { url: string; name: "mail" | "github" | "instagram" | "linkedin" | "x" }[];
