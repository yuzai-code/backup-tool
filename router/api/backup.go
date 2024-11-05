// router/api/backup.go
package api

import (
	"backup-tool/internal/handler"
	"backup-tool/internal/repository"
	"backup-tool/internal/service/backup"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupBackupRouters(r *gin.RouterGroup, db *gorm.DB, logger *zap.Logger) {
	repo := repository.NewPathRepository(db)
	service := backup.NewPathRepository(logger, repo)
	handler := handler.NewBackupHandler(service)
	// backup router
	backupRouter := r.Group("/backup")
	{
		backupRouter.POST("/:id", handler.HandleBackup)
	}
}
