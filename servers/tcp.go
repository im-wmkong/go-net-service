package servers

import (
	"fmt"
	"go-net-service/connections"
	"go-net-service/pkg/config"
	"go-net-service/pkg/logger"
	"net"
)

type TcpServer struct {
	Host    string
	Port    int
	Network string
}

func NewTcpServer() TcpServer {
	return TcpServer{
		Host:    config.GetString("servers.tcp.host"),
		Port:    config.GetInt("servers.tcp.port"),
		Network: "tcp4",
	}
}

func (s TcpServer) Start() {
	logger.Infof("Starting Tcp Server at Host: %s, Port:%d", s.Host, s.Port)

	// 获取一个tcp addr
	address, err := net.ResolveTCPAddr(s.Network, fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		logger.Errorw("Resolve tcp address error: ", err)
		return
	}

	// 监听服务器地址
	listener, err := net.ListenTCP(s.Network, address)
	if err != nil {
		logger.Errorf("Listen %s error: %s", s.Host, err)
		return
	}

	for {
		connection, err := listener.AcceptTCP()
		if err != nil {
			logger.Errorf("Accept err %s", err)
			continue
		}
		connections.NewTcpConnection(connection)
	}
}
