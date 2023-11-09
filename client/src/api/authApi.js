import api from "@/api/axios"

export const register = async (fullName, email, password, community) => {
  const { data } = await api.post("/auth/register", {
    fullName,
    email,
    password,
    community
  })

  return data
}

export const login = async (email, password, community) => {
  const { data } = await api.post("/auth/login", {
    email,
    password,
    community
  })

  return data
}

