import request from './request';

// 获取备份列表
export function getBackupList() {
    return request.get('/path/');
}

// 创建备份
export function createBackup(data) {
    return request.post('/path/', data)
}

// 删除备份
export function deleteBackup(id) {
    return request.delete(`/path/${id}/`)
}
