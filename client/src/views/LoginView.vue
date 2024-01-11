<template>
  <div :class="isRegisterMode ? 'register' : 'login'">
    <v-container class="form-container justify-space-between flex-column" v-if="!loading">
      <h1 class="text-center primary--text">{{ isRegisterMode ? 'הרשמה' : 'כניסה' }}</h1>
      <form @submit.prevent="submit">
        <v-row justify="center" v-if="isRegisterMode">
          <v-col md="8" cols="10">
            <v-text-field
              v-model="form.fullName"
              :label="'full name'"
              :rules="[rules.usernameRequired]"
              required
            />
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col md="8" cols="10">
            <v-text-field
              v-model="form.email"
              :label="'email'"
              :rules="[rules.emailRequired]"
              required
            />
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col md="8" cols="10">
            <v-text-field
              v-model="form.password"
              :append-inner-icon="showPass ? 'mdi-eye' : 'mdi-eye-off'"
              :label="'סיסמה'"
              :type="showPass ? 'text' : 'password'"
              :rules="[rules.passRequired, rules.min]"
              @click:append-inner="showPass = !showPass"
              minlength="6"
              required
            />
          </v-col>
        </v-row>
        <v-row justify="center" v-if="isRegisterMode">
          <v-col md="8" cols="10">
            <v-text-field
              v-model="passwordValidation"
              :append-inner-icon="showPassAuth ? 'mdi-eye' : 'mdi-eye-off'"
              :label="'סיסמה'"
              :type="showPassAuth ? 'text' : 'password'"
              :rules="[rules.passValid, rules.min]"
              @click:append-inner="showPassAuth = !showPassAuth"
              minlength="6"
              required
            />
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col class="d-flex justify-center">
            <v-btn @click="submit" text color="primary" :disabled="!isValidForm">
              {{ isRegisterMode ? 'הרשמה' : 'כניסה' }}
            </v-btn>
            <v-btn
              variant="plain"
              :text="isRegisterMode ? 'רשום כבר? לחץ לכניסה' : 'לא רשום? עבור להרשמה'"
              @click="isRegisterMode = !isRegisterMode"
            ></v-btn>
          </v-col>
        </v-row>
      </form>
      <v-alert v-if="isAlertMessage" outlined type="error" class="d-flex justify-center mt-5">
        {{ error }}
      </v-alert>
    </v-container>
    <div class="progress text-center" v-else>
      <v-progress-circular :size="150" :width="2" color="primary" indeterminate />
    </div>
  </div>
</template>

<script>
import { useAuthStore } from '@/store/auth'

export default {
  name: 'CampaignManagerLogin',
  data: () => {
    return {
      form: {
        fullName: null,
        email: null,
        password: null,
      },
      isRegisterMode: false,
      passwordValidation: '',
      isAlertMessage: false,
      loading: false,
      showPass: false,
      showPassAuth: false,
    }
  },
  computed: {
    isValidForm() {
      let registerValid = true
      if (this.isRegisterMode) {
        registerValid = this.form.fullName && this.form.password === this.passwordValidation
      }
      return (
        registerValid && this.form.email && this.form.password && this.form.password.length >= 6
      )
    },
    rules() {
      return {
        usernameRequired: (v) => !!v || 'errors.auth.username',
        emailRequired: (v) => !!v || 'errors.auth.email',
        passRequired: (v) => !!v || 'errors.auth.pass',
        passValid: (v) => v === this.form.password || 'הסיסמאות לא תואמות',
        min: (v) => v.length >= 8 || 'מינימום 8 תווים',
      }
    },
  },
  methods: {
    async submit() {
      try {
        this.loading = true
        const { login, register } = useAuthStore()
        if (this.isRegisterMode) {
          await register(this.form)
        } else {
          await login(this.form)
        }

        this.$router.push({ name: 'Dashboard' })
      } catch (e) {
        console.log(e)
      } finally {
        this.loading = false
      }
    },
  },
}
</script>
<style lang="scss" scoped>
.login {
  margin: 10% 20% 0 20%;
  .progress {
    margin: 20% auto;
  }
}
.register {
  margin: 5% 20% 0 20%;
  .progress {
    margin: 20% auto;
  }
}
@media screen and (max-width: 600px) {
  .login {
    margin: 40% 10% 0 10%;
  }
  .register {
    margin: 20% 10% 0 10%;
  }
}
</style>
