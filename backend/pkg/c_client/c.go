package c_client

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"

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
		"--port",
		PORT,
	)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	var conn net.Conn
	var err error
	tries := 0
	for {
		time.Sleep(500 * time.Millisecond)
		conn, err = net.Dial("tcp", net.JoinHostPort("localhost", PORT))

		if err == nil {
			break
		}

		log.Print(err)
		tries += 1
		if tries > 10 {
			return errors.New("Could not connect to the debugger for codelldb on port " + PORT)
		}
	}

	c.Debugger = debugger.NewDebugger(
		bufio.NewWriter(conn),
		bufio.NewReader(conn),
	)

	err = c.SendMessage(
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
		})

	res, err := c.ReadMessage()

	if err != nil {
		return err
	}

	fmt.Print(res)

	// Override the default value, it should always be true
	launchArgs["stopOnEntry"] = true

	launch, err := json.Marshal(launchArgs)
	if err != nil {
		return nil
	}

	err = c.SendMessage(&dap.LaunchRequest{Arguments: launch})
	res, err = c.ReadMessage()
	fmt.Print(res)

	if err != nil {
		return err
	}

	fmt.Print(res)

	return nil
}

func (c *CClient) Kill() {
}
