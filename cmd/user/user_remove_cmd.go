package user

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param"
	"qshell/qn_shell_error"
	"qshell/iqshell/user"
)

type userRemoveCMD struct {
	*param.ParamCMD

	name string
}

func (cmd *userRemoveCMD) Check(context *common.QShellContext) qn_shell_error.IQShellError {
	if cmd.name == "" {
		return qn_shell_error.NewInvalidUserParamError("name can not empty")
	}
	return nil
}

func (cmd *userRemoveCMD) Execute(context *common.QShellContext) qn_shell_error.IQShellError {
	return user.RemoveCredential(cmd.name)
}

func configUserRemoveCMD(root param.IParamCMD) {
	cmd := &userRemoveCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "remove",
		Short: "remove credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
}
