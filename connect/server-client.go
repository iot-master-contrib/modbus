package connect

import (
	"errors"
	"modbus/types"
	"net"
)

// ServerClient 网络连接
type ServerClient struct {
	tunnelBase
	model *types.ServerClient
}

func newServerClient(client *types.ServerClient, conn net.Conn) *ServerClient {
	return &ServerClient{
		model: client,
		tunnelBase: tunnelBase{
			link: conn,
		}}
}

func (l *ServerClient) Open() error {
	return errors.New("ServerClient cannot open")
}
