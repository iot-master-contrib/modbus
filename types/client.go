package types

import "github.com/zgwit/iot-master/v3/model"

type Client struct {
	Id        string
	Name      string
	Addr      string
	Port      uint16
	Heartbeat string //心跳包
	Period    uint   //采集周期
	Interval  uint
	Created   model.Time
}
