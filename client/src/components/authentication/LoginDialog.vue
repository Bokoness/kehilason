<template>
  <VForm ref="form">
    <VTextField
      bg-color="white"
      label="אימייל"
      type="email"
      v-model="email"
      :rules="[validation.required, validation.email]"
    />

    <VTextField
      bg-color="white"
      label="סיסמה"
      v-model="password"
      :rules="[validation.required, validation.password]"
      :append-icon="displayPassword ? 'mdi-eye' : 'mdi-eye-off'"
      :type="displayPassword ? 'text' : 'password'"
      @click:append="displayPassword = !displayPassword"
    />

    <VSelect
      v-model="community"
      :items="communities"
      item-title="name"
      item-value="id"
      label="קהילה"
      persistent-hint
      single-line
      :rules="[validation.community]"
    />

    <div class="d-flex justify-center">
      <VBtn text="התחברות" flat @click="submit" />
    </div>
  </VForm>
</template>

<script setup>
import { onMounted, ref } from "vue"
import { getCommunities } from "@/api/communityApi"
import validation from "@/validationRules"
import { useAuthStore } from "@/store/auth"

const auth = useAuthStore()

let email = ref("")
let password = ref("")
let community = ref("")
let communities = ref([])
let form = ref(null)

let displayPassword = ref(false)

async function submit() {
  const validation = await form.value.validate()

  if (validation.valid) {
    try {
      await auth.loginUser(email.value, password.value, community.value)
      this.$router.push({ name: "home" })
    } catch (error) {
      console.error(error)
      //todo : do something with the error
    }
  }
}

onMounted(async () => {
  communities.value = await getCommunities()
})
</script>
