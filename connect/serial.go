package connect

import (
	"errors"
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/v3/pkg/log"
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

	opts := serial.OpenOptions{
		PortName:              s.model.Options.PortName,
		BaudRate:              s.model.Options.BaudRate,
		DataBits:              s.model.Options.DataBits,
		StopBits:              s.model.Options.StopBits,
		ParityMode:            serial.ParityMode(s.model.Options.ParityMode),
		InterCharacterTimeout: s.model.Options.InterCharacterTimeout,
		MinimumReadSize:       s.model.Options.MinimumReadSize,
		Rs485Enable:           s.model.Options.Rs485Enable,
	}

	port, err := serial.Open(opts)
	if err != nil {
		//TODO 串口重试
		s.Retry()
		return err
	}
	s.running = true
	s.online = true
	s.link = port
	s.retry = 0

	//清空重连计数
	s.retry = 0

	return nil
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
