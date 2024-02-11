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
          <VCol v-if="!isAuthenticated">
            <VBtn color="indigo" @click="openDialog">
              <VIcon>mdi-account</VIcon>
              <span>התחברו</span>
            </VBtn>
          </VCol>
        </VRow>
      </VCol>
    </VRow>
    <VDialog v-model="dialogModel" max-width="600" persistent>
      <AuthDialog @close-dialog="closeDialog" />
    </VDialog>
  </VAppBar>
</template>

<script setup>
import AuthDialog from "@/components/authentication/AuthDialog.vue"
import { ref } from "vue"
import { useAuthStore } from "@/store/auth"
import { onMounted } from "vue"

const auth = useAuthStore()

const dialogModel = ref(false)

const isAuthenticated = ref(false)

onMounted(async () => {
  isAuthenticated.value = await auth.isAuthenticated()
})

const openDialog = () => {
  dialogModel.value = true
}

const closeDialog = () => {
  dialogModel.value = false
}
</script>
