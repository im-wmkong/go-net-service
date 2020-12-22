package main

import (
	"go-net-service/config"
	"go-net-service/pkg/logger"
	"go-net-service/servers"
)

func init() {
	config.Initialize()
	logger.Initialize()
}

func main() {
	servers.Start()
}
