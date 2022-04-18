<template>
  <v-snackbar
      v-model="$store.state.auth.snackbar.show"
      color="green"
  >
    {{ $store.state.auth.snackbar.text }}

    <template v-slot:action="{ attrs }">
      <v-btn
          @click="$store.commit('auth/hideSnackBar')"
          text
          v-bind="attrs"

      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
export default {

}
</script>

<style>

</style>