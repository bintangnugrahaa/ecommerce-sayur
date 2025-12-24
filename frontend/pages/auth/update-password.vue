<template>
    <section class="my-10">
        <div class="container">
            <div class="flex flex-wrap justify-center items-center gap-8 lg:gap-16">
                <div class="w-full md:w-1/3 xl:w-1/3 lg:order-1 order-2">
                    <!-- img -->
                    <img src="~/assets/images/svg-graphics/fp-g.svg" alt="" class="max-w-full h-auto" />
                </div>

                <div class="w-full md:w-1/2 lg:mx-1/6 xl:w-1/3 lg:order-2 order-1 flex flex-col gap-6">
                    <div class="flex flex-col gap-1">
                        <h1 class="text-xl">Update Your Password</h1>
                        <p>Welcome to FreshCart! Enter your email to get started.</p>
                    </div>
                    <form @submit.prevent="handleSubmit" class="needs-validation" novalidate>
                        <div v-if="showSuccessMessage" class="mb-4 p-4 text-sm text-green-800 rounded-lg bg-green-50" role="alert">
                            <span class="font-medium">Berhasil!</span> Akun anda telah berhasil dibuat. Silahkan login.
                        </div>

                        <!-- Alert Error -->
                        <div v-if="authStore.error" class="mb-4 p-4 text-sm text-red-800 rounded-lg bg-red-50" role="alert">
                            <span class="font-medium">Gagal!</span> {{ authStore.error }}
                        </div>
                        <div class="flex flex-wrap gap-4">
                            <div class="w-full">
                                <div class="password-field relative">
                                    <label for="formSignupPassword" class="invisible hidden">New Password</label>
                                    <div class="password-field relative">
                                        <input v-model="password_new" :type="showPassword ? 'text' : 'password'"
                                            class="form-control border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base fakePassword"
                                            id="formSignupPassword" placeholder="*****" required />
                                        <span @click="togglePassword">
                                            <i :class="showPassword ? 'ti ti-eye' : 'ti ti-eye-off'" class="passwordToggler"></i>
                                        </span>
                                        <div class="invalid-feedback">Please enter password.</div>
                                    </div>
                                </div>
                            </div>
                            <div class="w-full">
                                <div class="password-field relative">
                                    <label for="formSignupPasswordConfirm" class="invisible hidden">Password Confirmation</label>
                                    <div class="password-field relative">
                                        <input v-model="password_confirmation" :type="showPasswordConfirm ? 'text' : 'password'"
                                            class="form-control border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base fakePassword"
                                            id="formSignupPassword" placeholder="*****" required />
                                        <span @click="togglePasswordConfirm">
                                            <i :class="showPasswordConfirm ? 'ti ti-eye' : 'ti ti-eye-off'" class="passwordToggler"></i>
                                        </span>
                                        <div class="invalid-feedback">Please enter password.</div>
                                    </div>
                                </div>
                            </div>
                            <!-- btn -->
                            <div class="w-full grid">
                                <button :disabled="authStore.loading"  type="submit"
                                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300">
                                    <span v-if="authStore.loading">Loading...</span>
                                    <span v-else>Submit</span>
                                </button>
                            </div>

                            <!-- text -->
                        </div>
                        <div class="mt-2">
                            <p>
                                <small>
                                    By continuing, you agree to our
                                    <a href="#!">Terms of Service</a>
                                    &
                                    <a href="#!" class="text-green-600">Privacy Policy</a>
                                </small>
                            </p>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
import { ref } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useRouter, useRoute } from 'vue-router'

definePageMeta({
  layout: 'auth',
})

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const showPassword = ref(false)
const showPasswordConfirm = ref(false)
const showSuccessMessage = ref(false)

const password_new = ref('')
const password_confirmation = ref('')

const token = route.query.token

if (!token) {
  router.push('/auth/signin')
}

const handleSubmit = async () => {
  try {
    
    await authStore.updatePasswordNoAuth(password_new.value, password_confirmation.value, token)
    showSuccessMessage.value = true

    password_new = ''
    password_confirmation = ''

    setTimeout(() => {
      router.push('/auth/signin')
    }, 2000)
  } catch (error) {
    // Error sudah ditangani di store
    console.error('Update password error:', error)
  }
}

const togglePassword = () => {
  showPassword.value = !showPassword.value
}

const togglePasswordConfirm = () => {
  showPasswordConfirm.value = !showPasswordConfirm.value
}
</script>