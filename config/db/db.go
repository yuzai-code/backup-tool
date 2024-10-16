// config/db/mysql.go
package db

import (
	"backup-tool/config"
	"backup-tool/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDBConnection 初始化数据库连接
func InitDBConnection() (*gorm.DB, error) {
	// 根据数据库类型选择不同的驱动
	var (
		db         *gorm.DB
		err        error
		confDBType string = config.DatabaseDefaultConfig.Type
	)
	// 配置日志级别
	logLevel := logger.Silent
	if gin.Mode() != gin.ReleaseMode { // 非发布模式输出详细日志
		logLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logLevel,    // 设置日志级别
			Colorful:      true,        // 彩色日志
		},
	)
	// 兼容sqlite3配置
	if confDBType == "sqlite3" {
		confDBType = "sqlite"
	}

	// 如果是测试环境，使用内存数据库
	if gin.Mode() == gin.TestMode {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	} else {
		switch confDBType {
		case "UNSET", "sqlite":
			// 未设置数据库类型，使用SQLite数据库
			db, err = gorm.Open(sqlite.Open(config.DatabaseDefaultConfig.DBFile), &gorm.Config{})
			fmt.Println("数据库使用sqlite")
		case "mysql":
			// 使用MySQL数据库
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.DatabaseDefaultConfig.User,
				config.DatabaseDefaultConfig.Password,
				config.DatabaseDefaultConfig.Host,
				config.DatabaseDefaultConfig.Port,
				config.DatabaseDefaultConfig.Database,
			)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: newLogger,
			})
		default:
			return nil, fmt.Errorf("没有支持的数据库类型: %s", confDBType)
		}
		if err != nil {
			return nil, err
		}
		// 配置连接池参数
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}

		sqlDB.SetMaxOpenConns(100)          // 最大连接数
		sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
		sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

		// 自动迁移数据库结构
		err = db.AutoMigrate(&model.User{})
		err = db.AutoMigrate(&model.Path{})
		if err != nil {
			return nil, err
		}
	}
	return db, err
}
