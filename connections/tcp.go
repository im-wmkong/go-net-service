package connections

import (
	"go-net-service/pkg/uuid"
	"net"
)

type TcpConnection struct {
	Id string
	Connection *net.TCPConn
}

func NewTcpConnection(Connection *net.TCPConn) TcpConnection {
	return TcpConnection{
		Id: uuid.Generate(),
		Connection: Connection,
	}
}
