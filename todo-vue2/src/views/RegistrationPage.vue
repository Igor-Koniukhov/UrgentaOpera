<template>
<RegistrationForm/>
</template>

<script>
export default {
  name: "Registration",
  components:{
    'RegistrationForm': require('@/components/Forms/RegistrationForm.vue').default,
  }
}
</script>

<style scoped>

</style>