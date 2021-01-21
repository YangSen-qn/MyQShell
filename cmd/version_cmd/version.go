package version_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/message"
	"qshell/cmd/param_cmd"
	error2 "qshell/qn_error"
)

type versionCMD struct {
	*param_cmd.ParamCMD
}

func (cmd *versionCMD) Execute(context *common.QShellContext) error2.IError {
	output.OutputResult(cmd, message.NewStringOutputData(common.GetVersion()))
	return nil
}

func LoadCMD(root param_cmd.IParamCMD) {
	cmd := &versionCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "version",
		Short: "show version",
		Long:  "",
	})

	root.AddCMD(cmd)
}
