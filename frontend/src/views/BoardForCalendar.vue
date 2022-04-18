<template>
  <div class="calendar pa-2 mt-9">
<calendar></calendar>
  </div>
</template>

<script>
export default {
    name: "BoardForCalendar",
    components: {
        'calendar': require('@/components/Calendar/Calendar.vue').default,
    }

}
</script>


<style>

</style>