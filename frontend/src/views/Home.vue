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
                        <button class="btn btn-ghost btn-xs" @click="confirmDelete(item.id)">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24"
                                stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M6 18L18 6M6 6l12 12" />
                            </svg>
                        </button>
                    </th>
                </tr>
            </tbody>
        </table>
    </div>
    <div class="fixed bottom-4 right-4">
        <button class="btn btn-primary" @click="showCard = true">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
            </svg>
        </button>
    </div>
    <VueFinalModal v-model="showCard" class="fixed inset-0 z-50 flex items-center justify-center"
        content-class="w-96 bg-base-100 shadow-xl rounded-lg">
        <div class="card bg-base-100 w-96 shadow-xl">
            <div class="card-body">
                <label class="input input-bordered flex items-center gap-2">
                    文件名
                    <input type="text" class="grow" placeholder="为需要备份的文件起一个文件名" />
                </label>
                <label class="input input-bordered flex items-center gap-2">
                    文件路径
                    <input type="text" class="grow" placeholder="文件现在所在的路径" />
                </label>
                <label class="input input-bordered flex items-center gap-2">
                    备份路径
                    <input type="text" class="grow" placeholder="文件备份需要存储的路径" />
                </label>
                <div class="card-actions justify-end">
                    <button class="btn btn-primary" @click="showCard = false">添加</button>
                </div>
            </div>
        </div>
    </VueFinalModal>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { deleteBackup, getBackupList } from "../api/backup";
import { BackupItem } from "../models/home";
import { VueFinalModal } from "vue-final-modal";

const backupList = ref(<BackupItem[]>[]);
const selectAll = ref(false);
const showCard = ref(false);  // 控制显示

// fetchData 调用后端api获取数据
async function fetchData() {
    // 获取备份文件列表
    const response = await getBackupList();
    backupList.value = response.data.map((item: BackupItem) => ({ ...item, selected: false }));
}

fetchData()

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

// 删除确认窗口
function confirmDelete(id: number) {
    if (window.confirm("确认要删除此项吗？")) {
        deleteItem(id)
    }
}

// 删除项的方法
async function deleteItem(id: number) {
    // 调用后端接口删除配置项
    const response = await deleteBackup(id);
    backupList.value = backupList.value.filter(item => item.id !== id);
}

// 创建备份文件配置
function createBackup() {
    console.log('createBackup');
}
</script>
