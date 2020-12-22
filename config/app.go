package config

import "go-net-service/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "GoNetService"),
	})
}
