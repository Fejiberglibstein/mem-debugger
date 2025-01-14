package c_debugger

import (
	"bufio"
	"fmt"
	"os/exec"
	"reflect"

	debugger "github.com/Fejiberglibstein/mem-debugger/debugger/pkg"
	"github.com/google/go-dap"
)

type CClient struct {
	debugger.Debugger
	cmd *exec.Cmd
}

func (c *CClient) Start() error {
	c.cmd = exec.Command("../installed_debuggers/codelldb/extension/adapter/codelldb")

	c.Debugger = debugger.NewDebugger(
		bufio.NewWriter(c.cmd.Stdout),
		bufio.NewReader(c.cmd.Stdin),
	)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	res, err := c.SendAndWait(&dap.InitializeRequest{
		Arguments: dap.InitializeRequestArguments{
			ClientID:                 "mem-debugger",
			ClientName:               "mem-debugger",
			AdapterID:                "mem-debugger",
			Locale:                   "en-US",
			LinesStartAt1:            false,
			ColumnsStartAt1:          false,
			SupportsMemoryReferences: true,
			PathFormat:               "path",
			SupportsVariableType:     true,
			SupportsVariablePaging:   false,

			// Ignore these for now, maybe they are useful though
			SupportsMemoryEvent:                 false,
			SupportsRunInTerminalRequest:        false,
			SupportsArgsCanBeInterpretedByShell: false,
			SupportsProgressReporting:           false,
			SupportsInvalidatedEvent:            false,
			SupportsStartDebuggingRequest:       false,
		},
	}, reflect.TypeOf(dap.InitializeResponse{}))
	if err != nil {
		return err
	}

	fmt.Print(res)

	return nil
}

func (c *CClient) Kill() {
}
