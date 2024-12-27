<template>
    <VueFinalModal v-model="show" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50"
        content-class="bg-base-100 shadow-xl rounded-lg">
        <div class="flex flex-col items-center p-6 max-w-md w-full">
            <div class="p-3 rounded-full bg-red-100">
                <svg class="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
            </div>
            <h3 class="mt-4 text-lg font-medium">
                {{ title }}
            </h3>
            <p class="mt-2 text-sm text-gray-500 text-center">
                {{ message }}
            </p>
            <div class="flex justify-end gap-3 mt-6 w-full">
                <button class="btn btn-ghost" @click="handleCancel">
                    取消
                </button>
                <button class="btn btn-error" @click="handleConfirm">
                    确认删除
                </button>
            </div>
        </div>
    </VueFinalModal>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { VueFinalModal } from 'vue-final-modal'

interface Props {
    modelValue: boolean
    title?: string
    message?: string
}

const props = withDefaults(defineProps<Props>(), {
    title: '确认删除',
    message: '你确定要删除这个项目吗？此操作不可撤销。'
})

const emit = defineEmits<{
    (e: 'update:modelValue', value: boolean): void
    (e: 'confirm'): void
    (e: 'cancel'): void
}>()

const show = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
})

const handleConfirm = () => {
    emit('confirm')
    emit('update:modelValue', false)
}

const handleCancel = () => {
    emit('cancel')
    emit('update:modelValue', false)
}
</script>
