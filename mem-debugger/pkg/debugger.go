package debugger

import (
	"bufio"
	"errors"
	"fmt"
	"reflect"

	"github.com/google/go-dap"
)

// Simple wrapper around the dap.ErrorResponse struct
type DebuggerError struct {
	dap.ErrorResponse
}

func (e *DebuggerError) Error() string {
	return e.Message
}

// Struct Wrapper to handle all the sending and receiving of dap Messages
//
// Each Client is expected to embed this struct inside it so that all of the
// Debugger's methods can be used
type Debugger struct {
	reader *bufio.Reader
	writer *bufio.Writer

	OnEvent func(dap.EventMessage)
}

func NewDebugger(writer *bufio.Writer, reader *bufio.Reader) Debugger {
	return Debugger{
		reader:  reader,
		writer:  writer,
		OnEvent: func(dap.EventMessage) { return },
	}
}

// Clients are implemented per language
type Client interface {
	// Start the debugger client by running whatever commands needed to set it
	// up, and sending the `Initialize` and `Launch` commands
	//
	// The passed in parameters will be the launch configuration for the
	// debugger, these should be passed in directly to the `launch` request
	Start(map[string]interface{}) (*dap.InitializeResponse, error)
	Kill() error

	// These are implemented by `Debugger`, so by embedding the struct into your
	// client you won't need to implement these

	SendMessage(dap.RequestMessage) error
	ReadMessage() (dap.Message, error)
	SendAndWait(dap.RequestMessage, reflect.Type) (dap.Message, error)
	WaitFor(reflect.Type) (dap.Message, error)
}

// Send a request to the debugger
//
// This function will handle the setting of the response's `seq` and `command`,
// so the request only needs the arguments filled in:
//
// ```golang
//
//	client.SendMessage(&dap.VariablesRequest {
//		// This line isn't needed:
//		// Request: dap.Request { /* ... */ }
//
//		Arguments: dap.VariablesArguments { /* ... */ }
//	})
//
// ```
func (c *Debugger) SendMessage(m dap.RequestMessage) error {
	err := dap.WriteProtocolMessage(c.writer, constructRequest(m))
	err2 := c.writer.Flush()
	if err != nil || err2 != nil {
		return fmt.Errorf("Could not send message: \n%w\n%w\n", err, err2)
	}

	return nil
}

// Read the first message that the debugger returns
func (c *Debugger) ReadMessage() (dap.Message, error) {
	res, err := dap.ReadProtocolMessage(c.reader)
	if err != nil {
		return nil, fmt.Errorf("Could not read message: %w", err)
	}
	return res, nil
}

// Send a request, and wait for a specific response.
//
// In almost every case, you will be waiting for the Response variant to the
// request you sent:
//
// ```golang
//
//	client.SendAndWait(
//		&dap.VariablesRequest{ /* ... */ },
//		reflect.TypeOf(&dap.VariablesResponse),
//	)
//
// ```
//
// See [WaitFor] for more information.
func (c *Debugger) SendAndWait(
	req dap.RequestMessage,
	desiredType reflect.Type,
) (dap.Message, error) {
	err := c.SendMessage(req)
	res, err2 := c.WaitFor(desiredType)
	return res, errors.Join(err, err2)
}

// Wait for *any* message from the debugger, accepting only the message that
// matches the type passed in.
//
// Usually, you will want to use the more powerful [SendAndWait] function,
// though this function is useful if you would like to wait for a specific event
// to happen.
//
// Note that if you are waiting for a specific event to happen, the event may
// end up getting skipped and passed through the `OnEvent` callback. This
// behavior may be changed in the future, though.
func (c *Debugger) WaitFor(desiredType reflect.Type) (dap.Message, error) {
	for {
		resp, err := c.ReadMessage()
		if err != nil {
			return nil, err
		}

		if reflect.TypeOf(resp) == desiredType {
			return resp, nil
		}

		if er, ok := resp.(*dap.ErrorResponse); ok {
			return nil, &DebuggerError{*er}
		}

		if ev, ok := resp.(dap.EventMessage); ok {
			c.OnEvent(ev)
		}
	}
}
