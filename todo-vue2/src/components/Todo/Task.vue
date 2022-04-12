<template>
  <div>
    <v-list-item
      @click="doneTask(task.id)"
      class="white"
      :class="task.done ? 'yellow lighten-4' : 'indigo lighten-4'"
      :ripple="false"
    >
      <template v-slot:default>
        <v-list-item-action>
          <v-checkbox :input-value="task.done" color="primary"></v-checkbox>
        </v-list-item-action>
        <v-list-item-content>
          <div
            class="text-in-task"
            :class="{ 'text-decoration-line-through': task.done }"
          >
            {{ task.title }}
          </div>
        </v-list-item-content>

        <v-list-item-action v-if="task.dueDate">
          <v-list-item-action-text>
            <v-icon small>mdi-calendar</v-icon>
            {{ task.dueDate | niceDate }}
          </v-list-item-action-text>
        </v-list-item-action>
        <v-list-item-action v-if="$store.state.todo.sorting">
          <v-btn icon class="handle">
            <v-icon>mdi-drag-horizontal</v-icon>
          </v-btn>
        </v-list-item-action>
        <v-list-item-action>
          <task-menu :task="task" />
        </v-list-item-action>
      </template>
    </v-list-item>

    <v-divider></v-divider>
  </div>
</template>

<script>
import { format } from "date-fns";
export default {
  props: ["task", "board"],
  filters: {
    niceDate(value) {
      return format(new Date(value), "MMM d");
    },
  },
  data() {
    return {};
  },
  components: {
    "task-menu": require("@/components/Todo/TaskMenu.vue").default,
  },
  methods: {
    doneTask(id) {
      this.$store.dispatch("todo/doneTask", id);
    },
  },
};
</script>

<style lang="scss">
.v-list-item__content > * {
  font-size: 12px;
}
.sortable-ghost {
  opacity: 0.4;
  
}
.sortable-drag{
  box-shadow:0 0 10px rgba(0,0,0,.3);
}
</style>
