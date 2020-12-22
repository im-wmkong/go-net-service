package servers

import "go-net-service/pkg/logger"

func Start() {
	logger.Debug("Start all servers")
	TcpServer{}.Start()
	WsServer{}.Start()
}
