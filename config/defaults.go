package config

// DatabaseDefaultConfig 是默认的数据库配置
var DatabaseDefaultConfig = &DatabaseConfig{
	Port:       3306,
	Type:       "UNSET",
	Charset:    "utf8",
	DBFile:     "backup.db",
	UnixSocket: false,
}
