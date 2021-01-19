package user

import (
	"qshell/cmd/param"
	"qshell/common"
	"qshell/execute"
	"qshell/iqshell/user"
	"qshell/output"
)

type userListCMD struct {
	*param.ParamCMD
}

func (cmd *userListCMD) Execute(context *common.QShellContext) common.IQShellError {
	userList := user.CredentialList()

	u := &user.Credential{
		IsCurrent: false,
		Name:      "kodo",
		AccessKey: "accessKey",
		SecretKey: "secretKey",
	}
	output.OutputResult(cmd, u)

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
		Short: "user list",
		Long:  "",
	})

	root.AddCMD(cmd)
}
