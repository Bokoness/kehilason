import axios from "axios"
import eventBus from "@/eventBus"

const api = axios.create({
  baseURL: "/api",
})

api.interceptors.response.use(null, function (error) {
  if (error.response) {
    eventBus.emit("error", error.response?.data || "בעיה כללית")
  }

  return Promise.reject(error)
})

export default api
