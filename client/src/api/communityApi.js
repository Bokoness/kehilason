import api from "@/api/axios"

export const getCommunities = async () => {
  const { data } = await api.get("/communities")

  return data
}
