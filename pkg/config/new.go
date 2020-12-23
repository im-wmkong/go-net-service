package config

import "github.com/spf13/viper"

type Configuration map[string]interface{}

func init() {
	// 1. 初始化 viper 库
	v = viper.New()
	// 2. 设置配置文件名称
	v.SetConfigName("config.json")
	// 3. 配置类型，支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	v.SetConfigType("json")
	// 4. 环境变量配置文件查找的路径，相对于 main.go
	v.AddConfigPath(".")

	// 5. 开始读根目录下的配置文件，读不到会报错
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 6. 设置环境变量前缀，用以区分 Go 的系统环境变量
	v.SetEnvPrefix("viper")
	// 7. v.Get() 时，优先读取环境变量
	v.AutomaticEnv()
}

// Init 配置默认值
func Init(name string, configuration Configuration) {
	for key, value := range configuration {
		key = name + "." + key
		if !IsSet(key) {
			v.SetDefault(key, value)
		}
	}
}

// Env 读取环境变量，支持默认值
func Env(key string, defaultValue ...interface{}) interface{} {
	if !IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return v.Get(key)
}
