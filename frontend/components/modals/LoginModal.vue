<template>
    <div class="modal-overlay" @click.self="$emit('close')">
      <div class="modal-content p-8 bg-white rounded-lg">
        <div class="flex justify-between items-center mb-4">
          <h3 class="font-bold text-gray-800">Sign Up</h3>
          <button type="button" class="btn-close" @click="$emit('close')">
            <Icon name="tabler:x" class="text-gray-700" size="24" />
          </button>
        </div>
  
        <div class="modal-body">
          <form @submit.prevent="handleSubmit">
            <div v-if="authStore.error" class="mb-4 p-4 text-sm text-red-800 rounded-lg bg-red-50 relative" role="alert">
              <button 
                  @click="authStore.error = null" 
                  class="absolute -top-2 -right-2 p-1.5 inline-flex items-center justify-center text-red-800 bg-red-50 hover:bg-red-200 rounded-full"
              >
                  <span class="sr-only">Tutup</span>
                  <i class="ti ti-x text-lg"></i>
              </button>
              <div class="flex items-center pr-8">
                  <span class="font-medium mr-2">Gagal!</span> {{ authStore.error }}
              </div>
          </div>
            <div class="mb-3">
              <label for="fullName" class="mb-2 block text-gray-800">Nama</label>
              <input
                v-model="form.fullName"
                type="text"
                class="form-control border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="fullName"
                placeholder="Masukkan Nama Anda"
                required
              />
            </div>
            
            <div class="mb-3">
              <label for="email" class="mb-2 block text-gray-800">Email</label>
              <input
                v-model="form.email"
                type="email"
                class="form-control border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="email"
                placeholder="Masukkan alamat email"
                required
              />
            </div>

            <div class="mb-5">
              <label for="password" class="mb-2 block text-gray-800">Password</label>
              <input
                v-model="form.password"
                type="password"
                class="form-control border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="password"
                placeholder="Masukkan Password"
                required
              />
            </div>
  
            <div class="mb-5">
              <label for="password" class="mb-2 block text-gray-800">Password Confirmation</label>
              <input
                v-model="form.password_confirmation"
                type="password"
                class="form-control border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="password"
                placeholder="Masukkan Password Konfirmasi"
                required
              />
              <span class="block mt-1 text-sm text-gray-500">
                Dengan mendaftar, Anda menyetujui
                <NuxtLink to="/terms" class="text-green-600">Syarat dan Ketentuan</NuxtLink>
                &
                <NuxtLink to="/privacy" class="text-green-600">Kebijakan Privasi</NuxtLink>
              </span>
            </div>
  
            <button
              :disabled="authStore.loading"
              type="submit"
              class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 justify-center"
            >
              <span v-if="authStore.loading">Loading...</span>
              <span v-else>Daftar</span>
            </button>
          </form>
        </div>
  
        <div class="modal-footer flex border-0 justify-center mt-3">
          Sudah punya akun?
          <NuxtLink to="/auth/signin" class="text-green-600 ml-1">Masuk</NuxtLink>
        </div>
      </div>
    </div>
  </template>
  
<script setup>
import { ref } from 'vue'
import { useAuthStore } from '~/stores/auth'


const handleLoginSuccess = async () => {
    // Proses item yang pending setelah login berhasil
    const pendingResult = processPendingCartItem()
    if (pendingResult) {
        alert(pendingResult.message)
    }
    
    // Tutup modal
    emit('close')
}

const authStore = useAuthStore()
const form = ref({
  fullName: '',
  email: '',
  password: '',
  password_confirmation: '',
})
  
const handleSubmit = async() => {
  // Implementasi logika pendaftaran di sini
  try {
    const userData = {
        name: form.value.fullName,
        email: form.value.email,
        password: form.value.password,
        password_confirmation: form.value.password_confirmation
    }
    
    await authStore.signup(userData)
    // Reset form setelah berhasil
    form.value = {
        fullName: '',
        email: '',
        password: '',
        password_confirmation: ''
    }
    
    // Tutup modal dan arahkan ke halaman signin
    $emit('close')
    navigateTo('/auth/signin')
    
} catch (error) {
    console.error('Signup error:', error)
    // Error sudah ditangani di store
}
}
</script>
  
<style scoped>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    }
  
    .modal-content {
        max-width: 500px;
        width: 90%;
    }
  </style>