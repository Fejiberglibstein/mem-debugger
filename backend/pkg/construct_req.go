package debugger

import (
	"github.com/google/go-dap"
	"log"
)

var seq int

func ConstructRequest(t any) dap.Message {
	req := dap.Request{
		ProtocolMessage: dap.ProtocolMessage{
			Seq:  seq,
			Type: "request",
		},
	}
	seq += 1
	switch m := t.(type) {
	case dap.CancelRequest:
		req.Command = "cancel"
		m.Request = req
		return &m
	case dap.RunInTerminalRequest:
		req.Command = "runInTerminal"
		m.Request = req
		return &m
	case dap.StartDebuggingRequest:
		req.Command = "startDebugging"
		m.Request = req
		return &m
	case dap.InitializeRequest:
		req.Command = "initialize"
		m.Request = req
		return &m
	case dap.ConfigurationDoneRequest:
		req.Command = "configurationDone"
		m.Request = req
		return &m
	case dap.LaunchRequest:
		req.Command = "launch"
		m.Request = req
		return &m
	case dap.AttachRequest:
		req.Command = "attach"
		m.Request = req
		return &m
	case dap.RestartRequest:
		req.Command = "restart"
		m.Request = req
		return &m
	case dap.DisconnectRequest:
		req.Command = "disconnect"
		m.Request = req
		return &m
	case dap.TerminateRequest:
		req.Command = "terminate"
		m.Request = req
		return &m
	case dap.BreakpointLocationsRequest:
		req.Command = "breakpointLocations"
		m.Request = req
		return &m
	case dap.SetBreakpointsRequest:
		req.Command = "setBreakpoints"
		m.Request = req
		return &m
	case dap.SetFunctionBreakpointsRequest:
		req.Command = "setFunctionBreakpoints"
		m.Request = req
		return &m
	case dap.SetExceptionBreakpointsRequest:
		req.Command = "setExceptionBreakpoints"
		m.Request = req
		return &m
	case dap.DataBreakpointInfoRequest:
		req.Command = "dataBreakpointInfo"
		m.Request = req
		return &m
	case dap.SetDataBreakpointsRequest:
		req.Command = "setDataBreakpoints"
		m.Request = req
		return &m
	case dap.SetInstructionBreakpointsRequest:
		req.Command = "setInstructionBreakpoints"
		m.Request = req
		return &m
	case dap.ContinueRequest:
		req.Command = "continue"
		m.Request = req
		return &m
	case dap.NextRequest:
		req.Command = "next"
		m.Request = req
		return &m
	case dap.StepInRequest:
		req.Command = "stepIn"
		m.Request = req
		return &m
	case dap.StepOutRequest:
		req.Command = "stepOut"
		m.Request = req
		return &m
	case dap.StepBackRequest:
		req.Command = "stepBack"
		m.Request = req
		return &m
	case dap.ReverseContinueRequest:
		req.Command = "reverseContinue"
		m.Request = req
		return &m
	case dap.RestartFrameRequest:
		req.Command = "restartFrame"
		m.Request = req
		return &m
	case dap.GotoRequest:
		req.Command = "goto"
		m.Request = req
		return &m
	case dap.PauseRequest:
		req.Command = "pause"
		m.Request = req
		return &m
	case dap.StackTraceRequest:
		req.Command = "stackTrace"
		m.Request = req
		return &m
	case dap.ScopesRequest:
		req.Command = "scopes"
		m.Request = req
		return &m
	case dap.VariablesRequest:
		req.Command = "variables"
		m.Request = req
		return &m
	case dap.SetVariableRequest:
		req.Command = "setVariable"
		m.Request = req
		return &m
	case dap.SourceRequest:
		req.Command = "source"
		m.Request = req
		return &m
	case dap.ThreadsRequest:
		req.Command = "threads"
		m.Request = req
		return &m
	case dap.TerminateThreadsRequest:
		req.Command = "terminateThreads"
		m.Request = req
		return &m
	case dap.ModulesRequest:
		req.Command = "modules"
		m.Request = req
		return &m
	case dap.LoadedSourcesRequest:
		req.Command = "loadedSources"
		m.Request = req
		return &m
	case dap.EvaluateRequest:
		req.Command = "evaluate"
		m.Request = req
		return &m
	case dap.SetExpressionRequest:
		req.Command = "setExpression"
		m.Request = req
		return &m
	case dap.StepInTargetsRequest:
		req.Command = "stepInTargets"
		m.Request = req
		return &m
	case dap.GotoTargetsRequest:
		req.Command = "gotoTargets"
		m.Request = req
		return &m
	case dap.CompletionsRequest:
		req.Command = "completions"
		m.Request = req
		return &m
	case dap.ExceptionInfoRequest:
		req.Command = "exceptionInfo"
		m.Request = req
		return &m
	case dap.ReadMemoryRequest:
		req.Command = "readMemory"
		m.Request = req
		return &m
	case dap.WriteMemoryRequest:
		req.Command = "writeMemory"
		m.Request = req
		return &m
	case dap.DisassembleRequest:
		req.Command = "disassemble"
		m.Request = req
		return &m
	default:
		log.Fatal("Not a valid request")
		return nil
	}
}
