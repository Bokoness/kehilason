/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import { useAuthStore } from './store/auth'

const createVueApp = async () => {
  const app = createApp(App)

  registerPlugins(app)

  try {
    const { checkLogin } = useAuthStore()
    await checkLogin()
  } catch (error) {
    console.log(error)
  }

  app.mount('#app')
}
createVueApp()
