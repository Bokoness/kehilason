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

    <VTextField
      v-model="fullName"
      bg-color="white"
      label="שם מלא"
      :rules="[validation.required]"
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
      <VBtn text="הרשמה" flat @click="submit" />
    </div>
  </VForm>
</template>

<script setup>
import { onMounted, ref, defineEmits } from "vue"
import { register } from "@/api/authApi"
import { getCommunities } from "@/api/communityApi"
import validation from "@/validationRules"

import eventBus from "@/eventBus"

let fullName = ref("")
let email = ref("")
let password = ref("")
let communities = ref([])
let community = ref("")
let form = ref(null)

let displayPassword = ref(false)

const emit = defineEmits(["goLogin"])

async function submit() {
  const validation = await form.value.validate()

  if (validation.valid) {
    await register(fullName.value, email.value, password.value, community.value)

    eventBus.emit("snackbar", {
      message: "ההרשמה בוצעה בהצלחה",
      color: "green",
    })

    emit("goLogin")
  }
}

onMounted(async () => {
  communities.value = await getCommunities()
})
</script>
>
