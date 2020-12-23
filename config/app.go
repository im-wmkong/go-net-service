package config

import "go-net-service/pkg/config"

func init() {
	config.Init("app", config.Configuration{
		// 应用名称，暂时没有使用到
		"name": config.Env("app.name", "GoNetService"),
	})
}
