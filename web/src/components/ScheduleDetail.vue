<template>
  <v-card
    v-if="data.schedule !== undefined"
  >
    <v-card-item>
      <v-card-title>{{ data.schedule.metadata.name }}</v-card-title>
      <v-card-subtitle v-if="'velero.io/schedule-name' in data.schedule.metadata.labels">
            Schedule: {{ data.schedule.metadata.labels["velero.io/schedule-name"] }}
      </v-card-subtitle>
    </v-card-item>
  </v-card>

  <v-container :fluid=true>
    <v-row v-if="data.schedule">
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
            v-for="(value, key) in data.schedule.metadata"
          >
            <td v-if="!['managedFields', 'annotations'].includes(key)"><b>{{key}}</b></td>
            <td v-if="!['managedFields', 'annotations'].includes(key)">
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
            v-for="(value, key) in data.schedule.spec"
          >
            <td v-if="JSON.stringify(value) !== '{}'"><b>{{key}}</b></td>
            <td v-if="JSON.stringify(value) !== '{}'">
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
              Status
            </th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(value, key) in data.schedule.status">
            <td><b>{{key}}</b></td>
            <td v-if="key==='progress'">
              {{ data.schedule.status.progress.itemsBackedUp }} / {{ data.schedule.status.progress.totalItems }}
              <div v-if="data.schedule.status.phase === 'InProgress'">
                <v-progress-linear striped color="primary" :model-value="(100*data.schedule.status.progress.itemsBackedUp)/data.schedule.status.progress.totalItems"/>
              </div>
            </td>
            <td v-else>{{value}}</td>
          </tr>
          </tbody>
        </v-table>
        </v-sheet>
      </v-col>
      <v-col>
        <p>Backups</p>
        <backup-line
          v-if="data.backups !== undefined"
          v-for="item in in_progress_backups"
          key="item.metadata.name"
          :data="item"
        />
        <backup-line
          v-if="data.backups !== undefined"
          v-for="item in completed_backups"
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
import BackupLine from "@/components/BackupLine.vue";

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
    fetch("/api/v1/schedule/" + route.params.name)
      .then(response => response.json())
      .then(response => data.value = response.result)
  }

  const completed_backups = computed(() => {
    return data.value.backups
      .filter(item => item.status.phase !== "InProgress" )
      .sort((a, b) => new Date(b.status.startTimestamp) - new Date(a.status.startTimestamp))
  })
  const in_progress_backups = computed(() => {
    return data.value.backups
      .filter(item => item.status.phase === "InProgress" )
      .sort((a, b) => new Date(b.status.startTimestamp) - new Date(a.status.startTimestamp))
  })
</script>
