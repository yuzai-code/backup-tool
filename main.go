package main

import (
	"backup-tool/config"
	"backup-tool/config/db"
	"backup-tool/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	env := os.Getenv("GIN_ENV")
	if env == "" {
		env = "dev" // 默认为开发环境
	}

	config, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 初始化数据库链接
	db, err := db.NewDBConnection(config)
	if err != nil {
		log.Fatalf("链接数据库失败: %v", err)
	}

	// 创建 zap 日志实例
	logger, err := zap.NewProduction() // 使用生产配置，如果是开发环境，可以用 zap.NewDevelopment()
	if err != nil {
		log.Fatalf("初始化 zap 日志器失败: %v", err)
	}
	defer logger.Sync() // 确保日志缓冲区被刷新

	// 创建 Gin 实例
	r := gin.New()

	// 使用 zap 日志中间件
	r.Use(middleware.NewZapLoggerMiddleware(logger))

	// 使用恢复中间件
	r.Use(gin.Recovery())

	// 将数据库连接传递给需要的处理函数或中间件
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	// 启动服务
	r.Run(config.Server.Port) // 默认在 0.0.0.0:8080
}
