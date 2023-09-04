<template>
  <v-card
    class="pa-0 ma-2"
    :style="properties.data.status.phase.endsWith('Failed')?
      'border-left: 3px solid #FF7F00': // red
      properties.data.status.phase==='Completed'?
        'border-left: 3px solid #7FFF00': // green
        'border-left: 3px solid #FFC107' // orange
      "
    :to="`/backup/${properties.data.metadata.name}`"
  >
    <v-card-text class="pa-0 ma-0">
      <v-row no-gutters class="pa-0 ma-0">
        <v-col>
          <v-sheet class="pa-0 ma-0 pl-1">
            <b>Name:</b>
            <br/>
            <b>Namespace:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{ properties.data.metadata.name}}
            <br/>
            {{ properties.data.metadata.namespace}}
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            <b>Storage:</b>
            <br/>
            <b>Progress:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{ properties.data.spec.storageLocation}}
            <br/>
            <div v-if="properties.data.status.phase === 'InProgress'" >
              <v-progress-linear
                striped
                color="primary"
                :model-value=100*properties.data.status.progress.itemsBackedUp/properties.data.status.progress.totalItems
              >
              </v-progress-linear>
            </div>
            <div v-else>{{ properties.data.status.phase }}</div>
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
