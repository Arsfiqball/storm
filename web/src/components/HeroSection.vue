<template>
  <section class="px-4 py-24 md:py-32 max-w-7xl mx-auto text-center relative overflow-hidden">
    <!-- Honeycomb background -->
    <div class="honeycomb-background"></div>

    <!-- Storm Effects -->
    <StormEffects />

    <div class="mb-8 inline-block p-2 bg-indigo-900/30 rounded-full relative z-10">
      <i class="pi pi-bolt text-2xl text-indigo-400"></i>
    </div>
    <h1
      class="text-4xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-indigo-400 to-indigo-200 bg-clip-text text-transparent relative z-10"
    >
      Storm Framework
    </h1>
    <p class="text-xl md:text-2xl mb-10 max-w-3xl mx-auto text-slate-300 relative z-10">
      The AI-friendly fullstack framework for building modern web applications
    </p>
    <div class="flex flex-wrap justify-center gap-4 relative z-10">
      <button
        class="p-button p-button-lg bg-indigo-600 hover:bg-indigo-700 border-indigo-600 text-white px-8 py-3 rounded-lg font-medium flex items-center gap-2"
      >
        <i class="pi pi-play"></i>
        Get Started
      </button>
      <a
        href="https://github.com/Arsfiqball/storm"
        target="_blank"
        class="p-button p-button-lg p-button-outlined text-indigo-300 border-indigo-600/50 hover:border-indigo-400 px-8 py-3 rounded-lg font-medium flex items-center gap-2"
      >
        <i class="pi pi-github"></i>
        GitHub
      </a>
    </div>
  </section>
</template>

<script setup lang="ts">
import StormEffects from './StormEffects.vue'
</script>

<style scoped>
.p-button {
  transition: all 0.2s;
}

.honeycomb-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: transparent;
  background-image:
    radial-gradient(circle at center, rgba(99, 91, 255, 0.1) 0%, transparent 70%),
    url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='56' height='100' viewBox='0 0 56 100'%3E%3Cpath d='M28 66L0 50L0 16L28 0L56 16L56 50L28 66L28 100' fill='none' stroke='rgba(120, 110, 255, 0.18)'/%3E%3C/svg%3E");
  opacity: 0.8;
  mask-image: radial-gradient(ellipse at center, rgba(0, 0, 0, 1) 30%, transparent 80%);
  -webkit-mask-image: radial-gradient(ellipse at center, rgba(0, 0, 0, 1) 30%, transparent 80%);
  z-index: 1;
  pointer-events: none;
}
</style>
