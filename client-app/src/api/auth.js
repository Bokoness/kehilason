import axios from "./axios"

export default class AuthApi {
  static async register(fullName, email, password) {
    const { data } = await axios.post("/auth/register", { fullName, email, password })
    return data
  }

  static async login(email, password) {
    const { data } = await axios.post("/auth/login", { email, password })
    return data
  }
}
