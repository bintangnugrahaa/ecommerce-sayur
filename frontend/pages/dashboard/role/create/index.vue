<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
<!-- page header -->
      <div class="md:flex justify-between items-center">
        <div>
          <h2 class="text-xl">Add New Role</h2>
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
              <li class="inline-block text-green-600">
                <a href="/dashboard/role">
                    Role
                  <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor"
                    stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-slash inline-block mx-2">
                  <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                  <path d="M17 5l-10 14" />
                  </svg>
                </a>
              </li>
              <li class="inline-block text-gray-500 active" aria-current="page">Add New Role</li>
            </ol>
          </nav>
        </div>
        <!-- button -->
      </div>
    </div>
					<!-- row -->
    <div class="grid grid-cols-1">
      <div class="card card-lg border-0">
        <div class="card-body flex flex-col gap-8 p-7">
          <div class="flex flex-col gap-4">
            <h3 class="mb-0 text-md">Role</h3>
            <form @submit.prevent="handleSubmit" class="grid grid-cols-12 gap-6 needs-validation" novalidate>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <!-- input -->
                  <label for="createRoleName"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Role Name
                    <span class="text-red-600">*</span>
                  </label>
                  <input v-model="name" type="text"
                    class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                    id="createRoleName" placeholder="Role Name" required />
                  <div class="invalid-feedback">Please enter role name</div>
                </div>
              </div>
              <div class="col-span-12 mt-3">
                <div class="flex flex-col md:flex-row gap-2">
                  <button :disabled="useRole.loading"
                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100"
                    type="submit">
                    <span v-if="useRole.loading">Loading...</span>
                    <span v-else>Create New Role</span>
                  </button>
                  <a href="/dashboard/role"
                    class="btn inline-flex items-center gap-x-2 bg-gray-200 text-gray-800 border-gray-200 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300"
                    type="submit">
                    Cancel
                  </a>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
	</div>
</template>

<script setup>
import { useRoleStore } from '~/stores/role'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const useRole = useRoleStore()

definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

const router = useRouter()
const name = ref('')

const handleSubmit = async () => {
  try {
    await useRole.createRole(name.value)

    setTimeout(() => {
      router.push('/dashboard/role')
    }, 1000)
  } catch (error) {
    console.log(error)
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