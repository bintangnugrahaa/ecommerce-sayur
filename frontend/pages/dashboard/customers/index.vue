<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
      <!-- page header -->
      <div class="md:flex justify-between items-center">
        <div>
          <h2 class="text-xl">Customers</h2>
          <!-- breacrumb -->
          <nav aria-label="breadcrumb">
            <ol class="flex flex-wrap">
              <li class="inline-block text-green-600">
                <a href="/dashboard">
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

              <li class="inline-block text-gray-500 active" aria-current="page">Customers</li>
            </ol>
          </nav>
        </div>
        <!-- button -->
        <div class="mt-3 lg:mt-0">
          <a href="/dashboard/customers/create"
            class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100">
            Add New Customer
          </a>
        </div>
      </div>
    </div>
    <!-- row -->
    <div class="grid grid-cols-1">
      <!-- card -->
      <div class="card h-full card-lg">
        <div class="px-6 py-6">
          <div class="grid grid-cols-12 justify-between gap-2">
            <div class="lg:col-span-3 md:col-span-6 col-span-12">
              <!-- form -->
              <form class="flex" role="search">
                <input
                  v-model="searchFilter"
                  @input="handleSearch"
                  class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                  type="search" placeholder="Search Customers" aria-label="Search" />
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
                  <th scope="col" class="px-6 py-3">Name</th>
                  <th scope="col" class="px-6 py-3">Email</th>
                  <th scope="col" class="px-6 py-3">Phone</th>

                  <th scope="col" class="px-6 py-3"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="response?.loading">
                    <td colspan="8" class="text-center py-4">Loading...</td>
                </tr>
                <tr v-else-if="response?.error">
                    <td colspan="8" class="text-center py-4 text-red-500">{{ customerStore.error }}</td>
                </tr>
                <tr v-else-if="response?.data?.length === 0">
                    <td colspan="8" class="text-center py-4">No data available</td>
                </tr>
                <tr v-else v-for="item in customers">
                  <td class="py-3 px-6 text-left">
                    <div class="flex items-center">
                      <img :src="item.photo" alt="" class="h-8 w-8 rounded-full" />
                    </div>
                  </td>
                  <td class="py-3 px-6 text-left">{{ item.email }}</td>
                  <td class="py-3 px-6 text-left">{{ item.phone }}</td>

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
                          <a
                            @click="handleDelete(item.id)"
                            class="dropdown-item" href="#">
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
                          <NuxtLink class="dropdown-item" :to="`/dashboard/customers/edit/${item.id}`">
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
        <div v-if="customerStore.pagination?.total_count > 0" class="border-t border-gray-300 flex flex-col md:flex-row justify-between items-center px-6 py-6 gap-3">
          <span>Showing  {{ pagination.page }} to  {{ pagination.total_page }} of  {{ pagination.total_count }} entries</span>
          <nav class="flex items-center gap-x-1">
            <button
              :disabled="pagination.page === 1"
              @click="handlePageChange(pagination.page - 1)"
              type="button"
              class="leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300 disabled:opacity-50 disabled:pointer-events-none"
              disabled="">
              Previous
            </button>
            <div class="flex items-center gap-x-1">
              <button v-for="page in pagination.total_page" :key="page"
                @click="handlePageChange(page)"
                type="button"
                :class="[
                    'leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border',
                    page === pagination.page 
                        ? 'text-white border bg-green-600 border-green-600 hover:bg-green-600 focus:outline-none focus:bg-green-600'
                        : 'bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300'
                ]"
                aria-current="page">
                {{ page }}
              </button>
            </div>
            <button
              :disabled="pagination.page === pagination.total_page"
              @click="handlePageChange(pagination.page + 1)"
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

import {useCustomerStore} from '~/stores/customer'
import { ref, onMounted } from 'vue'

const customerStore = useCustomerStore()
const customers = ref([])
const pagination = ref({})
const searchFilter = ref('')

onMounted(async () => {
  const response = await customerStore.fetchCustomers()
  customers.value = response.data
  pagination.value = response.pagination
})

const handlePageChange = async (page) => {
  const responseProds = await customerStore.fetchCustomers({
    page: page,
    search: searchFilter.value,
  })
  
  customers.value = responseProds.data
  pagination.value = responseProds.pagination
}

const handleSearch = async () => {
  const responseProds = await customerStore.fetchCustomers({
    page: 1,
    search: searchFilter.value,
  })

  customers.value = responseProds.data
  pagination.value = responseProds.pagination
}

const handleDelete = async (id) => {
  try {
    const isConfirmed = window.confirm('Apakah Anda yakin ingin menghapus kategori ini?')
    
    if (isConfirmed) {
      await customerStore.deleteCustomer(id)
      const response = await customerStore.fetchCustomers()
      customers.value = response.data
      pagination.value = response.pagination
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