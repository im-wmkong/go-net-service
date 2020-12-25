package config

import "github.com/spf13/viper"

var v *viper.Viper

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

// IsSet 是否设置过
func IsSet(key string) bool {
	return v.IsSet(key)
}

// GetString 获取 string 类型的配置信息
func GetString(key string) string {
	return v.GetString(key)
}

// GetInt 获取 int 类型的配置信息
func GetInt(key string) int {
	return v.GetInt(key)
}

// GetBool 获取 bool 类型的配置信息
func GetBool(key string) bool {
	return v.GetBool(key)
}
