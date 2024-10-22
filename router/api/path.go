package api

import (
	"backup-tool/internal/handler"
	"backup-tool/internal/repository"
	"backup-tool/internal/service/path"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupPathRoutes 设置路径相关的路由
func SetupPathRoutes(router *gin.RouterGroup, db *gorm.DB) {
	// 创建路径仓库
	pathRepo := repository.NewPathRepository(db)

	// 创建路径服务
	pathService := path.NewPathService(pathRepo)

	// 创建路径处理器
	pathHandler := handler.NewPathHandler(pathService)

	// 设置路由
	pathGroup := router.Group("/path")
	{
		pathGroup.POST("", pathHandler.PathConfig)
		pathGroup.GET("", pathHandler.GetAllPaths)
		pathGroup.DELETE("/:id", pathHandler.DeletePath)
		pathGroup.GET("/:id", pathHandler.GetPathByID)
	}
}
