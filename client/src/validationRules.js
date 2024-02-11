import * as yup from "yup"

export default {
  required: (v) => !!v || "יש למלא שדה זה",
  email: (v) => yup.string().email().isValidSync(v) || "אימייל לא תקין",
  password: (v) => yup.string().min(8).isValidSync(v) || "מינימום 8 תווים",
  community: (v) => !!v || "נא לבחור קהילה",
}
