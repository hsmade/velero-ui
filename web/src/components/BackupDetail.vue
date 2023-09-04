<template>

  <v-card
    v-if="data.Backup !== undefined"
  >
    <v-card-item>
      <v-card-title>{{ data.Backup.metadata.name }} in ns:{{ data.Backup.metadata.namespace }} => {{ data.Backup.spec.storageLocation }}</v-card-title>
      <v-card-subtitle>{{ data.Backup.status.phase }}</v-card-subtitle>
    </v-card-item>
    <v-card-text>
      <BackupStatus :data="data.Backup"/>
    </v-card-text>
  </v-card>

  <pod-volume-backup-line
    v-if="data.PodVolumeBackups !== undefined"
    v-for="item in in_progress_pods"
    key="item.metadata.name"
    :data="item"
  />
  <pod-volume-backup-line
    v-if="data.PodVolumeBackups !== undefined"
    v-for="item in completed_pods"
    key="item.metadata.name"
    :data="item"
  />
</template>

<script setup>
import {computed, onBeforeUnmount, onMounted, ref} from 'vue'
import {useRoute} from "vue-router";
import BackupStatus from "@/components/BackupStatus.vue";
import PodVolumeBackupLine from "@/components/PodVolumeBackupLine.vue";

  const route = useRoute()
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

  let data = ref(Object)

  function update() {
    fetch("/api/v1/backup/" + route.params.name)
      .then(response => response.json())
      .then(response => data.value = response.result)
  }

  const completed_pods = computed(() => {
    return data.value.PodVolumeBackups
      .filter(item => item.status.phase !== "InProgress" )
      .sort((a, b) => new Date(b.status.startTimestamp) - new Date(a.status.startTimestamp))
  })
  const in_progress_pods = computed(() => {
    return data.value.PodVolumeBackups
      .filter(item => item.status.phase === "InProgress" )
      .sort((a, b) => new Date(b.status.startTimestamp) - new Date(a.status.startTimestamp))
  })
</script>
