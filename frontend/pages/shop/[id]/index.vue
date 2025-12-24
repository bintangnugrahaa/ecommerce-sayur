<template>
    <main>
        <div class="container">
			<div class="flex flex-wrap mt-4">
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
								<a href="shop">
									{{ featuredProduct.product_name }}

									<svg xmlns="http://www.w3.org/2000/svg"
										class="icon icon-tabler icon-tabler-chevron-right inline-block" width="14"
										height="14" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"
										fill="none" stroke-linecap="round" stroke-linejoin="round">
										<path stroke="none" d="M0 0h24v24H0z" fill="none" />
										<path d="M9 6l6 6l-6 6" />
									</svg>
								</a>
							</li>
							<li class="inline-block text-gray-500 active" aria-current="page">{{ featuredProduct.product_name }}</li>
						</ol>
					</nav>
				</div>
			</div>
		</div>

        <section class="my-10">
			<div class="container">
				<div class="flex flex-wrap">
					<ProductSingleGallery :product="featuredProduct" />
					<div class="lg:w-1/2 pr-4 pl-4">
						<div class="lg:pl-10 mt-6 md:mt-0">
							<div class="flex flex-col gap-4">
								<!-- content -->
								<a href="#!" class="block text-green-600">{{featuredProduct.category_name}}</a>
								<!-- heading -->
								<div class="flex flex-col">
									<h1>{{ featuredProduct.product_name }}</h1>
									<div class="flex flex-col gap-4">
										<div class="text-md">
											<span class="text-gray-900 font-semibold">
												Rp. {{ formatPrice(selectedChild ? selectedChild.reguler_price : featuredProduct.reguler_price) }}
											</span>
											<span v-if="selectedChild ? (selectedChild.sale_price !== selectedChild.reguler_price) : (featuredProduct.sale_price !== featuredProduct.reguler_price)" 
											class="line-through text-gray-500 ml-2">
												Rp. {{ formatPrice(selectedChild ? selectedChild.sale_price : featuredProduct.sale_price) }}
											</span>
										</div>
									</div>
								</div>
								<!-- hr -->
								<div class="flex flex-col gap-6">
									<hr />
									<div>
										<button 
											type="button"
											:class="[
											'btn inline-flex items-center gap-x-2 bg-white text-gray-800 border-gray-300 border mr-2',
											!selectedChild ? 'bg-green-600 text-white hover:text-white' : ''
											]"
											@click="handleVariantSelect(null)">
											{{ featuredProduct.weight }}{{ featuredProduct.unit }}
										</button>
										<button 
											v-for="child in featuredProduct.child" 
											:key="child.id"
											type="button"
											:class="[
											'btn inline-flex items-center gap-x-2 bg-white text-gray-800 border-gray-300 border mr-2',
											selectedChild?.id === child.id ? 'bg-green-600 text-white hover:text-white' : ''
											]"
											@click="handleVariantSelect(child)">
											{{ child.weight }}{{ featuredProduct.unit }}
										</button>
									</div>
									<div>
										<!-- input -->
										<div class="w-1/3 md:w-1/4 lg:w-1/5">
											<!-- input -->
											<div class="input-group input-spinner rounded-lg flex justify-between items-center">
												<button type="button" 
													@click="handleQuantityChange('decrement')"
													class="button-minus w-8 py-1 border-r cursor-pointer border-gray-300">
													-
												</button>
												<input 
													type="number" 
													v-model="quantity" 
													min="1"
													max="10"
													class="quantity-field w-9 px-2 text-center h-7 border-0 bg-transparent" 
												/>
												<button type="button"
													@click="handleQuantityChange('increment')"
													class="button-plus w-8 py-1 border-l cursor-pointer border-gray-300">
													+
												</button>
											</div>
										</div>
									</div>
									<div class="flex flex-wrap justify-start gap-2 items-center">
										<div class="lg:w-1/3 md:w-2/5 w-full grid">
											<!-- button -->
											<!-- btn -->
											<button
												@click="addToCart"
												type="button"
												class="btn gap-x-1 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 justify-center">
												<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18"
													viewBox="0 0 24 24" fill="none" stroke="currentColor"
													stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
													class="icon icon-tabler icons-tabler-outline icon-tabler-shopping-bag">
													<path stroke="none" d="M0 0h24v24H0z" fill="none" />
													<path
														d="M6.331 8h11.339a2 2 0 0 1 1.977 2.304l-1.255 8.152a3 3 0 0 1 -2.966 2.544h-6.852a3 3 0 0 1 -2.965 -2.544l-1.255 -8.152a2 2 0 0 1 1.977 -2.304z" />
													<path d="M9 11v-5a3 3 0 0 1 6 0v5" />
												</svg>
												Add to cart
											</button>
										</div>
										<div class="md:w-1/3 w-full">
										</div>
									</div>
									<!-- hr -->
									<hr />
								</div>
								<div>
									<!-- table -->
									<table class="text-left w-full">
										<tbody>
											<tr>
												<td class="px-6 py-3">Availability:</td>
												<td class="px-6 py-3"> {{ selectedChild ? selectedChild.stock : featuredProduct.stock }}</td>
											</tr>
											<tr>
												<td class="px-6 py-3">Type:</td>
												<td class="px-6 py-3">{{ featuredProduct.category_name }}</td>
											</tr>
										</tbody>
									</table>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>
        <section class="mb-10">
			<div class="container">
				<div class="flex flex-wrap">
					<div class="w-full">
						<ul class="nav nav-line-bottom border-b border-gray-300 pl-0" id="myTab" role="tablist">
							<!-- nav item -->
							<li class="nav-item" role="presentation">
								<!-- btn -->
								<button class="inline-block py-3 font-semibold px-4 no-underline nav-link active"
									id="product-tab" data-bs-toggle="tab" data-bs-target="#product-tab-pane"
									type="button" role="tab" aria-controls="product-tab-pane" aria-selected="true">
									Product Details
								</button>
							</li>
						</ul>
						<!-- tab content -->
						<div class="tab-content" id="myTabContent">
							<!-- tab pane -->
							<div class="tab-pane active opacity-100 block" id="product-tab-pane" role="tabpanel"
								aria-labelledby="product-tab" tabindex="0">
								<div class="my-8 flex flex-col gap-5">
									<div class="flex flex-col gap-1">
										<!-- text -->
										<h3 class="text-lg font-semibold">Nutrient Value &amp; Benefits</h3>
										<p>
											{{featuredProduct.description}}
										</p>
									</div>
									<!-- content -->
									<div class="flex flex-col gap-1">
										<h4 class="text-md font-semibold">Variant</h4>
										<!-- <p>{{featuredProduct.child + 1}} variant</p> -->
									</div>
									<div class="flex flex-col gap-1">
										<h4 class="text-md font-semibold">Disclaimer</h4>
										<p>Image shown is a representation and may slightly vary from the actual
											product. Every effort is made to maintain accuracy of all information
											displayed.</p>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>
    </main>

</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useRouter, useRoute } from 'vue-router'
import { useProductStore } from '~/stores/product'
import { useCartStore } from '~/stores/cart'
import ProductSingleGallery from '~/components/home/ProductSingleGallery.vue'

const featuredProduct = ref({})
const productStore = useProductStore()
const authStore = useAuthStore()
const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()
const productId = route.params.id

const selectedChild = ref(null)
const quantity = ref(1)

const handleQuantityChange = (type) => {
    if (type === 'increment' && quantity.value < 10) {
        quantity.value++
    } else if (type === 'decrement' && quantity.value > 1) {
        quantity.value--
    }
}

const handleVariantSelect = (child) => {
    if (child === null) {
        selectedChild.value = null
        return
    }
    selectedChild.value = child
}


onMounted(async () => {
	const response = await productStore.fetchProductDetailHome(productId)
  	featuredProduct.value = response.data
  	// Set default selected variant to first child if exists
  	selectedChild.value = null
})

const formatPrice = (price) => {
    return new Intl.NumberFormat('id-ID').format(price)
}

const addToCart =async () => {
    if (!authStore.token) {
      // Jika belum login, arahkan ke halaman login
      alert('Silakan login terlebih dahulu')
      router.push('/auth/signin')
      return
    }
    try {
      await cartStore.addToCart(parseInt(productId), quantity.value)
      alert('Produk berhasil ditambahkan ke keranjang')
    } catch (error) {
      console.error('Error adding to cart:', error)
      alert('Terjadi kesalahan saat menambahkan ke keranjang')
    }
}


useHead({
  script: [
    {
      src: '/js/vendors/tns-slider.js',
      body: true
    },
    {
      src: '/js/vendors/zoom.js',
      body: true
    },
    {
      src: '/js/theme.min.js',
      body: true
    },
    {
      src: '/js/vendors/countdown.js',
      body: true
    },
  ],
})

</script>