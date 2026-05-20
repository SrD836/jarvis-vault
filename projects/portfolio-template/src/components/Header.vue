<script setup lang="ts">
import type { Locale } from '../types';
defineProps<{ locale: Locale }>();
defineEmits<{ (e: 'set-locale', l: Locale): void }>();
</script>

<template>
  <header class="site-header">
    <div class="container header-inner">
      <a href="#" class="logo">DG</a>
      <nav class="lang-switch">
        <button :class="{ on: locale === 'es' }" @click="$emit('set-locale','es')">ES</button>
        <span>·</span>
        <button :class="{ on: locale === 'en' }" @click="$emit('set-locale','en')">EN</button>
      </nav>
    </div>
  </header>
</template>
