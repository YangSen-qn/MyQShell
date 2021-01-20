package version

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	error2 "qshell/qn_shell_error"
)

type versionCMD struct {
	*param.ParamCMD
}

func (cmd *versionCMD) Execute(context *common.QShellContext) error2.IQShellError {
	output.OutputResult(cmd, output.NewStringOutputData(common.GetVersion()))
	return nil
}

func ConfigCMD(root param.IParamCMD) {
	cmd := &versionCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "version",
		Short: "show version",
		Long:  "",
	})

	root.AddCMD(cmd)
}
