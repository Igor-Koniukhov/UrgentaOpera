<template>
  <v-btn
   @click="$store.commit(toggleSorting.comm)"
   color="primary"
   class="button-done-sorting"
   > {{toggleSorting.buttonTitle}} </v-btn>
</template>

<script>
export default {
props:['toggleSorting']

};
</script>

<style lang="scss">
.button-done-sorting{
  position: fixed;
  bottom: 20px;
  left:50%;
  transform: transition(-50%);

}
</style>
