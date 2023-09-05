<template>
  <v-card
    v-if="data.Backup !== undefined"
  >
    <v-card-item>
      <v-card-title>{{ data.Backup.metadata.name }}</v-card-title>
      <v-card-subtitle v-if="'velero.io/schedule-name' in data.Backup.metadata.labels">
            Schedule: {{ data.Backup.metadata.labels["velero.io/schedule-name"] }}
      </v-card-subtitle>
    </v-card-item>
  </v-card>

  <v-container fluid="true">
    <v-row v-if="data.Backup">
      <v-col>
        <p>Backup resource</p>
        <v-sheet elevation="1" class="ma-2">
        <v-table density="compact">
          <thead>
          <tr>
            <th colspan="2" style="text-align: center">
              Metadata
            </th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="(value, key) in data.Backup.metadata"
          >
            <td v-if="!['managedFields'].includes(key)"><b>{{key}}</b></td>
            <td v-if="!['managedFields'].includes(key)">
              <v-table v-if="typeof value === 'object'">
                <tbody>
                <tr v-for="(m_value, m_key) in value">
                  <td><b>{{ m_key }}</b></td>
                  <td>{{ m_value }}</td>
                </tr>
                </tbody>
              </v-table>
              <div v-else>
                {{value}}
              </div>
            </td>
          </tr>
          </tbody>

          <thead>
          <tr>
            <th colspan="2" style="text-align: center">
              Spec
            </th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="(value, key) in data.Backup.spec"
          >
            <td v-if="JSON.stringify(value) !== '{}'"><b>{{key}}</b></td>
            <td v-if="JSON.stringify(value) !== '{}'">{{value}}</td>
          </tr>
          </tbody>

          <thead>
          <tr>
            <th colspan="2" style="text-align: center">
              Status
            </th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(value, key) in data.Backup.status">
            <td><b>{{key}}</b></td>
            <td v-if="key==='progress'">
              {{ data.Backup.status.progress.itemsBackedUp }} / {{ data.Backup.status.progress.totalItems }}
              <div v-if="data.Backup.status.phase === 'InProgress'">
                <v-progress-linear striped color="primary" :model-value="(100*data.Backup.status.progress.itemsBackedUp)/data.Backup.status.progress.totalItems"/>
              </div>
            </td>
            <td v-else>{{value}}</td>
          </tr>
          </tbody>
        </v-table>
        </v-sheet>
      </v-col>
      <v-col>
        <p>Pod Volume Backups</p>
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
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import {computed, onBeforeUnmount, onMounted, ref} from 'vue'
import {useRoute} from "vue-router";
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
