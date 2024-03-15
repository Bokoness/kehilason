import { login, logout, checkAuth } from "@/api/authApi"
import { defineStore } from "pinia"

export const useAuthStore = defineStore("auth", {
  state: () => ({ user: undefined }),
  actions: {
    async loginUser(email, password, community) {
      const user = await login(email, password, community)

      this.user = user
    },

    async logoutUser() {
      await logout()

      this.user = undefined
    },

    async isAuthenticated() {
      if (this.user) return true

      const user = await checkAuth()

      if (!user) return false

      this.user = user

      return true
    },
  },
})
