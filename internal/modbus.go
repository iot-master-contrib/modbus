package internal

type Modbus interface {
	Read(station uint8, code uint8, addr uint16, size uint16) ([]byte, error)
	Write(station uint8, code uint8, addr uint16, buf []byte) error
}
