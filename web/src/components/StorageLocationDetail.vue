<template>
  <v-card
    v-if="data.storage_location !== undefined"
  >
    <v-card-item>
      <v-card-title>{{ data.storage_location.metadata.name }}</v-card-title>
    </v-card-item>
  </v-card>

  <v-container :fluid=true>
    <v-row v-if="data.storage_location">
      <v-col>
        <p>Storage Location resource</p>
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
            v-for="(value, key) in data.storage_location.metadata"
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
            v-for="(value, key) in data.storage_location.spec"
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
          <tr v-for="(value, key) in data.storage_location.status">
            <td><b>{{key}}</b></td>
            <td>{{value}}</td>
          </tr>
          </tbody>
        </v-table>
        </v-sheet>
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
    fetch("/api/v1/storagelocation/" + route.params.name)
      .then(response => response.json())
      .then(response => data.value = response.result)
  }
</script>
