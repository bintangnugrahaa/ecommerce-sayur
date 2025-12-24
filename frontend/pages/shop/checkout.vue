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
								<li class="inline-block text-gray-500 active" aria-current="page">Shop Checkout</li>
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

						<h1 class="text-xl">Checkout</h1>
					</div>
				</div>

				<div class="flex flex-wrap lg:flex-nowrap gap-10">
					<div class="lg:w-3/5 md:w-full">
						<!-- accordion -->
						<div class="accordion accordion-flush" id="accordionFlushExample">
							<!-- accordion item -->
							<div class="border-b border-gray-300 py-4">
									<div class="flex justify-between items-center">
										<button 
										@click="toggleSection('address')"
										class="flex flex-row gap-2 items-center text-gray-900 text-md font-bold"
										:aria-expanded="activeSection === 'address'"
										>
										<svg xmlns="http://www.w3.org/2000/svg"
											class="icon icon-tabler icon-tabler-map-pin inline-block text-gray-500"
											width="20" height="20" viewBox="0 0 24 24" stroke-width="1.5"
											stroke="currentColor" fill="none" stroke-linecap="round"
											stroke-linejoin="round">
											<path stroke="none" d="M0 0h24v24H0z" fill="none" />
											<path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" />
											<path d="M17.657 16.657l-4.243 4.243a2 2 0 0 1 -2.827 0l-4.244 -4.243a8 8 0 1 1 11.314 0z" />
										</svg>
										Delivery address
										</button>
									</div>
									<div v-show="activeSection === 'address'" class="my-6">
										<div class="md:flex gap-6">
										<div class="lg:w-1/2 w-full">
											<div class="card card-body flex-col gap-4">
											<div class="relative flex items-center gap-2">
												<input
												v-model="selectedAddress"
												class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded-full focus:ring-green-600 focus:outline-none focus:ring-2"
												type="radio" 
												value="home"
												id="homeRadio"
												/>
												<label class="text-gray-800 inline-block" for="homeRadio">Home</label>
											</div>
											<address class="not-italic">
												{{ formData.address }}
												<br>
												<abbr title="Phone">{{ formData.phone }}</abbr>
											</address>
											</div>
										</div>
										</div>
									</div>
								</div>

								<!-- Delivery Instructions Section -->
								<div class="border-b border-gray-300 py-4">
								<button 
									@click="toggleSection('instructions')"
									class="flex flex-row gap-2 items-center text-gray-900 text-md font-bold"
									:aria-expanded="activeSection === 'instructions'"
								>
									<svg xmlns="http://www.w3.org/2000/svg"
									class="icon icon-tabler icon-tabler-shopping-bag inline-block text-gray-500"
									width="20" height="20" viewBox="0 0 24 24" stroke-width="1.5"
									stroke="currentColor" fill="none" stroke-linecap="round"
									stroke-linejoin="round">
									<path stroke="none" d="M0 0h24v24H0z" fill="none" />
									<path d="M6.331 8h11.339a2 2 0 0 1 1.977 2.304l-1.255 8.152a3 3 0 0 1 -2.966 2.544h-6.852a3 3 0 0 1 -2.965 -2.544l-1.255 -8.152a2 2 0 0 1 1.977 -2.304z" />
									<path d="M9 11v-5a3 3 0 0 1 6 0v5" />
									</svg>
									Delivery Method
								</button>
								<div v-show="activeSection === 'instructions'" class="my-6">
									<!-- Delivery Method Options -->
									<div class="mb-4">
										<div class="flex gap-4">
											<div class="flex items-center">
												<input
												v-model="deliveryMethod"
												type="radio"
												value="Delivery"
												id="delivery"
												class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded-full focus:ring-green-600 focus:outline-none focus:ring-2"
											/>
											<label for="delivery" class="ml-2 text-gray-800">Delivery</label>
											</div>
											<div class="flex items-center">
												<input
												v-model="deliveryMethod"
												type="radio"
												value="Pickup"
												id="pickup"
												class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded-full focus:ring-green-600 focus:outline-none focus:ring-2"
											/>
											<label for="pickup" class="ml-2 text-gray-800">Pickup</label>
											</div>
										</div>
									</div>
									
									<textarea
									v-model="deliveryRemarks"
									class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
									rows="3"
									:placeholder="deliveryMethod === 'pickup' ? 'Tulis catatan untuk pickup' : 'Tulis instruksi pengiriman'"
									></textarea>
									<p class="block mt-1 small text-gray-500">
									{{ deliveryMethod === 'pickup' ? 'Tambahkan catatan untuk pengambilan pesanan' : 'Tambahkan instruksi untuk pengiriman pesanan Anda' }}
									</p>
									<div class="mt-5 flex justify-end gap-2">
									<button
										@click="prevSection"
										class="btn inline-flex items-center gap-x-2 bg-white text-gray-800 border-gray-300 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300"
									>
										Prev
									</button>
									<button
										@click="nextSection"
										class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300"
									>
										Next
									</button>
									</div>
								</div>
								</div>

								<!-- Payment Method Section -->
								<div class="py-4">
								<button 
									@click="toggleSection('payment')"
									class="flex flex-row gap-2 items-center text-gray-900 text-md font-bold"
									:aria-expanded="activeSection === 'payment'"
								>
									<svg xmlns="http://www.w3.org/2000/svg"
									class="icon icon-tabler icon-tabler-credit-card inline-block text-gray-500"
									width="20" height="20" viewBox="0 0 24 24" stroke-width="1.5"
									stroke="currentColor" fill="none" stroke-linecap="round"
									stroke-linejoin="round">
									<path stroke="none" d="M0 0h24v24H0z" fill="none" />
									<path d="M3 5m0 3a3 3 0 0 1 3 -3h12a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-12a3 3 0 0 1 -3 -3z" />
									<path d="M3 10l18 0" />
									<path d="M7 15l.01 0" />
									<path d="M11 15l2 0" />
									</svg>
									Payment Method
								</button>
								<div v-show="activeSection === 'payment'" class="mt-6 flex flex-col gap-4">
									<div class="card">
									<div class="flex-col flex p-6 gap-4">
										<div class="flex gap-3">
										<div class="relative flex mt-2">
											<input
											v-model="selectedPayment"
											class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded-full focus:ring-green-600 focus:outline-none focus:ring-2"
											type="radio"
											value="cod"
											id="cashonDelivery"
											/>
											<label class="text-gray-800 inline-block" for="cashonDelivery"></label>
										</div>
										<div class="flex flex-col gap-1">
											<h5 class="text-base">Cash on Delivery</h5>
											<p class="text-sm">Pay with cash when your order is delivered.</p>
										</div>
										</div>
									</div>
									</div>
									<div class="card">
									<div class="flex-col flex p-6 gap-4">
										<div class="flex gap-3">
										<div class="relative flex mt-2">
											<input
											v-model="selectedPayment"
											class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded-full focus:ring-green-600 focus:outline-none focus:ring-2"
											type="radio"
											value="midtrans"
											id="payoneer"
											/>
											<label class="text-gray-800 inline-block" for="payoneer"></label>
										</div>
										<div class="flex flex-col gap-1">
											<h5 class="text-base">Pay with Midtrans</h5>
											<p class="text-sm">You will be redirected to Midtrans website to complete your purchase securely.</p>
										</div>
										</div>
									</div>
									</div>
									<div class="flex justify-end">
									<button
										@click="prevSection"
										class="btn inline-flex items-center gap-x-2 bg-white text-gray-800 border-gray-300 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300"
									>
										Prev
									</button>
									<button
										@click="placeOrder"
										class="ml-3 btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-300"
									>
										Place Order
									</button>
									</div>
								</div>
								</div>
						</div>
					</div>

					<div class="w-full md:w-full lg:w-1/3 lg:ml-14">
						<div>
							<div class="card shadow-sm">
								<h5 class="px-6 py-4 border-b border-gray-300">Order Details</h5>
								<ul class="flex flex-col">
									<!-- list group item -->
									<li v-for="item in cartItems" class="py-3 px-6 border-b border-gray-300">
										<div class="flex flex-wrap items-center">
											<div class="w-1/5 md:w-1/5">
												<img :src="item.product_image" alt="Ecommerce"
													class="w-10" />
											</div>
											<div class="w-2/5 md:w-2/5 flex flex-col flex-wrap gap-1">
												<h6>{{item.product_name}}</h6>
												<span class="text-gray-700 text-sm">{{item.weight}} / {{ item.unit }}</span>
											</div>
											<div class="w-1/5 md:w-1/5 text-center text-gray-700">
												<span>{{ item.quantity }}</span>
											</div>
											<div class="w-1/5 text-center md:w-1/5">
												<span class="font-bold text-gray-800">Rp. {{ formatPrice(item.sale_price) }}</span>
											</div>
										</div>
									</li>
									<!-- list group item -->
									<li class="py-3 px-6">
										<div class="flex items-center justify-between font-bold text-gray-800">
											<div>Subtotal</div>
											<div>Rp. {{ formatPrice(subtotal) }}</div>
										</div>
									</li>
								</ul>
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
import { useOrderStore } from '~/stores/orders'
import { usePaymentStore } from '~/stores/payment'
const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()
const orderStore = useOrderStore()
const paymentStore = usePaymentStore()
const route = useRoute()
const cartItems = ref([])
const deliveryRemarks = ref('')
const deliveryMethod = ref('')

const sections = ['address', 'instructions', 'payment']
const activeSection = ref('address')
const selectedAddress = ref('home')
const selectedPayment = ref('')
const formData = ref({
    name: authStore.user?.name || '',
    email: authStore.user?.email || '',
    phone: authStore.user?.phone || '',
    address: authStore.user?.address || '',
    lat: authStore.user?.lat || null,
    lng: authStore.user?.lng || null,
    photo: authStore.user?.photo || ''
})

const toggleSection = (section) => {
  activeSection.value = section
}

const nextSection = () => {
  const currentIndex = sections.indexOf(activeSection.value)
  if (currentIndex < sections.length - 1) {
    activeSection.value = sections[currentIndex + 1]
  }
}

const prevSection = () => {
  const currentIndex = sections.indexOf(activeSection.value)
  if (currentIndex > 0) {
    activeSection.value = sections[currentIndex - 1]
  }
}

const placeOrder = async () => {
  try {
	let total_amount = subtotal.value
	if (deliveryMethod.value == 'Delivery') {
		total_amount = total_amount + 5000
	}

	const currentTime = new Date();
	const formattedDate = currentTime.toISOString().split('T')[0]
    const formattedTime = currentTime.toTimeString().split(' ')[0]

    const orderData = {
		buyer_id: authStore.user?.id,
		order_date: formattedDate,
		total_amount: total_amount,
		shipping_type: deliveryMethod.value,
		order_time: formattedTime,
		payment_type: selectedPayment.value,
		remarks: deliveryRemarks.value,
		order_details: cartItems.value.map(item => ({
			product_id: item.id,
			quantity: item.quantity
		}))
    }

	const response = await orderStore.createOrders(orderData, formData.value.lat, formData.value.lng)
    await cartStore.deleteAllCart()

    const paymentData = {
		order_id: response.data.order_id,
		remarks: deliveryRemarks.value,
		payment_method: selectedPayment.value,
		gross_amount: total_amount,
		user_id: authStore.user?.id,

    }

    const respPayment = await paymentStore.createPayment(paymentData)
	if (respPayment.data.payment_token == "") {
		router.push('/shop/success-payment')
		return
	}

	window.snap.pay(respPayment.data.payment_token, {
        onSuccess: function(result) {
          router.push('/shop/success-payment')
        },
        onPending: function(result) {
          alert('Pembayaran dalam proses')
        },
        onError: function(result) {
          alert('Pembayaran gagal')
        },
        onClose: function() {
          alert('Anda menutup popup tanpa menyelesaikan pembayaran')
        }
	})
  } catch (error) {
    console.error('Error placing order:', error)
  }
}

const mapsLink = computed(() => {
  return `https://www.google.com/maps?q=${formData.lat.value},${formData.lat.value}`
})

onMounted(async () => {
    if (!authStore.isAuthenticated) {
        router.push('/auth/signin')
        return
    }
    const response = await cartStore.fetchCarts()
    cartItems.value = response.data

	const profile = await authStore.getProfile()
	if (profile.data.address == "") {
		router.push('/account/setting')
        return
	}

    formData.value = {
		name: profile.data.name || '',
		email: profile.data.email || '',
		phone: profile.data.phone || '',
		address: profile.data.address || '',
		lat: profile.data.lat || null,
		lng: profile.data.lng || null,
		photo: profile.data.photo || ''
	}
})

const formatPrice = (price) => {
  return new Intl.NumberFormat('id-ID').format(price)
}

const subtotal = computed(() => {
  return cartItems.value.reduce((total, item) => total + (item.sale_price * item.quantity), 0)
})
</script>