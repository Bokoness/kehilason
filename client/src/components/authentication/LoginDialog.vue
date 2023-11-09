<template>
  <VForm ref="form">
    <VTextField bg-color="white" label="אימייל" type="email" v-bind="email"
                :rules="[validation.required,validation.email]"/>
    <VTextField bg-color="white" label="סיסמה" type="password" v-bind="password"
                :rules="[validation.required,validation.password]"/>

    <VSelect
      v-bind="community"
      :items="communities"
      item-title="name"
      item-value="id"
      label="קהילה"
      persistent-hint
      return-object
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
    await login(email.value, password.value)
  }
}

onMounted(async () => {
  communities.value = await getCommunities()
})

</script>
