package connect

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"modbus/model"
	"net"
	"time"
)

// Client 网络链接
type Client struct {
	tunnelBase
	model *model.Client
}

func NewClient(client *model.Client) *Client {
	return &Client{
		model: client,
	}
}

// Open 打开
func (client *Client) Open() error {
	if client.running {
		return errors.New("client is opened")
	}

	//发起连接
	addr := fmt.Sprintf("%s:%d", client.model.Addr, client.model.Port)
	conn, err := net.Dial(client.model.Net, addr)
	if err != nil {
		client.Retry()
		return err
	}
	client.retry = 0
	client.link = conn

	//启动轮询
	return client.start(&client.model.Tunnel)
}

func (client *Client) Retry() {
	//重连
	retry := &client.model.Retry
	if retry.Enable && (retry.Maximum == 0 || client.retry < retry.Maximum) {
		client.retry++
		client.retryTimer = time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			client.retryTimer = nil
			err := client.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}

// Close 关闭
func (client *Client) Close() error {
	client.running = false

	if client.link != nil {
		link := client.link
		client.link = nil
		return link.Close()
	}
	return errors.New("model is closed")
}
