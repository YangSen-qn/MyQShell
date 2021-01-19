package version

import (
	"qshell/cmd/param"
	"qshell/common"
	"qshell/execute"
	"qshell/output"
)

type versionCMD struct {
	*param.ParamCMD
}

func (cmd *versionCMD) Execute(context *common.QShellContext) common.IQShellError {
	output.OutputResult(cmd, output.NewStringOutputData(common.GetVersion()))
	return nil
}

func ConfigCMD(root param.IParamCMD) {
	versionCmd := &versionCMD{
		ParamCMD: param.NewParamCMD(),
	}

	versionCmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: versionCmd.Execute,
	})

	versionCmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "version",
		Short: "show version",
		Long:  "",
	})

	root.AddCMD(versionCmd)
}
