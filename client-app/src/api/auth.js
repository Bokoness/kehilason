import axios from "axios";
export default class AuthApi {
  static async register(fullName, email, password) {
    try {
      const { data } = await axios.post("api/auth/register", {
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
      const { data } = await axios.post("api/auth/login", { email, password });
      return data;
    } catch (e) {
      console.log(e);
    }
  }
}
