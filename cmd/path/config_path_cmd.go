package path

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	error2 "qshell/qn_shell_error"
	"qshell/iqshell/config"
)

type configPathCMD struct {
	*param.ParamCMD
}

func (cmd *configPathCMD) Execute(context *common.QShellContext) error2.IQShellError {
	output.OutputResult(cmd, output.NewStringOutputData(config.GetConfigPath()))
	return nil
}

func configConfigPathCMD(root param.IParamCMD) {
	cmd := &configPathCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "config",
		Short: "manager config path",
		Long:  "",
	})

	root.AddCMD(cmd)
}