<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
      <!-- page header -->
      <div class="md:flex justify-between items-center">
        <div>
          <h2 class="text-xl">Create Customer</h2>
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

              <li class="inline-block text-gray-500 active" aria-current="page">Create
                Customer</li>
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
          <div v-if="customerStore.error" class="mb-4 p-4 text-sm text-red-800 rounded-lg bg-red-50 relative" role="alert">
            <div class="flex items-center pr-8">
                <span class="font-medium mr-2">Gagal!</span> {{ customerStore.error }}
            </div>
          </div>
          <div class="flex flex-col md:flex-row items-center mb-4 file-input-wrapper gap-2">
            <div>
              <img class="image h-16 w-16 rounded-lg" v-if="imageUrl" :src="imageUrl" alt="Image" />
            </div>

            <div
              class="file-upload btn inline-flex items-center gap-x-2 bg-gray-200 text-gray-800 border-gray-200 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300 md:ml-4">
              <input @change="handleFileUpload" type="file" class="file-input opacity-0" accept="image/*"/>
              Upload Photo
            </div>

            <span class="ms-2">JPG, GIF or PNG. 1MB Max.</span>
          </div>
          <div class="flex flex-col gap-4">
            <h3 class="mb-0 text-md">Customer Information</h3>
            <form @submit.prevent="handlerSubmit" class="grid grid-cols-12 gap-6 needs-validation" novalidate>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <!-- input -->
                  <label for="creatCustomerName"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Name
                    <span class="text-red-600">*</span>
                  </label>
                  <input type="text"
                    v-model="form.name"
                    class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                    id="creatCustomerName" placeholder="Customer Name" required />
                  <div class="invalid-feedback">Please enter customer name</div>
                </div>
              </div>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <!-- input -->
                  <label for="creatCustomerEmail"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Email
                    <span class="text-red-600">*</span>
                  </label>
                  <input
                     v-model="form.email"
                    type="email"
                    class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                    id="creatCustomerEmail" placeholder="Email Address" required />
                  <div class="invalid-feedback">Please enter email</div>
                </div>
              </div>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <!-- input -->
                  <label for="creatCustomerPhone"
                    class="inline-block text-gray-800 font-medium mb-2">Phone<span class="text-red-600">*</span></label>
                  <input
                     v-model="form.phone"
                    type="text"
                    class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                    id="creatCustomerPhone" placeholder="Number" required />
                  <div class="invalid-feedback">Please enter phone</div>
                </div>
              </div>

              <div class="lg:col-span-6 col-span-12">
                <div class="password-field relative">
                  <label class="inline-block text-gray-800 font-medium mb-2"
                    for="password">Password</label>
                  <div class="password-field relative">
                    <input
                      v-model="form.password"
                      :type="showPasswordConfirm ? 'text' : 'password'"
                      class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base flatpickr"
                      id="password" />
                      <span @click="togglePassword">
                        <i :class="showPassword ? 'ti ti-eye' : 'ti ti-eye-off'" class="passwordToggler"></i>
                      </span>
                      <div class="invalid-feedback">Please enter password</div>
                    </div>
                </div>
              </div>

              <div class="lg:col-span-6 col-span-12">
                <div class="password-field relative">
                  <label class="inline-block text-gray-800 font-medium mb-2"
                    for="password_confirmation">Password Confirmation</label>
                  <div class="password-field relative">
                    <input
                      v-model="form.password_confirmation"
                      :type="showPasswordConfirm ? 'text' : 'password'"
                      class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base flatpickr"
                      id="password_confirmation" />
                      <span @click="togglePasswordConfirmation">
                        <i :class="showPasswordConfirm ? 'ti ti-eye' : 'ti ti-eye-off'" class="passwordToggler"></i>
                      </span>
                      <div class="invalid-feedback">Please enter password confirmation</div>
                    </div>
                </div>
              </div>

              <div class="lg:col-span-6 col-span-12">
                <label class="inline-block text-gray-800 font-medium mb-2"
                  for="address">Address<span class="text-red-600">*</span></label>
                <input
                  v-model="form.address"
                  type="text"
                  class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base flatpickr"
                  id="address" required />
                <div class="invalid-feedback">Please enter address</div>
              </div>

              <div class="lg:col-span-6 col-span-12">
                <label class="inline-block text-gray-800 font-medium mb-2"
                  for="lat">Latitude<span class="text-red-600">*</span></label>
                <input
                  v-model="form.lat"
                  type="text"
                  class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base flatpickr"
                  id="lat" required />
                <div class="invalid-feedback">Please enter latitude</div>
              </div>

              <div class="lg:col-span-6 col-span-12">
                <label class="inline-block text-gray-800 font-medium mb-2"
                  for="lng">Longitude<span class="text-red-600">*</span></label>
                <input
                  v-model="form.lng"
                  type="text"
                  class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base flatpickr"
                  id="lng" required />
                <div class="invalid-feedback">Please enter longitude</div>
              </div>

              <div class="col-span-12 mt-3">
                <div class="flex flex-col md:flex-row gap-2">
                  <button
                    :disabled="customerStore.loading"
                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100"
                    type="submit">
                    <span v-if="customerStore.loading">Loading...</span>
                    <span v-else>Create New Customer</span>
                  </button>
                  <a href="/dashboard/customers"
                    class="btn inline-flex items-center gap-x-2 bg-gray-200 text-gray-800 border-gray-200 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300">
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
definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

import {useCustomerStore} from '~/stores/customer'
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const customerStore = useCustomerStore()
const router = useRouter()
const route = useRoute()
const file = ref(null)
const showPassword = ref(false)
const showPasswordConfirm = ref(false)
const imageUrl = ref('')
const customerId = route.params.id

const form = ref({
  name: '',
  email: '',
  phone: '',
  photo: '',
  password: '',
  address: '',
  lat: '',
  lng: '',
  role_id: 2,
  password_confirmation: '',
})
const password = ref('')
const password_confirmation = ref('')

onMounted(async () => {
  const response = await customerStore.fetchCustomerByID(customerId)
  imageUrl.value = response.data.photo
  password.value = response.data.password
  password_confirmation.value = response.data.password
  form.value = {
    name: response.data.name || '',
    email: response.data.email || '',
    phone: response.data.phone || '',
    address: response.data.address || '',
    lat: response.data.lat || '',
    lng: response.data.lng || '',
  }
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
      const result = await customerStore.uploadImage(selectedFile)
      
      // Update preview dan simpan URL
      imageUrl.value = result.data.image_url
      file.value = selectedFile
    }
  } catch (error) {
    console.error('Error upload:', error)
    alert(error.message || 'Gagal mengupload gambar')
  }
}


const handlerSubmit = async () => {
  try {
    if (form.value.password != "") {
      password.value = form.value.password
    }

    if (form.value.password_confirmation != "") {
      password_confirmation.value = form.value.password_confirmation
    }

    const customerData = {
      name: form.value.name,
      email: form.value.email,
      phone: form.value.phone,
      photo: imageUrl.value,
      password: password.value,
      address: form.value.address,
      lat: parseFloat(form.value.lat),
      lng: parseFloat(form.value.lng),
      password_confirmation: password_confirmation.value,
      role_id: 2
    }
    await customerStore.updateCustomer(customerData, customerId)

    router.push('/dashboard/customers')
  } catch (error) {
    console.error('Error creating customer:', error)
    alert(error.message || 'Gagal membuat customer')
  }
}


const togglePassword = () => {
  showPassword.value = !showPassword.value
}

const togglePasswordConfirmation = () => {
  showPasswordConfirm.value = !showPasswordConfirm.value
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