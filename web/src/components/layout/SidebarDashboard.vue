<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import Sidebar from 'primevue/sidebar'
import Menu from 'primevue/menu'
import Button from 'primevue/button'
import Avatar from 'primevue/avatar'

const router = useRouter()
const visible = ref(true)
const miniSidebar = ref(false)
const isMobile = ref(false)
const menuRef = ref()

// Menu items
const menuItems = ref([
  {
    label: 'Dashboard',
    icon: 'pi pi-home',
    command: () => router.push('/dashboard'),
  },
  {
    label: 'Analytics',
    icon: 'pi pi-chart-line',
    items: [
      {
        label: 'Overview',
        icon: 'pi pi-chart-pie',
        command: () => router.push('/analytics/overview'),
      },
      {
        label: 'Reports',
        icon: 'pi pi-file-pdf',
        command: () => router.push('/analytics/reports'),
      },
      {
        label: 'Advanced',
        icon: 'pi pi-sliders-v',
        items: [
          {
            label: 'Forecasting',
            icon: 'pi pi-chart-line',
            command: () => router.push('/analytics/advanced/forecasting'),
          },
          {
            label: 'Trends',
            icon: 'pi pi-arrow-up',
            command: () => router.push('/analytics/advanced/trends'),
          },
        ],
      },
    ],
  },
  {
    label: 'Services',
    icon: 'pi pi-server',
    items: [
      {
        label: 'Active Services',
        icon: 'pi pi-check-circle',
        command: () => router.push('/services/active'),
      },
      {
        label: 'Service Manager',
        icon: 'pi pi-cog',
        command: () => router.push('/services/manager'),
      },
      {
        label: 'API Gateway',
        icon: 'pi pi-arrow-right-arrow-left',
        command: () => router.push('/services/api-gateway'),
      },
    ],
  },
  {
    label: 'Resources',
    icon: 'pi pi-database',
    items: [
      {
        label: 'Storage',
        icon: 'pi pi-save',
        command: () => router.push('/resources/storage'),
      },
      {
        label: 'Networking',
        icon: 'pi pi-wifi',
        items: [
          {
            label: 'Connections',
            icon: 'pi pi-link',
            command: () => router.push('/resources/networking/connections'),
          },
          {
            label: 'Security Groups',
            icon: 'pi pi-shield',
            command: () => router.push('/resources/networking/security-groups'),
          },
        ],
      },
      {
        label: 'Compute',
        icon: 'pi pi-server',
        command: () => router.push('/resources/compute'),
      },
    ],
  },
  {
    label: 'Settings',
    icon: 'pi pi-cog',
    items: [
      {
        label: 'User Settings',
        icon: 'pi pi-user-edit',
        command: () => router.push('/settings/user'),
      },
      {
        label: 'System Settings',
        icon: 'pi pi-sliders-h',
        command: () => router.push('/settings/system'),
      },
      {
        label: 'Notifications',
        icon: 'pi pi-bell',
        command: () => router.push('/settings/notifications'),
      },
      {
        label: 'Security',
        icon: 'pi pi-shield',
        items: [
          {
            label: 'Authentication',
            icon: 'pi pi-key',
            command: () => router.push('/settings/security/authentication'),
          },
          {
            label: 'Permissions',
            icon: 'pi pi-lock',
            command: () => router.push('/settings/security/permissions'),
          },
          {
            label: 'Audit Logs',
            icon: 'pi pi-history',
            command: () => router.push('/settings/security/audit'),
          },
        ],
      },
    ],
  },
  {
    label: 'Monitoring',
    icon: 'pi pi-chart-bar',
    items: [
      {
        label: 'System Health',
        icon: 'pi pi-heart',
        command: () => router.push('/monitoring/health'),
      },
      {
        label: 'Logs',
        icon: 'pi pi-list',
        command: () => router.push('/monitoring/logs'),
      },
      {
        label: 'Alerts',
        icon: 'pi pi-exclamation-triangle',
        command: () => router.push('/monitoring/alerts'),
      },
    ],
  },
  {
    label: 'Support',
    icon: 'pi pi-question-circle',
    items: [
      {
        label: 'Documentation',
        icon: 'pi pi-book',
        command: () => router.push('/support/docs'),
      },
      {
        label: 'Help Center',
        icon: 'pi pi-info-circle',
        command: () => router.push('/support/help'),
      },
      {
        label: 'Contact Support',
        icon: 'pi pi-envelope',
        command: () => router.push('/support/contact'),
      },
    ],
  },
])

// Toggle sidebar visibility (for mobile)
const toggleSidebar = () => {
  visible.value = !visible.value
}

// Toggle between full and mini sidebar (for desktop)
const toggleMiniSidebar = () => {
  miniSidebar.value = !miniSidebar.value
}

// Check screen size and set mobile state
const checkScreenSize = () => {
  isMobile.value = window.innerWidth < 768
  if (!isMobile.value && !visible.value) {
    visible.value = true
  }
}

onMounted(() => {
  checkScreenSize()
  window.addEventListener('resize', checkScreenSize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', checkScreenSize)
})
</script>

<template>
  <div class="flex h-screen w-full relative bg-slate-50 dark:bg-slate-900">
    <!-- Top navbar for mobile -->
    <div
      v-if="isMobile"
      class="fixed top-0 left-0 right-0 flex items-center p-4 bg-white dark:bg-slate-800 shadow-md z-50"
    >
      <Button
        icon="pi pi-bars"
        @click="toggleSidebar"
        text
        severity="secondary"
        aria-label="Menu"
        class="mr-4"
      />
      <h2 class="text-xl font-semibold">Storm Dashboard</h2>
    </div>

    <!-- Sidebar using PrimeVue Sidebar component for mobile -->
    <Sidebar v-if="isMobile" v-model:visible="visible" :modal="true" class="w-[280px] p-0">
      <div class="flex flex-col h-full">
        <div class="flex items-center justify-between p-4">
          <div class="flex items-center">
            <i class="pi pi-bolt text-indigo-500 text-2xl mr-3"></i>
            <span class="text-xl font-semibold">Storm Dashboard</span>
          </div>
        </div>

        <div class="flex-grow overflow-y-auto">
          <Menu :model="menuItems" class="w-full mobile-sidebar-menu" />
        </div>

        <div class="mt-auto p-4 flex items-center border-t border-slate-200 dark:border-slate-700">
          <Avatar icon="pi pi-user" size="large" shape="circle" />
          <div class="ml-3">
            <div class="font-medium">Admin User</div>
            <div class="text-sm text-slate-600 dark:text-slate-400">Administrator</div>
          </div>
        </div>
      </div>
    </Sidebar>

    <!-- Desktop sidebar -->
    <div
      v-if="!isMobile"
      class="h-screen transition-all duration-300 border-r border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 flex flex-col"
      :class="miniSidebar ? 'w-[70px]' : 'w-[280px]'"
    >
      <div
        class="flex items-center justify-between p-4 border-b border-slate-200 dark:border-slate-700"
      >
        <div class="flex items-center overflow-hidden">
          <i class="pi pi-bolt text-indigo-500 text-2xl" :class="miniSidebar ? '' : 'mr-3'"></i>
          <span
            class="text-xl font-semibold whitespace-nowrap overflow-hidden transition-opacity duration-300"
            :class="miniSidebar ? 'opacity-0 w-0' : 'opacity-100'"
          >
            Storm Dashboard
          </span>
        </div>
        <Button
          :icon="miniSidebar ? 'pi pi-angle-right' : 'pi pi-angle-left'"
          @click="toggleMiniSidebar"
          text
          severity="secondary"
          aria-label="Toggle sidebar"
          class="p-button-sm"
        />
      </div>

      <div class="flex-grow overflow-y-auto">
        <!-- Custom mini sidebar menu when collapsed -->
        <div v-if="miniSidebar" class="py-2">
          <div
            v-for="(item, i) in menuItems"
            :key="i"
            class="flex justify-center py-3 cursor-pointer hover:bg-indigo-50 dark:hover:bg-indigo-900/30"
            @click="item.command ? item.command() : ''"
            @mouseover="item.items && menuRef.value && menuRef.value.toggle($event, item)"
          >
            <i :class="[item.icon, 'text-lg']"></i>
          </div>
        </div>
        <!-- Regular Menu when expanded -->
        <div v-else class="p-2">
          <Menu :model="menuItems" class="w-full sidebar-menu" />
        </div>
      </div>

      <!-- User profile section - fixed at bottom -->
      <div
        class="p-4 flex items-center border-t border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800"
      >
        <Avatar icon="pi pi-user" :size="miniSidebar ? 'normal' : 'large'" shape="circle" />
        <div
          class="ml-3 whitespace-nowrap overflow-hidden transition-opacity duration-300"
          :class="miniSidebar ? 'opacity-0 w-0' : 'opacity-100'"
        >
          <div class="font-medium">Admin User</div>
          <div class="text-sm text-slate-600 dark:text-slate-400">Administrator</div>
        </div>
      </div>
    </div>

    <!-- Main content area -->
    <div
      class="flex-1 overflow-y-auto transition-all duration-300 p-6 bg-slate-50 dark:bg-slate-900"
      :class="isMobile ? 'pt-20' : ''"
    >
      <slot></slot>
    </div>
  </div>
</template>

<style scoped lang="postcss">
:deep(.sidebar-menu),
:deep(.mobile-sidebar-menu) {
  width: 100%;
  border: none;
  background-color: transparent;
}

:deep(.sidebar-menu .p-menuitem-link),
:deep(.mobile-sidebar-menu .p-menuitem-link) {
  padding: 0.75rem 1rem;
}

:deep(.sidebar-menu .p-menu-list),
:deep(.mobile-sidebar-menu .p-menu-list) {
  padding: 0;
}

/* Style for nested submenu */
:deep(.p-submenu-list) {
  background-color: inherit !important;
  margin-left: 1rem;
}

:deep(.p-menuitem-active > .p-menuitem-link) {
  background-color: rgba(99, 102, 241, 0.1) !important;
}

/* Fix for mini sidebar submenu display */
:deep(.p-menu-overlay) {
  position: fixed;
  z-index: 999;
}
</style>
