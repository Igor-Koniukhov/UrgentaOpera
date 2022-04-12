<template>
  <v-dialog
      :value="true"
      persistent
      max-width="290"
    >
   
      <v-card>
        <v-card-title class="text-h5">
         Delete board?
        </v-card-title>
        <v-card-text>Are you sure you wanna delete this board?</v-card-text>
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
            @click="deleteBoard(board.list)"
          >
            Yes
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script>
export default {
    props: ['board'],
    methods:{
         deleteBoard(id) {
      this.$store.dispatch("todo/deleteBoard", id);
    },
    }

}
</script>

<style>

</style>