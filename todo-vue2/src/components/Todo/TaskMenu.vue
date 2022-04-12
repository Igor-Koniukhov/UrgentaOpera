<template>
  <div>
    <v-menu bottom left>
      <template v-slot:activator="{ on, attrs }">
        <v-btn icon v-bind="attrs" v-on="on">
          <v-icon>mdi-dots-vertical</v-icon>
        </v-btn>
      </template>

      <v-list>
        <v-list-item
          v-for="(item, index) in items"
          :key="index"
          @click="handleClick(index)"
        >
          <v-list-item-icon>
            <v-icon v-text="item.icon"></v-icon>
          </v-list-item-icon>
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
    <dialog-delete
      v-if="dialogs.delete"
      @close="dialogs.delete = false"
      :task="task"
    />
    <dialog-due-date
      v-if="dialogs.dueDate"
      @close="dialogs.dueDate = false"
      :task="task"
    />
     <dialog-edit
      v-if="dialogs.edit"
      @close="dialogs.edit = false"
      :task="task"
    /> 
    
  </div>
</template>

<script>
export default {
  props: ["task"],
  data: () => ({
    dialogs: {
      delete: false,
      edit: false,
      dueDate: false,
     
    },
    items: [
      {
        title: "Edit",
        icon: "mdi-pencil-outline",
        click() {
         this.dialogs.edit = true;
        },
      },
      {
        title: "Due date",
        icon: "mdi-calendar-month-outline",
        click() {
         this.dialogs.dueDate = true;
        },
      },
      {
        title: "Delete",
        icon: "mdi-trash-can-outline",
        click() {
          this.dialogs.delete = true;
        },
      },
      {
        title: "Sort",
        icon: "mdi-drag-horizontal-variant",
        click() {
          this.$store.commit('todo/toggleSorting')
        },
      },     
    ],
  }),
  methods: {
    handleClick(index) {
      this.items[index].click.call(this);
    },
  },
  components: {
    "dialog-delete": require("@/components/Dialogs/DialogDelete.vue").default,
    "dialog-edit": require("@/components/Dialogs/DialogEdit.vue").default,
    'dialog-due-date': require("@/components/Dialogs/DialogDueDate.vue").default,    
  },
};
</script>

<style></style>
