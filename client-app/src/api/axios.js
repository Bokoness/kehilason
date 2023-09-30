import axios from "axios"

const MyAxios = axios.create({
  baseURL: "/api",
})

export default MyAxios
