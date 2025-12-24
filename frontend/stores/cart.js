import { defineStore } from 'pinia'

export const useCartStore = defineStore('carts', {
    state: () => ({
        carts: [],
        loading: false,
        error: null,
    }),

    actions: {
        async fetchCarts() {
            try {
                this.loading = true
                this.error = null
                const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/auth/cart`, {
                    method: 'GET',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    }
                })
                const result = await response.json()
                this.carts = result.data

                return result
            } catch (err) {
                this.error = err.message
                throw error
            } finally {
                this.loading = false
            }
        },

        async addToCart(productId, quantity) {
            try {
                this.loading = true
                this.error = null
                const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/auth/cart`, {
                    method: 'POST',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        product_id: productId,
                        quantity: quantity
                    })
                })
                const result = await response.json()

                await this.fetchCarts()
                
                return result
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async deleteCart(productId) {
            try {
                this.loading = true
                this.error = null
                const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/auth/cart?product_id=${productId}`, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    }
                })
                const result = await response.json()

                await this.fetchCarts()
                
                return result
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async deleteAllCart() {
            try {
                this.loading = true
                this.error = null
                const response = await fetch(`${useRuntimeConfig().public.apiProductBaseUrl}/auth/cart/all`, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    }
                })
                const result = await response.json()

                await this.fetchCarts()
                
                return result
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        }
    },

    getters: {
        totalItems: (state) =>  state.carts.length || 0
    },
})