<template>
  <VForm ref="form">
    <VTextField bg-color="white" label="אימייל" type="email" v-bind="email"
                :rules="[validation.required,validation.email]"/>
    <VTextField bg-color="white" label="סיסמה" type="password" v-bind="password"
                :rules="[validation.required,validation.password]"/>

    <VTextField v-model="fullName" bg-color="white" label="שם מלא" :rules="[validation.required]"/>

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
      <VBtn text="הרשמה" flat @click="submit"/>
    </div>
  </VForm>
</template>

<script setup>
import { onMounted, ref } from "vue"
import { register } from "@/api/authApi"
import { getCommunities } from "@/api/communityApi"
import validation from "@/validationRules"

let fullName = ref("")
let email = ref("")
let password = ref("")
let communities = ref([])
let community = ref("")
let form = ref(null)

async function submit() {
  if (form.value.validate()) {
    await register(fullName.value, email.value, password.value)
  }
}

onMounted(async () => {
  communities.value = await getCommunities()
})

</script>>



