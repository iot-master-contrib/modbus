package connect

import (
	"errors"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"go.bug.st/serial"
	"modbus/model"
	"time"
)

// Serial 串口
type Serial struct {
	tunnelBase
	model *model.Serial
}

func NewSerial(model *model.Serial) *Serial {
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
		BaudRate: int(s.model.Options.BaudRate),
		DataBits: int(s.model.Options.DataBits),
		StopBits: serial.StopBits(s.model.Options.StopBits),
		Parity:   serial.Parity(s.model.Options.ParityMode),
	}

	port, err := serial.Open(s.model.Port, &opts)
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
	return s.start(s.model.Id, s.model.Protocol, s.model.ProtocolOps)
}

func (s *Serial) Retry() {
	retry := &s.model.Retry
	if retry.Enable && (retry.Maximum == 0 || s.retry < retry.Maximum) {
		s.retry++
		s.retryTimer = time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			s.retryTimer = nil
			err := s.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}
