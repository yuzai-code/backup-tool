// package handle /internal/handle/backup.go
package handler

import (
	"backup-tool/internal/service/backup"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BackupServiceImpl struct {
	backupService backup.BackupService
}

func (b *BackupServiceImpl) HandleBackup(c *gin.Context) {
	// 获取配置备份文件的id
	id := c.Param("id")
	pathID, err := strconv.Atoi(id) // 转换为int类型
	if err != nil {
		c.JSON(400, gin.H{"error": "无效的id"})
	}
	// 调用service层的BackupSerivce方法进行备份
	err = b.backupService.Backupservice(pathID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}
