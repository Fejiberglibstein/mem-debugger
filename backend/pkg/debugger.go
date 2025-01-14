package debugger

import (
	"bufio"
	"reflect"

	"github.com/google/go-dap"
)

type Debugger struct {
	Reader *bufio.Reader
	Writer *bufio.Writer

	OnEvent func(dap.Message)
}

type WaitFor func(req dap.Message, res dap.Message) bool
type Client interface {
	Start() error
	Kill() error

	SendMessage(dap.RequestMessage) error
	ReadMessage() (dap.Message, error)
	SendAndWait(dap.RequestMessage, reflect.Type) (dap.Message, error)
}

func (c *Debugger) SendMessage(m dap.RequestMessage) error {
	return dap.WriteProtocolMessage(c.Writer, ConstructRequest(m))
}
func (c *Debugger) ReadMessage() (dap.Message, error) {
	return dap.ReadProtocolMessage(c.Reader)
}

func (c *Debugger) SendAndWait(req dap.RequestMessage, desiredType reflect.Type) (dap.Message, error) {
	c.SendMessage(req)
	return c.WaitFor(desiredType)
}

func (c *Debugger) WaitFor(desiredType reflect.Type) (dap.Message, error) {
	for {
		resp, err := c.ReadMessage()
		if err != nil {
			return nil, err
		}

		if reflect.TypeOf(resp) == desiredType {
			return resp, nil
		}
	}
}
