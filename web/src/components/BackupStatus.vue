<template>
  <div v-if="properties.data.status.progress.itemsBackedUp !== undefined">
    {{ properties.data.status.progress.itemsBackedUp }} / {{ properties.data.status.progress.totalItems }} items done
    <div v-if="properties.data.status.phase === 'InProgress'" >
      <v-progress-linear striped color="primary" :model-value="(100*properties.data.status.progress.itemsBackedUp)/properties.data.status.progress.totalItems"/>
      <br />
    </div>
  </div>
  <div>
    <b>Storage location:</b>
    {{ properties.data.spec.storageLocation}}
  </div>
  <div>
    <b>Started:</b>
    {{ properties.data.status.startTimestamp}}
  </div>
  <div>
    <b>Completed:</b>
    {{ properties.data.status.completionTimestamp}}
  </div>
  <div v-if="properties.data.status.expiration !== undefined">
    <b>Expires:</b>
    {{ properties.data.status.expiration}}
  </div>
  <div v-if="properties.data.status.warnings > 0" >
    <b>Warnings:</b>
    {{ properties.data.status.warnings }}
  </div>
  <div v-if="properties.data.status.errors > 0" >
    <b>Errors:</b>
    {{ properties.data.status.errors }}
  </div>
</template>

<!--tbd: more detail -->
<!--namespaces / resources filters-->
<!--label selectors-->
<!--link to errors/logs-->
<!--raw yaml popup-->

<script setup>
const properties = defineProps({
  data: Object
})
</script>
