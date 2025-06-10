<script setup lang="ts">
import { ref } from 'vue'
import HeroSection from '@/components/HeroSection.vue'
import FeatureCard from '@/components/FeatureCard.vue'
import CLITerminal from '@/components/CLITerminal.vue'
import AppFooter from '@/components/AppFooter.vue'

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
    icon: 'pi pi-server',
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
    <HeroSection />

    <!-- Features Section -->
    <section class="px-4 py-16 max-w-7xl mx-auto">
      <h2 class="text-3xl font-bold mb-12 text-center">Why Storm Framework?</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        <FeatureCard
          v-for="(feature, index) in features"
          :key="index"
          :icon="feature.icon"
          :title="feature.title"
          :description="feature.description"
        />
      </div>
    </section>

    <!-- CLI Section -->
    <section class="px-4 py-16 max-w-7xl mx-auto">
      <CLITerminal />
    </section>

    <!-- Footer -->
    <AppFooter />
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
