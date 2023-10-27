<template>
  <div>
    <VContainer>

      <h1 class="text-h4 d-flex justify-center">מעניין אתכם? צרו קשר</h1>
      <h3 align="center" class="text-h6">הרע או לשון הרע לאו הוא האוסרים מרכילות לעבירה ולפגוע, להבדיל אמירה כל בני הרע
        שיסודה איסור או או שדיברו.</h3>
      <VRow class="mt-5">
        <VCol md="4" cols="12">
          <VCard>
            <VCardTitle>
              <VIcon>mdi-email</VIcon>
            </VCardTitle>
            <VCardText>
              <div>שלחו מייל</div>
              <div class="text-caption">נשתדל לענות בזמן</div>
            </VCardText>
            <VCardActions>mail@gmail.com</VCardActions>
          </VCard>
        </VCol>


        <VCol md="4" cols="12">
          <VCard>
            <VCardTitle>
              <VIcon>mdi-phone</VIcon>
            </VCardTitle>
            <VCardText>
              <div>חייגו אלינו</div>
              <div class="text-caption">נשתדל לענות בזמן</div>
            </VCardText>
            <VCardActions>0523480844</VCardActions>
          </VCard>
        </VCol>


        <VCol md="4" cols="12">
          <VCard>
            <VCardTitle>
              <VIcon>mdi-whatsapp</VIcon>
            </VCardTitle>
            <VCardText>
              <div>שלחו וואטסאפ</div>
              <div class="text-caption">נשתדל לענות בזמן</div>
            </VCardText>
            <VCardActions>0523480844</VCardActions>
          </VCard>
        </VCol>
      </VRow>
    </VContainer>

    <section class="mt-10 bg-grey-darken-1">
      <VContainer>
        <h2  class="text-h5">השאירו פרטים ונחזור אליכם</h2>

        <VRow class="mt-5">
          <VCol>
            <VTextField label="שם מלא (חובה)" bg-color="white"/>
          </VCol>

          <VCol>
            <VTextField label="נייד (חובה)" bg-color="white"/>
          </VCol>

          <VCol>
            <VTextField label="מייל (חובה)" bg-color="white"/>
          </VCol>

          <VCol>
            <VTextField label="תוכן הפניה (חובה)" bg-color="white"/>
          </VCol>
        </VRow>
        <VRow justify="end" class="mb-5 ml-1">
          <VBtn>שלחו</VBtn>
        </VRow>
      </VContainer>
    </section>
  </div>
</template>
<script setup>
</script>
