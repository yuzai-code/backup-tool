<template>
        <h3>dsfafdsagsdagasdg</h3>
        <ul>
          <li v-for="item in backupList" :key="item.id">{{ item.dir_name }}</li>
        </ul>
</template>
      
<script setup lang="ts">
import { ref } from 'vue';
import { getBackupList } from '../api/backup';
import { BackupItem } from '../models/home';

const backupList = ref(<BackupItem[]>([]));
      
async function fetchData() {
const response = await getBackupList();
backupList.value = response.data;
}
      
fetchData();
</script>