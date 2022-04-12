<template>
  <div class="home">
    <field-add-task />
    <div class="pa-3" v-if="$store.state.todo.tasks.length">
      <draggable handle=".handle-board" group="board" v-model="boards">
        <list-tasks
          class="ma-1 board"
          v-for="(board, index) in boards"
          :key="board.list"
          :index="index"
          :board="board"
          :list="board.list"
        />
      </draggable>
    </div>

    <no-tasks v-else />
    <button-done-sorting
      v-if="$store.state.todo.sorting"
      :toggleSorting="toggleSorting"
    />
    <button-done-sorting
      v-if="$store.state.todo.boardSorting"
      :toggleSorting="toggleBoardSorting"
    />
  </div>
</template>

<script>
import draggable from "vuedraggable";
// @ is an alias to /src

export default {
  name: "ToDo",
  components: {
    FieldAddTask: require("@/components/Todo/FieldAddTask.vue").default,
    ListTasks: require("@/components/Todo/ListTasks.vue").default,
    NoTasks: require("@/components/Todo/NoTasks.vue").default,
    ButtonDoneSorting: require("@/components/Todo/ButtonDoneSorting.vue")
      .default,
    draggable,
  },
  data() {
    return {
      toggleSorting: {
        commitMessage: "todo/toggleSorting",
        buttonTitle: "finish task sorting",
      },
      toggleBoardSorting: {
        commitMessage: "todo/toggleBoardSorting",
        buttonTitle: "finish board sorting",
      },
    };
  },
  computed: {
    boards: {
      get() {
        return this.$store.state.todo.boards;
      },
      set(value) {
        this.$store.dispatch("todo/setBoards", value);
      },
    },
  },
};
</script>
<style lang="scss">
.v-list.board {
  border: 1px solid rgba(0, 0, 0, 0.1);
  max-width: 300px;
  width: 100%;
  display: inline-flex;
  flex-direction: column;
}
</style>
