package connect

import (
	"errors"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"go.bug.st/serial"
	"modbus/types"
	"time"
)

// Serial 串口
type Serial struct {
	tunnelBase
	model *types.Serial
}

func NewSerial(model *types.Serial) *Serial {
	return &Serial{
		model: model,
	}
}

// Open 打开
func (s *Serial) Open() error {
	if s.running {
		return errors.New("serial is opened")
	}

	opts := serial.Mode{
		BaudRate: int(s.model.BaudRate),
		DataBits: int(s.model.DataBits),
		StopBits: serial.StopBits(s.model.StopBits),
		Parity:   serial.Parity(s.model.ParityMode),
	}

	port, err := serial.Open(s.model.PortName, &opts)
	if err != nil {
		//TODO 串口重试
		s.Retry()
		return err
	}

	//读超时
	err = port.SetReadTimeout(time.Second * 5)
	if err != nil {
		return err
	}

	s.running = true
	s.online = true
	s.link = port
	s.retry = 0

	//清空重连计数
	s.retry = 0

	//启动轮询
	return s.start(&s.model.Tunnel)
}

func (s *Serial) Retry() {
	retry := &s.model.Retry
	if retry.RetryMaximum == 0 || s.retry < retry.RetryMaximum {
		s.retry++
		timeout := retry.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		s.retryTimer = time.AfterFunc(time.Second*time.Duration(timeout), func() {
			s.retryTimer = nil
			err := s.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}
