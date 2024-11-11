<template>
    <div class="overflow-x-auto">
        <table class="table">
            <!-- head -->
            <thead>
                <tr>
                    <th>
                        <label>
                            <input type="checkbox" class="checkbox" v-model="selectAll" @change="toggleSelectAll" />
                        </label>
                    </th>
                    <th>文件名</th>
                    <th>文件所在路径</th>
                    <th>文件备份路径</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in backupList" :key="item.id">
                    <th>
                        <label>
                            <input type="checkbox" class="checkbox" v-model="item.selected" />
                        </label>
                    </th>
                    <td>
                        <div class="flex items-center gap-3">
                            <div class="font-bold">{{ item.dir_name }}</div>
                        </div>
                    </td>
                    <td>
                        {{ item.back_path }}
                        <!-- Zemlak, Daniel and Leannon -->
                        <!-- <br /> -->
                        <!-- <span class="badge badge-ghost badge-sm">Desktop Support Technician</span> -->
                    </td>
                    <td>{{ item.file_path }}</td>
                    <th>
                        <button class="btn btn-ghost btn-xs">details</button>
                    </th>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { getBackupList } from "../api/backup";
import { BackupItem } from "../models/home";

const backupList = ref(<BackupItem[]>[]);
const selectAll = ref(false);

// fetchData 调用后端api获取数据
async function fetchData() {
    // 获取备份文件列表
    const response = await getBackupList();
    backupList.value = response.data.map((item: BackupItem) => ({ ...item, selected: false }));
}

fetchData();

// 计算属性：检查所有项是否被选择
const allSelected = computed(() => {
    return backupList.value.every(item => item.selected);
});

// 监听selectAll 的变化，更新所有项的状态
function toggleSelectAll() {
    backupList.value.forEach(item => {
        item.selected = selectAll.value;
    });
}

// 监听backupList的变化，更新selectAll 的状态
watch(allSelected, (newValue) => {
    selectAll.value = newValue;
})
</script>
