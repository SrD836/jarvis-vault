import { ref } from "vue";

const EMAIL = "davidgnjunior@gmail.com";

export const toastVisible = ref(false);
export const toastMessage = ref("");

let hideTimer: ReturnType<typeof setTimeout> | null = null;

export async function copyEmail(): Promise<void> {
  let ok = false;
  try {
    if (navigator?.clipboard?.writeText) {
      await navigator.clipboard.writeText(EMAIL);
      ok = true;
    }
  } catch {
    ok = false;
  }
  if (!ok) {
    try {
      const ta = document.createElement("textarea");
      ta.value = EMAIL;
      ta.style.position = "fixed";
      ta.style.opacity = "0";
      document.body.appendChild(ta);
      ta.select();
      document.execCommand("copy");
      document.body.removeChild(ta);
      ok = true;
    } catch {}
  }
  toastMessage.value = ok ? `Email copied: ${EMAIL}` : `Email: ${EMAIL}`;
  toastVisible.value = true;
  if (hideTimer) clearTimeout(hideTimer);
  hideTimer = setTimeout(() => {
    toastVisible.value = false;
  }, 2800);
}

export function getEmail(): string {
  return EMAIL;
}
