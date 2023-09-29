import axios from "axios";
export default class AuthApi {
  static async register(fullName, email, password) {
    try {
      const { data } = await axios.post("/auth/register", {
        fullName,
        email,
        password,
      });
      return data;
    } catch (e) {
      console.log(e);
    }
  }

  static async login(email, password) {
    try {
      const { data } = await axios.post("/auth/login", { email, password });
      return data;
    } catch (e) {
      console.log(e);
    }
  }
}
