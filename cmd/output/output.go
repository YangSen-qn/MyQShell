package output

import (
	"qshell/cmd/common"
	"qshell/cmd/output/message"
	"qshell/qn_error"
)

type OutputType int8

const (
	OutputTypeInit OutputType = iota
	OutputTypeProgress
	OutputTypeError
	OutputTypeResult
	OutputTypeComplete
	OutputTypeDebug
)

type IOutput interface {
	Output(outputType OutputType, data message.IOutputMessage, err qn_error.IError)
}

func OutputInit(output IOutput, err qn_error.IError) {
	output.Output(OutputTypeInit, nil, err)
}

func OutputProgress(output IOutput, data message.IOutputMessage) {
	output.Output(OutputTypeProgress, data, nil)
}

func OutputError(output IOutput, err qn_error.IError) {
	output.Output(OutputTypeError, nil, err)
}

func OutputResult(output IOutput, result message.IOutputMessage) {
	output.Output(OutputTypeResult, result, nil)
}

func OutputComplete(output IOutput, err qn_error.IError) {
	output.Output(OutputTypeComplete, nil, err)
}

func OutputDebug(output IOutput, message message.IOutputMessage) {
	output.Output(OutputTypeDebug, message, nil)
}

func Output(config *common.Config) IOutput {
	if config == nil {
		return &DefaultOutput{
			true,
		}
	}

	switch config.OutputFormatValue {
	case common.OutputFormatJSON:
		return &JsonOutput{
			pretty: true,
		}
	default:
		return &DefaultOutput{
			true,
		}
	}
}
