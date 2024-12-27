// main.go
package main

import (
	"backup-tool/config"
	"backup-tool/config/db"
	"backup-tool/router"
	"backup-tool/utils/logger"
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist
var distFS embed.FS

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
	logger.InitLogger(cfg.Logger.Path, cfg.Logger.Level)
	defer logger.Log.Sync()

	// 初始化路由
	r := router.InitRouter(dbConn)

	// 添加gzip压缩
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// 获取嵌入的前端文件系统
	frontend, err := fs.Sub(distFS, "frontend/dist")
	if err != nil {
		log.Fatalf("无法加载前端资源: %v", err)
	}

	// 静态文件处理中间件
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/assets/") {
			c.Header("Cache-Control", "public, max-age=31536000")
		}
	})

	// 使用http.FileServer提供静态文件服务
	staticHandler := http.FileServer(http.FS(frontend))
	r.GET("/assets/*filepath", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		staticHandler.ServeHTTP(c.Writer, c.Request)
	})

	// 处理前端路由
	r.NoRoute(func(c *gin.Context) {
		// 如果请求的是API路径，返回404
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// 设置适当的缓存控制
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		// 从嵌入式文件系统中读取index.html
		indexHTML, err := frontend.Open("index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "无法加载页面")
			return
		}
		defer indexHTML.Close()

		// 设置正确的内容类型
		c.Header("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(c.Writer, c.Request, "index.html", time.Now(), indexHTML.(io.ReadSeeker))
	})

	// 启动服务
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
