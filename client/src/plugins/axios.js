import axios from 'axios'

console.log(import.meta.env.VITE_APP_URL)
console.log(process.env.VITE_APP_URL)

const myAxios = axios.create({
  baseURL: import.meta.env.VITE_APP_URL + 'api',
  withCredentials: true,
})

myAxios.interceptors.request.use((config) => {
  config.headers['Authorization'] = `Bearer ${localStorage.getItem('authToken')}`
  return config
})
export default myAxios
