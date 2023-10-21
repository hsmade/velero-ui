<template>
  <storage-location-line
    v-if="data !== null"
    v-for="location in data"
    key="location.metadata.uid"
    :data="location"
  />
</template>

<script setup>
import {onBeforeUnmount, onMounted, ref} from 'vue'
import StorageLocationLine from "@/components/StorageLocationLine.vue";

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
    fetch("/api/v1/storagelocations")
      .then(response => response.json())
      .then(response => data.value = response.result
        .sort((a, b) => new Date(b.status.lastSyncedTime) - new Date(a.status.lastSyncedTime))
      )
  }
</script>
