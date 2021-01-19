package user

import (
	"qshell/cmd/param"
	"qshell/common"
	"qshell/execute"
	"qshell/iqshell/user"
	"qshell/output"
)

type userCMD struct {
	*param.ParamCMD
}

func (cmd *userCMD) Execute(context *common.QShellContext) common.IQShellError {
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


func configUserCMD(root param.IParamCMD) param.IParamCMD  {
	cmd := &userCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:                    "user",
		Short:                  "manager user",
		Long:                   "",
	})

	root.AddCMD(cmd)
	return cmd
}

func ConfigCMD(root param.IParamCMD) {
	userCMD := configUserCMD(root)
	configUserListCMD(userCMD)
	configUserAddCMD(userCMD)
	configUserRemoveCMD(userCMD)
	configUserChangeCMD(userCMD)
}