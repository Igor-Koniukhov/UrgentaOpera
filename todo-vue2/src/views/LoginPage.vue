<template>
<v-col lg="8" md="10">
  <login-form/>
</v-col>
</template>

<script>
export default {
  name: "LoginPage",
  components: {
    'login-form': require('@/components/Forms/LoginForm.vue').default,
}
}
</script>

<style scoped>

</style>