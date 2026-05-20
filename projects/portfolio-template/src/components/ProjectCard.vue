<script setup lang="ts">
import type { Locale, ProjectContent } from '../types';
defineProps<{ project: ProjectContent; locale: Locale; index: number }>();
</script>

<template>
  <article class="project-card">
    <div class="project-thumb">
      <div class="thumb-gradient" :data-idx="(index % 5).toString()">
        <span class="thumb-title">{{ project.title }}</span>
      </div>
    </div>
    <div class="project-body">
      <h3 class="project-title">{{ project.title }}</h3>
      <p class="project-summary">
        {{ locale === 'es' ? project.summary_es : project.summary_en }}
      </p>
      <ul class="project-tech">
        <li v-for="t in project.tech" :key="t" :class="{ highlighted: project.highlighted_tech?.includes(t) }">{{ t }}</li>
      </ul>
      <div class="project-links">
        <a v-if="project.live" :href="project.live" target="_blank" rel="noopener">
          {{ locale === 'es' ? 'Ver live' : 'View live' }} →
        </a>
        <a v-if="project.source" :href="project.source" target="_blank" rel="noopener">
          {{ locale === 'es' ? 'Código' : 'Source' }} →
        </a>
      </div>
    </div>
  </article>
</template>
