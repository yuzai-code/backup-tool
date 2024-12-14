import { useToast, POSITION } from 'vue-toastification'
import type { PluginOptions as ToastOptions } from 'vue-toastification'

const toast = useToast()

const defaultOptions: ToastOptions = {
    position: POSITION.TOP_RIGHT,
    timeout: 3000,
    closeOnClick: true,
    pauseOnFocusLoss: true,
    pauseOnHover: true,
    draggable: true,
    draggablePercent: 0.6,
    showCloseButtonOnHover: false,
    hideProgressBar: true,
    closeButton: "button",
    icon: true,
    rtl: false
}

export const showToast = {
    success(message: string, options?: ToastOptions) {
        toast.success(message, { ...defaultOptions, ...options })
    },
    error(message: string, options?: ToastOptions) {
        toast.error(message, { ...defaultOptions, ...options })
    },
    warning(message: string, options?: ToastOptions) {
        toast.warning(message, { ...defaultOptions, ...options })
    },
    info(message: string, options?: ToastOptions) {
        toast.info(message, { ...defaultOptions, ...options })
    }
} 