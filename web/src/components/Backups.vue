<template>
  backups
  <v-card
    v-if="data !== null"
    v-for="backup in data"
    key="backup.metadata.uid"
    max-width="344"
    :to="`/backup/${backup.metadata.name}`"
  >
    <v-card-item>
      <v-card-title>{{ backup.metadata.name }} in ns:{{ backup.metadata.namespace }} => {{ backup.spec.storageLocation }}</v-card-title>
      <v-card-subtitle>{{ backup.status.phase }}</v-card-subtitle>
    </v-card-item>
    <v-card-text>
      <BackupStatus :status="backup.status"/>
    </v-card-text>
  </v-card>
</template>

<script setup>
import {onBeforeUnmount, onMounted, ref} from 'vue'
import BackupStatus from "@/components/BackupStatus.vue";
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
