import request from './request';

// 获取备份列表
export function getBackupList() {
    return request.get('/path/');
}

// 创建备份
export function createBackup() {
    return request.post('/path/')
}
