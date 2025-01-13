package debugger

import (
	"github.com/google/go-dap"
)

type Debugger interface {
	Start() error
	Kill() error
	Send(any) error
	Read() (dap.Message, error)
}
