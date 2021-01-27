package path

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/iqshell/config"
)

type rootPathCMD struct {
	*param.ParamCMD
}

func (cmd *rootPathCMD) Execute(context *common.QShellContext) error {
	path, err := config.RootPath()
	if err != nil {
		return err
	}

	output_utils.OutputResultWithString(cmd, path)
	return nil
}

func loadRootPathCMD(root param.IParamCMD) {
	cmd := &rootPathCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "root",
		Short: "manager root path",
		Long:  "",
	})

	root.AddCMD(cmd)
}
