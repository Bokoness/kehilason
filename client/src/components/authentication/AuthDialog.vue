<template>
  <VCard>
    <VCardTitle>
      <div class="d-flex justify-space-between">
        <span v-if="currentTab === 'login'">התחברות</span>
        <span v-else>הרשמה</span>
        <VIcon @click="closeDialog">mdi-close</VIcon>
      </div>
    </VCardTitle>

    <VCardText>
      <VWindow v-model="currentTab">
        <VWindowItem value="login">
          <LoginDialog />
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
</template>

<script setup>
import { ref } from "vue"
import LoginDialog from "@/components/authentication/LoginDialog.vue"
import RegisterDialog from "@/components/authentication/RegisterDialog.vue"

const emit = defineEmits(["close-dialog"])

const currentTab = ref("login")

const closeDialog = () => {
  emit("close-dialog")
}

const toggleTab = () => {
  currentTab.value = currentTab.value === "register" ? "login" : "register"
}
</script>

<style lang="scss" scoped>
section {
  cursor: pointer;
}
</style>
