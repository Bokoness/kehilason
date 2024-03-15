<template>
  <VAppBar flat color="secondaryBG" class="pl-15">
    <VRow align="center" justify="space-between" class="px-3">
      <VCol cols="3">
        <VImg src="icon.png" max-height="80" max-width="=250" />
      </VCol>

      <VCol lg="4" md="5 " cols="7">
        <VRow align="center">
          <VCol>קהילות</VCol>
          <VCol>אודות</VCol>
          <VCol>צור קשר</VCol>
          <VCol>
            <AuthDialog v-if="!auth.user" />
            <VBtn v-else text color="red" @click="logout"> התנתקות </VBtn>
          </VCol>
        </VRow>
      </VCol>
    </VRow>
  </VAppBar>
</template>

<script setup>
import AuthDialog from "@/components/authentication/AuthDialog.vue"
import { ref } from "vue"
import { useAuthStore } from "@/store/auth"
import { onMounted } from "vue"
import { useRouter } from "vue-router"

const auth = useAuthStore()
const router = useRouter()

const isAuthenticated = ref(false)

onMounted(async () => {
  isAuthenticated.value = await auth.isAuthenticated()
})

const logout = async () => {
  await auth.logoutUser()

  router.push({ name: "Home" })
}
</script>
