package types

import "github.com/zgwit/iot-master/v3/model"

type Device struct {
	Id string `json:"id" xorm:"pk"`

	Name string `json:"name,omitempty"` //名称
	Desc string `json:"desc,omitempty"` //说明

	TunnelId  string `json:"tunnel_id"`  //通道
	ProductId string `json:"product_id"` //产品ID

	Slave    uint8 `json:"slave"` //从站号
	Disabled bool  `json:"disabled"`

	Created model.Time `json:"created" xorm:"created"`
}
