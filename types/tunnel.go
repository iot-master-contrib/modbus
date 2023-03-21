package types

import "github.com/zgwit/iot-master/v3/model"

type Tunnel struct {
	Id   string `json:"id,omitempty" xorm:"pk"` //ID
	Name string `json:"name,omitempty"`         //名称
	Desc string `json:"desc,omitempty"`         //说明

	Heartbeat string `json:"heartbeat,omitempty"` //心跳包

	Period   uint `json:"period,omitempty"`   //采集周期
	Interval uint `json:"interval,omitempty"` //采集间隔

	Created model.Time `json:"created" xorm:"created"` //创建时间
}

type Retry struct {
	Enable  bool `json:"enable,omitempty"`
	Minimum uint `json:"minimum,omitempty"`
	Maximum uint `json:"maximum,omitempty"`
	Timeout uint `json:"timeout,omitempty"`
}
