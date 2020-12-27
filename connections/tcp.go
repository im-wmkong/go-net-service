package connections

import (
	"go-net-service/pkg/uuid"
	"net"
)

type TcpConnection struct {
	Connection *net.TCPConn
	Id         string
	MsgChan    chan []byte
}

func NewTcpConnection(Connection *net.TCPConn) *TcpConnection {
	return &TcpConnection{
		Connection: Connection,
		Id:         uuid.Generate(),
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
	return tc.Id
}

func (tc *TcpConnection) startReader() {

}

func (tc *TcpConnection) startWriter() {

}
