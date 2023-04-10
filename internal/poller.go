package internal

import (
	"encoding/json"
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"io"
	"modbus/define"
	"modbus/model"
	"time"
)

func init() {

	define.RegisterFactory("rtu", func(tunnel io.ReadWriteCloser, opts string) (define.Poller, error) {
		p := &poller{}
		p.modbus = NewRTU(tunnel, opts)
		return p, nil
	})

	define.RegisterFactory("tcp", func(tunnel io.ReadWriteCloser, opts string) (define.Poller, error) {
		p := &poller{}
		p.modbus = NewTCP(tunnel, opts)
		return p, nil
	})

	define.RegisterFactory("parallel-tcp", func(tunnel io.ReadWriteCloser, opts string) (define.Poller, error) {
		p := &poller{}
		p.modbus = NewParallelTCP(tunnel, opts)
		return p, nil
	})
}

type poller struct {
	modbus  Modbus
	devices []model.Device
}

func (p *poller) Load(tunnel string) error {
	return db.Engine.Where("tunnel_id=?", tunnel).Find(&p.devices)
}

func (p *poller) Poll() bool {

	cnt := 0

	//TODO 将迭代器提升到p中，单次调用只查询一个设备
	for _, device := range p.devices {
		values := make(map[string]interface{})
		product := Products.Load(device.ProductId)
		if product == nil {
			continue
		}

		//统计加1
		cnt++
		cnt2 := 0

		for _, mapper := range product.Mappers {
			r, e := p.modbus.Read(device.Slave, mapper.Code, mapper.Addr, mapper.Size)
			if e != nil {
				//连接关闭就退出
				if e == io.EOF {
					return false
				}

				log.Error(e)
				continue
			}
			mapper.Parse(r, values)
			cnt2++
		}

		if cnt2 > 0 {
			//mqtt上传数据，暂定使用Object方式，简单
			topic := fmt.Sprintf("up/property/%s/%s", product.Id, device.Id)
			payload, _ := json.Marshal(values)
			err := mqtt.Publish(topic, payload, false, 0)
			if err != nil {
				log.Error(err)
			}
		}
	}

	//如果没有设备，就睡眠1分钟
	if cnt == 0 {
		time.Sleep(time.Minute)
		//return errors.New("没有设备")
	}

	return true
}

func (p *poller) Close() error {

	for _, device := range p.devices {
		//mqtt上传数据，暂定使用Object方式，简单
		topic := fmt.Sprintf("up/event/%s/%s", device.ProductId, device.Id)
		//payload, _ := json.Marshal(values)
		err := mqtt.Publish(topic, nil, false, 0)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
