// config/logger.go
package config

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZapLogger 初始化带彩色输出的 zap 日志器
func InitZapLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()                                // 使用开发环境配置
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 启用彩色日志输出
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 格式化时间
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder      // 短路径格式的调用者

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("初始化 zap 日志器失败: %v", err)
	}

	return logger
}
