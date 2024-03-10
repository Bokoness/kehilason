<template>
  <div>
    <VSnackbar v-model="snackbar" :color="snackColor">
      {{ snackMessage }}
    </VSnackbar>
  </div>
</template>

<script setup>
import { ref } from "vue"
import eventBus from "@/eventBus"

const snackMessage = ref("")
const snackColor = ref("red")
const snackbar = ref(false)

eventBus.on("error", (message) => {
  snackMessage.value = message
  snackbar.value = true
})

eventBus.on("snackbar", ({ message, color }) => {
  snackMessage.value = message
  snackColor.value = color
  snackbar.value = true
})
</script>
