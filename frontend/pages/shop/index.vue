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
									<a href="shop">
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
								<li class="inline-block text-gray-500 active" aria-current="page">{{ selectedCategory?.name || 'All Products' }}</li>
							</ol>
						</nav>
					</div>
				</div>
			</div>
		</div>
        <div class="my-10">
			<div class="container">
				<div class="flex lg:gap-8">
                    <aside class="lg:w-1/4 mb-6 md:">
                        <div class="offcanvas offcanvas-left offcanvas-collapse" tabindex="-1" id="offcanvasCategory"
                        aria-labelledby="offcanvasCategoryLabel">
                            <div class="lg:invisible lg:hidden flex items-center p-4 justify-between">
                                <h5 class="offcanvas-title" id="offcanvasCategoryLabel">Filter</h5>
                                <button type="button" class="btn-close" @click="closeOffcanvas">
                                <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-x text-gray-700" width="24"
                                    height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                                    <path d="M18 6l-12 12"></path>
                                    <path d="M6 6l12 12"></path>
                                </svg>
                                </button>
                            </div>
                            <div class="offcanvas-body flex flex-col gap-8">
                                <div class="flex flex-col gap-3">
                                    <h5>Categories</h5>
                                    <ul class="flex flex-wrap nav-category">
                                        <CategoryMenuItem  v-for="(category, index) in categories" 
                                        :key="index"
                                        :category="category"
										:is-active="categoryFilter === category.slug"
  										@select="handleCategoryFilter" />
                                    </ul>
                                </div>
                                <div class="flex flex-col gap-3">
                                    <!-- price -->
                                    <h5>Price</h5>
                                    <div class="flex flex-col gap-3">
                                    <!-- range -->
                                        <div ref="priceRangeRef"></div>
                                        <div class="flex flex-row gap-2 items-center">
                                            <span class="text-gray-800">Price:</span>
                                            <span class="text-xs">{{ priceRangeValue }}</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </aside>
                    <section class="lg:w-3/4 w-full">
						<!-- card -->
						<div v-if="categoryFilter && selectedCategory" class="relative flex flex-col min-w-0 rounded-lg break-words bg-gray-100 p-8 mb-6">
							<!-- card body -->
							<div class="flex-auto">
								<h1 class="text-xl">{{ selectedCategory?.name || 'All Products' }}</h1>
							</div>
						</div>
						<!-- list icon -->
                        <div v-if="paginateProds.total_count > 0" class="flex flex-col md:flex-row justify-between lg:items-center mb-6 gap-3">
							<div>
								<p>
									<span class="text-gray-900">{{ paginateProds.total_count }}</span>
									Products found
								</p>
							</div>
                            <div class="flex flex-col md:flex-row justify-between md:items-center gap-3">
								<div class="flex items-center justify-between">
									<div class="ml-3 lg:hidden">
										<a class="btn inline-flex items-center gap-x-2 bg-white text-gray-800 border-gray-300 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300"
											data-bs-toggle="offcanvas" href="#offcanvasCategory" role="button"
											aria-controls="offcanvasCategory">
											<svg xmlns="http://www.w3.org/2000/svg"
												class="icon icon-tabler icon-tabler-filter inline-block" width="16"
												height="16" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"
												fill="none" stroke-linecap="round" stroke-linejoin="round">
												<path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
												<path
													d="M4 4h16v2.172a2 2 0 0 1 -.586 1.414l-4.414 4.414v7l-6 2v-8.5l-4.48 -4.928a2 2 0 0 1 -.52 -1.345v-2.227z">
												</path>
											</svg>
											Filters
										</a>
									</div>
								</div>

								<div class="flex gap-3">
									<div class="flex-grow-1">
										<!-- select option -->
										<select
											v-model="limit"
                							@change="handleLimitChange"
											class="text-base py-2 block w-full border-gray-300 rounded-lg focus:border-green-600 focus:ring-green-600 disabled:opacity-50 disabled:pointer-events-none">
											<option selected="">Show: 50</option>
											<option value="10">10</option>
											<option value="20">20</option>
											<option value="30">30</option>
										</select>
									</div>
									<div>
										<!-- select option -->
										<select
											v-model="sort"
											@change="handleSortChange"
											class="text-base py-2 block w-full border-gray-300 rounded-lg focus:border-green-600 focus:ring-green-600 disabled:opacity-50 disabled:pointer-events-none">
											<option selected="">Sort by: Featured</option>
											<option value="price_asc">Price: Low to High</option>
											<option value="price_desc">Price: High to Low</option>
											<option value="newest">Release Date</option>
										</select>
									</div>
								</div>
							</div>
                        </div>
                        <div class="grid lg:grid-cols-4 md:grid-cols-3 gap-4">
							<div v-for="product in featuredProducts" :key="product.id" class="relative rounded-lg break-words border bg-white border-gray-300 card-product">
                                <div  class="flex-auto p-4">
									<div class="text-center relative flex justify-center">
										<div class="absolute top-0 left-0">
											<span
												class="inline-block p-1 text-center font-semibold text-sm align-baseline leading-none rounded bg-red-600 text-white">Sale</span>
										</div>
										<NuxtLink :to="`/shop/${product.id}`">
											<img :src="product.product_image"
												alt="Grocery Ecommerce Template" class="w-full h-auto" />
										</NuxtLink>
									</div>
									<div class="flex flex-col gap-3">
										<a href="#!" class="text-decoration-none text-gray-500">
											<small>{{ product.category_name }}</small></a>
										<div class="flex flex-col gap-2">
											<h3 class="text-base truncate">
                                                <NuxtLink :to="`/shop/${product.id}`">
                                                    {{ product.product_name }}
												</NuxtLink>
                                            </h3>
										</div>
										<div class="flex justify-between items-center">
											<div>
												<span class="text-gray-900 font-semibold">Rp. {{ formatPrice(product.reguler_price) }}</span>
												<span v-if="product.sale_price != product.reguler_price" class="line-through text-gray-500">Rp. {{ formatPrice(product.sale_price) }}</span>
											</div>
											<div>
												<button @click="addToCart(product.id)" type="button"
													class="btn inline-flex items-center gap-x-1 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300 btn-sm">
													<svg xmlns="http://www.w3.org/2000/svg"
														class="icon icon-tabler icon-tabler-plus" width="14" height="14"
														viewBox="0 0 24 24" stroke-width="3" stroke="currentColor"
														fill="none" stroke-linecap="round" stroke-linejoin="round">
														<path stroke="none" d="M0 0h24v24H0z" fill="none" />
														<path d="M12 5l0 14" />
														<path d="M5 12l14 0" />
													</svg>
													<span>Add</span>
												</button>
											</div>
										</div>
									</div>
								</div>
                            </div>
                        </div>
                        <div class="flex my-8">
							<nav class="flex items-center gap-x-1">
								<button type="button"
									:disabled="paginateProds.page === 1"
									@click="handlePageChange(paginateProds.page - 1)"
									class="leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300 disabled:opacity-50 disabled:pointer-events-none">
									<svg xmlns="http://www.w3.org/2000/svg"
										class="icon icon-tabler icon-tabler-chevron-left" width="16" height="16"
										viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none"
										stroke-linecap="round" stroke-linejoin="round">
										<path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
										<path d="M15 6l-6 6l6 6"></path>
									</svg>
								</button>
								<div class="flex items-center gap-x-1">
									<button v-for="page in paginateProds.total_page" 
                                        :key="page"
                                        @click="handlePageChange(page)"
										:class="[
                                            'leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border',
                                            page === paginateProds.page 
                                                ? 'text-white border bg-green-600 border-green-600 hover:bg-green-600 focus:outline-none focus:bg-green-600'
                                                : 'bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300'
                                        ]">
										{{ page }}
									</button>
								</div>
								<button type="button"
									:disabled="paginateProds.page === paginateProds.total_page"
									@click="handlePageChange(paginateProds.page + 1)"
									class="leading-none min-h-[36px] min-w-[36px] py-2 px-2.5 inline-flex justify-center items-center gap-x-1.5 rounded-md border bg-white border-gray-300 text-gray-800 hover:bg-gray-300 focus:outline-none focus:bg-gray-300 disabled:opacity-50 disabled:pointer-events-none">
									<svg xmlns="http://www.w3.org/2000/svg"
										class="icon icon-tabler icon-tabler-chevron-right" width="16" height="16"
										viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none"
										stroke-linecap="round" stroke-linejoin="round">
										<path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
										<path d="M9 6l6 6l-6 6"></path>
									</svg>
								</button>
							</nav>
						</div>
                    </section>
                </div>
            </div>
        </div>
    </main>
	<!-- <QuickView v-if="show" @close="showLoginModal = false" /> -->
	<LoginModal v-if="showLoginModal" @close="showLoginModal = false" />
</template>

<script setup>
import CategoryMenuItem from '~/components/home/CategoryMenuItem.vue'
import { ref, onMounted } from 'vue'
import noUiSlider from 'nouislider'
import wNumb from 'wnumb'
import QuickView from '~/components/modals/QuickView.vue'
import { useCategoryStore } from '~/stores/category'
import { useProductStore } from '~/stores/product'
import { useAuthStore } from '~/stores/auth'
import { useRouter, useRoute } from 'vue-router'
import LoginModal from '~/components/modals/LoginModal.vue'
import { useCartStore } from '~/stores/cart'

const categories = ref([])
const categoryStore = useCategoryStore()
const featuredProducts = ref([])
const productStore = useProductStore()
const authStore = useAuthStore()
const cartStore = useCartStore()
const showLoginModal = ref(false)

const priceRangeRef = ref(null)
const priceRangeValue = ref('')
const showQickViewModal = ref(false)
const route = useRoute()
const paginateProds = ref({})
const limit = ref(10)
const sort = ref('')
const router = useRouter()

const search = route.query.search ?? ''
const categoryFilter = route.query.category?? ''

const showQuickView = () => {
  showQuickViewModal.value = true
}

const selectedCategory = ref(null)

watch(() => route.query.search, async (newSearch) => {
    const responseProds = await productStore.fetchProductsShop({
			search: newSearch,
			category: categoryFilter,
			page: 1,
			limit: limit.value,
			sort: sort.value,
		})
			
	featuredProducts.value = responseProds.data
	paginateProds.value = responseProds.pagination
}, { immediate: true })


onMounted(async () => {
	const slider = await noUiSlider.create(priceRangeRef.value, {
			start: [0, 100000],
			connect: true,
			range: {
				'min': 0,
				'max': 2000000
			},
			format: wNumb({
			decimals: 0,
			thousand: '.',
			prefix: 'Rp. '
		})
	})

	slider.on('update', (values) => {
		priceRangeValue.value = values.join(' - ')
	})

	slider.on('change', async (values) => {
		const min = parseInt(values[0].replace('Rp. ', '').replace(/\./g, ''))
    	const max = parseInt(values[1].replace('Rp. ', '').replace(/\./g, ''))

		const responseProds = await productStore.fetchProductsShop({
			search: search,
			category: categoryFilter,
			page: 1,
			limit: limit.value,
			sort: sort.value,
			price: `${min} - ${max}`,
		})
			
		featuredProducts.value = responseProds.data
		paginateProds.value = responseProds.pagination
	})

	const responseCat = await categoryStore.fetchCategoriesHome()

	categories.value = responseCat.data

  	const responseProds = await productStore.fetchProductsShop({
		orderBy: sort.value
	})
	featuredProducts.value = responseProds.data
	paginateProds.value = responseProds.pagination
	if (categoryFilter) {
		const category = categories.value.find(cat => cat.slug === categoryFilter)
		if (category) {
		selectedCategory.value = category
		}
	}
})

const handleAddToCart = async () => {

}

const handlePageChange = async (page) => {
	const responseProds = await productStore.fetchProductsShop({
		search: search,
		page: page,
		category: selectedCategory.value?.slug || ''
	})
  
  featuredProducts.value = responseProds.data
  paginateProds.value = responseProds.pagination
}

const handleLimitChange = async () => {
    const responseProds = await productStore.fetchProductsShop({
        search: search,
        page: 1,
        limit: limit.value,
        orderBy: sort.value,
        category: selectedCategory.value?.slug || ''
    })
    
    featuredProducts.value = responseProds.data
    paginateProds.value = responseProds.pagination
}

const handleSortChange = async () => {
    const responseProds = await productStore.fetchProductsShop({
        search: search,
        page: paginateProds.value?.page,
        limit: limit.value,
        orderBy: sort.value,
		category: selectedCategory.value?.slug || ''
    })
    
    featuredProducts.value = responseProds.data
    paginateProds.value = responseProds.pagination
}

const handleCategoryFilter = async (categoryId) => {
	selectedCategory.value = categories.value.find(cat => cat.slug === categoryId)

	const responseProds = await productStore.fetchProductsShop({
		category: categoryId,
		// page: 1,
		// limit: limit.value,
	})
	
	featuredProducts.value = responseProds.data
	paginateProds.value = responseProds.pagination
}

const addToCart =async (productId) => {
    if (!authStore.token) {
      // Jika belum login, arahkan ke halaman login
      alert('Silakan login terlebih dahulu')
      router.push('/auth/signin')
      return
    }
    try {
      await cartStore.addToCart(productId, 1)
      alert('Produk berhasil ditambahkan ke keranjang')
    } catch (error) {
      console.error('Error adding to cart:', error)
      alert('Terjadi kesalahan saat menambahkan ke keranjang')
    }
}


useHead({
  script: [
  	{
		src: '/libs/bootstrap/dist/js/bootstrap.bundle.min.js',
		body: true
    },
  	{
		src: '/libs/nouislider/dist/nouislider.min.js',
		body: true
    },
	{
		src: '/libs/wnumb/wNumb.min.js',
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
	{
		src: '/libs/simplebar/dist/simplebar.min.js',
		body: true
    },
  ],
})

const formatPrice = (price) => {
    return new Intl.NumberFormat('id-ID').format(price)
}

const closeOffcanvas = () => {
  // implementasi untuk menutup offcanvas
}
</script>

<style scoped>
.card-product:hover .card-product-action {
  opacity: 1 !important;
  visibility: visible !important;
  transition: all 0.3s ease;
}

.card-product .card-product-action {
  transition: all 0.3s ease;
}
</style>