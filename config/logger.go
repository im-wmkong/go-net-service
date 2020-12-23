package config

import "go-net-service/pkg/config"

func init() {
	config.Init("logger", config.Configuration{
		// 日志级别
		"level": config.Env("logger.level", "debug"),
		// 每个日志文件保存的最大尺寸 单位：M
		"max_size": config.Env("logger.max_size", 256),
		// 日志文件最多保存多少个备份
		"max_backup": config.Env("logger.max_backup", 10),
		// 文件最多保存多少天
		"max_age": config.Env("logger.max_age", 7),
	})
}
