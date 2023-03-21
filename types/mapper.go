package types

import "github.com/zgwit/iot-master/v3/model"

type Mapper struct {
	Id   string
	Name string
	Desc string

	Code   uint8
	Size   uint16
	Points []Point

	Created model.Time
}

type Point struct {
	Name      string
	Type      string
	Offset    uint16
	BigEndian bool
	Rate      float64
}
