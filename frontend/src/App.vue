<template>
    <div class="main-content">
        <Sidebar />
        <div class="content">
            <router-view />
        </div>
    </div>
</template>

<script setup>
import Sidebar from "./components/layout/sidebar";
</script>
<style scoped>
.main-content {
    flex: 1;
    display: flex;
}

.content {
    flex: 1;
    padding: 16px;
    background-color: #fafafa;
}
</style>
