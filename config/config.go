package config

import "fmt"

// Initialize 配置信息初始化
// 触发加载本目录下其他文件中的 init 方法
func Initialize() {
	fmt.Println("Initialize config")
}
