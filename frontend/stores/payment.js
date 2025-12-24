import { defineStore } from 'pinia'

export const usePaymentStore = defineStore('payment', {
  state: () => ({
      token: null,
      loading: false,
      error: null,
      payment: null,
      midToken: null
  }),

  actions: {
    async createPayment(payload) {
      try {
        this.loading = true
        this.error = null

        const response = await fetch(`${useRuntimeConfig().public.apiPaymentBaseUrl}/auth/payments`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${useCookie('token').value}`
          },
          body: JSON.stringify(payload)
        })
        const data = await response.json()

        return data
      } catch (error) {
        this.error = error
      } finally {
        this.loading = false
      }
    }
  }
})