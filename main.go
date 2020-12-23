package main

import (
	"go-net-service/config"
	"go-net-service/servers"
)

func init() {
	config.Initialize()
}

func main() {
	servers.Start()
}
