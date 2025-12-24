<script setup>
import Navbar from '~/components/admin/Navbar.vue';
import AdminSidebarNav from '~/components/admin/SidebarNav.vue';
import { useAuthStore } from '~/stores/auth'

const authStore = useAuthStore()
authStore.checkAuth()

const toggleSidebar = () => {
  // Implementasi toggle sidebar
}

</script>

<template>
    <div>
    <Navbar @toggle-sidebar="toggleSidebar" />
    <div class="main-wrapper">
      <AdminSidebarNav/>
      <main class="main-content-wrapper">
        <slot />
      </main>
    </div>
  </div>
</template>