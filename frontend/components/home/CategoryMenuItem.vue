<template>
    <li class="border-b w-full">
      <a :class="{ 'active': isActive }"
      @click.prevent="$emit('select', category.slug)" href="#" class="flex justify-between py-2 collapsed items-center text-gray-800 hover:text-green-600"
        :data-bs-toggle="'collapse'" :data-bs-target="`#${category.slug}`" aria-expanded="false"
        :aria-controls="category.slug">
        {{ category.name }}
        <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-chevron-right inline-block"
          width="14" height="14" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none"
          stroke-linecap="round" stroke-linejoin="round">
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <path d="M9 6l6 6l-6 6" />
        </svg>
      </a>
      <div :id="category.slug" class="accordion-collapse collapse" data-bs-parent="#categoryCollapseMenu">
        <div class="text-gray-800">
          <ul class="flex flex-wrap flex-col ml-3">
            <li v-for="(item, index) in category.childs" :key="index" >
              <a href="#!" class="inline-block py-2 no-underline hover:text-green-600"
              :class="{ 'active': isActive }"
              @click.prevent="$emit('select', item.slug)">
                {{ item.name }}
              </a>
            </li>
          </ul>
        </div>
      </div>
    </li>
  </template>
  
  <script setup>
  defineProps({
    category: {
      type: Object,
      required: true
    },
    isActive: {
      type: Boolean,
      default: false
    }
  })
  </script>