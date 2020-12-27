package servers

type Servers struct {
	Tcp *TcpServer
	Ws  *WsServer
}

func NewServers() *Servers {
	return &Servers{
		Tcp: NewTcpServer(),
		Ws:  NewWsServer(),
	}
}

func (s *Servers) Start() {
	s.Tcp.Start()
	s.Ws.Start()
}
