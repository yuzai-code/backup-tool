// router/router.go
package router

import (
	"backup-tool/router/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitRouter 初始化 Gin 路由，并注册中间件和路由
func InitRouter(db *gorm.DB) *gin.Engine {
	// 创建 Gin 实例
	r := gin.New()

	// 使用 zap 日志中间件
	// r.Use(middleware.NewZapLoggerMiddleware(logger))

	// 使用恢复中间件
	r.Use(gin.Recovery())

	// 将数据库连接传递给所有处理函数
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	// 注册路由组
	apiBase := r.Group("/api")
	// api.RegisterBackupRouter(r.Group("/api"))
	api.SetupUserRoutes(apiBase, db)    // 初始化用户路由
	api.SetupPathRoutes(apiBase, db)    // 初始化路径路由
	api.SetupBackupRouters(apiBase, db) // 初始化路径路由

	return r
}
