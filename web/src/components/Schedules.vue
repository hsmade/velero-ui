<template>
  <schedule-line
    v-if="data !== null"
    v-for="schedule in data"
    key="schedule.metadata.name"
    :data="schedule"
  />
</template>

<script setup>
import {onBeforeUnmount, onMounted, ref} from 'vue'
import ScheduleLine from "@/components/ScheduleLine.vue";

let timer
onMounted(() => {
  timer = setInterval(()=>{
    update()
  }, 3000)
  update()
})
onBeforeUnmount(() => {
  clearInterval(timer)
})

let data = ref(Array)

function update() {
  fetch("/api/v1/schedules")
    .then(response => response.json())
    .then(response => data.value = response.result
      .sort((a, b) => new Date(b.status.lastBackup) - new Date(a.status.lastBackup))
    )
}
</script>
