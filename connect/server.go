package connect

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"modbus/model"
	"net"
)

// Server TCP服务器
type Server struct {
	model *model.Server

	children map[string]*ServerClient

	listener *net.TCPListener

	running bool
}

func NewServer(model *model.Server) *Server {
	s := &Server{
		model:    model,
		children: make(map[string]*ServerClient),
	}
	return s
}

// Open 打开
func (s *Server) Open() error {
	if s.running {
		return errors.New("s is opened")
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.model.Port))
	if err != nil {
		return err
	}
	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	s.running = true
	go func() {
		for {
			c, err := s.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			//单例模式，关闭之前的连接
			if s.model.Standalone {
				const k = "internal"
				if cc, ok := s.children[k]; ok {
					_ = cc.Close()
				}
				s.children[k] = newServerClient(&model.ServerClient{
					Tunnel:   s.model.Tunnel,
					ServerId: s.model.Id,
					Remote:   c.RemoteAddr().String(),
				}, c)
				continue
			}

			buf := make([]byte, 128)
			n := 0
			n, err = c.Read(buf)
			if err != nil {
				_ = c.Close()
				continue
			}
			data := buf[:n]
			sn := string(data)

			var client model.ServerClient
			//get, err := db.Engine.Where("server_id=?", s.model.Id).And("sn=?", sn).Get(&client)
			get, err := db.Engine.ID(sn).Get(&client)
			if err != nil {
				_, _ = c.Write([]byte(err.Error()))
				_ = c.Close()
				return
			}
			if !get {
				client = model.ServerClient{
					Tunnel:   s.model.Tunnel,
					ServerId: s.model.Id,
					Remote:   c.RemoteAddr().String(),
				}
				client.Id = sn
				_, err := db.Engine.InsertOne(&client)
				if err != nil {
					_, _ = c.Write([]byte(err.Error()))
					_ = c.Close()
					return
				}
			}

			tnl := newServerClient(&client, c)
			s.children[sn] = tnl
		}

		s.running = false
	}()

	return nil
}

// Close 关闭
func (s *Server) Close() (err error) {
	//close tunnels
	if s.children != nil {
		for _, l := range s.children {
			_ = l.Close()
		}
	}
	return s.listener.Close()
}

// GetTunnel 获取连接
func (s *Server) GetTunnel(id string) Tunnel {
	return s.children[id]
}

func (s *Server) Running() bool {
	return s.running
}
