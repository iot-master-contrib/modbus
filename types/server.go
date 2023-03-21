package types

import "github.com/zgwit/iot-master/v3/model"

type Server struct {
	Id          string
	Name        string
	Port        uint16
	Heartbeat   string //心跳包
	Registrable bool   //自动注册
	Created     model.Time
}

type ServerClient struct {
	SN       string //sn
	ServerId string
	Created  model.Time
}
