import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: null,
        loading: false,
        error: null,
        user: null
    }),

    actions: {
        async signin (email, password) {
            try {
                this.loading = true
                this.error = null

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/signin`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email, password })
                })

                const data = await response.json()
                if (!response.ok) {
                    throw new Error(data.message || 'Something went wrong!');
                }

                this.user = {
                    ...data.data,
                    role: data.data.role
                }
                this.token = data.data.access_token

                const tokenCookie = useCookie('token')
                tokenCookie.value = this.token

                const userCookie = useCookie('user')
                userCookie.value = JSON.stringify(this.user)
                // localStorage.setItem('token', this.token)
                // localStorage.setItem('user', JSON.stringify(this.user))

                return data
            } catch (error) {
                this.error = error.message
                throw error
            } finally {
                this.loading = false
            }
        },

        isSuperAdmin() {
            return this.user?.role === 'Super Admin'
        },
        
        logout() {
            this.user = null
            this.token = null
            const tokenCookie = useCookie('token')
            tokenCookie.value = null
            const userCookie = useCookie('user')
            userCookie.value = null
            // localStorage.removeItem('token')
        },

        checkAuth() {
            const tokenCookie = useCookie('token')
            const userCookie = useCookie('user')
            if (tokenCookie.value && userCookie.value) {
                this.token = tokenCookie.value
                this.user = userCookie.value 
            }
        },

        async signup(userData) {
            try {
                this.loading = true
                this.error = null
                
                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/signup`, {
                    method: 'POST',
                    headers: {
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(userData)
                })
                
                const data = await response.json()
                
                if (!response.ok) {
                    console.log(userData)
                    throw new Error(data.message || 'Terjadi kesalahan saat registrasi')
                }
                
                return data
            } catch (error) {
                this.error = error.message
                throw error
            } finally {
                this.loading = false
            }
        },

        async forgotPassword (email) {
            try {
                this.loading = true
                this.error = null

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/forgot-password`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email })
                })

                const data = await response.json()
                if (!response.ok) {
                    throw new Error(data.message || 'Something went wrong!');
                }

                return data
            } catch (error) {
                this.error = error.message
                throw error
            } finally {
                this.loading = false
            }
        },

        async verifyAccount ( token) {
            try {
                this.loading = true
                this.error = null

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/verify-account?token=${token}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                })

                const data = await response.json()
                if (!response.ok) {
                    throw new Error(data.message || 'Something went wrong!');
                }

                return data
            } catch (error) {
                this.error = error.message
                throw error
            } finally {
                this.loading = false
            }
        },
        
        async updatePasswordNoAuth (password_new, password_confirmation, token) {
            try {
                this.loading = true
                this.error = null

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/update-password?token=${token}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ password_new, password_confirmation  })
                })

                const data = await response.json()
                if (!response.ok) {
                    throw new Error(data.message || 'Something went wrong!');
                }

                return data
            } catch (error) {
                this.error = error.message
                throw error
            } finally {
                this.loading = false
            }
        },
        async getProfile () {
            try {
                this.loading = true
        
                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/auth/profile`, {
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    }
                })
        
                const result = await response.json()
                
                this.user = result.data

                return result
                
            } catch (err) {
                this.error = err.message
                throw error
            } finally {
                this.loading = false
            }
        },

        async updateProfile (userData) {
            try {
                this.loading = true
                this.error = null

                const response = await fetch(`${useRuntimeConfig().public.apiUserBaseUrl}/auth/profile`, {
                    method: 'PUT',
                    headers: {
                        'Authorization': `Bearer ${useCookie('token').value}`,
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(userData)
                })
                const data = await response.json()
                if (!response.ok) {
                    throw new Error(data.message || 'Something went wrong!');
                }
                return data
            } catch (error) {
                this.error = error.message
                throw error
            } finally {
                this.loading = false
            }
        }
    },

    getters: {
        isAuthenticated: (state) => !!state.token,
        getUser: (state) => state.user,
        userRole: (state) => state.user?.role
    }
})