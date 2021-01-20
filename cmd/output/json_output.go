package output

import (
	"encoding/json"
	error2 "qshell/qn_shell_error"
)

/// 业务相关
type JsonOutput struct {
	pretty bool
}

func (output *JsonOutput) Output(outputType OutputType, data IOutputData, err error2.IQShellError) {
	if outputType == OutputTypeInit ||
		outputType == OutputTypeComplete {
		return
	}

	msg := ""
	format := NewPrintFormat()
	format.IsColorful = false
	switch outputType {
	case OutputTypeProgress, OutputTypeResult, OutputTypeDebug:
		msgByte, err := json.Marshal(data)
		if err != nil {

		} else {
			msg = string(msgByte)
		}
	case OutputTypeError, OutputTypeInit, OutputTypeComplete:
		if err != nil {

		} else {

		}
	}
	printBeautiful(msg, format)
	printBeautifulNewLine()
}
