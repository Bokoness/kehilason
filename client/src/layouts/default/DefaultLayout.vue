<template>
  <VApp>
    <DefaultBar v-if="viewBar" />

    <DefaultView />
  </VApp>
</template>

<script setup>
import DefaultBar from './AppBar.vue'
import DefaultView from './DefaultView.vue'

import { onBeforeRouteUpdate, useRoute } from 'vue-router'
import { ref } from 'vue'
const route = useRoute()

const viewBar = ref(!route.name.includes(['Login']))

onBeforeRouteUpdate(async (to) => {
  if (to.name.includes(['Login'])) {
    viewBar.value = false
  } else {
    viewBar.value = true
  }
})
</script>
