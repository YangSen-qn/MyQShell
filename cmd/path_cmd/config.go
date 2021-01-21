package path_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/message"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/config"
	"qshell/qn_error"
)

type configPathCMD struct {
	*param_cmd.ParamCMD
}

func (cmd *configPathCMD) Execute(context *common.QShellContext) qn_error.IError {
	path, err := config.GetConfigPath()
	if err != nil {
		return nil
	}
	output.OutputResult(cmd, message.NewStringOutputData(path))
	return nil
}

func loadConfigPathCMD(root param_cmd.IParamCMD) {
	cmd := &configPathCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "config",
		Short: "manager config path",
		Long:  "",
	})

	root.AddCMD(cmd)
}