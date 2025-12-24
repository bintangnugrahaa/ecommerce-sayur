<template>
    <main>
        <!-- section -->
        <section>
          <div class="container">
            <div class="grid grid-cols-12">
                <div class="col-span-12">
                    <div class="flex justify-between items-center md:hidden py-4">
                        <!-- heading -->
                        <h3 class="text-md mb-0">Account Setting</h3>
                        <!-- button -->
                        <button
                            class="btn inline-flex items-center gap-x-2 bg-transparent text-gray-600 border-gray-300 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300 md:hidden btn-sm ms-3"
                            type="button" data-bs-toggle="offcanvas" data-bs-target="#offcanvasAccount"
                            aria-controls="offcanvasAccount">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round"
                                class="icon icon-tabler icons-tabler-outline icon-tabler-indent-increase">
                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                <path d="M20 6l-11 0" />
                                <path d="M20 12l-7 0" />
                                <path d="M20 18l-11 0" />
                                <path d="M4 8l4 4l-4 4" />
                            </svg>
                        </button>
                    </div>
                </div>
                <div class="lg:col-span-2 md:col-span-4 col-span-12 border-r hidden md:block">
                    <div class="pt-10 md:pr-10">
                        <!-- nav -->
                        <ul class="nav flex-col nav-pills nav-pills-dark">
                            <!-- nav item -->
                            <li class="nav-item">
                                <a class="nav-link active gap-2" href="./account">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                                        viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round"
                                        class="icon icon-tabler icons-tabler-outline icon-tabler-shopping-bag">
                                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                        <path
                                            d="M6.331 8h11.339a2 2 0 0 1 1.977 2.304l-1.255 8.152a3 3 0 0 1 -2.966 2.544h-6.852a3 3 0 0 1 -2.965 -2.544l-1.255 -8.152a2 2 0 0 1 1.977 -2.304z" />
                                        <path d="M9 11v-5a3 3 0 0 1 6 0v5" />
                                    </svg>
                                    Your Orders
                                </a>
                            </li>
                            <!-- nav item -->
                            <li class="nav-item">
                                <a class="nav-link gap-2" aria-current="page" href="./account/setting">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                                        viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round"
                                        class="icon icon-tabler icons-tabler-outline icon-tabler-settings">
                                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                        <path
                                            d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c1 .608 2.296 .07 2.572 -1.065z" />
                                        <path d="M9 12a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" />
                                    </svg>
                                    Settings
                                </a>
                            </li>
                            <!-- nav item -->
                            <li class="nav-item">
                                <hr class="pt-3 mt-5" />
                            </li>
                            <!-- nav item -->
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
                <div class="lg:col-span-10 md:col-span-8 col-span-12">
                    <div class="py-6 md:p-6 lg:p-10">
							<!-- heading -->
                        <h2 class="mb-6 text-lg">Your Orders</h2>

                        <div class="block w-full overflow-auto scrolling-touch">
                            <!-- Table -->
                            <table class="w-full max-w-full bg-transparent whitespace-nowrap table-with-checkbox">
                                <!-- Table Head -->
                                <thead class="bg-gray-100">
                                    <tr class="border-gray-300 border-b">
                                        <th class="py-3 pl-6 text-left">&nbsp;</th>
                                        <th class="py-3 pl-6 text-left">Product</th>
                                        <th class="py-3 pl-6 text-left">Order</th>
                                        <th class="py-3 pl-6 text-left">Date</th>
                                        <th class="py-3 pl-6 text-left">Items</th>
                                        <th class="py-3 pl-6 text-left">Status</th>
                                        <th class="py-3 pl-6 text-left">Amount</th>

                                        <th></th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr v-if="response?.loading">
                                        <td colspan="8" class="text-center py-4">Loading...</td>
                                    </tr>
                                    <tr v-else-if="response?.error">
                                        <td colspan="8" class="text-center py-4 text-red-500">{{ orderStore.error }}</td>
                                    </tr>
                                    <!-- Table body -->
                                    <tr v-else-if="response?.orders.length === 0">
                                        <td colspan="8" class="text-center py-4">Tidak ada order</td>
                                    </tr>
                                    <tr v-else v-for="order in orderList" :key="order.id" class="border-gray-300 border-b">
                                        <td class="align-middle px-6 py-3">
                                            <a href="#" class="">
                                                <img
                                                    :src="order.product_image" :alt="order.order_code"
                                                    class="w-16 h-16 max-w-screen-lg" />
                                            </a>
                                        </td>
                                        <td class="align-middle px-6 py-3">
                                            <a href="#" class="font-semibold text-inherit">
                                                <h6 class="mb-0">{{ order.order_code }}</h6>
                                            </a>
                                            <span><small class="text-gray-500">{{ order.weight }} {{ order.unit }}</small></span>
                                        </td>
                                        <td class="align-middle px-6 py-3">
                                            <a href="#" class="text-inherit">{{ order.order_code }}</a>
                                        </td>
                                        <td class="align-middle px-6 py-3">{{ order.order_datetime }}</td>
                                        <td class="align-middle px-6 py-3">{{ order.quantity }}</td>
                                        <td class="align-middle px-6 py-3">
                                            <span :class="getStatusClass(order.status)">{{ order.status }}</span>
                                        </td>
                                        <td class="align-middle px-6 py-3">Rp {{ order.total_amount }}</td>
                                        <td class="text-gray-500 align-middle px-6 py-3">
                                            <a href="#" class="text-inherit" data-bs-toggle="tooltip"
                                                data-bs-placement="top" data-bs-title="View"><i
                                                    class="feather-icon icon-eye"></i></a>
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
          </div>
        </section>
    </main>

</template>
  
<script setup>
import { useAuthStore } from '~/stores/auth'
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'
import { useOrderStore } from '~/stores/orders'

definePageMeta({
  middleware: ['auth'],
})
const router = useRouter()
const authStore = useAuthStore()
const orderStore = useOrderStore()
const orderList = ref([])

const handleLogout = () => {
    try {
        authStore.logout()
        router.push('/auth/signin')
    } catch (error) {
        console.error('Error during logout:', error)
    }
}

const getStatusClass = (status) => {
  const classes = {
    'processing': 'inline-block p-1 text-sm align-baseline leading-none rounded bg-yellow-100 text-yellow-800 border border-yellow-200 font-semibold',
    'completed': 'inline-block p-1 text-sm align-baseline leading-none rounded bg-green-100 text-green-800 border border-green-200 font-semibold',
    'cancelled': 'inline-block p-1 text-sm align-baseline leading-none rounded bg-red-100 text-red-800 border border-red-200 font-semibold'
  }
  return classes[status.toLowerCase()] || classes['processing']
}

onMounted(async () => {
    const response = await orderStore.fetchOrders()
    orderList.value = response.data
})

</script>