package path

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/iqshell/config"
)

type configPathCMD struct {
	*param.ParamCMD
}

func (cmd *configPathCMD) Execute(context *common.QShellContext) error {
	path, err := config.GetConfigPath()
	if err != nil {
		return nil
	}
	output_utils.OutputResultWithString(cmd, path)
	return nil
}

func loadConfigPathCMD(root param.IParamCMD) {
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
