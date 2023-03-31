package model

import (
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/bin"
)

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
	Addr   uint16  `json:"addr"`   //地址
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

func (m *Mapper) Parse(buf []byte, ret map[string]interface{}) {
	l := uint16(len(buf))
	for _, p := range m.Points {
		if p.Offset > l {
			continue
		}
		switch p.Type {
		case "word":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint16(buf[p.Offset:])
			} else {
				ret[p.Name] = bin.ParseUint16LittleEndian(buf[p.Offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint16)) * p.Rate
			}
		case "qword":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint32(buf[p.Offset:])
			} else {
				ret[p.Name] = bin.ParseUint32LittleEndian(buf[p.Offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint16)) * p.Rate
			}
		case "float":
			if p.BigEndian {
				ret[p.Name] = bin.ParseFloat32(buf[p.Offset:])
			} else {
				ret[p.Name] = bin.ParseFloat32LittleEndian(buf[p.Offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(float32)) * p.Rate
			}
		case "double":
			if p.BigEndian {
				ret[p.Name] = bin.ParseFloat64(buf[p.Offset:])
			} else {
				ret[p.Name] = bin.ParseFloat64LittleEndian(buf[p.Offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(float64)) * p.Rate
			}

		}
	}
}
