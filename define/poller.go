package define

import (
	"fmt"
	"io"
)

type Poller interface {
	Load(tunnel string) error
	Poll() bool
	Close() error
}

type Factory func(link io.ReadWriteCloser, opts string) (Poller, error)

//
//type Factory interface {
//	Create(define Tunnel, protocol string, opts string) (Poller, error)
//}

var factories = map[string]Factory{}

func RegisterFactory(protocol string, factory Factory) {
	factories[protocol] = factory
}

func CreatePoller(link io.ReadWriteCloser, protocol string, opts string) (Poller, error) {
	if f, ok := factories[protocol]; ok {
		return f(link, opts)
	}
	return nil, fmt.Errorf("unkown protocol %s", protocol)
}
