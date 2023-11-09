import axios from "axios"

const api = axios.create({
  baseURL: "/api/communities"
})

export const getCommunities = async () => {
  const { data } = await api.get("/")
  return data
}
