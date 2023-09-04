<template>
  <backup-line
    v-if="data !== null"
    v-for="backup in data"
    key="backup.metadata.uid"
    :data="backup"
  />
</template>

<script setup>
import {onBeforeUnmount, onMounted, ref} from 'vue'
import BackupLine from "@/components/BackupLine.vue";

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
    fetch("/api/v1/backups")
      .then(response => response.json())
      .then(response => data.value = response.result)
  }
</script>
