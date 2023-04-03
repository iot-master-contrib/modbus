package model

type Serial struct {
	Tunnel  `xorm:"extends"`
	Port    string        `json:"port"` //port, e.g. COM1 "/dev/ttySerial1".
	Options SerialOptions `json:"options" xorm:"json"`
	Retry   Retry         `json:"retry" xorm:"json"`
}

type SerialOptions struct {
	BaudRate   uint `json:"baud_rate,omitempty"`   //9600 115200
	DataBits   uint `json:"data_bits,omitempty"`   //5 6 7 8
	StopBits   uint `json:"stop_bits,omitempty"`   //1 2
	ParityMode int  `json:"parity_mode,omitempty"` //0 1 2 NONE ODD EVEN
}
