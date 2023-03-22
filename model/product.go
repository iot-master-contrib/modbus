package model

type Product struct {
	Id string `json:"id" xorm:"pk"`
	//model.Product `xorm:"extends"`
	Mappers []Mapper `json:"mappers" xorm:"json"`
}

type Mapper struct {
	Code   uint8   `json:"code"` //指令
	Size   uint16  `json:"size"`
	Points []Point `json:"points"`
}

type Point struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Offset    uint16  `json:"offset"`
	BigEndian bool    `json:"be,omitempty"`
	Rate      float64 `json:"rate,omitempty"`
}
