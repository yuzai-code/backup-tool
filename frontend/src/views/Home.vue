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
                        {{ item.file_path }}
                        <!-- Zemlak, Daniel and Leannon -->
                        <!-- <br /> -->
                        <!-- <span class="badge badge-ghost badge-sm">Desktop Support Technician</span> -->
                    </td>
                    <td>{{ item.back_path }}</td>
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
                <div class="form-control w-full">
                    <label class="input input-bordered flex items-center gap-2" :class="{ 'input-error': errors.dirName }">
                        文件名
                        <input 
                            type="text" 
                            class="grow" 
                            placeholder="为需要备份的文件起一个文件名" 
                            v-model="formData.dir_name"
                        />
                    </label>
                    <div class="text-error text-sm mt-1" v-if="errors.dirName">{{ errors.dirName }}</div>
                </div>

                <div class="form-control w-full">
                    <label class="input input-bordered flex items-center gap-2" :class="{ 'input-error': errors.filePath }">
                        文件路径
                        <input 
                            type="text" 
                            class="grow" 
                            placeholder="文件现在所在的路径" 
                            v-model="formData.file_path"
                        />
                    </label>
                    <div class="text-error text-sm mt-1" v-if="errors.filePath">{{ errors.filePath }}</div>
                </div>

                <div class="form-control w-full">
                    <label class="input input-bordered flex items-center gap-2" :class="{ 'input-error': errors.backPath }">
                        备份路径
                        <input 
                            type="text" 
                            class="grow" 
                            placeholder="文件备份需要存储的路径" 
                            v-model="formData.back_path"
                        />
                    </label>
                    <div class="text-error text-sm mt-1" v-if="errors.backPath">{{ errors.backPath }}</div>
                </div>

                <div class="card-actions justify-end">
                    <button class="btn btn-primary" @click="handleSubmit">添加</button>
                </div>
            </div>
        </div>
    </VueFinalModal>
    <ConfirmDialog
        v-model="showConfirmDialog"
        :title="'确认删除'"
        :message="'确定要删除这个备份配置吗？此操作不可撤销。'"
        @confirm="handleDeleteConfirm"
    />
</template>

<script setup lang="ts">
import { ref, computed, watch } from "vue";
import { deleteBackup, getBackupList, createBackup } from "../api/backup";
import { BackupItem } from "../models/home";
import { VueFinalModal } from "vue-final-modal";
import { showToast } from '../utils/toast'
import { ErrorCode } from '../config/errorCode'
import ConfirmDialog from '../components/ConfirmDialog.vue'
import { getErrorMessage } from '../utils/error';

// 定义状态
const backupList = ref<BackupItem[]>([]);
const selectAll = ref(false);
const showCard = ref(false);
const showConfirmDialog = ref(false);
const pendingDeleteId = ref<number | null>(null);

// 表单数据
const formData = ref<Omit<BackupItem, 'id' | 'selected'>>({
    dir_name: '',
    file_path: '',
    back_path: ''
});

// 表单错误信息
const errors = ref({
    dirName: '',
    filePath: '',
    backPath: ''
});

// 计算属性
const allSelected = computed(() => {
    return backupList.value.length > 0 && backupList.value.every(item => item.selected);
});

// 表单验证
const validateForm = (): boolean => {
    const { dir_name, file_path, back_path } = formData.value;
    let isValid = true;
    
    errors.value = {
        dirName: '',
        filePath: '',
        backPath: ''
    };

    if (!dir_name.trim()) {
        errors.value.dirName = '文件名不能为空';
        isValid = false;
    }

    if (!file_path.trim()) {
        errors.value.filePath = '文件路径不能为空';
        isValid = false;
    }

    if (!back_path.trim()) {
        errors.value.backPath = '备份路径不能为空';
        isValid = false;
    }

    return isValid;
};

// 重置表单
const resetForm = () => {
    formData.value = {
        dir_name: '',
        file_path: '',
        back_path: ''
    };
};

// 获取备份列表数据
const fetchData = async () => {
    try {
        const response = await getBackupList();
        if (response.status === ErrorCode.SUCCESS) {
            backupList.value = response.data.map((item: BackupItem) => ({ ...item, selected: false }));
        } else {
            showToast.error(response.data.error || response.data.message || '获取数据失败');
        }
    } catch (error) {
        showToast.error(getErrorMessage(error));
    }
};

// 提交表单
const handleSubmit = async () => {
    if (validateForm()) {
        try {
            const response = await createBackup(formData.value);
            if (response.status === ErrorCode.SUCCESS) {
                showToast.success('添加成功');
                await fetchData();
                showCard.value = false;
                resetForm();
            } else {
                showToast.error(response.data.error || response.data.message || '添加失败');
            }
        } catch (error) {
            showToast.error(getErrorMessage(error));
        }
    }
};

// 处理全选/取消全选
const toggleSelectAll = () => {
    backupList.value.forEach(item => {
        item.selected = selectAll.value;
    });
};

// 确认删除
const confirmDelete = (id: number) => {
    pendingDeleteId.value = id;
    showConfirmDialog.value = true;
};

// 执行删除操作
const handleDeleteConfirm = async () => {
    if (pendingDeleteId.value === null) return;
    
    try {
        const response = await deleteBackup(pendingDeleteId.value);
        if (response.status === ErrorCode.SUCCESS) {
            showToast.success('删除成功');
            backupList.value = backupList.value.filter(item => item.id !== pendingDeleteId.value);
        } else {
            showToast.error(response.data.error || response.data.message || '删除失败');
        }
    } catch (error) {
        showToast.error(getErrorMessage(error));
    } finally {
        pendingDeleteId.value = null;
    }
};

// 监听器
watch(allSelected, (newValue) => {
    selectAll.value = newValue;
});

// 始化
fetchData();
</script>
