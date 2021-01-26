package output_utils

import (
	"qshell/cmd/output"
	"qshell/cmd/output/message"
)

func OutputProgressWithString(op output.IOutput, stringData string) {
	op.Output(output.OutputTypeProgress, message.NewStringOutputData(stringData), nil)
}

func OutputResultWithString(op output.IOutput, stringResult string) {
	op.Output(output.OutputTypeResult, message.NewStringOutputData(stringResult), nil)
}

func OutputDebugWithString(op output.IOutput, stringMessage string) {
	op.Output(output.OutputTypeDebug, message.NewStringOutputData(stringMessage), nil)
}
