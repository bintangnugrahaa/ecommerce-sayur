<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
      <!-- page header -->
      <div class="md:flex justify-between items-center">
        <div>
          <h2 class="text-xl">Add New Category</h2>
          <!-- breacrumb -->
          <nav aria-label="breadcrumb">
            <ol class="flex flex-wrap">
              <li class="inline-block text-green-600">
                <a href="/dashboard/categories">
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
                <a href="../dashboard/categories.html">
                  Categories
                  <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor"
                    stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-slash inline-block mx-2">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M17 5l-10 14" />
                  </svg>
                </a>
              </li>

              <li class="inline-block text-gray-500 active" aria-current="page">Add New Category</li>
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
          <div v-if="useCategory.error" class="mb-4 p-4 text-sm text-red-800 rounded-lg bg-red-50 relative" role="alert">
            <div class="flex items-center pr-8">
                <span class="font-medium mr-2">Gagal!</span> {{ useCategory.error }}
            </div>
          </div>
          <div class="flex flex-col md:flex-row items-center mb-4 file-input-wrapper gap-2">
            <div>
              <img class="image h-16 w-16 rounded-lg" v-if="imageUrl" :src="imageUrl" alt="Image" />
            </div>

            <div
              class="file-upload btn inline-flex items-center gap-x-2 bg-gray-200 text-gray-800 border-gray-200 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300 md:ml-4">
              <input @change="handleFileUpload" id="icon" type="file" class="file-input opacity-0" accept="image/*"/>
              Upload Photo
            </div>

            <span class="ms-2">JPG, GIF or PNG. 1MB Max.</span>
          </div>
          <div class="flex flex-col gap-4">
            <h3 class="mb-0 text-md">Category Information</h3>
            <form @submit.prevent="handleSubmit" class="grid grid-cols-12 gap-6 needs-validation" novalidate>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <!-- input -->
                  <label for="creatCustomerName"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Category Name
                    <span class="text-red-600">*</span>
                  </label>
                  <input type="text"
                    v-model="form.name"
                    class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                    id="creatCustomerName" placeholder="Customer Name" required />
                  <div class="invalid-feedback">Please enter category name</div>
                </div>
              </div>
              <div class="lg:col-span-12 col-span-12">
                <div>
                  <!-- input -->
                  <label for="creatCustomerPhone" class="inline-block text-gray-800 font-medium mb-2">Description</label>
                  <textarea v-model="form.description" name="description" class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base" id=""></textarea>
                </div>
              </div>

              <div class="mb-3 col-lg-12">
                <label class="form-label" for="creatCustomerDate">Status</label>
                <div class="flex space-x-6">
                    <!-- Active Radio Button -->
                  <label class="inline-flex items-center">
                    <input v-model="form.status" type="radio" class="form-radio focus:outline-none text-green-600" name="status" :value="1">
                    <span class="ml-2">Active</span>
                  </label>
                    
                    <!-- Disabled Radio Button -->
                  <label class="inline-flex items-center">
                    <input v-model="form.status" type="radio" class="form-radio focus:outline-none text-green-600" name="status" :value="0">
                    <span class="ml-2">Disabled</span>
                  </label>
                </div>
              </div>

              <div class="col-span-12 mt-3">
                <div class="flex flex-col md:flex-row gap-2">
                  <button
                    :disabled="useCategory.loading"
                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100"
                    type="submit">
                    <span v-if="useCategory.loading">Loading...</span>
                    <span v-else>Create New Category</span>
                  </button>
                  <a href="/dashboard/categories"
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
import { useCategoryStore } from '~/stores/category'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

const router = useRouter()
const file = ref(null)
const imageUrl = ref('')
const useCategory = useCategoryStore()

const form = reactive({
  name: '',
  description: '',
  status: 1,
  icon: '',
  parent_id: '',
})

const handleFileUpload = async (event) => {
  try {
    const selectedFile = event.target.files[0]
    if (selectedFile) {
      // Validasi ukuran file (max 1MB)
      if (selectedFile.size > 1024 * 1024) {
        alert('Ukuran file maksimal 1MB')
        return
      }
      
      // Validasi tipe file
      if (!['image/jpeg', 'image/png', 'image/gif'].includes(selectedFile.type)) {
        alert('File harus berupa JPG, PNG, atau GIF')
        return
      }

      // Upload file
      const result = await useCategory.uploadImage(selectedFile)
      
      // Update preview dan simpan URL
      imageUrl.value = result.data.image_url
      file.value = selectedFile
    }
  } catch (error) {
    console.error('Error upload:', error)
    alert(error.message || 'Gagal mengupload gambar')
  }
}

const handleSubmit = async () => {
  try {
    let status = 'Published'
    if (form.status === 0) {
      status = 'Unpublished'
    }

    const category = {
      name: form.name,
      description: form.description,
      status: status,
      icon: imageUrl.value,
      parent_id: null
    }

    await useCategory.createCategory(category)

    form.name = ''
    form.description = ''
    form.status = 1
    form.icon = ''
    form.parent_id = ''

    router.push('/dashboard/categories')
  } catch(error) {
      console.error('Error', error)
  }
}

</script>

<style>

</style>