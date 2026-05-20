<script setup lang="ts">
import { toastVisible, toastMessage } from "../composables/useCopyEmail";
</script>

<template>
  <Transition name="email-toast">
    <div v-if="toastVisible" class="email-toast" role="status" aria-live="polite">
      {{ toastMessage }}
    </div>
  </Transition>
</template>

<style scoped lang="scss">
.email-toast {
  position: fixed;
  bottom: 32px;
  left: 50%;
  transform: translateX(-50%);
  background: var(--color-background-100, #111);
  color: var(--color-text-100, #fff);
  padding: 12px 24px;
  border-radius: 999px;
  font-size: var(--font-size-sm, 14px);
  font-weight: 500;
  letter-spacing: 0.02em;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.35);
  z-index: 9999;
  pointer-events: none;
  white-space: nowrap;
}

.email-toast-enter-active,
.email-toast-leave-active {
  transition:
    opacity 220ms ease,
    transform 220ms ease;
}

.email-toast-enter-from,
.email-toast-leave-to {
  opacity: 0;
  transform: translate(-50%, 12px);
}
</style>
