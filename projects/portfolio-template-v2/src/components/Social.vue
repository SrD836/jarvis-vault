<script setup lang="ts">
import Github from "./icons/Github.vue";
import Linkedin from "./icons/Linkedin.vue";
import Instagram from "./icons/Instagram.vue";
import Mail from "./icons/Mail.vue";
import X from "./icons/X.vue";
import Link from "./Link.vue";
import { t } from "../i18n/utils/translate";
import ButtonRound from "./ButtonRound.vue";
import { copyEmail } from "../composables/useCopyEmail";

import { social } from "../content/social";

const props = defineProps<{
  variant?: "theme" | "background";
}>();

const icons = {
  mail: Mail,
  github: Github,
  linkedin: Linkedin,
  x: X,
  instagram: Instagram,
} as const;

const getAriaLabel = (name: string) => `${t("go-to")} ${name.charAt(0).toUpperCase() + name.slice(1)}`;
</script>

<template>
  <div class="social">
    <template v-for="item in social" :key="item.name">
      <button
        v-if="item.name === 'mail'"
        type="button"
        :aria-label="t('get-in-touch')"
        class="social-link social-link--button"
        data-cursor="circle-white"
        data-sound="click"
        @click="copyEmail"
      >
        <ButtonRound
          renderAs="div"
          :variant="props.variant ?? 'theme'"
          class="children-unclickable"
          data-hoversound="hover"
        >
          <component :is="icons[item.name]" :aria-label="t('get-in-touch')" />
        </ButtonRound>
      </button>
      <Link
        v-else
        external
        :href="item.url"
        :aria-label="getAriaLabel(item.name)"
        class="social-link"
        data-cursor="circle-white"
      >
        <ButtonRound
          renderAs="div"
          :variant="props.variant ?? 'theme'"
          class="children-unclickable"
          data-hoversound="hover"
        >
          <component :is="icons[item.name]" :aria-label="getAriaLabel(item.name)" external />
        </ButtonRound>
      </Link>
    </template>
  </div>
</template>

<style scoped lang="scss">
.social {
  display: flex;
  gap: var(--space-md);
}

.social-link--button {
  background: none;
  border: none;
  padding: 0;
  cursor: pointer;
  font: inherit;
  color: inherit;
}
</style>
