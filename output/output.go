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
	Output(outputType OutputType, data IOutputData, err error)
}

func OutputInit(output IOutput, err error) {
	output.Output(OutputTypeInit, nil, err)
}

func OutputProgress(output IOutput, data IOutputData) {
	output.Output(OutputTypeProgress, data, nil)
}

func OutputError(output IOutput, err error) {
	output.Output(OutputTypeError, nil, err)
}

func OutputResult(output IOutput, result IOutputData) {
	output.Output(OutputTypeResult, result, nil)
}

func OutputComplete(output IOutput, err error) {
	output.Output(OutputTypeComplete, nil, err)
}

func OutputDebug(output IOutput, message IOutputData) {
	output.Output(OutputTypeDebug, message, nil)
}

func Output(config *common.Config) IOutput {
	switch config.OutputFormatValue {
	case common.OutputFormatJSON:
		return &JsonOutput{
			pretty:true,
		}
	default:
		return &StdOutput{
			true,
		}
	}
}