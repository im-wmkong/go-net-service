package main

import (
	_ "go-net-service/bootstrap"
	"go-net-service/servers"
)

func main() {
	servers.NewServers().Start()
}
