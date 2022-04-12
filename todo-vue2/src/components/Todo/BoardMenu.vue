<template>
  <div>
    <v-menu bottom left>
      <template v-slot:activator="{ on, attrs }">
        <v-btn icon v-bind="attrs" v-on="on">
          <v-icon
           :style="{color: board.color}"
          >mdi-dots-vertical-circle-outline</v-icon>
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
    <dialog-delete-board
      v-if="dialogs.deleteBoard"
      @close="dialogs.deleteBoard = false"
      :board="board"
    />   
     <dialog-edit-board
      v-if="dialogs.editBoard"
      @close="dialogs.editBoard = false"
      :board="board"
    /> 
    <dialog-add-board
      v-if="dialogs.board"
      @close="dialogs.board = false"
      :board="board"
    />    
  </div>
</template>

<script>
export default {
  props: ["board"],
  data: () => ({
    dialogs: {
      deleteBoard: false,
      editBoard: false,      
      board: false
    },
    items: [
      {
        title: "Edit",
        icon: "mdi-pencil-outline",
        click() {
         this.dialogs.editBoard = true;
        },
      },      
      {
        title: "Delete",
        icon: "mdi-trash-can-outline",
        click() {
          this.dialogs.deleteBoard = true;
        },
      },
      {
        title: "Sort Board",
        icon: "mdi-drag-horizontal-variant",
        click() {
          this.$store.commit('todo/toggleBoardSorting')
        },
      },
      {
        title: "Add Board",
        icon: "mdi-table-column-plus-before",
        click() {
          console.log("add board");
          this.dialogs.board = true;
          this.$store.commit('todo/toggleAddBoard')
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
    "dialog-delete-board": require("@/components/Dialogs/DialogDeleteBoard.vue").default,
    "dialog-edit-board": require("@/components/Dialogs/DialogEditBoard.vue").default,    
    'dialog-add-board': require("@/components/Dialogs/DialogAddBoard.vue").default,
  },
};
</script>

<style></style>
