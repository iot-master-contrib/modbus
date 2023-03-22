package connect

import (
	"errors"
	"modbus/model"
	"net"
)

// Link 网络连接
type Link struct {
	tunnelBase
	model *model.Link
}

func newLink(client *model.Link, conn net.Conn) *Link {
	return &Link{
		model: client,
		tunnelBase: tunnelBase{
			link: conn,
		}}
}

func (l *Link) Open() error {
	return errors.New("Link cannot open")
}
