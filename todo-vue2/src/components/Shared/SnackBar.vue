<template>
   <v-snackbar
      v-model="$store.state.todo.snackbar.show"
    >
      {{ $store.state.todo.snackbar.text }}

      <template v-slot:action="{ attrs }">
        <v-btn
         @click="$store.commit('todo/hideSnackBar')"
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