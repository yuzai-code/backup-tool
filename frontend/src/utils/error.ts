import { AxiosError } from 'axios';

export const getErrorMessage = (error: unknown): string => {
    if (error instanceof AxiosError) {
        return error.response?.data?.error 
            || error.response?.data?.message 
            || error.message 
            || '系统错误，请稍后重试';
    }
    
    if (error instanceof Error) {
        return error.message;
    }
    
    return '系统错误，请稍后重试';
}; 