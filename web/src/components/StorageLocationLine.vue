<template>
  <v-card
    class="pa-0 ma-2"
    :style="properties.data.status.phase?.endsWith('Failed')?
      'border-left: 3px solid #FF7F00': // red
      properties.data.status.phase==='Available'?
        'border-left: 3px solid #7FFF00': // green
        'border-left: 3px solid #FFC107' // orange
      "
    :to="`/storagelocation/${properties.data.metadata.name}`"
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
            {{properties.data.metadata.name}}
            <br/>
            {{properties.data.metadata.namespace}}
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            <b>Provider:</b>
            <br/>
            <b>Access Mode:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{properties.data.spec.provider}}
            <br/>
            {{properties.data.spec.accessMode}}
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            <b>Bucket:</b>
            <br/>
            <b>Prefix:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{properties.data.spec.objectStorage.bucket}}
            <br/>
            {{properties.data.spec.objectStorage.prefix}}
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
