package types

import "github.com/zgwit/iot-master/v3/model"

type Poller struct {
	Id       string
	TunnelId string

	Slave uint8

	ProductId string //产品ID
	DeviceId  string //子设备号

	Created model.Time
}
