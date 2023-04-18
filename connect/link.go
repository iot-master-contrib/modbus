package connect

import (
	"errors"
	"modbus/types"
	"net"
)

// Link 网络连接
type Link struct {
	tunnelBase
	model *types.Link
}

func newLink(client *types.Link, conn net.Conn) *Link {
	return &Link{
		model: client,
		tunnelBase: tunnelBase{
			link: conn,
		}}
}

func (l *Link) Open() error {
	return errors.New("Link cannot open")
}
