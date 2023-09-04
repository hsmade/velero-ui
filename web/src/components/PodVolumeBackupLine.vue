<template>
  <v-card
    class="pa-0 ma-2"
    :style="properties.data.status.phase==='Failed'?'border-left: 3px solid #FF7F00':properties.data.status.phase==='Completed'?'border-left: 3px solid #7FFF00':'border-left: 3px solid #0000FF'"
  >
    <v-card-text class="pa-0 ma-0">
      <v-row no-gutters class="pa-0 ma-0">
        <v-col>
          <v-sheet class="pa-0 ma-0 pl-1">
            <b>Pod:</b>
            <br/>
            <b>Namespace:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{ properties.data.spec.pod.name}}
            <br/>
            {{ properties.data.spec.pod.namespace}}
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
            {{ properties.data.spec.volume}}
            <br/>
            <v-hover v-slot="{ isHovering, props }">
              <div v-bind="props">
                <div v-if="properties.data.status.phase === 'InProgress'" >
                  <v-progress-linear striped color="primary" :model-value=100*properties.data.status.progress.bytesDone/properties.data.status.progress.totalBytes />
                </div>
                <div v-else>{{ properties.data.status.phase }}</div>
                <v-expand-transition>
                  <div
                    v-if="isHovering"
                    class="d-flex transition-fast-in-fast-out"
                    style="height: 100%;"
                  >
                    <div v-if="properties.data.status.progress.bytesDone !== undefined">
                      {{ properties.data.status.progress.bytesDone }} / {{ properties.data.status.progress.totalBytes }} bytes
                    </div>
                    <div v-else>
                      {{ properties.data.status.message }}
                    </div>
                  </div>
                </v-expand-transition>
              </div>
            </v-hover>
          </v-sheet>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup>
const properties = defineProps({
  data: Object
})
</script>
