package debugger

import (
	"bufio"

	"github.com/google/go-dap"
)

type Debugger struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

type Client interface {
	Start() error
	Kill() error
	Send(any) error
	Read() (dap.Message, error)
}

func (c *Debugger) Send(m any) error {
	return dap.WriteProtocolMessage(c.Writer, ConstructRequest(m))
}
func (c *Debugger) Read() (dap.Message, error) {
	return dap.ReadProtocolMessage(c.Reader)
}
