<template>
    <nav class="navbar-vertical-nav hidden xl:block">
      <div class="navbar-vertical">
        <div class="px-4 py-5">
          <NuxtLink to="/" class="navbar-brand">
            <img src="~/assets/images/logo/freshcart-logo.svg" alt="FreshCart Logo" />
          </NuxtLink>
        </div>
        <div class="navbar-vertical-content flex-grow-1">
          <ul class="navbar-nav flex-col" id="sideNavbar">
            <li class="nav-item">
              <NuxtLink 
                class="nav-link" 
                :class="{ 'active': $route.path === '/dashboard' }"
                to="/dashboard"
              >
                <div class="flex items-center">
                  <span class="nav-link-icon">
                    <IconHome />
                  </span>
                  <span class="nav-link-text">Dashboard</span>
                </div>
              </NuxtLink>
            </li>
  
            <li class="nav-item mt-6 mb-3">
              <span class="nav-label">Store Managements</span>
            </li>
  
            <li v-for="(item, index) in storeManagementItems" :key="index" class="nav-item">
              <NuxtLink 
                class="nav-link" 
                :class="{ 'active': $route.path === item.path }"
                :to="item.path"
              >
                <div class="flex items-center">
                  <span class="nav-link-icon">
                    <component :is="item.icon" />
                  </span>
                  <span class="nav-link-text">{{ item.label }}</span>
                </div>
              </NuxtLink>
            </li>
  
            <li class="nav-item mt-6 mb-3">
              <span class="nav-label">Site Settings</span>
            </li>
  
            <li v-for="(item, index) in settingsItems" :key="index" class="nav-item">
              <NuxtLink 
                class="nav-link" 
                :class="{ 'active': $route.path === item.path }"
                :to="item.path"
              >
                <div class="flex items-center">
                  <span class="nav-link-icon">
                    <component :is="item.icon" />
                  </span>
                  <span class="nav-link-text">{{ item.label }}</span>
                </div>
              </NuxtLink>
            </li>
            <li  class="nav-item">
              <button type="button" class="nav-link gap-2" @click="handleLogout">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                      viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                      stroke-linecap="round" stroke-linejoin="round"
                      class="icon icon-tabler icons-tabler-outline icon-tabler-logout">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                      <path
                          d="M14 8v-2a2 2 0 0 0 -2 -2h-7a2 2 0 0 0 -2 2v12a2 2 0 0 0 2 2h7a2 2 0 0 0 2 -2v-2" />
                      <path d="M9 12h12l-3 -3" />
                      <path d="M18 15l3 -3" />
                  </svg>
                  Log out
              </button>
            </li>
          </ul>
        </div>
      </div>
    </nav>
</template>
  
<script setup>
import IconHome from './IconHome.vue';
import IconStore from './IconStore.vue';
import IconList from './IconList.vue';
import IconShoppingBag from './IconShoppingBag.vue';
import IconUsers from './IconUsers.vue';
import IconUserCog from './IconUserCog.vue';
import { useAuthStore } from '~/stores/auth'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter() // Assuming you have a router instance set up in your Vue app

const handleLogout = () => {
    try {
        authStore.logout()
        router.push('/auth/signin')
    } catch (error) {
        console.error('Error during logout:', error)
    }
}

const storeManagementItems = [
{
    label: 'Products',
    path: '/dashboard/products',
    icon: IconStore
},
{
    label: 'Categories',
    path: '/dashboard/categories',
    icon: IconList
},
{
    label: 'Orders',
    path: '/dashboard/orders',
    icon: IconShoppingBag
},
{
    label: 'Customers',
    path: '/dashboard/customers',
    icon: IconUsers
}
]

const settingsItems = [
{
    label: 'Role',
    path: '/dashboard/role',
    icon: IconUserCog
}
]
</script>