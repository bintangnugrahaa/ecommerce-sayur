<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
        <!-- page header -->
        <div class="flex flex-col md:flex-row md:justify-between md:items-center gap-4">
            <div>
                <h2 class="text-xl">Categories</h2>
                <!-- breacrumb -->
                <nav aria-label="breadcrumb">
                    <ol class="flex flex-wrap">
                        <li class="inline-block text-green-600">
                            <a href="#!">
                                Dashboard
                                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14"
                                    viewBox="0 0 24 24" fill="none" stroke="currentColor"
                                    stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                                    class="icon icon-tabler icons-tabler-outline icon-tabler-slash inline-block mx-2">
                                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                    <path d="M17 5l-10 14" />
                                </svg>
                            </a>
                        </li>

                        <li class="inline-block text-gray-500 active" aria-current="page">Categories</li>
                    </ol>
                </nav>
            </div>
            <!-- button -->
            <div>
                <a href="/dashboard/categories/create"
                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100">
                    Add Categories
                </a>
            </div>
        </div>
    </div>
    <!-- row -->
    <div class="grid grid-cols-1">
        <!-- card -->
        <div class="card h-full card-lg">
            <div class="px-6 py-6">
                <div class="grid grid-cols-12 justify-between">
                    <!-- form -->
                    <div class="lg:col-span-3 md:col-span-6 col-span-12 mb-2 lg:mb-0">
                        <form class="flex" role="search">
                            <input
                                v-model="searchFilter"
                                @input="handleSearchFilter"
                                class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                                type="search" placeholder="Search Products" aria-label="Search" />
                        </form>
                    </div>
                </div>
            </div>
            <!-- card body -->
            <div class="card-body p-0">
                <!-- table -->
                <div class="relative overflow-x-auto">
                    <table class="text-left w-full whitespace-nowrap table-with-checkbox table-hover">
                        <thead class="bg-gray-200 text-gray-700">
                            <tr class="border-transparent !border-b-0">
                                <th scope="col" class="px-6 py-3">
                                    <div class="flex items-center">
                                        <input
                                            class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded focus:ring-green-600 focus:outline-none focus:ring-2"
                                            type="checkbox" value="" id="checkAll" />
                                        <label class="text-gray-800 ms-3" for="checkAll"></label>
                                    </div>
                                </th>
                                <th scope="col" class="px-6 py-3">Icon</th>
                                <th scope="col" class="px-6 py-3">Name</th>
                                <th scope="col" class="px-6 py-3">Total Product</th>
                                <th scope="col" class="px-6 py-3">Status</th>
                                <th scope="col" class="px-6 py-3"></th>
                            </tr>
                        </thead>
                        <tbody class="divide-y">
                            <tr v-if="response?.loading">
                                <td colspan="8" class="text-center py-4">Loading...</td>
                            </tr>
                            <tr v-else-if="response?.error">
                                <td colspan="8" class="text-center py-4 text-red-500">{{ useCategory.error }}</td>
                            </tr>
                            <tr v-else-if="response?.data?.length === 0">
                                <td colspan="8" class="text-center py-4">No data available</td>
                            </tr>
                            <tr v-else v-for="item in categoryDatas" class="border-transparent !border-b-0">
                                <td class="py-3 px-6 text-center">
                                    <div class="flex items-center">
                                        <input
                                            class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded focus:ring-green-600 focus:outline-none focus:ring-2"
                                            type="checkbox" value="" id="productOne" />
                                        <label class="text-gray-800 ms-3" for="productOne"></label>
                                    </div>
                                </td>
                                <td class="py-3 px-6 text-left">
                                    <NuxtLink :to="`/dashboard/products/${item.id}`">
                                        <img :src="item.icon" alt="" class="h-12 w-12" />
                                    </NuxtLink>
                                </td>
                                <td class="py-3 px-6 text-left">
                                    <NuxtLink :to="`/dashboard/products/${item.id}`" class="text-inherit">
                                        {{ item.name }}
                                </NuxtLink>
                                </td>
                                <td class="py-3 px-6 text-left">
                                  {{ item.total_product }}
                                </td>
                                <td class="py-3 px-6 text-left">
                                  <span class="inline-block p-1 text-sm align-baseline leading-none rounded border font-semibold"
                                  :class="[
                                    item.status === 'Published' 
                                      ? 'bg-green-100 text-green-800 border-green-200' 
                                      : 'bg-red-100 text-red-800 border-red-200'
                                  ]">
                                    {{ item.status }}
                                  </span>
                                </td>
                                <td class="py-3 px-6 text-left">
                                    <div class="dropdown">
                                        <a href="#" class="text-inherit" data-bs-toggle="dropdown"
                                            aria-expanded="false">
                                            <svg xmlns="http://www.w3.org/2000/svg" width="20"
                                                height="20" viewBox="0 0 24 24" fill="none"
                                                stroke="currentColor" stroke-width="2"
                                                stroke-linecap="round" stroke-linejoin="round"
                                                class="icon icon-tabler icons-tabler-outline icon-tabler-dots-vertical">
                                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                                <path d="M12 12m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
                                                <path d="M12 19m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
                                                <path d="M12 5m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
                                            </svg>
                                        </a>
                                        <ul class="dropdown-menu">
                                            <li>
                                                <a @click.prevent="handleDelete(item.id)" class="dropdown-item" href="#">
                                                    <svg xmlns="http://www.w3.org/2000/svg" width="14"
                                                        height="14" viewBox="0 0 24 24" fill="none"
                                                        stroke="currentColor" stroke-width="2"
                                                        stroke-linecap="round" stroke-linejoin="round"
                                                        class="icon icon-tabler icons-tabler-outline icon-tabler-trash">
                                                        <path stroke="none" d="M0 0h24v24H0z"
                                                            fill="none" />
                                                        <path d="M4 7l16 0" />
                                                        <path d="M10 11l0 6" />
                                                        <path d="M14 11l0 6" />
                                                        <path
                                                            d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
                                                        <path
                                                            d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
                                                    </svg>
                                                    Delete
                                                </a>
                                            </li>
                                            <li>
                                                <NuxtLink class="dropdown-item" :to="`/dashboard/categories/edit/${item.id}`">
                                                    <svg xmlns="http://www.w3.org/2000/svg" width="14"
                                                        height="14" viewBox="0 0 24 24" fill="none"
                                                        stroke="currentColor" stroke-width="2"
                                                        stroke-linecap="round" stroke-linejoin="round"
                                                        class="icon icon-tabler icons-tabler-outline icon-tabler-edit">
                                                        <path stroke="none" d="M0 0h24v24H0z"
                                                            fill="none" />
                                                        <path
                                                            d="M7 7h-1a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-1" />
                                                        <path
                                                            d="M20.385 6.585a2.1 2.1 0 0 0 -2.97 -2.97l-8.415 8.385v3h3l8.385 -8.415z" />
                                                        <path d="M16 5l3 3" />
                                                    </svg>
                                                    Edit
                                                </NuxtLink>
                                            </li>
                                        </ul>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
            <div v-if="useCategory.pagination?.total_count > 0" class="border-t border-gray-300 flex flex-col md:flex-row justify-between items-center px-6 py-6 gap-3">
                <span>Showing {{ paginateProds.page }} to {{ paginateProds.total_page }} of {{ paginateProds.total_count }} entries</span>
                <nav class="flex items-center gap-x-1">
                    <button :disabled="paginateProds.page === 1"
                        @click="handlePageChange(paginateProds.page - 1)"
                        type="button"
                        class="leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300 disabled:opacity-50 disabled:pointer-events-none"
                        disabled="">
                        Previous
                    </button>
                    <div class="flex items-center gap-x-1">
                        <button v-for="page in paginateProds.total_page" :key="page"
                            @click="handlePageChange(page)"
                            type="button"
                            :class="[
                                'leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border',
                                page === paginateProds.page 
                                    ? 'text-white border bg-green-600 border-green-600 hover:bg-green-600 focus:outline-none focus:bg-green-600'
                                    : 'bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300'
                            ]"
                            aria-current="page">
                            {{ page }}
                        </button>
                    </div>
                    <button
                        :disabled="paginateProds.page === paginateProds.total_page"
                        @click="handlePageChange(paginateProds.page + 1)"
                        type="button"
                        class="leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300 disabled:opacity-50 disabled:pointer-events-none">
                        Next
                    </button>
                </nav>
            </div>
        </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

import { useCategoryStore } from '~/stores/category'
import { ref, onMounted } from 'vue'

const useCategory = useCategoryStore()
const categoryDatas = ref([])
const paginateProds = ref({})
const searchFilter = ref('')

const handleSearchFilter = async () => {
  const response = await useCategory.fetchCategoriesAdmin({
    search: searchFilter.value,
    page: 1,
  })
  categoryDatas.value = response.data
  paginateProds.value = response.pagination
}

onMounted(async () => {
  const response = await useCategory.fetchCategoriesAdmin()
  categoryDatas.value = response.data
  paginateProds.value = response.pagination
})

const handlePageChange = async (page) => {
  const responseProds = await useCategory.fetchCategoriesAdmin({
    page: page,
  })
  
  categoryDatas.value = responseProds.data
  paginateProds.value = responseProds.pagination
}

const handleDelete = async (id) => {
  try {
    const isConfirmed = window.confirm('Apakah Anda yakin ingin menghapus kategori ini?')
    
    if (isConfirmed) {
      await useCategory.deleteCategory(id)
      const response = await useCategory.fetchCategoriesAdmin()
      categoryDatas.value = response.data
      paginateProds.value = response.pagination
    }
  } catch (error) {
    alert(error.message || 'Terjadi kesalahan saat menghapus kategori')
  }
}

useHead({
  script: [
  	{
		src: '/libs/simplebar/dist/simplebar.min.js',
		body: true
    },
    {
		src: '/js/theme.min.js',
		body: true
    },
  ],
})
</script>

<style>

</style>