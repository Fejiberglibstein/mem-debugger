package c

import (
	"bufio"
	"os/exec"

	debugger "github.com/Fejiberglibstein/mem-debugger/debugger/pkg"
	"github.com/google/go-dap"
)

type CDebugger struct {
	cmd   *exec.Cmd
	write *bufio.Writer
	read  *bufio.Reader
}

func (c *CDebugger) Start() error {
	c.cmd = exec.Command("../installed_debuggers/codelldb/extension/adapter/codelldb")
	c.write = bufio.NewWriter(c.cmd.Stdout)
	c.read = bufio.NewReader(c.cmd.Stdin)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	c.Send(dap.InitializeRequest{})

	return nil
}

func (c *CDebugger) Send(m any) error {
	return dap.WriteProtocolMessage(c.write, debugger.ConstructRequest(m))
}
func (c *CDebugger) Read() (dap.Message, error) {
	return dap.ReadProtocolMessage(c.read)
}

func (c *CDebugger) Kill() {
}
