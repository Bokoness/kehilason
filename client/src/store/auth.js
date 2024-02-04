// Utilities
import { defineStore } from 'pinia'
import { errorHandler } from '@/handlers/errorHandler.js'
import router from '@/router'
import myAxios from '@/plugins/axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    userDetails: {},
    isLoggedIn: false,
    isSuperuser: false,
  }),
  getters: {
    isAdmin() {
      return this.userDetails.isSuperuser
    },
  },
  actions: {
    loginSuccess(user) {
      this.isLoggedIn = true
      this.userDetails = user
      this.isSuperuser = user.isSuperuser
    },
    clearData() {
      this.$reset()
    },
    async register(payload) {
      try {
        const { data } = await myAxios.post('/auth/register', payload)
        this.loginSuccess(data)
      } catch (e) {
        errorHandler(e, `${this.$id}: register`)
        throw new Error (e.message)
      }
    },
    async login(payload) {
      try {
        const { data } = await myAxios.post('/auth/login', payload)
        this.loginSuccess(data)
      } catch (e) {
        errorHandler(e, `${this.$id}: login`)
        throw new Error (e.message)
      }
    },
    async checkLogin() {
      try {
        const { data } = await myAxios.get('/auth/check-login')
        this.loginSuccess({ user: data })
      } catch (e) {
        errorHandler(e, `${this.$id}: checkLogin`)
      }
    },
    async logout() {
      try {
        this.clearData()
        await myAxios.post('/auth/logout')
        router.push({ name: 'Login' })
      } catch (e) {
        errorHandler(e, `${this.$id}: logout`)
      }
    },
  },
})
