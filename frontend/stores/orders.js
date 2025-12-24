import { defineStore } from 'pinia'

export const useOrderStore = defineStore('orders', {
    state: () => ({
        orders: [],
        loading: false,
        error: null,
        order: {},
        pagination: {
            page: 1,
            total_count: 0,
            per_page: 10,
            total_pages: 0
      }
    }),
    
    actions: {
      async fetchOrders(params = {}) {
        try {
          this.loading = true
          const { search, page, limit, status } = params
          
          const query = new URLSearchParams()
          if (search) query.append('search', search)
          if (page) query.append('page', page)
          if (limit) query.append('limit', limit) 
          if (status) query.append('status', status)
  
          const response = await fetch(`${useRuntimeConfig().public.apiOrderBaseUrl}/auth/orders?${query.toString()}`, {
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.orders = result.data
          this.pagination = result.pagination
          
          return result
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async fetchOrdersAdmin(params = {}) {
        try {
          this.loading = true
          const { search, page, limit, status } = params
          
          const query = new URLSearchParams()
          if (search) query.append('search', search)
          if (page) query.append('page', page)
          if (limit) query.append('limit', limit) 
          if (status) query.append('status', status)
  
          const response = await fetch(`${useRuntimeConfig().public.apiOrderBaseUrl}/admin/orders?${query.toString()}`, {
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            }
          })
  
          const result = await response.json()
          
          this.orders = result.data
          this.pagination = result.pagination

          return result
          
        } catch (err) {
          this.error = err.message
          throw error
        } finally {
          this.loading = false
        }
      },

      async createOrders(orderData, lat,lng) {
        try {
          this.loading = true
  
          const response = await fetch(`${useRuntimeConfig().public.apiOrderBaseUrl}/auth/orders?lat=${lat}&lng=${lng}`, {
            method: 'POST',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(orderData)
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

      async getDetailOrders(orderId, isAdmin = false) {
        try {
          this.loading = true
  
          let url = `${useRuntimeConfig().public.apiOrderBaseUrl}/auth/orders/${orderId}`
          if (isAdmin) {
            url = `${useRuntimeConfig().public.apiOrderBaseUrl}/admin/orders/${orderId}`
          }

          const response = await fetch(url, {
            method: 'GET',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
          })
  
          const result = await response.json()
          this.order = result.data
          
          return result
        } catch (err) {
          this.error = err.message
          throw err
        } finally {
          this.loading = false
        }
      },

      async updateStatusOrder(orderId, dataUpdate) {
        try {
          this.loading = true

          const response = await fetch(`${useRuntimeConfig().public.apiOrderBaseUrl}/admin/orders/${orderId}/status`,
          {
            method: 'PUT',
            headers: {
              'Authorization': `Bearer ${useCookie('token').value}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(dataUpdate)
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

    }
  })