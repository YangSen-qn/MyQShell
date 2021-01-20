package user

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param"
	"qshell/qn_shell_error"
	"qshell/iqshell/user"
)

type userChangeCMD struct {
	*param.ParamCMD

	name string
}

func (cmd *userChangeCMD) Check(context *common.QShellContext) qn_shell_error.IQShellError {
	if cmd.name == "" {
		return qn_shell_error.NewInvalidUserParamError("name can not empty")
	}
	return nil
}

func (cmd *userChangeCMD) Execute(context *common.QShellContext) qn_shell_error.IQShellError {
	return user.SetCurrentCredential(cmd.name)
}

func configUserChangeCMD(root param.IParamCMD) {
	cmd := &userChangeCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "change",
		Short: "change current credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
}
