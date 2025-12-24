<template>
  <div class="cart-sidebar fixed right-0 top-0 h-full bg-white shadow-lg w-96 z-50">
    <div class="h-full flex flex-col">
      <div class="border-b p-4 flex justify-between items-center">
        <h5>Keranjang Belanja</h5>
        <button type="button" class="text-gray-700" @click="$emit('close')">
          <Icon name="tabler:x" size="24" />
        </button>
      </div>

      <div class="flex-grow overflow-y-auto p-4">
        <div class="text-right mb-3">
          <NuxtLink to="/shop/cart" class="text-blue-600">Lihat Semua</NuxtLink>
        </div>

        <ul class="list-none">
          <li v-for="item in cartItems" :key="item.id" class="py-3 border-t">
            <div class="flex items-center">
              <div class="w-1/2 md:w-1/2 lg:w-3/5">
                <div class="flex">
                  <img :src="item.product_image" :alt="item.name" class="w-16 h-16 object-cover" />
                  <div class="ml-3">
                    <NuxtLink :to="`/product/${item.id}`" class="text-inherit">
                      <h6>{{ item.product_name }}</h6>
                    </NuxtLink>
                    <span><small class="text-gray-500">{{ item.weight }} {{ item.unit }}</small></span>
                    <div class="mt-2 small leading-none">
                      <button @click="removeItem(item.id)" class="text-green-600 flex items-center">
                        <Icon name="tabler:trash" size="14" class="mr-1" />
                        <span class="text-gray-500 text-sm">Hapus</span>
                      </button>
                    </div>
                  </div>
                </div>
              </div>

              <div class="w-1/3 md:w-1/4 lg:w-1/5">
                <div class="input-spinner rounded-lg flex justify-between items-center">
                  <button
                    @click="decreaseQuantity(item)"
                    class="w-8 py-1 border-r cursor-pointer border-gray-300"
                  >
                    -
                  </button>
                  <input
                    type="number"
                    v-model="item.quantity"
                    class="w-9 px-2 text-center h-7 border-0 bg-transparent"
                    min="1"
                    max="10"
                  />
                  <button
                    @click="increaseQuantity(item)"
                    class="w-8 py-1 border-l cursor-pointer border-gray-300"
                  >
                    +
                  </button>
                </div>
              </div>

              <div class="w-1/5 text-center">
                <span class="font-bold text-gray-800">Rp {{ formatPrice(item.sale_price) }}</span>
              </div>
            </div>
          </li>
        </ul>
      </div>

      <div class="border-t p-4">
        <div class="flex justify-between mb-3">
          <span>Subtotal:</span>
          <span class="font-bold">Rp {{ formatPrice(subtotal) }}</span>
        </div>
        <NuxtLink
          :to="`/shop/checkout`"
          class="btn w-full bg-green-600 text-white hover:bg-green-700 focus:ring-4 focus:ring-green-300"
        >
          Checkout
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '~/stores/cart'
import { onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '~/stores/auth'

const cartStore = useCartStore()
const cartItems = ref([])
const authStore = useAuthStore()
const router = useRouter()
const route = useRoute()

const formatPrice = (price) => {
  return new Intl.NumberFormat('id-ID').format(price)
}

onMounted(async () => {
  const response = await cartStore.fetchCarts()
  cartItems.value = response.data.slice(0, 5)
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

const removeItem = async (itemId) => {
  if (!authStore.token) {
    // Jika belum login, arahkan ke halaman login
    alert('Silakan login terlebih dahulu')
    return
  }
  try {
    cartStore.deleteCart(itemId)
    const response = await cartStore.fetchCarts()
    cartItems.value = response.data.slice(0, 5)
  } catch (error) {
    console.error('Error deleting item:', error)
    alert('Terjadi kesalahan saat menghapus ke keranjang')
  }
}
</script>