package internal

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
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

func (p *Poller) execute() {
	for _, d := range p.devices {
		values := make(map[string]interface{})
		product := Products.Load(d.ProductId)
		for _, m := range product.Mappers {
			read, err := p.modbus.Read(d.Slave, m.Code, m.Addr, m.Size)
			if err != nil {
				log.Error(err)
				continue
			}
			m.Parse(read, values)
		}
		//TODO mqtt
	}
}
