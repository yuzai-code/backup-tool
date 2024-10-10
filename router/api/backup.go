// router/api/backup.go
package api

import (
	"backup-tool/internal/service"

	"github.com/gin-gonic/gin"
)

func BackupRouter(r *gin.RouterGroup) {
	// backup router
	backupRouter := r.Group("/backup")
	{
		backupRouter.POST("/", service.HandleBackup)
	}
}
