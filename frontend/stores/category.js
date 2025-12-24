import { defineStore } from 'pinia'

export const useCategoryStore = defineStore('categories', {
    state: () => ({
        categories: [],
        category: {},
        loading: false,
        error: null,
        imageUrl: null,
        pagination: {
          page: 1,
          total_count: 0,
          per_page: 10,
          total_pages: 0
        }
    }),
    
    actions: {
      async fetchCategoriesHome() {
        try {
          this.loading = true
        
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/categories/home`, {
            headers: {
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.categories = result.data

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchCategoriesShop() {
        try {
          this.loading = true
        
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/categories/shop`, {
            headers: {
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.categories = result.data

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchCategoriesAdmin(params={}) {
        try {
          this.loading = true
        
          const { search, page, limit,  orderBy } = params
          
          const query = new URLSearchParams()
          if (search) query.append('search', search)
          if (page) query.append('page', page)
          if (limit) query.append('limit', limit) 
          if (orderBy) query.append('orderBy', orderBy)
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/categories`, {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${useCookie('token').value}`
            }
          })
  
          const result = await response.json()
          
          this.categories = result.data
          this.pagination = result.pagination

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchCategoryByIDAdmin(id) {
        try {
          this.loading = true
        
  
          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/categories/${id}`, {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${useCookie('token').value}`
            }
          })
  
          const result = await response.json()
          
          this.category = result.data

          return result
          
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async uploadImage(file) {
        try {
          this.loading = true
          this.error = null

          if (!file) {
            throw new Error('File tidak ditemukan')
          }

          const formData = new FormData()
          formData.append('image', file)

          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/image-upload`, {
            method: 'POST',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
            },
            body: formData
          })

          const result = await response.json()
          this.imageUrl = result.data.imageUrl

          return result
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async createCategory(formData) {
        try {
          this.loading = true
          this.error = null

          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/categories`, {
            method: 'POST',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
          })

          const result = await response.json()
          
          if (!response.ok) {
            throw new Error(result.message || 'Terjadi kesalahan saat membuat kategori')
          }

          return result
          
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async editCategory(formData, id) {
        try {
          this.loading = true
          this.error = null

          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/categories/${id}`, {
            method: 'PUT',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
          })

          const result = await response.json()
          
          if (!response.ok) {
            throw new Error(result.message || 'Terjadi kesalahan saat membuat kategori')
          }

          return result
          
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async deleteCategory(id) {
        try {
          this.loading = true
          this.error = null

          const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/admin/categories/${id}`, {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
          })

          const result = await response.json()
          
          if (!response.ok) {
            throw new Error(result.message || 'Terjadi kesalahan saat membuat kategori')
          }

          return result
          
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },
    },
})