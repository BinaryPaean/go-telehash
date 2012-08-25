package exchange

import (
	"net"
)

type Listener struct {
	port   int
	socket *net.UDPConn
}

func NewListener(port int) *Listener {
	skt, err := net.ListenUDP(net, laddr)
}
