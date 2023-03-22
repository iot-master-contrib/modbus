package connect

import (
	"errors"
	"modbus/model"
	"net"
)

// ServerClient 网络连接
type ServerClient struct {
	tunnelBase
	model *model.ServerClient
}

func newServerClient(client *model.ServerClient, conn net.Conn) *ServerClient {
	return &ServerClient{
		model: client,
		tunnelBase: tunnelBase{
			link: conn,
		}}
}

func (l *ServerClient) Open() error {
	return errors.New("ServerClient cannot open")
}
