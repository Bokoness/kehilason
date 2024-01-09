<template>
  <VAppBar flat color="secondaryBG" class="pl-15">
    <VRow align="center" justify="space-between" class="px-3">
      <VCol cols="3">
        <VImg src="@/assets/logoWithBlackWords.png" max-height="80" max-width="=250" />
      </VCol>

      <VCol lg="4" md="5 " cols="7">
        <VRow align="center">
          <VCol>קהילות</VCol>
          <VCol>אודות</VCol>
          <VCol>צור קשר</VCol>
          <VCol>
            <router-link to="/login">
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
