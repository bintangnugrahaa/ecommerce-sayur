export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStore()
  
  await authStore.checkAuth()

  // Cek role user
  if (!authStore.user?.role || authStore.user.role !== 'Super Admin') {
    return navigateTo('/')
  }
})