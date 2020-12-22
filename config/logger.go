package config

import "go-net-service/pkg/config"

func init() {
	config.Add("logger", config.StrMap{
		// 日志级别
		"level": config.Env("LOG_LEVEL", "debug"),
		// 每个日志文件保存的最大尺寸 单位：M
		"max_size": config.Env("LOG_MAX_SIZE", 256),
		// 日志文件最多保存多少个备份
		"max_backup": config.Env("LOG_MAX_BACKUP", 10),
		// 文件最多保存多少天
		"max_age": config.Env("LOG_MAX_AGE", 7),
	})
}
