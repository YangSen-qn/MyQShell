package path

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	"qshell/iqshell/config"
	"qshell/qn_shell_error"
)

type configPathCMD struct {
	*param.ParamCMD
}

func (cmd *configPathCMD) Execute(context *common.QShellContext) qn_shell_error.IQShellError {
	path, err := config.GetConfigPath()
	if err != nil {
		return nil
	}
	output.OutputResult(cmd, output.NewStringOutputData(path))
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