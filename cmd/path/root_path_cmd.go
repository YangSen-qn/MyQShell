package path

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	"qshell/iqshell/config"
	qn_shell_error "qshell/qn_shell_error"
)

type rootPathCMD struct {
	*param.ParamCMD
}

func (cmd *rootPathCMD) Execute(context *common.QShellContext) qn_shell_error.IQShellError {
	path, err := config.RootPath()
	if err != nil {
		return err
	}

	output.OutputResult(cmd, output.NewStringOutputData(path))
	return nil
}

func configRootPathCMD(root param.IParamCMD) {
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