<script setup lang="ts">
import { ref } from 'vue'

const features = ref([
  {
    icon: 'pi pi-bolt',
    title: 'Lightning Fast',
    description: 'Built on Go for high-performance backend operations with minimal resource usage',
  },
  {
    icon: 'pi pi-cog',
    title: 'Docker Ready',
    description: 'Containerized development and testing environments with isolated components',
  },
  {
    icon: 'pi pi-server', // or pi-cog-wheel, pi-database, pi-desktop, pi-star
    title: 'AI Friendly',
    description: 'Architected to work smoothly with AI tools and services out of the box',
  },
  {
    icon: 'pi pi-code',
    title: 'Full Stack',
    description: 'Integrated frontend and backend development with modern tooling',
  },
])
</script>

<template>
  <div class="min-h-screen bg-gradient-to-b from-slate-900 to-slate-800 text-white">
    <!-- Hero Section -->
    <section class="px-4 py-24 md:py-32 max-w-7xl mx-auto text-center relative overflow-hidden">
      <!-- Honeycomb background -->
      <div class="honeycomb-background"></div>

      <!-- Lightning Storm effects -->
      <div class="storm-container">
        <div v-for="n in 5" :key="n" class="lightning" :class="`lightning-${n}`"></div>
        <div class="storm-clouds top-left"><i class="pi pi-cloud text-indigo-800/30"></i></div>
        <div class="storm-clouds top-right"><i class="pi pi-cloud text-indigo-800/30"></i></div>
        <div class="storm-clouds bottom-left"><i class="pi pi-cloud text-indigo-800/30"></i></div>
        <div class="storm-clouds bottom-right"><i class="pi pi-cloud text-indigo-800/30"></i></div>

        <div class="storm-bolts">
          <i
            v-for="n in 8"
            :key="`bolt-${n}`"
            class="pi pi-bolt storm-bolt"
            :class="`bolt-${n}`"
          ></i>
        </div>
      </div>

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

    <!-- Features Section -->
    <section class="px-4 py-16 max-w-7xl mx-auto">
      <h2 class="text-3xl font-bold mb-12 text-center">Why Storm Framework?</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        <div
          v-for="(feature, index) in features"
          :key="index"
          class="bg-slate-800/50 p-6 rounded-xl border border-slate-700"
        >
          <div class="w-12 h-12 bg-indigo-900/50 rounded-lg flex items-center justify-center mb-4">
            <i :class="feature.icon + ' text-xl text-indigo-400'"></i>
          </div>
          <h3 class="text-xl font-semibold mb-2">{{ feature.title }}</h3>
          <p class="text-slate-400">{{ feature.description }}</p>
        </div>
      </div>
    </section>

    <!-- CLI Section -->
    <section class="px-4 py-16 max-w-7xl mx-auto">
      <div class="bg-slate-800 border border-slate-700 rounded-lg p-4 md:p-6">
        <div class="flex items-center mb-4">
          <div class="flex space-x-2">
            <div class="w-3 h-3 rounded-full bg-red-500"></div>
            <div class="w-3 h-3 rounded-full bg-yellow-500"></div>
            <div class="w-3 h-3 rounded-full bg-green-500"></div>
          </div>
        </div>
        <div class="font-mono text-sm md:text-base text-slate-300 overflow-x-auto">
          <p class="pb-2"><span class="text-indigo-400">$</span> storm init-config</p>
          <p class="text-slate-500 pb-2">Configuration file created at ./.storm.yaml</p>
          <p class="pb-2"><span class="text-indigo-400">$</span> storm isolate up --expose --web</p>
          <p class="text-slate-500 pb-2">Starting isolated environment with web service...</p>
          <p class="pb-2"><span class="text-indigo-400">$</span> docker ps</p>
          <p class="text-slate-500">Web frontend running on :5173, backend API on :3000</p>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="px-4 py-8 border-t border-slate-800 mt-16">
      <div class="max-w-7xl mx-auto flex flex-col md:flex-row justify-between items-center">
        <div class="flex items-center mb-4 md:mb-0">
          <i class="pi pi-bolt text-indigo-400 mr-2"></i>
          <span class="font-semibold">Storm Framework</span>
        </div>
        <div class="text-sm text-slate-500">Built with ⚡ for AI-friendly development</div>
      </div>
    </footer>
  </div>
</template>

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

/* Storm Effects */
.storm-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  pointer-events: none;
  z-index: 2;
}

.lightning {
  position: absolute;
  background: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 0) 0%,
    rgba(157, 142, 255, 0.8) 20%,
    rgba(157, 142, 255, 0.3) 70%,
    rgba(157, 142, 255, 0) 100%
  );
  opacity: 0;
  width: 10px;
  height: 150px;
  transform: rotate(20deg) translateY(-100px);
  filter: blur(3px);
  animation: lightning-flash 8s linear infinite;
}

.lightning-1 {
  left: 20%;
  animation-delay: 0s;
}

.lightning-2 {
  left: 40%;
  animation-delay: 2.5s;
}

.lightning-3 {
  left: 60%;
  animation-delay: 4.7s;
}

.lightning-4 {
  right: 15%;
  animation-delay: 6.2s;
}

.lightning-5 {
  left: 30%;
  animation-delay: 7.5s;
}

@keyframes lightning-flash {
  0% {
    opacity: 0;
    transform: rotate(20deg) translateY(-100px);
  }
  2% {
    opacity: 0.8;
    transform: rotate(20deg) translateY(20px);
  }
  4% {
    opacity: 0.2;
  }
  6% {
    opacity: 0.9;
  }
  8% {
    opacity: 0;
    transform: rotate(20deg) translateY(100px);
  }
  100% {
    opacity: 0;
  }
}

.storm-clouds {
  position: absolute;
  font-size: 4rem;
  opacity: 0.6;
  animation: cloud-drift 20s linear infinite;
}

.top-left {
  top: 10%;
  left: 5%;
  animation-delay: 0s;
}

.top-right {
  top: 15%;
  right: 5%;
  animation-delay: 5s;
}

.bottom-left {
  bottom: 10%;
  left: 10%;
  animation-delay: 10s;
}

.bottom-right {
  bottom: 15%;
  right: 10%;
  animation-delay: 15s;
}

@keyframes cloud-drift {
  0% {
    transform: translateX(0) scale(1);
  }
  50% {
    transform: translateX(20px) scale(1.05);
  }
  100% {
    transform: translateX(0) scale(1);
  }
}

.storm-bolts {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
}

.storm-bolt {
  position: absolute;
  color: rgba(147, 197, 253, 0.8);
  font-size: 1.5rem;
  animation: bolt-flash 5s ease-in-out infinite;
}

.bolt-1 {
  top: 25%;
  left: 15%;
  animation-delay: 0.5s;
}
.bolt-2 {
  top: 40%;
  left: 80%;
  animation-delay: 1.5s;
}
.bolt-3 {
  top: 70%;
  left: 25%;
  animation-delay: 2.7s;
}
.bolt-4 {
  top: 20%;
  left: 60%;
  animation-delay: 3.2s;
}
.bolt-5 {
  top: 60%;
  left: 70%;
  animation-delay: 4.1s;
}
.bolt-6 {
  top: 35%;
  left: 30%;
  animation-delay: 1.8s;
}
.bolt-7 {
  top: 55%;
  left: 10%;
  animation-delay: 3.5s;
}
.bolt-8 {
  top: 15%;
  left: 85%;
  animation-delay: 0.2s;
}

@keyframes bolt-flash {
  0% {
    opacity: 0;
    transform: scale(0.8) rotate(0deg);
  }
  3% {
    opacity: 0.9;
    transform: scale(1.2) rotate(5deg);
  }
  6% {
    opacity: 0.3;
    transform: scale(0.9) rotate(-5deg);
  }
  9% {
    opacity: 0.8;
    transform: scale(1.1) rotate(0deg);
  }
  12% {
    opacity: 0;
    transform: scale(1) rotate(0deg);
  }
  100% {
    opacity: 0;
  }
}
</style>
