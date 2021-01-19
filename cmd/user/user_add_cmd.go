package user

import (
	"qshell/cmd/param"
	"qshell/common"
	"qshell/execute"
)

type userAddCMD struct {
	*param.ParamCMD

	name      string
	accessKey string
	secretKey string
}

func (cmd *userAddCMD) Check(context *common.QShellContext) common.IQShellError {
	if cmd.name == "" || cmd.accessKey == "" || cmd.secretKey == "" {
		return common.NewQShellError(-1, "name or accessKey or secretKey invalid")
	}
	return nil
}

func (cmd *userAddCMD) Execute(context *common.QShellContext) common.IQShellError {

	return nil
}

func configUserAddCMD(root param.IParamCMD) {
	cmd := &userAddCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		CheckFunction: cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "add",
		Short: "manager user",
		Long:  "",
	})

	root.AddCMD(cmd)
}
