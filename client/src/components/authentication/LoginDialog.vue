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
import { onMounted, ref, defineEmits } from "vue"
import { getCommunities } from "@/api/communityApi"
import validation from "@/validationRules"
import { useAuthStore } from "@/store/auth"
import { useRouter } from "vue-router"
import eventBus from "@/eventBus"

const router = useRouter()
const auth = useAuthStore()

let email = ref("")
let password = ref("")
let community = ref("")
let communities = ref([])
let form = ref(null)

let displayPassword = ref(false)

const emit = defineEmits(["closeDialog"])

async function submit() {
  const validation = await form.value.validate()

  if (validation.valid) {
    await auth.loginUser(email.value, password.value, community.value)

    eventBus.emit("snackbar", {
      message: "ההתחברות בוצעה בהצלחה",
      color: "green",
    })

    router.push({ name: "Dashboard" })

    emit("closeDialog")
  }
}

onMounted(async () => {
  communities.value = await getCommunities()
})
</script>
