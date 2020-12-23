package config

import "github.com/spf13/viper"

var v *viper.Viper

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
