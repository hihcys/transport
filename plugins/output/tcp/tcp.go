package tcp

import (
	"github.com/luopengift/transport"
	"net"
)

const (
	VERSION = "0.0.1"
)

type TcpOutput struct {
	Addr string `json:"addr"`
}

func NewTcpOutput() *TcpOutput {
	return new(TcpOutput)
}

func (out *TcpOutput) Init(config transport.Configer) error {
	err := config.Parse(out)
	return err
}

func (out *TcpOutput) Write(p []byte) (int, error) {
	conn, err := net.Dial("tcp", out.Addr)
	if err != nil {
		return 0, err
	}
	n, err := conn.Write(p)
	if err != nil {
		return 0, err
	}
	return n, conn.Close()
}

func (out *TcpOutput) Close() error {
	return nil
}

func (out *TcpOutput) Start() error {
	return nil
}

func (out *TcpOutput) Version() string {
	return VERSION
}

func init() {
	transport.RegistOutputer("tcp", NewTcpOutput())
}
