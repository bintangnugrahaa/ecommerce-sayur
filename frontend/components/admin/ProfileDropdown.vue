<template>
    <li class="dropdown">
      <button @click="toggleDropdown" class="flex items-center">
        <img :src="userAvatar" alt="User avatar" class="h-12 w-12 rounded-full" />
      </button>
  
      <div v-if="isOpen" class="dropdown-menu dropdown-menu-end !p-0">
        <div class="leading-snug px-5 py-4 border-b border-gray-300">
          <h5 class="mb-1 text-base">{{ userName }}</h5>
          <small>{{ userEmail }}</small>
        </div>
  
        <ul class="list-unstyled px-2 py-3">
          <li v-for="item in menuItems" :key="item.path">
            <NuxtLink 
              :to="item.path"
              class="dropdown-item"
            >
              {{ item.label }}
            </NuxtLink>
          </li>
        </ul>
        <div class="border-t px-5 py-3">
          <button @click="handleLogout">Log Out</button>
        </div>
      </div>
    </li>
  </template>
  
  <script setup>
  const isOpen = ref(false)
  import { ref } from 'vue'
  import { onClickOutside } from '@vueuse/core'
  import avatar from '~/assets/images/avatar/avatar-1.jpg'
  
  const userAvatar = avatar
  const userName = 'FreshCart Admin'
  const userEmail = 'admindemo@email.com'
  
  const menuItems = [
    { label: 'Home', path: '/' },
    { label: 'Profile', path: '/profile' },
    { label: 'Settings', path: '/settings' }
  ]
  
  const toggleDropdown = () => {
    isOpen.value = !isOpen.value
  }
  
  const handleLogout = () => {
    // Implementasi logout
  }
  
  onClickOutside(isOpen, () => {
    isOpen.value = false
  })
  </script>