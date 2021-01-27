package utils

import (
	"fmt"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/util"
)

type timestampCMD struct {
	*param.ParamCMD
}

func (cmd *timestampCMD) Execute(context *common.QShellContext) error {
	timestamp := util.Timestamp()
	timestampString := fmt.Sprintf("%d", timestamp)
	output_utils.OutputResultWithString(cmd, timestampString)
	return nil
}

func loadTimestampCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &timestampCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "timestamp",
		Short: "show current timestamp by seconds",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
