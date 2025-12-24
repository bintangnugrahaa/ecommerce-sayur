<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
        <!-- page header -->
        <div class="flex flex-col md:flex-row md:justify-between md:items-center gap-4">
            <div>
                <h2 class="text-xl">Products</h2>
                <!-- breacrumb -->
                <nav aria-label="breadcrumb">
                    <ol class="flex flex-wrap">
                        <li class="inline-block text-green-600">
                            <a href="../../dashboard">
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

                        <li class="inline-block text-gray-500 active" aria-current="page">Role</li>
                    </ol>
                </nav>
            </div>
            <!-- button -->
            <div>
                <!-- <a href="/dashboard/role/create"
                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100">
                    Add Role
                </a> -->
            </div>
        </div>
    </div>
    <!-- row -->
    <div class="grid grid-cols-1">
        <!-- card -->
        <div class="card h-full card-lg">
            <!-- card body -->
          <div class="card-body p-0">
              <!-- table -->
              <div class="relative overflow-x-auto">
                  <table class="text-left w-full whitespace-nowrap table-with-checkbox table-hover">
                      <thead class="bg-gray-200 text-gray-700">
                          <tr class="border-transparent !border-b-0">
                              <th scope="col" class="px-6 py-3">Name</th>
                              <th scope="col" class="px-6 py-3 text-right">Action</th>
                          </tr>
                      </thead>
                      <tbody class="divide-y">
                        <tr v-if="response?.loading">
                            <td colspan="8" class="text-center py-4">Loading...</td>
                        </tr>
                        <tr v-else-if="response?.error">
                            <td colspan="8" class="text-center py-4 text-red-500">{{ useRole.error }}</td>
                        </tr>
                        <tr v-else-if="response?.data?.length === 0">
                            <td colspan="8" class="text-center py-4">No data available</td>
                        </tr>
                        <tr v-else v-for="role in rolesData" :key="role.id" class="border-transparent !border-b-0">

                            <td class="py-3 px-6 text-left">{{ role.name }}</td>
                            <td class="py-3 px-6 text-right flex justify-end items-center gap-2">
                              <div class="inline-flex items-center gap-x-2 ">
                                <a class="text-gray-700 hover:text-green-600" href="#">
                                  <svg xmlns="http://www.w3.org/2000/svg" width="14"
                                      height="14" viewBox="0 0 24 24" fill="none"
                                      stroke="currentColor" stroke-width="2"
                                      stroke-linecap="round" stroke-linejoin="round"
                                      class="icon icon-tabler icons-tabler-outline icon-tabler-edit">
                                      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                      <path d="M7 7h-1a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-1" />
                                      <path d="M20.385 6.585a2.1 2.1 0 0 0 -2.97 -2.97l-8.415 8.385v3h3l8.385 -8.415z" />
                                      <path d="M16 5l3 3" />
                                  </svg>
                                  Edit
                                </a>
                                <a class="text-gray-700 hover:text-red-600" href="#">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="14"
                                        height="14" viewBox="0 0 24 24" fill="none"
                                        stroke="currentColor" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round"
                                        class="icon icon-tabler icons-tabler-outline icon-tabler-trash">
                                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                        <path d="M4 7l16 0" />
                                        <path d="M10 11l0 6" />
                                        <path d="M14 11l0 6" />
                                        <path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
                                        <path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
                                    </svg>
                                    Delete
                                </a>
                              </div>
                          </td>
                        </tr>
                      </tbody>
                  </table>
              </div>
          </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRoleStore } from '~/stores/role'
import { ref, onMounted } from 'vue'

const useRole = useRoleStore()
const rolesData = ref([])

onMounted(async () => {
  const response = await useRole.fetchRoles()
  rolesData.value = response.data
})

definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

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