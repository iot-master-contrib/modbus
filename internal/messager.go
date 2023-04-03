package internal

import (
	"errors"
	"io"
	"sync"
	"time"
)

type Messenger struct {
	Timeout time.Duration
	mu      sync.Mutex
	Conn    io.ReadWriter
}

func (m *Messenger) Ask(request []byte, response []byte) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//先写
	_, err := m.Conn.Write(request)
	if err != nil {
		return 0, err
	}

	//接收返回值
	n := make(chan int)
	e := make(chan error)

	//TODO 此处开启了新协程，不太好
	go func() {
		//读
		nn, ee := m.Conn.Read(response)
		if err != nil {
			return
		}
		if ee != nil {
			e <- ee
			return
		}
		n <- nn
	}()

	select {
	case <-time.After(m.Timeout):
		return 0, errors.New("timeout")
	case nn := <-n:
		return nn, nil
	case ee := <-e:
		return 0, ee
	}
}

func (m *Messenger) AskAtLeast(request []byte, response []byte, min int) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//先写
	_, err := m.Conn.Write(request)
	if err != nil {
		return 0, err
	}

	//接收返回值
	n := make(chan int)
	e := make(chan error)

	//TODO 此处开启了新协程，不太好
	go func() {
		//读
		nn, ee := io.ReadAtLeast(m.Conn, response, min)
		if err != nil {
			return
		}
		if ee != nil {
			e <- ee
			return
		}
		n <- nn
	}()

	select {
	case <-time.After(m.Timeout):
		return 0, errors.New("timeout")
	case nn := <-n:
		return nn, nil
	case ee := <-e:
		return 0, ee
	}
}
