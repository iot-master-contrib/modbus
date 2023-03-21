package types

import "github.com/zgwit/iot-master/v3/model"

type Server struct {
	Tunnel     `xorm:"extends"`
	Port       uint16     `json:"port,omitempty"`         //监听端口
	Standalone bool       `json:"standalone,omitempty"`   //单例模式（不支持注册）
	Created    model.Time `json:"created" xorm:"created"` //创建时间
}

type ServerClient struct {
	Tunnel   `xorm:"extends"`
	ServerId string `json:"server_id" xorm:"index"` //服务器ID
	Remote   string `json:"remote,omitempty"`       //远程地址
}
