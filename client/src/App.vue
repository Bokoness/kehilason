<template>
  <div>
    <RouterView />
    <VSnackbar v-model="snackbar" color="red">{{ errorMessage }}</VSnackbar>
  </div>
</template>

<script setup>
import { ref } from "vue"
import eventBus from "@/eventBus"

const errorMessage = ref("")
const snackbar = ref(false)

eventBus.on("axios-error", (message) => {
  errorMessage.value = message
  snackbar.value = true
})
</script>
