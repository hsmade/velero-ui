<template>
  <v-card
    max-width="344"
    :to="`/backup/${props.data.metadata.name}`"
  >
    <v-card-item>
      <v-card-title>{{ props.data.metadata.name }} in ns:{{ props.data.metadata.namespace }} => {{ props.data.spec.storageLocation }}</v-card-title>
      <v-card-subtitle>{{ props.data.status.phase }}</v-card-subtitle>
    </v-card-item>
    <v-card-text>
      <BackupStatus :status="props.data.status"/>
    </v-card-text>
  </v-card>
</template>

<script setup>
import BackupStatus from "@/components/BackupStatus.vue";

const props = defineProps({
  data: Object
})
</script>
