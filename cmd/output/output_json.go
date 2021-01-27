package output

import (
	"encoding/json"
	"qshell/cmd/output/message"
	"qshell/qn_error"
)

/// 业务相关
type JsonOutput struct {
	pretty bool
}

func (o *JsonOutput) Output(outputType OutputType, data message.IOutputMessage, err error) {
	if outputType == OutputTypeInit ||
		outputType == OutputTypeComplete {
		return
	}

	msg := ""
	format := NewPrintFormat()
	format.IsColorful = false
	switch outputType {
	case OutputTypeProgress, OutputTypeResult, OutputTypeDebug:
		msgByte, _ := json.Marshal(data)
		if err == nil {
			msg = string(msgByte)
		}
	case OutputTypeError, OutputTypeInit, OutputTypeComplete:
		if err != nil {
			msg = qn_error.ToJson(err)
		}
	}
	printBeautiful(msg, format)
	printBeautifulNewLine()
}
