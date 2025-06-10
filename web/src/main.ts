import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import Lara from '@primeuix/themes/lara'

import App from './App.vue'
import router from './router'
import { definePreset } from '@primeuix/themes'

const app = createApp(App)

app.use(createPinia())

const stormPreset = definePreset(Lara, {
  semantic: {
    primary: {
      50: '{slate.50}',
      100: '{slate.100}',
      200: '{slate.200}',
      300: '{slate.300}',
      400: '{slate.400}',
      500: '{slate.500}',
      600: '{slate.600}',
      700: '{slate.700}',
      800: '{slate.800}',
      900: '{slate.900}',
      950: '{slate.950}',
    },
  },
})

app.use(PrimeVue, {
  theme: {
    preset: stormPreset,
  },
})

app.use(router)

app.mount('#app')
