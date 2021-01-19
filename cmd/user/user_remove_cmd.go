package user

import (
	"qshell/cmd/param"
	"qshell/common"
	"qshell/execute"
	"qshell/iqshell/user"
	"qshell/output"
)

type userRemoveCMD struct {
	*param.ParamCMD
}

func (cmd *userRemoveCMD) Execute(context *common.QShellContext) common.IQShellError {
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

func configUserRemoveCMD(root param.IParamCMD) {
	cmd := &userRemoveCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "remove",
		Short: "manager user",
		Long:  "",
	})

	root.AddCMD(cmd)
}
