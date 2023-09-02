<template>
  <v-card
    v-if="data.Backup !== undefined"
  >
    <v-card-item>
      test
      <v-card-title>{{ data.Backup.metadata.name }} in ns:{{ data.Backup.metadata.namespace }} => {{ data.Backup.spec.storageLocation }}</v-card-title>
      <v-card-subtitle>{{ data.Backup.status.phase }}</v-card-subtitle>
    </v-card-item>
    <v-card-text>
      <BackupStatus :status="data.Backup.status"/>
    </v-card-text>
  </v-card>

  <v-card
    v-if="data.PodVolumeBackups !== undefined"
    v-for="item in data.PodVolumeBackups"
    key="item.metadata.name"
    >
    <v-card-item>
      <v-card-title>{{ item.metadata.name }} - {{item.status.phase}}</v-card-title>
      <v-row no-gutters>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            <b>Pod:</b>
            <br/>
            <b>Namespace:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{ item.spec.pod.name}}
            <br/>
            {{ item.spec.pod.namespace}}
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            <b>Volume:</b>
            <br/>
            <b>Progress:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{ item.spec.volume}}
            <br/>
            <v-hover v-slot="{ isHovering, props }">
              <div v-bind="props">
              <div v-if="item.status.phase === 'InProgress'" >
                <v-progress-linear striped color="primary" model-value=100*item.status.progress.bytesDone/item.status.progress.totalBytes />
              </div>
              <div v-else>{{ item.status.phase }}</div>
              <v-expand-transition>
                <div
                  v-if="isHovering"
                  class="d-flex transition-fast-in-fast-out"
                  style="height: 100%;"
                >
                  <div v-if="item.status.progress.bytesDone !== undefined">
                    {{ item.status.progress.bytesDone }} / {{ item.status.progress.totalBytes }} bytes
                  </div>
                  <div v-else>
                    {{ item.status.message }}
                  </div>
                </div>
              </v-expand-transition>
              </div>
            </v-hover>
          </v-sheet>
        </v-col>
      </v-row>
    </v-card-item>
  </v-card>
</template>

<script setup>
import {onBeforeUnmount, onMounted, ref} from 'vue'
import {useRoute} from "vue-router";
import BackupStatus from "@/components/BackupStatus.vue";
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
</script>
