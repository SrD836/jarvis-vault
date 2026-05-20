<script setup lang="ts">
import type { Locale } from '../types';
import portfolio from '../data/portfolio';
defineProps<{ locale: Locale }>();
const gen = portfolio.generated_for;
</script>

<template>
  <footer class="site-footer">
    <div class="container">
      <p>© 2026 David González Nuez</p>
      <p v-if="gen && gen.offer_title" class="muted small">
        {{ locale === 'es' ? 'Generado para' : 'Generated for' }}
        <strong>{{ gen.offer_title }}</strong>
        <template v-if="gen.company"> @ {{ gen.company }}</template>
      </p>
    </div>
  </footer>
</template>
