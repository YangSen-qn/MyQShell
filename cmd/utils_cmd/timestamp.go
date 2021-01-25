package utils_cmd

import (
	"fmt"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/message"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
)

type timestampCMD struct {
	*param_cmd.ParamCMD
}

func (cmd *timestampCMD) Execute(context *common.QShellContext) qn_error.IError {
	timestamp := qn_util.Timestamp()
	timestampString := fmt.Sprintf("%d", timestamp)
	output.OutputResult(cmd, message.NewStringOutputData(timestampString))
	return nil
}

func loadTimestampCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &timestampCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "timestamp",
		Short: "show current timestamp by seconds",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
