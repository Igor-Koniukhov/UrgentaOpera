<template>
  <v-dialog :value="true" persistent max-width="320">
    <v-card>
      <v-card-title class="text-h5"> Update the fields: </v-card-title>
      <v-card-text>
        <v-text-field
          class="pb-3 mr-auto"
          label="Title"
          v-model="boardTitle"
          outlined
          hide-details
          clearable
        />
        <div class="pa-3">
          <v-row>
            <div>Choose task background :</div>
            <input
              class="ml-auto"
              type="color"
              v-model="boardBackgroundColor"
            />
          </v-row>
          <v-row>
            <div>Choose text color:</div>
            <input class="ml-auto" v-model="color" type="color" />
          </v-row>
        </div>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn text @click="$emit('close')"> close </v-btn>
        <v-btn
          @click="updateBoard"
          :disabled="boardTitleInvalid"
          color="primary"
        >
          Update Board
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: ["board"],
  components: {},
  data() {
    return {
      boardTitle: null,
      boardBackgroundColor: null,
      color: null,
    };
  },
  methods: {
    updateBoard() {
      if (!this.boardTitleInvalid) {
        let payload = {
          list: this.board.list,
          title: this.boardTitle,
          backgroundColor: this.boardBackgroundColor,
          color: this.color,
        };
        this.$store.dispatch("todo/updateBoard", payload);
        setTimeout(() => {
          this.$emit("close");
        }, 300);
      }
    },
  },
  computed: {
    boardTitleInvalid() {
      return !this.boardTitle;
    },
  },
  mounted() {
    this.boardTitle = this.board.title;
    this.boardBackgroundColor = this.board.backgroundColor;
    this.color = this.board.color;
  },
};
</script>

<style></style>
