package c

import (
	"bufio"
	"os/exec"

	debugger "github.com/Fejiberglibstein/mem-debugger/debugger/pkg"
	"github.com/google/go-dap"
)

type CClient struct {
	debugger.Debugger
	cmd   *exec.Cmd
}

func (c *CClient) Start() error {
	c.cmd = exec.Command("../installed_debuggers/codelldb/extension/adapter/codelldb")
	c.Writer = bufio.NewWriter(c.cmd.Stdout)
	c.Reader = bufio.NewReader(c.cmd.Stdin)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	c.Send(dap.InitializeRequest{
		Arguments: dap.InitializeRequestArguments{
			ClientID:                            "",
			ClientName:                          "",
			AdapterID:                           "",
			Locale:                              "",
			LinesStartAt1:                       false,
			ColumnsStartAt1:                     false,
			PathFormat:                          "",
			SupportsVariableType:                false,
			SupportsVariablePaging:              false,
			SupportsRunInTerminalRequest:        false,
			SupportsMemoryReferences:            false,
			SupportsProgressReporting:           false,
			SupportsInvalidatedEvent:            false,
			SupportsMemoryEvent:                 false,
			SupportsArgsCanBeInterpretedByShell: false,
			SupportsStartDebuggingRequest:       false,
		},
	})

	return nil
}

func (c *CClient) Kill() {
}
