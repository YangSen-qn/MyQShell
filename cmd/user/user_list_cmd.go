package user

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	error2 "qshell/qn_shell_error"
	"qshell/iqshell/user"
)

type userListCMD struct {
	*param.ParamCMD
}

func (cmd *userListCMD) Execute(context *common.QShellContext) error2.IQShellError {
	userList := user.CredentialList()
	if len(userList) == 0 {
		output.OutputResult(cmd, output.NewStringOutputData("credential list is empty"))
		return nil
	}

	for _, u := range userList {
		output.OutputResult(cmd, u)
	}
	return nil
}

func configUserListCMD(root param.IParamCMD) {
	cmd := &userListCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "list",
		Short: "list credential",
		Long:  "",
	})

	root.AddCMD(cmd)
}
