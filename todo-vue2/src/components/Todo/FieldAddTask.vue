<template>
  <v-text-field
    class="pa-3"
    v-model="newTaskTitle"
    @click:append="addTask"
    @keyup.enter="addTask"
    outlined
    label="+Task"
    hide-details
    clearable
  >
    <template v-slot:append>
      <v-icon @click="addTask" color="primary" :disabled="newTaskTitleInvalid">
        mdi-plus
      </v-icon>
    </template>
  </v-text-field>
</template>

<script>
export default {
  data() {
    return {
      newTaskTitle: "",
    };
  },
  methods: {
    addTask() {
      if (!this.newTaskTitleInvalid) {
        this.$store.dispatch("addTask", this.newTaskTitle)
        this.newTaskTitle = ''
      }
    },
  },
  computed: {
    newTaskTitleInvalid() {
      return !this.newTaskTitle
    },
  },
};
</script>

<style></style>
