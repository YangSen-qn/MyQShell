package user

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param"
	"qshell/qn_shell_error"
	"qshell/iqshell/user"
)

type userAddCMD struct {
	*param.ParamCMD

	name      string
	accessKey string
	secretKey string
}

func (cmd *userAddCMD) Check(context *common.QShellContext) qn_shell_error.IQShellError {
	if cmd.name == "" || cmd.accessKey == "" || cmd.secretKey == "" {
		return qn_shell_error.NewInvalidUserParamError("name or accessKey or secretKey invalid")
	}
	return nil
}

func (cmd *userAddCMD) Execute(context *common.QShellContext) qn_shell_error.IQShellError {
	credential := &user.Credential{
		Name:      cmd.name,
		AccessKey: cmd.accessKey,
		SecretKey: cmd.secretKey,
	}
	return user.AddCredential(credential)
}

func configUserAddCMD(root param.IParamCMD) {
	cmd := &userAddCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "add",
		Short: "add credential",
		Long:  "",
		Example: "qshell credential add name \"credential_name\" accessKey \"ak\" secretKey \"sk\"",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")
	cmd.FlagsStringVar(&cmd.accessKey, "accessKey", "", "", "access key")
	cmd.FlagsStringVar(&cmd.secretKey, "secretKey", "", "", "secret key")

	root.AddCMD(cmd)
}
