<template>
  <v-dialog :value="true" persistent max-width="320">
    <v-card>
      <v-card-title class="text-h5"> Fill the feilds: </v-card-title>
      <v-card-text>
        <v-text-field
          class="pb-3 mr-auto"
          label="Title"
          v-model="board.title"
          outlined
          hide-details
          clearable
        />
        <div class="pa-3">
          <v-row>
            <div>Choose task backgraund :</div>
            <input
              class="ml-auto"
              type="color"
              v-model="board.backgroundColor"
            />
          </v-row>
          <v-row>
            <div>Choose text color:</div>
            <input class="ml-auto" v-model="board.color" type="color" />
          </v-row>
        </div>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn text @click="$emit('close')"> close </v-btn>
        <v-btn @click="addBoard" color="primary"> Add Board </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  components: {},
  data() {
    return {
      list: this.$store.state.boards.length,
      board: {
        list: "",
        title: "",
        backgroundColor: "",
        color: "",
      },
    };
  },
  methods: {
    addBoard() {
      this.list++;
      this.$store.dispatch("addBoard", {
        list: this.list,
        title: this.board.title,
        backgroundColor: this.board.backgroundColor,
        color: this.board.color,
      });
      this.title = "";
      this.backgroundColor = "";
      this.color = "";
    },
  },
};
</script>

<style></style>
