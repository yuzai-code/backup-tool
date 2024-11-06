<template>
    <Header />
  <div class="container-fluid">
    <div class="main-content">
      <Sidebar />
      <div class="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Sidebar from "./components/layout/Sidebar.vue";
import Header from "./components/layout/Header.vue";


</script>
<style scoped>
.container-fluid {
  display: flex;
}
.main-content {
  flex: 1;
  display: flex;
}
.content {
  flex: 1;
  padding: 16px;
}
</style>
