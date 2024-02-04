<template>
  <VAppBar flat color="secondaryBG" class="pl-15">
    <VRow align="center" justify="space-between" class="px-3">
      <VCol cols="3">
        <VImg src="@/assets/logo.png" max-height="80" max-width="=250" />
      </VCol>

      <VCol lg="4" md="5 " cols="7">
        <VRow align="center">
          <VCol>קהילות</VCol>
          <VCol>אודות</VCol>
          <VCol>צור קשר</VCol>
          <VCol>
            <v-avatar v-if="isLoggedIn" color="info">
              <v-icon icon="mdi-account-circle"></v-icon>
            </v-avatar>
            <router-link v-else to="/login">
              <VBtn color="indigo">
                <VIcon>mdi-account</VIcon>
                <span>התחברו</span>
              </VBtn>
            </router-link>
          </VCol>
        </VRow>
      </VCol>
    </VRow>
  </VAppBar>
</template>
<script setup>
import { useAuthStore } from "@/store/auth";

const { isLoggedIn } = useAuthStore();
</script>
