// config/config.go
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config结构体定义了整个配置文件的结构，包括Server和Database的配置
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
}

// ServerConfig结构体定义了服务器的配置，包括端口号
type ServerConfig struct {
	Port  string // 端口号
	Debug bool   // 调试模式
}

// DatabaseConfig结构体定义了数据库的配置，包括主机地址、端口号、用户名、密码和数据库名称
type DatabaseConfig struct {
	Host       string // 主机地址
	User       string // 用户名
	Password   string // 密码
	Database   string // 数据库名称
	Charset    string // 字符集
	DBFile     string // 数据库文件
	Type       string // 数据库类型
	Port       int    // 端口号
	UnixSocket bool
}

type LoggerConfig struct {
	Level string
	Path  string
}

// LoadConfig 函数根据给定的环境名称加载相应的配置文件，并将其解析为Config结构体。
// 参数:
//   - env: 环境名称，用于确定配置文件的名称。
//
// 返回值:
//   - *Config: 解析后的配置结构体指针。
//   - error: 如果加载或解析配置文件时发生错误，则返回相应的错误信息。
func LoadConfig(env string) (*Config, error) {
	filename := fmt.Sprintf("config/%s.yaml", env)
	config := &Config{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
