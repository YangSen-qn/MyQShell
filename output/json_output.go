package output

import (
	"encoding/json"
	"qshell/common"
)

/// 业务相关
type JsonOutput struct {
	pretty bool
}

func (output *JsonOutput) Output(outputType OutputType, data IOutputData, err common.IQShellError) {
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
