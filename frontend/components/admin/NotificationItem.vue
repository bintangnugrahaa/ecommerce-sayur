<template>
    <li :class="[
      'px-5 py-4 hover:bg-gray-100 border-b border-gray-300',
      { 'bg-gray-100 active': notification.isRead === false }
    ]">
      <NuxtLink :to="notification.link" class="text-gray-500">
        <div class="flex">
          <img 
            :src="notification.avatar" 
            :alt="notification.title"
            class="h-10 w-10 rounded-full" 
          />
          <div class="ms-4">
            <p class="mb-1">
              <span class="text-gray-900">{{ notification.title }}</span>
              {{ notification.message }}
            </p>
            <span class="flex items-center">
              <svg 
                xmlns="http://www.w3.org/2000/svg" 
                width="12" 
                height="12" 
                fill="currentColor" 
                class="bi bi-clock text-gray-500" 
                viewBox="0 0 16 16"
              >
                <path d="M8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71V3.5z" />
                <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm7-8A7 7 0 1 1 1 8a7 7 0 0 1 14 0z" />
              </svg>
              <small class="ms-2">{{ notification.time }}</small>
            </span>
          </div>
        </div>
      </NuxtLink>
    </li>
</template>
  
<script setup>
  defineProps({
    notification: {
      type: Object,
      required: true,
      default: () => ({
        id: '',
        avatar: '',
        title: '',
        message: '',
        time: '',
        isRead: false,
        link: '#'
      })
    }
  })
</script>