<template>
  <div class="storm-container">
    <div v-for="n in 5" :key="n" class="lightning" :class="`lightning-${n}`"></div>
    <div class="storm-clouds top-left"><i class="pi pi-cloud text-indigo-800/30"></i></div>
    <div class="storm-clouds top-right"><i class="pi pi-cloud text-indigo-800/30"></i></div>
    <div class="storm-clouds bottom-left"><i class="pi pi-cloud text-indigo-800/30"></i></div>
    <div class="storm-clouds bottom-right"><i class="pi pi-cloud text-indigo-800/30"></i></div>

    <div class="storm-bolts">
      <i v-for="n in 8" :key="`bolt-${n}`" class="pi pi-bolt storm-bolt" :class="`bolt-${n}`"></i>
    </div>
  </div>
</template>

<style scoped>
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
