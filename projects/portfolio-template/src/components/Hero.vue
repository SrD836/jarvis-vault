<script setup lang="ts">
import type { Locale, Hero } from '../types';
defineProps<{ hero: Hero; locale: Locale }>();
</script>

<template>
  <section class="hero">
    <div class="hero-bg" aria-hidden="true">
      <div class="orb orb-1"></div>
      <div class="orb orb-2"></div>
      <div class="grid-overlay"></div>
    </div>
    <div class="container hero-content">
      <p class="tagline">{{ locale === 'es' ? hero.tagline_es : hero.tagline_en }}</p>
      <h1 class="name">{{ hero.name }}</h1>
      <p class="intro">{{ locale === 'es' ? hero.intro_es : hero.intro_en }}</p>
    </div>
    <div class="scroll-cue">{{ locale === 'es' ? 'Desliza' : 'Scroll' }} ↓</div>
  </section>
</template>
