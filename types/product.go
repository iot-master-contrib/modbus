package types

import (
	"github.com/zgwit/iot-master/v3/pkg/bin"
	"github.com/zgwit/iot-master/v3/pkg/convert"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"time"
)

type Product struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name,omitempty"` //名称
	Desc string `json:"desc,omitempty"` //说明
	//model.Product `xorm:"extends"`
	Mappers     []Mapper     `json:"mappers" xorm:"json"`
	Filters     []Filter     `json:"filters" xorm:"json"`
	Calculators []Calculator `json:"calculators" xorm:"json"`
	Created     time.Time    `json:"created" xorm:"created"` //创建时间
}

type Filter struct {
	Name       string `json:"name"`       //字段
	Expression string `json:"expression"` //表达式
	//Entire     bool   `json:"entire"`
}

type Calculator struct {
	Name       string `json:"name"`       //赋值
	Expression string `json:"expression"` //表达式
}

type Mapper struct {
	Code   uint8   `json:"code"`   //指令
	Addr   uint16  `json:"addr"`   //地址
	Size   uint16  `json:"size"`   //长度
	Points []Point `json:"points"` //数据点
}

type Point struct {
	Name      string  `json:"name"`            //名称
	Type      string  `json:"type"`            //类型
	Offset    uint16  `json:"offset"`          //偏移
	Bits      uint16  `json:"bits,omitempty"`  //位，1 2 3...
	BigEndian bool    `json:"be,omitempty"`    //大端模式
	Order     string  `json:"order,omitempty"` //字序：abcd badc cdab dcba
	Rate      float64 `json:"rate,omitempty"`  //倍率
}

// reorder 根据字序重排前 n 字节，返回重排后的字节切片。
// 仅对多字节类型（32位/64位）有意义，n 为字节长度（4 或 8）。
// 四种字序：
//   - abcd：大端字序 + 大端字节序（默认，不重排）
//   - badc：大端字序 + 小端字节序（字内字节交换）
//   - cdab：小端字序 + 大端字节序（字序交换）
//   - dcba：小端字序 + 小端字节序（完全反转）
//
// 当 order 为空时，回退到 BigEndian 字段的旧语义（false 等价于 dcba）。
func (p *Point) reorder(buf []byte, n int) []byte {
	byteSwap := false
	wordSwap := false
	switch p.Order {
	case "badc":
		byteSwap = true
	case "cdab":
		wordSwap = true
	case "dcba":
		byteSwap = true
		wordSwap = true
	default:
		if p.Order == "" && !p.BigEndian {
			byteSwap = true
			wordSwap = true
		}
	}

	if !byteSwap && !wordSwap {
		if len(buf) >= n {
			return buf[:n]
		}
		return buf
	}

	b := make([]byte, n)
	copy(b, buf[:n])
	if byteSwap {
		for i := 0; i+1 < n; i += 2 {
			b[i], b[i+1] = b[i+1], b[i]
		}
	}
	if wordSwap {
		for i, j := 0, n-2; i < j; i, j = i+2, j-2 {
			b[i], b[j] = b[j], b[i]
			b[i+1], b[j+1] = b[j+1], b[i+1]
		}
	}
	return b
}

func (m *Mapper) Encode(name string, data any) []byte {
	//TODO 支持 布尔，数值，数组
	for _, p := range m.Points {
		if p.Name == name {
			continue
		}

		switch m.Code {
		case 1:
			//convert.ToBool(data) 太范了
			val, ok := data.(bool)
			if !ok {
				//TODO error
				return nil
			}
			if val {
				return []byte{0xFF, 00}
			} else {
				return []byte{0x00, 00}
			}
		case 3:
			var ret []byte

			//倍率逆转换
			if p.Rate != 0 && p.Rate != 1 {
				val, ok := data.(float64)
				if ok {
					data = val / p.Rate
				} else {
					//error ?
				}
			}

			switch p.Type {
			case "short", "int16":
				ret = make([]byte, 2)
				val := convert.ToInt16(data)
				if p.BigEndian {
					bin.WriteUint16(ret, uint16(val))
				} else {
					bin.WriteUint16LittleEndian(ret, uint16(val))
				}
			case "word", "uint16":
				ret = make([]byte, 2)
				val := convert.ToUint16(data)
				if p.BigEndian {
					bin.WriteUint16(ret, val)
				} else {
					bin.WriteUint16LittleEndian(ret, val)
				}
			case "int32", "int":
				ret = make([]byte, 4)
				val := convert.ToInt32(data)
				bin.WriteUint32(ret, uint32(val))
				ret = p.reorder(ret, 4)
			case "qword", "uint32", "uint":
				ret = make([]byte, 4)
				val := convert.ToUint32(data)
				bin.WriteUint32(ret, val)
				ret = p.reorder(ret, 4)
			case "int64":
				ret = make([]byte, 8)
				val := convert.ToInt64(data)
				bin.WriteUint64(ret, uint64(val))
				ret = p.reorder(ret, 8)
			case "uint64":
				ret = make([]byte, 8)
				val := convert.ToUint64(data)
				bin.WriteUint64(ret, val)
				ret = p.reorder(ret, 8)
			case "float", "float32":
				ret = make([]byte, 4)
				val := convert.ToFloat32(data)
				bin.WriteFloat32(ret, val)
				ret = p.reorder(ret, 4)
			case "double", "float64":
				ret = make([]byte, 8)
				val := convert.ToFloat64(data)
				bin.WriteFloat64(ret, val)
				ret = p.reorder(ret, 8)
			}

			return ret
		}
	}

	return nil
}

func (m *Mapper) Parse(buf []byte, ret map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	l := uint16(len(buf))

	//识别位
	if m.Code == 1 || m.Code == 2 {
		bytes := bin.ExpandBool(buf, int(m.Size))
		l = uint16(len(bytes))
		for _, p := range m.Points {
			offset := p.Offset
			if offset >= l {
				continue
			}
			ret[p.Name] = bytes[p.Offset] > 0
		}
		return
	}

	//解析16位
	for _, p := range m.Points {
		//offset := p.Offset * 2
		offset := p.Offset << 1
		if offset >= l {
			continue
		}
		switch p.Type {
		case "bit", "bool", "boolean":
			var v uint16
			if p.BigEndian {
				v = bin.ParseUint16(buf[offset:])
			} else {
				v = bin.ParseUint16LittleEndian(buf[offset:])
			}
			ret[p.Name] = 1<<(p.Bits-1)&v != 0
		case "short", "int16":
			if p.BigEndian {
				ret[p.Name] = int16(bin.ParseUint16(buf[offset:]))
			} else {
				ret[p.Name] = int16(bin.ParseUint16LittleEndian(buf[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int16)) * p.Rate
			}
		case "word", "uint16":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint16(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseUint16LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint16)) * p.Rate
			}
		case "int32", "int":
			ret[p.Name] = int32(bin.ParseUint32(p.reorder(buf[offset:], 4)))
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int32)) * p.Rate
			}
		case "qword", "uint32", "uint":
			ret[p.Name] = bin.ParseUint32(p.reorder(buf[offset:], 4))
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint32)) * p.Rate
			}
		case "int64":
			ret[p.Name] = int64(bin.ParseUint64(p.reorder(buf[offset:], 8)))
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int64)) * p.Rate
			}
		case "uint64":
			ret[p.Name] = bin.ParseUint64(p.reorder(buf[offset:], 8))
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint64)) * p.Rate
			}
		case "float", "float32":
			ret[p.Name] = bin.ParseFloat32(p.reorder(buf[offset:], 4))
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(float32)) * p.Rate
			}
		case "double", "float64":
			ret[p.Name] = bin.ParseFloat64(p.reorder(buf[offset:], 8))
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = ret[p.Name].(float64) * p.Rate
			}
		}
	}
}
