package path_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/config"
	qn_shell_error "qshell/qn_error"
)

type rootPathCMD struct {
	*param_cmd.ParamCMD
}

func (cmd *rootPathCMD) Execute(context *common.QShellContext) qn_shell_error.IError {
	path, err := config.RootPath()
	if err != nil {
		return err
	}

	output_utils.OutputResultWithString(cmd, path)
	return nil
}

func loadRootPathCMD(root param_cmd.IParamCMD) {
	cmd := &rootPathCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "root",
		Short: "manager root path",
		Long:  "",
	})

	root.AddCMD(cmd)
}