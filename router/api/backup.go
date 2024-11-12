// router/api/backup.go
package api

import (
	"backup-tool/internal/handler"
	"backup-tool/internal/repository"
	"backup-tool/internal/service/backup"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBackupRouters(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewPathRepository(db)
	service := backup.NewPathRepository(repo)
	handler := handler.NewBackupHandler(service)
	// backup router
	backupRouter := r.Group("/backup")
	{
		backupRouter.POST("/:id", handler.HandleBackup)
	}
}
