package debugger

import (
	"github.com/google/go-dap"
	"log"
)

var seq int

func constructRequest(t dap.RequestMessage) dap.Message {
	req := t.GetRequest()

	req.Seq = seq
	req.Type = "request"
	seq += 1

	switch t.(type) {
	case *dap.CancelRequest:
		req.Command = "cancel"
	case *dap.RunInTerminalRequest:
		req.Command = "runInTerminal"
	case *dap.StartDebuggingRequest:
		req.Command = "startDebugging"
	case *dap.InitializeRequest:
		req.Command = "initialize"
	case *dap.ConfigurationDoneRequest:
		req.Command = "configurationDone"
	case *dap.LaunchRequest:
		req.Command = "launch"
	case *dap.AttachRequest:
		req.Command = "attach"
	case *dap.RestartRequest:
		req.Command = "restart"
	case *dap.DisconnectRequest:
		req.Command = "disconnect"
	case *dap.TerminateRequest:
		req.Command = "terminate"
	case *dap.BreakpointLocationsRequest:
		req.Command = "breakpointLocations"
	case *dap.SetBreakpointsRequest:
		req.Command = "setBreakpoints"
	case *dap.SetFunctionBreakpointsRequest:
		req.Command = "setFunctionBreakpoints"
	case *dap.SetExceptionBreakpointsRequest:
		req.Command = "setExceptionBreakpoints"
	case *dap.DataBreakpointInfoRequest:
		req.Command = "dataBreakpointInfo"
	case *dap.SetDataBreakpointsRequest:
		req.Command = "setDataBreakpoints"
	case *dap.SetInstructionBreakpointsRequest:
		req.Command = "setInstructionBreakpoints"
	case *dap.ContinueRequest:
		req.Command = "continue"
	case *dap.NextRequest:
		req.Command = "next"
	case *dap.StepInRequest:
		req.Command = "stepIn"
	case *dap.StepOutRequest:
		req.Command = "stepOut"
	case *dap.StepBackRequest:
		req.Command = "stepBack"
	case *dap.ReverseContinueRequest:
		req.Command = "reverseContinue"
	case *dap.RestartFrameRequest:
		req.Command = "restartFrame"
	case *dap.GotoRequest:
		req.Command = "goto"
	case *dap.PauseRequest:
		req.Command = "pause"
	case *dap.StackTraceRequest:
		req.Command = "stackTrace"
	case *dap.ScopesRequest:
		req.Command = "scopes"
	case *dap.VariablesRequest:
		req.Command = "variables"
	case *dap.SetVariableRequest:
		req.Command = "setVariable"
	case *dap.SourceRequest:
		req.Command = "source"
	case *dap.ThreadsRequest:
		req.Command = "threads"
	case *dap.TerminateThreadsRequest:
		req.Command = "terminateThreads"
	case *dap.ModulesRequest:
		req.Command = "modules"
	case *dap.LoadedSourcesRequest:
		req.Command = "loadedSources"
	case *dap.EvaluateRequest:
		req.Command = "evaluate"
	case *dap.SetExpressionRequest:
		req.Command = "setExpression"
	case *dap.StepInTargetsRequest:
		req.Command = "stepInTargets"
	case *dap.GotoTargetsRequest:
		req.Command = "gotoTargets"
	case *dap.CompletionsRequest:
		req.Command = "completions"
	case *dap.ExceptionInfoRequest:
		req.Command = "exceptionInfo"
	case *dap.ReadMemoryRequest:
		req.Command = "readMemory"
	case *dap.WriteMemoryRequest:
		req.Command = "writeMemory"
	case *dap.DisassembleRequest:
		req.Command = "disassemble"
	default:
		log.Fatal("Not a valid request")
	}
	return t
}
