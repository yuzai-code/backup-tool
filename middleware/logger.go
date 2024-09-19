package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewZapLoggerMiddleware 创建一个基于 zap 的日志中间件
func NewZapLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 计算响应时间
		latency := time.Since(start)

		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// 根据状态码设置日志级别
		switch {
		case status >= 500:
			logger.Error("Server error",
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("clientIP", clientIP),
				zap.Duration("latency", latency),
				zap.String("userAgent", userAgent),
				zap.String("errorMessage", errorMessage),
			)
		case status >= 400:
			logger.Warn("Client error",
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("clientIP", clientIP),
				zap.Duration("latency", latency),
				zap.String("userAgent", userAgent),
				zap.String("errorMessage", errorMessage),
			)
		default:
			logger.Info("Request handled",
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("clientIP", clientIP),
				zap.Duration("latency", latency),
				zap.String("userAgent", userAgent),
			)
		}
	}
}
