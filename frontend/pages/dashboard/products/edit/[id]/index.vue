<template>
  <div class="container">
    <div class="grid grid-cols-1 mb-8">
      <!-- page header -->
      <div class="md:flex justify-between items-center">
        <div>
          <h2 class="text-xl">Edit Product</h2>
          <!-- breacrumb -->
          <nav aria-label="breadcrumb">
            <ol class="flex flex-wrap">
              <li class="inline-block text-green-600">
                <a href="/dashboard">
                  Dashboard
                  <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor"
                    stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-slash inline-block mx-2">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M17 5l-10 14" />
                  </svg>
                </a>
              </li>
              <li class="inline-block text-green-600">
                  <a href="/dashboard/products">
                      Products
                      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14"
                      viewBox="0 0 24 24" fill="none" stroke="currentColor"
                      stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                      class="icon icon-tabler icons-tabler-outline icon-tabler-slash inline-block mx-2">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                        <path d="M17 5l-10 14" />
                      </svg>
                  </a>
              </li>
              <li class="inline-block text-gray-500 active" aria-current="page">Edit Product</li>
            </ol>
          </nav>
        </div>
        <!-- button -->
        <div>
            <a href="/dashboard/products" class="btn inline-flex items-center gap-x-2 bg-gray-300 text-black border-gray-300 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-100">Back to Product</a>
        </div>
      </div>
    </div>
    <!-- row -->
    <div class="grid mb-5">
      <div class="card card-lg border-0 row-span-4 col-span-2 ">
        <div class="card-body flex flex-col gap-8 p-7">
          <div class="flex flex-col gap-4">
            <h3 class="mb-0 text-md">Product Information</h3>
            <form @submit.prevent="handleSubmit" class="grid grid-cols-12 gap-6 needs-validation" novalidate>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <label for="creatCustomerName"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Product Name
                    <span class="text-red-600">*</span>
                  </label>
                  <input
                    v-model="form.name"
                    type="text"
                    class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                    id="creatCustomerName" placeholder="Customer Name" required />
                  <div class="invalid-feedback">Please enter product name</div>
                </div>
              </div>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <label class="inline-block text-gray-800 font-medium mb-2" for="selectOne">Product Category
                      <span class="text-red-600">*</span>
                  </label>
                  <select v-model="form.category_slug" class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base" aria-label="Default select example" required>
                      <option selected>Open this select menu</option>
                      <option v-for="category in categoryDatas" :key="category.slug" :value="category.slug">{{ category.name }}</option>
                  </select>
                </div>
              </div>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <label for="creatCustomerName"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Unit
                    <span class="text-red-600">*</span>
                  </label>
                  <select required v-model="form.unit" class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base" aria-label="Default select example">
                    <option selected>Open this select unit</option>
                    <option value="gram">Gr</option>
                    <option value="kilo">Kg</option>
                  </select>
                </div>
              </div>
              <div class="lg:col-span-12 col-span-12">
                <div>
                  <label class="inline-block text-gray-800 font-medium mb-2">Product descriptions</label>
                  <textarea v-model="form.description" class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base" name="product_description" id="product_description" required></textarea>
                </div>
              </div>
              <div class="lg:col-span-6 col-span-12">
                <div>
                  <label for="creatCustomerName"
                    class="inline-block text-gray-800 font-medium mb-2">
                    Product Status
                    <span class="text-red-600">*</span>
                  </label>
                  <select v-model="form.status" class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base" aria-label="Default select example" required>
                    <option selected>Open this select status</option>
                    <option value="DRAFT">Draft</option>
                    <option value="ACTIVE">Active</option>
                    <option value="INACTIVE">In Active</option>
                  </select>
                </div>
              </div>

              <div class="col-span-12 mt-3">
                <div class="flex flex-col md:flex-row gap-2">
                  <button
                    :disabled="useProduct.loading"
                    class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100"
                    type="submit">
                    <span v-if="useProduct.loading">Loading...</span>
                    <span v-else>Update Product</span>
                  </button>
                  <a href="/dashboard/products"
                    class="btn inline-flex items-center gap-x-2 bg-gray-200 text-gray-800 border-gray-200 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300"
                    type="submit">
                    Cancel
                  </a>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="card h-full card-lg">
      <div class="p-6">
        <div class="flex justify-between flex-row items-center">
          <div>
            <h3 class=" text-md">Variant</h3>
          </div>
          <div>
            <button type="button"
              class="btn inline-flex items-center gap-x-2 bg-cyan-600 text-white border-cyan-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-cyan-700 hover:border-cyan-700 active:bg-cyan-700 active:border-cyan-700 focus:outline-none focus:ring-4 focus:ring-cyan-100"
              data-bs-toggle="modal" data-bs-target="#variantModal">
              Add Variant
              <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-plus"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 5l0 14" /><path d="M5 12l14 0" /></svg>
            </button>
          </div>
        </div>
      </div>
      <div class="card-body p-0">
        <div class="relative overflow-x-auto">
          <table class="text-left w-full whitespace-nowrap table-with-checkbox">
            <thead class="bg-gray-200 text-gray-700">
              <tr class="border-transparent !border-b-0">
                <th scope="col" class="px-6 py-3">
                  <div class="flex items-center">
                    <input
                      class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded focus:ring-green-600 focus:outline-none focus:ring-2"
                      type="checkbox" value="" id="checkAll" />
                    <label class="text-gray-800 ms-3"
                      for="checkAll"></label>
                  </div>
                </th>
                <th scope="col" class="px-6 py-3">Stock</th>
                <th scope="col" class="px-6 py-3">Weight</th>
                <th scope="col" class="px-6 py-3">Reguler Price</th>
                <th scope="col" class="px-6 py-3">Sale Price</th>
                <th scope="col" class="px-6 py-3">Product Image</th>

                <th scope="col" class="px-6 py-3">
                  Action
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="variant in variants" :key="variant.id" class="border-transparent !border-b-0">
                <td class="py-3 px-6 text-left">
                  <div class="flex items-center">
                    <input
                      class="w-4 h-4 text-green-600 bg-white border-gray-300 rounded focus:ring-green-600 focus:outline-none focus:ring-2"
                      type="checkbox"
                    />
                  </div>
                </td>
                <td class="py-3 px-6 text-left">{{ variant.stock }}</td>
                <td class="py-3 px-6 text-left">{{ variant.weight }}</td>
                <td class="py-3 px-6 text-left">Rp {{ variant.reguler_price }}</td>
                <td class="py-3 px-6 text-left">Rp {{ variant.sale_price }}</td>
                <td class="py-3 px-6 text-left">
                  <img :src="variant.product_image" class="h-12 w-12 object-cover rounded" alt="Variant Image">
                </td>
                <td class="py-3 px-6 text-left">
                  <div class="dropdown">
                    <a href="#" class="text-inherit" data-bs-toggle="dropdown" aria-expanded="false">
                      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                        class="icon icon-tabler icons-tabler-outline icon-tabler-dots-vertical">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                        <path d="M12 12m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
                        <path d="M12 19m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
                        <path d="M12 5m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
                      </svg>
                    </a>
                    <ul class="dropdown-menu">
                      <li>
                        <a class="dropdown-item" href="#" @click.prevent="deleteVariant(variant.id)">
                          <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" 
                            fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" 
                            stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-trash">
                            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                            <path d="M4 7l16 0" />
                            <path d="M10 11l0 6" />
                            <path d="M14 11l0 6" />
                            <path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
                            <path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
                          </svg>
                          Delete
                        </a>
                      </li>
                    </ul>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <div class="modal fade" id="variantModal" tabindex="-1" aria-labelledby="variantModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog-centered">
      <div class="modal-content p-6 flex flex-col gap-6">
        <div class="flex flex-row items-center justify-between">
          <h5 class="modal-title" id="addressLabel">Add Variant</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close">
            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-x text-gray-700"
              width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"
              fill="none" stroke-linecap="round" stroke-linejoin="round">
              <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
              <path d="M18 6l-12 12"></path>
              <path d="M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        <div class="modal-body p-0">
          <form @submit.prevent="addVariant" class="grid grid-cols-12 needs-validation gap-3" novalidate>
            <div class="lg:col-span-6 col-span-12">
              <label for="customerEditAdd" class="inline-block text-gray-800 font-medium mb-2">
                Stock <span class="text-red-600">*</span>
              </label>
              <input 
                v-model="variantForm.stock"
                type="number"
                class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="customerEditAdd" 
                placeholder="10" 
                required 
              />
              <div class="invalid-feedback">Please enter stock</div>
            </div>

            <div class="lg:col-span-6 col-span-12">
              <label for="customerZip" class="inline-block text-gray-800 font-medium mb-2">Weight</label>
              <input 
                v-model="variantForm.weight"
                type="number"
                class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="customerZip" 
                placeholder="100" 
                required 
              />
              <div class="invalid-feedback">Please enter weight</div>
            </div>

            <div class="lg:col-span-6 col-span-12">
              <label for="customerCity" class="inline-block text-gray-800 font-medium mb-2">Regular Price</label>
              <input 
                v-model="variantForm.reguler_price"
                type="number"
                class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="customerCity" 
                placeholder="10000" 
                required 
              />
              <div class="invalid-feedback">Please enter regular price</div>
            </div>

            <div class="lg:col-span-6 col-span-12">
              <label for="customerState" class="inline-block text-gray-800 font-medium mb-2">Sale Price</label>
              <input 
                v-model="variantForm.sale_price"
                type="number"
                class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                id="customerState" 
                placeholder="9000" 
                required 
              />
              <div class="invalid-feedback">Please enter sale price</div>
            </div>

            <div class="lg:col-span-8 col-span-12">
              <label class="inline-block text-gray-800 font-medium mb-2">Image Product</label>
              <input 
                type="file" 
                @change="handleFileUpload"
                class="border border-gray-300 text-gray-900 rounded-lg focus:shadow-[0_0_0_.25rem_rgba(10,173,10,.25)] focus:ring-green-600 focus:ring-0 focus:border-green-600 block p-2 px-3 disabled:opacity-50 disabled:pointer-events-none w-full text-base"
                accept="image/*"
              />
              <img v-if="variantForm.imageUrl" :src="variantForm.imageUrl" class="mt-2 h-20 w-20 object-cover rounded" />
            </div>
          </form>
        </div>
        <div class="flex flex-row gap-3">
          <button 
            type="button"
            @click="addVariant"
            class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100"
          >
            Add
          </button>
          <button 
            type="button"
            class="btn inline-flex items-center gap-x-2 bg-gray-200 text-gray-800 border-gray-200 border disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-gray-700 hover:border-gray-700 active:bg-gray-700 active:border-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-300"
            data-bs-dismiss="modal"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

import { useProductStore } from '~/stores/product'
import { ref, onMounted } from 'vue'
import { useCategoryStore } from '~/stores/category'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const useProduct = useProductStore()
const useCategory = useCategoryStore()
const categoryDatas = ref([])

const variants = ref([])
const variantForm = ref({
  stock: 0,
  weight: 0,
  reguler_price: 0,
  sale_price: 0,
  image: null,
  imageUrl: '',
  product_image: '',
})

const form = ref({
  name: '',
  category_slug: '',
  unit: 'gram',
  description: '',
  status: 'ACTIVE',
})

onMounted(async () => {
  try {
    // Fetch categories
    const respCat = await useCategory.fetchCategoriesAdmin()
    categoryDatas.value = respCat.data

    // Fetch product detail
    const productId = route.params.id
    const response = await useProduct.fetchProductDetail(productId)
    
    // Populate form with product data
    form.value = {
      name: response.data.product_name,
      category_slug: response.data.category_slug,
      unit: response.data.unit,
      description: response.data.product_description,
      status: response.data.product_status,
    }

    variantForm.value = {
      stock: response.data.stock,
      weight: response.data.weight,
      reguler_price: response.data.reguler_price,
      sale_price: response.data.sale_price,
      imageUrl: response.data.product_image,
      product_image: response.data.product_image,
    }

    variants.value.push({
      ...variantForm.value,
      id: Date.now()
    })

    if (response.data.child.length > 0) {
      response.data.child.forEach((child) => {
        variants.value.push({
          stock: child.stock,
          weight: child.weight,
          reguler_price: child.reguler_price,
          sale_price: child.sale_price,
          imageUrl: child.product_image,
          product_image: child.product_image,
          id: Date.now()
        })
      })   
    }

  } catch (error) {
    console.error('Error fetching data:', error)
    alert('Gagal memuat data produk')
  }
})

const addVariant = async () => {
    variants.value.push({
      ...variantForm.value,
      id: Date.now() // temporary id
    })
    
    // Reset form
    variantForm.value = {
      stock: 0,
      weight: 0,
      reguler_price: 0,
      sale_price: 0,
      image: null,
      imageUrl: '',
      product_image: '',
    }
    
    // Tutup modal
    // const modal = document.getElementById('variantModal')
    // const modalInstance = bootstrap.Modal.getInstance(modal)
    // modalInstance.hide()
}

const deleteVariant = (id) => {
  const isConfirmed = window.confirm('Apakah Anda yakin ingin menghapus variant ini?')
  if (isConfirmed) {
    variants.value = variants.value.filter(variant => variant.id !== id)
  }
}

const handleFileUpload = async (event) => {
  try {
    const file = event.target.files[0]
    if (file) {
      if (file.size > 1024 * 1024) {
        alert('Ukuran file maksimal 1MB')
        return
      }
      
      // Validasi tipe file
      if (!['image/jpeg', 'image/png', 'image/gif'].includes(file.type)) {
        alert('File harus berupa JPG, PNG, atau GIF')
        return
      }
      variantForm.value.image = file
      const resImage = await useCategory.uploadImage(file)
      variantForm.value.product_image = resImage.data.image_url
    }
  } catch (error) {
    console.log(error)
  }
}

const handleSubmit = async () => {
  try {
    const productData = {
      product_name: form.value.name,
      category_slug: form.value.category_slug,
      unit: form.value.unit,
      product_description: form.value.description,
      status: form.value.status,
      variant_detail: variants.value,
      variant: variants.value.length,
    }

    const productId = route.params.id
    await useProduct.updateProduct(productId, productData)
    router.push('/dashboard/products')
  } catch (error) {
    console.error('Error updating product:', error)
    alert(error.message || 'Gagal mengupdate produk')
  }
}
</script>