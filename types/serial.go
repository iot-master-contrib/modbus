package types

import "github.com/zgwit/iot-master/v3/model"

type Serial struct {
	Id      string
	Name    string
	Port    string
	Created model.Time
}
