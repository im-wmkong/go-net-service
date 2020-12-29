package connections

import (
	"go-net-service/pkg/logger"
	"go-net-service/pkg/uuid"
	"net"
)

type TcpConnection struct {
	Connection *net.TCPConn
	ID         string
	MsgChan    chan []byte
}

func NewTcpConnection(Connection *net.TCPConn) *TcpConnection {
	return &TcpConnection{
		Connection: Connection,
		ID:         uuid.Generate(),
		MsgChan:    make(chan []byte),
	}
}

func (tc *TcpConnection) Start() {
	go tc.startReader()
	go tc.startWriter()
}

func (tc *TcpConnection) Stop() {

}

func (tc *TcpConnection) GetTcpConn() *net.TCPConn {
	return tc.Connection
}

func (tc *TcpConnection) GetRemoteAddr() net.Addr {
	return tc.Connection.RemoteAddr()
}

func (tc *TcpConnection) GetConnId() string {
	return tc.ID
}

func (tc *TcpConnection) startReader() {
	logger.Info("Connection Start Reader")
}

func (tc *TcpConnection) startWriter() {
	logger.Info("Connection Start Writer")
}
