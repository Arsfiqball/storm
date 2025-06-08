<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Button from 'primevue/button'
import Card from 'primevue/card'
import Divider from 'primevue/divider'
import Timeline from 'primevue/timeline'

const timelineItems = ref([
  {
    title: 'Framework Inception',
    description: 'The Storm project was born, with a vision of simplified development.',
  },
  {
    title: 'First Release',
    description: 'Version 1.0 released with core components and documentation.',
  },
  {
    title: 'Community Growth',
    description: 'Developer community expanded with contributions and feedback.',
  },
  {
    title: 'Enterprise Adoption',
    description: 'Major organizations began deploying Storm-based solutions.',
  },
  {
    title: 'Future Roadmap',
    description: 'Continued innovation with AI integration and expanded tooling.',
  },
])

// Animate lightning flashes randomly
const triggerLightning = () => {
  const lightning = document.querySelector('.lightning-effect')
  if (lightning) {
    lightning.classList.add('active')
    setTimeout(() => {
      lightning.classList.remove('active')
    }, 200)
  }
}

onMounted(() => {
  setInterval(
    () => {
      // Random lightning flashes (25% chance every 5-10 seconds)
      if (Math.random() < 0.25) {
        triggerLightning()
      }
    },
    Math.random() * 5000 + 5000,
  )
})
</script>

<template>
  <div class="surface-ground px-4 py-8 md:px-6 lg:px-8">
    <div class="hexagon-background px-4 py-8 md:px-6 lg:px-8 mb-6">
      <!-- Lightning flash overlay -->
      <div class="lightning-effect"></div>

      <!-- Rain drops effect -->
      <div class="rain-container">
        <div
          v-for="n in 20"
          :key="n"
          class="rain-drop"
          :style="`--delay: ${(n * 0.1).toFixed(2)}s; --position: ${Math.random() * 100}%`"
        ></div>
      </div>

      <div class="content-wrapper">
        <div class="text-center mb-5">
          <div class="text-900 font-bold text-5xl mb-3 storm-text">Welcome to Storm Framework</div>
          <div class="text-600 text-xl mb-6">
            <i class="pi pi-bolt thunderbolt-icon"></i> An AI-friendly fullstack framework built for
            modern development <i class="pi pi-bolt thunderbolt-icon"></i>
          </div>
          <p class="text-700 line-height-3 mb-4 max-w-6 mx-auto">
            Storm combines a powerful Go backend with a pre-configured Vue SPA frontend using
            PrimeVue components, enabling developers to build sophisticated applications with
            built-in AI capabilities.
          </p>
          <div class="flex justify-content-center">
            <Button
              label="Get Started"
              icon="pi pi-rocket"
              class="mr-3 p-button-raised pulse-btn"
            />
            <Button label="Documentation" icon="pi pi-book" class="p-button-outlined" />
          </div>
        </div>

        <!-- Cloud with lightning icon -->
        <div class="storm-cloud">
          <i class="pi pi-cloud"></i>
          <span class="cloud-lightning"></span>
        </div>

        <div class="grid">
          <div class="col-12 md:col-4">
            <Card class="shadow-2 p-3 h-full feature-card">
              <template #title>Fast Development</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-bolt text-xl text-primary mr-2 lightning-icon"></i>
                  <span class="text-xl font-medium">Speed Optimized</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Build applications rapidly with our streamlined workflows and ready-to-use
                  components.
                </p>
              </template>
            </Card>
          </div>

          <div class="col-12 md:col-4">
            <Card class="shadow-2 p-3 h-full feature-card">
              <template #title>Modern Architecture</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-cloud text-xl text-primary mr-2 cloud-icon"></i>
                  <span class="text-xl font-medium">Future-Proof</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Built on the latest technologies and standards to ensure your projects remain
                  maintainable.
                </p>
              </template>
            </Card>
          </div>

          <div class="col-12 md:col-4">
            <Card class="shadow-2 p-3 h-full feature-card">
              <template #title>Responsive Design</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-mobile text-xl text-primary mr-2"></i>
                  <span class="text-xl font-medium">Any Device</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Create applications that look and function beautifully across desktop, tablet, and
                  mobile.
                </p>
              </template>
            </Card>
          </div>

          <!-- Add a fourth card about power -->
          <div class="col-12 md:col-6 lg:col-3 mt-4">
            <Card class="shadow-2 p-3 h-full feature-card power-card">
              <template #title>Powerful Backend</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-server text-xl text-primary mr-2 thunder-icon"></i>
                  <span class="text-xl font-medium">Unlimited Power</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Harness the lightning-fast Go backend to handle any workload with exceptional
                  performance.
                </p>
              </template>
            </Card>
          </div>

          <!-- Add a fifth card about AI -->
          <div class="col-12 md:col-6 lg:col-3 mt-4">
            <Card class="shadow-2 p-3 h-full feature-card ai-card">
              <template #title>AI Integration</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-bolt-circle text-xl text-primary mr-2 brain-icon"></i>
                  <span class="text-xl font-medium">Smart Tools</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Seamlessly integrate AI capabilities that supercharge your application's
                  intelligence.
                </p>
              </template>
            </Card>
          </div>

          <!-- Add a sixth card about security -->
          <div class="col-12 md:col-6 lg:col-3 mt-4">
            <Card class="shadow-2 p-3 h-full feature-card security-card">
              <template #title>Enhanced Security</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-shield text-xl text-primary mr-2 shield-icon"></i>
                  <span class="text-xl font-medium">Built-in Protection</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Robust security features to protect your applications from common vulnerabilities
                  and threats.
                </p>
              </template>
            </Card>
          </div>

          <!-- Add a seventh card about extensibility -->
          <div class="col-12 md:col-6 lg:col-3 mt-4">
            <Card class="shadow-2 p-3 h-full feature-card extension-card">
              <template #title>Extensible Platform</template>
              <template #content>
                <div class="flex align-items-center mb-3">
                  <i class="pi pi-cog text-xl text-primary mr-2 gear-icon"></i>
                  <span class="text-xl font-medium">Customize & Extend</span>
                </div>
                <p class="text-700 line-height-3 m-0">
                  Easily extend functionality with plugins and customize every aspect to meet your
                  project needs.
                </p>
              </template>
            </Card>
          </div>
        </div>
      </div>
    </div>

    <Divider class="my-6" />

    <div class="content-wrapper">
      <div class="text-center">
        <div class="text-900 font-bold text-4xl mb-4">Framework Timeline</div>
        <p class="text-600 mb-5">Journey through Storm Framework's development and future</p>
      </div>

      <div class="card storm-timeline-card">
        <Timeline :value="timelineItems" align="alternate" class="my-5">
          <template #content="slotProps">
            <div class="timeline-event">
              <div class="text-900 font-bold mb-1">{{ slotProps.item.title }}</div>
              <p class="m-0 text-600">{{ slotProps.item.description }}</p>
            </div>
          </template>
        </Timeline>
      </div>

      <div class="mt-8 text-center">
        <Button label="Join Our Community" icon="pi pi-users" class="p-button-lg" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.timeline-event {
  padding: 1rem;
  border-radius: 12px;
  background-color: var(--surface-card);
  box-shadow:
    0 2px 1px -1px rgba(0, 0, 0, 0.2),
    0 1px 1px 0 rgba(0, 0, 0, 0.14),
    0 1px 3px 0 rgba(0, 0, 0, 0.12);
  transition: transform 0.2s;
}

.timeline-event:hover {
  transform: translateY(-3px);
}

.storm-event {
  border-left: 3px solid var(--primary-color);
}

.timeline-icon {
  position: absolute;
  top: 10px;
  right: 10px;
  color: var(--primary-color);
  font-size: 1.2rem;
  opacity: 0.7;
}

.hexagon-background {
  position: relative;
  overflow: hidden;
  background-color: var(--surface-ground);
  border-radius: 1rem;
}

/* Simplified base layer for honeycomb pattern with reduced opacity */
.hexagon-background::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='56' height='100' viewBox='0 0 56 100'%3E%3Cpath d='M28 66L0 50L0 16L28 0L56 16L56 50L28 66L28 100' fill='none' stroke='rgba(89, 129, 255, 0.05)' stroke-width='1'/%3E%3Cpath d='M28 0L56 16L56 50L28 66L0 50L0 16L28 0Z' fill='rgba(89, 129, 255, 0.01)'/%3E%3C/svg%3E");
  background-size: 56px 100px;
  z-index: 0;
  /* Removed mask-image gradient that caused the cloudy effect */
}

/* Removed the secondary offset layer that created additional cloudy effect */

.content-wrapper {
  position: relative;
  z-index: 1;
}

:root {
  --primary-color-rgb: 89, 129, 255; /* Default primary color */
}

/* Lightning flash effect */
.lightning-effect {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0);
  z-index: 2;
  pointer-events: none;
  transition: background-color 0.1s;
}

.lightning-effect.active {
  background-color: rgba(255, 255, 255, 0.4);
  box-shadow: 0 0 80px 30px rgba(89, 129, 255, 0.8);
}

/* Rain drops effect */
.rain-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  pointer-events: none;
  z-index: 1;
  opacity: 0.4;
}

.rain-drop {
  position: absolute;
  width: 1px;
  height: 20px;
  background: linear-gradient(to bottom, rgba(89, 129, 255, 0), rgba(89, 129, 255, 0.7));
  top: -20px;
  left: var(--position);
  animation: rain 1s linear infinite;
  animation-delay: var(--delay);
}

@keyframes rain {
  0% {
    transform: translateY(-20px);
  }
  100% {
    transform: translateY(calc(100vh + 20px));
  }
}

/* Storm text with gradient effect */
.storm-text {
  background: linear-gradient(45deg, #3a6eff, #5981ff, #82a4ff);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  text-shadow: 0px 3px 7px rgba(58, 110, 255, 0.2);
  position: relative;
}

/* Thunderbolt icon animation */
.thunderbolt-icon {
  color: #5981ff;
  animation: pulse 2s infinite;
  display: inline-block;
}

@keyframes pulse {
  0%,
  100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.6;
    transform: scale(0.92);
  }
}

/* Storm cloud with lightning */
.storm-cloud {
  position: absolute;
  top: 10px;
  right: 20px;
  font-size: 2rem;
  color: rgba(89, 129, 255, 0.4);
  z-index: 2;
}

.cloud-lightning {
  position: absolute;
  top: 60%;
  left: 50%;
  width: 2px;
  height: 15px;
  background-color: rgba(255, 255, 255, 0.9);
  animation: strike 5s infinite;
  transform-origin: top;
  z-index: 3;
}

@keyframes strike {
  0%,
  95%,
  100% {
    opacity: 0;
  }
  96%,
  99% {
    opacity: 1;
  }
}

/* Feature cards enhancement */
.feature-card {
  transition:
    transform 0.3s,
    box-shadow 0.3s;
  position: relative;
  overflow: hidden;
}

.feature-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  background: linear-gradient(90deg, #3a6eff, #82a4ff);
  z-index: 1;
  opacity: 0;
  transition: opacity 0.3s;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(89, 129, 255, 0.15);
}

.feature-card:hover::after {
  opacity: 1;
}

/* Lightning icon animation */
.lightning-icon {
  animation: flash 3s infinite;
}

@keyframes flash {
  0%,
  80%,
  100% {
    opacity: 1;
  }
  85%,
  95% {
    opacity: 0.4;
    transform: scale(1.2);
    color: #3a6eff;
  }
}

/* Cloud icon animation */
.cloud-icon {
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-5px);
  }
}

/* Pulsing button effect */
.pulse-btn {
  animation: pulse-border 2s infinite;
  position: relative;
}

@keyframes pulse-border {
  0% {
    box-shadow: 0 0 0 0 rgba(89, 129, 255, 0.4);
  }
  70% {
    box-shadow: 0 0 0 6px rgba(89, 129, 255, 0);
  }
  100% {
    box-shadow: 0 0 0 0 rgba(89, 129, 255, 0);
  }
}

/* Thunder button style */
.thunder-btn {
  background: linear-gradient(45deg, #3a6eff, #5981ff);
  border: none;
  position: relative;
  overflow: hidden;
  transition: all 0.3s;
}

.thunder-btn::after {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: all 0.5s;
}

.thunder-btn:hover::after {
  left: 100%;
}

/* Power card special styling */
.power-card {
  background: linear-gradient(135deg, rgba(255, 255, 255, 1), rgba(240, 246, 255, 1));
}

.thunder-icon {
  animation: rumble 2s infinite;
}

@keyframes rumble {
  0%,
  100% {
    transform: translateX(0);
  }
  10%,
  30%,
  50%,
  70%,
  90% {
    transform: translateX(-2px);
  }
  20%,
  40%,
  60%,
  80% {
    transform: translateX(2px);
  }
}

/* AI card special styling */
.ai-card {
  background: linear-gradient(135deg, rgba(255, 255, 255, 1), rgba(245, 248, 255, 1));
}

.brain-icon {
  animation: glow 3s infinite;
}

@keyframes glow {
  0%,
  100% {
    filter: drop-shadow(0 0 0 rgba(89, 129, 255, 0.5));
  }
  50% {
    filter: drop-shadow(0 0 5px rgba(89, 129, 255, 0.8));
  }
}

/* Storm timeline card */
.storm-timeline-card {
  background-color: var(--surface-card);
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  position: relative;
}
</style>
