<template>
  <v-dialog
      :value="true"
      persistent
      max-width="290"
  >

    <v-card>
      <v-card-title class="text-h5">
        Confirm logout
      </v-card-title>
      <v-card-text>Are you sure you wanna logout?</v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn

            text
            @click="$emit('close')"
        >
          No
        </v-btn>
        <v-btn
            color="red darken-1"
            text
            @click="Logout"
        >
          Yes
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  methods:{
    Logout() {
      this.$store.commit('auth/setIsNotLogin')
      this.$emit('close')
      this.$router.push('/login')
    },
  }

}
</script>

<style>

</style>