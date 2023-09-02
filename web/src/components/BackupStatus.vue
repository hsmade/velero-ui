<template>
  <div v-if="props.status.progress.itemsBackedUp !== undefined">
    {{ props.status.progress.itemsBackedUp }} / {{ props.status.progress.totalItems }} items done
    <div v-if="props.status.phase === 'InProgress'" >
      <v-progress-linear striped color="primary" model-value=100*props.status.progress?.itemsBackedUp/props.status.progress?.totalItems />
      <br />
    </div>
  </div>
  <div v-if="props.status.progress.bytesDone !== undefined">
    {{ props.status.progress.bytesDone }} / {{ props.status.progress.totalBytes }} bytes done
    <div v-if="props.status.phase === 'InProgress'" >
      <v-progress-linear striped color="primary" model-value=100*props.status.progress?.bytesDone/props.status.progress?.totalBytes />
      <br />
    </div>
  </div>
  <div>
    <b>Started:</b>
    {{ props.status.startTimestamp}}
  </div>
  <div v-if="props.status.phase === 'Completed'">
    <b>Completed:</b>
    {{ props.status.completionTimestamp}}
  </div>
  <div v-if="props.status.expiration !== undefined">
    <b>Expires:</b>
    {{ props.status.expiration}}
  </div>
</template>

<script setup>
const props = defineProps({
  status: Object
})
</script>
