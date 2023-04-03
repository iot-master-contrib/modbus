package define

import (
	"io"
)

// Tunnel 通道
type Tunnel interface {
	io.ReadWriteCloser

	Open() error

	Running() bool

	Online() bool
}
