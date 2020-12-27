package servers

import "go-net-service/pkg/logger"

type WsServer struct {
}

func NewWsServer() *WsServer {
	return &WsServer{}
}

func (s *WsServer) Start() {
	logger.Info("Starting WebSocket Server")
}
