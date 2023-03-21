package types

import "github.com/zgwit/iot-master/v3/model"

type Product struct {
	//Id   string `json:"id" xorm:"pk"`
	model.Product `xorm:"extends"`
	Mappers       []Mapper `json:"mappers" xorm:"json"`
}

type Mapper struct {
	Code   uint8 //指令
	Size   uint16
	Points []Point
}

type Point struct {
	Name      string
	Type      string
	Offset    uint16
	BigEndian bool
	Rate      float64
}
