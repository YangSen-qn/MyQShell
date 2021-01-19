package user

import (
	"qshell/cmd/param"
	"qshell/common"
	"qshell/execute"
	"qshell/iqshell/user"
	"qshell/output"
)

type userChangeCMD struct {
	*param.ParamCMD
}

func (cmd *userChangeCMD) Execute(context *common.QShellContext) common.IQShellError {
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

func configUserChangeCMD(root param.IParamCMD) {
	cmd := &userChangeCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "change",
		Short: "manager user",
		Long:  "",
	})

	root.AddCMD(cmd)
}
