package internal

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"io"
	"modbus/model"
)

func CreatePoller(typ string, link io.ReadWriteCloser, opts string) (*Poller, error) {
	p := &Poller{}
	switch typ {
	case "modbus-rtu":
		p.modbus = NewRTU(link, opts)
	case "modbus-tcp":
		p.modbus = NewTCP(link, opts)
	case "parallel-tcp":
		p.modbus = NewParallelTCP(link, opts)
	default:
		return nil, fmt.Errorf("未知类型 %s", typ)
	}
	return p, nil
}

var Products lib.Map[model.Product]

type Poller struct {
	devices []model.Device
	modbus  Modbus
}
