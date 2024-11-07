<template>
  <div class="main-content">
    <Sidebar />
    <div class="content">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
import Sidebar from "./components/layout/Sidebar.vue";
</script>
<style scoped>
.main-content {
  flex: 1;
  display: flex;
}
.content {
  flex: 1;
  padding: 16px;
}
</style>
