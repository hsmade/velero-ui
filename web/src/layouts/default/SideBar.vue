<template>
    <v-layout>
      <v-navigation-drawer
        permanent
        style="background-color: #3F51B5; color:white;"
      >
        <v-list nav density="compact">
          <v-list-item to="/" prepend-icon="mdi-email" title="Home" value="Home"></v-list-item>
          <v-list-item to="/backups" prepend-icon="mdi-email" title="Backups" value="Backups"></v-list-item>
        </v-list>
      </v-navigation-drawer>
      <v-main style="height: 100%">
        <router-view />
      </v-main>
    </v-layout>
</template>
<script setup>
</script>
