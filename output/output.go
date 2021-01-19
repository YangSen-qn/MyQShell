package output

import (
	"qshell/common"
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
	Output(outputType OutputType, data IOutputData, err common.IQShellError)
}

func OutputInit(output IOutput, err common.IQShellError) {
	output.Output(OutputTypeInit, nil, err)
}

func OutputProgress(output IOutput, data IOutputData) {
	output.Output(OutputTypeProgress, data, nil)
}

func OutputError(output IOutput, err common.IQShellError) {
	output.Output(OutputTypeError, nil, err)
}

func OutputResult(output IOutput, result IOutputData) {
	output.Output(OutputTypeResult, result, nil)
}

func OutputComplete(output IOutput, err common.IQShellError) {
	output.Output(OutputTypeComplete, nil, err)
}

func OutputDebug(output IOutput, message IOutputData) {
	output.Output(OutputTypeDebug, message, nil)
}

func Output(config *common.Config) IOutput {
	if config == nil {
		return &StdOutput{
			true,
		}
	}

	switch config.OutputFormatValue {
	case common.OutputFormatJSON:
		return &JsonOutput{
			pretty: true,
		}
	default:
		return &StdOutput{
			true,
		}
	}
}
