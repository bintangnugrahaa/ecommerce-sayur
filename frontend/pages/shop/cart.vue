<template>
    <main>
        <div class="mt-4">
			<div class="container">
				<div class="flex flex-wrap">
					<div class="w-full">
						<!-- breadcrumb -->
						<nav aria-label="breadcrumb">
							<ol class="flex flex-wrap">
								<li class="inline-block text-green-600 mr-2">
									<a href="/">
										Home
										<svg xmlns="http://www.w3.org/2000/svg"
											class="icon icon-tabler icon-tabler-chevron-right inline-block" width="14"
											height="14" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"
											fill="none" stroke-linecap="round" stroke-linejoin="round">
											<path stroke="none" d="M0 0h24v24H0z" fill="none" />
											<path d="M9 6l6 6l-6 6" />
										</svg>
									</a>
								</li>
								<li class="inline-block text-green-600 mr-2">
									<a href="../shop">
										Shop

										<svg xmlns="http://www.w3.org/2000/svg"
											class="icon icon-tabler icon-tabler-chevron-right inline-block" width="14"
											height="14" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"
											fill="none" stroke-linecap="round" stroke-linejoin="round">
											<path stroke="none" d="M0 0h24v24H0z" fill="none" />
											<path d="M9 6l6 6l-6 6" />
										</svg>
									</a>
								</li>
								<li class="inline-block text-gray-500 active" aria-current="page">Shop Cart</li>
							</ol>
						</nav>
					</div>
				</div>
			</div>
		</div>
        <div class="my-10">
			<div class="container">
				<div class="flex flex-wrap">
					<div class="w-full mb-6">
						<!-- card -->

						<h1 class="text-xl">Shop Cart</h1>
					</div>
				</div>

				<div class="flex flex-wrap lg:flex-nowrap lg:gap-x-12 gap-y-6">
					<div class="lg:w-2/3 w-full">
						<div class="flex flex-col gap-5">
							<ul class="list-none">
								<!-- list group -->
								<li v-for="item in cartItems" class="py-3 border-gray-300 border-t">
									<div class="flex items-center justify-between">
										<div class="w-1/2 md:w-1/2 lg:w-3/5">
											<div class="flex flex-row gap-5">
												<img :src="item.product_image" alt="Ecommerce"
													class="w-16 h-16" />
												<div class="flex flex-col gap-2">
													<div>
														<!-- title -->
														<NuxtLink :to="`../shop/${item.id}`" class="text-inherit">
															<h6>{{item.product_name}}</h6>
														</NuxtLink>
														<span class="text-gray-500 text-sm">{{ item.weight }} / {{item.unit}}</span>
													</div>
													<!-- text -->
													<div class="text-sm leading-none">
														<a href="#!" class="text-green-600 flex items-center gap-1">
															<span class="align-text-bottom">
																<svg xmlns="http://www.w3.org/2000/svg"
																	class="icon icon-tabler icon-tabler-trash"
																	width="14" height="14" viewBox="0 0 24 24"
																	stroke-width="2" stroke="currentColor" fill="none"
																	stroke-linecap="round" stroke-linejoin="round">
																	<path stroke="none" d="M0 0h24v24H0z" fill="none">
																	</path>
																	<path d="M4 7l16 0"></path>
																	<path d="M10 11l0 6"></path>
																	<path d="M14 11l0 6"></path>
																	<path
																		d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12">
																	</path>
																	<path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3">
																	</path>
																</svg>
															</span>
															<span @click="removeItem(item.id)" class="text-gray-500 text-sm">Remove</span>
														</a>
													</div>
												</div>
											</div>
										</div>
										<!-- input group -->
										<div class="w-1/3 md:w-1/4 lg:w-1/6">
											<!-- input -->
											<div
												class="input-group input-spinner rounded-lg flex justify-between items-center">
												<button
                                                    @click="decreaseQuantity(item)"
                                                    class="w-8 py-1 border-r cursor-pointer border-gray-300"
                                                >
                                                    -
                                                </button>
												<input v-model="item.quantity" type="number" step="1" max="10"
													class="quantity-field w-9 px-2 text-center h-7 border-0 bg-transparent" />
                                                    <button
                                                    @click="increaseQuantity(item)"
                                                    class="w-8 py-1 border-l cursor-pointer border-gray-300"
                                                >
                                                    +
                                                </button>
											</div>
										</div>
										<!-- price -->
										<div class="w-1/5 md:w-1/5 text-right">
											<span class="font-bold text-gray-800">Rp. {{ formatPrice(item.sale_price) }}</span>
										</div>
									</div>
								</li>
							</ul>
							<!-- btn -->
							<div class="flex justify-between">
								<a href="../shop"
									class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300">
									Continue Shopping
								</a>
								<a href="#!"
									class="btn inline-flex items-center gap-x-2 bg-gray-800 text-white border-gray-800 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-900 hover:border-gray-900 active:bg-gray-900 active:border-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300">
									Update Cart
								</a>
							</div>
						</div>
					</div>

					<!-- sidebar -->
					<div class="w-full lg:w-1/3 md:w-full">
						<!-- card -->
						<div class="relative card min-w-0">
							<div class="card-body flex flex-col gap-4">
								<!-- heading -->
								<h2 class="text-md">Summary</h2>
								<div
									class="relative flex flex-col min-w-0 rounded-lg break-words border bg-white border-1 border-gray-300">
									<!-- list group -->
									<ul class="flex flex-col">
										<li
											class="relative py-3 px-4 -mb-px border-r-0 border-l-0 border-gray-300 no-underline flex justify-between items-start">
											<div>
												<div class="font-bold text-gray-800">Subtotal</div>
											</div>
											<span class="font-bold text-gray-800">Rp. {{formatPrice(subtotal)}}</span>
										</li>
									</ul>
								</div>
								<div>
									<div class="grid">
										<!-- btn -->
										<a href="../shop/checkout"
											class="btn flex justify-between bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 btn-lg"
											type="submit">
											Go to Checkout
											<span class="font-bold">Rp. {{formatPrice(subtotal)}}</span>
										</a>
									</div>
									<!-- text -->
									<p class="mt-1">
										<span class="text-sm">
											By placing your order, you agree to be bound by the Freshcart
											<a href="#!" class="text-green-600">Terms of Service</a>
											and
											<a href="#!" class="text-green-600">Privacy Policy.</a>
										</span>
									</p>
								</div>

							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
    </main>
</template>

<script setup>
import { useCartStore } from '~/stores/cart'
import { onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '~/stores/auth'
const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()
const route = useRoute()
const cartItems = ref([])

onMounted(async () => {
    if (!authStore.isAuthenticated) {
        router.push('/auth/signin')
        return
    }
    const response = await cartStore.fetchCarts()
    cartItems.value = response.data
})

const subtotal = computed(() => {
  return cartItems.value.reduce((total, item) => total + (item.sale_price * item.quantity), 0)
})

const increaseQuantity = (item) => {
  if (item.quantity < 10) item.quantity++
}

const decreaseQuantity = (item) => {
  if (item.quantity > 1) item.quantity--
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('id-ID').format(price)
}

const removeItem = async (itemId) => {
  if (!authStore.token) {
    // Jika belum login, arahkan ke halaman login
    alert('Silakan login terlebih dahulu')
    router.push('/auth/signin')
    return
  }
  try {
    cartStore.deleteCart(itemId)
    router.replace(route.fullPath)
  } catch (error) {
    console.error('Error deleting item:', error)
    alert('Terjadi kesalahan saat menghapus ke keranjang')
  }
}

</script>