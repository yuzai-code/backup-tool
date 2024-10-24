// package handle /internal/handle/backup.go
// 处理备份请求的处理器
package handler

import (
	"backup-tool/internal/service/backup"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BackupHandler 结构体，包含一个 BackupService
type BackupHandler struct {
	backupService backup.BackupService
}

// NewBackupHandler 创建一个新的 BackupHandler 实例
func NewBackupHandler(backupService backup.BackupService) *BackupHandler {
	return &BackupHandler{backupService: backupService}
}

// HandleBackup 处理备份请求
func (b *BackupHandler) HandleBackup(c *gin.Context) {
	// 获取配置备份文件的id
	id := c.Param("id")
	pathID, err := strconv.Atoi(id) // 转换为int类型
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
		return
	}

	// 调用service层的BackupService方法进行备份
	err = b.backupService.BackupService(pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "备份成功"})
}
