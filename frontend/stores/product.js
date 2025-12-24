import { pagination } from '#build/ui'
import { defineStore } from 'pinia'

export const useProductStore = defineStore('products', {
    state: () => ({
        products: [],
        product: {},
        loading: false,
        error: null,
        pagination: {
          page: 1,
          total_count: 0,
          per_page: 10,
          total_pages: 0
        }
    }),
    
    actions: {
      async fetchProductsHome() {
        try {
          this.loading = true
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/products/home`, {
            headers: {
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.products = result.data

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchProductsShop(params = {}) {
        try {
          this.loading = true

          const { search, page, limit, price, orderBy, category } = params
          
          const query = new URLSearchParams()
          if (search) query.append('search', search)
          if (page) query.append('page', page)
          if (limit) query.append('limit', limit) 
          if (price) query.append('price', price)
          if (orderBy) query.append('orderBy', orderBy)
          if (category) query.append('category', category)
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/products/shop?${query.toString()}`, {
            headers: {
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.products = result.data
          this.pagination = result.pagination

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchProductDetailHome(productId) {
        try {
          this.loading = true
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/products/home/${productId}`, {
            headers: {
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.product = result.data

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchProductsAdmin(params = {}) {
        try {
          this.loading = true

          const { search, page, limit, status, orderBy, category } = params
          
          const query = new URLSearchParams()
          if (search) query.append('search', search)
          if (page) query.append('page', page)
          if (limit) query.append('limit', limit) 
          if (status) query.append('status', status)
          if (orderBy) query.append('orderBy', orderBy)
          if (category) query.append('category', category)
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/products?${query.toString()}`, {
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${useCookie('token').value}`
            }
          })
  
          const result = await response.json()
          
          this.products = result.data
          this.pagination = result.pagination

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async createProductAdmin(formData) {
        try {
          this.loading = true

  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/products`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${useCookie('token').value}`
            },
            body: JSON.stringify(formData)
          })
  
          const result = await response.json()

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchProductDetail(id) {
        try {
          this.loading = true
          this.error = null
          
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/products/${id}`, {
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`
            }
          })
      
          if (!response.ok) {
            const error = await response.json()
            throw new Error(error.message || 'Gagal mengambil detail produk')
          }
      
          return await response.json()
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async updateProduct(id, data) {
        try {
          this.loading = true
          this.error = null
          
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/products/${id}`, {
            method: 'PUT',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
          })
    
      
          return await response.json()
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async deleteProductAdmin(productId) {
        try {
          this.loading = true
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/products/${productId}`, {
            method: 'DELETE',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${useCookie('token').value}`
            }
          })
  
          const result = await response.json()

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },
    }
})