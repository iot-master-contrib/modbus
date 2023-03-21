package types

import "github.com/zgwit/iot-master/v3/model"

type PollerOptions struct {
	Period   uint //采集周期
	Interval uint //采集间隔

}

type Poller struct {
	Id       string
	TunnelId string

	ProductId string //产品ID
	DeviceId  string //子设备号

	Slave   uint8
	Mappers []string

	Created model.Time
}
