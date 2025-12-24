import { defineStore } from 'pinia'

export const useCustomerStore = defineStore('customers', {
    state: () => ({
        customers: [],
        loading: false,
        error: null,
        customer: {},
        pagination: {
            page: 1,
            total_count: 0,
            per_page: 10,
            total_pages: 0
        }
    }),
    actions: {
        async fetchCustomers(params = {}) {
            try {
                this.loading = true
                const { search, page, limit, orderBy } = params
          
                const query = new URLSearchParams()
                if (search) query.append('search', search)
                if (page) query.append('page', page)
                if (limit) query.append('limit', limit) 
                if (orderBy) query.append('orderBy', orderBy)

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/admin/customers?${query.toString()}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${useCookie('token').value}`
                    }
                })
                const result = await response.json()
                this.customers = result.data
                this.pagination = result.pagination

                return result
            } catch (error) {
                this.error = error.message
            } finally {
                this.loading = false
            }
        },

        async fetchCustomerByID(id) {
            try {
                this.loading = true

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/admin/customers/${id}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${useCookie('token').value}`
                    }
                })
                const result = await response.json()
                this.customer = result.data

                return result
            } catch (error) {
                this.error = error.message
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
                formData.append('photo', file)
        
                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/auth/profile/image-upload`, {
                    method: 'POST',
                    headers: {
                    'Authorization': `Bearer ${useCookie('token').value}`,
                    },
                    body: formData
                })
        
                if (!response.ok) {
                    const error = await response.json()
                    throw new Error(error.message || 'Gagal mengupload gambar')
                }
        
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

        async createCustomer(customerData) {
            try {
                this.loading = true
                this.error = null
    
        
                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/admin/customers`, {
                    method: 'POST',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(customerData)
                })
    
        
                const result = await response.json()
        
                return result
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async updateCustomer(customerData, id) {
            try {
                this.loading = true
                this.error = null
        
                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/admin/customers/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(customerData)
                })
    
        
                const result = await response.json()
        
                return result
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async deleteCustomer(id) {
            try {
                this.loading = true
                this.error = null
        
                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/admin/customers/${id}`, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    },
                })
    
        
                const result = await response.json()
        
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