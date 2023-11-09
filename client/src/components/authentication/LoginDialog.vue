<template>
  <VForm ref="form">
    <VTextField bg-color="white" label="אימייל" type="email" v-model="email"
                :rules="[validation.required,validation.email]"/>
    <VTextField bg-color="white" label="סיסמה" type="password" v-model="password"
                :rules="[validation.required,validation.password]"/>

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
      <VBtn text="התחברות" flat @click="submit"/>
    </div>
  </VForm>
</template>


<script setup>

import { onMounted, ref } from "vue"
import { login } from "@/api/authApi"
import { getCommunities } from "@/api/communityApi"
import validation from "@/validationRules"

let email = ref("")
let password = ref("")
let community = ref("")
let communities = ref([])
let form = ref(null)

async function submit() {
  if (form.value.validate()) {
    await login(email.value, password.value, community.value)
  }
}

onMounted(async () => {
  communities.value = await getCommunities()
})

</script>
