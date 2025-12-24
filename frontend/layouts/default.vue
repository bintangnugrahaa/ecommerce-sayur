<template>
  <div>
    <Header />
    <slot />
    <Footer />
  </div>
</template>

<script setup>
import Footer from '~/components/common/Footer.vue';
import Header from '~/components/common/Header.vue';
import { useAuthStore } from '~/stores/auth'

const authStore = useAuthStore()
authStore.checkAuth()
</script>