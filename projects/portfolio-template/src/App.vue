<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue';
import Lenis from 'lenis';
import gsap from 'gsap';
import { ScrollTrigger } from 'gsap/ScrollTrigger';
import portfolio from './data/portfolio';
import Header from './components/Header.vue';
import Hero from './components/Hero.vue';
import ProjectCard from './components/ProjectCard.vue';
import Footer from './components/Footer.vue';
import type { Locale } from './types';

gsap.registerPlugin(ScrollTrigger);

const locale = ref<Locale>('es');
const setLocale = (l: Locale) => { locale.value = l; };

const projects = computed(() => portfolio.projects);
const hero = computed(() => portfolio.hero);

let lenis: Lenis | null = null;
let rafId = 0;

onMounted(() => {
  document.body.classList.remove('is-loading');
  lenis = new Lenis({ duration: 1.1, smoothWheel: true });
  const raf = (time: number) => { lenis?.raf(time); rafId = requestAnimationFrame(raf); };
  rafId = requestAnimationFrame(raf);
  lenis.on('scroll', () => ScrollTrigger.update());

  gsap.from('.hero-content > *', {
    opacity: 0, y: 40, stagger: 0.12, duration: 0.9, ease: 'power3.out', delay: 0.2,
  });
  gsap.utils.toArray<HTMLElement>('.project-card').forEach((el) => {
    gsap.from(el, {
      opacity: 0, y: 60, duration: 0.7, ease: 'power2.out',
      scrollTrigger: { trigger: el, start: 'top 85%' },
    });
  });
});

onUnmounted(() => {
  cancelAnimationFrame(rafId);
  lenis?.destroy();
});
</script>

<template>
  <div class="app">
    <Header :locale="locale" @set-locale="setLocale" />
    <main>
      <Hero :hero="hero" :locale="locale" />
      <section class="projects-section">
        <div class="container">
          <h2 class="section-title">{{ locale === 'es' ? 'Proyectos' : 'Projects' }}</h2>
          <div class="project-grid">
            <ProjectCard
              v-for="(p, i) in projects"
              :key="p.slug"
              :project="p"
              :locale="locale"
              :index="i"
            />
          </div>
        </div>
      </section>
    </main>
    <Footer :locale="locale" />
  </div>
</template>
