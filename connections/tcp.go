package connections

import "net"

type TcpConnection struct {
	Connection *net.TCPConn
}

func NewTcpConnection(Connection *net.TCPConn) TcpConnection {
	return TcpConnection{Connection: Connection}
}
