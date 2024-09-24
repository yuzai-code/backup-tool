// main.go
package main

import (
	"backup-tool/config" // 导入 config 包
	"backup-tool/config/db"
	"backup-tool/router"
	"backup-tool/utils"
	"log"
	"os"
)

func main() {
	// 加载环境配置
	env := os.Getenv("GIN_ENV")
	if env == "" {
		env = "dev" // 默认为开发环境
	}

	// 加载配置文件
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 初始化数据库连接
	dbConn, err := db.InitDBConnection()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 正确初始化 zap 日志器
	logger := utils.InitZapLogger() // 从 config 包直接调用 InitZapLogger，而不是从 cfg 结构体

	// 初始化路由
	r := router.InitRouter(dbConn, logger)

	// 启动服务
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
