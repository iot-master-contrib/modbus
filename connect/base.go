package connect

import (
	"errors"
	"io"
	"modbus/define"
	"sync"
	"time"
)

func (l *tunnelBase) start(id, protocol, opts string) (err error) {
	l.poller, err = define.CreatePoller(l, protocol, opts)
	if err != nil {
		return
	}

	//加载设备
	err = l.poller.Load(id)
	if err != nil {
		return
	}

	//开启线程，在回调中完成一次询问
	go func() {
		l.running = true

		for {
			if !l.running {
				break
			}

			//轮询
			if !l.poller.Poll() {
				break
			}
		}

		l.running = false
	}()

	return
}

type tunnelBase struct {
	link   io.ReadWriteCloser
	poller define.Poller

	lock sync.Mutex

	running bool
	online  bool

	retry      uint
	retryTimer *time.Timer

	//透传
	pipe io.ReadWriteCloser
}

func (l *tunnelBase) Running() bool {
	return l.running
}

func (l *tunnelBase) Online() bool {
	return l.online
}

// Close 关闭
func (l *tunnelBase) Close() error {
	if l.retryTimer != nil {
		l.retryTimer.Stop()
	}
	if !l.running {
		return errors.New("tunnel closed")
	}

	l.onClose()
	return l.link.Close()
}

func (l *tunnelBase) onClose() {
	l.running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
}

// Write 写
func (l *tunnelBase) Write(data []byte) (int, error) {
	if !l.running {
		return 0, errors.New("model closed")
	}
	if l.pipe != nil {
		return 0, nil //透传模式下，直接抛弃
	}
	return l.link.Write(data)
}

// Read 读
func (l *tunnelBase) Read(data []byte) (int, error) {
	if !l.running {
		return 0, errors.New("model closed")
	}

	//if port, ok := l.link.(serial.Port); ok {
	//	_ = port.SetReadTimeout(time.Second)
	//}
	//
	//if conn, ok := l.link.(net.Conn); ok {
	//	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	//}

	if l.pipe != nil {
		//TODO 先read，然后透传
		return 0, nil //透传模式下，直接抛弃
	}
	return l.link.Read(data)
}

func (l *tunnelBase) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}

	l.pipe = pipe
	//传入空，则关闭
	if pipe == nil {
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := pipe.Read(buf)
		if err != nil {
			//if err == io.EOF {
			//	continue
			//}
			//pipe关闭，则不再透传
			break
		}
		//将收到的数据转发出去
		n, err = l.link.Write(buf[:n])
		if err != nil {
			//发送失败，说明连接失效
			_ = pipe.Close()
			break
		}
	}
	l.pipe = nil

	//TODO 使用io.copy
	//go io.Copy(pipe, l.link)
	//go io.Copy(l.link, pipe)
}
