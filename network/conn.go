package network

import (
	"net"
)

type Conn interface {
	ReadMsg() ([]byte, error)
	ReadMsgToLogicServer() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
