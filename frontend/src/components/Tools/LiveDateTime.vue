<template>
  <div
    class="text-subtitle-1 mt-1 ml-1 ">
   <small>  {{ date }}</small>
  </div>
</template>

<script>
import { format } from "date-fns";
export default {
  props: ["drawer"],
  data() {
    return {
      date: "",
    };
  },
  methods: {
    getDate() {
      this.date = format(new Date(), "MMMM d - iiii, H:mm:ss");
      setTimeout(this.getDate, 1000);
    },
  },
  mounted() {
    this.getDate();
  },
};
</script>

<style lang="scss">
.federica-font {
  &__date {
    color: rgba(250, 250, 250, 0.9);
  }
}
</style>
