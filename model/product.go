package model

import "github.com/zgwit/iot-master/v3/model"

type Product struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name,omitempty"` //名称
	Desc string `json:"desc,omitempty"` //说明
	//model.Product `xorm:"extends"`
	Mappers []Mapper   `json:"mappers" xorm:"json"`
	Created model.Time `json:"created" xorm:"created"` //创建时间
}

type Mapper struct {
	Code   uint8   `json:"code"`   //指令
	Size   uint16  `json:"size"`   //长度
	Points []Point `json:"points"` //数据点
}

type Point struct {
	Name      string  `json:"name"`           //名称
	Type      string  `json:"type"`           //类型
	Offset    uint16  `json:"offset"`         //偏移
	BigEndian bool    `json:"be,omitempty"`   //大端模式
	Rate      float64 `json:"rate,omitempty"` //倍率
}
