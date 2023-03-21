package types

import "github.com/jacobsa/go-serial/serial"

type Serial struct {
	Tunnel             `xorm:"extends"`
	Retry              `xorm:"extends"`
	serial.OpenOptions `xorm:"extends"`
}
