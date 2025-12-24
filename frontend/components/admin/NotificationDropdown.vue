<template>
    <li class="dropdown-center">
      <a 
        class="relative btn text-gray-600 hover:bg-gray-100 hover:text-green-600 h-12 w-12 p-0 rounded-full" 
        href="#" 
        @click="toggleDropdown"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          class="icon icon-tabler icons-tabler-outline icon-tabler-bell"
        >
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6" />
          <path d="M9 17v1a3 3 0 0 0 6 0v-1" />
        </svg>
        <span class="absolute -top-2 right-0 translate-middle inline-block text-sm align-baseline bg-red-600 text-white font-semibold rounded-full h-5 w-5 text-center">
          {{ unreadCount }}
        </span>
      </a>
  
      <div v-if="isOpen" class="dropdown-menu dropdown-menu-end dropdown-menu-lg !p-0 border-0">
        <div class="border-b p-5 flex justify-between items-center">
          <div>
            <h5 class="mb-1">Notifications</h5>
            <p class="mb-0 small">You have {{ unreadCount }} unread messages</p>
          </div>
          <button
            class="btn text-gray-600 hover:bg-gray-100 hover:text-green-600 h-10 w-10 p-0 rounded-full"
            @click="markAllAsRead"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="icon icon-tabler icons-tabler-outline icon-tabler-checks"
            >
              <path stroke="none" d="M0 0h24v24H0z" fill="none" />
              <path d="M7 12l5 5l10 -10" />
              <path d="M2 12l5 5m5 -5l5 -5" />
            </svg>
          </button>
        </div>
        <div class="notification-list max-h-[250px] overflow-y-auto">
          <NotificationItem 
            v-for="notification in notifications" 
            :key="notification.id"
            :notification="notification"
          />
        </div>
        <div class="px-5 py-4 text-center">
          <NuxtLink to="/notifications">View All</NuxtLink>
        </div>
      </div>
    </li>
  </template>
  
<script setup>
  import avatar from '~/assets/images/avatar/avatar-1.jpg'
  import { ref } from 'vue'
  import { onClickOutside } from '@vueuse/core'
  import NotificationItem from './NotificationItem.vue'
  const isOpen = ref(false)
  const unreadCount = ref(2)
  
  const notifications = ref([
    {
      id: 1,
      avatar: avatar,
      title: 'Your order is placed',
      message: 'waiting for shipping',
      time: '1 minute ago',
      isRead: false
    },
    // ... tambahkan notifikasi lainnya
  ])
  
  const toggleDropdown = () => {
    isOpen.value = !isOpen.value
  }
  
  const markAllAsRead = () => {
    notifications.value = notifications.value.map(notification => ({
      ...notification,
      isRead: true
    }))
    unreadCount.value = 0
  }
  
  onClickOutside(isOpen, () => {
    isOpen.value = false
  })
</script>