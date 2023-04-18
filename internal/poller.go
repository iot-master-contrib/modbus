package internal

import (
	"encoding/json"
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"io"
	"modbus/define"
	"modbus/types"
	"time"
)

func init() {

	define.RegisterFactory("rtu", func(tunnel io.ReadWriteCloser, opts string) (define.Poller, error) {
		p := &poller{online: map[string]bool{}, offline: map[string]bool{}}
		p.modbus = NewRTU(tunnel, opts)
		return p, nil
	})

	define.RegisterFactory("tcp", func(tunnel io.ReadWriteCloser, opts string) (define.Poller, error) {
		p := &poller{online: map[string]bool{}, offline: map[string]bool{}}
		p.modbus = NewTCP(tunnel, opts)
		return p, nil
	})

	define.RegisterFactory("parallel-tcp", func(tunnel io.ReadWriteCloser, opts string) (define.Poller, error) {
		p := &poller{online: map[string]bool{}, offline: map[string]bool{}}
		p.modbus = NewParallelTCP(tunnel, opts)
		return p, nil
	})
}

type poller struct {
	modbus  Modbus
	devices []types.Device
	online  map[string]bool
	offline map[string]bool
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

			//mqtt上传数据，暂定使用Object方式，简单
			topic := fmt.Sprintf("up/property/%s/%s", product.Id, device.Id)
			payload, _ := json.Marshal(values)
			err := mqtt.Publish(topic, payload, false, 0)
			if err != nil {
				log.Error(err)
			}

			//上线提醒
			if v, ok := p.online[device.Id]; !v || !ok {
				p.online[device.Id] = true
				topic := fmt.Sprintf("online/%s/%s", device.ProductId, device.Id)
				_ = mqtt.Publish(topic, nil, false, 0)
			}
			p.offline[device.Id] = false
		} else {
			//掉线提醒
			if v, ok := p.offline[device.Id]; !v || !ok {
				p.offline[device.Id] = true
				topic := fmt.Sprintf("offline/%s/%s", device.ProductId, device.Id)
				_ = mqtt.Publish(topic, nil, false, 0)
			}
			p.online[device.Id] = false
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
