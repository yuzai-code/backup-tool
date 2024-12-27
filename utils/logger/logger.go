// utils/logger/logger.go
package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// InitLogger 初始化带彩色输出的 zap 日志器
func InitLogger(logPath string, level string) {
	// 设置日志级别
	logLevel := zapcore.InfoLevel
	switch level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	// 配置 Console 输出
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",                             // 时间字段的键名
		LevelKey:       "level",                            // 日志级别字段的键名
		NameKey:        "logger",                           // 日志名字段的键名
		CallerKey:      "caller",                           // 调用者字段的键名
		MessageKey:     "msg",                              // 日志消息字段的键名
		StacktraceKey:  "stacktrace",                       // 堆栈跟踪字段的键名
		LineEnding:     zapcore.DefaultLineEnding,          // 换行符
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // 日志级别的编码器
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, // 持续时间的编码器
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 调用者的编码器
	})

	// 配置文件输出
	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",                             // 时间字段的键名
		LevelKey:       "level",                            // 日志级别字段的键名
		NameKey:        "logger",                           // 日志名字段的键名
		CallerKey:      "caller",                           // 调用者字段的键名
		MessageKey:     "msg",                              // 日志消息字段的键名
		StacktraceKey:  "stacktrace",                       // 堆栈跟踪字段的键名
		LineEnding:     zapcore.DefaultLineEnding,          // 换行符
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // 日志级别的编码器
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, // 持续时间的编码器
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 调用者的编码器
	})

	// 配置日志轮转
	writer := &lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    100,     // 每个日志文件最大 100MB
		MaxBackups: 30,      // 最多保留 30 个日志文件
		MaxAge:     7,       // 日志文件最长保存时间 7 天
		Compress:   true,    // 是否压缩
	}

	// 配置输出
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(writer), logLevel),
	)

	// 创建Logger
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Sugar returns sugared logger
// func Sugar() *zap.SugaredLogger {
// 	return Logger.Sugar()
// }
