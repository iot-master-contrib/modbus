package model

import "github.com/zgwit/iot-master/v3/model"

type Tunnel struct {
	Id   string `json:"id,omitempty" xorm:"pk"` //ID
	Name string `json:"name,omitempty"`         //名称
	Desc string `json:"desc,omitempty"`         //说明

	Heartbeat string `json:"heartbeat,omitempty"` //心跳包

	Protocol    string `json:"protocol,omitempty"`     //协议 rtu tcp parallel-tcp
	ProtocolOps string `json:"protocol_ops,omitempty"` //协议参数

	Period   uint `json:"period,omitempty"`   //采集周期
	Interval uint `json:"interval,omitempty"` //采集间隔

	Disabled bool       `json:"disabled"`
	Created  model.Time `json:"created" xorm:"created"` //创建时间
}

type Retry struct {
	Timeout uint `json:"timeout,omitempty"` //重试时间
	Maximum uint `json:"maximum,omitempty"` //最大次数
}
