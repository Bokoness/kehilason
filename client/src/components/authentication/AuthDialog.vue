<template>
  <VDialog v-model="dialog" max-width="600" persistent>
    <template #activator>
      <VBtn color="indigo" @click="dialog = true">
        <VIcon>mdi-account</VIcon>
        <span>התחברו</span>
      </VBtn>
    </template>
    <VCard>
      <VCardTitle>
        <div class="d-flex justify-space-between">
          <span v-if="currentTab === 'login'">התחברות</span>
          <span v-else>הרשמה</span>
          <VIcon @click="dialog = false">mdi-close</VIcon>
        </div>
      </VCardTitle>

      <VCardText>
        <VWindow v-model="currentTab">
          <VWindowItem value="login">
            <LoginDialog @closeDialog="dialog = false" />
          </VWindowItem>

          <VWindowItem value="register">
            <RegisterDialog @goLogin="toggleTab" />
          </VWindowItem>
        </VWindow>
      </VCardText>

      <VCardActions class="d-flex justify-center">
        <section v-if="currentTab === 'register'" @click="toggleTab">
          רשומים? התחברו עכשיו
        </section>
        <section v-else @click="toggleTab">עוד לא רשומים? הרשמו עכשיו</section>
      </VCardActions>
    </VCard>
  </VDialog>
</template>

<script setup>
import { ref } from "vue"
import LoginDialog from "@/components/authentication/LoginDialog.vue"
import RegisterDialog from "@/components/authentication/RegisterDialog.vue"

const currentTab = ref("login")

const dialog = ref(false)

const toggleTab = () => {
  currentTab.value = currentTab.value === "register" ? "login" : "register"
}
</script>

<style lang="scss" scoped>
section {
  cursor: pointer;
}
</style>
