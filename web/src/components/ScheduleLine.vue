<template>
  <v-card
    class="pa-0 ma-2"
    :to="`/schedule/${properties.data.metadata.name}`"
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
            <b>Last backup:</b>
            <br/>
            <b>foo:</b>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-0 ma-0">
            {{ properties.data.status.lastBackup }}
            <br/>
            bar
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
