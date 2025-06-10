<script setup lang="ts">
defineProps<{
  icon: string
  title: string
  description: string
}>()
</script>

<template>
  <div class="bg-slate-800/50 p-6 rounded-xl border border-slate-700">
    <div class="w-12 h-12 bg-indigo-900/50 rounded-lg flex items-center justify-center mb-4">
      <i :class="icon + ' text-xl text-indigo-400'"></i>
    </div>
    <h3 class="text-xl font-semibold mb-2">{{ title }}</h3>
    <p class="text-slate-400">{{ description }}</p>
  </div>
</template>
