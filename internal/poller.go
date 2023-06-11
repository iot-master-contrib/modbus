package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iot-master-contrib/modbus/define"
	"github.com/iot-master-contrib/modbus/types"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"io"
	"time"
)

func init() {

	define.RegisterFactory("rtu", func(tunnel define.Conn, opts string) (define.Poller, error) {
		p := &poller{}
		p.modbus = NewRTU(tunnel, opts)
		return p, nil
	})

	define.RegisterFactory("tcp", func(tunnel define.Conn, opts string) (define.Poller, error) {
		p := &poller{}
		p.modbus = NewTCP(tunnel, opts)
		return p, nil
	})

	define.RegisterFactory("parallel-tcp", func(tunnel define.Conn, opts string) (define.Poller, error) {
		p := &poller{}
		p.modbus = NewParallelTCP(tunnel, opts)
		return p, nil
	})
}

type poller struct {
	modbus  Modbus
	devices []*types.Device
}

func (p *poller) Load(tunnel string) error {
	return db.Engine.Where("tunnel_id=?", tunnel).Find(&p.devices)
}

func (p *poller) Poll() bool {
	total := 0

	//TODO 将迭代器提升到p中，单次调用只查询一个设备
	for _, device := range p.devices {
		values := make(map[string]interface{})
		product := Products.Load(device.ProductId)
		if product == nil {
			continue
		}

		//统计加1
		sum := 0

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
			sum++
		}

		if sum > 0 {
			total += sum

			//过滤字段
			for i, c := range product.filters {
				ret, err := c.EvalBool(context.Background(), values)
				if err != nil {
					log.Error(err)
					continue
				}
				if !ret {
					name := product.Filters[i].Name
					if name == "*" {
						//break xx
						//TODO 不上传数据
					} else {
						delete(values, name)
					}
				}
			}

			//计算数据
			for i, c := range product.calculators {
				ret, err := c(context.Background(), values)
				if err != nil {
					log.Error(err)
					continue
				}
				name := product.Filters[i].Name
				values[name] = ret
			}

			//mqtt上传数据，暂定使用Object方式，简单
			topic := fmt.Sprintf("up/property/%s/%s", product.Id, device.Id)
			payload, _ := json.Marshal(values)
			_ = mqtt.Publish(topic, payload, false, 0)

			//上线提醒
			if !device.Online {
				device.Online = true
				topic := fmt.Sprintf("online/%s/%s", device.ProductId, device.Id)
				_ = mqtt.Publish(topic, nil, false, 0)
			}
		} else {
			//掉线提醒
			if device.Online {
				device.Online = false
				topic := fmt.Sprintf("offline/%s/%s", device.ProductId, device.Id)
				_ = mqtt.Publish(topic, nil, false, 0)
			}
		}
	}

	//如果没有设备，就睡眠1分钟
	if total == 0 {
		time.Sleep(time.Second * 5)
		//return errors.New("没有设备")
	}

	return true
}

func (p *poller) Close() error {

	for _, device := range p.devices {
		//mqtt上传数据，暂定使用Object方式，简单
		topic := fmt.Sprintf("offline/%s/%s", device.ProductId, device.Id)
		//payload, _ := json.Marshal(values)
		err := mqtt.Publish(topic, nil, false, 0)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
