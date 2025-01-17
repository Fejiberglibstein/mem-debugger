package c_client

import (
	"bufio"
	"encoding/json"
	"os/exec"
	"reflect"

	debugger "github.com/Fejiberglibstein/mem-debugger/debugger/pkg"
	"github.com/google/go-dap"
)

const PORT = "8921"

type CClient struct {
	debugger.Debugger
	cmd *exec.Cmd
}

func (c *CClient) Start(launchArgs map[string]interface{}) error {
	c.cmd = exec.Command(
		"../installed_debuggers/codelldb/extension/adapter/codelldb",
	)

	stdout, _ := c.cmd.StdoutPipe()
	stdin, _ := c.cmd.StdinPipe()

	// stderr, _ := c.cmd.StderrPipe()
	// c.cmd.Env = []string{"RUST_LOG=trace"}
	// go func() {
	// 	io.Copy(os.Stdout, stderr)
	// }()

	c.Debugger = debugger.NewDebugger(
		bufio.NewWriter(stdin),
		bufio.NewReader(stdout),
	)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	_, err := c.SendAndWait(
		&dap.InitializeRequest{
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
		},
		reflect.TypeOf(&dap.InitializeResponse{}),
	)

	if err != nil {
		return err
	}

	// Override the default value, it should always be true
	launchArgs["stopOnEntry"] = true

	launch, err := json.Marshal(launchArgs)
	if err != nil {
		return nil
	}

	_, err = c.SendAndWait(
		&dap.LaunchRequest{Arguments: launch},
		reflect.TypeOf(&dap.InitializedEvent{}),
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *CClient) Kill() {
}
