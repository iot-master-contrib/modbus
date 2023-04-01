package internal

import (
	"encoding/json"
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
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
	for _, device := range p.devices {
		values := make(map[string]interface{})
		product := Products.Load(device.ProductId)
		for _, mapper := range product.Mappers {
			read, err := p.modbus.Read(device.Slave, mapper.Code, mapper.Addr, mapper.Size)
			if err != nil {
				log.Error(err)
				continue
			}
			mapper.Parse(read, values)
		}

		//mqtt上传数据，暂定使用Object方式，简单
		topic := fmt.Sprintf("up/property/%s/%s", product.Id, device.Id)
		payload, _ := json.Marshal(values)
		err := mqtt.Publish(topic, payload, false, 0)
		if err != nil {
			log.Error(err)
		}
	}
}
